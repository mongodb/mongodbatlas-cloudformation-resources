// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// LegacyAtlasCluster Group of settings that configure a MongoDB cluster.
type LegacyAtlasCluster struct {
	// If reconfiguration is necessary to regain a primary due to a regional outage, submit this field alongside your topology reconfiguration to request a new regional outage resistant topology. Forced reconfigurations during an outage of the majority of electable nodes carry a risk of data loss if replicated writes (even majority committed writes) have not been replicated to the new primary node. MongoDB Atlas docs contain more information. To proceed with an operation which carries that risk, set **acceptDataRisksAndForceReplicaSetReconfig** to the current date.
	AcceptDataRisksAndForceReplicaSetReconfig *time.Time                  `json:"acceptDataRisksAndForceReplicaSetReconfig,omitempty"`
	AutoScaling                               *ClusterAutoScalingSettings `json:"autoScaling,omitempty"`
	// Flag that indicates whether the cluster can perform backups. If set to `true`, the cluster can perform backups. You must set this value to `true` for NVMe clusters. Backup uses Cloud Backups for dedicated clusters and Shared Cluster Backups for tenant clusters. If set to `false`, the cluster doesn't use MongoDB Cloud backups.
	BackupEnabled *bool        `json:"backupEnabled,omitempty"`
	BiConnector   *BiConnector `json:"biConnector,omitempty"`
	// Configuration of nodes that comprise the cluster.
	ClusterType       *string                   `json:"clusterType,omitempty"`
	ConnectionStrings *ClusterConnectionStrings `json:"connectionStrings,omitempty"`
	// Date and time when MongoDB Cloud created this serverless instance. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.
	// Read only field.
	CreateDate *time.Time `json:"createDate,omitempty"`
	// Storage capacity that the host's root volume possesses expressed in gigabytes. Increase this number to add capacity. MongoDB Cloud requires this parameter if you set **replicationSpecs**. If you specify a disk size below the minimum (10 GB), this parameter defaults to the minimum disk size value. Storage charge calculations depend on whether you choose the default value or a custom value.  The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require more storage space, consider upgrading your cluster to a higher tier.
	DiskSizeGB *float64 `json:"diskSizeGB,omitempty"`
	// Disk warming mode selection.
	DiskWarmingMode *string `json:"diskWarmingMode,omitempty"`
	// Cloud service provider that manages your customer keys to provide an additional layer of Encryption at Rest for the cluster.
	EncryptionAtRestProvider *string `json:"encryptionAtRestProvider,omitempty"`
	// Set this field to configure the Sharding Management Mode when creating a new Global Cluster.  When set to false, the management mode is set to Atlas-Managed Sharding. This mode fully manages the sharding of your Global Cluster and is built to provide a seamless deployment experience.  When set to true, the management mode is set to Self-Managed Sharding. This mode leaves the management of shards in your hands and is built to provide an advanced and flexible deployment experience.  This setting cannot be changed once the cluster is deployed.
	GlobalClusterSelfManagedSharding *bool `json:"globalClusterSelfManagedSharding,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the cluster.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// Collection of key-value pairs between 1 to 255 characters in length that tag and categorize the cluster. The MongoDB Cloud console doesn't display your labels.  Cluster labels are deprecated and will be removed in a future release. We strongly recommend that you use [resource tags](https://dochub.mongodb.org/core/add-cluster-tag-atlas) instead.
	// Deprecated
	Labels *[]ComponentLabel `json:"labels,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Major MongoDB version of the cluster. MongoDB Cloud deploys the cluster with the latest stable release of the specified version.
	MongoDBMajorVersion *string `json:"mongoDBMajorVersion,omitempty"`
	// Version of MongoDB that the cluster runs.
	MongoDBVersion *string `json:"mongoDBVersion,omitempty"`
	// Base connection string that you can use to connect to the cluster. MongoDB Cloud displays the string only after the cluster starts, not while it builds the cluster.
	// Read only field.
	MongoURI *string `json:"mongoURI,omitempty"`
	// Date and time when someone last updated the connection string. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.
	// Read only field.
	MongoURIUpdated *time.Time `json:"mongoURIUpdated,omitempty"`
	// Connection string that you can use to connect to the cluster including the `replicaSet`, `ssl`, and `authSource` query parameters with values appropriate for the cluster. You may need to add MongoDB database users. The response returns this parameter once the cluster can receive requests, not while it builds the cluster.
	// Read only field.
	MongoURIWithOptions *string `json:"mongoURIWithOptions,omitempty"`
	// Human-readable label that identifies the cluster.
	Name *string `json:"name,omitempty"`
	// Number of shards up to 50 to deploy for a sharded cluster. The resource returns `1` to indicate a replica set and values of `2` and higher to indicate a sharded cluster. The returned value equals the number of shards in the cluster.
	NumShards *int `json:"numShards,omitempty"`
	// Flag that indicates whether the cluster is paused.
	Paused *bool `json:"paused,omitempty"`
	// Flag that indicates whether the cluster uses continuous cloud backups.
	PitEnabled *bool `json:"pitEnabled,omitempty"`
	// Flag that indicates whether the M10 or higher cluster can perform Cloud Backups. If set to `true`, the cluster can perform backups. If this and **backupEnabled** are set to `false`, the cluster doesn't use MongoDB Cloud backups.
	ProviderBackupEnabled *bool                    `json:"providerBackupEnabled,omitempty"`
	ProviderSettings      *ClusterProviderSettings `json:"providerSettings,omitempty"`
	// Number of members that belong to the replica set. Each member retains a copy of your databases, providing high availability and data redundancy. Use **replicationSpecs** instead.
	// Deprecated
	ReplicationFactor *int `json:"replicationFactor,omitempty"`
	// Physical location where MongoDB Cloud provisions cluster nodes.
	ReplicationSpec *map[string]RegionSpec `json:"replicationSpec,omitempty"`
	// List of settings that configure your cluster regions.  - For Global Clusters, each object in the array represents one zone where MongoDB Cloud deploys your clusters nodes. - For non-Global sharded clusters and replica sets, the single object represents where MongoDB Cloud deploys your clusters nodes.
	ReplicationSpecs *[]LegacyReplicationSpec `json:"replicationSpecs,omitempty"`
	// Root Certificate Authority that MongoDB Atlas clusters uses. MongoDB Cloud supports Internet Security Research Group.
	RootCertType *string `json:"rootCertType,omitempty"`
	// Connection string that you can use to connect to the cluster. The `+srv` modifier forces the connection to use Transport Layer Security (TLS). The `mongoURI` parameter lists additional options.
	// Read only field.
	SrvAddress *string `json:"srvAddress,omitempty"`
	// Human-readable label that indicates the current operating condition of the cluster.
	// Read only field.
	StateName *string `json:"stateName,omitempty"`
	// List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the cluster.
	Tags *[]ResourceTag `json:"tags,omitempty"`
	// Flag that indicates whether termination protection is enabled on the cluster. If set to `true`, MongoDB Cloud won't delete the cluster. If set to `false`, MongoDB Cloud will delete the cluster.
	TerminationProtectionEnabled *bool `json:"terminationProtectionEnabled,omitempty"`
	// Method by which the cluster maintains the MongoDB versions. If value is `CONTINUOUS`, you must not specify **mongoDBMajorVersion**.
	VersionReleaseSystem *string `json:"versionReleaseSystem,omitempty"`
}

