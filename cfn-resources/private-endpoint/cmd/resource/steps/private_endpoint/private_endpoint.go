package private_endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	ProviderName            = "AWS"
	StatusPendingAcceptance = "PENDING_ACCEPTANCE"
	StatusPending           = "PENDING"
	StatusAvailable         = "AVAILABLE"
)

// Todo: im not convinced about this resource, maybe we can find another way
type privateEndpointCreationCallBackContext struct {
	StateName   constants.EventStatus
	Id          string
	InterfaceId string
}

func (s *privateEndpointCreationCallBackContext) FillStruct(m map[string]interface{}) error {
	//Todo: we can unify this logic
	s.Id = fmt.Sprint(m["Id"])
	s.InterfaceId = fmt.Sprint(m["InterfaceId"])
	eventStatusParam := fmt.Sprint(m["StateName"])
	eventStatus, err := constants.ParseEventStatus(eventStatusParam)
	if err != nil {
		return err
	}

	s.StateName = eventStatus

	return nil
}

func Create(mongodbClient *mongodbatlas.Client, groupId string, interfaceEndpointID string, endpointServiceID string) handler.ProgressEvent {
	interfaceEndpointRequest := &mongodbatlas.InterfaceEndpointConnection{
		ID: interfaceEndpointID,
	}

	_, response, err := mongodbClient.PrivateEndpoints.AddOnePrivateEndpoint(context.Background(),
		groupId,
		ProviderName,
		endpointServiceID,
		interfaceEndpointRequest)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
			response.Response)
	}

	callBackContext := privateEndpointCreationCallBackContext{
		StateName:   constants.CreatingPrivateEndpoint,
		Id:          endpointServiceID,
		InterfaceId: interfaceEndpointID,
	}

	var callBackMap map[string]interface{}
	data, _ := json.Marshal(callBackContext)
	json.Unmarshal(data, &callBackMap)

	return progress_events.GetInProgressProgressEvent("Adding private endpoint", callBackMap, nil, 20)
}

func ValidateCreationCompletion(mongodbClient *mongodbatlas.Client, groupID string, req handler.Request) (*ValidationResponse, *handler.ProgressEvent) {
	callBackContext := privateEndpointCreationCallBackContext{}

	err := callBackContext.FillStruct(req.CallbackContext)
	if err != nil {
		pe := progress_events.GetFailedEventByCode(fmt.Sprintf("Error parsing PrivateEndpointCallBackContext : %s", err.Error()), cloudformation.HandlerErrorCodeServiceInternalError)
		return nil, &pe
	}

	privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.GetOnePrivateEndpoint(context.Background(),
		groupID,
		ProviderName,
		callBackContext.Id,
		callBackContext.InterfaceId)
	if err != nil {
		pe := progress_events.GetFailedEventByResponse(fmt.Sprintf("Error validating private endpoint create : %s", err.Error()),
			response.Response)
		return nil, &pe
	}

	switch privateEndpointResponse.AWSConnectionStatus {
	case StatusPendingAcceptance, StatusPending:
		{
			pe := progress_events.GetInProgressProgressEvent("Adding private endpoint in progress",
				req.CallbackContext, nil, 20)
			return nil, &pe
		}
	case StatusAvailable:
		{
			vr := ValidationResponse{
				Id:                 callBackContext.Id,
				InterfaceEndpoints: []string{callBackContext.InterfaceId},
			}
			return &vr, nil
		}
	}

	pe := handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          fmt.Sprintf("Resource is in status : %s", privateEndpointResponse.AWSConnectionStatus),
		HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}
	return nil, &pe
}

func DeletePrivateEndpoints(mongodbClient *mongodbatlas.Client, groupID string, id string, interfaceEndpoints []string) *handler.ProgressEvent {
	for _, intEndpoints := range interfaceEndpoints {
		response, err := mongodbClient.PrivateEndpoints.DeleteOnePrivateEndpoint(context.Background(),
			groupID,
			ProviderName,
			id,
			intEndpoints)
		if err != nil {
			pe := progress_events.GetFailedEventByResponse(fmt.Sprintf("Error deleting private endpoint : %s",
				err.Error()),
				response.Response)
			return &pe
		}
	}

	return nil
}

type ValidationResponse struct {
	Id                 string
	InterfaceEndpoints []string
}
