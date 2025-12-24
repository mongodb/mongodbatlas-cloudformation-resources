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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20250312010/admin"
)

var CreateRequiredFields = []string{"WorkspaceName", constants.ProjectID, constants.DataProcessRegion}
var ReadRequiredFields = []string{"WorkspaceName", constants.ProjectID}
var UpdateRequiredFields = []string{"WorkspaceName", constants.ProjectID, constants.DataProcessRegion}
var DeleteRequiredFields = []string{"WorkspaceName", constants.ProjectID}
var ListRequiredFields = []string{constants.ProjectID}

const Kafka = "Kafka"
const Cluster = "Cluster"
const DefaultItemsPerPage = 100

var InitEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin.APIClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-stream-workspace")

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

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := InitEnvWithLatestClient(req, currentModel, CreateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	streamWorkspaceCreateReq := NewStreamWorkspaceCreateReq(currentModel)

	createdStreamWorkspace, resp, err := conn.StreamsApi.CreateStreamWorkspace(ctx, *currentModel.ProjectId, streamWorkspaceCreateReq).Execute()
	if err != nil {
		return HandleError(resp, constants.CREATE, err)
	}

	model := GetStreamWorkspaceModel(createdStreamWorkspace, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   model,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := InitEnvWithLatestClient(req, currentModel, ReadRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	streamWorkspace, resp, err := conn.StreamsApi.GetStreamWorkspace(ctx, *currentModel.ProjectId, *currentModel.WorkspaceName).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "StreamWorkspace not found",
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}, nil
		}
		return HandleError(resp, constants.READ, err)
	}

	model := GetStreamWorkspaceModel(streamWorkspace, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   model,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := InitEnvWithLatestClient(req, currentModel, UpdateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	streamWorkspaceUpdateReq := NewStreamWorkspaceUpdateReq(currentModel)
	updatedStreamWorkspace, resp, err := conn.StreamsApi.UpdateStreamWorkspace(ctx, *currentModel.ProjectId, *currentModel.WorkspaceName, streamWorkspaceUpdateReq).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "StreamWorkspace not found",
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}, nil
		}
		return HandleError(resp, constants.UPDATE, err)
	}

	model := GetStreamWorkspaceModel(updatedStreamWorkspace, currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   model,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := InitEnvWithLatestClient(req, currentModel, DeleteRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	resp, err := conn.StreamsApi.DeleteStreamWorkspace(ctx, *currentModel.ProjectId, *currentModel.WorkspaceName).Execute()
	if err != nil {
		return HandleError(resp, constants.DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := InitEnvWithLatestClient(req, currentModel, ListRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	accumulatedStreamWorkspaces, apiResp, err := getAllStreamWorkspaces(ctx, conn, *currentModel.ProjectId)
	if err != nil {
		return HandleError(apiResp, constants.LIST, err)
	}

	response := make([]interface{}, 0)
	for i := range accumulatedStreamWorkspaces {
		model := GetStreamWorkspaceModel(&accumulatedStreamWorkspaces[i], nil)
		model.ProjectId = currentModel.ProjectId
		model.Profile = currentModel.Profile
		response = append(response, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModels:  response,
	}, nil
}

func getAllStreamWorkspaces(ctx context.Context, conn *admin.APIClient, projectID string) ([]admin.StreamsTenant, *http.Response, error) {
	pageNum := 1
	accumulatedStreamWorkspaces := make([]admin.StreamsTenant, 0)

	for allStreamWorkspacesRetrieved := false; !allStreamWorkspacesRetrieved; {
		streamWorkspaces, resp, err := conn.StreamsApi.ListStreamWorkspacesWithParams(ctx, &admin.ListStreamWorkspacesApiParams{
			GroupId:      projectID,
			ItemsPerPage: util.Pointer(DefaultItemsPerPage),
			PageNum:      util.Pointer(pageNum),
		}).Execute()

		if err != nil {
			return nil, resp, err
		}
		accumulatedStreamWorkspaces = append(accumulatedStreamWorkspaces, streamWorkspaces.GetResults()...)
		allStreamWorkspacesRetrieved = streamWorkspaces.GetTotalCount() <= len(accumulatedStreamWorkspaces)
		pageNum++
	}

	return accumulatedStreamWorkspaces, nil, nil
}

func HandleError(response *http.Response, method constants.CfnFunctions, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())
	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}
