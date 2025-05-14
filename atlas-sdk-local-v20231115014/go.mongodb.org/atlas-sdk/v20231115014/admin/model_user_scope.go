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

// UserScope Range of resources available to this database user.
type UserScope struct {
	// Human-readable label that identifies the cluster or MongoDB Atlas Data Lake that this database user can access.
	Name string `json:"name"`
	// Category of resource that this database user can access.
	Type string `json:"type"`
}

// NewUserScope instantiates a new UserScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserScope(name string, type_ string) *UserScope {
	this := UserScope{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewUserScopeWithDefaults instantiates a new UserScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserScopeWithDefaults() *UserScope {
	this := UserScope{}
	return &this
}

// GetName returns the Name field value
func (o *UserScope) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserScope) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserScope) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *UserScope) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserScope) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserScope) SetType(v string) {
	o.Type = v
}

func (o UserScope) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o UserScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}
