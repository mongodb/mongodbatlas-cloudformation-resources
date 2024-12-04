// Copyright 2024 MongoDB Inc
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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"go.mongodb.org/atlas-sdk/v20241113002/admin"
)

func NewResourcePolicyCreateReq(model *Model) *admin.ApiAtlasResourcePolicyCreate {
	if model == nil {
		return nil
	}
	return &admin.ApiAtlasResourcePolicyCreate{
		Name:     *model.Name,
		Policies: *modelPoliciesToSDKPolicies(model.Policies),
	}
}

func modelPoliciesToSDKPolicies(policy []ApiAtlasPolicy) *[]admin.ApiAtlasPolicyCreate {
	policiesCreate := make([]admin.ApiAtlasPolicyCreate, len(policy))
	for i, v := range policy {
		policiesCreate[i].Body = *v.Body
	}
	return &policiesCreate
}

func NewResourcePolicyUpdateReq(model *Model) *admin.ApiAtlasResourcePolicyEdit {
	if model == nil {
		return nil
	}
	return &admin.ApiAtlasResourcePolicyEdit{
		Name:     model.Name,
		Policies: modelPoliciesToSDKPolicies(model.Policies),
	}
}

func GetResourcePolicyModel(resourcePolicyResp *admin.ApiAtlasResourcePolicy, currentModel *Model) *Model {
	model := new(Model)

	if currentModel != nil {
		model = currentModel
	}
	if resourcePolicyResp != nil {
		model.CreatedByUser = newAPIAtlasUserMetadata(resourcePolicyResp.CreatedByUser)
		model.CreatedDate = util.TimePtrToStringPtr(resourcePolicyResp.CreatedDate)
		model.Id = resourcePolicyResp.Id
		model.LastUpdatedByUser = newAPIAtlasUserMetadata(resourcePolicyResp.LastUpdatedByUser)
		model.LastUpdatedDate = util.TimePtrToStringPtr(resourcePolicyResp.LastUpdatedDate)
		model.Name = resourcePolicyResp.Name
		model.OrgId = resourcePolicyResp.OrgId
		model.Version = resourcePolicyResp.Version
		model.Policies = sdkPoliciesToModelPolicies(resourcePolicyResp.Policies)
	}
	return model
}

func newAPIAtlasUserMetadata(userMetadata *admin.ApiAtlasUserMetadata) *ApiAtlasUserMetadata {
	if userMetadata == nil {
		return nil
	}
	return &ApiAtlasUserMetadata{
		Id:   userMetadata.Id,
		Name: userMetadata.Name,
	}
}

func sdkPoliciesToModelPolicies(policies *[]admin.ApiAtlasPolicy) []ApiAtlasPolicy {
	if policies == nil {
		return nil
	}
	sdkPolicies := make([]ApiAtlasPolicy, len(*policies))
	for i, v := range *policies {
		sdkPolicies[i].Body = v.Body
		sdkPolicies[i].Id = v.Id
	}
	return sdkPolicies
}
