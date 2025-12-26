// Copyright 2024 MongoDB Inc
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

package resource_test

import (
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-workspace/cmd/resource"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20250312010/admin"
)

var (
	workspaceName    = "name"
	projectID        = "projectId"
	cloudProvider    = "AWS"
	region           = "VIRGINIA_USA"
	tier             = "SP30"
	maxTierSize      = "100"
	name             = "name"
	kafkaType        = "Kafka"
	clusterType      = "Cluster"
	bootstrapServers = "bootstrapServers"
	clusterName      = "clusterName"
)

func TestNewStreamsTenant(t *testing.T) {
	testCases := []struct {
		input    *resource.Model
		expected *admin.StreamsTenant
		name     string
	}{
		{
			name: "Model with StreamConfig including Tier and MaxTierSize",
			input: &resource.Model{
				WorkspaceName: &workspaceName,
				ProjectId:     &projectID,
				DataProcessRegion: &resource.StreamsDataProcessRegion{
					CloudProvider: &cloudProvider,
					Region:        &region,
				},
				StreamConfig: &resource.StreamConfig{
					Tier:        &tier,
					MaxTierSize: &maxTierSize,
				},
			},
			expected: &admin.StreamsTenant{
				Name:    &workspaceName,
				GroupId: &projectID,
				DataProcessRegion: &admin.StreamsDataProcessRegion{
					CloudProvider: cloudProvider,
					Region:        region,
				},
				StreamConfig: &admin.StreamConfig{
					Tier:        &tier,
					MaxTierSize: &maxTierSize,
				},
			},
		},
		{
			name: "Model with StreamConfig with only Tier",
			input: &resource.Model{
				WorkspaceName: &workspaceName,
				ProjectId:     &projectID,
				DataProcessRegion: &resource.StreamsDataProcessRegion{
					CloudProvider: &cloudProvider,
					Region:        &region,
				},
				StreamConfig: &resource.StreamConfig{
					Tier: &tier,
				},
			},
			expected: &admin.StreamsTenant{
				Name:    &workspaceName,
				GroupId: &projectID,
				DataProcessRegion: &admin.StreamsDataProcessRegion{
					CloudProvider: cloudProvider,
					Region:        region,
				},
				StreamConfig: &admin.StreamConfig{
					Tier: &tier,
				},
			},
		},
		{
			name: "Model without StreamConfig",
			input: &resource.Model{
				WorkspaceName: &workspaceName,
				ProjectId:     &projectID,
				DataProcessRegion: &resource.StreamsDataProcessRegion{
					CloudProvider: &cloudProvider,
					Region:        &region,
				},
			},
			expected: &admin.StreamsTenant{
				Name:    &workspaceName,
				GroupId: &projectID,
				DataProcessRegion: &admin.StreamsDataProcessRegion{
					CloudProvider: cloudProvider,
					Region:        region,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := resource.NewStreamsTenant(tc.input)
			assert.Equal(t, tc.expected, result, "created model did not match expected output")
		})
	}
}

func TestNewModelConnections(t *testing.T) {
	testCases := []struct {
		name     string
		input    *[]admin.StreamsConnection
		expected []resource.StreamsConnection
	}{
		{
			name:     "StreamConfig is nil",
			input:    nil,
			expected: nil,
		},
		{
			name:     "StreamConfig is empty",
			input:    &[]admin.StreamsConnection{},
			expected: nil,
		},
		{
			name: "Connection type Kafka",
			input: &[]admin.StreamsConnection{
				{
					Name:             &name,
					Type:             &kafkaType,
					BootstrapServers: &bootstrapServers,
					Authentication:   &admin.StreamsKafkaAuthentication{},
					Security:         &admin.StreamsKafkaSecurity{},
				},
			},
			expected: []resource.StreamsConnection{
				{
					Name:             &name,
					Type:             &kafkaType,
					BootstrapServers: &bootstrapServers,
					Authentication:   &resource.StreamsKafkaAuthentication{},
					Security:         &resource.StreamsKafkaSecurity{},
				},
			},
		},
		{
			name: "Connection type Cluster",
			input: &[]admin.StreamsConnection{
				{
					Name:            &name,
					Type:            &clusterType,
					ClusterName:     &clusterName,
					DbRoleToExecute: &admin.DBRoleToExecute{},
				},
			},
			expected: []resource.StreamsConnection{
				{
					Name:            &name,
					Type:            &clusterType,
					ClusterName:     &clusterName,
					DbRoleToExecute: &resource.DBRoleToExecute{},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := resource.NewModelConnections(tc.input)
			assert.Equal(t, tc.expected, result, "created model did not match expected output")
		})
	}
}
