// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	CloudProviderConfig *CloudProviderConfig `json:",omitempty"`
	DataProcessRegion   *DataProcessRegion   `json:",omitempty"`
	ProjectId           *string              `json:",omitempty"`
	TenantName          *string              `json:",omitempty"`
	SkipRoleValidation  *bool                `json:",omitempty"`
	Storage             *Storage             `json:",omitempty"`
	State               *string              `json:",omitempty"`
	HostNames           []string             `json:",omitempty"`
	Profile             *string              `json:",omitempty"`
}

// CloudProviderConfig is autogenerated from the json schema
type CloudProviderConfig struct {
	ExternalId        *string `json:",omitempty"`
	IamAssumedRoleARN *string `json:",omitempty"`
	IamUserARN        *string `json:",omitempty"`
	RoleId            *string `json:",omitempty"`
	TestS3Bucket      *string `json:",omitempty"`
}

// DataProcessRegion is autogenerated from the json schema
type DataProcessRegion struct {
	CloudProvider *string `json:",omitempty"`
	Region        *string `json:",omitempty"`
}

// Storage is autogenerated from the json schema
type Storage struct {
	Databases []Database `json:",omitempty"`
	Stores    []Store    `json:",omitempty"`
}

// Database is autogenerated from the json schema
type Database struct {
	Collections            []Collection `json:",omitempty"`
	MaxWildcardCollections *string      `json:",omitempty"`
	Name                   *string      `json:",omitempty"`
	Views                  []View       `json:",omitempty"`
}

// Collection is autogenerated from the json schema
type Collection struct {
	DataSources []DataSource `json:",omitempty"`
	Name        *string      `json:",omitempty"`
}

// DataSource is autogenerated from the json schema
type DataSource struct {
	AllowInsecure       *bool    `json:",omitempty"`
	Collection          *string  `json:",omitempty"`
	CollectionRegex     *string  `json:",omitempty"`
	Database            *string  `json:",omitempty"`
	DatabaseRegex       *string  `json:",omitempty"`
	DefaultFormat       *string  `json:",omitempty"`
	Path                *string  `json:",omitempty"`
	ProvenanceFieldName *string  `json:",omitempty"`
	StoreName           *string  `json:",omitempty"`
	Urls                []string `json:",omitempty"`
}

// View is autogenerated from the json schema
type View struct {
	Name     *string `json:",omitempty"`
	Pipeline *string `json:",omitempty"`
	Source   *string `json:",omitempty"`
}

// Store is autogenerated from the json schema
type Store struct {
	Name           *string         `json:",omitempty"`
	Provider       *string         `json:",omitempty"`
	ClusterName    *string         `json:",omitempty"`
	ProjectId      *string         `json:",omitempty"`
	ReadPreference *ReadPreference `json:",omitempty"`
}

// ReadPreference is autogenerated from the json schema
type ReadPreference struct {
	Mode                *string    `json:",omitempty"`
	MaxStalenessSeconds *string    `json:",omitempty"`
	TagSets             [][]TagSet `json:",omitempty"`
}

// TagSet is autogenerated from the json schema
type TagSet struct {
	Name  *string `json:",omitempty"`
	Value *string `json:",omitempty"`
}
