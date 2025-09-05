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

const callBackSeconds = 10

var (
	createRequiredFields           = []string{constants.ProjectID, constants.Name, "ProviderSettings"}
	readUpdateDeleteRequiredFields = []string{constants.ProjectID, constants.Name}
	listRequiredFields             = []string{constants.ProjectID}
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, createRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		return handleCallback(client, model, false, "Create Success"), nil
	}
	flexReq := &admin.FlexClusterDescriptionCreate20241113{
		Name: *model.Name,
		ProviderSettings: admin.FlexProviderSettingsCreate20241113{
			BackingProviderName: *model.ProviderSettings.BackingProviderName,
			RegionName:          *model.ProviderSettings.RegionName,
		},
		TerminationProtectionEnabled: model.TerminationProtectionEnabled,
		Tags:                         expandTags(model.Tags),
	}
	flexResp, resp, err := client.AtlasSDK.FlexClustersApi.CreateFlexCluster(context.Background(), *model.ProjectId, flexReq).Execute()
	if err != nil {
		if apiError, ok := admin.AsError(err); ok && apiError.Error == http.StatusBadRequest && strings.Contains(apiError.ErrorCode, constants.Duplicate) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
		}
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}
	updateModel(model, flexResp)
	return inProgressEvent(model, flexResp.GetStateName(), fmt.Sprintf("Create FlexCluster `%s`", flexResp.GetStateName())), nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, readUpdateDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	flexResp, resp, err := client.AtlasSDK.FlexClustersApi.GetFlexCluster(context.Background(), *model.ProjectId, *model.Name).Execute()
	if err != nil {
		return handleReadError(err, resp), nil
	}
	updateModel(model, flexResp)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   model,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, readUpdateDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		return handleCallback(client, model, false, "Update Success"), nil
	}
	updateReq := &admin.FlexClusterDescriptionUpdate20241113{
		TerminationProtectionEnabled: model.TerminationProtectionEnabled,
		Tags:                         expandTags(model.Tags),
	}
	flexResp, resp, err := client.AtlasSDK.FlexClustersApi.UpdateFlexCluster(context.Background(), *model.ProjectId, *model.Name, updateReq).Execute()
	if err != nil {
		return handleUpdateError(err, resp), nil
	}
	updateModel(model, flexResp)
	return inProgressEvent(model, flexResp.GetStateName(), fmt.Sprintf("Update Cluster, state: %s", flexResp.GetStateName())), nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, readUpdateDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		return validateProgress(client, model, true), nil
	}
	resp, err := client.AtlasSDK.FlexClustersApi.DeleteFlexCluster(context.Background(), *model.ProjectId, *model.Name).Execute()
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
		ResourceModel:        model,
		CallbackDelaySeconds: callBackSeconds,
		CallbackContext: map[string]any{
			constants.StateName: constants.DeletingState,
		}}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, listRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	listOptions := &admin.ListFlexClustersApiParams{
		GroupId:      *model.ProjectId,
		ItemsPerPage: admin.PtrInt(100),
		PageNum:      admin.PtrInt(1),
	}
	flexListResp, resp, err := client.AtlasSDK.FlexClustersApi.ListFlexClustersWithParams(context.Background(), listOptions).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error listing flex clusters: %s", err.Error()), resp), nil
	}
	results := flexListResp.GetResults()
	models := make([]*Model, len(results))
	for i := range results {
		modelItem := &Model{}
		updateModel(modelItem, &results[i])
		models[i] = modelItem
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List",
		ResourceModel:   models,
	}, nil
}

func setupRequest(req handler.Request, model *Model, requiredFields []string) (*util.MongoDBClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-flexcluster")
	if requiredFields != nil {
		if modelValidation := validator.ValidateModel(requiredFields, model); modelValidation != nil {
			return nil, modelValidation
		}
	}
	util.SetDefaultProfileIfNotDefined(&model.Profile)
	client, peErr := util.NewAtlasClient(&req, model.Profile)
	if peErr != nil {
		return nil, peErr
	}
	return client, nil
}

