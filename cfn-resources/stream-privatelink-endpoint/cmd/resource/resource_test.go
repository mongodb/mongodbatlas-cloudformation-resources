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
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-privatelink-endpoint/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312010/admin"
	"go.mongodb.org/atlas-sdk/v20250312010/mockadmin"
)

func createTestModel() *resource.Model {
	projectID := "507f1f77bcf86cd799439011"
	providerName := "AWS"
	vendor := "MSK"
	arn := "arn:aws:kafka:us-east-1:123456789012:cluster/test-cluster/12345678-1234-1234-1234-123456789012-1"

	return &resource.Model{
		Profile:      util.StringPtr("default"),
		ProjectId:    &projectID,
		ProviderName: &providerName,
		Vendor:       &vendor,
		Arn:          &arn,
	}
}

func createTestAPIResponse() *admin.StreamsPrivateLinkConnection {
	connectionID := "507f1f77bcf86cd799439012"
	providerName := "AWS"
	vendor := "MSK"
	state := "DONE"
	arn := "arn:aws:kafka:us-east-1:123456789012:cluster/test-cluster/12345678-1234-1234-1234-123456789012-1"

	return &admin.StreamsPrivateLinkConnection{
		Id:       &connectionID,
		Provider: providerName,
		Vendor:   &vendor,
		State:    &state,
		Arn:      &arn,
	}
}

func TestCRUDOperations(t *testing.T) {
	originalInitEnv := resource.InitEnvWithLatestClient
	defer func() {
		resource.InitEnvWithLatestClient = originalInitEnv
	}()

	testCases := map[string]struct {
		currentModel   *resource.Model
		mockSetup      func(*mockadmin.StreamsApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		operation      string
		expectedStatus handler.Status
		req            handler.Request
	}{
		"Create_success": {
			operation:    "CREATE",
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.CreatePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().CreatePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				doneState := resource.StateDone
				apiResp.State = &doneState
				m.EXPECT().CreatePrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"Create_inProgress": {
			operation:    "CREATE",
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.CreatePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().CreatePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				workingState := resource.StateWorking
				apiResp.State = &workingState
				m.EXPECT().CreatePrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"Create_apiError": {
			operation:    "CREATE",
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.CreatePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().CreatePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().CreatePrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("API error"))
			},
			expectedStatus: handler.Failed,
		},
		"Create_failedState": {
			operation:    "CREATE",
			currentModel: createTestModel(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.CreatePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().CreatePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				failedState := resource.StateFailed
				errorMsg := "Connection failed"
				apiResp.State = &failedState
				apiResp.ErrorMessage = &errorMsg
				m.EXPECT().CreatePrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"Read_success": {
			operation: "READ",
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"Read_notFound": {
			operation: "READ",
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
			},
		},
		"Update_notSupported": {
			operation:      "UPDATE",
			currentModel:   createTestModel(),
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Contains(t, event.Message, "not supported")
				assert.Equal(t, string(types.HandlerErrorCodeInvalidRequest), event.HandlerErrorCode)
			},
		},
		"Delete_success": {
			operation: "DELETE",
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				getReq := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(getReq)
				apiResp := createTestAPIResponse()
				doneState := resource.StateDone
				apiResp.State = &doneState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)

				delReq := admin.DeletePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().DeletePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(delReq)
				m.EXPECT().DeletePrivateLinkConnectionExecute(mock.Anything).Return(&http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"Delete_notFound": {
			operation: "DELETE",
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
			},
		},
		"Create_withCallback": {
			operation:    "CREATE",
			currentModel: createTestModel(),
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamPrivatelinkEndpoint": true,
					"projectID":                         "507f1f77bcf86cd799439011",
					"connectionID":                      "507f1f77bcf86cd799439012",
				},
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				doneState := resource.StateDone
				apiResp.State = &doneState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"Delete_withCallback": {
			operation: "DELETE",
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamPrivatelinkEndpoint": true,
					"projectID":                         "507f1f77bcf86cd799439011",
					"connectionID":                      "507f1f77bcf86cd799439012",
				},
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				deletedState := resource.StateDeleted
				apiResp.State = &deletedState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"Delete_alreadyDeleted": {
			operation: "DELETE",
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				deletedState := resource.StateDeleted
				apiResp.State = &deletedState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
			},
		},
		"Delete_deleteApiError": {
			operation: "DELETE",
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				getReq := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(getReq)
				apiResp := createTestAPIResponse()
				doneState := resource.StateDone
				apiResp.State = &doneState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)

				delReq := admin.DeletePrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().DeletePrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(delReq)
				m.EXPECT().DeletePrivateLinkConnectionExecute(mock.Anything).Return(&http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
		},
		"Delete_getApiServerError": {
			operation: "DELETE",
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
		},
		"Delete_getApiOtherError": {
			operation: "DELETE",
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 400}, fmt.Errorf("bad request"))
			},
			expectedStatus: handler.Failed,
		},
		"Read_apiError": {
			operation: "READ",
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
		},
		"List_success": {
			operation:      "LIST",
			currentModel:   createTestModel(),
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.NotNil(t, event.ResourceModels)
				assert.Len(t, event.ResourceModels, 2)
			},
		},
		"List_apiError": {
			operation:      "LIST",
			currentModel:   createTestModel(),
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			runCRUDTest(t, name, tc)
		})
	}
}

