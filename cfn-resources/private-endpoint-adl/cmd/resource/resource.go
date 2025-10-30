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
	"net/http"

	admin20231115014 "go.mongodb.org/atlas-sdk/v20250312006/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

var RequiredFields = []string{constants.ProjectID, constants.EndpointID}
var ListRequiredFields = []string{constants.ProjectID}

// function to validate inputs to all actions
func validateAndDefaultRequest(fields []string, model *Model) *handler.ProgressEvent {
	if model.Type == nil {
		model.Type = aws.String(constants.DataLake)
	}
	if model.Provider == nil {
		model.Provider = aws.String(constants.AWS)
	}
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-private-endpoint-adl")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	validationError := validateAndDefaultRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	alreadyExists, pe := resourceAlreadyExists(*client, *currentModel)
	if pe != nil {
		return *pe, nil
	}

	if alreadyExists {
		return progressevent.GetFailedEventByCode("resource Already exists", cloudformation.HandlerErrorCodeAlreadyExists), nil
	}

	ctx := context.Background()

	requestBody := admin20231115014.PrivateNetworkEndpointIdEntry{
		Provider:   currentModel.Provider,
		Type:       currentModel.Type,
		EndpointId: *currentModel.EndpointId,
		Comment:    currentModel.Comment,
	}
	_, resp, err := client.AtlasSDK.DataFederationApi.CreateDataFederationPrivateEndpoint(ctx, *currentModel.ProjectId, &requestBody).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Created Private Link ADL",
		ResourceModel:   currentModel,
	}
	return event, nil
}

func resourceAlreadyExists(client util.MongoDBClient, currentModel Model) (bool, *handler.ProgressEvent) {
	_, resp, err := client.AtlasSDK.DataFederationApi.GetDataFederationPrivateEndpoint(context.Background(), *currentModel.ProjectId, *currentModel.EndpointId).Execute()
	if err != nil {
		if apiError, ok := admin20231115014.AsError(err); ok && apiError.Error == http.StatusNotFound {
			return false, nil
		}

		pe := progressevent.GetFailedEventByResponse(err.Error(), resp)
		return false, &pe
	}

	return true, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if currentModel.EndpointId == nil {
		return progressevent.GetFailedEventByResponse("required field missing. Resource not found", &http.Response{
			StatusCode: http.StatusNotFound,
		}), nil
	}
	validationError := validateAndDefaultRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()
	dlEndpoint, resp, err := client.AtlasSDK.DataFederationApi.GetDataFederationPrivateEndpoint(ctx, *currentModel.ProjectId, *currentModel.EndpointId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.Comment = dlEndpoint.Comment
	currentModel.Type = dlEndpoint.Type
	currentModel.Provider = dlEndpoint.Provider
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Private Link ADL",
		ResourceModel:   currentModel,
	}
	return event, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	validationError := validateAndDefaultRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()
	resp, err := client.AtlasSDK.DataFederationApi.DeleteDataFederationPrivateEndpoint(ctx, *currentModel.ProjectId, *currentModel.EndpointId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "delete data lake endpoint",
	}
	return event, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	validationError := validateAndDefaultRequest(ListRequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()
	list, resp, err := client.AtlasSDK.DataFederationApi.ListDataFederationPrivateEndpoints(ctx, *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}
	models := make([]any, 0, len(list.GetResults()))
	for _, v := range list.GetResults() {
		models = append(models, &Model{
			ProjectId:  currentModel.ProjectId,
			Profile:    currentModel.Profile,
			Comment:    v.Comment,
			EndpointId: admin20231115014.PtrString(v.GetEndpointId()),
			Provider:   v.Provider,
			Type:       v.Type,
		})
	}
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "list data lake endpoints",
		ResourceModels:  models,
	}
	return event, nil
}
