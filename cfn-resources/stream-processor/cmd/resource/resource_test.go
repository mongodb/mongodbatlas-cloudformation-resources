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
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"
	"go.mongodb.org/atlas-sdk/v20250312010/mockadmin"
)

func TestIsCallback(t *testing.T) {
	testCases := map[string]struct {
		req            handler.Request
		expectedResult bool
	}{
		"isCallback": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamProcessor": true,
				},
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
			result := isCallback(&tc.req)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestGetCallbackData(t *testing.T) {
	testCases := map[string]struct {
		req            handler.Request
		expectedResult *callbackData
	}{
		"allFieldsPresent": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"projectID":             "507f1f77bcf86cd799439011",
					"workspaceName":         "workspace-1",
					"processorName":         "processor-1",
					"needsStarting":         true,
					"plannedState":          "STARTED",
					"startTime":             "2024-01-01T00:00:00Z",
					"timeoutDuration":       "20m",
					"deleteOnCreateTimeout": true,
				},
			},
			expectedResult: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           true,
				PlannedState:            "STARTED",
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
			expectedResult: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
			},
		},
		"emptyContext": {
			req: handler.Request{
				CallbackContext: map[string]any{},
			},
			expectedResult: &callbackData{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := getCallbackData(tc.req)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestValidateCallbackData(t *testing.T) {
	testCases := map[string]struct {
		callbackCtx        *callbackData
		expectedError      bool
		expectedMsgContain string
	}{
		"valid": {
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
			},
			expectedError: false,
		},
		"missingProjectID": {
			callbackCtx: &callbackData{
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
			},
			expectedError:      true,
			expectedMsgContain: "Missing required values",
		},
		"missingWorkspaceName": {
			callbackCtx: &callbackData{
				ProjectID:     "507f1f77bcf86cd799439011",
				ProcessorName: "processor-1",
			},
			expectedError:      true,
			expectedMsgContain: "Missing required values",
		},
		"missingProcessorName": {
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
			},
			expectedError:      true,
			expectedMsgContain: "Missing required values",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			peErr := validateCallbackData(tc.callbackCtx)
			if tc.expectedError {
				require.NotNil(t, peErr)
				assert.Contains(t, peErr.Message, tc.expectedMsgContain)
			} else {
				require.Nil(t, peErr)
			}
		})
	}
}

