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

// CloudRegionConfig Cloud service provider on which MongoDB Cloud provisions the hosts.
type CloudRegionConfig struct {
	ElectableSpecs *HardwareSpec `json:"electableSpecs,omitempty"`
	// Precedence is given to this region when a primary election occurs. If your **regionConfigs** has only **readOnlySpecs**, **analyticsSpecs**, or both, set this value to `0`. If you have multiple **regionConfigs** objects (your cluster is multi-region or multi-cloud), they must have priorities in descending order. The highest priority is `7`.  **Example:** If you have three regions, their priorities would be `7`, `6`, and `5` respectively. If you added two more regions for supporting electable nodes, the priorities of those regions would be `4` and `3` respectively.
	Priority *int `json:"priority,omitempty"`
	// Cloud service provider on which MongoDB Cloud provisions the hosts. Set dedicated clusters to `AWS`, `GCP`, `AZURE` or `TENANT`.
	ProviderName *string `json:"providerName,omitempty"`
	// Physical location of your MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. The region name is only returned in the response for single-region clusters. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. It assigns the VPC a Classless Inter-Domain Routing (CIDR) block. To limit a new VPC peering connection to one Classless Inter-Domain Routing (CIDR) block and region, create the connection first. Deploy the cluster after the connection starts. GCP Clusters and Multi-region clusters require one VPC peering connection for each region. MongoDB nodes can use only the peering connection that resides in the same region as the nodes to communicate with the peered VPC.
	RegionName           *string                      `json:"regionName,omitempty"`
	AnalyticsAutoScaling *AdvancedAutoScalingSettings `json:"analyticsAutoScaling,omitempty"`
	AnalyticsSpecs       *DedicatedHardwareSpec       `json:"analyticsSpecs,omitempty"`
	AutoScaling          *AdvancedAutoScalingSettings `json:"autoScaling,omitempty"`
	ReadOnlySpecs        *DedicatedHardwareSpec       `json:"readOnlySpecs,omitempty"`
	// Cloud service provider on which MongoDB Cloud provisioned the multi-tenant cluster. The resource returns this parameter when **providerName** is `TENANT` and **electableSpecs.instanceSize** is `M0`, `M2` or `M5`.
	BackingProviderName *string `json:"backingProviderName,omitempty"`
}

// NewCloudRegionConfig instantiates a new CloudRegionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRegionConfig() *CloudRegionConfig {
	this := CloudRegionConfig{}
	return &this
}

// NewCloudRegionConfigWithDefaults instantiates a new CloudRegionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRegionConfigWithDefaults() *CloudRegionConfig {
	this := CloudRegionConfig{}
	return &this
}

// GetElectableSpecs returns the ElectableSpecs field value if set, zero value otherwise
func (o *CloudRegionConfig) GetElectableSpecs() HardwareSpec {
	if o == nil || IsNil(o.ElectableSpecs) {
		var ret HardwareSpec
		return ret
	}
	return *o.ElectableSpecs
}

// GetElectableSpecsOk returns a tuple with the ElectableSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionConfig) GetElectableSpecsOk() (*HardwareSpec, bool) {
	if o == nil || IsNil(o.ElectableSpecs) {
		return nil, false
	}

	return o.ElectableSpecs, true
}

// HasElectableSpecs returns a boolean if a field has been set.
func (o *CloudRegionConfig) HasElectableSpecs() bool {
	if o != nil && !IsNil(o.ElectableSpecs) {
		return true
	}

	return false
}

// SetElectableSpecs gets a reference to the given HardwareSpec and assigns it to the ElectableSpecs field.
func (o *CloudRegionConfig) SetElectableSpecs(v HardwareSpec) {
	o.ElectableSpecs = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise
func (o *CloudRegionConfig) GetPriority() int {
	if o == nil || IsNil(o.Priority) {
		var ret int
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionConfig) GetPriorityOk() (*int, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}

	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CloudRegionConfig) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int and assigns it to the Priority field.
func (o *CloudRegionConfig) SetPriority(v int) {
	o.Priority = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise
func (o *CloudRegionConfig) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionConfig) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}

	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *CloudRegionConfig) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *CloudRegionConfig) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *CloudRegionConfig) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionConfig) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *CloudRegionConfig) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *CloudRegionConfig) SetRegionName(v string) {
	o.RegionName = &v
}

// GetAnalyticsAutoScaling returns the AnalyticsAutoScaling field value if set, zero value otherwise
func (o *CloudRegionConfig) GetAnalyticsAutoScaling() AdvancedAutoScalingSettings {
	if o == nil || IsNil(o.AnalyticsAutoScaling) {
		var ret AdvancedAutoScalingSettings
		return ret
	}
	return *o.AnalyticsAutoScaling
}

// GetAnalyticsAutoScalingOk returns a tuple with the AnalyticsAutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionConfig) GetAnalyticsAutoScalingOk() (*AdvancedAutoScalingSettings, bool) {
	if o == nil || IsNil(o.AnalyticsAutoScaling) {
		return nil, false
	}

	return o.AnalyticsAutoScaling, true
}

