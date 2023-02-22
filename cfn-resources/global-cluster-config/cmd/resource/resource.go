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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/openlyinc/pointy"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

var RequiredFields = []string{constants.ClusterName, constants.ProjectID}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	// Create Atlas API Request Object
	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName

	// create namespaces
	nameSpaces := currentModel.ManagedNamespaces
	addManagedNamespaces(context.Background(), client, nameSpaces, projectID, clusterName)
	// add zone mappings
	customZoneMappings := modelToCustomZoneMappings(currentModel.CustomZoneMappings)

	// API call to create
	_, _, err := client.GlobalClusters.AddCustomZoneMappings(context.Background(), projectID, clusterName, &mongodbatlas.CustomZoneMappingsRequest{
		CustomZoneMappings: customZoneMappings,
	})
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
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	// Check if  already exist
	if !isExist(*client, currentModel) {
		return progressevents.GetFailedEventByCode("Resource Not Found", cloudformation.HandlerErrorCodeNotFound), nil
	}
	// method call to read configuration
	config, event, err := ReadConfig(*client, currentModel)
	if err != nil {
		_, _ = logger.Warnf("error reading MongoDB Global Cluster Configuration (%s): %v", *currentModel.ClusterName, err)
		return event, nil
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   config,
	}, nil
}
func ReadConfig(client mongodbatlas.Client, currentModel *Model) (*Model, handler.ProgressEvent, error) {
	projectID := *currentModel.ProjectId
	ClusterName := *currentModel.ClusterName

	// API call to read configuration
	globalCluster, resp, err := client.GlobalClusters.Get(context.Background(), projectID, ClusterName)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil, handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		return nil, progressevents.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			resp.Response), err
	}
	nameSpaces := globalCluster.ManagedNamespaces
	zoneMappings := globalCluster.CustomZoneMapping

	if len(nameSpaces) == 0 && len(zoneMappings) == 0 {
		return nil, handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, errors.New("resource not found")
	}
	readModel := newModel(globalCluster, currentModel)
	return readModel, handler.ProgressEvent{}, nil
}

func newModel(globalCluster *mongodbatlas.GlobalCluster, currentModel *Model) *Model {
	readModel := new(Model)
	readModel.ProjectId = currentModel.ProjectId
	readModel.ClusterName = currentModel.ClusterName
	maps := flattenManagedNamespaces(globalCluster.ManagedNamespaces)
	zones := customZoneToModelMappings(globalCluster.CustomZoneMapping)
	readModel.CustomZoneMappings = zones
	readModel.ManagedNamespaces = maps
	readModel.Profile = currentModel.Profile
	readModel.RemoveAllZoneMapping = currentModel.RemoveAllZoneMapping
	return readModel
}

func flattenManagedNamespaces(managedNamespaces []mongodbatlas.ManagedNamespace) []ManagedNamespace {
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
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create MongoDb Atlas Client using keys
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}
	// Check if  already exist
	if !isExist(*client, currentModel) {
		return progressevents.GetFailedEventByCode("Resource Not Found", cloudformation.HandlerErrorCodeNotFound), nil
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
		_, _, err := client.GlobalClusters.DeleteCustomZoneMappings(context.Background(), projectID, clusterName)
		if err != nil {
			return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to remove custom zones : %s", err.Error()),
				cloudformation.HandlerErrorCodeInvalidRequest), nil
		}
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func setup() {
	util.SetupLogger("mongodb-atlas-global-cluster-config")
}

func isExist(client mongodbatlas.Client, currentModel *Model) bool {
	config, _, err := ReadConfig(client, currentModel)
	return err == nil && config != nil
}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func removeManagedNamespaces(ctx context.Context, conn *mongodbatlas.Client, remove []ManagedNamespace, projectID, clusterName string) {
	for _, m := range remove {
		addManagedNamespace := &mongodbatlas.ManagedNamespace{
			Collection:     cast.ToString(m.Collection),
			Db:             cast.ToString(m.Db),
			CustomShardKey: cast.ToString(m.CustomShardKey),
		}
		addManagedNamespace.IsCustomShardKeyHashed = pointy.Bool(*m.IsCustomShardKeyHashed)
		addManagedNamespace.IsShardKeyUnique = pointy.Bool(*m.IsShardKeyUnique)
		_, _, err := conn.GlobalClusters.DeleteManagedNamespace(ctx, projectID, clusterName, addManagedNamespace)
		if err != nil {
			_, _ = logger.Warnf("error while removing namespace:%+v", err)
		}
	}
}

func addManagedNamespaces(ctx context.Context, client *mongodbatlas.Client, nameSpaces []ManagedNamespace, projectID, clusterName string) {
	for _, mn := range nameSpaces {
		addManagedNamespace := &mongodbatlas.ManagedNamespace{
			Collection:     cast.ToString(mn.Collection),
			Db:             cast.ToString(mn.Db),
			CustomShardKey: cast.ToString(mn.CustomShardKey),
		}
		addManagedNamespace.IsCustomShardKeyHashed = mn.IsCustomShardKeyHashed
		addManagedNamespace.IsShardKeyUnique = mn.IsShardKeyUnique
		_, _, err := client.GlobalClusters.AddManagedNamespace(ctx, projectID, clusterName, addManagedNamespace)
		if err != nil {
			_, _ = logger.Warnf("error while adding namespace:%+v", err)
		}
	}
}

func modelToCustomZoneMapping(tfMap ZoneMapping) *mongodbatlas.CustomZoneMapping {
	return &mongodbatlas.CustomZoneMapping{
		Location: cast.ToString(tfMap.Location),
		Zone:     cast.ToString(tfMap.Zone),
	}
}

func modelToCustomZoneMappings(tfList []ZoneMapping) []mongodbatlas.CustomZoneMapping {
	apiObjects := make([]mongodbatlas.CustomZoneMapping, len(tfList))
	for i, tfMapRaw := range tfList {
		if tfMapRaw == (ZoneMapping{}) || tfMapRaw.Location == nil || tfMapRaw.Zone == nil {
			continue
		}
		apiObject := modelToCustomZoneMapping(tfMapRaw)
		apiObjects[i] = *apiObject
	}
	return apiObjects
}
func customZoneToModelMapping(location string, zone string) *ZoneMapping {
	if location == "" {
		return nil
	}
	return &ZoneMapping{
		Location: &location,
		Zone:     &zone,
	}
}

func customZoneToModelMappings(tfList map[string]string) []ZoneMapping {
	apiObjects := make([]ZoneMapping, len(tfList))
	var i = 0
	for k, v := range tfList {
		fmt.Printf("key[%s] value[%s]\n", k, v)
		apiObject := customZoneToModelMapping(k, v)
		if apiObject == nil {
			continue
		}
		apiObjects[i] = *apiObject
		i++
	}
	return apiObjects
}
