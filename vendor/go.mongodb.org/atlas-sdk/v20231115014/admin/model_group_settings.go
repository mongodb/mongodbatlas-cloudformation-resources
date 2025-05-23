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

// GroupSettings Collection of settings that configures the project.
type GroupSettings struct {
	// Flag that indicates whether to collect database-specific metrics  for the specified project.
	IsCollectDatabaseSpecificsStatisticsEnabled *bool `json:"isCollectDatabaseSpecificsStatisticsEnabled,omitempty"`
	// Flag that indicates whether to enable the Data Explorer for the specified project.
	IsDataExplorerEnabled *bool `json:"isDataExplorerEnabled,omitempty"`
	// Flag that indicates whether to enable extended storage sizes  for the specified project.
	IsExtendedStorageSizesEnabled *bool `json:"isExtendedStorageSizesEnabled,omitempty"`
	// Flag that indicates whether to enable the Performance Advisor and Profiler  for the specified project.
	IsPerformanceAdvisorEnabled *bool `json:"isPerformanceAdvisorEnabled,omitempty"`
	// Flag that indicates whether to enable the Real Time Performance Panel for the specified project.
	IsRealtimePerformancePanelEnabled *bool `json:"isRealtimePerformancePanelEnabled,omitempty"`
	// Flag that indicates whether to enable the Schema Advisor for the specified project.
	IsSchemaAdvisorEnabled *bool `json:"isSchemaAdvisorEnabled,omitempty"`
}

// NewGroupSettings instantiates a new GroupSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSettings() *GroupSettings {
	this := GroupSettings{}
	return &this
}

// NewGroupSettingsWithDefaults instantiates a new GroupSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSettingsWithDefaults() *GroupSettings {
	this := GroupSettings{}
	return &this
}

// GetIsCollectDatabaseSpecificsStatisticsEnabled returns the IsCollectDatabaseSpecificsStatisticsEnabled field value if set, zero value otherwise
func (o *GroupSettings) GetIsCollectDatabaseSpecificsStatisticsEnabled() bool {
	if o == nil || IsNil(o.IsCollectDatabaseSpecificsStatisticsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsCollectDatabaseSpecificsStatisticsEnabled
}

// GetIsCollectDatabaseSpecificsStatisticsEnabledOk returns a tuple with the IsCollectDatabaseSpecificsStatisticsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettings) GetIsCollectDatabaseSpecificsStatisticsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCollectDatabaseSpecificsStatisticsEnabled) {
		return nil, false
	}

	return o.IsCollectDatabaseSpecificsStatisticsEnabled, true
}

// HasIsCollectDatabaseSpecificsStatisticsEnabled returns a boolean if a field has been set.
func (o *GroupSettings) HasIsCollectDatabaseSpecificsStatisticsEnabled() bool {
	if o != nil && !IsNil(o.IsCollectDatabaseSpecificsStatisticsEnabled) {
		return true
	}

	return false
}

// SetIsCollectDatabaseSpecificsStatisticsEnabled gets a reference to the given bool and assigns it to the IsCollectDatabaseSpecificsStatisticsEnabled field.
func (o *GroupSettings) SetIsCollectDatabaseSpecificsStatisticsEnabled(v bool) {
	o.IsCollectDatabaseSpecificsStatisticsEnabled = &v
}

// GetIsDataExplorerEnabled returns the IsDataExplorerEnabled field value if set, zero value otherwise
func (o *GroupSettings) GetIsDataExplorerEnabled() bool {
	if o == nil || IsNil(o.IsDataExplorerEnabled) {
		var ret bool
		return ret
	}
	return *o.IsDataExplorerEnabled
}

// GetIsDataExplorerEnabledOk returns a tuple with the IsDataExplorerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettings) GetIsDataExplorerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDataExplorerEnabled) {
		return nil, false
	}

	return o.IsDataExplorerEnabled, true
}

// HasIsDataExplorerEnabled returns a boolean if a field has been set.
func (o *GroupSettings) HasIsDataExplorerEnabled() bool {
	if o != nil && !IsNil(o.IsDataExplorerEnabled) {
		return true
	}

	return false
}

// SetIsDataExplorerEnabled gets a reference to the given bool and assigns it to the IsDataExplorerEnabled field.
func (o *GroupSettings) SetIsDataExplorerEnabled(v bool) {
	o.IsDataExplorerEnabled = &v
}

// GetIsExtendedStorageSizesEnabled returns the IsExtendedStorageSizesEnabled field value if set, zero value otherwise
func (o *GroupSettings) GetIsExtendedStorageSizesEnabled() bool {
	if o == nil || IsNil(o.IsExtendedStorageSizesEnabled) {
		var ret bool
		return ret
	}
	return *o.IsExtendedStorageSizesEnabled
}

// GetIsExtendedStorageSizesEnabledOk returns a tuple with the IsExtendedStorageSizesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettings) GetIsExtendedStorageSizesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExtendedStorageSizesEnabled) {
		return nil, false
	}

	return o.IsExtendedStorageSizesEnabled, true
}

// HasIsExtendedStorageSizesEnabled returns a boolean if a field has been set.
func (o *GroupSettings) HasIsExtendedStorageSizesEnabled() bool {
	if o != nil && !IsNil(o.IsExtendedStorageSizesEnabled) {
		return true
	}

	return false
}

