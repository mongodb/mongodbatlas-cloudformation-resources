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
	"errors"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115014/admin"
)

var CreateRequiredFields = []string{constants.ProjectID}
var ReadRequiredFields = []string{constants.ProjectID}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.ProjectID}
var ListRequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-private-endpoint-regional-mode")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	mongodbClient, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if isRegModeSettingExists(currentModel, mongodbClient) {
		return progressevent.GetFailedEventByCode(fmt.Sprintf("Regionalized Setting for Private Endpoint already enabled for : %s", *currentModel.ProjectId),
			cloudformation.HandlerErrorCodeAlreadyExists), nil
	}

	// API call to Add Regional Mode for Private Endpoint
	return resourcePrivateEndpointRegionalModeUpdate(currentModel, mongodbClient, true)
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	regPrivateEndpointSetting, response, err := client.AtlasSDK.PrivateEndpointServicesApi.GetRegionalizedPrivateEndpointSetting(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), response), nil
	}
	enabled := regPrivateEndpointSetting.Enabled
	if !enabled {
		return progressevent.GetFailedEventByCode(fmt.Sprintf("Regionalized Setting for Private Endpoint not found for Project : %s", *currentModel.ProjectId),
			cloudformation.HandlerErrorCodeNotFound), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "READ Complete",
		ResourceModel:   newResponseModel(*currentModel),
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	mongodbClient, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if isRegModeSettingExists(currentModel, mongodbClient) {
		events, err := resourcePrivateEndpointRegionalModeUpdate(currentModel, mongodbClient, false)
		if err != nil {
			return progressevent.GetFailedEventByCode(fmt.Sprintf("Error in disabling regionalized mode for private endpoint for Project : %s", *currentModel.ProjectId),
				events.HandlerErrorCode), nil
		}

		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Delete Complete",
		}, nil
	}

	return progressevent.GetFailedEventByCode(fmt.Sprintf("Error in disabling regionalized mode for private endpoint for Project : %s", *currentModel.ProjectId),
		cloudformation.HandlerErrorCodeNotFound), nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func resourcePrivateEndpointRegionalModeUpdate(currentModel *Model, client *util.MongoDBClient, enabled bool) (handler.ProgressEvent, error) {
	_, response, err := client.AtlasSDK.PrivateEndpointServicesApi.ToggleRegionalizedPrivateEndpointSetting(context.Background(), *currentModel.ProjectId,
		&admin.ProjectSettingItem{
			Enabled: enabled,
		}).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   newResponseModel(*currentModel),
	}, nil
}

func isRegModeSettingExists(currentModel *Model, client *util.MongoDBClient) bool {
	regModeSetting, _, err := client.AtlasSDK.PrivateEndpointServicesApi.GetRegionalizedPrivateEndpointSetting(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return false
	}

	return regModeSetting.Enabled
}

func newResponseModel(currentModel Model) *Model {
	out := &Model{
		ProjectId: currentModel.ProjectId,
		Profile:   currentModel.Profile,
	}
	return out
}
