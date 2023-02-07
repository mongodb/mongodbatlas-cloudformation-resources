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
	ApiKeys              *ApiKeyDefinition  `json:",omitempty"`
	ProjectId            *string            `json:",omitempty"`
	ClusterName          *string            `json:",omitempty"`
	ManagedNamespaces    []ManagedNamespace `json:",omitempty"`
	RemoveAllZoneMapping *bool              `json:",omitempty"`
	CustomZoneMappings   []ZoneMapping      `json:",omitempty"`
}

// ApiKeyDefinition is autogenerated from the json schema
type ApiKeyDefinition struct {
	PrivateKey *string `json:",omitempty"`
	PublicKey  *string `json:",omitempty"`
}

// ManagedNamespace is autogenerated from the json schema
type ManagedNamespace struct {
	Collection             *string `json:",omitempty"`
	CustomShardKey         *string `json:",omitempty"`
	Db                     *string `json:",omitempty"`
	IsCustomShardKeyHashed *bool   `json:",omitempty"`
	IsShardKeyUnique       *bool   `json:",omitempty"`
}

// ZoneMapping is autogenerated from the json schema
type ZoneMapping struct {
	Location *string `json:",omitempty"`
	Zone     *string `json:",omitempty"`
}
