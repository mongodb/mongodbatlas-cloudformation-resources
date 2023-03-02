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
	"errors"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	localconstants "github.com/mongodb/mongodbatlas-cloudformation-resources/cloud-backup-snapshot-export-bucket/cmd/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.ProjectID, localconstants.BucketName, localconstants.IamRoleID}
var ReadRequiredFields = []string{constants.ProjectID, constants.ID}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ID}
var ListRequiredFields = []string{constants.ProjectID}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-CloudBackupSnapshotExportBucket")
}

func (m *Model) completeByAtlasModel(bucket *mongodbatlas.CloudProviderSnapshotExportBucket) {
	m.BucketName = &bucket.BucketName
	m.IamRoleID = &bucket.IAMRoleID
	m.Id = &bucket.ID
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = log.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}
	var res *mongodbatlas.Response

	input := &mongodbatlas.CloudProviderSnapshotExportBucket{
		BucketName:    *currentModel.BucketName,
		CloudProvider: constants.AWS,
		IAMRoleID:     *currentModel.IamRoleID,
	}

	output, res, err := client.CloudProviderSnapshotExportBuckets.Create(context.Background(), *currentModel.ProjectId, input)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.Id = &output.ID

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create successful",
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = log.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}
	var res *mongodbatlas.Response

	output, res, err := client.CloudProviderSnapshotExportBuckets.Get(context.Background(), *currentModel.ProjectId, *currentModel.Id)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.completeByAtlasModel(output)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Get successful",
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Warnf("Create cluster model : %+v", currentModel)

	// Validation
	if modelValidation := validateModel(DeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = log.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}
	var err error
	var res *mongodbatlas.Response

	res, err = client.CloudProviderSnapshotExportBuckets.Delete(context.Background(), *currentModel.ProjectId, *currentModel.Id)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete successful",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = log.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	output, res, err := client.CloudProviderSnapshotExportBuckets.List(context.Background(), *currentModel.ProjectId, nil)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	resultList := make([]interface{}, 0)

	for i := range output.Results {
		model := Model{}
		model.completeByAtlasModel(output.Results[i])
		model.ProjectId = currentModel.ProjectId
		model.Profile = currentModel.Profile
		resultList = append(resultList, model)
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List successful",
		ResourceModels:  resultList,
	}, nil
}
