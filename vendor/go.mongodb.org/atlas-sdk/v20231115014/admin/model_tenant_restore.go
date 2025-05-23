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

// TenantRestore struct for TenantRestore
type TenantRestore struct {
	// Human-readable label that identifies the source cluster.
	// Read only field.
	ClusterName *string `json:"clusterName,omitempty"`
	// Means by which this resource returns the snapshot to the requesting MongoDB Cloud user.
	// Read only field.
	DeliveryType *string `json:"deliveryType,omitempty"`
	// Date and time when the download link no longer works. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the restore job.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project from which the restore job originated.
	// Read only field.
	ProjectId *string `json:"projectId,omitempty"`
	// Date and time when MongoDB Cloud completed writing this snapshot. MongoDB Cloud changes the status of the restore job to `CLOSED`. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	RestoreFinishedDate *time.Time `json:"restoreFinishedDate,omitempty"`
	// Date and time when MongoDB Cloud will restore this snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	RestoreScheduledDate *time.Time `json:"restoreScheduledDate,omitempty"`
	// Date and time when MongoDB Cloud completed writing this snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	SnapshotFinishedDate *time.Time `json:"snapshotFinishedDate,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the snapshot to restore.
	SnapshotId string `json:"snapshotId"`
	// Internet address from which you can download the compressed snapshot files. The resource returns this parameter when  `\"deliveryType\" : \"DOWNLOAD\"`.
	// Read only field.
	SnapshotUrl *string `json:"snapshotUrl,omitempty"`
	// Phase of the restore workflow for this job at the time this resource made this request.
	// Read only field.
	Status *string `json:"status,omitempty"`
	// Human-readable label that identifies the cluster on the target project to which you want to restore the snapshot. You can restore the snapshot to a cluster tier *M2* or greater.
	TargetDeploymentItemName string `json:"targetDeploymentItemName"`
	// Unique 24-hexadecimal digit string that identifies the project that contains the cluster to which you want to restore the snapshot.
	TargetProjectId *string `json:"targetProjectId,omitempty"`
}

// NewTenantRestore instantiates a new TenantRestore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantRestore(snapshotId string, targetDeploymentItemName string) *TenantRestore {
	this := TenantRestore{}
	this.SnapshotId = snapshotId
	this.TargetDeploymentItemName = targetDeploymentItemName
	return &this
}

// NewTenantRestoreWithDefaults instantiates a new TenantRestore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantRestoreWithDefaults() *TenantRestore {
	this := TenantRestore{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *TenantRestore) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *TenantRestore) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *TenantRestore) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetDeliveryType returns the DeliveryType field value if set, zero value otherwise
func (o *TenantRestore) GetDeliveryType() string {
	if o == nil || IsNil(o.DeliveryType) {
		var ret string
		return ret
	}
	return *o.DeliveryType
}

// GetDeliveryTypeOk returns a tuple with the DeliveryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetDeliveryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryType) {
		return nil, false
	}

	return o.DeliveryType, true
}

// HasDeliveryType returns a boolean if a field has been set.
func (o *TenantRestore) HasDeliveryType() bool {
	if o != nil && !IsNil(o.DeliveryType) {
		return true
	}

	return false
}

// SetDeliveryType gets a reference to the given string and assigns it to the DeliveryType field.
func (o *TenantRestore) SetDeliveryType(v string) {
	o.DeliveryType = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise
func (o *TenantRestore) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}

	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *TenantRestore) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *TenantRestore) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *TenantRestore) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TenantRestore) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TenantRestore) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *TenantRestore) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TenantRestore) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *TenantRestore) SetLinks(v []Link) {
	o.Links = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise
func (o *TenantRestore) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}

	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *TenantRestore) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *TenantRestore) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRestoreFinishedDate returns the RestoreFinishedDate field value if set, zero value otherwise
func (o *TenantRestore) GetRestoreFinishedDate() time.Time {
	if o == nil || IsNil(o.RestoreFinishedDate) {
		var ret time.Time
		return ret
	}
	return *o.RestoreFinishedDate
}

// GetRestoreFinishedDateOk returns a tuple with the RestoreFinishedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetRestoreFinishedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RestoreFinishedDate) {
		return nil, false
	}

	return o.RestoreFinishedDate, true
}

// HasRestoreFinishedDate returns a boolean if a field has been set.
func (o *TenantRestore) HasRestoreFinishedDate() bool {
	if o != nil && !IsNil(o.RestoreFinishedDate) {
		return true
	}

	return false
}

// SetRestoreFinishedDate gets a reference to the given time.Time and assigns it to the RestoreFinishedDate field.
func (o *TenantRestore) SetRestoreFinishedDate(v time.Time) {
	o.RestoreFinishedDate = &v
}

// GetRestoreScheduledDate returns the RestoreScheduledDate field value if set, zero value otherwise
func (o *TenantRestore) GetRestoreScheduledDate() time.Time {
	if o == nil || IsNil(o.RestoreScheduledDate) {
		var ret time.Time
		return ret
	}
	return *o.RestoreScheduledDate
}

