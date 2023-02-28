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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

var RequiredFields = []string{constants.ProjectID, constants.TenantName}
var ListRequiredFields = []string{constants.ProjectID}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	validationError := validateRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// progress callback setup
	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["project_id"].(string)
		currentModel.ProjectId = &sid
		return validateProgress(client, currentModel, "ACTIVE")
	}

	// Create Atlas API Request Object
	projectID := *currentModel.ProjectId
	dataLakeReq := &mongodbatlas.DataLakeCreateRequest{
		CloudProviderConfig: expandCloudProviderConfig(currentModel),
		Name:                *currentModel.TenantName,
	}

	// API call to create data lake
	dataLake, resp, err := client.DataLakes.Create(context.Background(), projectID, dataLakeReq)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}
	currentModel.ProjectId = &dataLake.GroupID

	_, _ = logger.Debugf("Created Successfully - (%s)", *currentModel.ProjectId)

	// track progress
	event := handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create cloud provider snapshots : %s", dataLake.State),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 65,
		CallbackContext: map[string]interface{}{
			"status":     dataLake.State,
			"project_id": currentModel.ProjectId,
		},
	}
	return event, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	if validationError := validateRequest(RequiredFields, currentModel); validationError != nil {
		return *validationError, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if !isExist(currentModel, client) {
		_, _ = logger.Warnf("resource not exist for Id: %s", *currentModel.TenantName)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	// API call to get
	projectID := *currentModel.ProjectId
	dataLake, resp, err := client.DataLakes.Get(context.Background(), projectID, *currentModel.TenantName)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	readModel := convertToModel(dataLake, currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   readModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if validationError := validateRequest(RequiredFields, currentModel); validationError != nil {
		return *validationError, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if !isExist(currentModel, client) {
		_, _ = logger.Warnf("resource not exist for Id: %s", *currentModel.TenantName)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	// create request object
	projectID := *currentModel.ProjectId
	dataLakeReq := &mongodbatlas.DataLakeUpdateRequest{
		CloudProviderConfig: expandCloudProviderConfig(currentModel),
		DataProcessRegion:   expandDataLakeDataProcessRegion(currentModel.DataProcessRegion),
	}

	// API call to update
	dataLake, resp, err := client.DataLakes.Update(context.Background(), projectID, *currentModel.TenantName, dataLakeReq)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}
	currentModel.ProjectId = &dataLake.GroupID

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if validationError := validateRequest(RequiredFields, currentModel); validationError != nil {
		return *validationError, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if !isExist(currentModel, client) {
		_, _ = logger.Warnf("resource not exist for Id: %s", *currentModel.TenantName)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	// API call to delete
	resp, err := client.DataLakes.Delete(context.Background(), *currentModel.ProjectId, *currentModel.TenantName)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

// List  handles the list event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if validationError := validateRequest(ListRequiredFields, currentModel); validationError != nil {
		return *validationError, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// API call to list
	result, resp, err := client.DataLakes.List(context.Background(), *currentModel.ProjectId)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	var models []interface{}
	for ind := range result {
		models = append(models, convertToModel(&result[ind], currentModel))
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}
func flattenAWSBlock(awsConfig *mongodbatlas.CloudProviderConfig) *DataLakeAWSCloudProviderConfigView {
	if awsConfig == nil {
		return nil
	}
	return &DataLakeAWSCloudProviderConfigView{
		RoleId:            &awsConfig.AWSConfig.RoleID,
		IamAssumedRoleARN: &awsConfig.AWSConfig.IAMAssumedRoleARN,
		IamUserARN:        &awsConfig.AWSConfig.IAMUserARN,
		ExternalId:        &awsConfig.AWSConfig.ExternalID,
		TestS3Bucket:      &awsConfig.AWSConfig.TestS3Bucket,
	}
}

func flattenDataLakeProcessRegion(processRegion *mongodbatlas.DataProcessRegion) *DataLakeDataProcessRegionView {
	if processRegion != nil && (processRegion.Region != "" || processRegion.CloudProvider != "") {
		return &DataLakeDataProcessRegionView{
			CloudProvider: &processRegion.CloudProvider,
			Region:        &processRegion.Region,
		}
	}
	return nil
}

func flattenDataLakeStorageDatabases(databases []mongodbatlas.DataLakeDatabase) []DataLakeDatabaseView {
	database := make([]DataLakeDatabaseView, len(databases))
	for ind := range databases {
		database = append(database, DataLakeDatabaseView{
			Name:                   &databases[ind].Name,
			Collections:            flattenDataLakeStorageDatabaseCollections(databases[ind].Collections),
			Views:                  flattenDataLakeStorageDatabaseViews(databases[ind].Views),
			MaxWildcardCollections: castNO64(databases[ind].MaxWildcardCollections),
		})
	}
	return database
}
func castNO64(i *int64) *int {
	x := cast.ToInt(&i)
	return &x
}

func flattenDataLakeStorageDatabaseCollections(collections []mongodbatlas.DataLakeCollection) []DataLakeDatabaseCollectionView {
	database := make([]DataLakeDatabaseCollectionView, 0)
	for ind := range collections {
		database = append(database, DataLakeDatabaseCollectionView{
			Name:        &collections[ind].Name,
			DataSources: flattenDataLakeStorageDatabaseCollectionsDataSources(collections[ind].DataSources),
		})
	}
	return database
}

func flattenDataLakeStorageDatabaseCollectionsDataSources(dataSources []mongodbatlas.DataLakeDataSource) []DataLakeDatabaseDataSourceView {
	database := make([]DataLakeDatabaseDataSourceView, 0)
	for ind := range dataSources {
		database = append(database, DataLakeDatabaseDataSourceView{
			StoreName:     &dataSources[ind].StoreName,
			DefaultFormat: &dataSources[ind].DefaultFormat,
			Path:          &dataSources[ind].Path,
		})
	}
	return database
}

func flattenDataLakeStorageDatabaseViews(views []mongodbatlas.DataLakeDatabaseView) []DataLakeViewView {
	view := make([]DataLakeViewView, 0)
	for ind := range views {
		view = append(view, DataLakeViewView{
			Name:     &views[ind].Name,
			Source:   &views[ind].Source,
			Pipeline: &views[ind].Pipeline,
		})
	}
	return view
}

func flattenDataLakeStorageStores(stores []mongodbatlas.DataLakeStore) []StoreDetail {
	store := make([]StoreDetail, 0)
	for ind := range stores {
		store = append(store, StoreDetail{
			Name:                     &stores[ind].Name,
			Provider:                 &stores[ind].Provider,
			Region:                   &stores[ind].Region,
			Bucket:                   &stores[ind].Bucket,
			Prefix:                   &stores[ind].Prefix,
			Delimiter:                &stores[ind].Delimiter,
			IncludeTags:              stores[ind].IncludeTags,
			AdditionalStorageClasses: stores[ind].AdditionalStorageClasses,
		})
	}
	return store
}

func expandDataLakeAwsBlock(cloudProviderConfig DataLakeCloudProviderConfigView) mongodbatlas.AwsCloudProviderConfig {
	awsConfig := mongodbatlas.AwsCloudProviderConfig{}
	if cloudProviderConfig.Aws != nil {
		awsConfig.ExternalID = cast.ToString(cloudProviderConfig.Aws.ExternalId)
		awsConfig.IAMAssumedRoleARN = cast.ToString(cloudProviderConfig.Aws.IamAssumedRoleARN)
		awsConfig.IAMUserARN = cast.ToString(cloudProviderConfig.Aws.IamUserARN)
		awsConfig.RoleID = cast.ToString(cloudProviderConfig.Aws.RoleId)
		awsConfig.TestS3Bucket = cast.ToString(cloudProviderConfig.Aws.TestS3Bucket)
	}
	return awsConfig
}
func expandCloudProviderConfig(currentModel *Model) *mongodbatlas.CloudProviderConfig {
	if currentModel.CloudProviderConfig != nil {
		return &mongodbatlas.CloudProviderConfig{
			AWSConfig: expandDataLakeAwsBlock(*currentModel.CloudProviderConfig),
		}
	}
	return nil
}

func expandDataLakeDataProcessRegion(dataProcessRegion *DataLakeDataProcessRegionView) *mongodbatlas.DataProcessRegion {
	if dataProcessRegion != nil && dataProcessRegion.Region != nil {
		return &mongodbatlas.DataProcessRegion{
			CloudProvider: cast.ToString(dataProcessRegion.CloudProvider),
			Region:        cast.ToString(dataProcessRegion.Region),
		}
	}
	return nil
}

// logger setup function
func setup() {
	util.SetupLogger("mongodb-atlas-data-lakes")
}

// function to validate inputs to all actions
func validateRequest(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// function to track snapshot creation status
func validateProgress(client *mongodbatlas.Client, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	projectID := *currentModel.ProjectId
	tenantName := *currentModel.TenantName
	isReady, state, err := dataLakeIsReady(client, projectID, tenantName, targetState)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	if !isReady {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = handler.InProgress
		p.CallbackDelaySeconds = 35
		p.Message = "Pending"
		p.CallbackContext = map[string]interface{}{
			"status":     state,
			"project_id": *currentModel.ProjectId,
		}
		return p, nil
	}

	p := handler.NewProgressEvent()
	p.ResourceModel = currentModel
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	return p, nil
}

// function to check if record already exist
func isExist(currentModel *Model, client *mongodbatlas.Client) bool {
	projectID := *currentModel.ProjectId
	tenantName := *currentModel.TenantName
	dataLake, _, err := client.DataLakes.Get(context.Background(), projectID, tenantName)
	if err != nil {
		return false
	}
	if dataLake != nil {
		return true
	}

	return false
}

// function to check if snapshot already exist in atlas
func dataLakeIsReady(client *mongodbatlas.Client, projectID, name, targetState string) (isReady bool, status string, err error) {
	dataLake, resp, err := client.DataLakes.Get(context.Background(), projectID, name)
	if err != nil {
		return false, "", err
	}
	if err != nil {
		if dataLake == nil && resp == nil {
			return false, "", err
		}
		if resp != nil && resp.StatusCode == 404 {
			return true, "deleted", nil
		}
		return false, "", err
	}
	return dataLake.State == targetState, dataLake.State, nil
}

func convertToModel(dataLake *mongodbatlas.DataLake, currentModel *Model) *Model {
	var result = new(Model)

	result.Profile = currentModel.Profile // cfn test
	result.TenantName = &dataLake.Name
	result.State = &dataLake.State
	result.ProjectId = &dataLake.GroupID
	result.CloudProviderConfig = &DataLakeCloudProviderConfigView{
		Aws: flattenAWSBlock(&dataLake.CloudProviderConfig),
	}
	result.DataProcessRegion = flattenDataLakeProcessRegion(&dataLake.DataProcessRegion)
	result.Storage = &DataLakeStorageView{
		Databases: flattenDataLakeStorageDatabases(dataLake.Storage.Databases),
		Stores:    flattenDataLakeStorageStores(dataLake.Storage.Stores),
	}
	result.Hostnames = dataLake.Hostnames
	result.State = &dataLake.State

	return result
}
