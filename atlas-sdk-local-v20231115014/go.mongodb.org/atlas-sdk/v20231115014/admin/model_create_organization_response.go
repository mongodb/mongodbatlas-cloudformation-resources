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

// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// CreateOrganizationResponse struct for CreateOrganizationResponse
type CreateOrganizationResponse struct {
	ApiKey *ApiKeyUserDetails `json:"apiKey,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the federation that you linked the newly created organization to.
	// Read only field.
	FederationSettingsId *string `json:"federationSettingsId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user that you assigned the Organization Owner role in the new organization.
	// Read only field.
	OrgOwnerId   *string            `json:"orgOwnerId,omitempty"`
	Organization *AtlasOrganization `json:"organization,omitempty"`
}

// NewCreateOrganizationResponse instantiates a new CreateOrganizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationResponse() *CreateOrganizationResponse {
	this := CreateOrganizationResponse{}
	return &this
}

// NewCreateOrganizationResponseWithDefaults instantiates a new CreateOrganizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationResponseWithDefaults() *CreateOrganizationResponse {
	this := CreateOrganizationResponse{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise
func (o *CreateOrganizationResponse) GetApiKey() ApiKeyUserDetails {
	if o == nil || IsNil(o.ApiKey) {
		var ret ApiKeyUserDetails
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationResponse) GetApiKeyOk() (*ApiKeyUserDetails, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}

	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *CreateOrganizationResponse) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given ApiKeyUserDetails and assigns it to the ApiKey field.
func (o *CreateOrganizationResponse) SetApiKey(v ApiKeyUserDetails) {
	o.ApiKey = &v
}

// GetFederationSettingsId returns the FederationSettingsId field value if set, zero value otherwise
func (o *CreateOrganizationResponse) GetFederationSettingsId() string {
	if o == nil || IsNil(o.FederationSettingsId) {
		var ret string
		return ret
	}
	return *o.FederationSettingsId
}

// GetFederationSettingsIdOk returns a tuple with the FederationSettingsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationResponse) GetFederationSettingsIdOk() (*string, bool) {
	if o == nil || IsNil(o.FederationSettingsId) {
		return nil, false
	}

	return o.FederationSettingsId, true
}

// HasFederationSettingsId returns a boolean if a field has been set.
func (o *CreateOrganizationResponse) HasFederationSettingsId() bool {
	if o != nil && !IsNil(o.FederationSettingsId) {
		return true
	}

	return false
}

// SetFederationSettingsId gets a reference to the given string and assigns it to the FederationSettingsId field.
func (o *CreateOrganizationResponse) SetFederationSettingsId(v string) {
	o.FederationSettingsId = &v
}

// GetOrgOwnerId returns the OrgOwnerId field value if set, zero value otherwise
func (o *CreateOrganizationResponse) GetOrgOwnerId() string {
	if o == nil || IsNil(o.OrgOwnerId) {
		var ret string
		return ret
	}
	return *o.OrgOwnerId
}

// GetOrgOwnerIdOk returns a tuple with the OrgOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationResponse) GetOrgOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgOwnerId) {
		return nil, false
	}

	return o.OrgOwnerId, true
}

// HasOrgOwnerId returns a boolean if a field has been set.
func (o *CreateOrganizationResponse) HasOrgOwnerId() bool {
	if o != nil && !IsNil(o.OrgOwnerId) {
		return true
	}

	return false
}

// SetOrgOwnerId gets a reference to the given string and assigns it to the OrgOwnerId field.
func (o *CreateOrganizationResponse) SetOrgOwnerId(v string) {
	o.OrgOwnerId = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise
func (o *CreateOrganizationResponse) GetOrganization() AtlasOrganization {
	if o == nil || IsNil(o.Organization) {
		var ret AtlasOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationResponse) GetOrganizationOk() (*AtlasOrganization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}

	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CreateOrganizationResponse) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given AtlasOrganization and assigns it to the Organization field.
func (o *CreateOrganizationResponse) SetOrganization(v AtlasOrganization) {
	o.Organization = &v
}

func (o CreateOrganizationResponse) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateOrganizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}
