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
	"fmt"
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
	CallBackSeconds                    = 40
	SearchDeploymentDoesNotExistsError = "ATLAS_FTS_DEPLOYMENT_DOES_NOT_EXIST"
	SearchDeploymentAlreadyExistsError = "ATLAS_FTS_DEPLOYMENT_ALREADY_EXISTS"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.Specs}
var ReadRequiredFields = []string{constants.ProjectID, constants.ClusterName}
var UpdateRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.Specs}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ClusterName}

func setup() {
	util.SetupLogger("mongodb-atlas-searchdeployment")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel); modelValidation != nil {
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
	apiResp, res, err := connV2.AtlasSearchApi.CreateAtlasSearchDeployment(context.Background(), projectID, clusterName, &apiReq).Execute()
	if err != nil {
		if apiError, ok := admin.AsError(err); ok && *apiError.Error == http.StatusBadRequest && strings.Contains(*apiError.ErrorCode, SearchDeploymentAlreadyExistsError) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
		}
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Search Deployment: %s", err.Error()),
			res), nil
	}

	newModel := newCFNSearchDeployment(currentModel, apiResp)
	return inProgressEvent("Creating Search Deployment", &newModel), nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	client, progressErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if progressErr != nil {
		return *progressErr, nil
	}
	connV2 := client.AtlasSDK

	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	apiResp, res, err := connV2.AtlasSearchApi.GetAtlasSearchDeployment(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   newCFNSearchDeployment(currentModel, apiResp),
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel); modelValidation != nil {
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
		// specific handling is done here as NotFound error is returned with 400 status code
		if apiError, ok := admin.AsError(err); ok && *apiError.Error == http.StatusBadRequest && strings.Contains(*apiError.ErrorCode, SearchDeploymentDoesNotExistsError) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Failed to Update Search Deployment: %s", err.Error()),
			res), nil
	}

	newModel := newCFNSearchDeployment(currentModel, apiResp)
	return inProgressEvent("Updating Search Deployment", &newModel), nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel); modelValidation != nil {
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
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
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
		if targetState == constants.DeletedState && resp.StatusCode == 400 && strings.Contains(err.Error(), SearchDeploymentDoesNotExistsError) {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				ResourceModel:   currentModel,
				Message:         constants.Complete,
			}
		}

		return progressevent.GetFailedEventByResponse(err.Error(), resp)
	}

	newModel := newCFNSearchDeployment(currentModel, apiResp)
	if util.SafeString(currentModel.StateName) == targetState {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			ResourceModel:   newModel,
			Message:         constants.Complete,
		}
	}

	return inProgressEvent(constants.Pending, &newModel)
}

func newCFNSearchDeployment(prevModel *Model, apiResp *admin.ApiSearchDeploymentResponse) Model {
	respSpecs := apiResp.GetSpecs()
	resultSpecs := make([]ApiSearchDeploymentSpec, len(respSpecs))
	for i := range respSpecs {
		resultSpecs[i] = ApiSearchDeploymentSpec{
			InstanceSize: &respSpecs[i].InstanceSize,
			NodeCount:    &respSpecs[i].NodeCount,
		}
	}
	return Model{
		Profile:     prevModel.Profile,
		ClusterName: prevModel.ClusterName,
		ProjectId:   prevModel.ProjectId,
		Id:          apiResp.Id,
		Specs:       resultSpecs,
		StateName:   apiResp.StateName,
	}
}

func newSearchDeploymentReq(model *Model) admin.ApiSearchDeploymentRequest {
	modelSpecs := model.Specs
	requestSpecs := make([]admin.ApiSearchDeploymentSpec, len(modelSpecs))
	for i, spec := range modelSpecs {
		// Both spec fields are required in CFN model and will be defined
		requestSpecs[i] = admin.ApiSearchDeploymentSpec{
			InstanceSize: *spec.InstanceSize,
			NodeCount:    *spec.NodeCount,
		}
	}
	return admin.ApiSearchDeploymentRequest{Specs: &requestSpecs}
}

func inProgressEvent(message string, model *Model) handler.ProgressEvent {
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              message,
		ResourceModel:        model,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.ID: model.Id,
		}}
}
