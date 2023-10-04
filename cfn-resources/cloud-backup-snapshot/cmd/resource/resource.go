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
	"strconv"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
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

	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["snapshot_id"].(string)
		currentModel.SnapshotId = &sid
		return validateProgress(client, currentModel, "completed")
	}

	clusterName := cast.ToString(currentModel.ClusterName)
	projectID := cast.ToString(currentModel.ProjectId)

	requestParameters := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}
	snapshotRequest := &mongodbatlas.CloudProviderSnapshot{
		RetentionInDays: cast.ToInt(currentModel.RetentionInDays),
		Description:     cast.ToString(currentModel.Description),
	}

	snapshot, res, err := client.CloudProviderSnapshots.Create(context.Background(), requestParameters, snapshotRequest)
	if err != nil {
		if res.Response.StatusCode == 400 && strings.Contains(err.Error(), constants.UnfinishedOnDemandSnapshot) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
		}
		return progressevent.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	currentModel.SnapshotId = &snapshot.ID

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create cloud provider snapshots : %s", snapshot.Status),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 65,
		CallbackContext: map[string]interface{}{
			"status":      snapshot.Status,
			"snapshot_id": snapshot.ID,
		},
	}, nil
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

	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	clusterName := cast.ToString(currentModel.ClusterName)
	instanceName := cast.ToString(currentModel.InstanceName)
	projectID := cast.ToString(currentModel.ProjectId)

	if !isSnapshotExist(client, currentModel) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	snapshotID := cast.ToString(currentModel.SnapshotId)
	var snapshot *mongodbatlas.CloudProviderSnapshot
	var res *mongodbatlas.Response
	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:    projectID,
		SnapshotID: snapshotID,
	}
	var err error

	if clusterName != "" {
		snapshotRequest.ClusterName = clusterName
		snapshot, res, err = client.CloudProviderSnapshots.GetOneCloudProviderSnapshot(context.Background(), snapshotRequest)
	} else {
		// read serverless instance snapshot
		snapshotRequest.InstanceName = instanceName
		snapshot, res, err = client.CloudProviderSnapshots.GetOneServerlessSnapshot(context.Background(), snapshotRequest)
	}

	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	currentModel = convertToUIModel(snapshot, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
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

	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	if !isSnapshotExist(client, currentModel) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	clusterName := cast.ToString(currentModel.ClusterName)
	projectID := cast.ToString(currentModel.ProjectId)
	snapshotID := cast.ToString(currentModel.SnapshotId)

	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		SnapshotID:  snapshotID,
		ClusterName: clusterName,
	}

	resp, err := client.CloudProviderSnapshots.Delete(context.Background(), snapshotRequest)

	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp.Response), nil
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

	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	// Create Atlas API Request Object
	projectID := *currentModel.ProjectId
	clusterName := cast.ToString(currentModel.ClusterName)
	instanceName := cast.ToString(currentModel.InstanceName)
	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID: projectID,
	}
	models := make([]interface{}, 0)
	var snapshots *mongodbatlas.CloudProviderSnapshots
	var res *mongodbatlas.Response
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	var err error
	// API call to list snapshot
	if clusterName != "" {
		snapshotRequest.ClusterName = clusterName
		snapshots, res, err = client.CloudProviderSnapshots.GetAllCloudProviderSnapshots(context.Background(), snapshotRequest, params)
	} else {
		// read serverless instance snapshot
		snapshotRequest.InstanceName = instanceName
		snapshots, res, err = client.CloudProviderSnapshots.GetAllServerlessSnapshots(context.Background(), snapshotRequest, params)
	}

	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// populate list
	snapshotsList := snapshots.Results
	for ind := range snapshotsList {
		var model Model
		model.ProjectId = currentModel.ProjectId
		model.Profile = currentModel.Profile
		model.ClusterName = currentModel.ClusterName
		models = append(models, convertToUIModel(snapshotsList[ind], &model))
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}

