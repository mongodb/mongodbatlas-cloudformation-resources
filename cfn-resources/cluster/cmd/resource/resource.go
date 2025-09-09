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
	"net/http"
	"strings"

	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/spf13/cast"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

const (
	LabelError      = "you should not set `Infrastructure Tool` label, it is used for internal purposes"
	CallBackSeconds = 40
)

var (
	defaultLabel                         = Labels{Key: aws.String("Infrastructure Tool"), Value: aws.String("MongoDB Atlas CloudFormation Provider")}
	createReadUpdareDeleteRequiredFields = []string{constants.ProjectID, constants.Name}
	listRequiredFields                   = []string{constants.ProjectID}
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, _ *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, currentModel, createReadUpdareDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		return clusterCallback(client, currentModel, *currentModel.ProjectId)
	}
	currentModel.validateDefaultLabel()
	clusterRequest, errEvent := setClusterRequest(currentModel)
	if errEvent != nil {
		return *errEvent, nil
	}

	var err error
	cluster, res, err := client.Atlas20231115014.ClustersApi.CreateCluster(context.Background(), *currentModel.ProjectId, clusterRequest).Execute()
	if err != nil {
		if apiError, ok := admin20231115014.AsError(err); ok && *apiError.Error == http.StatusBadRequest && strings.Contains(*apiError.ErrorCode, constants.Duplicate) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
		}

		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	currentModel.StateName = cluster.StateName

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create Cluster `%s`", *cluster.StateName),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.StateName: cluster.StateName,
		},
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, currentModel, createReadUpdareDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	model, resp, err := readCluster(context.Background(), client, currentModel)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}

		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   model,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, currentModel, createReadUpdareDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	if _, ok := req.CallbackContext[constants.StateName]; ok {
		return updateClusterCallback(client, currentModel, *currentModel.ProjectId)
	}
	currentModel.validateDefaultLabel()
	adminCluster, errEvent := setClusterRequest(currentModel)
	if len(adminCluster.GetReplicationSpecs()) > 0 {
		currentCluster, _, _ := client.Atlas20231115014.ClustersApi.GetCluster(context.Background(), *currentModel.ProjectId, *currentModel.Name).Execute()
		if currentCluster != nil {
			adminCluster.ReplicationSpecs = AddReplicationSpecIDs(currentCluster.GetReplicationSpecs(), adminCluster.GetReplicationSpecs())
		}
	}
	if errEvent != nil {
		return *errEvent, nil
	}

	model, _, err := updateCluster(context.Background(), client, currentModel, adminCluster)
	if err != nil {
		if apiError, ok := admin20231115014.AsError(err); ok && *apiError.Error == http.StatusNotFound {
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}

		code := cloudformation.HandlerErrorCodeServiceInternalError
		if strings.Contains(err.Error(), "not exist") { // cfn test needs 404
			code = cloudformation.HandlerErrorCodeNotFound
		}
		if strings.Contains(err.Error(), "being deleted") {
			code = cloudformation.HandlerErrorCodeNotFound // cfn test needs 404
		}
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: code}, nil
	}

	var state string
	if model.StateName != nil {
		state = *model.StateName
	}
	event := handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Update Cluster %s", state),
		ResourceModel:        model,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.StateName: state,
		},
	}
	return event, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, currentModel, createReadUpdareDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	if _, ok := req.CallbackContext[constants.StateName]; ok {
		return validateProgress(client, currentModel, constants.DeletedState)
	}

	params := &admin20231115014.DeleteClusterApiParams{
		RetainBackups: util.Pointer(false),
		GroupId:       *currentModel.ProjectId,
		ClusterName:   *currentModel.Name,
	}

	_, err := client.Atlas20231115014.ClustersApi.DeleteClusterWithParams(context.Background(), params).Execute()
	if err != nil {
		if apiError, ok := admin20231115014.AsError(err); ok && *apiError.Error == http.StatusNotFound {
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}

		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              constants.DeleteInProgress,
		ResourceModel:        currentModel,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.StateName: constants.DeletingState,
		}}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, currentModel, listRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}

	listOptions := &admin20231115014.ListClustersApiParams{
		ItemsPerPage: admin20231115014.PtrInt(100),
		PageNum:      admin20231115014.PtrInt(1),
		GroupId:      *currentModel.ProjectId,
		IncludeCount: admin20231115014.PtrBool(true),
	}
	// List call
	clustersResponse, res, err := client.Atlas20231115014.ClustersApi.ListClustersWithParams(context.Background(), listOptions).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
			res), nil
	}
	models := make([]*Model, *clustersResponse.TotalCount)
	clusterResults := clustersResponse.GetResults()
	for i := range clusterResults {
		model := &Model{}
		mapClusterToModel(model, &clusterResults[i])
		// Call AdvancedSettings
		processArgs, res, err := client.Atlas20231115014.ClustersApi.GetClusterAdvancedConfiguration(context.Background(), *model.ProjectId, *model.Name).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
				res), nil
		}
		model.AdvancedSettings = flattenProcessArgs(processArgs, &clusterResults[i])
		models[i] = model
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List",
		ResourceModel:   models}, nil
}

