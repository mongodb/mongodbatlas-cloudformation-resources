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
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/federated-query-limit/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312010/admin"
	"go.mongodb.org/atlas-sdk/v20250312010/mockadmin"
)

const (
	testProjectID  = "507f1f77bcf86cd799439011"
	testTenantName = "test-tenant"
	testLimitName  = "bytesProcessed.query"
	testValue      = "2000000000"
)

func createTestModel() *resource.Model {
	profile := "default"
	return &resource.Model{
		ProjectId:     util.StringPtr(testProjectID),
		TenantName:    util.StringPtr(testTenantName),
		LimitName:     util.StringPtr(testLimitName),
		OverrunPolicy: util.StringPtr("BLOCK"),
		Value:         util.StringPtr(testValue),
		Profile:       &profile,
	}
}

func createTestQueryLimit() *admin.DataFederationTenantQueryLimit {
	tenantName := testTenantName
	overrunPolicy := "BLOCK"
	currentUsage := int64(1000000000)
	defaultLimit := int64(1000000000)
	maximumLimit := int64(5000000000)
	lastModified := time.Now()

	return &admin.DataFederationTenantQueryLimit{
		TenantName:       &tenantName,
		Name:             testLimitName,
		OverrunPolicy:    &overrunPolicy,
		Value:            int64(2000000000),
		CurrentUsage:     &currentUsage,
		DefaultLimit:     &defaultLimit,
		MaximumLimit:     &maximumLimit,
		LastModifiedDate: &lastModified,
	}
}

func TestConstants(t *testing.T) {
	assert.Equal(t, []string{constants.ProjectID, constants.TenantName, constants.LimitName, constants.OverrunPolicy, constants.Value}, resource.CreateOrUpdateRequiredFields)
	assert.Equal(t, []string{constants.ProjectID, constants.TenantName, constants.LimitName}, resource.ReadRequiredFields)
	assert.Equal(t, []string{constants.ProjectID, constants.TenantName, constants.LimitName}, resource.DeleteRequiredFields)
	assert.Equal(t, []string{constants.ProjectID, constants.TenantName}, resource.ListRequiredFields)
	assert.Equal(t, "already exists", resource.AlreadyExists)
	assert.Equal(t, "does not exist", resource.DoesntExists)
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
				TenantName:    util.StringPtr(testTenantName),
				LimitName:     util.StringPtr(testLimitName),
				OverrunPolicy: util.StringPtr("BLOCK"),
				Value:         util.StringPtr(testValue),
			},
			expectedMsg: "required",
		},
		"Create_missingTenantName": {
			operation: resource.Create,
			currentModel: &resource.Model{
				ProjectId:     util.StringPtr(testProjectID),
				LimitName:     util.StringPtr(testLimitName),
				OverrunPolicy: util.StringPtr("BLOCK"),
				Value:         util.StringPtr(testValue),
			},
			expectedMsg: "required",
		},
		"Create_missingLimitName": {
			operation: resource.Create,
			currentModel: &resource.Model{
				ProjectId:     util.StringPtr(testProjectID),
				TenantName:    util.StringPtr(testTenantName),
				OverrunPolicy: util.StringPtr("BLOCK"),
				Value:         util.StringPtr(testValue),
			},
			expectedMsg: "required",
		},
		"Create_missingOverrunPolicy": {
			operation: resource.Create,
			currentModel: &resource.Model{
				ProjectId:  util.StringPtr(testProjectID),
				TenantName: util.StringPtr(testTenantName),
				LimitName:  util.StringPtr(testLimitName),
				Value:      util.StringPtr(testValue),
			},
			expectedMsg: "required",
		},
		"Create_missingValue": {
			operation: resource.Create,
			currentModel: &resource.Model{
				ProjectId:     util.StringPtr(testProjectID),
				TenantName:    util.StringPtr(testTenantName),
				LimitName:     util.StringPtr(testLimitName),
				OverrunPolicy: util.StringPtr("BLOCK"),
			},
			expectedMsg: "required",
		},
		"Read_missingProjectId": {
			operation:    resource.Read,
			currentModel: &resource.Model{TenantName: util.StringPtr(testTenantName), LimitName: util.StringPtr(testLimitName)},
			expectedMsg:  "required",
		},
		"Read_missingTenantName": {
			operation:    resource.Read,
			currentModel: &resource.Model{ProjectId: util.StringPtr(testProjectID), LimitName: util.StringPtr(testLimitName)},
			expectedMsg:  "required",
		},
		"Read_missingLimitName": {
			operation:    resource.Read,
			currentModel: &resource.Model{ProjectId: util.StringPtr(testProjectID), TenantName: util.StringPtr(testTenantName)},
			expectedMsg:  "required",
		},
		"Update_missingProjectId": {
			operation: resource.Update,
			currentModel: &resource.Model{
				TenantName:    util.StringPtr(testTenantName),
				LimitName:     util.StringPtr(testLimitName),
				OverrunPolicy: util.StringPtr("BLOCK"),
				Value:         util.StringPtr(testValue),
			},
			expectedMsg: "required",
		},
		"Delete_missingProjectId": {
			operation:    resource.Delete,
			currentModel: &resource.Model{TenantName: util.StringPtr(testTenantName), LimitName: util.StringPtr(testLimitName)},
			expectedMsg:  "required",
		},
		"List_missingProjectId": {
			operation:    resource.List,
			currentModel: &resource.Model{TenantName: util.StringPtr(testTenantName)},
			expectedMsg:  "required",
		},
		"List_missingTenantName": {
			operation:    resource.List,
			currentModel: &resource.Model{ProjectId: util.StringPtr(testProjectID)},
			expectedMsg:  "required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := tc.operation(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func setupMockClient(t *testing.T, mockSetup func(*mockadmin.DataFederationApi)) func() {
	t.Helper()
	originalFunc := resource.NewAtlasClientFunc
	mockAPI := mockadmin.NewDataFederationApi(t)
	mockSetup(mockAPI)

	mockClient := &util.MongoDBClient{
		AtlasSDK: &admin.APIClient{
			DataFederationApi: mockAPI,
		},
	}
	// Override the variable to inject mock client for testing
	resource.NewAtlasClientFunc = func(req *handler.Request, profile *string) (*util.MongoDBClient, *handler.ProgressEvent) {
		return mockClient, nil
	}

	return func() { resource.NewAtlasClientFunc = originalFunc }
}

