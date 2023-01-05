package resource

import (
	"context"
	"fmt"

	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

var ListRequiredFields = []string{constants.FederationSettingsID, constants.PvtKey, constants.PubKey}
var RequiredFields = []string{constants.FederationSettingsID, constants.OrgID, constants.PvtKey, constants.PubKey}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-FederatedSettingOrgConfigs")
}

func isExist(client *mongodbatlas.Client, currentModel *Model) bool {
	setup()

	federationSettingsID := *currentModel.FederationSettingsId
	orgID := *currentModel.OrgId
	fedSettingsConnectedOrg, _, err := client.FederatedSettings.GetConnectedOrg(context.Background(), federationSettingsID, orgID)
	return err == nil && fedSettingsConnectedOrg != nil
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// logic included for CFN Test starts
	if currentModel.TestMode != nil {
		_, err := util.PutKey(*currentModel.OrgId, "created", "x509", req.Session)
		if err != nil {
			return handler.ProgressEvent{}, err
		}
	}
	// logic included for CFN Test ends

	return Update(req, prevModel, currentModel)
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// logic included for CFN Test starts
	if currentModel.TestMode != nil && util.Get(*currentModel.OrgId, "x509", req.Session) == "" {
		return progressevents.GetFailedEventByCode("Resource Not Found",
			cloudformation.HandlerErrorCodeNotFound), nil
	}
	// logic included for CFN Test ends

	// Validate required fields in the request
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), err
	}

	// Check if  already exist
	if !isExist(client, currentModel) {
		return progressevents.GetFailedEventByCode("Resource Not Found", cloudformation.HandlerErrorCodeNotFound), nil
	}
	// Create Atlas API Request Object
	federationSettingsID := *currentModel.FederationSettingsId
	orgID := *currentModel.OrgId
	federatedSettingsOrg, resp, err := client.FederatedSettings.GetConnectedOrg(context.Background(), federationSettingsID, orgID)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	// convert response to client model
	currentModel = convertToUIModel(federatedSettingsOrg, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// logic included for CFN Test starts
	if currentModel.TestMode != nil && util.Get(*currentModel.OrgId, "x509", req.Session) == "" {
		return progressevents.GetFailedEventByCode("Resource Not Found",
			cloudformation.HandlerErrorCodeNotFound), nil
	}
	// logic included for CFN Test ends

	// Validate required fields in the request
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), err
	}

	// Check if  already exist
	if !isExist(client, currentModel) {
		return progressevents.GetFailedEventByCode("Resource Not Found", cloudformation.HandlerErrorCodeNotFound), nil
	}

	// Create Atlas API Request Object
	federationSettingsID := *currentModel.FederationSettingsId
	orgID := *currentModel.OrgId
	federatedSettingsOrgUpdate := convertToRequestModel(currentModel)
	_, resp, err := client.FederatedSettings.UpdateConnectedOrg(context.Background(), federationSettingsID, orgID, federatedSettingsOrgUpdate)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// logic included for CFN Test starts
	if currentModel.TestMode != nil {
		params := util.Get(*currentModel.OrgId, "x509", req.Session)
		if params == "" {
			return progressevents.GetFailedEventByCode("Resource Not Found", cloudformation.HandlerErrorCodeNotFound), nil
		}
		_, _ = util.DeleteKey(*currentModel.OrgId, "x509", req.Session)
		return handler.ProgressEvent{
			Message:         "Delete Complete",
			OperationStatus: handler.Success,
		}, nil
	}
	// logic included for CFN Test ends

	// Validate required fields in the request
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), err
	}

	// Check if  already exist
	if !isExist(client, currentModel) {
		return progressevents.GetFailedEventByCode("Resource Not Found", cloudformation.HandlerErrorCodeNotFound), nil
	}

	// Create Atlas API Request Object
	var res mongodbatlas.Response
	federationSettingsID := *currentModel.FederationSettingsId
	orgID := *currentModel.OrgId
	_, err = client.FederatedSettings.DeleteConnectedOrg(context.Background(), federationSettingsID, orgID)

	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	return handler.ProgressEvent{
		Message:         "Delete Complete",
		OperationStatus: handler.Success,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validate required fields in the request
	if modelValidation := validateModel(ListRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	federationSettingsID := *currentModel.FederationSettingsId
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}
	federatedSettingsOrgList, resp, err := client.FederatedSettings.ListConnectedOrgs(context.Background(), federationSettingsID, params)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}
	resultList := federatedSettingsOrgList.Results
	var models = make([]interface{}, len(resultList))
	if len(resultList) > 0 {
		for _, federatedSettingsOrg := range resultList {
			var model Model
			model.FederationSettingsId = &federationSettingsID
			models = append(models, convertToUIModel(federatedSettingsOrg, &model))
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModels:  models,
	}, nil
}