func expandTags(modelTags []Tag) *[]admin.ResourceTag {
	tags := make([]admin.ResourceTag, len(modelTags))
	for i, tag := range modelTags {
		tags[i] = admin.ResourceTag{
			Key:   *tag.Key,
			Value: *tag.Value,
		}
	}
	return &tags
}

func flattenTags(atlasTags *[]admin.ResourceTag) []Tag {
	if atlasTags == nil {
		return []Tag{}
	}
	tags := make([]Tag, len(*atlasTags))
	for i, tag := range *atlasTags {
		tags[i] = Tag{
			Key:   &tag.Key,
			Value: &tag.Value,
		}
	}
	return tags
}

func updateModel(model *Model, flexResp *admin.FlexClusterDescription20241113) {
	model.ProjectId = flexResp.GroupId
	model.Name = flexResp.Name
	model.Id = flexResp.Id
	model.StateName = flexResp.StateName
	model.ClusterType = flexResp.ClusterType
	model.CreateDate = util.TimePtrToStringPtr(flexResp.CreateDate)
	model.MongoDBVersion = flexResp.MongoDBVersion
	model.VersionReleaseSystem = flexResp.VersionReleaseSystem
	model.TerminationProtectionEnabled = flexResp.TerminationProtectionEnabled
	model.ProviderSettings = &ProviderSettings{
		BackingProviderName: flexResp.ProviderSettings.BackingProviderName,
		RegionName:          flexResp.ProviderSettings.RegionName,
		DiskSizeGB:          flexResp.ProviderSettings.DiskSizeGB,
		ProviderName:        flexResp.ProviderSettings.ProviderName,
	}
	if flexResp.BackupSettings != nil {
		model.BackupSettings = &BackupSettings{
			Enabled: flexResp.BackupSettings.Enabled,
		}
	}
	if flexResp.ConnectionStrings != nil {
		model.ConnectionStrings = &ConnectionStrings{
			Standard:    flexResp.ConnectionStrings.Standard,
			StandardSrv: flexResp.ConnectionStrings.StandardSrv,
		}
	}
	model.Tags = flattenTags(flexResp.Tags)
}

func handleCallback(client *util.MongoDBClient, model *Model, isDelete bool, successMessage string) handler.ProgressEvent {
	progressEvent := validateProgress(client, model, isDelete)
	if progressEvent.Message == constants.Complete {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         successMessage,
			ResourceModel:   model,
		}
	}
	return progressEvent
}

func handleReadError(err error, resp *http.Response) handler.ProgressEvent {
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
		}
	}
	return handler.ProgressEvent{
		Message:          err.Error(),
		OperationStatus:  handler.Failed,
		HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError,
	}
}

func handleUpdateError(err error, resp *http.Response) handler.ProgressEvent {
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
		}
	}
	code := cloudformation.HandlerErrorCodeServiceInternalError
	if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "being deleted") {
		code = cloudformation.HandlerErrorCodeNotFound
	}
	return handler.ProgressEvent{
		Message:          err.Error(),
		OperationStatus:  handler.Failed,
		HandlerErrorCode: code,
	}
}

func inProgressEvent(model *Model, stateName, message string) handler.ProgressEvent {
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              message,
		ResourceModel:        model,
		CallbackDelaySeconds: callBackSeconds,
		CallbackContext: map[string]any{
			constants.StateName: stateName,
		},
	}
}

func validateProgress(client *util.MongoDBClient, model *Model, isDelete bool) handler.ProgressEvent {
	state, err := getState(client, *model.ProjectId, *model.Name)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError,
		}
	}
	model.StateName = &state
	targetState := constants.IdleState
	if isDelete {
		targetState = constants.DeletedState
		model = nil // Delete event shouldn't have model in the response.
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
		ResourceModel:        model,
		CallbackDelaySeconds: callBackSeconds,
		CallbackContext: map[string]any{
			constants.StateName: state,
		},
	}
}

func getState(client *util.MongoDBClient, projectID, clusterName string) (string, error) {
	cluster, resp, err := client.AtlasSDK.FlexClustersApi.GetFlexCluster(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return constants.DeletedState, nil
		}
		return constants.Error, fmt.Errorf("error fetching flex cluster info (%s): %s", clusterName, err)
	}
	return *cluster.StateName, nil
}
