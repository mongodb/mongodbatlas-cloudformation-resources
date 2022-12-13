package resource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	realmAuth "go.mongodb.org/realm/auth"
	"go.mongodb.org/realm/realm"
)

type TriggerType string

const (
	DATABASE       TriggerType = "DATABASE"
	SCHEDULED      TriggerType = "SCHEDULE"
	AUTHENTICATION TriggerType = "AUTHENTICATION"
	ua                         = "terraform-provider-mongodbatlas"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.AppID, constants.RealmPubKey, constants.RealmPvtKey}
var ReadRequiredFields = []string{constants.ProjectID, constants.AppID, constants.ID, constants.RealmPubKey, constants.RealmPvtKey}
var UpdateRequiredFields = []string{constants.ProjectID, constants.AppID, constants.RealmPubKey, constants.RealmPvtKey}
var DeleteRequiredFields = []string{constants.ProjectID, constants.AppID, constants.ID, constants.RealmPubKey, constants.RealmPvtKey}
var ListRequiredFields = []string{constants.ProjectID, constants.AppID, constants.RealmPubKey, constants.RealmPvtKey}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}
	ctx := context.Background()
	client, err := GetRealmClient(ctx, currentModel.RealmConfig)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	eventTrigger, err := newEventTrigger(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	et, _, err := client.EventTriggers.Create(ctx, *currentModel.ProjectId, *currentModel.AppId, eventTrigger)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
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
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}
	ctx := context.Background()
	client, err := GetRealmClient(ctx, currentModel.RealmConfig)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	trigger, res, err := client.EventTriggers.Get(ctx, *currentModel.ProjectId, *currentModel.AppId, *currentModel.Id)
	if err != nil {
		if res != nil && res.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}

		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
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
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}
	ctx := context.Background()
	client, err := GetRealmClient(ctx, currentModel.RealmConfig)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	eventTrigger, err := newEventTrigger(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	_, res, err := client.EventTriggers.Update(ctx, *currentModel.ProjectId, *currentModel.AppId, *currentModel.Id, eventTrigger)
	if err != nil {
		if res != nil && res.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if currentModel.Id == nil {
		err := errors.New("no Id found in currentModel")
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	if errEvent := validateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	ctx := context.Background()
	client, err := GetRealmClient(ctx, currentModel.RealmConfig)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	res, err := client.EventTriggers.Delete(ctx, *currentModel.ProjectId, *currentModel.AppId, *currentModel.Id)
	if err != nil {
		if res != nil && res.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}
	ctx := context.Background()
	client, err := GetRealmClient(ctx, currentModel.RealmConfig)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	triggers, _, err := client.EventTriggers.List(ctx, *currentModel.ProjectId, *currentModel.AppId)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   triggers,
	}, nil
}

func GetRealmClient(ctx context.Context, c *RealmConfig) (*realm.Client, error) {
	optsRealm := []realm.ClientOpt{realm.SetUserAgent(ua)}
	if c.BaseURL != nil && c.RealmBaseURL != nil {
		optsRealm = append(optsRealm, realm.SetBaseURL(*c.RealmBaseURL))
	}
	authConfig := realmAuth.NewConfig(nil)
	token, err := authConfig.NewTokenFromCredentials(ctx, *c.PublicKey, *c.PrivateKey)
	if err != nil {
		return nil, err
	}
	clientRealm := realmAuth.NewClient(realmAuth.BasicTokenSource(token))
	realmClient, err := realm.New(clientRealm, optsRealm...)
	if err != nil {
		return nil, err
	}
	return realmClient, nil
}

func newEventTrigger(model *Model) (*realm.EventTriggerRequest, error) {
	et := realm.EventTriggerRequest{Disabled: model.Disabled}
	if model.Name != nil {
		et.Name = *model.Name
	}
	if model.Type != nil {
		et.Type = *model.Type
	}
	if model.FunctionId != nil {
		et.FunctionID = *model.FunctionId
	}
	conf := realm.EventTriggerConfig{}
	if model.DatabaseTrigger != nil {
		conf.Database = *model.DatabaseTrigger.Database
	}
	dTrigger := model.DatabaseTrigger
	if dTrigger != nil {
		conf.Collection = aws.StringValue(dTrigger.Collection)
		conf.ServiceID = aws.StringValue(dTrigger.ServiceId)
		conf.OperationTypes = dTrigger.OperationTypes
		conf.Match = dTrigger.Match
		conf.FullDocument = dTrigger.FullDocument
		conf.FullDocumentBeforeChange = dTrigger.FullDocumentBeforeChange
		conf.Unordered = dTrigger.Unordered
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

func newEventProcessor(model *Model, et realm.EventTriggerRequest) (realm.EventTriggerRequest, error) {
	ep := EventProcess{}
	if model.EventProcessors.FUNCTION != nil && model.EventProcessors.FUNCTION.FuncConfig != nil {
		ep.FUNCTION = new(FUNC)
		ep.FUNCTION.FuncConf = new(FuncConf)
		if model.EventProcessors.FUNCTION.FuncConfig.FunctionName != nil {
			ep.FUNCTION.FuncConf.FunctionName = model.EventProcessors.FUNCTION.FuncConfig.FunctionName
		}
		if model.EventProcessors.FUNCTION.FuncConfig.FunctionId != nil {
			ep.FUNCTION.FuncConf.FunctionId = model.EventProcessors.FUNCTION.FuncConfig.FunctionId
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
		log.Printf("error in unmarshal %v", err)
		return et, err
	}
	et.EventProcessors = inInterface
	return et, nil
}

// EventProcess These structs are created because the client has generic map for event processor
// and cfn generate doesn't support tags
type EventProcess struct {
	FUNCTION       *FUNC      `json:",omitempty"`
	AWSEVENTBRIDGE *AWSEVENTB `json:",omitempty"`
}

type FUNC struct {
	FuncConf *FuncConf `json:"config,omitempty"`
}

type FuncConf struct {
	FunctionId   *string `json:"function_id,omitempty"`
	FunctionName *string `json:"function_name,omitempty"`
}

type AWSEVENTB struct {
	AWSConfig *AWSConf `json:"config,omitempty"`
}

type AWSConf struct {
	AccountId           *string `json:"account_id,omitempty"`
	Region              *string `json:",omitempty"`
	ExtendedJsonEnabled *bool   `json:"extended_json_enabled,omitempty"`
}
