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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-processor/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312010/admin"
	"go.mongodb.org/atlas-sdk/v20250312010/mockadmin"
)

var (
	baseResourceModel = &resource.Model{
		ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
		ProcessorName: util.StringPtr("processor-1"),
		WorkspaceName: util.StringPtr("workspace-1"),
		Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
	}
	validProcessor = &admin.StreamsProcessorWithStats{
		Name:  "processor-1",
		Id:    "507f1f77bcf86cd799439011",
		State: resource.CreatedState,
	}
)

func TestList(t *testing.T) {
	originalInitEnv := resource.InitEnvWithLatestClient
	defer func() {
		resource.InitEnvWithLatestClient = originalInitEnv
	}()

	testCases := map[string]struct {
		currentModel   *resource.Model
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
		expectedCount  int
	}{
		"successfulListSinglePage": {
			currentModel: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorsApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorsWithParams(mock.Anything, mock.Anything).Return(req)
				processors := &admin.PaginatedApiStreamsStreamProcessorWithStats{
					Results: &[]admin.StreamsProcessorWithStats{
						{Name: "processor-1", Id: "507f1f77bcf86cd799439011", State: resource.CreatedState},
						{Name: "processor-2", Id: "507f1f77bcf86cd799439012", State: resource.StartedState},
					},
					TotalCount: util.Pointer(2),
				}
				m.EXPECT().GetStreamProcessorsExecute(mock.Anything).Return(processors, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			expectedCount:  2,
		},
		"listWithApiError": {
			currentModel: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorsApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorsWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorsExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
			expectedCount:  0,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsAPI := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsAPI)

			mockClient := &admin.APIClient{StreamsApi: mockStreamsAPI}
			resource.InitEnvWithLatestClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := resource.List(handler.Request{RequestContext: handler.RequestContext{}}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)

			if tc.expectedStatus == handler.Success {
				require.NotNil(t, event.ResourceModels)
				assert.Len(t, event.ResourceModels, tc.expectedCount)
			}
		})
	}
}

func TestSetup(t *testing.T) {
	assert.NotPanics(t, func() {
		resource.Setup()
	})
}

