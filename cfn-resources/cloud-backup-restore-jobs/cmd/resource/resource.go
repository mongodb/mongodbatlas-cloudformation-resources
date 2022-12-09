package resource

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.PubKey, constants.SnapshotID, constants.PvtKey, constants.ClusterName, constants.ProjectID}
var ReadAndDeleteRequiredFields = []string{constants.PubKey, constants.ID, constants.PvtKey, constants.ClusterName, constants.ProjectID}
var ListRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ClusterName, constants.ProjectID}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("Create snapshot restore for Request() currentModel:%+v", currentModel)
	// Validate required fields in the request
	if modelValidation := validateModel(CreateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	deliveryType := currentModel.DeliveryType
	if deliveryType == nil || (*deliveryType != constants.Automated && *deliveryType != constants.Download) {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Error - creating cloud backup  snapshot restore job: you must specify either `automated` or `download` delivery types",
			ResourceModel:   currentModel,
		}, nil
	}
	requestParameters := &mongodbatlas.SnapshotReqPathParameters{}
	snapshotRequest := &mongodbatlas.CloudProviderSnapshotRestoreJob{}
	targetClusterName := currentModel.TargetClusterName
	targetProjectID := currentModel.TargetProjectId
	if *deliveryType == constants.Automated {
		if targetClusterName == nil || *targetClusterName == "" {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "Error - creating cloud backup  snapshot restore job: `target_cluster_name` must be set if delivery type is `automated`",
				ResourceModel:   currentModel,
			}, nil
		}
		if targetProjectID == nil || *targetProjectID == "" {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "Error - creating cloud backup  snapshot restore job: `target_project_id` must be set if delivery type is `automated`",
				ResourceModel:   currentModel,
			}, nil
		}
		snapshotRequest.TargetClusterName = *targetClusterName
		snapshotRequest.TargetGroupID = *targetProjectID
	}
	// Create Atlas API Request Object
	snapshotID := *currentModel.SnapshotId
	requestParameters.GroupID = *currentModel.ProjectId
	requestParameters.ClusterName = *currentModel.ClusterName
	snapshotRequest.SnapshotID = snapshotID
	snapshotRequest.DeliveryType = *deliveryType
	// API call to create job
	restoreJob, _, err := client.CloudProviderSnapshotRestoreJobs.Create(context.Background(), requestParameters, snapshotRequest)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateCloudBackupRestoreJob, snapshotID, err)
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Error - creating cloud backup  snapshot restore job",
			ResourceModel:   currentModel,
		}, nil
	}
	currentModel.Id = &restoreJob.ID
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// ValidateRequest function to validate the request
func ValidateRequest(currentModel *Model) (handler.ProgressEvent, *mongodbatlas.Client, error) {
	setup() // logger setup

	_, _ = logger.Debugf("Request - snapshot restore starts currentModel:%+v", currentModel)
	// Validate required fields in the request
	if modelValidation := validateModel(ReadAndDeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil, errors.New("required field not found")
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil, err
	}

	// Check if job already exist
	if !isRestoreJobExist(currentModel) {
		_, _ = logger.Warnf(constants.ErrorReadCloudBackUpRestoreJob, *currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil, errors.New(constants.ResourceNotFound)
	}
	return handler.ProgressEvent{}, client, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// validate the request
	event, client, err := ValidateRequest(currentModel)
	if err != nil {
		if err.Error() == constants.ResourceNotFound {
			return event, nil
		}
		return event, err
	}

	// Create Atlas API Request Object
	clusterName := *currentModel.ClusterName
	projectID := *currentModel.ProjectId
	jobID := *currentModel.Id
	requestParameters := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
		JobID:       jobID,
	}

	// API call
	restoreJob, _, err := client.CloudProviderSnapshotRestoreJobs.Get(context.Background(), requestParameters)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading cloud provider snapshot restore job with id(project: %s, job: %s): %s", projectID, jobID, err)
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
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// NO-OP
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// validate the request
	event, client, err := ValidateRequest(currentModel)
	if err != nil {
		if err.Error() == constants.ResourceNotFound {
			return event, nil
		}
		return event, err
	}

	// Create API Request Object
	clusterName := *currentModel.ClusterName
	projectID := *currentModel.ProjectId
	jobID := *currentModel.Id
	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
		JobID:       jobID,
	}

	if *currentModel.DeliveryType == "Automated" {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Automated restore cannot be cancelled",
			ResourceModel:   currentModel,
		}, nil
	}
	// API call to delete
	_, err = client.CloudProviderSnapshotRestoreJobs.Delete(context.Background(), snapshotRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting cloud provider snapshot restore job with id(project: %s, job: %s): %s", projectID, jobID, err)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete complete",
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("return all snapshot restore jobs for Request() currentModel:%+v", currentModel)
	// Validate required fields in the request
	if modelValidation := validateModel(ListRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// Create Atlas API Request Object
	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName

	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: clusterName,
	}
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}
	// API call
	restoreJobs, _, err := client.CloudProviderSnapshotRestoreJobs.List(context.Background(), snapshotRequest, params)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading cloud provider snapshot restore job list with id(project: %s): %s", projectID, err)
	}

	models := make([]interface{}, 0)
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
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List complete",
		ResourceModel:   models,
	}, nil
}

// function to check id restore job already exist
func isRestoreJobExist(currentModel *Model) (isExist bool) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return false
	}

	projectID := *currentModel.ProjectId
	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID:     projectID,
		ClusterName: *currentModel.ClusterName,
	}

	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	restoreJobs, _, err := client.CloudProviderSnapshotRestoreJobs.List(context.Background(), snapshotRequest, params)
	if err != nil {
		return false
	}
	for _, restoreJob := range restoreJobs.Results {
		_, _ = logger.Debugf("Read - Restore Job with id(%s): %s", restoreJob.ID, *currentModel.Id)
		if restoreJob.ID == *currentModel.Id && !restoreJob.Expired && !restoreJob.Cancelled {
			return true
		}
	}
	return false
}

// convert mongodb links to model links
func flattenLinks(linksResult []*mongodbatlas.Link) []Links {
	links := make([]Links, 0)
	for _, link := range linksResult {
		var lin Links
		lin.Href = &link.Href
		lin.Rel = &link.Rel
		links = append(links, lin)
	}
	return links
}

// function to set the logger
func setup() {
	util.SetupLogger("mongodb-atlas-project")
}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}