// NewLegacyAtlasCluster instantiates a new LegacyAtlasCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyAtlasCluster() *LegacyAtlasCluster {
	this := LegacyAtlasCluster{}
	var diskWarmingMode string = "FULLY_WARMED"
	this.DiskWarmingMode = &diskWarmingMode
	var mongoDBMajorVersion string = "7.0"
	this.MongoDBMajorVersion = &mongoDBMajorVersion
	var numShards int = 1
	this.NumShards = &numShards
	var replicationFactor int = 3
	this.ReplicationFactor = &replicationFactor
	var rootCertType string = "ISRGROOTX1"
	this.RootCertType = &rootCertType
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	var versionReleaseSystem string = "LTS"
	this.VersionReleaseSystem = &versionReleaseSystem
	return &this
}

// NewLegacyAtlasClusterWithDefaults instantiates a new LegacyAtlasCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyAtlasClusterWithDefaults() *LegacyAtlasCluster {
	this := LegacyAtlasCluster{}
	var diskWarmingMode string = "FULLY_WARMED"
	this.DiskWarmingMode = &diskWarmingMode
	var mongoDBMajorVersion string = "7.0"
	this.MongoDBMajorVersion = &mongoDBMajorVersion
	var numShards int = 1
	this.NumShards = &numShards
	var replicationFactor int = 3
	this.ReplicationFactor = &replicationFactor
	var rootCertType string = "ISRGROOTX1"
	this.RootCertType = &rootCertType
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	var versionReleaseSystem string = "LTS"
	this.VersionReleaseSystem = &versionReleaseSystem
	return &this
}

