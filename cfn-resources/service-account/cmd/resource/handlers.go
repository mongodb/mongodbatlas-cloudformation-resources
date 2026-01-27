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
	orgID := model.OrgId
	serviceAccountReq := NewOrgServiceAccountCreateReq(model)

	serviceAccountResp, apiResp, err := client.AtlasSDK.ServiceAccountsApi.CreateOrgServiceAccount(ctx, *orgID, serviceAccountReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}

	resourceModel := GetOrgServiceAccountModel(serviceAccountResp, model)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModel:   resourceModel,
	}
}

func handleRead(client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	orgID := model.OrgId
	clientID := model.ClientId

	serviceAccount, apiResp, err := client.AtlasSDK.ServiceAccountsApi.GetOrgServiceAccount(ctx, *orgID, *clientID).Execute()
	if err != nil {
		return handleError(apiResp, constants.READ, err)
	}

	resourceModel := GetOrgServiceAccountModel(serviceAccount, model)
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
	orgID := model.OrgId
	clientID := model.ClientId

	serviceAccountReq := NewOrgServiceAccountUpdateReq(model)
	serviceAccountResp, apiResp, err := client.AtlasSDK.ServiceAccountsApi.UpdateOrgServiceAccount(ctx, *clientID, *orgID, serviceAccountReq).Execute()
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

	resourceModel := GetOrgServiceAccountModel(serviceAccountResp, model)
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
	orgID := model.OrgId
	clientID := model.ClientId

	apiResp, err := client.AtlasSDK.ServiceAccountsApi.DeleteOrgServiceAccount(ctx, *clientID, *orgID).Execute()
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
	orgID := model.OrgId

	serviceAccounts, apiResp, err := client.AtlasSDK.ServiceAccountsApi.ListOrgServiceAccounts(ctx, *orgID).Execute()
	if err != nil {
		return handleError(apiResp, constants.LIST, err)
	}

	response := make([]interface{}, 0)
	if serviceAccounts != nil && serviceAccounts.Results != nil {
		for i := range *serviceAccounts.Results {
			itemModel := &Model{}
			resourceModel := GetOrgServiceAccountModel(&(*serviceAccounts.Results)[i], itemModel)
			resourceModel.OrgId = model.OrgId
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
