package resource

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{"FederationSettingsId", "OrgId", "ExternalGroupName", "RoleAssignments", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var ReadRequiredFields = []string{"FederationSettingsId", "Id", "OrgId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var UpdateRequiredFields = []string{"FederationSettingsId", "Id", "OrgId", "ExternalGroupName", "RoleAssignments", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var DeleteRequiredFields = []string{"FederationSettingsId", "Id", "OrgId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
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
		log.Debugf("Create - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	federationSettingsId := currentModel.FederationSettingsId
	orgId := currentModel.OrgId

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
	federatedSettingsOrganizationRoleMapping, resp, err := client.FederatedSettings.CreateRoleMapping(context.Background(), *federationSettingsId, *orgId, requestBody)

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			_, _ = log.Warnf("Create 404 err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = log.Warnf("create err: %+v", err)
		code := cloudformation.HandlerErrorCodeServiceInternalError
		if strings.Contains(err.Error(), "not exist") { // cfn test needs 404
			code = cloudformation.HandlerErrorCodeNotFound
		}
		if strings.Contains(err.Error(), "being deleted") {
			code = cloudformation.HandlerErrorCodeNotFound // cfn test needs 404
		}
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: code}, nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   roleMappingToModel(*currentModel, federatedSettingsOrganizationRoleMapping),
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Read() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Read - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	federationSettingsId := currentModel.FederationSettingsId
	orgId := currentModel.OrgId
	roleMappingID := currentModel.Id

	federatedSettingsOrganizationRoleMapping, resp, err := client.FederatedSettings.GetRoleMapping(context.Background(), *federationSettingsId, *orgId, *roleMappingID)

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Warnf("error 404- err:%+v resp:%+v", err, resp)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		log.Warnf("error cloud backup policy get- err:%+v resp:%+v", err, resp)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   roleMappingToModel(*currentModel, federatedSettingsOrganizationRoleMapping),
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Update() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Update - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	federationSettingsId := currentModel.FederationSettingsId
	orgId := currentModel.OrgId
	roleMappingId := currentModel.Id

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

	federatedSettingsOrganizationRoleMapping, resp, err := client.FederatedSettings.UpdateRoleMapping(context.Background(), *federationSettingsId, *orgId, *roleMappingId, requestBody)

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			_, _ = log.Warnf("Update 404 err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = log.Warnf("update err: %+v", err)
		code := cloudformation.HandlerErrorCodeServiceInternalError
		if strings.Contains(err.Error(), "not exist") { // cfn test needs 404
			code = cloudformation.HandlerErrorCodeNotFound
		}
		if strings.Contains(err.Error(), "being deleted") {
			code = cloudformation.HandlerErrorCodeNotFound // cfn test needs 404
		}
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: code}, nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   roleMappingToModel(*currentModel, federatedSettingsOrganizationRoleMapping),
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Delete() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Delete - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	federationSettingsId := currentModel.FederationSettingsId
	orgId := currentModel.OrgId
	roleMappingID := currentModel.Id

	resp, err := client.FederatedSettings.DeleteRoleMapping(context.Background(), *federationSettingsId, *orgId, *roleMappingID)

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Warnf("error 404- err:%+v resp:%+v", err, resp)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		log.Warnf("error cloud backup policy get- err:%+v resp:%+v", err, resp)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	log.Debugf("Atlas Client %v", client)
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("List() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("List - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	federationSettingsId := currentModel.FederationSettingsId
	orgId := currentModel.OrgId

	listOptions := &mongodbatlas.ListOptions{ItemsPerPage: 100, PageNum: 1}
	//listOptions := &mongodbatlas.ListOptions{
	//	PageNum:      *currentModel.ListOptions.PageNum,
	//	IncludeCount: *currentModel.ListOptions.IncludeCount,
	//	ItemsPerPage: *currentModel.ListOptions.ItemsPerPage,
	//}
	federatedSettingsOrganizationRoleMappings, resp, err := client.FederatedSettings.ListRoleMappings(context.Background(), *federationSettingsId, *orgId, listOptions)

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Warnf("error 404- err:%+v resp:%+v", err, resp)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		log.Warnf("error cloud backup policy get- err:%+v resp:%+v", err, resp)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	log.Debugf("Atlas Client %v", client)

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
	//if currentModel.Id != nil {
	//	roleMappingRequest.ID = *currentModel.Id
	//}
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
		var groupId string
		if assignments[i].GroupId != nil {
			groupId = *assignments[i].GroupId
		}
		var orgId string
		if assignments[i].OrgId != nil {
			orgId = *assignments[i].OrgId
		}
		roles[i] = &mongodbatlas.RoleAssignments{
			Role:    role,
			GroupID: groupId,
			OrgID:   orgId,
		}
	}
	fmt.Printf("roles: len %d %+v", len(roles), roles)
	return roles
}

func roleMappingToModel(currentModel Model, roleMapping *mongodbatlas.FederatedSettingsOrganizationRoleMapping) *Model {
	out := &Model{
		ApiKeys:              currentModel.ApiKeys,
		FederationSettingsId: currentModel.FederationSettingsId,
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
