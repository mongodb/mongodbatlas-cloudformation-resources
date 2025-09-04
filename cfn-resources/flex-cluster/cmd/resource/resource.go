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

var createRequiredFields = []string{constants.ProjectID, constants.Name, "ProviderSettings"}
var updateDeleteRequiredFields = []string{constants.ProjectID, constants.Name}

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
	if modelValidation := validateModel(createRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		progressEvent := validateProgress(client, currentModel, false)
		if progressEvent.Message == constants.Complete {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Create Success",
				ResourceModel:   currentModel,
			}, nil
		}
		return progressEvent, nil
	}

	// Build create request from model
	flexReq := &admin.FlexClusterDescriptionCreate20241113{
		Name: *currentModel.Name,
	}

	// Set ProviderSettings
	if currentModel.ProviderSettings != nil {
		flexReq.ProviderSettings = admin.FlexProviderSettingsCreate20241113{
			BackingProviderName: *currentModel.ProviderSettings.BackingProviderName,
			RegionName:          *currentModel.ProviderSettings.RegionName,
		}
	}

	// Set Tags if provided
	if len(currentModel.Tags) > 0 {
		tags := make([]admin.ResourceTag, len(currentModel.Tags))
		for i, tag := range currentModel.Tags {
			tags[i] = admin.ResourceTag{
				Key:   *tag.Key,
				Value: *tag.Value,
			}
		}
		flexReq.Tags = &tags
	}

	// Set TerminationProtectionEnabled if provided
	if currentModel.TerminationProtectionEnabled != nil {
		flexReq.TerminationProtectionEnabled = currentModel.TerminationProtectionEnabled
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
	updateModel(currentModel, flexResp)
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create FlexCluster `%s`", flexResp.GetStateName()),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: callBackSeconds,
		CallbackContext: map[string]any{
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
	updateModel(currentModel, flexResp)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if modelValidation := validateModel(updateDeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		progressEvent := validateProgress(client, currentModel, false)
		if progressEvent.Message == constants.Complete {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Update Success",
				ResourceModel:   currentModel,
			}, nil
		}
		return progressEvent, nil
	}

	// Build update request - only Tags and TerminationProtectionEnabled are updatable
	updateReq := &admin.FlexClusterDescriptionUpdate20241113{}

	// Update Tags if provided
	if len(currentModel.Tags) > 0 {
		tags := make([]admin.ResourceTag, len(currentModel.Tags))
		for i, tag := range currentModel.Tags {
			tags[i] = admin.ResourceTag{
				Key:   *tag.Key,
				Value: *tag.Value,
			}
		}
		updateReq.Tags = &tags
	}

	// Update TerminationProtectionEnabled if provided
	if currentModel.TerminationProtectionEnabled != nil {
		updateReq.TerminationProtectionEnabled = currentModel.TerminationProtectionEnabled
	}

	// Perform the update
	flexResp, resp, err := client.AtlasSDK.FlexClustersApi.UpdateFlexCluster(context.Background(), *currentModel.ProjectId, *currentModel.Name, updateReq).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		code := cloudformation.HandlerErrorCodeServiceInternalError
		if strings.Contains(err.Error(), "not exist") { // cfn test needs 404
			code = cloudformation.HandlerErrorCodeNotFound
		}
		if strings.Contains(err.Error(), "being deleted") {
			code = cloudformation.HandlerErrorCodeNotFound // cfn test needs 404
		}
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: code}, nil
	}
	updateModel(currentModel, flexResp)
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Update Cluster, state: %s", flexResp.GetStateName()),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: callBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.StateName: flexResp.GetStateName(),
		},
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if modelValidation := validateModel(updateDeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		return validateProgress(client, currentModel, true), nil
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

func validateProgress(client *util.MongoDBClient, currentModel *Model, isDelete bool) handler.ProgressEvent {
	state, err := getState(client, *currentModel.ProjectId, *currentModel.Name)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}
	}
	currentModel.StateName = &state
	targetState := constants.IdleState
	model := currentModel
	if isDelete {
		targetState = constants.DeletedState
		model = nil // Delete event shouldn't have model in the response
	}
	if state == targetState {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         constants.Complete,
			ResourceModel:   model,
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

func getState(client *util.MongoDBClient, projectID, clusterName string) (stateName string, err error) {
	cluster, resp, err := client.AtlasSDK.FlexClustersApi.GetFlexCluster(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return constants.DeletedState, nil
		}
		return constants.Error, fmt.Errorf("error fetching flex cluster info (%s): %s", clusterName, err)
	}
	return *cluster.StateName, nil
}

func updateModel(currentModel *Model, flexResp *admin.FlexClusterDescription20241113) {
	currentModel.ProjectId = flexResp.GroupId
	currentModel.Name = flexResp.Name
	currentModel.Id = flexResp.Id
	currentModel.StateName = flexResp.StateName
	currentModel.ClusterType = flexResp.ClusterType
	currentModel.CreateDate = util.TimePtrToStringPtr(flexResp.CreateDate)
	currentModel.MongoDBVersion = flexResp.MongoDBVersion
	currentModel.VersionReleaseSystem = flexResp.VersionReleaseSystem
	currentModel.TerminationProtectionEnabled = flexResp.TerminationProtectionEnabled

	// Update ProviderSettings (not a pointer type in the SDK)
	if currentModel.ProviderSettings == nil {
		currentModel.ProviderSettings = &ProviderSettings{}
	}
	currentModel.ProviderSettings.BackingProviderName = flexResp.ProviderSettings.BackingProviderName
	currentModel.ProviderSettings.RegionName = flexResp.ProviderSettings.RegionName
	currentModel.ProviderSettings.DiskSizeGB = flexResp.ProviderSettings.DiskSizeGB
	currentModel.ProviderSettings.ProviderName = flexResp.ProviderSettings.ProviderName

	// Update BackupSettings
	if flexResp.BackupSettings != nil {
		currentModel.BackupSettings = &BackupSettings{
			Enabled: flexResp.BackupSettings.Enabled,
		}
	}

	// Update ConnectionStrings
	if flexResp.ConnectionStrings != nil {
		currentModel.ConnectionStrings = &ConnectionStrings{
			Standard:    flexResp.ConnectionStrings.Standard,
			StandardSrv: flexResp.ConnectionStrings.StandardSrv,
		}
	}

	// Update Tags
	if flexResp.Tags != nil && len(*flexResp.Tags) > 0 {
		tags := make([]Tag, len(*flexResp.Tags))
		for i, tag := range *flexResp.Tags {
			tags[i] = Tag{
				Key:   &tag.Key,
				Value: &tag.Value,
			}
		}
		currentModel.Tags = tags
	}
}
