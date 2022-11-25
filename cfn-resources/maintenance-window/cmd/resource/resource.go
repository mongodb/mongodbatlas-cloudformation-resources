package resource

import (
	"context"
	"errors"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.GroupID, constants.PvtKey, constants.PubKey}
var ReadRequiredFields = []string{constants.GroupID, constants.PvtKey, constants.PubKey}
var UpdateRequiredFields = []string{constants.GroupID, constants.PvtKey, constants.PubKey}
var DeleteRequiredFields = []string{constants.GroupID, constants.PvtKey, constants.PubKey}
var ListRequiredFields []string

func setup() {
	util.SetupLogger("mongodb-atlas-maintenance-window")
}

func (m Model) toAtlasModel() mongodbatlas.MaintenanceWindow {
	return mongodbatlas.MaintenanceWindow{
		DayOfWeek:            *m.DayOfWeek,
		HourOfDay:            m.HourOfDay,
		StartASAP:            m.StartASAP,
		AutoDeferOnceEnabled: m.AutoDeferOnceEnabled,
	}
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf("Create - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	atlasModel := currentModel.toAtlasModel()
	startASP := false
	atlasModel.StartASAP = &startASP

	res, err = client.MaintenanceWindows.Update(context.Background(), *currentModel.GroupId, &atlasModel)

	if err != nil {
		_, _ = logger.Warnf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   *currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	// Validation
	if errEvent := validator.ValidateModel(ReadRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf("Read - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	maintenanceWindow, errorProgressEvent := get(client, *currentModel)
	if errorProgressEvent != nil {
		return *errorProgressEvent, nil
	}

	currentModel.AutoDeferOnceEnabled = maintenanceWindow.AutoDeferOnceEnabled
	currentModel.DayOfWeek = &maintenanceWindow.DayOfWeek
	currentModel.HourOfDay = maintenanceWindow.HourOfDay
	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}

func get(client *mongodbatlas.Client, currentModel Model) (*mongodbatlas.MaintenanceWindow, *handler.ProgressEvent) {
	maintenanceWindow, res, err := client.MaintenanceWindows.Get(context.Background(), *currentModel.GroupId)
	if err != nil {
		_, _ = logger.Warnf("Read - error: %+v", err)
		ev := progress_events.GetFailedEventByResponse(err.Error(), res.Response)
		return nil, &ev
	}

	if isResponseEmpty(maintenanceWindow) {
		_, _ = logger.Warnf("Read - resource is empty: %+v", err)
		ev := progress_events.GetFailedEventByCode("resource not found", cloudformation.HandlerErrorCodeNotFound)
		return nil, &ev
	}

	return maintenanceWindow, nil
}

func isResponseEmpty(maintenanceWindow *mongodbatlas.MaintenanceWindow) bool {
	return (maintenanceWindow != nil) && (maintenanceWindow != nil && maintenanceWindow.DayOfWeek == 0)
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	// Validation
	if errEvent := validator.ValidateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf("Update - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	_, handlerError := get(client, *currentModel)
	if handlerError != nil {
		return *handlerError, nil
	}

	var res *mongodbatlas.Response

	atlasModel := currentModel.toAtlasModel()
	startASP := false
	atlasModel.StartASAP = &startASP

	res, err = client.MaintenanceWindows.Update(context.Background(), *currentModel.GroupId, &atlasModel)

	if err != nil {
		_, _ = logger.Warnf("Update - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	// Validation
	if errEvent := validator.ValidateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf("Delete - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	_, handlerError := get(client, *currentModel)
	if handlerError != nil {
		return *handlerError, nil
	}

	var res *mongodbatlas.Response
	res, err = client.MaintenanceWindows.Reset(context.Background(), *currentModel.GroupId)

	if err != nil {
		_, _ = logger.Warnf("Delete - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "delete successful",
	}
	return event, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("Not implemented: List")
}
