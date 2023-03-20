// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	Profile                  *string             `json:",omitempty"`
	DomainAllowList          []string            `json:",omitempty"`
	DomainRestrictionEnabled *bool               `json:",omitempty"`
	TestMode                 *string             `json:",omitempty"`
	FederationSettingsId     *string             `json:",omitempty"`
	IdentityProviderId       *string             `json:",omitempty"`
	OrgId                    *string             `json:",omitempty"`
	PostAuthRoleGrants       []string            `json:",omitempty"`
	RoleMappings             []RoleMappingView   `json:",omitempty"`
	UserConflicts            []FederatedUserView `json:",omitempty"`
}

// RoleMappingView is autogenerated from the json schema
type RoleMappingView struct {
	ExternalGroupName *string          `json:",omitempty"`
	Id                *string          `json:",omitempty"`
	RoleAssignments   []RoleAssignment `json:",omitempty"`
}

// RoleAssignment is autogenerated from the json schema
type RoleAssignment struct {
	ProjectId *string `json:",omitempty"`
	OrgId     *string `json:",omitempty"`
	Role      *string `json:",omitempty"`
}

// FederatedUserView is autogenerated from the json schema
type FederatedUserView struct {
	EmailAddress         *string `json:",omitempty"`
	FederationSettingsId *string `json:",omitempty"`
	FirstName            *string `json:",omitempty"`
	LastName             *string `json:",omitempty"`
	UserId               *string `json:",omitempty"`
}
