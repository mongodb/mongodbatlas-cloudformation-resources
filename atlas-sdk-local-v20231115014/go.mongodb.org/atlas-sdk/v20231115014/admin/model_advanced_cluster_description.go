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

// AdvancedClusterDescription struct for AdvancedClusterDescription
type AdvancedClusterDescription struct {
	// If reconfiguration is necessary to regain a primary due to a regional outage, submit this field alongside your topology reconfiguration to request a new regional outage resistant topology. Forced reconfigurations during an outage of the majority of electable nodes carry a risk of data loss if replicated writes (even majority committed writes) have not been replicated to the new primary node. MongoDB Atlas docs contain more information. To proceed with an operation which carries that risk, set **acceptDataRisksAndForceReplicaSetReconfig** to the current date.
	AcceptDataRisksAndForceReplicaSetReconfig *time.Time                            `json:"acceptDataRisksAndForceReplicaSetReconfig,omitempty"`
	AdvancedConfiguration                     *ApiAtlasClusterAdvancedConfiguration `json:"advancedConfiguration,omitempty"`
	// Flag that indicates whether the cluster can perform backups. If set to `true`, the cluster can perform backups. You must set this value to `true` for NVMe clusters. Backup uses [Cloud Backups](https://docs.atlas.mongodb.com/backup/cloud-backup/overview/) for dedicated clusters and [Shared Cluster Backups](https://docs.atlas.mongodb.com/backup/shared-tier/overview/) for tenant clusters. If set to `false`, the cluster doesn't use backups.
	BackupEnabled *bool        `json:"backupEnabled,omitempty"`
	BiConnector   *BiConnector `json:"biConnector,omitempty"`
	// Configuration of nodes that comprise the cluster.
	ClusterType *string `json:"clusterType,omitempty"`
	// Config Server Management Mode for creating or updating a sharded cluster.  When configured as ATLAS_MANAGED, atlas may automatically switch the cluster's config server type for optimal performance and savings.  When configured as FIXED_TO_DEDICATED, the cluster will always use a dedicated config server.
	ConfigServerManagementMode *string `json:"configServerManagementMode,omitempty"`
	// Describes a sharded cluster's config server type.
	// Read only field.
	ConfigServerType  *string                   `json:"configServerType,omitempty"`
	ConnectionStrings *ClusterConnectionStrings `json:"connectionStrings,omitempty"`
	// Date and time when MongoDB Cloud created this cluster. This parameter expresses its value in ISO 8601 format in UTC.
	// Read only field.
	CreateDate *time.Time `json:"createDate,omitempty"`
	// Storage capacity of instance data volumes expressed in gigabytes. Increase this number to add capacity.   This value is not configurable on M0/M2/M5 clusters.   MongoDB Cloud requires this parameter if you set **replicationSpecs**.   If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value.    Storage charge calculations depend on whether you choose the default value or a custom value.   The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier.
	DiskSizeGB *float64 `json:"diskSizeGB,omitempty"`
	// Disk warming mode selection.
	DiskWarmingMode *string `json:"diskWarmingMode,omitempty"`
	// Cloud service provider that manages your customer keys to provide an additional layer of encryption at rest for the cluster. To enable customer key management for encryption at rest, the cluster **replicationSpecs[n].regionConfigs[m].{type}Specs.instanceSize** setting must be `M10` or higher and `\"backupEnabled\" : false` or omitted entirely.
	EncryptionAtRestProvider *string `json:"encryptionAtRestProvider,omitempty"`
	// Feature compatibility version of the cluster.
	// Read only field.
	FeatureCompatibilityVersion *string `json:"featureCompatibilityVersion,omitempty"`
	// Feature compatibility version expiration date.
	// Read only field.
	FeatureCompatibilityVersionExpirationDate *time.Time `json:"featureCompatibilityVersionExpirationDate,omitempty"`
	// Set this field to configure the Sharding Management Mode when creating a new Global Cluster.  When set to false, the management mode is set to Atlas-Managed Sharding. This mode fully manages the sharding of your Global Cluster and is built to provide a seamless deployment experience.  When set to true, the management mode is set to Self-Managed Sharding. This mode leaves the management of shards in your hands and is built to provide an advanced and flexible deployment experience.  This setting cannot be changed once the cluster is deployed.
	GlobalClusterSelfManagedSharding *bool `json:"globalClusterSelfManagedSharding,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the cluster.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster. The MongoDB Cloud console doesn't display your labels.  Cluster labels are deprecated and will be removed in a future release. We strongly recommend that you use Resource tags instead.
	// Deprecated
	Labels *[]ComponentLabel `json:"labels,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links                      *[]Link              `json:"links,omitempty"`
	MongoDBEmployeeAccessGrant *EmployeeAccessGrant `json:"mongoDBEmployeeAccessGrant,omitempty"`
	// MongoDB major version of the cluster.  On creation: Choose from the available versions of MongoDB, or leave unspecified for the current recommended default in the MongoDB Cloud platform. The recommended version is a recent Long Term Support version. The default is not guaranteed to be the most recently released version throughout the entire release cycle. For versions available in a specific project, see the linked documentation or use the API endpoint for [project LTS versions endpoint](#tag/Projects/operation/getProjectLTSVersions).   On update: Increase version only by 1 major version at a time. If the cluster is pinned to a MongoDB feature compatibility version exactly one major version below the current MongoDB version, the MongoDB version can be downgraded to the previous major version.
	MongoDBMajorVersion *string `json:"mongoDBMajorVersion,omitempty"`
	// Version of MongoDB that the cluster runs.
	// Read only field.
	MongoDBVersion *string `json:"mongoDBVersion,omitempty"`
	// Human-readable label that identifies the advanced cluster.
	Name *string `json:"name,omitempty"`
	// Flag that indicates whether the cluster is paused.
	Paused *bool `json:"paused,omitempty"`
	// Flag that indicates whether the cluster uses continuous cloud backups.
	PitEnabled *bool `json:"pitEnabled,omitempty"`
	// Set this field to configure the replica set scaling mode for your cluster.  By default, Atlas scales under WORKLOAD_TYPE. This mode allows Atlas to scale your analytics nodes in parallel to your operational nodes.  When configured as SEQUENTIAL, Atlas scales all nodes sequentially. This mode is intended for steady-state workloads and applications performing latency-sensitive secondary reads.  When configured as NODE_TYPE, Atlas scales your electable nodes in parallel with your read-only and analytics nodes. This mode is intended for large, dynamic workloads requiring frequent and timely cluster tier scaling. This is the fastest scaling strategy, but it might impact latency of workloads when performing extensive secondary reads.
	ReplicaSetScalingStrategy *string `json:"replicaSetScalingStrategy,omitempty"`
	// List of settings that configure your cluster regions. For Global Clusters, each object in the array represents a zone where your clusters nodes deploy. For non-Global sharded clusters and replica sets, this array has one object representing where your clusters nodes deploy.
	ReplicationSpecs *[]ReplicationSpec `json:"replicationSpecs,omitempty"`
	// Root Certificate Authority that MongoDB Cloud cluster uses. MongoDB Cloud supports Internet Security Research Group.
	RootCertType *string `json:"rootCertType,omitempty"`
	// Human-readable label that indicates the current operating condition of this cluster.
	// Read only field.
	StateName *string `json:"stateName,omitempty"`
	// List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the cluster.
	Tags *[]ResourceTag `json:"tags,omitempty"`
	// Flag that indicates whether termination protection is enabled on the cluster. If set to `true`, MongoDB Cloud won't delete the cluster. If set to `false`, MongoDB Cloud will delete the cluster.
	TerminationProtectionEnabled *bool `json:"terminationProtectionEnabled,omitempty"`
	// Method by which the cluster maintains the MongoDB versions. If value is `CONTINUOUS`, you must not specify **mongoDBMajorVersion**.
	VersionReleaseSystem *string `json:"versionReleaseSystem,omitempty"`
}

