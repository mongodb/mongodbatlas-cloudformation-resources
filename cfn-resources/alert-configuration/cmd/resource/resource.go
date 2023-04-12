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
	"go.mongodb.org/atlas/mongodbatlas"
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
	setup() // logger setup

	validationError := validateRequest(CreateRequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Check if  already exist
	if currentModel.Id != nil && *currentModel.Id != "" {
		_, _ = logger.Warnf("resource already exists for Id: %s", *currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Already Exists",
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}

	// API Request creation
	alertConfigRequest := &mongodbatlas.AlertConfiguration{
		GroupID:         cast.ToString(currentModel.ProjectId),
		EventTypeName:   cast.ToString(currentModel.EventTypeName),
		Enabled:         currentModel.Enabled,
		Matchers:        expandAlertConfigurationMatchers(currentModel.Matchers),
		MetricThreshold: expandAlertConfigurationMetricThresholdConfig(currentModel),
		Threshold:       expandAlertConfigurationThreshold(currentModel.Threshold),
	}
	if currentModel.Notifications != nil {
		alertConfigRequest.Notifications, _ = expandAlertConfigurationNotification(currentModel.Notifications)
	}

	projectID := *currentModel.ProjectId
	// API call to create
	var res *mongodbatlas.Response
	alertConfig, res, err := client.AlertConfigurations.Create(context.Background(), projectID, alertConfigRequest)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel = convertToUIModel(alertConfig, currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	validationError := validateRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Check if  already exist
	if !isExist(currentModel, client) {
		_, _ = logger.Warnf("resource not exist for Id: %s", *currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	// API call to read resource
	alertConfig, resp, err := client.AlertConfigurations.GetAnAlertConfig(context.Background(), *currentModel.ProjectId, *currentModel.Id)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	// populate response model
	currentModel = convertToUIModel(alertConfig, currentModel)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	validationError := validateRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Check if  already exist
	if !isExist(currentModel, client) {
		_, _ = logger.Warnf("resource not exist for Id: %s", *currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	var res *mongodbatlas.Response

	// In order to update an alert config it is necessary to send the original alert configuration request again, if not the
	// server returns an error 500
	projectID := *currentModel.ProjectId
	id := *currentModel.Id
	alertReq, res, err := client.AlertConfigurations.GetAnAlertConfig(context.Background(), projectID, id)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// create request object
	alertReq = convertToMongoModel(alertReq, currentModel)

	// Removing the computed attributes to recreate the original request
	alertReq.GroupID = ""
	alertReq.Created = ""
	alertReq.Updated = ""
	var alertModel *mongodbatlas.AlertConfiguration

	// Cannot enable/disable ONLY via update (if only send enable as changed field server returns a 500 error)
	// so have to use different method to change enabled.
	if reflect.DeepEqual(alertReq, &mongodbatlas.AlertConfiguration{Enabled: aws.Bool(true)}) ||
		reflect.DeepEqual(alertReq, &mongodbatlas.AlertConfiguration{Enabled: aws.Bool(false)}) {
		alertModel, res, err = client.AlertConfigurations.EnableAnAlertConfig(context.Background(), projectID, id, alertReq.Enabled)
	} else {
		alertModel, res, err = client.AlertConfigurations.Update(context.Background(), projectID, id, alertReq)
	}

	if err != nil {
		_, _ = logger.Warnf("Update - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	currentModel = convertToUIModel(alertModel, currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	validationError := validateRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Check if  already exist
	if !isExist(currentModel, client) {
		_, _ = logger.Warnf("resource not exist for Id: %s", *currentModel.Id)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	// API call to delete
	res, err := client.AlertConfigurations.Delete(context.Background(), *currentModel.ProjectId, *currentModel.Id)

	if err != nil {
		_, _ = logger.Warnf("Delete - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Response
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

// function to check if record already exist
func isExist(currentModel *Model, client *mongodbatlas.Client) bool {
	alert, _, err := client.AlertConfigurations.GetAnAlertConfig(context.Background(), *currentModel.ProjectId, *currentModel.Id)
	return err == nil && alert != nil
}

func expandAlertConfigurationMatchers(matchers []Matcher) []mongodbatlas.Matcher {
	mts := make([]mongodbatlas.Matcher, 0)
	for ind := range matchers {
		mMatcher := mongodbatlas.Matcher{
			FieldName: cast.ToString(matchers[ind].FieldName),
			Operator:  cast.ToString(matchers[ind].Operator),
			Value:     cast.ToString(matchers[ind].Value),
		}
		mts = append(mts, mMatcher)
	}
	return mts
}

func expandAlertConfigurationMetricThresholdConfig(currentModel *Model) *mongodbatlas.MetricThreshold {
	threshold := currentModel.MetricThreshold
	if threshold == nil {
		return nil
	}
	return &mongodbatlas.MetricThreshold{
		MetricName: cast.ToString(threshold.MetricName),
		Operator:   cast.ToString(threshold.Operator),
		Threshold:  cast.ToFloat64(threshold.Threshold),
		Units:      cast.ToString(threshold.Units),
		Mode:       cast.ToString(threshold.Mode),
	}
}

func expandAlertConfigurationThreshold(threshold *IntegerThresholdView) *mongodbatlas.Threshold {
	if threshold == nil {
		return nil
	}
	return &mongodbatlas.Threshold{
		Operator:  cast.ToString(threshold.Operator),
		Threshold: cast.ToFloat64(threshold.Threshold),
		Units:     cast.ToString(threshold.Units),
	}
}

// convert  model notification to mongodb atlas  notification
func expandAlertConfigurationNotification(notificationList []NotificationView) ([]mongodbatlas.Notification, error) {
	notifications := make([]mongodbatlas.Notification, 0)

	for ind := range notificationList {
		if *notificationList[ind].IntervalMin > cast.ToFloat64(0) {
			typeName := *notificationList[ind].TypeName
			if strings.EqualFold(typeName, constants.PagerDuty) || strings.EqualFold(typeName, constants.OpsGenie) || strings.EqualFold(typeName, constants.VictorOps) {
				return nil, fmt.Errorf(`'interval_min' doesn't need to be set if type_name is 'PAGER_DUTY', 'OPS_GENIE' or 'VICTOR_OPS'`)
			}
		}
	}

	for ind := range notificationList {
		notification := mongodbatlas.Notification{
			APIToken:            cast.ToString(notificationList[ind].ApiToken),
			ChannelName:         cast.ToString(notificationList[ind].ChannelName),
			DatadogAPIKey:       cast.ToString(notificationList[ind].DatadogApiKey),
			DatadogRegion:       cast.ToString(notificationList[ind].DatadogRegion),
			EmailAddress:        cast.ToString(notificationList[ind].EmailAddress),
			EmailEnabled:        notificationList[ind].EmailEnabled,
			FlowdockAPIToken:    cast.ToString(notificationList[ind].FlowdockApiToken),
			FlowName:            cast.ToString(notificationList[ind].FlowName),
			IntervalMin:         cast.ToInt(notificationList[ind].IntervalMin),
			MobileNumber:        cast.ToString(notificationList[ind].MobileNumber),
			OpsGenieAPIKey:      cast.ToString(notificationList[ind].OpsGenieApiKey),
			OpsGenieRegion:      cast.ToString(notificationList[ind].OpsGenieRegion),
			OrgName:             cast.ToString(notificationList[ind].OrgName),
			ServiceKey:          cast.ToString(notificationList[ind].ServiceKey),
			SMSEnabled:          notificationList[ind].SmsEnabled,
			TeamID:              cast.ToString(notificationList[ind].TeamId),
			TypeName:            cast.ToString(notificationList[ind].TypeName),
			Username:            cast.ToString(notificationList[ind].Username),
			VictorOpsAPIKey:     cast.ToString(notificationList[ind].VictorOpsApiKey),
			VictorOpsRoutingKey: cast.ToString(notificationList[ind].VictorOpsRoutingKey),
			Roles:               cast.ToStringSlice(notificationList[ind].Roles),
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

func convertToMongoModel(reqModel *mongodbatlas.AlertConfiguration, currentModel *Model) *mongodbatlas.AlertConfiguration {
	if reqModel == nil {
		reqModel = &mongodbatlas.AlertConfiguration{}
	}

	// Only change the updated fields
	if currentModel.Enabled != nil {
		reqModel.Enabled = currentModel.Enabled
	}
	if currentModel.EventTypeName != nil {
		reqModel.EventTypeName = *currentModel.EventTypeName
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

func convertToUIModel(alertConfig *mongodbatlas.AlertConfiguration, currentModel *Model) *Model {
	currentModel.Id = &alertConfig.ID
	if alertConfig.Created != "" {
		currentModel.Created = &alertConfig.Created
	}

	if alertConfig.Updated != "" {
		currentModel.Updated = &alertConfig.Updated
	}

	return currentModel
}