func TestCRUDOperations(t *testing.T) {
	testCases := map[string]struct {
		operation      func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error)
		mockSetup      func(*mockadmin.DataFederationApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
	}{
		"Create_alreadyExists": {
			operation: resource.Create,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(createTestQueryLimit(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, string(types.HandlerErrorCodeAlreadyExists), event.HandlerErrorCode)
				assert.Contains(t, event.Message, "already exists")
			},
		},
		"Create_success": {
			operation: resource.Create,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
				setReq := admin.SetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().SetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName, mock.Anything).Return(setReq)
				m.EXPECT().SetDataFederationLimitExecute(mock.Anything).Return(createTestQueryLimit(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Contains(t, event.Message, "CREATE completed")
			},
		},
		"Create_apiError": {
			operation: resource.Create,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
				setReq := admin.SetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().SetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName, mock.Anything).Return(setReq)
				m.EXPECT().SetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("API error"))
			},
			expectedStatus: handler.Failed,
		},
		"Read_success": {
			operation: resource.Read,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(createTestQueryLimit(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Contains(t, event.Message, "Read Completed")
				require.NotNil(t, event.ResourceModel)
			},
		},
		"Read_notFound": {
			operation: resource.Read,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
		"Read_apiError": {
			operation: resource.Read,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
		"Update_notFound": {
			operation: resource.Update,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
				assert.Contains(t, event.Message, "does not exist")
			},
		},
		"Update_success": {
			operation: resource.Update,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(createTestQueryLimit(), &http.Response{StatusCode: 200}, nil)
				setReq := admin.SetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().SetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName, mock.Anything).Return(setReq)
				m.EXPECT().SetDataFederationLimitExecute(mock.Anything).Return(createTestQueryLimit(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Contains(t, event.Message, "UPDATE completed")
			},
		},
		"Update_apiError": {
			operation: resource.Update,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
				setReq := admin.SetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().SetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName, mock.Anything).Return(setReq)
				m.EXPECT().SetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("set failed"))
			},
			expectedStatus: handler.Failed,
		},
		"Delete_success": {
			operation: resource.Delete,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				m.EXPECT().DeleteDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(admin.DeleteDataFederationLimitApiRequest{ApiService: m})
				m.EXPECT().DeleteDataFederationLimitExecute(mock.Anything).Return(&http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Contains(t, event.Message, "Delete Completed")
			},
		},
		"Delete_apiError": {
			operation: resource.Delete,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				m.EXPECT().DeleteDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(admin.DeleteDataFederationLimitApiRequest{ApiService: m})
				m.EXPECT().DeleteDataFederationLimitExecute(mock.Anything).Return(&http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
		"List_success": {
			operation: resource.List,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				limits := []admin.DataFederationTenantQueryLimit{*createTestQueryLimit()}
				m.EXPECT().ListDataFederationLimits(mock.Anything, testProjectID, testTenantName).Return(admin.ListDataFederationLimitsApiRequest{ApiService: m})
				m.EXPECT().ListDataFederationLimitsExecute(mock.Anything).Return(limits, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Contains(t, event.Message, "List Completed")
				require.NotNil(t, event.ResourceModels)
				assert.GreaterOrEqual(t, len(event.ResourceModels), 1)
			},
		},
		"List_apiError": {
			operation: resource.List,
			mockSetup: func(m *mockadmin.DataFederationApi) {
				m.EXPECT().ListDataFederationLimits(mock.Anything, testProjectID, testTenantName).Return(admin.ListDataFederationLimitsApiRequest{ApiService: m})
				m.EXPECT().ListDataFederationLimitsExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			cleanup := setupMockClient(t, tc.mockSetup)
			defer cleanup()

			event, err := tc.operation(handler.Request{}, nil, createTestModel())
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.validateResult != nil {
				tc.validateResult(t, event)
			}
		})
	}
}

