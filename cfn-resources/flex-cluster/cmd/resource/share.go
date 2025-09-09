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
	"go.mongodb.org/atlas-sdk/v20250312006/admin"
)

func HandleCreate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	if util.IsCallback(req) {
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

func HandleDelete(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	if util.IsCallback(req) {
		return validateProgress(client, model, true)
	}
	resp, err := client.AtlasSDK.FlexClustersApi.DeleteFlexCluster(context.Background(), *model.ProjectId, *model.Name).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe
	}
	return inProgressEvent(model, nil)
}
