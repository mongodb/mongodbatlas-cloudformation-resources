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

	"go.mongodb.org/atlas-sdk/v20231115007/admin"

	"github.com/aws/smithy-go/ptr"
	"github.com/stretchr/testify/assert"
)

func TestNewModelDBRoleToExecute(t *testing.T) {
	t.Run("Nil Input", func(t *testing.T) {
		assert.Nil(t, newModelDBRoleToExecute(nil))
	})

	t.Run("Valid Input", func(t *testing.T) {
		dbRole := &admin.DBRoleToExecute{
			Role: ptr.String("atlasAdmin"),
			Type: ptr.String("BUILT_IN"),
		}
		expected := &DBRoleToExecute{
			Role: ptr.String("atlasAdmin"),
			Type: ptr.String("BUILT_IN"),
		}
		assert.Equal(t, expected, newModelDBRoleToExecute(dbRole))
	})
}

func TestNewModelAuthentication(t *testing.T) {
	t.Run("Nil Input", func(t *testing.T) {
		assert.Nil(t, newModelAuthentication(nil))
	})

	t.Run("Valid Input", func(t *testing.T) {
		auth := &admin.StreamsKafkaAuthentication{
			Mechanism: ptr.String("PLAIN"),
			Username:  ptr.String("user1"),
			Password:  ptr.String("passwrd"),
		}
		expected := &StreamsKafkaAuthentication{
			Mechanism: ptr.String("PLAIN"),
			Username:  ptr.String("user1"),
			Password:  ptr.String("passwrd"),
		}
		assert.Equal(t, expected, newModelAuthentication(auth))
	})
}

func TestNewModelSecurity(t *testing.T) {
	t.Run("Nil Input", func(t *testing.T) {
		assert.Nil(t, newModelSecurity(nil))
	})

	t.Run("Valid Input", func(t *testing.T) {
		sec := &admin.StreamsKafkaSecurity{
			BrokerPublicCertificate: ptr.String("cert1"),
			Protocol:                ptr.String("SSL"),
		}
		expected := &StreamsKafkaSecurity{
			BrokerPublicCertificate: ptr.String("cert1"),
			Protocol:                ptr.String("SSL"),
		}
		assert.Equal(t, expected, newModelSecurity(sec))
	})
}

func TestNewDBRoleToExecute(t *testing.T) {
	t.Run("Nil Input", func(t *testing.T) {
		assert.Nil(t, newDBRoleToExecute(nil))
	})

	t.Run("Valid Input", func(t *testing.T) {
		dbRole := &DBRoleToExecute{
			Role: ptr.String("admin"),
			Type: ptr.String("CUSTOM"),
		}
		expected := &admin.DBRoleToExecute{
			Role: ptr.String("admin"),
			Type: ptr.String("CUSTOM"),
		}
		assert.Equal(t, expected, newDBRoleToExecute(dbRole))
	})
}

func TestGetStreamConnectionKafkaTypeModel(t *testing.T) {
	streamsConnKafka := &admin.StreamsConnection{
		Name:             ptr.String("TestConnection"),
		Type:             ptr.String("Kafka"),
		BootstrapServers: ptr.String("local.example.com:9192"),
		Authentication: &admin.StreamsKafkaAuthentication{
			Mechanism: ptr.String("PLAIN"),
			Username:  ptr.String("user1"),
			Password:  ptr.String("passwrd"),
		},
		Security: &admin.StreamsKafkaSecurity{
			BrokerPublicCertificate: ptr.String("cert1"),
			Protocol:                ptr.String("SSL"),
		},
		Config: mapPtr(map[string]string{"retention.test": "60000"}),
	}

	t.Run("With Nil Current Model", func(t *testing.T) {
		result := getStreamConnectionModel(streamsConnKafka, nil)

		assert.NotNil(t, result)
		assert.Equal(t, *streamsConnKafka.Name, *result.ConnectionName)
		assert.Equal(t, *streamsConnKafka.Type, *result.Type)
		assert.Equal(t, *streamsConnKafka.BootstrapServers, *result.BootstrapServers)
		assert.Equal(t, *streamsConnKafka.Authentication.Mechanism, *result.Authentication.Mechanism)
		assert.Equal(t, *streamsConnKafka.Security.Protocol, *result.Security.Protocol)
		assert.Equal(t, map[string]string{"retention.test": "60000"}, result.Config)
	})

	t.Run("With Non-Null Current Model", func(t *testing.T) {
		currentModel := &Model{
			Profile:          ptr.String("default"),
			ProjectId:        ptr.String("testProjectID"),
			InstanceName:     ptr.String("TestInstance"),
			ConnectionName:   ptr.String("TestConnection"),
			Type:             ptr.String("Kafka"),
			BootstrapServers: ptr.String("local.example.com:9192"),
			Authentication: &StreamsKafkaAuthentication{
				Mechanism: ptr.String("PLAIN"),
				Username:  ptr.String("user1"),
				Password:  ptr.String("passwrd"),
			},
			Security: &StreamsKafkaSecurity{
				BrokerPublicCertificate: ptr.String("cert1"),
				Protocol:                ptr.String("SSL"),
			},
			Config: map[string]string{"retention.test": "60000"},
		}
		result := getStreamConnectionModel(streamsConnKafka, currentModel)

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
	streamsConnKafka := &admin.StreamsConnection{
		Name:        ptr.String("TestConnection"),
		Type:        ptr.String("Cluster"),
		ClusterName: ptr.String("TestCluster"),
		DbRoleToExecute: &admin.DBRoleToExecute{
			Role: ptr.String("admin"),
			Type: ptr.String("Custom"),
		},
	}

	t.Run("With Nil Current Model", func(t *testing.T) {
		result := getStreamConnectionModel(streamsConnKafka, nil)

		assert.NotNil(t, result)
		assert.Equal(t, *streamsConnKafka.Name, *result.ConnectionName)
		assert.Equal(t, *streamsConnKafka.Type, *result.Type)
		assert.Equal(t, streamsConnKafka.DbRoleToExecute.GetRole(), *result.DbRoleToExecute.Role)
		assert.Equal(t, streamsConnKafka.DbRoleToExecute.GetType(), *result.DbRoleToExecute.Type)
	})

	t.Run("With Non-Null Current Model", func(t *testing.T) {
		currentModel := &Model{
			Profile:        ptr.String("default"),
			ProjectId:      ptr.String("testProjectID"),
			InstanceName:   ptr.String("TestInstance"),
			ConnectionName: ptr.String("TestConnection"),
			Type:           ptr.String("Kafka"),
			ClusterName:    ptr.String("TestCluster"),
			DbRoleToExecute: &DBRoleToExecute{
				Role: ptr.String("admin"),
				Type: ptr.String("Custom"),
			},
		}
		result := getStreamConnectionModel(streamsConnKafka, currentModel)

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

//nolint:gocritic
func mapPtr(m map[string]string) *map[string]string {
	return &m
}