func TestProfileDefaulting(t *testing.T) {
	testCases := map[string]struct {
		profile         *string
		expectedDefault bool
	}{
		"nilProfile": {
			profile:         nil,
			expectedDefault: true,
		},
		"emptyProfile": {
			profile:         util.StringPtr(""),
			expectedDefault: true,
		},
		"customProfile": {
			profile:         util.StringPtr("custom"),
			expectedDefault: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			model := createTestModel()
			model.Profile = tc.profile

			mockAPI := mockadmin.NewDataFederationApi(t)
			getReq := admin.GetDataFederationLimitApiRequest{ApiService: mockAPI}
			mockAPI.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
			mockAPI.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(createTestQueryLimit(), &http.Response{StatusCode: 200}, nil)

			originalFunc := resource.NewAtlasClientFunc
			profileUsed := ""
			resource.NewAtlasClientFunc = func(req *handler.Request, profile *string) (*util.MongoDBClient, *handler.ProgressEvent) {
				if profile != nil {
					profileUsed = *profile
				}
				return &util.MongoDBClient{
					AtlasSDK: &admin.APIClient{
						DataFederationApi: mockAPI,
					},
				}, nil
			}
			defer func() { resource.NewAtlasClientFunc = originalFunc }()

			req := handler.Request{RequestContext: handler.RequestContext{}}
			event, err := resource.Read(req, nil, model)
			require.NoError(t, err)
			assert.Equal(t, handler.Success, event.OperationStatus)
			if tc.expectedDefault {
				assert.Equal(t, "default", profileUsed)
			} else {
				assert.Equal(t, "custom", profileUsed)
			}
		})
	}
}

func TestErrorHandling(t *testing.T) {
	testCases := map[string]struct {
		currentModel      *resource.Model
		mockSetup         func(*mockadmin.DataFederationApi)
		operation         func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error)
		expectedStatus    handler.Status
		expectedMsg       string
		expectedErrorCode string
	}{
		"Create_conflictError": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
				setReq := admin.SetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().SetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName, mock.Anything).Return(setReq)
				m.EXPECT().SetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: http.StatusConflict}, fmt.Errorf("conflict"))
			},
			expectedStatus:    handler.Failed,
			expectedErrorCode: string(types.HandlerErrorCodeAlreadyExists),
			expectedMsg:       "CREATE error",
			operation:         resource.Create,
		},
		"Read_unauthorizedError": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: http.StatusUnauthorized}, fmt.Errorf("unauthorized"))
			},
			expectedStatus:    handler.Failed,
			expectedMsg:       "Not found",
			expectedErrorCode: string(types.HandlerErrorCodeNotFound),
			operation:         resource.Read,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			cleanup := setupMockClient(t, tc.mockSetup)
			defer cleanup()

			event, err := tc.operation(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
			if tc.expectedErrorCode != "" {
				assert.Equal(t, tc.expectedErrorCode, event.HandlerErrorCode)
			}
		})
	}
}
