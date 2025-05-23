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

// TargetOrg struct for TargetOrg
type TargetOrg struct {
	// Link token that contains all the information required to complete the link.
	LinkToken string `json:"linkToken"`
}

// NewTargetOrg instantiates a new TargetOrg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetOrg(linkToken string) *TargetOrg {
	this := TargetOrg{}
	this.LinkToken = linkToken
	return &this
}

// NewTargetOrgWithDefaults instantiates a new TargetOrg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetOrgWithDefaults() *TargetOrg {
	this := TargetOrg{}
	return &this
}

// GetLinkToken returns the LinkToken field value
func (o *TargetOrg) GetLinkToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkToken
}

// GetLinkTokenOk returns a tuple with the LinkToken field value
// and a boolean to check if the value has been set.
func (o *TargetOrg) GetLinkTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkToken, true
}

// SetLinkToken sets field value
func (o *TargetOrg) SetLinkToken(v string) {
	o.LinkToken = v
}

func (o TargetOrg) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TargetOrg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["linkToken"] = o.LinkToken
	return toSerialize, nil
}