func TestCopyIdentifyingFields(t *testing.T) {
	testCases := map[string]struct {
		resourceModel *Model
		currentModel  *Model
		validateFunc  func(t *testing.T, resourceModel *Model)
	}{
		"withWorkspaceName": {
			resourceModel: &Model{},
			currentModel: &Model{
				Profile:       util.StringPtr("default"),
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			validateFunc: func(t *testing.T, rm *Model) {
				assert.Equal(t, "default", util.SafeString(rm.Profile))
				assert.Equal(t, "507f1f77bcf86cd799439011", util.SafeString(rm.ProjectId))
				assert.Equal(t, "processor-1", util.SafeString(rm.ProcessorName))
				assert.Equal(t, "workspace-1", util.SafeString(rm.WorkspaceName))
				// Primary identifier requires both fields - InstanceName should be set from WorkspaceName
				assert.Equal(t, "workspace-1", util.SafeString(rm.InstanceName))
			},
		},
		"withInstanceName": {
			resourceModel: &Model{},
			currentModel: &Model{
				Profile:       util.StringPtr("default"),
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				InstanceName:  util.StringPtr("instance-1"),
			},
			validateFunc: func(t *testing.T, rm *Model) {
				assert.Equal(t, "default", util.SafeString(rm.Profile))
				assert.Equal(t, "507f1f77bcf86cd799439011", util.SafeString(rm.ProjectId))
				assert.Equal(t, "processor-1", util.SafeString(rm.ProcessorName))
				assert.Equal(t, "instance-1", util.SafeString(rm.InstanceName))
				// Primary identifier requires both fields - WorkspaceName should be set from InstanceName
				assert.Equal(t, "instance-1", util.SafeString(rm.WorkspaceName))
			},
		},
		"emptyWorkspaceName": {
			resourceModel: &Model{},
			currentModel: &Model{
				Profile:       util.StringPtr("default"),
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr(""),
			},
			validateFunc: func(t *testing.T, rm *Model) {
				assert.Nil(t, rm.WorkspaceName)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			copyIdentifyingFields(tc.resourceModel, tc.currentModel)
			if tc.validateFunc != nil {
				tc.validateFunc(t, tc.resourceModel)
			}
		})
	}
}

func TestBuildCallbackContext(t *testing.T) {
	testCases := map[string]struct {
		projectID               string
		workspaceOrInstanceName string
		processorName           string
		additionalFields        map[string]any
		validateFunc            func(t *testing.T, ctx map[string]any)
	}{
		"basic": {
			projectID:               "507f1f77bcf86cd799439011",
			workspaceOrInstanceName: "workspace-1",
			processorName:           "processor-1",
			additionalFields:        map[string]any{},
			validateFunc: func(t *testing.T, ctx map[string]any) {
				assert.True(t, ctx["callbackStreamProcessor"].(bool))
				assert.Equal(t, "507f1f77bcf86cd799439011", ctx["projectID"])
				assert.Equal(t, "workspace-1", ctx["workspaceName"])
				assert.Equal(t, "processor-1", ctx["processorName"])
			},
		},
		"withAdditionalFields": {
			projectID:               "507f1f77bcf86cd799439011",
			workspaceOrInstanceName: "workspace-1",
			processorName:           "processor-1",
			additionalFields: map[string]any{
				"needsStarting": true,
				"plannedState":  "STARTED",
			},
			validateFunc: func(t *testing.T, ctx map[string]any) {
				assert.True(t, ctx["callbackStreamProcessor"].(bool))
				assert.Equal(t, "507f1f77bcf86cd799439011", ctx["projectID"])
				assert.True(t, ctx["needsStarting"].(bool))
				assert.Equal(t, "STARTED", ctx["plannedState"])
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := buildCallbackContext(tc.projectID, tc.workspaceOrInstanceName, tc.processorName, tc.additionalFields)
			if tc.validateFunc != nil {
				tc.validateFunc(t, ctx)
			}
		})
	}
}

func TestParseTimeout(t *testing.T) {
	testCases := map[string]struct {
		timeoutStr     string
		expectedResult time.Duration
	}{
		"validDuration": {
			timeoutStr:     "20m",
			expectedResult: 20 * time.Minute,
		},
		"validSeconds": {
			timeoutStr:     "30s",
			expectedResult: 30 * time.Second,
		},
		"emptyString": {
			timeoutStr:     "",
			expectedResult: defaultCreateTimeout,
		},
		"invalidFormat": {
			timeoutStr:     "invalid",
			expectedResult: defaultCreateTimeout,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := parseTimeout(tc.timeoutStr)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestIsTimeoutExceeded(t *testing.T) {
	testCases := map[string]struct {
		startTimeStr       string
		timeoutDurationStr string
		expectedResult     bool
	}{
		"timeoutExceeded": {
			startTimeStr:       time.Now().Add(-25 * time.Minute).Format(time.RFC3339),
			timeoutDurationStr: "20m",
			expectedResult:     true,
		},
		"timeoutNotExceeded": {
			startTimeStr:       time.Now().Add(-10 * time.Minute).Format(time.RFC3339),
			timeoutDurationStr: "20m",
			expectedResult:     false,
		},
		"emptyStartTime": {
			startTimeStr:       "",
			timeoutDurationStr: "20m",
			expectedResult:     false,
		},
		"emptyTimeoutDuration": {
			startTimeStr:       time.Now().Format(time.RFC3339),
			timeoutDurationStr: "",
			expectedResult:     false,
		},
		"invalidStartTime": {
			startTimeStr:       "invalid",
			timeoutDurationStr: "20m",
			expectedResult:     false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := isTimeoutExceeded(tc.startTimeStr, tc.timeoutDurationStr)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestCreateInProgressEvent(t *testing.T) {
	testCases := map[string]struct {
		message         string
		currentModel    *Model
		callbackContext map[string]any
		validateFunc    func(t *testing.T, event handler.ProgressEvent)
	}{
		"basic": {
			message:      "Creating stream processor",
			currentModel: &Model{ProjectId: util.StringPtr("507f1f77bcf86cd799439011")},
			callbackContext: map[string]any{
				"projectID": "507f1f77bcf86cd799439011",
			},
			validateFunc: func(t *testing.T, event handler.ProgressEvent) {
				assert.Equal(t, handler.InProgress, event.OperationStatus)
				assert.Equal(t, "Creating stream processor", event.Message)
				assert.Equal(t, int64(defaultCallbackDelaySeconds), event.CallbackDelaySeconds)
				assert.NotNil(t, event.CallbackContext)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event := createInProgressEvent(tc.message, tc.currentModel, tc.callbackContext)
			if tc.validateFunc != nil {
				tc.validateFunc(t, event)
			}
		})
	}
}

func TestValidateUpdateStateTransition(t *testing.T) {
	testCases := map[string]struct {
		currentState    string
		plannedState    string
		expectedIsValid bool
		expectedErrMsg  string
	}{
		"validCREATEDtoSTARTED": {
			currentState:    CreatedState,
			plannedState:    StartedState,
			expectedIsValid: true,
		},
		"invalidSTARTEDtoCREATED": {
			currentState:    StartedState,
			plannedState:    CreatedState,
			expectedIsValid: false,
			expectedErrMsg:  "cannot transition from STARTED to CREATED",
		},
		"validSTARTEDtoSTOPPED": {
			currentState:    StartedState,
			plannedState:    StoppedState,
			expectedIsValid: true,
		},
		"invalidCREATEDtoSTOPPED": {
			currentState:    CreatedState,
			plannedState:    StoppedState,
			expectedIsValid: false,
			expectedErrMsg:  "must be in STARTED state",
		},
		"sameState": {
			currentState:    CreatedState,
			plannedState:    CreatedState,
			expectedIsValid: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			errMsg, isValid := validateUpdateStateTransition(tc.currentState, tc.plannedState)
			assert.Equal(t, tc.expectedIsValid, isValid)
			if !tc.expectedIsValid {
				assert.Contains(t, errMsg, tc.expectedErrMsg)
			}
		})
	}
}

func TestList(t *testing.T) {
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
		"successfulListSinglePage": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorsApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorsWithParams(mock.Anything, mock.Anything).Return(req)
				processors := &admin20250312010.PaginatedApiStreamsStreamProcessorWithStats{
					Results: &[]admin20250312010.StreamsProcessorWithStats{
						{
							Name:  "processor-1",
							Id:    "507f1f77bcf86cd799439011",
							State: CreatedState,
						},
						{
							Name:  "processor-2",
							Id:    "507f1f77bcf86cd799439012",
							State: StartedState,
						},
					},
					TotalCount: util.Pointer(2),
				}
				m.EXPECT().GetStreamProcessorsExecute(mock.Anything).Return(processors, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			expectedCount:  2,
		},
		"successfulListMultiplePages": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				// First page - returns full page (100 items) but totalCount is 101, so we need to fetch page 2
				// Create 100 processors for first page
				firstPageResults := make([]admin20250312010.StreamsProcessorWithStats, 100)
				for i := 0; i < 100; i++ {
					firstPageResults[i] = admin20250312010.StreamsProcessorWithStats{
						Name:  fmt.Sprintf("processor-%d", i+1),
						Id:    fmt.Sprintf("507f1f77bcf86cd79943%03d", i),
						State: CreatedState,
					}
				}
				req1 := admin20250312010.GetStreamProcessorsApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorsWithParams(mock.Anything, mock.MatchedBy(func(params *admin20250312010.GetStreamProcessorsApiParams) bool {
					return params.PageNum != nil && *params.PageNum == 1
				})).Return(req1).Once()
				processors1 := &admin20250312010.PaginatedApiStreamsStreamProcessorWithStats{
					Results:    &firstPageResults,
					TotalCount: util.Pointer(101),
				}
				m.EXPECT().GetStreamProcessorsExecute(mock.Anything).Return(processors1, &http.Response{StatusCode: 200}, nil).Once()

				// Second page - returns 1 remaining item
				req2 := admin20250312010.GetStreamProcessorsApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorsWithParams(mock.Anything, mock.MatchedBy(func(params *admin20250312010.GetStreamProcessorsApiParams) bool {
					return params.PageNum != nil && *params.PageNum == 2
				})).Return(req2).Once()
				processors2 := &admin20250312010.PaginatedApiStreamsStreamProcessorWithStats{
					Results: &[]admin20250312010.StreamsProcessorWithStats{
						{
							Name:  "processor-101",
							Id:    "507f1f77bcf86cd799439101",
							State: StartedState,
						},
					},
					TotalCount: util.Pointer(101),
				}
				m.EXPECT().GetStreamProcessorsExecute(mock.Anything).Return(processors2, &http.Response{StatusCode: 200}, nil).Once()
			},
			expectedStatus: handler.Success,
			expectedCount:  101,
		},
		"successfulListEmpty": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorsApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorsWithParams(mock.Anything, mock.Anything).Return(req)
				processors := &admin20250312010.PaginatedApiStreamsStreamProcessorWithStats{
					Results:    &[]admin20250312010.StreamsProcessorWithStats{},
					TotalCount: util.Pointer(0),
				}
				m.EXPECT().GetStreamProcessorsExecute(mock.Anything).Return(processors, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			expectedCount:  0,
		},
		"listWithInstanceName": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId:    util.StringPtr("507f1f77bcf86cd799439011"),
				InstanceName: util.StringPtr("instance-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorsApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorsWithParams(mock.Anything, mock.Anything).Return(req)
				processors := &admin20250312010.PaginatedApiStreamsStreamProcessorWithStats{
					Results: &[]admin20250312010.StreamsProcessorWithStats{
						{
							Name:  "processor-1",
							Id:    "507f1f77bcf86cd799439011",
							State: CreatedState,
						},
					},
					TotalCount: util.Pointer(1),
				}
				m.EXPECT().GetStreamProcessorsExecute(mock.Anything).Return(processors, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			expectedCount:  1,
		},
		"listWithApiError": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorsApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorsWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorsExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
			expectedCount:  0,
		},
		"listWithMissingWorkspaceAndInstance": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
			},
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Failed,
			expectedCount:  0,
		},
		"listWithPipelineAndOptions": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorsApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorsWithParams(mock.Anything, mock.Anything).Return(req)
				pipeline := []any{
					map[string]any{"$match": map[string]any{"status": "active"}},
				}
				stats := map[string]any{
					"processed": 100,
					"errors":    0,
				}
				processors := &admin20250312010.PaginatedApiStreamsStreamProcessorWithStats{
					Results: &[]admin20250312010.StreamsProcessorWithStats{
						{
							Name:     "processor-1",
							Id:       "507f1f77bcf86cd799439011",
							State:    CreatedState,
							Pipeline: pipeline,
							Stats:    stats,
							Options: &admin20250312010.StreamsOptions{
								Dlq: &admin20250312010.StreamsDLQ{
									Coll:           util.StringPtr("dlq_collection"),
									ConnectionName: util.StringPtr("connection-1"),
									Db:             util.StringPtr("dlq_db"),
								},
							},
						},
					},
					TotalCount: util.Pointer(1),
				}
				m.EXPECT().GetStreamProcessorsExecute(mock.Anything).Return(processors, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			expectedCount:  1,
		},
		"listValidationError": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Failed,
			expectedCount:  0,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// For validation error test cases, validation should fail before API calls
			if name == "listValidationError" || name == "listWithMissingWorkspaceAndInstance" {
				// Mock initEnvWithLatestClient to run real validation but return error
				initEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
					// Run validation - this should fail for these test cases
					if errEvent := validator.ValidateModel(requiredFields, currentModel); errEvent != nil {
						return nil, errEvent
					}
					// If validation passes (shouldn't happen for these cases), return nil client
					return nil, &handler.ProgressEvent{
						OperationStatus: handler.Failed,
						Message:         "unexpected validation success",
					}
				}
				// Validation will fail and return early, so no API calls should be made
				event, err := List(tc.req, nil, tc.currentModel)
				require.NoError(t, err)
				assert.Equal(t, tc.expectedStatus, event.OperationStatus)
				return
			}

			// Mock initEnvWithLatestClient for other test cases
			mockStreamsApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockStreamsApi

			initEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := List(tc.req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)

			if tc.expectedStatus == handler.Success {
				require.NotNil(t, event.ResourceModels)
				assert.Equal(t, tc.expectedCount, len(event.ResourceModels))

				// Verify that each model has identifying fields set
				for i, rm := range event.ResourceModels {
					model, ok := rm.(*Model)
					require.True(t, ok, "ResourceModel[%d] should be *Model", i)
					assert.Equal(t, tc.currentModel.ProjectId, model.ProjectId)
					if tc.currentModel.WorkspaceName != nil {
						assert.Equal(t, tc.currentModel.WorkspaceName, model.WorkspaceName)
					}
					if tc.currentModel.InstanceName != nil {
						assert.Equal(t, tc.currentModel.InstanceName, model.InstanceName)
					}
				}
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
		"conflictError": {
			response: &http.Response{
				StatusCode: http.StatusConflict,
			},
			method:             constants.CREATE,
			err:                fmt.Errorf("resource already exists"),
			expectedStatus:     handler.Failed,
			expectedErrorCode:  "AlreadyExists",
			expectedMsgContain: "CREATE error:resource already exists",
		},
		"otherError": {
			response: &http.Response{
				StatusCode: http.StatusBadRequest,
			},
			method:             constants.UPDATE,
			err:                fmt.Errorf("invalid request"),
			expectedStatus:     handler.Failed,
			expectedErrorCode:  "",
			expectedMsgContain: "UPDATE error:invalid request",
		},
		"nilResponse": {
			response:           nil,
			method:             constants.DELETE,
			err:                fmt.Errorf("network error"),
			expectedStatus:     handler.Failed,
			expectedErrorCode:  "",
			expectedMsgContain: "DELETE error:network error",
		},
		"notFoundError": {
			response: &http.Response{
				StatusCode: http.StatusNotFound,
			},
			method:             constants.READ,
			err:                fmt.Errorf("resource not found"),
			expectedStatus:     handler.Failed,
			expectedErrorCode:  "",
			expectedMsgContain: "READ error:resource not found",
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

func TestFinalizeModel(t *testing.T) {
	testCases := map[string]struct {
		streamProcessor *admin20250312010.StreamsProcessorWithStats
		currentModel    *Model
		message         string
		expectedStatus  handler.Status
		expectedError   bool
	}{
		"successfulFinalize": {
			streamProcessor: &admin20250312010.StreamsProcessorWithStats{
				Name:  "test-processor",
				Id:    "507f1f77bcf86cd799439011",
				State: "CREATED",
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
			},
			message:        "Create Complete",
			expectedStatus: handler.Success,
			expectedError:  false,
		},
		"withPipeline": {
			streamProcessor: func() *admin20250312010.StreamsProcessorWithStats {
				pipeline := []any{
					map[string]any{"$match": map[string]any{"status": "active"}},
				}
				return &admin20250312010.StreamsProcessorWithStats{
					Name:     "test-processor",
					Id:       "507f1f77bcf86cd799439011",
					State:    "STARTED",
					Pipeline: pipeline,
				}
			}(),
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
			},
			message:        "Update Complete",
			expectedStatus: handler.Success,
			expectedError:  false,
		},
		"withStats": {
			streamProcessor: func() *admin20250312010.StreamsProcessorWithStats {
				stats := map[string]any{
					"processed": 100,
					"errors":    0,
				}
				return &admin20250312010.StreamsProcessorWithStats{
					Name:  "test-processor",
					Id:    "507f1f77bcf86cd799439011",
					State: "CREATED",
					Stats: stats,
				}
			}(),
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
			},
			message:        "Read Complete",
			expectedStatus: handler.Success,
			expectedError:  false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := finalizeModel(tc.streamProcessor, tc.currentModel, tc.message)

			if tc.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedStatus, event.OperationStatus)
				assert.Equal(t, tc.message, event.Message)
				require.NotNil(t, event.ResourceModel)
				resourceModel, ok := event.ResourceModel.(*Model)
				require.True(t, ok, "ResourceModel should be *Model")
				assert.Equal(t, tc.currentModel.ProjectId, resourceModel.ProjectId)
				assert.Equal(t, tc.currentModel.ProcessorName, resourceModel.ProcessorName)
			}
		})
	}
}

