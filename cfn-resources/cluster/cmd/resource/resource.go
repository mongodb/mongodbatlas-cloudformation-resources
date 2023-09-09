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
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	LabelError      = "you should not set `Infrastructure Tool` label, it is used for internal purposes"
	CallBackSeconds = 40
)

var defaultLabel = Labels{Key: aws.String("Infrastructure Tool"), Value: aws.String("MongoDB Atlas CloudFormation Provider")}

var CreateRequiredFields = []string{constants.ProjectID, constants.Name}
var UpdateRequiredFields = []string{constants.ProjectID, constants.Name}
var DeleteRequiredFields = []string{constants.ProjectID, constants.Name}
var ListRequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-cluster")
}

func castNO64(i *int64) *int {
	x := cast.ToInt(&i)
	return &x
}

func cast64(i *int) *int64 {
	x := cast.ToInt64(&i)
	return &x
}

// validateModel inputs based on the method
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, _ *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Callback
	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		return clusterCallback(client, currentModel, *currentModel.ProjectId)
	}

	var err error

	currentModel.validateDefaultLabel()

	// Prepare cluster request
	clusterRequest, event, err := setClusterRequest(currentModel)
	if err != nil {
		return event, nil
	}

	// Create Cluster
	cluster, res, err := client.AdvancedClusters.Create(context.Background(), *currentModel.ProjectId, clusterRequest)
	if err != nil {
		_, _ = log.Warnf("Create - Cluster.Create() - error: %+v", err)
		if res.Response.StatusCode == 400 && strings.Contains(err.Error(), constants.Duplicate) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
		}
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.StateName = &cluster.StateName

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create Cluster `%s`", cluster.StateName),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.StateName: cluster.StateName,
		},
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	_, _ = log.Debugf("Read() currentModel:%+v", currentModel)

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Read call
	model, resp, err := readCluster(context.Background(), client, currentModel)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = log.Warnf("error cluster get- err:%+v resp:%+v", err, resp)
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
	setup()
	_, _ = log.Debugf("Update() currentModel:%+v", currentModel)

	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Update callback
	if _, ok := req.CallbackContext[constants.StateName]; ok {
		return updateClusterCallback(client, currentModel, *currentModel.ProjectId)
	}

	currentModel.validateDefaultLabel()

	// Update Cluster
	model, resp, err := updateCluster(context.Background(), client, currentModel)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			_, _ = log.Warnf("update 404 err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}

		_, _ = log.Warnf("update err: %+v", err)
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
	_, _ = log.Debugf("state: %+v", state)
	event := handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Update Cluster %s", state),
		ResourceModel:        model,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.StateName: state,
		},
	}
	_, _ = log.Debugf("Update() return event:%+v", event)
	return event, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	_, _ = log.Debugf("Delete() currentModel:%+v", currentModel)

	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	ctx := context.Background()

	if _, ok := req.CallbackContext[constants.StateName]; ok {
		return validateProgress(client, currentModel, constants.DeletingState, constants.DeletedState)
	}

	options := &mongodbatlas.DeleteAdvanceClusterOptions{RetainBackups: util.Pointer(false)}
	resp, err := client.AdvancedClusters.Delete(ctx, *currentModel.ProjectId, *currentModel.Name, options)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			_, _ = log.Warnf("Delete 404 err: %+v", err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}

		_, _ = log.Warnf("Delete err: %+v", err)
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
	setup()
	_, _ = log.Debugf("List() currentModel:%+v", currentModel)

	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	listOptions := &mongodbatlas.ListOptions{ItemsPerPage: 100, PageNum: 1}
	// List call
	clustersResponse, res, err := client.AdvancedClusters.List(context.Background(), *currentModel.ProjectId, listOptions)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
			res.Response), nil
	}
	models := make([]*Model, clustersResponse.TotalCount)
	for i := range clustersResponse.Results {
		model := &Model{}
		mapClusterToModel(model, clustersResponse.Results[i])
		// Call AdvancedSettings
		processArgs, res, err := client.Clusters.GetProcessArgs(context.Background(), *model.ProjectId, *model.Name)
		if err != nil {
			return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
				res.Response), nil
		}
		model.AdvancedSettings = flattenProcessArgs(processArgs)
		models = append(models, model)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List",
		ResourceModel:   models}, nil
}

