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

// OrganizationSettings Collection of settings that configures the organization.
type OrganizationSettings struct {
	// Flag that indicates whether to require API operations to originate from an IP Address added to the API access list for the specified organization.
	ApiAccessListRequired *bool `json:"apiAccessListRequired,omitempty"`
	// Flag that indicates whether to require users to set up Multi-Factor Authentication (MFA) before accessing the specified organization. To learn more, see: https://www.mongodb.com/docs/atlas/security-multi-factor-authentication/.
	MultiFactorAuthRequired *bool `json:"multiFactorAuthRequired,omitempty"`
	// Flag that indicates whether to block MongoDB Support from accessing Atlas infrastructure for any deployment in the specified organization without explicit permission. Once this setting is turned on, you can grant MongoDB Support a 24-hour bypass access to the Atlas deployment to resolve support issues. To learn more, see: https://www.mongodb.com/docs/atlas/security-restrict-support-access/.
	RestrictEmployeeAccess *bool `json:"restrictEmployeeAccess,omitempty"`
}

// NewOrganizationSettings instantiates a new OrganizationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSettings() *OrganizationSettings {
	this := OrganizationSettings{}
	return &this
}

// NewOrganizationSettingsWithDefaults instantiates a new OrganizationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSettingsWithDefaults() *OrganizationSettings {
	this := OrganizationSettings{}
	return &this
}

// GetApiAccessListRequired returns the ApiAccessListRequired field value if set, zero value otherwise
func (o *OrganizationSettings) GetApiAccessListRequired() bool {
	if o == nil || IsNil(o.ApiAccessListRequired) {
		var ret bool
		return ret
	}
	return *o.ApiAccessListRequired
}

// GetApiAccessListRequiredOk returns a tuple with the ApiAccessListRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettings) GetApiAccessListRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ApiAccessListRequired) {
		return nil, false
	}

	return o.ApiAccessListRequired, true
}

// HasApiAccessListRequired returns a boolean if a field has been set.
func (o *OrganizationSettings) HasApiAccessListRequired() bool {
	if o != nil && !IsNil(o.ApiAccessListRequired) {
		return true
	}

	return false
}

// SetApiAccessListRequired gets a reference to the given bool and assigns it to the ApiAccessListRequired field.
func (o *OrganizationSettings) SetApiAccessListRequired(v bool) {
	o.ApiAccessListRequired = &v
}

// GetMultiFactorAuthRequired returns the MultiFactorAuthRequired field value if set, zero value otherwise
func (o *OrganizationSettings) GetMultiFactorAuthRequired() bool {
	if o == nil || IsNil(o.MultiFactorAuthRequired) {
		var ret bool
		return ret
	}
	return *o.MultiFactorAuthRequired
}

// GetMultiFactorAuthRequiredOk returns a tuple with the MultiFactorAuthRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettings) GetMultiFactorAuthRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiFactorAuthRequired) {
		return nil, false
	}

	return o.MultiFactorAuthRequired, true
}

// HasMultiFactorAuthRequired returns a boolean if a field has been set.
func (o *OrganizationSettings) HasMultiFactorAuthRequired() bool {
	if o != nil && !IsNil(o.MultiFactorAuthRequired) {
		return true
	}

	return false
}

// SetMultiFactorAuthRequired gets a reference to the given bool and assigns it to the MultiFactorAuthRequired field.
func (o *OrganizationSettings) SetMultiFactorAuthRequired(v bool) {
	o.MultiFactorAuthRequired = &v
}

// GetRestrictEmployeeAccess returns the RestrictEmployeeAccess field value if set, zero value otherwise
func (o *OrganizationSettings) GetRestrictEmployeeAccess() bool {
	if o == nil || IsNil(o.RestrictEmployeeAccess) {
		var ret bool
		return ret
	}
	return *o.RestrictEmployeeAccess
}

// GetRestrictEmployeeAccessOk returns a tuple with the RestrictEmployeeAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettings) GetRestrictEmployeeAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictEmployeeAccess) {
		return nil, false
	}

	return o.RestrictEmployeeAccess, true
}

// HasRestrictEmployeeAccess returns a boolean if a field has been set.
func (o *OrganizationSettings) HasRestrictEmployeeAccess() bool {
	if o != nil && !IsNil(o.RestrictEmployeeAccess) {
		return true
	}

	return false
}

// SetRestrictEmployeeAccess gets a reference to the given bool and assigns it to the RestrictEmployeeAccess field.
func (o *OrganizationSettings) SetRestrictEmployeeAccess(v bool) {
	o.RestrictEmployeeAccess = &v
}

func (o OrganizationSettings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o OrganizationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiAccessListRequired) {
		toSerialize["apiAccessListRequired"] = o.ApiAccessListRequired
	}
	if !IsNil(o.MultiFactorAuthRequired) {
		toSerialize["multiFactorAuthRequired"] = o.MultiFactorAuthRequired
	}
	if !IsNil(o.RestrictEmployeeAccess) {
		toSerialize["restrictEmployeeAccess"] = o.RestrictEmployeeAccess
	}
	return toSerialize, nil
}
