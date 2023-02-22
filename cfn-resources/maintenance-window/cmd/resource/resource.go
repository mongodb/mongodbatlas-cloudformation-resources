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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

var RequiredFields = []string{constants.GroupID}

func setup() {
	util.SetupLogger("mongodb-atlas-maintenance-window")
}

func (m Model) toAtlasModel() mongodbatlas.MaintenanceWindow {
	return mongodbatlas.MaintenanceWindow{
		DayOfWeek:            *m.DayOfWeek,
		HourOfDay:            m.HourOfDay,
		StartASAP:            m.StartASAP,
		AutoDeferOnceEnabled: m.AutoDeferOnceEnabled,
	}
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	// Validation
	if errEvent := validator.ValidateModel(RequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}
	// Create atlas client
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}
	var res *mongodbatlas.Response

	atlasModel := currentModel.toAtlasModel()
	startASP := false
	atlasModel.StartASAP = &startASP

	res, err := client.MaintenanceWindows.Update(context.Background(), *currentModel.GroupId, &atlasModel)
	if err != nil {
		_, _ = logger.Warnf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   *currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Validation
	if errEvent := validator.ValidateModel(RequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}
	// Create atlas client
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	maintenanceWindow, errorProgressEvent := get(client, *currentModel)
	if errorProgressEvent != nil {
		return *errorProgressEvent, nil
	}

	currentModel.AutoDeferOnceEnabled = maintenanceWindow.AutoDeferOnceEnabled
	currentModel.DayOfWeek = &maintenanceWindow.DayOfWeek
	currentModel.HourOfDay = maintenanceWindow.HourOfDay
	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}

func get(client *mongodbatlas.Client, currentModel Model) (*mongodbatlas.MaintenanceWindow, *handler.ProgressEvent) {
	maintenanceWindow, res, err := client.MaintenanceWindows.Get(context.Background(), *currentModel.GroupId)
	if err != nil {
		_, _ = logger.Warnf("Read - error: %+v", err)
		ev := progress_events.GetFailedEventByResponse(err.Error(), res.Response)
		return nil, &ev
	}

	if isResponseEmpty(maintenanceWindow) {
		_, _ = logger.Warnf("Read - resource is empty: %+v", err)
		ev := progress_events.GetFailedEventByCode("resource not found", cloudformation.HandlerErrorCodeNotFound)
		return nil, &ev
	}

	return maintenanceWindow, nil
}

func isResponseEmpty(maintenanceWindow *mongodbatlas.MaintenanceWindow) bool {
	return maintenanceWindow != nil && maintenanceWindow.DayOfWeek == 0
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Validation
	if errEvent := validator.ValidateModel(RequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}
	// Create atlas client
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	_, handlerError := get(client, *currentModel)
	if handlerError != nil {
		return *handlerError, nil
	}

	var res *mongodbatlas.Response

	atlasModel := currentModel.toAtlasModel()
	startASP := false
	atlasModel.StartASAP = &startASP

	res, err := client.MaintenanceWindows.Update(context.Background(), *currentModel.GroupId, &atlasModel)
	if err != nil {
		_, _ = logger.Warnf("Update - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Validation
	if errEvent := validator.ValidateModel(RequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}
	// Create atlas client
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	_, handlerError := get(client, *currentModel)
	if handlerError != nil {
		return *handlerError, nil
	}

	var res *mongodbatlas.Response
	res, err := client.MaintenanceWindows.Reset(context.Background(), *currentModel.GroupId)

	if err != nil {
		_, _ = logger.Warnf("Delete - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Response
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "delete successful",
	}
	return event, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}
