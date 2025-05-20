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

// LDAPVerifyConnectivityJobRequestValidation One test that MongoDB Cloud runs to test verification of the provided Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration details.
type LDAPVerifyConnectivityJobRequestValidation struct {
	// Human-readable string that indicates the result of this verification test.
	// Read only field.
	Status *string `json:"status,omitempty"`
	// Human-readable label that identifies this verification test that MongoDB Cloud runs.
	// Read only field.
	ValidationType *string `json:"validationType,omitempty"`
}

// NewLDAPVerifyConnectivityJobRequestValidation instantiates a new LDAPVerifyConnectivityJobRequestValidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPVerifyConnectivityJobRequestValidation() *LDAPVerifyConnectivityJobRequestValidation {
	this := LDAPVerifyConnectivityJobRequestValidation{}
	return &this
}

// NewLDAPVerifyConnectivityJobRequestValidationWithDefaults instantiates a new LDAPVerifyConnectivityJobRequestValidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPVerifyConnectivityJobRequestValidationWithDefaults() *LDAPVerifyConnectivityJobRequestValidation {
	this := LDAPVerifyConnectivityJobRequestValidation{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *LDAPVerifyConnectivityJobRequestValidation) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPVerifyConnectivityJobRequestValidation) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LDAPVerifyConnectivityJobRequestValidation) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LDAPVerifyConnectivityJobRequestValidation) SetStatus(v string) {
	o.Status = &v
}

// GetValidationType returns the ValidationType field value if set, zero value otherwise
func (o *LDAPVerifyConnectivityJobRequestValidation) GetValidationType() string {
	if o == nil || IsNil(o.ValidationType) {
		var ret string
		return ret
	}
	return *o.ValidationType
}

// GetValidationTypeOk returns a tuple with the ValidationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPVerifyConnectivityJobRequestValidation) GetValidationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationType) {
		return nil, false
	}

	return o.ValidationType, true
}

// HasValidationType returns a boolean if a field has been set.
func (o *LDAPVerifyConnectivityJobRequestValidation) HasValidationType() bool {
	if o != nil && !IsNil(o.ValidationType) {
		return true
	}

	return false
}

// SetValidationType gets a reference to the given string and assigns it to the ValidationType field.
func (o *LDAPVerifyConnectivityJobRequestValidation) SetValidationType(v string) {
	o.ValidationType = &v
}

func (o LDAPVerifyConnectivityJobRequestValidation) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o LDAPVerifyConnectivityJobRequestValidation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
