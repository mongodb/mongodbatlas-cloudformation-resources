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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	atlasSDK "go.mongodb.org/atlas-sdk/v20230201002/admin"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.Name, constants.DataFederationRoleID, constants.DataFederationTestS3Bucket, constants.DataProcessRegion}
var ReadRequiredFields = []string{constants.ProjectID, constants.Name}
var UpdateRequiredFields = []string{constants.ProjectID, constants.Name, constants.DataFederationRoleID, constants.DataFederationTestS3Bucket, constants.SkipRoleValidation}
var DeleteRequiredFields = []string{constants.ProjectID, constants.Name}
var ListRequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-data-federation")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	_, _ = logger.Warnf("Create Request: %+v\n", currentModel)

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	dataLakeTenantInput, _ := currentModel.setDataLakeTenant()
	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	createFederatedDatabaseAPIRequest := atlas.AtlasV2.DataFederationApi.CreateFederatedDatabase(context.Background(), *currentModel.ProjectId, &dataLakeTenantInput)
	dataLakeTenant, response, err := createFederatedDatabaseAPIRequest.Execute()
	defer CloseResponse(response)
	if err != nil {
		_, _ = logger.Warnf("Create failed, error: %s", err.Error())

		return handleError(response, err)
	}
	currentModel.getDataLakeTenant(*dataLakeTenant)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   currentModel}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	_, _ = logger.Debugf("Read Request: %+v", currentModel)

	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	_, _ = logger.Debugf("Initiating Read Execute: %+v", currentModel)

	getFederatedDatabaseAPIRequest := atlas.AtlasV2.DataFederationApi.GetFederatedDatabase(context.Background(), *currentModel.ProjectId, *currentModel.Name)
	dataLakeTenant, response, err := getFederatedDatabaseAPIRequest.Execute()
	defer CloseResponse(response)

	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, err)
	}
	currentModel.getDataLakeTenant(*dataLakeTenant)
	_, _ = logger.Debugf("Read Response: %+v", currentModel)

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

	dataLakeTenantInput, _ := currentModel.setDataLakeTenant()
	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	updateFederatedDatabaseAPIRequest := atlas.AtlasV2.DataFederationApi.UpdateFederatedDatabase(context.Background(), *currentModel.ProjectId, *currentModel.Name, &dataLakeTenantInput)
	updateFederatedDatabaseAPIRequest = updateFederatedDatabaseAPIRequest.SkipRoleValidation(*currentModel.SkipRoleValidation)
	dataLakeTenant, response, err := updateFederatedDatabaseAPIRequest.Execute()
	defer CloseResponse(response)
	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, err)
	}
	currentModel.getDataLakeTenant(*dataLakeTenant)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   currentModel}, nil
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

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	deleteFederatedDatabaseAPIRequest := atlas.AtlasV2.DataFederationApi.DeleteFederatedDatabase(context.Background(), *currentModel.ProjectId, *currentModel.Name)
	_, response, err := deleteFederatedDatabaseAPIRequest.Execute()
	defer CloseResponse(response)

	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, err)
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

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	listFederatedDatabaseAPIRequest := atlas.AtlasV2.DataFederationApi.ListFederatedDatabases(context.Background(), *currentModel.ProjectId)

	dataLakeTenants, response, err := listFederatedDatabaseAPIRequest.Execute()
	defer CloseResponse(response)

	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, err)
	}
	var tenants = make([]Model, len(dataLakeTenants))
	for i := range dataLakeTenants {
		model := Model{}
		model.getDataLakeTenant(dataLakeTenants[i])
		tenants[i] = model
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Completed",
		ResourceModel:   tenants}, nil
}

func CloseResponse(response *http.Response) {
	if response != nil {
		response.Body.Close()
	}
	return
}

