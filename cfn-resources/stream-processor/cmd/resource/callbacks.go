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
	"maps"

	"go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
)

type CallbackData struct {
	ProjectID               string
	WorkspaceOrInstanceName string
	ProcessorName           string
	DesiredState            string
	StartTime               string
	TimeoutDuration         string
	NeedsStarting           bool
	DeleteOnCreateTimeout   bool
}

func IsCallback(req *handler.Request) bool {
	_, found := req.CallbackContext["callbackStreamProcessor"]
	return found
}

func GetCallbackData(req handler.Request) *CallbackData {
	ctx := &CallbackData{}

	if val, ok := req.CallbackContext["projectID"].(string); ok {
		ctx.ProjectID = val
	}
	if val, ok := req.CallbackContext["workspaceName"].(string); ok {
		ctx.WorkspaceOrInstanceName = val
	}
	if val, ok := req.CallbackContext["processorName"].(string); ok {
		ctx.ProcessorName = val
	}
	if val, ok := req.CallbackContext["needsStarting"].(bool); ok {
		ctx.NeedsStarting = val
	}
	if val, ok := req.CallbackContext["desiredState"].(string); ok {
		ctx.DesiredState = val
	}
	if val, ok := req.CallbackContext["startTime"].(string); ok {
		ctx.StartTime = val
	}
	if val, ok := req.CallbackContext["timeoutDuration"].(string); ok {
		ctx.TimeoutDuration = val
	}
	if val, ok := req.CallbackContext["deleteOnCreateTimeout"].(bool); ok {
		ctx.DeleteOnCreateTimeout = val
	}

	return ctx
}

func ValidateCallbackData(ctx *CallbackData) *handler.ProgressEvent {
	if ctx.ProjectID == "" || ctx.WorkspaceOrInstanceName == "" || ctx.ProcessorName == "" {
		return &handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Missing required values in callback context",
		}
	}
	return nil
}

func BuildCallbackContext(projectID, workspaceOrInstanceName, processorName string, additionalFields map[string]any) map[string]any {
	ctx := map[string]any{
		"callbackStreamProcessor": true,
		"projectID":               projectID,
		"workspaceName":           workspaceOrInstanceName,
		"processorName":           processorName,
	}

	maps.Copy(ctx, additionalFields)

	return ctx
}

func cleanupOnCreateTimeout(ctx context.Context, atlasClient *admin.APIClient, callbackCtx *CallbackData) error {
	if !callbackCtx.DeleteOnCreateTimeout {
		return nil
	}

	_, err := atlasClient.StreamsApi.DeleteStreamProcessor(ctx, callbackCtx.ProjectID, callbackCtx.WorkspaceOrInstanceName, callbackCtx.ProcessorName).Execute()
	if err != nil {
		_, _ = logger.Warnf("Cleanup delete failed: %v", err)
	}
	return nil
}

