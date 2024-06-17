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
	"reflect"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20231115014/admin"
)

var CreateRequiredFields = []string{constants.EventTypeName, constants.ProjectID}
var RequiredFields = []string{constants.ID, constants.ProjectID}

func validateRequest(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-alert-configuration")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	validationError := validateRequest(CreateRequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.AtlasSDK

	if currentModel.Id != nil && *currentModel.Id != "" {
		_, _ = logger.Warnf("resource already exists for Id: %s", *currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Already Exists",
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}

	notifications, err := expandAlertConfigurationNotification(currentModel.Notifications)
	if err != nil {
		return progressevents.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), err
	}

	alertConfigRequest := admin.GroupAlertsConfig{
		GroupId:         currentModel.ProjectId,
		EventTypeName:   currentModel.EventTypeName,
		Enabled:         currentModel.Enabled,
		Matchers:        expandAlertConfigurationMatchers(currentModel.Matchers),
		MetricThreshold: expandAlertConfigurationMetricThresholdConfig(currentModel),
		Threshold:       expandAlertConfigurationThreshold(currentModel.Threshold),
		Notifications:   notifications,
	}

	projectID := *currentModel.ProjectId
	var res *http.Response
	alertConfig, res, err := atlasV2.AlertConfigurationsApi.CreateAlertConfiguration(context.Background(), projectID, &alertConfigRequest).Execute()
	defer res.Body.Close()
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res), nil
	}

	currentModel = convertToUIModel(alertConfig, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	validationError := validateRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.AtlasSDK

	if !isExist(currentModel, atlasV2) {
		_, _ = logger.Warnf("resource not exist for Id: %s", *currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	alertConfig, resp, err := atlasV2.AlertConfigurationsApi.GetAlertConfiguration(context.Background(), *currentModel.ProjectId, *currentModel.Id).Execute()
	defer resp.Body.Close()
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel = convertToUIModel(alertConfig, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	validationError := validateRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.AtlasSDK

	if !isExist(currentModel, atlasV2) {
		_, _ = logger.Warnf("resource not exist for Id: %s", *currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	// In order to update an alert config it is necessary to send the original alert configuration request again, if not the
	// server returns an error 500
	projectID := *currentModel.ProjectId
	id := *currentModel.Id
	alertReq, res, err := atlasV2.AlertConfigurationsApi.GetAlertConfiguration(context.Background(), projectID, id).Execute()
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res), nil
	}

	alertReq = convertToMongoModel(alertReq, currentModel)

	// Removing the computed attributes to recreate the original request
	alertReq.Created = nil
	alertReq.Updated = nil
	var alertModel *admin.GroupAlertsConfig

	// Cannot enable/disable ONLY via update (if only send enable as changed field server returns a 500 error)
	// so have to use different method to change enabled.
	if reflect.DeepEqual(alertReq, &admin.GroupAlertsConfig{Enabled: aws.Bool(true)}) ||
		reflect.DeepEqual(alertReq, &admin.GroupAlertsConfig{Enabled: aws.Bool(false)}) {
		alertModel, res, err = atlasV2.AlertConfigurationsApi.ToggleAlertConfiguration(context.Background(), projectID, id, &admin.AlertsToggle{Enabled: alertReq.Enabled}).Execute()
	} else {
		alertModel, res, err = atlasV2.AlertConfigurationsApi.UpdateAlertConfiguration(context.Background(), projectID, id, alertReq).Execute()
	}

	if err != nil {
		_, _ = logger.Warnf("Update - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res), nil
	}
	defer res.Body.Close()

	currentModel = convertToUIModel(alertModel, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	validationError := validateRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.AtlasSDK

	if !isExist(currentModel, atlasV2) {
		_, _ = logger.Warnf("resource not exist for Id: %s", *currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	res, err := atlasV2.AlertConfigurationsApi.DeleteAlertConfiguration(context.Background(), *currentModel.ProjectId, *currentModel.Id).Execute()

	if err != nil {
		_, _ = logger.Warnf("Delete - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          "List operation is not supported",
		HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
}

func isExist(currentModel *Model, client *admin.APIClient) bool {
	alert, _, err := client.AlertConfigurationsApi.GetAlertConfiguration(context.Background(), *currentModel.ProjectId, *currentModel.Id).Execute()
	return err == nil && alert != nil
}

func expandAlertConfigurationMatchers(matchers []Matcher) *[]map[string]interface{} {
	mts := make([]map[string]interface{}, 0)
	for ind := range matchers {
		mMatcher := map[string]interface{}{
			"fieldName": cast.ToString(matchers[ind].FieldName),
			"operator":  cast.ToString(matchers[ind].Operator),
			"value":     cast.ToString(matchers[ind].Value),
		}
		mts = append(mts, mMatcher)
	}
	return &mts
}

func expandAlertConfigurationMetricThresholdConfig(currentModel *Model) *admin.ServerlessMetricThreshold {
	threshold := currentModel.MetricThreshold
	if threshold == nil {
		return nil
	}
	return &admin.ServerlessMetricThreshold{
		MetricName: cast.ToString(threshold.MetricName),
		Operator:   threshold.Operator,
		Threshold:  threshold.Threshold,
		Units:      threshold.Units,
		Mode:       threshold.Mode,
	}
}

func expandAlertConfigurationThreshold(threshold *IntegerThresholdView) *admin.GreaterThanRawThreshold {
	if threshold == nil {
		return nil
	}
	return &admin.GreaterThanRawThreshold{
		Operator:  threshold.Operator,
		Threshold: util.Pointer(int(*threshold.Threshold)),
		Units:     threshold.Units,
	}
}

func expandAlertConfigurationNotification(notificationList []NotificationView) (*[]admin.AlertsNotificationRootForGroup, error) {
	notifications := make([]admin.AlertsNotificationRootForGroup, 0)

	for ind := range notificationList {
		if notificationList[ind].IntervalMin != nil && *notificationList[ind].IntervalMin > cast.ToFloat64(0) {
			typeName := *notificationList[ind].TypeName
			if strings.EqualFold(typeName, constants.PagerDuty) || strings.EqualFold(typeName, constants.OpsGenie) || strings.EqualFold(typeName, constants.VictorOps) {
				return nil, fmt.Errorf(`'interval_min' doesn't need to be set if type_name is 'PAGER_DUTY', 'OPS_GENIE' or 'VICTOR_OPS'`)
			}
		}
	}

	for ind := range notificationList {
		notification := admin.AlertsNotificationRootForGroup{
			ApiToken:                 notificationList[ind].ApiToken,
			ChannelName:              notificationList[ind].ChannelName,
			DatadogApiKey:            notificationList[ind].DatadogApiKey,
			DatadogRegion:            notificationList[ind].DatadogRegion,
			EmailAddress:             notificationList[ind].EmailAddress,
			EmailEnabled:             notificationList[ind].EmailEnabled,
			IntervalMin:              util.Pointer(int(*notificationList[ind].IntervalMin)),
			MicrosoftTeamsWebhookUrl: notificationList[ind].MicrosoftTeamsWebhookUrl,
			MobileNumber:             notificationList[ind].MobileNumber,
			OpsGenieApiKey:           notificationList[ind].OpsGenieApiKey,
			OpsGenieRegion:           notificationList[ind].OpsGenieRegion,
			ServiceKey:               notificationList[ind].ServiceKey,
			SmsEnabled:               notificationList[ind].SmsEnabled,
			TeamId:                   notificationList[ind].TeamId,
			TypeName:                 notificationList[ind].TypeName,
			Username:                 notificationList[ind].Username,
			VictorOpsApiKey:          notificationList[ind].VictorOpsApiKey,
			VictorOpsRoutingKey:      notificationList[ind].VictorOpsRoutingKey,
			Roles:                    &notificationList[ind].Roles,
			DelayMin:                 notificationList[ind].DelayMin,
		}
		notifications = append(notifications, notification)
	}
	return &notifications, nil
}

func convertToMongoModel(reqModel *admin.GroupAlertsConfig, currentModel *Model) *admin.GroupAlertsConfig {
	if reqModel == nil {
		reqModel = &admin.GroupAlertsConfig{}
	}

	// Only change the updated fields
	if currentModel.Enabled != nil {
		reqModel.Enabled = currentModel.Enabled
	}
	if currentModel.EventTypeName != nil {
		reqModel.EventTypeName = currentModel.EventTypeName
	}
	if currentModel.Matchers != nil {
		reqModel.Matchers = expandAlertConfigurationMatchers(currentModel.Matchers)
	}
	if currentModel.MetricThreshold != nil {
		reqModel.MetricThreshold = expandAlertConfigurationMetricThresholdConfig(currentModel)
	}
	if currentModel.Threshold != nil {
		reqModel.Threshold = expandAlertConfigurationThreshold(currentModel.Threshold)
	}
	if currentModel.Notifications != nil {
		reqModel.Notifications, _ = expandAlertConfigurationNotification(currentModel.Notifications)
	}
	return reqModel
}

func convertToUIModel(alertConfig *admin.GroupAlertsConfig, currentModel *Model) *Model {
	currentModel.Id = alertConfig.Id
	if alertConfig.Created != nil {
		currentModel.Created = util.TimePtrToStringPtr(alertConfig.Created)
	}

	if alertConfig.Updated != nil {
		currentModel.Updated = util.TimePtrToStringPtr(alertConfig.Updated)
	}

	if alertConfig.Enabled != nil {
		currentModel.Enabled = alertConfig.Enabled
	}

	return currentModel
}
