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
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	resource_constats "github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint-service/cmd/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint-service/cmd/resource/steps/privateendpointservice"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

func setup() {
	util.SetupLogger("mongodb-atlas-private-endpoint")
}

var CreateRequiredFields = []string{constants.ProjectID, constants.Region, constants.CloudProvider}
var ReadRequiredFields = []string{constants.ProjectID, constants.ID, constants.Region, constants.CloudProvider}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.ProjectID, constants.ID, constants.CloudProvider}
var ListRequiredFields = []string{constants.ProjectID}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	mongodbClient, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	status, pe := getProcessStatus(req)
	if pe != nil {
		return *pe, nil
	}

	switch status {
	case resource_constats.Init:
		pe := privateendpointservice.Create(*mongodbClient, *currentModel.Region, *currentModel.ProjectId, *currentModel.CloudProvider)
		return addModelToProgressEvent(&pe, currentModel), nil
	default:
		peConnection, completionValidation := privateendpointservice.ValidateCreationCompletion(mongodbClient,
			*currentModel.ProjectId, *currentModel.CloudProvider, req)
		if completionValidation != nil {
			return addModelToProgressEvent(completionValidation, currentModel), nil
		}

		currentModel.EndpointServiceName = &peConnection.EndpointServiceName
		currentModel.Id = &peConnection.ID
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Create Completed",
			ResourceModel:   currentModel}, nil
	}
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(ReadRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	mongodbClient, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.Get(context.Background(), *currentModel.ProjectId, *currentModel.CloudProvider, *currentModel.Id)
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response), nil
	}

	currentModel.completeByConnection(*privateEndpointResponse)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Get successful",
		ResourceModel:   currentModel}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	mongodbClient, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}
	privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.Get(context.Background(),
		*currentModel.ProjectId, *currentModel.CloudProvider, *currentModel.Id)

	if isDeleting(req) {
		if response.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Delete success"}, nil
		}

		if privateEndpointResponse != nil {
			return progressevent.GetInProgressProgressEvent("Delete in progress",
				map[string]interface{}{"stateName": "DELETING"}, currentModel, 20), nil
		}
	}
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response), nil
	}

	if privateEndpointResponse == nil {
		return progressevent.GetFailedEventByCode(fmt.Sprintf("Error deleting resource, private Endpoint Response is null : %s", err.Error()),
			cloudformation.HandlerErrorCodeNotFound), nil
	}

	response, err = mongodbClient.PrivateEndpoints.Delete(context.Background(), *currentModel.ProjectId,
		*currentModel.CloudProvider,
		*currentModel.Id)

	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              "Delete in progress",
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 20,
		CallbackContext: map[string]interface{}{
			"stateName": "DELETING",
		}}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(ListRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	mongodbClient, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.List(context.Background(),
		*currentModel.ProjectId,
		*currentModel.CloudProvider,
		params)
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error listing resource : %s", err.Error()),
			response.Response), nil
	}

	mm := make([]interface{}, 0, len(privateEndpointResponse))
	for i := range privateEndpointResponse {
		var m Model
		m.completeByConnection(privateEndpointResponse[i])
		m.Region = currentModel.Region
		m.Profile = currentModel.Profile
		m.ProjectId = currentModel.ProjectId
		mm = append(mm, m)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List successful",
		ResourceModels:  mm}, nil
}

func isDeleting(req handler.Request) bool {
	callback := req.CallbackContext["stateName"]
	if callback == nil {
		return false
	}

	callbackValue := fmt.Sprintf("%v", callback)
	return callbackValue == "DELETING"
}

func (m *Model) completeByConnection(c mongodbatlas.PrivateEndpointConnection) {
	m.Id = &c.ID
	m.EndpointServiceName = &c.EndpointServiceName
	m.ErrorMessage = &c.ErrorMessage
	m.Status = &c.Status
	m.InterfaceEndpoints = c.InterfaceEndpoints
	m.CloudProvider = &c.ProviderName
}

func getProcessStatus(req handler.Request) (resource_constats.EventStatus, *handler.ProgressEvent) {
	callback := req.CallbackContext["StateName"]
	if callback == nil {
		return resource_constats.Init, nil
	}

	eventStatus, err := resource_constats.ParseEventStatus(fmt.Sprintf("%v", callback))

	if err != nil {
		pe := progressevent.GetFailedEventByCode(fmt.Sprintf("Error parsing callback status : %s", err.Error()),
			cloudformation.HandlerErrorCodeServiceInternalError)
		return "", &pe
	}

	return eventStatus, nil
}

func addModelToProgressEvent(progressEvent *handler.ProgressEvent, model *Model) handler.ProgressEvent {
	if progressEvent.OperationStatus == handler.InProgress {
		progressEvent.ResourceModel = model

		callbackID := progressEvent.CallbackContext["ID"]

		if callbackID != nil {
			id := fmt.Sprint(callbackID)
			model.Id = &id
		}
	}

	return *progressEvent
}
