package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/cloud-backup/cmd/validator_def"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	log "github.com/sirupsen/logrus"
	matlasClient "go.mongodb.org/atlas/mongodbatlas"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() //logger setup

	log.Debugf("Create snapshot for Request() currentModel:%+v", currentModel)
	// Validate required fields in the request
	modelValidation := validateModel(constants.Create, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	// progress callback setup
	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["snapshot_id"].(string)
		currentModel.Id = &sid
		return validateProgress(client, currentModel, "completed")
	}
	// Create Atlas API Request Object
	projectId := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName

	requestParameters := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		ClusterName: clusterName,
	}
	snapshotRequest := &matlasClient.CloudProviderSnapshot{
		RetentionInDays: int(*currentModel.RetentionInDays),
		Description:     *currentModel.Description,
	}
	// API call to create snapshot
	snapshot, _, err := client.CloudProviderSnapshots.Create(context.Background(), requestParameters, snapshotRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating cloud provider snapshot: %s", err)
	}
	currentModel.Id = &snapshot.ID
	log.Info("Created Successfully - (%s)", currentModel.Id)
	// track progress
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
	return event, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() //logger setup

	log.Debugf("Read snapshot for Request() currentModel:%+v", currentModel)
	// Validate required fields in the request
	modelValidation := validateModel(constants.Read, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	isExist := isSnapshotExist(currentModel)
	// Check if snapshot already exist due to this issue https://github.com/mongodb/go-client-mongodb-atlas/issues/315
	if !isExist {
		log.Infof("Read - errors reading snapshot with id(%s)", currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	// Create Atlas API Request Object
	projectId := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName
	snapshotId := *currentModel.Id

	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		SnapshotID:  snapshotId,
		ClusterName: clusterName,
	}
	// API call to read snapshot
	snapshot, res, err := client.CloudProviderSnapshots.GetOneCloudProviderSnapshot(context.Background(), snapshotRequest)
	if err != nil {
		log.Infof("Read - errors reading snapshot with id(%s): %s", snapshotId, err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	log.Infof("Read -reading snapshot status (%d)", res.StatusCode)

	currentModel.Id = &snapshot.ID
	currentModel.Description = &snapshot.Description
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
	setup() //logger setup

	log.Debugf("Delete snapshot for Request() currentModel:%+v", currentModel)
	// Validate required fields in the request
	modelValidation := validateModel(constants.Delete, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	isExist := isSnapshotExist(currentModel)
	// Check if snapshot already exist due to this issue https://github.com/mongodb/go-client-mongodb-atlas/issues/315
	if !isExist {
		log.Infof("Delete - errors Delete snapshot with id(%s)", currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	// Create Atlas API Request Object
	projectId := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName
	snapshotId := *currentModel.Id

	log.Infof("Deleting snapshotId (%s)", snapshotId)

	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		SnapshotID:  snapshotId,
		ClusterName: clusterName,
	}
	// API call to delete snapshot
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
		Message:         "Delete Complete",
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
	setup() //logger setup

	log.Debugf("return all snapshot for Request() currentModel:%+v", currentModel)
	// Validate required fields in the request
	modelValidation := validateModel(constants.List, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	// Create Atlas API Request Object
	projectId := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName
	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		ClusterName: clusterName,
	}
	params := &matlasClient.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}
	// API call to read snapshot
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

	if len(models) == 0 {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "List Complete",
			ResourceModels:  nil,
		}, nil
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}

// function to track snapshot creation status
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

// function to check if snapshot already exist in atlas
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

// logger setup function
func setup() {
	util.SetupLogger("mongodb-atlas-project")

}

// function to validate inputs to all actions
func validateModel(event constants.Event, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(event, validator_def.ModelValidator{}, model)
}
