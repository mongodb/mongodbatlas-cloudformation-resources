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
	"slices"
	"strings"

	"go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

func ValidateVendorRequirements(model *Model) *handler.ProgressEvent {
	if model.Vendor == nil {
		return nil
	}

	vendor := *model.Vendor

	switch vendor {
	case VendorConfluent:
		if model.DnsDomain == nil || *model.DnsDomain == "" {
			return &handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          fmt.Sprintf("DnsDomain is required for vendor %s", VendorConfluent),
				HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
			}
		}
		if model.Region == nil || *model.Region == "" {
			return &handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          fmt.Sprintf("Region is required for vendor %s", VendorConfluent),
				HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
			}
		}
		if model.ServiceEndpointId == nil || *model.ServiceEndpointId == "" {
			return &handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          fmt.Sprintf("ServiceEndpointId is required for vendor %s with AWS provider", VendorConfluent),
				HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
			}
		}

	case VendorMSK:
		if model.Arn == nil || *model.Arn == "" {
			return &handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          fmt.Sprintf("Arn is required for vendor %s", VendorMSK),
				HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
			}
		}
		if model.Region != nil && *model.Region != "" {
			return &handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          fmt.Sprintf("Region cannot be set for vendor %s (it is computed from ARN)", VendorMSK),
				HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
			}
		}

	case VendorS3:
		if model.Region == nil || *model.Region == "" {
			return &handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          fmt.Sprintf("Region is required for vendor %s", VendorS3),
				HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
			}
		}
		if model.ServiceEndpointId == nil || *model.ServiceEndpointId == "" {
			return &handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          fmt.Sprintf("ServiceEndpointId is required for vendor %s. It should follow the format 'com.amazonaws.<region>.s3', for example 'com.amazonaws.us-east-1.s3'", VendorS3),
				HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
			}
		}
	}

	return nil
}

var GetAllPrivateLinkConnections = func(ctx context.Context, conn *admin.APIClient, projectID string) ([]admin.StreamsPrivateLinkConnection, *http.Response, error) {
	pageNum := 1
	accumulatedConnections := make([]admin.StreamsPrivateLinkConnection, 0)

	for allRecordsRetrieved := false; !allRecordsRetrieved; {
		connections, apiResp, err := conn.StreamsApi.ListPrivateLinkConnectionsWithParams(ctx, &admin.ListPrivateLinkConnectionsApiParams{
			GroupId:      projectID,
			ItemsPerPage: util.Pointer(constants.DefaultListItemsPerPage),
			PageNum:      util.Pointer(pageNum),
		}).Execute()

		if err != nil {
			return nil, apiResp, err
		}
		accumulatedConnections = append(accumulatedConnections, connections.GetResults()...)
		allRecordsRetrieved = connections.GetTotalCount() <= len(accumulatedConnections)
		pageNum++
	}

	return accumulatedConnections, nil, nil
}

func WaitForStateTransition(ctx context.Context, conn *admin.APIClient, currentModel *Model, pendingStates, targetStates []string, isDelete bool) (handler.ProgressEvent, error) {
	projectID := util.SafeString(currentModel.ProjectId)
	connectionID := util.SafeString(currentModel.Id)

	streamsPrivateLinkConnection, apiResp, err := conn.StreamsApi.GetPrivateLinkConnection(ctx, projectID, connectionID).Execute()
	if err != nil {
		if isDelete && apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Resource deleted",
			}, nil
		}
		if !isDelete && apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:      handler.InProgress,
				Message:              "Resource not yet available, retrying...",
				ResourceModel:        currentModel,
				CallbackDelaySeconds: CallbackDelaySeconds,
				CallbackContext:      BuildCallbackContext(projectID, connectionID),
			}, nil
		}
		if isDelete {
			if apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
				return handler.ProgressEvent{
					OperationStatus: handler.Success,
					Message:         "Resource already deleted",
				}, nil
			}
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          constants.ResourceNotFound,
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}, nil
		}
		return HandleError(apiResp, constants.CREATE, err)
	}

	currentState := ""
	if streamsPrivateLinkConnection.State != nil {
		currentState = *streamsPrivateLinkConnection.State
	}

	if currentState == "" {
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              "Resource state not yet available, waiting...",
			ResourceModel:        currentModel,
			CallbackDelaySeconds: CallbackDelaySeconds,
			CallbackContext:      BuildCallbackContext(projectID, connectionID),
		}, nil
	}

	for _, targetState := range targetStates {
		if currentState != targetState {
			continue
		}

		if currentState == StateFailed {
			errorMsg := "Private endpoint is in a failed status"
			if streamsPrivateLinkConnection.ErrorMessage != nil {
				errorMsg = fmt.Sprintf("%s: %s", errorMsg, *streamsPrivateLinkConnection.ErrorMessage)
			}
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          errorMsg,
				HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
			}, nil
		}

		if isDelete {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Delete completed",
			}, nil
		}

		resourceModel := GetStreamPrivatelinkEndpointModel(streamsPrivateLinkConnection, currentModel)
		resourceModel.ProjectId = currentModel.ProjectId

		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Create completed",
			ResourceModel:   resourceModel,
		}, nil
	}

	if slices.Contains(pendingStates, currentState) {
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              fmt.Sprintf("Waiting for state transition. Current state: %s", currentState),
			ResourceModel:        currentModel,
			CallbackDelaySeconds: CallbackDelaySeconds,
			CallbackContext:      BuildCallbackContext(projectID, connectionID),
		}, nil
	}

	if isDelete {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          constants.ResourceNotFound,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
		}, nil
	}
	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          fmt.Sprintf("Unexpected state: %s", currentState),
		HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
	}, nil
}

func HandleError(response *http.Response, method constants.CfnFunctions, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error: %s", method, err.Error())

	if response != nil {
		if response.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          constants.ResourceNotFound,
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}, nil
		}
		if response.StatusCode == http.StatusBadRequest {
			errStr := err.Error()
			if strings.Contains(errStr, "STREAM_PRIVATE_LINK_ALREADY_EXISTS") || strings.Contains(errStr, "already exists") {
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          errMsg,
					HandlerErrorCode: string(types.HandlerErrorCodeAlreadyExists),
				}, nil
			}
		}
	}

	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}
