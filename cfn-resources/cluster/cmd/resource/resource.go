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

	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"
	admin20250312006 "go.mongodb.org/atlas-sdk/v20250312006/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/spf13/cast"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
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

	if util.IsFlexCluster(currentModel.ClusterType) {
		return createFlexCluster(req, client, currentModel)
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
	if util.IsFlexCluster(currentModel.ClusterType) {
		flexResp, resp, err := client.AtlasSDK.FlexClustersApi.GetFlexCluster(context.Background(), *currentModel.ProjectId, *currentModel.Name).Execute()
		if pe := util.HandleClusterError(err, resp); pe != nil {
			return *pe, nil
		}
		updateModelFromFlexCluster(currentModel, flexResp)
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         constants.ReadComplete,
			ResourceModel:   currentModel,
		}, nil
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
	if util.IsFlexCluster(currentModel.ClusterType) {
		return updateFlexCluster(req, client, currentModel)
	}

	// Update callback
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

	// Update Cluster
	model, _, err := updateCluster(context.Background(), client, currentModel, adminCluster)
	if err != nil {
		if apiError, ok := admin20231115014.AsError(err); ok && *apiError.Error == http.StatusNotFound {
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
	client, setupErr := setupRequest(req, currentModel, createReadUpdareDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	if util.IsFlexCluster(currentModel.ClusterType) {
		return deleteFlexCluster(req, client, currentModel)
	}

	ctx := context.Background()

	if _, ok := req.CallbackContext[constants.StateName]; ok {
		return validateProgress(client, currentModel, constants.DeletedState)
	}

	params := &admin20231115014.DeleteClusterApiParams{
		RetainBackups: util.Pointer(false),
		GroupId:       *currentModel.ProjectId,
		ClusterName:   *currentModel.Name,
	}

	_, err := client.Atlas20231115014.ClustersApi.DeleteClusterWithParams(ctx, params).Execute()
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

		_, _ = log.Debugf("Cluster Creation completed:%s", *currentModel.Name)

		cluster, res, err := client.Atlas20231115014.ClustersApi.GetCluster(context.Background(), projectID, *currentModel.Name).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
				res), nil
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

func isClusterInTargetState(client *util.MongoDBClient, projectID, clusterName, targetState string) (isReady bool, stateName string, mongoCluster *admin20231115014.AdvancedClusterDescription, err error) {
	cluster, resp, err := client.Atlas20231115014.ClustersApi.GetCluster(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return constants.DeletedState == targetState, constants.DeletedState, nil, nil
		}
		return false, constants.Error, nil, fmt.Errorf("error fetching cluster info (%s): %s", clusterName, err)
	}
	_, _ = log.Debugf("Cluster state: %s, targetState : %s", *cluster.StateName, targetState)
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
	_, _ = log.Debugf("params : %+v %+v %+v", ctx, client, clusterRequest)
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
		_, _ = log.Debugf("compelted updation:%s", *currentModel.Name)
		cluster, res, err := client.Atlas20231115014.ClustersApi.GetCluster(context.Background(), projectID, *currentModel.Name).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error in Get Cluster : %s", err.Error()),
				res), nil
		}

		_, _ = log.Debugf("Updating cluster :%s", *currentModel.Name)

		return updateClusterSettings(currentModel, client, projectID, cluster, &progressEvent)
	}
	return progressEvent, nil
}

func updateClusterSettings(currentModel *Model, client *util.MongoDBClient,
	projectID string, cluster *admin20231115014.AdvancedClusterDescription, pe *handler.ProgressEvent) (handler.ProgressEvent, error) {
	// Update advanced configuration
	if currentModel.AdvancedSettings != nil {
		_, _ = log.Debugf("AdvancedSettings: %+v", *currentModel.AdvancedSettings)

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
			_, _ = log.Warnf("Cluster Pause - error: %+v", err)
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Cluster Pause error : %s", err.Error()),
				res), err
		}
	}

	jsonStr, _ := json.Marshal(currentModel)
	_, _ = log.Debugf("Cluster Response --- value: %s ", jsonStr)
	return *pe, nil
}

func validateProgress(client *util.MongoDBClient, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
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

// createFlexCluster handles creation of flex clusters
func createFlexCluster(req handler.Request, client *util.MongoDBClient, currentModel *Model) (handler.ProgressEvent, error) {
	// Check if this is a callback
	if _, isCallback := req.CallbackContext["callback"]; isCallback {
		return validateFlexClusterProgress(client, currentModel, false)
	}

	// Extract provider settings from ReplicationSpecs if available
	backingProvider := "AWS" // Default
	region := "US_EAST_1"    // Default

	if len(currentModel.ReplicationSpecs) > 0 && len(currentModel.ReplicationSpecs[0].AdvancedRegionConfigs) > 0 {
		regionConfig := currentModel.ReplicationSpecs[0].AdvancedRegionConfigs[0]
		if regionConfig.ProviderName != nil {
			backingProvider = *regionConfig.ProviderName
		}
		if regionConfig.RegionName != nil {
			region = *regionConfig.RegionName
		}
	}

	// Convert tags to the new SDK format
	var tags *[]admin20250312006.ResourceTag
	if len(currentModel.Tags) > 0 {
		convertedTags := make([]util.TagModel, len(currentModel.Tags))
		for i, tag := range currentModel.Tags {
			convertedTags[i] = util.TagModel{
				Key:   tag.Key,
				Value: tag.Value,
			}
		}
		tags = util.ExpandTags(convertedTags)
	}

	flexReq := util.CreateFlexClusterRequest(
		*currentModel.Name,
		backingProvider,
		region,
		currentModel.TerminationProtectionEnabled,
		tags,
	)

	flexResp, resp, err := client.AtlasSDK.FlexClustersApi.CreateFlexCluster(context.Background(), *currentModel.ProjectId, flexReq).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe, nil
	}

	updateModelFromFlexCluster(currentModel, flexResp)

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              constants.Pending,
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 10,
		CallbackContext:      map[string]interface{}{"callback": true},
	}, nil
}

