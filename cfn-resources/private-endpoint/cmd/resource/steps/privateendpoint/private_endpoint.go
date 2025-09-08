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

package privateendpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

const (
	ProviderName            = "AWS"
	StatusPendingAcceptance = "PENDING_ACCEPTANCE"
	StatusPending           = "PENDING"
	StatusAvailable         = "AVAILABLE"
	StatusInitiating        = "INITIATING"
)

// Todo: im not convinced about this resource, maybe we can find another way
type privateEndpointCreationCallBackContext struct {
	StateName        constants.EventStatus
	ID               string
	PrivateEndpoints []AtlasPrivateEndpointCallBack
}

type AtlasPrivateEndpointCallBack struct {
	VpcID               string
	InterfaceEndpointID string
	Status              string
	SubnetIDs           []string
}

type AtlasPrivateEndpointInput struct {
	Status              *string
	VpcID               string
	InterfaceEndpointID string
	SubnetIDs           []string
}

func (s *privateEndpointCreationCallBackContext) FillStruct(m map[string]interface{}) error {
	s.ID = fmt.Sprint(m["ID"])
	eventStatusParam := fmt.Sprint(m["StateName"])
	eventStatus, err := constants.ParseEventStatus(eventStatusParam)
	if err != nil {
		return err
	}
	s.StateName = eventStatus

	privateEndpointI := m["PrivateEndpoints"]
	a := reflect.ValueOf(privateEndpointI)
	if a.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]AtlasPrivateEndpointCallBack, a.Len())

	for i := 0; i < a.Len(); i++ {
		v := reflect.ValueOf(a.Index(i).Interface())
		peCallback := AtlasPrivateEndpointCallBack{}

		for _, key := range v.MapKeys() {
			valStr := fmt.Sprint(v.MapIndex(key).Interface())
			switch key.String() {
			case "VpcID":
				peCallback.VpcID = valStr
			case "SubnetIDs":
				subnets := strings.Replace(strings.Replace(valStr, "[", "", 1), "]", "", 1)
				peCallback.SubnetIDs = strings.Split(subnets, " ")
			case "InterfaceEndpointID":
				peCallback.InterfaceEndpointID = valStr
			case "Status":
				peCallback.Status = valStr
			}
		}
		ret[i] = peCallback
	}

	s.PrivateEndpoints = ret

	return nil
}

func GetCallback(privateEndpointInput []AtlasPrivateEndpointInput, endpointServiceID string, state constants.EventStatus) (map[string]interface{}, error) {
	endpointCallBacks := make([]AtlasPrivateEndpointCallBack, len(privateEndpointInput))

	for i, pe := range privateEndpointInput {
		callBack := AtlasPrivateEndpointCallBack{
			VpcID:               pe.VpcID,
			SubnetIDs:           pe.SubnetIDs,
			InterfaceEndpointID: pe.InterfaceEndpointID,
		}

		if pe.Status != nil {
			callBack.Status = *pe.Status
		}

		endpointCallBacks[i] = callBack
	}

	callBackContext := privateEndpointCreationCallBackContext{
		StateName:        state,
		ID:               endpointServiceID,
		PrivateEndpoints: endpointCallBacks,
	}

	return getMapFromCallBackContext(callBackContext)
}

func getMapFromCallBackContext(callBackContext privateEndpointCreationCallBackContext) (map[string]interface{}, error) {
	var callBackMap map[string]interface{}
	data, _ := json.Marshal(callBackContext)
	err := json.Unmarshal(data, &callBackMap)

	return callBackMap, err
}