// GetAcceptDataRisksAndForceReplicaSetReconfig returns the AcceptDataRisksAndForceReplicaSetReconfig field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetAcceptDataRisksAndForceReplicaSetReconfig() time.Time {
	if o == nil || IsNil(o.AcceptDataRisksAndForceReplicaSetReconfig) {
		var ret time.Time
		return ret
	}
	return *o.AcceptDataRisksAndForceReplicaSetReconfig
}

// GetAcceptDataRisksAndForceReplicaSetReconfigOk returns a tuple with the AcceptDataRisksAndForceReplicaSetReconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetAcceptDataRisksAndForceReplicaSetReconfigOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AcceptDataRisksAndForceReplicaSetReconfig) {
		return nil, false
	}

	return o.AcceptDataRisksAndForceReplicaSetReconfig, true
}

// HasAcceptDataRisksAndForceReplicaSetReconfig returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasAcceptDataRisksAndForceReplicaSetReconfig() bool {
	if o != nil && !IsNil(o.AcceptDataRisksAndForceReplicaSetReconfig) {
		return true
	}

	return false
}

// SetAcceptDataRisksAndForceReplicaSetReconfig gets a reference to the given time.Time and assigns it to the AcceptDataRisksAndForceReplicaSetReconfig field.
func (o *LegacyAtlasCluster) SetAcceptDataRisksAndForceReplicaSetReconfig(v time.Time) {
	o.AcceptDataRisksAndForceReplicaSetReconfig = &v
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetAutoScaling() ClusterAutoScalingSettings {
	if o == nil || IsNil(o.AutoScaling) {
		var ret ClusterAutoScalingSettings
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetAutoScalingOk() (*ClusterAutoScalingSettings, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}

	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given ClusterAutoScalingSettings and assigns it to the AutoScaling field.
func (o *LegacyAtlasCluster) SetAutoScaling(v ClusterAutoScalingSettings) {
	o.AutoScaling = &v
}

// GetBackupEnabled returns the BackupEnabled field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetBackupEnabled() bool {
	if o == nil || IsNil(o.BackupEnabled) {
		var ret bool
		return ret
	}
	return *o.BackupEnabled
}

// GetBackupEnabledOk returns a tuple with the BackupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetBackupEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BackupEnabled) {
		return nil, false
	}

	return o.BackupEnabled, true
}

// HasBackupEnabled returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasBackupEnabled() bool {
	if o != nil && !IsNil(o.BackupEnabled) {
		return true
	}

	return false
}

// SetBackupEnabled gets a reference to the given bool and assigns it to the BackupEnabled field.
func (o *LegacyAtlasCluster) SetBackupEnabled(v bool) {
	o.BackupEnabled = &v
}

// GetBiConnector returns the BiConnector field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetBiConnector() BiConnector {
	if o == nil || IsNil(o.BiConnector) {
		var ret BiConnector
		return ret
	}
	return *o.BiConnector
}

// GetBiConnectorOk returns a tuple with the BiConnector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetBiConnectorOk() (*BiConnector, bool) {
	if o == nil || IsNil(o.BiConnector) {
		return nil, false
	}

	return o.BiConnector, true
}

// HasBiConnector returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasBiConnector() bool {
	if o != nil && !IsNil(o.BiConnector) {
		return true
	}

	return false
}

// SetBiConnector gets a reference to the given BiConnector and assigns it to the BiConnector field.
func (o *LegacyAtlasCluster) SetBiConnector(v BiConnector) {
	o.BiConnector = &v
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetClusterType() string {
	if o == nil || IsNil(o.ClusterType) {
		var ret string
		return ret
	}
	return *o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetClusterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterType) {
		return nil, false
	}

	return o.ClusterType, true
}

// HasClusterType returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasClusterType() bool {
	if o != nil && !IsNil(o.ClusterType) {
		return true
	}

	return false
}

// SetClusterType gets a reference to the given string and assigns it to the ClusterType field.
func (o *LegacyAtlasCluster) SetClusterType(v string) {
	o.ClusterType = &v
}

// GetConnectionStrings returns the ConnectionStrings field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetConnectionStrings() ClusterConnectionStrings {
	if o == nil || IsNil(o.ConnectionStrings) {
		var ret ClusterConnectionStrings
		return ret
	}
	return *o.ConnectionStrings
}

