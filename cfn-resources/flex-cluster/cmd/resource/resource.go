// Copyright 2025 MongoDB Inc
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

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"go.mongodb.org/atlas-sdk/v20250312006/admin"
)

const (
	callBackSeconds = 10
)

var createUpdateDeleteRequiredFields = []string{constants.ProjectID, constants.Name}

func setup() {
	util.SetupLogger("mongodb-atlas-flexcluster")
}

// validateModel inputs based on the method
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if modelValidation := validateModel(createUpdateDeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		progressEvent := validateProgress(client, currentModel, constants.IdleState)
		if progressEvent.Message == constants.Complete {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Create Success",
				ResourceModel:   currentModel,
			}, nil
		}
		return progressEvent, nil
	}

	// TODO: fill from model
	flexReq := &admin.FlexClusterDescriptionCreate20241113{
		Name: *currentModel.Name,
		ProviderSettings: admin.FlexProviderSettingsCreate20241113{
			BackingProviderName: "AWS",       // Write only field.
			RegionName:          "US_EAST_1", // Write only field.
			DiskSizeGB:          nil,         // Read only field.
			ProviderName:        nil,         // Read only field.
		},
		Tags:                         nil,
		TerminationProtectionEnabled: nil,
	}
	flexResp, resp, err := client.AtlasSDK.FlexClustersApi.CreateFlexCluster(context.Background(), *currentModel.ProjectId, flexReq).Execute()
	if err != nil {
		if apiError, ok := admin.AsError(err); ok && apiError.Error == http.StatusBadRequest && strings.Contains(apiError.ErrorCode, constants.Duplicate) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
		}
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}
	currentModel.StateName = flexResp.StateName

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create FlexCluster `%s`", flexResp.GetStateName()),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: callBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.StateName: flexResp.GetStateName(),
		},
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	flexResp, resp, err := client.AtlasSDK.FlexClustersApi.GetFlexCluster(context.Background(), *currentModel.ProjectId, *currentModel.Name).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	currentModel.StateName = flexResp.StateName
	currentModel.Id = flexResp.Id
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if modelValidation := validateModel(createUpdateDeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	_ = client // TODO: implement, not updatable attributes yet

	return handler.ProgressEvent{}, errors.New("not implemented: Update, not updatable attributes yet")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if modelValidation := validateModel(createUpdateDeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		return validateProgress(client, currentModel, constants.DeletedState), nil
	}
	resp, err := client.AtlasSDK.FlexClustersApi.DeleteFlexCluster(context.Background(), *currentModel.ProjectId, *currentModel.Name).Execute()
	if err != nil {
		if apiError, ok := admin.AsError(err); ok && apiError.Error == http.StatusBadRequest && strings.Contains(apiError.ErrorCode, constants.Duplicate) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
		}
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              constants.DeleteInProgress,
		ResourceModel:        currentModel,
		CallbackDelaySeconds: callBackSeconds,
		CallbackContext: map[string]any{
			constants.StateName: constants.DeletingState,
		}}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func validateProgress(client *util.MongoDBClient, currentModel *Model, targetState string) handler.ProgressEvent {
	state, err := isClusterInTargetState(client, *currentModel.ProjectId, *currentModel.Name)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}
	}
	currentModel.StateName = &state
	if state == targetState {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         constants.Complete,
			ResourceModel:   currentModel,
		}
	}
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              constants.Pending,
		ResourceModel:        currentModel,
		CallbackDelaySeconds: callBackSeconds,
		CallbackContext: map[string]any{
			constants.StateName: state,
		},
	}
}

func isClusterInTargetState(client *util.MongoDBClient, projectID, clusterName string) (stateName string, err error) {
	cluster, resp, err := client.AtlasSDK.FlexClustersApi.GetFlexCluster(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return constants.DeletedState, nil
		}
		return constants.Error, fmt.Errorf("error fetching flex cluster info (%s): %s", clusterName, err)
	}
	return *cluster.StateName, nil
}
