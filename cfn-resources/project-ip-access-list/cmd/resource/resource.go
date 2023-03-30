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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
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

	if len(currentModel.AccessList) == 0 {
		return progressevents.GetFailedEventByCode("AccessList must not be empty", cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
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

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	result, resp, err := client.ProjectIPAccessList.List(context.Background(), *currentModel.ProjectId, nil)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	currentModel.TotalCount = &result.TotalCount
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

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

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

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	event := deleteEntries(currentModel, client)
	if event.OperationStatus == handler.Failed {
		_, _ = logger.Warnf("Delete error: %+v", event.Message)
		return event, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func (m *AccessListDefinition) completeByConnection(c mongodbatlas.ProjectIPAccessList) {
	m.IPAddress = &c.IPAddress
	m.CIDRBlock = &c.CIDRBlock
	m.Comment = &c.Comment
	m.AwsSecurityGroup = &c.AwsSecurityGroup
}

func getProjectIPAccessListRequest(model *Model) []*mongodbatlas.ProjectIPAccessList {
	var accesslist []*mongodbatlas.ProjectIPAccessList
	for i := range model.AccessList {
		modelAccessList := model.AccessList[i]
		projectIPAccessList := &mongodbatlas.ProjectIPAccessList{}

		if modelAccessList.DeleteAfterDate != nil {
			projectIPAccessList.DeleteAfterDate = *modelAccessList.DeleteAfterDate
		}

		if modelAccessList.Comment != nil {
			projectIPAccessList.Comment = *modelAccessList.Comment
		}

		if modelAccessList.CIDRBlock != nil {
			projectIPAccessList.CIDRBlock = *modelAccessList.CIDRBlock
		}

		if modelAccessList.IPAddress != nil {
			projectIPAccessList.IPAddress = *modelAccessList.IPAddress
		}

		if modelAccessList.AwsSecurityGroup != nil {
			projectIPAccessList.AwsSecurityGroup = *modelAccessList.AwsSecurityGroup
		}

		accesslist = append(accesslist, projectIPAccessList)
	}

	_, _ = logger.Debugf("getProjectIPAccessListRequest accesslist:%v", accesslist)
	return accesslist
}

func getEntry(entry AccessListDefinition) (string, error) {
	if entry.CIDRBlock != nil && *entry.CIDRBlock != "" {
		return *entry.CIDRBlock, nil
	}

	if entry.AwsSecurityGroup != nil && *entry.AwsSecurityGroup != "" {
		return *entry.AwsSecurityGroup, nil
	}

	if entry.IPAddress != nil && *entry.IPAddress != "" {
		return *entry.IPAddress, nil
	}

	return "", fmt.Errorf("AccessList entry must have one of the following fields: cidrBlock, awsSecurityGroup, ipAddress")
}

func createEntries(model *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {
	request := getProjectIPAccessListRequest(model)
	projectID := *model.ProjectId

	if _, _, err := client.ProjectIPAccessList.Create(context.Background(), projectID, request); err != nil {
		_, _ = logger.Warnf("Error createEntries projectId:%s, err:%+v", projectID, err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, err
	}

	return handler.ProgressEvent{}, nil
}

func deleteEntries(model *Model, client *mongodbatlas.Client) handler.ProgressEvent {
	for _, accessListEntry := range model.AccessList {
		entry, err := getEntry(accessListEntry)
		if err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting the resource : %s", err.Error()),
				nil)
		}

		resp, err := client.ProjectIPAccessList.Delete(context.Background(), *model.ProjectId, entry)
		if err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error deleting the resource: %s", err.Error()),
				resp.Response)
		}
	}

	return handler.ProgressEvent{}
}
