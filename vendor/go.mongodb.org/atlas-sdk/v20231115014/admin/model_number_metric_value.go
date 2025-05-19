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

// NumberMetricValue Measurement of the **metricName** recorded at the time of the event.
type NumberMetricValue struct {
	// Amount of the **metricName** recorded at the time of the event. This value triggered the alert.
	// Read only field.
	Number *float64 `json:"number,omitempty"`
	// Element used to express the quantity. This can be an element of time, storage capacity, and the like.
	Units *string `json:"units,omitempty"`
}

// NewNumberMetricValue instantiates a new NumberMetricValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumberMetricValue() *NumberMetricValue {
	this := NumberMetricValue{}
	return &this
}

// NewNumberMetricValueWithDefaults instantiates a new NumberMetricValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumberMetricValueWithDefaults() *NumberMetricValue {
	this := NumberMetricValue{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise
func (o *NumberMetricValue) GetNumber() float64 {
	if o == nil || IsNil(o.Number) {
		var ret float64
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberMetricValue) GetNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}

	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *NumberMetricValue) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given float64 and assigns it to the Number field.
func (o *NumberMetricValue) SetNumber(v float64) {
	o.Number = &v
}

// GetUnits returns the Units field value if set, zero value otherwise
func (o *NumberMetricValue) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberMetricValue) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}

	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *NumberMetricValue) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *NumberMetricValue) SetUnits(v string) {
	o.Units = &v
}

func (o NumberMetricValue) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o NumberMetricValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	return toSerialize, nil
}
