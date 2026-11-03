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

// SetModelFromRole copies an API response into the resource model. The identifier
// and the requested ARN are preserved when the response omits them.
func SetModelFromRole(model *Model, role *admin.CloudProviderAccessRole) {
	if role == nil {
		return
	}
	if role.RoleId != nil {
		model.RoleId = role.RoleId
	}
	if role.IamAssumedRoleArn != nil {
		model.IamAssumedRoleArn = role.IamAssumedRoleArn
	}
	model.AuthorizedDate = util.TimePtrToStringPtr(role.AuthorizedDate)
	model.FeatureUsages = FeatureUsagesFromRole(role)
}

// FeatureUsagesFromRole maps the feature usages of a role into the resource model.
// It uses the shared conversion so both Cloud Provider Access resources expose the
// same shape, and leaves the attribute unset when Atlas reports no usage.
func FeatureUsagesFromRole(role *admin.CloudProviderAccessRole) []FeatureUsage {
	return cloudprovideraccess.MapFeatureUsages(role.FeatureUsages, func(usage cloudprovideraccess.FeatureUsage) FeatureUsage {
		item := FeatureUsage{FeatureType: usage.FeatureType}
		if usage.ProjectID != nil || usage.BucketName != nil {
			item.FeatureId = &FeatureId{
				ProjectId:  usage.ProjectID,
				BucketName: usage.BucketName,
			}
		}
		return item
	})
}
