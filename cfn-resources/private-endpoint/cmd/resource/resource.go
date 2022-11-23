package resource

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/resource/steps/awsvpcendpoint"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/resource/steps/privateendpoint"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/resource/steps/privateendpointservice"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	resource_constats "github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	providerName = "AWS"
)

func setup() {
	util.SetupLogger("mongodb-atlas-private-endpoint")
}

var CreateRequiredFields = []string{constants.GroupID, constants.Region, constants.PubKey, constants.PvtKey, constants.SubnetID, constants.VPCID}
var ReadRequiredFields = []string{constants.GroupID, constants.ID, constants.PubKey, constants.PvtKey}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.GroupID, constants.ID, constants.PubKey, constants.PvtKey}
var ListRequiredFields = []string{constants.GroupID, constants.PubKey, constants.PvtKey}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	mongodbClient, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	status, pe := getProcessStatus(req)
	if pe != nil {
		return *pe, nil
	}

	switch status {
	case resource_constats.CreationInit:
		pe := privateendpointservice.Create(*mongodbClient, *currentModel.Region, *currentModel.GroupId)
		return addModelToProgressEvent(&pe, currentModel), nil
	case resource_constats.CreatingPrivateEndpointService:
		peConnection, completionValidation := privateendpointservice.ValidateCreationCompletion(mongodbClient,
			*currentModel.GroupId, req)
		if completionValidation != nil {
			return addModelToProgressEvent(completionValidation, currentModel), nil
		}

		vpcEndpointID, progressEvent := awsvpcendpoint.Create(req, *peConnection, *currentModel.Region,
			*currentModel.SubnetId, *currentModel.VpcId)
		if progressEvent != nil {
			return addModelToProgressEvent(progressEvent, currentModel), nil
		}

		pe := privateendpoint.Create(mongodbClient, *currentModel.GroupId, *vpcEndpointID, peConnection.ID)

		return addModelToProgressEvent(&pe, currentModel), nil
	default:
		ValidationOutput, progressEvent := privateendpoint.ValidateCreationCompletion(mongodbClient, *currentModel.GroupId, req)
		if progressEvent != nil {
			return addModelToProgressEvent(progressEvent, currentModel), nil
		}
		currentModel.Id = &ValidationOutput.ID
		currentModel.InterfaceEndpoints = ValidationOutput.InterfaceEndpoints
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

	mongodbClient, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.Get(context.Background(), *currentModel.GroupId, providerName, *currentModel.Id)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
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

	mongodbClient, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeNotFound), nil
	}

	privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.Get(context.Background(),
		*currentModel.GroupId, providerName, *currentModel.Id)

	//Todo: we can move this functionality in the private endpoint step
	if isDeleting(req) {
		if response.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Delete success"}, nil
		}

		if privateEndpointResponse != nil {
			return progress_events.GetInProgressProgressEvent("Delete in progress",
				map[string]interface{}{"stateName": "DELETING"}, currentModel, 20), nil
		}
	}

	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response), nil
	}

	currentModel.completeByConnection(*privateEndpointResponse)

	if currentModel.HasInterfaceEndpoints() {
		epr := privateendpoint.DeletePrivateEndpoints(mongodbClient, *currentModel.GroupId, *currentModel.Id,
			currentModel.InterfaceEndpoints)
		if epr != nil {
			return *epr, nil
		}

		_, epr = awsvpcendpoint.Delete(req, currentModel.InterfaceEndpoints, *currentModel.Region)
		if epr != nil {
			return *epr, nil
		}
	} else {
		response, err = mongodbClient.PrivateEndpoints.Delete(context.Background(), *currentModel.GroupId,
			providerName,
			*currentModel.Id)

		if err != nil {
			return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
				response.Response), nil
		}
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              "Delete in progress",
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 20,
		CallbackContext: map[string]interface{}{
			"stateName":         "DELETING",
			"AwsVpcEndpointIds": currentModel.InterfaceEndpoints,
		}}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(ListRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	mongodbClient, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.List(context.Background(),
		*currentModel.GroupId,
		providerName,
		params)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error listing resource : %s", err.Error()),
			response.Response), nil
	}

	mm := make([]interface{}, 0)
	for i := range privateEndpointResponse {
		var m Model
		m.completeByConnection(privateEndpointResponse[i])
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

func (m *Model) HasInterfaceEndpoints() bool {
	return len(m.InterfaceEndpoints) != 0
}

func (m *Model) completeByConnection(c mongodbatlas.PrivateEndpointConnection) {
	m.Id = &c.ID
	m.EndpointServiceName = &c.EndpointServiceName
	m.ErrorMessage = &c.ErrorMessage
	m.InterfaceEndpoints = c.InterfaceEndpoints
	m.Status = &c.Status
}

func getProcessStatus(req handler.Request) (resource_constats.EventStatus, *handler.ProgressEvent) {
	callback := req.CallbackContext["StateName"]
	if callback == nil {
		return resource_constats.CreationInit, nil
	}

	eventStatus, err := resource_constats.ParseEventStatus(fmt.Sprintf("%v", callback))

	if err != nil {
		pe := progress_events.GetFailedEventByCode(fmt.Sprintf("Error parsing callback status : %s", err.Error()),
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
