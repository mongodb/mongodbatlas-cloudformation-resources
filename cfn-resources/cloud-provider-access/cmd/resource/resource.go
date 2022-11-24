package resource

import (
	"context"
	"crypto/rand"
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
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
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
	desiredState := processJsonData(currentModel, constants.RolePolicyJson)
	_, errInRoleCreation := awsRoleCreationUsingCloudControlApi(req.Session, currentModel, desiredState)

	if errInRoleCreation != nil {
		log.Printf("Tring to use local role because of the following error \n %+v", err)
		localSession := session.Must(session.NewSession())
		_, errInRoleCreation = awsRoleCreationUsingCloudControlApi(localSession, currentModel, desiredState)
		if errInRoleCreation != nil {
			log.Debugf("Create - error: %+v", err)
			return handler.ProgressEvent{
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError,
				Message:          errInRoleCreation.Error(),
				OperationStatus:  handler.Failed,
			}, nil
		}
	}

	cloudProviderAuthorizationRequest := &mongodbatlas.CloudProviderAuthorizationRequest{
		ProviderName:      constants.AwsProviderName,
		IAMAssumedRoleARN: *currentModel.IamAssumedRoleArn,
	}

	role, res, err := client.CloudProviderAccess.AuthorizeRole(context.Background(), *currentModel.ProjectId, *currentModel.RoleId, cloudProviderAuthorizationRequest)
	if err != nil {
		log.Debugf("Update - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	currentModel.completeByConnection(role)

	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}

func awsRoleCreationUsingCloudControlApi(session *session.Session, currentModel *Model, desiredState string) (*Model, error) {
	clientCloudControl := cloudcontrolapi.New(session)
	typeName := constants.TypeName
	createResourceInput := &cloudcontrolapi.CreateResourceInput{
		DesiredState: &desiredState,
		TypeName:     &typeName,
	}
	outputCreate, err := clientCloudControl.CreateResource(createResourceInput)
	if err != nil {
		return currentModel, err
	}

	if !pollOutput(*outputCreate.ProgressEvent.RequestToken, clientCloudControl) {
		log.Warnf("Error while polling")
		return currentModel, errors.New("error in AWS Role Creation")
	}

	log.Warnf("Started Reading")

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

func pollOutput(reqToken string, clientCloudControl *cloudcontrolapi.CloudControlApi) bool {

	var i = 0
	var err error
	var status *cloudcontrolapi.GetResourceRequestStatusOutput
	for range time.Tick(time.Second * 5) {
		status, err = clientCloudControl.GetResourceRequestStatus(&cloudcontrolapi.GetResourceRequestStatusInput{
			RequestToken: aws.String(reqToken),
		})
		if err != nil {
			break
		}
		log.Warnf(*status.ProgressEvent.OperationStatus)
		if *status.ProgressEvent.OperationStatus == "SUCCESS" {
			log.Warnf("Role Creation Success")
			return true
		}
		if i == 4 {
			break
		}
		if *status.ProgressEvent.OperationStatus == "FAILED" {
			log.Warnf("Role Creation Failed")
			break
		}

		fmt.Println("Waiting for AWS Role Creation Status, Current Status is")
		i++
	}

	log.Warnf("%+v  Check error", err)
	return false
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
	randomNumber, _ := generateRandomNumber(10)

	roleName := constants.AwsRolePrefix + strconv.Itoa(randomNumber)
	log.Printf("the expected rolename is   %+v", roleName)

	return roleName
}

// generateRandomNumber generates random integer of n digits.
func generateRandomNumber(numberOfDigits int) (int, error) {
	maxLimit := int64(int(math.Pow10(numberOfDigits)) - 1)
	lowLimit := int(math.Pow10(numberOfDigits - 1))

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return 0, err
	}
	randomNumberInt := int(randomNumber.Int64())

	// Handling integers between 0, 10^(n-1) .. for n=4, handling cases between (0, 999)
	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}

	// Never likely to occur, kust for safe side.
	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}
	return randomNumberInt, nil
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

	foundRole := false
	for i := range roles.AWSIAMRoles {
		role := &(roles.AWSIAMRoles[i])
		if role.RoleID == *currentModel.RoleId && role.ProviderName == constants.AwsProviderName {
			currentModel.completeByConnection(role)
			foundRole = true
			break
		}
	}

	if !foundRole {
		log.Printf("The read returned no Matching result")
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
			Message:          "NotFound",
			OperationStatus:  handler.Failed,
		}, nil
	}

	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
		Message:         "Read Complete",
	}
	return event, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

//func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
//	log.Debugf("Update() currentModel:%+v", currentModel)
//
//	// Validation
//	modelValidation := validateModel(UpdateRequiredFields, currentModel)
//	if modelValidation != nil {
//		return *modelValidation, nil
//	}
//
//	userProvided := currentModel.IamAssumedRoleArn
//	readEvent, err := Read(req, prevModel, currentModel)
//	if readEvent.HandlerErrorCode == cloudformation.HandlerErrorCodeNotFound {
//		log.Printf("Didnt find the object while read. So we cant update")
//		return readEvent, nil
//	}
//	existingIamAssumedRoleArn := currentModel.IamAssumedRoleArn
//
//	if userProvided == existingIamAssumedRoleArn {
//		// Response
//		event := handler.ProgressEvent{
//			OperationStatus: handler.Success,
//			ResourceModel:   currentModel,
//			Message:         "No Change Detected, Update complete",
//		}
//		return event, nil
//	}
//	// Create atlas client
//	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
//	if err != nil {
//		log.Debugf("Update - error: %+v", err)
//		return handler.ProgressEvent{
//			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
//			Message:          err.Error(),
//			OperationStatus:  handler.Failed,
//		}, nil
//	}
//
//	cloudProviderAuthorizationRequest := &mongodbatlas.CloudProviderAuthorizationRequest{
//		ProviderName:      constants.AwsProviderName,
//		IAMAssumedRoleARN: *currentModel.IamAssumedRoleArn,
//	}
//
//	var res *mongodbatlas.Response
//
//	role, res, err := client.CloudProviderAccess.AuthorizeRole(context.Background(), *currentModel.ProjectId, *currentModel.RoleId, cloudProviderAuthorizationRequest)
//	if err != nil {
//		log.Debugf("Update - error: %+v", err)
//		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
//	}
//	log.Debugf("Atlas Client %v", client)
//
//	currentModel.completeByConnection(role)
//
//	// Response
//	event := handler.ProgressEvent{
//		OperationStatus: handler.Success,
//		ResourceModel:   currentModel,
//	}
//	return event, nil
//}

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

	_, errRead := Read(req, prevModel, currentModel)
	if errRead != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Yet to  implement deletion of role

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
