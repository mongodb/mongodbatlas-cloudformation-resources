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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"
)

const (
	BucketName = "BucketName"
	IamRoleID  = "IamRoleID"
)

var CreateRequiredFields = []string{constants.ProjectID, BucketName, IamRoleID}
var ReadRequiredFields = []string{constants.ProjectID, constants.ID}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ID}
var ListRequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-cloud-backup-snapshot-export-bucket")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(CreateRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	params := &admin.DiskBackupSnapshotExportBucketRequest{
		BucketName:               currentModel.BucketName,
		CloudProvider:            constants.AWS,
		IamRoleId:                currentModel.IamRoleID,
		RequirePrivateNetworking: currentModel.RequirePrivateNetworking,
	}
	output, resp, err := client.AtlasSDK.CloudBackupsApi.CreateExportBucket(context.Background(), *currentModel.ProjectId, params).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.Id = &output.Id

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create successful",
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(ReadRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	output, resp, err := client.AtlasSDK.CloudBackupsApi.GetExportBucket(context.Background(), *currentModel.ProjectId, *currentModel.Id).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.updateModel(output)

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
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(DeleteRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	resp, err := client.AtlasSDK.CloudBackupsApi.DeleteExportBucket(context.Background(), *currentModel.ProjectId, *currentModel.Id).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete successful",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(ListRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	output, resp, err := client.AtlasSDK.CloudBackupsApi.ListExportBuckets(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	resultList := make([]any, 0)

	if output.Results != nil {
		for i := range *output.Results {
			model := Model{
				ProjectId: currentModel.ProjectId,
				Profile:   currentModel.Profile,
			}
			model.updateModel(&(*output.Results)[i])
			resultList = append(resultList, model)
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List successful",
		ResourceModels:  resultList,
	}, nil
}

func (m *Model) updateModel(bucket *admin.DiskBackupSnapshotExportBucketResponse) {
	m.Id = &bucket.Id
	m.BucketName = &bucket.BucketName
	m.IamRoleID = bucket.IamRoleId
	m.RequirePrivateNetworking = bucket.RequirePrivateNetworking
}
