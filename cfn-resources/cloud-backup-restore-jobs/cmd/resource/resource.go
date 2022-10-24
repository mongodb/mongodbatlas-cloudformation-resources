package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/cloud-backup-restore-jobs/cmd/validator_def"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	log "github.com/sirupsen/logrus"
	matlasClient "go.mongodb.org/atlas/mongodbatlas"
)

const (
	automated = "automated"
	download  = "download"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() //logger setup

	log.Debugf("Create snapshot restore for Request() currentModel:%+v", currentModel)
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	// Validate required fields in the request
	modelValidation := validateModel(constants.Create, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	deliveryType := currentModel.DeliveryType
	if deliveryType == nil || (*deliveryType != automated && *deliveryType != download) {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "error creating cloud provider snapshot restore job: you must specify either `automated` or `download` delivery types",
			ResourceModel:   currentModel,
		}, nil
	}
	requestParameters := &matlasClient.SnapshotReqPathParameters{}
	snapshotRequest := &matlasClient.CloudProviderSnapshotRestoreJob{}
	targetClusterName := currentModel.TargetClusterName
	targetProjectId := currentModel.TargetProjectId
	if *deliveryType == automated {
		if targetClusterName == nil || *targetClusterName == "" {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "error creating cloud provider snapshot restore job: `target_cluster_name` must be set if delivery type is `automated`",
				ResourceModel:   currentModel,
			}, nil
		}
		if targetProjectId == nil || *targetProjectId == "" {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "error creating cloud provider snapshot restore job: `target_project_id` must be set if delivery type is `automated`",
				ResourceModel:   currentModel,
			}, nil
		}
		snapshotRequest.TargetClusterName = *targetClusterName
		snapshotRequest.TargetGroupID = *targetProjectId
	}
	// Create Atlas API Request Object
	requestParameters.GroupID = *currentModel.ProjectId
	requestParameters.ClusterName = *currentModel.ClusterName
	snapshotRequest.SnapshotID = *currentModel.SnapshotId
	snapshotRequest.DeliveryType = *deliveryType
	// API call to create job
	restoreJob, _, err := client.CloudProviderSnapshotRestoreJobs.Create(context.Background(), requestParameters, snapshotRequest)

	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating cloud provider snapshot restore job: %s", err)
	}

	currentModel.Id = &restoreJob.ID
	log.Info("Create snapshot restore ends")

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() //logger setup

	log.Debugf("Read snapshot restore starts  Request() currentModel:%+v", currentModel)
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	// Validate required fields in the request
	modelValidation := validateModel(constants.Read, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	// Create Atlas API Request Object
	clusterName := *currentModel.ClusterName
	projectId := *currentModel.ProjectId
	jobId := *currentModel.Id
	requestParameters := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		ClusterName: clusterName,
		JobID:       jobId,
	}
	isExist := isSnapshotExist(currentModel)
	// Check if job already exist
	if !isExist {
		log.Infof("Read - errors reading restore with id(%s)", jobId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	// API call
	restoreJob, _, err := client.CloudProviderSnapshotRestoreJobs.Get(context.Background(), requestParameters)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading cloud provider snapshot restore job with id(project: %s, job: %s): %s", projectId, jobId, err)
	}

	currentModel.TargetClusterName = &restoreJob.TargetClusterName
	currentModel.DeliveryType = &restoreJob.DeliveryType
	currentModel.ExpiresAt = &restoreJob.ExpiresAt
	currentModel.CreatedAt = &restoreJob.CreatedAt
	currentModel.Id = &restoreJob.ID
	currentModel.FinishedAt = &restoreJob.FinishedAt
	currentModel.SnapshotId = &restoreJob.SnapshotID
	currentModel.TargetProjectId = &restoreJob.TargetGroupID
	currentModel.Timestamp = &restoreJob.Timestamp
	currentModel.Cancelled = &restoreJob.Cancelled
	currentModel.Expired = &restoreJob.Expired
	currentModel.DeliveryUrl = restoreJob.DeliveryURL
	currentModel.Links = flattenLinks(restoreJob.Links)
	log.Info("Read snapshot restore ends")
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	//NO-OP
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() //logger setup

	log.Debugf("Delete snapshot restore starts for Request() currentModel:%+v", currentModel)
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	// Validate required fields in the request
	modelValidation := validateModel(constants.Delete, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	// Create API Request Object
	clusterName := *currentModel.ClusterName
	projectId := *currentModel.ProjectId
	jobId := *currentModel.Id
	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		ClusterName: clusterName,
		JobID:       jobId,
	}
	isExist := isSnapshotExist(currentModel)
	// Check if snapshot already exist
	if !isExist {
		log.Infof("Delete - errors reading restore with id(%s)", jobId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	if *currentModel.DeliveryType == "automated" {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Automated restore cannot be cancelled",
			ResourceModel:   currentModel,
		}, nil
	}
	//API call to delete
	_, err = client.CloudProviderSnapshotRestoreJobs.Delete(context.Background(), snapshotRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting cloud provider snapshot restore job with id(project: %s, job: %s): %s", projectId, jobId, err)
	}
	log.Info("Delete snapshot restore ends")
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete complete",
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() //logger setup

	log.Debugf("return all snapshot restore jobs for Request() currentModel:%+v", currentModel)
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	// Validate required fields in the request
	modelValidation := validateModel(constants.List, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
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
	// API call
	restoreJobs, _, err := client.CloudProviderSnapshotRestoreJobs.List(context.Background(), snapshotRequest, params)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading cloud provider snapshot restore job list with id(project: %s): %s", projectId, err)
	}

	var models []Model
	for _, restoreJob := range restoreJobs.Results {
		var model Model
		model.TargetClusterName = &restoreJob.TargetClusterName
		model.DeliveryType = &restoreJob.DeliveryType
		model.ExpiresAt = &restoreJob.ExpiresAt
		model.CreatedAt = &restoreJob.CreatedAt
		model.Id = &restoreJob.ID
		model.FinishedAt = &restoreJob.FinishedAt
		model.SnapshotId = &restoreJob.SnapshotID
		model.TargetProjectId = &restoreJob.TargetGroupID
		model.Timestamp = &restoreJob.Timestamp
		model.Cancelled = &restoreJob.Cancelled
		model.Expired = &restoreJob.Expired
		model.DeliveryUrl = restoreJob.DeliveryURL
		model.Links = flattenLinks(restoreJob.Links)
		models = append(models, model)
	}
	log.Info("List snapshot restore ends")
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List complete",
		ResourceModel:   models,
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

	snapshots, _, err := client.CloudProviderSnapshotRestoreJobs.List(context.Background(), snapshotRequest, params)
	if err != nil {
		return false
	}
	for _, snapshot := range snapshots.Results {
		log.Infof("Read - errors reading snapshot with id(%s): %s", snapshot.ID, *currentModel.Id)
		if snapshot.ID == *currentModel.Id && !snapshot.Expired && !snapshot.Cancelled {

			return true
		}

	}

	return false
}
func flattenDeliveryUrl(deliveryUrlResult []string) []*string {
	deliveryUrls := make([]*string, 0)
	for _, deliveryUrl := range deliveryUrlResult {
		deliveryUrls = append(deliveryUrls, &deliveryUrl)
	}
	return deliveryUrls
}

func flattenLinks(linksResult []*matlasClient.Link) []Links {
	links := make([]Links, 0)
	for _, link := range linksResult {
		var lin Links
		lin.Href = &link.Href
		lin.Rel = &link.Rel
		links = append(links, lin)
	}
	return links
}
func setup() {
	util.SetupLogger("mongodb-atlas-project")

}

// function to validate inputs to all actions
func validateModel(event constants.Event, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(event, validator_def.ModelValidator{}, model)
}
