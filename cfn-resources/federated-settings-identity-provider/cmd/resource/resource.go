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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	OktaIdpID = "OktaIdpId"
)

var CreateRequiredFields = []string{constants.PvtKey, constants.PubKey}
var ReadRequiredFields = []string{constants.FederationSettingsID, OktaIdpID, constants.PvtKey, constants.PubKey}
var UpdateRequiredFields = []string{constants.FederationSettingsID, OktaIdpID, constants.PvtKey, constants.PubKey}
var DeleteRequiredFields = []string{constants.PvtKey, constants.PubKey}
var ListRequiredFields = []string{constants.PvtKey, constants.PubKey, constants.FederationSettingsID}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-federatedsettingsidentityprovider")
}

func (m *Model) CompleteByResponse(response mongodbatlas.FederatedSettingsIdentityProvider) {
	m.SsoDebugEnabled = response.SsoDebugEnabled
	m.AssociatedDomains = response.AssociatedDomains
	m.Status = &response.Status
	m.IssuerUri = &response.IssuerURI
	m.RequestBinding = &response.RequestBinding
	m.ResponseSignatureAlgorithm = &response.ResponseSignatureAlgorithm
	m.SsoUrl = &response.SsoURL
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Debugf("Read - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	federatedSettingsIdentityProvider, resp, err := client.FederatedSettings.GetIdentityProvider(context.Background(), *currentModel.FederationSettingsId, *currentModel.OktaIdpId)
	if err != nil {
		_, _ = logger.Debugf("Read - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	currentModel.CompleteByResponse(*federatedSettingsIdentityProvider)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Debugf("Update - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	federatedSettingsIdentityProviderUpdate, resp, err := client.FederatedSettings.GetIdentityProvider(context.Background(), *currentModel.FederationSettingsId, *currentModel.OktaIdpId)
	if err != nil {
		_, _ = logger.Debugf("Update - error getting Identity Provider: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	federatedSettingsIdentityProviderUpdate.SsoDebugEnabled = currentModel.SsoDebugEnabled
	federatedSettingsIdentityProviderUpdate.AssociatedDomains = currentModel.AssociatedDomains

	if currentModel.DisplayName != nil {
		federatedSettingsIdentityProviderUpdate.DisplayName = *currentModel.DisplayName
	}
	if currentModel.Status != nil {
		federatedSettingsIdentityProviderUpdate.Status = *currentModel.Status
	}

	if currentModel.IssuerUri != nil {
		federatedSettingsIdentityProviderUpdate.IssuerURI = *currentModel.IssuerUri
	}

	if currentModel.RequestBinding != nil {
		federatedSettingsIdentityProviderUpdate.RequestBinding = *currentModel.RequestBinding
	}

	if currentModel.ResponseSignatureAlgorithm != nil {
		federatedSettingsIdentityProviderUpdate.ResponseSignatureAlgorithm = *currentModel.ResponseSignatureAlgorithm
	}

	if currentModel.SsoUrl != nil {
		federatedSettingsIdentityProviderUpdate.SsoURL = *currentModel.SsoUrl
	}

	federatedSettingsIdentityProviderUpdate.PemFileInfo = nil
	_, _, err = client.FederatedSettings.UpdateIdentityProvider(context.Background(), *currentModel.FederationSettingsId, *currentModel.OktaIdpId, federatedSettingsIdentityProviderUpdate)
	if err != nil {
		_, _ = logger.Debugf("Update - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Debugf("List - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	federatedSettingsIdentityProviders, resp, err := client.FederatedSettings.ListIdentityProviders(context.Background(), *currentModel.FederationSettingsId, nil)
	if err != nil {
		_, _ = logger.Debugf("Read - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	mm := make([]interface{}, 0)
	for i := range federatedSettingsIdentityProviders {
		var m Model
		m.CompleteByResponse(federatedSettingsIdentityProviders[i])
		mm = append(mm, m)
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModels:  mm,
	}, nil
}
