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

// Package cloudprovideraccess holds the Cloud Provider Access logic shared by the resource types.
package cloudprovideraccess

import (
	"net/http"

	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

const (
	// ErrorCodeRoleInUse is returned by Atlas when an Atlas feature still uses the role.
	ErrorCodeRoleInUse = "AWS_IAM_ROLE_IN_USE"
	// ErrorCodeRoleNotFound is returned by Atlas when the registration is gone.
	ErrorCodeRoleNotFound = "CLOUD_PROVIDER_ACCESS_ROLE_NOT_FOUND"
	// RoleInUseMessage tells the customer how to unblock a delete, because the dependent feature has
	// to be removed before the role can be deleted.
	RoleInUseMessage = "Cloud Provider Access role %s cannot be deleted because an Atlas feature still uses it. " +
		"Remove the dependent feature, for example by disabling Encryption at Rest for the project, then retry the delete."
)

// FeatureUsage is one Atlas feature that uses the role, with the identifiers Atlas returns for it.
// ProjectID and BucketName are only set for the features Atlas identifies.
type FeatureUsage struct {
	FeatureType *string
	ProjectID   *string
	BucketName  *string
}

// FeatureUsages converts the feature usages of a role into the shared view. It returns nil when
// Atlas reports no usage, so callers can leave the attribute out of the model.
func FeatureUsages(usages *[]admin.CloudProviderAccessFeatureUsage) []FeatureUsage {
	if usages == nil || len(*usages) == 0 {
		return nil
	}
	result := make([]FeatureUsage, 0, len(*usages))
	for i := range *usages {
		usage := FeatureUsage{FeatureType: (*usages)[i].FeatureType}
		if id := (*usages)[i].FeatureId; id != nil && (id.BucketName != nil || id.GroupId != nil) {
			usage.ProjectID = id.GroupId
			usage.BucketName = id.BucketName
		}
		result = append(result, usage)
	}
	return result
}

// IsRoleInUse reports whether Atlas rejected the call because a feature still uses the role.
func IsRoleInUse(err error) bool {
	return admin.IsErrorCode(err, ErrorCodeRoleInUse)
}

// IsRoleNotFound reports whether the call failed because the registration is gone.
func IsRoleNotFound(err error, response *http.Response) bool {
	return util.StatusNotFound(response) || admin.IsErrorCode(err, ErrorCodeRoleNotFound)
}