// HasAnalyticsAutoScaling returns a boolean if a field has been set.
func (o *CloudRegionConfig) HasAnalyticsAutoScaling() bool {
	if o != nil && !IsNil(o.AnalyticsAutoScaling) {
		return true
	}

	return false
}

// SetAnalyticsAutoScaling gets a reference to the given AdvancedAutoScalingSettings and assigns it to the AnalyticsAutoScaling field.
func (o *CloudRegionConfig) SetAnalyticsAutoScaling(v AdvancedAutoScalingSettings) {
	o.AnalyticsAutoScaling = &v
}

// GetAnalyticsSpecs returns the AnalyticsSpecs field value if set, zero value otherwise
func (o *CloudRegionConfig) GetAnalyticsSpecs() DedicatedHardwareSpec {
	if o == nil || IsNil(o.AnalyticsSpecs) {
		var ret DedicatedHardwareSpec
		return ret
	}
	return *o.AnalyticsSpecs
}

// GetAnalyticsSpecsOk returns a tuple with the AnalyticsSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionConfig) GetAnalyticsSpecsOk() (*DedicatedHardwareSpec, bool) {
	if o == nil || IsNil(o.AnalyticsSpecs) {
		return nil, false
	}

	return o.AnalyticsSpecs, true
}

// HasAnalyticsSpecs returns a boolean if a field has been set.
func (o *CloudRegionConfig) HasAnalyticsSpecs() bool {
	if o != nil && !IsNil(o.AnalyticsSpecs) {
		return true
	}

	return false
}

// SetAnalyticsSpecs gets a reference to the given DedicatedHardwareSpec and assigns it to the AnalyticsSpecs field.
func (o *CloudRegionConfig) SetAnalyticsSpecs(v DedicatedHardwareSpec) {
	o.AnalyticsSpecs = &v
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise
func (o *CloudRegionConfig) GetAutoScaling() AdvancedAutoScalingSettings {
	if o == nil || IsNil(o.AutoScaling) {
		var ret AdvancedAutoScalingSettings
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionConfig) GetAutoScalingOk() (*AdvancedAutoScalingSettings, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}

	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *CloudRegionConfig) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given AdvancedAutoScalingSettings and assigns it to the AutoScaling field.
func (o *CloudRegionConfig) SetAutoScaling(v AdvancedAutoScalingSettings) {
	o.AutoScaling = &v
}

// GetReadOnlySpecs returns the ReadOnlySpecs field value if set, zero value otherwise
func (o *CloudRegionConfig) GetReadOnlySpecs() DedicatedHardwareSpec {
	if o == nil || IsNil(o.ReadOnlySpecs) {
		var ret DedicatedHardwareSpec
		return ret
	}
	return *o.ReadOnlySpecs
}

// GetReadOnlySpecsOk returns a tuple with the ReadOnlySpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionConfig) GetReadOnlySpecsOk() (*DedicatedHardwareSpec, bool) {
	if o == nil || IsNil(o.ReadOnlySpecs) {
		return nil, false
	}

	return o.ReadOnlySpecs, true
}

// HasReadOnlySpecs returns a boolean if a field has been set.
func (o *CloudRegionConfig) HasReadOnlySpecs() bool {
	if o != nil && !IsNil(o.ReadOnlySpecs) {
		return true
	}

	return false
}

// SetReadOnlySpecs gets a reference to the given DedicatedHardwareSpec and assigns it to the ReadOnlySpecs field.
func (o *CloudRegionConfig) SetReadOnlySpecs(v DedicatedHardwareSpec) {
	o.ReadOnlySpecs = &v
}

// GetBackingProviderName returns the BackingProviderName field value if set, zero value otherwise
func (o *CloudRegionConfig) GetBackingProviderName() string {
	if o == nil || IsNil(o.BackingProviderName) {
		var ret string
		return ret
	}
	return *o.BackingProviderName
}

// GetBackingProviderNameOk returns a tuple with the BackingProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRegionConfig) GetBackingProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.BackingProviderName) {
		return nil, false
	}

	return o.BackingProviderName, true
}

// HasBackingProviderName returns a boolean if a field has been set.
func (o *CloudRegionConfig) HasBackingProviderName() bool {
	if o != nil && !IsNil(o.BackingProviderName) {
		return true
	}

	return false
}

// SetBackingProviderName gets a reference to the given string and assigns it to the BackingProviderName field.
func (o *CloudRegionConfig) SetBackingProviderName(v string) {
	o.BackingProviderName = &v
}

func (o CloudRegionConfig) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudRegionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ElectableSpecs) {
		toSerialize["electableSpecs"] = o.ElectableSpecs
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	if !IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !IsNil(o.AnalyticsAutoScaling) {
		toSerialize["analyticsAutoScaling"] = o.AnalyticsAutoScaling
	}
	if !IsNil(o.AnalyticsSpecs) {
		toSerialize["analyticsSpecs"] = o.AnalyticsSpecs
	}
	if !IsNil(o.AutoScaling) {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if !IsNil(o.ReadOnlySpecs) {
		toSerialize["readOnlySpecs"] = o.ReadOnlySpecs
	}
	if !IsNil(o.BackingProviderName) {
		toSerialize["backingProviderName"] = o.BackingProviderName
	}
	return toSerialize, nil
}