func handleError(response *http.Response, err error) (handler.ProgressEvent, error) {
	if response.StatusCode == http.StatusConflict {
		return progress_events.GetFailedEventByCode("Resource already exists",
			cloudformation.HandlerErrorCodeAlreadyExists), nil
	}
	if response != nil && response.StatusCode == 404 {
		_, _ = logger.Warnf("update 404 err: %+v", err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error while deleting resource : %s", err.Error()),
		response), nil
}

func (model *Model) setDataLakeTenant() (dataLakeTenant atlasSDK.DataLakeTenant, err error) {
	dataLakeTenant = atlasSDK.DataLakeTenant{
		Name: model.Name,
		CloudProviderConfig: &atlasSDK.DataLakeCloudProviderConfig{
			Aws: atlasSDK.DataLakeAWSCloudProviderConfig{
				TestS3Bucket: *model.CloudProviderConfig.TestS3Bucket,
				RoleId:       *model.CloudProviderConfig.RoleId,
			},
		},
		DataProcessRegion: &atlasSDK.DataLakeDataProcessRegion{
			Region:        *model.DataProcessRegion.Region,
			CloudProvider: constants.AWS,
		},

		Storage: model.newDataFederationDataStorage(),
	}
	return dataLakeTenant, nil
}

func (model *Model) newDataFederationDataStorage() *atlasSDK.DataLakeStorage {
	return &atlasSDK.DataLakeStorage{
		Databases: model.newDataFederationDatabase(),
		Stores:    model.newStores(),
	}
}

func (model *Model) newDataFederationDatabase() []atlasSDK.DataLakeDatabaseInstance {
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
			maxWildColl := cast.ToInt(storageDBs[i].MaxWildcardCollections)
			dbs[i].MaxWildcardCollections = &maxWildColl
		}
	}
	return dbs
}

func newDataFederationCollections(storageDBColls []Collection) []atlasSDK.DataLakeDatabaseCollection {
	if len(storageDBColls) == 0 {
		return nil
	}

	collections := make([]atlasSDK.DataLakeDatabaseCollection, len(storageDBColls))
	for i := range storageDBColls {
		collections[i] = atlasSDK.DataLakeDatabaseCollection{
			Name:        storageDBColls[i].Name,
			DataSources: newDataFederationDataSource(storageDBColls[i].DataSources),
		}
	}

	return collections
}

func newDataFederationDataSource(dataSrcs []DataSource) []atlasSDK.DataLakeDatabaseDataSourceSettings {
	if len(dataSrcs) == 0 {
		return nil
	}
	dataSources := make([]atlasSDK.DataLakeDatabaseDataSourceSettings, len(dataSrcs))
	for i := range dataSrcs {
		dataSources[i] = atlasSDK.DataLakeDatabaseDataSourceSettings{
			AllowInsecure:       dataSrcs[i].AllowInsecure,
			Database:            dataSrcs[i].Database,
			Collection:          dataSrcs[i].Collection,
			CollectionRegex:     dataSrcs[i].CollectionRegex,
			DefaultFormat:       dataSrcs[i].DefaultFormat,
			Path:                dataSrcs[i].Path,
			ProvenanceFieldName: dataSrcs[i].ProvenanceFieldName,
			StoreName:           dataSrcs[i].StoreName,
			Urls:                dataSrcs[i].Urls,
		}
	}
	return dataSources
}

func (model *Model) newStores() []atlasSDK.DataLakeStoreSettings {
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
			Name:                     stores[i].Name,
			Region:                   stores[i].Region,
			ProjectId:                stores[i].ProjectId,
			Bucket:                   stores[i].Bucket,
			ClusterName:              stores[i].ClusterName,
			Prefix:                   stores[i].Prefix,
			Delimiter:                stores[i].Delimiter,
			IncludeTags:              stores[i].IncludeTags,
			AdditionalStorageClasses: stores[i].AdditionalStorageClasses,
		}
		if stores[i].Provider != nil {
			dataLakeStores[i].Provider = *stores[i].Provider
		}
	}

	return dataLakeStores
}

