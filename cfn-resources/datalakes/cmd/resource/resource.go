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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20231115002/admin"
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

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["project_id"].(string)
		currentModel.ProjectId = &sid
		return validateProgress(client, currentModel, "ACTIVE")
	}

	projectID := *currentModel.ProjectId
	dataLakeReq := &admin.DataLakeTenant{
		CloudProviderConfig: expandCloudProviderConfig(currentModel),
		Name:                currentModel.TenantName,
	}

	dataLake, resp, err := client.Atlas20231115002.DataFederationApi.CreateFederatedDatabase(context.Background(), projectID, dataLakeReq).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}
	currentModel.ProjectId = dataLake.GroupId
	event := handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create cloud provider snapshots : %s", *dataLake.State),
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

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	projectID := *currentModel.ProjectId
	dataLake, resp, err := client.Atlas20231115002.DataFederationApi.GetFederatedDatabase(context.Background(), projectID, *currentModel.TenantName).Execute()
	if err != nil {
		if apiError, ok := admin.AsError(err); ok && *apiError.Error == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Resource Not Found",
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}

		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
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

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if !isExist(currentModel, client) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	projectID := *currentModel.ProjectId
	bodyRequest := &admin.UpdateFederatedDatabaseApiParams{
		GroupId:            projectID,
		TenantName:         *currentModel.TenantName,
		SkipRoleValidation: admin.PtrBool(false),
		DataLakeTenant: &admin.DataLakeTenant{
			CloudProviderConfig: expandCloudProviderConfig(currentModel),
			DataProcessRegion:   expandDataLakeDataProcessRegion(currentModel.DataProcessRegion),
		},
	}

	dataLake, resp, err := client.Atlas20231115002.DataFederationApi.UpdateFederatedDatabaseWithParams(context.Background(), bodyRequest).Execute()
	if err != nil {
		if resp != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}

		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Error in updating the resource",
			HandlerErrorCode: cloudformation.HandlerErrorCodeHandlerInternalFailure}, nil
	}

	currentModel.ProjectId = dataLake.GroupId

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

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if !isExist(currentModel, client) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	_, resp, err := client.Atlas20231115002.DataFederationApi.DeleteFederatedDatabase(context.Background(), *currentModel.ProjectId, *currentModel.TenantName).Execute()
	if err != nil {
		if resp != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}

		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Error in deleting the resource",
			HandlerErrorCode: cloudformation.HandlerErrorCodeHandlerInternalFailure}, nil
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

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	result, resp, err := client.Atlas20231115002.DataFederationApi.ListFederatedDatabases(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
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
func flattenAWSBlock(awsConfig *admin.DataLakeCloudProviderConfig) *DataLakeAWSCloudProviderConfigView {
	if awsConfig == nil {
		return nil
	}
	return &DataLakeAWSCloudProviderConfigView{
		RoleId:            &awsConfig.Aws.RoleId,
		IamAssumedRoleARN: awsConfig.Aws.IamAssumedRoleARN,
		IamUserARN:        awsConfig.Aws.IamUserARN,
		ExternalId:        awsConfig.Aws.ExternalId,
		TestS3Bucket:      &awsConfig.Aws.TestS3Bucket,
	}
}

func flattenDataLakeProcessRegion(processRegion *admin.DataLakeDataProcessRegion) *DataLakeDataProcessRegionView {
	if processRegion != nil && (processRegion.Region != "" || processRegion.CloudProvider != "") {
		return &DataLakeDataProcessRegionView{
			CloudProvider: &processRegion.CloudProvider,
			Region:        &processRegion.Region,
		}
	}
	return nil
}

func flattenDataLakeStorageDatabases(databases []admin.DataLakeDatabaseInstance) []DataLakeDatabaseView {
	views := make([]DataLakeDatabaseView, len(databases))
	for i := range databases {
		views[i] = DataLakeDatabaseView{
			Name:                   databases[i].Name,
			Collections:            flattenDataLakeStorageDatabaseCollections(databases[i].Collections),
			Views:                  flattenDataLakeStorageDatabaseViews(databases[i].Views),
			MaxWildcardCollections: databases[i].MaxWildcardCollections,
		}
	}
	return views
}

func flattenDataLakeStorageDatabaseCollections(collections []admin.DataLakeDatabaseCollection) []DataLakeDatabaseCollectionView {
	database := make([]DataLakeDatabaseCollectionView, 0)
	for ind := range collections {
		database = append(database, DataLakeDatabaseCollectionView{
			Name:        collections[ind].Name,
			DataSources: flattenDataLakeStorageDatabaseCollectionsDataSources(collections[ind].DataSources),
		})
	}
	return database
}

func flattenDataLakeStorageDatabaseCollectionsDataSources(dataSources []admin.DataLakeDatabaseDataSourceSettings) []DataLakeDatabaseDataSourceView {
	database := make([]DataLakeDatabaseDataSourceView, 0)
	for ind := range dataSources {
		database = append(database, DataLakeDatabaseDataSourceView{
			StoreName:     dataSources[ind].StoreName,
			DefaultFormat: dataSources[ind].DefaultFormat,
			Path:          dataSources[ind].Path,
		})
	}
	return database
}

func flattenDataLakeStorageDatabaseViews(views []admin.DataLakeApiBase) []DataLakeViewView {
	view := make([]DataLakeViewView, 0)
	for ind := range views {
		view = append(view, DataLakeViewView{
			Name:     views[ind].Name,
			Source:   views[ind].Source,
			Pipeline: views[ind].Pipeline,
		})
	}
	return view
}

func flattenDataLakeStorageStores(stores []admin.DataLakeStoreSettings) []StoreDetail {
	store := make([]StoreDetail, 0)
	for ind := range stores {
		store = append(store, StoreDetail{
			Name:                     stores[ind].Name,
			Provider:                 &stores[ind].Provider,
			Region:                   stores[ind].Region,
			Bucket:                   stores[ind].Bucket,
			Prefix:                   stores[ind].Prefix,
			Delimiter:                stores[ind].Delimiter,
			IncludeTags:              stores[ind].IncludeTags,
			AdditionalStorageClasses: stores[ind].AdditionalStorageClasses,
		})
	}
	return store
}

func expandDataLakeAwsBlock(cloudProviderConfig DataLakeCloudProviderConfigView) admin.DataLakeAWSCloudProviderConfig {
	awsConfig := admin.DataLakeAWSCloudProviderConfig{}
	if cloudProviderConfig.Aws != nil {
		awsConfig.ExternalId = cloudProviderConfig.Aws.ExternalId
		awsConfig.IamAssumedRoleARN = cloudProviderConfig.Aws.IamAssumedRoleARN
		awsConfig.IamUserARN = cloudProviderConfig.Aws.IamUserARN
		awsConfig.RoleId = cast.ToString(cloudProviderConfig.Aws.RoleId)
		awsConfig.TestS3Bucket = cast.ToString(cloudProviderConfig.Aws.TestS3Bucket)
	}
	return awsConfig
}
func expandCloudProviderConfig(currentModel *Model) *admin.DataLakeCloudProviderConfig {
	if currentModel.CloudProviderConfig != nil {
		return &admin.DataLakeCloudProviderConfig{
			Aws: expandDataLakeAwsBlock(*currentModel.CloudProviderConfig),
		}
	}
	return nil
}

func expandDataLakeDataProcessRegion(dataProcessRegion *DataLakeDataProcessRegionView) *admin.DataLakeDataProcessRegion {
	if dataProcessRegion != nil && dataProcessRegion.Region != nil {
		return &admin.DataLakeDataProcessRegion{
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
func validateProgress(client *util.MongoDBClient, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
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

func isExist(currentModel *Model, client *util.MongoDBClient) bool {
	projectID := *currentModel.ProjectId
	tenantName := *currentModel.TenantName
	dataLake, _, err := client.Atlas20231115002.DataFederationApi.GetFederatedDatabase(context.Background(), projectID, tenantName).Execute()
	return err == nil && dataLake != nil
}

// function to check if snapshot already exist in atlas
func dataLakeIsReady(client *util.MongoDBClient, projectID, name, targetState string) (isReady bool, status string, err error) {
	dataLake, _, err := client.Atlas20231115002.DataFederationApi.GetFederatedDatabase(context.Background(), projectID, name).Execute()
	if err != nil {
		return false, "", err
	}

	if dataLake != nil {
		return *dataLake.State == targetState, *dataLake.State, nil
	}

	return false, "", nil
}

func convertToModel(dataLake *admin.DataLakeTenant, currentModel *Model) *Model {
	var result = new(Model)

	result.Profile = currentModel.Profile // cfn test
	result.TenantName = dataLake.Name
	result.State = dataLake.State
	result.ProjectId = dataLake.GroupId
	result.CloudProviderConfig = &DataLakeCloudProviderConfigView{
		Aws: flattenAWSBlock(dataLake.CloudProviderConfig),
	}
	result.DataProcessRegion = flattenDataLakeProcessRegion(dataLake.DataProcessRegion)
	result.Storage = &DataLakeStorageView{
		Databases: flattenDataLakeStorageDatabases(dataLake.Storage.Databases),
		Stores:    flattenDataLakeStorageStores(dataLake.Storage.Stores),
	}
	result.Hostnames = dataLake.Hostnames
	result.State = dataLake.State

	return result
}
