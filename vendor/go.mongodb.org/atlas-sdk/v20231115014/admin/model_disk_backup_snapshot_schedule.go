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

// DiskBackupSnapshotSchedule struct for DiskBackupSnapshotSchedule
type DiskBackupSnapshotSchedule struct {
	// Flag that indicates whether MongoDB Cloud automatically exports cloud backup snapshots to the AWS bucket.
	AutoExportEnabled *bool `json:"autoExportEnabled,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the cluster with the snapshot you want to return.
	// Read only field.
	ClusterId *string `json:"clusterId,omitempty"`
	// Human-readable label that identifies the cluster with the snapshot you want to return.
	// Read only field.
	ClusterName *string `json:"clusterName,omitempty"`
	// List that contains a document for each copy setting item in the desired backup policy.
	CopySettings *[]DiskBackupCopySetting `json:"copySettings,omitempty"`
	// List that contains a document for each deleted copy setting whose backup copies you want to delete.
	// Write only field.
	DeleteCopiedBackups *[]DeleteCopiedBackups `json:"deleteCopiedBackups,omitempty"`
	Export              *AutoExportPolicy      `json:"export,omitempty"`
	// List that contains a document for each extra retention setting item in the desired backup policy.
	ExtraRetentionSettings *[]ExtraRetentionSetting `json:"extraRetentionSettings,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Date and time when MongoDB Cloud takes the next snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	NextSnapshot *time.Time `json:"nextSnapshot,omitempty"`
	// Rules set for this backup schedule.
	Policies *[]AdvancedDiskBackupSnapshotSchedulePolicy `json:"policies,omitempty"`
	// Hour of day in Coordinated Universal Time (UTC) that represents when MongoDB Cloud takes the snapshot.
	ReferenceHourOfDay *int `json:"referenceHourOfDay,omitempty"`
	// Minute of the **referenceHourOfDay** that represents when MongoDB Cloud takes the snapshot.
	ReferenceMinuteOfHour *int `json:"referenceMinuteOfHour,omitempty"`
	// Number of previous days that you can restore back to with Continuous Cloud Backup accuracy. You must specify a positive, non-zero integer. This parameter applies to continuous cloud backups only.
	RestoreWindowDays *int `json:"restoreWindowDays,omitempty"`
	// Flag that indicates whether to apply the retention changes in the updated backup policy to snapshots that MongoDB Cloud took previously.
	// Write only field.
	UpdateSnapshots *bool `json:"updateSnapshots,omitempty"`
	// Flag that indicates whether to use organization and project names instead of organization and project UUIDs in the path to the metadata files that MongoDB Cloud uploads to your AWS bucket.
	UseOrgAndGroupNamesInExportPrefix *bool `json:"useOrgAndGroupNamesInExportPrefix,omitempty"`
}

// NewDiskBackupSnapshotSchedule instantiates a new DiskBackupSnapshotSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupSnapshotSchedule() *DiskBackupSnapshotSchedule {
	this := DiskBackupSnapshotSchedule{}
	return &this
}

// NewDiskBackupSnapshotScheduleWithDefaults instantiates a new DiskBackupSnapshotSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupSnapshotScheduleWithDefaults() *DiskBackupSnapshotSchedule {
	this := DiskBackupSnapshotSchedule{}
	return &this
}

// GetAutoExportEnabled returns the AutoExportEnabled field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetAutoExportEnabled() bool {
	if o == nil || IsNil(o.AutoExportEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoExportEnabled
}

// GetAutoExportEnabledOk returns a tuple with the AutoExportEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetAutoExportEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoExportEnabled) {
		return nil, false
	}

	return o.AutoExportEnabled, true
}

// HasAutoExportEnabled returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasAutoExportEnabled() bool {
	if o != nil && !IsNil(o.AutoExportEnabled) {
		return true
	}

	return false
}