func TestCleanupOnCreateTimeout(t *testing.T) {
	testCases := map[string]struct {
		callbackCtx       *callbackData
		mockSetup         func(*mockadmin.StreamsApi)
		expectedNoAPICall bool
	}{
		"deleteOnCreateTimeoutFalse": {
			callbackCtx: &callbackData{
				DeleteOnCreateTimeout:   false,
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
			},
			mockSetup:         func(m *mockadmin.StreamsApi) {},
			expectedNoAPICall: true,
		},
		"deleteOnCreateTimeoutTrue": {
			callbackCtx: &callbackData{
				DeleteOnCreateTimeout:   true,
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.DeleteStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().DeleteStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeleteStreamProcessorExecute(mock.Anything).Return(nil, nil)
			},
			expectedNoAPICall: false,
		},
		"deleteOnCreateTimeoutTrueWithError": {
			callbackCtx: &callbackData{
				DeleteOnCreateTimeout:   true,
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.DeleteStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().DeleteStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeleteStreamProcessorExecute(mock.Anything).Return(nil, fmt.Errorf("delete failed"))
			},
			expectedNoAPICall: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockStreamsApi

			err := cleanupOnCreateTimeout(context.Background(), mockClient, tc.callbackCtx)
			assert.NoError(t, err)
		})
	}
}

