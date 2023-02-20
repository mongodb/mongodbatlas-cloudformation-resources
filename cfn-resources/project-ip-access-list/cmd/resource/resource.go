// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"context"
	"fmt"
	"github.com/openlyinc/pointy"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

func setup() {
	util.SetupLogger("mongodb-atlas-project-ip-access-list")
}

var CreateRequiredFields = []string{constants.ProjectID, constants.AccessList}
var ReadRequiredFields = []string{constants.ProjectID}
var UpdateRequiredFields = []string{constants.ProjectID, constants.AccessList}
var DeleteRequiredFields = []string{constants.ProjectID}
var ListRequiredFields = []string{constants.ProjectID}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(util.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	exists, peErr := existsEntries(currentModel, client)
	if peErr != nil {
		return *peErr, nil
	}

	if exists {
		return progressevents.GetFailedEventByCode("resource already exists", cloudformation.HandlerErrorCodeAlreadyExists), nil
	}

	event, err := createEntries(currentModel, client)
	if err != nil {
		_, _ = logger.Warnf("Create err:%v", err)
		return event, nil
	}

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

	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(util.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	projectID := *currentModel.ProjectId

	listOptions := &mongodbatlas.ListOptions{
		PageNum:      0,
		IncludeCount: true,
		ItemsPerPage: 200,
	}

	result, resp, err := client.ProjectIPAccessList.List(context.Background(), projectID, listOptions)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	if !accessListHasEntries(*result) {
		return progressevents.GetFailedEventByCode("Resource not found", cloudformation.HandlerErrorCodeNotFound), nil
	}

	mm := make([]AccessListDefinition, 0)
	for i := range result.Results {
		var m AccessListDefinition
		m.completeByConnection(result.Results[i])
		mm = append(mm, m)
	}
	currentModel.AccessList = mm

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

	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(util.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	progressEvent := deleteEntries(currentModel, client)
	if progressEvent.OperationStatus == handler.Failed {
		_, _ = logger.Warnf("Update deleteEntries error:%+v", progressEvent.Message)
		return progressEvent, nil
	}

	progressEvent, err := createEntries(currentModel, client)
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

	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(util.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	event := deleteEntries(currentModel, client)
	if event.OperationStatus == handler.Failed {
		_, _ = logger.Warnf("Delete deleteEntries error:%+v", event.Message)
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

	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(util.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
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

func getEntry(wl mongodbatlas.ProjectIPAccessList) string {
	if wl.CIDRBlock != "" {
		return wl.CIDRBlock
	}
	if wl.AwsSecurityGroup != "" {
		return wl.AwsSecurityGroup
	}
	if wl.IPAddress != "" {
		return wl.IPAddress
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

func deleteEntries(model *Model, client *mongodbatlas.Client) handler.ProgressEvent {
	projectID := *model.ProjectId

	listOptions := &mongodbatlas.ListOptions{
		PageNum:      0,
		IncludeCount: true,
		ItemsPerPage: 200,
	}

	result, resp, err := client.ProjectIPAccessList.List(context.Background(), projectID, listOptions)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response)
	}

	if !accessListHasEntries(*result) {
		return progressevents.GetFailedEventByCode("error trying to delete Resource not found", cloudformation.HandlerErrorCodeNotFound)
	}

	for i := range result.Results {
		wl := result.Results[i]
		entry := getEntry(wl)
		resp, err := client.ProjectIPAccessList.Delete(context.Background(), projectID, entry)
		if err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
				resp.Response)
		}
	}

	return handler.ProgressEvent{}
}

func existsEntries(model *Model, client *mongodbatlas.Client) (bool, *handler.ProgressEvent) {
	listOptions := &mongodbatlas.ListOptions{
		PageNum:      0,
		IncludeCount: true,
		ItemsPerPage: 1,
	}

	result, resp, err := client.ProjectIPAccessList.List(context.Background(), *model.ProjectId, listOptions)
	if err != nil {
		pe := progressevents.GetFailedEventByResponse(fmt.Sprintf("Validating if the resource already exists : %s", err.Error()),
			resp.Response)
		return false, &pe
	}

	return accessListHasEntries(*result), nil
}

func accessListHasEntries(list mongodbatlas.ProjectIPAccessLists) bool {
	return list.TotalCount != 0
}
