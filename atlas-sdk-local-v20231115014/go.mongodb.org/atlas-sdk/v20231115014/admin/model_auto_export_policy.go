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

// AutoExportPolicy Policy for automatically exporting cloud backup snapshots.
type AutoExportPolicy struct {
	// Unique 24-hexadecimal character string that identifies the AWS Bucket.
	ExportBucketId *string `json:"exportBucketId,omitempty"`
	// Human-readable label that indicates the rate at which the export policy item occurs.
	FrequencyType *string `json:"frequencyType,omitempty"`
}

// NewAutoExportPolicy instantiates a new AutoExportPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoExportPolicy() *AutoExportPolicy {
	this := AutoExportPolicy{}
	return &this
}

// NewAutoExportPolicyWithDefaults instantiates a new AutoExportPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoExportPolicyWithDefaults() *AutoExportPolicy {
	this := AutoExportPolicy{}
	return &this
}

// GetExportBucketId returns the ExportBucketId field value if set, zero value otherwise
func (o *AutoExportPolicy) GetExportBucketId() string {
	if o == nil || IsNil(o.ExportBucketId) {
		var ret string
		return ret
	}
	return *o.ExportBucketId
}

// GetExportBucketIdOk returns a tuple with the ExportBucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoExportPolicy) GetExportBucketIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExportBucketId) {
		return nil, false
	}

	return o.ExportBucketId, true
}

// HasExportBucketId returns a boolean if a field has been set.
func (o *AutoExportPolicy) HasExportBucketId() bool {
	if o != nil && !IsNil(o.ExportBucketId) {
		return true
	}

	return false
}

// SetExportBucketId gets a reference to the given string and assigns it to the ExportBucketId field.
func (o *AutoExportPolicy) SetExportBucketId(v string) {
	o.ExportBucketId = &v
}

// GetFrequencyType returns the FrequencyType field value if set, zero value otherwise
func (o *AutoExportPolicy) GetFrequencyType() string {
	if o == nil || IsNil(o.FrequencyType) {
		var ret string
		return ret
	}
	return *o.FrequencyType
}

// GetFrequencyTypeOk returns a tuple with the FrequencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoExportPolicy) GetFrequencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FrequencyType) {
		return nil, false
	}

	return o.FrequencyType, true
}

// HasFrequencyType returns a boolean if a field has been set.
func (o *AutoExportPolicy) HasFrequencyType() bool {
	if o != nil && !IsNil(o.FrequencyType) {
		return true
	}

	return false
}

// SetFrequencyType gets a reference to the given string and assigns it to the FrequencyType field.
func (o *AutoExportPolicy) SetFrequencyType(v string) {
	o.FrequencyType = &v
}

func (o AutoExportPolicy) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AutoExportPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExportBucketId) {
		toSerialize["exportBucketId"] = o.ExportBucketId
	}
	if !IsNil(o.FrequencyType) {
		toSerialize["frequencyType"] = o.FrequencyType
	}
	return toSerialize, nil
}
