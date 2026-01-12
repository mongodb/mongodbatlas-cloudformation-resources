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
	"encoding/json"
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-processor/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312010/admin"
)

func assertJSONEqual(t *testing.T, expected, actual string) {
	t.Helper()
	var expectedJSON, actualJSON any
	require.NoError(t, json.Unmarshal([]byte(expected), &expectedJSON))
	require.NoError(t, json.Unmarshal([]byte(actual), &actualJSON))
	assert.Equal(t, expectedJSON, actualJSON)
}

func TestGetWorkspaceOrInstanceName(t *testing.T) {
	testCases := map[string]struct {
		model          *resource.Model
		expectedResult string
		expectedError  string
	}{
		"workspaceName": {
			model:          &resource.Model{WorkspaceName: util.StringPtr("workspace-1")},
			expectedResult: "workspace-1",
		},
		"instanceName": {
			model:          &resource.Model{InstanceName: util.StringPtr("instance-1")},
			expectedResult: "instance-1",
		},
		"neitherSet": {
			model:         &resource.Model{},
			expectedError: "either WorkspaceName or InstanceName must be provided",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := resource.GetWorkspaceOrInstanceName(tc.model)
			if tc.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}

func TestConvertPipelineToSdk(t *testing.T) {
	testCases := map[string]struct {
		pipeline      string
		expectedError bool
	}{
		"validPipeline": {
			pipeline: `[{"$match": {"status": "active"}}]`,
		},
		"invalidJSON": {
			pipeline:      `[{"$match": {"status": "active"}`,
			expectedError: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := resource.ConvertPipelineToSdk(tc.pipeline)
			if tc.expectedError {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestConvertPipelineToString(t *testing.T) {
	testCases := map[string]struct {
		expectedJSON string
		pipeline     []any
	}{
		"validPipeline": {
			pipeline:     []any{map[string]any{"$match": map[string]any{"status": "active"}}},
			expectedJSON: `[{"$match":{"status":"active"}}]`,
		},
		"nilPipeline": {
			pipeline:     nil,
			expectedJSON: `null`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := resource.ConvertPipelineToString(tc.pipeline)
			require.NoError(t, err)
			assertJSONEqual(t, tc.expectedJSON, result)
		})
	}
}

func TestConvertStatsToString(t *testing.T) {
	result, err := resource.ConvertStatsToString(map[string]any{"bytesProcessed": 1000})
	require.NoError(t, err)
	assertJSONEqual(t, `{"bytesProcessed":1000}`, result)

	result, err = resource.ConvertStatsToString(nil)
	require.NoError(t, err)
	assert.Empty(t, result)
}

func TestNewStreamProcessorReq(t *testing.T) {
	validPipeline := `[{"$match": {"status": "active"}}]`
	validDLQ := &resource.StreamsOptions{
		Dlq: &resource.StreamsDLQ{
			Coll:           util.StringPtr("dlq-collection"),
			ConnectionName: util.StringPtr("dlq-connection"),
			Db:             util.StringPtr("dlq-db"),
		},
	}

	testCases := map[string]struct {
		model         *resource.Model
		expectedError bool
		checkOptions  bool
	}{
		"minimalRequest": {
			model: &resource.Model{
				ProcessorName: util.StringPtr("test-processor"),
				Pipeline:      util.StringPtr(validPipeline),
			},
		},
		"withOptions": {
			model: &resource.Model{
				ProcessorName: util.StringPtr("test-processor"),
				Pipeline:      util.StringPtr(validPipeline),
				Options:       validDLQ,
			},
			checkOptions: true,
		},
		"invalidPipeline": {
			model: &resource.Model{
				ProcessorName: util.StringPtr("test-processor"),
				Pipeline:      util.StringPtr(`invalid json`),
			},
			expectedError: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := resource.NewStreamProcessorReq(tc.model)
			if tc.expectedError {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, "test-processor", result.GetName())
				if tc.checkOptions {
					require.NotNil(t, result.Options.Dlq)
					assert.Equal(t, "dlq-collection", util.SafeString(result.Options.Dlq.Coll))
				}
			}
		})
	}
}

func TestNewStreamProcessorUpdateReq(t *testing.T) {
	validPipeline := `[{"$match": {"status": "active"}}]`
	validDLQ := &resource.StreamsOptions{
		Dlq: &resource.StreamsDLQ{
			Coll:           util.StringPtr("dlq-collection"),
			ConnectionName: util.StringPtr("dlq-connection"),
			Db:             util.StringPtr("dlq-db"),
		},
	}

	testCases := map[string]struct {
		model         *resource.Model
		checkTenant   string
		expectedError bool
		checkOptions  bool
	}{
		"minimalRequest": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(validPipeline),
			},
			checkTenant: "workspace-1",
		},
		"withOptions": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(validPipeline),
				Options:       validDLQ,
			},
			checkTenant:  "workspace-1",
			checkOptions: true,
		},
		"invalidPipeline": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`invalid json`),
			},
			expectedError: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := resource.NewStreamProcessorUpdateReq(tc.model)
			if tc.expectedError {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, "507f1f77bcf86cd799439011", result.GroupId)
				assert.Equal(t, "test-processor", result.ProcessorName)
				if tc.checkTenant != "" {
					assert.Equal(t, tc.checkTenant, result.TenantName)
				}
				if tc.checkOptions {
					require.NotNil(t, result.StreamsModifyStreamProcessor.Options.Dlq)
					assert.Equal(t, "dlq-collection", result.StreamsModifyStreamProcessor.Options.Dlq.GetColl())
				}
			}
		})
	}
}

