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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20240805001/admin"
)

const (
	CallBackSeconds = 30
)

var CreateRequiredFields = []string{constants.ProjID, constants.Name}
var ReadRequiredFields = []string{constants.ProjID, constants.Name}
var UpdateRequiredFields = []string{constants.ProjID, constants.Name}
var DeleteRequiredFields = []string{constants.ProjID, constants.Name}
var ListRequiredFields = []string{constants.ProjID}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-ServerlessInstance")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	// Create atlas client
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Callback
	if stateName, ok := req.CallbackContext[constants.StateName]; ok {
		_, _ = log.Debugf("Callback state: %s", stateName)
		return serverlessCallback(client, currentModel, constants.IdleState)
	}

	serverlessInstanceRequest := &admin.ServerlessInstanceDescriptionCreate{
		Name:                         *currentModel.Name,
		ProviderSettings:             *setProviderSettings(currentModel),
		ServerlessBackupOptions:      setBackupOptions(currentModel),
		TerminationProtectionEnabled: currentModel.TerminationProtectionEnabled,
	}

	serverless, res, err := client.Atlas20231115002.ServerlessInstancesApi.CreateServerlessInstance(context.Background(), *currentModel.ProjectID, serverlessInstanceRequest).Execute()
	if err != nil {
		if apiError, ok := admin.AsError(err); ok && *apiError.Error == http.StatusBadRequest && strings.Contains(*apiError.ErrorCode, constants.Duplicate) {
			_, _ = log.Debugf("Serverless - Create() - error 400: %+v", err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
		}

		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create ServerlessInstance `%s`", *serverless.StateName),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.StateName: serverless.StateName,
		},
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	cluster, res, err := client.Atlas20231115002.ServerlessInstancesApi.GetServerlessInstance(context.Background(), *currentModel.ProjectID, *currentModel.Name).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}
	// Read Instance
	model := readServerlessInstance(cluster, currentModel.Profile)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   model,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Callback
	if _, ok := req.CallbackContext[constants.StateName]; ok {
		return serverlessCallback(client, currentModel, constants.IdleState)
	}

	// CFN TEST : currently Update is throwing 500 Error instead of 404 if resource not exists
	_, res, err := client.Atlas20231115002.ServerlessInstancesApi.GetServerlessInstance(context.Background(), *currentModel.ProjectID, *currentModel.Name).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	serverlessInstanceRequest := &admin.UpdateServerlessInstanceApiParams{
		GroupId: *currentModel.ProjectID,
		ServerlessInstanceDescriptionUpdate: &admin.ServerlessInstanceDescriptionUpdate{
			TerminationProtectionEnabled: currentModel.TerminationProtectionEnabled,
			ServerlessBackupOptions: &admin.ClusterServerlessBackupOptions{
				ServerlessContinuousBackupEnabled: currentModel.ContinuousBackupEnabled,
			},
		},
		Name: *currentModel.Name,
	}

	serverless, res, err := client.Atlas20231115002.ServerlessInstancesApi.UpdateServerlessInstanceWithParams(context.Background(), serverlessInstanceRequest).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              fmt.Sprintf("Create ServerlessInstance `%s`", *serverless.StateName),
		ResourceModel:        currentModel,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.StateName: serverless.StateName,
			constants.ID:        serverless.Id,
		},
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if _, ok := req.CallbackContext[constants.StateName]; ok {
		return serverlessCallback(client, currentModel, constants.DeletedState)
	}

	_, res, err := client.Atlas20231115002.ServerlessInstancesApi.DeleteServerlessInstance(context.Background(), *currentModel.ProjectID, *currentModel.Name).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              "Deleting ServerlessInstance",
		ResourceModel:        currentModel,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.StateName: constants.DeletingState,
		},
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	listOptions := &admin.ListServerlessInstancesApiParams{
		GroupId:      *currentModel.ProjectID,
		PageNum:      admin.PtrInt(0),
		ItemsPerPage: admin.PtrInt(1000),
	}
	clustersResp, res, err := client.Atlas20231115002.ServerlessInstancesApi.ListServerlessInstancesWithParams(context.Background(), listOptions).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	instances := []interface{}{} // cfn test needs empty array instead nil, when items entries found
	for i := range clustersResp.Results {
		cluster := readServerlessInstance(&clustersResp.Results[i], currentModel.Profile)
		instances = append(instances, cluster)
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   instances,
	}, nil
}

func setBackupOptions(currentModel *Model) (serverlessBackupOptions *admin.ClusterServerlessBackupOptions) {
	if currentModel.ContinuousBackupEnabled == nil {
		return nil
	}
	serverlessBackupOptions = &admin.ClusterServerlessBackupOptions{
		ServerlessContinuousBackupEnabled: currentModel.ContinuousBackupEnabled,
	}
	return serverlessBackupOptions
}

