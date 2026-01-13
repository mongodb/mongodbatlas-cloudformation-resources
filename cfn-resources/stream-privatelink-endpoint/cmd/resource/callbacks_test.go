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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-privatelink-endpoint/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312010/admin"
	"go.mongodb.org/atlas-sdk/v20250312010/mockadmin"
)

func TestIsCallback(t *testing.T) {
	testCases := map[string]struct {
		req            *handler.Request
		expectedResult bool
	}{
		"isCallback": {
			req: &handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamPrivatelinkEndpoint": true,
				},
			},
			expectedResult: true,
		},
		"notCallback": {
			req: &handler.Request{
				CallbackContext: map[string]any{},
			},
			expectedResult: false,
		},
		"nilCallbackContext": {
			req: &handler.Request{
				CallbackContext: nil,
			},
			expectedResult: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := resource.IsCallback(tc.req)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestBuildCallbackContext(t *testing.T) {
	testCases := map[string]struct {
		validateFunc func(t *testing.T, ctx map[string]interface{})
		projectID    string
		connectionID string
	}{
		"basic": {
			projectID:    "507f1f77bcf86cd799439011",
			connectionID: "507f1f77bcf86cd799439012",
			validateFunc: func(t *testing.T, ctx map[string]interface{}) {
				t.Helper()
				assert.True(t, ctx["callbackStreamPrivatelinkEndpoint"].(bool))
				assert.Equal(t, "507f1f77bcf86cd799439011", ctx["projectID"])
				assert.Equal(t, "507f1f77bcf86cd799439012", ctx["connectionID"])
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := resource.BuildCallbackContext(tc.projectID, tc.connectionID)
			if tc.validateFunc != nil {
				tc.validateFunc(t, ctx)
			}
		})
	}
}

func TestHandleCreateCallback(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *resource.Model
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
		req            handler.Request
	}{
		"doneState": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamPrivatelinkEndpoint": true,
					"projectID":                         "507f1f77bcf86cd799439011",
					"connectionID":                      "507f1f77bcf86cd799439012",
				},
			},
			currentModel: &resource.Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
				Id:        util.StringPtr("507f1f77bcf86cd799439012"),
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
		"failedState": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamPrivatelinkEndpoint": true,
					"projectID":                         "507f1f77bcf86cd799439011",
					"connectionID":                      "507f1f77bcf86cd799439012",
				},
			},
			currentModel: &resource.Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
				Id:        util.StringPtr("507f1f77bcf86cd799439012"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				failedState := resource.StateFailed
				errorMsg := "Connection failed"
				apiResp.State = &failedState
				apiResp.ErrorMessage = &errorMsg
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"inProgressState": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamPrivatelinkEndpoint": true,
					"projectID":                         "507f1f77bcf86cd799439011",
					"connectionID":                      "507f1f77bcf86cd799439012",
				},
			},
			currentModel: &resource.Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
				Id:        util.StringPtr("507f1f77bcf86cd799439012"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				workingState := resource.StateWorking
				apiResp.State = &workingState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"missingConnectionID": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamPrivatelinkEndpoint": true,
					"projectID":                         "507f1f77bcf86cd799439011",
				},
			},
			currentModel: &resource.Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
			},
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsAPI := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsAPI)

			mockClient := &admin.APIClient{}
			mockClient.StreamsApi = mockStreamsAPI

			event, err := resource.HandleCreateCallback(context.Background(), mockClient, tc.req, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

func TestHandleDeleteCallback(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *resource.Model
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
		req            handler.Request
	}{
		"deletedState": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamPrivatelinkEndpoint": true,
					"projectID":                         "507f1f77bcf86cd799439011",
					"connectionID":                      "507f1f77bcf86cd799439012",
				},
			},
			currentModel: &resource.Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
				Id:        util.StringPtr("507f1f77bcf86cd799439012"),
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
		"inProgressState": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamPrivatelinkEndpoint": true,
					"projectID":                         "507f1f77bcf86cd799439011",
					"connectionID":                      "507f1f77bcf86cd799439012",
				},
			},
			currentModel: &resource.Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
				Id:        util.StringPtr("507f1f77bcf86cd799439012"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				deletingState := resource.StateDeleting
				apiResp.State = &deletingState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"notFoundSuccess": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamPrivatelinkEndpoint": true,
					"projectID":                         "507f1f77bcf86cd799439011",
					"connectionID":                      "507f1f77bcf86cd799439012",
				},
			},
			currentModel: &resource.Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
				Id:        util.StringPtr("507f1f77bcf86cd799439012"),
			},
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Success,
		},
		"missingIDs": {
			req: handler.Request{
				CallbackContext: map[string]any{
					"callbackStreamPrivatelinkEndpoint": true,
				},
			},
			currentModel:   &resource.Model{},
			mockSetup:      func(m *mockadmin.StreamsApi) {},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsAPI := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsAPI)

			mockClient := &admin.APIClient{}
			mockClient.StreamsApi = mockStreamsAPI

			event, err := resource.HandleDeleteCallback(context.Background(), mockClient, tc.req, tc.currentModel)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}
