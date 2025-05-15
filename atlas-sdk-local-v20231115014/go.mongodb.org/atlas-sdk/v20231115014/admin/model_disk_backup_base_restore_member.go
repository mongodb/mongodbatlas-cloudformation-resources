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

// DiskBackupBaseRestoreMember struct for DiskBackupBaseRestoreMember
type DiskBackupBaseRestoreMember struct {
	// Human-readable label that identifies the replica set on the sharded cluster.
	// Read only field.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
}

// NewDiskBackupBaseRestoreMember instantiates a new DiskBackupBaseRestoreMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupBaseRestoreMember() *DiskBackupBaseRestoreMember {
	this := DiskBackupBaseRestoreMember{}
	return &this
}

// NewDiskBackupBaseRestoreMemberWithDefaults instantiates a new DiskBackupBaseRestoreMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupBaseRestoreMemberWithDefaults() *DiskBackupBaseRestoreMember {
	this := DiskBackupBaseRestoreMember{}
	return &this
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise
func (o *DiskBackupBaseRestoreMember) GetReplicaSetName() string {
	if o == nil || IsNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupBaseRestoreMember) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetName) {
		return nil, false
	}

	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *DiskBackupBaseRestoreMember) HasReplicaSetName() bool {
	if o != nil && !IsNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *DiskBackupBaseRestoreMember) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

func (o DiskBackupBaseRestoreMember) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DiskBackupBaseRestoreMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
