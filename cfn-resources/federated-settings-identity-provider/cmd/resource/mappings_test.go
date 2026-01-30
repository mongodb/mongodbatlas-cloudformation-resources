// Copyright 2025 MongoDB Inc
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

package resource_test

import (
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/federated-settings-identity-provider/cmd/resource"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20250312012/admin"
)

const (
	testIdpID  = "test-idp-id"
	testOktaID = "test-okta-id"
)

func TestGetFederatedSettingsIdentityProviderModel_SAML(t *testing.T) {
	protocol := "SAML"
	displayName := "saml-name"
	issuerURI := "https://issuer.example.com"
	api := &admin.FederationIdentityProvider{
		Id:                         testIdpID,
		OktaIdpId:                  testOktaID,
		Protocol:                   &protocol,
		DisplayName:                &displayName,
		IssuerUri:                  &issuerURI,
		RequestBinding:             func() *string { s := "HTTP-POST"; return &s }(),
		ResponseSignatureAlgorithm: func() *string { s := "RSA-SHA256"; return &s }(),
		SsoDebugEnabled:            func() *bool { b := true; return &b }(),
		SsoUrl:                     func() *string { s := "https://sso.example.com"; return &s }(),
		Status:                     func() *string { s := "ACTIVE"; return &s }(),
		AssociatedDomains:          func() *[]string { s := []string{"example.com"}; return &s }(),
	}

	model := resource.GetFederatedSettingsIdentityProviderModel(api, &resource.Model{})

	assert.Equal(t, testIdpID, *model.IdpId)
	assert.Equal(t, testOktaID, *model.OktaIdpId)
	assert.Equal(t, "SAML", *model.Protocol)
	assert.Equal(t, "saml-name", *model.Name)
	assert.Equal(t, "https://issuer.example.com", *model.IssuerUri)
	assert.Equal(t, "HTTP-POST", *model.RequestBinding)
	assert.Equal(t, "RSA-SHA256", *model.ResponseSignatureAlgorithm)
	assert.True(t, *model.SsoDebugEnabled)
	assert.Equal(t, "https://sso.example.com", *model.SsoUrl)
	assert.Equal(t, "ACTIVE", *model.Status)
	assert.Equal(t, []string{"example.com"}, model.AssociatedDomains)
	// OIDC-only fields should not be set by the SAML branch
	assert.Nil(t, model.ClientId)
	assert.Nil(t, model.UserClaim)
}

func TestGetFederatedSettingsIdentityProviderModel_OIDC(t *testing.T) {
	protocol := "OIDC"
	displayName := "oidc-name"
	issuerURI := "https://issuer.oidc.example.com"
	api := &admin.FederationIdentityProvider{
		Id:                testIdpID,
		OktaIdpId:         testOktaID,
		Protocol:          &protocol,
		DisplayName:       &displayName,
		IssuerUri:         &issuerURI,
		Audience:          func() *string { s := "aud"; return &s }(),
		ClientId:          func() *string { s := "client"; return &s }(),
		GroupsClaim:       func() *string { s := "groups"; return &s }(),
		RequestedScopes:   func() *[]string { s := []string{"openid", "profile"}; return &s }(),
		UserClaim:         func() *string { s := "sub"; return &s }(),
		AssociatedDomains: func() *[]string { s := []string{"oidc.example.com"}; return &s }(),
	}

	model := resource.GetFederatedSettingsIdentityProviderModel(api, &resource.Model{})

	assert.Equal(t, testIdpID, *model.IdpId)
	assert.Equal(t, testOktaID, *model.OktaIdpId)
	assert.Equal(t, "OIDC", *model.Protocol)
	assert.Equal(t, "oidc-name", *model.Name)
	assert.Equal(t, "https://issuer.oidc.example.com", *model.IssuerUri)
	assert.Equal(t, "aud", *model.Audience)
	assert.Equal(t, "client", *model.ClientId)
	assert.Equal(t, "groups", *model.GroupsClaim)
	assert.Equal(t, []string{"openid", "profile"}, model.RequestedScopes)
	assert.Equal(t, "sub", *model.UserClaim)
	assert.Equal(t, []string{"oidc.example.com"}, model.AssociatedDomains)
	// SAML-only fields should not be set by the OIDC branch
	assert.Nil(t, model.RequestBinding)
	assert.Nil(t, model.SsoUrl)
}

func TestExpandOIDCCreateRequest_DefaultSlices(t *testing.T) {
	protocol := resource.ProtocolOIDC
	name := "n"
	issuer := "i"
	m := &resource.Model{
		Protocol:  &protocol,
		Name:      &name,
		IssuerUri: &issuer,
		// AssociatedDomains and RequestedScopes intentionally nil to exercise defaults
	}

	req := resource.ExpandOIDCCreateRequest(m)
	assert.NotNil(t, req)
	assert.NotNil(t, req.AssociatedDomains)
	assert.NotNil(t, req.RequestedScopes)
	assert.Empty(t, *req.AssociatedDomains)
	assert.Empty(t, *req.RequestedScopes)
}
