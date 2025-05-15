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

// AdvancedDiskBackupSnapshotSchedulePolicy List that contains a document for each backup policy item in the desired backup policy.
type AdvancedDiskBackupSnapshotSchedulePolicy struct {
	// Unique 24-hexadecimal digit string that identifies this backup policy.
	Id *string `json:"id,omitempty"`
	// List that contains the specifications for one policy.
	PolicyItems *[]DiskBackupApiPolicyItem `json:"policyItems,omitempty"`
}

// NewAdvancedDiskBackupSnapshotSchedulePolicy instantiates a new AdvancedDiskBackupSnapshotSchedulePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedDiskBackupSnapshotSchedulePolicy() *AdvancedDiskBackupSnapshotSchedulePolicy {
	this := AdvancedDiskBackupSnapshotSchedulePolicy{}
	return &this
}

// NewAdvancedDiskBackupSnapshotSchedulePolicyWithDefaults instantiates a new AdvancedDiskBackupSnapshotSchedulePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedDiskBackupSnapshotSchedulePolicyWithDefaults() *AdvancedDiskBackupSnapshotSchedulePolicy {
	this := AdvancedDiskBackupSnapshotSchedulePolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *AdvancedDiskBackupSnapshotSchedulePolicy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedDiskBackupSnapshotSchedulePolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvancedDiskBackupSnapshotSchedulePolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdvancedDiskBackupSnapshotSchedulePolicy) SetId(v string) {
	o.Id = &v
}

// GetPolicyItems returns the PolicyItems field value if set, zero value otherwise
func (o *AdvancedDiskBackupSnapshotSchedulePolicy) GetPolicyItems() []DiskBackupApiPolicyItem {
	if o == nil || IsNil(o.PolicyItems) {
		var ret []DiskBackupApiPolicyItem
		return ret
	}
	return *o.PolicyItems
}

// GetPolicyItemsOk returns a tuple with the PolicyItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedDiskBackupSnapshotSchedulePolicy) GetPolicyItemsOk() (*[]DiskBackupApiPolicyItem, bool) {
	if o == nil || IsNil(o.PolicyItems) {
		return nil, false
	}

	return o.PolicyItems, true
}

// HasPolicyItems returns a boolean if a field has been set.
func (o *AdvancedDiskBackupSnapshotSchedulePolicy) HasPolicyItems() bool {
	if o != nil && !IsNil(o.PolicyItems) {
		return true
	}

	return false
}

// SetPolicyItems gets a reference to the given []DiskBackupApiPolicyItem and assigns it to the PolicyItems field.
func (o *AdvancedDiskBackupSnapshotSchedulePolicy) SetPolicyItems(v []DiskBackupApiPolicyItem) {
	o.PolicyItems = &v
}

func (o AdvancedDiskBackupSnapshotSchedulePolicy) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AdvancedDiskBackupSnapshotSchedulePolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PolicyItems) {
		toSerialize["policyItems"] = o.PolicyItems
	}
	return toSerialize, nil
}