func convertToUIModel(responseModel *mongodbatlas.FederatedSettingsConnectedOrganization, currentModel *Model) *Model {
	if responseModel.DomainRestrictionEnabled != nil {
		currentModel.DomainRestrictionEnabled = responseModel.DomainRestrictionEnabled
	}
	if responseModel.RoleMappings != nil {
		currentModel.RoleMappings = flattenRoleMappings(responseModel.RoleMappings)
	}
	if responseModel.DomainAllowList != nil {
		currentModel.DomainAllowList = responseModel.DomainAllowList
	}
	if responseModel.PostAuthRoleGrants != nil {
		currentModel.PostAuthRoleGrants = responseModel.PostAuthRoleGrants
	}
	if responseModel.UserConflicts != nil {
		currentModel.UserConflicts = flattenUserConflict(*responseModel.UserConflicts)
	}
	currentModel.IdentityProviderId = &responseModel.IdentityProviderID

	return currentModel
}
func convertToRequestModel(currentModel *Model) *mongodbatlas.FederatedSettingsConnectedOrganization {
	federatedSettingsOrgUpdate := &mongodbatlas.FederatedSettingsConnectedOrganization{
		OrgID: *currentModel.OrgId,
	}
	federatedSettingsOrgUpdate.IdentityProviderID = cast.ToString(currentModel.IdentityProviderId)

	if currentModel.DomainRestrictionEnabled != nil {
		federatedSettingsOrgUpdate.DomainRestrictionEnabled = currentModel.DomainRestrictionEnabled
	}

	if len(currentModel.RoleMappings) != 0 {
		federatedSettingsOrgUpdate.RoleMappings = createRoleMappings(currentModel.RoleMappings)
	}
	if len(currentModel.DomainAllowList) != 0 {
		federatedSettingsOrgUpdate.DomainAllowList = currentModel.DomainAllowList
	}
	if len(currentModel.PostAuthRoleGrants) != 0 {
		federatedSettingsOrgUpdate.PostAuthRoleGrants = currentModel.PostAuthRoleGrants
	}
	return federatedSettingsOrgUpdate
}
func createRoleMappings(roleMappings []RoleMappingView) []*mongodbatlas.RoleMappings {
	results := make([]*mongodbatlas.RoleMappings, 0)
	for k := range roleMappings {
		resultRoleMapping := &mongodbatlas.RoleMappings{
			ExternalGroupName: cast.ToString(&roleMappings[k].ExternalGroupName),
			RoleAssignments:   createRoleAssignment(roleMappings[k].RoleAssignments),
		}
		results = append(results, resultRoleMapping)
	}
	return results
}

func createRoleAssignment(roleAssignments []RoleAssignment) []*mongodbatlas.RoleAssignments {
	results := make([]*mongodbatlas.RoleAssignments, len(roleAssignments))
	for k, roleMapping := range roleAssignments {
		if roleMapping.GroupId != nil {
			results[k] = &mongodbatlas.RoleAssignments{
				GroupID: cast.ToString(roleMapping.GroupId),
				Role:    cast.ToString(roleMapping.Role),
			}
		} else {
			results[k] = &mongodbatlas.RoleAssignments{
				OrgID: cast.ToString(roleMapping.OrgId),
				Role:  cast.ToString(roleMapping.Role),
			}
		}
	}
	return results
}
func flattenRoleMappings(roleMappings []*mongodbatlas.RoleMappings) []RoleMappingView {
	results := make([]RoleMappingView, 0)
	for k := range roleMappings {
		resultRoleMapping := RoleMappingView{
			Id:                &roleMappings[k].ID,
			ExternalGroupName: &roleMappings[k].ExternalGroupName,
			RoleAssignments:   flattenRoleAssignment(roleMappings[k].RoleAssignments),
		}
		results = append(results, resultRoleMapping)
	}
	return results
}

func flattenRoleAssignment(roleAssignments []*mongodbatlas.RoleAssignments) []RoleAssignment {
	results := make([]RoleAssignment, 0)
	for k := range roleAssignments {
		resultRoleAssignment := RoleAssignment{
			OrgId:   &roleAssignments[k].OrgID,
			GroupId: &roleAssignments[k].GroupID,
			Role:    &roleAssignments[k].Role,
		}
		results = append(results, resultRoleAssignment)
	}
	return results
}
func flattenUserConflict(conflicts mongodbatlas.UserConflicts) []FederatedUserView {
	results := make([]FederatedUserView, 0)
	for k := range conflicts {
		userView := FederatedUserView{
			EmailAddress:         &conflicts[k].EmailAddress,
			FederationSettingsId: &conflicts[k].FederationSettingsID,
			FirstName:            &conflicts[k].FirstName,
			LastName:             &conflicts[k].LastName,
			UserId:               &conflicts[k].UserID,
		}
		results = append(results, userView)
	}
	return results
}
