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
	baseModel = &resource.Model{
		ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
		ProcessorName: util.StringPtr("processor-1"),
	}
	baseCallbackCtx = &resource.CallbackData{
		ProjectID:               "507f1f77bcf86cd799439011",
		WorkspaceOrInstanceName: "workspace-1",
		ProcessorName:           "processor-1",
		StartTime:               time.Now().Format(time.RFC3339),
		TimeoutDuration:         "20m",
	}
)

func TestIsCallback(t *testing.T) {
	testCases := map[string]struct {
		req            handler.Request
		expectedResult bool
	}{
		"isCallback": {
			req: handler.Request{
				CallbackContext: map[string]any{"callbackStreamProcessor": true},
			},
			expectedResult: true,
		},
		"notCallback": {
			req: handler.Request{
				CallbackContext: map[string]any{},
			},
			expectedResult: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, resource.IsCallback(&tc.req))
		})
	}
}

func TestGetCallbackData(t *testing.T) {
	testCases := map[string]struct {
		expectedResult *resource.CallbackData
		req            handler.Request
	}{
		"allFieldsPresent": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"projectID":             "507f1f77bcf86cd799439011",
					"workspaceName":         "workspace-1",
					"processorName":         "processor-1",
					"needsStarting":         true,
					"desiredState":          "STARTED",
					"startTime":             "2024-01-01T00:00:00Z",
					"timeoutDuration":       "20m",
					"deleteOnCreateTimeout": true,
				},
			},
			expectedResult: &resource.CallbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           true,
				DesiredState:            "STARTED",
				StartTime:               "2024-01-01T00:00:00Z",
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   true,
			},
		},
		"partialFields": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"projectID":     "507f1f77bcf86cd799439011",
					"workspaceName": "workspace-1",
					"processorName": "processor-1",
				},
			},
			expectedResult: &resource.CallbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
			},
		},
		"emptyContext": {
			req: handler.Request{
				CallbackContext: map[string]any{},
			},
			expectedResult: &resource.CallbackData{},
		},
		"typeAssertionFailures": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"projectID":             123,
					"workspaceName":         true,
					"processorName":         []string{"invalid"},
					"needsStarting":         "not a bool",
					"desiredState":          456,
					"startTime":             struct{}{},
					"timeoutDuration":       nil,
					"deleteOnCreateTimeout": "not a bool",
				},
			},
			expectedResult: &resource.CallbackData{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, resource.GetCallbackData(tc.req))
		})
	}
}

func TestValidateCallbackData(t *testing.T) {
	validCtx := &resource.CallbackData{
		ProjectID:               "507f1f77bcf86cd799439011",
		WorkspaceOrInstanceName: "workspace-1",
		ProcessorName:           "processor-1",
	}

	testCases := map[string]struct {
		callbackCtx        *resource.CallbackData
		expectedMsgContain string
		expectedError      bool
	}{
		"valid": {
			callbackCtx:   validCtx,
			expectedError: false,
		},
		"missingProjectID": {
			callbackCtx:        &resource.CallbackData{WorkspaceOrInstanceName: "workspace-1", ProcessorName: "processor-1"},
			expectedError:      true,
			expectedMsgContain: "Missing required values",
		},
		"missingWorkspaceName": {
			callbackCtx:        &resource.CallbackData{ProjectID: "507f1f77bcf86cd799439011", ProcessorName: "processor-1"},
			expectedError:      true,
			expectedMsgContain: "Missing required values",
		},
		"missingProcessorName": {
			callbackCtx:        &resource.CallbackData{ProjectID: "507f1f77bcf86cd799439011", WorkspaceOrInstanceName: "workspace-1"},
			expectedError:      true,
			expectedMsgContain: "Missing required values",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			peErr := resource.ValidateCallbackData(tc.callbackCtx)
			if tc.expectedError {
				require.NotNil(t, peErr)
				assert.Contains(t, peErr.Message, tc.expectedMsgContain)
			} else {
				require.Nil(t, peErr)
			}
		})
	}
}

