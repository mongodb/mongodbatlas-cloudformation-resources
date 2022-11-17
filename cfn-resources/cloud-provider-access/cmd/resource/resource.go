package resource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudcontrolapi"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	log "github.com/sirupsen/logrus"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	"math/rand"
	"strings"
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
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	currentModel.completeByConnection(roleResponse)
	//_, errInRoleCreation := awsRoleCreationUsingCloudControlApi(req.Session, currentModel)
	_, errInRoleCreation := awsRoleCreationGoSdk(req, currentModel)

	if errInRoleCreation != nil {
		log.Debugf("Create - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError,
			Message:          errInRoleCreation.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}

func awsRoleCreationUsingCloudControlApi(session *session.Session, currentModel *Model) (*Model, error) {
	clientCloudControl := cloudcontrolapi.New(session)
	typeName := constants.TypeName
	desiredState := processJsonData(currentModel, constants.RolePolicyJson)

	createResourceInput := &cloudcontrolapi.CreateResourceInput{
		DesiredState: &desiredState,
		TypeName:     &typeName,
	}
	outputCreate, err := clientCloudControl.CreateResource(createResourceInput)
	if err != nil {
		return currentModel, err
	}

	getResourceInput := &cloudcontrolapi.GetResourceInput{
		TypeName:   &typeName,
		Identifier: outputCreate.ProgressEvent.Identifier,
	}

	outputGet, err := clientCloudControl.GetResource(getResourceInput)
	if err != nil {
		return currentModel, err
	}

	roleArn, err := getRoleArn(outputGet)
	if err != nil {
		return currentModel, err
	}

	currentModel.IamAssumedRoleArn = &roleArn
	return currentModel, nil
}

func awsRoleCreationGoSdk(req handler.Request, currentModel *Model) (*Model, error) {

	iamAwsRegion := currentModel.IamAwsRegion
	clientIam := iam.New(req.Session, aws.NewConfig().WithRegion(*iamAwsRegion))
	assumedRoleJson := processJsonData(currentModel, constants.RolePolicyJsonGoSdk)
	roleName := generateRandomRoleName()
	fmt.Printf("RoleName processed %+v", roleName)
	createRoleInput := iam.CreateRoleInput{
		AssumeRolePolicyDocument: &assumedRoleJson,
		RoleName:                 &roleName,
	}

	iamRole, err := clientIam.CreateRole(&createRoleInput)
	if err != nil {
		return currentModel, err
	}
	currentModel.IamAssumedRoleArn = iamRole.Role.Arn
	return currentModel, nil
}

func getRoleArn(outputGet *cloudcontrolapi.GetResourceOutput) (string, error) {
	if outputGet != nil {
		if outputGet.ResourceDescription != nil {
			if outputGet.ResourceDescription.Properties != nil {
				var IAMRoleARN IAMRoleArn
				json.Unmarshal([]byte(*outputGet.ResourceDescription.Properties), &IAMRoleARN)
				fmt.Printf("%+v", IAMRoleARN.Arn)
				if IAMRoleARN.Arn != nil {
					return *IAMRoleARN.Arn, nil
				}
			}
		}
	}
	return "Error in Fetching the Arn from Response", errors.New("error in Fetching the Arn from Response")
}

type IAMRoleArn struct {
	_   struct{} `type:"structure"`
	Arn *string  `min:"20" type:"string"`
}

//func createRoleInAWS(req *handler.Request, currentModel *Model, err error) *iam.CreateRoleOutput {
//	iamAwsRegion := currentModel.IamAwsRegion
//	fmt.Println("whats the req here")
//	//spew.Dump(req.Session)
//	clientIam := iam.New(req.Session, aws.NewConfig().WithRegion(*iamAwsRegion))
//	fmt.Println("Session created")
//	assumedRoleJson := processJsonData(currentModel)
//	fmt.Printf("json processed %+v", assumedRoleJson)
//	roleName := generateRandomRoleName()
//	fmt.Printf("RoleName processed %+v", roleName)
//	createRoleInput := iam.CreateRoleInput{
//		AssumeRolePolicyDocument: &assumedRoleJson,
//		RoleName:                 &roleName,
//	}
//	fmt.Println("Create Role input processed")
//	spew.Dump(createRoleInput)
//	iamRole, err := clientIam.CreateRole(&createRoleInput)
//	if err != nil {
//		spew.Dump(err)
//	}
//	fmt.Println("Role Created")
//
//	spew.Dump(iamRole.Role.Arn)
//	return iamRole
//}

func generateRandomRoleName() string {
	var letterRunes = []rune("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 13)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	roleName := constants.AwsRolePrefix + string(b)
	return roleName
}

func processJsonData(currentModel *Model, RolePolicyJson string) string {
	// an arbitrary json string

	RolePolicyJson = strings.Replace(RolePolicyJson, "$ATLAS_AWS_ACCOUNT_ARN", *currentModel.AtlasAWSAccountArn, 1)
	RolePolicyJson = strings.Replace(RolePolicyJson, "$ATLAS_ASSUMEDROLE_EXTERNAL_ID", *currentModel.AtlasAssumedRoleExternalId, 1)
	RolePolicyJson = strings.Replace(RolePolicyJson, "$ROLE_NAME", generateRandomRoleName(), 1)

	return RolePolicyJson
}

//func processJsonData(currentModel *Model) string {
//	// an arbitrary json string
//
//	RolePolicyJsonUpdated := constants.RolePolicyJson
//	RolePolicyJsonUpdated = strings.Replace(RolePolicyJsonUpdated, "$ATLAS_AWS_ACCOUNT_ARN", *currentModel.AtlasAWSAccountArn, 1)
//	RolePolicyJsonUpdated = strings.Replace(RolePolicyJsonUpdated, "$ATLAS_ASSUMEDROLE_EXTERNAL_ID", *currentModel.AtlasAssumedRoleExternalId, 1)
//	RolePolicyJsonUpdated = strings.Replace(RolePolicyJsonUpdated, "$ROLE_NAME", generateRandomRoleName(), 1)
//
//	return RolePolicyJsonUpdated
//}

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
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
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
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
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
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
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
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
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
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
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
