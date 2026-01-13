// Copyright 2026 MongoDB Inc
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
	"fmt"
	"net/http"
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-connection/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312012/admin"
	"go.mongodb.org/atlas-sdk/v20250312012/mockadmin"
)

const (
	testProjectID        = "507f1f77bcf86cd799439011"
	testConnectionName   = "test-connection"
	testWorkspaceName    = "test-workspace"
	testProfile          = "default"
	testClusterName      = "test-cluster"
	testRole             = "atlasAdmin"
	testRoleType         = "BUILT_IN"
	msgRequired          = "required"
	msgWorkspaceRequired = "Either WorkspaceName or InstanceName must be provided"
)

func createTestClusterConnectionModel() *resource.Model {
	return &resource.Model{
		Profile:        util.StringPtr(testProfile),
		ProjectId:      util.StringPtr(testProjectID),
		ConnectionName: util.StringPtr(testConnectionName),
		WorkspaceName:  util.StringPtr(testWorkspaceName),
		Type:           util.StringPtr(resource.ClusterConnectionType),
		ClusterName:    util.StringPtr(testClusterName),
		DbRoleToExecute: &resource.DBRoleToExecute{
			Role: util.StringPtr(testRole),
			Type: util.StringPtr(testRoleType),
		},
	}
}

func createTestStreamConnectionResponse(connType string) *admin.StreamsConnection {
	name := testConnectionName
	response := &admin.StreamsConnection{
		Name: &name,
		Type: &connType,
	}
	if connType == resource.ClusterConnectionType {
		response.ClusterName = util.StringPtr(testClusterName)
		response.DbRoleToExecute = &admin.DBRoleToExecute{
			Role: admin.PtrString(testRole),
			Type: admin.PtrString(testRoleType),
		}
	}
	return response
}

func TestConstants(t *testing.T) {
	assert.Equal(t, "Cluster", resource.ClusterConnectionType)
	assert.Equal(t, "Kafka", resource.KafkaConnectionType)
	assert.Equal(t, []string{constants.ProjectID, constants.ConnectionName, constants.Type}, resource.CreateRequiredFields)
	assert.Equal(t, []string{constants.ProjectID, constants.ConnectionName}, resource.ReadRequiredFields)
	assert.Equal(t, []string{constants.ProjectID, constants.ConnectionName, constants.Type}, resource.UpdateRequiredFields)
	assert.Equal(t, []string{constants.ProjectID, constants.ConnectionName}, resource.DeleteRequiredFields)
	assert.Equal(t, []string{constants.ProjectID}, resource.ListRequiredFields)
}

func TestValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		operation    func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error)
		currentModel *resource.Model
		expectedMsg  string
	}{
		"Create_missingProjectId": {
			operation: resource.Create,
			currentModel: &resource.Model{
				Profile: util.StringPtr(testProfile), ConnectionName: util.StringPtr(testConnectionName),
				WorkspaceName: util.StringPtr(testWorkspaceName), Type: util.StringPtr(resource.ClusterConnectionType),
			},
			expectedMsg: msgRequired,
		},
		"Create_missingConnectionName": {
			operation: resource.Create,
			currentModel: &resource.Model{
				Profile: util.StringPtr(testProfile), ProjectId: util.StringPtr(testProjectID),
				WorkspaceName: util.StringPtr(testWorkspaceName), Type: util.StringPtr(resource.ClusterConnectionType),
			},
			expectedMsg: msgRequired,
		},
		"Create_missingWorkspaceOrInstanceName": {
			operation: resource.Create,
			currentModel: &resource.Model{
				Profile: util.StringPtr(testProfile), ProjectId: util.StringPtr(testProjectID),
				ConnectionName: util.StringPtr(testConnectionName), Type: util.StringPtr(resource.ClusterConnectionType),
			},
			expectedMsg: msgWorkspaceRequired,
		},
		"Read_missingProjectId": {
			operation:    resource.Read,
			currentModel: &resource.Model{ConnectionName: util.StringPtr(testConnectionName), WorkspaceName: util.StringPtr(testWorkspaceName)},
			expectedMsg:  msgRequired,
		},
		"Read_missingWorkspaceOrInstanceName": {
			operation:    resource.Read,
			currentModel: &resource.Model{ProjectId: util.StringPtr(testProjectID), ConnectionName: util.StringPtr(testConnectionName)},
			expectedMsg:  msgWorkspaceRequired,
		},
		"Update_missingProjectId": {
			operation: resource.Update,
			currentModel: &resource.Model{
				Profile: util.StringPtr(testProfile), ConnectionName: util.StringPtr(testConnectionName),
				WorkspaceName: util.StringPtr(testWorkspaceName), Type: util.StringPtr(resource.ClusterConnectionType),
			},
			expectedMsg: msgRequired,
		},
		"Delete_missingProjectId": {
			operation:    resource.Delete,
			currentModel: &resource.Model{ConnectionName: util.StringPtr(testConnectionName), WorkspaceName: util.StringPtr(testWorkspaceName)},
			expectedMsg:  msgRequired,
		},
		"List_missingProjectId": {
			operation:    resource.List,
			currentModel: &resource.Model{WorkspaceName: util.StringPtr(testWorkspaceName)},
			expectedMsg:  msgRequired,
		},
		"List_missingWorkspaceOrInstanceName": {
			operation:    resource.List,
			currentModel: &resource.Model{ProjectId: util.StringPtr(testProjectID)},
			expectedMsg:  msgWorkspaceRequired,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			needsMock := name == "Create_missingWorkspaceOrInstanceName" ||
				name == "Read_missingWorkspaceOrInstanceName" ||
				name == "List_missingWorkspaceOrInstanceName"
			if needsMock {
				cleanup := setupMockClient(t, func(*mockadmin.StreamsApi) {})
				defer cleanup()
			}
			event, err := tc.operation(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func setupMockClient(t *testing.T, mockSetup func(*mockadmin.StreamsApi)) func() {
	t.Helper()
	originalInitEnv := resource.InitEnvWithLatestClient
	mockStreamsAPI := mockadmin.NewStreamsApi(t)
	mockSetup(mockStreamsAPI)
	mockClient := &admin.APIClient{StreamsApi: mockStreamsAPI}
	resource.InitEnvWithLatestClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin.APIClient, *handler.ProgressEvent) {
		return mockClient, nil
	}
	return func() { resource.InitEnvWithLatestClient = originalInitEnv }
}

func TestCRUDOperations(t *testing.T) {
	testCases := map[string]struct {
		operation      func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error)
		model          *resource.Model
		mockSetup      func(*mockadmin.StreamsApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
	}{
		"Create_success": {
			operation: resource.Create,
			model:     createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				response := createTestStreamConnectionResponse(resource.ClusterConnectionType)
				m.EXPECT().CreateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin.CreateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().CreateStreamConnectionExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 201}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, "Create Completed", event.Message)
				assert.NotNil(t, event.ResourceModel)
			},
		},
		"Create_kafka": {
			operation: resource.Create,
			model: &resource.Model{
				Profile: util.StringPtr(testProfile), ProjectId: util.StringPtr(testProjectID),
				ConnectionName: util.StringPtr(testConnectionName), WorkspaceName: util.StringPtr(testWorkspaceName),
				Type: util.StringPtr(resource.KafkaConnectionType), BootstrapServers: util.StringPtr("broker1:9092"),
				Security:       &resource.StreamsKafkaSecurity{Protocol: util.StringPtr("SSL")},
				Authentication: &resource.StreamsKafkaAuthentication{Mechanism: util.StringPtr("PLAIN")},
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				bootstrap := "broker1:9092"
				response := &admin.StreamsConnection{
					Name: util.StringPtr(testConnectionName), Type: util.StringPtr(resource.KafkaConnectionType),
					BootstrapServers: &bootstrap,
				}
				m.EXPECT().CreateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin.CreateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().CreateStreamConnectionExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 201}, nil)
			},
			expectedStatus: handler.Success,
		},
		"Create_apiError": {
			operation: resource.Create,
			model:     createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().CreateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin.CreateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().CreateStreamConnectionExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
		"Read_success": {
			operation: resource.Read,
			model:     createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				response := createTestStreamConnectionResponse(resource.ClusterConnectionType)
				m.EXPECT().GetStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().GetStreamConnectionExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"Read_notFound": {
			operation: resource.Read,
			model:     createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().GetStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().GetStreamConnectionExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
		"Update_success": {
			operation: resource.Update,
			model:     createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				response := createTestStreamConnectionResponse(resource.ClusterConnectionType)
				m.EXPECT().UpdateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin.UpdateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().UpdateStreamConnectionExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"Update_apiError": {
			operation: resource.Update,
			model:     createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().UpdateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin.UpdateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().UpdateStreamConnectionExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
		"Delete_success": {
			operation: resource.Delete,
			model:     createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().DeleteStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin.DeleteStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().DeleteStreamConnectionExecute(mock.Anything).
					Return(&http.Response{StatusCode: 204}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, "Delete Completed", event.Message)
				assert.Nil(t, event.ResourceModel)
			},
		},
		"Delete_notFound": {
			operation: resource.Delete,
			model:     createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().DeleteStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin.DeleteStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().DeleteStreamConnectionExecute(mock.Anything).
					Return(&http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
			},
		},
		"List_success": {
			operation: resource.List,
			model: &resource.Model{
				ProjectId: util.StringPtr(testProjectID), WorkspaceName: util.StringPtr(testWorkspaceName),
				Profile: util.StringPtr(testProfile),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				connections := []admin.StreamsConnection{*createTestStreamConnectionResponse(resource.ClusterConnectionType)}
				totalCount := 1
				response := &admin.PaginatedApiStreamsConnection{
					Results: &connections, TotalCount: &totalCount,
				}
				m.EXPECT().ListStreamConnectionsWithParams(mock.Anything, mock.Anything).
					Return(admin.ListStreamConnectionsApiRequest{ApiService: m})
				m.EXPECT().ListStreamConnectionsExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Len(t, event.ResourceModels, 1)
			},
		},
		"List_withPagination": {
			operation: resource.List,
			model: &resource.Model{
				ProjectId: util.StringPtr(testProjectID), WorkspaceName: util.StringPtr(testWorkspaceName),
				Profile: util.StringPtr(testProfile),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				conn1 := []admin.StreamsConnection{*createTestStreamConnectionResponse(resource.ClusterConnectionType)}
				conn2 := []admin.StreamsConnection{*createTestStreamConnectionResponse(resource.ClusterConnectionType)}
				totalCount := 2
				m.EXPECT().ListStreamConnectionsWithParams(mock.Anything, mock.Anything).
					Return(admin.ListStreamConnectionsApiRequest{ApiService: m}).Times(2)
				m.EXPECT().ListStreamConnectionsExecute(mock.Anything).
					Return(&admin.PaginatedApiStreamsConnection{Results: &conn1, TotalCount: &totalCount},
						&http.Response{StatusCode: 200}, nil).Once()
				m.EXPECT().ListStreamConnectionsExecute(mock.Anything).
					Return(&admin.PaginatedApiStreamsConnection{Results: &conn2, TotalCount: &totalCount},
						&http.Response{StatusCode: 200}, nil).Once()
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.GreaterOrEqual(t, len(event.ResourceModels), 2)
			},
		},
		"List_apiError": {
			operation: resource.List,
			model: &resource.Model{
				ProjectId: util.StringPtr(testProjectID), WorkspaceName: util.StringPtr(testWorkspaceName),
				Profile: util.StringPtr(testProfile),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().ListStreamConnectionsWithParams(mock.Anything, mock.Anything).
					Return(admin.ListStreamConnectionsApiRequest{ApiService: m})
				m.EXPECT().ListStreamConnectionsExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			cleanup := setupMockClient(t, tc.mockSetup)
			defer cleanup()
			event, err := tc.operation(handler.Request{}, nil, tc.model)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.validateResult != nil {
				tc.validateResult(t, event)
			}
		})
	}
}
