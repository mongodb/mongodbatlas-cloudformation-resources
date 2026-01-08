// Copyright 2025 MongoDB Inc
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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"
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

func createTestQueryLimit() *admin20250312010.DataFederationTenantQueryLimit {
	tenantName := testTenantName
	overrunPolicy := "BLOCK"
	currentUsage := int64(1000000000)
	defaultLimit := int64(1000000000)
	maximumLimit := int64(5000000000)
	lastModified := time.Now()

	return &admin20250312010.DataFederationTenantQueryLimit{
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

func createMockClient(mockAPI *mockadmin.DataFederationApi) *util.MongoDBClient {
	return &util.MongoDBClient{
		AtlasSDK: &admin20250312010.APIClient{
			DataFederationApi: mockAPI,
		},
	}
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		currentModel   *resource.Model
		mockSetup      func(*mockadmin.DataFederationApi)
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"validationFailure": {
			currentModel:   &resource.Model{},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"alreadyExists": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin20250312010.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(createTestQueryLimit(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "already exists",
		},
		"success": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin20250312010.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
				setReq := admin20250312010.SetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().SetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName, mock.Anything).Return(setReq)
				m.EXPECT().SetDataFederationLimitExecute(mock.Anything).Return(createTestQueryLimit(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			expectedMsg:    "CREATE completed",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			req := handler.Request{RequestContext: handler.RequestContext{}}

			if tc.mockSetup != nil {
				mockAPI := mockadmin.NewDataFederationApi(t)
				tc.mockSetup(mockAPI)
				originalFunc := resource.NewAtlasClientFunc
				resource.NewAtlasClientFunc = func(req *handler.Request, profile *string) (*util.MongoDBClient, *handler.ProgressEvent) {
					return createMockClient(mockAPI), nil
				}
				defer func() { resource.NewAtlasClientFunc = originalFunc }()
			}

			event, err := resource.Create(req, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.expectedMsg != "" {
				assert.Contains(t, event.Message, tc.expectedMsg)
			}
		})
	}
}

func TestRead(t *testing.T) {
	tests := map[string]struct {
		currentModel   *resource.Model
		mockSetup      func(*mockadmin.DataFederationApi)
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"validationFailure": {
			currentModel:   &resource.Model{},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"success": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin20250312010.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(createTestQueryLimit(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			expectedMsg:    "Read Completed",
		},
		"notFound": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin20250312010.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "READ error",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			req := handler.Request{RequestContext: handler.RequestContext{}}

			if tc.mockSetup != nil {
				mockAPI := mockadmin.NewDataFederationApi(t)
				tc.mockSetup(mockAPI)
				originalFunc := resource.NewAtlasClientFunc
				resource.NewAtlasClientFunc = func(req *handler.Request, profile *string) (*util.MongoDBClient, *handler.ProgressEvent) {
					return createMockClient(mockAPI), nil
				}
				defer func() { resource.NewAtlasClientFunc = originalFunc }()
			}

			event, err := resource.Read(req, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := map[string]struct {
		currentModel   *resource.Model
		mockSetup      func(*mockadmin.DataFederationApi)
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"validationFailure": {
			currentModel:   &resource.Model{},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"notFound": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin20250312010.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "does not exist",
		},
		"success": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin20250312010.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(createTestQueryLimit(), &http.Response{StatusCode: 200}, nil)
				setReq := admin20250312010.SetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().SetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName, mock.Anything).Return(setReq)
				m.EXPECT().SetDataFederationLimitExecute(mock.Anything).Return(createTestQueryLimit(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			expectedMsg:    "UPDATE completed",
		},
		"otherError": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin20250312010.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "UPDATE error",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			req := handler.Request{RequestContext: handler.RequestContext{}}

			if tc.mockSetup != nil {
				mockAPI := mockadmin.NewDataFederationApi(t)
				tc.mockSetup(mockAPI)
				originalFunc := resource.NewAtlasClientFunc
				resource.NewAtlasClientFunc = func(req *handler.Request, profile *string) (*util.MongoDBClient, *handler.ProgressEvent) {
					return createMockClient(mockAPI), nil
				}
				defer func() { resource.NewAtlasClientFunc = originalFunc }()
			}

			event, err := resource.Update(req, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestDelete(t *testing.T) {
	tests := map[string]struct {
		currentModel   *resource.Model
		mockSetup      func(*mockadmin.DataFederationApi)
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"validationFailure": {
			currentModel:   &resource.Model{},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"success": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				m.EXPECT().DeleteDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(admin20250312010.DeleteDataFederationLimitApiRequest{ApiService: m})
				m.EXPECT().DeleteDataFederationLimitExecute(mock.Anything).Return(&http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			expectedMsg:    "Delete Completed",
		},
		"error": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				m.EXPECT().DeleteDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(admin20250312010.DeleteDataFederationLimitApiRequest{ApiService: m})
				m.EXPECT().DeleteDataFederationLimitExecute(mock.Anything).Return(&http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "DELETE error",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			req := handler.Request{RequestContext: handler.RequestContext{}}

			if tc.mockSetup != nil {
				mockAPI := mockadmin.NewDataFederationApi(t)
				tc.mockSetup(mockAPI)
				originalFunc := resource.NewAtlasClientFunc
				resource.NewAtlasClientFunc = func(req *handler.Request, profile *string) (*util.MongoDBClient, *handler.ProgressEvent) {
					return createMockClient(mockAPI), nil
				}
				defer func() { resource.NewAtlasClientFunc = originalFunc }()
			}

			event, err := resource.Delete(req, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestList(t *testing.T) {
	tests := map[string]struct {
		currentModel   *resource.Model
		mockSetup      func(*mockadmin.DataFederationApi)
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"validationFailure": {
			currentModel:   &resource.Model{},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"success": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				limits := []admin20250312010.DataFederationTenantQueryLimit{*createTestQueryLimit()}
				m.EXPECT().ListDataFederationLimits(mock.Anything, testProjectID, testTenantName).Return(admin20250312010.ListDataFederationLimitsApiRequest{ApiService: m})
				m.EXPECT().ListDataFederationLimitsExecute(mock.Anything).Return(limits, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			expectedMsg:    "List Completed",
		},
		"error": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				m.EXPECT().ListDataFederationLimits(mock.Anything, testProjectID, testTenantName).Return(admin20250312010.ListDataFederationLimitsApiRequest{ApiService: m})
				m.EXPECT().ListDataFederationLimitsExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "LIST error",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			req := handler.Request{RequestContext: handler.RequestContext{}}

			if tc.mockSetup != nil {
				mockAPI := mockadmin.NewDataFederationApi(t)
				tc.mockSetup(mockAPI)
				originalFunc := resource.NewAtlasClientFunc
				resource.NewAtlasClientFunc = func(req *handler.Request, profile *string) (*util.MongoDBClient, *handler.ProgressEvent) {
					return createMockClient(mockAPI), nil
				}
				defer func() { resource.NewAtlasClientFunc = originalFunc }()
			}

			event, err := resource.List(req, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestValidationErrors(t *testing.T) {
	requiredFields := map[string]struct {
		createFields []string
		readFields   []string
		deleteFields []string
		listFields   []string
	}{
		"Create": {
			createFields: []string{"ProjectId", "TenantName", "LimitName", "Value", "OverrunPolicy"},
		},
		"Read": {
			readFields: []string{"ProjectId", "TenantName", "LimitName"},
		},
		"Delete": {
			deleteFields: []string{"ProjectId", "TenantName", "LimitName"},
		},
		"List": {
			listFields: []string{"ProjectId", "TenantName"},
		},
	}

	for op, fields := range requiredFields {
		t.Run(op, func(t *testing.T) {
			req := handler.Request{RequestContext: handler.RequestContext{}}
			var event handler.ProgressEvent
			var err error

			switch {
			case len(fields.createFields) > 0:
				event, err = resource.Create(req, nil, &resource.Model{})
			case len(fields.readFields) > 0:
				event, err = resource.Read(req, nil, &resource.Model{})
			case len(fields.deleteFields) > 0:
				event, err = resource.Delete(req, nil, &resource.Model{})
			case len(fields.listFields) > 0:
				event, err = resource.List(req, nil, &resource.Model{})
			}

			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, "required")
		})
	}
}

func TestProfileDefaulting(t *testing.T) {
	tests := map[string]struct {
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

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model := createTestModel()
			model.Profile = tc.profile

			mockAPI := mockadmin.NewDataFederationApi(t)
			getReq := admin20250312010.GetDataFederationLimitApiRequest{ApiService: mockAPI}
			mockAPI.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
			mockAPI.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(createTestQueryLimit(), &http.Response{StatusCode: 200}, nil)

			originalFunc := resource.NewAtlasClientFunc
			profileUsed := ""
			resource.NewAtlasClientFunc = func(req *handler.Request, profile *string) (*util.MongoDBClient, *handler.ProgressEvent) {
				if profile != nil {
					profileUsed = *profile
				}
				return createMockClient(mockAPI), nil
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
	tests := map[string]struct {
		currentModel      *resource.Model
		mockSetup         func(*mockadmin.DataFederationApi)
		operation         func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error)
		expectedStatus    handler.Status
		expectedMsg       string
		expectedErrorCode string
	}{
		"conflictError": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin20250312010.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
				setReq := admin20250312010.SetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().SetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName, mock.Anything).Return(setReq)
				m.EXPECT().SetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: http.StatusConflict}, fmt.Errorf("conflict"))
			},
			expectedStatus:    handler.Failed,
			expectedErrorCode: string(types.HandlerErrorCodeAlreadyExists),
			expectedMsg:       "CREATE error",
			operation:         resource.Create,
		},
		"unauthorizedError": {
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.DataFederationApi) {
				getReq := admin20250312010.GetDataFederationLimitApiRequest{ApiService: m}
				m.EXPECT().GetDataFederationLimit(mock.Anything, testProjectID, testTenantName, testLimitName).Return(getReq)
				m.EXPECT().GetDataFederationLimitExecute(mock.Anything).Return(nil, &http.Response{StatusCode: http.StatusUnauthorized}, fmt.Errorf("unauthorized"))
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "Not found",
			operation:      resource.Read,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			req := handler.Request{RequestContext: handler.RequestContext{}}

			if tc.mockSetup != nil {
				mockAPI := mockadmin.NewDataFederationApi(t)
				tc.mockSetup(mockAPI)
				originalFunc := resource.NewAtlasClientFunc
				resource.NewAtlasClientFunc = func(req *handler.Request, profile *string) (*util.MongoDBClient, *handler.ProgressEvent) {
					return createMockClient(mockAPI), nil
				}
				defer func() { resource.NewAtlasClientFunc = originalFunc }()
			}

			event, err := tc.operation(req, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
			if tc.expectedErrorCode != "" {
				assert.Equal(t, tc.expectedErrorCode, event.HandlerErrorCode)
			}
		})
	}
}