// GetConnectionStringsOk returns a tuple with the ConnectionStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetConnectionStringsOk() (*ClusterConnectionStrings, bool) {
	if o == nil || IsNil(o.ConnectionStrings) {
		return nil, false
	}

	return o.ConnectionStrings, true
}

// HasConnectionStrings returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasConnectionStrings() bool {
	if o != nil && !IsNil(o.ConnectionStrings) {
		return true
	}

	return false
}

// SetConnectionStrings gets a reference to the given ClusterConnectionStrings and assigns it to the ConnectionStrings field.
func (o *LegacyAtlasCluster) SetConnectionStrings(v ClusterConnectionStrings) {
	o.ConnectionStrings = &v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetCreateDate() time.Time {
	if o == nil || IsNil(o.CreateDate) {
		var ret time.Time
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetCreateDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDate) {
		return nil, false
	}

	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasCreateDate() bool {
	if o != nil && !IsNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given time.Time and assigns it to the CreateDate field.
func (o *LegacyAtlasCluster) SetCreateDate(v time.Time) {
	o.CreateDate = &v
}

// GetDiskSizeGB returns the DiskSizeGB field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetDiskSizeGB() float64 {
	if o == nil || IsNil(o.DiskSizeGB) {
		var ret float64
		return ret
	}
	return *o.DiskSizeGB
}

// GetDiskSizeGBOk returns a tuple with the DiskSizeGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetDiskSizeGBOk() (*float64, bool) {
	if o == nil || IsNil(o.DiskSizeGB) {
		return nil, false
	}

	return o.DiskSizeGB, true
}

// HasDiskSizeGB returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasDiskSizeGB() bool {
	if o != nil && !IsNil(o.DiskSizeGB) {
		return true
	}

	return false
}

// SetDiskSizeGB gets a reference to the given float64 and assigns it to the DiskSizeGB field.
func (o *LegacyAtlasCluster) SetDiskSizeGB(v float64) {
	o.DiskSizeGB = &v
}

// GetDiskWarmingMode returns the DiskWarmingMode field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetDiskWarmingMode() string {
	if o == nil || IsNil(o.DiskWarmingMode) {
		var ret string
		return ret
	}
	return *o.DiskWarmingMode
}

// GetDiskWarmingModeOk returns a tuple with the DiskWarmingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetDiskWarmingModeOk() (*string, bool) {
	if o == nil || IsNil(o.DiskWarmingMode) {
		return nil, false
	}

	return o.DiskWarmingMode, true
}

// HasDiskWarmingMode returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasDiskWarmingMode() bool {
	if o != nil && !IsNil(o.DiskWarmingMode) {
		return true
	}

	return false
}

// SetDiskWarmingMode gets a reference to the given string and assigns it to the DiskWarmingMode field.
func (o *LegacyAtlasCluster) SetDiskWarmingMode(v string) {
	o.DiskWarmingMode = &v
}

// GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetEncryptionAtRestProvider() string {
	if o == nil || IsNil(o.EncryptionAtRestProvider) {
		var ret string
		return ret
	}
	return *o.EncryptionAtRestProvider
}

// GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetEncryptionAtRestProviderOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionAtRestProvider) {
		return nil, false
	}

	return o.EncryptionAtRestProvider, true
}

// HasEncryptionAtRestProvider returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasEncryptionAtRestProvider() bool {
	if o != nil && !IsNil(o.EncryptionAtRestProvider) {
		return true
	}

	return false
}

// SetEncryptionAtRestProvider gets a reference to the given string and assigns it to the EncryptionAtRestProvider field.
func (o *LegacyAtlasCluster) SetEncryptionAtRestProvider(v string) {
	o.EncryptionAtRestProvider = &v
}

// GetGlobalClusterSelfManagedSharding returns the GlobalClusterSelfManagedSharding field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetGlobalClusterSelfManagedSharding() bool {
	if o == nil || IsNil(o.GlobalClusterSelfManagedSharding) {
		var ret bool
		return ret
	}
	return *o.GlobalClusterSelfManagedSharding
}

// GetGlobalClusterSelfManagedShardingOk returns a tuple with the GlobalClusterSelfManagedSharding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetGlobalClusterSelfManagedShardingOk() (*bool, bool) {
	if o == nil || IsNil(o.GlobalClusterSelfManagedSharding) {
		return nil, false
	}

	return o.GlobalClusterSelfManagedSharding, true
}

