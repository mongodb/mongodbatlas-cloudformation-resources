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
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func setup() {
	util.SetupLogger("mongodb-atlas-private-endpoint")
}

const (
	Available         = "AVAILABLE"
	Rejected          = "REJECTED"
	EndpointServiceID = "EndpointServiceId"
	CloudProvider     = "AWS"
)

func IsTerminalStatus(status string) bool {
	// Convert the status to uppercase to handle case-insensitivity
	status = strings.ToUpper(status)

	// Check if the status is "AVAILABLE" or "REJECTED"
	return status == Available || status == Rejected
}

var CreateRequiredFields = []string{constants.ProjectID, EndpointServiceID, constants.ID}
var ReadRequiredFields = []string{constants.ProjectID, constants.ID}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.ProjectID, constants.ID, EndpointServiceID}
var ListRequiredFields = []string{constants.GroupID}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if currentModel.EnforceConnectionSuccess == nil {
		currentModel.EnforceConnectionSuccess = aws.Bool(true)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// progress callback setup
	if _, ok := req.CallbackContext["state"]; ok {
		privateEndpoint, response, peError := getPrivateEndpoint(client, currentModel)
		defer response.Body.Close()
		if peError != nil {
			return progress_events.GetFailedEventByResponse("Error getting Private Endpoint", response), nil
		}

		if IsTerminalStatus(*privateEndpoint.ConnectionStatus) {
			if currentModel.EnforceConnectionSuccess != nil && *currentModel.EnforceConnectionSuccess &&
				*privateEndpoint.ConnectionStatus == Rejected {
				return handler.ProgressEvent{
					OperationStatus: handler.Failed,
					Message:         fmt.Sprintf("Connection was Rejected : %s", *privateEndpoint.ErrorMessage),
					ResourceModel:   currentModel,
				}, nil
			}

			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Create Success",
				ResourceModel:   currentModel,
			}, nil
		}

		status := *privateEndpoint.ConnectionStatus
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              "Create in progress",
			ResourceModel:        currentModel,
			CallbackDelaySeconds: 20,
			CallbackContext: map[string]interface{}{
				"state": status,
			}}, nil
	}

	endpointRequest := admin.CreateEndpointRequest{
		Id: currentModel.Id,
	}

	privateEndpointRequest := client.AtlasSDK.PrivateEndpointServicesApi.CreatePrivateEndpoint(context.Background(), *currentModel.ProjectId,
		CloudProvider, *currentModel.EndpointServiceId, &endpointRequest)

	_, response, err := privateEndpointRequest.Execute()
	defer response.Body.Close()
	if err != nil {
		if response.StatusCode == http.StatusConflict {
			return progress_events.GetFailedEventByCode(fmt.Sprintf("error creating Serverless Private Endpoint %s",
					err.Error()), cloudformation.HandlerErrorCodeAlreadyExists),
				nil
		}
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("error creating Serverless Private Endpoint %s",
				err.Error()), response),
			nil
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              "Create in progress",
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 10,
		CallbackContext: map[string]interface{}{
			"state": "Pending",
		}}, nil
}

func getPrivateEndpoint(client *util.MongoDBClient, model *Model) (*admin.PrivateLinkEndpoint, *http.Response, error) {
	privateEndpointRequest := client.AtlasSDK.PrivateEndpointServicesApi.GetPrivateEndpoint(context.Background(), *model.ProjectId,
		CloudProvider, *model.Id, *model.EndpointServiceId)
	privateEndpoint, response, err := privateEndpointRequest.Execute()

	return privateEndpoint, response, err
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(ReadRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	privateEndpoint, response, err := getPrivateEndpoint(client, currentModel)
	defer response.Body.Close()
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("READ: Error getting private endpoint: %s", err.Error()), response), nil
	}

	currentModel.completeByAtlasModel(*privateEndpoint)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func (m *Model) completeByAtlasModel(privateEndpoint admin.PrivateLinkEndpoint) {
	m.ErrorMessage = privateEndpoint.ErrorMessage
	m.ConnectionStatus = privateEndpoint.ConnectionStatus
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

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// progress callback setup
	if _, ok := req.CallbackContext["state"]; ok {
		_, response, peError := getPrivateEndpoint(client, currentModel)
		defer response.Body.Close()
		if peError != nil {
			if response.StatusCode == http.StatusNotFound {
				return handler.ProgressEvent{
					OperationStatus: handler.Success,
					Message:         "Create Success",
				}, nil
			}
			return progress_events.GetFailedEventByResponse("Error validating Private Endpoint deletion progress", response), nil
		}

		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              "Create in progress",
			CallbackDelaySeconds: 20,
			CallbackContext: map[string]interface{}{
				"state": "deleting",
			}}, nil
	}

	privateEndpointRequest := client.AtlasSDK.PrivateEndpointServicesApi.DeletePrivateEndpoint(context.Background(), *currentModel.ProjectId,
		CloudProvider, *currentModel.Id, *currentModel.EndpointServiceId)
	_, response, err := privateEndpointRequest.Execute()
	defer response.Body.Close()
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("error creating Serverless Private Endpoint %s",
				err.Error()), response),
			nil
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              "Create in progress",
		CallbackDelaySeconds: 20,
		ResourceModel:        currentModel,
		CallbackContext: map[string]interface{}{
			"state": "deleting",
		}}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}
