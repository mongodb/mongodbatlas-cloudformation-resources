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

// OrganizationInvitationGroupRoleAssignmentsRequest struct for OrganizationInvitationGroupRoleAssignmentsRequest
type OrganizationInvitationGroupRoleAssignmentsRequest struct {
	// Unique 24-hexadecimal digit string that identifies the project to which these roles belong.
	GroupId *string `json:"groupId,omitempty"`
	// One or more project level roles to assign to the MongoDB Cloud user.
	Roles *[]string `json:"roles,omitempty"`
}

// NewOrganizationInvitationGroupRoleAssignmentsRequest instantiates a new OrganizationInvitationGroupRoleAssignmentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInvitationGroupRoleAssignmentsRequest() *OrganizationInvitationGroupRoleAssignmentsRequest {
	this := OrganizationInvitationGroupRoleAssignmentsRequest{}
	return &this
}

// NewOrganizationInvitationGroupRoleAssignmentsRequestWithDefaults instantiates a new OrganizationInvitationGroupRoleAssignmentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInvitationGroupRoleAssignmentsRequestWithDefaults() *OrganizationInvitationGroupRoleAssignmentsRequest {
	this := OrganizationInvitationGroupRoleAssignmentsRequest{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) SetGroupId(v string) {
	o.GroupId = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *OrganizationInvitationGroupRoleAssignmentsRequest) SetRoles(v []string) {
	o.Roles = &v
}

func (o OrganizationInvitationGroupRoleAssignmentsRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o OrganizationInvitationGroupRoleAssignmentsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}
