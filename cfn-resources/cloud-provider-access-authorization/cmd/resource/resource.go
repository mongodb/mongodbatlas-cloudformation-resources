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
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

const (
	iamAssumedRoleArnField = "IamAssumedRoleArn"

	// Retry IAM propagation across callbacks instead of sleeping inside a handler.
	authorizeTimeout      = 5 * time.Minute
	authorizeRetrySeconds = 10

	// Atlas API error codes this resource reacts to.
	errorCodeCannotAssumeRole = "CANNOT_ASSUME_ROLE"
	errorCodeRoleInUse        = "AWS_IAM_ROLE_IN_USE"
	errorCodeRoleNotFound     = "CLOUD_PROVIDER_ACCESS_ROLE_NOT_FOUND"
)

var (
	createAndUpdateRequiredFields = []string{constants.ProjectID, constants.CloudProviderAccessRoleID, iamAssumedRoleArnField}
	readAndDeleteRequiredFields   = []string{constants.ProjectID, constants.CloudProviderAccessRoleID}
)

func setup() {
	util.SetupLogger("mongodb-atlas-cloud-provider-access-authorization")
}

// Create authorizes the Cloud Provider Access role with the given ARN. The role
// registration must already exist, so it is created by the
// MongoDB::Atlas::CloudProviderAccessSetup resource first.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if pe := validator.ValidateModel(createAndUpdateRequiredFields, currentModel); pe != nil {
		return *pe, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	return HandleCreate(req, client, currentModel), nil
}

// HandleCreate checkpoints the operation before authorizing so a retried callback
// can recognize its own completed authorization without taking over an existing one.
func HandleCreate(req handler.Request, client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	startedAt, continuing := authorizationCallback(req, currentModel, "CREATE")
	role, resp, err := getRole(client, currentModel)
	if err != nil {
		event := roleLookupFailed(resp, err)
		if continuing {
			event.ResourceModel = currentModel
		}
		return event
	}
	if IsAuthorized(role) {
		if continuing && role.GetIamAssumedRoleArn() == util.SafeString(currentModel.IamAssumedRoleArn) {
			return authorizationComplete(currentModel, role)
		}
		return progressevent.GetFailedEventByCode(
			fmt.Sprintf("The role %s is already authorized to %s.", util.SafeString(currentModel.RoleId), role.GetIamAssumedRoleArn()),
			string(types.HandlerErrorCodeAlreadyExists))
	}

	if !continuing {
		return pendingAuthorization(currentModel, "CREATE", time.Now())
	}
	return authorizeRole(client, currentModel, "CREATE", startedAt)
}

// Update re-authorizes the same role in place. Changing IamAssumedRoleArn does
// not replace the resource, so the role registration is preserved.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if pe := validator.ValidateModel(createAndUpdateRequiredFields, currentModel); pe != nil {
		return *pe, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	return HandleUpdate(req, client, currentModel), nil
}

// HandleUpdate changes an existing authorization without creating a missing one.
func HandleUpdate(req handler.Request, client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	role, resp, err := getRole(client, currentModel)
	if err != nil {
		event := roleLookupFailed(resp, err)
		event.ResourceModel = currentModel
		return event
	}
	if !IsAuthorized(role) {
		event := notFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s is registered but not authorized.", util.SafeString(currentModel.RoleId)))
		event.ResourceModel = currentModel
		return event
	}
	if role.GetIamAssumedRoleArn() == util.SafeString(currentModel.IamAssumedRoleArn) {
		return authorizationComplete(currentModel, role)
	}
	startedAt, continuing := authorizationCallback(req, currentModel, "UPDATE")
	if !continuing {
		startedAt = time.Now()
	}
	return authorizeRole(client, currentModel, "UPDATE", startedAt)
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if pe := validator.ValidateModel(readAndDeleteRequiredFields, currentModel); pe != nil {
		return *pe, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	return HandleRead(client, currentModel), nil
}

// HandleRead returns the existing authorization, if any.
func HandleRead(client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	role, resp, err := getRole(client, currentModel)
	if err != nil {
		return roleLookupFailed(resp, err)
	}
	if !IsAuthorized(role) {
		return notFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s is registered but not authorized.", util.SafeString(currentModel.RoleId)))
	}

	SetModelFromRole(currentModel, role)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}
}

// Delete removes the Cloud Provider Access role registration, not only the
// authorization. Atlas has no revoke-only operation, so the paired
// MongoDB::Atlas::CloudProviderAccessSetup resource is left orphaned.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if pe := validator.ValidateModel(readAndDeleteRequiredFields, currentModel); pe != nil {
		return *pe, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	return HandleDelete(client, currentModel), nil
}