func clusterCallback(client *util.MongoDBClient, currentModel *Model, projectID string) (handler.ProgressEvent, error) {
	progressEvent, err := validateProgress(client, currentModel, constants.IdleState)
	if err != nil {
		return progressEvent, nil
	}

	if progressEvent.Message == constants.Complete {
		if !currentModel.HasAdvanceSettings() {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Create Success",
				ResourceModel:   currentModel}, nil
		}

		cluster, res, err := client.Atlas20231115014.ClustersApi.GetCluster(context.Background(), projectID, *currentModel.Name).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
				res), nil
		}

		return updateClusterSettings(currentModel, client, projectID, cluster, &progressEvent)
	}
	return progressEvent, nil
}

func (m *Model) HasAdvanceSettings() bool {
	/*This logic is because of a bug un Cloud Formation, when we return in_progress in the CREATE
	,the second time the CREATE gets executed
	it returns the AdvancedSettings is not nil but its fields are nil*/
	return m.AdvancedSettings != nil && (m.AdvancedSettings.DefaultReadConcern != nil ||
		m.AdvancedSettings.DefaultWriteConcern != nil ||
		m.AdvancedSettings.FailIndexKeyTooLong != nil ||
		m.AdvancedSettings.JavascriptEnabled != nil ||
		m.AdvancedSettings.MinimumEnabledTLSProtocol != nil ||
		m.AdvancedSettings.NoTableScan != nil ||
		m.AdvancedSettings.OplogSizeMB != nil ||
		m.AdvancedSettings.SampleSizeBIConnector != nil ||
		m.AdvancedSettings.SampleRefreshIntervalBIConnector != nil ||
		m.AdvancedSettings.OplogMinRetentionHours != nil)
}

func formatMongoDBMajorVersion(val interface{}) string {
	if strings.Contains(val.(string), ".") {
		return val.(string)
	}
	return fmt.Sprintf("%.1f", cast.ToFloat32(val))
}

func isClusterInTargetState(client *util.MongoDBClient, projectID, clusterName, targetState string) (isReady bool, stateName string, mongoCluster *admin20231115014.AdvancedClusterDescription, err error) {
	cluster, resp, err := client.Atlas20231115014.ClustersApi.GetCluster(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return constants.DeletedState == targetState, constants.DeletedState, nil, nil
		}
		return false, constants.Error, nil, fmt.Errorf("error fetching cluster info (%s): %s", clusterName, err)
	}
	return *cluster.StateName == targetState, *cluster.StateName, cluster, nil
}

func readCluster(ctx context.Context, client *util.MongoDBClient, currentModel *Model) (*Model, *http.Response, error) {
	cluster, res, err := client.Atlas20231115014.ClustersApi.GetCluster(ctx, *currentModel.ProjectId, *currentModel.Name).Execute()
	if err != nil || res.StatusCode != http.StatusOK {
		return currentModel, res, err
	}

	setClusterData(currentModel, cluster)

	if currentModel.AdvancedSettings != nil {
		processArgs, resp, errr := client.Atlas20231115014.ClustersApi.GetClusterAdvancedConfiguration(ctx, *currentModel.ProjectId, *currentModel.Name).Execute()
		if errr != nil || resp.StatusCode != http.StatusOK {
			return currentModel, resp, errr
		}
		currentModel.AdvancedSettings = flattenProcessArgs(processArgs, cluster)
	}
	return currentModel, res, err
}

