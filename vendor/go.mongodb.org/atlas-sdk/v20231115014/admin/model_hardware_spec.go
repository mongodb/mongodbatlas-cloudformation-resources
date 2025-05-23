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

// HardwareSpec Hardware specifications for all electable nodes deployed in the region. Electable nodes can become the primary and can enable local reads. If you don't specify this option, MongoDB Cloud deploys no electable nodes to the region.
type HardwareSpec struct {
	// Target throughput desired for storage attached to your AWS-provisioned cluster. Change this parameter only if you:  - set `\"replicationSpecs[n].regionConfigs[m].providerName\" : \"AWS\"`. - set `\"replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize\" : \"M30\"` or greater not including `Mxx_NVME` tiers.  The maximum input/output operations per second (IOPS) depend on the selected **.instanceSize** and **.diskSizeGB**. This parameter defaults to the cluster tier's standard IOPS value. Changing this value impacts cluster cost. MongoDB Cloud enforces minimum ratios of storage capacity to system memory for given cluster tiers. This keeps cluster performance consistent with large datasets.  - Instance sizes `M10` to `M40` have a ratio of disk capacity to system memory of 60:1. - Instance sizes greater than `M40` have a ratio of 120:1.
	DiskIOPS *int `json:"diskIOPS,omitempty"`
	// Type of storage you want to attach to your AWS-provisioned cluster.  - `STANDARD` volume types can't exceed the default input/output operations per second (IOPS) rate for the selected volume size.   - `PROVISIONED` volume types must fall within the allowable IOPS range for the selected volume size. You must set this value to (`PROVISIONED`) for NVMe clusters.
	EbsVolumeType *string `json:"ebsVolumeType,omitempty"`
	// Hardware specification for the instance sizes in this region. Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size.
	InstanceSize *string `json:"instanceSize,omitempty"`
	// Number of nodes of the given type for MongoDB Cloud to deploy to the region.
	NodeCount *int `json:"nodeCount,omitempty"`
}

// NewHardwareSpec instantiates a new HardwareSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHardwareSpec() *HardwareSpec {
	this := HardwareSpec{}
	var ebsVolumeType string = "STANDARD"
	this.EbsVolumeType = &ebsVolumeType
	return &this
}

// NewHardwareSpecWithDefaults instantiates a new HardwareSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHardwareSpecWithDefaults() *HardwareSpec {
	this := HardwareSpec{}
	var ebsVolumeType string = "STANDARD"
	this.EbsVolumeType = &ebsVolumeType
	return &this
}

// GetDiskIOPS returns the DiskIOPS field value if set, zero value otherwise
func (o *HardwareSpec) GetDiskIOPS() int {
	if o == nil || IsNil(o.DiskIOPS) {
		var ret int
		return ret
	}
	return *o.DiskIOPS
}

// GetDiskIOPSOk returns a tuple with the DiskIOPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareSpec) GetDiskIOPSOk() (*int, bool) {
	if o == nil || IsNil(o.DiskIOPS) {
		return nil, false
	}

	return o.DiskIOPS, true
}

// HasDiskIOPS returns a boolean if a field has been set.
func (o *HardwareSpec) HasDiskIOPS() bool {
	if o != nil && !IsNil(o.DiskIOPS) {
		return true
	}

	return false
}

// SetDiskIOPS gets a reference to the given int and assigns it to the DiskIOPS field.
func (o *HardwareSpec) SetDiskIOPS(v int) {
	o.DiskIOPS = &v
}

// GetEbsVolumeType returns the EbsVolumeType field value if set, zero value otherwise
func (o *HardwareSpec) GetEbsVolumeType() string {
	if o == nil || IsNil(o.EbsVolumeType) {
		var ret string
		return ret
	}
	return *o.EbsVolumeType
}

// GetEbsVolumeTypeOk returns a tuple with the EbsVolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareSpec) GetEbsVolumeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EbsVolumeType) {
		return nil, false
	}

	return o.EbsVolumeType, true
}

// HasEbsVolumeType returns a boolean if a field has been set.
func (o *HardwareSpec) HasEbsVolumeType() bool {
	if o != nil && !IsNil(o.EbsVolumeType) {
		return true
	}

	return false
}

// SetEbsVolumeType gets a reference to the given string and assigns it to the EbsVolumeType field.
func (o *HardwareSpec) SetEbsVolumeType(v string) {
	o.EbsVolumeType = &v
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise
func (o *HardwareSpec) GetInstanceSize() string {
	if o == nil || IsNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareSpec) GetInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSize) {
		return nil, false
	}

	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *HardwareSpec) HasInstanceSize() bool {
	if o != nil && !IsNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *HardwareSpec) SetInstanceSize(v string) {
	o.InstanceSize = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise
func (o *HardwareSpec) GetNodeCount() int {
	if o == nil || IsNil(o.NodeCount) {
		var ret int
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareSpec) GetNodeCountOk() (*int, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}

	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *HardwareSpec) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int and assigns it to the NodeCount field.
func (o *HardwareSpec) SetNodeCount(v int) {
	o.NodeCount = &v
}

func (o HardwareSpec) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o HardwareSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiskIOPS) {
		toSerialize["diskIOPS"] = o.DiskIOPS
	}
	if !IsNil(o.EbsVolumeType) {
		toSerialize["ebsVolumeType"] = o.EbsVolumeType
	}
	if !IsNil(o.InstanceSize) {
		toSerialize["instanceSize"] = o.InstanceSize
	}
	if !IsNil(o.NodeCount) {
		toSerialize["nodeCount"] = o.NodeCount
	}
	return toSerialize, nil
}
