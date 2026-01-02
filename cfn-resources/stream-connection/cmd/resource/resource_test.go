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
	"fmt"
	"net/http"
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-connection/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"
	"go.mongodb.org/atlas-sdk/v20250312010/mockadmin"
)

// Test helper functions to create test models

func createTestClusterConnectionModel() *resource.Model {
	projectID := "507f1f77bcf86cd799439011"
	connectionName := "test-cluster-connection"
	workspaceName := "test-workspace"
	profile := "default"
	connType := "Cluster"
	clusterName := "test-cluster"
	role := "atlasAdmin"
	roleType := "BUILT_IN"

	return &resource.Model{
		Profile:        &profile,
		ProjectId:      &projectID,
		ConnectionName: &connectionName,
		WorkspaceName:  &workspaceName,
		Type:           &connType,
		ClusterName:    &clusterName,
		DbRoleToExecute: &resource.DBRoleToExecute{
			Role: &role,
			Type: &roleType,
		},
	}
}

func createTestKafkaConnectionModel() *resource.Model {
	projectID := "507f1f77bcf86cd799439011"
	connectionName := "test-kafka-connection"
	workspaceName := "test-workspace"
	profile := "default"
	connType := "Kafka"
	bootstrapServers := "broker1:9092,broker2:9092"
	protocol := "SSL"
	mechanism := "PLAIN"
	username := "test-user"
	password := "test-password"

	return &resource.Model{
		Profile:          &profile,
		ProjectId:        &projectID,
		ConnectionName:   &connectionName,
		WorkspaceName:    &workspaceName,
		Type:             &connType,
		BootstrapServers: &bootstrapServers,
		Security: &resource.StreamsKafkaSecurity{
			Protocol: &protocol,
		},
		Authentication: &resource.StreamsKafkaAuthentication{
			Mechanism: &mechanism,
			Username:  &username,
			Password:  &password,
		},
	}
}

func createTestAWSLambdaConnectionModel() *resource.Model {
	projectID := "507f1f77bcf86cd799439011"
	connectionName := "test-lambda-connection"
	workspaceName := "test-workspace"
	profile := "default"
	connType := "AWSLambda"
	roleArn := "arn:aws:iam::123456789012:role/test-role"

	return &resource.Model{
		Profile:        &profile,
		ProjectId:      &projectID,
		ConnectionName: &connectionName,
		WorkspaceName:  &workspaceName,
		Type:           &connType,
		Aws: &resource.Aws{
			RoleArn: &roleArn,
		},
	}
}

func createTestHTTPSConnectionModel() *resource.Model {
	projectID := "507f1f77bcf86cd799439011"
	connectionName := "test-https-connection"
	workspaceName := "test-workspace"
	profile := "default"
	connType := "Https"
	url := "https://example.com/webhook"
	headers := map[string]string{
		"Authorization": "Bearer token",
		"Content-Type":  "application/json",
	}

	return &resource.Model{
		Profile:        &profile,
		ProjectId:      &projectID,
		ConnectionName: &connectionName,
		WorkspaceName:  &workspaceName,
		Type:           &connType,
		Url:            &url,
		Headers:        headers,
	}
}

func createTestStreamConnectionResponse(connType string) *admin20250312010.StreamsConnection {
	name := "test-connection"
	response := &admin20250312010.StreamsConnection{
		Name: &name,
		Type: &connType,
	}

	switch connType {
	case "Cluster":
		clusterName := "test-cluster"
		response.ClusterName = &clusterName
		response.DbRoleToExecute = &admin20250312010.DBRoleToExecute{
			Role: admin20250312010.PtrString("atlasAdmin"),
			Type: admin20250312010.PtrString("BUILT_IN"),
		}
	case "Kafka":
		bootstrapServers := "broker1:9092,broker2:9092"
		response.BootstrapServers = &bootstrapServers
		response.Security = &admin20250312010.StreamsKafkaSecurity{
			Protocol: admin20250312010.PtrString("SSL"),
		}
		response.Authentication = &admin20250312010.StreamsKafkaAuthentication{
			Mechanism: admin20250312010.PtrString("PLAIN"),
			Username:  admin20250312010.PtrString("test-user"),
		}
	case "AWSLambda":
		response.Aws = &admin20250312010.StreamsAWSConnectionConfig{
			RoleArn: admin20250312010.PtrString("arn:aws:iam::123456789012:role/test-role"),
		}
	case "Https":
		url := "https://example.com/webhook"
		response.Url = &url
		headers := map[string]string{
			"Authorization": "Bearer token",
		}
		response.Headers = &headers
	}

	return response
}

