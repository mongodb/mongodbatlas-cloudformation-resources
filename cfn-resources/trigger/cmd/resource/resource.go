// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb-labs/go-client-mongodb-atlas-app-services/appservices"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

type TriggerType string

const (
	DATABASE       TriggerType = "DATABASE"
	SCHEDULED      TriggerType = "SCHEDULE"
	AUTHENTICATION TriggerType = "AUTHENTICATION"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.AppID}
var ReadRequiredFields = []string{constants.ProjectID, constants.AppID, constants.ID}
var UpdateRequiredFields = []string{constants.ProjectID, constants.AppID, constants.ID}
var DeleteRequiredFields = []string{constants.ProjectID, constants.AppID, constants.ID}
var ListRequiredFields = []string{constants.ProjectID, constants.AppID}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("trigger")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}
	setProfileIfAbsent(currentModel)

	ctx := context.Background()
	client, err := util.GetAppServicesClient(ctx, req, currentModel.Profile)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating App Services client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	eventTrigger, err := newEventTrigger(currentModel)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating event trigger request : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	et, resp, err := client.EventTriggers.Create(ctx, *currentModel.ProjectId, *currentModel.AppId, eventTrigger)
	if err != nil {
		_, _ = logger.Warnf("error in creating event trigger %v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}
	currentModel.Id = &et.ID

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if currentModel.Id == nil {
		err := errors.New("no Id found in currentModel")
		return progressevents.GetFailedEventByCode(err.Error(),
			cloudformation.HandlerErrorCodeNotFound), nil
	}
	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}
	setProfileIfAbsent(currentModel)

	ctx := context.Background()
	client, err := util.GetAppServicesClient(ctx, req, currentModel.Profile)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating App Services client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	trigger, resp, err := client.EventTriggers.Get(ctx, *currentModel.ProjectId, *currentModel.AppId, *currentModel.Id)
	if err != nil {
		_, _ = logger.Warnf("error in getting event trigger %v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}
	currentModel.Id = &trigger.ID

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if currentModel.Id == nil {
		err := errors.New("no Id found in currentModel")
		return progressevents.GetFailedEventByCode(err.Error(),
			cloudformation.HandlerErrorCodeNotFound), nil
	}
	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}
	setProfileIfAbsent(currentModel)

	ctx := context.Background()
	client, err := util.GetAppServicesClient(ctx, req, currentModel.Profile)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating App Services client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	eventTrigger, err := newEventTrigger(currentModel)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating trigger request : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	_, resp, err := client.EventTriggers.Update(ctx, *currentModel.ProjectId, *currentModel.AppId, *currentModel.Id, eventTrigger)
	if err != nil {
		_, _ = logger.Warnf("error in updating event trigger %v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if currentModel.Id == nil {
		err := errors.New("no Id found in currentModel")
		return progressevents.GetFailedEventByCode(err.Error(),
			cloudformation.HandlerErrorCodeNotFound), nil
	}
	if errEvent := validateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}
	setProfileIfAbsent(currentModel)

	ctx := context.Background()
	client, err := util.GetAppServicesClient(ctx, req, currentModel.Profile)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating App Services client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	resp, err := client.EventTriggers.Delete(ctx, *currentModel.ProjectId, *currentModel.AppId, *currentModel.Id)
	if err != nil {
		_, _ = logger.Warnf("error in deleting event trigger %v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}
	setProfileIfAbsent(currentModel)

	ctx := context.Background()
	client, err := util.GetAppServicesClient(ctx, req, currentModel.Profile)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating App Services client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	triggers, resp, err := client.EventTriggers.List(ctx, *currentModel.ProjectId, *currentModel.AppId)
	if err != nil {
		_, _ = logger.Warnf("error in listing event trigger %v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   triggers,
	}, nil
}

func setProfileIfAbsent(model *Model) {
	if model.Profile == nil || *model.Profile == "" {
		model.Profile = aws.String(profile.DefaultProfile)
	}
}

func newEventTrigger(model *Model) (*appservices.EventTriggerRequest, error) {
	et := appservices.EventTriggerRequest{Disabled: model.Disabled}
	if model.Name != nil {
		et.Name = *model.Name
	}
	if model.Type != nil {
		et.Type = *model.Type
	}
	if model.FunctionId != nil {
		et.FunctionID = *model.FunctionId
	}
	conf := appservices.EventTriggerConfig{}
	if model.DatabaseTrigger != nil {
		conf.Database = *model.DatabaseTrigger.Database
	}
	dTrigger := model.DatabaseTrigger
	if dTrigger != nil {
		if dTrigger.Match != nil {
			jsonData := []byte(*dTrigger.Match)
			// convert the JSON string to a map
			var m interface{}
			if err := json.Unmarshal(jsonData, &m); err != nil {
				return nil, errors.New("error unmarshalling Match field - " + err.Error())
			}
			conf.Match = m
		}

		if dTrigger.Project != nil {
			jsonData := []byte(*dTrigger.Project)
			// convert the JSON string to a map
			var m interface{}
			if err := json.Unmarshal(jsonData, &m); err != nil {
				return nil, errors.New("error unmarshalling Project field - " + err.Error())
			}
			conf.Project = m
		}
		conf.Collection = aws.StringValue(dTrigger.Collection)
		conf.ServiceID = aws.StringValue(dTrigger.ServiceId)
		conf.OperationTypes = dTrigger.OperationTypes
		conf.FullDocument = dTrigger.FullDocument
		conf.FullDocumentBeforeChange = dTrigger.FullDocumentBeforeChange
		conf.Unordered = dTrigger.Unordered
		conf.TolerateResumeErrors = dTrigger.TolerateResumeErrors
		conf.SkipCatchupEvents = dTrigger.SkipCatchupEvents
		conf.MaximumThroughput = dTrigger.MaximumThroughput
	}

	if model.ScheduleTrigger != nil &&
		model.ScheduleTrigger.Schedule != nil {
		conf.Schedule = *model.ScheduleTrigger.Schedule
	}

	if model.AuthTrigger != nil {
		if model.AuthTrigger.Providers != nil {
			conf.Providers = model.AuthTrigger.Providers
		}
		if model.AuthTrigger.OperationType != nil {
			conf.OperationType = *model.AuthTrigger.OperationType
		}
	}

	et.Config = &conf
	var err2 error
	if model.EventProcessors != nil {
		et, err2 = newEventProcessor(model, et)
		if err2 != nil {
			return &et, err2
		}
	}
	switch *model.Type {
	case string(DATABASE):
		if len(et.Config.OperationTypes) == 0 || et.Config.Database == "" || et.Config.Collection == "" || et.Config.ServiceID == "" {
			return &et, fmt.Errorf("`config_operation_types`, `config_database`,`config_collection`,`config_service_id` must be provided if type is DATABASE")
		}
	case string(AUTHENTICATION):
		if len(et.Config.OperationTypes) == 0 || len(et.Config.Providers) == 0 {
			return &et, fmt.Errorf("`config_operation_type`, `config_providers` must be provided if type is AUTHENTICATION")
		}
	case string(SCHEDULED):
		if et.Config.Schedule == "" {
			return &et, fmt.Errorf("`config_schedule` must be provided if type is SCHEDULED")
		}
	}
	return &et, nil
}

func newEventProcessor(model *Model, et appservices.EventTriggerRequest) (appservices.EventTriggerRequest, error) {
	ep := EventProcess{}
	if model.EventProcessors.FUNCTION != nil && model.EventProcessors.FUNCTION.FuncConfig != nil {
		ep.FUNCTION = new(FUNC)
		ep.FUNCTION.FuncConf = new(FuncConf)
		if model.EventProcessors.FUNCTION.FuncConfig.FunctionName != nil {
			ep.FUNCTION.FuncConf.FunctionName = model.EventProcessors.FUNCTION.FuncConfig.FunctionName
		}
		if model.EventProcessors.FUNCTION.FuncConfig.FunctionId != nil {
			ep.FUNCTION.FuncConf.FunctionID = model.EventProcessors.FUNCTION.FuncConfig.FunctionId
		}
	}
	awsEventBridge := model.EventProcessors.AWSEVENTBRIDGE
	if awsEventBridge != nil && awsEventBridge.AWSConfig.AccountId != nil {
		ep.AWSEVENTBRIDGE = new(AWSEVENTB)
		if awsEventBridge.AWSConfig != nil {
			ep.AWSEVENTBRIDGE.AWSConfig = new(AWSConf)
			ep.AWSEVENTBRIDGE.AWSConfig.AccountID = awsEventBridge.AWSConfig.AccountId
			ep.AWSEVENTBRIDGE.AWSConfig.ExtendedJSONEnabled = awsEventBridge.AWSConfig.ExtendedJsonEnabled
			ep.AWSEVENTBRIDGE.AWSConfig.Region = awsEventBridge.AWSConfig.Region
		}
	}
	var inInterface map[string]interface{}
	inrec, err := json.Marshal(ep)
	if err != nil {
		log.Printf("error in marshal %v", err)
		return et, err
	}
	err = json.Unmarshal(inrec, &inInterface)
	if err != nil {
		_, _ = logger.Warnf("error %v", err)
		return et, err
	}
	et.EventProcessors = inInterface
	return et, nil
}

// EventProcess These structs are created because the client has generic map for event processor
// and cfn generate doesn't support tags
type EventProcess struct {
	FUNCTION       *FUNC      `json:",omitempty"`
	AWSEVENTBRIDGE *AWSEVENTB `json:"AWS_EVENTBRIDGE,omitempty"`
}

type FUNC struct {
	FuncConf *FuncConf `json:"config,omitempty"`
}

type FuncConf struct {
	FunctionID   *string `json:"function_id,omitempty"`
	FunctionName *string `json:"function_name,omitempty"`
}

type AWSEVENTB struct {
	AWSConfig *AWSConf `json:"config,omitempty"`
}

type AWSConf struct {
	AccountID           *string `json:"account_id,omitempty"`
	Region              *string `json:"region,omitempty"`
	ExtendedJSONEnabled *bool   `json:"extended_json_enabled,omitempty"`
}
