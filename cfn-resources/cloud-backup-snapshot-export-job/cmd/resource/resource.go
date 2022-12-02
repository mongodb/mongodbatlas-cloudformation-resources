package resource

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"

	"github.com/spf13/cast"

	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/openlyinc/pointy"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.GroupID, constants.ExportBucketID, constants.SnapshotID}
var ReadRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.GroupID, constants.ExportID, constants.ClusterName}
var ListRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.GroupID, constants.UserName}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-CloudBackupSnapshotExportJob")
}

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
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// create API request object
	projectID := *currentModel.GroupId
	clusterName := *currentModel.ClusterName
	request := &mongodbatlas.CloudProviderSnapshotExportJob{
		SnapshotID:     *currentModel.SnapshotId,
		ExportBucketID: *currentModel.ExportBucketId,
		CustomData:     expandExportJobCustomData(currentModel),
	}

	// progress callback setup
	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["export_id"].(string)
		currentModel.ExportId = &sid
		return validateProgress(client, currentModel, "Successful")
	}
	// API call to create export job
	jobResponse, _, err := client.CloudProviderSnapshotExportJobs.Create(context.Background(), projectID, clusterName, request)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorExportJobCreate, projectID, err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	currentModel.ExportId = &jobResponse.ID
	_, _ = logger.Debugf("created successfully  %v", &jobResponse.State)

	// track progress
	event := handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create export snapshots : %s", jobResponse.ID),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 25,
		CallbackContext: map[string]interface{}{
			"status":    jobResponse.State,
			"export_id": jobResponse.ID,
		},
	}
	return event, nil
}

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
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	projectID := *currentModel.GroupId
	clusterName := *currentModel.ClusterName
	exportJobID := *currentModel.ExportId

	if !isExist(client, projectID, clusterName, clusterName) {
		_, _ = logger.Warnf(constants.ErrorExportJobRead, projectID, exportJobID, errors.New("resource Not Found"))
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	// API call to read export job
	exportJob, resp, err := client.CloudProviderSnapshotExportJobs.Get(context.Background(), projectID, clusterName, exportJobID)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorExportJobRead, projectID, exportJobID, err)
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}
	readModel := convertToModel(exportJob, resp.Links)
	_, _ = logger.Debugf("Read Result : %v", currentModel)

	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   readModel,
	}
	return event, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// NO OP

	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	_, _ = logger.Debugf("Delete() currentModel:%+v", currentModel)
	// NO OP

	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("List snapshot for Request() currentModel:%+v", currentModel)

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
	projectID := *currentModel.GroupId
	clusterName := *currentModel.ClusterName

	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}
	// API call to get the list of export jobes for the project and cluster
	exportJobs, _, err := client.CloudProviderSnapshotExportJobs.List(context.Background(), projectID, clusterName, params)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading cloud provider snapshot restore job list with id(project: %s): %s", projectID, err)
	}
	// create list model to return
	var models []Model
	for _, exportJob := range exportJobs.Results {
		models = append(models, convertToModel(exportJob, exportJobs.Links))
	}
	_, _ = logger.Debug("List cloud backup export job ends")

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List complete",
		ResourceModel:   models,
	}, nil
}

// function to track snapshot creation status
func validateProgress(client *mongodbatlas.Client, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	fmt.Printf("mmmmmmmmmmmmmmbbbbbbbbbbbbbbbbbbbbbbbbbbbbb%+v", client)
	fmt.Printf("mmmmmmmmmmmmmmbbbbbbbbbbbbbbbbbbbbbcurrentModelbbbbbbbb%+v", currentModel)
	fmt.Printf("mmmmmmmmmmmmmmbbbbbbbbbbbbbtargetStatetargetStatebbbbbbbbcurrentModelbbbbbbbb%+v", targetState)

	exportID := *currentModel.ExportId
	projectID := *currentModel.GroupId
	clusterName := *currentModel.ClusterName
	isReady, state, err := isJobInTargetState(client, projectID, exportID, clusterName, targetState)
	if err != nil || state == "Cancelled" {
		return handler.ProgressEvent{}, err
	}
	// if not ready retry
	if !isReady {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = handler.InProgress
		p.CallbackDelaySeconds = 20
		p.Message = "Pending"
		p.CallbackContext = map[string]interface{}{
			"status":    state,
			"export_id": *currentModel.ExportId,
		}
		return p, nil
	}
	fmt.Printf("kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk ")

	// API call to get export job details
	exportJob, resp, err := client.CloudProviderSnapshotExportJobs.Get(context.Background(), projectID, clusterName, exportID)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorExportJobRead, projectID, exportID, err)
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}
	fmt.Printf("llllllllllllllllllllllllllllllllllllllll ")
	resultModel := convertToModel(exportJob, resp.Links)
	fmt.Printf("mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm ")
	p := handler.NewProgressEvent()
	p.ResourceModel = resultModel
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	return p, nil
}

