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

func TestValidateVendorRequirements(t *testing.T) {
	testCases := map[string]struct {
		model         *resource.Model
		expectedMsg   string
		expectedError bool
	}{
		"nilVendor": {
			model: &resource.Model{
				ProviderName: util.StringPtr("AWS"),
			},
			expectedError: false,
		},
		"validMSK": {
			model: &resource.Model{
				Vendor: util.StringPtr("MSK"),
				Arn:    util.StringPtr("arn:aws:kafka:us-east-1:123456789012:cluster/test"),
			},
			expectedError: false,
		},
		"validS3": {
			model: &resource.Model{
				Vendor:            util.StringPtr("S3"),
				Region:            util.StringPtr("us-east-1"),
				ServiceEndpointId: util.StringPtr("com.amazonaws.us-east-1.s3"),
			},
			expectedError: false,
		},
		"validConfluent": {
			model: &resource.Model{
				Vendor:            util.StringPtr("CONFLUENT"),
				Region:            util.StringPtr("us-east-1"),
				DnsDomain:         util.StringPtr("test.example.com"),
				ServiceEndpointId: util.StringPtr("com.amazonaws.vpce.us-east-1.vpce-svc-12345678"),
			},
			expectedError: false,
		},
		"confluentMissingDnsDomain": {
			model: &resource.Model{
				Vendor:            util.StringPtr("CONFLUENT"),
				Region:            util.StringPtr("us-east-1"),
				ServiceEndpointId: util.StringPtr("com.amazonaws.vpce.us-east-1.vpce-svc-12345678"),
			},
			expectedError: true,
			expectedMsg:   "DnsDomain is required",
		},
		"confluentMissingRegion": {
			model: &resource.Model{
				Vendor:            util.StringPtr("CONFLUENT"),
				DnsDomain:         util.StringPtr("test.example.com"),
				ServiceEndpointId: util.StringPtr("com.amazonaws.vpce.us-east-1.vpce-svc-12345678"),
			},
			expectedError: true,
			expectedMsg:   "Region is required",
		},
		"confluentMissingServiceEndpointId": {
			model: &resource.Model{
				Vendor:    util.StringPtr("CONFLUENT"),
				DnsDomain: util.StringPtr("test.example.com"),
				Region:    util.StringPtr("us-east-1"),
			},
			expectedError: true,
			expectedMsg:   "ServiceEndpointId is required",
		},
		"mskMissingArn": {
			model: &resource.Model{
				Vendor: util.StringPtr("MSK"),
			},
			expectedError: true,
			expectedMsg:   "Arn is required",
		},
		"mskWithRegion": {
			model: &resource.Model{
				Vendor: util.StringPtr("MSK"),
				Arn:    util.StringPtr("arn:aws:kafka:us-east-1:123456789012:cluster/test"),
				Region: util.StringPtr("us-east-1"),
			},
			expectedError: true,
			expectedMsg:   "Region cannot be set for vendor MSK",
		},
		"s3MissingRegion": {
			model: &resource.Model{
				Vendor:            util.StringPtr("S3"),
				ServiceEndpointId: util.StringPtr("com.amazonaws.us-east-1.s3"),
			},
			expectedError: true,
			expectedMsg:   "Region is required",
		},
		"s3MissingServiceEndpointId": {
			model: &resource.Model{
				Vendor: util.StringPtr("S3"),
				Region: util.StringPtr("us-east-1"),
			},
			expectedError: true,
			expectedMsg:   "ServiceEndpointId is required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := resource.ValidateVendorRequirements(tc.model)
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

func TestWaitForStateTransition(t *testing.T) {
	testCases := map[string]struct {
		currentModel   *resource.Model
		mockSetup      func(*mockadmin.StreamsApi)
		expectedStatus handler.Status
		pendingStates  []string
		targetStates   []string
		isDelete       bool
	}{
		"reachedTargetStateDone": {
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{resource.StateIdle, resource.StateWorking},
			targetStates:  []string{resource.StateDone, resource.StateFailed},
			isDelete:      false,
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
		"reachedTargetStateFailed": {
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{resource.StateIdle, resource.StateWorking},
			targetStates:  []string{resource.StateDone, resource.StateFailed},
			isDelete:      false,
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
		"inPendingState": {
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{resource.StateIdle, resource.StateWorking},
			targetStates:  []string{resource.StateDone, resource.StateFailed},
			isDelete:      false,
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
		"deleteReachedDeletedState": {
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{resource.StateDeleteRequested, resource.StateDeleting},
			targetStates:  []string{resource.StateDeleted, resource.StateFailed},
			isDelete:      true,
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
		"deleteNotFound": {
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{resource.StateDeleteRequested, resource.StateDeleting},
			targetStates:  []string{resource.StateDeleted, resource.StateFailed},
			isDelete:      true,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Success,
		},
		"emptyState": {
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{resource.StateIdle, resource.StateWorking},
			targetStates:  []string{resource.StateDone},
			isDelete:      false,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				apiResp.State = nil
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"unexpectedState": {
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{resource.StateIdle},
			targetStates:  []string{resource.StateDone},
			isDelete:      false,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				unexpectedState := "UNKNOWN"
				apiResp.State = &unexpectedState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"deleteUnexpectedState": {
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{resource.StateDeleteRequested},
			targetStates:  []string{resource.StateDeleted},
			isDelete:      true,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				apiResp := createTestAPIResponse()
				unexpectedState := "UNKNOWN"
				apiResp.State = &unexpectedState
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(apiResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"nonDeleteError": {
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{resource.StateIdle},
			targetStates:  []string{resource.StateDone},
			isDelete:      false,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
		},
		"deleteErrorWithResponse": {
			currentModel: func() *resource.Model {
				m := createTestModel()
				id := "507f1f77bcf86cd799439012"
				m.Id = &id
				return m
			}(),
			pendingStates: []string{resource.StateDeleteRequested},
			targetStates:  []string{resource.StateDeleted},
			isDelete:      true,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.GetPrivateLinkConnectionApiRequest{ApiService: m}
				m.EXPECT().GetPrivateLinkConnection(mock.Anything, mock.Anything, mock.Anything).Return(req)
				m.EXPECT().GetPrivateLinkConnectionExecute(mock.Anything).Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("server error"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockStreamsAPI := mockadmin.NewStreamsApi(t)
			tc.mockSetup(mockStreamsAPI)

			mockClient := &admin.APIClient{}
			mockClient.StreamsApi = mockStreamsAPI

			event, err := resource.WaitForStateTransition(context.Background(), mockClient, tc.currentModel, tc.pendingStates, tc.targetStates, tc.isDelete)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
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
		"alreadyExistsError": {
			response: &http.Response{
				StatusCode: http.StatusBadRequest,
			},
			method:             constants.CREATE,
			err:                fmt.Errorf("STREAM_PRIVATE_LINK_ALREADY_EXISTS: resource already exists"),
			expectedStatus:     handler.Failed,
			expectedErrorCode:  string(types.HandlerErrorCodeAlreadyExists),
			expectedMsgContain: "CREATE error",
		},
		"otherError": {
			response: &http.Response{
				StatusCode: http.StatusBadRequest,
			},
			method:             constants.CREATE,
			err:                fmt.Errorf("invalid request"),
			expectedStatus:     handler.Failed,
			expectedErrorCode:  "",
			expectedMsgContain: "CREATE error",
		},
		"nilResponse": {
			response:          nil,
			method:            constants.CREATE,
			err:               fmt.Errorf("network error"),
			expectedStatus:    handler.Failed,
			expectedErrorCode: "",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.HandleError(tc.response, tc.method, tc.err)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.expectedErrorCode != "" {
				assert.Equal(t, tc.expectedErrorCode, event.HandlerErrorCode)
			}
			if tc.expectedMsgContain != "" {
				assert.Contains(t, event.Message, tc.expectedMsgContain)
			}
		})
	}
}
