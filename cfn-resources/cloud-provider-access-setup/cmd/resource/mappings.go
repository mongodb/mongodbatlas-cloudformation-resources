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
	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/cloudprovideraccess"
)

// NewCFNCloudProviderAccessSetup builds the resource model from a create or read API response.
func NewCFNCloudProviderAccessSetup(model *Model, role *admin.CloudProviderAccessRole) Model {
	return Model{
		Profile:                    model.Profile,
		ProjectId:                  model.ProjectId,
		RoleId:                     role.RoleId,
		AtlasAWSAccountArn:         role.AtlasAWSAccountArn,
		AtlasAssumedRoleExternalId: role.AtlasAssumedRoleExternalId,
		CreatedDate:                util.TimePtrToStringPtr(role.CreatedDate),
		FeatureUsages:              cloudprovideraccess.MapFeatureUsages(role.FeatureUsages, newFeatureUsage),
	}
}

// NewCFNCloudProviderAccessSetupFromAWSIAMRole builds the resource model from one AWS entry of the
// list API response.
func NewCFNCloudProviderAccessSetupFromAWSIAMRole(model *Model, role *admin.CloudProviderAccessAWSIAMRole) Model {
	return Model{
		Profile:                    model.Profile,
		ProjectId:                  model.ProjectId,
		RoleId:                     role.RoleId,
		AtlasAWSAccountArn:         role.AtlasAWSAccountArn,
		AtlasAssumedRoleExternalId: role.AtlasAssumedRoleExternalId,
		CreatedDate:                util.TimePtrToStringPtr(role.CreatedDate),
		FeatureUsages:              cloudprovideraccess.MapFeatureUsages(role.FeatureUsages, newFeatureUsage),
	}
}

// newFeatureUsage maps one shared feature usage into the resource model. It uses the shared
// conversion so both Cloud Provider Access resources expose the same shape.
func newFeatureUsage(usage cloudprovideraccess.FeatureUsage) FeatureUsage {
	item := FeatureUsage{FeatureType: usage.FeatureType}
	if usage.ProjectID != nil || usage.BucketName != nil {
		item.FeatureId = &FeatureId{
			ProjectId:  usage.ProjectID,
			BucketName: usage.BucketName,
		}
	}
	return item
}