// NewAdvancedClusterDescription instantiates a new AdvancedClusterDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedClusterDescription() *AdvancedClusterDescription {
	this := AdvancedClusterDescription{}
	var backupEnabled bool = false
	this.BackupEnabled = &backupEnabled
	var configServerManagementMode string = "ATLAS_MANAGED"
	this.ConfigServerManagementMode = &configServerManagementMode
	var diskWarmingMode string = "FULLY_WARMED"
	this.DiskWarmingMode = &diskWarmingMode
	var replicaSetScalingStrategy string = "WORKLOAD_TYPE"
	this.ReplicaSetScalingStrategy = &replicaSetScalingStrategy
	var rootCertType string = "ISRGROOTX1"
	this.RootCertType = &rootCertType
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	var versionReleaseSystem string = "LTS"
	this.VersionReleaseSystem = &versionReleaseSystem
	return &this
}

// NewAdvancedClusterDescriptionWithDefaults instantiates a new AdvancedClusterDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedClusterDescriptionWithDefaults() *AdvancedClusterDescription {
	this := AdvancedClusterDescription{}
	var backupEnabled bool = false
	this.BackupEnabled = &backupEnabled
	var configServerManagementMode string = "ATLAS_MANAGED"
	this.ConfigServerManagementMode = &configServerManagementMode
	var diskWarmingMode string = "FULLY_WARMED"
	this.DiskWarmingMode = &diskWarmingMode
	var replicaSetScalingStrategy string = "WORKLOAD_TYPE"
	this.ReplicaSetScalingStrategy = &replicaSetScalingStrategy
	var rootCertType string = "ISRGROOTX1"
	this.RootCertType = &rootCertType
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	var versionReleaseSystem string = "LTS"
	this.VersionReleaseSystem = &versionReleaseSystem
	return &this
}

