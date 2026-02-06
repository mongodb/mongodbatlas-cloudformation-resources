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

import "go.mongodb.org/atlas-sdk/v20250312013/admin"

func NewLogIntegrationCreateRequest(model *Model) *admin.S3LogIntegrationRequest {
	if model == nil {
		return nil
	}
	req := &admin.S3LogIntegrationRequest{
		Type:       *model.Type,
		BucketName: *model.BucketName,
		IamRoleId:  *model.IamRoleId,
		PrefixPath: *model.PrefixPath,
		LogTypes:   model.LogTypes,
	}
	if model.KmsKey != nil && *model.KmsKey != "" {
		req.KmsKey = model.KmsKey
	}
	return req
}

func NewLogIntegrationUpdateRequest(model *Model) *admin.LogIntegrationRequest {
	if model == nil {
		return nil
	}
	req := &admin.LogIntegrationRequest{
		Type:       *model.Type,
		BucketName: model.BucketName,
		IamRoleId:  model.IamRoleId,
		PrefixPath: model.PrefixPath,
		LogTypes:   &model.LogTypes,
	}
	if model.KmsKey != nil && *model.KmsKey != "" {
		req.KmsKey = model.KmsKey
	}
	return req
}

func UpdateLogIntegrationModel(model *Model, logIntegrationResp *admin.LogIntegrationResponse) {
	if logIntegrationResp == nil {
		return
	}
	model.IntegrationId = &logIntegrationResp.Id
	model.BucketName = logIntegrationResp.BucketName
	model.IamRoleId = logIntegrationResp.IamRoleId
	model.PrefixPath = logIntegrationResp.PrefixPath
	model.Type = &logIntegrationResp.Type
	model.KmsKey = logIntegrationResp.KmsKey
	if logIntegrationResp.LogTypes != nil {
		model.LogTypes = *logIntegrationResp.LogTypes
	}
}
