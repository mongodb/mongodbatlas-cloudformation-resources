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
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"
	"go.mongodb.org/atlas-sdk/v20250312010/mockadmin"
)

// Helper function to create a test model
func createTestModel() *Model {
	projectID := "507f1f77bcf86cd799439011"
	providerName := "AWS"
	vendor := "MSK"
	arn := "arn:aws:kafka:us-east-1:123456789012:cluster/test-cluster/12345678-1234-1234-1234-123456789012-1"

	return &Model{
		Profile:      util.StringPtr("default"),
		ProjectId:    &projectID,
		ProviderName: &providerName,
		Vendor:       &vendor,
		Arn:          &arn,
	}
}

func createTestModelConfluent() *Model {
	projectID := "507f1f77bcf86cd799439011"
	providerName := "AWS"
	vendor := "CONFLUENT"
	region := "us-east-1"
	serviceEndpointId := "com.amazonaws.vpce.us-east-1.vpce-svc-12345678"
	dnsDomain := "test.example.com"
	dnsSubDomain := []string{"az1", "az2"}

	return &Model{
		Profile:           util.StringPtr("default"),
		ProjectId:         &projectID,
		ProviderName:      &providerName,
		Vendor:            &vendor,
		Region:            &region,
		ServiceEndpointId: &serviceEndpointId,
		DnsDomain:         &dnsDomain,
		DnsSubDomain:      dnsSubDomain,
	}
}

func createTestModelS3() *Model {
	projectID := "507f1f77bcf86cd799439011"
	providerName := "AWS"
	vendor := "S3"
	region := "us-east-1"
	serviceEndpointId := "com.amazonaws.us-east-1.s3"

	return &Model{
		Profile:           util.StringPtr("default"),
		ProjectId:         &projectID,
		ProviderName:      &providerName,
		Vendor:            &vendor,
		Region:            &region,
		ServiceEndpointId: &serviceEndpointId,
	}
}

// Helper function to create a test API response
func createTestAPIResponse() *admin20250312010.StreamsPrivateLinkConnection {
	connectionID := "507f1f77bcf86cd799439012"
	providerName := "AWS"
	vendor := "MSK"
	state := "DONE"
	arn := "arn:aws:kafka:us-east-1:123456789012:cluster/test-cluster/12345678-1234-1234-1234-123456789012-1"

	return &admin20250312010.StreamsPrivateLinkConnection{
		Id:       &connectionID,
		Provider: providerName,
		Vendor:   &vendor,
		State:    &state,
		Arn:      &arn,
	}
}