func convertToModel(exportJob *mongodbatlas.CloudProviderSnapshotExportJob, links []*mongodbatlas.Link) Model {
	var result Model
	if exportJob != nil {
		result.ExportId = &exportJob.ID
		result.ExportBucketId = &exportJob.ExportBucketID
		result.CreatedAt = &exportJob.CreatedAt
		result.FinishedAt = &exportJob.FinishedAt
		result.CreatedAt = &exportJob.CreatedAt
		result.Prefix = &exportJob.Prefix
		result.State = &exportJob.State
		result.SnapshotId = &exportJob.SnapshotID
		result.Links = flattenLinks(links)
		result.ExportStatus = flattenStatus(exportJob.ExportStatus)
		result.CustomDataSet = flattenExportJobsCustomData(exportJob.CustomData)
		result.Components = flattenExportComponent(exportJob.Components)
	}
	return result
}

func readExportJob(client *mongodbatlas.Client, projectID, clusterName, exportJobID string) (*mongodbatlas.CloudProviderSnapshotExportJob, error) {
	exportJob, _, err := client.CloudProviderSnapshotExportJobs.Get(context.Background(), projectID, clusterName, exportJobID)
	return exportJob, err
}

// function to check if export job is in target state
func isJobInTargetState(client *mongodbatlas.Client, projectID, exportJobID, clusterName, targetState string) (isReady bool, state string, err error) {
	fmt.Printf("ssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss")
	exportJob, err := readExportJob(client, projectID, clusterName, exportJobID)
	fmt.Printf("mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm %+v", exportJob)

	if err != nil {
		return false, "", err
	}
	return exportJob.State == targetState, exportJob.State, nil
}

// function to check if snapshot already exist in atlas
func isExist(client *mongodbatlas.Client, projectID, exportJobID, clusterName string) bool {
	exportJob, err := readExportJob(client, projectID, clusterName, exportJobID)
	if err != nil {
		return false
	} else if exportJob == nil {
		return false
	}
	return true
}

// function to convert custom metadata from request to mongodb atlas object
func expandExportJobCustomData(currentModel *Model) []*mongodbatlas.CloudProviderSnapshotExportJobCustomData {
	customData := currentModel.CustomDataSet
	if customData != nil {
		res := make([]*mongodbatlas.CloudProviderSnapshotExportJobCustomData, len(customData))
		for i, val := range customData {
			res[i] = &mongodbatlas.CloudProviderSnapshotExportJobCustomData{
				Key:   cast.ToString(val.Key),
				Value: cast.ToString(val.Value),
			}
		}
		return res
	}
	return nil
}
func flattenLinks(linksResult []*mongodbatlas.Link) []Link {
	if len(linksResult) == 0 {
		return nil
	}
	links := make([]Link, 0)
	for _, link := range linksResult {
		var lin Link
		lin.Href = &link.Href
		lin.Rel = &link.Rel
		links = append(links, lin)
	}
	return links
}
func flattenStatus(v *mongodbatlas.CloudProviderSnapshotExportJobStatus) *ApiExportStatusView {
	if v != nil {
		status := ApiExportStatusView{
			ExportedCollections: &v.ExportedCollections,
			TotalCollections:    &v.TotalCollections,
		}
		return &status
	}
	return nil
}

func flattenExportJobsCustomData(metaData []*mongodbatlas.CloudProviderSnapshotExportJobCustomData) []CustomData {
	if len(metaData) == 0 {
		return nil
	}
	statusList := make([]CustomData, len(metaData))
	for i := range metaData {
		v := metaData[i]
		role := CustomData{
			Key:   &v.Key,
			Value: &v.Value,
		}
		statusList = append(statusList, role)
	}
	return statusList
}
func flattenExportComponent(components []*mongodbatlas.CloudProviderSnapshotExportJobComponent) []ApiAtlasDiskBackupBaseRestoreMemberView {
	if len(components) == 0 {
		return nil
	}
	componentList := make([]ApiAtlasDiskBackupBaseRestoreMemberView, len(components))
	for i := range components {
		v := components[i]
		role := ApiAtlasDiskBackupBaseRestoreMemberView{
			ReplicaSetName: pointy.String(v.ReplicaSetName),
			ExportID:       pointy.String(v.ExportID),
		}

		componentList = append(componentList, role)
	}
	return componentList
}