func TestSetup(t *testing.T) {
	assert.NotPanics(t, func() {
		setup()
	})
}

func TestGetStreamProcessor(t *testing.T) {
	testCases := map[string]struct {
		projectID               string
		workspaceOrInstanceName string
		processorName           string
		mockSetup               func(*mockadmin.StreamsApi)
		expectedError           bool
		expectedState           string
	}{
		"success": {
			projectID:               "507f1f77bcf86cd799439011",
			workspaceOrInstanceName: "workspace-1",
			processorName:           "processor-1",
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: "CREATED",
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedError: false,
			expectedState: "CREATED",
		},
		"notFound": {
			projectID:               "507f1f77bcf86cd799439011",
			workspaceOrInstanceName: "workspace-1",
			processorName:           "processor-1",
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedError: true,
		},
		"apiError": {
			projectID:               "507f1f77bcf86cd799439011",
			workspaceOrInstanceName: "workspace-1",
			processorName:           "processor-1",
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedError: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockStreamsApi

			processor, peErr := getStreamProcessor(context.Background(), mockClient, tc.projectID, tc.workspaceOrInstanceName, tc.processorName)

			if tc.expectedError {
				require.NotNil(t, peErr)
				assert.Nil(t, processor)
			} else {
				require.Nil(t, peErr)
				require.NotNil(t, processor)
				assert.Equal(t, tc.expectedState, processor.GetState())
			}
		})
	}
}