// SetAutoExportEnabled gets a reference to the given bool and assigns it to the AutoExportEnabled field.
func (o *DiskBackupSnapshotSchedule) SetAutoExportEnabled(v bool) {
	o.AutoExportEnabled = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}

	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *DiskBackupSnapshotSchedule) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *DiskBackupSnapshotSchedule) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCopySettings returns the CopySettings field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetCopySettings() []DiskBackupCopySetting {
	if o == nil || IsNil(o.CopySettings) {
		var ret []DiskBackupCopySetting
		return ret
	}
	return *o.CopySettings
}

// GetCopySettingsOk returns a tuple with the CopySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetCopySettingsOk() (*[]DiskBackupCopySetting, bool) {
	if o == nil || IsNil(o.CopySettings) {
		return nil, false
	}

	return o.CopySettings, true
}

// HasCopySettings returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasCopySettings() bool {
	if o != nil && !IsNil(o.CopySettings) {
		return true
	}

	return false
}

// SetCopySettings gets a reference to the given []DiskBackupCopySetting and assigns it to the CopySettings field.
func (o *DiskBackupSnapshotSchedule) SetCopySettings(v []DiskBackupCopySetting) {
	o.CopySettings = &v
}

// GetDeleteCopiedBackups returns the DeleteCopiedBackups field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetDeleteCopiedBackups() []DeleteCopiedBackups {
	if o == nil || IsNil(o.DeleteCopiedBackups) {
		var ret []DeleteCopiedBackups
		return ret
	}
	return *o.DeleteCopiedBackups
}

// GetDeleteCopiedBackupsOk returns a tuple with the DeleteCopiedBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetDeleteCopiedBackupsOk() (*[]DeleteCopiedBackups, bool) {
	if o == nil || IsNil(o.DeleteCopiedBackups) {
		return nil, false
	}

	return o.DeleteCopiedBackups, true
}

// HasDeleteCopiedBackups returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasDeleteCopiedBackups() bool {
	if o != nil && !IsNil(o.DeleteCopiedBackups) {
		return true
	}

	return false
}

// SetDeleteCopiedBackups gets a reference to the given []DeleteCopiedBackups and assigns it to the DeleteCopiedBackups field.
func (o *DiskBackupSnapshotSchedule) SetDeleteCopiedBackups(v []DeleteCopiedBackups) {
	o.DeleteCopiedBackups = &v
}

// GetExport returns the Export field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetExport() AutoExportPolicy {
	if o == nil || IsNil(o.Export) {
		var ret AutoExportPolicy
		return ret
	}
	return *o.Export
}

// GetExportOk returns a tuple with the Export field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetExportOk() (*AutoExportPolicy, bool) {
	if o == nil || IsNil(o.Export) {
		return nil, false
	}

	return o.Export, true
}

// HasExport returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasExport() bool {
	if o != nil && !IsNil(o.Export) {
		return true
	}

	return false
}

// SetExport gets a reference to the given AutoExportPolicy and assigns it to the Export field.
func (o *DiskBackupSnapshotSchedule) SetExport(v AutoExportPolicy) {
	o.Export = &v
}

// GetExtraRetentionSettings returns the ExtraRetentionSettings field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetExtraRetentionSettings() []ExtraRetentionSetting {
	if o == nil || IsNil(o.ExtraRetentionSettings) {
		var ret []ExtraRetentionSetting
		return ret
	}
	return *o.ExtraRetentionSettings
}

// GetExtraRetentionSettingsOk returns a tuple with the ExtraRetentionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetExtraRetentionSettingsOk() (*[]ExtraRetentionSetting, bool) {
	if o == nil || IsNil(o.ExtraRetentionSettings) {
		return nil, false
	}

	return o.ExtraRetentionSettings, true
}

// HasExtraRetentionSettings returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasExtraRetentionSettings() bool {
	if o != nil && !IsNil(o.ExtraRetentionSettings) {
		return true
	}

	return false
}

// SetExtraRetentionSettings gets a reference to the given []ExtraRetentionSetting and assigns it to the ExtraRetentionSettings field.
func (o *DiskBackupSnapshotSchedule) SetExtraRetentionSettings(v []ExtraRetentionSetting) {
	o.ExtraRetentionSettings = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *DiskBackupSnapshotSchedule) SetLinks(v []Link) {
	o.Links = &v
}