// GetAcceptDataRisksAndForceReplicaSetReconfig returns the AcceptDataRisksAndForceReplicaSetReconfig field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetAcceptDataRisksAndForceReplicaSetReconfig() time.Time {
	if o == nil || IsNil(o.AcceptDataRisksAndForceReplicaSetReconfig) {
		var ret time.Time
		return ret
	}
	return *o.AcceptDataRisksAndForceReplicaSetReconfig
}

// GetAcceptDataRisksAndForceReplicaSetReconfigOk returns a tuple with the AcceptDataRisksAndForceReplicaSetReconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetAcceptDataRisksAndForceReplicaSetReconfigOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AcceptDataRisksAndForceReplicaSetReconfig) {
		return nil, false
	}

	return o.AcceptDataRisksAndForceReplicaSetReconfig, true
}

// HasAcceptDataRisksAndForceReplicaSetReconfig returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasAcceptDataRisksAndForceReplicaSetReconfig() bool {
	if o != nil && !IsNil(o.AcceptDataRisksAndForceReplicaSetReconfig) {
		return true
	}

	return false
}

// SetAcceptDataRisksAndForceReplicaSetReconfig gets a reference to the given time.Time and assigns it to the AcceptDataRisksAndForceReplicaSetReconfig field.
func (o *AdvancedClusterDescription) SetAcceptDataRisksAndForceReplicaSetReconfig(v time.Time) {
	o.AcceptDataRisksAndForceReplicaSetReconfig = &v
}

// GetAdvancedConfiguration returns the AdvancedConfiguration field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetAdvancedConfiguration() ApiAtlasClusterAdvancedConfiguration {
	if o == nil || IsNil(o.AdvancedConfiguration) {
		var ret ApiAtlasClusterAdvancedConfiguration
		return ret
	}
	return *o.AdvancedConfiguration
}

// GetAdvancedConfigurationOk returns a tuple with the AdvancedConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetAdvancedConfigurationOk() (*ApiAtlasClusterAdvancedConfiguration, bool) {
	if o == nil || IsNil(o.AdvancedConfiguration) {
		return nil, false
	}

	return o.AdvancedConfiguration, true
}

// HasAdvancedConfiguration returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasAdvancedConfiguration() bool {
	if o != nil && !IsNil(o.AdvancedConfiguration) {
		return true
	}

	return false
}

// SetAdvancedConfiguration gets a reference to the given ApiAtlasClusterAdvancedConfiguration and assigns it to the AdvancedConfiguration field.
func (o *AdvancedClusterDescription) SetAdvancedConfiguration(v ApiAtlasClusterAdvancedConfiguration) {
	o.AdvancedConfiguration = &v
}

// GetBackupEnabled returns the BackupEnabled field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetBackupEnabled() bool {
	if o == nil || IsNil(o.BackupEnabled) {
		var ret bool
		return ret
	}
	return *o.BackupEnabled
}

// GetBackupEnabledOk returns a tuple with the BackupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetBackupEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BackupEnabled) {
		return nil, false
	}

	return o.BackupEnabled, true
}

// HasBackupEnabled returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasBackupEnabled() bool {
	if o != nil && !IsNil(o.BackupEnabled) {
		return true
	}

	return false
}

// SetBackupEnabled gets a reference to the given bool and assigns it to the BackupEnabled field.
func (o *AdvancedClusterDescription) SetBackupEnabled(v bool) {
	o.BackupEnabled = &v
}

// GetBiConnector returns the BiConnector field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetBiConnector() BiConnector {
	if o == nil || IsNil(o.BiConnector) {
		var ret BiConnector
		return ret
	}
	return *o.BiConnector
}

// GetBiConnectorOk returns a tuple with the BiConnector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetBiConnectorOk() (*BiConnector, bool) {
	if o == nil || IsNil(o.BiConnector) {
		return nil, false
	}

	return o.BiConnector, true
}

// HasBiConnector returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasBiConnector() bool {
	if o != nil && !IsNil(o.BiConnector) {
		return true
	}

	return false
}

