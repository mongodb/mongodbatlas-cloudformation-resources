package resource

import (
	"context"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/openlyinc/pointy"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	publicKey              = "ApiKeys.PublicKey"
	privateKey             = "ApiKeys.PrivateKey"
	projectID              = "GroupId"
	clusterName            = "ClusterName"
	id                     = "SnapshotId"
	cloudProvider          = "AWS"
	errorCreateMongoClient = "Error - Create MongoDB Client- Details: %+v"
	errorCreateCloudBackup = "Error - Create Cloud Backup snapshot for Project(%s) and Cluster(%s)- Details: %+v"
	errorReadCloudBackUp   = "Error - Read snapshot with id(%s)"
)

var CreateRequiredFields = []string{publicKey, privateKey, clusterName, projectID}
var ReadRequiredFields = []string{publicKey, privateKey, clusterName, projectID, id}
var DeleteRequiredFields = []string{publicKey, privateKey, clusterName, projectID, id}
var ListRequiredFields = []string{publicKey, privateKey, clusterName, projectID}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	_, _ = logger.Debugf("Create snapshot for Request() currentModel:%+v", currentModel)

	// Validate required fields in the request
	if modelValidation := validateModel(CreateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(errorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// progress callback setup
	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["snapshot_id"].(string)
		currentModel.SnapshotId = &sid
		return validateProgress(client, currentModel, "completed")
	}

	// Create Atlas API Request Object
	projectID := *currentModel.GroupId
	clusterName := *currentModel.ClusterName

	requestParameters := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}

	snapshotRequest := &mongodbatlas.CloudProviderSnapshot{
		RetentionInDays: *currentModel.RetentionInDays,
		Description:     *currentModel.Description,
	}

	// API call to create snapshot
	snapshot, _, err := client.CloudProviderSnapshots.Create(context.Background(), requestParameters, snapshotRequest)
	if err != nil {
		_, _ = logger.Warnf(errorCreateCloudBackup, projectID, clusterName, err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	currentModel.SnapshotId = pointy.String(snapshot.ID)

	_, _ = logger.Debugf("Created Successfully - (%s)", *currentModel.SnapshotId)

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
	_, _ = logger.Debugf("Create() return event:%+v", event)

	return event, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	_, _ = logger.Debugf("Read snapshot for Request() currentModel:%+v", currentModel)

	// Validate required fields in the request
	if modelValidation := validateModel(ReadRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(errorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	if !isSnapshotExist(currentModel) {
		_, _ = logger.Warnf(errorReadCloudBackUp, *currentModel.SnapshotId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	// Create Atlas API Request Object
	projectID := *currentModel.GroupId
	clusterName := *currentModel.ClusterName
	snapshotID := *currentModel.SnapshotId

	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		SnapshotID:  snapshotID,
		ClusterName: clusterName,
	}
	// API call to read snapshot
	snapshot, res, err := client.CloudProviderSnapshots.GetOneCloudProviderSnapshot(context.Background(), snapshotRequest)
	if err != nil {
		_, _ = logger.Warnf("Read - errors reading snapshot with id(%s): %s", snapshotID, err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	_, _ = logger.Debugf("Read -reading snapshot status (%d)", res.StatusCode)

	currentModel.SnapshotId = &snapshot.ID
	currentModel.Description = &snapshot.Description
	currentModel.RetentionInDays = &snapshot.RetentionInDays
	currentModel.Status = &snapshot.Status
	currentModel.Type = &snapshot.Type
	currentModel.CreatedAt = &snapshot.CreatedAt
	currentModel.ExpiresAt = &snapshot.ExpiresAt
	currentModel.ReplicaSetName = &snapshot.ReplicaSetName
	currentModel.MasterKeyUUID = &snapshot.MasterKeyUUID
	currentModel.MongodVersion = &snapshot.MongodVersion
	currentModel.StorageSizeBytes = &snapshot.StorageSizeBytes
	currentModel.Links = flattenLinks(snapshot.Links)
	currentModel.CloudProvider = pointy.String(cloudProvider)
	currentModel.SnapshotIds = snapshot.SnapshotsIds
	currentModel.Members = flattenCloudMembers(snapshot.Members)
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
	setup() // logger setup
	_, _ = logger.Debugf("Delete snapshot for Request() currentModel:%+v", currentModel)

	// Validate required fields in the request
	if modelValidation := validateModel(DeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(errorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	if !isSnapshotExist(currentModel) {
		_, _ = logger.Warnf(errorReadCloudBackUp, *currentModel.SnapshotId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	// Create Atlas API Request Object
	projectID := *currentModel.GroupId
	clusterName := *currentModel.ClusterName
	snapshotID := *currentModel.SnapshotId

	_, _ = logger.Debugf("Deleting snapshotID (%s)", snapshotID)

	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		SnapshotID:  snapshotID,
		ClusterName: clusterName,
	}
	// API call to delete snapshot
	resp, err := client.CloudProviderSnapshots.Delete(context.Background(), snapshotRequest)

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			_, _ = logger.Warnf("Delete 404 err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = logger.Warnf("Delete err: %+v", err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

// function to check if snapshot already available for the snapshot id
func isSnapshotExist(currentModel *Model) bool {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return false
	}

	projectID := *currentModel.GroupId
	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: *currentModel.ClusterName,
	}

	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	snapshots, _, err := client.CloudProviderSnapshots.GetAllCloudProviderSnapshots(context.Background(), snapshotRequest, params)
	if err != nil {
		return false
	}
	for _, snapshot := range snapshots.Results {
		_, _ = logger.Debugf("Read - errors reading snapshot with id(%s): %s", snapshot.ID, *currentModel.SnapshotId)
		if snapshot.ID == *currentModel.SnapshotId {
			return true
		}
	}
	return false
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	_, _ = logger.Debugf("return all snapshot for Request() currentModel:%+v", currentModel)

	// Validate required fields in the request
	if modelValidation := validateModel(ListRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(errorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	// Create Atlas API Request Object
	projectID := *currentModel.GroupId
	clusterName := *currentModel.ClusterName
	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}
	// API call to read snapshot
	snapshots, _, err := client.CloudProviderSnapshots.GetAllCloudProviderSnapshots(context.Background(), snapshotRequest, params)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading cloud provider snapshot list with id(project: %s): %s", projectID, err)
	}
	models := make([]interface{}, 0)
	for _, snapshot := range snapshots.Results {
		var model Model
		model.SnapshotId = &snapshot.ID
		model.Description = &snapshot.Description
		model.Status = &snapshot.Status
		model.Type = &snapshot.Type
		model.CreatedAt = &snapshot.CreatedAt
		model.ExpiresAt = &snapshot.ExpiresAt
		model.ReplicaSetName = &snapshot.ReplicaSetName
		model.MasterKeyUUID = &snapshot.MasterKeyUUID
		model.MongodVersion = &snapshot.MongodVersion
		model.StorageSizeBytes = &snapshot.StorageSizeBytes
		model.Links = flattenLinks(snapshot.Links)
		model.CloudProvider = pointy.String(cloudProvider)
		model.SnapshotIds = snapshot.SnapshotsIds
		model.Members = flattenCloudMembers(snapshot.Members)
		models = append(models, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}

// function to track snapshot creation status
func validateProgress(client *mongodbatlas.Client, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	snapshotID := *currentModel.SnapshotId
	projectID := *currentModel.GroupId
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
	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		SnapshotID:  snapshotID,
		ClusterName: clusterName,
	}

	snapshot, _, _ := client.CloudProviderSnapshots.GetOneCloudProviderSnapshot(context.Background(), snapshotRequest)
	currentModel.SnapshotId = &snapshot.ID
	currentModel.Description = &snapshot.Description
	currentModel.Status = &snapshot.Status
	currentModel.Type = &snapshot.Type
	currentModel.CreatedAt = &snapshot.CreatedAt
	currentModel.ExpiresAt = &snapshot.ExpiresAt
	currentModel.ReplicaSetName = &snapshot.ReplicaSetName
	currentModel.MasterKeyUUID = &snapshot.MasterKeyUUID
	currentModel.MongodVersion = &snapshot.MongodVersion
	currentModel.StorageSizeBytes = &snapshot.StorageSizeBytes
	currentModel.Links = flattenLinks(snapshot.Links)
	currentModel.CloudProvider = pointy.String(cloudProvider)
	currentModel.SnapshotIds = snapshot.SnapshotsIds
	currentModel.Members = flattenCloudMembers(snapshot.Members)
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

// logger setup function
func setup() {
	util.SetupLogger("mongodb-atlas-BackupSnapshot")
}

// converts mongodb link class to model link class
func flattenLinks(linksResult []*mongodbatlas.Link) []Link {
	links := make([]Link, 0)
	for _, link := range linksResult {
		var lin Link
		lin.Href = &link.Href
		lin.Rel = &link.Rel
		links = append(links, lin)
	}
	return links
}
func flattenCloudMember(apiObject *mongodbatlas.Member) ApiAtlasDiskBackupShardedClusterSnapshotMemberView {
	tfMap := ApiAtlasDiskBackupShardedClusterSnapshotMemberView{}
	tfMap.CloudProvider = pointy.String(apiObject.CloudProvider)
	tfMap.Id = pointy.String(apiObject.ID)
	tfMap.ReplicaSetName = pointy.String(apiObject.ReplicaSetName)

	return tfMap
}

func flattenCloudMembers(apiObjects []*mongodbatlas.Member) []ApiAtlasDiskBackupShardedClusterSnapshotMemberView {
	if len(apiObjects) == 0 {
		return nil
	}
	tfList := make([]ApiAtlasDiskBackupShardedClusterSnapshotMemberView, 0)

	for _, apiObject := range apiObjects {
		if apiObject != nil {
			tfList = append(tfList, flattenCloudMember(apiObject))
		}
	}
	return tfList
}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}
