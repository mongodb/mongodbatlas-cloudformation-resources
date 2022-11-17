package private_endpoint_service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/constants"
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
	Id        string
}

func (s *privateEndpointCreationCallBackContext) FillStruct(m map[string]interface{}) error {
	s.Id = fmt.Sprint(m["Id"])
	eventStatusParam := fmt.Sprint(m["StateName"])
	eventStatus, err := constants.ParseEventStatus(eventStatusParam)
	if err != nil {
		return err
	}

	s.StateName = eventStatus

	return nil
}

func Create(mongodbClient mongodbatlas.Client, region string, groupId string) handler.ProgressEvent {

	privateEndpointRequest := &mongodbatlas.PrivateEndpointConnection{
		ProviderName: ProviderName,
		Region:       region,
	}

	privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.Create(context.Background(),
		groupId,
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
		Id:        privateEndpointResponse.ID,
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

func ValidateCreationCompletion(mongodbClient *mongodbatlas.Client, groupId string, req handler.Request) (*mongodbatlas.PrivateEndpointConnection, *handler.ProgressEvent) {

	PrivateEndpointCallBackContext := privateEndpointCreationCallBackContext{}

	err := PrivateEndpointCallBackContext.FillStruct(req.CallbackContext)
	if err != nil {
		ev := progress_events.GetFailedEventByCode(fmt.Sprintf("Error parsing PrivateEndpointCallBackContext : %s", err.Error()),
			cloudformation.HandlerErrorCodeServiceInternalError)
		return nil, &ev
	}

	privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.Get(context.Background(), groupId,
		ProviderName, PrivateEndpointCallBackContext.Id)
	if err != nil {
		ev := progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response)
		return nil, &ev
	}

	if privateEndpointResponse.Status == InitiatingStatus {
		callBackContext := privateEndpointCreationCallBackContext{
			StateName: constants.CreatingPrivateEndpointService,
			Id:        privateEndpointResponse.ID,
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
	} else if privateEndpointResponse.Status == AvailableStatus {
		return privateEndpointResponse, nil
	} else {
		ev := progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating private endpoint in status : %s",
			privateEndpointResponse.Status),
			cloudformation.HandlerErrorCodeInvalidRequest)

		return nil, &ev
	}
}
