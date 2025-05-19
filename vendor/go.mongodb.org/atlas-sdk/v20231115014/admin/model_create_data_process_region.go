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

// CreateDataProcessRegion Settings to configure the region where you wish to store your archived data.
type CreateDataProcessRegion struct {
	// Human-readable label that identifies the Cloud service provider where you wish to store your archived data. **AZURE** may be selected only if **AZURE** is the Cloud service provider for the cluster and no **AWS** online archive has been created for the cluster.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// Human-readable label that identifies the geographic location of the region where you wish to store your archived data.
	Region *string `json:"region,omitempty"`
}

// NewCreateDataProcessRegion instantiates a new CreateDataProcessRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDataProcessRegion() *CreateDataProcessRegion {
	this := CreateDataProcessRegion{}
	return &this
}

// NewCreateDataProcessRegionWithDefaults instantiates a new CreateDataProcessRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDataProcessRegionWithDefaults() *CreateDataProcessRegion {
	this := CreateDataProcessRegion{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *CreateDataProcessRegion) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataProcessRegion) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *CreateDataProcessRegion) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *CreateDataProcessRegion) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetRegion returns the Region field value if set, zero value otherwise
func (o *CreateDataProcessRegion) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataProcessRegion) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}

	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CreateDataProcessRegion) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CreateDataProcessRegion) SetRegion(v string) {
	o.Region = &v
}

func (o CreateDataProcessRegion) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateDataProcessRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}
