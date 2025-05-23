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

// ServerlessProviderSettings Group of cloud provider settings that configure the provisioned MongoDB serverless instance.
type ServerlessProviderSettings struct {
	// Cloud service provider on which MongoDB Cloud provisioned the serverless instance.
	BackingProviderName string `json:"backingProviderName"`
	// Human-readable label that identifies the cloud service provider.
	ProviderName *string `json:"providerName,omitempty"`
	// Human-readable label that identifies the geographic location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
	RegionName string `json:"regionName"`
}

// NewServerlessProviderSettings instantiates a new ServerlessProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessProviderSettings(backingProviderName string, regionName string) *ServerlessProviderSettings {
	this := ServerlessProviderSettings{}
	this.BackingProviderName = backingProviderName
	var providerName string = "SERVERLESS"
	this.ProviderName = &providerName
	this.RegionName = regionName
	return &this
}

// NewServerlessProviderSettingsWithDefaults instantiates a new ServerlessProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessProviderSettingsWithDefaults() *ServerlessProviderSettings {
	this := ServerlessProviderSettings{}
	var providerName string = "SERVERLESS"
	this.ProviderName = &providerName
	return &this
}

// GetBackingProviderName returns the BackingProviderName field value
func (o *ServerlessProviderSettings) GetBackingProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackingProviderName
}

// GetBackingProviderNameOk returns a tuple with the BackingProviderName field value
// and a boolean to check if the value has been set.
func (o *ServerlessProviderSettings) GetBackingProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackingProviderName, true
}

// SetBackingProviderName sets field value
func (o *ServerlessProviderSettings) SetBackingProviderName(v string) {
	o.BackingProviderName = v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise
func (o *ServerlessProviderSettings) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessProviderSettings) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}

	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *ServerlessProviderSettings) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *ServerlessProviderSettings) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetRegionName returns the RegionName field value
func (o *ServerlessProviderSettings) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *ServerlessProviderSettings) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *ServerlessProviderSettings) SetRegionName(v string) {
	o.RegionName = v
}

func (o ServerlessProviderSettings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServerlessProviderSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["backingProviderName"] = o.BackingProviderName
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	toSerialize["regionName"] = o.RegionName
	return toSerialize, nil
}
