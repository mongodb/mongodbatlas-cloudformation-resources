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
	"context"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
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

// MapFeatureUsages converts the feature usages of a role and maps each one with the caller's
// constructor, so both Cloud Provider Access resources share the conversion.
func MapFeatureUsages[T any](usages *[]admin.CloudProviderAccessFeatureUsage, makeItem func(FeatureUsage) T) []T {
	converted := FeatureUsages(usages)
	if converted == nil {
		return nil
	}
	result := make([]T, 0, len(converted))
	for _, usage := range converted {
		result = append(result, makeItem(usage))
	}
	return result
}

// GetRole returns the registration with the given id.
func GetRole(ctx context.Context, client *util.MongoDBClient, projectID, roleID string) (*admin.CloudProviderAccessRole, *http.Response, error) {
	return client.AtlasSDK.CloudProviderAccessApi.
		GetCloudProviderAccess(ctx, projectID, roleID).
		Execute()
}

// IsAWSRole reports whether the role is an AWS registration. The lookup endpoint also returns
// Azure and GCP registrations, which the AWS-only resources do not manage.
func IsAWSRole(role *admin.CloudProviderAccessRole) bool {
	return role != nil && role.GetProviderName() == constants.AWS
}

// IsRoleInUse reports whether Atlas rejected the call because a feature still uses the role.
func IsRoleInUse(err error) bool {
	return admin.IsErrorCode(err, ErrorCodeRoleInUse)
}

// IsRoleNotFound reports whether the call failed because the registration is gone.
func IsRoleNotFound(err error, response *http.Response) bool {
	return util.StatusNotFound(response) || admin.IsErrorCode(err, ErrorCodeRoleNotFound)
}

// NotFoundEvent builds the failed event used when a registration is missing, and logs the message.
func NotFoundEvent(message string) handler.ProgressEvent {
	_, _ = logger.Warnf("%s", message)
	return progressevent.GetFailedEventByCode(message, string(types.HandlerErrorCodeNotFound))
}

// DeauthorizeRole removes the registration. It returns nil on success, including a registration
// that was removed concurrently, and a failed event otherwise: ResourceConflict when an Atlas
// feature still uses the role, and the mapped API error for anything else.
func DeauthorizeRole(ctx context.Context, client *util.MongoDBClient, projectID, roleID string) *handler.ProgressEvent {
	response, err := client.AtlasSDK.CloudProviderAccessApi.
		DeauthorizeProviderAccessRole(ctx, projectID, constants.AWS, roleID).
		Execute()
	if err == nil || IsRoleNotFound(err, response) {
		return nil
	}

	var event handler.ProgressEvent
	if IsRoleInUse(err) {
		event = progressevent.GetFailedEventByCode(
			fmt.Sprintf(RoleInUseMessage, roleID),
			string(types.HandlerErrorCodeResourceConflict))
	} else {
		event = progressevent.GetFailedEventByResponse(err.Error(), response)
	}
	return &event
}
