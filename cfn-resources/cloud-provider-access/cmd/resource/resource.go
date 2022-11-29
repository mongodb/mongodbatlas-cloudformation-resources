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
	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

const (
	LocalSessionType   = "LOCAL_SESSION"
	RequestSessionType = "REQUEST_SESSION"
)

var awsRoleARNRegex = regexp.MustCompile(`arn:[a-z-]+:iam::(\d{12}):role/(.*)`)
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

	if req.CallbackContext != nil {
		return processCreateCallBack(req, currentModel, client)
	}
	var res *mongodbatlas.Response
	var roleResponse *mongodbatlas.AWSIAMRole

	cloudProviderAccessRequest := &mongodbatlas.CloudProviderAccessRoleRequest{
		ProviderName: constants.AwsProviderName,
	}
	// Step 1
	roleResponse, res, err = client.CloudProviderAccess.CreateRole(context.Background(), *currentModel.ProjectId, cloudProviderAccessRequest)
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	currentModel.completeByConnection(roleResponse)

	desiredState := processJsonData(currentModel, constants.RolePolicyJson)
	// Step 2
	event, errInRoleCreation := awsRoleCreationUsingCloudControlApi(req.Session, currentModel, desiredState, RequestSessionType)

	if errInRoleCreation != nil {
		fmt.Printf("Trying to use local role because of the following error \n %+v", errInRoleCreation)
		localSession := session.Must(session.NewSession())

		// Step 3
		event, errInRoleCreation = awsRoleCreationUsingCloudControlApi(localSession, currentModel, desiredState, LocalSessionType)
		if errInRoleCreation != nil {
			log.Debugf("Create - error: %+v", errInRoleCreation)
			return progressevents.GetFailedEventByCode(errInRoleCreation.Error(), cloudformation.HandlerErrorCodeServiceInternalError), errInRoleCreation
		}
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
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	// searching in roles
	if len(roles.AWSIAMRoles) == 0 {
		// Response
		log.Warnf("The read returned no result so nothing to return")
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
		log.Warnf("The read returned no Matching result")
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

	if req.CallbackContext != nil {
		return processDeleteCallBack(req, currentModel, client)
	}

	var res *mongodbatlas.Response

	_, errRead := Read(req, prevModel, currentModel)
	if errRead != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	cloudProviderDeAuthorizationRequest := &mongodbatlas.CloudProviderDeauthorizationRequest{
		ProviderName: constants.AwsProviderName,
		RoleID:       *currentModel.RoleId,
		GroupID:      *currentModel.ProjectId}

	res, err = client.CloudProviderAccess.DeauthorizeRole(context.Background(), cloudProviderDeAuthorizationRequest)
	if err != nil {
		log.Debugf("Delete - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	event, errInRoleDeletion := awsRoleDeletionUsingCloudControlApi(req.Session, currentModel, RequestSessionType)
	if errInRoleDeletion != nil {
		log.Warnf("Trying to use local role because of the following error \n %+v", errInRoleDeletion)
		localSession := session.Must(session.NewSession())
		event, errInRoleDeletion = awsRoleDeletionUsingCloudControlApi(localSession, currentModel, LocalSessionType)
		if errInRoleDeletion != nil {
			return progressevents.GetFailedEventByCode(errInRoleDeletion.Error(), cloudformation.HandlerErrorCodeServiceInternalError), errInRoleDeletion
		}
	}

	// Response
	return event, nil
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

func awsRoleCreationUsingCloudControlApi(session *session.Session, currentModel *Model, desiredState string, sessionType string) (handler.ProgressEvent, error) {
	log.Warnf("Inside AWS Role creation")

	clientCloudControl := cloudcontrolapi.New(session)
	typeName := constants.TypeName
	createResourceInput := &cloudcontrolapi.CreateResourceInput{
		DesiredState: &desiredState,
		TypeName:     &typeName,
	}
	outputCreate, err := clientCloudControl.CreateResource(createResourceInput)
	if err != nil {
		return progressevents.GetFailedEventByCode("AWS Role Creation error", *outputCreate.ProgressEvent.ErrorCode), err
	}
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Aws Role Creating..."),
		CallbackDelaySeconds: 2,
		ResourceModel:        currentModel,
		CallbackContext: map[string]interface{}{
			"Token":          outputCreate.ProgressEvent.RequestToken,
			"SessionType":    sessionType,
			"Identifier":     outputCreate.ProgressEvent.Identifier,
			"DesiredState":   desiredState,
			"AtlasRoleId":    currentModel.RoleId,
			"AtlasProjectId": currentModel.ProjectId,
		},
	}, nil

}

func authorizeRole(currentModel *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {
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

func processCreateCallBack(req handler.Request, currentModel *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {
	fmt.Println("Check for callback object here")
	spew.Dump(req.CallbackContext)

	//Variable Processing
	reqToken := fmt.Sprintf("%v", req.CallbackContext["Token"])
	identifier := fmt.Sprintf("%v", req.CallbackContext["Identifier"])
	roleSession := req.Session
	sessionType := fmt.Sprintf("%v", req.CallbackContext["SessionType"])
	projectId := fmt.Sprintf("%v", req.CallbackContext["AtlasProjectId"])
	roleId := fmt.Sprintf("%v", req.CallbackContext["AtlasRoleId"])
	desiredState := fmt.Sprintf("%v", req.CallbackContext["DesiredState"])
	currentModel.ProjectId = &projectId
	currentModel.RoleId = &roleId

	//Decision of Session type
	if sessionType == LocalSessionType {
		roleSession = session.Must(session.NewSession())
	}

	//Cloud Control API Client Creation
	clientCloudControl := cloudcontrolapi.New(roleSession)
	status, err := clientCloudControl.GetResourceRequestStatus(&cloudcontrolapi.GetResourceRequestStatusInput{
		RequestToken: aws.String(reqToken),
	})
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating aws role : %s", err.Error()),
			*status.ProgressEvent.ErrorCode), err
	}
	operationStatus := *status.ProgressEvent.OperationStatus
	switch operationStatus {
	case "SUCCESS":
		log.Warnf("AWS Role is created lets execute a GET request to fetch the Role ARN")
		roleArn := fetchAWSRole(identifier, clientCloudControl)
		currentModel.IamAssumedRoleArn = &roleArn
		return authorizeRole(currentModel, client)
	case "IN_PROGRESS":
		log.Warnf("Role Creation In-Progress")
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              fmt.Sprintf("Aws Role Creating..."),
			CallbackDelaySeconds: 2,
			ResourceModel:        currentModel,
			CallbackContext: map[string]interface{}{
				"Token":          reqToken,
				"SessionType":    sessionType,
				"Identifier":     identifier,
				"DesiredState":   desiredState,
				"AtlasRoleId":    currentModel.RoleId,
				"AtlasProjectId": currentModel.ProjectId,
			},
		}, nil
	case "FAILED":
		if sessionType == RequestSessionType {
			localSession := session.Must(session.NewSession())
			event, errInRoleCreation := awsRoleCreationUsingCloudControlApi(localSession, currentModel, desiredState, LocalSessionType)
			if errInRoleCreation != nil {
				log.Debugf("Create - error: %+v", errInRoleCreation)
				return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating aws role : %s", errInRoleCreation.Error()),
					*status.ProgressEvent.ErrorCode), errInRoleCreation
			}
			return event, nil
		}
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating aws role : %s", err.Error()),
			*status.ProgressEvent.ErrorCode), err
	}

	// If callback context is not null if operation status is not set we return error in creating
	return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating aws role : %s", err.Error()),
		*status.ProgressEvent.ErrorCode), err

}

