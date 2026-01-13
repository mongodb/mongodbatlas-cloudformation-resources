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
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20250312010/admin"
)

var RequiredFields = []string{constants.IntegrationType, constants.ProjectID}
var ListRequiredFields = []string{constants.ProjectID}

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
	"PROMETHEUS":      {"UserName", "Password", "ServiceDiscovery", "Enabled"},
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
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	ProjectID := currentModel.ProjectId
	IntegrationType := currentModel.Type

	// checking per type fields
	requiredInputs := requiredPerType[*IntegrationType]
	if validationErr := validateModel(requiredInputs, currentModel); validationErr != nil {
		return *validationErr, nil
	}

	requestBody := modelToIntegration(currentModel)
	integrations, resModel, err := client.AtlasSDK.ThirdPartyIntegrationsApi.CreateGroupIntegration(context.Background(), *IntegrationType, *ProjectID, requestBody).Execute()
	if err != nil {
		if apiError, ok := admin.AsError(err); ok && apiError.Error == http.StatusConflict {
			return progressevent.GetFailedEventByCode("INTEGRATION_ALREADY_CONFIGURED.", cloudformation.HandlerErrorCodeAlreadyExists), nil
		}

		return progressevent.GetFailedEventByResponse(err.Error(), resModel), nil
	}

	if integrations == nil || len(integrations.GetResults()) == 0 {
		return progressevent.GetFailedEventByResponse("No integration returned from create", resModel), nil
	}

	results := integrations.GetResults()
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   integrationToModel(*currentModel, &results[0]),
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Read() currentModel:%+v", currentModel)
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	ProjectID := currentModel.ProjectId
	IntegrationType := currentModel.Type

	integration, res, err := client.AtlasSDK.ThirdPartyIntegrationsApi.GetGroupIntegration(context.Background(), *ProjectID, *IntegrationType).Execute()

	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}
	_, _ = log.Debugf("Atlas Client %v", client)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   integrationToModel(*currentModel, integration),
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Update() currentModel:%+v", currentModel)

	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	ProjectID := currentModel.ProjectId
	IntegrationType := currentModel.Type

	integration, res, err := client.AtlasSDK.ThirdPartyIntegrationsApi.GetGroupIntegration(context.Background(), *ProjectID, *IntegrationType).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	updateIntegrationFromSchema(currentModel, integration)
	integrations, res, err := client.AtlasSDK.ThirdPartyIntegrationsApi.UpdateGroupIntegration(context.Background(), *IntegrationType, *ProjectID, integration).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	if integrations == nil || len(integrations.GetResults()) == 0 {
		return progressevent.GetFailedEventByResponse("No integration returned from update", res), nil
	}

	results := integrations.GetResults()
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   integrationToModel(*currentModel, &results[0]),
	}, nil
}

