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

	"go.mongodb.org/atlas-sdk/v20250312012/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
)

func IsCallback(req *handler.Request) bool {
	_, found := req.CallbackContext["callbackStreamProcessor"]
	return found
}

func HandleCreate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	if IsCallback(req) {
		callbackCtx := getCallbackData(*req)
		if peErr := validateCallbackData(callbackCtx); peErr != nil {
			return *peErr
		}
		return handleCreateCallback(
			context.Background(),
			client,
			model,
			callbackCtx,
		)
	}

	workspaceOrInstanceName, err := GetWorkspaceOrInstanceName(model)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
		}
	}

	ctx := context.Background()
	projectID := util.SafeString(model.ProjectId)
	processorName := util.SafeString(model.ProcessorName)

	var needsStarting bool
	if model.DesiredState != nil {
		state := *model.DesiredState
		switch state {
		case StartedState:
			needsStarting = true
		case CreatedState:
			needsStarting = false
		default:
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "When creating a stream processor, the only valid states are CREATED and STARTED",
			}
		}
	}

	streamProcessorReq, err := NewStreamProcessorReq(model)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         fmt.Sprintf("Error creating stream processor request: %s", err.Error()),
		}
	}

	_, apiResp, err := client.AtlasSDK.StreamsApi.CreateStreamProcessor(ctx, projectID, workspaceOrInstanceName, streamProcessorReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}

	timeoutStr := ""
	if model.Timeouts != nil && model.Timeouts.Create != nil {
		timeoutStr = *model.Timeouts.Create
	}

	deleteOnCreateTimeout := true
	if model.DeleteOnCreateTimeout != nil {
		deleteOnCreateTimeout = *model.DeleteOnCreateTimeout
	}

	inProgressModel := &Model{}
	if model != nil {
		*inProgressModel = *model
		inProgressModel.DeleteOnCreateTimeout = nil
	}
	copyIdentifyingFields(inProgressModel, model)

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              constants.Pending,
		ResourceModel:        inProgressModel,
		CallbackDelaySeconds: defaultCallbackDelaySeconds,
		CallbackContext: buildCallbackContext(projectID, workspaceOrInstanceName, processorName, map[string]any{
			"needsStarting":         needsStarting,
			"startTime":             time.Now().Format(time.RFC3339),
			"timeoutDuration":       timeoutStr,
			"deleteOnCreateTimeout": deleteOnCreateTimeout,
		}),
	}
}

func HandleRead(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	workspaceOrInstanceName, err := GetWorkspaceOrInstanceName(model)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
		}
	}

	projectID := util.SafeString(model.ProjectId)
	processorName := util.SafeString(model.ProcessorName)

	streamProcessor, apiResp, err := client.AtlasSDK.StreamsApi.GetStreamProcessorWithParams(context.Background(),
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
			}
		}
		return handleError(apiResp, constants.READ, err)
	}

	resourceModel, err := GetStreamProcessorModel(streamProcessor, model)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         fmt.Sprintf("Error converting stream processor model: %s", err.Error()),
		}
	}

	copyIdentifyingFields(resourceModel, model)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   resourceModel,
	}
}

func HandleUpdate(req *handler.Request, client *util.MongoDBClient, prevModel *Model, model *Model) handler.ProgressEvent {
	if IsCallback(req) {
		callbackCtx := getCallbackData(*req)
		if peErr := validateCallbackData(callbackCtx); peErr != nil {
			return *peErr
		}
		return handleUpdateCallback(
			context.Background(),
			client,
			model,
			callbackCtx,
		)
	}

	workspaceOrInstanceName, err := GetWorkspaceOrInstanceName(model)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
		}
	}

	ctx := context.Background()
	projectID := util.SafeString(model.ProjectId)
	processorName := util.SafeString(model.ProcessorName)

	requestParams := &admin.GetStreamProcessorApiParams{
		GroupId:       projectID,
		TenantName:    workspaceOrInstanceName,
		ProcessorName: processorName,
	}

	currentStreamProcessor, apiResp, err := client.AtlasSDK.StreamsApi.GetStreamProcessorWithParams(ctx, requestParams).Execute()
	if err != nil {
		if apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Resource not found",
				HandlerErrorCode: "NotFound",
			}
		}
		return handleError(apiResp, constants.READ, err)
	}

	currentState := currentStreamProcessor.GetState()

	desiredState := currentState
	if model.DesiredState != nil && *model.DesiredState != "" {
		desiredState = *model.DesiredState
	} else if prevModel != nil && prevModel.DesiredState != nil && *prevModel.DesiredState != "" {
		desiredState = *prevModel.DesiredState
	}

	if errMsg, isValid := validateUpdateStateTransition(currentState, desiredState); !isValid {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         errMsg,
		}
	}

	if currentState == StartedState {
		_, err := client.AtlasSDK.StreamsApi.StopStreamProcessorWithParams(ctx,
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
			}
		}

		inProgressModel := &Model{}
		if model != nil {
			*inProgressModel = *model
			inProgressModel.DeleteOnCreateTimeout = nil
		}
		copyIdentifyingFields(inProgressModel, model)

		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              constants.Pending,
			ResourceModel:        inProgressModel,
			CallbackDelaySeconds: defaultCallbackDelaySeconds,
			CallbackContext: buildCallbackContext(projectID, workspaceOrInstanceName, processorName, map[string]any{
				"desiredState": desiredState,
			}),
		}
	}

	modifyAPIRequestParams, err := NewStreamProcessorUpdateReq(model)
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
		_, err := client.AtlasSDK.StreamsApi.StartStreamProcessorWithParams(ctx,
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
			}
		}

		inProgressModel := &Model{}
		if model != nil {
			*inProgressModel = *model
			inProgressModel.DeleteOnCreateTimeout = nil
		}
		copyIdentifyingFields(inProgressModel, model)

		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              constants.Pending,
			ResourceModel:        inProgressModel,
			CallbackDelaySeconds: defaultCallbackDelaySeconds,
			CallbackContext: buildCallbackContext(projectID, workspaceOrInstanceName, processorName, map[string]any{
				"desiredState": desiredState,
			}),
		}
	}

	return finalizeModel(streamProcessorResp, model, constants.Complete)
}

func HandleDelete(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	workspaceOrInstanceName, err := GetWorkspaceOrInstanceName(model)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
		}
	}

	ctx := context.Background()
	projectID := util.SafeString(model.ProjectId)
	processorName := util.SafeString(model.ProcessorName)

	apiResp, err := client.AtlasSDK.StreamsApi.DeleteStreamProcessor(ctx, projectID, workspaceOrInstanceName, processorName).Execute()
	if err != nil {
		if apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Resource not found",
				HandlerErrorCode: "NotFound",
			}
		}
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         fmt.Sprintf("Error deleting stream processor: %s", err.Error()),
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
	}
}

func HandleList(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	workspaceOrInstanceName, err := GetWorkspaceOrInstanceName(model)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
		}
	}

	ctx := context.Background()
	projectID := util.SafeString(model.ProjectId)

	accumulatedProcessors, apiResp, err := getAllStreamProcessors(ctx, client.AtlasSDK, projectID, workspaceOrInstanceName)
	if err != nil {
		return handleError(apiResp, constants.LIST, err)
	}

	response := make([]interface{}, 0, len(accumulatedProcessors))
	for i := range accumulatedProcessors {
		modelItem, err := GetStreamProcessorModel(&accumulatedProcessors[i], model)
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         fmt.Sprintf("Error converting stream processor model: %s", err.Error()),
			}
		}

		copyIdentifyingFields(modelItem, model)
		response = append(response, modelItem)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModels:  response,
	}
}