// HasGlobalClusterSelfManagedSharding returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasGlobalClusterSelfManagedSharding() bool {
	if o != nil && !IsNil(o.GlobalClusterSelfManagedSharding) {
		return true
	}

	return false
}

// SetGlobalClusterSelfManagedSharding gets a reference to the given bool and assigns it to the GlobalClusterSelfManagedSharding field.
func (o *LegacyAtlasCluster) SetGlobalClusterSelfManagedSharding(v bool) {
	o.GlobalClusterSelfManagedSharding = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *LegacyAtlasCluster) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LegacyAtlasCluster) SetId(v string) {
	o.Id = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise
// Deprecated
func (o *LegacyAtlasCluster) GetLabels() []ComponentLabel {
	if o == nil || IsNil(o.Labels) {
		var ret []ComponentLabel
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LegacyAtlasCluster) GetLabelsOk() (*[]ComponentLabel, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}

	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []ComponentLabel and assigns it to the Labels field.
// Deprecated
func (o *LegacyAtlasCluster) SetLabels(v []ComponentLabel) {
	o.Labels = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *LegacyAtlasCluster) SetLinks(v []Link) {
	o.Links = &v
}

// GetMongoDBMajorVersion returns the MongoDBMajorVersion field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetMongoDBMajorVersion() string {
	if o == nil || IsNil(o.MongoDBMajorVersion) {
		var ret string
		return ret
	}
	return *o.MongoDBMajorVersion
}

// GetMongoDBMajorVersionOk returns a tuple with the MongoDBMajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetMongoDBMajorVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongoDBMajorVersion) {
		return nil, false
	}

	return o.MongoDBMajorVersion, true
}

// HasMongoDBMajorVersion returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasMongoDBMajorVersion() bool {
	if o != nil && !IsNil(o.MongoDBMajorVersion) {
		return true
	}

	return false
}

// SetMongoDBMajorVersion gets a reference to the given string and assigns it to the MongoDBMajorVersion field.
func (o *LegacyAtlasCluster) SetMongoDBMajorVersion(v string) {
	o.MongoDBMajorVersion = &v
}

// GetMongoDBVersion returns the MongoDBVersion field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetMongoDBVersion() string {
	if o == nil || IsNil(o.MongoDBVersion) {
		var ret string
		return ret
	}
	return *o.MongoDBVersion
}

// GetMongoDBVersionOk returns a tuple with the MongoDBVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetMongoDBVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongoDBVersion) {
		return nil, false
	}

	return o.MongoDBVersion, true
}

// HasMongoDBVersion returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasMongoDBVersion() bool {
	if o != nil && !IsNil(o.MongoDBVersion) {
		return true
	}

	return false
}

// SetMongoDBVersion gets a reference to the given string and assigns it to the MongoDBVersion field.
func (o *LegacyAtlasCluster) SetMongoDBVersion(v string) {
	o.MongoDBVersion = &v
}

// GetMongoURI returns the MongoURI field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetMongoURI() string {
	if o == nil || IsNil(o.MongoURI) {
		var ret string
		return ret
	}
	return *o.MongoURI
}

// GetMongoURIOk returns a tuple with the MongoURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetMongoURIOk() (*string, bool) {
	if o == nil || IsNil(o.MongoURI) {
		return nil, false
	}

	return o.MongoURI, true
}

// HasMongoURI returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasMongoURI() bool {
	if o != nil && !IsNil(o.MongoURI) {
		return true
	}

	return false
}

// SetMongoURI gets a reference to the given string and assigns it to the MongoURI field.
func (o *LegacyAtlasCluster) SetMongoURI(v string) {
	o.MongoURI = &v
}

// GetMongoURIUpdated returns the MongoURIUpdated field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetMongoURIUpdated() time.Time {
	if o == nil || IsNil(o.MongoURIUpdated) {
		var ret time.Time
		return ret
	}
	return *o.MongoURIUpdated
}

// GetMongoURIUpdatedOk returns a tuple with the MongoURIUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetMongoURIUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MongoURIUpdated) {
		return nil, false
	}

	return o.MongoURIUpdated, true
}

// HasMongoURIUpdated returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasMongoURIUpdated() bool {
	if o != nil && !IsNil(o.MongoURIUpdated) {
		return true
	}

	return false
}

// SetMongoURIUpdated gets a reference to the given time.Time and assigns it to the MongoURIUpdated field.
func (o *LegacyAtlasCluster) SetMongoURIUpdated(v time.Time) {
	o.MongoURIUpdated = &v
}