func setProviderSettings(currentModel *Model) (serverlessProviderSettings *admin.ServerlessProviderSettings) {
	if currentModel.ProviderSettings == nil {
		return &admin.ServerlessProviderSettings{
			ProviderName:        admin.PtrString(constants.Serverless),
			BackingProviderName: constants.AWS,
		}
	}

	serverlessProviderSettings = &admin.ServerlessProviderSettings{
		BackingProviderName: constants.AWS,
	}

	if currentModel.ProviderSettings.ProviderName != nil {
		serverlessProviderSettings.ProviderName = currentModel.ProviderSettings.ProviderName
	}
	if currentModel.ProviderSettings.RegionName != nil {
		serverlessProviderSettings.RegionName = *currentModel.ProviderSettings.RegionName
	}
	return serverlessProviderSettings
}

func readServerlessInstance(cluster *admin.ServerlessInstanceDescription, profile *string) (serverless *Model) {
	serverless = &Model{}
	serverless.Name = cluster.Name
	serverless.Id = cluster.Id
	serverless.ProjectID = cluster.GroupId
	serverless.Profile = profile

	serverless.ProviderSettings = &ServerlessInstanceProviderSettings{
		ProviderName: cluster.ProviderSettings.ProviderName,
		RegionName:   &cluster.ProviderSettings.RegionName,
	}

	if cluster.ServerlessBackupOptions != nil {
		serverless.ContinuousBackupEnabled = cluster.ServerlessBackupOptions.ServerlessContinuousBackupEnabled
	}

	if cluster.ConnectionStrings != nil {
		serverless.ConnectionStrings = &ServerlessInstanceConnectionStrings{
			StandardSrv:     cluster.ConnectionStrings.StandardSrv,
			PrivateEndpoint: readPrivateEndpoint(cluster.ConnectionStrings.PrivateEndpoint),
		}
	}
	serverless.CreateDate = util.TimePtrToStringPtr(cluster.CreateDate)
	serverless.MongoDBVersion = cluster.MongoDBVersion
	serverless.TerminationProtectionEnabled = cluster.TerminationProtectionEnabled
	serverless.StateName = cluster.StateName
	return serverless
}

func readPrivateEndpoint(privateEPs []admin.ServerlessConnectionStringsPrivateEndpointList) (endPoints []ServerlessInstancePrivateEndpoint) {
	for i := range privateEPs {
		var pep = ServerlessInstancePrivateEndpoint{}
		pep.Endpoints = readPrivateEndpointEndpoints(privateEPs[i].Endpoints)
		pep.Type = privateEPs[i].Type
		pep.SrvConnectionString = privateEPs[i].SrvConnectionString
		endPoints = append(endPoints, pep)
	}
	return
}

func readPrivateEndpointEndpoints(peEndpoints []admin.ServerlessConnectionStringsPrivateEndpointItem) (epEndpoints []ServerlessInstancePrivateEndpointEndpoint) {
	for i := range peEndpoints {
		epEndpoints = append(epEndpoints, ServerlessInstancePrivateEndpointEndpoint{
			EndpointId:   peEndpoints[i].EndpointId,
			ProviderName: peEndpoints[i].ProviderName,
			Region:       peEndpoints[i].Region,
		})
	}
	return
}

func serverlessCallback(client *util.MongoDBClient, currentModel *Model, targtStatus string) (progressEvent handler.ProgressEvent, err error) {
	serverless, resp, err := client.Atlas20231115002.ServerlessInstancesApi.GetServerlessInstance(context.Background(), *currentModel.ProjectID, *currentModel.Name).Execute()
	if err != nil {
		if apiError, ok := admin.AsError(err); ok && *apiError.Error == http.StatusNotFound {
			_, _ = log.Debugf("404: No instance found")
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Deleted ServerlessInstance",
				ResourceModel:   nil,
			}, nil
		}
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.Id = serverless.Id
	if *serverless.StateName != constants.IdleState {
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              fmt.Sprintf("Create ServerlessInstance `%s`", *serverless.StateName),
			ResourceModel:        currentModel,
			CallbackDelaySeconds: CallBackSeconds,
			CallbackContext: map[string]interface{}{
				constants.StateName: serverless.StateName,
			},
		}, nil
	}

	model := readServerlessInstance(serverless, currentModel.Profile)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         fmt.Sprintf("Create ServerlessInstance `%s`", *serverless.StateName),
		ResourceModel:   model,
	}, nil
}
