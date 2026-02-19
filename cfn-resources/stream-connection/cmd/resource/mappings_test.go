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

	"go.mongodb.org/atlas-sdk/v20250312014/admin"

	"github.com/aws/smithy-go/ptr"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-connection/cmd/resource"
	"github.com/stretchr/testify/assert"
)

const (
	testRoleValue     = "readWrite"
	testRoleTypeValue = "BUILT_IN"
	testUsername      = "testuser111"
	testMechanism     = "PLAIN"
	testProtocol      = "SSL"
	testCert          = "testcert"
	testCustomRole    = "customroleadmin"
	testCustomType    = "CUSTOM"
	testConnection    = "TestConnection"
	testCluster       = "TestCluster"
	testBootstrap     = "local.example.com:9192"
	testUser          = "user1"
	testSampleName    = "sample_stream_solar"
	testRoleArn       = "arn:aws:iam::123456789012:role/test-lambda-role"
	testURL           = "https://api.example.com/stream"
)

func TestMappings(t *testing.T) {
	testCases := map[string]struct {
		testFunc func(*testing.T)
	}{
		"NewModelDBRoleToExecute": {
			testFunc: func(t *testing.T) {
				t.Helper()
				input := &admin.DBRoleToExecute{Role: ptr.String(testRoleValue), Type: ptr.String(testRoleTypeValue)}
				result := resource.NewModelDBRoleToExecute(input)
				assert.Equal(t, testRoleValue, *result.Role)
				assert.Equal(t, testRoleTypeValue, *result.Type)
				assert.Nil(t, resource.NewModelDBRoleToExecute(nil))
			},
		},
		"NewModelAuthentication": {
			testFunc: func(t *testing.T) {
				t.Helper()
				input := &admin.StreamsKafkaAuthentication{
					Mechanism: ptr.String(testMechanism), Username: ptr.String(testUsername),
					Password: ptr.String("test-password-placeholder"),
				}
				result := resource.NewModelAuthentication(input, nil)
				assert.Equal(t, testMechanism, *result.Mechanism)
				assert.Equal(t, testUsername, *result.Username)
				assert.Nil(t, result.Password)
				assert.Nil(t, resource.NewModelAuthentication(nil, nil))
			},
		},
		"NewModelSecurity": {
			testFunc: func(t *testing.T) {
				t.Helper()
				input := &admin.StreamsKafkaSecurity{
					BrokerPublicCertificate: ptr.String(testCert), Protocol: ptr.String(testProtocol),
				}
				result := resource.NewModelSecurity(input)
				assert.Equal(t, testCert, *result.BrokerPublicCertificate)
				assert.Equal(t, testProtocol, *result.Protocol)
				assert.Nil(t, resource.NewModelSecurity(nil))
			},
		},
		"NewDBRoleToExecute": {
			testFunc: func(t *testing.T) {
				t.Helper()
				input := &resource.DBRoleToExecute{Role: ptr.String(testCustomRole), Type: ptr.String(testCustomType)}
				result := resource.NewDBRoleToExecute(input)
				assert.Equal(t, testCustomRole, *result.Role)
				assert.Equal(t, testCustomType, *result.Type)
				assert.Nil(t, resource.NewDBRoleToExecute(nil))
			},
		},
		"GetStreamConnectionModel_kafka": {
			testFunc: func(t *testing.T) {
				t.Helper()
				streamsConn := &admin.StreamsConnection{
					Name: ptr.String(testConnection), Type: ptr.String(resource.KafkaConnectionType),
					BootstrapServers: ptr.String(testBootstrap),
					Authentication: &admin.StreamsKafkaAuthentication{
						Mechanism: ptr.String(testMechanism), Username: ptr.String(testUser),
					},
					Security: &admin.StreamsKafkaSecurity{Protocol: ptr.String(testProtocol)},
				}
				result := resource.GetStreamConnectionModel(streamsConn, nil)
				assert.Equal(t, testConnection, *result.ConnectionName)
				assert.Equal(t, resource.KafkaConnectionType, *result.Type)
				assert.Equal(t, testBootstrap, *result.BootstrapServers)
			},
		},
		"GetStreamConnectionModel_cluster": {
			testFunc: func(t *testing.T) {
				t.Helper()
				streamsConn := &admin.StreamsConnection{
					Name: ptr.String(testConnection), Type: ptr.String(resource.ClusterConnectionType),
					ClusterName:     ptr.String(testCluster),
					DbRoleToExecute: &admin.DBRoleToExecute{Role: ptr.String("admin"), Type: ptr.String("Custom")},
				}
				result := resource.GetStreamConnectionModel(streamsConn, nil)
				assert.Equal(t, testConnection, *result.ConnectionName)
				assert.Equal(t, resource.ClusterConnectionType, *result.Type)
				assert.Equal(t, "admin", *result.DbRoleToExecute.Role)
			},
		},
		"GetStreamConnectionModel_sample": {
			testFunc: func(t *testing.T) {
				t.Helper()
				streamsConn := &admin.StreamsConnection{
					Name: ptr.String(testSampleName), Type: ptr.String("Sample"),
				}
				result := resource.GetStreamConnectionModel(streamsConn, nil)
				assert.Equal(t, testSampleName, *result.ConnectionName)
				assert.Equal(t, "Sample", *result.Type)
				assert.Nil(t, result.DbRoleToExecute)
			},
		},
		"GetStreamConnectionModel_awsLambda": {
			testFunc: func(t *testing.T) {
				t.Helper()
				streamsConn := &admin.StreamsConnection{
					Name: ptr.String(testConnection), Type: ptr.String(resource.AWSLambdaType),
					Aws: &admin.StreamsAWSConnectionConfig{
						RoleArn: ptr.String(testRoleArn),
					},
				}
				result := resource.GetStreamConnectionModel(streamsConn, nil)
				assert.Equal(t, testConnection, *result.ConnectionName)
				assert.Equal(t, resource.AWSLambdaType, *result.Type)
				assert.NotNil(t, result.Aws)
				assert.Equal(t, testRoleArn, *result.Aws.RoleArn)
			},
		},
		"GetStreamConnectionModel_https": {
			testFunc: func(t *testing.T) {
				t.Helper()
				testHeaders := map[string]string{
					"Authorization": "Bearer token123",
					"Content-Type":  "application/json",
				}
				streamsConn := &admin.StreamsConnection{
					Name: ptr.String(testConnection), Type: ptr.String(resource.HTTPSType),
					Url:     ptr.String(testURL),
					Headers: &testHeaders,
				}
				result := resource.GetStreamConnectionModel(streamsConn, nil)
				assert.Equal(t, testConnection, *result.ConnectionName)
				assert.Equal(t, resource.HTTPSType, *result.Type)
				assert.Equal(t, testURL, *result.Url)
				assert.NotNil(t, result.Headers)
				assert.Equal(t, "Bearer token123", result.Headers["Authorization"])
				assert.Equal(t, "application/json", result.Headers["Content-Type"])
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, tc.testFunc)
	}
}