func updateCluster(ctx context.Context, client *util.MongoDBClient, currentModel *Model, clusterRequest *admin20231115014.AdvancedClusterDescription) (*Model, *http.Response, error) {
	cluster, resp, err := client.Atlas20231115014.ClustersApi.UpdateCluster(ctx, *currentModel.ProjectId, *currentModel.Name, clusterRequest).Execute()

	if cluster != nil {
		currentModel.StateName = cluster.StateName
	}

	return currentModel, resp, err
}

func updateAdvancedCluster(ctx context.Context, client *util.MongoDBClient,
	request *admin20231115014.AdvancedClusterDescription, projectID, name string) (*admin20231115014.AdvancedClusterDescription, *http.Response, error) {
	return client.Atlas20231115014.ClustersApi.UpdateCluster(ctx, projectID, name, request).Execute()
}

func updateClusterCallback(client *util.MongoDBClient, currentModel *Model, projectID string) (handler.ProgressEvent, error) {
	progressEvent, err := validateProgress(client, currentModel, constants.IdleState)
	if err != nil {
		return progressEvent, nil
	}

	if progressEvent.Message == constants.Complete {
		cluster, res, err := client.Atlas20231115014.ClustersApi.GetCluster(context.Background(), projectID, *currentModel.Name).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error in Get Cluster : %s", err.Error()),
				res), nil
		}
		return updateClusterSettings(currentModel, client, projectID, cluster, &progressEvent)
	}
	return progressEvent, nil
}

func updateClusterSettings(currentModel *Model, client *util.MongoDBClient,
	projectID string, cluster *admin20231115014.AdvancedClusterDescription, pe *handler.ProgressEvent) (handler.ProgressEvent, error) {
	// Update advanced configuration
	if currentModel.AdvancedSettings != nil {
		advancedConfig := expandAdvancedSettings(*currentModel.AdvancedSettings)
		_, res, err := client.Atlas20231115014.ClustersApi.UpdateClusterAdvancedConfiguration(context.Background(), projectID, *cluster.Name, advancedConfig).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
				res), err
		}
	}

	// Update pause
	if (currentModel.Paused != nil) && (*currentModel.Paused != *cluster.Paused) {
		_, res, err := updateAdvancedCluster(context.Background(), client, &admin20231115014.AdvancedClusterDescription{Paused: currentModel.Paused}, projectID, *currentModel.Name)
		if err != nil {
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Cluster Pause error : %s", err.Error()),
				res), err
		}
	}
	return *pe, nil
}

func validateProgress(client *util.MongoDBClient, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	isReady, state, cluster, err := isClusterInTargetState(client, *currentModel.ProjectId, *currentModel.Name, targetState)
	if err != nil {
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}

	if !isReady {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = handler.InProgress
		p.CallbackDelaySeconds = CallBackSeconds
		p.Message = constants.Pending
		p.CallbackContext = map[string]interface{}{
			constants.StateName: state,
		}
		return p, nil
	}

	p := handler.NewProgressEvent()
	p.OperationStatus = handler.Success
	p.Message = constants.Complete
	// Delete event shouldn't have model in the response
	if targetState == constants.IdleState {
		currentModel.StateName = cluster.StateName
		currentModel.ConnectionStrings = flattenConnectionStrings(cluster.ConnectionStrings)
		p.ResourceModel = currentModel
	}

	return p, nil
}

func (m *Model) validateDefaultLabel() {
	if !containsLabelOrKey(m.Labels, defaultLabel) {
		m.Labels = append(m.Labels, defaultLabel)
	}
}

func setupRequest(req handler.Request, model *Model, requiredFields []string) (*util.MongoDBClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-cluster")
	if modelValidation := validator.ValidateModel(requiredFields, model); modelValidation != nil {
		return nil, modelValidation
	}
	util.SetDefaultProfileIfNotDefined(&model.Profile)
	client, peErr := util.NewAtlasClient(&req, model.Profile)
	if peErr != nil {
		return nil, peErr
	}
	return client, nil
}
