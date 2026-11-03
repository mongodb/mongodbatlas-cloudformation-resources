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
