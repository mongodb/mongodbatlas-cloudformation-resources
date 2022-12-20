package resource

import (
	"context"
	"errors"

	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.GroupID, constants.PvtKey, constants.PubKey}
var ReadRequiredFields = []string{constants.GroupID, constants.PvtKey, constants.PubKey}
var UpdateRequiredFields = []string{constants.GroupID, constants.PvtKey, constants.PubKey}
var DeleteRequiredFields = []string{constants.GroupID, constants.PvtKey, constants.PubKey}
var ListRequiredFields []string

func setup() {
	util.SetupLogger("mongodb-atlas-auditing")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("CRATE Validation Error")
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	atlasAuditing, res, err := client.Auditing.Get(context.Background(), *currentModel.GroupId)
	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	if *atlasAuditing.Enabled {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists,
			OperationStatus:  handler.Failed,
		}, nil
	}

	enabled := true

	auditingInput := mongodbatlas.Auditing{
		Enabled: &enabled,
	}

	if currentModel.AuditAuthorizationSuccess != nil {
		auditingInput.AuditAuthorizationSuccess = currentModel.AuditAuthorizationSuccess
	}

	if currentModel.AuditFilter != nil {
		auditingInput.AuditFilter = *currentModel.AuditFilter
	}

	atlasAuditing, res, err = client.Auditing.Configure(context.Background(), *currentModel.GroupId, &auditingInput)

	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.ConfigurationType = &atlasAuditing.ConfigurationType

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("READ Validation Error")
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Read - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	atlasAuditing, res, err := client.Auditing.Get(context.Background(), *currentModel.GroupId)
	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	if !*atlasAuditing.Enabled {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
			OperationStatus:  handler.Failed,
		}, nil
	}

	currentModel.ConfigurationType = &atlasAuditing.ConfigurationType

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "get successful",
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("UPDATE Validation Error")
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Update - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	resourceEnabled, handlerEvent := isEnabled(*client, *currentModel)
	if handlerEvent != nil {
		return *handlerEvent, nil
	}
	if !resourceEnabled {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
			OperationStatus:  handler.Failed,
			Message:          "resource not found",
		}, nil
	}

	var res *mongodbatlas.Response

	auditingInput := mongodbatlas.Auditing{}

	modified := false

	if currentModel.AuditAuthorizationSuccess != nil {
		modified = true
		auditingInput.AuditAuthorizationSuccess = currentModel.AuditAuthorizationSuccess
	}

	if currentModel.AuditFilter != nil {
		modified = true
		auditingInput.AuditFilter = *currentModel.AuditFilter
	}

	if !modified {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Update success (no properties were changed)",
			ResourceModel:   currentModel,
		}, nil
	}

	atlasAuditing, res, err := client.Auditing.Configure(context.Background(), *currentModel.GroupId, &auditingInput)

	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	if err != nil {
		_, _ = log.Debugf("Update - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	_, _ = log.Debugf("Atlas Client %v", client)

	currentModel.ConfigurationType = &atlasAuditing.ConfigurationType

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update success",
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("DELETE Validation Error")
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Delete - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	resourceEnabled, handlerEvent := isEnabled(*client, *currentModel)
	if handlerEvent != nil {
		return *handlerEvent, nil
	}

	if !resourceEnabled {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
			OperationStatus:  handler.Failed,
		}, nil
	}

	var res *mongodbatlas.Response

	enabled := false

	auditingInput := mongodbatlas.Auditing{
		Enabled: &enabled,
	}

	_, res, err = client.Auditing.Configure(context.Background(), *currentModel.GroupId, &auditingInput)

	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	if err != nil {
		_, _ = log.Debugf("Delete - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	_, _ = log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
	}, nil
}

func isEnabled(client mongodbatlas.Client, currentModel Model) (bool, *handler.ProgressEvent) {
	atlasAuditing, res, err := client.Auditing.Get(context.Background(), *currentModel.GroupId)

	if err != nil {
		_, _ = log.Debugf("Validating enabled - error: %+v", err)
		er := progress_events.GetFailedEventByResponse(err.Error(), res.Response)
		return false, &er
	}

	return *atlasAuditing.Enabled, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}
