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

// FeatureUsagesFromRole maps the feature type of each usage and, when Atlas
// identifies the feature, its project and bucket. Features that Atlas does not
// identify, for example ENCRYPTION_AT_REST, carry no feature id.
func FeatureUsagesFromRole(role *admin.CloudProviderAccessRole) []FeatureUsage {
	usages := role.GetFeatureUsages()
	result := make([]FeatureUsage, 0, len(usages))
	for i := range usages {
		usage := FeatureUsage{FeatureType: usages[i].FeatureType}
		if id := usages[i].FeatureId; id != nil && (id.BucketName != nil || id.GroupId != nil) {
			usage.FeatureId = &FeatureId{
				ProjectId:  id.GroupId,
				BucketName: id.BucketName,
			}
		}
		result = append(result, usage)
	}
	return result
}
