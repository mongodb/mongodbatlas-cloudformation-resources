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
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func setup() {
	util.SetupLogger("mongodb-atlas-custom-db-role")
}

var CreateRequiredFields = []string{constants.ProjectID, constants.RoleName}
var ReadRequiredFields = []string{constants.ProjectID, constants.RoleName}
var UpdateRequiredFields = []string{constants.ProjectID, constants.RoleName}
var DeleteRequiredFields = []string{constants.ProjectID, constants.RoleName}
var ListRequiredFields = []string{constants.ProjectID}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	atlasCustomDBRole := currentModel.ToCustomDBRole()
	customDBRole, response, err := client.Atlas20231115002.CustomDatabaseRolesApi.CreateCustomDatabaseRole(context.Background(), *currentModel.ProjectId, atlasCustomDBRole).Execute()
	if err != nil {
		if apiError, ok := admin20231115002.AsError(err); ok && *apiError.Error == http.StatusConflict {
			return progress_events.GetFailedEventByCode("Resource already exists",
				cloudformation.HandlerErrorCodeAlreadyExists), nil
		}

		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
			response), nil
	}

	currentModel.completeByAtlasRole(*customDBRole)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   currentModel}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	atlasCustomDdRole, response, err := client.Atlas20231115002.CustomDatabaseRolesApi.GetCustomDatabaseRole(context.Background(), *currentModel.ProjectId, *currentModel.RoleName).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response), nil
	}

	currentModel.completeByAtlasRole(*atlasCustomDdRole)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Get successful",
		ResourceModel:   currentModel}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	var actions []admin20231115002.DatabasePrivilegeAction
	for _, a := range currentModel.Actions {
		actions = append(actions, a.toAtlasAction())
	}

	var inheritedRoles []admin20231115002.DatabaseInheritedRole
	for _, ir := range currentModel.InheritedRoles {
		inheritedRoles = append(inheritedRoles, ir.toAtlasInheritedRole())
	}

	inputCustomDBRole := admin20231115002.UpdateCustomDBRole{
		Actions:        actions,
		InheritedRoles: inheritedRoles,
	}

	atlasCustomDdRole, response, err := client.Atlas20231115002.CustomDatabaseRolesApi.UpdateCustomDatabaseRole(context.Background(), *currentModel.ProjectId,
		*currentModel.RoleName, &inputCustomDBRole).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response), nil
	}

	currentModel.completeByAtlasRole(*atlasCustomDdRole)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update successful",
		ResourceModel:   currentModel}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	response, err := client.Atlas20231115002.CustomDatabaseRolesApi.DeleteCustomDatabaseRole(context.Background(), *currentModel.ProjectId, *currentModel.RoleName).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error deleting resource : %s", err.Error()),
			response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete success"}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	customDBRoleResponse, response, err := client.Atlas20231115002.CustomDatabaseRolesApi.ListCustomDatabaseRoles(context.Background(),
		*currentModel.ProjectId).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error listing resource : %s", err.Error()),
			response), nil
	}

	mm := make([]interface{}, 0)
	for _, customDBRole := range customDBRoleResponse {
		var m Model
		m.completeByAtlasRole(customDBRole)
		m.ProjectId = currentModel.ProjectId
		m.Profile = currentModel.Profile
		mm = append(mm, m)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List successful",
		ResourceModels:  mm}, nil
}

func (m *Model) ToCustomDBRole() *admin20231115002.UserCustomDBRole {
	var actions []admin20231115002.DatabasePrivilegeAction
	for _, a := range m.Actions {
		actions = append(actions, a.toAtlasAction())
	}

	var inheritedRoles []admin20231115002.DatabaseInheritedRole
	for _, ir := range m.InheritedRoles {
		inheritedRoles = append(inheritedRoles, ir.toAtlasInheritedRole())
	}

	return &admin20231115002.UserCustomDBRole{
		Actions:        actions,
		InheritedRoles: inheritedRoles,
		RoleName:       *m.RoleName,
	}
}

func (a InheritedRole) toAtlasInheritedRole() admin20231115002.DatabaseInheritedRole {
	return admin20231115002.DatabaseInheritedRole{
		Db:   *a.Db,
		Role: *a.Role,
	}
}

func (a Action) toAtlasAction() admin20231115002.DatabasePrivilegeAction {
	var resources []admin20231115002.DatabasePermittedNamespaceResource
	for _, r := range a.Resources {
		resources = append(resources, r.toAtlasResource())
	}

	return admin20231115002.DatabasePrivilegeAction{
		Action:    *a.Action,
		Resources: resources,
	}
}

func (r Resource) toAtlasResource() admin20231115002.DatabasePermittedNamespaceResource {
	out := admin20231115002.DatabasePermittedNamespaceResource{
		Cluster: false,
	}
	if r.Collection != nil {
		out.Collection = *r.Collection
	}

	if r.DB != nil {
		out.Db = *r.DB
	}

	if r.Cluster != nil && *r.Cluster {
		out.Cluster = *r.Cluster
	}

	return out
}

func (m *Model) completeByAtlasRole(role admin20231115002.UserCustomDBRole) {
	var actions []Action
	for _, a := range role.Actions {
		actions = append(actions, atlasActionToModel(a))
	}

	var inheritedRoles []InheritedRole
	for _, ir := range role.InheritedRoles {
		inheritedRoles = append(inheritedRoles, atlasInheritedRoleToModel(ir))
	}

	m.Actions = actions
	m.InheritedRoles = inheritedRoles
	m.RoleName = &role.RoleName
}

func atlasActionToModel(action admin20231115002.DatabasePrivilegeAction) Action {
	var resources []Resource
	for _, r := range action.Resources {
		resources = append(resources, atlasResourceToModel(r))
	}

	return Action{
		Action:    &action.Action,
		Resources: resources,
	}
}

func atlasResourceToModel(resource admin20231115002.DatabasePermittedNamespaceResource) Resource {
	return Resource{
		Collection: &resource.Collection,
		DB:         &resource.Db,
		Cluster:    &resource.Cluster,
	}
}

func atlasInheritedRoleToModel(inheritedRole admin20231115002.DatabaseInheritedRole) InheritedRole {
	return InheritedRole{
		Db:   &inheritedRole.Db,
		Role: &inheritedRole.Role,
	}
}
