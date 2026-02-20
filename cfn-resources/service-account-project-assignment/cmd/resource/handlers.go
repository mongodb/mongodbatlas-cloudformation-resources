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
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"
)

const (
	InvalidRoleAssignmentError        = "INVALID_ROLE_ASSIGNMENT"
	ServiceAccountAlreadyInGroupError = "SERVICE_ACCOUNT_ALREADY_IN_GROUP"
)

func updateModel(model *Model, apiResp *admin.GroupServiceAccount) {
	if apiResp == nil {
		return
	}
	model.ClientId = apiResp.ClientId
	if apiResp.Roles != nil {
		model.Roles = *apiResp.Roles
	}
}

func HandleCreate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()

	apiReq := admin.NewGroupServiceAccountRoleAssignment(model.Roles)

	apiResp, resp, err := client.AtlasSDK.ServiceAccountsApi.InviteGroupServiceAccount(ctx, *model.ClientId, *model.ProjectId, apiReq).Execute()
	if err != nil {
		if apiError, ok := admin.AsError(err); ok && apiError.Error == http.StatusBadRequest {
			if apiError.ErrorCode == InvalidRoleAssignmentError && strings.Contains(err.Error(), ServiceAccountAlreadyInGroupError) {
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          err.Error(),
					HandlerErrorCode: string(types.HandlerErrorCodeAlreadyExists),
				}
			}
		}
		return HandleError(resp, constants.CREATE, err)
	}

	updateModel(model, apiResp)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModel:   model,
	}
}

func HandleRead(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()

	apiResp, resp, err := client.AtlasSDK.ServiceAccountsApi.GetGroupServiceAccount(ctx, *model.ProjectId, *model.ClientId).Execute()
	if err != nil {
		return HandleError(resp, constants.READ, err)
	}

	updateModel(model, apiResp)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   model,
	}
}

func HandleUpdate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()

	apiReq := admin.NewGroupServiceAccountUpdateRequest()
	apiReq.SetRoles(model.Roles)

	apiResp, resp, err := client.AtlasSDK.ServiceAccountsApi.UpdateGroupServiceAccount(ctx, *model.ClientId, *model.ProjectId, apiReq).Execute()
	if err != nil {
		return HandleError(resp, constants.UPDATE, err)
	}

	updateModel(model, apiResp)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModel:   model,
	}
}

func HandleDelete(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()

	resp, err := client.AtlasSDK.ServiceAccountsApi.DeleteGroupServiceAccount(ctx, *model.ClientId, *model.ProjectId).Execute()
	if err != nil {
		return HandleError(resp, constants.DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
	}
}

func HandleList(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	allModels := make([]interface{}, 0)
	const itemsPerPage = 100

	for pageNum := 1; ; pageNum++ {
		listResp, resp, err := client.AtlasSDK.ServiceAccountsApi.
			GetServiceAccountGroups(ctx, *model.OrgId, *model.ClientId).
			ItemsPerPage(itemsPerPage).
			PageNum(pageNum).
			Execute()

		if err != nil {
			return HandleError(resp, constants.LIST, err)
		}

		results := listResp.GetResults()
		for i := range results {
			if results[i].GroupId == nil {
				continue
			}

			modelItem := &Model{
				OrgId:     model.OrgId,
				ProjectId: results[i].GroupId,
				ClientId:  model.ClientId,
				Profile:   model.Profile,
				Roles:     []string{}, // API doesn't return roles in this endpoint
			}
			allModels = append(allModels, modelItem)
		}

		if len(results) == 0 || len(results) < itemsPerPage {
			break
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModels:  allModels,
	}
}

func HandleError(response *http.Response, method constants.CfnFunctions, err error) handler.ProgressEvent {
	errMsg := fmt.Sprintf("%s error: %s", method, err.Error())
	return progress_events.GetFailedEventByResponse(errMsg, response)
}
