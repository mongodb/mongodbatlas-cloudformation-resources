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
	"github.com/openlyinc/pointy"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
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
	_, _ = logger.Debugf(" currentModel: %#+v, prevModel: %#+v", currentModel, prevModel)

	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(util.DefaultProfile)
	}
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	groupID, dbUser, event, err := setModel(currentModel)
	if err != nil {
		return event, nil
	}
	_, _ = logger.Debugf("Arguments: Project ID: %s, Request %#+v", groupID, dbUser)
	newUser, res, err := client.DatabaseUsers.Create(context.Background(), groupID, dbUser)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			res.Response), nil
	}
	_, _ = logger.Debugf("newUser: %s", newUser)
	cfnid := fmt.Sprintf("%s-%s", *currentModel.Username, groupID)
	currentModel.UserCFNIdentifier = &cfnid

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(util.DefaultProfile)
	}
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	groupID := *currentModel.ProjectId
	username := *currentModel.Username
	dbName := *currentModel.DatabaseName
	databaseUser, resp, err := client.DatabaseUsers.Get(context.Background(), dbName, groupID, username)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	currentModel.DatabaseName = &databaseUser.DatabaseName

	if currentModel.LdapAuthType != nil {
		currentModel.LdapAuthType = &databaseUser.LDAPAuthType
	}
	if currentModel.AWSIAMType != nil {
		currentModel.AWSIAMType = &databaseUser.AWSIAMType
	}
	if currentModel.X509Type != nil {
		currentModel.X509Type = &databaseUser.X509Type
	}
	currentModel.Username = &databaseUser.Username
	_, _ = logger.Debugf("databaseUser:%+v", databaseUser)
	var roles []RoleDefinition

	for i := range databaseUser.Roles {
		r := databaseUser.Roles[i]
		role := RoleDefinition{
			CollectionName: &r.CollectionName,
			DatabaseName:   &r.DatabaseName,
			RoleName:       &r.RoleName,
		}

		roles = append(roles, role)
	}
	currentModel.Roles = roles
	_, _ = logger.Debugf("currentModel.Roles:%+v", roles)
	var labels []LabelDefinition

	for i := range databaseUser.Labels {
		l := databaseUser.Labels[i]
		label := LabelDefinition{
			Key:   &l.Key,
			Value: &l.Value,
		}

		labels = append(labels, label)
	}
	currentModel.Labels = labels

	cfnid := fmt.Sprintf("%s-%s", *currentModel.Username, groupID)
	currentModel.UserCFNIdentifier = &cfnid
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(util.DefaultProfile)
	}
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	groupID, dbUser, event, err := setModel(currentModel)
	if err != nil {
		return event, nil
	}
	_, resp, err := client.DatabaseUsers.Update(context.Background(), groupID, *currentModel.Username, dbUser)

	_, _ = logger.Debugf("Update resp:%+v", resp)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	cfnid := fmt.Sprintf("%s-%s", *currentModel.Username, groupID)
	currentModel.UserCFNIdentifier = &cfnid

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if errEvent := validateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(util.DefaultProfile)
	}
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	groupID := *currentModel.ProjectId
	username := *currentModel.Username
	dbName := *currentModel.DatabaseName

	resp, err := client.DatabaseUsers.Delete(context.Background(), dbName, groupID, username)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	cfnid := fmt.Sprintf("%s-%s", *currentModel.Username, groupID)
	currentModel.UserCFNIdentifier = &cfnid
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

