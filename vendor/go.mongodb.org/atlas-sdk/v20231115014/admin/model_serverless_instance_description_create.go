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

// ServerlessInstanceDescriptionCreate Settings that you can specify when you create a serverless instance.
type ServerlessInstanceDescriptionCreate struct {
	// Human-readable label that identifies the serverless instance.
	// Write only field.
	Name                    string                          `json:"name"`
	ProviderSettings        ServerlessProviderSettings      `json:"providerSettings"`
	ServerlessBackupOptions *ClusterServerlessBackupOptions `json:"serverlessBackupOptions,omitempty"`
	// Human-readable label that indicates the current operating condition of the serverless instance.
	// Read only field.
	StateName *string `json:"stateName,omitempty"`
	// List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the serverless instance.
	Tags *[]ResourceTag `json:"tags,omitempty"`
	// Flag that indicates whether termination protection is enabled on the serverless instance. If set to `true`, MongoDB Cloud won't delete the serverless instance. If set to `false`, MongoDB Cloud will delete the serverless instance.
	TerminationProtectionEnabled *bool `json:"terminationProtectionEnabled,omitempty"`
}

// NewServerlessInstanceDescriptionCreate instantiates a new ServerlessInstanceDescriptionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessInstanceDescriptionCreate(name string, providerSettings ServerlessProviderSettings) *ServerlessInstanceDescriptionCreate {
	this := ServerlessInstanceDescriptionCreate{}
	this.Name = name
	this.ProviderSettings = providerSettings
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	return &this
}

// NewServerlessInstanceDescriptionCreateWithDefaults instantiates a new ServerlessInstanceDescriptionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessInstanceDescriptionCreateWithDefaults() *ServerlessInstanceDescriptionCreate {
	this := ServerlessInstanceDescriptionCreate{}
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	return &this
}

// GetName returns the Name field value
func (o *ServerlessInstanceDescriptionCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServerlessInstanceDescriptionCreate) SetName(v string) {
	o.Name = v
}

// GetProviderSettings returns the ProviderSettings field value
func (o *ServerlessInstanceDescriptionCreate) GetProviderSettings() ServerlessProviderSettings {
	if o == nil {
		var ret ServerlessProviderSettings
		return ret
	}

	return o.ProviderSettings
}

// GetProviderSettingsOk returns a tuple with the ProviderSettings field value
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionCreate) GetProviderSettingsOk() (*ServerlessProviderSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderSettings, true
}

// SetProviderSettings sets field value
func (o *ServerlessInstanceDescriptionCreate) SetProviderSettings(v ServerlessProviderSettings) {
	o.ProviderSettings = v
}

// GetServerlessBackupOptions returns the ServerlessBackupOptions field value if set, zero value otherwise
func (o *ServerlessInstanceDescriptionCreate) GetServerlessBackupOptions() ClusterServerlessBackupOptions {
	if o == nil || IsNil(o.ServerlessBackupOptions) {
		var ret ClusterServerlessBackupOptions
		return ret
	}
	return *o.ServerlessBackupOptions
}

// GetServerlessBackupOptionsOk returns a tuple with the ServerlessBackupOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionCreate) GetServerlessBackupOptionsOk() (*ClusterServerlessBackupOptions, bool) {
	if o == nil || IsNil(o.ServerlessBackupOptions) {
		return nil, false
	}

	return o.ServerlessBackupOptions, true
}

// HasServerlessBackupOptions returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionCreate) HasServerlessBackupOptions() bool {
	if o != nil && !IsNil(o.ServerlessBackupOptions) {
		return true
	}

	return false
}

// SetServerlessBackupOptions gets a reference to the given ClusterServerlessBackupOptions and assigns it to the ServerlessBackupOptions field.
func (o *ServerlessInstanceDescriptionCreate) SetServerlessBackupOptions(v ClusterServerlessBackupOptions) {
	o.ServerlessBackupOptions = &v
}

// GetStateName returns the StateName field value if set, zero value otherwise
func (o *ServerlessInstanceDescriptionCreate) GetStateName() string {
	if o == nil || IsNil(o.StateName) {
		var ret string
		return ret
	}
	return *o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionCreate) GetStateNameOk() (*string, bool) {
	if o == nil || IsNil(o.StateName) {
		return nil, false
	}

	return o.StateName, true
}

// HasStateName returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionCreate) HasStateName() bool {
	if o != nil && !IsNil(o.StateName) {
		return true
	}

	return false
}

// SetStateName gets a reference to the given string and assigns it to the StateName field.
func (o *ServerlessInstanceDescriptionCreate) SetStateName(v string) {
	o.StateName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise
func (o *ServerlessInstanceDescriptionCreate) GetTags() []ResourceTag {
	if o == nil || IsNil(o.Tags) {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionCreate) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}

	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionCreate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *ServerlessInstanceDescriptionCreate) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field value if set, zero value otherwise
func (o *ServerlessInstanceDescriptionCreate) GetTerminationProtectionEnabled() bool {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.TerminationProtectionEnabled
}

// GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionCreate) GetTerminationProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		return nil, false
	}

	return o.TerminationProtectionEnabled, true
}

// HasTerminationProtectionEnabled returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionCreate) HasTerminationProtectionEnabled() bool {
	if o != nil && !IsNil(o.TerminationProtectionEnabled) {
		return true
	}

	return false
}

// SetTerminationProtectionEnabled gets a reference to the given bool and assigns it to the TerminationProtectionEnabled field.
func (o *ServerlessInstanceDescriptionCreate) SetTerminationProtectionEnabled(v bool) {
	o.TerminationProtectionEnabled = &v
}

func (o ServerlessInstanceDescriptionCreate) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServerlessInstanceDescriptionCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["providerSettings"] = o.ProviderSettings
	if !IsNil(o.ServerlessBackupOptions) {
		toSerialize["serverlessBackupOptions"] = o.ServerlessBackupOptions
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TerminationProtectionEnabled) {
		toSerialize["terminationProtectionEnabled"] = o.TerminationProtectionEnabled
	}
	return toSerialize, nil
}
