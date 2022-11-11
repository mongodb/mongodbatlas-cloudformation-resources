package resource

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/openlyinc/pointy"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	publicKey              = "ApiKeys.PublicKey"
	privateKey             = "ApiKeys.PrivateKey"
	projectID              = "ProjectId"
	customMasterKey        = "AwsKms.CustomerMasterKeyID"
	roleID                 = "AwsKms.RoleID"
	region                 = "AwsKms.Region"
	errorCreateMongoClient = "error - Create MongoDB Client- Details: %+v"
	errorCreateEncryption  = "error - Create Encryption  for Project(%s)- Details: %+v"
)

var CreateRequiredFields = []string{publicKey, privateKey, roleID, customMasterKey, region, projectID}
var ReadRequiredFields = []string{publicKey, privateKey, projectID}
var DeleteRequiredFields = []string{publicKey, privateKey, projectID}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	_, _ = logger.Debugf("Create - Encryption for Request() currentModel:%+v", currentModel)

	// Validate required fields in the request
	if modelValidation := validateModel(CreateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(errorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// Create Atlas API Request Object
	_, _ = logger.Debug("Create - Encryption at rest starts ")
	projectID := *currentModel.ProjectId
	encryptionAtRest := &mongodbatlas.EncryptionAtRest{
		AwsKms: mongodbatlas.AwsKms{
			Enabled:             currentModel.AwsKms.Enabled,
			CustomerMasterKeyID: *currentModel.AwsKms.CustomerMasterKeyID,
			RoleID:              *currentModel.AwsKms.RoleID,
			Region:              *currentModel.AwsKms.Region,
		},
		GroupID: projectID,
	}
	deploySecretString, _ := json.Marshal(encryptionAtRest)
	log.Printf("Response Object: %s", deploySecretString)

	// API call to create
	_, _, err = client.EncryptionsAtRest.Create(context.Background(), encryptionAtRest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf(errorCreateEncryption, projectID, err)
	}
	currentModel.Id = pointy.String(strconv.FormatInt(randInt64(), 10))

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	_, _ = logger.Debugf("Read snapshot for Request() currentModel:%+v", currentModel)

	// Validate required fields in the request
	if modelValidation := validateModel(ReadRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(errorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// Check if already  exist
	projectID := *currentModel.ProjectId
	if !isExist(currentModel) {
		_, _ = logger.Debugf("Read - error in encryption at rest with id(%s)", projectID)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	// API call
	encryptionAtRest, _, err := client.EncryptionsAtRest.Get(context.Background(), projectID)
	if err != nil {
		return handler.NewProgressEvent(), fmt.Errorf("error fetching encryption at rest configuration for project (%s): %s", projectID, err)
	}
	currentModel.AwsKms.CustomerMasterKeyID = &encryptionAtRest.AwsKms.CustomerMasterKeyID
	currentModel.AwsKms.Enabled = encryptionAtRest.AwsKms.Enabled
	currentModel.AwsKms.RoleID = &encryptionAtRest.AwsKms.RoleID
	currentModel.AwsKms.Region = &encryptionAtRest.AwsKms.Region

	currentModelString, _ := json.Marshal(currentModel)
	log.Printf("Response Object: %s", currentModelString)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	_, _ = logger.Debugf("Update - Encryption for Request() currentModel:%+v", currentModel)

	// Validate required fields in the request
	if modelValidation := validateModel(CreateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(errorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	// delete existing configuration
	projectID := *currentModel.ProjectId
	if !isExist(currentModel) {
		_, _ = logger.Debugf("Delete encryption - errors encryption at rest with id(%s)", projectID)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	// API call to delete
	_, err = client.EncryptionsAtRest.Delete(context.Background(), projectID)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting encryption at rest configuration for project (%s): %s", projectID, err)
	}
	// Create Atlas API Request Object
	_, _ = logger.Debug("Create - Encryption at rest starts ")

	encryptionAtRest := &mongodbatlas.EncryptionAtRest{
		AwsKms: mongodbatlas.AwsKms{
			Enabled:             currentModel.AwsKms.Enabled,
			CustomerMasterKeyID: *currentModel.AwsKms.CustomerMasterKeyID,
			RoleID:              *currentModel.AwsKms.RoleID,
			Region:              *currentModel.AwsKms.Region,
		},
		GroupID: projectID,
	}
	deploySecretString, _ := json.Marshal(encryptionAtRest)
	log.Printf("Response Object: %s", deploySecretString)

	// API call to create
	_, _, err = client.EncryptionsAtRest.Create(context.Background(), encryptionAtRest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf(errorCreateEncryption, projectID, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	_, _ = logger.Debugf("Delete encryption for Request() currentModel:%+v", currentModel)

	// Validate required fields in the request
	if modelValidation := validateModel(DeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(errorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	projectID := *currentModel.ProjectId

	// Check if  already exist
	if !isExist(currentModel) {
		_, _ = logger.Debugf("Delete encryption - errors encryption at rest with id(%s)", projectID)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	// API call to delete
	_, err = client.EncryptionsAtRest.Delete(context.Background(), projectID)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting encryption at rest configuration for project (%s): %s", projectID, err)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}
func isExist(currentModel *Model) bool {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return false
	}
	setup()
	projectID := *currentModel.ProjectId

	encryptionAtRest, _, err := client.EncryptionsAtRest.Get(context.Background(), projectID)
	if err != nil {
		return false
	}
	if encryptionAtRest != nil && *encryptionAtRest.AwsKms.Enabled {
		return true
	}

	return false
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// no-op
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   currentModel,
	}, nil
}

// function to set logger
func setup() {
	util.SetupLogger("mongodb-atlas-project")
}

func randInt64() int64 {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		return 0
	}
	return val.Int64()
}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}
