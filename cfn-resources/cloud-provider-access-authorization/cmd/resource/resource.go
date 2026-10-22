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

	// authorizeAttempts and authorizeRetryDelay bound the retry of the authorize
	// call while AWS is still propagating a newly created IAM role. Three
	// attempts with a ten second pause stay well inside the CloudFormation
	// progress-event budget.
	authorizeAttempts   = 3
	authorizeRetryDelay = 10 * time.Second

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

	// A second create for the same role cannot produce a new distinct
	// identifier, so an existing authorization to the same ARN is reported as
	// AlreadyExists. This also makes a retried create idempotent.
	role, resp, err := getRole(client, currentModel)
	if err != nil {
		return roleLookupFailed(resp, err), nil
	}
	if IsAuthorized(role) && role.GetIamAssumedRoleArn() == util.SafeString(currentModel.IamAssumedRoleArn) {
		return progressevent.GetFailedEventByCode(
			fmt.Sprintf("The role %s is already authorized to %s.", util.SafeString(currentModel.RoleId), util.SafeString(currentModel.IamAssumedRoleArn)),
			string(types.HandlerErrorCodeAlreadyExists)), nil
	}

	return authorizeRole(client, currentModel)
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

	return authorizeRole(client, currentModel)
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

	role, resp, err := getRole(client, currentModel)
	if err != nil {
		return roleLookupFailed(resp, err), nil
	}
	if !IsAuthorized(role) {
		return notFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s is registered but not authorized.", util.SafeString(currentModel.RoleId))), nil
	}

	SetModelFromRole(currentModel, role)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
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

	role, resp, err := getRole(client, currentModel)
	if err != nil {
		return roleLookupFailed(resp, err), nil
	}
	if !IsAuthorized(role) {
		return notFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s is registered but not authorized.", util.SafeString(currentModel.RoleId))), nil
	}

	resp, err = client.Atlas20231115014.CloudProviderAccessApi.
		DeauthorizeCloudProviderAccessRole(context.Background(), util.SafeString(currentModel.ProjectId), constants.AWS, util.SafeString(currentModel.RoleId)).
		Execute()
	if err != nil {
		if util.StatusNotFound(resp) || HasErrorCode(err, errorCodeRoleNotFound) {
			return notFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s was already removed.", util.SafeString(currentModel.RoleId))), nil
		}
		if HasErrorCode(err, errorCodeRoleInUse) {
			return progressevent.GetFailedEventByCode(
				"Cannot remove the AWS IAM role because it is still in use. Remove the feature that uses it first, for example disable Encryption at Rest, and retry the deletion.",
				string(types.HandlerErrorCodeInvalidRequest)), nil
		}
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func getRole(client *util.MongoDBClient, model *Model) (*admin20231115014.CloudProviderAccessRole, *http.Response, error) {
	return client.Atlas20231115014.CloudProviderAccessApi.
		GetCloudProviderAccessRole(context.Background(), util.SafeString(model.ProjectId), util.SafeString(model.RoleId)).
		Execute()
}

func authorizeRole(client *util.MongoDBClient, model *Model) (handler.ProgressEvent, error) {
	projectID := util.SafeString(model.ProjectId)
	roleID := util.SafeString(model.RoleId)
	arn := util.SafeString(model.IamAssumedRoleArn)

	role, resp, err := AuthorizeWithRetry(authorizeAttempts, authorizeRetryDelay, time.Sleep, func() (*admin20231115014.CloudProviderAccessRole, *http.Response, error) {
		return client.Atlas20231115014.CloudProviderAccessApi.
			AuthorizeCloudProviderAccessRole(context.Background(), projectID, roleID, &admin20231115014.CloudProviderAccessRole{
				ProviderName:      constants.AWS,
				IamAssumedRoleArn: &arn,
			}).
			Execute()
	})
	if err != nil {
		if util.StatusNotFound(resp) || HasErrorCode(err, errorCodeRoleNotFound) {
			return notFoundEvent(fmt.Sprintf("The Cloud Provider Access role %s was not found in project %s. Create it with MongoDB::Atlas::CloudProviderAccessSetup before authorizing it.", roleID, projectID)), nil
		}
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	SetModelFromRole(model, role)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Authorization Complete",
		ResourceModel:   model,
	}, nil
}

// AuthorizeWithRetry runs attempt until it succeeds, fails with something other
// than CANNOT_ASSUME_ROLE, or exhausts the attempt budget. Atlas returns
// CANNOT_ASSUME_ROLE while AWS is still propagating a newly created IAM role,
// so a short retry usually succeeds. sleep is injectable so the retry can be
// unit tested without waiting.
func AuthorizeWithRetry(attempts int, delay time.Duration, sleep func(time.Duration), attempt func() (*admin20231115014.CloudProviderAccessRole, *http.Response, error)) (*admin20231115014.CloudProviderAccessRole, *http.Response, error) {
	var (
		role *admin20231115014.CloudProviderAccessRole
		resp *http.Response
		err  error
	)
	for i := range attempts {
		role, resp, err = attempt()
		if err == nil || !HasErrorCode(err, errorCodeCannotAssumeRole) {
			return role, resp, err
		}
		if i < attempts-1 {
			sleep(delay)
		}
	}
	return role, resp, err
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

// FeatureUsagesFromRole always returns a non-nil slice so the read-only
// FeatureUsages property is never null.
func FeatureUsagesFromRole(role *admin20231115014.CloudProviderAccessRole) []FeatureUsage {
	usages := role.GetFeatureUsages()
	result := make([]FeatureUsage, 0, len(usages))
	for i := range usages {
		featureID := usages[i].GetFeatureId()
		result = append(result, FeatureUsage{
			FeatureType: usages[i].FeatureType,
			FeatureId: &FeatureUsageId{
				GroupId:    featureID.GroupId,
				BucketName: featureID.BucketName,
			},
		})
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