func updateIntegrationFromSchema(currentModel *Model, integration *admin.ThirdPartyIntegration) {
	if util.IsStringPresent(currentModel.Url) && !util.AreStringPtrEqual(currentModel.Url, integration.Url) {
		integration.Url = currentModel.Url
	}
	if util.IsStringPresent(currentModel.ApiKey) && !util.AreStringPtrEqual(currentModel.ApiKey, integration.ApiKey) {
		integration.ApiKey = currentModel.ApiKey
	}
	if util.IsStringPresent(currentModel.Region) && !util.AreStringPtrEqual(currentModel.Region, integration.Region) {
		integration.Region = currentModel.Region
	}
	if util.IsStringPresent(currentModel.ServiceKey) && !util.AreStringPtrEqual(currentModel.ServiceKey, integration.ServiceKey) {
		integration.ServiceKey = currentModel.ServiceKey
	}
	if util.IsStringPresent(currentModel.ApiToken) && !util.AreStringPtrEqual(currentModel.ApiToken, integration.ApiToken) {
		integration.ApiToken = currentModel.ApiToken
	}
	if util.IsStringPresent(currentModel.TeamName) && !util.AreStringPtrEqual(currentModel.TeamName, integration.TeamName) {
		integration.TeamName = currentModel.TeamName
	}
	if util.IsStringPresent(currentModel.ChannelName) && !util.AreStringPtrEqual(currentModel.ChannelName, integration.ChannelName) {
		integration.ChannelName = currentModel.ChannelName
	}
	if util.IsStringPresent(currentModel.RoutingKey) && !util.AreStringPtrEqual(currentModel.RoutingKey, integration.RoutingKey) {
		integration.RoutingKey = currentModel.RoutingKey
	}
	if util.IsStringPresent(currentModel.Secret) && !util.AreStringPtrEqual(currentModel.Secret, integration.Secret) {
		integration.Secret = currentModel.Secret
	}
	if util.IsStringPresent(currentModel.MicrosoftTeamsWebhookUrl) && !util.AreStringPtrEqual(currentModel.MicrosoftTeamsWebhookUrl, integration.MicrosoftTeamsWebhookUrl) {
		integration.MicrosoftTeamsWebhookUrl = currentModel.MicrosoftTeamsWebhookUrl
	}
	if util.IsStringPresent(currentModel.UserName) && !util.AreStringPtrEqual(currentModel.UserName, integration.Username) {
		integration.Username = currentModel.UserName
	}
	if util.IsStringPresent(currentModel.Password) && !util.AreStringPtrEqual(currentModel.Password, integration.Password) {
		integration.Password = currentModel.Password
	}
	if util.IsStringPresent(currentModel.ServiceDiscovery) && !util.AreStringPtrEqual(currentModel.ServiceDiscovery, integration.ServiceDiscovery) {
		integration.ServiceDiscovery = currentModel.ServiceDiscovery
	}
	if currentModel.Enabled != nil && currentModel.Enabled != integration.Enabled {
		integration.Enabled = currentModel.Enabled
	}

	if currentModel.SendUserProvidedResourceTags != nil {
		integration.SendUserProvidedResourceTags = currentModel.SendUserProvidedResourceTags
	}
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Delete() currentModel:%+v", currentModel)
	if modelValidation := validateModel(RequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	var res *http.Response
	var err error

	ProjectID := currentModel.ProjectId
	IntegrationType := currentModel.Type

	res, err = client.AtlasSDK.ThirdPartyIntegrationsApi.DeleteGroupIntegration(context.Background(), *IntegrationType, *ProjectID).Execute()

	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("List() currentModel:%+v", currentModel)
	if modelValidation := validateModel(ListRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	var res *http.Response
	ProjectID := currentModel.ProjectId
	integrations, res, err := client.AtlasSDK.ThirdPartyIntegrationsApi.ListGroupIntegrations(context.Background(), *ProjectID).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	mm := make([]interface{}, 0)
	if integrations != nil {
		results := integrations.GetResults()
		for i := range results {
			m := integrationToModel(*currentModel, &results[i])
			mm = append(mm, m)
		}
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List successful",
		ResourceModels:  mm,
	}, nil
}

func modelToIntegration(currentModel *Model) (out *admin.ThirdPartyIntegration) {
	out = &admin.ThirdPartyIntegration{}

	if util.IsStringPresent(currentModel.Type) {
		out.Type = currentModel.Type
	}
	if currentModel.Enabled != nil {
		out.Enabled = currentModel.Enabled
	}
	if util.IsStringPresent(currentModel.ServiceDiscovery) {
		out.ServiceDiscovery = currentModel.ServiceDiscovery
	}
	if util.IsStringPresent(currentModel.Password) {
		out.Password = currentModel.Password
	}
	if util.IsStringPresent(currentModel.UserName) {
		out.Username = currentModel.UserName
	}
	if util.IsStringPresent(currentModel.MicrosoftTeamsWebhookUrl) {
		out.MicrosoftTeamsWebhookUrl = currentModel.MicrosoftTeamsWebhookUrl
	}
	if util.IsStringPresent(currentModel.Secret) {
		out.Secret = currentModel.Secret
	}
	if util.IsStringPresent(currentModel.Url) {
		out.Url = currentModel.Url
	}
	if util.IsStringPresent(currentModel.RoutingKey) {
		out.RoutingKey = currentModel.RoutingKey
	}
	if util.IsStringPresent(currentModel.ChannelName) {
		out.ChannelName = currentModel.ChannelName
	}
	if util.IsStringPresent(currentModel.TeamName) {
		out.TeamName = currentModel.TeamName
	}
	if util.IsStringPresent(currentModel.ApiToken) {
		out.ApiToken = currentModel.ApiToken
	}
	if util.IsStringPresent(currentModel.ServiceKey) {
		out.ServiceKey = currentModel.ServiceKey
	}
	if util.IsStringPresent(currentModel.Region) {
		out.Region = currentModel.Region
	}
	if util.IsStringPresent(currentModel.ApiKey) {
		out.ApiKey = currentModel.ApiKey
	}
	if currentModel.SendUserProvidedResourceTags != nil {
		out.SendUserProvidedResourceTags = currentModel.SendUserProvidedResourceTags
	}
	return out
}

func integrationToModel(currentModel Model, integration *admin.ThirdPartyIntegration) Model {
	// if "Enabled" is not set in the inputs we dont want to return "Enabled" in outputs
	enabled := currentModel.Enabled != nil

	/*
	   The variables from the thirdparty integration are not returned back in reposnse because most of the variables are sensitive variables.
	*/
	out := Model{
		Type:      integration.Type,
		ProjectId: currentModel.ProjectId,
		Profile:   currentModel.Profile,
	}

	if !enabled {
		out.Enabled = nil
	}

	return out
}
