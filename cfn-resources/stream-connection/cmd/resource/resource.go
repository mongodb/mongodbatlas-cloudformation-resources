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

	"go.mongodb.org/atlas-sdk/v20250312014/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

const (
	ClusterConnectionType = "Cluster"
	KafkaConnectionType   = "Kafka"
	AWSLambdaType         = "AWSLambda"
	HTTPSType             = "Https"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.ConnectionName, constants.Type}
var ReadRequiredFields = []string{constants.ProjectID, constants.ConnectionName}
var UpdateRequiredFields = []string{constants.ProjectID, constants.ConnectionName, constants.Type}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ConnectionName}
var ListRequiredFields = []string{constants.ProjectID}

func normalizeWorkspaceName(model *Model) *handler.ProgressEvent {
	var workspaceOrInstanceName *string

	// Validate that at least one of WorkspaceName or InstanceName is provided
	switch {
	case model.WorkspaceName != nil && *model.WorkspaceName != "":
		workspaceOrInstanceName = model.WorkspaceName
	case model.InstanceName != nil && *model.InstanceName != "":
		workspaceOrInstanceName = model.InstanceName
	default:
		return &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Either WorkspaceName or InstanceName must be provided",
			HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
		}
	}

	// Ensure both WorkspaceName and InstanceName are set for backward compatibility
	model.WorkspaceName = workspaceOrInstanceName
	model.InstanceName = workspaceOrInstanceName
	return nil
}

var InitEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin.APIClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-stream-connection")

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if errEvent := validator.ValidateModel(requiredFields, currentModel); errEvent != nil {
		return nil, errEvent
	}

	if peErr := normalizeWorkspaceName(currentModel); peErr != nil {
		return nil, peErr
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return nil, peErr
	}
	return client.AtlasSDK, nil
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := InitEnvWithLatestClient(req, currentModel, CreateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	workspaceOrInstanceName := currentModel.WorkspaceName

	ctx := context.Background()

	projectID := currentModel.ProjectId
	streamConnectionReq := newStreamConnectionReq(currentModel)

	streamConnResp, apiResp, err := conn.StreamsApi.CreateStreamConnection(ctx, *projectID, *workspaceOrInstanceName, streamConnectionReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}

	resourceModel := GetStreamConnectionModel(streamConnResp, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   resourceModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := InitEnvWithLatestClient(req, currentModel, ReadRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	workspaceOrInstanceName := currentModel.WorkspaceName

	projectID := currentModel.ProjectId
	connectionName := currentModel.ConnectionName
	streamConnResp, apiResp, err := conn.StreamsApi.GetStreamConnection(context.Background(), *projectID, *workspaceOrInstanceName, *connectionName).Execute()
	if err != nil {
		return handleError(apiResp, constants.READ, err)
	}

	resourceModel := GetStreamConnectionModel(streamConnResp, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   resourceModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := InitEnvWithLatestClient(req, currentModel, UpdateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	workspaceOrInstanceName := currentModel.WorkspaceName

	ctx := context.Background()

	projectID := currentModel.ProjectId
	connectionName := currentModel.ConnectionName
	streamConnectionReq := newStreamConnectionReq(currentModel)
	streamConnResp, apiResp, err := conn.StreamsApi.UpdateStreamConnection(ctx, *projectID, *workspaceOrInstanceName, *connectionName, streamConnectionReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.UPDATE, err)
	}

	resourceModel := GetStreamConnectionModel(streamConnResp, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   resourceModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := InitEnvWithLatestClient(req, currentModel, DeleteRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	workspaceOrInstanceName := currentModel.WorkspaceName

	ctx := context.Background()

	projectID := currentModel.ProjectId
	connectionName := currentModel.ConnectionName
	apiResp, err := conn.StreamsApi.DeleteStreamConnection(ctx, *projectID, *workspaceOrInstanceName, *connectionName).Execute()
	if err != nil {
		return handleError(apiResp, constants.DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
		ResourceModel:   nil}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := InitEnvWithLatestClient(req, currentModel, ListRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	workspaceOrInstanceName := currentModel.WorkspaceName

	ctx := context.Background()

	projectID := currentModel.ProjectId

	accumulatedStreamConns, apiResp, err := getAllStreamConnections(ctx, conn, *projectID, *workspaceOrInstanceName)
	if err != nil {
		return handleError(apiResp, constants.LIST, err)
	}

	response := make([]interface{}, 0)
	for i := range accumulatedStreamConns {
		model := GetStreamConnectionModel(&accumulatedStreamConns[i], nil)
		model.ProjectId = currentModel.ProjectId
		// InstanceName is deprecated but we maintain it for backward compatibility
		model.WorkspaceName = workspaceOrInstanceName
		model.InstanceName = workspaceOrInstanceName
		model.Profile = currentModel.Profile

		response = append(response, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModels:  response,
	}, nil
}

func getAllStreamConnections(ctx context.Context, conn *admin.APIClient, projectID, workspaceOrInstanceName string) ([]admin.StreamsConnection, *http.Response, error) {
	pageNum := 1
	accumulatedStreamConns := make([]admin.StreamsConnection, 0)

	for allRecordsRetrieved := false; !allRecordsRetrieved; {
		streamConns, apiResp, err := conn.StreamsApi.ListStreamConnectionsWithParams(ctx, &admin.ListStreamConnectionsApiParams{
			GroupId:      projectID,
			TenantName:   workspaceOrInstanceName,
			ItemsPerPage: util.Pointer(constants.DefaultListItemsPerPage),
			PageNum:      util.Pointer(pageNum),
		}).Execute()

		if err != nil {
			return nil, apiResp, err
		}
		accumulatedStreamConns = append(accumulatedStreamConns, streamConns.GetResults()...)
		allRecordsRetrieved = streamConns.GetTotalCount() <= len(accumulatedStreamConns)
		pageNum++
	}

	return accumulatedStreamConns, nil, nil
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())

	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}
