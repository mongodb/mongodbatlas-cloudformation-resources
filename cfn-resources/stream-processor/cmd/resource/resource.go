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
	"fmt"
	"maps"
	"net/http"
	"time"

	"go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
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
	DefaultCallbackDelaySeconds = 3
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

func CopyIdentifyingFields(resourceModel, currentModel *Model) {
	resourceModel.Profile = currentModel.Profile
	resourceModel.ProjectId = currentModel.ProjectId
	resourceModel.ProcessorName = currentModel.ProcessorName

	switch {
	case currentModel.WorkspaceName != nil && *currentModel.WorkspaceName != "":
		resourceModel.WorkspaceName = currentModel.WorkspaceName
		resourceModel.InstanceName = util.Pointer(*currentModel.WorkspaceName)
	case currentModel.InstanceName != nil && *currentModel.InstanceName != "":
		resourceModel.InstanceName = currentModel.InstanceName
		resourceModel.WorkspaceName = util.Pointer(*currentModel.InstanceName)
	default:
		resourceModel.WorkspaceName = currentModel.WorkspaceName
		resourceModel.InstanceName = currentModel.InstanceName
	}
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

func ParseTimeout(timeoutStr string) time.Duration {
	if timeoutStr == "" {
		return DefaultCreateTimeout
	}
	duration, err := time.ParseDuration(timeoutStr)
	if err != nil {
		_, _ = logger.Warnf("Invalid timeout format '%s', using default: %v", timeoutStr, err)
		return DefaultCreateTimeout
	}
	return duration
}

func IsTimeoutExceeded(startTimeStr, timeoutDurationStr string) bool {
	if startTimeStr == "" || timeoutDurationStr == "" {
		return false
	}

	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		_, _ = logger.Warnf("Invalid start time format '%s': %v", startTimeStr, err)
		return false
	}

	timeoutDuration := ParseTimeout(timeoutDurationStr)
	elapsed := time.Since(startTime)

	return elapsed >= timeoutDuration
}

func CleanupOnCreateTimeout(ctx context.Context, atlasClient *admin.APIClient, callbackCtx *CallbackData) error {
	if !callbackCtx.DeleteOnCreateTimeout {
		return nil
	}

	_, err := atlasClient.StreamsApi.DeleteStreamProcessor(ctx, callbackCtx.ProjectID, callbackCtx.WorkspaceOrInstanceName, callbackCtx.ProcessorName).Execute()
	if err != nil {
		_, _ = logger.Warnf("Cleanup delete failed: %v", err)
	}
	return nil
}

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
		CallbackDelaySeconds: DefaultCallbackDelaySeconds,
		CallbackContext: BuildCallbackContext(projectID, workspaceOrInstanceName, processorName, map[string]any{
			"needsStarting":         needsStarting,
			"startTime":             time.Now().Format(time.RFC3339),
			"timeoutDuration":       timeoutStr,
			"deleteOnCreateTimeout": deleteOnCreateTimeout,
		}),
	}, nil
}

func HandleCreateCallback(ctx context.Context, atlasClient *admin.APIClient, currentModel *Model, callbackCtx *CallbackData) (handler.ProgressEvent, error) {
	needsStarting := callbackCtx.NeedsStarting

	if IsTimeoutExceeded(callbackCtx.StartTime, callbackCtx.TimeoutDuration) {
		if err := CleanupOnCreateTimeout(context.Background(), atlasClient, callbackCtx); err != nil {
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

	streamProcessor, peErr := GetStreamProcessor(ctx, atlasClient, callbackCtx.ProjectID, callbackCtx.WorkspaceOrInstanceName, callbackCtx.ProcessorName)
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
			if peErr := StartStreamProcessor(ctx, atlasClient, callbackCtx.ProjectID, callbackCtx.WorkspaceOrInstanceName, callbackCtx.ProcessorName); peErr != nil {
				return *peErr, nil
			}
			return CreateInProgressEvent("Starting stream processor", currentModel, callbackContext), nil
		}
		return FinalizeModel(streamProcessor, currentModel, "Create Completed")

	case StartedState:
		return FinalizeModel(streamProcessor, currentModel, "Create Completed")

	case InitiatingState, CreatingState:
		return CreateInProgressEvent(fmt.Sprintf("Creating stream processor (current state: %s)", currentState), currentModel, callbackContext), nil

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

func FinalizeModel(streamProcessor *admin.StreamsProcessorWithStats, currentModel *Model, message string) (handler.ProgressEvent, error) {
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
		Message:         message,
		ResourceModel:   resourceModel,
	}, nil
}

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
			CallbackDelaySeconds: DefaultCallbackDelaySeconds,
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
			CallbackDelaySeconds: DefaultCallbackDelaySeconds,
			CallbackContext: BuildCallbackContext(projectID, workspaceOrInstanceName, processorName, map[string]any{
				"desiredState": desiredState,
			}),
		}, nil
	}

	return FinalizeModel(streamProcessorResp, currentModel, "Update Completed")
}

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

