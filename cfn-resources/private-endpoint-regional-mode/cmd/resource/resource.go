package resource

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.GroupID, constants.PubKey, constants.PvtKey}
var ReadRequiredFields = []string{constants.GroupID, constants.PubKey, constants.PvtKey}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.GroupID, constants.PubKey, constants.PvtKey}
var ListRequiredFields = []string{constants.GroupID, constants.PubKey, constants.PvtKey}

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

	mongodbClient, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// API call to Create cloud backup schedule
	return resourcePrivateEndpointRegionalModeUpdate(req, prevModel, currentModel, mongodbClient)
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
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

	rationalizedPrivateEndpointSetting, resp, err := mongodbClient.PrivateEndpoints.GetRegionalizedPrivateEndpointSetting(context.Background(), *currentModel.GroupId)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			_, _ = logger.Warnf("update 404 err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = logger.Warnf("update err: %+v", err)
		code := cloudformation.HandlerErrorCodeServiceInternalError
		if strings.Contains(err.Error(), "not exist") { // cfn test needs 404
			code = cloudformation.HandlerErrorCodeNotFound
		}
		if strings.Contains(err.Error(), "being deleted") {
			code = cloudformation.HandlerErrorCodeNotFound // cfn test needs 404
		}
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: code}, nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "PATCH Complete",
		ResourceModel:   regionalPrivateEndpointToModel(*currentModel, rationalizedPrivateEndpointSetting),
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
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

	// API call to Create cloud backup schedule
	return resourcePrivateEndpointRegionalModeUpdate(req, prevModel, currentModel, mongodbClient)
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("Not implemented: Delete")
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("Not implemented: List")
}

func resourcePrivateEndpointRegionalModeUpdate(req handler.Request, prevModel *Model, currentModel *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {

	if (currentModel.Enabled) == nil {
		err := errors.New("error updating Private Endpoint Regional Mode: Enabled should be set when `Export` is set")
		_, _ = logger.Warnf("Update - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	regionalizedPrivateEndpointSetting, resp, err := client.PrivateEndpoints.UpdateRegionalizedPrivateEndpointSetting(context.Background(), *currentModel.GroupId, *currentModel.Enabled)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			_, _ = logger.Warnf("update 404 err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = logger.Warnf("update err: %+v", err)
		code := cloudformation.HandlerErrorCodeServiceInternalError
		if strings.Contains(err.Error(), "not exist") { // cfn test needs 404
			code = cloudformation.HandlerErrorCodeNotFound
		}
		if strings.Contains(err.Error(), "being deleted") {
			code = cloudformation.HandlerErrorCodeNotFound // cfn test needs 404
		}
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: code}, nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "PATCH Complete",
		ResourceModel:   regionalPrivateEndpointToModel(*currentModel, regionalizedPrivateEndpointSetting),
	}, nil
}

func regionalPrivateEndpointToModel(currentModel Model, regPrivateMode *mongodbatlas.RegionalizedPrivateEndpointSetting) *Model {
	out := &Model{
		ApiKeys: currentModel.ApiKeys,
		Enabled: &regPrivateMode.Enabled,
	}
	return out
}
