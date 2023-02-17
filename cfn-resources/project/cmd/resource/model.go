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
	Name                      *string          `json:",omitempty"`
	OrgId                     *string          `json:",omitempty"`
	ProjectOwnerId            *string          `json:",omitempty"`
	WithDefaultAlertsSettings *bool            `json:",omitempty"`
	Id                        *string          `json:",omitempty"`
	Created                   *string          `json:",omitempty"`
	ClusterCount              *int             `json:",omitempty"`
	ProjectSettings           *ProjectSettings `json:",omitempty"`
	Profile                   *string          `json:",omitempty"`
	ProjectTeams              []ProjectTeam    `json:",omitempty"`
	ProjectApiKeys            []ProjectApiKey  `json:",omitempty"`
}

// ProjectSettings is autogenerated from the json schema
type ProjectSettings struct {
	IsCollectDatabaseSpecificsStatisticsEnabled *bool `json:",omitempty"`
	IsDataExplorerEnabled                       *bool `json:",omitempty"`
	IsPerformanceAdvisorEnabled                 *bool `json:",omitempty"`
	IsRealtimePerformancePanelEnabled           *bool `json:",omitempty"`
	IsSchemaAdvisorEnabled                      *bool `json:",omitempty"`
}

// ProjectTeam is autogenerated from the json schema
type ProjectTeam struct {
	TeamId    *string  `json:",omitempty"`
	RoleNames []string `json:",omitempty"`
}

// ProjectApiKey is autogenerated from the json schema
type ProjectApiKey struct {
	Key       *string  `json:",omitempty"`
	RoleNames []string `json:",omitempty"`
}
