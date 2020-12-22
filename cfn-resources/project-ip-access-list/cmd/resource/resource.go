package resource

import (
	//"errors"
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/rs/xid"
	"go.mongodb.org/atlas/mongodbatlas"
	"log"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	log.Printf("%#+v\n", currentModel)

	err = createEntries(currentModel, client)
	log.Printf("Create err:%v", err)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	guid := xid.New()

	x := guid.String()
	currentModel.Id = &x

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId

	entries := []string{}
	for _, wl := range currentModel.AccessList {
		entry := getEntry(wl)
		entries = append(entries, entry)
	}

	accesslist, err := getProjectIPAccessList(projectID, entries, client)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	currentModel.AccessList = flattenAccessList(accesslist)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil

}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
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
	log.Printf("Got list request - returning read - %v", currentModel)
	return Read(req, prevModel, currentModel)
	//return handler.ProgressEvent{
	//	OperationStatus: handler.Success,
	//	Message:         "List Complete",
	//	ResourceModel:   currentModel,
	//}, nil

}

func getProjectIPAccessList(projectID string, entries []string, conn *mongodbatlas.Client) ([]*mongodbatlas.ProjectIPAccessList, error) {

	var accesslist []*mongodbatlas.ProjectIPAccessList
	for _, entry := range entries {
		res, _, err := conn.ProjectIPAccessList.Get(context.Background(), projectID, entry)
		if err != nil {
			return nil, fmt.Errorf("error getting project IP accesslist information: %s", err)
		}
		accesslist = append(accesslist, res)
	}
	return accesslist, nil
}

func getProjectIPAccessListRequest(model *Model) []*mongodbatlas.ProjectIPAccessList {
	var accesslist []*mongodbatlas.ProjectIPAccessList
	for _, w := range model.AccessList {
		wl := &mongodbatlas.ProjectIPAccessList{}
		if w.Comment != nil {
			wl.Comment = *w.Comment
		}
		if w.CIDRBlock != nil {
			wl.CIDRBlock = *w.CIDRBlock
		}
		if w.IPAddress != nil {
			wl.IPAddress = *w.IPAddress
		}
		if w.AwsSecurityGroup != nil {
			wl.AwsSecurityGroup = *w.AwsSecurityGroup
		}

		log.Printf("getProjectIPAccessListRequest: %+#v\n", wl)

		accesslist = append(accesslist, wl)
	}
	log.Printf("getProjectIPAccessListRequest accesslist:%v", accesslist)
	return accesslist
}

func getEntry(wl AccessListDefinition) string {
	if wl.IPAddress != nil {
		return *wl.IPAddress
	}
	if wl.CIDRBlock != nil {
		return *wl.CIDRBlock
	}
	if wl.AwsSecurityGroup != nil {
		return *wl.AwsSecurityGroup
	}
	return ""
}

func flattenAccessList(accesslist []*mongodbatlas.ProjectIPAccessList) []AccessListDefinition {
	var results []AccessListDefinition
	for _, wl := range accesslist {
		r := AccessListDefinition{
			IPAddress:        &wl.IPAddress,
			CIDRBlock:        &wl.CIDRBlock,
			AwsSecurityGroup: &wl.AwsSecurityGroup,
			Comment:          &wl.Comment,
			ProjectId:        &wl.GroupID,
		}
		results = append(results, r)
	}
	return results
}

func createEntries(model *Model, client *mongodbatlas.Client) error {
	request := getProjectIPAccessListRequest(model)
	projectID := *model.ProjectId
	log.Printf("createEntries : projectID:%s, model:%v", projectID, model)
	_, _, err := client.ProjectIPAccessList.Create(context.Background(), projectID, request)
	return err
}

func deleteEntries(model *Model, client *mongodbatlas.Client) error {
	projectID := *model.ProjectId
	var err error

	for _, wl := range model.AccessList {
		entry := getEntry(wl)
		_, errDelete := client.ProjectIPAccessList.Delete(context.Background(), projectID, entry)
		if errDelete != nil {
			err = fmt.Errorf("error deleting(%s) %w ", entry, errDelete)
		}
	}

	return err
}
