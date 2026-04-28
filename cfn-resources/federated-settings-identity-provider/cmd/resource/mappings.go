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
	"go.mongodb.org/atlas-sdk/v20250312013/admin"

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

func getStringSliceOrEmpty(slice []string) []string {
	if slice != nil {
		return slice
	}
	return []string{}
}

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

	model.IdpId = new(api.GetId())
	model.Name = new(api.GetDisplayName())
	model.IssuerUri = new(api.GetIssuerUri())
	model.Protocol = new(api.GetProtocol())
	model.Description = new(api.GetDescription())
	model.AuthorizationType = new(api.GetAuthorizationType())
	model.IdpType = new(api.GetIdpType())

	protocol := api.GetProtocol()
	switch protocol {
	case ProtocolSAML:
		model.RequestBinding = new(api.GetRequestBinding())
		model.ResponseSignatureAlgorithm = new(api.GetResponseSignatureAlgorithm())
		model.SsoDebugEnabled = api.SsoDebugEnabled
		model.SsoUrl = new(api.GetSsoUrl())
		model.Status = new(api.GetStatus())

		associatedDomains := api.GetAssociatedDomains()
		if len(associatedDomains) == 0 && currentModel != nil && len(currentModel.AssociatedDomains) > 0 {
			associatedDomains = currentModel.AssociatedDomains
		}
		model.AssociatedDomains = associatedDomains
	case ProtocolOIDC:
		model.Audience = new(api.GetAudience())
		model.ClientId = new(api.GetClientId())
		model.GroupsClaim = new(api.GetGroupsClaim())

		requestedScopes := api.GetRequestedScopes()
		if len(requestedScopes) == 0 && currentModel != nil && len(currentModel.RequestedScopes) > 0 {
			requestedScopes = currentModel.RequestedScopes
		}
		model.RequestedScopes = requestedScopes

		model.UserClaim = new(api.GetUserClaim())

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
	associatedDomains := getStringSliceOrEmpty(model.AssociatedDomains)
	requestedScopes := getStringSliceOrEmpty(model.RequestedScopes)

	return &admin.FederationOidcIdentityProviderUpdate{
		Audience:          new(util.SafeString(model.Audience)),
		AssociatedDomains: &associatedDomains,
		AuthorizationType: new(util.SafeString(model.AuthorizationType)),
		ClientId:          new(util.SafeString(model.ClientId)),
		Description:       new(util.SafeString(model.Description)),
		DisplayName:       new(util.SafeString(model.Name)),
		GroupsClaim:       new(util.SafeString(model.GroupsClaim)),
		IdpType:           new(util.SafeString(model.IdpType)),
		IssuerUri:         new(util.SafeString(model.IssuerUri)),
		Protocol:          new(util.SafeString(model.Protocol)),
		RequestedScopes:   &requestedScopes,
		UserClaim:         new(util.SafeString(model.UserClaim)),
	}
}