func TestStartStreamProcessor(t *testing.T) {
	testCases := map[string]struct {
		projectID               string
		workspaceOrInstanceName string
		processorName           string
		mockSetup               func(*mockadmin.StreamsApi)
		expectedError           bool
	}{
		"success": {
			projectID:               "507f1f77bcf86cd799439011",
			workspaceOrInstanceName: "workspace-1",
			processorName:           "processor-1",
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.StartStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().StartStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().StartStreamProcessorExecute(mock.Anything).Return(nil, nil)
			},
			expectedError: false,
		},
		"apiError": {
			projectID:               "507f1f77bcf86cd799439011",
			workspaceOrInstanceName: "workspace-1",
			processorName:           "processor-1",
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.StartStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().StartStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().StartStreamProcessorExecute(mock.Anything).Return(nil, fmt.Errorf("start failed"))
			},
			expectedError: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockStreamsApi

			peErr := startStreamProcessor(context.Background(), mockClient, tc.projectID, tc.workspaceOrInstanceName, tc.processorName)

			if tc.expectedError {
				require.NotNil(t, peErr)
			} else {
				require.Nil(t, peErr)
			}
		})
	}
}

func TestHandleCreateCallback(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *Model
		callbackCtx    *callbackData
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
	}{
		"createdStateNoStart": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           false,
				StartTime:               time.Now().Format(time.RFC3339),
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   false,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: CreatedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"createdStateNeedsStart": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           true,
				StartTime:               time.Now().Format(time.RFC3339),
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   false,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req1 := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req1).Once()
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: CreatedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil).Once()

				startReq := admin20250312010.StartStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().StartStreamProcessorWithParams(mock.Anything, mock.Anything).Return(startReq)
				m.EXPECT().StartStreamProcessorExecute(mock.Anything).Return(nil, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"startedState": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           false,
				StartTime:               time.Now().Format(time.RFC3339),
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   false,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: StartedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"failedState": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           false,
				StartTime:               time.Now().Format(time.RFC3339),
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   false,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: FailedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"initiatingState": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           false,
				StartTime:               time.Now().Format(time.RFC3339),
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   false,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: InitiatingState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"creatingState": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           false,
				StartTime:               time.Now().Format(time.RFC3339),
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   false,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: CreatingState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"timeoutExceeded": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           false,
				StartTime:               time.Now().Add(-25 * time.Minute).Format(time.RFC3339),
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   false,
			},
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Failed,
		},
		"timeoutExceededWithCleanup": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           false,
				StartTime:               time.Now().Add(-25 * time.Minute).Format(time.RFC3339),
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   true,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().DeleteStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(admin20250312010.DeleteStreamProcessorApiRequest{ApiService: m})
				m.EXPECT().DeleteStreamProcessorExecute(mock.Anything).Return(nil, nil)
			},
			expectedStatus: handler.Failed,
		},
		"timeoutWithCleanupError": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           false,
				StartTime:               time.Now().Add(-25 * time.Minute).Format(time.RFC3339),
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   true,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.DeleteStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().DeleteStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeleteStreamProcessorExecute(mock.Anything).Return(nil, fmt.Errorf("delete failed"))
			},
			expectedStatus: handler.Failed,
		},
		"unexpectedState": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           false,
				StartTime:               time.Now().Format(time.RFC3339),
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   false,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: "UNEXPECTED_STATE",
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"getStreamProcessorError": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           false,
				StartTime:               time.Now().Format(time.RFC3339),
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   false,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
		},
		"startProcessorError": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				NeedsStarting:           true,
				StartTime:               time.Now().Format(time.RFC3339),
				TimeoutDuration:         "20m",
				DeleteOnCreateTimeout:   false,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req1 := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req1).Once()
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: CreatedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil).Once()

				startReq := admin20250312010.StartStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().StartStreamProcessorWithParams(mock.Anything, mock.Anything).Return(startReq)
				m.EXPECT().StartStreamProcessorExecute(mock.Anything).Return(nil, fmt.Errorf("start failed"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockStreamsApi

			event, err := handleCreateCallback(context.Background(), mockClient, tc.currentModel, tc.callbackCtx)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}
func TestHandleUpdateCallback(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *Model
		callbackCtx    *callbackData
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
	}{
		"stoppedStateUpdate": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				PlannedState:            CreatedState,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req1 := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req1)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: StoppedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)

				updateReq := admin20250312010.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				updatedProcessor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: CreatedState,
				}
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(updatedProcessor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"startedStateNeedsStop": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				PlannedState:            StoppedState,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req1 := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req1)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: StartedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)

				stopReq := admin20250312010.StopStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().StopStreamProcessorWithParams(mock.Anything, mock.Anything).Return(stopReq)
				m.EXPECT().StopStreamProcessorExecute(mock.Anything).Return(nil, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"startedStateSamePlannedState": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				PlannedState:            StartedState,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: StartedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)

				// When state is already STARTED and planned state is STARTED, code still calls UpdateStreamProcessorWithParams
				// to allow updates to other fields (like Pipeline)
				updateReq := admin20250312010.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				updatedProcessor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: StartedState,
				}
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(updatedProcessor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"stoppedStateUpdateAndStart": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				PlannedState:            StartedState,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req1 := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req1)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: StoppedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)

				updateReq := admin20250312010.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				updatedProcessor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: StoppedState,
				}
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(updatedProcessor, &http.Response{StatusCode: 200}, nil)

				startReq := admin20250312010.StartStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().StartStreamProcessorWithParams(mock.Anything, mock.Anything).Return(startReq)
				m.EXPECT().StartStreamProcessorExecute(mock.Anything).Return(nil, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"failedState": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				PlannedState:            CreatedState,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: FailedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"defaultTransitioningState": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				PlannedState:            CreatedState,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: "STOPPING",
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"emptyPlannedState": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				PlannedState:            "",
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: CreatedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)

				updateReq := admin20250312010.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				updatedProcessor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: CreatedState,
				}
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(updatedProcessor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"updateError": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				PlannedState:            CreatedState,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req1 := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req1)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: StoppedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)

				updateReq := admin20250312010.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("update failed"))
			},
			expectedStatus: handler.Failed,
		},
		"stopError": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				PlannedState:            StoppedState,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req1 := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req1)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: StartedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)

				stopReq := admin20250312010.StopStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().StopStreamProcessorWithParams(mock.Anything, mock.Anything).Return(stopReq)
				m.EXPECT().StopStreamProcessorExecute(mock.Anything).Return(nil, fmt.Errorf("stop failed"))
			},
			expectedStatus: handler.Failed,
		},
		"startAfterUpdateError": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				PlannedState:            StartedState,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req1 := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req1)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: StoppedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)

				updateReq := admin20250312010.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				updatedProcessor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: StoppedState,
				}
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(updatedProcessor, &http.Response{StatusCode: 200}, nil)

				startReq := admin20250312010.StartStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().StartStreamProcessorWithParams(mock.Anything, mock.Anything).Return(startReq)
				m.EXPECT().StartStreamProcessorExecute(mock.Anything).Return(nil, fmt.Errorf("start failed"))
			},
			expectedStatus: handler.Failed,
		},
		"getStreamProcessorErrorInUpdate": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				PlannedState:            CreatedState,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("get failed"))
			},
			expectedStatus: handler.Failed,
		},
		"updateRequestError": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`invalid json`), // Invalid JSON will cause error
			},
			callbackCtx: &callbackData{
				ProjectID:               "507f1f77bcf86cd799439011",
				WorkspaceOrInstanceName: "workspace-1",
				ProcessorName:           "processor-1",
				PlannedState:            CreatedState,
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req1 := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req1)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: StoppedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockStreamsApi

			event, err := handleUpdateCallback(context.Background(), mockClient, tc.currentModel, tc.callbackCtx)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

