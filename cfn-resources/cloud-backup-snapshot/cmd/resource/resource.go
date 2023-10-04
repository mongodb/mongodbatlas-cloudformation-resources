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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

var CreateRequiredFields = []string{constants.ClusterName, constants.ProjectID}
var DeleteRequiredFields = []string{constants.ClusterName, constants.ProjectID, constants.SnapshotID}
var ReadRequiredFields = []string{constants.ProjectID, constants.SnapshotID}
var ListRequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-cloud-backup-snapshot")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(CreateRequiredFields, currentModel); err != nil {
		return *err, nil
	}
	if err := clusterOrInstance(currentModel); err != nil {
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

	if util.IsStringPresent(currentModel.ClusterName) {
		params := admin.DiskBackupOnDemandSnapshotRequest{
			Description:     currentModel.Description,
			RetentionInDays: currentModel.RetentionInDays,
		}
		snapshot, resp, err := client.AtlasV2.CloudBackupsApi.TakeSnapshot(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, &params).Execute()
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
	return handler.ProgressEvent{}, errors.New("not implemented: Create for serverless snapshots")
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(ReadRequiredFields, currentModel); err != nil {
		return *err, nil
	}
	if err := clusterOrInstance(currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	if pe := validateExist(client, currentModel); pe != nil {
		return *pe, nil
	}

	if util.IsStringPresent(currentModel.ClusterName) {
		server, resp, err := client.AtlasV2.CloudBackupsApi.GetReplicaSetBackup(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.SnapshotId).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}
		currentModel.updateModelServer(server)
	} else {
		serverless, resp, err := client.AtlasV2.CloudBackupsApi.GetServerlessBackup(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, *currentModel.SnapshotId).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}
		currentModel.updateModelServerless(serverless)
	}

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
	if err := validator.ValidateModel(DeleteRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	if pe := validateExist(client, currentModel); pe != nil {
		return *pe, nil
	}

	if util.IsStringPresent(currentModel.ClusterName) {
		_, resp, err := client.AtlasV2.CloudBackupsApi.DeleteReplicaSetBackup(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.SnapshotId).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(ListRequiredFields, currentModel); err != nil {
		return *err, nil
	}
	if err := clusterOrInstance(currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	models := make([]interface{}, 0)

	if util.IsStringPresent(currentModel.ClusterName) {
		server, resp, err := client.AtlasV2.CloudBackupsApi.ListReplicaSetBackups(aws.BackgroundContext(), *currentModel.ProjectId, *currentModel.ClusterName).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}
		for i := range server.Results {
			model := Model{
				ProjectId:   currentModel.ProjectId,
				Profile:     currentModel.Profile,
				ClusterName: currentModel.ClusterName,
			}
			model.updateModelServer(&server.Results[i])
			models = append(models, &model)
		}
	} else {
		serverless, resp, err := client.AtlasV2.CloudBackupsApi.ListServerlessBackups(aws.BackgroundContext(), *currentModel.ProjectId, *currentModel.InstanceName).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}
		for i := range serverless.Results {
			model := Model{
				ProjectId:    currentModel.ProjectId,
				Profile:      currentModel.Profile,
				InstanceName: currentModel.InstanceName,
			}
			model.updateModelServerless(&serverless.Results[i])
			models = append(models, &model)
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil

}

func validateExist(client *util.MongoDBClient, model *Model) *handler.ProgressEvent {
	if util.IsStringPresent(model.ClusterName) {
		server, resp, err := client.AtlasV2.CloudBackupsApi.ListReplicaSetBackups(aws.BackgroundContext(), *model.ProjectId, *model.ClusterName).Execute()
		if err != nil {
			pe := progressevent.GetFailedEventByResponse(err.Error(), resp)
			return &pe
		}
		for i := range server.Results {
			if util.AreStringPtrEqual(model.SnapshotId, server.Results[i].Id) {
				return nil
			}
		}
	} else {
		serverless, resp, err := client.AtlasV2.CloudBackupsApi.ListServerlessBackups(aws.BackgroundContext(), *model.ProjectId, *model.InstanceName).Execute()
		if err != nil {
			pe := progressevent.GetFailedEventByResponse(err.Error(), resp)
			return &pe
		}
		for i := range serverless.Results {
			if util.AreStringPtrEqual(model.SnapshotId, serverless.Results[i].Id) {
				return nil
			}
		}
	}

	return &handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          "Resource Not Found",
		HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}
}

func validateProgress(client *util.MongoDBClient, currentModel *Model, targetState string) (handler.ProgressEvent, error) {

	snapshotID := *currentModel.SnapshotId
	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName
	snapshot, _, err := client.AtlasV2.CloudBackupsApi.GetReplicaSetBackup(context.Background(), projectID, clusterName, snapshotID).Execute()
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

func (m *Model) updateModelServer(snapShot *admin.DiskBackupReplicaSet) {
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

func (m *Model) updateModelServerless(snapShot *admin.ServerlessBackupSnapshot) {
	m.SnapshotId = snapShot.Id
	m.Status = snapShot.Status
	m.CreatedAt = util.TimePtrToStringPtr(snapShot.CreatedAt)
	m.ExpiresAt = util.TimePtrToStringPtr(snapShot.ExpiresAt)
	m.MongodVersion = snapShot.MongodVersion
	m.StorageSizeBytes = util.IntPtrToStrPtr(util.Int64PtrToIntPtr(snapShot.StorageSizeBytes))
	m.CloudProvider = aws.String(constants.AWS)
}

func clusterOrInstance(model *Model) *handler.ProgressEvent {
	is1, is2 := util.IsStringPresent(model.ClusterName), util.IsStringPresent(model.InstanceName)
	if is1 && is2 || (!is1 && !is2) {
		return &handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Cluster name or instance name must be set, and not both",
			ResourceModel:   model,
		}
	}
	return nil
}