// GetNextSnapshot returns the NextSnapshot field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetNextSnapshot() time.Time {
	if o == nil || IsNil(o.NextSnapshot) {
		var ret time.Time
		return ret
	}
	return *o.NextSnapshot
}

// GetNextSnapshotOk returns a tuple with the NextSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetNextSnapshotOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NextSnapshot) {
		return nil, false
	}

	return o.NextSnapshot, true
}

// HasNextSnapshot returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasNextSnapshot() bool {
	if o != nil && !IsNil(o.NextSnapshot) {
		return true
	}

	return false
}

// SetNextSnapshot gets a reference to the given time.Time and assigns it to the NextSnapshot field.
func (o *DiskBackupSnapshotSchedule) SetNextSnapshot(v time.Time) {
	o.NextSnapshot = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetPolicies() []AdvancedDiskBackupSnapshotSchedulePolicy {
	if o == nil || IsNil(o.Policies) {
		var ret []AdvancedDiskBackupSnapshotSchedulePolicy
		return ret
	}
	return *o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetPoliciesOk() (*[]AdvancedDiskBackupSnapshotSchedulePolicy, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}

	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []AdvancedDiskBackupSnapshotSchedulePolicy and assigns it to the Policies field.
func (o *DiskBackupSnapshotSchedule) SetPolicies(v []AdvancedDiskBackupSnapshotSchedulePolicy) {
	o.Policies = &v
}

// GetReferenceHourOfDay returns the ReferenceHourOfDay field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetReferenceHourOfDay() int {
	if o == nil || IsNil(o.ReferenceHourOfDay) {
		var ret int
		return ret
	}
	return *o.ReferenceHourOfDay
}

// GetReferenceHourOfDayOk returns a tuple with the ReferenceHourOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetReferenceHourOfDayOk() (*int, bool) {
	if o == nil || IsNil(o.ReferenceHourOfDay) {
		return nil, false
	}

	return o.ReferenceHourOfDay, true
}

// HasReferenceHourOfDay returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasReferenceHourOfDay() bool {
	if o != nil && !IsNil(o.ReferenceHourOfDay) {
		return true
	}

	return false
}

// SetReferenceHourOfDay gets a reference to the given int and assigns it to the ReferenceHourOfDay field.
func (o *DiskBackupSnapshotSchedule) SetReferenceHourOfDay(v int) {
	o.ReferenceHourOfDay = &v
}

// GetReferenceMinuteOfHour returns the ReferenceMinuteOfHour field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetReferenceMinuteOfHour() int {
	if o == nil || IsNil(o.ReferenceMinuteOfHour) {
		var ret int
		return ret
	}
	return *o.ReferenceMinuteOfHour
}

// GetReferenceMinuteOfHourOk returns a tuple with the ReferenceMinuteOfHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetReferenceMinuteOfHourOk() (*int, bool) {
	if o == nil || IsNil(o.ReferenceMinuteOfHour) {
		return nil, false
	}

	return o.ReferenceMinuteOfHour, true
}

// HasReferenceMinuteOfHour returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasReferenceMinuteOfHour() bool {
	if o != nil && !IsNil(o.ReferenceMinuteOfHour) {
		return true
	}

	return false
}

// SetReferenceMinuteOfHour gets a reference to the given int and assigns it to the ReferenceMinuteOfHour field.
func (o *DiskBackupSnapshotSchedule) SetReferenceMinuteOfHour(v int) {
	o.ReferenceMinuteOfHour = &v
}

// GetRestoreWindowDays returns the RestoreWindowDays field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetRestoreWindowDays() int {
	if o == nil || IsNil(o.RestoreWindowDays) {
		var ret int
		return ret
	}
	return *o.RestoreWindowDays
}

// GetRestoreWindowDaysOk returns a tuple with the RestoreWindowDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetRestoreWindowDaysOk() (*int, bool) {
	if o == nil || IsNil(o.RestoreWindowDays) {
		return nil, false
	}

	return o.RestoreWindowDays, true
}

// HasRestoreWindowDays returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasRestoreWindowDays() bool {
	if o != nil && !IsNil(o.RestoreWindowDays) {
		return true
	}

	return false
}

