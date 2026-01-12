// Copyright 2026 MongoDB Inc
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
	"time"

	"go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

const (
	InitiatingState = "INIT"
	CreatingState   = "CREATING"
	CreatedState    = "CREATED"
	StartedState    = "STARTED"
	StoppedState    = "STOPPED"
	DroppedState    = "DROPPED"
	FailedState     = "FAILED"
)

const (
	defaultCallbackDelaySeconds = 3
	DefaultCreateTimeout        = 20 * time.Minute
)

func Setup() {
	util.SetupLogger("mongodb-atlas-stream-processor")
}

var CreateRequiredFields = []string{constants.ProjectID, constants.ProcessorName, constants.Pipeline}
var ReadRequiredFields = []string{constants.ProjectID, constants.ProcessorName}
var UpdateRequiredFields = []string{constants.ProjectID, constants.ProcessorName, constants.Pipeline}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ProcessorName}
var ListRequiredFields = []string{constants.ProjectID}

var InitEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin.APIClient, *handler.ProgressEvent) {
	Setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if errEvent := validator.ValidateModel(requiredFields, currentModel); errEvent != nil {
		return nil, errEvent
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return nil, peErr
	}
	return client.AtlasSDK, nil
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	atlasClient, peErr := InitEnvWithLatestClient(req, currentModel, CreateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	if IsCallback(&req) {
		callbackCtx := GetCallbackData(req)
		if peErr := ValidateCallbackData(callbackCtx); peErr != nil {
			return *peErr, nil
		}
		return HandleCreateCallback(
			context.Background(),
			atlasClient,
			currentModel,
			callbackCtx,
		)
	}

	workspaceOrInstanceName, err := GetWorkspaceOrInstanceName(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
		}, nil
	}

	ctx := context.Background()
	projectID := util.SafeString(currentModel.ProjectId)
	processorName := util.SafeString(currentModel.ProcessorName)

	var needsStarting bool
	if currentModel.DesiredState != nil {
		state := *currentModel.DesiredState
		switch state {
		case StartedState:
			needsStarting = true
		case CreatedState:
			needsStarting = false
		default:
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "When creating a stream processor, the only valid states are CREATED and STARTED",
			}, nil
		}
	}

	streamProcessorReq, err := NewStreamProcessorReq(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         fmt.Sprintf("Error creating stream processor request: %s", err.Error()),
		}, nil
	}

	_, apiResp, err := atlasClient.StreamsApi.CreateStreamProcessor(ctx, projectID, workspaceOrInstanceName, streamProcessorReq).Execute()
	if err != nil {
		return HandleError(apiResp, constants.CREATE, err)
	}

	timeoutStr := ""
	if currentModel.Timeouts != nil && currentModel.Timeouts.Create != nil {
		timeoutStr = *currentModel.Timeouts.Create
	}

	deleteOnCreateTimeout := true
	if currentModel.DeleteOnCreateTimeout != nil {
		deleteOnCreateTimeout = *currentModel.DeleteOnCreateTimeout
	}

	inProgressModel := &Model{}
	if currentModel != nil {
		*inProgressModel = *currentModel
		inProgressModel.DeleteOnCreateTimeout = nil
	}
	CopyIdentifyingFields(inProgressModel, currentModel)

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              "Creating stream processor",
		ResourceModel:        inProgressModel,
		CallbackDelaySeconds: defaultCallbackDelaySeconds,
		CallbackContext: BuildCallbackContext(projectID, workspaceOrInstanceName, processorName, map[string]any{
			"needsStarting":         needsStarting,
			"startTime":             time.Now().Format(time.RFC3339),
			"timeoutDuration":       timeoutStr,
			"deleteOnCreateTimeout": deleteOnCreateTimeout,
		}),
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	atlasClient, peErr := InitEnvWithLatestClient(req, currentModel, ReadRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	workspaceOrInstanceName, err := GetWorkspaceOrInstanceName(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
		}, nil
	}

	projectID := util.SafeString(currentModel.ProjectId)
	processorName := util.SafeString(currentModel.ProcessorName)

	streamProcessor, apiResp, err := atlasClient.StreamsApi.GetStreamProcessorWithParams(context.Background(),
		&admin.GetStreamProcessorApiParams{
			GroupId:       projectID,
			TenantName:    workspaceOrInstanceName,
			ProcessorName: processorName,
		}).Execute()
	if err != nil {
		if apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Resource not found",
				HandlerErrorCode: "NotFound",
			}, nil
		}
		return HandleError(apiResp, constants.READ, err)
	}

	resourceModel, err := GetStreamProcessorModel(streamProcessor, currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         fmt.Sprintf("Error converting stream processor model: %s", err.Error()),
		}, nil
	}

	CopyIdentifyingFields(resourceModel, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
		ResourceModel:   resourceModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	atlasClient, peErr := InitEnvWithLatestClient(req, currentModel, UpdateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	if IsCallback(&req) {
		callbackCtx := GetCallbackData(req)
		if peErr := ValidateCallbackData(callbackCtx); peErr != nil {
			return *peErr, nil
		}
		return HandleUpdateCallback(
			context.Background(),
			atlasClient,
			currentModel,
			callbackCtx,
		)
	}

	workspaceOrInstanceName, err := GetWorkspaceOrInstanceName(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
		}, nil
	}

	ctx := context.Background()
	projectID := util.SafeString(currentModel.ProjectId)
	processorName := util.SafeString(currentModel.ProcessorName)

	requestParams := &admin.GetStreamProcessorApiParams{
		GroupId:       projectID,
		TenantName:    workspaceOrInstanceName,
		ProcessorName: processorName,
	}

	currentStreamProcessor, apiResp, err := atlasClient.StreamsApi.GetStreamProcessorWithParams(ctx, requestParams).Execute()
	if err != nil {
		if apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Resource not found",
				HandlerErrorCode: "NotFound",
			}, nil
		}
		return HandleError(apiResp, constants.READ, err)
	}

	currentState := currentStreamProcessor.GetState()

	desiredState := currentState
	if currentModel.DesiredState != nil && *currentModel.DesiredState != "" {
		desiredState = *currentModel.DesiredState
	} else if prevModel != nil && prevModel.DesiredState != nil && *prevModel.DesiredState != "" {
		desiredState = *prevModel.DesiredState
	}

	if errMsg, isValid := ValidateUpdateStateTransition(currentState, desiredState); !isValid {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         errMsg,
		}, nil
	}

	if currentState == StartedState {
		_, err := atlasClient.StreamsApi.StopStreamProcessorWithParams(ctx,
			&admin.StopStreamProcessorApiParams{
				GroupId:       projectID,
				TenantName:    workspaceOrInstanceName,
				ProcessorName: processorName,
			},
		).Execute()
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         fmt.Sprintf("Error stopping stream processor: %s", err.Error()),
			}, nil
		}

		inProgressModel := &Model{}
		if currentModel != nil {
			*inProgressModel = *currentModel
			inProgressModel.DeleteOnCreateTimeout = nil
		}
		CopyIdentifyingFields(inProgressModel, currentModel)

		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              "Stopping stream processor",
			ResourceModel:        inProgressModel,
			CallbackDelaySeconds: defaultCallbackDelaySeconds,
			CallbackContext: BuildCallbackContext(projectID, workspaceOrInstanceName, processorName, map[string]any{
				"desiredState": desiredState,
			}),
		}, nil
	}

	modifyAPIRequestParams, err := NewStreamProcessorUpdateReq(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         fmt.Sprintf("Error creating update request: %s", err.Error()),
		}, nil
	}

	streamProcessorResp, apiResp, err := atlasClient.StreamsApi.UpdateStreamProcessorWithParams(ctx, modifyAPIRequestParams).Execute()
	if err != nil {
		return HandleError(apiResp, constants.UPDATE, err)
	}

	if desiredState == StartedState {
		_, err := atlasClient.StreamsApi.StartStreamProcessorWithParams(ctx,
			&admin.StartStreamProcessorApiParams{
				GroupId:       projectID,
				TenantName:    workspaceOrInstanceName,
				ProcessorName: processorName,
			},
		).Execute()
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         fmt.Sprintf("Error starting stream processor: %s", err.Error()),
			}, nil
		}

		inProgressModel := &Model{}
		if currentModel != nil {
			*inProgressModel = *currentModel
			inProgressModel.DeleteOnCreateTimeout = nil
		}
		CopyIdentifyingFields(inProgressModel, currentModel)

		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              "Starting stream processor",
			ResourceModel:        inProgressModel,
			CallbackDelaySeconds: defaultCallbackDelaySeconds,
			CallbackContext: BuildCallbackContext(projectID, workspaceOrInstanceName, processorName, map[string]any{
				"desiredState": desiredState,
			}),
		}, nil
	}

	return FinalizeModel(streamProcessorResp, currentModel, "Update Completed")
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	atlasClient, peErr := InitEnvWithLatestClient(req, currentModel, ListRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	workspaceOrInstanceName, err := GetWorkspaceOrInstanceName(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
		}, nil
	}

	ctx := context.Background()
	projectID := util.SafeString(currentModel.ProjectId)

	accumulatedProcessors, apiResp, err := getAllStreamProcessors(ctx, atlasClient, projectID, workspaceOrInstanceName)
	if err != nil {
		return HandleError(apiResp, constants.LIST, err)
	}

	response := make([]interface{}, 0, len(accumulatedProcessors))
	for i := range accumulatedProcessors {
		model, err := GetStreamProcessorModel(&accumulatedProcessors[i], currentModel)
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         fmt.Sprintf("Error converting stream processor model: %s", err.Error()),
			}, nil
		}

		CopyIdentifyingFields(model, currentModel)
		response = append(response, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Completed",
		ResourceModels:  response,
	}, nil
}

// Delete handles the Delete event from the CloudFormation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	atlasClient, peErr := InitEnvWithLatestClient(req, currentModel, DeleteRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	workspaceOrInstanceName, err := GetWorkspaceOrInstanceName(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
		}, nil
	}

	ctx := context.Background()
	projectID := util.SafeString(currentModel.ProjectId)
	processorName := util.SafeString(currentModel.ProcessorName)

	apiResp, err := atlasClient.StreamsApi.DeleteStreamProcessor(ctx, projectID, workspaceOrInstanceName, processorName).Execute()
	if err != nil {
		if apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Resource not found",
				HandlerErrorCode: "NotFound",
			}, nil
		}
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         fmt.Sprintf("Error deleting stream processor: %s", err.Error()),
		}, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
	}, nil
}
