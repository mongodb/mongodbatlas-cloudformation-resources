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

	"go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
)

func IsCallback(req *handler.Request) bool {
	if req.CallbackContext == nil {
		return false
	}
	_, found := req.CallbackContext["callbackStreamPrivatelinkEndpoint"]
	return found
}

func BuildCallbackContext(projectID, connectionID string) map[string]interface{} {
	return map[string]interface{}{
		"callbackStreamPrivatelinkEndpoint": true,
		"projectID":                         projectID,
		"connectionID":                      connectionID,
	}
}

func HandleCreateCallback(ctx context.Context, atlasClient *admin.APIClient, req handler.Request, currentModel *Model) (handler.ProgressEvent, error) {
	projectID := util.SafeString(currentModel.ProjectId)
	if projectID == "" {
		if pid, ok := req.CallbackContext["projectID"].(string); ok {
			projectID = pid
			if currentModel.ProjectId == nil {
				currentModel.ProjectId = &pid
			}
		}
	}

	connectionID := util.SafeString(currentModel.Id)
	if connectionID == "" {
		if id, ok := req.CallbackContext["connectionID"].(string); ok {
			connectionID = id
			currentModel.Id = &id
		} else {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "Missing connection ID in callback context",
			}, nil
		}
	}

	streamsPrivateLinkConnection, apiResp, err := atlasClient.StreamsApi.GetPrivateLinkConnection(ctx, projectID, connectionID).Execute()
	if err != nil {
		if apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:      handler.InProgress,
				Message:              "Resource not yet available, retrying...",
				ResourceModel:        currentModel,
				CallbackDelaySeconds: CallbackDelaySeconds,
				CallbackContext:      BuildCallbackContext(projectID, connectionID),
			}, nil
		}
		return HandleError(apiResp, constants.CREATE, err)
	}

	currentState := ""
	if streamsPrivateLinkConnection.State != nil {
		currentState = *streamsPrivateLinkConnection.State
	}

	switch currentState {
	case StateDone:
		resourceModel := GetStreamPrivatelinkEndpointModel(streamsPrivateLinkConnection, currentModel)
		resourceModel.ProjectId = currentModel.ProjectId
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Create completed",
			ResourceModel:   resourceModel,
		}, nil

	case StateFailed:
		errorMsg := "Private endpoint is in a failed status"
		if streamsPrivateLinkConnection.ErrorMessage != nil {
			errorMsg = fmt.Sprintf("%s: %s", errorMsg, *streamsPrivateLinkConnection.ErrorMessage)
		}
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errorMsg,
			HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
		}, nil

	case StateIdle, StateWorking, "":
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              fmt.Sprintf("Waiting for state transition. Current state: %s", currentState),
			ResourceModel:        currentModel,
			CallbackDelaySeconds: CallbackDelaySeconds,
			CallbackContext:      BuildCallbackContext(projectID, connectionID),
		}, nil

	default:
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              fmt.Sprintf("Waiting for state transition. Current state: %s", currentState),
			ResourceModel:        currentModel,
			CallbackDelaySeconds: CallbackDelaySeconds,
			CallbackContext:      BuildCallbackContext(projectID, connectionID),
		}, nil
	}
}

func HandleDeleteCallback(ctx context.Context, atlasClient *admin.APIClient, req handler.Request, currentModel *Model) (handler.ProgressEvent, error) {
	projectID := util.SafeString(currentModel.ProjectId)
	connectionID := util.SafeString(currentModel.Id)

	if projectID == "" {
		if pid, ok := req.CallbackContext["projectID"].(string); ok {
			projectID = pid
		}
	}
	if connectionID == "" {
		if id, ok := req.CallbackContext["connectionID"].(string); ok {
			connectionID = id
			currentModel.Id = &id
		}
	}

	if projectID == "" || connectionID == "" {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          constants.ResourceNotFound,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
		}, nil
	}

	return WaitForStateTransition(ctx, atlasClient, currentModel, []string{StateDeleteRequested, StateDeleting}, []string{StateDeleted, StateFailed}, true)
}
