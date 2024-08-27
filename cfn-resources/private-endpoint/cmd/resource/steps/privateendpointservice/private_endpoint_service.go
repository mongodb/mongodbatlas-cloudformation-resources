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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"go.mongodb.org/atlas-sdk/v20240805001/admin"
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

func Create(client util.MongoDBClient, region string, groupID string) handler.ProgressEvent {
	privateEndpointRequest := &admin.CloudProviderEndpointServiceRequest{
		ProviderName: ProviderName,
		Region:       region,
	}

	privateEndpointResponse, response, err := client.Atlas20231115002.PrivateEndpointServicesApi.CreatePrivateEndpointService(
		context.Background(),
		groupID,
		privateEndpointRequest).Execute()

	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource already exists",
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}
	}

	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
			response)
	}

	callBackContext := privateEndpointCreationCallBackContext{
		StateName: constants.CreatingPrivateEndpointService,
		ID:        *privateEndpointResponse.Id,
	}

	var callBackMap map[string]interface{}
	data, _ := json.Marshal(callBackContext)
	err = json.Unmarshal(data, &callBackMap)
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error Unmarshalling callback map : %s", err.Error()),
			response)
	}

	return progressevent.GetInProgressProgressEvent("Creating private endpoint service", callBackMap,
		nil, 20)
}

func ValidateCreationCompletion(client *util.MongoDBClient, groupID string, req handler.Request) (*admin.EndpointService, *handler.ProgressEvent) {
	PrivateEndpointCallBackContext := privateEndpointCreationCallBackContext{}

	err := PrivateEndpointCallBackContext.FillStruct(req.CallbackContext)
	if err != nil {
		ev := progressevent.GetFailedEventByCode(fmt.Sprintf("Error parsing PrivateEndpointCallBackContext : %s", err.Error()),
			cloudformation.HandlerErrorCodeServiceInternalError)
		return nil, &ev
	}

	privateEndpointResponse, response, err := client.Atlas20231115002.PrivateEndpointServicesApi.GetPrivateEndpointService(context.Background(), groupID,
		ProviderName, PrivateEndpointCallBackContext.ID).Execute()
	if err != nil {
		ev := progressevent.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response)
		return nil, &ev
	}

	switch *privateEndpointResponse.Status {
	case InitiatingStatus:
		callBackContext := privateEndpointCreationCallBackContext{
			StateName: constants.CreatingPrivateEndpointService,
			ID:        *privateEndpointResponse.Id,
		}

		var callBackMap map[string]interface{}
		data, _ := json.Marshal(callBackContext)
		err = json.Unmarshal(data, &callBackMap)
		if err != nil {
			ev := progressevent.GetFailedEventByResponse(fmt.Sprintf("Error Unmarshalling callback map : %s", err.Error()),
				response)
			return nil, &ev
		}

		ev := progressevent.GetInProgressProgressEvent("Private endpoint service initiating", callBackMap,
			nil, 20)
		return nil, &ev
	case AvailableStatus:
		return privateEndpointResponse, nil
	default:
		ev := progressevent.GetFailedEventByCode(fmt.Sprintf("Error creating private endpoint in status : %s",
			*privateEndpointResponse.Status),
			cloudformation.HandlerErrorCodeInvalidRequest)
		return nil, &ev
	}
}