func processDeleteCallBack(req handler.Request, currentModel *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {
	fmt.Println("Check for callback object here")
	spew.Dump(req.CallbackContext)

	//Variable Processing
	reqToken := fmt.Sprintf("%v", req.CallbackContext["Token"])
	identifier := fmt.Sprintf("%v", req.CallbackContext["Identifier"])
	roleSession := req.Session
	sessionType := fmt.Sprintf("%v", req.CallbackContext["SessionType"])
	iamAssumedRoleArn := fmt.Sprintf("%v", req.CallbackContext["IamAssumedRoleArn"])
	currentModel.IamAssumedRoleArn = &iamAssumedRoleArn

	//Decision of Session type
	if sessionType == LocalSessionType {
		roleSession = session.Must(session.NewSession())
	}

	//Cloud Control API Client Creation
	clientCloudControl := cloudcontrolapi.New(roleSession)
	status, err := clientCloudControl.GetResourceRequestStatus(&cloudcontrolapi.GetResourceRequestStatusInput{
		RequestToken: aws.String(reqToken),
	})
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating aws role : %s", err.Error()),
			*status.ProgressEvent.ErrorCode), err
	}
	operationStatus := *status.ProgressEvent.OperationStatus
	switch operationStatus {
	case "SUCCESS":
		log.Warnf("AWS Role is created lets execute a GET request to fetch the Role ARN")
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Delete Complete",
		}, nil
	case "IN_PROGRESS":
		log.Warnf("Role Deletion In-Progress")
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              fmt.Sprintf("Delete in progress."),
			CallbackDelaySeconds: 2,
			ResourceModel:        currentModel,
			CallbackContext: map[string]interface{}{
				"Token":             reqToken,
				"SessionType":       sessionType,
				"Identifier":        identifier,
				"IamAssumedRoleArn": iamAssumedRoleArn,
			},
		}, nil
	case "FAILED":
		if sessionType == RequestSessionType {
			localSession := session.Must(session.NewSession())
			event, errInRoleCreation := awsRoleDeletionUsingCloudControlApi(localSession, currentModel, LocalSessionType)
			if errInRoleCreation != nil {
				log.Debugf("Create - error: %+v", errInRoleCreation)
				return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating aws role : %s", errInRoleCreation.Error()),
					*status.ProgressEvent.ErrorCode), errInRoleCreation
			}
			return event, nil
		}
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating aws role : %s", err.Error()),
			*status.ProgressEvent.ErrorCode), err
	}

	// If callback context is not null if operation status is not set we return error in creating
	return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating aws role : %s", err.Error()),
		*status.ProgressEvent.ErrorCode), err

}

