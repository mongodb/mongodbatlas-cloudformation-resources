// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	Profile                           *string           `json:",omitempty"`
	Id                                *string           `json:",omitempty"`
	EndpointServiceName               *string           `json:",omitempty"`
	ErrorMessage                      *string           `json:",omitempty"`
	Status                            *string           `json:",omitempty"`
	GroupId                           *string           `json:",omitempty"`
	Region                            *string           `json:",omitempty"`
	CreateAndAssignAWSPrivateEndpoint *bool             `json:",omitempty"`
	PrivateEndpoints                  []PrivateEndpoint `json:",omitempty"`
	InterfaceEndpoints                []string          `json:",omitempty"`
}

// PrivateEndpoint is autogenerated from the json schema
type PrivateEndpoint struct {
	VpcId                      *string  `json:",omitempty"`
	SubnetIds                  []string `json:",omitempty"`
	InterfaceEndpointId        *string  `json:",omitempty"`
	AWSPrivateEndpointStatus   *string  `json:",omitempty"`
	AtlasPrivateEndpointStatus *string  `json:",omitempty"`
}
