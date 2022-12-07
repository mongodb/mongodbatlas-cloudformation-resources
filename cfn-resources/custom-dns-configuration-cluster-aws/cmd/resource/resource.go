package resource

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.PubKey, constants.PvtKey}
var ReadRequiredFields = []string{constants.ProjectID, constants.PubKey, constants.PvtKey}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.ProjectID, constants.PubKey, constants.PvtKey}
var ListRequiredFields = []string{constants.ProjectID, constants.PubKey, constants.PvtKey}

func setup() {
	util.SetupLogger("mongodb-atlas-custom-dns-configuration-cluster-aws")
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
	if isCustomAWSDNSSettingExists(currentModel, mongodbClient) {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Custom AWS dns settings already enabled for : %s", *currentModel.ProjectId),
			cloudformation.HandlerErrorCodeAlreadyExists), nil
	}
	// API call to
	enabled := true
	currentModel.Enabled = &enabled
	return resourceCustomAWSDNSUpdate(req, prevModel, currentModel, mongodbClient)
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
	customAWSDNSSetting, response, err := mongodbClient.CustomAWSDNS.Get(context.Background(), *currentModel.ProjectId)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error reading  : %s", err.Error()),
			response.Response), nil
	}
	enabled := customAWSDNSSetting.Enabled
	if !enabled {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Custom AWS dns settings not found for Project : %s", *currentModel.ProjectId),
			cloudformation.HandlerErrorCodeNotFound), nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "READ Complete",
		ResourceModel:   customAWSDNSToModel(*currentModel, customAWSDNSSetting),
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
	if isCustomAWSDNSSettingExists(currentModel, mongodbClient) {
		enabled := false
		currentModel.Enabled = &enabled
		events, err := resourceCustomAWSDNSUpdate(req, prevModel, currentModel, mongodbClient)
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
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error in disabling Custom AWS DNS settings for Project : %s", *currentModel.ProjectId),
			cloudformation.HandlerErrorCodeNotFound), nil
	}
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("Not implemented: List")
}

func resourceCustomAWSDNSUpdate(req handler.Request, prevModel *Model, currentModel *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {

	customAWSDNSRequest := &mongodbatlas.AWSCustomDNSSetting{
		Enabled: *currentModel.Enabled,
	}
	customAWSDNSModel, response, err := client.CustomAWSDNS.Update(context.Background(), *currentModel.ProjectId, customAWSDNSRequest)
	if err != nil {
		return progress_events.GetFailedEventByResponse(
			fmt.Sprintf("Error in enabling Custom AWS DNS settings : %s", err.Error()),
			response.Response), nil
	}
	currentModel.Enabled = &customAWSDNSModel.Enabled
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   customAWSDNSToModel(*currentModel, customAWSDNSModel),
	}, nil
}

func isCustomAWSDNSSettingExists(currentModel *Model, client *mongodbatlas.Client) bool {
	var isExists bool
	customAWSDNSSetting, _, err := client.CustomAWSDNS.Get(context.Background(), *currentModel.ProjectId)
	if err != nil {
		return isExists
	}
	if customAWSDNSSetting.Enabled {
		isExists = true
	}
	return isExists
}

func customAWSDNSToModel(currentModel Model, regPrivateMode *mongodbatlas.AWSCustomDNSSetting) *Model {
	out := &Model{
		ApiKeys:   currentModel.ApiKeys,
		Enabled:   &regPrivateMode.Enabled,
		ProjectId: currentModel.ProjectId,
	}
	return out
}
