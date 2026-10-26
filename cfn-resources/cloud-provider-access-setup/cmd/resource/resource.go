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

	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

const (
	awsProviderName = "AWS"
	// roleInUseErrorCode is returned by Atlas when a feature still uses the role.
	roleInUseErrorCode = "AWS_IAM_ROLE_IN_USE"
	// roleInUseMessage tells the customer how to unblock the delete, because the dependent feature
	// has to be removed before the role can be deleted.
	roleInUseMessage = "Cloud Provider Access role %s cannot be deleted because an Atlas feature still uses it. " +
		"Remove the dependent feature, for example by disabling Encryption at Rest for the project, then retry the delete."
	// roleIDCallbackKey carries the registered role id across the create callback. CloudFormation
	// passes the callback context between invocations, so it is the reliable carrier of the id.
	roleIDCallbackKey = "roleId"
)

var (
	CreateRequiredFields = []string{constants.ProjectID}
	ReadRequiredFields   = []string{constants.ProjectID, constants.CloudProviderAccessRoleID}
	DeleteRequiredFields = []string{constants.ProjectID, constants.CloudProviderAccessRoleID}
	ListRequiredFields   = []string{constants.ProjectID}
)

func setup() {
	util.SetupLogger("mongodb-atlas-cloud-provider-access-setup")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	return HandleCreate(req.CallbackContext, client, currentModel), nil
}

// HandleCreate registers a role once and then reads the registered role on the callback. The role id
// travels in the callback context, which is the state CloudFormation carries between invocations; the
// returned model is the fallback when the callback context is absent.
func HandleCreate(callbackContext map[string]any, client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	if !util.IsStringPresent(currentModel.RoleId) {
		if roleID, ok := callbackContext[roleIDCallbackKey].(string); ok && roleID != "" {
			currentModel.RoleId = aws.String(roleID)
		}
	}

	if util.IsStringPresent(currentModel.RoleId) {
		event := HandleRead(client, currentModel)
		if event.OperationStatus == handler.Success {
			event.Message = "Create Complete"
		} else {
			// Preserve the identifier so CloudFormation can roll back this registration.
			event.ResourceModel = currentModel
		}
		return event
	}

	// Atlas cannot deduplicate a create whose response was lost before RoleId was
	// checkpointed. Do not adopt another role by project/provider: several independent
	// registrations may exist in the same project.
	roleRequest := &admin.CloudProviderAccessRoleRequest{ProviderName: awsProviderName}

	role, response, err := client.AtlasSDK.CloudProviderAccessApi.
		CreateCloudProviderAccess(context.Background(), *currentModel.ProjectId, roleRequest).
		Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(
			fmt.Sprintf("error creating Cloud Provider Access role: %s", err.Error()), response)
	}

	registered := NewCFNCloudProviderAccessSetup(currentModel, role)
	if !util.IsStringPresent(registered.RoleId) {
		// Without the role id the registration cannot be tracked or read back, and a retry would
		// register another role, so fail instead of returning an unusable checkpoint.
		return progress_events.GetFailedEventByCode(
			"Cloud Provider Access role registration did not return a role id",
			string(types.HandlerErrorCodeInternalFailure))
	}

	return progress_events.GetInProgressProgressEvent(
		"Cloud Provider Access role registered; verifying registration",
		map[string]any{roleIDCallbackKey: *registered.RoleId},
		registered, 1)
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(ReadRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	return HandleRead(client, currentModel), nil
}

// HandleRead reads a registration without changing its identity.
func HandleRead(client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	role, response, err := client.AtlasSDK.CloudProviderAccessApi.
		GetCloudProviderAccess(context.Background(), *currentModel.ProjectId, *currentModel.RoleId).
		Execute()
	if err != nil {
		if util.StatusNotFound(response) {
			return progress_events.GetFailedEventByCode(
				fmt.Sprintf("Cloud Provider Access role %s not found", *currentModel.RoleId),
				string(types.HandlerErrorCodeNotFound))
		}
		return progress_events.GetFailedEventByResponse(
			fmt.Sprintf("error reading Cloud Provider Access role: %s", err.Error()), response)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   NewCFNCloudProviderAccessSetup(currentModel, role),
	}
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	// Deleting removes the whole role registration, not just an authorization. The delete endpoint
	// reports a missing role with a message that names the project, so list the roles first and
	// report a missing registration as NotFound.
	roles, response, err := client.AtlasSDK.CloudProviderAccessApi.
		ListCloudProviderAccess(ctx, *currentModel.ProjectId).
		Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(
			fmt.Sprintf("error listing Cloud Provider Access roles: %s", err.Error()), response), nil
	}

	if FindAWSRole(roles, *currentModel.RoleId) == nil {
		return progress_events.GetFailedEventByCode(
			fmt.Sprintf("Cloud Provider Access role %s not found", *currentModel.RoleId),
			string(types.HandlerErrorCodeNotFound)), nil
	}

	response, err = client.AtlasSDK.CloudProviderAccessApi.
		DeauthorizeProviderAccessRole(ctx, *currentModel.ProjectId, awsProviderName, *currentModel.RoleId).
		Execute()
	if err != nil {
		// The role was removed between the list and the delete. The registration is gone, which is
		// the end state this handler wants, so tolerate the 404.
		if util.StatusNotFound(response) {
			return handler.ProgressEvent{OperationStatus: handler.Success, Message: "Delete Complete"}, nil
		}
		if admin.IsErrorCode(err, roleInUseErrorCode) {
			return progress_events.GetFailedEventByCode(
				fmt.Sprintf(roleInUseMessage, *currentModel.RoleId),
				string(types.HandlerErrorCodeResourceConflict)), nil
		}
		return progress_events.GetFailedEventByResponse(
			fmt.Sprintf("error deleting Cloud Provider Access role: %s", err.Error()), response), nil
	}

	return handler.ProgressEvent{OperationStatus: handler.Success, Message: "Delete Complete"}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(ListRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	roles, response, err := client.AtlasSDK.CloudProviderAccessApi.
		ListCloudProviderAccess(context.Background(), *currentModel.ProjectId).
		Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(
			fmt.Sprintf("error listing Cloud Provider Access roles: %s", err.Error()), response), nil
	}

	awsRoles := roles.GetAwsIamRoles()
	models := make([]any, 0, len(awsRoles))
	for i := range awsRoles {
		models = append(models, NewCFNCloudProviderAccessSetupFromAWSIAMRole(currentModel, &awsRoles[i]))
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}
