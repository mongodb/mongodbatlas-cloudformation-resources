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

// BackupSnapshotPart Characteristics that identify this snapshot.
type BackupSnapshotPart struct {
	// Unique 24-hexadecimal digit string that identifies the cluster with the snapshots you want to return.
	// Read only field.
	ClusterId *string `json:"clusterId,omitempty"`
	// Human-readable label that identifies the method of compression for the snapshot.
	// Read only field.
	CompressionSetting *string `json:"compressionSetting,omitempty"`
	// Total size of the data stored on each node in the cluster. This parameter expresses its value in bytes.
	// Read only field.
	DataSizeBytes *int64 `json:"dataSizeBytes,omitempty"`
	// Flag that indicates whether someone encrypted this snapshot.
	// Read only field.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty"`
	// Number that indicates the total size of the data files in bytes.
	// Read only field.
	FileSizeBytes *int64 `json:"fileSizeBytes,omitempty"`
	// Unique string that identifies the Key Management Interoperability (KMIP) master key used to encrypt the snapshot data. The resource returns this parameter when `\"parts.encryptionEnabled\" : true`.
	// Read only field.
	MasterKeyUUID *string `json:"masterKeyUUID,omitempty"`
	// Number that indicates the version of MongoDB that the replica set primary ran when MongoDB Cloud created the snapshot.
	// Read only field.
	MongodVersion *string `json:"mongodVersion,omitempty"`
	// Human-readable label that identifies the replica set.
	// Read only field.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
	// Number that indicates the total size of space allocated for document storage.
	// Read only field.
	StorageSizeBytes *int64 `json:"storageSizeBytes,omitempty"`
	// Human-readable label that identifies the type of server from which MongoDB Cloud took this snapshot.
	// Read only field.
	TypeName *string `json:"typeName,omitempty"`
}

// NewBackupSnapshotPart instantiates a new BackupSnapshotPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupSnapshotPart() *BackupSnapshotPart {
	this := BackupSnapshotPart{}
	return &this
}

// NewBackupSnapshotPartWithDefaults instantiates a new BackupSnapshotPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupSnapshotPartWithDefaults() *BackupSnapshotPart {
	this := BackupSnapshotPart{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}

	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *BackupSnapshotPart) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetCompressionSetting returns the CompressionSetting field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetCompressionSetting() string {
	if o == nil || IsNil(o.CompressionSetting) {
		var ret string
		return ret
	}
	return *o.CompressionSetting
}

// GetCompressionSettingOk returns a tuple with the CompressionSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetCompressionSettingOk() (*string, bool) {
	if o == nil || IsNil(o.CompressionSetting) {
		return nil, false
	}

	return o.CompressionSetting, true
}

// HasCompressionSetting returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasCompressionSetting() bool {
	if o != nil && !IsNil(o.CompressionSetting) {
		return true
	}

	return false
}

// SetCompressionSetting gets a reference to the given string and assigns it to the CompressionSetting field.
func (o *BackupSnapshotPart) SetCompressionSetting(v string) {
	o.CompressionSetting = &v
}

// GetDataSizeBytes returns the DataSizeBytes field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetDataSizeBytes() int64 {
	if o == nil || IsNil(o.DataSizeBytes) {
		var ret int64
		return ret
	}
	return *o.DataSizeBytes
}

// GetDataSizeBytesOk returns a tuple with the DataSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetDataSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DataSizeBytes) {
		return nil, false
	}

	return o.DataSizeBytes, true
}

// HasDataSizeBytes returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasDataSizeBytes() bool {
	if o != nil && !IsNil(o.DataSizeBytes) {
		return true
	}

	return false
}

// SetDataSizeBytes gets a reference to the given int64 and assigns it to the DataSizeBytes field.
func (o *BackupSnapshotPart) SetDataSizeBytes(v int64) {
	o.DataSizeBytes = &v
}

// GetEncryptionEnabled returns the EncryptionEnabled field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetEncryptionEnabled() bool {
	if o == nil || IsNil(o.EncryptionEnabled) {
		var ret bool
		return ret
	}
	return *o.EncryptionEnabled
}

// GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetEncryptionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptionEnabled) {
		return nil, false
	}

	return o.EncryptionEnabled, true
}

