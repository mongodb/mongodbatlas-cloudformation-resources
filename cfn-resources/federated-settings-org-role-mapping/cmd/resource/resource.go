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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
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

	_, _ = log.Debugf("Create() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = log.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId

	// preparing model request
	requestBody, _, _ := modelToRoleMappingRequest(currentModel)
	federatedSettingsOrganizationRoleMapping, resp, err := client.FederatedSettings.CreateRoleMapping(context.Background(), *federationSettingsID, *orgID, requestBody)
	if err != nil {
		_, _ = log.Warnf("error creating federated settings: %s", err)
		if resp.StatusCode == http.StatusBadRequest && strings.Contains(err.Error(), "DUPLICATE_ROLE_MAPPING") {
			return progressevents.GetFailedEventByCode("Resource already exists",
				cloudformation.HandlerErrorCodeAlreadyExists), nil
		}
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}
	currentModel.Id = &federatedSettingsOrganizationRoleMapping.ID
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   roleMappingToModel(*currentModel, federatedSettingsOrganizationRoleMapping),
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Read() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = log.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId
	roleMappingID := currentModel.Id

	federatedSettingsOrganizationRoleMapping, resp, err := client.FederatedSettings.GetRoleMapping(context.Background(), *federationSettingsID, *orgID, *roleMappingID)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}
	_, _ = log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   roleMappingToModel(*currentModel, federatedSettingsOrganizationRoleMapping),
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Update() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = log.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId
	roleMappingID := currentModel.Id

	// Check if  already exist
	if !isRoleMappingExists(currentModel, client) {
		return progressevents.GetFailedEventByCode("Not Found", cloudformation.HandlerErrorCodeNotFound), nil
	}
	if (currentModel.RoleAssignments) == nil || len(currentModel.RoleAssignments) == 0 {
		err := errors.New(RoleAssignementShouldBeSet)
		_, _ = log.Warnf("Create - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	// preparing model request
	requestBody, _, _ := modelToRoleMappingRequest(currentModel)
	federatedSettingsOrganizationRoleMapping, resp, err := client.FederatedSettings.UpdateRoleMapping(context.Background(), *federationSettingsID, *orgID, *roleMappingID, requestBody)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error updating federated settings : %s", err.Error()),
			resp.Response), nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   roleMappingToModel(*currentModel, federatedSettingsOrganizationRoleMapping),
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Delete() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = log.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	// Check if  already exist
	if !isRoleMappingExists(currentModel, client) {
		return progressevents.GetFailedEventByCode("Not Found", cloudformation.HandlerErrorCodeNotFound), nil
	}
	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId
	roleMappingID := currentModel.Id
	resp, err := client.FederatedSettings.DeleteRoleMapping(context.Background(), *federationSettingsID, *orgID, *roleMappingID)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error deleting federated settings : %s", err.Error()),
			resp.Response), nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("List() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = log.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId

	listOptions := &mongodbatlas.ListOptions{ItemsPerPage: 100, PageNum: 1}
	federatedSettingsOrganizationRoleMappings, resp, err := client.FederatedSettings.ListRoleMappings(context.Background(), *federationSettingsID, *orgID, listOptions)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting federated settings : %s", err.Error()),
			resp.Response), nil
	}

	models := make([]interface{}, 0) // cfn test
	for i := range federatedSettingsOrganizationRoleMappings.Results {
		model := Model{}
		model.Profile = currentModel.Profile
		model.OrgId = currentModel.OrgId
		model.FederationSettingsId = currentModel.FederationSettingsId
		model.Id = &federatedSettingsOrganizationRoleMappings.Results[i].ID
		model.ExternalGroupName = &federatedSettingsOrganizationRoleMappings.Results[i].ExternalGroupName
		model.RoleAssignments = flattenRoleAssignments(federatedSettingsOrganizationRoleMappings.Results[i].RoleAssignments)
		models = append(models, model)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List",
		ResourceModels:  models,
	}, nil
}

func modelToRoleMappingRequest(currentModel *Model) (*mongodbatlas.FederatedSettingsOrganizationRoleMapping, handler.ProgressEvent, error) {
	// Atlas client
	roleMappingRequest := &mongodbatlas.FederatedSettingsOrganizationRoleMapping{}
	if currentModel.Id != nil {
		roleMappingRequest.ID = *currentModel.Id
	}
	if currentModel.ExternalGroupName != nil {
		roleMappingRequest.ExternalGroupName = *currentModel.ExternalGroupName
	}
	if currentModel.RoleAssignments != nil {
		roleMappingRequest.RoleAssignments = expandRoleAssignments(currentModel.RoleAssignments)
	}
	return roleMappingRequest, handler.ProgressEvent{}, nil
}

func expandRoleAssignments(assignments []RoleAssignment) []*mongodbatlas.RoleAssignments {
	roles := make([]*mongodbatlas.RoleAssignments, len(assignments))
	for i := range assignments {
		var role string
		if assignments[i].Role != nil {
			role = *assignments[i].Role
		}
		var groupID string
		if assignments[i].ProjectId != nil {
			groupID = *assignments[i].ProjectId
		}
		var orgID string
		if assignments[i].OrgId != nil {
			orgID = *assignments[i].OrgId
		}
		roles[i] = &mongodbatlas.RoleAssignments{
			Role:    role,
			GroupID: groupID,
			OrgID:   orgID,
		}
	}
	fmt.Printf("roles: len %d %+v", len(roles), roles)
	return roles
}

func roleMappingToModel(currentModel Model, roleMapping *mongodbatlas.FederatedSettingsOrganizationRoleMapping) *Model {
	out := &Model{
		Profile:              currentModel.Profile,
		FederationSettingsId: currentModel.FederationSettingsId,
		OrgId:                currentModel.OrgId,
		Id:                   &roleMapping.ID,
		ExternalGroupName:    &roleMapping.ExternalGroupName,
		RoleAssignments:      flattenRoleAssignments(roleMapping.RoleAssignments),
	}
	return out
}

func flattenRoleAssignments(assignments []*mongodbatlas.RoleAssignments) []RoleAssignment {
	roleAssignments := make([]RoleAssignment, 0)
	for _, role := range assignments {
		roleAssignments = append(roleAssignments, RoleAssignment{
			Role:      &role.Role,
			OrgId:     &role.OrgID,
			ProjectId: &role.GroupID,
		})
	}
	return roleAssignments
}

func isRoleMappingExists(currentModel *Model, client *mongodbatlas.Client) bool {
	var isExists bool
	fedSettingsConnectedOrg, _, err := client.FederatedSettings.GetRoleMapping(context.Background(), *currentModel.FederationSettingsId, *currentModel.OrgId, *currentModel.Id)
	if err != nil {
		return isExists
	}
	if fedSettingsConnectedOrg != nil {
		isExists = true
	}
	return isExists
}
