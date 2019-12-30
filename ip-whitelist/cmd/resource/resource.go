package resource

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId.Value()

	request := getProjectIPWhitelistRequest(currentModel)

	_, _, err = client.ProjectIPWhitelist.Create(context.Background(), projectID, request)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	var withelist []string
	whiteListMap(request, func(entry string) {
		withelist = append(withelist, entry)
	})

	id := encodeStateID(map[string]string{
		"project_id": projectID,
		"entries":    strings.Join(withelist, ","),
	})

	currentModel.Id = encoding.NewString(id)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

func getProjectIPWhitelistRequest(model *Model) []*mongodbatlas.ProjectIPWhitelist {
	var whitelist []*mongodbatlas.ProjectIPWhitelist
	for _, w := range model.Whitelist {
		wl := &mongodbatlas.ProjectIPWhitelist{}
		if w.Comment != nil {
			wl.Comment = *w.Comment.Value()
		}
		if w.CidrBlock != nil {
			wl.CIDRBlock = *w.CidrBlock.Value()
		}
		if w.IpAddress != nil {
			wl.IPAddress = *w.IpAddress.Value()
		}

		whitelist = append(whitelist, wl)

	}
	return whitelist
}

func whiteListMap(whitelist []*mongodbatlas.ProjectIPWhitelist, f func(string)) {
	for _, entry := range whitelist {
		if entry.CIDRBlock != "" {
			f(entry.CIDRBlock)
		} else if entry.IPAddress != "" {
			f(entry.IPAddress)
		}
	}
}

func encodeStateID(values map[string]string) string {
	encode := func(e string) string { return base64.StdEncoding.EncodeToString([]byte(e)) }
	encodedValues := make([]string, 0)

	for key, value := range values {
		encodedValues = append(encodedValues, fmt.Sprintf("%s:%s", encode(key), encode(value)))
	}
	return strings.Join(encodedValues, "-")
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	ids := decodeStateID(*currentModel.Id.Value())

	whitelist, err := getProjectIPWhitelist(ids, client)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	currentModel.Whitelist = flattenWhitelist(whitelist)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

func flattenWhitelist(whitelist []*mongodbatlas.ProjectIPWhitelist) []WhitelistDefinition {
	var results []WhitelistDefinition
	for _, wl := range whitelist {
		r := WhitelistDefinition{
			IpAddress: encoding.NewString(wl.IPAddress),
			CidrBlock: encoding.NewString(wl.CIDRBlock),
			Comment:   encoding.NewString(wl.Comment),
			ProjectId: encoding.NewString(wl.GroupID),
		}
		results = append(results, r)
	}

	return results
}

func decodeStateID(stateID string) map[string]string {
	decode := func(d string) string {
		decodedString, err := base64.StdEncoding.DecodeString(d)
		if err != nil {
			log.Printf("[WARN] error decoding state ID: %s", err)
		}
		return string(decodedString)
	}
	decodedValues := make(map[string]string)
	encodedValues := strings.Split(stateID, "-")

	for _, value := range encodedValues {
		keyValue := strings.Split(value, ":")
		decodedValues[decode(keyValue[0])] = decode(keyValue[1])
	}
	return decodedValues
}

func getProjectIPWhitelist(ids map[string]string, conn *mongodbatlas.Client) ([]*mongodbatlas.ProjectIPWhitelist, error) {
	projectID := ids["project_id"]
	entries := strings.Split(ids["entries"], ",")

	var whitelist []*mongodbatlas.ProjectIPWhitelist
	for _, entry := range entries {
		res, _, err := conn.ProjectIPWhitelist.Get(context.Background(), projectID, entry)
		if err != nil {
			return nil, fmt.Errorf("error getting project IP whitelist information: %s", err)
		}
		whitelist = append(whitelist, res)
	}
	return whitelist, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	ids := decodeStateID(*currentModel.Id.Value())

	whitelist, err := getProjectIPWhitelist(ids, client)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	whiteListMap(whitelist, func(entry string) {
		_, err = client.ProjectIPWhitelist.Delete(context.Background(), ids["project_id"], entry)
	})
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting project IP whitelist: %s", err)
	}

	request := getProjectIPWhitelistRequest(currentModel)
	projectID := *currentModel.ProjectId.Value()

	_, _, err = client.ProjectIPWhitelist.Create(context.Background(), projectID, request)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	var withelist []string
	whiteListMap(request, func(entry string) {
		withelist = append(withelist, entry)
	})

	id := encodeStateID(map[string]string{
		"project_id": projectID,
		"entries":    strings.Join(withelist, ","),
	})

	currentModel.Id = encoding.NewString(id)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	ids := decodeStateID(*currentModel.Id.Value())

	whitelist, err := getProjectIPWhitelist(ids, client)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	whiteListMap(whitelist, func(entry string) {
		_, err = client.ProjectIPWhitelist.Delete(context.Background(), ids["project_id"], entry)
	})
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting project IP whitelist: %s", err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
// NO-OP
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   currentModel,
	}, nil
}
