package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/encryption-at-rest/cmd/validator_def"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/atlas/mongodbatlas"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() //logger setup

	log.Debugf("Create - Encryption for Request() currentModel:%+v", currentModel)
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	// Validate required fields in the request
	modelValidation := validateModel(constants.Create, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create Atlas API Request Object
	log.Info("Create - Encryption at rest starts ")
	encryptionAtRest := &mongodbatlas.EncryptionAtRest{
		AwsKms: mongodbatlas.AwsKms{
			Enabled:             currentModel.AwsKms.Enabled,
			CustomerMasterKeyID: *currentModel.AwsKms.CustomerMasterKeyID,
			RoleID:              *currentModel.AwsKms.RoleID,
			Region:              *currentModel.AwsKms.Region,
		},
		GroupID: *currentModel.ProjectId,
	}
	deploySecretString, err := json.Marshal(encryptionAtRest)
	log.Printf("Request Object: %s", deploySecretString)

	// API call to create
	_, _, err = client.EncryptionsAtRest.Create(context.Background(), encryptionAtRest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating encryption at rest: %s", err)
	}
	currentModelString, err := json.Marshal(currentModel)
	log.Printf("Request Object: %s", currentModelString)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() //logger setup

	log.Debugf("Read snapshot for Request() currentModel:%+v", currentModel)
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	// Validate required fields in the request
	modelValidation := validateModel(constants.Read, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create Atlas API Request Object
	projectID := *currentModel.ProjectId
	// API call
	encryptionAtRest, _, err := client.EncryptionsAtRest.Get(context.Background(), projectID)
	if err != nil {
		return handler.NewProgressEvent(), fmt.Errorf("error fetching encryption at rest configuration for project (%s): %s", projectID, err)
	}
	isExist := isExist(currentModel)
	// Check if already  exist
	if !isExist {
		log.Infof("Read - errors encryption at rest with id(%s)", projectID)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	currentModel.AwsKms.CustomerMasterKeyID = &encryptionAtRest.AwsKms.CustomerMasterKeyID
	currentModel.AwsKms.Enabled = encryptionAtRest.AwsKms.Enabled
	currentModel.AwsKms.RoleID = &encryptionAtRest.AwsKms.RoleID
	currentModel.AwsKms.Region = &encryptionAtRest.AwsKms.Region

	currentModelString, err := json.Marshal(currentModel)
	log.Printf("Response Object: %s", currentModelString)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// no-op
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() //logger setup

	log.Debugf("Delete encryption for Request() currentModel:%+v", currentModel)
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	// Validate required fields in the request
	modelValidation := validateModel(constants.Delete, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	projectID := *currentModel.ProjectId
	isExist := isExist(currentModel)
	// Check if  already exist
	if !isExist {
		log.Infof("Read - errors encryption at rest with id(%s)", projectID)
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
func setup() {
	util.SetupLogger("mongodb-atlas-project")

}

// function to validate inputs to all actions
func validateModel(event constants.Event, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(event, validator_def.ModelValidator{}, model)
}
