// Copyright 2024 MongoDB Inc
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
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func setup() {
	util.SetupLogger("mongodb-atlas-stream-instance")
}

var CreateRequiredFields = []string{constants.Name, constants.GroupID, constants.StreamConfig, constants.DataProcessRegion}
var ReadRequiredFields = []string{constants.Name, constants.GroupID}
var UpdateRequiredFields = []string{constants.Name, constants.GroupID, constants.DataProcessRegion}
var DeleteRequiredFields = []string{constants.Name, constants.GroupID}
var ListRequiredFields = []string{constants.GroupID}

const Kafka = "Kafka"
const Cluster = "Cluster"

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	streamInstanceCreateReq := newStreamsTenant(currentModel)

	atlasV2 := client.AtlasSDK

	createdStreamInstance, resp, err := atlasV2.StreamsApi.CreateStreamInstance(context.Background(), *currentModel.GroupId, streamInstanceCreateReq).Execute()
	if err != nil {
		return handleError(resp, constants.CREATE, err)
	}

	currentModel.Id = createdStreamInstance.Id

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	atlasV2 := client.AtlasSDK

	streamInstance, resp, err := atlasV2.StreamsApi.GetStreamInstance(context.Background(), *currentModel.GroupId, *currentModel.Name).Execute()
	if err != nil {
		return handleError(resp, constants.READ, err)
	}

	currentModel.Id = streamInstance.Id
	currentModel.Hostnames = streamInstance.GetHostnames()
	currentModel.DataProcessRegion = newModelDataRegion(streamInstance.DataProcessRegion)
	currentModel.StreamConfig = newModelStreamConfig(streamInstance.StreamConfig)
	currentModel.Connections = newModelConnections(streamInstance.GetConnections())

	return handler.ProgressEvent{
		OperationStatus: cloudformation.OperationStatusSuccess,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	updateRequest := &admin.StreamsDataProcessRegion{
		CloudProvider: *currentModel.DataProcessRegion.CloudProvider,
		Region:        *currentModel.DataProcessRegion.Region,
	}

	atlasV2 := client.AtlasSDK

	_, resp, err := atlasV2.StreamsApi.UpdateStreamInstance(context.Background(), *currentModel.GroupId, *currentModel.Name, updateRequest).Execute()
	if err != nil {
		return handleError(resp, constants.UPDATE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update success",
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	atlasV2 := client.AtlasSDK

	_, resp, err := atlasV2.StreamsApi.DeleteStreamInstance(context.Background(), *currentModel.GroupId, *currentModel.Name).Execute()
	if err != nil {
		return handleError(resp, constants.DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete success",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	atlasV2 := client.AtlasSDK

	accumulatedStreamInstances := make([]admin.StreamsTenant, 0)
	for ok := true; ok; {
		streamInstances, resp, err := atlasV2.StreamsApi.ListStreamInstances(context.Background(), *currentModel.GroupId).Execute() //TODO: increase pagenum
		if err != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}
		accumulatedStreamInstances = append(accumulatedStreamInstances, streamInstances.GetResults()...)
		ok = streamInstances.GetTotalCount() > len(accumulatedStreamInstances)
	}

	response := make([]any, 0, len(accumulatedStreamInstances))
	for i := range accumulatedStreamInstances {
		response = append(response, accumulatedStreamInstances[i])
	}

	return handler.ProgressEvent{
		OperationStatus: cloudformation.OperationStatusSuccess,
		Message:         "List Complete",
		ResourceModel:   response,
	}, nil
}

func newStreamsTenant(model *Model) *admin.StreamsTenant {
	return &admin.StreamsTenant{
		Name:    model.Name,
		GroupId: model.GroupId,
		DataProcessRegion: &admin.StreamsDataProcessRegion{
			CloudProvider: *model.DataProcessRegion.CloudProvider,
			Region:        *model.DataProcessRegion.Region,
		},
		StreamConfig: &admin.StreamConfig{
			Tier: model.StreamConfig.Tier,
		},
	}
}

func newModelDataRegion(dataProcessRegion *admin.StreamsDataProcessRegion) *StreamsDataProcessRegion {
	return &StreamsDataProcessRegion{
		CloudProvider: &dataProcessRegion.CloudProvider,
		Region:        &dataProcessRegion.Region,
	}
}

func newModelStreamConfig(streamConfig *admin.StreamConfig) *StreamConfig {
	return &StreamConfig{
		Tier: streamConfig.Tier,
	}
}

func newModelDBRoleToExecute(dbRole *admin.DBRoleToExecute) *DBRoleToExecute {
	return &DBRoleToExecute{
		Role: dbRole.Role,
		Type: dbRole.Type,
	}
}

func newModelAuthentication(authentication *admin.StreamsKafkaAuthentication) *StreamsKafkaAuthentication {
	return &StreamsKafkaAuthentication{
		Mechanism: authentication.Mechanism,
		Password:  authentication.Password,
		Username:  authentication.Username,
	}
}

func newModelSecurity(security *admin.StreamsKafkaSecurity) *StreamsKafkaSecurity {
	return &StreamsKafkaSecurity{
		BrokerPublicCertificate: security.BrokerPublicCertificate,
		Protocol:                security.Protocol,
	}
}

func newModelConnections(streamConfig []admin.StreamsConnection) []StreamsConnection {
	if len(streamConfig) == 0 {
		return nil
	}

	connections := make([]StreamsConnection, 0)
	for _, connection := range streamConfig {
		modelConnection := StreamsConnection{
			Name: connection.Name,
			Type: connection.Type,
		}
		if connection.GetType() == Cluster {
			modelConnection.ClusterName = connection.ClusterName
			modelConnection.DbRoleToExecute = newModelDBRoleToExecute(connection.DbRoleToExecute)
		} else if connection.GetType() == Kafka {
			modelConnection.BootstrapServers = connection.BootstrapServers
			modelConnection.Authentication = newModelAuthentication(connection.Authentication)
			modelConnection.Security = newModelSecurity(connection.Security)
		}
		connections = append(connections, modelConnection)
	}
	return connections
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())
	_, _ = logger.Warn(errMsg)
	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}
	if response.StatusCode == http.StatusBadRequest {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}
