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

// UpdateAtlasOrganizationApiKey struct for UpdateAtlasOrganizationApiKey
type UpdateAtlasOrganizationApiKey struct {
	// Purpose or explanation provided when someone creates this organization API key.
	Desc *string `json:"desc,omitempty"`
	// List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this organization.
	Roles *[]string `json:"roles,omitempty"`
}

// NewUpdateAtlasOrganizationApiKey instantiates a new UpdateAtlasOrganizationApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAtlasOrganizationApiKey() *UpdateAtlasOrganizationApiKey {
	this := UpdateAtlasOrganizationApiKey{}
	return &this
}

// NewUpdateAtlasOrganizationApiKeyWithDefaults instantiates a new UpdateAtlasOrganizationApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAtlasOrganizationApiKeyWithDefaults() *UpdateAtlasOrganizationApiKey {
	this := UpdateAtlasOrganizationApiKey{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise
func (o *UpdateAtlasOrganizationApiKey) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAtlasOrganizationApiKey) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}

	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *UpdateAtlasOrganizationApiKey) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *UpdateAtlasOrganizationApiKey) SetDesc(v string) {
	o.Desc = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *UpdateAtlasOrganizationApiKey) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAtlasOrganizationApiKey) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UpdateAtlasOrganizationApiKey) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UpdateAtlasOrganizationApiKey) SetRoles(v []string) {
	o.Roles = &v
}

func (o UpdateAtlasOrganizationApiKey) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o UpdateAtlasOrganizationApiKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}
