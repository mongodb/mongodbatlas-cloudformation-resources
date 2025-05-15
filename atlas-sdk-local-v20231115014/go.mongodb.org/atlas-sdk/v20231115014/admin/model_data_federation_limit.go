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
	"time"
)

// DataFederationLimit Details of user managed limits.
type DataFederationLimit struct {
	// Amount that indicates the current usage of the limit.
	// Read only field.
	CurrentUsage *int64 `json:"currentUsage,omitempty"`
	// Default value of the limit.
	// Read only field.
	DefaultLimit *int64 `json:"defaultLimit,omitempty"`
	// Maximum value of the limit.
	// Read only field.
	MaximumLimit *int64 `json:"maximumLimit,omitempty"`
	// Human-readable label that identifies the user-managed limit to modify.
	// Read only field.
	Name string `json:"name"`
	// Amount to set the limit to.
	Value int64 `json:"value"`
	// Only used for Data Federation limits. Timestamp that indicates when this usage limit was last modified. This field uses the ISO 8601 timestamp format in UTC.
	// Read only field.
	LastModifiedDate *time.Time `json:"lastModifiedDate,omitempty"`
	// Only used for Data Federation limits. Action to take when the usage limit is exceeded. If limit span is set to QUERY, this is ignored because MongoDB Cloud stops the query when it exceeds the usage limit.
	OverrunPolicy *string `json:"overrunPolicy,omitempty"`
}

// NewDataFederationLimit instantiates a new DataFederationLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataFederationLimit(name string, value int64) *DataFederationLimit {
	this := DataFederationLimit{}
	this.Name = name
	this.Value = value
	return &this
}

// NewDataFederationLimitWithDefaults instantiates a new DataFederationLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataFederationLimitWithDefaults() *DataFederationLimit {
	this := DataFederationLimit{}
	return &this
}

// GetCurrentUsage returns the CurrentUsage field value if set, zero value otherwise
func (o *DataFederationLimit) GetCurrentUsage() int64 {
	if o == nil || IsNil(o.CurrentUsage) {
		var ret int64
		return ret
	}
	return *o.CurrentUsage
}

// GetCurrentUsageOk returns a tuple with the CurrentUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationLimit) GetCurrentUsageOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentUsage) {
		return nil, false
	}

	return o.CurrentUsage, true
}

// HasCurrentUsage returns a boolean if a field has been set.
func (o *DataFederationLimit) HasCurrentUsage() bool {
	if o != nil && !IsNil(o.CurrentUsage) {
		return true
	}

	return false
}

// SetCurrentUsage gets a reference to the given int64 and assigns it to the CurrentUsage field.
func (o *DataFederationLimit) SetCurrentUsage(v int64) {
	o.CurrentUsage = &v
}

// GetDefaultLimit returns the DefaultLimit field value if set, zero value otherwise
func (o *DataFederationLimit) GetDefaultLimit() int64 {
	if o == nil || IsNil(o.DefaultLimit) {
		var ret int64
		return ret
	}
	return *o.DefaultLimit
}

// GetDefaultLimitOk returns a tuple with the DefaultLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationLimit) GetDefaultLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultLimit) {
		return nil, false
	}

	return o.DefaultLimit, true
}

// HasDefaultLimit returns a boolean if a field has been set.
func (o *DataFederationLimit) HasDefaultLimit() bool {
	if o != nil && !IsNil(o.DefaultLimit) {
		return true
	}

	return false
}

// SetDefaultLimit gets a reference to the given int64 and assigns it to the DefaultLimit field.
func (o *DataFederationLimit) SetDefaultLimit(v int64) {
	o.DefaultLimit = &v
}

// GetMaximumLimit returns the MaximumLimit field value if set, zero value otherwise
func (o *DataFederationLimit) GetMaximumLimit() int64 {
	if o == nil || IsNil(o.MaximumLimit) {
		var ret int64
		return ret
	}
	return *o.MaximumLimit
}

// GetMaximumLimitOk returns a tuple with the MaximumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationLimit) GetMaximumLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumLimit) {
		return nil, false
	}

	return o.MaximumLimit, true
}

// HasMaximumLimit returns a boolean if a field has been set.
func (o *DataFederationLimit) HasMaximumLimit() bool {
	if o != nil && !IsNil(o.MaximumLimit) {
		return true
	}

	return false
}

// SetMaximumLimit gets a reference to the given int64 and assigns it to the MaximumLimit field.
func (o *DataFederationLimit) SetMaximumLimit(v int64) {
	o.MaximumLimit = &v
}

// GetName returns the Name field value
func (o *DataFederationLimit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataFederationLimit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataFederationLimit) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *DataFederationLimit) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DataFederationLimit) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *DataFederationLimit) SetValue(v int64) {
	o.Value = v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise
func (o *DataFederationLimit) GetLastModifiedDate() time.Time {
	if o == nil || IsNil(o.LastModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDate
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationLimit) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedDate) {
		return nil, false
	}

	return o.LastModifiedDate, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *DataFederationLimit) HasLastModifiedDate() bool {
	if o != nil && !IsNil(o.LastModifiedDate) {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given time.Time and assigns it to the LastModifiedDate field.
func (o *DataFederationLimit) SetLastModifiedDate(v time.Time) {
	o.LastModifiedDate = &v
}

// GetOverrunPolicy returns the OverrunPolicy field value if set, zero value otherwise
func (o *DataFederationLimit) GetOverrunPolicy() string {
	if o == nil || IsNil(o.OverrunPolicy) {
		var ret string
		return ret
	}
	return *o.OverrunPolicy
}

// GetOverrunPolicyOk returns a tuple with the OverrunPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationLimit) GetOverrunPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.OverrunPolicy) {
		return nil, false
	}

	return o.OverrunPolicy, true
}

// HasOverrunPolicy returns a boolean if a field has been set.
func (o *DataFederationLimit) HasOverrunPolicy() bool {
	if o != nil && !IsNil(o.OverrunPolicy) {
		return true
	}

	return false
}

// SetOverrunPolicy gets a reference to the given string and assigns it to the OverrunPolicy field.
func (o *DataFederationLimit) SetOverrunPolicy(v string) {
	o.OverrunPolicy = &v
}

func (o DataFederationLimit) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataFederationLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.OverrunPolicy) {
		toSerialize["overrunPolicy"] = o.OverrunPolicy
	}
	return toSerialize, nil
}