// SetRestoreWindowDays gets a reference to the given int and assigns it to the RestoreWindowDays field.
func (o *DiskBackupSnapshotSchedule) SetRestoreWindowDays(v int) {
	o.RestoreWindowDays = &v
}

// GetUpdateSnapshots returns the UpdateSnapshots field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetUpdateSnapshots() bool {
	if o == nil || IsNil(o.UpdateSnapshots) {
		var ret bool
		return ret
	}
	return *o.UpdateSnapshots
}

// GetUpdateSnapshotsOk returns a tuple with the UpdateSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetUpdateSnapshotsOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateSnapshots) {
		return nil, false
	}

	return o.UpdateSnapshots, true
}

// HasUpdateSnapshots returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasUpdateSnapshots() bool {
	if o != nil && !IsNil(o.UpdateSnapshots) {
		return true
	}

	return false
}

// SetUpdateSnapshots gets a reference to the given bool and assigns it to the UpdateSnapshots field.
func (o *DiskBackupSnapshotSchedule) SetUpdateSnapshots(v bool) {
	o.UpdateSnapshots = &v
}

// GetUseOrgAndGroupNamesInExportPrefix returns the UseOrgAndGroupNamesInExportPrefix field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedule) GetUseOrgAndGroupNamesInExportPrefix() bool {
	if o == nil || IsNil(o.UseOrgAndGroupNamesInExportPrefix) {
		var ret bool
		return ret
	}
	return *o.UseOrgAndGroupNamesInExportPrefix
}

// GetUseOrgAndGroupNamesInExportPrefixOk returns a tuple with the UseOrgAndGroupNamesInExportPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedule) GetUseOrgAndGroupNamesInExportPrefixOk() (*bool, bool) {
	if o == nil || IsNil(o.UseOrgAndGroupNamesInExportPrefix) {
		return nil, false
	}

	return o.UseOrgAndGroupNamesInExportPrefix, true
}

// HasUseOrgAndGroupNamesInExportPrefix returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedule) HasUseOrgAndGroupNamesInExportPrefix() bool {
	if o != nil && !IsNil(o.UseOrgAndGroupNamesInExportPrefix) {
		return true
	}

	return false
}

// SetUseOrgAndGroupNamesInExportPrefix gets a reference to the given bool and assigns it to the UseOrgAndGroupNamesInExportPrefix field.
func (o *DiskBackupSnapshotSchedule) SetUseOrgAndGroupNamesInExportPrefix(v bool) {
	o.UseOrgAndGroupNamesInExportPrefix = &v
}

func (o DiskBackupSnapshotSchedule) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DiskBackupSnapshotSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoExportEnabled) {
		toSerialize["autoExportEnabled"] = o.AutoExportEnabled
	}
	if !IsNil(o.CopySettings) {
		toSerialize["copySettings"] = o.CopySettings
	}
	if !IsNil(o.DeleteCopiedBackups) {
		toSerialize["deleteCopiedBackups"] = o.DeleteCopiedBackups
	}
	if !IsNil(o.Export) {
		toSerialize["export"] = o.Export
	}
	if !IsNil(o.ExtraRetentionSettings) {
		toSerialize["extraRetentionSettings"] = o.ExtraRetentionSettings
	}
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	if !IsNil(o.ReferenceHourOfDay) {
		toSerialize["referenceHourOfDay"] = o.ReferenceHourOfDay
	}
	if !IsNil(o.ReferenceMinuteOfHour) {
		toSerialize["referenceMinuteOfHour"] = o.ReferenceMinuteOfHour
	}
	if !IsNil(o.RestoreWindowDays) {
		toSerialize["restoreWindowDays"] = o.RestoreWindowDays
	}
	if !IsNil(o.UpdateSnapshots) {
		toSerialize["updateSnapshots"] = o.UpdateSnapshots
	}
	if !IsNil(o.UseOrgAndGroupNamesInExportPrefix) {
		toSerialize["useOrgAndGroupNamesInExportPrefix"] = o.UseOrgAndGroupNamesInExportPrefix
	}
	return toSerialize, nil
}