// List handles listing database users
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(util.DefaultProfile)
	}
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	groupID := *currentModel.ProjectId

	dbUserModels := make([]interface{}, 0)

	databaseUsers, resp, err := client.DatabaseUsers.List(context.Background(), groupID, nil)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	if len(databaseUsers) > 0 {
		for i := range databaseUsers {
			var model Model

			databaseUser := databaseUsers[i]
			model.DatabaseName = &databaseUser.DatabaseName
			model.LdapAuthType = &databaseUser.LDAPAuthType
			model.X509Type = &databaseUser.X509Type
			model.Username = &databaseUser.Username
			model.ProjectId = currentModel.ProjectId
			var roles []RoleDefinition

			for i := range databaseUser.Roles {
				r := databaseUser.Roles[i]
				role := RoleDefinition{
					CollectionName: &r.CollectionName,
					DatabaseName:   &r.DatabaseName,
					RoleName:       &r.RoleName,
				}

				roles = append(roles, role)
			}
			model.Roles = roles

			var labels []LabelDefinition

			for i := range databaseUser.Labels {
				l := databaseUser.Labels[i]
				label := LabelDefinition{
					Key:   &l.Key,
					Value: &l.Value,
				}
				labels = append(labels, label)
			}

			model.Labels = labels
			cfnid := fmt.Sprintf("%s-%s", databaseUser.Username, databaseUser.GroupID)

			model.UserCFNIdentifier = &cfnid
			dbUserModels = append(dbUserModels, model)
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  dbUserModels,
	}, nil
}

func getDBUser(roles []mongodbatlas.Role, groupID string, currentModel *Model, labels []mongodbatlas.Label, scopes []mongodbatlas.Scope) *mongodbatlas.DatabaseUser {
	return &mongodbatlas.DatabaseUser{
		Roles:           roles,
		GroupID:         groupID,
		Username:        *currentModel.Username,
		Password:        *currentModel.Password,
		DatabaseName:    *currentModel.DatabaseName,
		Labels:          labels,
		Scopes:          scopes,
		LDAPAuthType:    *currentModel.LdapAuthType,
		AWSIAMType:      *currentModel.AWSIAMType,
		X509Type:        *currentModel.X509Type,
		DeleteAfterDate: *currentModel.DeleteAfterDate,
	}
}

func setModel(currentModel *Model) (string, *mongodbatlas.DatabaseUser, handler.ProgressEvent, error) {
	var roles []mongodbatlas.Role
	for i := range currentModel.Roles {
		r := currentModel.Roles[i]
		role := mongodbatlas.Role{}
		if r.CollectionName != nil {
			role.CollectionName = *r.CollectionName
		}
		if r.DatabaseName != nil {
			role.DatabaseName = *r.DatabaseName
		}
		if r.RoleName != nil {
			role.RoleName = *r.RoleName
		}
		roles = append(roles, role)
	}
	_, _ = logger.Debugf("roles: %#+v", roles)

	var labels []mongodbatlas.Label
	for i := range currentModel.Labels {
		l := currentModel.Labels[i]
		label := mongodbatlas.Label{
			Key:   *l.Key,
			Value: *l.Value,
		}
		labels = append(labels, label)
	}
	_, _ = logger.Debugf("labels: %#+v", labels)

	var scopes []mongodbatlas.Scope
	for i := range currentModel.Scopes {
		s := currentModel.Scopes[i]
		scope := mongodbatlas.Scope{
			Name: *s.Name,
			Type: *s.Type,
		}
		scopes = append(scopes, scope)
	}
	_, _ = logger.Debugf("scopes: %#+v", scopes)

	groupID := *currentModel.ProjectId
	_, _ = logger.Debugf("groupID: %#+v", groupID)

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
		if (currentModel.LdapAuthType == &none) && (currentModel.AWSIAMType == &none) && (currentModel.X509Type == &none) {
			err := fmt.Errorf("password cannot be empty if not LDAP or IAM or X509: %v", currentModel)
			return "", nil, handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		s := ""
		currentModel.Password = &s
	}
	if (currentModel.X509Type != &none) || (currentModel.DeleteAfterDate == nil) {
		s := ""
		currentModel.DeleteAfterDate = &s
	}
	_, _ = logger.Debugf("Check Delete after date here::???????")
	user := getDBUser(roles, groupID, currentModel, labels, scopes)
	return groupID, user, handler.ProgressEvent{}, nil
}
