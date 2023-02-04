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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.IntegrationType, constants.PubKey, constants.PvtKey, constants.ProjectID}
var ReadRequiredFields = []string{constants.IntegrationType, constants.PubKey, constants.PvtKey, constants.ProjectID}
var UpdateRequiredFields = []string{constants.IntegrationType, constants.PubKey, constants.PvtKey, constants.ProjectID}
var DeleteRequiredFields = []string{constants.IntegrationType, constants.PubKey, constants.PvtKey, constants.ProjectID}
var ListRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID}

// Custom validation only for ThirdPartyIntegrations
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
	util.SetupLogger("mongodb-atlas-thirdpartyintegration")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Warnf("Create() currentModel:%+v", currentModel)

	// Validation
	if modelValidation := validateModel(CreateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
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

	integrations, resModel, err := client.Integrations.Create(context.Background(), *ProjectID, *IntegrationType, requestBody)

	if err != nil {
		fmt.Printf("Create - error: %+v", err)
		_, _ = log.Debugf("Create - error: %+v", err)
		if resModel.Response.StatusCode == http.StatusConflict {
			return progressevents.GetFailedEventByCode("INTEGRATION_ALREADY_CONFIGURED.", cloudformation.HandlerErrorCodeAlreadyExists), nil
		}
		return progressevents.GetFailedEventByResponse(err.Error(), resModel.Response), nil
	}
	_, _ = log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   integrationToModel(*currentModel, integrations.Results[0]),
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Read() currentModel:%+v", currentModel)

	// Validation
	if modelValidation := validateModel(ReadRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Read - error: %+v", err)
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
		_, _ = log.Debugf("Read - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	_, _ = log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   integrationToModel(*currentModel, integration),
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Update() currentModel:%+v", currentModel)

	// Validation
	if modelValidation := validateModel(UpdateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Update - error: %+v", err)
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
		_, _ = log.Debugf("Update - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// check for changed attributes per type

	updateIntegrationFromSchema(currentModel, integration)

	integrations, res, err := client.Integrations.Replace(context.Background(), *ProjectID, *IntegrationType, integration)

	if err != nil {
		_, _ = log.Debugf("Update - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	_, _ = log.Debugf("Atlas Client %v", client)

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

	_, _ = log.Debugf("Delete() currentModel:%+v", currentModel)

	// Validation
	if modelValidation := validateModel(DeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Delete - error: %+v", err)
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
		_, _ = log.Debugf("Delete - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	_, _ = log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("List() currentModel:%+v", currentModel)

	// Validation
	if modelValidation := validateModel(ListRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("List - error: %+v", err)
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
		_, _ = log.Debugf("List - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	_, _ = log.Debugf("Atlas Client %v", client)

	mm := make([]interface{}, 0)

	for _, integration := range integrations.Results {
		m := integrationToModel(*currentModel, integration)
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

	return out
}

func integrationToModel(currentModel Model, integration *mongodbatlas.ThirdPartyIntegration) Model {
	enabled := false
	// if "Enabled" is not set in the inputs we dont want to return "Enabled" in outputs
	if currentModel.Enabled != nil {
		enabled = true
	}

	/*
	   The variables from the thirdparty integration are not returned back in reposnse because most of the variables are sensitive variables.
	*/
	out := Model{
		Type:      &integration.Type,
		ApiKeys:   currentModel.ApiKeys,
		ProjectId: currentModel.ProjectId,
	}

	if !enabled {
		out.Enabled = nil
	}
	return out
}
