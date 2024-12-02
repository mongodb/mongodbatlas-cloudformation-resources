// Copyright 2023 MongoDB Inc
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
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115014/admin"
)

var CreateRequiredFields = []string{constants.DatabaseName, constants.ProjectID, constants.Roles, constants.Username}
var ReadRequiredFields = []string{constants.ProjectID, constants.DatabaseName, constants.Username}
var UpdateRequiredFields = []string{constants.DatabaseName, constants.ProjectID, constants.Roles, constants.Username}
var DeleteRequiredFields = []string{constants.ProjectID, constants.DatabaseName, constants.Username}
var ListRequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-database-user")
}

// validateModel to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	dbUser, err := setModel(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          fmt.Sprintf("Error Creating resource: %s", err.Error()),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	groupID := *currentModel.ProjectId

	_, resp, err := client.Atlas20231115014.DatabaseUsersApi.CreateDatabaseUser(context.Background(), groupID, dbUser).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	updateUserCFNIdentifier(currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	groupID := *currentModel.ProjectId
	username := *currentModel.Username
	dbName := *currentModel.DatabaseName
	databaseUser, resp, err := client.Atlas20231115014.DatabaseUsersApi.GetDatabaseUser(context.Background(), groupID, dbName, username).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.DatabaseName = &databaseUser.DatabaseName

	if currentModel.LdapAuthType != nil {
		currentModel.LdapAuthType = databaseUser.LdapAuthType
	}
	if currentModel.AWSIAMType != nil {
		currentModel.AWSIAMType = databaseUser.AwsIAMType
	}
	if currentModel.X509Type != nil {
		currentModel.X509Type = databaseUser.X509Type
	}
	currentModel.Username = &databaseUser.Username
	_, _ = logger.Debugf("databaseUser:%+v", databaseUser)
	var roles []RoleDefinition

	for _, r := range databaseUser.GetRoles() {
		role := RoleDefinition{
			CollectionName: r.CollectionName,
			DatabaseName:   &r.DatabaseName,
			RoleName:       &r.RoleName,
		}

		roles = append(roles, role)
	}
	currentModel.Roles = roles

	var labels []LabelDefinition

	for _, l := range databaseUser.GetLabels() {
		label := LabelDefinition{
			Key:   l.Key,
			Value: l.Value,
		}

		labels = append(labels, label)
	}
	currentModel.Labels = labels
	var scopes []ScopeDefinition
	for _, s := range databaseUser.GetScopes() {
		scope := ScopeDefinition{
			Name: &s.Name,
			Type: &s.Type,
		}
		scopes = append(scopes, scope)
	}
	currentModel.Scopes = scopes
	updateUserCFNIdentifier(currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	dbUser, err := setModel(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          fmt.Sprintf("Error Creating resource: %s", err.Error()),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	groupID := *currentModel.ProjectId

	_, resp, err := client.Atlas20231115014.DatabaseUsersApi.UpdateDatabaseUser(context.Background(), groupID, *currentModel.DatabaseName, *currentModel.Username, dbUser).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	updateUserCFNIdentifier(currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if errEvent := validateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	groupID := *currentModel.ProjectId
	databaseName := *currentModel.DatabaseName
	username := *currentModel.Username
	_, resp, err := client.Atlas20231115014.DatabaseUsersApi.DeleteDatabaseUser(context.Background(), groupID, databaseName, username).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	updateUserCFNIdentifier(currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

// List handles listing database users
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	groupID := *currentModel.ProjectId

	dbUserModels := make([]interface{}, 0)

	databaseUsers, resp, err := client.Atlas20231115014.DatabaseUsersApi.ListDatabaseUsers(context.Background(), groupID).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	dbUserResults := databaseUsers.GetResults()
	for i := range dbUserResults {
		databaseUser := dbUserResults[i]
		var model = Model{
			DatabaseName: &databaseUser.DatabaseName,
			LdapAuthType: databaseUser.LdapAuthType,
			X509Type:     databaseUser.X509Type,
			Username:     &databaseUser.Username,
			ProjectId:    currentModel.ProjectId,
		}

		var roles []RoleDefinition

		for _, r := range databaseUser.GetRoles() {
			role := RoleDefinition{
				CollectionName: r.CollectionName,
				DatabaseName:   &r.DatabaseName,
				RoleName:       &r.RoleName,
			}

			roles = append(roles, role)
		}
		model.Roles = roles

		var labels []LabelDefinition

		for _, l := range databaseUser.GetLabels() {
			label := LabelDefinition{
				Key:   l.Key,
				Value: l.Value,
			}
			labels = append(labels, label)
		}

		model.Labels = labels
		updateUserCFNIdentifier(&model)
		dbUserModels = append(dbUserModels, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  dbUserModels,
	}, nil
}

func setModel(currentModel *Model) (*admin.CloudDatabaseUser, error) {
	roles := make([]admin.DatabaseUserRole, 0, len(currentModel.Roles))
	for i := range currentModel.Roles {
		r := currentModel.Roles[i]
		role := admin.DatabaseUserRole{}
		if r.CollectionName != nil {
			role.CollectionName = r.CollectionName
		}
		if r.DatabaseName != nil {
			role.DatabaseName = *r.DatabaseName
		}
		if r.RoleName != nil {
			role.RoleName = *r.RoleName
		}
		roles = append(roles, role)
	}

	labels := make([]admin.ComponentLabel, 0, len(currentModel.Labels))
	for i := range currentModel.Labels {
		l := currentModel.Labels[i]
		label := admin.ComponentLabel{
			Key:   l.Key,
			Value: l.Value,
		}
		labels = append(labels, label)
	}

	scopes := make([]admin.UserScope, 0, len(currentModel.Scopes))
	for i := range currentModel.Scopes {
		s := currentModel.Scopes[i]
		scope := admin.UserScope{
			Name: *s.Name,
			Type: *s.Type,
		}
		scopes = append(scopes, scope)
	}

	groupID := *currentModel.ProjectId

	none := "NONE"
	if currentModel.LdapAuthType == nil {
		currentModel.LdapAuthType = &none
	}
	if currentModel.AWSIAMType == nil {
		currentModel.AWSIAMType = &none
	}
	if currentModel.X509Type == nil {
		currentModel.X509Type = &none
	}

	if currentModel.Password == nil {
		if (*currentModel.LdapAuthType == none) && (*currentModel.AWSIAMType == none) && (*currentModel.X509Type == none) {
			err := fmt.Errorf("password cannot be empty if not LDAP or IAM or X509 is not provided")
			return nil, err
		}
		currentModel.Password = aws.String("")
	}

	if (currentModel.X509Type != &none) || (currentModel.DeleteAfterDate == nil) {
		currentModel.DeleteAfterDate = aws.String("")
	}

	user := &admin.CloudDatabaseUser{
		Roles:           &roles,
		GroupId:         groupID,
		Username:        *currentModel.Username,
		DatabaseName:    *currentModel.DatabaseName,
		Labels:          &labels,
		Scopes:          &scopes,
		LdapAuthType:    currentModel.LdapAuthType,
		AwsIAMType:      currentModel.AWSIAMType,
		X509Type:        currentModel.X509Type,
		DeleteAfterDate: util.StringPtrToTimePtr(currentModel.DeleteAfterDate),
	}

	if util.IsStringPresent(currentModel.Password) {
		user.Password = currentModel.Password
	}

	return user, nil
}

func updateUserCFNIdentifier(model *Model) {
	cfnid := fmt.Sprintf("%s-%s", *model.Username, *model.ProjectId)
	model.UserCFNIdentifier = &cfnid
}
