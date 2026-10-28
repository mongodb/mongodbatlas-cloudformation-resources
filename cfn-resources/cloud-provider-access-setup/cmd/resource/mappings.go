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
		FeatureUsages:              newFeatureUsages(role.FeatureUsages),
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
		FeatureUsages:              newFeatureUsages(role.FeatureUsages),
	}
}

func newFeatureUsages(usages *[]admin.CloudProviderAccessFeatureUsage) []FeatureUsage {
	converted := cloudprovideraccess.FeatureUsages(usages)
	if converted == nil {
		return nil
	}
	result := make([]FeatureUsage, 0, len(converted))
	for _, usage := range converted {
		item := FeatureUsage{FeatureType: usage.FeatureType}
		if usage.ProjectID != nil || usage.BucketName != nil {
			item.FeatureId = &FeatureId{
				ProjectId:  usage.ProjectID,
				BucketName: usage.BucketName,
			}
		}
		result = append(result, item)
	}
	return result
}

// FindAWSRole returns the AWS role with the given id from a list response, or nil when the project
// has no such registration.
func FindAWSRole(roles *admin.CloudProviderAccessRoles, roleID string) *admin.CloudProviderAccessAWSIAMRole {
	if roles == nil {
		return nil
	}
	awsRoles := roles.GetAwsIamRoles()
	for i := range awsRoles {
		if awsRoles[i].GetRoleId() == roleID {
			return &awsRoles[i]
		}
	}
	return nil
}
