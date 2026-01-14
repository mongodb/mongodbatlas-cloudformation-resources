// Copyright 2025 MongoDB Inc
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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"go.mongodb.org/atlas-sdk/v20250312012/admin"
)

var callbackContext = map[string]any{"callbackFlex": true}

func IsCallback(req *handler.Request) bool {
	_, found := req.CallbackContext["callbackFlex"]
	return found
}

func HandleCreate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	if IsCallback(req) {
		return validateProgress(client, model, false)
	}
	flexReq := &admin.FlexClusterDescriptionCreate20241113{
		Name: *model.Name,
		ProviderSettings: admin.FlexProviderSettingsCreate20241113{
			BackingProviderName: *model.ProviderSettings.BackingProviderName,
			RegionName:          *model.ProviderSettings.RegionName,
		},
		TerminationProtectionEnabled: model.TerminationProtectionEnabled,
		Tags:                         expandTags(model.Tags),
	}
	flexResp, resp, err := client.AtlasSDK.FlexClustersApi.CreateFlexCluster(context.Background(), *model.ProjectId, flexReq).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe
	}
	return inProgressEvent(model, flexResp)
}

func HandleRead(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	flexResp, resp, err := client.AtlasSDK.FlexClustersApi.GetFlexCluster(context.Background(), *model.ProjectId, *model.Name).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe
	}
	updateModel(model, flexResp)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   model,
	}
}

func HandleUpdate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	if IsCallback(req) {
		return validateProgress(client, model, false)
	}
	updateReq := &admin.FlexClusterDescriptionUpdate20241113{
		TerminationProtectionEnabled: model.TerminationProtectionEnabled,
		Tags:                         expandTags(model.Tags),
	}
	flexResp, resp, err := client.AtlasSDK.FlexClustersApi.UpdateFlexCluster(context.Background(), *model.ProjectId, *model.Name, updateReq).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe
	}
	return inProgressEvent(model, flexResp)
}

func HandleDelete(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	if IsCallback(req) {
		return validateProgress(client, model, true)
	}
	resp, err := client.AtlasSDK.FlexClustersApi.DeleteFlexCluster(context.Background(), *model.ProjectId, *model.Name).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe
	}
	return inProgressEvent(model, nil)
}

func HandleList(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	var allModels []*Model
	const itemsPerPage = 100
	for pageNum := 1; ; pageNum++ {
		listOptions := &admin.ListFlexClustersApiParams{
			GroupId:      *model.ProjectId,
			ItemsPerPage: admin.PtrInt(itemsPerPage),
			PageNum:      admin.PtrInt(pageNum),
			IncludeCount: admin.PtrBool(true),
		}
		flexListResp, resp, err := client.AtlasSDK.FlexClustersApi.ListFlexClustersWithParams(context.Background(), listOptions).Execute()
		if pe := util.HandleClusterError(err, resp); pe != nil {
			return *pe
		}
		results := flexListResp.GetResults()
		for i := range results {
			modelItem := &Model{}
			updateModel(modelItem, &results[i])
			allModels = append(allModels, modelItem)
		}
		if len(allModels) >= flexListResp.GetTotalCount() || len(results) < itemsPerPage {
			break
		}
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModel:   allModels,
	}
}
