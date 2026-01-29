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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20250312012/admin"
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

	maintenanceWindow, _ := get(client, *currentModel)
	if maintenanceWindow != nil {
		return progress_events.GetFailedEventByCode("resource already exists", string(types.HandlerErrorCodeAlreadyExists)), nil
	}

	// Handle Defer if requested
	if currentModel.Defer != nil && *currentModel.Defer {
		if pe := deferMaintenanceWindow(client, *currentModel.ProjectId); pe != nil {
			return *pe, nil
		}
	}

	atlasModel := currentModel.toAtlasModel()
	startASP := false
	atlasModel.StartASAP = &startASP

	resp, err := client.AtlasSDK.MaintenanceWindowsApi.UpdateMaintenanceWindow(context.Background(), *currentModel.ProjectId, &atlasModel).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(err.Error(), resp), nil
	}

	// Handle AutoDefer if requested
	if currentModel.AutoDefer != nil && *currentModel.AutoDefer {
		if pe := toggleAutoDefer(client, *currentModel.ProjectId); pe != nil {
			return *pe, nil
		}
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

	maintenanceWindow, errorProgressEvent := get(client, *currentModel)
	if errorProgressEvent != nil {
		return *errorProgressEvent, nil
	}

	currentModel.AutoDeferOnceEnabled = maintenanceWindow.AutoDeferOnceEnabled
	currentModel.DayOfWeek = &maintenanceWindow.DayOfWeek
	currentModel.HourOfDay = maintenanceWindow.HourOfDay
	currentModel.StartASAP = maintenanceWindow.StartASAP
	currentModel.NumberOfDeferrals = maintenanceWindow.NumberOfDeferrals
	currentModel.TimeZoneId = maintenanceWindow.TimeZoneId

	if maintenanceWindow.ProtectedHours != nil {
		currentModel.ProtectedHours = &ProtectedHours{
			StartHourOfDay: maintenanceWindow.ProtectedHours.StartHourOfDay,
			EndHourOfDay:   maintenanceWindow.ProtectedHours.EndHourOfDay,
		}
	}

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

	_, handlerError := get(client, *currentModel)
	if handlerError != nil {
		return *handlerError, nil
	}

	// Handle Defer if changed
	if currentModel.Defer != nil && *currentModel.Defer {
		if prevModel == nil || prevModel.Defer == nil || !*prevModel.Defer {
			if pe := deferMaintenanceWindow(client, *currentModel.ProjectId); pe != nil {
				return *pe, nil
			}
		}
	}

	atlasModel := currentModel.toAtlasModel()
	startASP := false
	atlasModel.StartASAP = &startASP

	resp, err := client.AtlasSDK.MaintenanceWindowsApi.UpdateMaintenanceWindow(context.Background(), *currentModel.ProjectId, &atlasModel).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(err.Error(), resp), nil
	}

	// Handle AutoDefer if changed
	if prevModel != nil {
		prevAutoDefer := prevModel.AutoDefer != nil && *prevModel.AutoDefer
		currAutoDefer := currentModel.AutoDefer != nil && *currentModel.AutoDefer
		if prevAutoDefer != currAutoDefer {
			if pe := toggleAutoDefer(client, *currentModel.ProjectId); pe != nil {
				return *pe, nil
			}
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
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

	_, handlerError := get(client, *currentModel)
	if handlerError != nil {
		return *handlerError, nil
	}

	resp, err := client.AtlasSDK.MaintenanceWindowsApi.ResetMaintenanceWindow(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "delete successful",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func (m Model) toAtlasModel() admin.GroupMaintenanceWindow {
	return admin.GroupMaintenanceWindow{
		DayOfWeek:            *m.DayOfWeek,
		HourOfDay:            m.HourOfDay,
		StartASAP:            m.StartASAP,
		AutoDeferOnceEnabled: m.AutoDeferOnceEnabled,
		ProtectedHours:       m.toProtectedHours(),
	}
}

func (m Model) toProtectedHours() *admin.ProtectedHours {
	if m.ProtectedHours == nil {
		return nil
	}
	return &admin.ProtectedHours{
		StartHourOfDay: m.ProtectedHours.StartHourOfDay,
		EndHourOfDay:   m.ProtectedHours.EndHourOfDay,
	}
}

func get(client *util.MongoDBClient, currentModel Model) (*admin.GroupMaintenanceWindow, *handler.ProgressEvent) {
	maintenanceWindow, resp, err := client.AtlasSDK.MaintenanceWindowsApi.GetMaintenanceWindow(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		_, _ = logger.Warnf("Read - error: %+v", err)
		ev := progress_events.GetFailedEventByResponse(err.Error(), resp)
		return nil, &ev
	}

	if isResponseEmpty(maintenanceWindow) {
		_, _ = logger.Warnf("Read - resource is empty: %+v", err)
		ev := progress_events.GetFailedEventByCode("resource not found", string(types.HandlerErrorCodeNotFound))
		return nil, &ev
	}

	return maintenanceWindow, nil
}

func isResponseEmpty(maintenanceWindow *admin.GroupMaintenanceWindow) bool {
	return maintenanceWindow != nil && maintenanceWindow.GetDayOfWeek() == 0
}

func deferMaintenanceWindow(client *util.MongoDBClient, projectID string) *handler.ProgressEvent {
	_, err := client.AtlasSDK.MaintenanceWindowsApi.DeferMaintenanceWindow(context.Background(), projectID).Execute()
	if err != nil {
		return &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: "GeneralServiceException",
		}
	}
	return nil
}

func toggleAutoDefer(client *util.MongoDBClient, projectID string) *handler.ProgressEvent {
	_, err := client.AtlasSDK.MaintenanceWindowsApi.ToggleMaintenanceAutoDefer(context.Background(), projectID).Execute()
	if err != nil {
		return &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: "GeneralServiceException",
		}
	}
	return nil
}
