// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	atlasSDK "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.TenantName}
var ReadRequiredFields = []string{constants.ProjectID, constants.TenantName}
var UpdateRequiredFields = []string{constants.ProjectID, constants.TenantName}
var DeleteRequiredFields = []string{constants.ProjectID, constants.TenantName}
var ListRequiredFields = []string{constants.ProjectID}

const (
	CREATE = "CREATE"
	READ   = "READ"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
	LIST   = "LIST"
)

func setup() {
	util.SetupLogger("mongodb-atlas-data-federation")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	dataLakeTenantInput := currentModel.setDataLakeTenant()
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	dataLakeTenant, response, err := client.AtlasSDK.DataFederationApi.CreateFederatedDatabaseWithParams(
		context.Background(),
		&atlasSDK.CreateFederatedDatabaseApiParams{
			GroupId:        *currentModel.ProjectId,
			DataLakeTenant: &dataLakeTenantInput,
		}).Execute()

	defer closeResponse(response)
	if err != nil {
		return handleError(response, CREATE, err)
	}
	readModel := Model{ProjectId: currentModel.ProjectId, TenantName: currentModel.TenantName, Profile: currentModel.Profile}
	readModel.getDataLakeTenant(*dataLakeTenant)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   readModel}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	dataLakeTenant, response, err := client.AtlasSDK.DataFederationApi.GetFederatedDatabase(context.Background(), *currentModel.ProjectId, *currentModel.TenantName).Execute()

	defer closeResponse(response)
	if err != nil {
		return handleError(response, READ, err)
	}
	currentModel.getDataLakeTenant(*dataLakeTenant)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
		ResourceModel:   currentModel}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	dataLakeTenantInput := currentModel.setDataLakeTenant()
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	_, checkExistsResponse, checkExistsErr := client.AtlasSDK.DataFederationApi.GetFederatedDatabase(context.Background(), *currentModel.ProjectId, *currentModel.TenantName).Execute()
	if checkExistsErr != nil {
		return handleError(checkExistsResponse, UPDATE, checkExistsErr)
	}

	updateFederatedDatabaseAPIRequest := client.AtlasSDK.DataFederationApi.UpdateFederatedDatabase(context.Background(), *currentModel.ProjectId, *currentModel.TenantName, &dataLakeTenantInput)
	updateFederatedDatabaseAPIRequest = updateFederatedDatabaseAPIRequest.SkipRoleValidation(*currentModel.SkipRoleValidation)
	dataLakeTenant, response, err := updateFederatedDatabaseAPIRequest.Execute()

	defer closeResponse(response)
	if err != nil {
		return handleError(response, UPDATE, err)
	}
	readModel := Model{ProjectId: currentModel.ProjectId, TenantName: currentModel.TenantName, Profile: currentModel.Profile}
	readModel.getDataLakeTenant(*dataLakeTenant)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   readModel}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	_, response, err := client.AtlasSDK.DataFederationApi.DeleteFederatedDatabase(context.Background(), *currentModel.ProjectId, *currentModel.TenantName).Execute()

	defer closeResponse(response)
	if err != nil {
		return handleError(response, DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
		ResourceModel:   nil}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	dataLakeTenants, response, err := client.AtlasSDK.DataFederationApi.ListFederatedDatabases(context.Background(), *currentModel.ProjectId).Execute()

	defer closeResponse(response)
	if err != nil {
		return handleError(response, LIST, err)
	}
	tenants := make([]interface{}, 0)
	for i := range dataLakeTenants {
		tenant := Model{Profile: currentModel.Profile}
		tenant.getDataLakeTenant(dataLakeTenants[i])
		tenants = append(tenants, tenant)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Completed",
		ResourceModels:  tenants}, nil
}

func closeResponse(response *http.Response) {
	if response != nil {
		response.Body.Close()
	}
}

func handleError(response *http.Response, method string, err error) (handler.ProgressEvent, error) {
	_, _ = logger.Warnf("%s failed, error: %s", method, err.Error())
	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          fmt.Sprintf("%s:%s", method, err.Error()),
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}
	if response.StatusCode == http.StatusNotFound {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          fmt.Sprintf("%s:%s", method, err.Error()),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error during execution : %s", err.Error()),
		response), nil
}

