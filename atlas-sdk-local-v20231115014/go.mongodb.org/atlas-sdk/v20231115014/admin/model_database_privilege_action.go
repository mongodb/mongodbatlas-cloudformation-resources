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

// DatabasePrivilegeAction Privilege action that the role grants.
type DatabasePrivilegeAction struct {
	// Human-readable label that identifies the privilege action.
	Action string `json:"action"`
	// List of resources on which you grant the action.
	Resources *[]DatabasePermittedNamespaceResource `json:"resources,omitempty"`
}

// NewDatabasePrivilegeAction instantiates a new DatabasePrivilegeAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabasePrivilegeAction(action string) *DatabasePrivilegeAction {
	this := DatabasePrivilegeAction{}
	this.Action = action
	return &this
}

// NewDatabasePrivilegeActionWithDefaults instantiates a new DatabasePrivilegeAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabasePrivilegeActionWithDefaults() *DatabasePrivilegeAction {
	this := DatabasePrivilegeAction{}
	return &this
}

// GetAction returns the Action field value
func (o *DatabasePrivilegeAction) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *DatabasePrivilegeAction) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *DatabasePrivilegeAction) SetAction(v string) {
	o.Action = v
}

// GetResources returns the Resources field value if set, zero value otherwise
func (o *DatabasePrivilegeAction) GetResources() []DatabasePermittedNamespaceResource {
	if o == nil || IsNil(o.Resources) {
		var ret []DatabasePermittedNamespaceResource
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabasePrivilegeAction) GetResourcesOk() (*[]DatabasePermittedNamespaceResource, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}

	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *DatabasePrivilegeAction) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []DatabasePermittedNamespaceResource and assigns it to the Resources field.
func (o *DatabasePrivilegeAction) SetResources(v []DatabasePermittedNamespaceResource) {
	o.Resources = &v
}

func (o DatabasePrivilegeAction) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DatabasePrivilegeAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}
