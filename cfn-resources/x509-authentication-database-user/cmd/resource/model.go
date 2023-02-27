// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	Profile               *string       `json:",omitempty"`
	Links                 []Links       `json:",omitempty"`
	TotalCount            *int          `json:",omitempty"`
	CustomerX509          *CustomerX509 `json:",omitempty"`
	UserName              *string       `json:",omitempty"`
	MonthsUntilExpiration *int          `json:",omitempty"`
	ProjectId             *string       `json:",omitempty"`
	Results               []Certificate `json:",omitempty"`
}

// Links is autogenerated from the json schema
type Links struct {
	Rel  *string `json:",omitempty"`
	Href *string `json:",omitempty"`
}

// CustomerX509 is autogenerated from the json schema
type CustomerX509 struct {
	Cas *string `json:",omitempty"`
}

// Certificate is autogenerated from the json schema
type Certificate struct {
	UserName              *string `json:",omitempty"`
	CreatedAt             *string `json:",omitempty"`
	MonthsUntilExpiration *int    `json:",omitempty"`
	NotAfter              *string `json:",omitempty"`
	Subject               *string `json:",omitempty"`
	GroupId               *string `json:",omitempty"`
	Id                    *string `json:",omitempty"`
}