func clusterCallback(client *mongodbatlas.Client, currentModel *Model, projectID string) (handler.ProgressEvent, error) {
	progressEvent, err := validateProgress(client, currentModel, constants.CreatingState, constants.IdleState)
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

		_, _ = log.Debugf("Cluster Creation completed:%s", *currentModel.Name)

		cluster, res, err := client.AdvancedClusters.Get(context.Background(), projectID, *currentModel.Name)
		if err != nil {
			return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
				res.Response), nil
		}

		_, _ = log.Debugf("Updating cluster settings:%s", *currentModel.Name)
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

func isClusterInTargetState(client *mongodbatlas.Client, projectID, clusterName, targetState string) (isReady bool, stateName string, mongoCluster *mongodbatlas.AdvancedCluster, err error) {
	cluster, resp, err := client.AdvancedClusters.Get(context.Background(), projectID, clusterName)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return constants.DeletedState == targetState, constants.DeletedState, nil, nil
		}
		return false, constants.Error, nil, fmt.Errorf("error fetching cluster info (%s): %s", clusterName, err)
	}
	_, _ = log.Debugf("Cluster state: %s, targetState : %s", cluster.StateName, targetState)
	return cluster.StateName == targetState, cluster.StateName, cluster, nil
}

func readCluster(ctx context.Context, client *mongodbatlas.Client, currentModel *Model) (*Model, *mongodbatlas.Response, error) {
	cluster, res, err := client.AdvancedClusters.Get(ctx, *currentModel.ProjectId, *currentModel.Name)

	if err != nil || res.StatusCode != http.StatusOK {
		return currentModel, res, err
	}
	setClusterData(currentModel, cluster)

	if currentModel.AdvancedSettings != nil {
		processArgs, resp, errr := client.Clusters.GetProcessArgs(ctx, *currentModel.ProjectId, *currentModel.Name)
		if errr != nil || resp.StatusCode != http.StatusOK {
			return currentModel, resp, errr
		}
		currentModel.AdvancedSettings = flattenProcessArgs(processArgs)
	}
	return currentModel, res, err
}

func updateCluster(ctx context.Context, client *mongodbatlas.Client, currentModel *Model) (*Model, *mongodbatlas.Response, error) {
	clusterRequest := &mongodbatlas.AdvancedCluster{}
	if currentModel.BackupEnabled != nil {
		clusterRequest.BackupEnabled = currentModel.BackupEnabled
	}

	if currentModel.BiConnector != nil {
		clusterRequest.BiConnector = expandBiConnector(currentModel.BiConnector)
	}

	if currentModel.ClusterType != nil {
		clusterRequest.ClusterType = *currentModel.ClusterType
	}

	if currentModel.DiskSizeGB != nil {
		clusterRequest.DiskSizeGB = currentModel.DiskSizeGB
	}

	if currentModel.EncryptionAtRestProvider != nil {
		clusterRequest.EncryptionAtRestProvider = *currentModel.EncryptionAtRestProvider
	}

	if len(currentModel.Labels) > 0 {
		clusterRequest.Labels = expandLabelSlice(currentModel.Labels)
	}

	if currentModel.MongoDBMajorVersion != nil {
		clusterRequest.MongoDBMajorVersion = formatMongoDBMajorVersion(*currentModel.MongoDBMajorVersion)
	}

	if currentModel.PitEnabled != nil {
		clusterRequest.PitEnabled = currentModel.PitEnabled
	}

	if currentModel.ReplicationSpecs != nil {
		clusterRequest.ReplicationSpecs = expandReplicationSpecs(currentModel.ReplicationSpecs)
	}

	if currentModel.RootCertType != nil {
		clusterRequest.RootCertType = *currentModel.RootCertType
	}

	if currentModel.VersionReleaseSystem != nil {
		clusterRequest.VersionReleaseSystem = *currentModel.VersionReleaseSystem
	}

	clusterRequest.TerminationProtectionEnabled = currentModel.TerminationProtectionEnabled

	_, _ = log.Debugf("params : %+v %+v %+v", ctx, client, clusterRequest)
	cluster, resp, err := client.AdvancedClusters.Update(ctx, *currentModel.ProjectId, *currentModel.Name, clusterRequest)

	if cluster != nil {
		currentModel.StateName = &cluster.StateName
	}

	return currentModel, resp, err
}

