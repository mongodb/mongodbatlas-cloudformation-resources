package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progress_event"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	log "github.com/sirupsen/logrus"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{"ProjectId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var ReadRequiredFields = []string{"ProjectId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey", "RoleId"}
var UpdateRequiredFields = []string{"ProjectId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey", "RoleId"}
var DeleteRequiredFields = []string{"ProjectId", "RoleId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey", "RoleId"}
var ListRequiredFields = []string{"ApiKeys.PrivateKey", "ApiKeys.PublicKey"}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func (m *Model) completeByConnection(c *mongodbatlas.AWSIAMRole) {
	m.RoleId = &c.RoleID
	m.IamAssumedRoleArn = &c.IAMAssumedRoleARN
	for ind, _ := range c.FeatureUsages {
		id := fmt.Sprintf("%v", c.FeatureUsages[ind].FeatureID)
		m.FeatureUsages = append(m.FeatureUsages, FeatureUsages{
			FeatureType: &c.FeatureUsages[ind].FeatureType,
			FeatureId:   &id,
		})
	}
	m.AuthorizedDate = &c.AuthorizedDate
	m.AtlasAWSAccountArn = &c.AtlasAWSAccountARN
	m.AtlasAssumedRoleExternalId = &c.AtlasAssumedRoleExternalID
	m.CreatedDate = &c.CreatedDate
	m.AuthorizedDate = &c.AuthorizedDate
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.Debugf("Create() currentModel:%+v", currentModel)

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
	var res *mongodbatlas.Response
	var roleResponse *mongodbatlas.AWSIAMRole

	cloudProviderAccessRequest := &mongodbatlas.CloudProviderAccessRoleRequest{
		ProviderName: constants.AwsProviderName,
	}
	roleResponse, res, err = client.CloudProviderAccess.CreateRole(context.Background(), *currentModel.ProjectId, cloudProviderAccessRequest)
	currentModel.completeByConnection(roleResponse)
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
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
	var res *mongodbatlas.Response
	var roles *mongodbatlas.CloudProviderAccessRoles
	roles, res, err = client.CloudProviderAccess.ListRoles(context.Background(), *currentModel.ProjectId)
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	// searching in roles
	if len(roles.AWSIAMRoles) == 0 {
		// Response
		log.Printf("The read returned no result so nothing to return")
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
			Message:          "NotFound",
			OperationStatus:  handler.Failed,
		}, nil
	}
	for i := range roles.AWSIAMRoles {
		role := &(roles.AWSIAMRoles[i])
		if role.RoleID == *currentModel.RoleId && role.ProviderName == constants.AwsProviderName {
			currentModel.completeByConnection(role)
			break
		}
	}

	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
		Message:         "Read Complete",
	}
	return event, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.Debugf("Update() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	userProvided := currentModel.IamAssumedRoleArn
	readEvent, err := Read(req, prevModel, currentModel)
	if readEvent.HandlerErrorCode == cloudformation.HandlerErrorCodeNotFound {
		log.Printf("Didnt find the object while read. So we cant update")
		return readEvent, nil
	}
	existingIamAssumedRoleArn := currentModel.IamAssumedRoleArn

	if userProvided == existingIamAssumedRoleArn {
		// Response
		event := handler.ProgressEvent{
			OperationStatus: handler.Success,
			ResourceModel:   currentModel,
			Message:         "No Change Detected, Update complete",
		}
		return event, nil
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

	cloudProviderAuthorizationRequest := &mongodbatlas.CloudProviderAuthorizationRequest{
		ProviderName:      constants.AwsProviderName,
		IAMAssumedRoleARN: *currentModel.IamAssumedRoleArn,
	}

	var res *mongodbatlas.Response

	role, res, err := client.CloudProviderAccess.AuthorizeRole(context.Background(), *currentModel.ProjectId, *currentModel.RoleId, cloudProviderAuthorizationRequest)
	if err != nil {
		log.Debugf("Update - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	currentModel.completeByConnection(role)

	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
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
	var res *mongodbatlas.Response

	cloudProviderDeAuthorizationRequest := &mongodbatlas.CloudProviderDeauthorizationRequest{
		ProviderName: constants.AwsProviderName,
		RoleID:       *currentModel.RoleId,
		GroupID:      *currentModel.ProjectId}

	res, err = client.CloudProviderAccess.DeauthorizeRole(context.Background(), cloudProviderDeAuthorizationRequest)
	if err != nil {
		log.Debugf("Delete - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
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

	var res *mongodbatlas.Response
	var roles *mongodbatlas.CloudProviderAccessRoles
	roles, res, err = client.CloudProviderAccess.ListRoles(context.Background(), *currentModel.ProjectId)
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	// searching in roles
	mm := make([]interface{}, 0)
	for i := range roles.AWSIAMRoles {
		var m Model
		role := &(roles.AWSIAMRoles[i])
		if role.ProviderName == constants.AwsProviderName {
			m.completeByConnection(role)
			mm = append(mm, m)
		}
	}
	if err != nil {
		log.Debugf("List - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List successful",
		ResourceModels:  mm,
	}
	return event, nil
}
