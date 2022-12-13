package resource

import (
	"context"
	"fmt"

	progressEvents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/openlyinc/pointy"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ClusterName, constants.GroupID}
var DeleteRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ClusterName, constants.GroupID, constants.SnapshotID}
var ReadRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.GroupID, constants.SnapshotID}
var ListRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.GroupID}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(CreateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progressEvents.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// progress callback setup
	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["snapshot_id"].(string)
		currentModel.SnapshotId = &sid
		return validateProgress(client, currentModel, "completed")
	}

	// Create Atlas API Request Object
	clusterName := cast.ToString(currentModel.ClusterName)
	projectID := cast.ToString(currentModel.GroupId)

	requestParameters := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}
	snapshotRequest := &mongodbatlas.CloudProviderSnapshot{
		RetentionInDays: cast.ToInt(currentModel.RetentionInDays),
		Description:     cast.ToString(currentModel.Description),
	}

	// API call to create snapshot
	snapshot, _, err := client.CloudProviderSnapshots.Create(context.Background(), requestParameters, snapshotRequest)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateCloudBackup, projectID, clusterName, err)
		return progressEvents.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeServiceInternalError), nil
	}
	currentModel.SnapshotId = &snapshot.ID

	_, _ = logger.Debugf("cloud backup created successfully %s", *currentModel.SnapshotId)

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
	setup() // logger setup
	_, _ = logger.Debugf("Read snapshot for Request() currentModel:%+v", currentModel)

	// Validate required fields in the request
	if modelValidation := validateModel(ReadRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	// cluster nae or instance name must present in the request
	if currentModel.ClusterName == nil && currentModel.InstanceName == nil {
		return progressEvents.GetFailedEventByCode("cluster Name or instance Name must present in the request",
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progressEvents.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	clusterName := cast.ToString(currentModel.ClusterName)
	instanceName := cast.ToString(currentModel.InstanceName)
	projectID := cast.ToString(currentModel.GroupId)

	if clusterName == "" && instanceName == "" {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Error - reading cloud backup  snapshot restore job: cluster name or instance name must be set ",
			ResourceModel:   currentModel,
		}, nil
	}
	// Check if  exist
	if !isSnapshotExist(currentModel) {
		_, _ = logger.Warnf("Error - Read snapshot with id(%s)", *currentModel.SnapshotId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	// Create Atlas API Request Object
	snapshotID := cast.ToString(currentModel.SnapshotId)
	var snapshot *mongodbatlas.CloudProviderSnapshot
	var res *mongodbatlas.Response
	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:    projectID,
		SnapshotID: snapshotID,
	}

	// API call to read snapshot
	if clusterName != "" {
		snapshotRequest.ClusterName = clusterName
		snapshot, res, err = client.CloudProviderSnapshots.GetOneCloudProviderSnapshot(context.Background(), snapshotRequest)
	} else {
		// read serverless instance snapshot
		snapshotRequest.InstanceName = instanceName
		snapshot, res, err = client.CloudProviderSnapshots.GetOneServerlessSnapshot(context.Background(), snapshotRequest)
	}

	if err != nil {
		return progressEvents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	currentModel = convertToUIModel(snapshot, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// NO OP
	return handler.ProgressEvent{
		OperationStatus: handler.Failed,
		Message:         "Not Implemented",
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(DeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressEvents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// check if exist
	if !isSnapshotExist(currentModel) {
		_, _ = logger.Warnf("Error - Read snapshot with id(%s)", *currentModel.SnapshotId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	// Create Atlas API Request Object
	clusterName := cast.ToString(currentModel.ClusterName)
	projectID := cast.ToString(currentModel.GroupId)
	snapshotID := cast.ToString(currentModel.SnapshotId)

	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		SnapshotID:  snapshotID,
		ClusterName: clusterName,
	}
	// API call to delete snapshot
	resp, err := client.CloudProviderSnapshots.Delete(context.Background(), snapshotRequest)

	if err != nil {
		return progressEvents.GetFailedEventByResponse(err.Error(), resp.Response), nil
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

	// Create Atlas API Request Object
	clusterName := cast.ToString(currentModel.ClusterName)
	instanceName := cast.ToString(currentModel.InstanceName)
	projectID := cast.ToString(currentModel.GroupId)
	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID: projectID,
	}
	var snapshots *mongodbatlas.CloudProviderSnapshots
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

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

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(ListRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	if currentModel.ClusterName == nil && currentModel.InstanceName == nil {
		return progressEvents.GetFailedEventByCode("cluster Name or instance Name must present in the request",
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressEvents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// Create Atlas API Request Object
	projectID := *currentModel.GroupId
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
		return progressEvents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// populate list
	snapshotsList := snapshots.Results
	for ind := range snapshotsList {
		var model Model
		models = append(models, convertToUIModel(snapshotsList[ind], &model))
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

// logger setup function
func setup() {
	util.SetupLogger("mongodb-atlas-backup-snapshot")
}

// converts mongodb link class to model link class
func flattenLinks(linksResult []*mongodbatlas.Link) []Link {
	links := make([]Link, 0)
	for ind := range linksResult {
		var lin Link
		lin.Href = &linksResult[ind].Href
		lin.Rel = &linksResult[ind].Rel
		links = append(links, lin)
	}
	return links
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

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
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
	currentModel.StorageSizeBytes = &snapShot.StorageSizeBytes
	currentModel.Links = flattenLinks(snapShot.Links)
	currentModel.CloudProvider = pointy.String(constants.AWS)
	currentModel.SnapshotIds = snapShot.SnapshotsIds
	currentModel.Members = flattenCloudMembers(snapShot.Members)
	return currentModel
}
