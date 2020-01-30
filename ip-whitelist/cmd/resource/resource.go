package resource

import (
	"context"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/rs/xid"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	err = createEntries(currentModel, client)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	guid := xid.New()

	currentModel.Id = encoding.NewString(guid.String())

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId.Value()

	entries := []string{}
	for _, wl := range currentModel.Whitelist {
		entry := getEntry(wl)
		entries = append(entries, entry)
	}

	whitelist, err := getProjectIPWhitelist(projectID, entries, client)
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

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	err = deleteEntries(currentModel, client)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Update Failed",
		}, err
	}

	err = createEntries(currentModel, client)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

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

	err = deleteEntries(currentModel, client)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Delete Failed",
		}, err
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

func getProjectIPWhitelist(projectID string, entries []string, conn *mongodbatlas.Client) ([]*mongodbatlas.ProjectIPWhitelist, error) {

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
		if w.AwsSecurityGroup != nil {
			wl.AwsSecurityGroup = *w.AwsSecurityGroup.Value()
		}

		whitelist = append(whitelist, wl)
	}
	return whitelist
}

func getEntry(wl WhitelistDefinition) string {
	if wl.IpAddress != nil {
		return *wl.IpAddress.Value()
	}
	if wl.CidrBlock != nil {
		return *wl.CidrBlock.Value()
	}
	if wl.AwsSecurityGroup != nil {
		return *wl.AwsSecurityGroup.Value()
	}
	return ""
}

func flattenWhitelist(whitelist []*mongodbatlas.ProjectIPWhitelist) []WhitelistDefinition {
	var results []WhitelistDefinition
	for _, wl := range whitelist {
		r := WhitelistDefinition{
			IpAddress:        encoding.NewString(wl.IPAddress),
			CidrBlock:        encoding.NewString(wl.CIDRBlock),
			AwsSecurityGroup: encoding.NewString(wl.AwsSecurityGroup),
			Comment:          encoding.NewString(wl.Comment),
			ProjectId:        encoding.NewString(wl.GroupID),
		}
		results = append(results, r)
	}
	return results
}

func createEntries(model *Model, client *mongodbatlas.Client) error {
	request := getProjectIPWhitelistRequest(model)
	projectID := *model.ProjectId.Value()

	_, _, err := client.ProjectIPWhitelist.Create(context.Background(), projectID, request)
	return err
}

func deleteEntries(model *Model, client *mongodbatlas.Client) error {
	projectID := *model.ProjectId.Value()
	var err error

	for _, wl := range model.Whitelist {
		entry := getEntry(wl)
		_, errDelete := client.ProjectIPWhitelist.Delete(context.Background(), projectID, entry)
		if errDelete != nil {
			err = fmt.Errorf("error deleting(%s) %w ", entry, errDelete)
		}
	}

	return err
}
