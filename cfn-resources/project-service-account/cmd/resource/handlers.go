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
)

func handleCreate(client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	projectID := model.ProjectId
	serviceAccountReq := NewGroupServiceAccountCreateReq(model)

	serviceAccountResp, apiResp, err := client.AtlasSDK.ServiceAccountsApi.CreateGroupServiceAccount(ctx, *projectID, serviceAccountReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}

	resourceModel := GetGroupServiceAccountModel(serviceAccountResp, model)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModel:   resourceModel,
	}
}

func handleRead(client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	projectID := model.ProjectId
	clientID := model.ClientId

	serviceAccount, apiResp, err := client.AtlasSDK.ServiceAccountsApi.GetGroupServiceAccount(ctx, *projectID, *clientID).Execute()
	if err != nil {
		return handleError(apiResp, constants.READ, err)
	}

	resourceModel := GetGroupServiceAccountModel(serviceAccount, model)
	// Mask secrets on read (writeOnly property)
	if resourceModel.Secrets != nil {
		for i := range resourceModel.Secrets {
			resourceModel.Secrets[i].Secret = nil
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   resourceModel,
	}
}

func handleUpdate(client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	projectID := model.ProjectId
	clientID := model.ClientId

	serviceAccountReq := NewGroupServiceAccountUpdateReq(model)
	serviceAccountResp, apiResp, err := client.AtlasSDK.ServiceAccountsApi.UpdateGroupServiceAccount(ctx, *clientID, *projectID, serviceAccountReq).Execute()
	if err != nil {
		if util.StatusNotFound(apiResp) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Resource not found",
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}
		}
		return handleError(apiResp, constants.UPDATE, err)
	}

	resourceModel := GetGroupServiceAccountModel(serviceAccountResp, model)
	// Mask secrets on update (writeOnly property)
	if resourceModel.Secrets != nil {
		for i := range resourceModel.Secrets {
			resourceModel.Secrets[i].Secret = nil
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModel:   resourceModel,
	}
}

func handleDelete(client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	projectID := model.ProjectId
	clientID := model.ClientId

	apiResp, err := client.AtlasSDK.ServiceAccountsApi.DeleteGroupServiceAccount(ctx, *clientID, *projectID).Execute()
	if err != nil {
		if util.StatusNotFound(apiResp) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Resource not found",
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
			}
		}
		return handleError(apiResp, constants.DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
	}
}

func handleList(client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	projectID := model.ProjectId

	serviceAccounts, apiResp, err := client.AtlasSDK.ServiceAccountsApi.ListGroupServiceAccounts(ctx, *projectID).Execute()
	if err != nil {
		return handleError(apiResp, constants.LIST, err)
	}

	response := make([]interface{}, 0)
	if serviceAccounts != nil && serviceAccounts.Results != nil {
		for i := range *serviceAccounts.Results {
			itemModel := &Model{}
			resourceModel := GetGroupServiceAccountModel(&(*serviceAccounts.Results)[i], itemModel)
			resourceModel.ProjectId = model.ProjectId
			resourceModel.Profile = model.Profile
			if resourceModel.Secrets != nil {
				for j := range resourceModel.Secrets {
					resourceModel.Secrets[j].Secret = nil
				}
			}
			response = append(response, resourceModel)
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModels:  response,
	}
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) handler.ProgressEvent {
	errMsg := fmt.Sprintf("%s error: %s", method, err.Error())
	return progress_events.GetFailedEventByResponse(errMsg, response)
}
