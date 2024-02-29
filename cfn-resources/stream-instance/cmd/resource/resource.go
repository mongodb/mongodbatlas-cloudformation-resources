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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func setup() {
	util.SetupLogger("mongodb-atlas-stream-instance")
}

var CreateRequiredFields = []string{constants.InstanceName, constants.ProjectID, constants.DataProcessRegion}
var ReadRequiredFields = []string{constants.InstanceName, constants.ProjectID}
var UpdateRequiredFields = []string{constants.InstanceName, constants.ProjectID, constants.DataProcessRegion}
var DeleteRequiredFields = []string{constants.InstanceName, constants.ProjectID}
var ListRequiredFields = []string{constants.ProjectID}

const Kafka = "Kafka"
const Cluster = "Cluster"

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	streamInstanceCreateReq := NewStreamsTenant(currentModel)

	atlasV2 := client.AtlasSDK

	createdStreamInstance, resp, err := atlasV2.StreamsApi.CreateStreamInstance(context.Background(), *currentModel.ProjectId, streamInstanceCreateReq).Execute()
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

	client, handlerError := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	atlasV2 := client.AtlasSDK

	streamInstance, resp, err := atlasV2.StreamsApi.GetStreamInstance(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName).Execute()
	if err != nil {
		return handleError(resp, constants.READ, err)
	}

	currentModel.Id = streamInstance.Id
	currentModel.Hostnames = streamInstance.GetHostnames()
	currentModel.DataProcessRegion = NewModelDataRegion(streamInstance.DataProcessRegion)
	currentModel.StreamConfig = NewModelStreamConfig(streamInstance.StreamConfig)
	currentModel.Connections = NewModelConnections(streamInstance.Connections)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
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

	client, handlerError := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	updateRequest := &admin.StreamsDataProcessRegion{
		CloudProvider: *currentModel.DataProcessRegion.CloudProvider,
		Region:        *currentModel.DataProcessRegion.Region,
	}

	atlasV2 := client.AtlasSDK

	_, resp, err := atlasV2.StreamsApi.UpdateStreamInstance(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, updateRequest).Execute()
	if err != nil {
		return handleError(resp, constants.UPDATE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update success",
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	atlasV2 := client.AtlasSDK

	_, resp, err := atlasV2.StreamsApi.DeleteStreamInstance(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName).Execute()
	if err != nil {
		return handleError(resp, constants.DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete success"}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	atlasV2 := client.AtlasSDK

	accumulatedStreamInstances := make([]admin.StreamsTenant, 0)
	pageNum := 1
	for ok := true; ok; {
		listStreamInstancesRequest := atlasV2.StreamsApi.ListStreamInstances(context.Background(), *currentModel.ProjectId)
		listStreamInstancesRequest.PageNum(pageNum)
		listStreamInstancesRequest.ItemsPerPage(100)
		streamInstances, resp, err := listStreamInstancesRequest.Execute()
		if err != nil {
			return handleError(resp, constants.LIST, err)
		}
		accumulatedStreamInstances = append(accumulatedStreamInstances, streamInstances.GetResults()...)
		ok = streamInstances.GetTotalCount() > len(accumulatedStreamInstances)
		pageNum++
	}

	response := make([]interface{}, 0)
	for _, stream := range accumulatedStreamInstances {
		cloudProvider := stream.DataProcessRegion.CloudProvider
		region := stream.DataProcessRegion.Region
		model := Model{
			InstanceName: stream.Name,
			DataProcessRegion: &StreamsDataProcessRegion{
				CloudProvider: &cloudProvider,
				Region:        &region,
			},
			StreamConfig: &StreamConfig{
				Tier: stream.StreamConfig.Tier,
			},
			ProjectId:   stream.GroupId,
			Id:          stream.Id,
			Hostnames:   *stream.Hostnames,
			Profile:     currentModel.Profile,
			Connections: NewModelConnections(stream.Connections),
		}
		response = append(response, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  response,
	}, nil
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
	return progressevent.GetFailedEventByResponse(errMsg, response), nil
}