func updateAdvancedCluster(ctx context.Context, conn *mongodbatlas.Client,
	request *mongodbatlas.AdvancedCluster, projectID, name string) (*mongodbatlas.AdvancedCluster, *mongodbatlas.Response, error) {
	return conn.AdvancedClusters.Update(ctx, projectID, name, request)
}

func updateClusterCallback(client *mongodbatlas.Client, currentModel *Model, projectID string) (handler.ProgressEvent, error) {
	progressEvent, err := validateProgress(client, currentModel, constants.UpdateState, constants.IdleState)
	if err != nil {
		return progressEvent, nil
	}

	if progressEvent.Message == constants.Complete {
		_, _ = log.Debugf("compelted updation:%s", *currentModel.Name)
		cluster, res, err := client.AdvancedClusters.Get(context.Background(), projectID, *currentModel.Name)
		if err != nil {
			return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error in Get Cluster : %s", err.Error()),
				res.Response), nil
		}

		_, _ = log.Debugf("Updating cluster :%s", *currentModel.Name)

		return updateClusterSettings(currentModel, client, projectID, cluster, &progressEvent)
	}
	return progressEvent, nil
}

func updateClusterSettings(currentModel *Model, client *mongodbatlas.Client,
	projectID string, cluster *mongodbatlas.AdvancedCluster, pe *handler.ProgressEvent) (handler.ProgressEvent, error) {
	// Update advanced configuration
	if currentModel.AdvancedSettings != nil {
		_, _ = log.Debugf("AdvancedSettings: %+v", *currentModel.AdvancedSettings)

		advancedConfig := expandAdvancedSettings(*currentModel.AdvancedSettings)
		_, res, err := client.Clusters.UpdateProcessArgs(context.Background(), projectID, cluster.Name, advancedConfig)
		if err != nil {
			_, _ = log.Warnf("Cluster UpdateProcessArgs - error: %+v", err)
			return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
				res.Response), err
		}
	}

	// Update pause
	if (currentModel.Paused != nil) && (*currentModel.Paused != *cluster.Paused) {
		_, res, err := updateAdvancedCluster(context.Background(), client, &mongodbatlas.AdvancedCluster{Paused: currentModel.Paused}, projectID, *currentModel.Name)
		if err != nil {
			_, _ = log.Warnf("Cluster Pause - error: %+v", err)
			return progress_events.GetFailedEventByResponse(fmt.Sprintf("Cluster Pause error : %s", err.Error()),
				res.Response), err
		}
	}

	jsonStr, _ := json.Marshal(currentModel)
	_, _ = log.Debugf("Cluster Response --- value: %s ", jsonStr)
	return *pe, nil
}

func validateProgress(client *mongodbatlas.Client, currentModel *Model, currentState, targetState string) (handler.ProgressEvent, error) {
	_, _ = log.Debugf(" Cluster validateProgress() currentModel:%+v", currentModel)

	isReady, state, cluster, err := isClusterInTargetState(client, *currentModel.ProjectId, *currentModel.Name, targetState)
	if err != nil {
		_, _ = log.Debugf("ERROR Cluster validateProgress() err:%+v", err)
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
		currentModel.StateName = &cluster.StateName
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