// SetBiConnector gets a reference to the given BiConnector and assigns it to the BiConnector field.
func (o *AdvancedClusterDescription) SetBiConnector(v BiConnector) {
	o.BiConnector = &v
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetClusterType() string {
	if o == nil || IsNil(o.ClusterType) {
		var ret string
		return ret
	}
	return *o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetClusterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterType) {
		return nil, false
	}

	return o.ClusterType, true
}

// HasClusterType returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasClusterType() bool {
	if o != nil && !IsNil(o.ClusterType) {
		return true
	}

	return false
}

// SetClusterType gets a reference to the given string and assigns it to the ClusterType field.
func (o *AdvancedClusterDescription) SetClusterType(v string) {
	o.ClusterType = &v
}

// GetConfigServerManagementMode returns the ConfigServerManagementMode field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetConfigServerManagementMode() string {
	if o == nil || IsNil(o.ConfigServerManagementMode) {
		var ret string
		return ret
	}
	return *o.ConfigServerManagementMode
}

// GetConfigServerManagementModeOk returns a tuple with the ConfigServerManagementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetConfigServerManagementModeOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigServerManagementMode) {
		return nil, false
	}

	return o.ConfigServerManagementMode, true
}

// HasConfigServerManagementMode returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasConfigServerManagementMode() bool {
	if o != nil && !IsNil(o.ConfigServerManagementMode) {
		return true
	}

	return false
}

// SetConfigServerManagementMode gets a reference to the given string and assigns it to the ConfigServerManagementMode field.
func (o *AdvancedClusterDescription) SetConfigServerManagementMode(v string) {
	o.ConfigServerManagementMode = &v
}

// GetConfigServerType returns the ConfigServerType field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetConfigServerType() string {
	if o == nil || IsNil(o.ConfigServerType) {
		var ret string
		return ret
	}
	return *o.ConfigServerType
}

// GetConfigServerTypeOk returns a tuple with the ConfigServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetConfigServerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigServerType) {
		return nil, false
	}

	return o.ConfigServerType, true
}

// HasConfigServerType returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasConfigServerType() bool {
	if o != nil && !IsNil(o.ConfigServerType) {
		return true
	}

	return false
}

// SetConfigServerType gets a reference to the given string and assigns it to the ConfigServerType field.
func (o *AdvancedClusterDescription) SetConfigServerType(v string) {
	o.ConfigServerType = &v
}

// GetConnectionStrings returns the ConnectionStrings field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetConnectionStrings() ClusterConnectionStrings {
	if o == nil || IsNil(o.ConnectionStrings) {
		var ret ClusterConnectionStrings
		return ret
	}
	return *o.ConnectionStrings
}

// GetConnectionStringsOk returns a tuple with the ConnectionStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetConnectionStringsOk() (*ClusterConnectionStrings, bool) {
	if o == nil || IsNil(o.ConnectionStrings) {
		return nil, false
	}

	return o.ConnectionStrings, true
}

// HasConnectionStrings returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasConnectionStrings() bool {
	if o != nil && !IsNil(o.ConnectionStrings) {
		return true
	}

	return false
}

// SetConnectionStrings gets a reference to the given ClusterConnectionStrings and assigns it to the ConnectionStrings field.
func (o *AdvancedClusterDescription) SetConnectionStrings(v ClusterConnectionStrings) {
	o.ConnectionStrings = &v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetCreateDate() time.Time {
	if o == nil || IsNil(o.CreateDate) {
		var ret time.Time
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetCreateDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDate) {
		return nil, false
	}

	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasCreateDate() bool {
	if o != nil && !IsNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given time.Time and assigns it to the CreateDate field.
func (o *AdvancedClusterDescription) SetCreateDate(v time.Time) {
	o.CreateDate = &v
}

// GetDiskSizeGB returns the DiskSizeGB field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetDiskSizeGB() float64 {
	if o == nil || IsNil(o.DiskSizeGB) {
		var ret float64
		return ret
	}
	return *o.DiskSizeGB
}

// GetDiskSizeGBOk returns a tuple with the DiskSizeGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetDiskSizeGBOk() (*float64, bool) {
	if o == nil || IsNil(o.DiskSizeGB) {
		return nil, false
	}

	return o.DiskSizeGB, true
}

// HasDiskSizeGB returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasDiskSizeGB() bool {
	if o != nil && !IsNil(o.DiskSizeGB) {
		return true
	}

	return false
}

// SetDiskSizeGB gets a reference to the given float64 and assigns it to the DiskSizeGB field.
func (o *AdvancedClusterDescription) SetDiskSizeGB(v float64) {
	o.DiskSizeGB = &v
}