// GetMongoURIWithOptions returns the MongoURIWithOptions field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetMongoURIWithOptions() string {
	if o == nil || IsNil(o.MongoURIWithOptions) {
		var ret string
		return ret
	}
	return *o.MongoURIWithOptions
}

// GetMongoURIWithOptionsOk returns a tuple with the MongoURIWithOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetMongoURIWithOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.MongoURIWithOptions) {
		return nil, false
	}

	return o.MongoURIWithOptions, true
}

// HasMongoURIWithOptions returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasMongoURIWithOptions() bool {
	if o != nil && !IsNil(o.MongoURIWithOptions) {
		return true
	}

	return false
}

// SetMongoURIWithOptions gets a reference to the given string and assigns it to the MongoURIWithOptions field.
func (o *LegacyAtlasCluster) SetMongoURIWithOptions(v string) {
	o.MongoURIWithOptions = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LegacyAtlasCluster) SetName(v string) {
	o.Name = &v
}

// GetNumShards returns the NumShards field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetNumShards() int {
	if o == nil || IsNil(o.NumShards) {
		var ret int
		return ret
	}
	return *o.NumShards
}

// GetNumShardsOk returns a tuple with the NumShards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetNumShardsOk() (*int, bool) {
	if o == nil || IsNil(o.NumShards) {
		return nil, false
	}

	return o.NumShards, true
}

// HasNumShards returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasNumShards() bool {
	if o != nil && !IsNil(o.NumShards) {
		return true
	}

	return false
}

// SetNumShards gets a reference to the given int and assigns it to the NumShards field.
func (o *LegacyAtlasCluster) SetNumShards(v int) {
	o.NumShards = &v
}

// GetPaused returns the Paused field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetPaused() bool {
	if o == nil || IsNil(o.Paused) {
		var ret bool
		return ret
	}
	return *o.Paused
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetPausedOk() (*bool, bool) {
	if o == nil || IsNil(o.Paused) {
		return nil, false
	}

	return o.Paused, true
}

// HasPaused returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasPaused() bool {
	if o != nil && !IsNil(o.Paused) {
		return true
	}

	return false
}

// SetPaused gets a reference to the given bool and assigns it to the Paused field.
func (o *LegacyAtlasCluster) SetPaused(v bool) {
	o.Paused = &v
}

// GetPitEnabled returns the PitEnabled field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetPitEnabled() bool {
	if o == nil || IsNil(o.PitEnabled) {
		var ret bool
		return ret
	}
	return *o.PitEnabled
}

// GetPitEnabledOk returns a tuple with the PitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetPitEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PitEnabled) {
		return nil, false
	}

	return o.PitEnabled, true
}

// HasPitEnabled returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasPitEnabled() bool {
	if o != nil && !IsNil(o.PitEnabled) {
		return true
	}

	return false
}

// SetPitEnabled gets a reference to the given bool and assigns it to the PitEnabled field.
func (o *LegacyAtlasCluster) SetPitEnabled(v bool) {
	o.PitEnabled = &v
}

// GetProviderBackupEnabled returns the ProviderBackupEnabled field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetProviderBackupEnabled() bool {
	if o == nil || IsNil(o.ProviderBackupEnabled) {
		var ret bool
		return ret
	}
	return *o.ProviderBackupEnabled
}

// GetProviderBackupEnabledOk returns a tuple with the ProviderBackupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetProviderBackupEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ProviderBackupEnabled) {
		return nil, false
	}

	return o.ProviderBackupEnabled, true
}

// HasProviderBackupEnabled returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasProviderBackupEnabled() bool {
	if o != nil && !IsNil(o.ProviderBackupEnabled) {
		return true
	}

	return false
}

// SetProviderBackupEnabled gets a reference to the given bool and assigns it to the ProviderBackupEnabled field.
func (o *LegacyAtlasCluster) SetProviderBackupEnabled(v bool) {
	o.ProviderBackupEnabled = &v
}

// GetProviderSettings returns the ProviderSettings field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetProviderSettings() ClusterProviderSettings {
	if o == nil || IsNil(o.ProviderSettings) {
		var ret ClusterProviderSettings
		return ret
	}
	return *o.ProviderSettings
}

// GetProviderSettingsOk returns a tuple with the ProviderSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetProviderSettingsOk() (*ClusterProviderSettings, bool) {
	if o == nil || IsNil(o.ProviderSettings) {
		return nil, false
	}

	return o.ProviderSettings, true
}

