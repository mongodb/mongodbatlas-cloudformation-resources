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

package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

var (
	instanceName     = "name"
	projectID        = "projectId"
	cloudProvider    = "AWS"
	region           = "VIRGINIA_USA"
	tier             = "SP30"
	name             = "name"
	kafkaType        = "Kafka"
	clusterType      = "Cluster"
	bootstrapServers = "bootstrapServers"
	clusterName      = "clusterName"
)

func TestNewStreamsTenant(t *testing.T) {
	testCases := []struct {
		name     string
		input    *Model
		expected *admin.StreamsTenant
	}{
		{
			name: "Model with StreamConfig",
			input: &Model{
				InstanceName: &instanceName,
				ProjectId:    &projectID,
				DataProcessRegion: &StreamsDataProcessRegion{
					CloudProvider: &cloudProvider,
					Region:        &region,
				},
				StreamConfig: &StreamConfig{
					Tier: &tier,
				},
			},
			expected: &admin.StreamsTenant{
				Name:    &instanceName,
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
			input: &Model{
				InstanceName: &instanceName,
				ProjectId:    &projectID,
				DataProcessRegion: &StreamsDataProcessRegion{
					CloudProvider: &cloudProvider,
					Region:        &region,
				},
			},
			expected: &admin.StreamsTenant{
				Name:    &instanceName,
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
			result := newStreamsTenant(tc.input)
			assert.Equal(t, tc.expected, result, "created model did not match expected output")
		})
	}
}

func TestNewModelConnections(t *testing.T) {
	testCases := []struct {
		name     string
		input    *[]admin.StreamsConnection
		expected []StreamsConnection
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
			expected: []StreamsConnection{
				{
					Name:             &name,
					Type:             &kafkaType,
					BootstrapServers: &bootstrapServers,
					Authentication:   &StreamsKafkaAuthentication{},
					Security:         &StreamsKafkaSecurity{},
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
			expected: []StreamsConnection{
				{
					Name:            &name,
					Type:            &clusterType,
					ClusterName:     &clusterName,
					DbRoleToExecute: &DBRoleToExecute{},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := newModelConnections(tc.input)
			assert.Equal(t, tc.expected, result, "created model did not match expected output")
		})
	}
}
