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
	"go.mongodb.org/atlas-sdk/v20250312013/admin"
)

func handleCreate(client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	orgID := *model.OrgId
	clientID := *model.ClientId

	createReq := &admin.ServiceAccountSecretRequest{}
	if model.SecretExpiresAfterHours != nil {
		createReq.SecretExpiresAfterHours = *model.SecretExpiresAfterHours
	}

	secretResp, apiResp, err := client.AtlasSDK.ServiceAccountsApi.CreateOrgSecret(ctx, orgID, clientID, createReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}

	UpdateModelFromSecret(model, secretResp)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModel:   model,
	}
}

func handleRead(client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	orgID := *model.OrgId
	clientID := *model.ClientId
	secretID := *model.SecretId

	serviceAccountResp, apiResp, err := client.AtlasSDK.ServiceAccountsApi.GetOrgServiceAccount(ctx, orgID, clientID).Execute()
	if err != nil {
		return handleError(apiResp, constants.READ, err)
	}

	var foundSecret *admin.ServiceAccountSecret
	if serviceAccountResp.Secrets != nil {
		for i := range *serviceAccountResp.Secrets {
			secret := &(*serviceAccountResp.Secrets)[i]
			if secret.Id == secretID {
				foundSecret = secret
				break
			}
		}
	}

	if foundSecret == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          fmt.Sprintf("Secret with ID %s not found in service account %s", secretID, clientID),
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
		}
	}

	UpdateModelFromSecret(model, foundSecret)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   model,
	}
}

func handleDelete(client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	orgID := *model.OrgId
	clientID := *model.ClientId
	secretID := *model.SecretId

	apiResp, err := client.AtlasSDK.ServiceAccountsApi.DeleteOrgSecret(ctx, clientID, secretID, orgID).Execute()
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
	orgID := *model.OrgId
	clientID := *model.ClientId

	serviceAccountResp, apiResp, err := client.AtlasSDK.ServiceAccountsApi.GetOrgServiceAccount(ctx, orgID, clientID).Execute()
	if err != nil {
		return handleError(apiResp, constants.LIST, err)
	}

	var secretModels []any
	if serviceAccountResp.Secrets != nil {
		for i := range *serviceAccountResp.Secrets {
			secret := &(*serviceAccountResp.Secrets)[i]
			secretModel := &Model{
				OrgId:    model.OrgId,
				ClientId: model.ClientId,
				Profile:  model.Profile,
			}
			UpdateModelFromSecret(secretModel, secret)
			secretModels = append(secretModels, secretModel)
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModels:  secretModels,
	}
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) handler.ProgressEvent {
	errMsg := fmt.Sprintf("%s error: %s", method, err.Error())
	return progress_events.GetFailedEventByResponse(errMsg, response)
}
