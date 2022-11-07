package resource

import (
	"context"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/rs/xid"
	"go.mongodb.org/atlas/mongodbatlas"
)

func setup() {
	util.SetupLogger("mongodb-atlas-project-ip-access-list")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	_, _ = logger.Debugf("currentModel: %+v, prevModel: %+v", currentModel, prevModel)

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	event, err := createEntries(currentModel, client)
	if err != nil {
		_, _ = logger.Warnf("Create err:%v", err)
		return event, nil
	}

	guid := xid.New()

	x := guid.String()
	currentModel.Id = &x
	_, _ = logger.Debugf("Create --- currentModel:%+v", currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	projectID := *currentModel.ProjectId

	_, _ = logger.Debugf("Read --- currentModel:%+v", currentModel)

	entries := []string{}
	for i := range currentModel.AccessList {
		wl := currentModel.AccessList[i]
		entry := getEntry(wl)
		entries = append(entries, entry)
	}

	_, _ = logger.Debugf("Read --- entries:%+v", entries)
	accesslist, progressEvent, err := getProjectIPAccessList(projectID, entries, client)
	_, _ = logger.Debugf("Read --- accesslist:%+v, progressEvent:%+v", accesslist, progressEvent)
	if err != nil {
		_, _ = logger.Warnf("error READ access list projectID:%s, error: %s, progressEvent: %+v", projectID, err, progressEvent)
		return progressEvent, nil
	}

	currentModel.AccessList = flattenAccessList(currentModel.AccessList, accesslist)
	_, _ = logger.Debugf("Read --- currentModel.AccessList:%+v", currentModel.AccessList)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	progressEvent, err := deleteEntries(currentModel, client)
	if err != nil {
		_, _ = logger.Warnf("Update deleteEntries error:%+v", err)
		return progressEvent, nil
	}

	progressEvent, err = createEntries(currentModel, client)
	if err != nil {
		_, _ = logger.Warnf("Update createEntries error:%+v", err)
		return progressEvent, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	event, err := deleteEntries(currentModel, client)
	if err != nil {
		_, _ = logger.Warnf("Delete deleteEntries error:%+v", err)
		return event, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

// List handles the List event from the Cloudformation service.
// NO-OP
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	_, _ = logger.Debugf("Got list request - returning read - %v", currentModel)
	readEvent, err := Read(req, prevModel, currentModel)
	_, _ = logger.Debugf("List readEvent:+%v   --------------------------- error:%+v", readEvent, err)
	if readEvent.OperationStatus == handler.Failed {
		return readEvent, nil
	}
	// create list with 1
	models := []interface{}{}
	models = append(models, readEvent.ResourceModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}

func getProjectIPAccessList(projectID string, entries []string, conn *mongodbatlas.Client) (accesslist []*mongodbatlas.ProjectIPAccessList,
	progressEvent handler.ProgressEvent, err error) {
	for i := range entries {
		entry := entries[i]
		result, resp, err := conn.ProjectIPAccessList.Get(context.Background(), projectID, entry)
		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				_, _ = logger.Warnf("Resource Not Found 404 for READ projectId:%s, entry:%+v, err:%+v", projectID, entry, err)
				return nil, handler.ProgressEvent{
					Message:          err.Error(),
					OperationStatus:  handler.Failed,
					HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, err
			}
			_, _ = logger.Warnf("Error READ projectId:%s, err:%+v", projectID, err)
			return nil, handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, err
		}
		_, _ = logger.Debugf("%+v", strings.Split(result.CIDRBlock, "/"))
		_, _ = logger.Debugf("getProjectIPAccessList result:%+v", result)
		accesslist = append(accesslist, result)
	}
	return accesslist, handler.ProgressEvent{}, nil
}

func getProjectIPAccessListRequest(model *Model) []*mongodbatlas.ProjectIPAccessList {
	var accesslist []*mongodbatlas.ProjectIPAccessList
	for i := range model.AccessList {
		w := model.AccessList[i]
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

		_, _ = logger.Debugf("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ getProjectIPAccessListRequest: %+v\n", wl)

		accesslist = append(accesslist, wl)
	}
	_, _ = logger.Debugf("getProjectIPAccessListRequest accesslist:%v", accesslist)
	return accesslist
}

func getEntry(wl AccessListDefinition) string {
	if wl.CIDRBlock != nil {
		return *wl.CIDRBlock
	}
	if wl.AwsSecurityGroup != nil {
		return *wl.AwsSecurityGroup
	}
	if wl.IPAddress != nil {
		return *wl.IPAddress
	}
	return ""
}

func flattenAccessList(original []AccessListDefinition, accesslist []*mongodbatlas.ProjectIPAccessList) []AccessListDefinition {
	var results []AccessListDefinition
	for i := range accesslist {
		wl := accesslist[i]
		// only add properties which were in model to begin with
		r := AccessListDefinition{
			IPAddress: &wl.IPAddress,
			Comment:   &wl.Comment,
		}
		if original[i].CIDRBlock != nil {
			r.CIDRBlock = &wl.CIDRBlock
		}
		if original[i].ProjectId != nil {
			r.ProjectId = &wl.GroupID
		}
		if original[i].AwsSecurityGroup != nil {
			r.AwsSecurityGroup = &wl.AwsSecurityGroup
		}
		results = append(results, r)
	}
	return results
}

func createEntries(model *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {
	request := getProjectIPAccessListRequest(model)
	projectID := *model.ProjectId
	_, _ = logger.Debugf("createEntries : projectID:%s, model:%+v, request:%+v", projectID, model, request)
	result, _, err := client.ProjectIPAccessList.Create(context.Background(), projectID, request)
	if err != nil {
		_, _ = logger.Warnf("Error createEntries projectId:%s,err:%+v", projectID, err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, err
	}
	_, _ = logger.Debugf("createEntries result:%+v", result)
	return handler.ProgressEvent{}, nil
}

func deleteEntries(model *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {
	projectID := *model.ProjectId

	for i := range model.AccessList {
		wl := model.AccessList[i]
		entry := getEntry(wl)
		resp, errDelete := client.ProjectIPAccessList.Delete(context.Background(), projectID, entry)
		if errDelete != nil {
			if resp != nil && resp.StatusCode == 404 {
				_, _ = logger.Warnf("Resource Not Found 404 deleteEntries projectId:%s, entry:%+v, err:%+v", projectID, entry, errDelete)
				return handler.ProgressEvent{
					Message:          errDelete.Error(),
					OperationStatus:  handler.Failed,
					HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, errDelete
			}
			_, _ = logger.Warnf("Error READ projectId:%s, err:%+v", projectID, errDelete)
			return handler.ProgressEvent{
				Message:          errDelete.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, errDelete
		}
	}
	return handler.ProgressEvent{}, nil
}
