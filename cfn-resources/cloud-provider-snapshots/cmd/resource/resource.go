package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	log "github.com/sirupsen/logrus"
	matlasClient "go.mongodb.org/atlas/mongodbatlas"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	//logger setup
	setup()

	var projectId string
	var clusterName string

	// Validate required fields in the request
	if currentModel.ProjectId == nil {
		log.Info("Read - Project Id can not be null ")
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Project Id is Nil in the Request",
			HandlerErrorCode: cloudformation.HandlerErrorCodeThrottling}, nil

	}

	if currentModel.ClusterName == nil {
		log.Info("Read - ClusterName can not be null ")
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "ClusterName Id is Nil in the Request",
			HandlerErrorCode: cloudformation.HandlerErrorCodeThrottling}, nil
	}

	projectId = *currentModel.ProjectId
	clusterName = *currentModel.ClusterName
	// progress callback setup
	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["snapshot_id"].(string)
		currentModel.Id = &sid

		return validateProgress(client, currentModel, "completed")
	}

	requestParameters := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		ClusterName: clusterName,
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

	log.Info("Created Successfully - (%s)", currentModel.Id)

	event := handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create cloud provider snapshots : %s", snapshot.Status),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 65,
		CallbackContext: map[string]interface{}{
			"status":      snapshot.Status,
			"snapshot_id": snapshot.ID,
		},
	}
	log.Infof("Create() return event:%+v", event)
	log.Infof("Create() return eventResourceModel:%+v", event.ResourceModel)
	return event, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModelRead *Model) (handler.ProgressEvent, error) {
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModelRead.ApiKeys.PublicKey, *currentModelRead.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	//logger setup
	setup()

	var snapshotId string
	var projectId string
	var clusterName string

	// Validate required fields in the request
	if currentModelRead.ProjectId == nil {
		log.Infof("Read - Project Id can not be null for snapshot with id(%s)", snapshotId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Project Id is Nil in the Request",
			HandlerErrorCode: cloudformation.HandlerErrorCodeThrottling}, nil

	}

	if currentModelRead.Id == nil {
		log.Infof("Read - SnapshotId can not be null ")
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "SnapshotId Id is Nil in the Request",
			HandlerErrorCode: cloudformation.HandlerErrorCodeThrottling}, nil
	}
	if currentModelRead.ClusterName == nil {
		log.Infof("Read - ClusterName can not be null ")
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "ClusterName Id is Nil in the Request",
			HandlerErrorCode: cloudformation.HandlerErrorCodeThrottling}, nil
	}
	projectId = *currentModelRead.ProjectId
	snapshotId = *currentModelRead.Id
	clusterName = *currentModelRead.ClusterName

	isExist := isSnapshotExist(currentModelRead)
	// Check if snapshot already exist
	if !isExist {
		log.Infof("Read - errors reading snapshot with id(%s)", snapshotId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		SnapshotID:  snapshotId,
		ClusterName: clusterName,
	}
	log.Infof("Read - error reading project with : %s", snapshotId)
	snapshot, res, err := client.CloudProviderSnapshots.GetOneCloudProviderSnapshot(context.TODO(), snapshotRequest)
	if err != nil {
		log.Infof("Read - errors reading snapshot with id(%s): %s", snapshotId, err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	log.Infof("Read -reading snapshot status (%d)", res.StatusCode)

	currentModelRead.Id = &snapshot.ID
	currentModelRead.Description = &snapshot.Description
	currentModelRead.Status = &snapshot.Status
	currentModelRead.Type = &snapshot.Type
	currentModelRead.CreatedAt = &snapshot.CreatedAt
	currentModelRead.MasterKeyUuid = &snapshot.MasterKeyUUID
	currentModelRead.MongoVersion = &snapshot.MongodVersion
	currentModelRead.StorageSizeBytes = &snapshot.StorageSizeBytes

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModelRead,
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
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	//logger setup
	setup()

	var snapshotId string
	var projectId string
	var clusterName string

	// Validate required fields in the request
	if currentModel.ProjectId == nil {
		log.Infof("Delete - Project Id can not be null for snapshot with id(%s)", snapshotId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Project Id is Nil in the Request",
			HandlerErrorCode: cloudformation.HandlerErrorCodeThrottling}, nil

	}

	if currentModel.Id == nil {
		log.Infof("Delete -  SnapshotId can not be null ")
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "SnapshotId Id is Nil in the Request",
			HandlerErrorCode: cloudformation.HandlerErrorCodeThrottling}, nil
	}
	if currentModel.ClusterName == nil {
		log.Infof("Delete - ClusterName can not be null ")
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "ClusterName Id is Nil in the Request",
			HandlerErrorCode: cloudformation.HandlerErrorCodeThrottling}, nil
	}
	projectId = *currentModel.ProjectId
	snapshotId = *currentModel.Id
	clusterName = *currentModel.ClusterName

	isExist := isSnapshotExist(currentModel)
	// Check if snapshot already exist
	if !isExist {
		log.Infof("Read - errors reading snapshot with id(%s)", snapshotId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	log.Infof("Deleting snapshotId (%s)", snapshotId)

	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		SnapshotID:  snapshotId,
		ClusterName: clusterName,
	}

	resp, err := client.CloudProviderSnapshots.Delete(context.Background(), snapshotRequest)

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Infof("Delete 404 err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		} else {
			log.Infof("Delete err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
	}, nil

}
func isSnapshotExist(currentModel *Model) bool {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return false
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
		return false
	}
	for _, snapshot := range snapshots.Results {
		log.Infof("Read - errors reading snapshot with id(%s): %s", snapshot.ID, *currentModel.Id)
		if snapshot.ID == *currentModel.Id {
			return true
		}

	}

	return false
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	var projectId string
	var clusterName string

	// Validate required fields in the request
	if currentModel.ProjectId == nil {
		log.Infof("Read - Project Id can not be null ")
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Project Id is Nil in the Request",
			HandlerErrorCode: cloudformation.HandlerErrorCodeThrottling}, nil

	}

	if currentModel.ClusterName == nil {
		log.Infof("Read - ClusterName can not be null ")
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "ClusterName Id is Nil in the Request",
			HandlerErrorCode: cloudformation.HandlerErrorCodeThrottling}, nil
	}

	projectId = *currentModel.ProjectId
	clusterName = *currentModel.ClusterName

	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		ClusterName: clusterName,
	}

	params := &matlasClient.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	snapshots, _, err := client.CloudProviderSnapshots.GetAllCloudProviderSnapshots(context.Background(), snapshotRequest, params)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading cloud provider snapshot list with id(project: %s): %s", projectId, err)
	}
	var models []interface{}
	for _, snapshot := range snapshots.Results {
		var model Model
		model.Description = &snapshot.Description
		model.Id = &snapshot.ID
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
		ResourceModels:  models,
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
func setup() {
	util.SetupLogger("mongodb-atlas-project")

}
