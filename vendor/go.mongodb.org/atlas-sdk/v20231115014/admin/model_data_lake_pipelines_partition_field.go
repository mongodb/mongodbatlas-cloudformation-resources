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

// DataLakePipelinesPartitionField Partition Field in the Data Lake Storage provider for a Data Lake Pipeline.
type DataLakePipelinesPartitionField struct {
	// Human-readable label that identifies the field name used to partition data.
	FieldName string `json:"fieldName"`
	// Sequence in which MongoDB Cloud slices the collection data to create partitions. The resource expresses this sequence starting with zero.
	Order int `json:"order"`
}

// NewDataLakePipelinesPartitionField instantiates a new DataLakePipelinesPartitionField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakePipelinesPartitionField(fieldName string, order int) *DataLakePipelinesPartitionField {
	this := DataLakePipelinesPartitionField{}
	this.FieldName = fieldName
	this.Order = order
	return &this
}

// NewDataLakePipelinesPartitionFieldWithDefaults instantiates a new DataLakePipelinesPartitionField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakePipelinesPartitionFieldWithDefaults() *DataLakePipelinesPartitionField {
	this := DataLakePipelinesPartitionField{}
	var order int = 0
	this.Order = order
	return &this
}

// GetFieldName returns the FieldName field value
func (o *DataLakePipelinesPartitionField) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *DataLakePipelinesPartitionField) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *DataLakePipelinesPartitionField) SetFieldName(v string) {
	o.FieldName = v
}

// GetOrder returns the Order field value
func (o *DataLakePipelinesPartitionField) GetOrder() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *DataLakePipelinesPartitionField) GetOrderOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *DataLakePipelinesPartitionField) SetOrder(v int) {
	o.Order = v
}

func (o DataLakePipelinesPartitionField) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataLakePipelinesPartitionField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldName"] = o.FieldName
	toSerialize["order"] = o.Order
	return toSerialize, nil
}