// Test validation errors
func TestCreateValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *Model
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"missingProjectId": {
			currentModel: &Model{
				ProviderName: util.StringPtr("AWS"),
				Vendor:       util.StringPtr("MSK"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"missingProviderName": {
			currentModel: &Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
				Vendor:    util.StringPtr("MSK"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"missingVendor": {
			currentModel: &Model{
				ProjectId:    util.StringPtr("507f1f77bcf86cd799439011"),
				ProviderName: util.StringPtr("AWS"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"missingArnForMSK": {
			currentModel: &Model{
				ProjectId:    util.StringPtr("507f1f77bcf86cd799439011"),
				ProviderName: util.StringPtr("AWS"),
				Vendor:       util.StringPtr("MSK"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "Arn is required",
		},
		"regionSetForMSK": {
			currentModel: &Model{
				ProjectId:    util.StringPtr("507f1f77bcf86cd799439011"),
				ProviderName: util.StringPtr("AWS"),
				Vendor:       util.StringPtr("MSK"),
				Arn:          util.StringPtr("arn:aws:kafka:us-east-1:123456789012:cluster/test"),
				Region:       util.StringPtr("us-east-1"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "Region cannot be set for vendor MSK",
		},
		"missingDnsDomainForConfluent": {
			currentModel: &Model{
				ProjectId:         util.StringPtr("507f1f77bcf86cd799439011"),
				ProviderName:      util.StringPtr("AWS"),
				Vendor:            util.StringPtr("CONFLUENT"),
				Region:            util.StringPtr("us-east-1"),
				ServiceEndpointId: util.StringPtr("com.amazonaws.vpce.us-east-1.vpce-svc-12345678"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "DnsDomain is required",
		},
		"missingRegionForConfluent": {
			currentModel: &Model{
				ProjectId:         util.StringPtr("507f1f77bcf86cd799439011"),
				ProviderName:      util.StringPtr("AWS"),
				Vendor:            util.StringPtr("CONFLUENT"),
				DnsDomain:         util.StringPtr("test.example.com"),
				ServiceEndpointId: util.StringPtr("com.amazonaws.vpce.us-east-1.vpce-svc-12345678"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "Region is required",
		},
		"missingServiceEndpointIdForConfluent": {
			currentModel: &Model{
				ProjectId:    util.StringPtr("507f1f77bcf86cd799439011"),
				ProviderName: util.StringPtr("AWS"),
				Vendor:       util.StringPtr("CONFLUENT"),
				DnsDomain:    util.StringPtr("test.example.com"),
				Region:       util.StringPtr("us-east-1"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "ServiceEndpointId is required",
		},
		"missingRegionForS3": {
			currentModel: &Model{
				ProjectId:         util.StringPtr("507f1f77bcf86cd799439011"),
				ProviderName:      util.StringPtr("AWS"),
				Vendor:            util.StringPtr("S3"),
				ServiceEndpointId: util.StringPtr("com.amazonaws.us-east-1.s3"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "Region is required",
		},
		"missingServiceEndpointIdForS3": {
			currentModel: &Model{
				ProjectId:    util.StringPtr("507f1f77bcf86cd799439011"),
				ProviderName: util.StringPtr("AWS"),
				Vendor:       util.StringPtr("S3"),
				Region:       util.StringPtr("us-east-1"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "ServiceEndpointId is required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			req := handler.Request{
				RequestContext: handler.RequestContext{},
			}
			event, err := Create(req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestReadValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *Model
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"missingProjectId": {
			currentModel: &Model{
				Id: util.StringPtr("507f1f77bcf86cd799439012"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"missingId": {
			currentModel: &Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			req := handler.Request{
				RequestContext: handler.RequestContext{},
			}
			event, err := Read(req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestDeleteValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *Model
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"missingProjectId": {
			currentModel: &Model{
				Id: util.StringPtr("507f1f77bcf86cd799439012"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"missingId": {
			currentModel: &Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			req := handler.Request{
				RequestContext: handler.RequestContext{},
			}
			event, err := Delete(req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestListValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *Model
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"missingProjectId": {
			currentModel:   &Model{},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			req := handler.Request{
				RequestContext: handler.RequestContext{},
			}
			event, err := List(req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestUpdate(t *testing.T) {
	req := handler.Request{
		RequestContext: handler.RequestContext{},
	}
	currentModel := createTestModel()

	event, err := Update(req, nil, currentModel)

	require.NoError(t, err)
	assert.Equal(t, handler.Failed, event.OperationStatus)
	assert.Contains(t, event.Message, "not supported")
	assert.Equal(t, string(types.HandlerErrorCodeInvalidRequest), event.HandlerErrorCode)
}

// Test CRUD operations with mocks
func TestCreateWithMocks(t *testing.T) {
	// Save original function
	originalInitEnv := initEnvWithLatestClient
	defer func() {
		initEnvWithLatestClient = originalInitEnv
	}()

	testCases := map[string]struct {
		req            handler.Request
		currentModel   *Model
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
	}{
		"successfulCreateMSK": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.CreatePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().CreatePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				m.EXPECT().CreatePrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)

				// State transition check
				getReq := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(getReq)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		// Note: createWithCallback test is skipped because handleCreateCallback calls util.NewAtlasClient
		// which requires a valid handler.Request with Session, which is complex to mock in unit tests.
		// The callback logic is tested indirectly through waitForStateTransition tests.
		"createWithInProgressState": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.CreatePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().CreatePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				workingState := "WORKING"
				apiResp.State = &workingState
				m.EXPECT().CreatePrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)

				// State transition check - still in progress
				getReq := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(getReq)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"createWithFailedState": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.CreatePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().CreatePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				failedState := "FAILED"
				errorMsg := "Connection failed"
				apiResp.State = &failedState
				apiResp.ErrorMessage = &errorMsg
				m.EXPECT().CreatePrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)

				// State transition check - failed
				getReq := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(getReq)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"createWithApiError": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.CreatePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().CreatePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().CreatePrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("API error"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockApi

			initEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := Create(tc.req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

func TestReadWithMocks(t *testing.T) {
	// Save original function
	originalInitEnv := initEnvWithLatestClient
	defer func() {
		initEnvWithLatestClient = originalInitEnv
	}()

	testCases := map[string]struct {
		req            handler.Request
		currentModel   *Model
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
	}{
		"successfulRead": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"readNotFound": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
		"readWithEmptyDnsSubDomain": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: func() *Model {
				m := createTestModelConfluent()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				m.DnsSubDomain = nil
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				emptyDnsSubDomain := []string{}
				apiResp.DnsSubDomain = &emptyDnsSubDomain
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"readApiError": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockApi

			initEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := Read(tc.req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

func TestDeleteWithMocks(t *testing.T) {
	// Save original function
	originalInitEnv := initEnvWithLatestClient
	defer func() {
		initEnvWithLatestClient = originalInitEnv
	}()

	testCases := map[string]struct {
		req            handler.Request
		currentModel   *Model
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
	}{
		"successfulDelete": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.DeletePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().DeletePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeletePrivateLinkConnectionExecute(mock.Anything).Return(&http.Response{StatusCode: 200}, nil)

				// State transition check
				getReq := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(getReq)
				deletedState := "DELETED"
				apiResp := &admin20250312010.StreamsPrivateLinkConnection{
					State: &deletedState,
				}
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		// Note: deleteWithCallback test is skipped because handleDeleteCallback calls util.NewAtlasClient
		// which requires a valid handler.Request with Session, which is complex to mock in unit tests.
		// The callback logic is tested indirectly through waitForStateTransition tests.
		"deleteWithInProgressState": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.DeletePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().DeletePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeletePrivateLinkConnectionExecute(mock.Anything).Return(&http.Response{StatusCode: 200}, nil)

				// State transition check - still deleting
				getReq := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(getReq)
				deletingState := "DELETING"
				apiResp := &admin20250312010.StreamsPrivateLinkConnection{
					State: &deletingState,
				}
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"deleteWithNotFound": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.DeletePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().DeletePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeletePrivateLinkConnectionExecute(mock.Anything).Return(&http.Response{StatusCode: 200}, nil)

				// State transition check - 404 means already deleted
				getReq := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(getReq)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Success,
		},
		"deleteWithFailedState": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.DeletePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().DeletePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeletePrivateLinkConnectionExecute(mock.Anything).Return(&http.Response{StatusCode: 200}, nil)

				// State transition check - failed
				getReq := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(getReq)
				failedState := "FAILED"
				errorMsg := "Delete failed"
				apiResp := &admin20250312010.StreamsPrivateLinkConnection{
					State:        &failedState,
					ErrorMessage: &errorMsg,
				}
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"deleteWithApiError": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.DeletePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().DeletePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeletePrivateLinkConnectionExecute(mock.Anything).Return(&http.Response{StatusCode: 500}, fmt.Errorf("delete failed"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockApi

			initEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := Delete(tc.req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

func TestListWithMocks(t *testing.T) {
	// Save original function
	originalInitEnv := initEnvWithLatestClient
	defer func() {
		initEnvWithLatestClient = originalInitEnv
	}()

	testCases := map[string]struct {
		req            handler.Request
		currentModel   *Model
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
		expectedCount  int
	}{
		"successfulList": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.ListPrivateLinkConnectionsApiRequest{ApiService: m}
				m.EXPECT().ListPrivateLinkConnectionsWithParams(mock.Anything, mock.Anything).Return(req)
				connection1 := createTestAPIResponse()
				connectionID2 := "507f1f77bcf86cd799439013"
				connection2 := createTestAPIResponse()
				connection2.Id = &connectionID2
				results := []admin20250312010.StreamsPrivateLinkConnection{*connection1, *connection2}
				paginated := &admin20250312010.PaginatedApiStreamsPrivateLink{
					Results:    &results,
					TotalCount: util.Pointer(2),
				}
				m.EXPECT().ListPrivateLinkConnectionsExecute(mock.Anything).Return(paginated, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			expectedCount:  2,
		},
		"listApiError": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.ListPrivateLinkConnectionsApiRequest{ApiService: m}
				m.EXPECT().ListPrivateLinkConnectionsWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().ListPrivateLinkConnectionsExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("list failed"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockApi

			initEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := List(tc.req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.expectedStatus == handler.Success {
				if event.ResourceModels != nil {
					assert.Equal(t, tc.expectedCount, len(event.ResourceModels))
				}
			}
		})
	}
}

func TestWaitForStateTransition(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *Model
		pendingStates  []string
		targetStates   []string
		isDelete       bool
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
	}{
		"reachedTargetStateDone": {
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{StateIdle, StateWorking},
			targetStates:  []string{StateDone, StateFailed},
			isDelete:      false,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				doneState := "DONE"
				apiResp.State = &doneState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"reachedTargetStateFailed": {
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{StateIdle, StateWorking},
			targetStates:  []string{StateDone, StateFailed},
			isDelete:      false,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				failedState := "FAILED"
				errorMsg := "Connection failed"
				apiResp.State = &failedState
				apiResp.ErrorMessage = &errorMsg
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"inPendingState": {
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{StateIdle, StateWorking},
			targetStates:  []string{StateDone, StateFailed},
			isDelete:      false,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				workingState := "WORKING"
				apiResp.State = &workingState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"deleteReachedDeletedState": {
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{StateDeleteRequested, StateDeleting},
			targetStates:  []string{StateDeleted, StateFailed},
			isDelete:      true,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				deletedState := "DELETED"
				apiResp.State = &deletedState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"deleteWithNotFound": {
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{StateDeleteRequested, StateDeleting},
			targetStates:  []string{StateDeleted, StateFailed},
			isDelete:      true,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Success,
		},
		"unexpectedState": {
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{StateIdle, StateWorking},
			targetStates:  []string{StateDone, StateFailed},
			isDelete:      false,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				unexpectedState := "UNEXPECTED"
				apiResp.State = &unexpectedState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"getStateError": {
			currentModel: func() *Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{StateIdle, StateWorking},
			targetStates:  []string{StateDone, StateFailed},
			isDelete:      false,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockApi

			event, err := waitForStateTransition(context.Background(), mockClient, tc.currentModel, tc.pendingStates, tc.targetStates, tc.isDelete)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

func TestValidateVendorRequirements(t *testing.T) {
	testCases := map[string]struct {
		model         *Model
		expectedError bool
		expectedMsg   string
	}{
		"nilVendor": {
			model: &Model{
				ProviderName: util.StringPtr("AWS"),
			},
			expectedError: false,
		},
		"validMSK": {
			model: &Model{
				Vendor: util.StringPtr("MSK"),
				Arn:    util.StringPtr("arn:aws:kafka:us-east-1:123456789012:cluster/test"),
			},
			expectedError: false,
		},
		"confluentMissingDnsDomain": {
			model: &Model{
				Vendor:            util.StringPtr("CONFLUENT"),
				Region:            util.StringPtr("us-east-1"),
				ServiceEndpointId: util.StringPtr("com.amazonaws.vpce.us-east-1.vpce-svc-12345678"),
			},
			expectedError: true,
			expectedMsg:   "DnsDomain is required",
		},
		"confluentMissingRegion": {
			model: &Model{
				Vendor:            util.StringPtr("CONFLUENT"),
				DnsDomain:         util.StringPtr("test.example.com"),
				ServiceEndpointId: util.StringPtr("com.amazonaws.vpce.us-east-1.vpce-svc-12345678"),
			},
			expectedError: true,
			expectedMsg:   "Region is required",
		},
		"confluentMissingServiceEndpointId": {
			model: &Model{
				Vendor:    util.StringPtr("CONFLUENT"),
				DnsDomain: util.StringPtr("test.example.com"),
				Region:    util.StringPtr("us-east-1"),
			},
			expectedError: true,
			expectedMsg:   "ServiceEndpointId is required",
		},
		"mskMissingArn": {
			model: &Model{
				Vendor: util.StringPtr("MSK"),
			},
			expectedError: true,
			expectedMsg:   "Arn is required",
		},
		"mskWithRegion": {
			model: &Model{
				Vendor: util.StringPtr("MSK"),
				Arn:    util.StringPtr("arn:aws:kafka:us-east-1:123456789012:cluster/test"),
				Region: util.StringPtr("us-east-1"),
			},
			expectedError: true,
			expectedMsg:   "Region cannot be set for vendor MSK",
		},
		"s3MissingRegion": {
			model: &Model{
				Vendor:            util.StringPtr("S3"),
				ServiceEndpointId: util.StringPtr("com.amazonaws.us-east-1.s3"),
			},
			expectedError: true,
			expectedMsg:   "Region is required",
		},
		"s3MissingServiceEndpointId": {
			model: &Model{
				Vendor: util.StringPtr("S3"),
				Region: util.StringPtr("us-east-1"),
			},
			expectedError: true,
			expectedMsg:   "ServiceEndpointId is required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := validateVendorRequirements(tc.model)
			if tc.expectedError {
				require.NotNil(t, result)
				assert.Equal(t, handler.Failed, result.OperationStatus)
				assert.Contains(t, result.Message, tc.expectedMsg)
			} else {
				require.Nil(t, result)
			}
		})
	}
}

func TestHandleError(t *testing.T) {
	testCases := map[string]struct {
		response           *http.Response
		method             constants.CfnFunctions
		err                error
		expectedStatus     handler.Status
		expectedErrorCode  string
		expectedMsgContain string
	}{
		"notFoundError": {
			response: &http.Response{
				StatusCode: http.StatusNotFound,
			},
			method:             constants.READ,
			err:                fmt.Errorf("resource not found"),
			expectedStatus:     handler.Failed,
			expectedErrorCode:  string(types.HandlerErrorCodeNotFound),
			expectedMsgContain: constants.ResourceNotFound,
		},
		"otherError": {
			response: &http.Response{
				StatusCode: http.StatusBadRequest,
			},
			method:             constants.CREATE,
			err:                fmt.Errorf("invalid request"),
			expectedStatus:     handler.Failed,
			expectedErrorCode:  "",
			expectedMsgContain: "CREATE error: invalid request",
		},
		"nilResponse": {
			response:           nil,
			method:             constants.DELETE,
			err:                fmt.Errorf("network error"),
			expectedStatus:     handler.Failed,
			expectedErrorCode:  "",
			expectedMsgContain: "DELETE error: network error",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := handleError(tc.response, tc.method, tc.err)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsgContain)
			if tc.expectedErrorCode != "" {
				assert.Equal(t, tc.expectedErrorCode, event.HandlerErrorCode)
			}
		})
	}
}

func TestGetAllPrivateLinkConnections(t *testing.T) {
	testCases := map[string]struct {
		projectID     string
		mockSetup     func(*mockadmin.StreamsApi)
		expectedCount int
		expectedError bool
	}{
		"singlePage": {
			projectID: "507f1f77bcf86cd799439011",
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.ListPrivateLinkConnectionsApiRequest{ApiService: m}
				m.EXPECT().ListPrivateLinkConnectionsWithParams(mock.Anything, mock.Anything).Return(req)
				connection1 := createTestAPIResponse()
				results := []admin20250312010.StreamsPrivateLinkConnection{*connection1}
				paginated := &admin20250312010.PaginatedApiStreamsPrivateLink{
					Results:    &results,
					TotalCount: util.Pointer(1),
				}
				m.EXPECT().ListPrivateLinkConnectionsExecute(mock.Anything).Return(paginated, &http.Response{StatusCode: 200}, nil)
			},
			expectedCount: 1,
			expectedError: false,
		},
		"multiplePages": {
			projectID: "507f1f77bcf86cd799439011",
			mockSetup: func(m *mockadmin.StreamsApi) {
				// First page
				req1 := admin20250312010.ListPrivateLinkConnectionsApiRequest{ApiService: m}
				m.EXPECT().ListPrivateLinkConnectionsWithParams(mock.Anything, mock.Anything).Return(req1).Once()
				connection1 := createTestAPIResponse()
				results1 := []admin20250312010.StreamsPrivateLinkConnection{*connection1}
				paginated1 := &admin20250312010.PaginatedApiStreamsPrivateLink{
					Results:    &results1,
					TotalCount: util.Pointer(2),
				}
				m.EXPECT().ListPrivateLinkConnectionsExecute(mock.Anything).Return(paginated1, &http.Response{StatusCode: 200}, nil).Once()

				// Second page
				req2 := admin20250312010.ListPrivateLinkConnectionsApiRequest{ApiService: m}
				m.EXPECT().ListPrivateLinkConnectionsWithParams(mock.Anything, mock.Anything).Return(req2).Once()
				connectionID2 := "507f1f77bcf86cd799439013"
				connection2 := createTestAPIResponse()
				connection2.Id = &connectionID2
				results2 := []admin20250312010.StreamsPrivateLinkConnection{*connection2}
				paginated2 := &admin20250312010.PaginatedApiStreamsPrivateLink{
					Results:    &results2,
					TotalCount: util.Pointer(2),
				}
				m.EXPECT().ListPrivateLinkConnectionsExecute(mock.Anything).Return(paginated2, &http.Response{StatusCode: 200}, nil).Once()
			},
			expectedCount: 2,
			expectedError: false,
		},
		"apiError": {
			projectID: "507f1f77bcf86cd799439011",
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.ListPrivateLinkConnectionsApiRequest{ApiService: m}
				m.EXPECT().ListPrivateLinkConnectionsWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().ListPrivateLinkConnectionsExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedError: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockApi

			result, resp, err := getAllPrivateLinkConnections(context.Background(), mockClient, tc.projectID)

			if tc.expectedError {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedCount, len(result))
				assert.Nil(t, resp)
			}
		})
	}
}
