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
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/cloudprovideraccess"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

const (
	iamAssumedRoleArnField = "IamAssumedRoleArn"

	// Retry IAM propagation across callbacks instead of sleeping inside a handler.
	authorizeTimeout      = 5 * time.Minute
	authorizeRetrySeconds = 10

	// Atlas API error code this resource retries on.
	errorCodeCannotAssumeRole = "CANNOT_ASSUME_ROLE"
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
		if _, continuing := authorizationCallback(req, currentModel, "CREATE"); continuing {
			// A previous callback may have authorized the role before its response
			// was lost. Preserve its identifiers so CloudFormation can roll back.
			pe.ResourceModel = currentModel
		}
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
		event := roleLookupFailed(currentModel, resp, err)
		if continuing {
			event.ResourceModel = currentModel
		}
		return event
	}
	if !cloudprovideraccess.IsAWSRole(role) {
		event := cloudprovideraccess.NotFoundEvent(fmt.Sprintf("The AWS Cloud Provider Access role %s was not found in project %s.", util.SafeString(currentModel.RoleId), util.SafeString(currentModel.ProjectId)))
		if continuing {
			event.ResourceModel = currentModel
		}
		return event
	}
	if IsAuthorized(role) {
		if continuing && role.GetIamAssumedRoleArn() == util.SafeString(currentModel.IamAssumedRoleArn) {
			return authorizationComplete(currentModel, role)
		}
		event := progressevent.GetFailedEventByCode(
			fmt.Sprintf("The role %s is already authorized to %s.", util.SafeString(currentModel.RoleId), role.GetIamAssumedRoleArn()),
			string(types.HandlerErrorCodeAlreadyExists))
		if continuing {
			event.ResourceModel = currentModel
		}
		return event
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
		event := roleLookupFailed(currentModel, resp, err)
		event.ResourceModel = currentModel
		return event
	}
	if !IsAuthorized(role) {
		event := cloudprovideraccess.NotFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s is registered but not authorized.", util.SafeString(currentModel.RoleId)))
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
		return roleLookupFailed(currentModel, resp, err)
	}
	if !IsAuthorized(role) {
		return cloudprovideraccess.NotFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s is registered but not authorized.", util.SafeString(currentModel.RoleId)))
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
		return roleLookupFailed(currentModel, resp, err)
	}
	if !IsAuthorized(role) {
		return cloudprovideraccess.NotFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s is registered but not authorized.", util.SafeString(currentModel.RoleId)))
	}

	if pe := cloudprovideraccess.DeauthorizeRole(context.Background(), client, util.SafeString(currentModel.ProjectId), util.SafeString(currentModel.RoleId)); pe != nil {
		return *pe
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func getRole(client *util.MongoDBClient, model *Model) (*admin.CloudProviderAccessRole, *http.Response, error) {
	return cloudprovideraccess.GetRole(context.Background(), client, util.SafeString(model.ProjectId), util.SafeString(model.RoleId))
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

	role, resp, err := client.AtlasSDK.CloudProviderAccessApi.
		AuthorizeProviderAccessRole(context.Background(), projectID, roleID, &admin.CloudProviderAccessRoleRequestUpdate{
			ProviderName:      constants.AWS,
			IamAssumedRoleArn: &arn,
		}).Execute()
	if err != nil {
		if admin.IsErrorCode(err, errorCodeCannotAssumeRole) {
			return pendingAuthorization(model, operation, startedAt)
		}
		if cloudprovideraccess.IsRoleNotFound(err, resp) {
			event := cloudprovideraccess.NotFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s was not found in project %s. Create it with MongoDB::Atlas::CloudProviderAccessSetup before authorizing it.", roleID, projectID))
			event.ResourceModel = model
			return event
		}
		event := progressevent.GetFailedEventByResponse(err.Error(), resp)
		event.ResourceModel = model
		return event
	}

	return authorizationComplete(model, role)
}

func authorizationComplete(model *Model, role *admin.CloudProviderAccessRole) handler.ProgressEvent {
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

// IsAuthorized reports whether the role is an authorized AWS registration. The
// lookup endpoint also returns Azure and GCP registrations, which this AWS-only
// resource does not manage.
func IsAuthorized(role *admin.CloudProviderAccessRole) bool {
	return cloudprovideraccess.IsAWSRole(role) && role.HasAuthorizedDate()
}

func roleLookupFailed(model *Model, resp *http.Response, err error) handler.ProgressEvent {
	if cloudprovideraccess.IsRoleNotFound(err, resp) {
		return cloudprovideraccess.NotFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s was not found in project %s: %s", util.SafeString(model.RoleId), util.SafeString(model.ProjectId), err.Error()))
	}
	return progressevent.GetFailedEventByResponse(err.Error(), resp)
}
