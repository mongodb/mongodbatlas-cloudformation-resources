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

// ResourceTag Key-value pair that tags and categorizes a MongoDB Cloud organization, project, or cluster. For example, `environment : production`.
type ResourceTag struct {
	// Constant that defines the set of the tag. For example, `environment` in the `environment : production` tag.
	Key string `json:"key"`
	// Variable that belongs to the set of the tag. For example, `production` in the `environment : production` tag.
	Value string `json:"value"`
}

// NewResourceTag instantiates a new ResourceTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTag(key string, value string) *ResourceTag {
	this := ResourceTag{}
	this.Key = key
	this.Value = value
	return &this
}

// NewResourceTagWithDefaults instantiates a new ResourceTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTagWithDefaults() *ResourceTag {
	this := ResourceTag{}
	return &this
}

// GetKey returns the Key field value
func (o *ResourceTag) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ResourceTag) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ResourceTag) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *ResourceTag) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ResourceTag) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ResourceTag) SetValue(v string) {
	o.Value = v
}

func (o ResourceTag) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ResourceTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}