// Test constants and required fields

func TestConstants(t *testing.T) {
	assert.Equal(t, "Cluster", resource.ClusterConnectionType)
	assert.Equal(t, "Kafka", resource.KafkaConnectionType)
	assert.Equal(t, "AWSLambda", resource.AWSLambdaType)
	assert.Equal(t, "Https", resource.HTTPSType)
}

func TestRequiredFields(t *testing.T) {
	assert.Equal(t, []string{constants.ProjectID, constants.ConnectionName, constants.Type}, resource.CreateRequiredFields)
	assert.Equal(t, []string{constants.ProjectID, constants.ConnectionName}, resource.ReadRequiredFields)
	assert.Equal(t, []string{constants.ProjectID, constants.ConnectionName, constants.Type}, resource.UpdateRequiredFields)
	assert.Equal(t, []string{constants.ProjectID, constants.ConnectionName}, resource.DeleteRequiredFields)
	assert.Equal(t, []string{constants.ProjectID}, resource.ListRequiredFields)
}

// Test validation errors for Create

func TestCreateValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel *resource.Model
		expectedMsg  string
	}{
		"missingProjectId": {
			&resource.Model{
				Profile:        util.StringPtr("default"),
				ConnectionName: util.StringPtr("test-connection"),
				WorkspaceName:  util.StringPtr("test-workspace"),
				Type:           util.StringPtr("Cluster"),
			},
			"required",
		},
		"missingConnectionName": {
			&resource.Model{
				Profile:       util.StringPtr("default"),
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("test-workspace"),
				Type:          util.StringPtr("Cluster"),
			},
			"required",
		},
		"missingType": {
			&resource.Model{
				Profile:        util.StringPtr("default"),
				ProjectId:      util.StringPtr("507f1f77bcf86cd799439011"),
				ConnectionName: util.StringPtr("test-connection"),
				WorkspaceName:  util.StringPtr("test-workspace"),
			},
			"required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.Create(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestCreateMissingWorkspaceValidation(t *testing.T) {
	// This test validates the workspace/instance name check that happens AFTER InitEnvWithLatestClient
	mockStreamsAPI := mockadmin.NewStreamsApi(t)
	mockClient := &admin20250312010.APIClient{StreamsApi: mockStreamsAPI}

	originalInitEnv := resource.InitEnvWithLatestClient
	defer func() { resource.InitEnvWithLatestClient = originalInitEnv }()

	resource.InitEnvWithLatestClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
		return mockClient, nil
	}

	model := &resource.Model{
		Profile:        util.StringPtr("default"),
		ProjectId:      util.StringPtr("507f1f77bcf86cd799439011"),
		ConnectionName: util.StringPtr("test-connection"),
		Type:           util.StringPtr("Cluster"),
		// Note: No WorkspaceName or InstanceName
	}

	event, err := resource.Create(handler.Request{}, nil, model)
	require.NoError(t, err)
	assert.Equal(t, handler.Failed, event.OperationStatus)
	assert.Contains(t, event.Message, "Either WorkspaceName or InstanceName must be provided")
}

// Test validation errors for Read

func TestReadValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel *resource.Model
		expectedMsg  string
	}{
		"missingProjectId": {
			&resource.Model{
				Profile:        util.StringPtr("default"),
				ConnectionName: util.StringPtr("test-connection"),
				WorkspaceName:  util.StringPtr("test-workspace"),
			},
			"required",
		},
		"missingConnectionName": {
			&resource.Model{
				Profile:       util.StringPtr("default"),
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("test-workspace"),
			},
			"required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.Read(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestReadMissingWorkspaceValidation(t *testing.T) {
	mockStreamsAPI := mockadmin.NewStreamsApi(t)
	mockClient := &admin20250312010.APIClient{StreamsApi: mockStreamsAPI}

	originalInitEnv := resource.InitEnvWithLatestClient
	defer func() { resource.InitEnvWithLatestClient = originalInitEnv }()

	resource.InitEnvWithLatestClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
		return mockClient, nil
	}

	model := &resource.Model{
		Profile:        util.StringPtr("default"),
		ProjectId:      util.StringPtr("507f1f77bcf86cd799439011"),
		ConnectionName: util.StringPtr("test-connection"),
	}

	event, err := resource.Read(handler.Request{}, nil, model)
	require.NoError(t, err)
	assert.Equal(t, handler.Failed, event.OperationStatus)
	assert.Contains(t, event.Message, "Either WorkspaceName or InstanceName must be provided")
}

// Test validation errors for Update

func TestUpdateValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel *resource.Model
		expectedMsg  string
	}{
		"missingProjectId": {
			&resource.Model{
				Profile:        util.StringPtr("default"),
				ConnectionName: util.StringPtr("test-connection"),
				WorkspaceName:  util.StringPtr("test-workspace"),
				Type:           util.StringPtr("Cluster"),
			},
			"required",
		},
		"missingConnectionName": {
			&resource.Model{
				Profile:       util.StringPtr("default"),
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("test-workspace"),
				Type:          util.StringPtr("Cluster"),
			},
			"required",
		},
		"missingType": {
			&resource.Model{
				Profile:        util.StringPtr("default"),
				ProjectId:      util.StringPtr("507f1f77bcf86cd799439011"),
				ConnectionName: util.StringPtr("test-connection"),
				WorkspaceName:  util.StringPtr("test-workspace"),
			},
			"required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.Update(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

// Test validation errors for Delete

func TestDeleteValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel *resource.Model
		expectedMsg  string
	}{
		"missingProjectId": {
			&resource.Model{
				Profile:        util.StringPtr("default"),
				ConnectionName: util.StringPtr("test-connection"),
				WorkspaceName:  util.StringPtr("test-workspace"),
			},
			"required",
		},
		"missingConnectionName": {
			&resource.Model{
				Profile:       util.StringPtr("default"),
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("test-workspace"),
			},
			"required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.Delete(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

// Test validation errors for List

func TestListValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel *resource.Model
		expectedMsg  string
	}{
		"missingProjectId": {
			&resource.Model{
				Profile:       util.StringPtr("default"),
				WorkspaceName: util.StringPtr("test-workspace"),
			},
			"required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.List(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestListMissingWorkspaceValidation(t *testing.T) {
	mockStreamsAPI := mockadmin.NewStreamsApi(t)
	mockClient := &admin20250312010.APIClient{StreamsApi: mockStreamsAPI}

	originalInitEnv := resource.InitEnvWithLatestClient
	defer func() { resource.InitEnvWithLatestClient = originalInitEnv }()

	resource.InitEnvWithLatestClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
		return mockClient, nil
	}

	model := &resource.Model{
		Profile:   util.StringPtr("default"),
		ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
	}

	event, err := resource.List(handler.Request{}, nil, model)
	require.NoError(t, err)
	assert.Equal(t, handler.Failed, event.OperationStatus)
	assert.Contains(t, event.Message, "Either WorkspaceName or InstanceName must be provided")
}

// Test Create with mocks

func TestCreateWithMocks(t *testing.T) {
	testCases := map[string]struct {
		model          *resource.Model
		mockSetup      func(*mockadmin.StreamsApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
	}{
		"successfulCreateClusterConnection": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				response := createTestStreamConnectionResponse("Cluster")
				m.EXPECT().CreateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateStreamConnectionApiRequest{ApiService: m})
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
		"successfulCreateKafkaConnection": {
			model: createTestKafkaConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				response := createTestStreamConnectionResponse("Kafka")
				m.EXPECT().CreateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().CreateStreamConnectionExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 201}, nil)
			},
			expectedStatus: handler.Success,
		},
		"successfulCreateAWSLambdaConnection": {
			model: createTestAWSLambdaConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				response := createTestStreamConnectionResponse("AWSLambda")
				m.EXPECT().CreateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().CreateStreamConnectionExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 201}, nil)
			},
			expectedStatus: handler.Success,
		},
		"successfulCreateHTTPSConnection": {
			model: createTestHTTPSConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				response := createTestStreamConnectionResponse("Https")
				m.EXPECT().CreateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().CreateStreamConnectionExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 201}, nil)
			},
			expectedStatus: handler.Success,
		},
		"createWithAPIError": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().CreateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().CreateStreamConnectionExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
		"createWithConflictError": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().CreateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().CreateStreamConnectionExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: http.StatusConflict}, fmt.Errorf("already exists"))
			},
			expectedStatus: handler.Failed,
		},
		"createWithInstanceNameInsteadOfWorkspaceName": {
			model: func() *resource.Model {
				model := createTestClusterConnectionModel()
				model.InstanceName = model.WorkspaceName
				model.WorkspaceName = nil
				return model
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				response := createTestStreamConnectionResponse("Cluster")
				m.EXPECT().CreateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().CreateStreamConnectionExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 201}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				model := event.ResourceModel.(*resource.Model)
				// WorkspaceName should be normalized from InstanceName
				assert.NotNil(t, model.WorkspaceName)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsAPI := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsAPI)

			mockClient := &admin20250312010.APIClient{StreamsApi: mockStreamsAPI}

			// Save original function and restore after test
			originalInitEnv := resource.InitEnvWithLatestClient
			defer func() { resource.InitEnvWithLatestClient = originalInitEnv }()

			// Mock the initialization function
			resource.InitEnvWithLatestClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := resource.Create(handler.Request{}, nil, tc.model)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.validateResult != nil {
				tc.validateResult(t, event)
			}
		})
	}
}

