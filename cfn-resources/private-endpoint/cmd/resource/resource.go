package resource

import (
	"context"
	"fmt"
	"log"
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

var CreateRequiredFields = []string{constants.GroupID, constants.Region, constants.PubKey, constants.PvtKey}
var ReadRequiredFields = []string{constants.GroupID, constants.ID, constants.PubKey, constants.PvtKey}
var UpdateRequiredFields = []string{constants.GroupID, constants.ID, constants.PubKey, constants.PvtKey}
var DeleteRequiredFields = []string{constants.GroupID, constants.ID, constants.PubKey, constants.PvtKey}
var ListRequiredFields = []string{constants.GroupID, constants.PubKey, constants.PvtKey}

func (m *Model) newAwsPrivateEndpointInput() []awsvpcendpoint.AwsPrivateEndpointInput {
	awsInput := make([]awsvpcendpoint.AwsPrivateEndpointInput, len(m.PrivateEndpoints))

	log.Printf("PrivateEndpoints %v, -- %v", len(m.PrivateEndpoints), m.PrivateEndpoints)

	for i, ep := range m.PrivateEndpoints {
		endpoint := awsvpcendpoint.AwsPrivateEndpointInput{
			VpcId:               *ep.VpcId,
			SubnetId:            *ep.SubnetId,
			InterfaceEndpointId: ep.InterfaceEndpointId,
		}

		awsInput[i] = endpoint
	}

	log.Printf("PrivateEndpoints %v, -- %v", len(awsInput), awsInput)

	return awsInput
}

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
	case resource_constats.Init:
		pe := privateendpointservice.Create(*mongodbClient, *currentModel.Region, *currentModel.GroupId)
		return addModelToProgressEvent(&pe, currentModel), nil
	case resource_constats.CreatingPrivateEndpointService:
		peConnection, completionValidation := privateendpointservice.ValidateCreationCompletion(mongodbClient,
			*currentModel.GroupId, req)
		if completionValidation != nil {
			return addModelToProgressEvent(completionValidation, currentModel), nil
		}

		awsPrivateEndpointOutput, progressEvent := awsvpcendpoint.Create(req, peConnection.EndpointServiceName, *currentModel.Region,
			currentModel.newAwsPrivateEndpointInput())
		if progressEvent != nil {
			return addModelToProgressEvent(progressEvent, currentModel), nil
		}

		privateEndpointInput := make([]privateendpoint.AtlasPrivateEndpointInput, len(awsPrivateEndpointOutput))

		for i, awsPe := range awsPrivateEndpointOutput {
			privateEndpointInput[i] = privateendpoint.AtlasPrivateEndpointInput{
				VpcId:               awsPe.VpcId,
				SubnetId:            awsPe.SubnetId,
				InterfaceEndpointId: awsPe.InterfaceEndpointId,
			}
		}

		pe := privateendpoint.Create(mongodbClient, *currentModel.GroupId, privateEndpointInput, peConnection.ID)

		return addModelToProgressEvent(&pe, currentModel), nil
	default:
		ValidationOutput, progressEvent := privateendpoint.ValidateCreationCompletion(mongodbClient, *currentModel.GroupId, req)
		if progressEvent != nil {
			return addModelToProgressEvent(progressEvent, currentModel), nil
		}

		currentModel.Id = &ValidationOutput.ID
		privateEndpoints := make([]PrivateEndpoint, len(ValidationOutput.Endpoints))
		for i, v := range ValidationOutput.Endpoints {
			privateEndpoints[i] = PrivateEndpoint{
				VpcId:               &v.VpcId,
				SubnetId:            &v.SubnetId,
				InterfaceEndpointId: &v.InterfaceEndpointId,
			}
		}

		currentModel.PrivateEndpoints = privateEndpoints
		privateEndpointResponse, response, err := mongodbClient.PrivateEndpoints.Get(context.Background(), *currentModel.GroupId, providerName, *currentModel.Id)
		if err != nil {
			return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
				response.Response), nil
		}
		currentModel.EndpointServiceName = &privateEndpointResponse.EndpointServiceName

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
	setup()
	log.Print("Entered Point 1")
	if errEvent := validator.ValidateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	mongodbClient, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	status, progressStatus := getProcessStatus(req)
	if progressStatus != nil {
		return *progressStatus, nil
	}

	log.Print("Entered Point 2")

	switch status {
	case resource_constats.Init:
		{
			log.Print("Entered Point 3")
			a := prevModel.newAwsPrivateEndpointInput()
			b := currentModel.newAwsPrivateEndpointInput()
			awsPrivateEndpointOutput, progressEvent := awsvpcendpoint.Update(req, *currentModel.EndpointServiceName, *currentModel.Region,
				a,
				b)
			if progressEvent != nil {
				return *progressEvent, nil
			}
			log.Print("Entered Point 15")
			toAdd := make([]privateendpoint.AtlasPrivateEndpointInput, 0)
			toDelete := make([]privateendpoint.AtlasPrivateEndpointInput, 0)

			log.Printf("Completed Aws UPDATE ToAdd:%v", len(awsPrivateEndpointOutput.ToAdd))
			for _, i := range awsPrivateEndpointOutput.ToAdd {
				toAdd = append(toAdd, privateendpoint.AtlasPrivateEndpointInput{
					VpcId:               i.VpcId,
					SubnetId:            i.SubnetId,
					InterfaceEndpointId: i.InterfaceEndpointId,
				})
			}

			log.Printf("Completed Aws UPDATE ToDelete:%v", len(awsPrivateEndpointOutput.ToDelete))
			for _, i := range awsPrivateEndpointOutput.ToDelete {
				toDelete = append(toDelete, privateendpoint.AtlasPrivateEndpointInput{
					VpcId:               i.VpcId,
					SubnetId:            i.SubnetId,
					InterfaceEndpointId: i.InterfaceEndpointId,
				})
			}
			log.Print("Entered Point 16")
			pe := privateendpoint.Update(mongodbClient, *currentModel.GroupId, *prevModel.Id, toAdd, toDelete)

			return addModelToProgressEvent(&pe, currentModel), nil
		}
	case resource_constats.UpdatingPrivateEndpoint:
		{
			log.Print("Entered Point 20")
			validationOutput, progressEvent := privateendpoint.ValidateUpdateCompletion(mongodbClient, *prevModel.GroupId, req, *prevModel.Id)
			if progressEvent != nil {
				return addModelToProgressEvent(progressEvent, currentModel), nil
			}

			for _, v := range validationOutput.Endpoints {
				for i, cm := range currentModel.PrivateEndpoints {
					if v.VpcId == *cm.VpcId && v.SubnetId == *cm.SubnetId {
						currentModel.PrivateEndpoints[i].InterfaceEndpointId = &v.InterfaceEndpointId
					}
				}
			}

			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Update Completed",
				ResourceModel:   currentModel}, nil
		}
	default:
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Unexpected Status",
			ResourceModel:   currentModel}, nil
	}

}

