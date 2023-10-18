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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	ProgressStatusCreating = "CREATING"
	ProgressStatusDeleting = "DELETING"
	AvailableStatus        = "AVAILABLE"
	InitiatingStatus       = "INITIATING"
)

func setup() {
	util.SetupLogger("mongodb-atlas-private-endpoint")
}

var CreateRequiredFields = []string{constants.ProjectID, constants.Region, constants.CloudProvider}
var ReadRequiredFields = []string{constants.ProjectID, constants.ID, constants.Region, constants.CloudProvider}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.ProjectID, constants.ID, constants.CloudProvider}
var ListRequiredFields = []string{constants.ProjectID, constants.CloudProvider}

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

	if isCreating(req) {
		print("In progress")
		return validateCreationCompletion(mongodbClient,
			currentModel, req), nil
	}

	return create(*mongodbClient, currentModel), nil
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
		m.CloudProvider = currentModel.CloudProvider
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
	return callbackValue == ProgressStatusDeleting
}

func isCreating(req handler.Request) bool {
	callback := req.CallbackContext["StateName"]
	if callback == nil {
		return false
	}
	println("got callback")

	callbackValue := fmt.Sprintf("%v", callback)
	println(callbackValue)
	return callbackValue == ProgressStatusCreating
}

func (m *Model) completeByConnection(c mongodbatlas.PrivateEndpointConnection) {
	m.Id = &c.ID
	m.EndpointServiceName = &c.EndpointServiceName
	m.ErrorMessage = &c.ErrorMessage
	m.Status = &c.Status
	m.InterfaceEndpoints = c.InterfaceEndpoints
}

type privateEndpointCreationCallBackContext struct {
	StateName string
	ID        string
}

func create(mongodbClient mongodbatlas.Client, currentModel *Model) handler.ProgressEvent {

	region := *currentModel.Region
	groupID := *currentModel.ProjectId
	cloudProvider := *currentModel.CloudProvider

	privateEndpointRequest := &mongodbatlas.PrivateEndpointConnection{
		ProviderName: cloudProvider,
		Region:       region,
	}

	privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.Create(context.Background(),
		groupID,
		privateEndpointRequest)

	if response.Response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource already exists",
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}
	}

	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
			response.Response)
	}

	callBackContext := privateEndpointCreationCallBackContext{
		StateName: ProgressStatusCreating,
		ID:        privateEndpointResponse.ID,
	}

	callBackMap, err := callBackContext.convertToInterface()
	if err != nil {
		progressevent.GetFailedEventByCode(fmt.Sprintf("Error Unmarshalling callback map : %s", err.Error()),
			cloudformation.HandlerErrorCodeServiceInternalError)
	}

	currentModel.completeByConnection(*privateEndpointResponse)
	return progressevent.GetInProgressProgressEvent("Creating private endpoint service", callBackMap,
		currentModel, 20)
}

func validateCreationCompletion(mongodbClient *mongodbatlas.Client, currentModel *Model, req handler.Request) handler.ProgressEvent {
	PrivateEndpointCallBackContext := privateEndpointCreationCallBackContext{}
	PrivateEndpointCallBackContext.fillStruct(req.CallbackContext)
	print("getting private endpoint")
	println(*currentModel.ProjectId)
	println(*currentModel.CloudProvider)
	println(PrivateEndpointCallBackContext.ID)
	privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.Get(context.Background(), *currentModel.ProjectId,
		*currentModel.CloudProvider, PrivateEndpointCallBackContext.ID)
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response)
	}

	println("Completing connection")

	currentModel.completeByConnection(*privateEndpointResponse)

	switch privateEndpointResponse.Status {
	case InitiatingStatus:
		callBackContext := privateEndpointCreationCallBackContext{
			StateName: ProgressStatusCreating,
			ID:        privateEndpointResponse.ID,
		}

		callBackMap, err := callBackContext.convertToInterface()
		if err != nil {
			progressevent.GetFailedEventByCode(fmt.Sprintf("Error Unmarshalling callback map : %s", err.Error()),
				cloudformation.HandlerErrorCodeServiceInternalError)
		}

		return progressevent.GetInProgressProgressEvent("Private endpoint service initiating", callBackMap,
			currentModel, 20)
	case AvailableStatus:
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Create Completed",
			ResourceModel:   currentModel}
	default:
		return progressevent.GetFailedEventByCode(fmt.Sprintf("Error creating private endpoint in status : %s",
			privateEndpointResponse.Status),
			cloudformation.HandlerErrorCodeInvalidRequest)
	}
}

func (callBackContext *privateEndpointCreationCallBackContext) convertToInterface() (map[string]interface{}, error) {
	var callBackMap map[string]interface{}
	data, _ := json.Marshal(callBackContext)
	err := json.Unmarshal(data, &callBackMap)
	if err != nil {
		return nil, err
	}

	return callBackMap, nil
}

func (callBackContext *privateEndpointCreationCallBackContext) fillStruct(m map[string]interface{}) {
	println("imprimiendo mierda")
	println(fmt.Sprint(m["ID"]))
	println(fmt.Sprint(m["StateName"]))
	callBackContext.ID = fmt.Sprint(m["ID"])
	callBackContext.StateName = fmt.Sprint(m["StateName"])
}
