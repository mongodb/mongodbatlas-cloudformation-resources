// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"context"
	"errors"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

var RequiredFields = []string{constants.GroupID}

func setup() {
	util.SetupLogger("mongodb-atlas-auditing")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validator.ValidateModel(RequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("CRATE Validation Error")
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
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
	modelValidation := validator.ValidateModel(RequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("READ Validation Error")
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
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
	currentModel.AuditFilter = &atlasAuditing.AuditFilter
	currentModel.AuditAuthorizationSuccess = atlasAuditing.AuditAuthorizationSuccess

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
	modelValidation := validator.ValidateModel(RequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("UPDATE Validation Error")
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
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

	modelValidation := validator.ValidateModel(RequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("DELETE Validation Error")
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
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

	_, res, err := client.Auditing.Configure(context.Background(), *currentModel.GroupId, &auditingInput)

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
