package resource

import (
	"context"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	matlasClient "go.mongodb.org/atlas/mongodbatlas"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["snapshot_id"].(string)
		currentModel.Id = &sid
		return validateProgress(client, currentModel, "completed")
	}

	requestParameters := &matlasClient.SnapshotReqPathParameters{
		GroupID:     *currentModel.ProjectId,
		ClusterName: *currentModel.ClusterName,
	}
	snapshotRequest := &matlasClient.CloudProviderSnapshot{
		RetentionInDays: int(*currentModel.RetentionInDays),
		Description:     *currentModel.Description,
	}

	snapshot, _, err := client.CloudProviderSnapshots.Create(context.Background(), requestParameters, snapshotRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating cloud provider snapshot: %s", err)
	}

	currentModel.Id = &snapshot.ID

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

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectId := *currentModel.ProjectId
	snapshotId := *currentModel.Id
	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		SnapshotID:  snapshotId,
		ClusterName: *currentModel.ClusterName,
	}

	snapshot, _, err := client.CloudProviderSnapshots.GetOneCloudProviderSnapshot(context.Background(), snapshotRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading cloud provider snapshot with id(project: %s, snapshot: %s): %s", projectId, snapshotId, err)
	}

	currentModel.Id = &snapshot.ID
	currentModel.Description = &snapshot.Description
	currentModel.RetentionInDays = &snapshot.RetentionInDays
	currentModel.Status = &snapshot.Status
	currentModel.Type = &snapshot.Type
	currentModel.CreatedAt = &snapshot.CreatedAt
	currentModel.MasterKeyUuid = &snapshot.MasterKeyUUID
	currentModel.MongoVersion = &snapshot.MongodVersion
	currentModel.StorageSizeBytes = &snapshot.StorageSizeBytes

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// NO-OP
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectId := *currentModel.ProjectId
	snapshotId := *currentModel.Id
	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		SnapshotID:  snapshotId,
		ClusterName: *currentModel.ClusterName,
	}

	_, err = client.CloudProviderSnapshots.Delete(context.Background(), snapshotRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting cloud provider snapshot with id(project: %s, snapshot: %s): %s", projectId, snapshotId, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectId := *currentModel.ProjectId
	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		ClusterName: *currentModel.ClusterName,
	}

	params := &matlasClient.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	snapshots, _, err := client.CloudProviderSnapshots.GetAllCloudProviderSnapshots(context.Background(), snapshotRequest, params)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading cloud provider snapshot list with id(project: %s): %s", projectId, err)
	}

	var models []Model
	for _, snapshot := range snapshots.Results {
		var model Model
		model.Description = &snapshot.Description
		model.RetentionInDays = &snapshot.RetentionInDays
		model.Status = &snapshot.Status
		model.Type = &snapshot.Type
		model.CreatedAt = &snapshot.CreatedAt
		model.MasterKeyUuid = &snapshot.MasterKeyUUID
		model.MongoVersion = &snapshot.MongodVersion
		model.StorageSizeBytes = &snapshot.StorageSizeBytes
		models = append(models, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   models,
	}, nil
}

func validateProgress(client *matlasClient.Client, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	isReady, state, err := snapshotIsReady(client, *currentModel.ProjectId, *currentModel.Id, *currentModel.ClusterName, targetState)
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
			"snapshot_id": *currentModel.Id,
		}
		return p, nil
	}

	p := handler.NewProgressEvent()
	p.ResourceModel = currentModel
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	return p, nil
}

func snapshotIsReady(client *matlasClient.Client, projectId, snapshotId, clusterName, targetState string) (bool, string, error) {
	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		SnapshotID:  snapshotId,
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
