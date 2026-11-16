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
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/cloudprovideraccess"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
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

	return HandleCreate(client, currentModel), nil
}

// HandleCreate registers the role and returns the registered model from the same request. Atlas
// cannot deduplicate a create whose response is lost, so a retry registers another role; that
// accepted risk is documented in the resource README.
func HandleCreate(client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	roleRequest := &admin.CloudProviderAccessRoleRequest{ProviderName: constants.AWS}

	role, response, err := client.AtlasSDK.CloudProviderAccessApi.
		CreateCloudProviderAccess(context.Background(), *currentModel.ProjectId, roleRequest).
		Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(
			fmt.Sprintf("error creating Cloud Provider Access role: %s", err.Error()), response)
	}

	registered := NewCFNCloudProviderAccessSetup(currentModel, role)
	if !util.IsStringPresent(registered.RoleId) {
		// Without the role id the registration cannot be tracked, so fail instead of returning an
		// unusable model.
		return progressevent.GetFailedEventByCode(
			"Cloud Provider Access role registration did not return a role id",
			string(types.HandlerErrorCodeInternalFailure))
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   registered,
	}
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
	role, response, err := cloudprovideraccess.GetRole(context.Background(), client, *currentModel.ProjectId, *currentModel.RoleId)
	if err != nil {
		if cloudprovideraccess.IsRoleNotFound(err, response) {
			return cloudprovideraccess.NotFoundEvent(fmt.Sprintf("Cloud Provider Access role %s not found", *currentModel.RoleId))
		}
		return progressevent.GetFailedEventByResponse(
			fmt.Sprintf("error reading Cloud Provider Access role: %s", err.Error()), response)
	}

	// The lookup endpoint also returns Azure and GCP registrations, which this AWS-only
	// resource neither lists nor deletes.
	if !cloudprovideraccess.IsAWSRole(role) {
		return cloudprovideraccess.NotFoundEvent(fmt.Sprintf("AWS Cloud Provider Access role %s not found", *currentModel.RoleId))
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

	return HandleDelete(client, currentModel), nil
}

// HandleDelete removes the role registration, tolerating a registration that is already gone.
func HandleDelete(client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	// Deleting removes the whole role registration, not just an authorization. A missing
	// registration is reported as NotFound before the delete; a registration that is removed
	// concurrently is the end state this handler wants, so it is tolerated as success.
	role, response, err := cloudprovideraccess.GetRole(context.Background(), client, *currentModel.ProjectId, *currentModel.RoleId)
	if err != nil {
		if cloudprovideraccess.IsRoleNotFound(err, response) {
			return cloudprovideraccess.NotFoundEvent(fmt.Sprintf("Cloud Provider Access role %s not found", *currentModel.RoleId))
		}
		return progressevent.GetFailedEventByResponse(
			fmt.Sprintf("error reading Cloud Provider Access role: %s", err.Error()), response)
	}
	if !cloudprovideraccess.IsAWSRole(role) {
		return cloudprovideraccess.NotFoundEvent(fmt.Sprintf("AWS Cloud Provider Access role %s not found", *currentModel.RoleId))
	}

	if pe := cloudprovideraccess.DeauthorizeRole(context.Background(), client, *currentModel.ProjectId, *currentModel.RoleId); pe != nil {
		return *pe
	}

	return handler.ProgressEvent{OperationStatus: handler.Success, Message: "Delete Complete"}
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

	return HandleList(client, currentModel), nil
}

// HandleList returns the project's AWS Cloud Provider Access roles.
func HandleList(client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	roles, response, err := client.AtlasSDK.CloudProviderAccessApi.
		ListCloudProviderAccess(context.Background(), *currentModel.ProjectId).
		Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(
			fmt.Sprintf("error listing Cloud Provider Access roles: %s", err.Error()), response)
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
	}
}
