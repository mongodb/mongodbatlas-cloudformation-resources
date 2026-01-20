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

	"go.mongodb.org/atlas-sdk/v20250312012/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
)

type CallbackData struct {
	ProjectID             string
	WorkspaceName         string
	ProcessorName         string
	DesiredState          string
	StartTime             string
	TimeoutDuration       string
	NeedsStarting         bool
	DeleteOnCreateTimeout bool
}

func getCallbackData(req handler.Request) *CallbackData {
	ctx := &CallbackData{}

	if val, ok := req.CallbackContext["projectID"].(string); ok {
		ctx.ProjectID = val
	}
	if val, ok := req.CallbackContext["workspaceName"].(string); ok {
		ctx.WorkspaceName = val
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

func validateCallbackData(ctx *CallbackData) *handler.ProgressEvent {
	if ctx.ProjectID == "" || ctx.WorkspaceName == "" || ctx.ProcessorName == "" {
		return &handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Missing required values in callback context",
		}
	}
	return nil
}

func buildCallbackContext(projectID, workspaceName, processorName string, additionalFields map[string]any) map[string]any {
	ctx := map[string]any{
		"callbackStreamProcessor": true,
		"projectID":               projectID,
		"workspaceName":           workspaceName,
		"processorName":           processorName,
	}

	maps.Copy(ctx, additionalFields)

	return ctx
}

func cleanupOnCreateTimeout(ctx context.Context, client *util.MongoDBClient, callbackCtx *CallbackData) error {
	if !callbackCtx.DeleteOnCreateTimeout {
		return nil
	}

	_, err := client.AtlasSDK.StreamsApi.DeleteStreamProcessor(ctx, callbackCtx.ProjectID, callbackCtx.WorkspaceName, callbackCtx.ProcessorName).Execute()
	if err != nil {
		_, _ = logger.Warnf("Cleanup delete failed: %v", err)
		return err
	}
	return nil
}

func handleCreateCallback(ctx context.Context, client *util.MongoDBClient, currentModel *Model, callbackCtx *CallbackData) handler.ProgressEvent {
	needsStarting := callbackCtx.NeedsStarting

	if isTimeoutExceeded(callbackCtx.StartTime, callbackCtx.TimeoutDuration) {
		if err := cleanupOnCreateTimeout(context.Background(), client, callbackCtx); err != nil {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         fmt.Sprintf("Timeout reached and cleanup failed: %s", err.Error()),
			}
		}
		cleanupMsg := "Timeout reached when waiting for stream processor creation"
		if callbackCtx.DeleteOnCreateTimeout {
			cleanupMsg += ". Deletion of resource has been triggered because delete_on_create_timeout is true. If you suspect a transient error, wait before retrying to allow resource deletion to finish."
		} else {
			cleanupMsg += ". Cleanup was not performed because delete_on_create_timeout is false."
		}
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         cleanupMsg,
		}
	}

	streamProcessor, peErr := getStreamProcessor(ctx, client.AtlasSDK, callbackCtx.ProjectID, callbackCtx.WorkspaceName, callbackCtx.ProcessorName)
	if peErr != nil {
		return *peErr
	}

	currentState := streamProcessor.GetState()

	callbackContext := buildCallbackContext(callbackCtx.ProjectID, callbackCtx.WorkspaceName, callbackCtx.ProcessorName, map[string]any{
		"needsStarting":         callbackCtx.NeedsStarting,
		"startTime":             callbackCtx.StartTime,
		"timeoutDuration":       callbackCtx.TimeoutDuration,
		"deleteOnCreateTimeout": callbackCtx.DeleteOnCreateTimeout,
	})

	switch currentState {
	case CreatedState:
		if needsStarting {
			if peErr := startStreamProcessor(ctx, client.AtlasSDK, callbackCtx.ProjectID, callbackCtx.WorkspaceName, callbackCtx.ProcessorName); peErr != nil {
				return *peErr
			}
			return createInProgressEvent(constants.Pending, currentModel, callbackContext)
		}
		return finalizeModel(streamProcessor, currentModel, constants.Complete)

	case StartedState:
		return finalizeModel(streamProcessor, currentModel, constants.Complete)

	case InitiatingState, CreatingState:
		return createInProgressEvent(constants.Pending, currentModel, callbackContext)

	case FailedState:
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Stream processor entered FAILED state",
		}

	default:
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         fmt.Sprintf("Unexpected state during creation: %s", currentState),
		}
	}
}

func handleUpdateCallback(ctx context.Context, client *util.MongoDBClient, currentModel *Model, callbackCtx *CallbackData) handler.ProgressEvent {
	streamProcessor, peErr := getStreamProcessor(ctx, client.AtlasSDK, callbackCtx.ProjectID, callbackCtx.WorkspaceName, callbackCtx.ProcessorName)
	if peErr != nil {
		return *peErr
	}

	desiredState := callbackCtx.DesiredState
	if desiredState == "" {
		desiredState = streamProcessor.GetState()
	}
	if desiredState == "" && currentModel != nil && currentModel.DesiredState != nil && *currentModel.DesiredState != "" {
		desiredState = *currentModel.DesiredState
	}
	if desiredState == "" {
		desiredState = CreatedState
	}

	currentState := streamProcessor.GetState()

	callbackContext := buildCallbackContext(callbackCtx.ProjectID, callbackCtx.WorkspaceName, callbackCtx.ProcessorName, map[string]any{
		"desiredState": desiredState,
	})

	switch currentState {
	case StoppedState, CreatedState:
		modifyAPIRequestParams, err := NewStreamProcessorUpdateReq(currentModel)
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         fmt.Sprintf("Error creating update request: %s", err.Error()),
			}
		}

		streamProcessorResp, apiResp, err := client.AtlasSDK.StreamsApi.UpdateStreamProcessorWithParams(ctx, modifyAPIRequestParams).Execute()
		if err != nil {
			return handleError(apiResp, constants.UPDATE, err)
		}

		if desiredState == StartedState {
			if peErr := startStreamProcessor(ctx, client.AtlasSDK, callbackCtx.ProjectID, callbackCtx.WorkspaceName, callbackCtx.ProcessorName); peErr != nil {
				return *peErr
			}
			return createInProgressEvent(constants.Pending, currentModel, callbackContext)
		}

		return finalizeModel(streamProcessorResp, currentModel, constants.Complete)

	case StartedState:
		if desiredState == StartedState {
			return finalizeModel(streamProcessor, currentModel, constants.Complete)
		}

		// Only StoppedState is a valid transition from StartedState
		// (CreatedState transitions are not allowed per validateUpdateStateTransition)
		if desiredState != StoppedState {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         fmt.Sprintf("Unexpected desired state %s when current state is %s. Only %s is allowed.", desiredState, StartedState, StoppedState),
			}
		}

		_, err := client.AtlasSDK.StreamsApi.StopStreamProcessorWithParams(ctx,
			&admin.StopStreamProcessorApiParams{
				GroupId:       callbackCtx.ProjectID,
				TenantName:    callbackCtx.WorkspaceName,
				ProcessorName: callbackCtx.ProcessorName,
			},
		).Execute()
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         fmt.Sprintf("Error stopping stream processor: %s", err.Error()),
			}
		}
		return createInProgressEvent(constants.Pending, currentModel, callbackContext)

	case FailedState:
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Stream processor entered FAILED state",
		}

	default:
		return createInProgressEvent(constants.Pending, currentModel, callbackContext)
	}
}
