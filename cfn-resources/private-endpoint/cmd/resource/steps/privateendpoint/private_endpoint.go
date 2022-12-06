package privateendpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"go.mongodb.org/atlas/mongodbatlas"
	"reflect"
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
	VpcId               string
	SubnetId            string
	InterfaceEndpointId string
	Status              string
}

type AtlasPrivateEndpointInput struct {
	VpcId               string
	SubnetId            string
	InterfaceEndpointId string
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
		v := reflect.ValueOf(a.Index(i).Interface().(interface{}))
		peCallback := AtlasPrivateEndpointCallBack{}

		//TODO: this can be done with Reflection

		for _, key := range v.MapKeys() {
			valStr := fmt.Sprint(v.MapIndex(key).Interface())
			switch key.String() {
			case "VpcId":
				peCallback.VpcId = valStr
			case "SubnetId":
				peCallback.SubnetId = valStr
			case "InterfaceEndpointId":
				peCallback.InterfaceEndpointId = valStr
			case "Status":
				peCallback.Status = valStr
			}
		}
		ret[i] = peCallback
	}

	s.PrivateEndpoints = ret

	return nil
}

func Create(mongodbClient *mongodbatlas.Client, groupID string, privateEndpointInput []AtlasPrivateEndpointInput, endpointServiceID string) handler.ProgressEvent {
	for _, endpoint := range privateEndpointInput {
		interfaceEndpointRequest := &mongodbatlas.InterfaceEndpointConnection{
			ID: endpoint.InterfaceEndpointId,
		}

		_, response, err := mongodbClient.PrivateEndpoints.AddOnePrivateEndpoint(context.Background(),
			groupID,
			ProviderName,
			endpointServiceID,
			interfaceEndpointRequest)
		if err != nil {
			return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
				response.Response)
		}
	}

	endpointCallBacks := make([]AtlasPrivateEndpointCallBack, len(privateEndpointInput))

	for i, pe := range privateEndpointInput {
		endpointCallBacks[i] = AtlasPrivateEndpointCallBack{
			VpcId:               pe.VpcId,
			SubnetId:            pe.SubnetId,
			InterfaceEndpointId: pe.InterfaceEndpointId,
			Status:              StatusInitiating,
		}
	}

	callBackContext := privateEndpointCreationCallBackContext{
		StateName:        constants.CreatingPrivateEndpoint,
		ID:               endpointServiceID,
		PrivateEndpoints: endpointCallBacks,
	}

	var callBackMap map[string]interface{}
	data, _ := json.Marshal(callBackContext)
	err := json.Unmarshal(data, &callBackMap)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error Unmarshalling callback map : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest)
	}

	return progress_events.GetInProgressProgressEvent("Adding private endpoint", callBackMap, nil, 20)
}

func ValidateCreationCompletion(mongodbClient *mongodbatlas.Client, groupID string, req handler.Request) (*ValidationResponse, *handler.ProgressEvent) {
	callBackContext := privateEndpointCreationCallBackContext{}

	err := callBackContext.FillStruct(req.CallbackContext)
	if err != nil {
		pe := progress_events.GetFailedEventByCode(fmt.Sprintf("Error parsing PrivateEndpointCallBackContext : %s", err.Error()), cloudformation.HandlerErrorCodeServiceInternalError)
		return nil, &pe
	}

	completed := true
	ids := make([]string, len(callBackContext.PrivateEndpoints))

	for i := range callBackContext.PrivateEndpoints {
		ids[i] = callBackContext.PrivateEndpoints[i].InterfaceEndpointId
		if callBackContext.PrivateEndpoints[i].Status != StatusAvailable {
			privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.GetOnePrivateEndpoint(context.Background(),
				groupID,
				ProviderName,
				callBackContext.ID,
				callBackContext.PrivateEndpoints[i].InterfaceEndpointId)
			if err != nil {
				pe := progress_events.GetFailedEventByResponse(fmt.Sprintf("Error validating private endpoint create : %s", err.Error()),
					response.Response)
				return nil, &pe
			}

			callBackContext.PrivateEndpoints[i].Status = privateEndpointResponse.AWSConnectionStatus

			switch privateEndpointResponse.AWSConnectionStatus {
			case StatusPendingAcceptance, StatusPending:
				completed = false
			case StatusAvailable:
				continue
			default:
				pe := progress_events.GetFailedEventByCode(fmt.Sprintf("Resource is in status : %s", privateEndpointResponse.AWSConnectionStatus),
					cloudformation.HandlerErrorCodeInternalFailure)
				return nil, &pe
			}
		}
	}

	if completed {
		endpoints := make([]AtlasPrivateEndpointCallBack, len(callBackContext.PrivateEndpoints))
		for i, v := range callBackContext.PrivateEndpoints {
			endpoints[i] = AtlasPrivateEndpointCallBack{
				VpcId:               v.VpcId,
				SubnetId:            v.SubnetId,
				InterfaceEndpointId: v.InterfaceEndpointId,
				Status:              v.Status,
			}
		}
		vr := ValidationResponse{
			ID:        callBackContext.ID,
			Endpoints: endpoints,
		}
		return &vr, nil
	} else {
		pe := progress_events.GetInProgressProgressEvent("Adding private endpoint in progress",
			req.CallbackContext, nil, 20)
		return nil, &pe
	}
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
	ID        string
	Endpoints []AtlasPrivateEndpointCallBack
}