// GetDiskWarmingMode returns the DiskWarmingMode field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetDiskWarmingMode() string {
	if o == nil || IsNil(o.DiskWarmingMode) {
		var ret string
		return ret
	}
	return *o.DiskWarmingMode
}

// GetDiskWarmingModeOk returns a tuple with the DiskWarmingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetDiskWarmingModeOk() (*string, bool) {
	if o == nil || IsNil(o.DiskWarmingMode) {
		return nil, false
	}

	return o.DiskWarmingMode, true
}

// HasDiskWarmingMode returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasDiskWarmingMode() bool {
	if o != nil && !IsNil(o.DiskWarmingMode) {
		return true
	}

	return false
}

// SetDiskWarmingMode gets a reference to the given string and assigns it to the DiskWarmingMode field.
func (o *AdvancedClusterDescription) SetDiskWarmingMode(v string) {
	o.DiskWarmingMode = &v
}

// GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetEncryptionAtRestProvider() string {
	if o == nil || IsNil(o.EncryptionAtRestProvider) {
		var ret string
		return ret
	}
	return *o.EncryptionAtRestProvider
}

// GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetEncryptionAtRestProviderOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionAtRestProvider) {
		return nil, false
	}

	return o.EncryptionAtRestProvider, true
}

// HasEncryptionAtRestProvider returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasEncryptionAtRestProvider() bool {
	if o != nil && !IsNil(o.EncryptionAtRestProvider) {
		return true
	}

	return false
}

// SetEncryptionAtRestProvider gets a reference to the given string and assigns it to the EncryptionAtRestProvider field.
func (o *AdvancedClusterDescription) SetEncryptionAtRestProvider(v string) {
	o.EncryptionAtRestProvider = &v
}

// GetFeatureCompatibilityVersion returns the FeatureCompatibilityVersion field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetFeatureCompatibilityVersion() string {
	if o == nil || IsNil(o.FeatureCompatibilityVersion) {
		var ret string
		return ret
	}
	return *o.FeatureCompatibilityVersion
}

// GetFeatureCompatibilityVersionOk returns a tuple with the FeatureCompatibilityVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetFeatureCompatibilityVersionOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureCompatibilityVersion) {
		return nil, false
	}

	return o.FeatureCompatibilityVersion, true
}

// HasFeatureCompatibilityVersion returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasFeatureCompatibilityVersion() bool {
	if o != nil && !IsNil(o.FeatureCompatibilityVersion) {
		return true
	}

	return false
}

// SetFeatureCompatibilityVersion gets a reference to the given string and assigns it to the FeatureCompatibilityVersion field.
func (o *AdvancedClusterDescription) SetFeatureCompatibilityVersion(v string) {
	o.FeatureCompatibilityVersion = &v
}

// GetFeatureCompatibilityVersionExpirationDate returns the FeatureCompatibilityVersionExpirationDate field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetFeatureCompatibilityVersionExpirationDate() time.Time {
	if o == nil || IsNil(o.FeatureCompatibilityVersionExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.FeatureCompatibilityVersionExpirationDate
}

// GetFeatureCompatibilityVersionExpirationDateOk returns a tuple with the FeatureCompatibilityVersionExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetFeatureCompatibilityVersionExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FeatureCompatibilityVersionExpirationDate) {
		return nil, false
	}

	return o.FeatureCompatibilityVersionExpirationDate, true
}

// HasFeatureCompatibilityVersionExpirationDate returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasFeatureCompatibilityVersionExpirationDate() bool {
	if o != nil && !IsNil(o.FeatureCompatibilityVersionExpirationDate) {
		return true
	}

	return false
}

// SetFeatureCompatibilityVersionExpirationDate gets a reference to the given time.Time and assigns it to the FeatureCompatibilityVersionExpirationDate field.
func (o *AdvancedClusterDescription) SetFeatureCompatibilityVersionExpirationDate(v time.Time) {
	o.FeatureCompatibilityVersionExpirationDate = &v
}

// GetGlobalClusterSelfManagedSharding returns the GlobalClusterSelfManagedSharding field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetGlobalClusterSelfManagedSharding() bool {
	if o == nil || IsNil(o.GlobalClusterSelfManagedSharding) {
		var ret bool
		return ret
	}
	return *o.GlobalClusterSelfManagedSharding
}

// GetGlobalClusterSelfManagedShardingOk returns a tuple with the GlobalClusterSelfManagedSharding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetGlobalClusterSelfManagedShardingOk() (*bool, bool) {
	if o == nil || IsNil(o.GlobalClusterSelfManagedSharding) {
		return nil, false
	}

	return o.GlobalClusterSelfManagedSharding, true
}