// validateFlexClusterProgress validates the progress of flex cluster operations
func validateFlexClusterProgress(client *util.MongoDBClient, model *Model, isDelete bool) (handler.ProgressEvent, error) {
	state, flexResp, err := util.ValidateFlexClusterProgress(client, *model.ProjectId, *model.Name, isDelete)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError,
		}, nil
	}

	targetState := constants.IdleState
	if isDelete {
		targetState = constants.DeletedState
	}

	if state != targetState {
		updateModelFromFlexCluster(model, flexResp)
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              constants.Pending,
			ResourceModel:        model,
			CallbackDelaySeconds: 10,
			CallbackContext:      map[string]interface{}{"callback": true},
		}, nil
	}

	if isDelete {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         constants.Complete,
		}, nil
	}

	updateModelFromFlexCluster(model, flexResp)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModel:   model,
	}, nil
}

// updateModelFromFlexCluster updates the model with flex cluster response data
func updateModelFromFlexCluster(model *Model, flexResp *admin20250312006.FlexClusterDescription20241113) {
	if flexResp == nil {
		return
	}

	model.Id = flexResp.Id
	model.StateName = flexResp.StateName
	model.CreatedDate = util.TimePtrToStringPtr(flexResp.CreateDate)
	model.MongoDBVersion = flexResp.MongoDBVersion
	model.VersionReleaseSystem = flexResp.VersionReleaseSystem
	model.TerminationProtectionEnabled = flexResp.TerminationProtectionEnabled
	model.ClusterType = flexResp.ClusterType

	if flexResp.BackupSettings != nil && flexResp.BackupSettings.Enabled != nil {
		model.BackupEnabled = flexResp.BackupSettings.Enabled
	}

	if flexResp.ConnectionStrings != nil {
		model.ConnectionStrings = &ConnectionStrings{
			Standard:    flexResp.ConnectionStrings.Standard,
			StandardSrv: flexResp.ConnectionStrings.StandardSrv,
		}
	}

	if flexResp.Tags != nil {
		convertedTags := util.FlattenTags(flexResp.Tags)
		model.Tags = make([]Tag, len(convertedTags))
		for i, tag := range convertedTags {
			model.Tags[i] = Tag{
				Key:   tag.Key,
				Value: tag.Value,
			}
		}
	}

	// Update ReplicationSpecs with flex cluster provider settings if needed
	if flexResp.ProviderSettings.BackingProviderName != nil {
		if flexResp.ProviderSettings.DiskSizeGB != nil {
			model.DiskSizeGB = flexResp.ProviderSettings.DiskSizeGB
		}
		// Keep ReplicationSpecs minimal for flex clusters
		if len(model.ReplicationSpecs) == 0 {
			model.ReplicationSpecs = []AdvancedReplicationSpec{{
				AdvancedRegionConfigs: []AdvancedRegionConfig{{
					ProviderName: flexResp.ProviderSettings.BackingProviderName,
					RegionName:   flexResp.ProviderSettings.RegionName,
				}},
			}}
		}
	}
}

// updateFlexCluster handles updating flex clusters
func updateFlexCluster(req handler.Request, client *util.MongoDBClient, model *Model) (handler.ProgressEvent, error) {
	// Check if this is a callback
	if _, isCallback := req.CallbackContext["callback"]; isCallback {
		return validateFlexClusterProgress(client, model, false)
	}

	// Convert tags to the new SDK format
	var tags *[]admin20250312006.ResourceTag
	if len(model.Tags) > 0 {
		convertedTags := make([]util.TagModel, len(model.Tags))
		for i, tag := range model.Tags {
			convertedTags[i] = util.TagModel{
				Key:   tag.Key,
				Value: tag.Value,
			}
		}
		tags = util.ExpandTags(convertedTags)
	}

	updateReq := &admin20250312006.FlexClusterDescriptionUpdate20241113{
		TerminationProtectionEnabled: model.TerminationProtectionEnabled,
		Tags:                         tags,
	}

	flexResp, resp, err := client.AtlasSDK.FlexClustersApi.UpdateFlexCluster(context.Background(), *model.ProjectId, *model.Name, updateReq).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe, nil
	}

	updateModelFromFlexCluster(model, flexResp)

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              constants.Pending,
		ResourceModel:        model,
		CallbackDelaySeconds: 10,
		CallbackContext:      map[string]interface{}{"callback": true},
	}, nil
}

// deleteFlexCluster handles deleting flex clusters
func deleteFlexCluster(req handler.Request, client *util.MongoDBClient, model *Model) (handler.ProgressEvent, error) {
	// Check if this is a callback
	if _, isCallback := req.CallbackContext["callback"]; isCallback {
		return validateFlexClusterProgress(client, model, true)
	}

	resp, err := client.AtlasSDK.FlexClustersApi.DeleteFlexCluster(context.Background(), *model.ProjectId, *model.Name).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe, nil
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              constants.DeleteInProgress,
		ResourceModel:        model,
		CallbackDelaySeconds: 10,
		CallbackContext:      map[string]interface{}{"callback": true},
	}, nil
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