// Test Read with mocks

func TestReadWithMocks(t *testing.T) {
	testCases := map[string]struct {
		model          *resource.Model
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
	}{
		"successfulReadClusterConnection": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				response := createTestStreamConnectionResponse("Cluster")
				m.EXPECT().GetStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().GetStreamConnectionExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"successfulReadKafkaConnection": {
			model: createTestKafkaConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				response := createTestStreamConnectionResponse("Kafka")
				m.EXPECT().GetStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().GetStreamConnectionExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"readNotFound": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().GetStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().GetStreamConnectionExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
		"readWithAPIError": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().GetStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().GetStreamConnectionExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsAPI := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsAPI)

			mockClient := &admin20250312010.APIClient{StreamsApi: mockStreamsAPI}

			originalInitEnv := resource.InitEnvWithLatestClient
			defer func() { resource.InitEnvWithLatestClient = originalInitEnv }()

			resource.InitEnvWithLatestClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := resource.Read(handler.Request{}, nil, tc.model)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

// Test Update with mocks

func TestUpdateWithMocks(t *testing.T) {
	testCases := map[string]struct {
		model          *resource.Model
		mockSetup      func(*mockadmin.StreamsApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
	}{
		"successfulUpdateClusterConnection": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				response := createTestStreamConnectionResponse("Cluster")
				m.EXPECT().UpdateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.UpdateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().UpdateStreamConnectionExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, "Update Completed", event.Message)
			},
		},
		"successfulUpdateKafkaConnection": {
			model: createTestKafkaConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				response := createTestStreamConnectionResponse("Kafka")
				m.EXPECT().UpdateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.UpdateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().UpdateStreamConnectionExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"updateNotFound": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().UpdateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.UpdateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().UpdateStreamConnectionExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
		"updateWithAPIError": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().UpdateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.UpdateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().UpdateStreamConnectionExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
		"updateWithBadRequest": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().UpdateStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.UpdateStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().UpdateStreamConnectionExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 400}, fmt.Errorf("bad request"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsAPI := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsAPI)

			mockClient := &admin20250312010.APIClient{StreamsApi: mockStreamsAPI}

			originalInitEnv := resource.InitEnvWithLatestClient
			defer func() { resource.InitEnvWithLatestClient = originalInitEnv }()

			resource.InitEnvWithLatestClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := resource.Update(handler.Request{}, nil, tc.model)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.validateResult != nil {
				tc.validateResult(t, event)
			}
		})
	}
}

// Test Delete with mocks

func TestDeleteWithMocks(t *testing.T) {
	testCases := map[string]struct {
		model          *resource.Model
		mockSetup      func(*mockadmin.StreamsApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
	}{
		"successfulDeleteClusterConnection": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().DeleteStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.DeleteStreamConnectionApiRequest{ApiService: m})
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
		"successfulDeleteKafkaConnection": {
			model: createTestKafkaConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().DeleteStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.DeleteStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().DeleteStreamConnectionExecute(mock.Anything).
					Return(&http.Response{StatusCode: 204}, nil)
			},
			expectedStatus: handler.Success,
		},
		"deleteNotFound": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().DeleteStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.DeleteStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().DeleteStreamConnectionExecute(mock.Anything).
					Return(&http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, "NotFound", event.HandlerErrorCode)
			},
		},
		"deleteWithAPIError": {
			model: createTestClusterConnectionModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().DeleteStreamConnection(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.DeleteStreamConnectionApiRequest{ApiService: m})
				m.EXPECT().DeleteStreamConnectionExecute(mock.Anything).
					Return(&http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsAPI := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsAPI)

			mockClient := &admin20250312010.APIClient{StreamsApi: mockStreamsAPI}

			originalInitEnv := resource.InitEnvWithLatestClient
			defer func() { resource.InitEnvWithLatestClient = originalInitEnv }()

			resource.InitEnvWithLatestClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := resource.Delete(handler.Request{}, nil, tc.model)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.validateResult != nil {
				tc.validateResult(t, event)
			}
		})
	}
}

// Test List with mocks