// HasGlobalClusterSelfManagedSharding returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasGlobalClusterSelfManagedSharding() bool {
	if o != nil && !IsNil(o.GlobalClusterSelfManagedSharding) {
		return true
	}

	return false
}

// SetGlobalClusterSelfManagedSharding gets a reference to the given bool and assigns it to the GlobalClusterSelfManagedSharding field.
func (o *AdvancedClusterDescription) SetGlobalClusterSelfManagedSharding(v bool) {
	o.GlobalClusterSelfManagedSharding = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AdvancedClusterDescription) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdvancedClusterDescription) SetId(v string) {
	o.Id = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise
// Deprecated
func (o *AdvancedClusterDescription) GetLabels() []ComponentLabel {
	if o == nil || IsNil(o.Labels) {
		var ret []ComponentLabel
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AdvancedClusterDescription) GetLabelsOk() (*[]ComponentLabel, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}

	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []ComponentLabel and assigns it to the Labels field.
// Deprecated
func (o *AdvancedClusterDescription) SetLabels(v []ComponentLabel) {
	o.Labels = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *AdvancedClusterDescription) SetLinks(v []Link) {
	o.Links = &v
}

// GetMongoDBEmployeeAccessGrant returns the MongoDBEmployeeAccessGrant field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetMongoDBEmployeeAccessGrant() EmployeeAccessGrant {
	if o == nil || IsNil(o.MongoDBEmployeeAccessGrant) {
		var ret EmployeeAccessGrant
		return ret
	}
	return *o.MongoDBEmployeeAccessGrant
}

// GetMongoDBEmployeeAccessGrantOk returns a tuple with the MongoDBEmployeeAccessGrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetMongoDBEmployeeAccessGrantOk() (*EmployeeAccessGrant, bool) {
	if o == nil || IsNil(o.MongoDBEmployeeAccessGrant) {
		return nil, false
	}

	return o.MongoDBEmployeeAccessGrant, true
}

// HasMongoDBEmployeeAccessGrant returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasMongoDBEmployeeAccessGrant() bool {
	if o != nil && !IsNil(o.MongoDBEmployeeAccessGrant) {
		return true
	}

	return false
}

// SetMongoDBEmployeeAccessGrant gets a reference to the given EmployeeAccessGrant and assigns it to the MongoDBEmployeeAccessGrant field.
func (o *AdvancedClusterDescription) SetMongoDBEmployeeAccessGrant(v EmployeeAccessGrant) {
	o.MongoDBEmployeeAccessGrant = &v
}

// GetMongoDBMajorVersion returns the MongoDBMajorVersion field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetMongoDBMajorVersion() string {
	if o == nil || IsNil(o.MongoDBMajorVersion) {
		var ret string
		return ret
	}
	return *o.MongoDBMajorVersion
}

// GetMongoDBMajorVersionOk returns a tuple with the MongoDBMajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetMongoDBMajorVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongoDBMajorVersion) {
		return nil, false
	}

	return o.MongoDBMajorVersion, true
}

// HasMongoDBMajorVersion returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasMongoDBMajorVersion() bool {
	if o != nil && !IsNil(o.MongoDBMajorVersion) {
		return true
	}

	return false
}

// SetMongoDBMajorVersion gets a reference to the given string and assigns it to the MongoDBMajorVersion field.
func (o *AdvancedClusterDescription) SetMongoDBMajorVersion(v string) {
	o.MongoDBMajorVersion = &v
}

// GetMongoDBVersion returns the MongoDBVersion field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetMongoDBVersion() string {
	if o == nil || IsNil(o.MongoDBVersion) {
		var ret string
		return ret
	}
	return *o.MongoDBVersion
}

// GetMongoDBVersionOk returns a tuple with the MongoDBVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetMongoDBVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongoDBVersion) {
		return nil, false
	}

	return o.MongoDBVersion, true
}

// HasMongoDBVersion returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasMongoDBVersion() bool {
	if o != nil && !IsNil(o.MongoDBVersion) {
		return true
	}

	return false
}

// SetMongoDBVersion gets a reference to the given string and assigns it to the MongoDBVersion field.
func (o *AdvancedClusterDescription) SetMongoDBVersion(v string) {
	o.MongoDBVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdvancedClusterDescription) SetName(v string) {
	o.Name = &v
}

// GetPaused returns the Paused field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetPaused() bool {
	if o == nil || IsNil(o.Paused) {
		var ret bool
		return ret
	}
	return *o.Paused
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetPausedOk() (*bool, bool) {
	if o == nil || IsNil(o.Paused) {
		return nil, false
	}

	return o.Paused, true
}

