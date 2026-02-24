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

	"go.mongodb.org/atlas-sdk/v20250312014/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
)

func HandleCreate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	if isCallback(req) {
		return validateProgress(client, model, false)
	}

	ctx := context.Background()
	projectID := util.SafeString(model.ProjectId)
	streamPrivatelinkEndpointReq := NewStreamPrivatelinkEndpointReq(model)

	streamsPrivateLinkConnection, apiResp, err := client.AtlasSDK.StreamsApi.CreatePrivateLinkConnection(ctx, projectID, streamPrivatelinkEndpointReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}

	model.Id = streamsPrivateLinkConnection.Id

	initialState := streamsPrivateLinkConnection.GetState()
	if initialState != "" {
		if initialState == stateDone {
			UpdateModel(model, streamsPrivateLinkConnection)
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         constants.Complete,
				ResourceModel:   model,
			}
		}
		if initialState == stateFailed {
			return handleFailedState(streamsPrivateLinkConnection)
		}
	}

	return inProgressEvent(model, streamsPrivateLinkConnection)
}

func HandleRead(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	projectID := util.SafeString(model.ProjectId)
	connectionID := util.SafeString(model.Id)

	streamsPrivateLinkConnection, apiResp, err := client.AtlasSDK.StreamsApi.GetPrivateLinkConnection(ctx, projectID, connectionID).Execute()
	if err != nil {
		if util.StatusNotFound(apiResp) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          constants.ResourceNotFound,
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}
		}
		return handleError(apiResp, constants.READ, err)
	}

	UpdateModel(model, streamsPrivateLinkConnection)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   model,
	}
}

func HandleUpdate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          errorMessageUpdateUnsupported,
		HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
	}
}

func HandleDelete(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	connectionID := util.SafeString(model.Id)
	if connectionID == "" {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          constants.ResourceNotFound,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
		}
	}

	if isCallback(req) {
		return validateProgress(client, model, true)
	}

	ctx := context.Background()
	projectID := util.SafeString(model.ProjectId)

	streamsPrivateLinkConnection, apiResp, err := client.AtlasSDK.StreamsApi.GetPrivateLinkConnection(ctx, projectID, connectionID).Execute()
	if err != nil {
		if util.StatusNotFound(apiResp) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          constants.ResourceNotFound,
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}
		}
		return handleError(apiResp, constants.READ, err)
	}

	if streamsPrivateLinkConnection.GetState() == stateDeleted {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         constants.Complete,
		}
	}

	deleteResp, err := client.AtlasSDK.StreamsApi.DeletePrivateLinkConnection(ctx, projectID, connectionID).Execute()
	if err != nil {
		return handleError(deleteResp, constants.DELETE, err)
	}

	return inProgressEvent(model, nil)
}

func HandleList(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	projectID := util.SafeString(model.ProjectId)

	allModels := make([]any, 0)
	const itemsPerPage = 100
	for pageNum := 1; ; pageNum++ {
		listOptions := &admin.ListPrivateLinkConnectionsApiParams{
			GroupId:      projectID,
			ItemsPerPage: util.Pointer(itemsPerPage),
			PageNum:      util.Pointer(pageNum),
		}
		connections, apiResp, err := client.AtlasSDK.StreamsApi.ListPrivateLinkConnectionsWithParams(ctx, listOptions).Execute()
		if err != nil {
			return handleError(apiResp, constants.LIST, err)
		}

		results := connections.GetResults()
		for i := range results {
			connectionState := ""
			if results[i].State != nil {
				connectionState = *results[i].State
			}
			if connectionState == stateDeleted || connectionState == stateDeleteRequested || connectionState == stateDeleting {
				continue
			}

			modelItem := &Model{
				ProjectId: model.ProjectId,
				Profile:   model.Profile,
			}
			UpdateModel(modelItem, &results[i])
			allModels = append(allModels, modelItem)
		}

		if connections.GetTotalCount() <= len(allModels) || len(results) < itemsPerPage {
			break
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModels:  allModels,
	}
}
