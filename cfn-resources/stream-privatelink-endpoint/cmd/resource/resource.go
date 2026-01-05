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
	"net/http"
	"slices"
	"strings"

	"go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

const (
	VendorConfluent = "CONFLUENT"
	VendorMSK       = "MSK"
	VendorS3        = "S3"

	StateIdle            = "IDLE"
	StateWorking         = "WORKING"
	StateDone            = "DONE"
	StateFailed          = "FAILED"
	StateDeleteRequested = "DELETE_REQUESTED"
	StateDeleting        = "DELETING"
	StateDeleted         = "DELETED"

	CallbackDelaySeconds = 3
)

var (
	CreateRequiredFields = []string{constants.ProjectID, "ProviderName", "Vendor"}
	ReadRequiredFields   = []string{constants.ProjectID, constants.ID}
	DeleteRequiredFields = []string{constants.ProjectID, constants.ID}
	ListRequiredFields   = []string{constants.ProjectID}
)

var InitEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin.APIClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-stream-privatelink-endpoint")

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if errEvent := validator.ValidateModel(requiredFields, currentModel); errEvent != nil {
		return nil, errEvent
	}

	if errEvent := ValidateVendorRequirements(currentModel); errEvent != nil {
		return nil, errEvent
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return nil, peErr
	}
	return client.AtlasSDK, nil
}

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

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if IsCallback(&req) {
		conn, peErr := InitEnvWithLatestClient(req, currentModel, CreateRequiredFields)
		if peErr != nil {
			return *peErr, nil
		}
		return HandleCreateCallback(context.Background(), conn, req, currentModel)
	}

	conn, peErr := InitEnvWithLatestClient(req, currentModel, CreateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()
	projectID := util.SafeString(currentModel.ProjectId)
	streamPrivatelinkEndpointReq := NewStreamPrivatelinkEndpointReq(currentModel)

	streamsPrivateLinkConnection, apiResp, err := conn.StreamsApi.CreatePrivateLinkConnection(ctx, projectID, streamPrivatelinkEndpointReq).Execute()
	if err != nil {
		return HandleError(apiResp, constants.CREATE, err)
	}

	currentModel.Id = streamsPrivateLinkConnection.Id

	if streamsPrivateLinkConnection.State != nil {
		initialState := *streamsPrivateLinkConnection.State
		if initialState == StateDone {
			resourceModel := GetStreamPrivatelinkEndpointModel(streamsPrivateLinkConnection, currentModel)
			resourceModel.ProjectId = currentModel.ProjectId
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Create completed",
				ResourceModel:   resourceModel,
			}, nil
		}
		if initialState == StateFailed {
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
	}

	callbackCtx := BuildCallbackContext(projectID, util.SafeString(currentModel.Id))
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              "Creating private link endpoint",
		ResourceModel:        currentModel,
		CallbackDelaySeconds: CallbackDelaySeconds,
		CallbackContext:      callbackCtx,
	}, nil
}

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

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := InitEnvWithLatestClient(req, currentModel, ReadRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()
	projectID := util.SafeString(currentModel.ProjectId)
	connectionID := util.SafeString(currentModel.Id)

	streamsPrivateLinkConnection, apiResp, err := conn.StreamsApi.GetPrivateLinkConnection(ctx, projectID, connectionID).Execute()
	if err != nil {
		if apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          constants.ResourceNotFound,
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}, nil
		}
		return HandleError(apiResp, constants.READ, err)
	}

	resourceModel := GetStreamPrivatelinkEndpointModel(streamsPrivateLinkConnection, currentModel)
	resourceModel.ProjectId = currentModel.ProjectId

	if currentModel.DnsSubDomain == nil && len(resourceModel.DnsSubDomain) == 0 {
		resourceModel.DnsSubDomain = nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   resourceModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          "Updating the private endpoint for streams is not supported. To modify your infrastructure, please delete the existing resource and create a new one with the necessary updates",
		HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	connectionID := util.SafeString(currentModel.Id)
	if connectionID == "" {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          constants.ResourceNotFound,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
		}, nil
	}

	if req.CallbackContext != nil {
		if _, ok := req.CallbackContext["callbackStreamPrivatelinkEndpoint"]; ok {
			conn, peErr := InitEnvWithLatestClient(req, currentModel, DeleteRequiredFields)
			if peErr != nil {
				return *peErr, nil
			}
			return HandleDeleteCallback(context.Background(), conn, req, currentModel)
		}
	}

	conn, peErr := InitEnvWithLatestClient(req, currentModel, DeleteRequiredFields)
	if peErr != nil {
		if peErr.HandlerErrorCode == string(types.HandlerErrorCodeInvalidRequest) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          constants.ResourceNotFound,
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}, nil
		}
		return *peErr, nil
	}

	ctx := context.Background()
	projectID := util.SafeString(currentModel.ProjectId)

	streamsPrivateLinkConnection, apiResp, err := conn.StreamsApi.GetPrivateLinkConnection(ctx, projectID, connectionID).Execute()
	if err != nil {
		if apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          constants.ResourceNotFound,
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}, nil
		}
		if apiResp != nil && apiResp.StatusCode >= http.StatusInternalServerError {
			return HandleError(apiResp, constants.DELETE, err)
		}
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          constants.ResourceNotFound,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
		}, nil
	}

	if streamsPrivateLinkConnection.State != nil {
		if *streamsPrivateLinkConnection.State == StateDeleted {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          constants.ResourceNotFound,
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}, nil
		}
	}

	deleteResp, err := conn.StreamsApi.DeletePrivateLinkConnection(ctx, projectID, connectionID).Execute()
	if err != nil {
		if deleteResp != nil && deleteResp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          constants.ResourceNotFound,
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}, nil
		}
		if deleteResp != nil && deleteResp.StatusCode >= http.StatusInternalServerError {
			return HandleError(deleteResp, constants.DELETE, err)
		}
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          constants.ResourceNotFound,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
		}, nil
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              "Deleting private link endpoint",
		ResourceModel:        currentModel,
		CallbackDelaySeconds: CallbackDelaySeconds,
		CallbackContext:      BuildCallbackContext(projectID, connectionID),
	}, nil
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

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := InitEnvWithLatestClient(req, currentModel, ListRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()
	projectID := util.SafeString(currentModel.ProjectId)

	accumulatedConnections, apiResp, err := GetAllPrivateLinkConnections(ctx, conn, projectID)
	if err != nil {
		return HandleError(apiResp, constants.LIST, err)
	}

	response := make([]interface{}, 0)
	for i := range accumulatedConnections {
		connectionState := ""
		if accumulatedConnections[i].State != nil {
			connectionState = *accumulatedConnections[i].State
		}
		if connectionState == StateDeleted || connectionState == StateDeleteRequested || connectionState == StateDeleting {
			continue
		}

		model := GetStreamPrivatelinkEndpointModel(&accumulatedConnections[i], nil)
		model.ProjectId = currentModel.ProjectId
		model.Profile = currentModel.Profile

		response = append(response, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModels:  response,
	}, nil
}

func GetAllPrivateLinkConnections(ctx context.Context, conn *admin.APIClient, projectID string) ([]admin.StreamsPrivateLinkConnection, *http.Response, error) {
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
