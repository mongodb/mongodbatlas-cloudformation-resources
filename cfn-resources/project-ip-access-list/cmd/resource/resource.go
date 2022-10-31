package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progress_event"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/atlas/mongodbatlas"
	"strings"
)

func setup() {
	util.SetupLogger("mongodb-atlas-project-ip-access-list")
}

var CreateRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId", "AccessList"}
var ReadRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId", "AccessList"}
var UpdateRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId", "AccessList"}
var DeleteRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId", "AccessList"}
var ListRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId"}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Print("Create handler called")

	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	log.Debugf("currentModel: %+v, prevModel: %+v", currentModel, prevModel)

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	event, err := createEntries(currentModel, client)
	if err != nil {
		log.Debugf("Create err:%v", err)
		return event, nil
	}

	guid := xid.New()

	x := guid.String()
	currentModel.Id = &x
	log.Debugf("Create --- currentModel:%+v", currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	projectID := *currentModel.ProjectId

	log.Debugf("Read --- currentModel:%+v", currentModel)

	entries := []string{}
	for i, _ := range currentModel.AccessList {
		wl := currentModel.AccessList[i]
		entry := getEntry(wl)
		entries = append(entries, entry)
	}

	log.Debugf("Read --- entries:%+v", entries)
	accesslist, progressEvent, err := getProjectIPAccessList(projectID, entries, client)
	log.Debugf("Read --- accesslist:%+v, progressEvent:%+v", accesslist, progressEvent)
	if err != nil {
		log.Debugf("error READ access list projectID:%s, error: %s, progressEvent: %+v", projectID, err, progressEvent)
		return progressEvent, nil
	}

	currentModel.AccessList = flattenAccessList(currentModel.AccessList, accesslist)
	log.Debugf("Read --- currentModel.AccessList:%+v", currentModel.AccessList)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil

}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	progressEvent, err := deleteEntries(currentModel, client)
	if err != nil {
		log.Debugf("Update deleteEntries error:%+v", err)
		return progressEvent, nil
	}

	progressEvent, err = createEntries(currentModel, client)
	if err != nil {
		log.Debugf("Update createEntries error:%+v", err)
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
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	event, err := deleteEntries(currentModel, client)
	if err != nil {
		log.Debugf("Delete deleteEntries error:%+v", err)
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
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
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
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	mm := make([]AccessListDefinition, 0)
	for _, projectIpAccessList := range result.Results {
		var m AccessListDefinition
		m.completeByConnection(projectIpAccessList)
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
	for i, _ := range entries {
		entry := entries[i]
		result, resp, err := conn.ProjectIPAccessList.Get(context.Background(), projectID, entry)
		if err != nil {
			return nil, progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
				resp.Response), err
		}
		log.Debugf("%+v", strings.Split(result.CIDRBlock, "/"))
		log.Debugf("getProjectIPAccessList result:%+v", result)
		accesslist = append(accesslist, result)
	}
	return accesslist, handler.ProgressEvent{}, nil
}

func getProjectIPAccessListRequest(model *Model) []*mongodbatlas.ProjectIPAccessList {
	var accesslist []*mongodbatlas.ProjectIPAccessList
	for i, _ := range model.AccessList {
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

		log.Debugf("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ getProjectIPAccessListRequest: %+v\n", wl)

		accesslist = append(accesslist, wl)
	}
	log.Debugf("getProjectIPAccessListRequest accesslist:%v", accesslist)
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
	for i, _ := range accesslist {
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
	log.Debugf("createEntries : projectID:%s, model:%+v, request:%+v", projectID, model, request)
	result, resp, err := client.ProjectIPAccessList.Create(context.Background(), projectID, request)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}
	log.Debugf("createEntries result:%+v", result)
	return handler.ProgressEvent{}, nil
}

func deleteEntries(model *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {
	projectID := *model.ProjectId

	for i, _ := range model.AccessList {
		wl := model.AccessList[i]
		entry := getEntry(wl)
		resp, err := client.ProjectIPAccessList.Delete(context.Background(), projectID, entry)
		if err != nil {
			return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
				resp.Response), err
		}
	}
	return handler.ProgressEvent{}, nil
}