func (model *Model) setDataLakeTenant() (dataLakeTenant atlasSDK.DataLakeTenant) {
	dataLakeTenant = atlasSDK.DataLakeTenant{
		Name:    model.TenantName,
		Storage: model.newDataFederationDataStorage(),
	}

	cloudProviderConfig := model.CloudProviderConfig
	if cloudProviderConfig != nil && cloudProviderConfig.TestS3Bucket != nil && cloudProviderConfig.RoleId != nil {
		dataLakeTenant.CloudProviderConfig = &atlasSDK.DataLakeCloudProviderConfig{
			Aws: atlasSDK.DataLakeAWSCloudProviderConfig{
				TestS3Bucket: *cloudProviderConfig.TestS3Bucket,
				RoleId:       *cloudProviderConfig.RoleId,
			},
		}
	}

	dataProcessRegion := model.DataProcessRegion
	if dataProcessRegion != nil && dataProcessRegion.CloudProvider != nil && dataProcessRegion.Region != nil {
		dataLakeTenant.DataProcessRegion.CloudProvider = *dataProcessRegion.CloudProvider
		dataLakeTenant.DataProcessRegion = &atlasSDK.DataLakeDataProcessRegion{
			CloudProvider: *dataProcessRegion.CloudProvider,
			Region:        *dataProcessRegion.Region,
		}
	}

	if dataProcessRegion != nil && dataProcessRegion.Region != nil {
		dataLakeTenant.DataProcessRegion = &atlasSDK.DataLakeDataProcessRegion{
			CloudProvider: constants.AWS,
			Region:        *dataProcessRegion.Region,
		}
	}

	return dataLakeTenant
}

func (model *Model) newDataFederationDataStorage() *atlasSDK.DataLakeStorage {
	return &atlasSDK.DataLakeStorage{
		Databases: model.newDataFederationDatabase(),
		Stores:    model.newStores(),
	}
}

func (model *Model) newDataFederationDatabase() *[]atlasSDK.DataLakeDatabaseInstance {
	if model.Storage == nil {
		return nil
	}
	storageDBs := model.Storage.Databases
	if len(storageDBs) == 0 {
		return nil
	}

	dbs := make([]atlasSDK.DataLakeDatabaseInstance, len(storageDBs))
	for i := range storageDBs {
		dbs[i] = atlasSDK.DataLakeDatabaseInstance{
			Name:        storageDBs[i].Name,
			Collections: newDataFederationCollections(storageDBs[i].Collections),
		}

		if storageDBs[i].MaxWildcardCollections != nil {
			dbs[i].MaxWildcardCollections = util.StrPtrToIntPtr(storageDBs[i].MaxWildcardCollections)
		}
	}
	return &dbs
}

func newDataFederationCollections(storageDBCollections []Collection) *[]atlasSDK.DataLakeDatabaseCollection {
	if len(storageDBCollections) == 0 {
		return nil
	}

	collections := make([]atlasSDK.DataLakeDatabaseCollection, len(storageDBCollections))
	for i := range storageDBCollections {
		collections[i] = atlasSDK.DataLakeDatabaseCollection{
			Name:        storageDBCollections[i].Name,
			DataSources: newDataFederationDataSource(storageDBCollections[i].DataSources),
		}
	}

	return &collections
}

func newDataFederationDataSource(dataSources []DataSource) *[]atlasSDK.DataLakeDatabaseDataSourceSettings {
	if len(dataSources) == 0 {
		return nil
	}
	dataSourceSettings := make([]atlasSDK.DataLakeDatabaseDataSourceSettings, len(dataSources))
	for i := range dataSources {
		dataSourceSettings[i] = atlasSDK.DataLakeDatabaseDataSourceSettings{
			AllowInsecure:       dataSources[i].AllowInsecure,
			Database:            dataSources[i].Database,
			Collection:          dataSources[i].Collection,
			CollectionRegex:     dataSources[i].CollectionRegex,
			DefaultFormat:       dataSources[i].DefaultFormat,
			Path:                dataSources[i].Path,
			ProvenanceFieldName: dataSources[i].ProvenanceFieldName,
			StoreName:           dataSources[i].StoreName,
			Urls:                &dataSources[i].Urls,
		}
	}
	return &dataSourceSettings
}

func (model *Model) newStores() *[]atlasSDK.DataLakeStoreSettings {
	if model.Storage == nil {
		return nil
	}
	stores := model.Storage.Stores
	if len(stores) == 0 {
		return nil
	}

	dataLakeStores := make([]atlasSDK.DataLakeStoreSettings, len(stores))
	for i := range stores {
		dataLakeStores[i] = atlasSDK.DataLakeStoreSettings{
			Name:        stores[i].Name,
			ProjectId:   stores[i].ProjectId,
			ClusterName: stores[i].ClusterName,
		}
		if stores[i].Provider != nil {
			dataLakeStores[i].Provider = *stores[i].Provider
		}
	}

	return &dataLakeStores
}

func (model *Model) getDataLakeTenant(dataLakeTenant atlasSDK.DataLakeTenant) {
	model.Storage = getDataLakeStorage(dataLakeTenant.Storage)
	model.TenantName = dataLakeTenant.Name
	model.ProjectId = dataLakeTenant.GroupId
	model.CloudProviderConfig = &CloudProviderConfig{
		ExternalId:        dataLakeTenant.GetCloudProviderConfig().Aws.ExternalId,
		IamAssumedRoleARN: dataLakeTenant.GetCloudProviderConfig().Aws.IamAssumedRoleARN,
		IamUserARN:        dataLakeTenant.GetCloudProviderConfig().Aws.IamUserARN,
		RoleId:            util.StringPtr(dataLakeTenant.GetCloudProviderConfig().Aws.RoleId),
		TestS3Bucket:      util.StringPtr(dataLakeTenant.GetCloudProviderConfig().Aws.TestS3Bucket),
	}
	model.DataProcessRegion = &DataProcessRegion{
		Region: util.StringPtr(dataLakeTenant.GetDataProcessRegion().Region),
	}
	model.State = dataLakeTenant.State
	model.HostNames = dataLakeTenant.GetHostnames()
}

