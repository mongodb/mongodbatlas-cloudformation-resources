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
	"errors"
	"math"
	"math/big"
	"strconv"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115002/admin"
)

var (
	CustomerMasterKeyID           = "AwsKmsConfig.CustomerMasterKeyID"
	RoleID                        = "AwsKmsConfig.RoleID"
	CreateAndUpdateRequiredFields = []string{RoleID, CustomerMasterKeyID, constants.ProjectID}
	ReadAndDeleteRequiredFields   = []string{constants.ProjectID}
)

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

	_, resp, err := client.Atlas20231115002.EncryptionAtRestUsingCustomerKeyManagementApi.UpdateEncryptionAtRest(context.Background(), *currentModel.ProjectId, currentModel.getParams()).Execute()
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

	info, resp, err := client.Atlas20231115002.EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRest(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if pe := validateExist(info); pe != nil {
		return *pe, nil
	}

	currentModel.AwsKmsConfig.CustomerMasterKeyID = info.AwsKms.CustomerMasterKeyID
	currentModel.AwsKmsConfig.Enabled = info.AwsKms.Enabled
	currentModel.AwsKmsConfig.RoleID = info.AwsKms.RoleId
	currentModel.AwsKmsConfig.Region = info.AwsKms.Region

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(CreateAndUpdateRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	info, resp, err := client.Atlas20231115002.EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRest(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if pe := validateExist(info); pe != nil {
		return *pe, nil
	}

	_, resp, err = client.Atlas20231115002.EncryptionAtRestUsingCustomerKeyManagementApi.UpdateEncryptionAtRest(context.Background(), *currentModel.ProjectId, currentModel.getParams()).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
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

	info, resp, err := client.Atlas20231115002.EncryptionAtRestUsingCustomerKeyManagementApi.GetEncryptionAtRest(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if pe := validateExist(info); pe != nil {
		return *pe, nil
	}

	params := &admin.EncryptionAtRest{
		AwsKms: &admin.AWSKMSConfiguration{Enabled: aws.Bool(false)},
	}
	_, resp, err = client.Atlas20231115002.EncryptionAtRestUsingCustomerKeyManagementApi.UpdateEncryptionAtRest(context.Background(), *currentModel.ProjectId, params).Execute()
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

func (m *Model) getParams() *admin.EncryptionAtRest {
	return &admin.EncryptionAtRest{
		AwsKms: &admin.AWSKMSConfiguration{
			Enabled:             m.AwsKmsConfig.Enabled,
			CustomerMasterKeyID: m.AwsKmsConfig.CustomerMasterKeyID,
			RoleId:              m.AwsKmsConfig.RoleID,
			Region:              m.AwsKmsConfig.Region,
		},
	}
}