// HasProviderSettings returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasProviderSettings() bool {
	if o != nil && !IsNil(o.ProviderSettings) {
		return true
	}

	return false
}

// SetProviderSettings gets a reference to the given ClusterProviderSettings and assigns it to the ProviderSettings field.
func (o *LegacyAtlasCluster) SetProviderSettings(v ClusterProviderSettings) {
	o.ProviderSettings = &v
}

// GetReplicationFactor returns the ReplicationFactor field value if set, zero value otherwise
// Deprecated
func (o *LegacyAtlasCluster) GetReplicationFactor() int {
	if o == nil || IsNil(o.ReplicationFactor) {
		var ret int
		return ret
	}
	return *o.ReplicationFactor
}

// GetReplicationFactorOk returns a tuple with the ReplicationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LegacyAtlasCluster) GetReplicationFactorOk() (*int, bool) {
	if o == nil || IsNil(o.ReplicationFactor) {
		return nil, false
	}

	return o.ReplicationFactor, true
}

// HasReplicationFactor returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasReplicationFactor() bool {
	if o != nil && !IsNil(o.ReplicationFactor) {
		return true
	}

	return false
}

// SetReplicationFactor gets a reference to the given int and assigns it to the ReplicationFactor field.
// Deprecated
func (o *LegacyAtlasCluster) SetReplicationFactor(v int) {
	o.ReplicationFactor = &v
}

// GetReplicationSpec returns the ReplicationSpec field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetReplicationSpec() map[string]RegionSpec {
	if o == nil || IsNil(o.ReplicationSpec) {
		var ret map[string]RegionSpec
		return ret
	}
	return *o.ReplicationSpec
}

// GetReplicationSpecOk returns a tuple with the ReplicationSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetReplicationSpecOk() (*map[string]RegionSpec, bool) {
	if o == nil || IsNil(o.ReplicationSpec) {
		return nil, false
	}

	return o.ReplicationSpec, true
}

// HasReplicationSpec returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasReplicationSpec() bool {
	if o != nil && !IsNil(o.ReplicationSpec) {
		return true
	}

	return false
}

// SetReplicationSpec gets a reference to the given map[string]RegionSpec and assigns it to the ReplicationSpec field.
func (o *LegacyAtlasCluster) SetReplicationSpec(v map[string]RegionSpec) {
	o.ReplicationSpec = &v
}

// GetReplicationSpecs returns the ReplicationSpecs field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetReplicationSpecs() []LegacyReplicationSpec {
	if o == nil || IsNil(o.ReplicationSpecs) {
		var ret []LegacyReplicationSpec
		return ret
	}
	return *o.ReplicationSpecs
}

// GetReplicationSpecsOk returns a tuple with the ReplicationSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetReplicationSpecsOk() (*[]LegacyReplicationSpec, bool) {
	if o == nil || IsNil(o.ReplicationSpecs) {
		return nil, false
	}

	return o.ReplicationSpecs, true
}

// HasReplicationSpecs returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasReplicationSpecs() bool {
	if o != nil && !IsNil(o.ReplicationSpecs) {
		return true
	}

	return false
}

// SetReplicationSpecs gets a reference to the given []LegacyReplicationSpec and assigns it to the ReplicationSpecs field.
func (o *LegacyAtlasCluster) SetReplicationSpecs(v []LegacyReplicationSpec) {
	o.ReplicationSpecs = &v
}

// GetRootCertType returns the RootCertType field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetRootCertType() string {
	if o == nil || IsNil(o.RootCertType) {
		var ret string
		return ret
	}
	return *o.RootCertType
}

// GetRootCertTypeOk returns a tuple with the RootCertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetRootCertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RootCertType) {
		return nil, false
	}

	return o.RootCertType, true
}

// HasRootCertType returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasRootCertType() bool {
	if o != nil && !IsNil(o.RootCertType) {
		return true
	}

	return false
}

// SetRootCertType gets a reference to the given string and assigns it to the RootCertType field.
func (o *LegacyAtlasCluster) SetRootCertType(v string) {
	o.RootCertType = &v
}

// GetSrvAddress returns the SrvAddress field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetSrvAddress() string {
	if o == nil || IsNil(o.SrvAddress) {
		var ret string
		return ret
	}
	return *o.SrvAddress
}

