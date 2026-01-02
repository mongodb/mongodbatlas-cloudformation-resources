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

	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/aws/smithy-go/ptr"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-connection/cmd/resource"
	"github.com/stretchr/testify/assert"
)

func TestNewModelDBRoleToExecute(t *testing.T) {
	tests := []struct {
		input    *admin20250312010.DBRoleToExecute
		expected *resource.DBRoleToExecute
		name     string
	}{
		{
			name:     "Nil Input",
			input:    nil,
			expected: nil,
		},
		{
			name: "Valid Input",
			input: &admin20250312010.DBRoleToExecute{
				Role: ptr.String("readWrite"),
				Type: ptr.String("BUILT_IN"),
			},
			expected: &resource.DBRoleToExecute{
				Role: ptr.String("readWrite"),
				Type: ptr.String("BUILT_IN"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := resource.NewModelDBRoleToExecute(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestNewModelAuthentication(t *testing.T) {
	tests := []struct {
		input    *admin20250312010.StreamsKafkaAuthentication
		expected *resource.StreamsKafkaAuthentication
		name     string
	}{
		{
			name:     "Nil Input",
			input:    nil,
			expected: nil,
		},
		{
			name: "Valid Input",
			input: &admin20250312010.StreamsKafkaAuthentication{
				Mechanism: ptr.String("PLAIN"),
				Username:  ptr.String("testuser111"),
				Password:  ptr.String("test-password-placeholder"),
			},
			expected: &resource.StreamsKafkaAuthentication{
				Mechanism: ptr.String("PLAIN"),
				Username:  ptr.String("testuser111"),
				Password:  nil, // Password is write-only, not returned from API
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := resource.NewModelAuthentication(tt.input, nil)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestNewModelSecurity(t *testing.T) {
	tests := []struct {
		input    *admin20250312010.StreamsKafkaSecurity
		expected *resource.StreamsKafkaSecurity
		name     string
	}{
		{
			name:     "Nil Input",
			input:    nil,
			expected: nil,
		},
		{
			name: "Valid Input",
			input: &admin20250312010.StreamsKafkaSecurity{
				BrokerPublicCertificate: ptr.String("testcert"),
				Protocol:                ptr.String("SSL"),
			},
			expected: &resource.StreamsKafkaSecurity{
				BrokerPublicCertificate: ptr.String("testcert"),
				Protocol:                ptr.String("SSL"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := resource.NewModelSecurity(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestNewDBRoleToExecute(t *testing.T) {
	tests := []struct {
		input    *resource.DBRoleToExecute
		expected *admin20250312010.DBRoleToExecute
		name     string
	}{
		{
			name:     "Nil Input",
			input:    nil,
			expected: nil,
		},
		{
			name: "Valid Input",
			input: &resource.DBRoleToExecute{
				Role: ptr.String("customroleadmin"),
				Type: ptr.String("CUSTOM"),
			},
			expected: &admin20250312010.DBRoleToExecute{
				Role: ptr.String("customroleadmin"),
				Type: ptr.String("CUSTOM"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := resource.NewDBRoleToExecute(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetStreamConnectionKafkaTypeModel(t *testing.T) {
	streamsConnKafka := &admin20250312010.StreamsConnection{
		Name:             ptr.String("TestConnection"),
		Type:             ptr.String("Kafka"),
		BootstrapServers: ptr.String("local.example.com:9192"),
		Authentication: &admin20250312010.StreamsKafkaAuthentication{
			Mechanism: ptr.String("PLAIN"),
			Username:  ptr.String("user1"),
			Password:  ptr.String("test-password-placeholder"),
		},
		Security: &admin20250312010.StreamsKafkaSecurity{
			BrokerPublicCertificate: ptr.String("cert1"),
			Protocol:                ptr.String("SSL"),
		},
		Config: &map[string]string{"retention.test": "60000"},
	}

	t.Run("With Nil Current Model", func(t *testing.T) {
		result := resource.GetStreamConnectionModel(streamsConnKafka, nil)

		assert.NotNil(t, result)
		assert.Equal(t, *streamsConnKafka.Name, *result.ConnectionName)
		assert.Equal(t, *streamsConnKafka.Type, *result.Type)
		assert.Equal(t, *streamsConnKafka.BootstrapServers, *result.BootstrapServers)
		assert.Equal(t, *streamsConnKafka.Authentication.Mechanism, *result.Authentication.Mechanism)
		assert.Equal(t, *streamsConnKafka.Security.Protocol, *result.Security.Protocol)
		assert.Equal(t, map[string]string{"retention.test": "60000"}, result.Config)
	})

	t.Run("With Non-Null Current Model", func(t *testing.T) {
		currentModel := &resource.Model{
			Profile:          ptr.String("default"),
			ProjectId:        ptr.String("testProjectID"),
			InstanceName:     ptr.String("TestInstance"),
			ConnectionName:   ptr.String("TestConnection"),
			Type:             ptr.String("Kafka"),
			BootstrapServers: ptr.String("local.example.com:9192"),
			Authentication: &resource.StreamsKafkaAuthentication{
				Mechanism: ptr.String("PLAIN"),
				Username:  ptr.String("user1"),
				Password:  ptr.String("test-password-placeholder"),
			},
			Security: &resource.StreamsKafkaSecurity{
				BrokerPublicCertificate: ptr.String("cert1"),
				Protocol:                ptr.String("SSL"),
			},
			Config: map[string]string{"retention.test": "60000"},
		}
		result := resource.GetStreamConnectionModel(streamsConnKafka, currentModel)

		assert.Equal(t, currentModel, result)
		assert.Equal(t, *streamsConnKafka.Name, *result.ConnectionName)
		assert.Equal(t, *currentModel.InstanceName, *result.InstanceName)
		assert.Equal(t, *currentModel.Profile, *result.Profile)
		assert.Equal(t, *currentModel.ProjectId, *result.ProjectId)
		assert.Equal(t, *streamsConnKafka.Type, *result.Type)
		assert.Equal(t, *streamsConnKafka.BootstrapServers, *result.BootstrapServers)
		assert.Equal(t, *streamsConnKafka.Authentication.Mechanism, *result.Authentication.Mechanism)
		assert.Equal(t, *streamsConnKafka.Security.Protocol, *result.Security.Protocol)
	})
}

func TestGetStreamConnectionClusterTypeModel(t *testing.T) {
	streamsConnKafka := &admin20250312010.StreamsConnection{
		Name:        ptr.String("TestConnection"),
		Type:        ptr.String("Cluster"),
		ClusterName: ptr.String("TestCluster"),
		DbRoleToExecute: &admin20250312010.DBRoleToExecute{
			Role: ptr.String("admin"),
			Type: ptr.String("Custom"),
		},
	}

	t.Run("With Nil Current Model", func(t *testing.T) {
		result := resource.GetStreamConnectionModel(streamsConnKafka, nil)

		assert.NotNil(t, result)
		assert.Equal(t, *streamsConnKafka.Name, *result.ConnectionName)
		assert.Equal(t, *streamsConnKafka.Type, *result.Type)
		assert.Equal(t, streamsConnKafka.DbRoleToExecute.GetRole(), *result.DbRoleToExecute.Role)
		assert.Equal(t, streamsConnKafka.DbRoleToExecute.GetType(), *result.DbRoleToExecute.Type)
	})

	t.Run("With Non-Null Current Model", func(t *testing.T) {
		currentModel := &resource.Model{
			Profile:        ptr.String("default"),
			ProjectId:      ptr.String("testProjectID"),
			InstanceName:   ptr.String("TestInstance"),
			ConnectionName: ptr.String("TestConnection"),
			Type:           ptr.String("Kafka"),
			ClusterName:    ptr.String("TestCluster"),
			DbRoleToExecute: &resource.DBRoleToExecute{
				Role: ptr.String("admin"),
				Type: ptr.String("Custom"),
			},
		}
		result := resource.GetStreamConnectionModel(streamsConnKafka, currentModel)

		assert.Equal(t, currentModel, result)
		assert.Equal(t, *streamsConnKafka.Name, *result.ConnectionName)
		assert.Equal(t, *currentModel.InstanceName, *result.InstanceName)
		assert.Equal(t, *currentModel.Profile, *result.Profile)
		assert.Equal(t, *currentModel.ProjectId, *result.ProjectId)
		assert.Equal(t, *streamsConnKafka.Type, *result.Type)
		assert.Equal(t, streamsConnKafka.DbRoleToExecute.GetRole(), *result.DbRoleToExecute.Role)
		assert.Equal(t, streamsConnKafka.DbRoleToExecute.GetType(), *result.DbRoleToExecute.Type)
	})
}

func TestGetStreamConnectionSampleTypeModel(t *testing.T) {
	streamsConnSample := &admin20250312010.StreamsConnection{
		Name: ptr.String("sample_stream_solar"),
		Type: ptr.String("Sample"),
	}
	testCases := []struct {
		model    *resource.Model
		asserter func(input, result *resource.Model, a *assert.Assertions)
		name     string
	}{
		{
			name:  "With Nil Current Model",
			model: nil,
			asserter: func(_, result *resource.Model, a *assert.Assertions) {
				a.NotNil(result)
				a.Equal(*streamsConnSample.Name, *result.ConnectionName)
				a.Equal(*streamsConnSample.Type, *result.Type)
				a.Nil(result.DbRoleToExecute)
			},
		},
		{
			name: "Sample Stream Solar dataset",
			model: &resource.Model{
				Profile:        ptr.String("default"),
				ProjectId:      ptr.String("testProjectID"),
				InstanceName:   ptr.String("TestInstance"),
				ConnectionName: ptr.String("sample_stream_solar"),
				Type:           ptr.String("Sample"),
			},
			asserter: func(input, result *resource.Model, a *assert.Assertions) {
				a.Equal(*input.InstanceName, *result.InstanceName)
				a.Equal(*input.Profile, *result.Profile)
				a.Equal(*input.ProjectId, *result.ProjectId)
				a.Equal(*input.ConnectionName, *result.ConnectionName)
				a.Equal(*input.Type, *result.Type)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := resource.GetStreamConnectionModel(streamsConnSample, tc.model)
			tc.asserter(tc.model, result, assert.New(t))
		})
	}
}
