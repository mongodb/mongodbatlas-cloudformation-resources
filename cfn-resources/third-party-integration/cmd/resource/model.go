// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	ProjectId                *string `json:",omitempty"`
	Profile                  *string `json:",omitempty"`
	Type                     *string `json:",omitempty"`
	ApiKey                   *string `json:",omitempty"`
	Region                   *string `json:",omitempty"`
	ServiceKey               *string `json:",omitempty"`
	ApiToken                 *string `json:",omitempty"`
	TeamName                 *string `json:",omitempty"`
	ChannelName              *string `json:",omitempty"`
	RoutingKey               *string `json:",omitempty"`
	Url                      *string `json:",omitempty"`
	Secret                   *string `json:",omitempty"`
	MicrosoftTeamsWebhookUrl *string `json:",omitempty"`
	UserName                 *string `json:",omitempty"`
	Password                 *string `json:",omitempty"`
	ServiceDiscovery         *string `json:",omitempty"`
	Scheme                   *string `json:",omitempty"`
	Enabled                  *bool   `json:",omitempty"`
	ListenAddress            *string `json:",omitempty"`
	TlsPemPath               *string `json:",omitempty"`
}
