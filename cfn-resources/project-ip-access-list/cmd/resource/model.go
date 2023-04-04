// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	AccessList  []AccessListDefinition `json:",omitempty"`
	ProjectId   *string                `json:",omitempty"`
	TotalCount  *int                   `json:",omitempty"`
	Profile     *string                `json:",omitempty"`
	ListOptions *ListOptions           `json:",omitempty"`
}

// AccessListDefinition is autogenerated from the json schema
type AccessListDefinition struct {
	DeleteAfterDate  *string `json:",omitempty"`
	AwsSecurityGroup *string `json:",omitempty"`
	CIDRBlock        *string `json:",omitempty"`
	Comment          *string `json:",omitempty"`
	IPAddress        *string `json:",omitempty"`
	ProjectId        *string `json:",omitempty"`
}

// ListOptions is autogenerated from the json schema
type ListOptions struct {
	PageNum      *int  `json:",omitempty"`
	ItemsPerPage *int  `json:",omitempty"`
	IncludeCount *bool `json:",omitempty"`
}
