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

// DataLakeDataProcessRegion Information about the cloud provider region to which the data lake routes client connections.
type DataLakeDataProcessRegion struct {
	// Name of the cloud service that hosts the data lake's data stores.
	CloudProvider string `json:"cloudProvider"`
	// Name of the region to which the data lake routes client connections.
	Region string `json:"region"`
}

// NewDataLakeDataProcessRegion instantiates a new DataLakeDataProcessRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeDataProcessRegion(cloudProvider string, region string) *DataLakeDataProcessRegion {
	this := DataLakeDataProcessRegion{}
	this.CloudProvider = cloudProvider
	this.Region = region
	return &this
}

// NewDataLakeDataProcessRegionWithDefaults instantiates a new DataLakeDataProcessRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeDataProcessRegionWithDefaults() *DataLakeDataProcessRegion {
	this := DataLakeDataProcessRegion{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *DataLakeDataProcessRegion) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *DataLakeDataProcessRegion) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *DataLakeDataProcessRegion) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetRegion returns the Region field value
func (o *DataLakeDataProcessRegion) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *DataLakeDataProcessRegion) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *DataLakeDataProcessRegion) SetRegion(v string) {
	o.Region = v
}

func (o DataLakeDataProcessRegion) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataLakeDataProcessRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProvider"] = o.CloudProvider
	toSerialize["region"] = o.Region
	return toSerialize, nil
}
