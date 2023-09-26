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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

var RequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-maintenance-window")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if err := validator.ValidateModel(RequiredFields, currentModel); err != nil {
		_, _ = logger.Warnf("Validation Error")
		return *err, nil
	}
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	maintenanceWindow, _ := get2(client, *currentModel)
	if maintenanceWindow != nil {
		return progress_events.GetFailedEventByCode("resource already exists", cloudformation.HandlerErrorCodeAlreadyExists), nil
	}

	atlasModel := currentModel.toAtlasModel2()
	startASP := false
	atlasModel.StartASAP = &startASP

	_, resp, err := client.AtlasV2.MaintenanceWindowsApi.UpdateMaintenanceWindow(context.Background(), *currentModel.ProjectId, &atlasModel).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   *currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if err := validator.ValidateModel(RequiredFields, currentModel); err != nil {
		_, _ = logger.Warnf("Validation Error")
		return *err, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	maintenanceWindow, errorProgressEvent := get2(client, *currentModel)
	if errorProgressEvent != nil {
		return *errorProgressEvent, nil
	}

	currentModel.AutoDeferOnceEnabled = maintenanceWindow.AutoDeferOnceEnabled
	currentModel.DayOfWeek = &maintenanceWindow.DayOfWeek
	currentModel.HourOfDay = &maintenanceWindow.HourOfDay

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if err := validator.ValidateModel(RequiredFields, currentModel); err != nil {
		_, _ = logger.Warnf("Validation Error")
		return *err, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	_, handlerError := get2(client, *currentModel)
	if handlerError != nil {
		return *handlerError, nil
	}

	atlasModel := currentModel.toAtlasModel2()
	startASP := false
	atlasModel.StartASAP = &startASP

	_, resp, err := client.AtlasV2.MaintenanceWindowsApi.UpdateMaintenanceWindow(context.Background(), *currentModel.ProjectId, &atlasModel).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Validation
	if errEvent := validator.ValidateModel(RequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
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
	res, err := client.MaintenanceWindows.Reset(context.Background(), *currentModel.ProjectId)

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

func (m Model) toAtlasModel() mongodbatlas.MaintenanceWindow {
	return mongodbatlas.MaintenanceWindow{
		DayOfWeek:            *m.DayOfWeek,
		HourOfDay:            m.HourOfDay,
		StartASAP:            m.StartASAP,
		AutoDeferOnceEnabled: m.AutoDeferOnceEnabled,
	}
}

func (m Model) toAtlasModel2() admin.GroupMaintenanceWindow {
	return admin.GroupMaintenanceWindow{
		DayOfWeek:            *m.DayOfWeek,
		HourOfDay:            *m.HourOfDay,
		StartASAP:            m.StartASAP,
		AutoDeferOnceEnabled: m.AutoDeferOnceEnabled,
	}
}

func get(client *mongodbatlas.Client, currentModel Model) (*mongodbatlas.MaintenanceWindow, *handler.ProgressEvent) {
	maintenanceWindow, res, err := client.MaintenanceWindows.Get(context.Background(), *currentModel.ProjectId)
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

func get2(client *util.MongoDBClient, currentModel Model) (*admin.GroupMaintenanceWindow, *handler.ProgressEvent) {
	maintenanceWindow, resp, err := client.AtlasV2.MaintenanceWindowsApi.GetMaintenanceWindow(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		_, _ = logger.Warnf("Read - error: %+v", err)
		ev := progress_events.GetFailedEventByResponse(err.Error(), resp)
		return nil, &ev
	}

	if isResponseEmpty2(maintenanceWindow) {
		_, _ = logger.Warnf("Read - resource is empty: %+v", err)
		ev := progress_events.GetFailedEventByCode("resource not found", cloudformation.HandlerErrorCodeNotFound)
		return nil, &ev
	}

	return maintenanceWindow, nil
}

func isResponseEmpty(maintenanceWindow *mongodbatlas.MaintenanceWindow) bool {
	return maintenanceWindow != nil && maintenanceWindow.DayOfWeek == 0
}

func isResponseEmpty2(maintenanceWindow *admin.GroupMaintenanceWindow) bool {
	return maintenanceWindow != nil && maintenanceWindow.DayOfWeek == 0
}