// HasEncryptionEnabled returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasEncryptionEnabled() bool {
	if o != nil && !IsNil(o.EncryptionEnabled) {
		return true
	}

	return false
}

// SetEncryptionEnabled gets a reference to the given bool and assigns it to the EncryptionEnabled field.
func (o *BackupSnapshotPart) SetEncryptionEnabled(v bool) {
	o.EncryptionEnabled = &v
}

// GetFileSizeBytes returns the FileSizeBytes field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetFileSizeBytes() int64 {
	if o == nil || IsNil(o.FileSizeBytes) {
		var ret int64
		return ret
	}
	return *o.FileSizeBytes
}

// GetFileSizeBytesOk returns a tuple with the FileSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetFileSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.FileSizeBytes) {
		return nil, false
	}

	return o.FileSizeBytes, true
}

// HasFileSizeBytes returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasFileSizeBytes() bool {
	if o != nil && !IsNil(o.FileSizeBytes) {
		return true
	}

	return false
}

// SetFileSizeBytes gets a reference to the given int64 and assigns it to the FileSizeBytes field.
func (o *BackupSnapshotPart) SetFileSizeBytes(v int64) {
	o.FileSizeBytes = &v
}

// GetMasterKeyUUID returns the MasterKeyUUID field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetMasterKeyUUID() string {
	if o == nil || IsNil(o.MasterKeyUUID) {
		var ret string
		return ret
	}
	return *o.MasterKeyUUID
}

// GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetMasterKeyUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.MasterKeyUUID) {
		return nil, false
	}

	return o.MasterKeyUUID, true
}

// HasMasterKeyUUID returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasMasterKeyUUID() bool {
	if o != nil && !IsNil(o.MasterKeyUUID) {
		return true
	}

	return false
}

// SetMasterKeyUUID gets a reference to the given string and assigns it to the MasterKeyUUID field.
func (o *BackupSnapshotPart) SetMasterKeyUUID(v string) {
	o.MasterKeyUUID = &v
}

// GetMongodVersion returns the MongodVersion field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetMongodVersion() string {
	if o == nil || IsNil(o.MongodVersion) {
		var ret string
		return ret
	}
	return *o.MongodVersion
}

// GetMongodVersionOk returns a tuple with the MongodVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetMongodVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongodVersion) {
		return nil, false
	}

	return o.MongodVersion, true
}

// HasMongodVersion returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasMongodVersion() bool {
	if o != nil && !IsNil(o.MongodVersion) {
		return true
	}

	return false
}

// SetMongodVersion gets a reference to the given string and assigns it to the MongodVersion field.
func (o *BackupSnapshotPart) SetMongodVersion(v string) {
	o.MongodVersion = &v
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetReplicaSetName() string {
	if o == nil || IsNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetName) {
		return nil, false
	}

	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasReplicaSetName() bool {
	if o != nil && !IsNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *BackupSnapshotPart) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

// GetStorageSizeBytes returns the StorageSizeBytes field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetStorageSizeBytes() int64 {
	if o == nil || IsNil(o.StorageSizeBytes) {
		var ret int64
		return ret
	}
	return *o.StorageSizeBytes
}

// GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetStorageSizeBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.StorageSizeBytes) {
		return nil, false
	}

	return o.StorageSizeBytes, true
}

// HasStorageSizeBytes returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasStorageSizeBytes() bool {
	if o != nil && !IsNil(o.StorageSizeBytes) {
		return true
	}

	return false
}

// SetStorageSizeBytes gets a reference to the given int64 and assigns it to the StorageSizeBytes field.
func (o *BackupSnapshotPart) SetStorageSizeBytes(v int64) {
	o.StorageSizeBytes = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise
func (o *BackupSnapshotPart) GetTypeName() string {
	if o == nil || IsNil(o.TypeName) {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupSnapshotPart) GetTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.TypeName) {
		return nil, false
	}

	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *BackupSnapshotPart) HasTypeName() bool {
	if o != nil && !IsNil(o.TypeName) {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *BackupSnapshotPart) SetTypeName(v string) {
	o.TypeName = &v
}

func (o BackupSnapshotPart) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o BackupSnapshotPart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
