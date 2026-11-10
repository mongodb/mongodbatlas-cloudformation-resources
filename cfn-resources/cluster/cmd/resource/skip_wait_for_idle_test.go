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
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/cluster/cmd/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldSkipWaitForIdle(t *testing.T) {
	testCases := map[string]struct {
		model    *resource.Model
		expected bool
	}{
		"absentMeansWait": {model: &resource.Model{}, expected: false},
		"trueSkipsTheWait": {
			model:    &resource.Model{SkipWaitForIdle: new(true)},
			expected: true,
		},
		"falseWaits": {
			model:    &resource.Model{SkipWaitForIdle: new(false)},
			expected: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, resource.ShouldSkipWaitForIdle(tc.model))
		})
	}
}

func TestRejectPauseWithSkipWaitForIdle(t *testing.T) {
	testCases := map[string]struct {
		paused          *bool
		skipWaitForIdle bool
		currentlyPaused bool
		expectRejection bool
	}{
		"pauseIsRejected": {
			paused:          new(true),
			skipWaitForIdle: true,
			currentlyPaused: false,
			expectRejection: true,
		},
		"pauseOfAnAlreadyPausedClusterIsAllowed": {
			paused:          new(true),
			skipWaitForIdle: true,
			currentlyPaused: true,
			expectRejection: false,
		},
		"unpauseIsAllowed": {
			paused:          new(false),
			skipWaitForIdle: true,
			currentlyPaused: true,
			expectRejection: false,
		},
		"absentPausedIsAllowed": {
			paused:          nil,
			skipWaitForIdle: true,
			currentlyPaused: false,
			expectRejection: false,
		},
		"pauseWithoutSkipWaitForIdleIsAllowed": {
			paused:          new(true),
			skipWaitForIdle: false,
			currentlyPaused: false,
			expectRejection: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			model := &resource.Model{Name: new("cluster"), Paused: tc.paused}
			if tc.skipWaitForIdle {
				model.SkipWaitForIdle = new(true)
			}

			pe := resource.RejectPauseWithSkipWaitForIdle(model, tc.currentlyPaused)

			if !tc.expectRejection {
				assert.Nil(t, pe)
				return
			}
			require.NotNil(t, pe)
			assert.Equal(t, handler.Failed, pe.OperationStatus)
			assert.Equal(t, string(types.HandlerErrorCodeInvalidRequest), pe.HandlerErrorCode)
			assert.Contains(t, pe.Message, "cluster")
			assert.Contains(t, pe.Message, "separate update")
		})
	}
}
