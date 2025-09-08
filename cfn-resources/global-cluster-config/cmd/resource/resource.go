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
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func setup() {
	util.SetupLogger("mongodb-atlas-global-cluster-config")
}

var RequiredFields = []string{constants.ClusterName, constants.ProjectID}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errValidation := validateModel(RequiredFields, currentModel); errValidation != nil {
		return *errValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName
	nameSpaces := currentModel.ManagedNamespaces

	if err := createManagedNamespaces(context.Background(), client, nameSpaces, projectID, clusterName); err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError,
		}, nil
	}

	_, _, err := client.Atlas20231115002.GlobalClustersApi.CreateCustomZoneMapping(context.Background(), projectID, clusterName, newCustomZoneMappings(currentModel)).Execute()
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError,
		}, nil
	}

	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   currentModel,
	}
	return event, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errValidation := validateModel(RequiredFields, currentModel); errValidation != nil {
		return *errValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	config, event, err := ReadConfig(client, currentModel)
	if err != nil {
		if config == nil {
			return progressevent.GetFailedEventByCode("Resource Not Found", cloudformation.HandlerErrorCodeNotFound), nil
		}

		return event, nil
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   config,
	}, nil
}

func ReadConfig(client *util.MongoDBClient, currentModel *Model) (*Model, handler.ProgressEvent, error) {
	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName

	globalCluster, resp, err := client.Atlas20231115002.GlobalClustersApi.GetManagedNamespace(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if apiError, ok := admin20231115002.AsError(err); ok && *apiError.Error == http.StatusNotFound {
			return nil, handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}

		return nil, progressevent.GetFailedEventByResponse(fmt.Sprintf("Failed to fetch managed namespace : %s", err.Error()),
			resp), err
	}

	nameSpaces := globalCluster.ManagedNamespaces
	zoneMappings := globalCluster.CustomZoneMapping
	if len(nameSpaces) == 0 && len(*zoneMappings) == 0 {
		return nil, handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, errors.New("resource not found")
	}
	readModel := newModel(globalCluster, currentModel)
	return readModel, handler.ProgressEvent{}, nil
}

func newModel(globalCluster *admin20231115002.GeoSharding, currentModel *Model) *Model {
	readModel := new(Model)
	readModel.ProjectId = currentModel.ProjectId
	readModel.ClusterName = currentModel.ClusterName
	maps := flattenManagedNamespaces(globalCluster.ManagedNamespaces)
	readModel.CustomZoneMappings = currentModel.CustomZoneMappings
	readModel.ManagedNamespaces = maps
	readModel.Profile = currentModel.Profile
	readModel.RemoveAllZoneMapping = currentModel.RemoveAllZoneMapping
	return readModel
}

func flattenManagedNamespaces(managedNamespaces []admin20231115002.ManagedNamespaces) []ManagedNamespace {
	var results []ManagedNamespace
	for ind := range managedNamespaces {
		namespace := ManagedNamespace{
			Db:                     &managedNamespaces[ind].Db,
			Collection:             &managedNamespaces[ind].Collection,
			CustomShardKey:         &managedNamespaces[ind].CustomShardKey,
			IsCustomShardKeyHashed: managedNamespaces[ind].IsCustomShardKeyHashed,
			IsShardKeyUnique:       managedNamespaces[ind].IsShardKeyUnique,
		}
		results = append(results, namespace)
	}
	return results
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// No OP
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// No OP
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	if !isExist(client, currentModel) {
		return progressevent.GetFailedEventByCode("Resource Not Found", cloudformation.HandlerErrorCodeNotFound), nil
	}

	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName
	remove := currentModel.ManagedNamespaces
	removeFlag := currentModel.RemoveAllZoneMapping
	// check if remove all zone mapping flag is enabled or name space array (to be removed) exist
	if (len(remove) == 0) && (removeFlag != nil && !*currentModel.RemoveAllZoneMapping) {
		// nothing to be removed .so raising invalid request
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "request doest not contain any item to remove",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if len(remove) > 0 {
		removeManagedNamespaces(context.Background(), client, remove, projectID, clusterName)
	}
	if currentModel.RemoveAllZoneMapping != nil && *currentModel.RemoveAllZoneMapping {
		_, _, err := client.Atlas20231115002.GlobalClustersApi.DeleteAllCustomZoneMappings(context.Background(), projectID, clusterName).Execute()
		if err != nil {
			return progressevent.GetFailedEventByCode(fmt.Sprintf("Failed to remove custom zones : %s", err.Error()),
				cloudformation.HandlerErrorCodeInvalidRequest), nil
		}
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func isExist(client *util.MongoDBClient, currentModel *Model) bool {
	config, _, err := ReadConfig(client, currentModel)
	return err == nil && config != nil
}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func removeManagedNamespaces(ctx context.Context, client *util.MongoDBClient, remove []ManagedNamespace, projectID, clusterName string) {
	for _, m := range remove {
		addManagedNamespace := &admin20231115002.DeleteManagedNamespaceApiParams{
			Collection:  m.Collection,
			Db:          m.Db,
			ClusterName: clusterName,
			GroupId:     projectID,
		}

		_, _, err := client.Atlas20231115002.GlobalClustersApi.DeleteManagedNamespaceWithParams(ctx, addManagedNamespace).Execute()
		if err != nil {
			_, _ = logger.Warnf("error while removing namespace:%+v", err)
		}
	}
}

func newCustomZoneMappings(currentModel *Model) *admin20231115002.CustomZoneMappings {
	return &admin20231115002.CustomZoneMappings{
		CustomZoneMappings: modelToCustomZoneMappings(currentModel.CustomZoneMappings),
	}
}

func createManagedNamespaces(ctx context.Context, client *util.MongoDBClient, nameSpaces []ManagedNamespace, projectID, clusterName string) error {
	for _, mn := range nameSpaces {
		addManagedNamespace := &admin20231115002.ManagedNamespace{
			Collection:     mn.Collection,
			Db:             mn.Db,
			CustomShardKey: mn.CustomShardKey,
		}
		addManagedNamespace.IsCustomShardKeyHashed = mn.IsCustomShardKeyHashed
		addManagedNamespace.IsShardKeyUnique = mn.IsShardKeyUnique
		_, _, err := client.Atlas20231115002.GlobalClustersApi.CreateManagedNamespace(ctx, projectID, clusterName, addManagedNamespace).Execute()
		if err != nil {
			return err
		}
	}

	return nil
}

func modelToCustomZoneMappings(tfList []ZoneMapping) []admin20231115002.ZoneMapping {
	apiObjects := make([]admin20231115002.ZoneMapping, len(tfList))
	for i, tfMapRaw := range tfList {
		if util.IsStringPresent(tfMapRaw.Location) || util.IsStringPresent(tfMapRaw.Zone) {
			apiObjects[i] = admin20231115002.ZoneMapping{
				Location: *tfMapRaw.Location,
				Zone:     *tfMapRaw.Zone,
			}
		}
	}

	return apiObjects
}