func getAllStreamProcessors(ctx context.Context, atlasClient *admin.APIClient, projectID, workspaceOrInstanceName string) ([]admin.StreamsProcessorWithStats, *http.Response, error) {
	pageNum := 1
	accumulatedProcessors := make([]admin.StreamsProcessorWithStats, 0)

	for allRecordsRetrieved := false; !allRecordsRetrieved; {
		processorsResp, apiResp, err := atlasClient.StreamsApi.GetStreamProcessorsWithParams(ctx, &admin.GetStreamProcessorsApiParams{
			GroupId:      projectID,
			TenantName:   workspaceOrInstanceName,
			ItemsPerPage: util.Pointer(constants.DefaultListItemsPerPage),
			PageNum:      util.Pointer(pageNum),
		}).Execute()

		if err != nil {
			return nil, apiResp, err
		}

		results := processorsResp.GetResults()
		accumulatedProcessors = append(accumulatedProcessors, results...)

		totalCount := processorsResp.GetTotalCount()
		allRecordsRetrieved = totalCount <= len(accumulatedProcessors) || len(results) < constants.DefaultListItemsPerPage
		pageNum++
	}

	return accumulatedProcessors, nil, nil
}

func HandleUpdateCallback(ctx context.Context, atlasClient *admin.APIClient, currentModel *Model, callbackCtx *CallbackData) (handler.ProgressEvent, error) {
	streamProcessor, peErr := GetStreamProcessor(ctx, atlasClient, callbackCtx.ProjectID, callbackCtx.WorkspaceOrInstanceName, callbackCtx.ProcessorName)
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
			if peErr := StartStreamProcessor(ctx, atlasClient, callbackCtx.ProjectID, callbackCtx.WorkspaceOrInstanceName, callbackCtx.ProcessorName); peErr != nil {
				return *peErr, nil
			}
			return CreateInProgressEvent("Starting stream processor", currentModel, callbackContext), nil
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
		return CreateInProgressEvent("Stopping stream processor", currentModel, callbackContext), nil

	case FailedState:
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Stream processor entered FAILED state",
		}, nil

	default:
		return CreateInProgressEvent(fmt.Sprintf("Updating stream processor (current state: %s)", currentState), currentModel, callbackContext), nil
	}
}

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

func GetStreamProcessor(ctx context.Context, atlasClient *admin.APIClient, projectID, workspaceOrInstanceName, processorName string) (*admin.StreamsProcessorWithStats, *handler.ProgressEvent) {
	requestParams := &admin.GetStreamProcessorApiParams{
		GroupId:       projectID,
		TenantName:    workspaceOrInstanceName,
		ProcessorName: processorName,
	}

	streamProcessor, resp, err := atlasClient.StreamsApi.GetStreamProcessorWithParams(ctx, requestParams).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil, &handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Stream processor not found",
				HandlerErrorCode: "NotFound",
			}
		}
		return nil, &handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         fmt.Sprintf("Error getting stream processor: %s", err.Error()),
		}
	}
	return streamProcessor, nil
}

func StartStreamProcessor(ctx context.Context, atlasClient *admin.APIClient, projectID, workspaceOrInstanceName, processorName string) *handler.ProgressEvent {
	_, err := atlasClient.StreamsApi.StartStreamProcessorWithParams(ctx,
		&admin.StartStreamProcessorApiParams{
			GroupId:       projectID,
			TenantName:    workspaceOrInstanceName,
			ProcessorName: processorName,
		},
	).Execute()
	if err != nil {
		return &handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         fmt.Sprintf("Error starting stream processor: %s", err.Error()),
		}
	}
	return nil
}

func CreateInProgressEvent(message string, currentModel *Model, callbackContext map[string]any) handler.ProgressEvent {
	inProgressModel := &Model{}
	if currentModel != nil {
		*inProgressModel = *currentModel
		inProgressModel.DeleteOnCreateTimeout = nil
	}
	CopyIdentifyingFields(inProgressModel, currentModel)

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              message,
		ResourceModel:        inProgressModel,
		CallbackDelaySeconds: DefaultCallbackDelaySeconds,
		CallbackContext:      callbackContext,
	}
}

func ValidateUpdateStateTransition(currentState, desiredState string) (errMsg string, isValidTransition bool) {
	if currentState == desiredState {
		return "", true
	}

	if desiredState == StoppedState && currentState != StartedState {
		return fmt.Sprintf("Stream Processor must be in %s state to transition to %s state", StartedState, StoppedState), false
	}

	if desiredState == CreatedState {
		return fmt.Sprintf("Stream Processor cannot transition from %s to CREATED", currentState), false
	}

	return "", true
}

func HandleError(response *http.Response, method constants.CfnFunctions, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())

	if response != nil && response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: "AlreadyExists",
		}, nil
	}

	return progressevent.GetFailedEventByResponse(errMsg, response), nil
}
