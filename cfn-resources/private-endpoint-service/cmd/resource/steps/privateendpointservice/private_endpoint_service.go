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

package privateendpointservice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint-service/cmd/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	ProviderName     = "AWS"
	AvailableStatus  = "AVAILABLE"
	InitiatingStatus = "INITIATING"
)

type privateEndpointCreationCallBackContext struct {
	StateName constants.EventStatus
	ID        string
}

func (s *privateEndpointCreationCallBackContext) FillStruct(m map[string]interface{}) error {
	s.ID = fmt.Sprint(m["ID"])
	eventStatusParam := fmt.Sprint(m["StateName"])
	eventStatus, err := constants.ParseEventStatus(eventStatusParam)
	if err != nil {
		return err
	}

	s.StateName = eventStatus

	return nil
}

func Create(mongodbClient mongodbatlas.Client, region string, groupID string) handler.ProgressEvent {
	privateEndpointRequest := &mongodbatlas.PrivateEndpointConnection{
		ProviderName: ProviderName,
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
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
			response.Response)
	}

	callBackContext := privateEndpointCreationCallBackContext{
		StateName: constants.CreatingPrivateEndpointService,
		ID:        privateEndpointResponse.ID,
	}

	var callBackMap map[string]interface{}
	data, _ := json.Marshal(callBackContext)
	err = json.Unmarshal(data, &callBackMap)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error Unmarshalling callback map : %s", err.Error()),
			response.Response)
	}

	return progress_events.GetInProgressProgressEvent("Creating private endpoint service", callBackMap,
		nil, 20)
}

func ValidateCreationCompletion(mongodbClient *mongodbatlas.Client, groupID string, req handler.Request) (*mongodbatlas.PrivateEndpointConnection, *handler.ProgressEvent) {
	PrivateEndpointCallBackContext := privateEndpointCreationCallBackContext{}

	err := PrivateEndpointCallBackContext.FillStruct(req.CallbackContext)
	if err != nil {
		ev := progress_events.GetFailedEventByCode(fmt.Sprintf("Error parsing PrivateEndpointCallBackContext : %s", err.Error()),
			cloudformation.HandlerErrorCodeServiceInternalError)
		return nil, &ev
	}

	privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.Get(context.Background(), groupID,
		ProviderName, PrivateEndpointCallBackContext.ID)
	if err != nil {
		ev := progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response)
		return nil, &ev
	}

	switch privateEndpointResponse.Status {
	case InitiatingStatus:
		callBackContext := privateEndpointCreationCallBackContext{
			StateName: constants.CreatingPrivateEndpointService,
			ID:        privateEndpointResponse.ID,
		}

		var callBackMap map[string]interface{}
		data, _ := json.Marshal(callBackContext)
		err = json.Unmarshal(data, &callBackMap)
		if err != nil {
			ev := progress_events.GetFailedEventByResponse(fmt.Sprintf("Error Unmarshalling callback map : %s", err.Error()),
				response.Response)
			return nil, &ev
		}

		ev := progress_events.GetInProgressProgressEvent("Private endpoint service initiating", callBackMap,
			nil, 20)
		return nil, &ev
	case AvailableStatus:
		return privateEndpointResponse, nil
	default:
		ev := progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating private endpoint in status : %s",
			privateEndpointResponse.Status),
			cloudformation.HandlerErrorCodeInvalidRequest)
		return nil, &ev
	}
}