// function to check if snapshot already available for the snapshot id
func isSnapshotExist(client *mongodbatlas.Client, currentModel *Model) bool {
	// Create Atlas API Request Object
	clusterName := cast.ToString(currentModel.ClusterName)
	instanceName := cast.ToString(currentModel.InstanceName)
	projectID := cast.ToString(currentModel.ProjectId)
	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID: projectID,
	}
	var snapshots *mongodbatlas.CloudProviderSnapshots
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	var err error
	// API call to list snapshot
	if clusterName != "" {
		snapshotRequest.ClusterName = clusterName
		snapshots, _, err = client.CloudProviderSnapshots.GetAllCloudProviderSnapshots(context.Background(), snapshotRequest, params)
	} else if instanceName != "" {
		// read serverless instance snapshot
		snapshotRequest.InstanceName = instanceName
		snapshots, _, err = client.CloudProviderSnapshots.GetAllServerlessSnapshots(context.Background(), snapshotRequest, params)
	}
	if err != nil {
		return false
	}
	for _, snapshot := range snapshots.Results {
		if snapshot.ID == *currentModel.SnapshotId {
			return true
		}
	}
	return false
}

// function to track snapshot creation status
func validateProgress(client *mongodbatlas.Client, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	snapshotID := *currentModel.SnapshotId
	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName
	isReady, state, err := snapshotIsReady(client, projectID, snapshotID, clusterName, targetState)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	if !isReady {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = handler.InProgress
		p.CallbackDelaySeconds = 35
		p.Message = "Pending"
		p.CallbackContext = map[string]interface{}{
			"status":      state,
			"snapshot_id": *currentModel.SnapshotId,
		}
		return p, nil
	}
	// Create Atlas API Request Object
	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		SnapshotID:  snapshotID,
		ClusterName: clusterName,
	}

	snapshot, _, _ := client.CloudProviderSnapshots.GetOneCloudProviderSnapshot(context.Background(), snapshotRequest)
	currentModel = convertToUIModel(snapshot, currentModel)
	p := handler.NewProgressEvent()
	p.ResourceModel = currentModel
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	return p, nil
}

// function to check if snapshot already exist in atlas
func snapshotIsReady(client *mongodbatlas.Client, projectID, snapshotID, clusterName, targetState string) (isExist bool, status string, err error) {
	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		SnapshotID:  snapshotID,
		ClusterName: clusterName,
	}

	snapshot, resp, err := client.CloudProviderSnapshots.GetOneCloudProviderSnapshot(context.Background(), snapshotRequest)
	if err != nil {
		if snapshot == nil && resp == nil {
			return false, "", err
		}
		if resp != nil && resp.StatusCode == 404 {
			return true, "deleted", nil
		}
		return false, "", err
	}
	return snapshot.Status == targetState, snapshot.Status, nil
}

func flattenCloudMember(apiObject *mongodbatlas.Member) ApiAtlasDiskBackupShardedClusterSnapshotMemberView {
	tfMap := ApiAtlasDiskBackupShardedClusterSnapshotMemberView{}
	tfMap.CloudProvider = &apiObject.CloudProvider
	tfMap.Id = &apiObject.ID
	tfMap.ReplicaSetName = &apiObject.ReplicaSetName

	return tfMap
}

func flattenCloudMembers(apiObjects []*mongodbatlas.Member) []ApiAtlasDiskBackupShardedClusterSnapshotMemberView {
	tfList := make([]ApiAtlasDiskBackupShardedClusterSnapshotMemberView, 0)

	for ind := range apiObjects {
		if apiObjects[ind] != nil {
			tfList = append(tfList, flattenCloudMember(apiObjects[ind]))
		}
	}
	return tfList
}

func convertToUIModel(snapShot *mongodbatlas.CloudProviderSnapshot, currentModel *Model) *Model {
	currentModel.SnapshotId = &snapShot.ID
	currentModel.Description = &snapShot.Description
	currentModel.Status = &snapShot.Status
	currentModel.Type = &snapShot.Type
	currentModel.CreatedAt = &snapShot.CreatedAt
	currentModel.ExpiresAt = &snapShot.ExpiresAt
	currentModel.ReplicaSetName = &snapShot.ReplicaSetName
	currentModel.MasterKeyUUID = &snapShot.MasterKeyUUID
	currentModel.MongodVersion = &snapShot.MongodVersion
	storageInBytesString := strconv.Itoa(snapShot.StorageSizeBytes)
	currentModel.StorageSizeBytes = &storageInBytesString
	currentModel.CloudProvider = aws.String(constants.AWS)
	currentModel.SnapshotIds = snapShot.SnapshotsIds
	currentModel.Members = flattenCloudMembers(snapShot.Members)
	return currentModel
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
