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

// TargetOrgRequest struct for TargetOrgRequest
type TargetOrgRequest struct {
	// IP address access list entries associated with the API key.
	AccessListIps *[]string `json:"accessListIps,omitempty"`
}

// NewTargetOrgRequest instantiates a new TargetOrgRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetOrgRequest() *TargetOrgRequest {
	this := TargetOrgRequest{}
	return &this
}

// NewTargetOrgRequestWithDefaults instantiates a new TargetOrgRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetOrgRequestWithDefaults() *TargetOrgRequest {
	this := TargetOrgRequest{}
	return &this
}

// GetAccessListIps returns the AccessListIps field value if set, zero value otherwise
func (o *TargetOrgRequest) GetAccessListIps() []string {
	if o == nil || IsNil(o.AccessListIps) {
		var ret []string
		return ret
	}
	return *o.AccessListIps
}

// GetAccessListIpsOk returns a tuple with the AccessListIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetOrgRequest) GetAccessListIpsOk() (*[]string, bool) {
	if o == nil || IsNil(o.AccessListIps) {
		return nil, false
	}

	return o.AccessListIps, true
}

// HasAccessListIps returns a boolean if a field has been set.
func (o *TargetOrgRequest) HasAccessListIps() bool {
	if o != nil && !IsNil(o.AccessListIps) {
		return true
	}

	return false
}

// SetAccessListIps gets a reference to the given []string and assigns it to the AccessListIps field.
func (o *TargetOrgRequest) SetAccessListIps(v []string) {
	o.AccessListIps = &v
}

func (o TargetOrgRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TargetOrgRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessListIps) {
		toSerialize["accessListIps"] = o.AccessListIps
	}
	return toSerialize, nil
}