// GetSrvAddressOk returns a tuple with the SrvAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetSrvAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SrvAddress) {
		return nil, false
	}

	return o.SrvAddress, true
}

// HasSrvAddress returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasSrvAddress() bool {
	if o != nil && !IsNil(o.SrvAddress) {
		return true
	}

	return false
}

// SetSrvAddress gets a reference to the given string and assigns it to the SrvAddress field.
func (o *LegacyAtlasCluster) SetSrvAddress(v string) {
	o.SrvAddress = &v
}

// GetStateName returns the StateName field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetStateName() string {
	if o == nil || IsNil(o.StateName) {
		var ret string
		return ret
	}
	return *o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetStateNameOk() (*string, bool) {
	if o == nil || IsNil(o.StateName) {
		return nil, false
	}

	return o.StateName, true
}

// HasStateName returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasStateName() bool {
	if o != nil && !IsNil(o.StateName) {
		return true
	}

	return false
}

// SetStateName gets a reference to the given string and assigns it to the StateName field.
func (o *LegacyAtlasCluster) SetStateName(v string) {
	o.StateName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetTags() []ResourceTag {
	if o == nil || IsNil(o.Tags) {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}

	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *LegacyAtlasCluster) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetTerminationProtectionEnabled() bool {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.TerminationProtectionEnabled
}

// GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetTerminationProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		return nil, false
	}

	return o.TerminationProtectionEnabled, true
}

// HasTerminationProtectionEnabled returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasTerminationProtectionEnabled() bool {
	if o != nil && !IsNil(o.TerminationProtectionEnabled) {
		return true
	}

	return false
}

// SetTerminationProtectionEnabled gets a reference to the given bool and assigns it to the TerminationProtectionEnabled field.
func (o *LegacyAtlasCluster) SetTerminationProtectionEnabled(v bool) {
	o.TerminationProtectionEnabled = &v
}

// GetVersionReleaseSystem returns the VersionReleaseSystem field value if set, zero value otherwise
func (o *LegacyAtlasCluster) GetVersionReleaseSystem() string {
	if o == nil || IsNil(o.VersionReleaseSystem) {
		var ret string
		return ret
	}
	return *o.VersionReleaseSystem
}

// GetVersionReleaseSystemOk returns a tuple with the VersionReleaseSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAtlasCluster) GetVersionReleaseSystemOk() (*string, bool) {
	if o == nil || IsNil(o.VersionReleaseSystem) {
		return nil, false
	}

	return o.VersionReleaseSystem, true
}

// HasVersionReleaseSystem returns a boolean if a field has been set.
func (o *LegacyAtlasCluster) HasVersionReleaseSystem() bool {
	if o != nil && !IsNil(o.VersionReleaseSystem) {
		return true
	}

	return false
}

// SetVersionReleaseSystem gets a reference to the given string and assigns it to the VersionReleaseSystem field.
func (o *LegacyAtlasCluster) SetVersionReleaseSystem(v string) {
	o.VersionReleaseSystem = &v
}

func (o LegacyAtlasCluster) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o LegacyAtlasCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptDataRisksAndForceReplicaSetReconfig) {
		toSerialize["acceptDataRisksAndForceReplicaSetReconfig"] = o.AcceptDataRisksAndForceReplicaSetReconfig
	}
	if !IsNil(o.AutoScaling) {
		toSerialize["autoScaling"] = o.AutoScaling
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
	if !IsNil(o.MongoDBMajorVersion) {
		toSerialize["mongoDBMajorVersion"] = o.MongoDBMajorVersion
	}
	if !IsNil(o.MongoDBVersion) {
		toSerialize["mongoDBVersion"] = o.MongoDBVersion
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumShards) {
		toSerialize["numShards"] = o.NumShards
	}
	if !IsNil(o.Paused) {
		toSerialize["paused"] = o.Paused
	}
	if !IsNil(o.PitEnabled) {
		toSerialize["pitEnabled"] = o.PitEnabled
	}
	if !IsNil(o.ProviderBackupEnabled) {
		toSerialize["providerBackupEnabled"] = o.ProviderBackupEnabled
	}
	if !IsNil(o.ProviderSettings) {
		toSerialize["providerSettings"] = o.ProviderSettings
	}
	if !IsNil(o.ReplicationFactor) {
		toSerialize["replicationFactor"] = o.ReplicationFactor
	}
	if !IsNil(o.ReplicationSpec) {
		toSerialize["replicationSpec"] = o.ReplicationSpec
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
