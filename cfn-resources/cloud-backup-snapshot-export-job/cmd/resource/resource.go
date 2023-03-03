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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.ExportBucketID, constants.SnapshotID}
var ReadRequiredFields = []string{constants.ProjectID, constants.ExportID, constants.ClusterName}
var ListRequiredFields = []string{constants.ProjectID, constants.ClusterName}

const (
	ErrorExportJobCreate = "error creating Export Job for the project(%s) : %s"
	ErrorExportJobRead   = "error reading export job for the projects(%s) : Job Id : %s with error :%+v"
)

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-CloudBackupSnapshotExportJob")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(CreateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == ""{
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// create API request object
	clusterName := cast.ToString(currentModel.ClusterName)
	snapshotID := cast.ToString(currentModel.SnapshotId)
	bucketID := cast.ToString(currentModel.ExportBucketId)
	projectID := cast.ToString(currentModel.ProjectId)
	request := &mongodbatlas.CloudProviderSnapshotExportJob{
		SnapshotID:     snapshotID,
		ExportBucketID: bucketID,
		CustomData:     expandExportJobCustomData(currentModel),
	}

	// progress callback setup
	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["export_id"].(string)
		currentModel.ExportId = &sid
		return validateProgress(client, currentModel, "Successful")
	}

	// API call to create export job
	jobResponse, resp, err := client.CloudProviderSnapshotExportJobs.Create(context.Background(), projectID, clusterName, request)
	if err != nil {
		_, _ = logger.Warnf(ErrorExportJobCreate, projectID, err)
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	currentModel.ExportId = &jobResponse.ID

	// logic included for CFN Test starts-a workaround for missing delete handler
	if currentModel.TestMode != nil {
		_, _ = util.DeleteKey(*currentModel.ProjectId, cast.ToString(currentModel.ExportId), req.Session)
	}
	// logic included for CFN Test ends

	_, _ = logger.Debugf("Created Successfully - (%s)", *currentModel.ExportId)

	// track progress
	event := handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create export snapshots : %s", jobResponse.ID),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 65,
		CallbackContext: map[string]interface{}{
			"status":    jobResponse.State,
			"export_id": jobResponse.ID,
		},
	}
	return event, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// logic included for CFN Test starts-a workaround for missing delete handler
	if getDeleteStatus(req, currentModel) {
		return progressevents.GetFailedEventByCode("Resource Not Found",
			cloudformation.HandlerErrorCodeNotFound), nil
	}
	//  logic included for CFN Test end

	// Validate required fields in the request
	if modelValidation := validateModel(ReadRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == ""{
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	clusterName := cast.ToString(currentModel.ClusterName)
	exportJobID := cast.ToString(currentModel.ExportId)
	projectID := cast.ToString(currentModel.ProjectId)

	// API call to read export job
	exportJob, resp, err := client.CloudProviderSnapshotExportJobs.Get(context.Background(), projectID, clusterName, exportJobID)
	if err != nil {
		_, _ = logger.Warnf(ErrorExportJobRead, projectID, exportJobID, err)
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}
	readModel := convertToModel(exportJob, *currentModel, resp.Links)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   readModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// NO OP
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// No OP .

	// logic included for CFN Test starts-a workaround for missing delete handler
	if currentModel.TestMode != nil && util.Get(*currentModel.ProjectId, cast.ToString(currentModel.ExportId), req.Session) != "" {
		return progressevents.GetFailedEventByCode("Resource Not Found", cloudformation.HandlerErrorCodeNotFound), nil
	}
	_, _ = util.PutKey(*currentModel.ProjectId, "deleted", cast.ToString(currentModel.ExportId), req.Session)
	// logic included for CFN Test ends

	return handler.ProgressEvent{
		Message:         "DeleteKey Complete",
		OperationStatus: handler.Success,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// logic included for CFN Test starts-a workaround for missing delete handler
	if getDeleteStatus(req, currentModel) {
		return handleCFNTestList(req, currentModel)
	}
	// logic included for CFN Test ends

	// Validate required fields in the request
	if modelValidation := validateModel(ListRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == ""{
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Create Atlas API Request Object
	clusterName := cast.ToString(currentModel.ClusterName)
	projectID := cast.ToString(currentModel.ProjectId)

	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	// API call to get the list of export jobs for the project and cluster
	exportJobs, resp, err := client.CloudProviderSnapshotExportJobs.List(context.Background(), projectID, clusterName, params)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	// create list model to return
	models := make([]interface{}, 0)
	for ind := range exportJobs.Results {
		models = append(models, convertToModel(exportJobs.Results[ind], *currentModel, exportJobs.Links))
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List complete",
		ResourceModels:  models,
	}, nil
}

// function to track snapshot creation status
func validateProgress(client *mongodbatlas.Client, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	clusterName := cast.ToString(currentModel.ClusterName)
	exportJobID := cast.ToString(currentModel.ExportId)
	projectID := cast.ToString(currentModel.ProjectId)
	isReady, state, err := isJobInTargetState(client, projectID, exportJobID, clusterName, targetState)
	if err != nil || state == "Cancelled" {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("error occurred while export job ,Details : %v", err),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// if not ready retry
	if !isReady {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = handler.InProgress
		p.CallbackDelaySeconds = 45
		p.Message = "Pending"
		p.CallbackContext = map[string]interface{}{
			"status":    state,
			"export_id": *currentModel.ExportId,
		}
		return p, nil
	}
	// API call to get export job details
	exportJob, resp, err := client.CloudProviderSnapshotExportJobs.Get(context.Background(), projectID, clusterName, exportJobID)
	if err != nil {
		_, _ = logger.Warnf(ErrorExportJobRead, projectID, exportJobID, err)
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	resultModel := convertToModel(exportJob, *currentModel, resp.Links)

	p := handler.NewProgressEvent()
	p.ResourceModel = resultModel
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	return p, nil
}

func convertToModel(exportJob *mongodbatlas.CloudProviderSnapshotExportJob, currentModel Model, links []*mongodbatlas.Link) Model {
	if exportJob != nil {
		currentModel.ExportId = &exportJob.ID
		currentModel.ExportBucketId = &exportJob.ExportBucketID
		currentModel.CreatedAt = &exportJob.CreatedAt
		currentModel.FinishedAt = &exportJob.FinishedAt
		currentModel.CreatedAt = &exportJob.CreatedAt
		currentModel.Prefix = &exportJob.Prefix
		currentModel.State = &exportJob.State
		currentModel.SnapshotId = &exportJob.SnapshotID
		currentModel.Links = flattenLinks(links)
		currentModel.ExportStatus = flattenStatus(exportJob.ExportStatus)
		currentModel.CustomDataSet = flattenExportJobsCustomData(exportJob.CustomData)
		currentModel.Components = flattenExportComponent(exportJob.Components)
	}
	return currentModel
}

func readExportJob(client *mongodbatlas.Client, projectID, clusterName, exportJobID string) (*mongodbatlas.CloudProviderSnapshotExportJob, error) {
	exportJob, _, err := client.CloudProviderSnapshotExportJobs.Get(context.Background(), projectID, clusterName, exportJobID)
	return exportJob, err
}

// function to check if export job is in target state
func isJobInTargetState(client *mongodbatlas.Client, projectID, exportJobID, clusterName, targetState string) (isReady bool, state string, err error) {
	exportJob, err := readExportJob(client, projectID, clusterName, exportJobID)
	if err != nil {
		return false, "", err
	}
	return exportJob.State == targetState, exportJob.State, nil
}

// function to convert custom metadata from request to mongodb atlas object
func expandExportJobCustomData(currentModel *Model) []*mongodbatlas.CloudProviderSnapshotExportJobCustomData {
	customData := currentModel.CustomDataSet
	result := make([]*mongodbatlas.CloudProviderSnapshotExportJobCustomData, 0)
	for i := range customData {
		res := &mongodbatlas.CloudProviderSnapshotExportJobCustomData{
			Key:   cast.ToString(customData[i].Key),
			Value: cast.ToString(customData[i].Value),
		}
		result = append(result, res)
	}
	return result
}
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
func flattenStatus(v *mongodbatlas.CloudProviderSnapshotExportJobStatus) *ExportStatus {
	if v != nil {
		status := ExportStatus{
			ExportedCollections: &v.ExportedCollections,
			TotalCollections:    &v.TotalCollections,
		}
		return &status
	}
	return nil
}

func flattenExportJobsCustomData(metaData []*mongodbatlas.CloudProviderSnapshotExportJobCustomData) []CustomData {
	statusList := make([]CustomData, 0)
	for i := range metaData {
		role := CustomData{
			Key:   &metaData[i].Key,
			Value: &metaData[i].Value,
		}
		statusList = append(statusList, role)
	}
	return statusList
}
func flattenExportComponent(components []*mongodbatlas.CloudProviderSnapshotExportJobComponent) []ApiAtlasDiskBackupBaseRestoreMemberView {
	componentList := make([]ApiAtlasDiskBackupBaseRestoreMemberView, 0)
	for i := range components {
		role := ApiAtlasDiskBackupBaseRestoreMemberView{
			ReplicaSetName: &components[i].ReplicaSetName,
			ExportID:       &components[i].ExportID,
		}
		componentList = append(componentList, role)
	}
	return componentList
}

func getDeleteStatus(req handler.Request, currentModel *Model) bool {
	return currentModel.TestMode != nil && util.Get(*currentModel.ProjectId, cast.ToString(currentModel.ExportId), req.Session) == "deleted"
}

func handleCFNTestList(req handler.Request, currentModel *Model) (handler.ProgressEvent, error) {
	var models []interface{}
	*currentModel.ExportId = "****Test63918a6ee35d0c1f078c76c8***"
	models = append(models, &currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List complete",
		ResourceModels:  models,
	}, nil
}