func TestBuildCallbackContext(t *testing.T) {
	testCases := map[string]struct {
		additionalFields map[string]any
		validateFunc     func(t *testing.T, ctx map[string]any)
	}{
		"basic": {
			additionalFields: map[string]any{},
			validateFunc: func(t *testing.T, ctx map[string]any) {
				t.Helper()
				assert.True(t, ctx["callbackStreamProcessor"].(bool))
				assert.Equal(t, "507f1f77bcf86cd799439011", ctx["projectID"])
				assert.Equal(t, "workspace-1", ctx["workspaceName"])
				assert.Equal(t, "processor-1", ctx["processorName"])
			},
		},
		"withAdditionalFields": {
			additionalFields: map[string]any{"needsStarting": true, "desiredState": "STARTED"},
			validateFunc: func(t *testing.T, ctx map[string]any) {
				t.Helper()
				assert.True(t, ctx["callbackStreamProcessor"].(bool))
				assert.Equal(t, "507f1f77bcf86cd799439011", ctx["projectID"])
				assert.True(t, ctx["needsStarting"].(bool))
				assert.Equal(t, "STARTED", ctx["desiredState"])
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := resource.BuildCallbackContext("507f1f77bcf86cd799439011", "workspace-1", "processor-1", tc.additionalFields)
			if tc.validateFunc != nil {
				tc.validateFunc(t, ctx)
			}
		})
	}
}

