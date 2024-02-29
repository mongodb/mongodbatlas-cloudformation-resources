package resource_test

import (
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-instance/cmd/resource"
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
		input    *resource.Model
		expected *admin.StreamsTenant
	}{
		{
			name: "Model with StreamConfig",
			input: &resource.Model{
				InstanceName: &instanceName,
				ProjectId:    &projectID,
				DataProcessRegion: &resource.StreamsDataProcessRegion{
					CloudProvider: &cloudProvider,
					Region:        &region,
				},
				StreamConfig: &resource.StreamConfig{
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
			input: &resource.Model{
				InstanceName: &instanceName,
				ProjectId:    &projectID,
				DataProcessRegion: &resource.StreamsDataProcessRegion{
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
