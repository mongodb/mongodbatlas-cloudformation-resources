// Copyright 2024 MongoDB Inc
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
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

const (
	callBackSeconds                    = 40
	searchDeploymentDoesNotExistsError = "ATLAS_FTS_DEPLOYMENT_DOES_NOT_EXIST"
	searchDeploymentAlreadyExistsError = "ATLAS_FTS_DEPLOYMENT_ALREADY_EXISTS"
)

var createRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.Specs}
var readRequiredFields = []string{constants.ProjectID, constants.ClusterName}
var updateRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.Specs}
var deleteRequiredFields = []string{constants.ProjectID, constants.ClusterName}

func setup() {
	util.SetupLogger("mongodb-atlas-searchdeployment")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if modelValidation := validator.ValidateModel(createRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	client, progressErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if progressErr != nil {
		return *progressErr, nil
	}
	connV2 := client.AtlasSDK

	// handling of subsequent retry calls
	if _, ok := req.CallbackContext[constants.ID]; ok {
		return handleStateTransition(*connV2, currentModel, constants.IdleState), nil
	}

	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	apiReq := newSearchDeploymentReq(currentModel)
	apiResp, resp, err := connV2.AtlasSearchApi.CreateAtlasSearchDeployment(context.Background(), projectID, clusterName, &apiReq).Execute()
	if err != nil {
		return handleError(resp, err)
	}

	newModel := newCFNSearchDeployment(currentModel, apiResp)
	return inProgressEvent("Creating Search Deployment", &newModel), nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if modelValidation := validator.ValidateModel(readRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	client, progressErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if progressErr != nil {
		return *progressErr, nil
	}
	connV2 := client.AtlasSDK

	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	apiResp, resp, err := connV2.AtlasSearchApi.GetAtlasSearchDeployment(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		return handleError(resp, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   newCFNSearchDeployment(currentModel, apiResp),
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if modelValidation := validator.ValidateModel(updateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	client, progressErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if progressErr != nil {
		return *progressErr, nil
	}
	connV2 := client.AtlasSDK

	// handling of subsequent retry calls
	if _, ok := req.CallbackContext[constants.ID]; ok {
		return handleStateTransition(*connV2, currentModel, constants.IdleState), nil
	}

	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	apiReq := newSearchDeploymentReq(currentModel)
	apiResp, res, err := connV2.AtlasSearchApi.UpdateAtlasSearchDeployment(context.Background(), projectID, clusterName, &apiReq).Execute()
	if err != nil {
		return handleError(res, err)
	}

	newModel := newCFNSearchDeployment(currentModel, apiResp)
	return inProgressEvent("Updating Search Deployment", &newModel), nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if modelValidation := validator.ValidateModel(deleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	client, progressErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if progressErr != nil {
		return *progressErr, nil
	}
	connV2 := client.AtlasSDK

	// handling of subsequent retry calls
	if _, ok := req.CallbackContext[constants.ID]; ok {
		return handleStateTransition(*connV2, currentModel, constants.DeletedState), nil
	}

	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	if resp, err := connV2.AtlasSearchApi.DeleteAtlasSearchDeployment(context.Background(), projectID, clusterName).Execute(); err != nil {
		return handleError(resp, err)
	}

	return inProgressEvent(constants.DeleteInProgress, currentModel), nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func handleStateTransition(connV2 admin.APIClient, currentModel *Model, targetState string) handler.ProgressEvent {
	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	apiResp, resp, err := connV2.AtlasSearchApi.GetAtlasSearchDeployment(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if targetState == constants.DeletedState && resp.StatusCode == http.StatusBadRequest && strings.Contains(err.Error(), searchDeploymentDoesNotExistsError) {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				ResourceModel:   nil,
				Message:         constants.Complete,
			}
		}
		return progressevent.GetFailedEventByResponse(err.Error(), resp)
	}

	newModel := newCFNSearchDeployment(currentModel, apiResp)
	if util.SafeString(newModel.StateName) == targetState {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			ResourceModel:   newModel,
			Message:         constants.Complete,
		}
	}

	return inProgressEvent(constants.Pending, &newModel)
}

// specific handling for search deployment API where 400 status code can include AlreadyExists or DoesNotExist that need specific mapping to CFN error codes
func handleError(res *http.Response, err error) (handler.ProgressEvent, error) {
	if apiError, ok := admin.AsError(err); ok && *apiError.Error == http.StatusBadRequest && strings.Contains(*apiError.ErrorCode, searchDeploymentAlreadyExistsError) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}
	if apiError, ok := admin.AsError(err); ok && *apiError.Error == http.StatusBadRequest && strings.Contains(*apiError.ErrorCode, searchDeploymentDoesNotExistsError) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	return progressevent.GetFailedEventByResponse(err.Error(), res), nil
}

func inProgressEvent(message string, model *Model) handler.ProgressEvent {
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              message,
		ResourceModel:        model,
		CallbackDelaySeconds: callBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.ID: model.Id,
		}}
}
