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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/organization/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
)

// A delete that outlives the synchronous Atlas call returns an in-progress
// callback together with a nil responseMsg. DeleteOutcome must return that
// callback instead of dereferencing the nil responseMsg, which panicked before.
func TestDeleteOutcomeReturnsTheInProgressCallbackWhenTheDeleteTimesOut(t *testing.T) {
	progress := handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              resource.DeleteInProgress,
		CallbackDelaySeconds: 20,
	}

	event, err := resource.DeleteOutcome(nil, &progress)

	require.NoError(t, err)
	assert.Equal(t, handler.InProgress, event.OperationStatus)
	assert.Equal(t, resource.DeleteInProgress, event.Message)
	assert.Positive(t, event.CallbackDelaySeconds)
}

func TestDeleteOutcomeCompletesWhenTheOrganizationIsGone(t *testing.T) {
	event, err := resource.DeleteOutcome(&resource.DeleteResponse{}, nil)

	require.NoError(t, err)
	assert.Equal(t, handler.Success, event.OperationStatus)
	assert.Equal(t, resource.DeleteCompleted, event.Message)
}

func TestDeleteOutcomeTreatsUnauthorizedAsCompleted(t *testing.T) {
	response := &http.Response{StatusCode: http.StatusUnauthorized}

	event, err := resource.DeleteOutcome(&resource.DeleteResponse{Error: errors.New("Unauthorized"), Response: response}, nil)

	require.NoError(t, err)
	assert.Equal(t, handler.Success, event.OperationStatus)
	assert.Equal(t, resource.DeleteCompleted, event.Message)
}

func TestDeleteOutcomeMapsDeleteErrors(t *testing.T) {
	testCases := []struct {
		name     string
		response *http.Response
		want     string
	}{
		{name: "Conflict", response: &http.Response{StatusCode: http.StatusConflict}, want: string(types.HandlerErrorCodeAlreadyExists)},
		{name: "Bad request", response: &http.Response{StatusCode: http.StatusBadRequest}, want: string(types.HandlerErrorCodeNotFound)},
		{name: "Server error", response: &http.Response{StatusCode: http.StatusInternalServerError}, want: string(types.HandlerErrorCodeServiceInternalError)},
		{name: "No response", response: nil, want: string(types.HandlerErrorCodeInternalFailure)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			event, err := resource.DeleteOutcome(&resource.DeleteResponse{Error: errors.New("delete failed"), Response: tc.response}, nil)

			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Equal(t, tc.want, event.HandlerErrorCode)
		})
	}
}

// A delete that outlives the synchronous Atlas call returns an in-progress
// callback with a nil response. The retry decision must return that callback
// without dereferencing the nil response, which panicked before.
func TestDeleteWithRetriesReturnsTheInProgressCallbackWithoutRetrying(t *testing.T) {
	progress := &handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              resource.DeleteInProgress,
		CallbackDelaySeconds: 20,
	}
	calls := 0

	responseMsg, progressEvent := resource.DeleteWithRetries(func() (*resource.DeleteResponse, *handler.ProgressEvent) {
		calls++
		return nil, progress
	}, 0)

	assert.Equal(t, 1, calls, "an in-progress callback is not retried")
	assert.Nil(t, responseMsg)
	require.NotNil(t, progressEvent)
	assert.Equal(t, handler.InProgress, progressEvent.OperationStatus)
	assert.Equal(t, resource.DeleteInProgress, progressEvent.Message)
}

func TestDeleteWithRetriesRetriesTransientServerErrors(t *testing.T) {
	calls := 0

	responseMsg, progressEvent := resource.DeleteWithRetries(func() (*resource.DeleteResponse, *handler.ProgressEvent) {
		calls++
		if calls < 3 {
			return &resource.DeleteResponse{Error: errors.New("Internal Server Error"), Response: &http.Response{StatusCode: http.StatusInternalServerError}}, nil
		}
		return &resource.DeleteResponse{}, nil
	}, 0)

	assert.Equal(t, 3, calls, "transient errors are retried until the delete succeeds")
	require.NotNil(t, responseMsg)
	require.NoError(t, responseMsg.Error)
	assert.Nil(t, progressEvent)
}

func TestDeleteWithRetriesStopsAtTheAttemptLimit(t *testing.T) {
	calls := 0

	responseMsg, progressEvent := resource.DeleteWithRetries(func() (*resource.DeleteResponse, *handler.ProgressEvent) {
		calls++
		return &resource.DeleteResponse{Error: errors.New("Internal Server Error"), Response: &http.Response{StatusCode: http.StatusInternalServerError}}, nil
	}, 0)

	assert.Equal(t, 3, calls, "the delete is attempted deleteRetryAttempts times")
	require.NotNil(t, responseMsg)
	require.Error(t, responseMsg.Error)
	assert.Nil(t, progressEvent)
}

func TestDeleteWithRetriesDoesNotRetryOtherErrors(t *testing.T) {
	calls := 0

	responseMsg, _ := resource.DeleteWithRetries(func() (*resource.DeleteResponse, *handler.ProgressEvent) {
		calls++
		return &resource.DeleteResponse{Error: errors.New("Unauthorized"), Response: &http.Response{StatusCode: http.StatusUnauthorized}}, nil
	}, 0)

	assert.Equal(t, 1, calls, "only transient server errors are retried")
	require.NotNil(t, responseMsg)
	require.Error(t, responseMsg.Error)
}

// A transient server error on the callback path must keep the in-progress
// callback alive so CloudFormation invokes the handler again and the delete is
// retried, instead of failing the stack deletion.
func TestDeleteCallbackOutcomeRetriesATransientErrorThroughTheCallback(t *testing.T) {
	responseMsg := &resource.DeleteResponse{
		Error:    errors.New("Internal Server Error"),
		Response: &http.Response{StatusCode: http.StatusInternalServerError},
	}
	model := &resource.Model{}

	event, err := resource.DeleteCallbackOutcome(responseMsg, model)

	require.NoError(t, err)
	assert.Equal(t, handler.InProgress, event.OperationStatus)
	assert.Equal(t, resource.DeleteInProgress, event.Message)
	assert.Same(t, model, event.ResourceModel)
	assert.Equal(t, int64(resource.CallBackSeconds), event.CallbackDelaySeconds)
	assert.Equal(t, resource.DeletingState, event.CallbackContext[constants.StateName])
}

func TestDeleteCallbackOutcomeCompletesWhenTheDeleteSucceeds(t *testing.T) {
	event, err := resource.DeleteCallbackOutcome(&resource.DeleteResponse{}, &resource.Model{})

	require.NoError(t, err)
	assert.Equal(t, handler.Success, event.OperationStatus)
	assert.Equal(t, resource.DeleteCompleted, event.Message)
}

func TestDeleteCallbackOutcomeFailsOnANonTransientError(t *testing.T) {
	responseMsg := &resource.DeleteResponse{
		Error:    errors.New("Bad Request"),
		Response: &http.Response{StatusCode: http.StatusBadRequest},
	}

	event, err := resource.DeleteCallbackOutcome(responseMsg, &resource.Model{})

	require.NoError(t, err)
	assert.Equal(t, handler.Failed, event.OperationStatus)
}
