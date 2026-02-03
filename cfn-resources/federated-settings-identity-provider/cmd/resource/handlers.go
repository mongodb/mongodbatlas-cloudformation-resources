// Copyright 2026 MongoDB Inc
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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

func HandleCreate(client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	if currentModel.Protocol == nil || *currentModel.Protocol != ProtocolOIDC {
		return progressevent.GetFailedEventByCode(
			fmt.Sprintf("create is only supported by %s, %s must be imported", ProtocolOIDC, ProtocolSAML),
			string(types.HandlerErrorCodeInvalidRequest),
		)
	}

	federationSettingsID := util.SafeString(currentModel.FederationSettingsId)

	createRequest := ExpandOIDCCreateRequest(currentModel)
	created, res, err := client.AtlasSDK.FederatedAuthenticationApi.CreateIdentityProvider(context.Background(), federationSettingsID, createRequest).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(
			fmt.Sprintf("error creating federation settings identity provider (%s): %s", federationSettingsID, err.Error()),
			res,
		)
	}

	currentModel.IdpId = util.Pointer(created.GetId())

	return HandleRead(client, currentModel)
}

func HandleRead(client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	federationSettingsID := util.SafeString(currentModel.FederationSettingsId)
	idpID := util.SafeString(currentModel.IdpId)

	idp, res, err := client.AtlasSDK.FederatedAuthenticationApi.GetIdentityProvider(context.Background(), federationSettingsID, idpID).Execute()
	if err != nil {
		if util.StatusNotFound(res) {
			return progressevent.GetFailedEventByCode("Resource not found", string(types.HandlerErrorCodeNotFound))
		}
		return progressevent.GetFailedEventByResponse(
			fmt.Sprintf("error getting federation settings identity provider (%s/%s): %s", federationSettingsID, idpID, err.Error()),
			res,
		)
	}

	model := GetFederatedSettingsIdentityProviderModel(idp, currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   model,
	}
}

func HandleUpdate(client *util.MongoDBClient, prevModel *Model, currentModel *Model) handler.ProgressEvent {
	federationSettingsID := util.SafeString(currentModel.FederationSettingsId)
	idpID := util.SafeString(currentModel.IdpId)

	associatedDomains := getStringSliceOrEmpty(currentModel.AssociatedDomains)
	requestedScopes := getStringSliceOrEmpty(currentModel.RequestedScopes)

	updateReq := &admin.FederationIdentityProviderUpdate{
		AssociatedDomains:          &associatedDomains,
		Audience:                   currentModel.Audience,
		AuthorizationType:          currentModel.AuthorizationType,
		ClientId:                   currentModel.ClientId,
		Description:                currentModel.Description,
		DisplayName:                currentModel.Name,
		GroupsClaim:                currentModel.GroupsClaim,
		IdpType:                    currentModel.IdpType,
		IssuerUri:                  currentModel.IssuerUri,
		Protocol:                   currentModel.Protocol,
		PemFileInfo:                nil,
		RequestBinding:             currentModel.RequestBinding,
		RequestedScopes:            &requestedScopes,
		ResponseSignatureAlgorithm: currentModel.ResponseSignatureAlgorithm,
		SsoDebugEnabled:            currentModel.SsoDebugEnabled,
		SsoUrl:                     currentModel.SsoUrl,
		Status:                     currentModel.Status,
		UserClaim:                  currentModel.UserClaim,
	}

	updated, updRes, err := client.AtlasSDK.FederatedAuthenticationApi.UpdateIdentityProvider(context.Background(), federationSettingsID, idpID, updateReq).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(
			fmt.Sprintf("error updating federation settings identity provider (%s): %s", federationSettingsID, err.Error()),
			updRes,
		)
	}

	model := GetFederatedSettingsIdentityProviderModel(updated, currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   model,
	}
}

func HandleDelete(client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	federationSettingsID := util.SafeString(currentModel.FederationSettingsId)
	idpID := util.SafeString(currentModel.IdpId)

	res, err := client.AtlasSDK.FederatedAuthenticationApi.DeleteIdentityProvider(context.Background(), federationSettingsID, idpID).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(
			fmt.Sprintf("error deleting federation settings identity provider (%s): %s, error: %s", federationSettingsID, idpID, err.Error()),
			res,
		)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}
}

func HandleList(client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	federationSettingsID := util.SafeString(currentModel.FederationSettingsId)

	params := &admin.ListIdentityProvidersApiParams{
		FederationSettingsId: federationSettingsID,
		Protocol:             &allProtocols,
		IdpType:              &allIdpTypes,
	}
	providers, res, err := client.AtlasSDK.FederatedAuthenticationApi.ListIdentityProvidersWithParams(context.Background(), params).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(
			fmt.Sprintf("error listing federation settings identity providers (%s): %s", federationSettingsID, err.Error()),
			res,
		)
	}

	results := providers.GetResults()
	models := make([]any, 0, len(results))
	for i := range results {
		m := &Model{
			Profile:              currentModel.Profile,
			FederationSettingsId: currentModel.FederationSettingsId,
		}
		models = append(models, GetFederatedSettingsIdentityProviderModel(&results[i], m))
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}
}
