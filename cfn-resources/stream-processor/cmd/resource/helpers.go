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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

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

func getStreamProcessor(ctx context.Context, atlasClient *admin.APIClient, projectID, workspaceOrInstanceName, processorName string) (*admin.StreamsProcessorWithStats, *handler.ProgressEvent) {
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

func startStreamProcessor(ctx context.Context, atlasClient *admin.APIClient, projectID, workspaceOrInstanceName, processorName string) *handler.ProgressEvent {
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

func createInProgressEvent(message string, currentModel *Model, callbackContext map[string]any) handler.ProgressEvent {
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
		CallbackDelaySeconds: defaultCallbackDelaySeconds,
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
