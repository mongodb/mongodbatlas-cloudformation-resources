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

// BackupTenantSnapshot struct for BackupTenantSnapshot
type BackupTenantSnapshot struct {
	// Date and time when the download link no longer works. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	Expiration *time.Time `json:"expiration,omitempty"`
	// Date and time when MongoDB Cloud completed writing this snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	FinishTime *time.Time `json:"finishTime,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the restore job.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// MongoDB host version that the snapshot runs.
	// Read only field.
	MongoDBVersion *string `json:"mongoDBVersion,omitempty"`
	// Date and time when MongoDB Cloud will take the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	ScheduledTime *time.Time `json:"scheduledTime,omitempty"`
	// Date and time when MongoDB Cloud began taking the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	StartTime *time.Time `json:"startTime,omitempty"`
	// Phase of the workflow for this snapshot at the time this resource made this request.
	// Read only field.
	Status *string `json:"status,omitempty"`
}

// NewBackupTenantSnapshot instantiates a new BackupTenantSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupTenantSnapshot() *BackupTenantSnapshot {
	this := BackupTenantSnapshot{}
	return &this
}

// NewBackupTenantSnapshotWithDefaults instantiates a new BackupTenantSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupTenantSnapshotWithDefaults() *BackupTenantSnapshot {
	this := BackupTenantSnapshot{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise
func (o *BackupTenantSnapshot) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupTenantSnapshot) GetExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}

	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *BackupTenantSnapshot) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *BackupTenantSnapshot) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetFinishTime returns the FinishTime field value if set, zero value otherwise
func (o *BackupTenantSnapshot) GetFinishTime() time.Time {
	if o == nil || IsNil(o.FinishTime) {
		var ret time.Time
		return ret
	}
	return *o.FinishTime
}

// GetFinishTimeOk returns a tuple with the FinishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupTenantSnapshot) GetFinishTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FinishTime) {
		return nil, false
	}

	return o.FinishTime, true
}

// HasFinishTime returns a boolean if a field has been set.
func (o *BackupTenantSnapshot) HasFinishTime() bool {
	if o != nil && !IsNil(o.FinishTime) {
		return true
	}

	return false
}

// SetFinishTime gets a reference to the given time.Time and assigns it to the FinishTime field.
func (o *BackupTenantSnapshot) SetFinishTime(v time.Time) {
	o.FinishTime = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *BackupTenantSnapshot) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupTenantSnapshot) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BackupTenantSnapshot) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BackupTenantSnapshot) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *BackupTenantSnapshot) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupTenantSnapshot) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BackupTenantSnapshot) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *BackupTenantSnapshot) SetLinks(v []Link) {
	o.Links = &v
}

// GetMongoDBVersion returns the MongoDBVersion field value if set, zero value otherwise
func (o *BackupTenantSnapshot) GetMongoDBVersion() string {
	if o == nil || IsNil(o.MongoDBVersion) {
		var ret string
		return ret
	}
	return *o.MongoDBVersion
}

// GetMongoDBVersionOk returns a tuple with the MongoDBVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupTenantSnapshot) GetMongoDBVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongoDBVersion) {
		return nil, false
	}

	return o.MongoDBVersion, true
}

// HasMongoDBVersion returns a boolean if a field has been set.
func (o *BackupTenantSnapshot) HasMongoDBVersion() bool {
	if o != nil && !IsNil(o.MongoDBVersion) {
		return true
	}

	return false
}

// SetMongoDBVersion gets a reference to the given string and assigns it to the MongoDBVersion field.
func (o *BackupTenantSnapshot) SetMongoDBVersion(v string) {
	o.MongoDBVersion = &v
}

// GetScheduledTime returns the ScheduledTime field value if set, zero value otherwise
func (o *BackupTenantSnapshot) GetScheduledTime() time.Time {
	if o == nil || IsNil(o.ScheduledTime) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledTime
}

// GetScheduledTimeOk returns a tuple with the ScheduledTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupTenantSnapshot) GetScheduledTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledTime) {
		return nil, false
	}

	return o.ScheduledTime, true
}

// HasScheduledTime returns a boolean if a field has been set.
func (o *BackupTenantSnapshot) HasScheduledTime() bool {
	if o != nil && !IsNil(o.ScheduledTime) {
		return true
	}

	return false
}

// SetScheduledTime gets a reference to the given time.Time and assigns it to the ScheduledTime field.
func (o *BackupTenantSnapshot) SetScheduledTime(v time.Time) {
	o.ScheduledTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise
func (o *BackupTenantSnapshot) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupTenantSnapshot) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}

	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *BackupTenantSnapshot) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *BackupTenantSnapshot) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *BackupTenantSnapshot) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupTenantSnapshot) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BackupTenantSnapshot) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BackupTenantSnapshot) SetStatus(v string) {
	o.Status = &v
}

func (o BackupTenantSnapshot) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o BackupTenantSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
