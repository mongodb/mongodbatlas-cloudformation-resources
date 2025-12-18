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
	"go.mongodb.org/atlas-sdk/v20250312012/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

const (
	ProtocolSAML = "SAML"
	ProtocolOIDC = "OIDC"

	IdpTypeWorkforce = "WORKFORCE"
	IdpTypeWorkload  = "WORKLOAD"
)

var (
	allProtocols = []string{ProtocolSAML, ProtocolOIDC}
	allIdpTypes  = []string{IdpTypeWorkforce, IdpTypeWorkload}
)

func GetFederatedSettingsIdentityProviderModel(api *admin.FederationIdentityProvider, currentModel *Model) *Model {
	var model *Model
	if currentModel != nil {
		model = currentModel
	} else {
		model = &Model{}
	}

	if api == nil {
		return model
	}

	if oktaID, ok := api.GetOktaIdpIdOk(); ok && *oktaID != "" {
		model.OktaIdpId = oktaID
	}

	idpID := api.GetId()
	model.IdpId = &idpID

	displayName := api.GetDisplayName()
	model.Name = &displayName
	issuerURI := api.GetIssuerUri()
	model.IssuerUri = &issuerURI
	protocol := api.GetProtocol()
	model.Protocol = &protocol

	description := api.GetDescription()
	model.Description = &description
	authorizationType := api.GetAuthorizationType()
	model.AuthorizationType = &authorizationType
	idpType := api.GetIdpType()
	model.IdpType = &idpType

	switch protocol {
	case ProtocolSAML:
		requestBinding := api.GetRequestBinding()
		model.RequestBinding = &requestBinding
		responseSignatureAlgorithm := api.GetResponseSignatureAlgorithm()
		model.ResponseSignatureAlgorithm = &responseSignatureAlgorithm
		model.SsoDebugEnabled = api.SsoDebugEnabled
		ssoURL := api.GetSsoUrl()
		model.SsoUrl = &ssoURL
		status := api.GetStatus()
		model.Status = &status

		associatedDomains := api.GetAssociatedDomains()
		if len(associatedDomains) == 0 && currentModel != nil && len(currentModel.AssociatedDomains) > 0 {
			associatedDomains = currentModel.AssociatedDomains
		}
		model.AssociatedDomains = associatedDomains
	case ProtocolOIDC:
		audience := api.GetAudience()
		model.Audience = &audience
		clientID := api.GetClientId()
		model.ClientId = &clientID
		groupsClaim := api.GetGroupsClaim()
		model.GroupsClaim = &groupsClaim

		requestedScopes := api.GetRequestedScopes()
		if len(requestedScopes) == 0 && currentModel != nil && len(currentModel.RequestedScopes) > 0 {
			requestedScopes = currentModel.RequestedScopes
		}
		model.RequestedScopes = requestedScopes

		userClaim := api.GetUserClaim()
		model.UserClaim = &userClaim

		associatedDomains := api.GetAssociatedDomains()
		if len(associatedDomains) == 0 && currentModel != nil && len(currentModel.AssociatedDomains) > 0 {
			associatedDomains = currentModel.AssociatedDomains
		}
		model.AssociatedDomains = associatedDomains
	default:
		return model
	}

	return model
}

func ExpandOIDCCreateRequest(model *Model) *admin.FederationOidcIdentityProviderUpdate {
	var associatedDomains []string
	if model.AssociatedDomains != nil {
		associatedDomains = model.AssociatedDomains
	} else {
		associatedDomains = []string{}
	}
	var requestedScopes []string
	if model.RequestedScopes != nil {
		requestedScopes = model.RequestedScopes
	} else {
		requestedScopes = []string{}
	}

	return &admin.FederationOidcIdentityProviderUpdate{
		Audience:          util.Pointer(util.SafeString(model.Audience)),
		AssociatedDomains: &associatedDomains,
		AuthorizationType: util.Pointer(util.SafeString(model.AuthorizationType)),
		ClientId:          util.Pointer(util.SafeString(model.ClientId)),
		Description:       util.Pointer(util.SafeString(model.Description)),
		DisplayName:       util.Pointer(util.SafeString(model.Name)),
		GroupsClaim:       util.Pointer(util.SafeString(model.GroupsClaim)),
		IdpType:           util.Pointer(util.SafeString(model.IdpType)),
		IssuerUri:         util.Pointer(util.SafeString(model.IssuerUri)),
		Protocol:          util.Pointer(util.SafeString(model.Protocol)),
		RequestedScopes:   &requestedScopes,
		UserClaim:         util.Pointer(util.SafeString(model.UserClaim)),
	}
}