// HasPaused returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasPaused() bool {
	if o != nil && !IsNil(o.Paused) {
		return true
	}

	return false
}

// SetPaused gets a reference to the given bool and assigns it to the Paused field.
func (o *AdvancedClusterDescription) SetPaused(v bool) {
	o.Paused = &v
}

// GetPitEnabled returns the PitEnabled field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetPitEnabled() bool {
	if o == nil || IsNil(o.PitEnabled) {
		var ret bool
		return ret
	}
	return *o.PitEnabled
}

// GetPitEnabledOk returns a tuple with the PitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetPitEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PitEnabled) {
		return nil, false
	}

	return o.PitEnabled, true
}

// HasPitEnabled returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasPitEnabled() bool {
	if o != nil && !IsNil(o.PitEnabled) {
		return true
	}

	return false
}

// SetPitEnabled gets a reference to the given bool and assigns it to the PitEnabled field.
func (o *AdvancedClusterDescription) SetPitEnabled(v bool) {
	o.PitEnabled = &v
}

// GetReplicaSetScalingStrategy returns the ReplicaSetScalingStrategy field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetReplicaSetScalingStrategy() string {
	if o == nil || IsNil(o.ReplicaSetScalingStrategy) {
		var ret string
		return ret
	}
	return *o.ReplicaSetScalingStrategy
}

// GetReplicaSetScalingStrategyOk returns a tuple with the ReplicaSetScalingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetReplicaSetScalingStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetScalingStrategy) {
		return nil, false
	}

	return o.ReplicaSetScalingStrategy, true
}

// HasReplicaSetScalingStrategy returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasReplicaSetScalingStrategy() bool {
	if o != nil && !IsNil(o.ReplicaSetScalingStrategy) {
		return true
	}

	return false
}

// SetReplicaSetScalingStrategy gets a reference to the given string and assigns it to the ReplicaSetScalingStrategy field.
func (o *AdvancedClusterDescription) SetReplicaSetScalingStrategy(v string) {
	o.ReplicaSetScalingStrategy = &v
}

// GetReplicationSpecs returns the ReplicationSpecs field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetReplicationSpecs() []ReplicationSpec {
	if o == nil || IsNil(o.ReplicationSpecs) {
		var ret []ReplicationSpec
		return ret
	}
	return *o.ReplicationSpecs
}

// GetReplicationSpecsOk returns a tuple with the ReplicationSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetReplicationSpecsOk() (*[]ReplicationSpec, bool) {
	if o == nil || IsNil(o.ReplicationSpecs) {
		return nil, false
	}

	return o.ReplicationSpecs, true
}

// HasReplicationSpecs returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasReplicationSpecs() bool {
	if o != nil && !IsNil(o.ReplicationSpecs) {
		return true
	}

	return false
}

// SetReplicationSpecs gets a reference to the given []ReplicationSpec and assigns it to the ReplicationSpecs field.
func (o *AdvancedClusterDescription) SetReplicationSpecs(v []ReplicationSpec) {
	o.ReplicationSpecs = &v
}

// GetRootCertType returns the RootCertType field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetRootCertType() string {
	if o == nil || IsNil(o.RootCertType) {
		var ret string
		return ret
	}
	return *o.RootCertType
}

// GetRootCertTypeOk returns a tuple with the RootCertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetRootCertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RootCertType) {
		return nil, false
	}

	return o.RootCertType, true
}

// HasRootCertType returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasRootCertType() bool {
	if o != nil && !IsNil(o.RootCertType) {
		return true
	}

	return false
}

// SetRootCertType gets a reference to the given string and assigns it to the RootCertType field.
func (o *AdvancedClusterDescription) SetRootCertType(v string) {
	o.RootCertType = &v
}

// GetStateName returns the StateName field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetStateName() string {
	if o == nil || IsNil(o.StateName) {
		var ret string
		return ret
	}
	return *o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetStateNameOk() (*string, bool) {
	if o == nil || IsNil(o.StateName) {
		return nil, false
	}

	return o.StateName, true
}

// HasStateName returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasStateName() bool {
	if o != nil && !IsNil(o.StateName) {
		return true
	}

	return false
}

// SetStateName gets a reference to the given string and assigns it to the StateName field.
func (o *AdvancedClusterDescription) SetStateName(v string) {
	o.StateName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetTags() []ResourceTag {
	if o == nil || IsNil(o.Tags) {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}

	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *AdvancedClusterDescription) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetTerminationProtectionEnabled() bool {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.TerminationProtectionEnabled
}

// GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetTerminationProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		return nil, false
	}

	return o.TerminationProtectionEnabled, true
}

