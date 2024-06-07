// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	Id                                        *string                   `json:",omitempty"`
	ProjectId                                 *string                   `json:",omitempty"`
	InstanceName                              *string                   `json:",omitempty"`
	Comment                                   *string                   `json:",omitempty"`
	Profile                                   *string                   `json:",omitempty"`
	EndpointServiceName                       *string                   `json:",omitempty"`
	ErrorMessage                              *string                   `json:",omitempty"`
	ProviderName                              *string                   `json:",omitempty"`
	Status                                    *string                   `json:",omitempty"`
	CloudProviderEndpointId                   *string                   `json:",omitempty"`
	AwsPrivateEndpointMetaData                *string                   `json:",omitempty"`
	PrivateEndpointIpAddress                  *string                   `json:",omitempty"`
	CreateAndAssignAWSPrivateEndpoint         *bool                     `json:",omitempty"`
	AwsPrivateEndpointConfigurationProperties *AwsPrivateEndpointConfig `json:",omitempty"`
}

// AwsPrivateEndpointConfig is autogenerated from the json schema
type AwsPrivateEndpointConfig struct {
	VpcId                    *string  `json:",omitempty"`
	SubnetIds                []string `json:",omitempty"`
	InterfaceEndpointId      *string  `json:",omitempty"`
	AWSPrivateEndpointStatus *string  `json:",omitempty"`
	Region                   *string  `json:",omitempty"`
}
