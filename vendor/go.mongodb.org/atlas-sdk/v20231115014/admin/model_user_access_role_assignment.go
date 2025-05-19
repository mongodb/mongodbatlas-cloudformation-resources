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

// UserAccessRoleAssignment struct for UserAccessRoleAssignment
type UserAccessRoleAssignment struct {
	// Unique 24-hexadecimal digit string that identifies the organization API key.
	// Read only field.
	ApiUserId *string `json:"apiUserId,omitempty"`
	// List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this project.
	Roles *[]string `json:"roles,omitempty"`
}

// NewUserAccessRoleAssignment instantiates a new UserAccessRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccessRoleAssignment() *UserAccessRoleAssignment {
	this := UserAccessRoleAssignment{}
	return &this
}

// NewUserAccessRoleAssignmentWithDefaults instantiates a new UserAccessRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccessRoleAssignmentWithDefaults() *UserAccessRoleAssignment {
	this := UserAccessRoleAssignment{}
	return &this
}

// GetApiUserId returns the ApiUserId field value if set, zero value otherwise
func (o *UserAccessRoleAssignment) GetApiUserId() string {
	if o == nil || IsNil(o.ApiUserId) {
		var ret string
		return ret
	}
	return *o.ApiUserId
}

// GetApiUserIdOk returns a tuple with the ApiUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessRoleAssignment) GetApiUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiUserId) {
		return nil, false
	}

	return o.ApiUserId, true
}

// HasApiUserId returns a boolean if a field has been set.
func (o *UserAccessRoleAssignment) HasApiUserId() bool {
	if o != nil && !IsNil(o.ApiUserId) {
		return true
	}

	return false
}

// SetApiUserId gets a reference to the given string and assigns it to the ApiUserId field.
func (o *UserAccessRoleAssignment) SetApiUserId(v string) {
	o.ApiUserId = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *UserAccessRoleAssignment) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessRoleAssignment) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserAccessRoleAssignment) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UserAccessRoleAssignment) SetRoles(v []string) {
	o.Roles = &v
}

func (o UserAccessRoleAssignment) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o UserAccessRoleAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}
