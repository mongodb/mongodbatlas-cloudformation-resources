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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	localconstants "github.com/mongodb/mongodbatlas-cloudformation-resources/cloud-outage-simulation/cmd/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	atlasSDK "go.mongodb.org/atlas-sdk/v20230201002/admin"

	"log"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
)

var RequiredFields = []string{constants.ClusterName, constants.ProjectID}
var SimulationStatus = []string{localconstants.Simulating, localconstants.Starting, localconstants.StartingRequested}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	if req.CallbackContext != nil {
		return validateProgress(atlas, currentModel, "SIMULATING")
	}

	clusterName := cast.ToString(currentModel.ClusterName)
	projectID := cast.ToString(currentModel.ProjectId)

	// check if already exist in active state
	active, _, _ := isActive(atlas, projectID, clusterName, "nil")
	if active {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Already Exist",
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}
	requestBody := atlasSDK.ClusterOutageSimulation{
		OutageFilters: newOutageFilters(currentModel),
	}
	_, _ = logger.Debugf("requestBody - : %+v", requestBody)
	atlas.AtlasV2.ClusterOutageSimulationApi.StartOutageSimulation(context.Background(), projectID, clusterName, &requestBody)
	simulationObject, res, err := atlas.AtlasV2.ClusterOutageSimulationApi.StartOutageSimulation(context.Background(), projectID, clusterName, &requestBody).Execute()
	if err != nil {
		_, _ = logger.Warnf("create Outage - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res), nil
	}
	_, _ = logger.Debugf("currentModel - error: %+v", currentModel)

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("outage simulation status : %s", *simulationObject.State),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 65,
		CallbackContext: map[string]interface{}{
			"status":       simulationObject.State,
			"cluster_name": simulationObject.ClusterName,
			"project_id":   simulationObject.GroupId,
		},
	}, nil
}

func newOutageFilters(currentModel *Model) []atlasSDK.AtlasClusterOutageSimulationOutageFilter {
	outageFilters := make([]atlasSDK.AtlasClusterOutageSimulationOutageFilter, 0)

	for ind := range currentModel.OutageFilters {
		mMatcher := atlasSDK.AtlasClusterOutageSimulationOutageFilter{
			CloudProvider: currentModel.OutageFilters[ind].CloudProvider,
			RegionName:    currentModel.OutageFilters[ind].Region,
			Type:          currentModel.OutageFilters[ind].Type,
		}
		outageFilters = append(outageFilters, mMatcher)
	}
	return outageFilters
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
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

	clusterName := cast.ToString(currentModel.ClusterName)
	projectID := cast.ToString(currentModel.ProjectId)
	// API call to read resource
	outageSimulation, resp, err := atlas.AtlasV2.ClusterOutageSimulationApi.GetOutageSimulation(context.Background(), projectID, clusterName).Execute()
	if err != nil || outageSimulation == nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp), nil
	}
	// check if simulation is in active state
	if !contains(SimulationStatus, *outageSimulation.State) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	convertToUIModel(*outageSimulation, currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup
	// Validate required fields in the request
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
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

	if req.CallbackContext != nil {
		return validateProgress(atlas, currentModel, "COMPLETE")
	}

	clusterName := cast.ToString(currentModel.ClusterName)
	projectID := cast.ToString(currentModel.ProjectId)

	active, _, _ := isActive(atlas, projectID, clusterName, "nil")
	if !active {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	simulationObject, res, err := atlas.AtlasV2.ClusterOutageSimulationApi.EndOutageSimulation(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		_, _ = logger.Warnf("Delete - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res), nil
	}

	log.Println("[INFO] Waiting for MongoDB Cluster Outage Simulation to end")
	// progress callback setup
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              constants.DeleteInProgress,
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 65,
		CallbackContext: map[string]interface{}{
			"status":       simulationObject.State,
			"cluster_name": simulationObject.ClusterName,
			"project_id":   simulationObject.GroupId,
		},
	}, nil
}

// function to track snapshot creation status
func validateProgress(client *util.MongoDBClient, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName
	isReady, state, err := isCompleted(client, projectID, clusterName, targetState)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	if !isReady {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = handler.InProgress
		p.CallbackDelaySeconds = 65
		p.Message = "Pending"
		p.CallbackContext = map[string]interface{}{
			"status":       state,
			"cluster_name": *currentModel.ClusterName,
			"project_id":   *currentModel.ProjectId,
		}
		return p, nil
	}
	p := handler.NewProgressEvent()
	if targetState != localconstants.Complete {
		p.ResourceModel = currentModel
	} else {
		p.ResourceModel = nil
	}
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	return p, nil
}

// function to check if resource action is completed
func isCompleted(client *util.MongoDBClient, projectID, clusterName, targetState string) (isExist bool, status string, err error) {
	outageSimulation, resp, err := client.AtlasV2.ClusterOutageSimulationApi.GetOutageSimulation(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if resp.StatusCode == 404 {
			return true, "COMPLETED", nil
		}
		return false, "", err
	}
	if *outageSimulation.State != "" {
		log.Printf("[DEBUG] status for MongoDB cluster outage simulation: %s: %s", clusterName, *outageSimulation.State)
	}
	return *outageSimulation.State == targetState, *outageSimulation.State, nil
}

// function to check if resource action is active
func isActive(client *util.MongoDBClient, projectID, clusterName, targetState string) (isExist bool, status string, err error) {
	outageSimulation, resp, err := client.AtlasV2.ClusterOutageSimulationApi.GetOutageSimulation(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if resp.StatusCode == 404 {
			return false, "COMPLETED", nil
		}
		return false, "", err
	}
	if !contains(SimulationStatus, *outageSimulation.State) {
		return false, "COMPLETED", nil
	}
	return true, "COMPLETED", nil
}

// logger setup function
func setup() {
	util.SetupLogger("mongodb-atlas-backup-snapshot")
}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}
func convertToUIModel(outageSimulation atlasSDK.ClusterOutageSimulation, currentModel *Model) *Model {
	currentModel.SimulationId = outageSimulation.Id
	if outageSimulation.StartRequestDate != nil {
		dateStr := outageSimulation.StartRequestDate.String()
		currentModel.StartRequestDate = &dateStr
	}
	if *outageSimulation.State != "" {
		currentModel.State = outageSimulation.State
	}
	currentModel.OutageFilters = convertOutageFiltersToModel(outageSimulation.OutageFilters)
	return currentModel
}

func convertOutageFiltersToModel(outageFilters []atlasSDK.AtlasClusterOutageSimulationOutageFilter) []Filter {
	outageFilterList := make([]Filter, 0)
	for ind := range outageFilters {
		outageFilterList = append(outageFilterList, Filter{
			CloudProvider: outageFilters[ind].CloudProvider,
			Region:        outageFilters[ind].RegionName,
			Type:          outageFilters[ind].Type,
		})
	}
	return outageFilterList
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}