func fetchAWSRole(identifier string, clientCloudControl *cloudcontrolapi.CloudControlApi) string {
	typeName := constants.TypeName
	getResourceInput := &cloudcontrolapi.GetResourceInput{
		TypeName:   &typeName,
		Identifier: &identifier,
	}

	outputGet, err := clientCloudControl.GetResource(getResourceInput)
	if err != nil {
		return ""
	}

	roleArn, err := unmarshalRoleArn(outputGet)
	if err != nil {
		return ""
	}
	return roleArn
}

func awsRoleDeletionUsingCloudControlApi(session *session.Session, currentModel *Model, sessionType string) (handler.ProgressEvent, error) {
	clientCloudControl := cloudcontrolapi.New(session)
	roleName := roleFromRoleARN(*currentModel.IamAssumedRoleArn)
	if len(roleName) == 0 {
		errors.New("error in Deleting AWS Role")
	}
	typeName := constants.TypeName
	deleteResourceInput := &cloudcontrolapi.DeleteResourceInput{
		Identifier: &roleName,
		TypeName:   &typeName,
	}

	outputDelete, err := clientCloudControl.DeleteResource(deleteResourceInput)

	if err != nil {
		return progressevents.GetFailedEventByCode("AWS Role Deletion error", *outputDelete.ProgressEvent.ErrorCode), err
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Aws Role Creating..."),
		CallbackDelaySeconds: 2,
		ResourceModel:        currentModel,
		CallbackContext: map[string]interface{}{
			"Token":             outputDelete.ProgressEvent.RequestToken,
			"SessionType":       sessionType,
			"Identifier":        outputDelete.ProgressEvent.Identifier,
			"IamAssumedRoleArn": currentModel.IamAssumedRoleArn,
		},
	}, nil
}

//func pollDeleteOutput(reqToken string, clientCloudControl *cloudcontrolapi.CloudControlApi) bool {
//
//	var i = 0
//	var err error
//	var status *cloudcontrolapi.GetResourceRequestStatusOutput
//	for range time.Tick(time.Second * 5) {
//		status, err = clientCloudControl.GetResourceRequestStatus(&cloudcontrolapi.GetResourceRequestStatusInput{
//			RequestToken: aws.String(reqToken),
//		})
//		if err != nil {
//			break
//		}
//		log.Warnf(*status.ProgressEvent.OperationStatus)
//		if *status.ProgressEvent.OperationStatus == "SUCCESS" {
//			log.Warnf("Role Creation/Deletion Success")
//			return true
//		}
//		if i == 4 {
//			break
//		}
//		if *status.ProgressEvent.OperationStatus == "FAILED" {
//			log.Warnf("Role Creation/Deletion Failed")
//			break
//		}
//
//		fmt.Println("Waiting for AWS Role Creation/Deletion Status, Current Status is")
//		i++
//	}
//
//	log.Warnf("%+v  Check error", err)
//	return false
//}

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

func roleFromRoleARN(roleARN string) string {
	matches := awsRoleARNRegex.FindStringSubmatch(roleARN)

	// matches will contain ("roleARN", "accountID", "roleName")
	if len(matches) == 3 {
		return matches[2]
	}

	// Getting here means we failed to extract accountID and roleName from
	// roleARN. It should "not" happen, but if it does, return empty string
	// as accountID and roleARN as roleName instead.
	return ""
}

func unmarshalRoleArn(outputGet *cloudcontrolapi.GetResourceOutput) (string, error) {
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

func generateRandomRoleName() string {
	randomNumber, _ := generateRandomNumber(10)

	roleName := constants.AwsRolePrefix + strconv.Itoa(randomNumber)
	log.Warnf("the expected rolename is   %+v", roleName)

	return roleName
}
