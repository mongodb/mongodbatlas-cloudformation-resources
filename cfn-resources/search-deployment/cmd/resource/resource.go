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

	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

const (
	CallBackSeconds                       = 40
	SearchDeploymentAlreadyExistsErrorAPI = "ATLAS_SEARCH_DEPLOYMENT_ALREADY_EXISTS"
	SearchDeploymentDoesNotExistsErrorAPI = "ATLAS_SEARCH_DEPLOYMENT_DOES_NOT_EXIST"
)

var callbackContext = map[string]any{"callbackSearchDeployment": true}

func IsCallback(req *handler.Request) bool {
	_, found := req.CallbackContext["callbackSearchDeployment"]
	return found
}

var (
	CreateRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.Specs}
	ReadRequiredFields   = []string{constants.ProjectID, constants.ClusterName}
	UpdateRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.Specs}
	DeleteRequiredFields = []string{constants.ProjectID, constants.ClusterName}
)

func setup() {
	util.SetupLogger("mongodb-atlas-searchdeployment")
}

var InitEnvWithClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if modelValidation := validator.ValidateModel(requiredFields, currentModel); modelValidation != nil {
		return nil, modelValidation
	}

	client, progressErr := util.NewAtlasClient(&req, currentModel.Profile)
	if progressErr != nil {
		return nil, progressErr
	}
	return client.AtlasSDK, nil
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	connV2, peErr := InitEnvWithClient(req, currentModel, CreateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	if IsCallback(&req) {
		return ValidateProgress(*connV2, currentModel, false), nil
	}

	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	apiReq := NewSearchDeploymentReq(currentModel)

	createResp, resp, err := connV2.AtlasSearchApi.CreateClusterSearchDeployment(context.Background(), projectID, clusterName, &apiReq).Execute()
	if err != nil {
		notFound := resp != nil && resp.StatusCode == http.StatusNotFound
		alreadyExists := resp != nil && resp.StatusCode == http.StatusConflict &&
			strings.Contains(err.Error(), SearchDeploymentAlreadyExistsErrorAPI)

		if alreadyExists || notFound {
			existingResp, _, getErr := connV2.AtlasSearchApi.GetClusterSearchDeployment(context.Background(), projectID, clusterName).Execute()
			if getErr != nil {
				return HandleError(resp, err)
			}
			existingModel := NewCFNSearchDeployment(currentModel, existingResp)
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				ResourceModel:   &existingModel,
				Message:         constants.Complete,
			}, nil
		}
		return HandleError(resp, err)
	}

	var apiResp *admin20250312010.ApiSearchDeploymentResponse
	if createResp != nil && createResp.Id != nil {
		apiResp = createResp
	} else {
		apiResp, resp, err = connV2.AtlasSearchApi.GetClusterSearchDeployment(context.Background(), projectID, clusterName).Execute()
		if err != nil {
			return HandleError(resp, err)
		}
		if apiResp == nil || apiResp.Id == nil {
			return handler.ProgressEvent{
				OperationStatus:      handler.InProgress,
				Message:              "Creating Search Deployment - waiting for deployment ID",
				ResourceModel:        currentModel,
				CallbackDelaySeconds: CallBackSeconds,
				CallbackContext:      callbackContext,
			}, nil
		}
	}

	newModel := NewCFNSearchDeployment(currentModel, apiResp)
	if newModel.Id == nil {
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              "Creating Search Deployment - waiting for deployment ID",
			ResourceModel:        currentModel,
			CallbackDelaySeconds: CallBackSeconds,
			CallbackContext:      callbackContext,
		}, nil
	}

	stateName := util.SafeString(newModel.StateName)
	if stateName == constants.IdleState {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			ResourceModel:   &newModel,
			Message:         constants.Complete,
		}, nil
	}

	return inProgressEvent("Creating Search Deployment", currentModel, apiResp), nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	connV2, peErr := InitEnvWithClient(req, currentModel, ReadRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	apiResp, resp, err := connV2.AtlasSearchApi.GetClusterSearchDeployment(context.Background(), projectID, clusterName).Execute()

	if err != nil {
		return HandleError(resp, err)
	}

	if apiResp == nil || apiResp.Id == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Search deployment not found",
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
		}, nil
	}

	newModel := NewCFNSearchDeployment(currentModel, apiResp)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   &newModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	connV2, peErr := InitEnvWithClient(req, currentModel, UpdateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	if IsCallback(&req) {
		return ValidateProgress(*connV2, currentModel, false), nil
	}

	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)

	// Check if resource exists before updating (required by contract tests)
	checkResp, checkHTTPResp, err := connV2.AtlasSearchApi.GetClusterSearchDeployment(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		// If resource doesn't exist, return NotFound (required by contract tests)
		if checkHTTPResp != nil && checkHTTPResp.StatusCode == http.StatusNotFound {
			return progressevent.GetFailedEventByResponse("Search deployment not found", checkHTTPResp), nil
		}
		return HandleError(checkHTTPResp, err)
	}
	if checkResp == nil || checkResp.Id == nil {
		// Resource doesn't exist - return NotFound with proper HTTP response
		notFoundResp := &http.Response{StatusCode: http.StatusNotFound}
		return progressevent.GetFailedEventByResponse("Search deployment not found", notFoundResp), nil
	}

	apiReq := NewSearchDeploymentReq(currentModel)
	_, res, err := connV2.AtlasSearchApi.UpdateClusterSearchDeployment(context.Background(), projectID, clusterName, &apiReq).Execute()
	if err != nil {
		return HandleError(res, err)
	}

	apiResp, resp, err := connV2.AtlasSearchApi.GetClusterSearchDeployment(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		modelWithID := GetModelWithID(currentModel, prevModel, checkResp)
		if modelWithID != nil {
			return handler.ProgressEvent{
				OperationStatus:      handler.InProgress,
				Message:              "Updating Search Deployment",
				ResourceModel:        modelWithID,
				CallbackDelaySeconds: CallBackSeconds,
				CallbackContext:      callbackContext,
			}, nil
		}
		return HandleError(resp, err)
	}

	if apiResp == nil || apiResp.Id == nil {
		modelWithID := GetModelWithID(currentModel, prevModel, checkResp)
		if modelWithID != nil {
			return handler.ProgressEvent{
				OperationStatus:      handler.InProgress,
				Message:              "Updating Search Deployment",
				ResourceModel:        modelWithID,
				CallbackDelaySeconds: CallBackSeconds,
				CallbackContext:      callbackContext,
			}, nil
		}
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              "Updating Search Deployment - waiting for deployment ID",
			ResourceModel:        currentModel,
			CallbackDelaySeconds: CallBackSeconds,
			CallbackContext:      callbackContext,
		}, nil
	}

	return inProgressEvent("Updating Search Deployment", currentModel, apiResp), nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if currentModel == nil || (currentModel.ProjectId == nil && currentModel.ClusterName == nil) {
		if prevModel != nil && prevModel.ProjectId != nil && prevModel.ClusterName != nil {
			currentModel = prevModel
		} else {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Search deployment not found",
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}, nil
		}
	}

	connV2, peErr := InitEnvWithClient(req, currentModel, DeleteRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	if IsCallback(&req) {
		return ValidateProgress(*connV2, currentModel, true), nil
	}

	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)

	resp, err := connV2.AtlasSearchApi.DeleteClusterSearchDeployment(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		return HandleError(resp, err)
	}

	apiResp, _, readErr := connV2.AtlasSearchApi.GetClusterSearchDeployment(context.Background(), projectID, clusterName).Execute()
	if readErr == nil && apiResp != nil && apiResp.Id != nil {
		updatedModel := NewCFNSearchDeployment(currentModel, apiResp)
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              constants.DeleteInProgress,
			ResourceModel:        &updatedModel,
			CallbackDelaySeconds: CallBackSeconds,
			CallbackContext:      callbackContext,
		}, nil
	}

	modelWithID := GetModelWithID(currentModel, prevModel, nil)
	if modelWithID != nil {
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              constants.DeleteInProgress,
			ResourceModel:        modelWithID,
			CallbackDelaySeconds: CallBackSeconds,
			CallbackContext:      callbackContext,
		}, nil
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              constants.DeleteInProgress,
		ResourceModel:        currentModel,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext:      callbackContext,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func HandleError(res *http.Response, err error) (handler.ProgressEvent, error) {
	pe := progressevent.GetFailedEventByResponse(err.Error(), res)

	// Search Deployment API returns 400 BadRequest for both AlreadyExists and NotFound
	// Need to check error code to distinguish them
	if res != nil && res.StatusCode == http.StatusBadRequest {
		if apiError, ok := admin20250312010.AsError(err); ok {
			if strings.Contains(apiError.ErrorCode, SearchDeploymentAlreadyExistsErrorAPI) {
				pe.HandlerErrorCode = string(types.HandlerErrorCodeAlreadyExists)
			} else if strings.Contains(apiError.ErrorCode, SearchDeploymentDoesNotExistsErrorAPI) {
				pe.HandlerErrorCode = string(types.HandlerErrorCodeNotFound)
			}
		}
	}

	if res != nil && res.StatusCode == http.StatusNotFound {
		pe.HandlerErrorCode = string(types.HandlerErrorCodeNotFound)
	}
	if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "being deleted") {
		pe.HandlerErrorCode = string(types.HandlerErrorCodeNotFound)
	}
	return pe, nil
}

func inProgressEvent(message string, model *Model, apiResp *admin20250312010.ApiSearchDeploymentResponse) handler.ProgressEvent {
	if apiResp != nil {
		newModel := NewCFNSearchDeployment(model, apiResp)
		model = &newModel
	}
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              message,
		ResourceModel:        model,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext:      callbackContext,
	}
}

func GetModelWithID(currentModel, prevModel *Model, checkResp *admin20250312010.ApiSearchDeploymentResponse) *Model {
	if currentModel != nil && currentModel.Id != nil {
		return currentModel
	}
	if prevModel != nil && prevModel.Id != nil {
		return prevModel
	}
	if checkResp != nil && checkResp.Id != nil {
		updatedModel := NewCFNSearchDeployment(currentModel, checkResp)
		return &updatedModel
	}
	return nil
}