// GetRestoreScheduledDateOk returns a tuple with the RestoreScheduledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetRestoreScheduledDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RestoreScheduledDate) {
		return nil, false
	}

	return o.RestoreScheduledDate, true
}

// HasRestoreScheduledDate returns a boolean if a field has been set.
func (o *TenantRestore) HasRestoreScheduledDate() bool {
	if o != nil && !IsNil(o.RestoreScheduledDate) {
		return true
	}

	return false
}

// SetRestoreScheduledDate gets a reference to the given time.Time and assigns it to the RestoreScheduledDate field.
func (o *TenantRestore) SetRestoreScheduledDate(v time.Time) {
	o.RestoreScheduledDate = &v
}

// GetSnapshotFinishedDate returns the SnapshotFinishedDate field value if set, zero value otherwise
func (o *TenantRestore) GetSnapshotFinishedDate() time.Time {
	if o == nil || IsNil(o.SnapshotFinishedDate) {
		var ret time.Time
		return ret
	}
	return *o.SnapshotFinishedDate
}

// GetSnapshotFinishedDateOk returns a tuple with the SnapshotFinishedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetSnapshotFinishedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SnapshotFinishedDate) {
		return nil, false
	}

	return o.SnapshotFinishedDate, true
}

// HasSnapshotFinishedDate returns a boolean if a field has been set.
func (o *TenantRestore) HasSnapshotFinishedDate() bool {
	if o != nil && !IsNil(o.SnapshotFinishedDate) {
		return true
	}

	return false
}

// SetSnapshotFinishedDate gets a reference to the given time.Time and assigns it to the SnapshotFinishedDate field.
func (o *TenantRestore) SetSnapshotFinishedDate(v time.Time) {
	o.SnapshotFinishedDate = &v
}

// GetSnapshotId returns the SnapshotId field value
func (o *TenantRestore) GetSnapshotId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetSnapshotIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotId, true
}

// SetSnapshotId sets field value
func (o *TenantRestore) SetSnapshotId(v string) {
	o.SnapshotId = v
}

// GetSnapshotUrl returns the SnapshotUrl field value if set, zero value otherwise
func (o *TenantRestore) GetSnapshotUrl() string {
	if o == nil || IsNil(o.SnapshotUrl) {
		var ret string
		return ret
	}
	return *o.SnapshotUrl
}

// GetSnapshotUrlOk returns a tuple with the SnapshotUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetSnapshotUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotUrl) {
		return nil, false
	}

	return o.SnapshotUrl, true
}

// HasSnapshotUrl returns a boolean if a field has been set.
func (o *TenantRestore) HasSnapshotUrl() bool {
	if o != nil && !IsNil(o.SnapshotUrl) {
		return true
	}

	return false
}

// SetSnapshotUrl gets a reference to the given string and assigns it to the SnapshotUrl field.
func (o *TenantRestore) SetSnapshotUrl(v string) {
	o.SnapshotUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *TenantRestore) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TenantRestore) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TenantRestore) SetStatus(v string) {
	o.Status = &v
}

// GetTargetDeploymentItemName returns the TargetDeploymentItemName field value
func (o *TenantRestore) GetTargetDeploymentItemName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetDeploymentItemName
}

// GetTargetDeploymentItemNameOk returns a tuple with the TargetDeploymentItemName field value
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetTargetDeploymentItemNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetDeploymentItemName, true
}

// SetTargetDeploymentItemName sets field value
func (o *TenantRestore) SetTargetDeploymentItemName(v string) {
	o.TargetDeploymentItemName = v
}

// GetTargetProjectId returns the TargetProjectId field value if set, zero value otherwise
func (o *TenantRestore) GetTargetProjectId() string {
	if o == nil || IsNil(o.TargetProjectId) {
		var ret string
		return ret
	}
	return *o.TargetProjectId
}

// GetTargetProjectIdOk returns a tuple with the TargetProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRestore) GetTargetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetProjectId) {
		return nil, false
	}

	return o.TargetProjectId, true
}

// HasTargetProjectId returns a boolean if a field has been set.
func (o *TenantRestore) HasTargetProjectId() bool {
	if o != nil && !IsNil(o.TargetProjectId) {
		return true
	}

	return false
}

// SetTargetProjectId gets a reference to the given string and assigns it to the TargetProjectId field.
func (o *TenantRestore) SetTargetProjectId(v string) {
	o.TargetProjectId = &v
}

func (o TenantRestore) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TenantRestore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["snapshotId"] = o.SnapshotId
	toSerialize["targetDeploymentItemName"] = o.TargetDeploymentItemName
	if !IsNil(o.TargetProjectId) {
		toSerialize["targetProjectId"] = o.TargetProjectId
	}
	return toSerialize, nil
}