// HasTerminationProtectionEnabled returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasTerminationProtectionEnabled() bool {
	if o != nil && !IsNil(o.TerminationProtectionEnabled) {
		return true
	}

	return false
}

// SetTerminationProtectionEnabled gets a reference to the given bool and assigns it to the TerminationProtectionEnabled field.
func (o *AdvancedClusterDescription) SetTerminationProtectionEnabled(v bool) {
	o.TerminationProtectionEnabled = &v
}

// GetVersionReleaseSystem returns the VersionReleaseSystem field value if set, zero value otherwise
func (o *AdvancedClusterDescription) GetVersionReleaseSystem() string {
	if o == nil || IsNil(o.VersionReleaseSystem) {
		var ret string
		return ret
	}
	return *o.VersionReleaseSystem
}

// GetVersionReleaseSystemOk returns a tuple with the VersionReleaseSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedClusterDescription) GetVersionReleaseSystemOk() (*string, bool) {
	if o == nil || IsNil(o.VersionReleaseSystem) {
		return nil, false
	}

	return o.VersionReleaseSystem, true
}

// HasVersionReleaseSystem returns a boolean if a field has been set.
func (o *AdvancedClusterDescription) HasVersionReleaseSystem() bool {
	if o != nil && !IsNil(o.VersionReleaseSystem) {
		return true
	}

	return false
}

// SetVersionReleaseSystem gets a reference to the given string and assigns it to the VersionReleaseSystem field.
func (o *AdvancedClusterDescription) SetVersionReleaseSystem(v string) {
	o.VersionReleaseSystem = &v
}

func (o AdvancedClusterDescription) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AdvancedClusterDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptDataRisksAndForceReplicaSetReconfig) {
		toSerialize["acceptDataRisksAndForceReplicaSetReconfig"] = o.AcceptDataRisksAndForceReplicaSetReconfig
	}
	if !IsNil(o.AdvancedConfiguration) {
		toSerialize["advancedConfiguration"] = o.AdvancedConfiguration
	}
	if !IsNil(o.BackupEnabled) {
		toSerialize["backupEnabled"] = o.BackupEnabled
	}
	if !IsNil(o.BiConnector) {
		toSerialize["biConnector"] = o.BiConnector
	}
	if !IsNil(o.ClusterType) {
		toSerialize["clusterType"] = o.ClusterType
	}
	if !IsNil(o.ConfigServerManagementMode) {
		toSerialize["configServerManagementMode"] = o.ConfigServerManagementMode
	}
	if !IsNil(o.ConnectionStrings) {
		toSerialize["connectionStrings"] = o.ConnectionStrings
	}
	if !IsNil(o.DiskSizeGB) {
		toSerialize["diskSizeGB"] = o.DiskSizeGB
	}
	if !IsNil(o.DiskWarmingMode) {
		toSerialize["diskWarmingMode"] = o.DiskWarmingMode
	}
	if !IsNil(o.EncryptionAtRestProvider) {
		toSerialize["encryptionAtRestProvider"] = o.EncryptionAtRestProvider
	}
	if !IsNil(o.GlobalClusterSelfManagedSharding) {
		toSerialize["globalClusterSelfManagedSharding"] = o.GlobalClusterSelfManagedSharding
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.MongoDBEmployeeAccessGrant) {
		toSerialize["mongoDBEmployeeAccessGrant"] = o.MongoDBEmployeeAccessGrant
	}
	if !IsNil(o.MongoDBMajorVersion) {
		toSerialize["mongoDBMajorVersion"] = o.MongoDBMajorVersion
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Paused) {
		toSerialize["paused"] = o.Paused
	}
	if !IsNil(o.PitEnabled) {
		toSerialize["pitEnabled"] = o.PitEnabled
	}
	if !IsNil(o.ReplicaSetScalingStrategy) {
		toSerialize["replicaSetScalingStrategy"] = o.ReplicaSetScalingStrategy
	}
	if !IsNil(o.ReplicationSpecs) {
		toSerialize["replicationSpecs"] = o.ReplicationSpecs
	}
	if !IsNil(o.RootCertType) {
		toSerialize["rootCertType"] = o.RootCertType
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TerminationProtectionEnabled) {
		toSerialize["terminationProtectionEnabled"] = o.TerminationProtectionEnabled
	}
	if !IsNil(o.VersionReleaseSystem) {
		toSerialize["versionReleaseSystem"] = o.VersionReleaseSystem
	}
	return toSerialize, nil
}