func (model *Model) getDataLakeTenant(dataLakeTenant atlasSDK.DataLakeTenant) {
	model.Storage = getDataLakeStorage(dataLakeTenant.Storage)
	model.Name = dataLakeTenant.Name
	model.CloudProviderConfig = &AtlasDataLakeCloudProviderConfig{
		ExternalId:        dataLakeTenant.CloudProviderConfig.Aws.ExternalId,
		IamAssumedRoleARN: dataLakeTenant.CloudProviderConfig.Aws.IamAssumedRoleARN,
		IamUserARN:        dataLakeTenant.CloudProviderConfig.Aws.IamUserARN,
		RoleId:            &dataLakeTenant.CloudProviderConfig.Aws.RoleId,
		TestS3Bucket:      &dataLakeTenant.CloudProviderConfig.Aws.TestS3Bucket,
	}
	model.DataProcessRegion = &AtlasDataLakeDataProcessRegion{
		Region: &dataLakeTenant.DataProcessRegion.Region,
	}
}

func getDataLakeStorage(storage *atlasSDK.DataLakeStorage) *AtlasDataLakeStorage {
	atlasDataLakeStorage := AtlasDataLakeStorage{
		Databases: getDataLakeDatabases(storage.Databases),
		Stores:    getDataLakeStores(storage.Stores),
	}
	return &atlasDataLakeStorage
}

func getDataLakeDatabases(dbs []atlasSDK.DataLakeDatabaseInstance) []AtlasDataLakeDatabase {
	dataLakeDbs := make([]AtlasDataLakeDatabase, len(dbs))
	for i := range dbs {
		dataLakeDbs[i] = getDataLakeDatabase(dbs[i])
	}
	return dataLakeDbs
}

func getDataLakeDatabase(db atlasSDK.DataLakeDatabaseInstance) AtlasDataLakeDatabase {
	atlasDataLakeDatabase := AtlasDataLakeDatabase{
		Collections: getCollections(db.Collections),
		Name:        db.Name,
		Views:       getViews(db.Views),
	}
	if db.MaxWildcardCollections != nil {
		maxWildCardCollCount := cast.ToString(*db.MaxWildcardCollections)
		atlasDataLakeDatabase.MaxWildcardCollections = &maxWildCardCollCount
	}

	return atlasDataLakeDatabase
}

func getCollections(dbCollections []atlasSDK.DataLakeDatabaseCollection) []Collection {
	collections := make([]Collection, len(dbCollections))

	for i := range dbCollections {
		collections[i] = Collection{
			Name:        dbCollections[i].Name,
			DataSources: getDataSources(dbCollections[i].DataSources),
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
			Urls:                dss[i].Urls,
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

func getDataLakeStores(storeSettings []atlasSDK.DataLakeStoreSettings) []AtlasDataLakeStore {
	var settings []AtlasDataLakeStore
	if storeSettings == nil {
		return settings
	}
	settings = make([]AtlasDataLakeStore, len(storeSettings))
	for i := range storeSettings {
		settings[i] = AtlasDataLakeStore{
			Name:                     storeSettings[i].Name,
			Provider:                 &storeSettings[i].Provider,
			AdditionalStorageClasses: storeSettings[i].AdditionalStorageClasses,
			Bucket:                   storeSettings[i].Bucket,
			Delimiter:                storeSettings[i].Delimiter,
			IncludeTags:              storeSettings[i].IncludeTags,
			Prefix:                   storeSettings[i].Prefix,
			Public:                   storeSettings[i].Public,
			Region:                   storeSettings[i].Region,
			ProjectId:                storeSettings[i].ProjectId,
			ClusterName:              storeSettings[i].ClusterName,
			ReadPreference:           getReadPreference(storeSettings[i].ReadPreference),
		}
	}
	return settings
}

func getReadPreference(storeReadPreference *atlasSDK.DataLakeAtlasStoreReadPreference) *ReadPreference {
	if storeReadPreference == nil {
		return nil
	}
	readPreference := &ReadPreference{
		Mode: storeReadPreference.Mode,
		Tags: getTagSets(storeReadPreference.TagSets),
	}
	if storeReadPreference.MaxStalenessSeconds != nil {
		maxStaleness := cast.ToString(storeReadPreference.MaxStalenessSeconds)
		readPreference.MaxStalenessSeconds = &maxStaleness
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
