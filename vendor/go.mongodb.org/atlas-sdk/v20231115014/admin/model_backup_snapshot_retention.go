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

// BackupSnapshotRetention struct for BackupSnapshotRetention
type BackupSnapshotRetention struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Quantity of time in which MongoDB Cloud measures snapshot retention.
	RetentionUnit string `json:"retentionUnit"`
	// Number that indicates the amount of days, weeks, months, or years that MongoDB Cloud retains the snapshot. For less frequent policy items, MongoDB Cloud requires that you specify a value greater than or equal to the value specified for more frequent policy items. If the hourly policy item specifies a retention of two days, specify two days or greater for the retention of the weekly policy item.
	RetentionValue int `json:"retentionValue"`
}

// NewBackupSnapshotRetention instantiates a new BackupSnapshotRetention object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupSnapshotRetention(retentionUnit string, retentionValue int) *BackupSnapshotRetention {
	this := BackupSnapshotRetention{}
	this.RetentionUnit = retentionUnit
	this.RetentionValue = retentionValue
	return &this
}

// NewBackupSnapshotRetentionWithDefaults instantiates a new BackupSnapshotRetention object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupSnapshotRetentionWithDefaults() *BackupSnapshotRetention {
	this := BackupSnapshotRetention{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *BackupSnapshotRetention) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotRetention) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BackupSnapshotRetention) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *BackupSnapshotRetention) SetLinks(v []Link) {
	o.Links = &v
}

// GetRetentionUnit returns the RetentionUnit field value
func (o *BackupSnapshotRetention) GetRetentionUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RetentionUnit
}

// GetRetentionUnitOk returns a tuple with the RetentionUnit field value
// and a boolean to check if the value has been set.
func (o *BackupSnapshotRetention) GetRetentionUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionUnit, true
}

// SetRetentionUnit sets field value
func (o *BackupSnapshotRetention) SetRetentionUnit(v string) {
	o.RetentionUnit = v
}

// GetRetentionValue returns the RetentionValue field value
func (o *BackupSnapshotRetention) GetRetentionValue() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.RetentionValue
}

// GetRetentionValueOk returns a tuple with the RetentionValue field value
// and a boolean to check if the value has been set.
func (o *BackupSnapshotRetention) GetRetentionValueOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionValue, true
}

// SetRetentionValue sets field value
func (o *BackupSnapshotRetention) SetRetentionValue(v int) {
	o.RetentionValue = v
}

func (o BackupSnapshotRetention) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o BackupSnapshotRetention) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["retentionUnit"] = o.RetentionUnit
	toSerialize["retentionValue"] = o.RetentionValue
	return toSerialize, nil
}
