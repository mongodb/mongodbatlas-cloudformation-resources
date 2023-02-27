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
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.SnapshotID, constants.DeliveryType}
var ReadDeleteRequiredFields = []string{constants.ID}
var ListRequiredFields = []string{constants.ProjectID}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(CreateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	targetClusterName := cast.ToString(currentModel.TargetClusterName)
	targetProjectID := cast.ToString(currentModel.TargetProjectId)
	deliveryType := cast.ToString(currentModel.DeliveryType)
	clusterName := cast.ToString(currentModel.ClusterName)
	instanceName := cast.ToString(currentModel.InstanceName)
	snapshotID := cast.ToString(currentModel.SnapshotId)

	// check target cluster and project set for automated download
	if deliveryType == constants.Automated {
		if targetClusterName == "" {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "Error - creating cloud backup  snapshot restore job: `TargetClusterName` must be set if delivery type is `automated`",
				ResourceModel:   currentModel,
			}, nil
		}
		if targetProjectID == "" {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "Error - creating cloud backup  snapshot restore job: `TargetProjectId` must be set if delivery type is `automated`",
				ResourceModel:   currentModel,
			}, nil
		}
	}

	if clusterName == "" && instanceName == "" {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Error - creating cloud backup  snapshot restore job: cluster name or instance name must be set ",
			ResourceModel:   currentModel,
		}, nil
	}
	// Create Atlas API Request Object
	snapshotReq := &mongodbatlas.CloudProviderSnapshotRestoreJob{
		SnapshotID:            snapshotID,
		DeliveryType:          deliveryType,
		TargetClusterName:     targetClusterName,
		TargetGroupID:         targetProjectID,
		OplogTs:               cast.ToInt64(currentModel.OpLogTs),
		OplogInc:              cast.ToInt64(currentModel.OpLogInc),
		PointInTimeUTCSeconds: cast.ToInt64(currentModel.PointInTimeUtcSeconds),
	}

	if clusterName != "" {
		requestParameters := &mongodbatlas.SnapshotReqPathParameters{
			GroupID:     cast.ToString(currentModel.ProjectId),
			SnapshotID:  snapshotID,
			ClusterName: clusterName,
		}
		// API call to create job
		restoreJob, resp, err := client.CloudProviderSnapshotRestoreJobs.Create(context.Background(), requestParameters, snapshotReq)
		if err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error - creating  snapshot restore job for dedicated cluster: %+v", err), resp.Response), nil
		}
		currentModel.Id = &restoreJob.ID
	} else {
		// API call to create job
		restoreJob, resp, err := client.CloudProviderSnapshotRestoreJobs.CreateForServerlessBackupRestore(context.Background(), *currentModel.ProjectId, instanceName, snapshotReq)
		if err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error - creating  snapshot restore job for serverless cluster: %+v", err), resp.Response), nil
		}
		currentModel.Id = &restoreJob.ID
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(ReadDeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	clusterName := cast.ToString(currentModel.ClusterName)
	instanceName := cast.ToString(currentModel.InstanceName)
	projectID := cast.ToString(currentModel.ProjectId)
	jobID := cast.ToString(currentModel.Id)

	if clusterName == "" && instanceName == "" {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Error - reading cloud backup  snapshot restore job: cluster name or instance name must be set ",
			ResourceModel:   currentModel,
		}, nil
	}
	// Check if job already exist
	if !isRestoreJobExist(client, currentModel) {
		_, _ = logger.Warnf(constants.ErrorReadCloudBackUpRestoreJob, *currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	// Create Atlas API Request Object
	if clusterName != "" {
		requestParameters := &mongodbatlas.SnapshotReqPathParameters{
			GroupID:     projectID,
			ClusterName: cast.ToString(currentModel.ClusterName),
			JobID:       jobID,
		}
		restoreJob, resp, err := client.CloudProviderSnapshotRestoreJobs.Get(context.Background(), requestParameters)
		if err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("error reading restore job with id(project: %s, job: %s): %+v", projectID, jobID, err), resp.Response), nil
		}
		currentModel = convertToUIModel(restoreJob, currentModel)
	} else {
		// API call to create job
		restoreJob, resp, err := client.CloudProviderSnapshotRestoreJobs.GetForServerlessBackupRestore(context.Background(), projectID, instanceName, jobID)
		if err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("error reading restore job for serverless instance with id(project: %s, job: %s): %+v", projectID, jobID, err), resp.Response), nil
		}
		currentModel = convertToUIModel(restoreJob, currentModel)
	}
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

	// Validate required fields in the request
	if modelValidation := validateModel(ReadDeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	// Check if job already exist
	if !isRestoreJobExist(client, currentModel) {
		_, _ = logger.Warnf("restore job not fund for id :%s", *currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	projectID := *currentModel.ProjectId
	jobID := *currentModel.Id
	clusterName := cast.ToString(currentModel.ClusterName)
	if clusterName != "" {
		// Check if delivery type is automated
		if currentModel.DeliveryType != nil && *currentModel.DeliveryType == "automated" {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "Automated restore cannot be cancelled",
				ResourceModel:   currentModel,
			}, nil
		}

		// Create API Request Object
		requestParameters := &mongodbatlas.SnapshotReqPathParameters{
			GroupID:     projectID,
			ClusterName: cast.ToString(currentModel.ClusterName),
			JobID:       jobID,
		}

		// API call to delete
		_, err := client.CloudProviderSnapshotRestoreJobs.Delete(context.Background(), requestParameters)
		if err != nil {
			return handler.ProgressEvent{}, fmt.Errorf("error deleting cloud provider snapshot restore job with id(project: %s, job: %s): %s", projectID, jobID, err)
		}
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete complete",
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	var err error

	// Validate required fields in the request
	if modelValidation := validateModel(ListRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	// Create Atlas API Request Object
	var restoreJobs *mongodbatlas.CloudProviderSnapshotRestoreJobs
	var resp *mongodbatlas.Response

	clusterName := cast.ToString(currentModel.ClusterName)
	instanceName := cast.ToString(currentModel.InstanceName)
	projectID := cast.ToString(currentModel.ProjectId)
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	if clusterName == "" && instanceName == "" {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Error - creating cloud backup  snapshot restore job: cluster name or instance name must be set ",
			ResourceModel:   currentModel,
		}, nil
	}

	if clusterName != "" {
		snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
			GroupID:     projectID,
			ClusterName: clusterName,
		}
		// API call to list dedicated cluster restore jobs
		restoreJobs, resp, err = client.CloudProviderSnapshotRestoreJobs.List(context.Background(), snapshotRequest, params)
	} else {
		// API call to list serverless instance jobs
		restoreJobs, resp, err = client.CloudProviderSnapshotRestoreJobs.ListForServerlessBackupRestore(context.Background(), *currentModel.ProjectId, instanceName, params)
	}

	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	models := make([]interface{}, 0)
	restoreJobsList := restoreJobs.Results
	for ind := range restoreJobsList {
		var model Model
		model.ProjectId = currentModel.ProjectId
		model.ClusterName = currentModel.ClusterName
		model.Profile = currentModel.Profile
		if !restoreJobsList[ind].Cancelled && !restoreJobsList[ind].Expired {
			models = append(models, *convertToUIModel(restoreJobsList[ind], &model))
		}
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List complete",
		ResourceModels:  models,
	}, nil
}

// function to check id restore job already exist
func isRestoreJobExist(client *mongodbatlas.Client, currentModel *Model) (isExist bool) {
	var restoreJobs *mongodbatlas.CloudProviderSnapshotRestoreJobs
	var err error
	clusterName := cast.ToString(currentModel.ClusterName)
	instanceName := cast.ToString(currentModel.InstanceName)
	projectID := cast.ToString(currentModel.ProjectId)
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	if clusterName != "" {
		snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
			GroupID:     projectID,
			ClusterName: clusterName,
		}
		// API call to list dedicated cluster restore jobs
		restoreJobs, _, err = client.CloudProviderSnapshotRestoreJobs.List(context.Background(), snapshotRequest, params)
	} else {
		// API call to list serverless instance jobs
		restoreJobs, _, err = client.CloudProviderSnapshotRestoreJobs.ListForServerlessBackupRestore(context.Background(), *currentModel.ProjectId, instanceName, params)
	}
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
func convertToUIModel(restoreJob *mongodbatlas.CloudProviderSnapshotRestoreJob, currentModel *Model) *Model {
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
	return currentModel
}

// function to set the logger
func setup() {
	util.SetupLogger("mongodb-atlas-backup-restore-job")
}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}
