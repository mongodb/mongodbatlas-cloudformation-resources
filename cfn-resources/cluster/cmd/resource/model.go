// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	AdvancedSettings             *ProcessArgs              `json:",omitempty"`
	BackupEnabled                *bool                     `json:",omitempty"`
	BiConnector                  *BiConnector              `json:",omitempty"`
	ClusterType                  *string                   `json:",omitempty"`
	CreatedDate                  *string                   `json:",omitempty"`
	ConnectionStrings            *ConnectionStrings        `json:",omitempty"`
	DiskSizeGB                   *float64                  `json:",omitempty"`
	EncryptionAtRestProvider     *string                   `json:",omitempty"`
	Profile                      *string                   `json:",omitempty"`
	ProjectId                    *string                   `json:",omitempty"`
	Id                           *string                   `json:",omitempty"`
	Labels                       []Labels                  `json:",omitempty"`
	MongoDBMajorVersion          *string                   `json:",omitempty"`
	MongoDBVersion               *string                   `json:",omitempty"`
	Name                         *string                   `json:",omitempty"`
	Paused                       *bool                     `json:",omitempty"`
	PitEnabled                   *bool                     `json:",omitempty"`
	ReplicationSpecs             []AdvancedReplicationSpec `json:",omitempty"`
	RootCertType                 *string                   `json:",omitempty"`
	StateName                    *string                   `json:",omitempty"`
	VersionReleaseSystem         *string                   `json:",omitempty"`
	TerminationProtectionEnabled *bool                     `json:",omitempty"`
}

// ProcessArgs is autogenerated from the json schema
type ProcessArgs struct {
	DefaultReadConcern               *string  `json:",omitempty"`
	DefaultWriteConcern              *string  `json:",omitempty"`
	FailIndexKeyTooLong              *bool    `json:",omitempty"`
	JavascriptEnabled                *bool    `json:",omitempty"`
	MinimumEnabledTLSProtocol        *string  `json:",omitempty"`
	NoTableScan                      *bool    `json:",omitempty"`
	OplogSizeMB                      *int     `json:",omitempty"`
	SampleSizeBIConnector            *int     `json:",omitempty"`
	SampleRefreshIntervalBIConnector *int     `json:",omitempty"`
	OplogMinRetentionHours           *float64 `json:",omitempty"`
}

// BiConnector is autogenerated from the json schema
type BiConnector struct {
	ReadPreference *string `json:",omitempty"`
	Enabled        *bool   `json:",omitempty"`
}

// ConnectionStrings is autogenerated from the json schema
type ConnectionStrings struct {
	Standard          *string           `json:",omitempty"`
	StandardSrv       *string           `json:",omitempty"`
	Private           *string           `json:",omitempty"`
	PrivateSrv        *string           `json:",omitempty"`
	PrivateEndpoint   []PrivateEndpoint `json:",omitempty"`
	AwsPrivateLinkSrv *string           `json:",omitempty"`
	AwsPrivateLink    *string           `json:",omitempty"`
}

// PrivateEndpoint is autogenerated from the json schema
type PrivateEndpoint struct {
	ConnectionString    *string    `json:",omitempty"`
	Endpoints           []Endpoint `json:",omitempty"`
	SRVConnectionString *string    `json:",omitempty"`
	Type                *string    `json:",omitempty"`
}

// Endpoint is autogenerated from the json schema
type Endpoint struct {
	EndpointID   *string `json:",omitempty"`
	ProviderName *string `json:",omitempty"`
	Region       *string `json:",omitempty"`
}

// Labels is autogenerated from the json schema
type Labels struct {
	Key   *string `json:",omitempty"`
	Value *string `json:",omitempty"`
}

// AdvancedReplicationSpec is autogenerated from the json schema
type AdvancedReplicationSpec struct {
	ID                    *string                `json:",omitempty"`
	NumShards             *int                   `json:",omitempty"`
	AdvancedRegionConfigs []AdvancedRegionConfig `json:",omitempty"`
	ZoneName              *string                `json:",omitempty"`
}

// AdvancedRegionConfig is autogenerated from the json schema
type AdvancedRegionConfig struct {
	AnalyticsAutoScaling *AdvancedAutoScaling `json:",omitempty"`
	AutoScaling          *AdvancedAutoScaling `json:",omitempty"`
	RegionName           *string              `json:",omitempty"`
	ProviderName         *string              `json:",omitempty"`
	AnalyticsSpecs       *Specs               `json:",omitempty"`
	ElectableSpecs       *Specs               `json:",omitempty"`
	Priority             *int                 `json:",omitempty"`
	ReadOnlySpecs        *Specs               `json:",omitempty"`
}

// AdvancedAutoScaling is autogenerated from the json schema
type AdvancedAutoScaling struct {
	DiskGB  *DiskGB  `json:",omitempty"`
	Compute *Compute `json:",omitempty"`
}

// DiskGB is autogenerated from the json schema
type DiskGB struct {
	Enabled *bool `json:",omitempty"`
}

// Compute is autogenerated from the json schema
type Compute struct {
	Enabled          *bool   `json:",omitempty"`
	ScaleDownEnabled *bool   `json:",omitempty"`
	MinInstanceSize  *string `json:",omitempty"`
	MaxInstanceSize  *string `json:",omitempty"`
}

// Specs is autogenerated from the json schema
type Specs struct {
	DiskIOPS      *string `json:",omitempty"`
	EbsVolumeType *string `json:",omitempty"`
	InstanceSize  *string `json:",omitempty"`
	NodeCount     *int    `json:",omitempty"`
}
