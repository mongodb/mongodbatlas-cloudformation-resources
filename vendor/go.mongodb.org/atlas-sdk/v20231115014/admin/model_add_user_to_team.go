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

// AddUserToTeam struct for AddUserToTeam
type AddUserToTeam struct {
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user.
	Id string `json:"id"`
}

// NewAddUserToTeam instantiates a new AddUserToTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUserToTeam(id string) *AddUserToTeam {
	this := AddUserToTeam{}
	this.Id = id
	return &this
}

// NewAddUserToTeamWithDefaults instantiates a new AddUserToTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUserToTeamWithDefaults() *AddUserToTeam {
	this := AddUserToTeam{}
	return &this
}

// GetId returns the Id field value
func (o *AddUserToTeam) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddUserToTeam) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddUserToTeam) SetId(v string) {
	o.Id = v
}

func (o AddUserToTeam) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AddUserToTeam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}
