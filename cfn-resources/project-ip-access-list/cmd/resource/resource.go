package resource

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/rs/xid"
	"go.mongodb.org/atlas/mongodbatlas"
)

func setup() {
	util.SetupLogger("mongodb-atlas-project-ip-access-list")
}

var CreateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.AccessList}
var ReadRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.AccessList}
var UpdateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.AccessList}
var DeleteRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.AccessList}
var ListRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	_, _ = logger.Debugf("currentModel: %+v, prevModel: %+v", currentModel, prevModel)

	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	_, _ = logger.Debugf("currentModel: %+v, prevModel: %+v", currentModel, prevModel)

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %s", err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
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

	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %s", err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	projectID := *currentModel.ProjectId

	_, _ = logger.Debugf("Read --- currentModel:%+v", currentModel)

	var entries []string
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
	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %s", err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
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
	if errEvent := validateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %s", err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
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

	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %s", err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	var pageNum, itemsPerPage int
	var includeCount bool

	if currentModel.ListOptions != nil {
		if currentModel.ListOptions.PageNum != nil {
			pageNum = *currentModel.ListOptions.PageNum
		}
		if currentModel.ListOptions.IncludeCount != nil {
			includeCount = *currentModel.ListOptions.IncludeCount
		}
		if currentModel.ListOptions.ItemsPerPage != nil {
			itemsPerPage = *currentModel.ListOptions.ItemsPerPage
		}
	}

	listOptions := &mongodbatlas.ListOptions{
		PageNum:      pageNum,
		IncludeCount: includeCount,
		ItemsPerPage: itemsPerPage,
	}

	result, resp, err := client.ProjectIPAccessList.List(context.Background(), *currentModel.ProjectId, listOptions)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	mm := make([]AccessListDefinition, 0)
	for i := range result.Results {
		var m AccessListDefinition
		m.completeByConnection(result.Results[i])
		mm = append(mm, m)
	}
	currentModel.AccessList = mm
	// create list with 1
	models := []interface{}{}
	models = append(models, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}

func (m *AccessListDefinition) completeByConnection(c mongodbatlas.ProjectIPAccessList) {
	m.IPAddress = &c.IPAddress
	m.CIDRBlock = &c.CIDRBlock
	m.Comment = &c.Comment
	m.AwsSecurityGroup = &c.AwsSecurityGroup
	m.ProjectId = &c.GroupID
}

func getProjectIPAccessList(projectID string, entries []string, conn *mongodbatlas.Client) ([]*mongodbatlas.ProjectIPAccessList, handler.ProgressEvent, error) {
	var accesslist []*mongodbatlas.ProjectIPAccessList
	for i := range entries {
		entry := entries[i]
		result, resp, err := conn.ProjectIPAccessList.Get(context.Background(), projectID, entry)
		if err != nil {
			return nil, progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
				resp.Response), err
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
		if w.DeleteAfterDate != nil {
			wl.DeleteAfterDate = *w.DeleteAfterDate
		}
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

		_, _ = logger.Debugf(" getProjectIPAccessListRequest: %+v\n", wl)

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
		resp, err := client.ProjectIPAccessList.Delete(context.Background(), projectID, entry)
		if err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
				resp.Response), err
		}
	}

	return handler.ProgressEvent{}, nil
}