func runCRUDTest(t *testing.T, name string, tc struct {
	currentModel   *resource.Model
	mockSetup      func(*mockadmin.StreamsApi)
	validateResult func(t *testing.T, event handler.ProgressEvent)
	operation      string
	expectedStatus handler.Status
	req            handler.Request
}) {
	t.Helper()
	originalGetAll := resource.GetAllPrivateLinkConnections
	defer func() {
		resource.GetAllPrivateLinkConnections = originalGetAll
	}()

	mockAPI := mockadmin.NewStreamsApi(t)
	tc.mockSetup(mockAPI)

	mockClient := &admin.APIClient{}
	mockClient.StreamsApi = mockAPI

	resource.InitEnvWithLatestClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin.APIClient, *handler.ProgressEvent) {
		return mockClient, nil
	}

	if tc.operation == "LIST" {
		setupListMock(name)
	}

	req := tc.req
	if req.CallbackContext == nil {
		req = handler.Request{RequestContext: handler.RequestContext{}}
	}

	event, err := executeOperation(tc.operation, req, tc.currentModel)
	require.NoError(t, err)
	assert.Equal(t, tc.expectedStatus, event.OperationStatus)
	if tc.validateResult != nil {
		tc.validateResult(t, event)
	}
}

func setupListMock(name string) {
	switch name {
	case "List_success":
		resource.GetAllPrivateLinkConnections = func(ctx context.Context, conn *admin.APIClient, projectID string) ([]admin.StreamsPrivateLinkConnection, *http.Response, error) {
			return []admin.StreamsPrivateLinkConnection{
				*createTestAPIResponse(),
				func() admin.StreamsPrivateLinkConnection {
					conn := *createTestAPIResponse()
					deletedState := resource.StateDeleted
					conn.State = &deletedState
					return conn
				}(),
				func() admin.StreamsPrivateLinkConnection {
					conn := *createTestAPIResponse()
					doneState := resource.StateDone
					conn.State = &doneState
					return conn
				}(),
			}, nil, nil
		}
	case "List_apiError":
		resource.GetAllPrivateLinkConnections = func(ctx context.Context, conn *admin.APIClient, projectID string) ([]admin.StreamsPrivateLinkConnection, *http.Response, error) {
			return nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error")
		}
	}
}

func executeOperation(operation string, req handler.Request, currentModel *resource.Model) (handler.ProgressEvent, error) {
	switch operation {
	case "CREATE":
		return resource.Create(req, nil, currentModel)
	case "READ":
		return resource.Read(req, nil, currentModel)
	case "UPDATE":
		return resource.Update(req, nil, currentModel)
	case "DELETE":
		return resource.Delete(req, nil, currentModel)
	case "LIST":
		return resource.List(req, nil, currentModel)
	default:
		return handler.ProgressEvent{}, fmt.Errorf("unknown operation: %s", operation)
	}
}

func TestValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		operation      string
		currentModel   *resource.Model
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"Create_missingProjectId": {
			operation: "CREATE",
			currentModel: &resource.Model{
				ProviderName: util.StringPtr("AWS"),
				Vendor:       util.StringPtr("MSK"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"Create_missingVendor": {
			operation: "CREATE",
			currentModel: &resource.Model{
				ProjectId:    util.StringPtr("507f1f77bcf86cd799439011"),
				ProviderName: util.StringPtr("AWS"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"Create_missingArnForMSK": {
			operation: "CREATE",
			currentModel: &resource.Model{
				ProjectId:    util.StringPtr("507f1f77bcf86cd799439011"),
				ProviderName: util.StringPtr("AWS"),
				Vendor:       util.StringPtr("MSK"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "Arn is required",
		},
		"Read_missingProjectId": {
			operation: "READ",
			currentModel: &resource.Model{
				Id: util.StringPtr("507f1f77bcf86cd799439012"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"Read_missingId": {
			operation: "READ",
			currentModel: &resource.Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"Delete_missingId": {
			operation: "DELETE",
			currentModel: &resource.Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    constants.ResourceNotFound,
		},
		"List_missingProjectId": {
			operation:      "LIST",
			currentModel:   &resource.Model{},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			req := handler.Request{RequestContext: handler.RequestContext{}}
			var event handler.ProgressEvent
			var err error

			switch tc.operation {
			case "CREATE":
				event, err = resource.Create(req, nil, tc.currentModel)
			case "READ":
				event, err = resource.Read(req, nil, tc.currentModel)
			case "DELETE":
				event, err = resource.Delete(req, nil, tc.currentModel)
			case "LIST":
				event, err = resource.List(req, nil, tc.currentModel)
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}