func Create(client *util.MongoDBClient, groupID string, privateEndpointInput []AtlasPrivateEndpointInput, endpointServiceID string) handler.ProgressEvent {
	for _, endpoint := range privateEndpointInput {
		interfaceEndpointRequest := &admin20231115002.CreateEndpointRequest{
			Id: &endpoint.InterfaceEndpointID,
		}

		_, response, err := client.Atlas20231115002.PrivateEndpointServicesApi.CreatePrivateEndpoint(context.Background(),
			groupID,
			ProviderName,
			endpointServiceID,
			interfaceEndpointRequest).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
				response)
		}
	}

	for i := range privateEndpointInput {
		status := StatusInitiating
		privateEndpointInput[i].Status = &status
	}

	callBackMap, err := GetCallback(privateEndpointInput, endpointServiceID, constants.CreatingPrivateEndpoint)
	if err != nil {
		return progressevent.GetFailedEventByCode(fmt.Sprintf("Error Unmarshalling callback map : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest)
	}

	return progressevent.GetInProgressProgressEvent("Adding private endpoint", callBackMap, nil, 20)
}

func ValidateCreationCompletion(client *util.MongoDBClient, groupID string, req handler.Request) (*ValidationResponse, *handler.ProgressEvent) {
	callBackContext := privateEndpointCreationCallBackContext{}

	err := callBackContext.FillStruct(req.CallbackContext)
	if err != nil {
		pe := progressevent.GetFailedEventByCode(fmt.Sprintf("Error parsing PrivateEndpointCallBackContext : %s", err.Error()), cloudformation.HandlerErrorCodeServiceInternalError)
		return nil, &pe
	}

	completed := true
	ids := make([]string, len(callBackContext.PrivateEndpoints))
	for i := range callBackContext.PrivateEndpoints {
		ids[i] = callBackContext.PrivateEndpoints[i].InterfaceEndpointID
		if callBackContext.PrivateEndpoints[i].Status != StatusAvailable {
			privateEndpointResponse, response, err := client.Atlas20231115002.PrivateEndpointServicesApi.GetPrivateEndpoint(context.Background(),
				groupID,
				ProviderName,
				callBackContext.PrivateEndpoints[i].InterfaceEndpointID,
				callBackContext.ID).Execute()
			if err != nil {
				pe := progressevent.GetFailedEventByResponse(fmt.Sprintf("Error validating private endpoint create : %s", err.Error()),
					response)
				return nil, &pe
			}
			callBackContext.PrivateEndpoints[i].Status = *privateEndpointResponse.ConnectionStatus

			switch *privateEndpointResponse.ConnectionStatus {
			case StatusPendingAcceptance, StatusPending:
				completed = false
			case StatusAvailable:
				continue
			default:
				pe := progressevent.GetFailedEventByCode(fmt.Sprintf("Resource is in status : %s", *privateEndpointResponse.Status),
					cloudformation.HandlerErrorCodeInternalFailure)
				return nil, &pe
			}
		}
	}

	if completed {
		endpoints := make([]AtlasPrivateEndpointCallBack, len(callBackContext.PrivateEndpoints))
		for i, v := range callBackContext.PrivateEndpoints {
			endpoints[i] = AtlasPrivateEndpointCallBack{
				VpcID:               v.VpcID,
				SubnetIDs:           v.SubnetIDs,
				InterfaceEndpointID: v.InterfaceEndpointID,
				Status:              v.Status,
			}
		}
		vr := ValidationResponse{
			ID:        callBackContext.ID,
			Endpoints: endpoints,
		}
		return &vr, nil
	}

	pe := progressevent.GetInProgressProgressEvent("Adding private endpoint in progress",
		req.CallbackContext, nil, 20)
	return nil, &pe
}

func (i AtlasPrivateEndpointInput) ToString() string {
	return fmt.Sprintf("%s%s", i.VpcID, i.SubnetIDs)
}

func Delete(client *util.MongoDBClient, groupID string, endpointServiceID string, interfaceEndpoints []string) *handler.ProgressEvent {
	for _, intEndpoints := range interfaceEndpoints {
		_, response, err := client.Atlas20231115002.PrivateEndpointServicesApi.DeletePrivateEndpoint(context.Background(),
			groupID,
			ProviderName,
			intEndpoints,
			endpointServiceID).Execute()
		if err != nil {
			pe := progressevent.GetFailedEventByResponse(fmt.Sprintf("Error deleting private endpoint : %s",
				err.Error()),
				response)
			return &pe
		}
	}

	return nil
}

type ValidationResponse struct {
	ID        string
	Endpoints []AtlasPrivateEndpointCallBack
}