// SetIsExtendedStorageSizesEnabled gets a reference to the given bool and assigns it to the IsExtendedStorageSizesEnabled field.
func (o *GroupSettings) SetIsExtendedStorageSizesEnabled(v bool) {
	o.IsExtendedStorageSizesEnabled = &v
}

// GetIsPerformanceAdvisorEnabled returns the IsPerformanceAdvisorEnabled field value if set, zero value otherwise
func (o *GroupSettings) GetIsPerformanceAdvisorEnabled() bool {
	if o == nil || IsNil(o.IsPerformanceAdvisorEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPerformanceAdvisorEnabled
}

// GetIsPerformanceAdvisorEnabledOk returns a tuple with the IsPerformanceAdvisorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettings) GetIsPerformanceAdvisorEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPerformanceAdvisorEnabled) {
		return nil, false
	}

	return o.IsPerformanceAdvisorEnabled, true
}

// HasIsPerformanceAdvisorEnabled returns a boolean if a field has been set.
func (o *GroupSettings) HasIsPerformanceAdvisorEnabled() bool {
	if o != nil && !IsNil(o.IsPerformanceAdvisorEnabled) {
		return true
	}

	return false
}

// SetIsPerformanceAdvisorEnabled gets a reference to the given bool and assigns it to the IsPerformanceAdvisorEnabled field.
func (o *GroupSettings) SetIsPerformanceAdvisorEnabled(v bool) {
	o.IsPerformanceAdvisorEnabled = &v
}

// GetIsRealtimePerformancePanelEnabled returns the IsRealtimePerformancePanelEnabled field value if set, zero value otherwise
func (o *GroupSettings) GetIsRealtimePerformancePanelEnabled() bool {
	if o == nil || IsNil(o.IsRealtimePerformancePanelEnabled) {
		var ret bool
		return ret
	}
	return *o.IsRealtimePerformancePanelEnabled
}

// GetIsRealtimePerformancePanelEnabledOk returns a tuple with the IsRealtimePerformancePanelEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettings) GetIsRealtimePerformancePanelEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRealtimePerformancePanelEnabled) {
		return nil, false
	}

	return o.IsRealtimePerformancePanelEnabled, true
}

// HasIsRealtimePerformancePanelEnabled returns a boolean if a field has been set.
func (o *GroupSettings) HasIsRealtimePerformancePanelEnabled() bool {
	if o != nil && !IsNil(o.IsRealtimePerformancePanelEnabled) {
		return true
	}

	return false
}

// SetIsRealtimePerformancePanelEnabled gets a reference to the given bool and assigns it to the IsRealtimePerformancePanelEnabled field.
func (o *GroupSettings) SetIsRealtimePerformancePanelEnabled(v bool) {
	o.IsRealtimePerformancePanelEnabled = &v
}

// GetIsSchemaAdvisorEnabled returns the IsSchemaAdvisorEnabled field value if set, zero value otherwise
func (o *GroupSettings) GetIsSchemaAdvisorEnabled() bool {
	if o == nil || IsNil(o.IsSchemaAdvisorEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSchemaAdvisorEnabled
}

// GetIsSchemaAdvisorEnabledOk returns a tuple with the IsSchemaAdvisorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSettings) GetIsSchemaAdvisorEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSchemaAdvisorEnabled) {
		return nil, false
	}

	return o.IsSchemaAdvisorEnabled, true
}

// HasIsSchemaAdvisorEnabled returns a boolean if a field has been set.
func (o *GroupSettings) HasIsSchemaAdvisorEnabled() bool {
	if o != nil && !IsNil(o.IsSchemaAdvisorEnabled) {
		return true
	}

	return false
}

// SetIsSchemaAdvisorEnabled gets a reference to the given bool and assigns it to the IsSchemaAdvisorEnabled field.
func (o *GroupSettings) SetIsSchemaAdvisorEnabled(v bool) {
	o.IsSchemaAdvisorEnabled = &v
}

func (o GroupSettings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsCollectDatabaseSpecificsStatisticsEnabled) {
		toSerialize["isCollectDatabaseSpecificsStatisticsEnabled"] = o.IsCollectDatabaseSpecificsStatisticsEnabled
	}
	if !IsNil(o.IsDataExplorerEnabled) {
		toSerialize["isDataExplorerEnabled"] = o.IsDataExplorerEnabled
	}
	if !IsNil(o.IsExtendedStorageSizesEnabled) {
		toSerialize["isExtendedStorageSizesEnabled"] = o.IsExtendedStorageSizesEnabled
	}
	if !IsNil(o.IsPerformanceAdvisorEnabled) {
		toSerialize["isPerformanceAdvisorEnabled"] = o.IsPerformanceAdvisorEnabled
	}
	if !IsNil(o.IsRealtimePerformancePanelEnabled) {
		toSerialize["isRealtimePerformancePanelEnabled"] = o.IsRealtimePerformancePanelEnabled
	}
	if !IsNil(o.IsSchemaAdvisorEnabled) {
		toSerialize["isSchemaAdvisorEnabled"] = o.IsSchemaAdvisorEnabled
	}
	return toSerialize, nil
}
