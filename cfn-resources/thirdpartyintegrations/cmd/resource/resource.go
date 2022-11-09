package resource

import (
	"context"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progress_event"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	log "github.com/sirupsen/logrus"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	"net/http"
)

var CreateRequiredFields = []string{"Type", "ProjectId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var ReadRequiredFields = []string{"ProjectId", "Type", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var UpdateRequiredFields = []string{"ProjectId", "Type", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var DeleteRequiredFields = []string{"Type", "ProjectId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var ListRequiredFields = []string{"ProjectId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}

var requiredPerType = map[string][]string{
	"PAGER_DUTY":      {"ServiceKey"},
	"DATADOG":         {"ApiKey", "Region"},
	"NEW_RELIC":       {"LicenseKey", "AccountId", "WriteToken", "ReadToken"},
	"OPS_GENIE":       {"ApiKey", "Region"},
	"VICTOR_OPS":      {"ApiKey"},
	"FLOWDOCK":        {"FlowName", "ApiToken", "OrgName"},
	"WEBHOOK":         {"Url"},
	"MICROSOFT_TEAMS": {"MicrosoftTeamsWebhookUrl"},
	"PROMETHEUS":      {"UserName", "Password", "ServiceDiscovery", "Scheme", "Enabled"},
}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-thirdpartyintegrations")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Create() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	ProjectID := currentModel.ProjectId
	IntegrationType := currentModel.Type

	// checking per type fields
	requiredInputs := requiredPerType[*IntegrationType]
	// Validation
	IntegrationTypeValidation := validateModel(requiredInputs, currentModel)
	if IntegrationTypeValidation != nil {
		return *IntegrationTypeValidation, nil
	}

	requestBody := modelToIntegration(currentModel)

	spew.Dump(requestBody, ProjectID)

	integrations, resModel, err := client.Integrations.Create(context.Background(), *ProjectID, *IntegrationType, requestBody)

	if err != nil {
		log.Debugf("Create - error: %+v", err)
		if resModel.Response.StatusCode == http.StatusConflict {
			return progress_events.GetFailedEventByCode("INTEGRATION_ALREADY_CONFIGURED.", cloudformation.HandlerErrorCodeAlreadyExists), nil
		}
		return progress_events.GetFailedEventByResponse(err.Error(), resModel.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   integrationToModel(*currentModel, integrations.Results[0]),
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Read() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Read - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	ProjectID := currentModel.ProjectId
	IntegrationType := currentModel.Type

	integration, res, err := client.Integrations.Get(context.Background(), *ProjectID, *IntegrationType)

	if err != nil {
		log.Debugf("Read - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   integrationToModel(*currentModel, integration),
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Update() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Update - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	ProjectID := currentModel.ProjectId
	IntegrationType := currentModel.Type

	integration, res, err := client.Integrations.Get(context.Background(), *ProjectID, *IntegrationType)
	if err != nil {
		log.Debugf("Update - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// check for changed attributes per type

	updateIntegrationFromSchema(currentModel, integration)

	integrations, res, err := client.Integrations.Replace(context.Background(), *ProjectID, *IntegrationType, integration)

	if err != nil {
		log.Debugf("Update - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   integrationToModel(*currentModel, integrations.Results[0]),
	}, nil
}

func updateIntegrationFromSchema(currentModel *Model, integration *mongodbatlas.ThirdPartyIntegration) {
	if currentModel.Url != nil && *currentModel.Url != integration.URL {
		integration.URL = *currentModel.Url
	}
	if currentModel.LicenseKey != nil && *currentModel.LicenseKey != integration.LicenseKey {
		integration.URL = *currentModel.Url
	}
	if currentModel.AccountId != nil && *currentModel.AccountId != integration.AccountID {
		integration.AccountID = *currentModel.AccountId
	}
	if currentModel.WriteToken != nil && *currentModel.WriteToken != integration.WriteToken {
		integration.WriteToken = *currentModel.WriteToken
	}
	if currentModel.ReadToken != nil && *currentModel.ReadToken != integration.ReadToken {
		integration.ReadToken = *currentModel.ReadToken
	}
	if currentModel.ApiKey != nil && *currentModel.ApiKey != integration.APIKey {
		integration.APIKey = *currentModel.ApiKey
	}
	if currentModel.Region != nil && *currentModel.Region != integration.Region {
		integration.Region = *currentModel.Region
	}
	if currentModel.ServiceKey != nil && *currentModel.ServiceKey != integration.ServiceKey {
		integration.ServiceKey = *currentModel.ServiceKey
	}
	if currentModel.ApiToken != nil && *currentModel.ApiToken != integration.APIToken {
		integration.APIToken = *currentModel.ApiToken
	}
	if currentModel.TeamName != nil && *currentModel.TeamName != integration.TeamName {
		integration.TeamName = *currentModel.TeamName
	}
	if currentModel.ChannelName != nil && *currentModel.ChannelName != integration.ChannelName {
		integration.ChannelName = *currentModel.ChannelName
	}
	if currentModel.RoutingKey != nil && *currentModel.RoutingKey != integration.RoutingKey {
		integration.RoutingKey = *currentModel.RoutingKey
	}
	if currentModel.FlowName != nil && *currentModel.FlowName != integration.FlowName {
		integration.FlowName = *currentModel.FlowName
	}
	if currentModel.OrgName != nil && *currentModel.OrgName != integration.OrgName {
		integration.OrgName = *currentModel.OrgName
	}
	if currentModel.Secret != nil && *currentModel.Secret != integration.Secret {
		integration.Secret = *currentModel.Secret
	}
	if currentModel.MicrosoftTeamsWebhookUrl != nil && *currentModel.MicrosoftTeamsWebhookUrl != integration.MicrosoftTeamsWebhookURL {
		integration.MicrosoftTeamsWebhookURL = *currentModel.MicrosoftTeamsWebhookUrl
	}
	if currentModel.UserName != nil && *currentModel.UserName != integration.UserName {
		integration.UserName = *currentModel.UserName
	}
	if currentModel.Password != nil && *currentModel.Password != integration.Password {
		integration.Password = *currentModel.Password
	}
	if currentModel.ServiceDiscovery != nil && *currentModel.ServiceDiscovery != integration.ServiceDiscovery {
		integration.ServiceDiscovery = *currentModel.ServiceDiscovery
	}
	if currentModel.Scheme != nil && *currentModel.Scheme != integration.Scheme {
		integration.Scheme = *currentModel.Scheme
	}
	if currentModel.Enabled != nil && *currentModel.Enabled != integration.Enabled {
		integration.Enabled = *currentModel.Enabled
	}
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Delete() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Delete - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	ProjectID := currentModel.ProjectId
	IntegrationType := currentModel.Type

	res, err = client.Integrations.Delete(context.Background(), *ProjectID, *IntegrationType)

	if err != nil {
		log.Debugf("Delete - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("List() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("List - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response
	ProjectID := currentModel.ProjectId
	integrations, res, err := client.Integrations.List(context.Background(), *ProjectID)
	if err != nil {
		log.Debugf("List - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	mm := make([]interface{}, 0)

	for _, integration := range integrations.Results {
		var m Model
		m = integrationToModel(m, integration)
		mm = append(mm, m)
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List successful",
		ResourceModels:  mm,
	}, nil
}

func modelToIntegration(currentModel *Model) (out *mongodbatlas.ThirdPartyIntegration) {
	out = &mongodbatlas.ThirdPartyIntegration{}

	if currentModel.Type != nil {
		out.Type = *currentModel.Type
	}
	if currentModel.Enabled != nil {
		out.Enabled = *currentModel.Enabled
	}
	if currentModel.Scheme != nil {
		out.Scheme = *currentModel.Scheme
	}
	if currentModel.ServiceDiscovery != nil {
		out.ServiceDiscovery = *currentModel.ServiceDiscovery
	}
	if currentModel.Password != nil {
		out.Password = *currentModel.Password
	}
	if currentModel.UserName != nil {
		out.UserName = *currentModel.UserName
	}
	if currentModel.MicrosoftTeamsWebhookUrl != nil {
		out.MicrosoftTeamsWebhookURL = *currentModel.MicrosoftTeamsWebhookUrl
	}
	if currentModel.Secret != nil {
		out.Secret = *currentModel.Secret
	}
	if currentModel.Url != nil {
		out.URL = *currentModel.Url
	}
	if currentModel.OrgName != nil {
		out.OrgName = *currentModel.OrgName
	}
	if currentModel.FlowName != nil {
		out.FlowName = *currentModel.FlowName
	}
	if currentModel.RoutingKey != nil {
		out.RoutingKey = *currentModel.RoutingKey
	}
	if currentModel.ChannelName != nil {
		out.ChannelName = *currentModel.ChannelName
	}
	if currentModel.TeamName != nil {
		out.TeamName = *currentModel.TeamName
	}
	if currentModel.ApiToken != nil {
		out.APIToken = *currentModel.ApiToken
	}
	if currentModel.ServiceKey != nil {
		out.ServiceKey = *currentModel.ServiceKey
	}
	if currentModel.Region != nil {
		out.Region = *currentModel.Region
	}
	if currentModel.ApiKey != nil {
		out.APIKey = *currentModel.ApiKey
	}
	if currentModel.ReadToken != nil {
		out.ReadToken = *currentModel.ReadToken
	}
	if currentModel.WriteToken != nil {
		out.WriteToken = *currentModel.WriteToken
	}
	if currentModel.AccountId != nil {
		out.AccountID = *currentModel.AccountId
	}
	if currentModel.LicenseKey != nil {
		out.LicenseKey = *currentModel.LicenseKey
	}

	return out
}

func integrationToModel(currentModel Model, integration *mongodbatlas.ThirdPartyIntegration) Model {

	enabled := false
	// if Enabled is not set we dont want to return false value to the response
	if currentModel.Enabled != nil {
		enabled = true
	}

	out := Model{
		Type:        &integration.Type,
		LicenseKey:  &integration.LicenseKey,
		AccountId:   &integration.AccountID,
		WriteToken:  &integration.WriteToken,
		ReadToken:   &integration.ReadToken,
		ApiKey:      &integration.APIKey,
		Region:      &integration.Region,
		ServiceKey:  &integration.ServiceKey,
		ApiToken:    &integration.APIToken,
		TeamName:    &integration.TeamName,
		ChannelName: &integration.ChannelName,
		RoutingKey:  &integration.RoutingKey,
		FlowName:    &integration.FlowName,
		OrgName:     &integration.OrgName,
		//Url:                      &integration.URL,
		//Secret:                   &integration.Secret,
		MicrosoftTeamsWebhookUrl: &integration.MicrosoftTeamsWebhookURL,
		UserName:                 &integration.UserName,
		//Password:                 &integration.Password,
		ServiceDiscovery: &integration.ServiceDiscovery,
		Scheme:           &integration.Scheme,
		Enabled:          &integration.Enabled,
		ProjectId:        currentModel.ProjectId,
		ApiKeys:          currentModel.ApiKeys,
	}

	if !enabled {
		out.Enabled = nil
	}
	return out
}