// TestCreateValidationErrors tests validation paths in Create handler
func TestCreateValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *Model
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"missingProjectId": {
			currentModel: &Model{
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"missingProcessorName": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("workspace-1"),
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
			event, err := Create(req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

// TestReadValidationErrors tests validation paths in Read handler
func TestReadValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *Model
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"missingProjectId": {
			currentModel: &Model{
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"missingProcessorName": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("workspace-1"),
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

// TestUpdateValidationErrors tests validation paths in Update handler
func TestUpdateValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *Model
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"missingProjectId": {
			currentModel: &Model{
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"missingProcessorName": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("workspace-1"),
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
			event, err := Update(req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

// TestDeleteValidationErrors tests validation paths in Delete handler
func TestDeleteValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *Model
		expectedStatus handler.Status
		expectedMsg    string
	}{
		"missingProjectId": {
			currentModel: &Model{
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			expectedStatus: handler.Failed,
			expectedMsg:    "required",
		},
		"missingProcessorName": {
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("workspace-1"),
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

// TestCreateWithMocks tests the Create handler with mocked client initialization
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
		"successfulCreate": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
				State:         util.StringPtr(CreatedState),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.CreateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().CreateStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().CreateStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"createWithStateStarted": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
				State:         util.StringPtr(StartedState),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.CreateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().CreateStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().CreateStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"createWithCallback": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
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
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: CreatedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"createWithInvalidState": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
				State:         util.StringPtr("INVALID_STATE"),
			},
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Failed,
		},
		"createWithApiError": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.CreateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().CreateStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().CreateStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("API error"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Mock initEnvWithLatestClient
			mockStreamsApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockStreamsApi

			initEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := Create(tc.req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

// TestReadWithMocks tests the Read handler with mocked client initialization
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
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				pipeline := []any{
					map[string]any{"$match": map[string]any{"status": "active"}},
				}
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:     "processor-1",
					Id:       "507f1f77bcf86cd799439011",
					State:    CreatedState,
					Pipeline: pipeline,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"readNotFound": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Mock initEnvWithLatestClient
			mockStreamsApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockStreamsApi

			initEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := Read(tc.req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

// TestUpdateWithMocks tests the Update handler with mocked client initialization
func TestUpdateWithMocks(t *testing.T) {
	// Save original function
	originalInitEnv := initEnvWithLatestClient
	defer func() {
		initEnvWithLatestClient = originalInitEnv
	}()

	testCases := map[string]struct {
		req            handler.Request
		prevModel      *Model
		currentModel   *Model
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
	}{
		"successfulUpdate": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			prevModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				State:         util.StringPtr(CreatedState),
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
				State:         util.StringPtr(CreatedState),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				// Get current state
				req1 := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req1)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: CreatedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)

				// Update
				updateReq := admin20250312010.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				updatedProcessor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: CreatedState,
				}
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(updatedProcessor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"updateWithCallback": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
				CallbackContext: map[string]any{
					"callbackStreamProcessor": true,
					"projectID":               "507f1f77bcf86cd799439011",
					"workspaceName":           "workspace-1",
					"processorName":           "processor-1",
					"plannedState":            CreatedState,
				},
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.GetStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().GetStreamProcessorWithParams(mock.Anything, mock.Anything).Return(req)
				processor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: CreatedState,
				}
				m.EXPECT().GetStreamProcessorExecute(mock.Anything).Return(processor, &http.Response{StatusCode: 200}, nil)

				updateReq := admin20250312010.UpdateStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().UpdateStreamProcessorWithParams(mock.Anything, mock.Anything).Return(updateReq)
				updatedProcessor := &admin20250312010.StreamsProcessorWithStats{
					Name:  "processor-1",
					Id:    "507f1f77bcf86cd799439011",
					State: CreatedState,
				}
				m.EXPECT().UpdateStreamProcessorExecute(mock.Anything).Return(updatedProcessor, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Mock initEnvWithLatestClient
			mockStreamsApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockStreamsApi

			initEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := Update(tc.req, tc.prevModel, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

// TestDeleteWithMocks tests the Delete handler with mocked client initialization
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
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.DeleteStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().DeleteStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeleteStreamProcessorExecute(mock.Anything).Return(nil, nil)
			},
			expectedStatus: handler.Success,
		},
		"deleteWithError": {
			req: handler.Request{
				RequestContext: handler.RequestContext{},
			},
			currentModel: &Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin20250312010.DeleteStreamProcessorApiRequest{ApiService: m}
				m.EXPECT().DeleteStreamProcessor(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().DeleteStreamProcessorExecute(mock.Anything).Return(nil, fmt.Errorf("delete failed"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Mock initEnvWithLatestClient
			mockStreamsApi := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsApi)

			mockClient := &admin20250312010.APIClient{}
			mockClient.StreamsApi = mockStreamsApi

			initEnvWithLatestClient = func(req handler.Request, currentModel *Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := Delete(tc.req, nil, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}
