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

	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	cloudformationtypes "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

var RequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-custom-dns-configuration-cluster-aws")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(RequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if isCustomAWSDNSSettingExists(currentModel, client) {
		return progressevent.GetFailedEventByCode(fmt.Sprintf("Custom AWS dns settings already enabled for : %s", *currentModel.ProjectId),
			string(cloudformationtypes.HandlerErrorCodeAlreadyExists)), nil
	}
	// API call to
	enabled := true
	currentModel.Enabled = &enabled
	return resourceCustomAWSDNSUpdate(req, prevModel, currentModel, client)
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(RequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	customAWSDNSSetting, response, err := client.Atlas20231115002.AWSClustersDNSApi.GetAWSCustomDNS(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error reading  : %s", err.Error()),
			response), nil
	}
	enabled := customAWSDNSSetting.Enabled
	if !enabled {
		return progressevent.GetFailedEventByCode(fmt.Sprintf("Custom AWS dns settings not found for Project : %s", *currentModel.ProjectId),
			string(cloudformationtypes.HandlerErrorCodeNotFound)), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "READ Complete",
		ResourceModel:   customAWSDNSToModel(*currentModel, customAWSDNSSetting),
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

	if errEvent := validator.ValidateModel(RequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if isCustomAWSDNSSettingExists(currentModel, client) {
		enabled := false
		currentModel.Enabled = &enabled
		events, err := resourceCustomAWSDNSUpdate(req, prevModel, currentModel, client)
		if err != nil {
			return progressevent.GetFailedEventByCode(fmt.Sprintf("Error in disabling regionalized mode for private endpoint for Project : %s", *currentModel.ProjectId),
				events.HandlerErrorCode), nil
		}

		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Delete Complete",
		}, nil
	}
	return progressevent.GetFailedEventByCode(fmt.Sprintf("Error in disabling Custom AWS DNS settings for Project : %s", *currentModel.ProjectId),
		string(cloudformationtypes.HandlerErrorCodeNotFound)), nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func resourceCustomAWSDNSUpdate(req handler.Request, prevModel *Model, currentModel *Model, client *util.MongoDBClient) (handler.ProgressEvent, error) {
	customAWSDNSRequest := &admin20231115002.AWSCustomDNSEnabled{
		Enabled: *currentModel.Enabled,
	}
	customAWSDNSModel, response, err := client.Atlas20231115002.AWSClustersDNSApi.ToggleAWSCustomDNS(context.Background(), *currentModel.ProjectId, customAWSDNSRequest).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(
			fmt.Sprintf("Error in enabling Custom AWS DNS settings : %s", err.Error()),
			response), nil
	}
	currentModel.Enabled = &customAWSDNSModel.Enabled

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   customAWSDNSToModel(*currentModel, customAWSDNSModel),
	}, nil
}

func isCustomAWSDNSSettingExists(currentModel *Model, client *util.MongoDBClient) bool {
	var isExists bool
	customAWSDNSSetting, _, err := client.Atlas20231115002.AWSClustersDNSApi.GetAWSCustomDNS(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return isExists
	}
	if customAWSDNSSetting.Enabled {
		isExists = true
	}
	return isExists
}

func customAWSDNSToModel(currentModel Model, regPrivateMode *admin20231115002.AWSCustomDNSEnabled) *Model {
	out := &Model{
		Profile:   currentModel.Profile,
		Enabled:   &regPrivateMode.Enabled,
		ProjectId: currentModel.ProjectId,
	}
	return out
}
