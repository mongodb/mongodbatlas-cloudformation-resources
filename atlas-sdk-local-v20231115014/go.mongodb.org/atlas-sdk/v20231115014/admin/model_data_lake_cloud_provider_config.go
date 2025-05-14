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

// DataLakeCloudProviderConfig Cloud provider linked to this data lake.
type DataLakeCloudProviderConfig struct {
	Aws DataLakeAWSCloudProviderConfig `json:"aws"`
}

// NewDataLakeCloudProviderConfig instantiates a new DataLakeCloudProviderConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeCloudProviderConfig(aws DataLakeAWSCloudProviderConfig) *DataLakeCloudProviderConfig {
	this := DataLakeCloudProviderConfig{}
	this.Aws = aws
	return &this
}

// NewDataLakeCloudProviderConfigWithDefaults instantiates a new DataLakeCloudProviderConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeCloudProviderConfigWithDefaults() *DataLakeCloudProviderConfig {
	this := DataLakeCloudProviderConfig{}
	return &this
}

// GetAws returns the Aws field value
func (o *DataLakeCloudProviderConfig) GetAws() DataLakeAWSCloudProviderConfig {
	if o == nil {
		var ret DataLakeAWSCloudProviderConfig
		return ret
	}

	return o.Aws
}

// GetAwsOk returns a tuple with the Aws field value
// and a boolean to check if the value has been set.
func (o *DataLakeCloudProviderConfig) GetAwsOk() (*DataLakeAWSCloudProviderConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aws, true
}

// SetAws sets field value
func (o *DataLakeCloudProviderConfig) SetAws(v DataLakeAWSCloudProviderConfig) {
	o.Aws = v
}

func (o DataLakeCloudProviderConfig) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataLakeCloudProviderConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aws"] = o.Aws
	return toSerialize, nil
}
