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
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	atlasSDK "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

var RequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-auditing")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("In CREATE handler...............")

	// Validation
	modelValidation := validator.ValidateModel(RequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("CREATE Validation Error")
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.AtlasV2
	var res *http.Response

	atlasAuditing, res, err := atlasV2.AuditingApi.GetAuditingConfiguration(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res), nil
	}

	if atlasAuditing.Enabled {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists,
			OperationStatus:  handler.Failed,
		}, nil
	}

	enabled := true

	auditingInput := atlasSDK.AuditLog{
		Enabled: enabled,
	}

	if currentModel.AuditAuthorizationSuccess != nil {
		auditingInput.AuditAuthorizationSuccess = *currentModel.AuditAuthorizationSuccess
	}

	if currentModel.AuditFilter != nil {
		auditingInput.AuditFilter = *currentModel.AuditFilter
	}

	atlasAuditing, res, err = atlasV2.AuditingApi.UpdateAuditingConfiguration(context.Background(), *currentModel.ProjectId, &auditingInput).Execute()

	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res), nil
	}

	currentModel.ConfigurationType = atlasAuditing.ConfigurationType

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("In READ handler...............")
	// Validation
	modelValidation := validator.ValidateModel(RequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("READ Validation Error")
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.AtlasV2
	var res *http.Response

	atlasAuditing, res, err := atlasV2.AuditingApi.GetAuditingConfiguration(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res), nil
	}

	if !atlasAuditing.Enabled {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
			OperationStatus:  handler.Failed,
		}, nil
	}

	currentModel.ConfigurationType = atlasAuditing.ConfigurationType
	currentModel.AuditFilter = &atlasAuditing.AuditFilter
	currentModel.AuditAuthorizationSuccess = &atlasAuditing.AuditAuthorizationSuccess

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "get successful",
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("\n \n In UPDATE handler...............")
	// Validation
	modelValidation := validator.ValidateModel(RequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("UPDATE Validation Error")
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.AtlasV2

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

	cm, _ := json.Marshal(currentModel)
	log.Debugf("\n \n Update current model " + string(cm))

	var res *http.Response

	auditingInput := atlasSDK.AuditLog{}

	modified := false

	if currentModel.AuditAuthorizationSuccess != nil {
		modified = true
		auditingInput.AuditAuthorizationSuccess = *currentModel.AuditAuthorizationSuccess
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

	log.Debugf("\n \n In UPDATE handler...............calling API")
	atlasAuditing, res, err := atlasV2.AuditingApi.UpdateAuditingConfiguration(context.Background(), *currentModel.ProjectId, &auditingInput).Execute()
	res2B, _ := json.Marshal(auditingInput)
	log.Debugf("\n \n Update request " + string(res2B))
	res3B, _ := json.Marshal(atlasAuditing)
	log.Debugf("\n \n Update response " + string(res3B))

	if err != nil {
		_, _ = log.Debugf("Update - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res), nil
	}

	currentModel.ConfigurationType = atlasAuditing.ConfigurationType

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update success",
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("\n \n In DELETE handler..............")
	modelValidation := validator.ValidateModel(RequiredFields, currentModel)
	if modelValidation != nil {
		_, _ = log.Debugf("DELETE Validation Error")
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.AtlasV2

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

	cm, _ := json.Marshal(currentModel)
	log.Debugf("\n \n Delete current model " + string(cm))

	var res *http.Response

	auditingInput := atlasSDK.AuditLog{
		Enabled:     false,
		AuditFilter: *currentModel.AuditFilter,
	}

	atlasAuditing, res, err := atlasV2.AuditingApi.UpdateAuditingConfiguration(context.Background(), *currentModel.ProjectId, &auditingInput).Execute()
	res2B, _ := json.Marshal(auditingInput)
	log.Debugf("\n \n Delete request " + string(res2B))
	res3B, _ := json.Marshal(atlasAuditing)
	log.Debugf("\n \n Delete response " + string(res3B))

	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res), nil
	}

	if err != nil {
		_, _ = log.Debugf("Delete - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res), nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
	}, nil
}

func isEnabled(client atlasSDK.APIClient, currentModel Model) (bool, *handler.ProgressEvent) {
	atlasAuditing, res, err := client.AuditingApi.GetAuditingConfiguration(context.Background(), *currentModel.ProjectId).Execute()

	if err != nil {
		_, _ = log.Debugf("Validating enabled - error: %+v", err)
		er := progress_events.GetFailedEventByResponse(err.Error(), res)
		return false, &er
	}

	return atlasAuditing.Enabled, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}
