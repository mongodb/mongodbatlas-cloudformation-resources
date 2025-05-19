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

// UserSecurity struct for UserSecurity
type UserSecurity struct {
	CustomerX509 *DBUserTLSX509Settings `json:"customerX509,omitempty"`
	Ldap         *LDAPSecuritySettings  `json:"ldap,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
}

// NewUserSecurity instantiates a new UserSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSecurity() *UserSecurity {
	this := UserSecurity{}
	return &this
}

// NewUserSecurityWithDefaults instantiates a new UserSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSecurityWithDefaults() *UserSecurity {
	this := UserSecurity{}
	return &this
}

// GetCustomerX509 returns the CustomerX509 field value if set, zero value otherwise
func (o *UserSecurity) GetCustomerX509() DBUserTLSX509Settings {
	if o == nil || IsNil(o.CustomerX509) {
		var ret DBUserTLSX509Settings
		return ret
	}
	return *o.CustomerX509
}

// GetCustomerX509Ok returns a tuple with the CustomerX509 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSecurity) GetCustomerX509Ok() (*DBUserTLSX509Settings, bool) {
	if o == nil || IsNil(o.CustomerX509) {
		return nil, false
	}

	return o.CustomerX509, true
}

// HasCustomerX509 returns a boolean if a field has been set.
func (o *UserSecurity) HasCustomerX509() bool {
	if o != nil && !IsNil(o.CustomerX509) {
		return true
	}

	return false
}

// SetCustomerX509 gets a reference to the given DBUserTLSX509Settings and assigns it to the CustomerX509 field.
func (o *UserSecurity) SetCustomerX509(v DBUserTLSX509Settings) {
	o.CustomerX509 = &v
}

// GetLdap returns the Ldap field value if set, zero value otherwise
func (o *UserSecurity) GetLdap() LDAPSecuritySettings {
	if o == nil || IsNil(o.Ldap) {
		var ret LDAPSecuritySettings
		return ret
	}
	return *o.Ldap
}

// GetLdapOk returns a tuple with the Ldap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSecurity) GetLdapOk() (*LDAPSecuritySettings, bool) {
	if o == nil || IsNil(o.Ldap) {
		return nil, false
	}

	return o.Ldap, true
}

// HasLdap returns a boolean if a field has been set.
func (o *UserSecurity) HasLdap() bool {
	if o != nil && !IsNil(o.Ldap) {
		return true
	}

	return false
}

// SetLdap gets a reference to the given LDAPSecuritySettings and assigns it to the Ldap field.
func (o *UserSecurity) SetLdap(v LDAPSecuritySettings) {
	o.Ldap = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *UserSecurity) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSecurity) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserSecurity) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *UserSecurity) SetLinks(v []Link) {
	o.Links = &v
}

func (o UserSecurity) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o UserSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerX509) {
		toSerialize["customerX509"] = o.CustomerX509
	}
	if !IsNil(o.Ldap) {
		toSerialize["ldap"] = o.Ldap
	}
	return toSerialize, nil
}
