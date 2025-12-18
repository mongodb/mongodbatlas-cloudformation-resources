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

	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"

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

	// State constants
	StateIdle            = "IDLE"
	StateWorking         = "WORKING"
	StateDone            = "DONE"
	StateFailed          = "FAILED"
	StateDeleteRequested = "DELETE_REQUESTED"
	StateDeleting        = "DELETING"
	StateDeleted         = "DELETED"

	CallbackDelaySeconds = 30
)

var (
	CreateRequiredFields = []string{constants.ProjectID, "ProviderName", "Vendor"}
	ReadRequiredFields   = []string{constants.ProjectID, constants.ID}
	DeleteRequiredFields = []string{constants.ProjectID, constants.ID}
	ListRequiredFields   = []string{constants.ProjectID}
)

// initEnvWithLatestClient is a variable to allow mocking in tests
var initEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-stream-privatelink-endpoint")

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if errEvent := validator.ValidateModel(requiredFields, currentModel); errEvent != nil {
		return nil, errEvent
	}

	// Validate vendor-specific requirements
	if errEvent := validateVendorRequirements(currentModel); errEvent != nil {
		return nil, errEvent
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return nil, peErr
	}
	return client.AtlasSDK, nil
}

func validateVendorRequirements(model *Model) *handler.ProgressEvent {
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
		// ServiceEndpointId is required for AWS CONFLUENT (GCP uses service_attachment_uris which we don't support)
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
		// Region cannot be set for MSK (it's computed from ARN)
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
	// Check if this is a callback from a previous Create operation
	if req.CallbackContext != nil {
		return handleCreateCallback(req, currentModel)
	}

	conn, peErr := initEnvWithLatestClient(req, currentModel, CreateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()
	projectID := util.SafeString(currentModel.ProjectId)
	streamPrivatelinkEndpointReq := NewStreamPrivatelinkEndpointReq(currentModel)

	streamsPrivateLinkConnection, apiResp, err := conn.StreamsApi.CreatePrivateLinkConnection(ctx, projectID, streamPrivatelinkEndpointReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}

	// Set the ID in the model for state tracking
	currentModel.Id = streamsPrivateLinkConnection.Id

	// Wait for state transition using callback mechanism
	return waitForStateTransition(ctx, conn, currentModel, []string{StateIdle, StateWorking}, []string{StateDone, StateFailed}, false)
}

func handleCreateCallback(req handler.Request, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()
	return waitForStateTransition(ctx, conn.AtlasSDK, currentModel, []string{StateIdle, StateWorking}, []string{StateDone, StateFailed}, false)
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := initEnvWithLatestClient(req, currentModel, ReadRequiredFields)
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
		return handleError(apiResp, constants.READ, err)
	}

	resourceModel := GetStreamPrivatelinkEndpointModel(streamsPrivateLinkConnection, currentModel)
	resourceModel.ProjectId = currentModel.ProjectId

	// Handle empty DnsSubDomain array - preserve null state if it was null in currentModel
	if currentModel.DnsSubDomain == nil && (resourceModel.DnsSubDomain == nil || len(resourceModel.DnsSubDomain) == 0) {
		resourceModel.DnsSubDomain = nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   resourceModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Update is not supported for stream private link endpoints
	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          "Updating the private endpoint for streams is not supported. To modify your infrastructure, please delete the existing resource and create a new one with the necessary updates",
		HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Check if this is a callback from a previous Delete operation
	if req.CallbackContext != nil {
		return handleDeleteCallback(req, currentModel)
	}

	conn, peErr := initEnvWithLatestClient(req, currentModel, DeleteRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()
	projectID := util.SafeString(currentModel.ProjectId)
	connectionID := util.SafeString(currentModel.Id)

	_, err := conn.StreamsApi.DeletePrivateLinkConnection(ctx, projectID, connectionID).Execute()
	if err != nil {
		// DeletePrivateLinkConnection doesn't return a response body, so we can't check for 404 here.
		// If the resource is already deleted, the waitForStateTransition will handle it via GetPrivateLinkConnection.
		return handleError(nil, constants.DELETE, err)
	}

	// Wait for deletion state transition using callback mechanism
	return waitForStateTransition(ctx, conn, currentModel, []string{StateDeleteRequested, StateDeleting}, []string{StateDeleted, StateFailed}, true)
}

func handleDeleteCallback(req handler.Request, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()
	return waitForStateTransition(ctx, conn.AtlasSDK, currentModel, []string{StateDeleteRequested, StateDeleting}, []string{StateDeleted, StateFailed}, true)
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := initEnvWithLatestClient(req, currentModel, ListRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()
	projectID := util.SafeString(currentModel.ProjectId)

	accumulatedConnections, apiResp, err := getAllPrivateLinkConnections(ctx, conn, projectID)
	if err != nil {
		return handleError(apiResp, constants.LIST, err)
	}

	response := make([]interface{}, 0)
	for i := range accumulatedConnections {
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

func getAllPrivateLinkConnections(ctx context.Context, conn *admin20250312010.APIClient, projectID string) ([]admin20250312010.StreamsPrivateLinkConnection, *http.Response, error) {
	pageNum := 1
	accumulatedConnections := make([]admin20250312010.StreamsPrivateLinkConnection, 0)

	for allRecordsRetrieved := false; !allRecordsRetrieved; {
		connections, apiResp, err := conn.StreamsApi.ListPrivateLinkConnectionsWithParams(ctx, &admin20250312010.ListPrivateLinkConnectionsApiParams{
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

func waitForStateTransition(ctx context.Context, conn *admin20250312010.APIClient, currentModel *Model, pendingStates, targetStates []string, isDelete bool) (handler.ProgressEvent, error) {
	projectID := util.SafeString(currentModel.ProjectId)
	connectionID := util.SafeString(currentModel.Id)

	// Get current state
	streamsPrivateLinkConnection, apiResp, err := conn.StreamsApi.GetPrivateLinkConnection(ctx, projectID, connectionID).Execute()
	if err != nil {
		// For delete operations, 404 means already deleted
		if isDelete && apiResp != nil && apiResp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Resource deleted",
			}, nil
		}
		return handleError(apiResp, constants.DELETE, err)
	}

	currentState := ""
	if streamsPrivateLinkConnection.State != nil {
		currentState = *streamsPrivateLinkConnection.State
	}

	// Check if we've reached a target state
	for _, targetState := range targetStates {
		if currentState == targetState {
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

			// For delete operations, don't return the model
			if isDelete {
				return handler.ProgressEvent{
					OperationStatus: handler.Success,
					Message:         "Delete completed",
				}, nil
			}

			// For create operations, return the model
			resourceModel := GetStreamPrivatelinkEndpointModel(streamsPrivateLinkConnection, currentModel)
			resourceModel.ProjectId = currentModel.ProjectId

			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Create completed",
				ResourceModel:   resourceModel,
			}, nil
		}
	}

	// Check if we're still in a pending state
	for _, pendingState := range pendingStates {
		if currentState == pendingState {
			// Return InProgress to trigger callback
			return handler.ProgressEvent{
				OperationStatus:      handler.InProgress,
				Message:              fmt.Sprintf("Waiting for state transition. Current state: %s", currentState),
				ResourceModel:        currentModel,
				CallbackDelaySeconds: CallbackDelaySeconds,
				CallbackContext: map[string]interface{}{
					"stateName": currentState,
					"id":        connectionID,
				},
			}, nil
		}
	}

	// If we're in an unexpected state, fail
	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          fmt.Sprintf("Unexpected state: %s", currentState),
		HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
	}, nil
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error: %s", method, err.Error())

	// Check for specific error conditions
	if response != nil {
		if response.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          constants.ResourceNotFound,
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}, nil
		}
	}

	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}