func TestListWithMocks(t *testing.T) {
	testCases := map[string]struct {
		model          *resource.Model
		mockSetup      func(*mockadmin.StreamsApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
	}{
		"successfulListWithMultipleConnections": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("test-workspace"),
				Profile:       util.StringPtr("default"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				connections := []admin20250312010.StreamsConnection{
					*createTestStreamConnectionResponse("Cluster"),
					*createTestStreamConnectionResponse("Kafka"),
				}
				totalCount := len(connections)
				response := &admin20250312010.PaginatedApiStreamsConnection{
					Results:    &connections,
					TotalCount: &totalCount,
				}
				m.EXPECT().ListStreamConnectionsWithParams(mock.Anything, mock.Anything).
					Return(admin20250312010.ListStreamConnectionsApiRequest{ApiService: m})
				m.EXPECT().ListStreamConnectionsExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Len(t, event.ResourceModels, 2)
			},
		},
		"successfulListWithEmptyResult": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("test-workspace"),
				Profile:       util.StringPtr("default"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				connections := []admin20250312010.StreamsConnection{}
				totalCount := 0
				response := &admin20250312010.PaginatedApiStreamsConnection{
					Results:    &connections,
					TotalCount: &totalCount,
				}
				m.EXPECT().ListStreamConnectionsWithParams(mock.Anything, mock.Anything).
					Return(admin20250312010.ListStreamConnectionsApiRequest{ApiService: m})
				m.EXPECT().ListStreamConnectionsExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Empty(t, event.ResourceModels)
			},
		},
		"successfulListWithPagination": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("test-workspace"),
				Profile:       util.StringPtr("default"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				// First page
				connections1 := []admin20250312010.StreamsConnection{
					*createTestStreamConnectionResponse("Cluster"),
				}
				totalCount1 := 2
				response1 := &admin20250312010.PaginatedApiStreamsConnection{
					Results:    &connections1,
					TotalCount: &totalCount1,
				}

				// Second page
				connections2 := []admin20250312010.StreamsConnection{
					*createTestStreamConnectionResponse("Kafka"),
				}
				totalCount2 := 2
				response2 := &admin20250312010.PaginatedApiStreamsConnection{
					Results:    &connections2,
					TotalCount: &totalCount2,
				}

				m.EXPECT().ListStreamConnectionsWithParams(mock.Anything, mock.Anything).
					Return(admin20250312010.ListStreamConnectionsApiRequest{ApiService: m}).Times(2)
				m.EXPECT().ListStreamConnectionsExecute(mock.Anything).
					Return(response1, &http.Response{StatusCode: 200}, nil).Once()
				m.EXPECT().ListStreamConnectionsExecute(mock.Anything).
					Return(response2, &http.Response{StatusCode: 200}, nil).Once()
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Len(t, event.ResourceModels, 2)
			},
		},
		"listWithAPIError": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("test-workspace"),
				Profile:       util.StringPtr("default"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().ListStreamConnectionsWithParams(mock.Anything, mock.Anything).
					Return(admin20250312010.ListStreamConnectionsApiRequest{ApiService: m})
				m.EXPECT().ListStreamConnectionsExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
		"listWithInstanceNameInsteadOfWorkspaceName": {
			model: &resource.Model{
				ProjectId:    util.StringPtr("507f1f77bcf86cd799439011"),
				InstanceName: util.StringPtr("test-workspace"),
				Profile:      util.StringPtr("default"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				connections := []admin20250312010.StreamsConnection{
					*createTestStreamConnectionResponse("Cluster"),
				}
				totalCount := 1
				response := &admin20250312010.PaginatedApiStreamsConnection{
					Results:    &connections,
					TotalCount: &totalCount,
				}
				m.EXPECT().ListStreamConnectionsWithParams(mock.Anything, mock.Anything).
					Return(admin20250312010.ListStreamConnectionsApiRequest{ApiService: m})
				m.EXPECT().ListStreamConnectionsExecute(mock.Anything).
					Return(response, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Len(t, event.ResourceModels, 1)
				// Verify InstanceName is preserved in response
				model := event.ResourceModels[0].(*resource.Model)
				assert.NotNil(t, model.InstanceName)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsAPI := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsAPI)

			mockClient := &admin20250312010.APIClient{StreamsApi: mockStreamsAPI}

			originalInitEnv := resource.InitEnvWithLatestClient
			defer func() { resource.InitEnvWithLatestClient = originalInitEnv }()

			resource.InitEnvWithLatestClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := resource.List(handler.Request{}, nil, tc.model)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.validateResult != nil {
				tc.validateResult(t, event)
			}
		})
	}
}
