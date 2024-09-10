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
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115014/admin"
)

var RequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-auditing")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	modelValidation := validator.ValidateModel(RequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("CREATE Validation Error")
		return *modelValidation, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.Atlas20231115014
	var res *http.Response

	atlasAuditing, res, err := atlasV2.AuditingApi.GetAuditingConfiguration(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	if aws.BoolValue(atlasAuditing.Enabled) {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists,
			OperationStatus:  handler.Failed,
		}, nil
	}

	enabled := true

	auditingInput := admin.AuditLog{
		Enabled: &enabled,
	}

	if currentModel.AuditAuthorizationSuccess != nil {
		auditingInput.AuditAuthorizationSuccess = currentModel.AuditAuthorizationSuccess
	}

	if currentModel.AuditFilter != nil {
		auditingInput.AuditFilter = currentModel.AuditFilter
	}

	atlasAuditing, res, err = atlasV2.AuditingApi.UpdateAuditingConfiguration(context.Background(), *currentModel.ProjectId, &auditingInput).Execute()

	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	currentModel.ConfigurationType = atlasAuditing.ConfigurationType

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	modelValidation := validator.ValidateModel(RequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("READ Validation Error")
		return *modelValidation, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.Atlas20231115014
	var res *http.Response

	atlasAuditing, res, err := atlasV2.AuditingApi.GetAuditingConfiguration(context.Background(), *currentModel.ProjectId).Execute()

	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	if !aws.BoolValue(atlasAuditing.Enabled) {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
			OperationStatus:  handler.Failed,
		}, nil
	}

	currentModel.ConfigurationType = atlasAuditing.ConfigurationType
	currentModel.AuditFilter = atlasAuditing.AuditFilter
	currentModel.AuditAuthorizationSuccess = atlasAuditing.AuditAuthorizationSuccess

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "get successful",
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	modelValidation := validator.ValidateModel(RequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("UPDATE Validation Error")
		return *modelValidation, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.Atlas20231115014

	resourceEnabled, handlerEvent := isEnabled(*atlasV2, *currentModel)
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

	var res *http.Response
	auditingInput := admin.AuditLog{}
	modified := false

	if currentModel.AuditAuthorizationSuccess != nil {
		modified = true
		auditingInput.AuditAuthorizationSuccess = currentModel.AuditAuthorizationSuccess
	}

	if currentModel.AuditFilter != nil {
		modified = true
		auditingInput.AuditFilter = currentModel.AuditFilter
	}

	if !modified {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Update success (no properties were changed)",
			ResourceModel:   currentModel,
		}, nil
	}

	atlasAuditing, res, err := atlasV2.AuditingApi.UpdateAuditingConfiguration(context.Background(), *currentModel.ProjectId, &auditingInput).Execute()

	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	currentModel.ConfigurationType = atlasAuditing.ConfigurationType

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update success",
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	modelValidation := validator.ValidateModel(RequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("DELETE Validation Error")
		return *modelValidation, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.Atlas20231115014

	resourceEnabled, handlerEvent := isEnabled(*atlasV2, *currentModel)
	if handlerEvent != nil {
		return *handlerEvent, nil
	}

	if !resourceEnabled {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
			OperationStatus:  handler.Failed,
		}, nil
	}

	var res *http.Response
	enabled := false

	auditingInput := admin.AuditLog{
		Enabled:     &enabled,
		AuditFilter: currentModel.AuditFilter,
	}

	_, res, err := atlasV2.AuditingApi.UpdateAuditingConfiguration(context.Background(), *currentModel.ProjectId, &auditingInput).Execute()

	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
	}, nil
}

func isEnabled(client admin.APIClient, currentModel Model) (bool, *handler.ProgressEvent) {
	atlasAuditing, res, err := client.AuditingApi.GetAuditingConfiguration(context.Background(), *currentModel.ProjectId).Execute()

	if err != nil {
		er := progressevent.GetFailedEventByResponse(err.Error(), res)
		return false, &er
	}

	return aws.BoolValue(atlasAuditing.Enabled), nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}