func getDataLakeStorage(storage *atlasSDK.DataLakeStorage) *Storage {
	atlasDataLakeStorage := &Storage{
		Databases: getDataLakeDatabases(storage.GetDatabases()),
		Stores:    getDataLakeStores(storage.GetStores()),
	}
	return atlasDataLakeStorage
}

func getDataLakeDatabases(dbs []atlasSDK.DataLakeDatabaseInstance) []Database {
	dataLakeDbs := make([]Database, len(dbs))
	for i := range dbs {
		dataLakeDbs[i] = getDataLakeDatabase(dbs[i])
	}
	return dataLakeDbs
}

func getDataLakeDatabase(db atlasSDK.DataLakeDatabaseInstance) Database {
	atlasDataLakeDatabase := Database{
		Collections: getCollections(db.GetCollections()),
		Name:        db.Name,
		Views:       getViews(db.GetViews()),
	}
	if db.MaxWildcardCollections != nil {
		atlasDataLakeDatabase.MaxWildcardCollections = util.IntPtrToStrPtr(db.MaxWildcardCollections)
	}

	return atlasDataLakeDatabase
}

func getCollections(dbCollections []atlasSDK.DataLakeDatabaseCollection) []Collection {
	collections := make([]Collection, len(dbCollections))

	for i := range dbCollections {
		collections[i] = Collection{
			Name:        dbCollections[i].Name,
			DataSources: getDataSources(dbCollections[i].GetDataSources()),
		}
	}
	return collections
}

func getDataSources(dss []atlasSDK.DataLakeDatabaseDataSourceSettings) []DataSource {
	dataSources := make([]DataSource, len(dss))

	for i := range dss {
		dataSources[i] = DataSource{
			AllowInsecure:       dss[i].AllowInsecure,
			Collection:          dss[i].Collection,
			CollectionRegex:     dss[i].CollectionRegex,
			Database:            dss[i].Database,
			DatabaseRegex:       dss[i].DatabaseRegex,
			DefaultFormat:       dss[i].DefaultFormat,
			Path:                dss[i].Path,
			ProvenanceFieldName: dss[i].ProvenanceFieldName,
			StoreName:           dss[i].StoreName,
			Urls:                dss[i].GetUrls(),
		}
	}
	return dataSources
}

func getViews(dlAPIBases []atlasSDK.DataLakeApiBase) []View {
	views := make([]View, len(dlAPIBases))
	for i := range dlAPIBases {
		views[i] = View{
			Name:     dlAPIBases[i].Name,
			Pipeline: dlAPIBases[i].Pipeline,
			Source:   dlAPIBases[i].Source,
		}
	}
	return views
}

func getDataLakeStores(storeSettings []atlasSDK.DataLakeStoreSettings) []Store {
	var settings []Store
	if storeSettings == nil {
		return settings
	}
	settings = make([]Store, len(storeSettings))
	for i := range storeSettings {
		settings[i] = Store{
			Name:           storeSettings[i].Name,
			Provider:       &storeSettings[i].Provider,
			ProjectId:      storeSettings[i].ProjectId,
			ClusterName:    storeSettings[i].ClusterName,
			ReadPreference: getReadPreference(storeSettings[i].ReadPreference),
		}
	}
	return settings
}

func getReadPreference(storeReadPreference *atlasSDK.DataLakeAtlasStoreReadPreference) *ReadPreference {
	if storeReadPreference == nil {
		return nil
	}
	readPreference := &ReadPreference{
		Mode:    storeReadPreference.Mode,
		TagSets: getTagSets(storeReadPreference.GetTagSets()),
	}
	if storeReadPreference.MaxStalenessSeconds != nil {
		readPreference.MaxStalenessSeconds = util.IntPtrToStrPtr(storeReadPreference.MaxStalenessSeconds)
	}
	return readPreference
}

func getTagSets(readRefTagSets [][]atlasSDK.DataLakeAtlasStoreReadPreferenceTag) [][]TagSet {
	tagSets := make([][]TagSet, len(readRefTagSets))
	for i := range readRefTagSets {
		tagSet := make([]TagSet, len(readRefTagSets[i]))
		for j := range readRefTagSets[i] {
			tagSet[j] = TagSet{
				Name:  readRefTagSets[i][j].Name,
				Value: readRefTagSets[i][j].Value,
			}
		}
		tagSets[i] = tagSet
	}
	return tagSets
}
