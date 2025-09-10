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

	flex "github.com/mongodb/mongodbatlas-cloudformation-resources/flex-cluster/cmd/resource"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/spf13/cast"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

const (
	callBackSeconds = 40
)

var (
	defaultLabel                         = Labels{Key: aws.String("Infrastructure Tool"), Value: aws.String("MongoDB Atlas CloudFormation Provider")}
	createReadUpdareDeleteRequiredFields = []string{constants.ProjectID, constants.Name}
	listRequiredFields                   = []string{constants.ProjectID}
	callbackContext                      = map[string]any{"callbackCluster": true}
)

func isCallback(req *handler.Request) bool {
	_, found := req.CallbackContext["callbackCluster"]
	return found
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, _ *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, currentModel, createReadUpdareDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	if flexModel := clusterToFlexModelFull(currentModel); flexModel != nil {
		pe := flex.HandleCreate(&req, client, flexModel)
		fillModelForFlex(&pe, currentModel)
		return pe, nil
	}
	if isCallback(&req) {
		return clusterCallback(client, currentModel, *currentModel.ProjectId)
	}
	currentModel.validateDefaultLabel()
	clusterRequest, errEvent := setClusterRequest(currentModel)
	if errEvent != nil {
		return *errEvent, nil
	}
	cluster, resp, err := client.Atlas20231115014.ClustersApi.CreateCluster(context.Background(), *currentModel.ProjectId, clusterRequest).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe, nil
	}
	currentModel.StateName = cluster.StateName
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create Cluster `%s`", *cluster.StateName),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: callBackSeconds,
		CallbackContext:      callbackContext,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, currentModel, createReadUpdareDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	if flexModel := clusterToFlexModelIdentifier(&req, client, currentModel); flexModel != nil {
		pe := flex.HandleRead(&req, client, flexModel)
		fillModelForFlex(&pe, currentModel)
		return pe, nil
	}
	model, resp, err := readCluster(context.Background(), client, currentModel)
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe, nil
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
	if flexModel := clusterToFlexModelFull(currentModel); flexModel != nil {
		pe := flex.HandleUpdate(&req, client, flexModel)
		fillModelForFlex(&pe, currentModel)
		return pe, nil
	}
	if isCallback(&req) {
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

	model, resp, err := updateCluster(context.Background(), client, currentModel, adminCluster)
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe, nil
	}

	var state string
	if model.StateName != nil {
		state = *model.StateName
	}
	event := handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Update Cluster %s", state),
		ResourceModel:        model,
		CallbackDelaySeconds: callBackSeconds,
		CallbackContext:      callbackContext,
	}
	return event, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, currentModel, createReadUpdareDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	if flexModel := clusterToFlexModelIdentifier(&req, client, currentModel); flexModel != nil {
		pe := flex.HandleDelete(&req, client, flexModel)
		fillModelForFlex(&pe, currentModel)
		return pe, nil
	}
	if isCallback(&req) {
		return validateProgress(client, currentModel, constants.DeletedState)
	}
	params := &admin20231115014.DeleteClusterApiParams{
		RetainBackups: util.Pointer(false),
		GroupId:       *currentModel.ProjectId,
		ClusterName:   *currentModel.Name,
	}
	resp, err := client.Atlas20231115014.ClustersApi.DeleteClusterWithParams(context.Background(), params).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe, nil
	}
	return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              constants.DeleteInProgress,
			ResourceModel:        currentModel,
			CallbackDelaySeconds: callBackSeconds,
			CallbackContext:      callbackContext,
		},
		nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, currentModel, listRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}

	var models []*Model
	const itemsPerPage = 100
	for pageNum := 1; ; pageNum++ {
		listOptions := &admin20231115014.ListClustersApiParams{
			ItemsPerPage: admin20231115014.PtrInt(itemsPerPage),
			PageNum:      admin20231115014.PtrInt(pageNum),
			GroupId:      *currentModel.ProjectId,
			IncludeCount: admin20231115014.PtrBool(true),
		}

		clustersResponse, resp, err := client.Atlas20231115014.ClustersApi.ListClustersWithParams(context.Background(), listOptions).Execute()
		if pe := util.HandleClusterError(err, resp); pe != nil {
			return *pe, nil
		}
		clusterResults := clustersResponse.GetResults()
		for i := range clusterResults {
			model := &Model{}
			mapClusterToModel(model, &clusterResults[i])

			processArgs, resp, err := client.Atlas20231115014.ClustersApi.GetClusterAdvancedConfiguration(context.Background(), *model.ProjectId, *model.Name).Execute()
			if pe := util.HandleClusterError(err, resp); pe != nil {
				return *pe, nil
			}
			model.AdvancedSettings = flattenProcessArgs(processArgs, &clusterResults[i])
			models = append(models, model)
		}

		// Check if we've retrieved all clusters or if the current page has fewer items than requested
		if len(models) >= clustersResponse.GetTotalCount() || len(clusterResults) < itemsPerPage {
			break
		}
	}

	// Get flex clusters and append them to the models list
	flexModel := &flex.Model{ProjectId: currentModel.ProjectId}
	pe := flex.HandleList(&req, client, flexModel)
	models = fillModelForFlexList(&pe, models)

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
		cluster, resp, err := client.Atlas20231115014.ClustersApi.GetCluster(context.Background(), projectID, *currentModel.Name).Execute()
		if pe := util.HandleClusterError(err, resp); pe != nil {
			return *pe, nil
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

func isClusterInTargetState(client *util.MongoDBClient, projectID, clusterName, targetState string) (isReady bool, mongoCluster *admin20231115014.AdvancedClusterDescription, err error) {
	cluster, resp, err := client.Atlas20231115014.ClustersApi.GetCluster(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return constants.DeletedState == targetState, nil, nil
		}
		return false, nil, fmt.Errorf("error fetching cluster info (%s): %s", clusterName, err)
	}
	return *cluster.StateName == targetState, cluster, nil
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
		cluster, resp, err := client.Atlas20231115014.ClustersApi.GetCluster(context.Background(), projectID, *currentModel.Name).Execute()
		if pe := util.HandleClusterError(err, resp); pe != nil {
			return *pe, nil
		}
		return updateClusterSettings(currentModel, client, projectID, cluster, &progressEvent)
	}
	return progressEvent, nil
}

func updateClusterSettings(currentModel *Model, client *util.MongoDBClient,
	projectID string, cluster *admin20231115014.AdvancedClusterDescription, pe *handler.ProgressEvent) (handler.ProgressEvent, error) {
	if currentModel.AdvancedSettings != nil {
		advancedConfig := expandAdvancedSettings(*currentModel.AdvancedSettings)
		_, resp, err := client.Atlas20231115014.ClustersApi.UpdateClusterAdvancedConfiguration(context.Background(), projectID, *cluster.Name, advancedConfig).Execute()
		if pe := util.HandleClusterError(err, resp); pe != nil {
			return *pe, nil
		}
	}

	if (currentModel.Paused != nil) && (*currentModel.Paused != *cluster.Paused) {
		_, resp, err := updateAdvancedCluster(context.Background(), client, &admin20231115014.AdvancedClusterDescription{Paused: currentModel.Paused}, projectID, *currentModel.Name)
		if pe := util.HandleClusterError(err, resp); pe != nil {
			return *pe, nil
		}
	}
	return *pe, nil
}

func validateProgress(client *util.MongoDBClient, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	isReady, cluster, err := isClusterInTargetState(client, *currentModel.ProjectId, *currentModel.Name, targetState)
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
		p.CallbackDelaySeconds = callBackSeconds
		p.Message = constants.Pending
		p.CallbackContext = callbackContext
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
