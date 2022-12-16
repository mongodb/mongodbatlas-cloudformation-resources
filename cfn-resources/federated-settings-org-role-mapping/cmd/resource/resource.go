package resource

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{"FederationSettingsId", "OrgId", "ExternalGroupName", "RoleAssignments", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var ReadRequiredFields = []string{"FederationSettingsId", "Id", "OrgId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var UpdateRequiredFields = []string{"FederationSettingsId", "OrgId", "Id", "ExternalGroupName", "RoleAssignments", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var DeleteRequiredFields = []string{"FederationSettingsId", "OrgId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var ListRequiredFields = []string{"ApiKeys.PrivateKey", "ApiKeys.PublicKey"}

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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId

	if (currentModel.ExternalGroupName) == nil {
		err := errors.New("error creating federated settings org role mapping: ExternalGroupName should be set when `Export` is set")
		_, _ = log.Warnf("Create - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	if (currentModel.RoleAssignments) == nil || len(currentModel.RoleAssignments) == 0 {
		err := errors.New("error creating federated settings org role mapping: RoleAssignments should be set when `Export` is set")
		_, _ = log.Warnf("Create - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Read - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Update - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId
	roleMappingID := currentModel.Id

	// Check if  already exist
	if !isRoleMappingExists(currentModel, client) {
		return progressevents.GetFailedEventByCode("Not Found", cloudformation.HandlerErrorCodeNotFound), nil
	}
	if (currentModel.RoleAssignments) == nil || len(currentModel.RoleAssignments) == 0 {
		err := errors.New("error creating federated settings org role mapping: RoleAssignments should be set when `Export` is set")
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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Delete - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("List - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	federationSettingsID := currentModel.FederationSettingsId
	orgID := currentModel.OrgId

	listOptions := &mongodbatlas.ListOptions{ItemsPerPage: 100, PageNum: 1}
	federatedSettingsOrganizationRoleMappings, resp, err := client.FederatedSettings.ListRoleMappings(context.Background(), *federationSettingsID, *orgID, listOptions)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting federated settings : %s", err.Error()),
			resp.Response), nil
	}

	models := make([]*Model, federatedSettingsOrganizationRoleMappings.TotalCount)
	for i := range federatedSettingsOrganizationRoleMappings.Results {
		model := &Model{}
		roleMappingToModel(*model, federatedSettingsOrganizationRoleMappings.Results[i])
		models = append(models, model)
	}
	return handler.ProgressEvent{
		OperationStatus:  handler.Success,
		Message:          "List",
		ResourceModel:    models,
		HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
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
		if assignments[i].GroupId != nil {
			groupID = *assignments[i].GroupId
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
		ApiKeys:              currentModel.ApiKeys,
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
			Role:    &role.Role,
			OrgId:   &role.OrgID,
			GroupId: &role.GroupID,
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