func TestHandleCreateCallback(t *testing.T) {
	mockClient := &admin.APIClient{StreamsApi: mockadmin.NewStreamsApi(t)}
	ctx := context.Background()

	timeoutCtx := func(deleteOnTimeout bool) *resource.CallbackData {
		ctx := *baseCallbackCtx
		ctx.StartTime = time.Now().Add(-25 * time.Minute).Format(time.RFC3339)
		ctx.DeleteOnCreateTimeout = deleteOnTimeout
		return &ctx
	}

	createMockProcessor := func(state string) *admin.StreamsProcessorWithStats {
		return &admin.StreamsProcessorWithStats{Name: "processor-1", State: state}
	}

	setupGetProcessor := func(m *mockadmin.StreamsApi, processor *admin.StreamsProcessorWithStats) {
		req := admin.GetStreamProcessorApiRequest{ApiService: m}
		m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
		m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
	}

	setupStartProcessor := func(m *mockadmin.StreamsApi) {
		startReq := admin.StartStreamProcessorApiRequest{ApiService: m}
		m.EXPECT().StartStreamProcessorWithParams(mock.Anything, mock.Anything).Return(startReq)
		m.EXPECT().StartStreamProcessorExecute(mock.Anything).Return(nil, nil)
	}

	testCases := map[string]struct {
		currentModel   *resource.Model
		callbackCtx    *resource.CallbackData
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"timeoutExceededWithCleanup": {
			currentModel: baseModel,
			callbackCtx:  timeoutCtx(true),
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.DeleteStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().DeleteStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeleteStreamProcessorExecute(mock.Anything).Return(nil, nil)
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "Timeout reached",
		},
		"timeoutExceededWithoutCleanup": {
			currentModel:   baseModel,
			callbackCtx:    timeoutCtx(false),
			expectedStatus: handler.Failed,
			expectedMsg:    "Timeout reached",
		},
		"createdStateNeedsStarting": {
			currentModel: baseModel,
			callbackCtx: func() *resource.CallbackData {
				ctx := *baseCallbackCtx
				ctx.NeedsStarting = true
				return &ctx
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				setupGetProcessor(m, createMockProcessor(resource.CreatedState))
				setupStartProcessor(m)
			},
			expectedStatus: handler.InProgress,
			expectedMsg:    "Starting stream processor",
		},
		"createdStateNoStarting": {
			currentModel: baseModel,
			callbackCtx: func() *resource.CallbackData {
				ctx := *baseCallbackCtx
				ctx.NeedsStarting = false
				return &ctx
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				setupGetProcessor(m, createMockProcessor(resource.CreatedState))
			},
			expectedStatus: handler.Success,
			expectedMsg:    "Create Completed",
		},
		"startedState": {
			currentModel: baseModel,
			callbackCtx:  baseCallbackCtx,
			mockSetup: func(m *mockadmin.StreamsApi) {
				setupGetProcessor(m, createMockProcessor(resource.StartedState))
			},
			expectedStatus: handler.Success,
			expectedMsg:    "Create Completed",
		},
		"creatingState": {
			currentModel: baseModel,
			callbackCtx:  baseCallbackCtx,
			mockSetup: func(m *mockadmin.StreamsApi) {
				setupGetProcessor(m, createMockProcessor(resource.CreatingState))
			},
			expectedStatus: handler.InProgress,
			expectedMsg:    "Creating stream processor",
		},
		"failedState": {
			currentModel: baseModel,
			callbackCtx:  baseCallbackCtx,
			mockSetup: func(m *mockadmin.StreamsApi) {
				setupGetProcessor(m, createMockProcessor(resource.FailedState))
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "FAILED state",
		},
		"unexpectedState": {
			currentModel: baseModel,
			callbackCtx:  baseCallbackCtx,
			mockSetup: func(m *mockadmin.StreamsApi) {
				setupGetProcessor(m, &admin.StreamsProcessorWithStats{Name: "processor-1", State: "UNKNOWN"})
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "Unexpected state",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsAPI := mockadmin.NewStreamsApi(t)
			if tc.mockSetup != nil {
				tc.mockSetup(mockStreamsAPI)
			}
			mockClient.StreamsApi = mockStreamsAPI

			event, err := resource.HandleCreateCallback(ctx, mockClient, tc.currentModel, tc.callbackCtx)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestHandleUpdateCallback(t *testing.T) {
	mockClient := &admin.APIClient{StreamsApi: mockadmin.NewStreamsApi(t)}
	ctx := context.Background()

	createMockProcessor := func(state string) *admin.StreamsProcessorWithStats {
		return &admin.StreamsProcessorWithStats{Name: "processor-1", State: state}
	}

	setupGetProcessor := func(m *mockadmin.StreamsApi, processor *admin.StreamsProcessorWithStats) {
		req := admin.GetStreamProcessorApiRequest{ApiService: m}
		m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
		m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
	}

	setupUpdateProcessor := func(m *mockadmin.StreamsApi, updatedState string) {
		updateReq := admin.UpdateStreamProcessorApiRequest{ApiService: m}
		m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
		m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(createMockProcessor(updatedState), &http.Response{StatusCode: 200}, nil)
	}

	setupStartProcessor := func(m *mockadmin.StreamsApi) {
		startReq := admin.StartStreamProcessorApiRequest{ApiService: m}
		m.EXPECT().StartStreamProcessorWithParams(mock.Anything, mock.Anything).Return(startReq)
		m.EXPECT().StartStreamProcessorExecute(mock.Anything).Return(nil, nil)
	}

	setupStopProcessor := func(m *mockadmin.StreamsApi) {
		stopReq := admin.StopStreamProcessorApiRequest{ApiService: m}
		m.EXPECT().StopStreamProcessorWithParams(mock.Anything, mock.Anything).Return(stopReq)
		m.EXPECT().StopStreamProcessorExecute(mock.Anything).Return(nil, nil)
	}

	testCases := map[string]struct {
		currentModel   *resource.Model
		callbackCtx    *resource.CallbackData
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"stoppedStateWithDesiredStarted": {
			currentModel: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			callbackCtx: func() *resource.CallbackData {
				ctx := *baseCallbackCtx
				ctx.DesiredState = resource.StartedState
				return &ctx
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				setupGetProcessor(m, createMockProcessor(resource.StoppedState))
				setupUpdateProcessor(m, resource.StoppedState)
				setupStartProcessor(m)
			},
			expectedStatus: handler.InProgress,
			expectedMsg:    "Starting stream processor",
		},
		"startedStateWithDesiredStarted": {
			currentModel: baseModel,
			callbackCtx: func() *resource.CallbackData {
				ctx := *baseCallbackCtx
				ctx.DesiredState = resource.StartedState
				return &ctx
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				setupGetProcessor(m, createMockProcessor(resource.StartedState))
			},
			expectedStatus: handler.Success,
			expectedMsg:    "Update Completed",
		},
		"startedStateWithDesiredStopped": {
			currentModel: baseModel,
			callbackCtx: func() *resource.CallbackData {
				ctx := *baseCallbackCtx
				ctx.DesiredState = resource.StoppedState
				return &ctx
			}(),
			mockSetup: func(m *mockadmin.StreamsApi) {
				setupGetProcessor(m, createMockProcessor(resource.StartedState))
				setupStopProcessor(m)
			},
			expectedStatus: handler.InProgress,
			expectedMsg:    "Stopping stream processor",
		},
		"failedState": {
			currentModel: baseModel,
			callbackCtx:  baseCallbackCtx,
			mockSetup: func(m *mockadmin.StreamsApi) {
				setupGetProcessor(m, createMockProcessor(resource.FailedState))
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "FAILED state",
		},
		"defaultState": {
			currentModel: baseModel,
			callbackCtx:  baseCallbackCtx,
			mockSetup: func(m *mockadmin.StreamsApi) {
				setupGetProcessor(m, &admin.StreamsProcessorWithStats{Name: "processor-1", State: "UNKNOWN"})
			},
			expectedStatus: handler.InProgress,
			expectedMsg:    "Updating stream processor",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsAPI := mockadmin.NewStreamsApi(t)
			if tc.mockSetup != nil {
				tc.mockSetup(mockStreamsAPI)
			}
			mockClient.StreamsApi = mockStreamsAPI

			event, err := resource.HandleUpdateCallback(ctx, mockClient, tc.currentModel, tc.callbackCtx)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}
