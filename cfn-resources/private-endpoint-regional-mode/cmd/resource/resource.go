package resource

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
	"log"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.PubKey, constants.PvtKey}
var ReadRequiredFields = []string{constants.ProjectID, constants.PubKey, constants.PvtKey}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.ProjectID, constants.PubKey, constants.PvtKey}
var ListRequiredFields = []string{constants.ProjectID, constants.PubKey, constants.PvtKey}

func setup() {
	util.SetupLogger("mongodb-atlas-private-endpoint-regional-mode")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}
	enabled := *currentModel.Enabled
	if !enabled {
		return handler.ProgressEvent{
			Message: "Value of 'enabled' can only be true during CREATE. In order to set 'enabled' as false," +
				"execute DELETE operation",
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	mongodbClient, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	// TODO: isRegModeExists()
	rationalizedPrivateEndpointSetting, response, err := mongodbClient.PrivateEndpoints.GetRegionalizedPrivateEndpointSetting(context.Background(), *currentModel.ProjectId)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error while reading regionalized mode for private endpoint : %s", err.Error()),
			response.Response), nil
	}
	enabled = rationalizedPrivateEndpointSetting.Enabled
	if enabled {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Private Endpoint Regionalized Mode already enabled for : %s", *currentModel.ProjectId),
			cloudformation.HandlerErrorCodeAlreadyExists), nil
	}
	// API call to
	return resourcePrivateEndpointRegionalModeUpdate(req, prevModel, currentModel, mongodbClient)
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
	// TODO: isRegModeExists()
	regionalizedPrivateEndpointSetting, response, err := mongodbClient.PrivateEndpoints.GetRegionalizedPrivateEndpointSetting(context.Background(), *currentModel.ProjectId)
	enabled := regionalizedPrivateEndpointSetting.Enabled
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error reading  : %s", err.Error()),
			response.Response), nil
	}

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			_, _ = logger.Warnf("error 404- err:%+v resp:%+v", err, response)
			return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
				cloudformation.HandlerErrorCodeNotFound), nil
		}
		_, _ = logger.Warnf("error cloud backup policy get- err:%+v resp:%+v", err, response)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	if !enabled {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Regionalized mode for private endpoint not found for Project : %s", *currentModel.ProjectId),
			cloudformation.HandlerErrorCodeNotFound), nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "READ Complete",
		ResourceModel:   regionalPrivateEndpointToModel(*currentModel, regionalizedPrivateEndpointSetting),
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("Not implemented: Update")
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
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	// TODO: isRegModeExists()
	regionalizedPrivateEndpointSetting, response, err := mongodbClient.PrivateEndpoints.GetRegionalizedPrivateEndpointSetting(context.Background(), *currentModel.ProjectId)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error while reading regionalized mode for private endpoint : %s", err.Error()),
			response.Response), nil
	}
	enabled := regionalizedPrivateEndpointSetting.Enabled
	if enabled {
		log.Printf("currentModel.Enabled flag is: %v", enabled)
		enabled := false
		currentModel.Enabled = &enabled
		events, err := resourcePrivateEndpointRegionalModeUpdate(req, prevModel, currentModel, mongodbClient)
		if err != nil {
			return progress_events.GetFailedEventByCode(fmt.Sprintf("Error in disabling regionalized mode for private endpoint for Project : %s", *currentModel.ProjectId),
				events.HandlerErrorCode), nil
		}
		// Response
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Delete Complete",
		}, nil
	} else {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error in disabling regionalized mode for private endpoint for Project : %s", *currentModel.ProjectId),
			cloudformation.HandlerErrorCodeNotFound), nil
	}
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("Not implemented: List")
}

func resourcePrivateEndpointRegionalModeUpdate(req handler.Request, prevModel *Model, currentModel *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {

	if (currentModel.Enabled) == nil {
		err := errors.New("error updating Private Endpoint Regional Mode: Enabled should be set in request")
		_, _ = logger.Warnf("Update - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	spew.Dump(currentModel)

	regionalizedPrivateEndpointSetting, response, err := client.PrivateEndpoints.UpdateRegionalizedPrivateEndpointSetting(context.Background(), *currentModel.ProjectId, *currentModel.Enabled)
	if err != nil {
		return progress_events.GetFailedEventByResponse(
			fmt.Sprintf("Error in enabling regionalized settings : %s", err.Error()),
			response.Response), nil
	}
	currentModel.Enabled = &regionalizedPrivateEndpointSetting.Enabled
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   regionalPrivateEndpointToModel(*currentModel, regionalizedPrivateEndpointSetting),
	}, nil
}

func regionalPrivateEndpointToModel(currentModel Model, regPrivateMode *mongodbatlas.RegionalizedPrivateEndpointSetting) *Model {
	out := &Model{
		ApiKeys:   currentModel.ApiKeys,
		Enabled:   &regPrivateMode.Enabled,
		ProjectId: currentModel.ProjectId,
	}
	return out
}