func TestGetStreamProcessorModel(t *testing.T) {
	validDLQ := &admin.StreamsDLQ{
		Coll:           admin.PtrString("dlq-collection"),
		ConnectionName: admin.PtrString("dlq-connection"),
		Db:             admin.PtrString("dlq-db"),
	}
	currentModelWithDLQ := &resource.Model{
		Options: &resource.StreamsOptions{
			Dlq: &resource.StreamsDLQ{
				Coll:           util.StringPtr("existing-dlq-collection"),
				ConnectionName: util.StringPtr("existing-dlq-connection"),
				Db:             util.StringPtr("existing-dlq-db"),
			},
		},
	}

	testCases := map[string]struct {
		streamProcessor *admin.StreamsProcessorWithStats
		currentModel    *resource.Model
		checkFields     []string
	}{
		"minimalConversion": {
			streamProcessor: &admin.StreamsProcessorWithStats{
				Name:  "test-processor",
				Id:    "507f1f77bcf86cd799439011",
				State: "CREATED",
			},
			checkFields: []string{"name", "id", "state"},
		},
		"withAllFields": {
			streamProcessor: &admin.StreamsProcessorWithStats{
				Name:     "test-processor",
				Id:       "507f1f77bcf86cd799439011",
				State:    "STARTED",
				Pipeline: []any{map[string]any{"$match": map[string]any{"status": "active"}}},
				Stats:    map[string]any{"bytesProcessed": 5000, "recordsProcessed": 500},
				Options:  &admin.StreamsOptions{Dlq: validDLQ},
			},
			currentModel: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
			},
			checkFields: []string{"all"},
		},
		"preserveCurrentModelOptions": {
			streamProcessor: &admin.StreamsProcessorWithStats{
				Name:  "test-processor",
				Id:    "507f1f77bcf86cd799439011",
				State: "CREATED",
			},
			currentModel: currentModelWithDLQ,
			checkFields:  []string{"preservedOptions"},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := resource.GetStreamProcessorModel(tc.streamProcessor, tc.currentModel)
			require.NoError(t, err)
			require.NotNil(t, result)

			for _, field := range tc.checkFields {
				switch field {
				case "name":
					assert.Equal(t, "test-processor", util.SafeString(result.ProcessorName))
				case "id":
					assert.Equal(t, "507f1f77bcf86cd799439011", util.SafeString(result.Id))
				case "state":
					assert.Equal(t, tc.streamProcessor.State, util.SafeString(result.State))
				case "all":
					assert.Equal(t, "test-processor", util.SafeString(result.ProcessorName))
					assert.NotNil(t, result.Pipeline)
					assert.NotNil(t, result.Stats)
					require.NotNil(t, result.Options.Dlq)
					assert.Equal(t, "dlq-collection", util.SafeString(result.Options.Dlq.Coll))
				case "preservedOptions":
					require.NotNil(t, result.Options.Dlq)
					assert.Equal(t, "existing-dlq-collection", util.SafeString(result.Options.Dlq.Coll))
				}
			}
		})
	}
}

func TestRoundTripConversions(t *testing.T) {
	t.Run("pipelineRoundTrip", func(t *testing.T) {
		originalPipeline := `[{"$match": {"status": "active"}}, {"$group": {"_id": "$category", "count": {"$sum": 1}}}]`
		sdkPipeline, err := resource.ConvertPipelineToSdk(originalPipeline)
		require.NoError(t, err)
		convertedBack, err := resource.ConvertPipelineToString(sdkPipeline)
		require.NoError(t, err)

		var original, converted any
		require.NoError(t, json.Unmarshal([]byte(originalPipeline), &original))
		require.NoError(t, json.Unmarshal([]byte(convertedBack), &converted))
		assert.Equal(t, original, converted)
	})

	t.Run("statsRoundTrip", func(t *testing.T) {
		originalStats := map[string]any{
			"bytesProcessed":   1000,
			"recordsProcessed": 100,
			"nested":           map[string]any{"value": 42},
		}
		statsString, err := resource.ConvertStatsToString(originalStats)
		require.NoError(t, err)

		var parsedStats any
		require.NoError(t, json.Unmarshal([]byte(statsString), &parsedStats))
		parsedMap := parsedStats.(map[string]any)

		assert.InDelta(t, float64(1000), parsedMap["bytesProcessed"], 0.01)
		assert.InDelta(t, float64(100), parsedMap["recordsProcessed"], 0.01)
		nested := parsedMap["nested"].(map[string]any)
		assert.InDelta(t, float64(42), nested["value"], 0.01)
	})
}
