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
	"go.mongodb.org/atlas-sdk/v20240805004/admin"
)

var CreateRequiredFields = []string{constants.SnapshotID, constants.DeliveryType, constants.InstanceType, constants.InstanceName}
var ReadDeleteRequiredFields = []string{constants.ID, constants.InstanceType, constants.InstanceName}
var ListRequiredFields = []string{constants.ProjectID, constants.InstanceType, constants.InstanceName}

const (
	defaultBackSeconds            = 30
	defaultTimeOutInSeconds       = 1200
	defaultReturnSuccessIfTimeOut = false
	clusterInstanceType           = "cluster"
	serverlessInstanceType        = "serverless"
)

func setup() {
	util.SetupLogger("mongodb-atlas-backup-restore-job")
}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	if pe := validator.ValidateModel(fields, model); pe != nil {
		return pe
	}

	if *model.InstanceType != clusterInstanceType && *model.InstanceType != serverlessInstanceType {
		pe := progressevent.GetFailedEventByCode(fmt.Sprintf("InstanceType must be %s or %s", clusterInstanceType, serverlessInstanceType),
			cloudformation.HandlerErrorCodeInvalidRequest)
		return &pe
	}

	return nil
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validateModel(CreateRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	err := currentModel.validateAsynchronousProperties()
	if err != nil {
		return progressevent.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), err
	}

	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		id := req.CallbackContext["id"].(string)
		startTime := req.CallbackContext["startTime"].(string)
		return createCallback(client, currentModel, id, startTime), nil
	}

	automated := constants.Automated
	if util.AreStringPtrEqual(currentModel.DeliveryType, &automated) {
		if !util.IsStringPresent(currentModel.TargetProjectId) || !util.IsStringPresent(currentModel.TargetClusterName) {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "Error - creating cloud backup  snapshot restore job: `TargetProjectId` and `TargetClusterName` must be set if delivery type is `automated`",
				ResourceModel:   currentModel,
			}, nil
		}
	}

	if *currentModel.InstanceType == serverlessInstanceType {
		params := paramsServerless(currentModel)
		serverless, resp, err := client.Atlas20231115014.CloudBackupsApi.CreateServerlessBackupRestoreJob(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, params).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}
		currentModel.Id = serverless.Id
	} else {
		params := paramsServer(currentModel)
		server, resp, err := client.Atlas20231115014.CloudBackupsApi.CreateBackupRestoreJob(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, params).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}
		currentModel.Id = server.Id
	}

	if aws.BoolValue(currentModel.EnableSynchronousCreation) {
		return progressevent.GetInProgressProgressEvent(
				"Create in progress",
				map[string]interface{}{
					constants.StateName: "in_progress",
					"id":                currentModel.Id,
					"startTime":         util.TimeToString(time.Now()),
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
	if err := validateModel(ReadDeleteRequiredFields, currentModel); err != nil {
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
	if err := validateModel(ReadDeleteRequiredFields, currentModel); err != nil {
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

	_, resp, err := client.Atlas20231115014.CloudBackupsApi.CancelBackupRestoreJob(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, *currentModel.Id).Execute()
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
	if err := validateModel(ListRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	models := make([]interface{}, 0)
	if *currentModel.InstanceType == serverlessInstanceType {
		serverless, resp, err := client.Atlas20231115014.CloudBackupsApi.ListServerlessBackupRestoreJobs(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}
		instanceType := serverlessInstanceType
		results := serverless.GetResults()
		for i := range results {
			job := &results[i]
			model := &Model{
				ProjectId:    currentModel.ProjectId,
				InstanceType: &instanceType,
				InstanceName: currentModel.InstanceName,
				Profile:      currentModel.Profile,
			}
			if !aws.BoolValue(job.Cancelled) && !aws.BoolValue(job.Expired) {
				updateModelServerless(model, job)
				models = append(models, model)
			}
		}
	} else {
		server, resp, err := client.Atlas20231115014.CloudBackupsApi.ListBackupRestoreJobs(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}
		instanceType := clusterInstanceType
		results := server.GetResults()
		for i := range results {
			job := &results[i]
			model := &Model{
				ProjectId:    currentModel.ProjectId,
				InstanceType: &instanceType,
				InstanceName: currentModel.InstanceName,
				Profile:      currentModel.Profile,
			}
			if !aws.BoolValue(job.Cancelled) && !aws.BoolValue(job.Expired) {
				updateModelServer(model, job)
				models = append(models, model)
			}
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

func createCallback(client *util.MongoDBClient, currentModel *Model, jobID, startTime string) handler.ProgressEvent {
	currentModel.Id = &jobID
	if err := updateModel(client, currentModel, false); err != nil {
		return *err
	}

	if util.IsStringPresent(currentModel.FinishedAt) {
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
	startDateTime := util.StringPtrToTimePtr(&startTime)
	if startDateTime == nil {
		return false // If there's an error parsing the start time, assume timeout is not reached
	}

	// Calculate the timeout time by adding timeOutInSeconds to the startDateTime
	timeoutTime := startDateTime.Add(time.Duration(timeOutInSeconds) * time.Second)

	// Get the current time
	currentTime := time.Now()

	// Compare the current time with the timeout time
	return currentTime.After(timeoutTime)
}

func updateModel(client *util.MongoDBClient, model *Model, checkFinish bool) *handler.ProgressEvent {
	if *model.InstanceType == serverlessInstanceType {
		serverless, resp, err := client.Atlas20231115014.CloudBackupsApi.GetServerlessBackupRestoreJob(context.Background(), *model.ProjectId, *model.InstanceName, *model.Id).Execute()
		if err != nil {
			pe := progressevent.GetFailedEventByResponse(err.Error(), resp)
			return &pe
		}
		updateModelServerless(model, serverless)
	} else {
		server, resp, err := client.Atlas20231115014.CloudBackupsApi.GetBackupRestoreJob(context.Background(), *model.ProjectId, *model.InstanceName, *model.Id).Execute()
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
	model.Id = job.Id
	model.FinishedAt = util.TimePtrToStringPtr(job.FinishedAt)
	model.SnapshotId = job.SnapshotId
	model.TargetProjectId = &job.TargetGroupId
	model.Timestamp = util.TimePtrToStringPtr(job.Timestamp)
	model.Cancelled = job.Cancelled
	model.Expired = job.Expired
	model.DeliveryUrl = job.GetDeliveryUrl()
	model.Links = flattenLinks(job.GetLinks())
}

func updateModelServer(model *Model, job *admin.DiskBackupSnapshotRestoreJob) {
	model.TargetClusterName = job.TargetClusterName
	model.DeliveryType = &job.DeliveryType
	model.ExpiresAt = util.TimePtrToStringPtr(job.ExpiresAt)
	model.Id = job.Id
	model.FinishedAt = util.TimePtrToStringPtr(job.FinishedAt)
	model.SnapshotId = job.SnapshotId
	model.TargetProjectId = job.TargetGroupId
	model.Timestamp = util.TimePtrToStringPtr(job.Timestamp)
	model.Cancelled = job.Cancelled
	model.Failed = job.Failed
	model.Expired = job.Expired
	model.DeliveryUrl = job.GetDeliveryUrl()
	model.Links = flattenLinks(job.GetLinks())
}

func flattenLinks(linksResult []admin.Link) []Links {
	links := make([]Links, 0)
	for _, link := range linksResult {
		var lin Links
		lin.Href = link.Href
		lin.Rel = link.Rel
		links = append(links, lin)
	}
	return links
}

func paramsServer(model *Model) *admin.DiskBackupSnapshotRestoreJob {
	return &admin.DiskBackupSnapshotRestoreJob{
		SnapshotId:            model.SnapshotId,
		DeliveryType:          *model.DeliveryType,
		TargetClusterName:     model.TargetClusterName,
		TargetGroupId:         model.TargetProjectId,
		OplogTs:               util.StrPtrToIntPtr(model.OpLogTs),
		OplogInc:              util.StrPtrToIntPtr(model.OpLogInc),
		PointInTimeUTCSeconds: model.PointInTimeUtcSeconds,
	}
}

func paramsServerless(model *Model) *admin.ServerlessBackupRestoreJob {
	return &admin.ServerlessBackupRestoreJob{
		SnapshotId:            model.SnapshotId,
		DeliveryType:          *model.DeliveryType,
		TargetClusterName:     *model.TargetClusterName,
		TargetGroupId:         *model.TargetProjectId,
		OplogTs:               util.StrPtrToIntPtr(model.OpLogTs),
		OplogInc:              util.StrPtrToIntPtr(model.OpLogInc),
		PointInTimeUTCSeconds: model.PointInTimeUtcSeconds,
	}
}