func TestValidationErrors(t *testing.T) {
	validationModels := map[string]*resource.Model{
		"missingProjectId": {
			ProcessorName: util.StringPtr("processor-1"),
			WorkspaceName: util.StringPtr("workspace-1"),
		},
		"missingProcessorName": {
			ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
			WorkspaceName: util.StringPtr("workspace-1"),
		},
	}

	operations := map[string]func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error){
		"Create": resource.Create,
		"Read":   resource.Read,
		"Update": resource.Update,
		"Delete": resource.Delete,
	}

	for opName, operation := range operations {
		for modelName, model := range validationModels {
			t.Run(opName+"_"+modelName, func(t *testing.T) {
				event, err := operation(handler.Request{RequestContext: handler.RequestContext{}}, nil, model)
				require.NoError(t, err)
				assert.Equal(t, handler.Failed, event.OperationStatus)
				assert.Contains(t, event.Message, "required")
			})
		}
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
		prevModel      *resource.Model
		currentModel   *resource.Model
		mockSetup      func(*mockadmin.StreamsApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
		req            handler.Request
	}{
		"Create_invalidState": {
			operation: resource.Create,
			currentModel: func() *resource.Model {
				m := *baseResourceModel
				m.DesiredState = util.StringPtr("INVALID_STATE")
				return &m
			}(),
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Failed,
		},
		"Create_apiError": {
			operation:    resource.Create,
			currentModel: baseResourceModel,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.CreateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().CreateStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().CreateStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("API error"))
			},
			expectedStatus: handler.Failed,
		},
		"Create_withCallback": {
			operation: resource.Create,
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamProcessor": true,
					"projectID":               "507f1f77bcf86cd799439011",
					"workspaceName":           "workspace-1",
					"processorName":           "processor-1",
					"needsStarting":           false,
					"startTime":               time.Now().Format(time.RFC3339),
					"timeoutDuration":         "20m",
					"deleteOnCreateTimeout":   false,
				},
			},
			currentModel: baseResourceModel,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(validProcessor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"Create_withDesiredStateStarted": {
			operation: resource.Create,
			currentModel: func() *resource.Model {
				m := *baseResourceModel
				m.DesiredState = util.StringPtr(resource.StartedState)
				return &m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.CreateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().CreateStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().CreateStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"Create_withTimeouts": {
			operation: resource.Create,
			currentModel: func() *resource.Model {
				m := *baseResourceModel
				m.Timeouts = &resource.Timeouts{Create: util.StringPtr("30m")}
				m.DeleteOnCreateTimeout = util.Pointer(false)
				return &m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.CreateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().CreateStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().CreateStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"Create_invalidPipeline": {
			operation: resource.Create,
			currentModel: func() *resource.Model {
				m := *baseResourceModel
				m.Pipeline = util.StringPtr("invalid json")
				return &m
			}(),
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Failed,
		},
		"Create_missingWorkspaceAndInstance": {
			operation: resource.Create,
			currentModel: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Failed,
		},
		"Read_success": {
			operation:    resource.Read,
			currentModel: baseResourceModel,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(validProcessor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"Read_notFound": {
			operation:    resource.Read,
			currentModel: baseResourceModel,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, "NotFound", event.HandlerErrorCode)
			},
		},
		"Read_apiError": {
			operation:    resource.Read,
			currentModel: baseResourceModel,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
		"Read_missingWorkspaceAndInstance": {
			operation: resource.Read,
			currentModel: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
			},
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Failed,
		},
		"Update_withCallback": {
			operation: resource.Update,
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamProcessor": true,
					"projectID":               "507f1f77bcf86cd799439011",
					"workspaceName":           "workspace-1",
					"processorName":           "processor-1",
					"desiredState":            resource.CreatedState,
				},
			},
			currentModel: baseResourceModel,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(validProcessor, &http.Response{StatusCode: 200}, nil)
				updateReq := admin.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(validProcessor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"Update_notFound": {
			operation:    resource.Update,
			currentModel: baseResourceModel,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, "NotFound", event.HandlerErrorCode)
			},
		},
		"Update_invalidStateTransition": {
			operation: resource.Update,
			currentModel: func() *resource.Model {
				m := *baseResourceModel
				m.DesiredState = util.StringPtr(resource.CreatedState)
				return &m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin.StreamsProcessorWithStats{Name: "processor-1", State: resource.StartedState}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"Update_stopError": {
			operation: resource.Update,
			currentModel: func() *resource.Model {
				m := *baseResourceModel
				m.DesiredState = util.StringPtr(resource.StoppedState)
				return &m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin.StreamsProcessorWithStats{Name: "processor-1", State: resource.StartedState}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
				stopReq := admin.StopStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().StopStreamProcessorWithParams(mock.Anything, mock.Anything).Return(stopReq)
				m.EXPECT().StopStreamProcessorExecute(mock.Anything).Return(nil, fmt.Errorf("stop failed"))
			},
			expectedStatus: handler.Failed,
		},
		"Update_stopSuccess": {
			operation: resource.Update,
			currentModel: func() *resource.Model {
				m := *baseResourceModel
				m.DesiredState = util.StringPtr(resource.StoppedState)
				return &m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin.StreamsProcessorWithStats{Name: "processor-1", State: resource.StartedState}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
				stopReq := admin.StopStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().StopStreamProcessorWithParams(mock.Anything, mock.Anything).Return(stopReq)
				m.EXPECT().StopStreamProcessorExecute(mock.Anything).Return(&http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"Update_startError": {
			operation: resource.Update,
			currentModel: func() *resource.Model {
				m := *baseResourceModel
				m.DesiredState = util.StringPtr(resource.StartedState)
				return &m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(validProcessor, &http.Response{StatusCode: 200}, nil)
				updateReq := admin.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(validProcessor, &http.Response{StatusCode: 200}, nil)
				startReq := admin.StartStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().StartStreamProcessorWithParams(mock.Anything, mock.Anything).Return(startReq)
				m.EXPECT().StartStreamProcessorExecute(mock.Anything).Return(nil, fmt.Errorf("start failed"))
			},
			expectedStatus: handler.Failed,
		},
		"Update_startSuccess": {
			operation: resource.Update,
			currentModel: func() *resource.Model {
				m := *baseResourceModel
				m.DesiredState = util.StringPtr(resource.StartedState)
				return &m
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(validProcessor, &http.Response{StatusCode: 200}, nil)
				updateReq := admin.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(validProcessor, &http.Response{StatusCode: 200}, nil)
				startReq := admin.StartStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().StartStreamProcessorWithParams(mock.Anything, mock.Anything).Return(startReq)
				m.EXPECT().StartStreamProcessorExecute(mock.Anything).Return(nil, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"Update_successWithoutStarting": {
			operation:    resource.Update,
			currentModel: baseResourceModel,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(validProcessor, &http.Response{StatusCode: 200}, nil)
				updateReq := admin.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(validProcessor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"Update_apiError": {
			operation:    resource.Update,
			currentModel: baseResourceModel,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(validProcessor, &http.Response{StatusCode: 200}, nil)
				updateReq := admin.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("update failed"))
			},
			expectedStatus: handler.Failed,
		},
		"Delete_success": {
			operation:    resource.Delete,
			currentModel: baseResourceModel,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.DeleteStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().DeleteStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeleteStreamProcessorExecute(mock.Anything).Return(&http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"Delete_notFound": {
			operation:    resource.Delete,
			currentModel: baseResourceModel,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.DeleteStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().DeleteStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeleteStreamProcessorExecute(mock.Anything).Return(&http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, "NotFound", event.HandlerErrorCode)
			},
		},
		"Delete_apiError": {
			operation:    resource.Delete,
			currentModel: baseResourceModel,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.DeleteStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().DeleteStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeleteStreamProcessorExecute(mock.Anything).Return(nil, fmt.Errorf("delete failed"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			cleanup := setupMockClient(t, tc.mockSetup)
			defer cleanup()

			event, err := tc.operation(tc.req, tc.prevModel, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.validateResult != nil {
				tc.validateResult(t, event)
			}
		})
	}
}
