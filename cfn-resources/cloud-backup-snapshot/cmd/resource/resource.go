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
	"fmt"

	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

const (
	clusterInstanceType = "cluster"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.InstanceType}
var DeleteRequiredFields = []string{constants.ProjectID, constants.SnapshotID, constants.InstanceName, constants.InstanceType}
var ReadRequiredFields = []string{constants.ProjectID, constants.SnapshotID, constants.InstanceName, constants.InstanceType}
var ListRequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-cloud-backup-snapshot")
}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	if pe := validator.ValidateModel(fields, model); pe != nil {
		return pe
	}

	if *model.InstanceType != clusterInstanceType {
		pe := progressevent.GetFailedEventByCode(fmt.Sprintf("InstanceType must be %s", clusterInstanceType),
			string(types.HandlerErrorCodeInvalidRequest))
		return &pe
	}

	return nil
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validateModel(CreateRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["snapshot_id"].(string)
		currentModel.SnapshotId = &sid
		return validateProgress(client, currentModel, "completed")
	}

	params := admin20231115002.DiskBackupOnDemandSnapshotRequest{
		Description:     currentModel.Description,
		RetentionInDays: currentModel.RetentionInDays,
	}
	snapshot, resp, err := client.Atlas20231115002.CloudBackupsApi.TakeSnapshot(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, &params).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.SnapshotId = snapshot.Id

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create cloud provider snapshots : %s", *snapshot.Status),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 65,
		CallbackContext: map[string]interface{}{
			"status":      snapshot.Status,
			"snapshot_id": snapshot.Id,
		},
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validateModel(ReadRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	if pe := validateExist(client, currentModel); pe != nil {
		return *pe, nil
	}

	server, resp, err := client.Atlas20231115002.CloudBackupsApi.GetReplicaSetBackup(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, *currentModel.SnapshotId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}
	currentModel.updateModelServer(server)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete ",
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validateModel(DeleteRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	if pe := validateExist(client, currentModel); pe != nil {
		return *pe, nil
	}

	_, resp, err := client.Atlas20231115002.CloudBackupsApi.DeleteReplicaSetBackup(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, *currentModel.SnapshotId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validateModel(ListRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	models := make([]interface{}, 0)

	server, resp, err := client.Atlas20231115002.CloudBackupsApi.ListReplicaSetBackups(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}
	for i := range server.Results {
		model := Model{
			ProjectId:    currentModel.ProjectId,
			Profile:      currentModel.Profile,
			InstanceName: currentModel.InstanceName,
			InstanceType: currentModel.InstanceType,
		}
		model.updateModelServer(&server.Results[i])
		models = append(models, &model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}

func validateExist(client *util.MongoDBClient, model *Model) *handler.ProgressEvent {
	server, resp, err := client.Atlas20231115002.CloudBackupsApi.ListReplicaSetBackups(context.Background(), *model.ProjectId, *model.InstanceName).Execute()
	if err != nil {
		pe := progressevent.GetFailedEventByResponse(err.Error(), resp)
		return &pe
	}
	for i := range server.Results {
		if util.AreStringPtrEqual(model.SnapshotId, server.Results[i].Id) {
			return nil
		}
	}

	return &handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          "Resource Not Found",
		HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}
}

func validateProgress(client *util.MongoDBClient, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	snapshotID := *currentModel.SnapshotId
	projectID := *currentModel.ProjectId
	clusterName := *currentModel.InstanceName
	snapshot, _, err := client.Atlas20231115002.CloudBackupsApi.GetReplicaSetBackup(context.Background(), projectID, clusterName, snapshotID).Execute()
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	status := util.SafeString(snapshot.Status)

	if status == targetState {
		currentModel.updateModelServer(snapshot)
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = handler.Success
		p.Message = "Complete"
		return p, nil
	}

	p := handler.NewProgressEvent()
	p.ResourceModel = currentModel
	p.OperationStatus = handler.InProgress
	p.CallbackDelaySeconds = 35
	p.Message = "Pending"
	p.CallbackContext = map[string]interface{}{
		"status":      status,
		"snapshot_id": *currentModel.SnapshotId,
	}
	return p, nil
}

func (m *Model) updateModelServer(snapShot *admin20231115002.DiskBackupReplicaSet) {
	m.SnapshotId = snapShot.Id
	m.Description = snapShot.Description
	m.Status = snapShot.Status
	m.Type = snapShot.Type
	m.CreatedAt = util.TimePtrToStringPtr(snapShot.CreatedAt)
	m.ExpiresAt = util.TimePtrToStringPtr(snapShot.ExpiresAt)
	m.ReplicaSetName = snapShot.ReplicaSetName
	m.MasterKeyUUID = snapShot.MasterKeyUUID
	m.MongodVersion = snapShot.MongodVersion
	m.StorageSizeBytes = util.IntPtrToStrPtr(util.Int64PtrToIntPtr(snapShot.StorageSizeBytes))
	m.CloudProvider = snapShot.CloudProvider
}