// HandleDelete removes the authorization and its backing registration.
func HandleDelete(client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	role, resp, err := getRole(client, currentModel)
	if err != nil {
		return roleLookupFailed(resp, err)
	}
	if !IsAuthorized(role) {
		return notFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s is registered but not authorized.", util.SafeString(currentModel.RoleId)))
	}

	resp, err = client.Atlas20231115014.CloudProviderAccessApi.
		DeauthorizeCloudProviderAccessRole(context.Background(), util.SafeString(currentModel.ProjectId), constants.AWS, util.SafeString(currentModel.RoleId)).
		Execute()
	if err != nil {
		if util.StatusNotFound(resp) || HasErrorCode(err, errorCodeRoleNotFound) {
			return notFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s was already removed.", util.SafeString(currentModel.RoleId)))
		}
		if HasErrorCode(err, errorCodeRoleInUse) {
			return progressevent.GetFailedEventByCode(
				"Cannot remove the AWS IAM role because it is still in use. Remove the feature that uses it first, for example disable Encryption at Rest, and retry the deletion.",
				string(types.HandlerErrorCodeResourceConflict))
		}
		return progressevent.GetFailedEventByResponse(err.Error(), resp)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func getRole(client *util.MongoDBClient, model *Model) (*admin20231115014.CloudProviderAccessRole, *http.Response, error) {
	return client.Atlas20231115014.CloudProviderAccessApi.
		GetCloudProviderAccessRole(context.Background(), util.SafeString(model.ProjectId), util.SafeString(model.RoleId)).
		Execute()
}

func authorizeRole(client *util.MongoDBClient, model *Model, operation string, startedAt time.Time) handler.ProgressEvent {
	if time.Since(startedAt) >= authorizeTimeout {
		// HandlerErrorCodeServiceTimeout carries the value "NotStabilized", the
		// CloudFormation code for a resource that never reached a stable state.
		event := progressevent.GetFailedEventByCode("Timed out waiting for Atlas to assume the AWS IAM role. Check its trust policy and external ID.", string(types.HandlerErrorCodeServiceTimeout))
		event.ResourceModel = model
		return event
	}
	projectID := util.SafeString(model.ProjectId)
	roleID := util.SafeString(model.RoleId)
	arn := util.SafeString(model.IamAssumedRoleArn)

	role, resp, err := client.Atlas20231115014.CloudProviderAccessApi.
		AuthorizeCloudProviderAccessRole(context.Background(), projectID, roleID, &admin20231115014.CloudProviderAccessRole{
			ProviderName:      constants.AWS,
			IamAssumedRoleArn: &arn,
		}).Execute()
	if err != nil {
		if HasErrorCode(err, errorCodeCannotAssumeRole) {
			return pendingAuthorization(model, operation, startedAt)
		}
		if util.StatusNotFound(resp) || HasErrorCode(err, errorCodeRoleNotFound) {
			event := notFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s was not found in project %s. Create it with MongoDB::Atlas::CloudProviderAccessSetup before authorizing it.", roleID, projectID))
			event.ResourceModel = model
			return event
		}
		event := progressevent.GetFailedEventByResponse(err.Error(), resp)
		event.ResourceModel = model
		return event
	}

	return authorizationComplete(model, role)
}

func authorizationComplete(model *Model, role *admin20231115014.CloudProviderAccessRole) handler.ProgressEvent {
	SetModelFromRole(model, role)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Authorization Complete",
		ResourceModel:   model,
	}
}

func pendingAuthorization(model *Model, operation string, startedAt time.Time) handler.ProgressEvent {
	return progressevent.GetInProgressProgressEvent("Waiting to authorize the AWS IAM role", map[string]any{
		"operation": operation,
		"projectId": util.SafeString(model.ProjectId),
		"roleId":    util.SafeString(model.RoleId),
		"profile":   util.SafeString(model.Profile),
		"arn":       util.SafeString(model.IamAssumedRoleArn),
		"startedAt": startedAt.UTC().Format(time.RFC3339Nano),
	}, model, authorizeRetrySeconds)
}

func authorizationCallback(req handler.Request, model *Model, operation string) (time.Time, bool) {
	for key, expected := range map[string]string{
		"operation": operation,
		"projectId": util.SafeString(model.ProjectId),
		"roleId":    util.SafeString(model.RoleId),
		"profile":   util.SafeString(model.Profile),
		"arn":       util.SafeString(model.IamAssumedRoleArn),
	} {
		if req.CallbackContext[key] != expected {
			return time.Time{}, false
		}
	}
	startedAt, _ := req.CallbackContext["startedAt"].(string)
	parsed, err := time.Parse(time.RFC3339Nano, startedAt)
	return parsed, err == nil
}

// HasErrorCode reports whether the Atlas API error carries the given
// application error code. The SDK exposes the decoded error through
// GenericOpenAPIError, and its formatted message also contains the code, so
// both are checked.
func HasErrorCode(err error, code string) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := errors.AsType[*admin20231115014.GenericOpenAPIError](err); ok {
		if model := apiErr.Model(); model.GetErrorCode() == code {
			return true
		}
	}
	return strings.Contains(err.Error(), code)
}

func IsAuthorized(role *admin20231115014.CloudProviderAccessRole) bool {
	return role != nil && role.HasAuthorizedDate()
}

func SetModelFromRole(model *Model, role *admin20231115014.CloudProviderAccessRole) {
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

// FeatureUsagesFromRole maps the feature type of each usage. The Atlas API models
// featureId as a polymorphic object that the typed SDK only covers for push-based
// log export, so it is deliberately not exposed.
func FeatureUsagesFromRole(role *admin20231115014.CloudProviderAccessRole) []FeatureUsage {
	usages := role.GetFeatureUsages()
	result := make([]FeatureUsage, 0, len(usages))
	for i := range usages {
		result = append(result, FeatureUsage{FeatureType: usages[i].FeatureType})
	}
	return result
}

func roleLookupFailed(resp *http.Response, err error) handler.ProgressEvent {
	if util.StatusNotFound(resp) || HasErrorCode(err, errorCodeRoleNotFound) {
		return notFoundEvent(fmt.Sprintf("The Cloud Provider Access role was not found. Create it with MongoDB::Atlas::CloudProviderAccessSetup before authorizing it: %s", err.Error()))
	}
	return progressevent.GetFailedEventByResponse(err.Error(), resp)
}

func notFoundEvent(message string) handler.ProgressEvent {
	_, _ = logger.Warnf("%s", message)
	return progressevent.GetFailedEventByCode(message, string(types.HandlerErrorCodeNotFound))
}
