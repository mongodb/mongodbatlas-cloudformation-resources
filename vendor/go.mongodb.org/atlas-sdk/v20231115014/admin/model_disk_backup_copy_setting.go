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

// DiskBackupCopySetting Copy setting item in the desired backup policy.
type DiskBackupCopySetting struct {
	// Human-readable label that identifies the cloud provider that stores the snapshot copy.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// List that describes which types of snapshots to copy.
	Frequencies *[]string `json:"frequencies,omitempty"`
	// Target region to copy snapshots belonging to replicationSpecId to. Please supply the 'Atlas Region' which can be found under [Cloud Providers](https://www.mongodb.com/docs/atlas/reference/cloud-providers/) 'regions' link.
	RegionName *string `json:"regionName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Replication Spec Id, do a GET request to Return One Cluster in One Project and consult the replicationSpecs array [Return One Cluster in One Project](#operation/getLegacyCluster).
	ReplicationSpecId *string `json:"replicationSpecId,omitempty"`
	// Flag that indicates whether to copy the oplogs to the target region. You can use the oplogs to perform point-in-time restores.
	ShouldCopyOplogs *bool `json:"shouldCopyOplogs,omitempty"`
}

// NewDiskBackupCopySetting instantiates a new DiskBackupCopySetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupCopySetting() *DiskBackupCopySetting {
	this := DiskBackupCopySetting{}
	return &this
}

// NewDiskBackupCopySettingWithDefaults instantiates a new DiskBackupCopySetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupCopySettingWithDefaults() *DiskBackupCopySetting {
	this := DiskBackupCopySetting{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *DiskBackupCopySetting) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupCopySetting) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *DiskBackupCopySetting) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *DiskBackupCopySetting) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetFrequencies returns the Frequencies field value if set, zero value otherwise
func (o *DiskBackupCopySetting) GetFrequencies() []string {
	if o == nil || IsNil(o.Frequencies) {
		var ret []string
		return ret
	}
	return *o.Frequencies
}

// GetFrequenciesOk returns a tuple with the Frequencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupCopySetting) GetFrequenciesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Frequencies) {
		return nil, false
	}

	return o.Frequencies, true
}

// HasFrequencies returns a boolean if a field has been set.
func (o *DiskBackupCopySetting) HasFrequencies() bool {
	if o != nil && !IsNil(o.Frequencies) {
		return true
	}

	return false
}

// SetFrequencies gets a reference to the given []string and assigns it to the Frequencies field.
func (o *DiskBackupCopySetting) SetFrequencies(v []string) {
	o.Frequencies = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *DiskBackupCopySetting) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupCopySetting) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *DiskBackupCopySetting) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *DiskBackupCopySetting) SetRegionName(v string) {
	o.RegionName = &v
}

// GetReplicationSpecId returns the ReplicationSpecId field value if set, zero value otherwise
func (o *DiskBackupCopySetting) GetReplicationSpecId() string {
	if o == nil || IsNil(o.ReplicationSpecId) {
		var ret string
		return ret
	}
	return *o.ReplicationSpecId
}

// GetReplicationSpecIdOk returns a tuple with the ReplicationSpecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupCopySetting) GetReplicationSpecIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationSpecId) {
		return nil, false
	}

	return o.ReplicationSpecId, true
}

// HasReplicationSpecId returns a boolean if a field has been set.
func (o *DiskBackupCopySetting) HasReplicationSpecId() bool {
	if o != nil && !IsNil(o.ReplicationSpecId) {
		return true
	}

	return false
}

// SetReplicationSpecId gets a reference to the given string and assigns it to the ReplicationSpecId field.
func (o *DiskBackupCopySetting) SetReplicationSpecId(v string) {
	o.ReplicationSpecId = &v
}

// GetShouldCopyOplogs returns the ShouldCopyOplogs field value if set, zero value otherwise
func (o *DiskBackupCopySetting) GetShouldCopyOplogs() bool {
	if o == nil || IsNil(o.ShouldCopyOplogs) {
		var ret bool
		return ret
	}
	return *o.ShouldCopyOplogs
}

// GetShouldCopyOplogsOk returns a tuple with the ShouldCopyOplogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupCopySetting) GetShouldCopyOplogsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldCopyOplogs) {
		return nil, false
	}

	return o.ShouldCopyOplogs, true
}

// HasShouldCopyOplogs returns a boolean if a field has been set.
func (o *DiskBackupCopySetting) HasShouldCopyOplogs() bool {
	if o != nil && !IsNil(o.ShouldCopyOplogs) {
		return true
	}

	return false
}

// SetShouldCopyOplogs gets a reference to the given bool and assigns it to the ShouldCopyOplogs field.
func (o *DiskBackupCopySetting) SetShouldCopyOplogs(v bool) {
	o.ShouldCopyOplogs = &v
}

func (o DiskBackupCopySetting) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DiskBackupCopySetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.Frequencies) {
		toSerialize["frequencies"] = o.Frequencies
	}
	if !IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !IsNil(o.ReplicationSpecId) {
		toSerialize["replicationSpecId"] = o.ReplicationSpecId
	}
	if !IsNil(o.ShouldCopyOplogs) {
		toSerialize["shouldCopyOplogs"] = o.ShouldCopyOplogs
	}
	return toSerialize, nil
}
