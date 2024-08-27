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
	resource_constats "github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/resource/steps/awsvpcendpoint"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/resource/steps/privateendpoint"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/resource/steps/privateendpointservice"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20240805001/admin"
)

const (
	providerName = "AWS"
)

func setup() {
	util.SetupLogger("mongodb-atlas-private-endpoint")
}

var CreateRequiredFields = []string{constants.GroupID, constants.Region}
var ReadRequiredFields = []string{constants.GroupID, constants.ID, constants.Region}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.GroupID, constants.ID}
var ListRequiredFields = []string{constants.GroupID}

func (m *Model) newAwsPrivateEndpointInput() []awsvpcendpoint.AwsPrivateEndpointInput {
	awsInput := make([]awsvpcendpoint.AwsPrivateEndpointInput, len(m.PrivateEndpoints))

	for i, ep := range m.PrivateEndpoints {
		subnetIds := make([]string, len(ep.SubnetIds))

		copy(subnetIds, m.PrivateEndpoints[i].SubnetIds)

		endpoint := awsvpcendpoint.AwsPrivateEndpointInput{
			VpcID:               *ep.VpcId,
			SubnetIDs:           subnetIds,
			InterfaceEndpointID: ep.InterfaceEndpointId,
		}

		awsInput[i] = endpoint
	}

	return awsInput
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	status, pe := getProcessStatus(req)
	if pe != nil {
		return *pe, nil
	}

	switch status {
	case resource_constats.Init:
		pe := privateendpointservice.Create(*client, *currentModel.Region, *currentModel.GroupId)
		return addModelToProgressEvent(&pe, currentModel), nil
	case resource_constats.CreatingPrivateEndpointService:
		peConnection, completionValidation := privateendpointservice.ValidateCreationCompletion(client,
			*currentModel.GroupId, req)
		if completionValidation != nil {
			return addModelToProgressEvent(completionValidation, currentModel), nil
		}

		awsPrivateEndpointOutput, progressEvent := awsvpcendpoint.Create(req, *peConnection.EndpointServiceName, *currentModel.Region,
			currentModel.newAwsPrivateEndpointInput())
		if progressEvent != nil {
			return addModelToProgressEvent(progressEvent, currentModel), nil
		}

		privateEndpointInput := make([]privateendpoint.AtlasPrivateEndpointInput, len(awsPrivateEndpointOutput))

		for i, awsPe := range awsPrivateEndpointOutput {
			privateEndpointInput[i] = privateendpoint.AtlasPrivateEndpointInput{
				VpcID:               awsPe.VpcID,
				SubnetIDs:           awsPe.SubnetIDs,
				InterfaceEndpointID: awsPe.InterfaceEndpointID,
			}
		}

		pe := privateendpoint.Create(client, *currentModel.GroupId, privateEndpointInput, *peConnection.Id)

		return addModelToProgressEvent(&pe, currentModel), nil
	default:
		ValidationOutput, progressEvent := privateendpoint.ValidateCreationCompletion(client, *currentModel.GroupId, req)
		if progressEvent != nil {
			return addModelToProgressEvent(progressEvent, currentModel), nil
		}

		currentModel.Id = &ValidationOutput.ID

		for _, cmpe := range currentModel.PrivateEndpoints {
			for i := range ValidationOutput.Endpoints {
				vpe := ValidationOutput.Endpoints[i]
				if vpe.VpcID == *cmpe.VpcId && CompareSlices(vpe.SubnetIDs, cmpe.SubnetIds) {
					currentModel.PrivateEndpoints[i].InterfaceEndpointId = &vpe.InterfaceEndpointID
				}
			}
		}

		privateEndpointResponse, response, err := client.Atlas20231115002.PrivateEndpointServicesApi.
			GetPrivateEndpointService(context.Background(), *currentModel.GroupId, providerName, *currentModel.Id).
			Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
				response), nil
		}
		currentModel.EndpointServiceName = privateEndpointResponse.EndpointServiceName

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
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	privateEndpointResponse, response, err := client.Atlas20231115002.PrivateEndpointServicesApi.
		GetPrivateEndpointService(context.Background(), *currentModel.GroupId, providerName, *currentModel.Id).
		Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response), nil
	}

	currentModel.completeByConnection(privateEndpointResponse)

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
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}
	privateEndpointResponse, response, err := client.Atlas20231115002.PrivateEndpointServicesApi.GetPrivateEndpointService(context.Background(),
		*currentModel.GroupId, providerName, *currentModel.Id).Execute()

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
			response), nil
	}

	if privateEndpointResponse == nil {
		return progressevent.GetFailedEventByCode(fmt.Sprintf("Error deleting resource, private Endpoint Response is null : %s", err.Error()),
			cloudformation.HandlerErrorCodeNotFound), nil
	}

	privateEndpoint := *privateEndpointResponse

	if hasInterfaceEndpoints(privateEndpoint) {
		epr := privateendpoint.Delete(client, *currentModel.GroupId, *currentModel.Id,
			privateEndpoint.InterfaceEndpoints)

		if epr != nil {
			return *epr, nil
		}
		epr = awsvpcendpoint.Delete(req, privateEndpoint.InterfaceEndpoints, *currentModel.Region)
		if epr != nil {
			return *epr, nil
		}
	} else {
		_, response, err = client.Atlas20231115002.PrivateEndpointServicesApi.DeletePrivateEndpointService(context.Background(), *currentModel.GroupId,
			providerName,
			*currentModel.Id).Execute()

		if err != nil {
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
				response), nil
		}
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
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	privateEndpointResponse, response, err := client.Atlas20231115002.PrivateEndpointServicesApi.ListPrivateEndpointServices(context.Background(),
		*currentModel.GroupId,
		providerName).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error listing resource : %s", err.Error()),
			response), nil
	}

	mm := make([]interface{}, 0, len(privateEndpointResponse))
	for i := range privateEndpointResponse {
		var m Model
		m.completeByConnection(&privateEndpointResponse[i])
		m.Region = currentModel.Region
		m.Profile = currentModel.Profile
		m.GroupId = currentModel.GroupId
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

func hasInterfaceEndpoints(p admin.EndpointService) bool {
	return len(p.InterfaceEndpoints) != 0
}

func (m *Model) completeByConnection(c *admin.EndpointService) {
	m.Id = c.Id
	m.EndpointServiceName = c.EndpointServiceName
	m.ErrorMessage = c.ErrorMessage
	m.Status = c.Status

	copy(m.InterfaceEndpoints, c.InterfaceEndpoints)
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

func CompareSlices(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, v := range a {
		found := false
		for _, c := range b {
			if v == c {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