func HandleCreateCallback(ctx context.Context, atlasClient *admin.APIClient, currentModel *Model, callbackCtx *CallbackData) (handler.ProgressEvent, error) {
	needsStarting := callbackCtx.NeedsStarting

	if IsTimeoutExceeded(callbackCtx.StartTime, callbackCtx.TimeoutDuration) {
		if err := cleanupOnCreateTimeout(context.Background(), atlasClient, callbackCtx); err != nil {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         fmt.Sprintf("Timeout reached and cleanup failed: %s", err.Error()),
			}, nil
		}
		cleanupMsg := "Timeout reached when waiting for stream processor creation"
		if callbackCtx.DeleteOnCreateTimeout {
			cleanupMsg += ". Resource has been deleted because delete_on_create_timeout is true. If you suspect a transient error, wait before retrying to allow resource deletion to finish."
		} else {
			cleanupMsg += ". Cleanup was not performed because delete_on_create_timeout is false."
		}
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         cleanupMsg,
		}, nil
	}

	streamProcessor, peErr := getStreamProcessor(ctx, atlasClient, callbackCtx.ProjectID, callbackCtx.WorkspaceOrInstanceName, callbackCtx.ProcessorName)
	if peErr != nil {
		return *peErr, nil
	}

	currentState := streamProcessor.GetState()

	callbackContext := BuildCallbackContext(callbackCtx.ProjectID, callbackCtx.WorkspaceOrInstanceName, callbackCtx.ProcessorName, map[string]any{
		"needsStarting":         callbackCtx.NeedsStarting,
		"startTime":             callbackCtx.StartTime,
		"timeoutDuration":       callbackCtx.TimeoutDuration,
		"deleteOnCreateTimeout": callbackCtx.DeleteOnCreateTimeout,
	})

	switch currentState {
	case CreatedState:
		if needsStarting {
			if peErr := startStreamProcessor(ctx, atlasClient, callbackCtx.ProjectID, callbackCtx.WorkspaceOrInstanceName, callbackCtx.ProcessorName); peErr != nil {
				return *peErr, nil
			}
			return createInProgressEvent("Starting stream processor", currentModel, callbackContext), nil
		}
		return FinalizeModel(streamProcessor, currentModel, "Create Completed")

	case StartedState:
		return FinalizeModel(streamProcessor, currentModel, "Create Completed")

	case InitiatingState, CreatingState:
		return createInProgressEvent(fmt.Sprintf("Creating stream processor (current state: %s)", currentState), currentModel, callbackContext), nil

	case FailedState:
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Stream processor entered FAILED state",
		}, nil

	default:
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         fmt.Sprintf("Unexpected state during creation: %s", currentState),
		}, nil
	}
}

func HandleUpdateCallback(ctx context.Context, atlasClient *admin.APIClient, currentModel *Model, callbackCtx *CallbackData) (handler.ProgressEvent, error) {
	streamProcessor, peErr := getStreamProcessor(ctx, atlasClient, callbackCtx.ProjectID, callbackCtx.WorkspaceOrInstanceName, callbackCtx.ProcessorName)
	if peErr != nil {
		return *peErr, nil
	}

	desiredState := callbackCtx.DesiredState
	if desiredState == "" {
		desiredState = streamProcessor.GetState()
		if desiredState == "" {
			if currentModel != nil && currentModel.DesiredState != nil && *currentModel.DesiredState != "" {
				desiredState = *currentModel.DesiredState
			} else {
				desiredState = CreatedState
			}
		}
	}

	currentState := streamProcessor.GetState()

	callbackContext := BuildCallbackContext(callbackCtx.ProjectID, callbackCtx.WorkspaceOrInstanceName, callbackCtx.ProcessorName, map[string]any{
		"desiredState": desiredState,
	})

	switch currentState {
	case StoppedState, CreatedState:
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
			if peErr := startStreamProcessor(ctx, atlasClient, callbackCtx.ProjectID, callbackCtx.WorkspaceOrInstanceName, callbackCtx.ProcessorName); peErr != nil {
				return *peErr, nil
			}
			return createInProgressEvent("Starting stream processor", currentModel, callbackContext), nil
		}

		return FinalizeModel(streamProcessorResp, currentModel, "Update Completed")

	case StartedState:
		if desiredState == StartedState {
			return FinalizeModel(streamProcessor, currentModel, "Update Completed")
		}

		_, err := atlasClient.StreamsApi.StopStreamProcessorWithParams(ctx,
			&admin.StopStreamProcessorApiParams{
				GroupId:       callbackCtx.ProjectID,
				TenantName:    callbackCtx.WorkspaceOrInstanceName,
				ProcessorName: callbackCtx.ProcessorName,
			},
		).Execute()
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         fmt.Sprintf("Error stopping stream processor: %s", err.Error()),
			}, nil
		}
		return createInProgressEvent("Stopping stream processor", currentModel, callbackContext), nil

	case FailedState:
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Stream processor entered FAILED state",
		}, nil

	default:
		return createInProgressEvent(fmt.Sprintf("Updating stream processor (current state: %s)", currentState), currentModel, callbackContext), nil
	}
}
