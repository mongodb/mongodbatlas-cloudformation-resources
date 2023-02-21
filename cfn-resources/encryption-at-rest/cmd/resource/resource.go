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
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateAndUpdateRequiredFields = []string{constants.RoleID, constants.CustomMasterKey, constants.RoleID, constants.ProjectID}
var ReadAndDeleteRequiredFields = []string{constants.ProjectID}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	_, _ = logger.Debugf("Create - Encryption for Request() currentModel:%+v", currentModel)

	// Validate required fields in the request
	if modelValidation := validateModel(CreateAndUpdateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	client, handlerError := util.NewMongoDBClient(req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
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

	if _, _, err := client.EncryptionsAtRest.Create(context.Background(), encryptionAtRest); err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error - Create Encryption  for Project(%s)- Details: %+v", projectID, err)
	}

	currentModel.Id = aws.String(strconv.FormatInt(randInt64(), 10))

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// validateRequest function to validate the request
func validateRequest(req *handler.Request, requiredFields []string, currentModel *Model) (handler.ProgressEvent, *mongodbatlas.Client, error) {
	// Validate required fields in the request
	if modelValidation := validateModel(requiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil, errors.New("required field not found")
	}

	client, handlerError := util.NewMongoDBClient(*req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, nil, errors.New(handlerError.Message)
	}

	// Check if  already exist
	if !isExist(client, currentModel) {
		_, _ = logger.Warnf("resource not found %s", *currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil, errors.New(constants.ResourceNotFound)
	}
	return handler.ProgressEvent{}, client, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	_, _ = logger.Debugf("Read snapshot for Request() currentModel:%+v", currentModel)

	// validate the request
	event, client, err := validateRequest(&req, ReadAndDeleteRequiredFields, currentModel)
	if err != nil {
		if err.Error() == constants.ResourceNotFound {
			return event, nil
		}
		return event, err
	}

	// API call
	projectID := *currentModel.ProjectId
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

	// validate the request
	event, client, err := validateRequest(&req, CreateAndUpdateRequiredFields, currentModel)
	if err != nil {
		if err.Error() == constants.ResourceNotFound {
			return event, nil
		}
		return event, err
	}
	// API call
	projectID := *currentModel.ProjectId
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
		return handler.ProgressEvent{}, fmt.Errorf("error - Create Encryption  for Project(%s)- Details: %+v", projectID, err)
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

	// validate the request
	event, client, err := validateRequest(&req, ReadAndDeleteRequiredFields, currentModel)
	if err != nil {
		if err.Error() == constants.ResourceNotFound {
			return event, nil
		}
		return event, err
	}

	// API call
	projectID := *currentModel.ProjectId
	_, err = client.EncryptionsAtRest.Delete(context.Background(), projectID)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting encryption at rest configuration for project (%s): %s", projectID, err)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}
func isExist(client *mongodbatlas.Client, currentModel *Model) bool {
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
