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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231001001/admin"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateAndUpdateRequiredFields = []string{constants.RoleID, constants.CustomMasterKey, constants.RoleID, constants.ProjectID}
var ReadAndDeleteRequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-encryption-at-rest")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(CreateAndUpdateRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	params := &admin.EncryptionAtRest{
		AwsKms: &admin.AWSKMSConfiguration{
			Enabled:             currentModel.AwsKms.Enabled,
			CustomerMasterKeyID: currentModel.AwsKms.CustomerMasterKeyID,
			RoleId:              currentModel.AwsKms.RoleID,
			Region:              currentModel.AwsKms.Region,
		},
	}
	_, resp, err := client.AtlasV2.EncryptionAtRestUsingCustomerKeyManagementApi.UpdateEncryptionAtRest(context.Background(), *currentModel.ProjectId, params).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}
	currentModel.Id = aws.String(strconv.FormatInt(randInt64(), 10))

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(ReadAndDeleteRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	info, resp, err := client.AtlasV2.EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRest(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if pe := validateExist(info); pe != nil {
		return *pe, nil
	}

	currentModel.AwsKms.CustomerMasterKeyID = info.AwsKms.CustomerMasterKeyID
	currentModel.AwsKms.Enabled = info.AwsKms.Enabled
	currentModel.AwsKms.RoleID = info.AwsKms.RoleId
	currentModel.AwsKms.Region = info.AwsKms.Region

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

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

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(ReadAndDeleteRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	info, resp, err := client.AtlasV2.EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRest(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if pe := validateExist(info); pe != nil {
		return *pe, nil
	}

	params := &admin.EncryptionAtRest{
		AwsKms: &admin.AWSKMSConfiguration{Enabled: aws.Bool(false)},
	}
	_, resp, err = client.AtlasV2.EncryptionAtRestUsingCustomerKeyManagementApi.UpdateEncryptionAtRest(context.Background(), *currentModel.ProjectId, params).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func validateRequest(req *handler.Request, requiredFields []string, currentModel *Model) (handler.ProgressEvent, *mongodbatlas.Client, error) {
	// Validate required fields in the request
	if modelValidation := validator.ValidateModel(requiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil, errors.New("required field not found")
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, handlerError := util.NewMongoDBClient(*req, currentModel.Profile)
	if handlerError != nil {
		return *handlerError, nil, errors.New(handlerError.Message)
	}

	// Check if  already exist
	if !isExist(client, currentModel) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil, errors.New(constants.ResourceNotFound)
	}
	return handler.ProgressEvent{}, client, nil
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

func validateExist(info *admin.EncryptionAtRest) *handler.ProgressEvent {
	if info != nil && info.AwsKms != nil && aws.BoolValue(info.AwsKms.Enabled) {
		return nil
	}
	return &handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          "Resource Not Found",
		HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}
}

func randInt64() int64 {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		return 0
	}
	return val.Int64()
}
