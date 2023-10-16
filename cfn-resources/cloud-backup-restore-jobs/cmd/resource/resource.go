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
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20231001001/admin"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.SnapshotID, constants.DeliveryType, constants.InstanceType, constants.InstanceName}
var ReadDeleteRequiredFields = []string{constants.ID, constants.InstanceType, constants.InstanceName}
var ListRequiredFields = []string{constants.ProjectID, constants.InstanceType, constants.InstanceName}

const (
	defaultBackSeconds            = 30
	defaultTimeOutInSeconds       = 1200
	defaultReturnSuccessIfTimeOut = false
	timeLayout                    = "2006-01-02 15:04:05"
	clusterInstanceType           = "cluster"
	serverlessInstanceType        = "serverless"
)

func setup() {
	util.SetupLogger("mongodb-atlas-backup-restore-job")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(CreateRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	err := currentModel.validateAsynchronousProperties()
	if err != nil {
		return progressevent.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), err
	}

	// Callback
	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		id := req.CallbackContext["id"]
		startTime := req.CallbackContext["startTime"]
		return createCallback(client, currentModel, cast.ToString(id), cast.ToString(startTime)), nil
	}

	targetClusterName := cast.ToString(currentModel.TargetClusterName)
	targetProjectID := cast.ToString(currentModel.TargetProjectId)
	deliveryType := cast.ToString(currentModel.DeliveryType)
	snapshotID := cast.ToString(currentModel.SnapshotId)

	clusterName := ""
	instanceName := ""

	if *currentModel.InstanceType == clusterInstanceType {
		clusterName = cast.ToString(currentModel.InstanceName)
	} else {
		instanceName = cast.ToString(currentModel.InstanceName)
	}

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

	if *currentModel.InstanceType == clusterInstanceType {
		requestParameters := &mongodbatlas.SnapshotReqPathParameters{
			GroupID:     cast.ToString(currentModel.ProjectId),
			SnapshotID:  snapshotID,
			ClusterName: clusterName,
		}
		// API call to create job
		restoreJob, resp, err := client.CloudProviderSnapshotRestoreJobs.Create(context.Background(), requestParameters, snapshotReq)
		if err != nil {
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error - creating  snapshot restore job for dedicated cluster: %+v", err), resp.Response), nil
		}
		currentModel.Id = &restoreJob.ID
	} else {
		// API call to create job
		restoreJob, resp, err := client.CloudProviderSnapshotRestoreJobs.CreateForServerlessBackupRestore(context.Background(), *currentModel.ProjectId, instanceName, snapshotReq)
		if err != nil {
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error - creating  snapshot restore job for serverless cluster: %+v", err), resp.Response), nil
		}
		currentModel.Id = &restoreJob.ID
	}

	if flowIsSynchronous(currentModel) {
		return progressevent.GetInProgressProgressEvent(
				"Create in progress",
				map[string]interface{}{
					constants.StateName: "in_progress",
					"id":                currentModel.Id,
					"startTime":         time.Now().Format(timeLayout),
				},
				currentModel,
				int64(*currentModel.SynchronousCreationOptions.CallbackDelaySeconds)),
			nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(ReadDeleteRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	if err := updateModel(client, currentModel, true); err != nil {
		return *err, nil
	}

	if aws.BoolValue(currentModel.Cancelled) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "The job is in status cancelled, Cannot read a cancelled job",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read complete",
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(ReadDeleteRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	if err := updateModel(client, currentModel, true); err != nil {
		return *err, nil
	}

	if aws.BoolValue(currentModel.Cancelled) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "The job is in status cancelled, Cannot delete a cancelled job",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	if util.IsStringPresent(currentModel.FinishedAt) || aws.BoolValue(currentModel.Failed) || aws.BoolValue(currentModel.Expired) {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "The resource is finished, failed, or expired",
		}, nil
	}

	automated := "automated"
	if util.AreStringPtrEqual(currentModel.DeliveryType, &automated) {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Automated restore cannot be cancelled, wait until the process is finished and try again",
			ResourceModel:   currentModel,
		}, nil
	}

	_, resp, err := client.AtlasV2.CloudBackupsApi.CancelBackupRestoreJob(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, *currentModel.Id).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(ListRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	// Create Atlas API Request Object
	var restoreJobs *mongodbatlas.CloudProviderSnapshotRestoreJobs
	var resp *mongodbatlas.Response

	clusterName := ""
	instanceName := ""

	if *currentModel.InstanceType == clusterInstanceType {
		clusterName = cast.ToString(currentModel.InstanceName)
	} else {
		instanceName = cast.ToString(currentModel.InstanceName)
	}

	projectID := cast.ToString(currentModel.ProjectId)
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	var err error
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
		return progressevent.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	models := make([]interface{}, 0)
	restoreJobsList := restoreJobs.Results
	for ind := range restoreJobsList {
		var model Model
		model.ProjectId = currentModel.ProjectId
		model.InstanceName = currentModel.InstanceName
		model.InstanceType = currentModel.InstanceType
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

func (model *Model) validateAsynchronousProperties() error {
	if model.EnableSynchronousCreation != nil && *model.EnableSynchronousCreation {
		if model.SynchronousCreationOptions == nil {
			model.SynchronousCreationOptions = &SynchronousCreationOptions{}
		}

		if model.SynchronousCreationOptions.CallbackDelaySeconds == nil {
			model.SynchronousCreationOptions.CallbackDelaySeconds = aws.Int(defaultBackSeconds)
		}

		if model.SynchronousCreationOptions.TimeOutInSeconds == nil {
			model.SynchronousCreationOptions.TimeOutInSeconds = aws.Int(defaultTimeOutInSeconds)
		}

		if model.SynchronousCreationOptions.ReturnSuccessIfTimeOut == nil {
			model.SynchronousCreationOptions.ReturnSuccessIfTimeOut = aws.Bool(defaultReturnSuccessIfTimeOut)
		}
	}

	return nil
}

func flowIsSynchronous(model *Model) bool {
	return model.EnableSynchronousCreation != nil && *model.EnableSynchronousCreation
}

func createCallback(client *mongodbatlas.Client, currentModel *Model, jobID, startTime string) handler.ProgressEvent {
	restoreJob, progressEvent := getRestoreJob(client, currentModel)
	if progressEvent != nil {
		return *progressEvent
	}

	currentModel.Id = &jobID

	if restoreJob.FinishedAt != "" {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Create Complete",
			ResourceModel:   currentModel,
		}
	}

	if isTimeOutReached(startTime, *currentModel.SynchronousCreationOptions.TimeOutInSeconds) {
		if *currentModel.SynchronousCreationOptions.ReturnSuccessIfTimeOut {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Create Complete - the resource was completed with timeout",
				ResourceModel:   currentModel,
			}
		}

		return progressevent.GetFailedEventByCode("Create failed with Timout", cloudformation.HandlerErrorCodeInternalFailure)
	}

	return progressevent.GetInProgressProgressEvent(
		"Create in progress",
		map[string]interface{}{
			constants.StateName: "in_progress",
			"id":                currentModel.Id,
			"startTime":         startTime,
		},
		currentModel,
		int64(*currentModel.SynchronousCreationOptions.CallbackDelaySeconds))
}

func isTimeOutReached(startTime string, timeOutInSeconds int) bool {
	startDateTime, err := time.Parse(timeLayout, startTime)
	if err != nil {
		return false // If there's an error parsing the start time, assume timeout is not reached
	}

	// Calculate the timeout time by adding timeOutInSeconds to the startDateTime
	timeoutTime := startDateTime.Add(time.Duration(timeOutInSeconds) * time.Second)

	// Get the current time
	currentTime := time.Now()

	// Compare the current time with the timeout time
	return currentTime.After(timeoutTime)
}

func getRestoreJob(client *mongodbatlas.Client, currentModel *Model) (*mongodbatlas.CloudProviderSnapshotRestoreJob, *handler.ProgressEvent) {
	if *currentModel.InstanceType == serverlessInstanceType {
		/*projectID, instanceName, jobID*/
		restoreJobs, resp, err := client.CloudProviderSnapshotRestoreJobs.GetForServerlessBackupRestore(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, *currentModel.Id)
		if err != nil {
			pe := progressevent.GetFailedEventByResponse("Error getting response job", resp.Response)
			return nil, &pe
		}
		return restoreJobs, nil
	}

	snapshotRequest := &mongodbatlas.SnapshotReqPathParameters{
		GroupID: *currentModel.ProjectId,
		JobID:   *currentModel.Id,
	}

	if *currentModel.InstanceType == clusterInstanceType {
		snapshotRequest.ClusterName = *currentModel.InstanceName
	} else {
		snapshotRequest.InstanceName = *currentModel.InstanceName
	}

	restoreJobs, resp, err := client.CloudProviderSnapshotRestoreJobs.Get(context.Background(), snapshotRequest)
	if err != nil {
		pe := progressevent.GetFailedEventByResponse("Error getting response job", resp.Response)
		return nil, &pe
	}
	return restoreJobs, nil
}

func isJobFinished(job mongodbatlas.CloudProviderSnapshotRestoreJob) bool {
	return job.FinishedAt != ""
}

func isJobFailed(job mongodbatlas.CloudProviderSnapshotRestoreJob) bool {
	return job.Failed != nil && *job.Failed
}

func updateModel(client *util.MongoDBClient, model *Model, checkFinish bool) *handler.ProgressEvent {
	if *model.InstanceType == serverlessInstanceType {
		serverless, resp, err := client.AtlasV2.CloudBackupsApi.GetServerlessBackupRestoreJob(context.Background(), *model.ProjectId, *model.InstanceName, *model.Id).Execute()
		if err != nil {
			pe := progressevent.GetFailedEventByResponse(err.Error(), resp)
			return &pe
		}
		updateModelServerless(model, serverless)
	} else {
		server, resp, err := client.AtlasV2.CloudBackupsApi.GetBackupRestoreJob(context.Background(), *model.ProjectId, *model.InstanceName, *model.Id).Execute()
		if err != nil {
			pe := progressevent.GetFailedEventByResponse(err.Error(), resp)
			return &pe
		}
		updateModelServer(model, server)
	}
	return nil
}

func updateModelServerless(model *Model, job *admin.ServerlessBackupRestoreJob) {
	model.TargetClusterName = &job.TargetClusterName
	model.DeliveryType = &job.DeliveryType
	model.ExpiresAt = util.TimePtrToStringPtr(job.ExpiresAt)
	model.CreatedAt = nil // not available if new SDK
	model.Id = job.Id
	model.FinishedAt = util.TimePtrToStringPtr(job.FinishedAt)
	model.SnapshotId = job.SnapshotId
	model.TargetProjectId = &job.TargetGroupId
	model.Timestamp = util.TimePtrToStringPtr(job.Timestamp)
	model.Cancelled = job.Cancelled
	model.Expired = job.Expired
	model.DeliveryUrl = job.DeliveryUrl
	model.Links = flattenLinks2(job.Links)
}

func updateModelServer(model *Model, job *admin.DiskBackupSnapshotRestoreJob) {
	model.TargetClusterName = job.TargetClusterName
	model.DeliveryType = &job.DeliveryType
	model.ExpiresAt = util.TimePtrToStringPtr(job.ExpiresAt)
	model.CreatedAt = nil // not available in new SDK
	model.Id = job.Id
	model.FinishedAt = util.TimePtrToStringPtr(job.FinishedAt)
	model.SnapshotId = job.SnapshotId
	model.TargetProjectId = job.TargetGroupId
	model.Timestamp = util.TimePtrToStringPtr(job.Timestamp)
	model.Cancelled = job.Cancelled
	model.Failed = job.Failed
	model.Expired = job.Expired
	model.DeliveryUrl = job.DeliveryUrl
	model.Links = flattenLinks2(job.Links)
}

func flattenLinks2(linksResult []admin.Link) []Links {
	links := make([]Links, 0)
	for _, link := range linksResult {
		var lin Links
		lin.Href = link.Href
		lin.Rel = link.Rel
		links = append(links, lin)
	}
	return links
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
