// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	RealmConfig     *RealmConfig    `json:",omitempty"`
	DatabaseTrigger *DatabaseConfig `json:",omitempty"`
	AuthTrigger     *AuthConfig     `json:",omitempty"`
	ScheduleTrigger *ScheduleConfig `json:",omitempty"`
	Id              *string         `json:",omitempty"`
	Name            *string         `json:",omitempty"`
	Type            *string         `json:",omitempty"`
	Disabled        *bool           `json:",omitempty"`
	FunctionId      *string         `json:",omitempty"`
	FunctionName    *string         `json:",omitempty"`
	EventProcessors *Event          `json:",omitempty"`
	AppId           *string         `json:",omitempty"`
	ProjectId       *string         `json:",omitempty"`
}

// RealmConfig is autogenerated from the json schema
type RealmConfig struct {
	PrivateKey   *string `json:",omitempty"`
	PublicKey    *string `json:",omitempty"`
	BaseURL      *string `json:",omitempty"`
	RealmBaseURL *string `json:",omitempty"`
}

// DatabaseConfig is autogenerated from the json schema
type DatabaseConfig struct {
	ServiceId                *string                `json:",omitempty"`
	Database                 *string                `json:",omitempty"`
	Collection               *string                `json:",omitempty"`
	OperationTypes           []string               `json:",omitempty"`
	Match                    map[string]interface{} `json:",omitempty"`
	FullDocument             *bool                  `json:",omitempty"`
	FullDocumentBeforeChange *bool                  `json:",omitempty"`
	SkipCatchupEvents        *bool                  `json:",omitempty"`
	TolerateResumeErrors     *bool                  `json:",omitempty"`
	Unordered                *bool                  `json:",omitempty"`
}

// AuthConfig is autogenerated from the json schema
type AuthConfig struct {
	OperationType *string  `json:",omitempty"`
	Providers     []string `json:",omitempty"`
}

// ScheduleConfig is autogenerated from the json schema
type ScheduleConfig struct {
	Schedule          *string `json:",omitempty"`
	SkipcatchupEvents *bool   `json:",omitempty"`
}

// Event is autogenerated from the json schema
type Event struct {
	FUNCTION       *FUNCTION       `json:",omitempty"`
	AWSEVENTBRIDGE *AWSEVENTBRIDGE `json:",omitempty"`
}

// FUNCTION is autogenerated from the json schema
type FUNCTION struct {
	FuncConfig *FuncConfig `json:",omitempty"`
}

// FuncConfig is autogenerated from the json schema
type FuncConfig struct {
	FunctionId   *string `json:",omitempty"`
	FunctionName *string `json:",omitempty"`
}

// AWSEVENTBRIDGE is autogenerated from the json schema
type AWSEVENTBRIDGE struct {
	AWSConfig *AWSConfig `json:",omitempty"`
}

// AWSConfig is autogenerated from the json schema
type AWSConfig struct {
	AccountId           *string `json:",omitempty"`
	Region              *string `json:",omitempty"`
	ExtendedJsonEnabled *bool   `json:",omitempty"`
}