/*func completeModels(client *mongodbatlas.Client, prevModel *Model, currentModel *Model) *handler.ProgressEvent {
	privateEndpointResponse, response, err := client.PrivateEndpoints.Get(context.Background(), *prevModel.GroupId, constants.AWS, *prevModel.Id)
	if err != nil {
		pe := progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response)
		return &pe
	}

	prevModel.EndpointServiceName = &privateEndpointResponse.EndpointServiceName
	currentModel.EndpointServiceName = &privateEndpointResponse.EndpointServiceName

	return nil
}*/

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
		interfaceEndpointIds := make([]string, 0, len(currentModel.PrivateEndpoints))
		for _, i := range currentModel.PrivateEndpoints {
			interfaceEndpointIds = append(interfaceEndpointIds, *i.InterfaceEndpointId)
		}
		epr := privateendpoint.Delete(mongodbClient, *currentModel.GroupId, *currentModel.Id,
			interfaceEndpointIds)
		if epr != nil {
			return *epr, nil
		}

		epr = awsvpcendpoint.Delete(req, interfaceEndpointIds, *currentModel.Region)
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

	mm := make([]interface{}, 0, len(privateEndpointResponse))
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
	return len(m.PrivateEndpoints) != 0
}

func (m *Model) completeByConnection(c mongodbatlas.PrivateEndpointConnection) {
	m.Id = &c.ID
	m.EndpointServiceName = &c.EndpointServiceName
	m.ErrorMessage = &c.ErrorMessage

	endpoints := make([]PrivateEndpoint, 0, len(c.InterfaceEndpoints))

	for i := range c.InterfaceEndpoints {
		endpoints = append(endpoints, PrivateEndpoint{
			InterfaceEndpointId: &c.InterfaceEndpoints[i],
		})
	}

	m.PrivateEndpoints = endpoints
	m.Status = &c.Status
}

func getProcessStatus(req handler.Request) (resource_constats.EventStatus, *handler.ProgressEvent) {
	callback := req.CallbackContext["StateName"]
	if callback == nil {
		return resource_constats.Init, nil
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
