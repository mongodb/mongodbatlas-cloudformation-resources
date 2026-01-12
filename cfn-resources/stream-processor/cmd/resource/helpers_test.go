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
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-processor/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312010/admin"
)

func TestCopyIdentifyingFields(t *testing.T) {
	testCases := map[string]struct {
		resourceModel *resource.Model
		currentModel  *resource.Model
		validateFunc  func(t *testing.T, resourceModel *resource.Model)
	}{
		"withWorkspaceName": {
			resourceModel: &resource.Model{},
			currentModel: &resource.Model{
				Profile:       util.StringPtr("default"),
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			validateFunc: func(t *testing.T, rm *resource.Model) {
				t.Helper()
				assert.Equal(t, "default", util.SafeString(rm.Profile))
				assert.Equal(t, "507f1f77bcf86cd799439011", util.SafeString(rm.ProjectId))
				assert.Equal(t, "processor-1", util.SafeString(rm.ProcessorName))
				assert.Equal(t, "workspace-1", util.SafeString(rm.WorkspaceName))
				assert.Equal(t, "workspace-1", util.SafeString(rm.InstanceName))
			},
		},
		"withInstanceName": {
			resourceModel: &resource.Model{},
			currentModel: &resource.Model{
				Profile:       util.StringPtr("default"),
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				InstanceName:  util.StringPtr("instance-1"),
			},
			validateFunc: func(t *testing.T, rm *resource.Model) {
				t.Helper()
				assert.Equal(t, "default", util.SafeString(rm.Profile))
				assert.Equal(t, "507f1f77bcf86cd799439011", util.SafeString(rm.ProjectId))
				assert.Equal(t, "processor-1", util.SafeString(rm.ProcessorName))
				assert.Equal(t, "instance-1", util.SafeString(rm.InstanceName))
				assert.Equal(t, "instance-1", util.SafeString(rm.WorkspaceName))
			},
		},
		"emptyWorkspaceName": {
			resourceModel: &resource.Model{},
			currentModel: &resource.Model{
				Profile:       util.StringPtr("default"),
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr(""),
			},
			validateFunc: func(t *testing.T, rm *resource.Model) {
				t.Helper()
				assert.Nil(t, rm.WorkspaceName)
			},
		},
		"bothNil": {
			resourceModel: &resource.Model{},
			currentModel: &resource.Model{
				Profile:       util.StringPtr("default"),
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: nil,
				InstanceName:  nil,
			},
			validateFunc: func(t *testing.T, rm *resource.Model) {
				t.Helper()
				assert.Nil(t, rm.WorkspaceName)
				assert.Nil(t, rm.InstanceName)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			resource.CopyIdentifyingFields(tc.resourceModel, tc.currentModel)
			if tc.validateFunc != nil {
				tc.validateFunc(t, tc.resourceModel)
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
			expectedResult: resource.DefaultCreateTimeout,
		},
		"invalidFormat": {
			timeoutStr:     "invalid",
			expectedResult: resource.DefaultCreateTimeout,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := resource.ParseTimeout(tc.timeoutStr)
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
			result := resource.IsTimeoutExceeded(tc.startTimeStr, tc.timeoutDurationStr)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestValidateUpdateStateTransition(t *testing.T) {
	testCases := map[string]struct {
		currentState    string
		desiredState    string
		expectedErrMsg  string
		expectedIsValid bool
	}{
		"validCREATEDtoSTARTED": {
			currentState:    resource.CreatedState,
			desiredState:    resource.StartedState,
			expectedIsValid: true,
		},
		"invalidSTARTEDtoCREATED": {
			currentState:    resource.StartedState,
			desiredState:    resource.CreatedState,
			expectedIsValid: false,
			expectedErrMsg:  "cannot transition from STARTED to CREATED",
		},
		"validSTARTEDtoSTOPPED": {
			currentState:    resource.StartedState,
			desiredState:    resource.StoppedState,
			expectedIsValid: true,
		},
		"invalidCREATEDtoSTOPPED": {
			currentState:    resource.CreatedState,
			desiredState:    resource.StoppedState,
			expectedIsValid: false,
			expectedErrMsg:  "must be in STARTED state",
		},
		"sameState": {
			currentState:    resource.CreatedState,
			desiredState:    resource.CreatedState,
			expectedIsValid: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			errMsg, isValid := resource.ValidateUpdateStateTransition(tc.currentState, tc.desiredState)
			assert.Equal(t, tc.expectedIsValid, isValid)
			if !tc.expectedIsValid {
				assert.Contains(t, errMsg, tc.expectedErrMsg)
			}
		})
	}
}

func TestHandleError(t *testing.T) {
	testCases := map[string]struct {
		response        *http.Response
		method          constants.CfnFunctions
		err             error
		expectedStatus  handler.Status
		expectedErrCode string
		expectedMsg     string
	}{
		"conflictError": {
			response: &http.Response{
				StatusCode: http.StatusConflict,
			},
			method:          constants.CREATE,
			err:             errors.New("already exists"),
			expectedStatus:  handler.Failed,
			expectedErrCode: "AlreadyExists",
			expectedMsg:     "CREATE error:already exists",
		},
		"genericError": {
			response: &http.Response{
				StatusCode: http.StatusBadRequest,
			},
			method:         constants.UPDATE,
			err:            errors.New("bad request"),
			expectedStatus: handler.Failed,
			expectedMsg:    "UPDATE error:bad request",
		},
		"nilResponse": {
			response:       nil,
			method:         constants.DELETE,
			err:            errors.New("network error"),
			expectedStatus: handler.Failed,
			expectedMsg:    "DELETE error:network error",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.HandleError(tc.response, tc.method, tc.err)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
			if tc.expectedErrCode != "" {
				assert.Equal(t, tc.expectedErrCode, event.HandlerErrorCode)
			}
		})
	}
}

func TestFinalizeModel(t *testing.T) {
	testCases := map[string]struct {
		streamProcessor *admin.StreamsProcessorWithStats
		currentModel    *resource.Model
		validateFunc    func(t *testing.T, event handler.ProgressEvent)
		message         string
		expectedStatus  handler.Status
		expectedMsg     string
	}{
		"success": {
			streamProcessor: &admin.StreamsProcessorWithStats{
				Name:  "processor-1",
				Id:    "507f1f77bcf86cd799439011",
				State: resource.CreatedState,
			},
			currentModel: &resource.Model{
				Profile:       util.StringPtr("default"),
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("processor-1"),
				WorkspaceName: util.StringPtr("workspace-1"),
			},
			message:        "Create Complete",
			expectedStatus: handler.Success,
			expectedMsg:    "Create Complete",
			validateFunc: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				model, ok := event.ResourceModel.(*resource.Model)
				require.True(t, ok, "ResourceModel should be *resource.Model")
				assert.Equal(t, "processor-1", util.SafeString(model.ProcessorName))
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.FinalizeModel(tc.streamProcessor, tc.currentModel, tc.message)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Equal(t, tc.expectedMsg, event.Message)
			if tc.validateFunc != nil {
				tc.validateFunc(t, event)
			}
		})
	}
}
