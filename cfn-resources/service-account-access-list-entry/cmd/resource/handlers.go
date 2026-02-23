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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/serviceaccountaccesslist"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"
)

func getCIDROrIP(model *Model) string {
	if model.IPAddress != nil && *model.IPAddress != "" {
		return *model.IPAddress
	}
	if model.CIDRBlock != nil && *model.CIDRBlock != "" {
		return *model.CIDRBlock
	}
	return ""
}

func newListPageFunc(client *util.MongoDBClient, orgID, clientID string) serviceaccountaccesslist.ListPageFunc {
	return func(ctx context.Context, pageNum int) (*admin.PaginatedServiceAccountIPAccessEntry, *http.Response, error) {
		return client.AtlasSDK.ServiceAccountsApi.ListOrgAccessList(ctx, orgID, clientID).
			PageNum(pageNum).ItemsPerPage(serviceaccountaccesslist.ItemsPerPage).Execute()
	}
}

func handleCreate(client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	orgID := *model.OrgId
	clientID := *model.ClientId

	createReq := []admin.ServiceAccountIPAccessListEntry{
		{
			CidrBlock: model.CIDRBlock,
			IpAddress: model.IPAddress,
		},
	}

	firstPage, apiResp, err := client.AtlasSDK.ServiceAccountsApi.CreateOrgAccessList(ctx, orgID, clientID, &createReq).ItemsPerPage(serviceaccountaccesslist.ItemsPerPage).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}

	cidrOrIP := getCIDROrIP(model)
	listPageFunc := newListPageFunc(client, orgID, clientID)

	entry, err := serviceaccountaccesslist.FindAccessListEntry(ctx, listPageFunc, cidrOrIP, firstPage)
	if err != nil || entry == nil {
		errMsg := "Created entry not found in response"
		if err != nil {
			errMsg = fmt.Sprintf("Error finding created entry: %s", err.Error())
		}
		return progress_events.GetFailedEventByCode(errMsg, string(types.HandlerErrorCodeInternalFailure))
	}

	UpdateModelFromEntry(model, entry)

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
	cidrOrIP := getCIDROrIP(model)
	listPageFunc := newListPageFunc(client, orgID, clientID)

	entry, err := serviceaccountaccesslist.FindAccessListEntry(ctx, listPageFunc, cidrOrIP, nil)
	if err != nil {
		return progress_events.GetFailedEventByCode(
			fmt.Sprintf("Error reading entry: %s", err.Error()),
			string(types.HandlerErrorCodeInternalFailure))
	}

	if entry == nil {
		return progress_events.GetFailedEventByCode(
			fmt.Sprintf("Access list entry %s not found for service account %s", cidrOrIP, clientID),
			string(types.HandlerErrorCodeNotFound))
	}

	UpdateModelFromEntry(model, entry)

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
	cidrOrIP := getCIDROrIP(model)

	apiResp, err := client.AtlasSDK.ServiceAccountsApi.DeleteOrgAccessEntry(ctx, orgID, clientID, cidrOrIP).Execute()
	if err != nil {
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
	listPageFunc := newListPageFunc(client, orgID, clientID)

	entries, err := serviceaccountaccesslist.ListAllAccessListEntries(ctx, listPageFunc)
	if err != nil {
		return progress_events.GetFailedEventByCode(
			fmt.Sprintf("Error listing entries: %s", err.Error()),
			string(types.HandlerErrorCodeInternalFailure))
	}

	entryModels := make([]any, 0, len(entries))
	for i := range entries {
		entryModel := &Model{
			OrgId:    model.OrgId,
			ClientId: model.ClientId,
			Profile:  model.Profile,
		}
		UpdateModelFromEntry(entryModel, &entries[i])
		entryModels = append(entryModels, entryModel)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModels:  entryModels,
	}
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) handler.ProgressEvent {
	errMsg := fmt.Sprintf("%s error: %s", method, err.Error())
	return progress_events.GetFailedEventByResponse(errMsg, response)
}
