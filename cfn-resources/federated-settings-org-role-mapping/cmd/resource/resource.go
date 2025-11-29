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
	"errors"
	"fmt"
	"net/http"
	"strings"

	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

var CreateRequiredFields = []string{constants.FederationSettingsID, constants.OrgID, constants.ExternalGroupName, constants.RoleAssignments}
var ReadRequiredFields = []string{constants.FederationSettingsID, constants.ID, constants.OrgID}
var UpdateRequiredFields = []string{constants.FederationSettingsID, constants.OrgID, constants.ID, constants.ExternalGroupName, constants.RoleAssignments}
var DeleteRequiredFields = []string{constants.FederationSettingsID, constants.OrgID}
var ListRequiredFields = []string{constants.FederationSettingsID, constants.OrgID}

const (
	RoleAssignementShouldBeSet = "error creating federated settings org role mapping: RoleAssignments should be set when `Export` is set"
)

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-FederatedSettingsOrgRoleMapping")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId

	requestBody, _, _ := modelToRoleMappingRequest(currentModel)
	federatedSettingsOrganizationRoleMapping, resp, err := client.Atlas20231115002.FederatedAuthenticationApi.CreateRoleMapping(context.Background(), *federationSettingsID, *orgID, requestBody).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusBadRequest && strings.Contains(err.Error(), "DUPLICATE_ROLE_MAPPING") {
			return progressevent.GetFailedEventByCode("Resource already exists",
				string(types.HandlerErrorCodeAlreadyExists)), nil
		}
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp), nil
	}
	currentModel.Id = federatedSettingsOrganizationRoleMapping.Id
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   roleMappingToModel(*currentModel, federatedSettingsOrganizationRoleMapping),
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId
	roleMappingID := currentModel.Id

	federatedSettingsOrganizationRoleMapping, resp, err := client.Atlas20231115002.FederatedAuthenticationApi.
		GetRoleMapping(context.Background(), *federationSettingsID, *roleMappingID, *orgID).
		Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   roleMappingToModel(*currentModel, federatedSettingsOrganizationRoleMapping),
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId
	roleMappingID := currentModel.Id

	if !isRoleMappingExists(currentModel, client) {
		return progressevent.GetFailedEventByCode("Not Found", string(types.HandlerErrorCodeNotFound)), nil
	}

	if (currentModel.RoleAssignments) == nil || len(currentModel.RoleAssignments) == 0 {
		err := errors.New(RoleAssignementShouldBeSet)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest)}, nil
	}
	// preparing model request
	requestBody, _, _ := modelToRoleMappingRequest(currentModel)
	federatedSettingsOrganizationRoleMapping, resp, err := client.Atlas20231115002.FederatedAuthenticationApi.
		UpdateRoleMapping(context.Background(), *federationSettingsID, *roleMappingID, *orgID, requestBody).
		Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error updating federated settings : %s", err.Error()),
			resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   roleMappingToModel(*currentModel, federatedSettingsOrganizationRoleMapping),
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	// Check if  already exist
	if !isRoleMappingExists(currentModel, client) {
		return progressevent.GetFailedEventByCode("Not Found", string(types.HandlerErrorCodeNotFound)), nil
	}

	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId
	roleMappingID := currentModel.Id
	resp, err := client.Atlas20231115002.FederatedAuthenticationApi.
		DeleteRoleMapping(context.Background(), *federationSettingsID, *roleMappingID, *orgID).
		Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error deleting federated settings : %s", err.Error()),
			resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId

	federatedSettingsOrganizationRoleMappings, resp, err := client.Atlas20231115002.
		FederatedAuthenticationApi.
		ListRoleMappings(context.Background(), *federationSettingsID, *orgID).
		Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error getting federated settings : %s", err.Error()),
			resp), nil
	}

	models := make([]interface{}, 0)
	for i := range federatedSettingsOrganizationRoleMappings.Results {
		model := Model{}
		model.Profile = currentModel.Profile
		model.OrgId = currentModel.OrgId
		model.FederationSettingsId = currentModel.FederationSettingsId
		model.Id = federatedSettingsOrganizationRoleMappings.Results[i].Id
		model.ExternalGroupName = &federatedSettingsOrganizationRoleMappings.Results[i].ExternalGroupName
		model.RoleAssignments = flattenRoleAssignments(federatedSettingsOrganizationRoleMappings.Results[i].RoleAssignments)
		models = append(models, model)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}

func modelToRoleMappingRequest(currentModel *Model) (*admin20231115002.AuthFederationRoleMapping, handler.ProgressEvent, error) {
	roleMappingRequest := &admin20231115002.AuthFederationRoleMapping{}
	if currentModel.Id != nil {
		roleMappingRequest.Id = currentModel.Id
	}
	if currentModel.ExternalGroupName != nil {
		roleMappingRequest.ExternalGroupName = *currentModel.ExternalGroupName
	}
	if currentModel.RoleAssignments != nil {
		roleMappingRequest.RoleAssignments = expandRoleAssignments(currentModel.RoleAssignments)
	}
	return roleMappingRequest, handler.ProgressEvent{}, nil
}

func expandRoleAssignments(assignments []RoleAssignment) []admin20231115002.RoleAssignment {
	roles := make([]admin20231115002.RoleAssignment, len(assignments))
	for i := range assignments {
		role := admin20231115002.RoleAssignment{}
		if util.IsStringPresent(assignments[i].Role) {
			role.Role = assignments[i].Role
		}

		if util.IsStringPresent(assignments[i].ProjectId) {
			role.GroupId = assignments[i].ProjectId
		}

		if util.IsStringPresent(assignments[i].OrgId) {
			role.OrgId = assignments[i].OrgId
		}
		roles[i] = role
	}

	return roles
}

func roleMappingToModel(currentModel Model, roleMapping *admin20231115002.AuthFederationRoleMapping) *Model {
	out := &Model{
		Profile:              currentModel.Profile,
		FederationSettingsId: currentModel.FederationSettingsId,
		OrgId:                currentModel.OrgId,
		Id:                   roleMapping.Id,
		ExternalGroupName:    &roleMapping.ExternalGroupName,
		RoleAssignments:      flattenRoleAssignments(roleMapping.RoleAssignments),
	}
	return out
}

func flattenRoleAssignments(assignments []admin20231115002.RoleAssignment) []RoleAssignment {
	roleAssignments := make([]RoleAssignment, 0)
	for _, role := range assignments {
		roleAssignments = append(roleAssignments, RoleAssignment{
			Role:      role.Role,
			OrgId:     role.OrgId,
			ProjectId: role.GroupId,
		})
	}
	return roleAssignments
}

func isRoleMappingExists(currentModel *Model, client *util.MongoDBClient) bool {
	var isExists bool
	fedSettingsConnectedOrg, _, err := client.Atlas20231115002.FederatedAuthenticationApi.
		GetRoleMapping(context.Background(), *currentModel.FederationSettingsId, *currentModel.Id, *currentModel.OrgId).
		Execute()
	if err != nil {
		return isExists
	}
	if fedSettingsConnectedOrg != nil {
		isExists = true
	}
	return isExists
}
