// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	Profile               *string                          `json:",omitempty"`
	BindUsername          *string                          `json:",omitempty"`
	Status                *string                          `json:",omitempty"`
	Hostname              *string                          `json:",omitempty"`
	AuthenticationEnabled *bool                            `json:",omitempty"`
	AuthorizationEnabled  *bool                            `json:",omitempty"`
	CaCertificate         *string                          `json:",omitempty"`
	AuthzQueryTemplate    *string                          `json:",omitempty"`
	BindPassword          *string                          `json:",omitempty"`
	ProjectId             *string                          `json:",omitempty"`
	Port                  *int                             `json:",omitempty"`
	UserToDNMapping       []ApiAtlasNDSUserToDNMappingView `json:",omitempty"`
}

// ApiAtlasNDSUserToDNMappingView is autogenerated from the json schema
type ApiAtlasNDSUserToDNMappingView struct {
	LdapQuery    *string `json:",omitempty"`
	Match        *string `json:",omitempty"`
	Substitution *string `json:",omitempty"`
}
