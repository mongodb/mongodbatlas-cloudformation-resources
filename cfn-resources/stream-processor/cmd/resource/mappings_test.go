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

package resource_test

import (
	"encoding/json"
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-processor/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"
)

func TestGetWorkspaceOrInstanceName(t *testing.T) {
	testCases := map[string]struct {
		model          *resource.Model
		expectedResult string
		expectedError  string
	}{
		"workspaceNameOnly": {
			model: &resource.Model{
				WorkspaceName: util.StringPtr("workspace-1"),
				InstanceName:  nil,
			},
			expectedResult: "workspace-1",
			expectedError:  "",
		},
		"instanceNameOnly": {
			model: &resource.Model{
				WorkspaceName: nil,
				InstanceName:  util.StringPtr("instance-1"),
			},
			expectedResult: "instance-1",
			expectedError:  "",
		},
		"bothSet": {
			model: &resource.Model{
				WorkspaceName: util.StringPtr("workspace-1"),
				InstanceName:  util.StringPtr("instance-1"),
			},
			expectedResult: "workspace-1", // WorkspaceName takes precedence
			expectedError:  "",
		},
		"neitherSet": {
			model: &resource.Model{
				WorkspaceName: nil,
				InstanceName:  nil,
			},
			expectedResult: "",
			expectedError:  "either WorkspaceName or InstanceName must be provided",
		},
		"workspaceNameEmptyString": {
			model: &resource.Model{
				WorkspaceName: util.StringPtr(""),
				InstanceName:  util.StringPtr("instance-1"),
			},
			expectedResult: "instance-1",
			expectedError:  "",
		},
		"instanceNameEmptyString": {
			model: &resource.Model{
				WorkspaceName: util.StringPtr("workspace-1"),
				InstanceName:  util.StringPtr(""),
			},
			expectedResult: "workspace-1",
			expectedError:  "",
		},
		"bothEmptyStrings": {
			model: &resource.Model{
				WorkspaceName: util.StringPtr(""),
				InstanceName:  util.StringPtr(""),
			},
			expectedResult: "",
			expectedError:  "either WorkspaceName or InstanceName must be provided",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := resource.GetWorkspaceOrInstanceName(tc.model)
			if tc.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Empty(t, result)
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
		validateFunc  func(t *testing.T, result []any)
	}{
		"validSimplePipeline": {
			pipeline:      `[{"$match": {"status": "active"}}]`,
			expectedError: false,
			validateFunc: func(t *testing.T, result []any) {
				require.Len(t, result, 1)
				stage, ok := result[0].(map[string]any)
				require.True(t, ok)
				assert.Equal(t, "active", stage["$match"].(map[string]any)["status"])
			},
		},
		"validComplexPipeline": {
			pipeline:      `[{"$match": {"status": "active"}}, {"$group": {"_id": "$category", "count": {"$sum": 1}}}]`,
			expectedError: false,
			validateFunc: func(t *testing.T, result []any) {
				require.Len(t, result, 2)
			},
		},
		"validEmptyPipeline": {
			pipeline:      `[]`,
			expectedError: false,
			validateFunc: func(t *testing.T, result []any) {
				assert.Empty(t, result)
			},
		},
		"invalidJSON": {
			pipeline:      `[{"$match": {"status": "active"}`,
			expectedError: true,
			validateFunc:  nil,
		},
		"notAnArray": {
			pipeline:      `{"$match": {"status": "active"}}`,
			expectedError: true, // JSON unmarshal fails when trying to unmarshal object into slice
			validateFunc:  nil,
		},
		"emptyString": {
			pipeline:      ``,
			expectedError: true,
			validateFunc:  nil,
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
				if tc.validateFunc != nil {
					tc.validateFunc(t, result)
				}
			}
		})
	}
}

func TestConvertPipelineToString(t *testing.T) {
	testCases := map[string]struct {
		pipeline      []any
		expectedJSON  string
		expectedError bool
	}{
		"validPipeline": {
			pipeline: []any{
				map[string]any{"$match": map[string]any{"status": "active"}},
			},
			expectedJSON:  `[{"$match":{"status":"active"}}]`,
			expectedError: false,
		},
		"validComplexPipeline": {
			pipeline: []any{
				map[string]any{"$match": map[string]any{"status": "active"}},
				map[string]any{"$group": map[string]any{"_id": "$category"}},
			},
			expectedJSON:  `[{"$match":{"status":"active"}},{"$group":{"_id":"$category"}}]`,
			expectedError: false,
		},
		"emptyPipeline": {
			pipeline:      []any{},
			expectedJSON:  `[]`,
			expectedError: false,
		},
		"nilPipeline": {
			pipeline:      nil,
			expectedJSON:  `null`,
			expectedError: false,
		},
		"pipelineWithNestedObjects": {
			pipeline: []any{
				map[string]any{
					"$match": map[string]any{
						"status": "active",
						"tags":   []any{"important", "urgent"},
					},
				},
			},
			expectedJSON:  `[{"$match":{"status":"active","tags":["important","urgent"]}}]`,
			expectedError: false,
		},
		"pipelineWithNumbers": {
			pipeline: []any{
				map[string]any{
					"$limit": 10,
					"$skip":  5,
				},
			},
			expectedJSON:  `[{"$limit":10,"$skip":5}]`,
			expectedError: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := resource.ConvertPipelineToString(tc.pipeline)
			if tc.expectedError {
				require.Error(t, err)
				assert.Empty(t, result)
			} else {
				require.NoError(t, err)
				// Parse both JSONs and compare to handle formatting differences
				var resultJSON, expectedJSON any
				require.NoError(t, json.Unmarshal([]byte(result), &resultJSON))
				require.NoError(t, json.Unmarshal([]byte(tc.expectedJSON), &expectedJSON))
				assert.Equal(t, expectedJSON, resultJSON)
			}
		})
	}
}

func TestConvertStatsToString(t *testing.T) {
	testCases := map[string]struct {
		stats         any
		expectedJSON  string
		expectedError bool
	}{
		"nilStats": {
			stats:         nil,
			expectedJSON:  "",
			expectedError: false,
		},
		"validStats": {
			stats: map[string]any{
				"bytesProcessed":   1000,
				"recordsProcessed": 100,
			},
			expectedJSON:  `{"bytesProcessed":1000,"recordsProcessed":100}`,
			expectedError: false,
		},
		"emptyMap": {
			stats:         map[string]any{},
			expectedJSON:  `{}`,
			expectedError: false,
		},
		"nestedStats": {
			stats: map[string]any{
				"input": map[string]any{
					"bytes": 1000,
				},
				"output": map[string]any{
					"records": 100,
				},
			},
			expectedJSON:  `{"input":{"bytes":1000},"output":{"records":100}}`,
			expectedError: false,
		},
		"statsWithArray": {
			stats: map[string]any{
				"errors": []any{"error1", "error2"},
				"count":  5,
			},
			expectedJSON:  `{"count":5,"errors":["error1","error2"]}`,
			expectedError: false,
		},
		"statsWithNumbers": {
			stats: map[string]any{
				"floatValue": 3.14,
				"intValue":   42,
			},
			expectedJSON:  `{"floatValue":3.14,"intValue":42}`,
			expectedError: false,
		},
		"statsWithBoolean": {
			stats: map[string]any{
				"enabled": true,
				"active":  false,
			},
			expectedJSON:  `{"active":false,"enabled":true}`,
			expectedError: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := resource.ConvertStatsToString(tc.stats)
			if tc.expectedError {
				require.Error(t, err)
				assert.Empty(t, result)
			} else {
				require.NoError(t, err)
				if tc.stats == nil {
					assert.Empty(t, result)
				} else {
					// Parse both JSONs and compare to handle formatting differences
					var resultJSON, expectedJSON any
					require.NoError(t, json.Unmarshal([]byte(result), &resultJSON))
					require.NoError(t, json.Unmarshal([]byte(tc.expectedJSON), &expectedJSON))
					assert.Equal(t, expectedJSON, resultJSON)
				}
			}
		})
	}
}

func TestNewStreamProcessorReq(t *testing.T) {
	testCases := map[string]struct {
		model         *resource.Model
		expectedError bool
		validateFunc  func(t *testing.T, result *admin20250312010.StreamsProcessor)
	}{
		"minimalRequest": {
			model: &resource.Model{
				ProcessorName: util.StringPtr("test-processor"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *admin20250312010.StreamsProcessor) {
				require.NotNil(t, result)
				assert.Equal(t, "test-processor", result.GetName())
				assert.NotNil(t, result.Pipeline)
				pipeline := result.GetPipeline()
				require.Len(t, pipeline, 1)
			},
		},
		"withOptions": {
			model: &resource.Model{
				ProcessorName: util.StringPtr("test-processor"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
				Options: &resource.StreamsOptions{
					Dlq: &resource.StreamsDLQ{
						Coll:           util.StringPtr("dlq-collection"),
						ConnectionName: util.StringPtr("dlq-connection"),
						Db:             util.StringPtr("dlq-db"),
					},
				},
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *admin20250312010.StreamsProcessor) {
				require.NotNil(t, result)
				assert.Equal(t, "test-processor", result.GetName())
				require.NotNil(t, result.Options)
				require.NotNil(t, result.Options.Dlq)
				assert.Equal(t, "dlq-collection", util.SafeString(result.Options.Dlq.Coll))
				assert.Equal(t, "dlq-connection", util.SafeString(result.Options.Dlq.ConnectionName))
				assert.Equal(t, "dlq-db", util.SafeString(result.Options.Dlq.Db))
			},
		},
		"invalidPipeline": {
			model: &resource.Model{
				ProcessorName: util.StringPtr("test-processor"),
				Pipeline:      util.StringPtr(`invalid json`),
			},
			expectedError: true,
			validateFunc:  nil,
		},
		"nilPipeline": {
			model: &resource.Model{
				ProcessorName: util.StringPtr("test-processor"),
				Pipeline:      nil,
			},
			expectedError: true,
			validateFunc:  nil,
		},
		"emptyPipeline": {
			model: &resource.Model{
				ProcessorName: util.StringPtr("test-processor"),
				Pipeline:      util.StringPtr(``),
			},
			expectedError: true,
			validateFunc:  nil,
		},
		"withOptionsButNilDlq": {
			model: &resource.Model{
				ProcessorName: util.StringPtr("test-processor"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
				Options: &resource.StreamsOptions{
					Dlq: nil,
				},
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *admin20250312010.StreamsProcessor) {
				require.NotNil(t, result)
				assert.Nil(t, result.Options)
			},
		},
		"withNilOptions": {
			model: &resource.Model{
				ProcessorName: util.StringPtr("test-processor"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
				Options:       nil,
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *admin20250312010.StreamsProcessor) {
				require.NotNil(t, result)
				assert.Nil(t, result.Options)
			},
		},
		"complexPipeline": {
			model: &resource.Model{
				ProcessorName: util.StringPtr("test-processor"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}, {"$group": {"_id": "$category", "count": {"$sum": 1}}}, {"$sort": {"count": -1}}]`),
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *admin20250312010.StreamsProcessor) {
				require.NotNil(t, result)
				pipeline := result.GetPipeline()
				require.Len(t, pipeline, 3)
			},
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
				if tc.validateFunc != nil {
					tc.validateFunc(t, result)
				}
			}
		})
	}
}

func TestNewStreamProcessorUpdateReq(t *testing.T) {
	testCases := map[string]struct {
		model         *resource.Model
		expectedError bool
		validateFunc  func(t *testing.T, result *admin20250312010.UpdateStreamProcessorApiParams)
	}{
		"minimalRequest": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *admin20250312010.UpdateStreamProcessorApiParams) {
				require.NotNil(t, result)
				assert.Equal(t, "507f1f77bcf86cd799439011", result.GroupId)
				assert.Equal(t, "workspace-1", result.TenantName)
				assert.Equal(t, "test-processor", result.ProcessorName)
				require.NotNil(t, result.StreamsModifyStreamProcessor)
				assert.Equal(t, "test-processor", result.StreamsModifyStreamProcessor.GetName())
				assert.NotNil(t, result.StreamsModifyStreamProcessor.Pipeline)
			},
		},
		"withInstanceName": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
				InstanceName:  util.StringPtr("instance-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *admin20250312010.UpdateStreamProcessorApiParams) {
				require.NotNil(t, result)
				assert.Equal(t, "instance-1", result.TenantName)
			},
		},
		"withOptions": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
				Options: &resource.StreamsOptions{
					Dlq: &resource.StreamsDLQ{
						Coll:           util.StringPtr("dlq-collection"),
						ConnectionName: util.StringPtr("dlq-connection"),
						Db:             util.StringPtr("dlq-db"),
					},
				},
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *admin20250312010.UpdateStreamProcessorApiParams) {
				require.NotNil(t, result)
				require.NotNil(t, result.StreamsModifyStreamProcessor)
				require.NotNil(t, result.StreamsModifyStreamProcessor.Options)
				require.NotNil(t, result.StreamsModifyStreamProcessor.Options.Dlq)
				assert.Equal(t, "dlq-collection", result.StreamsModifyStreamProcessor.Options.Dlq.GetColl())
				assert.Equal(t, "dlq-connection", result.StreamsModifyStreamProcessor.Options.Dlq.GetConnectionName())
				assert.Equal(t, "dlq-db", result.StreamsModifyStreamProcessor.Options.Dlq.GetDb())
			},
		},
		"invalidPipeline": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`invalid json`),
			},
			expectedError: true,
			validateFunc:  nil,
		},
		"bothWorkspaceAndInstanceName": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
				WorkspaceName: util.StringPtr("workspace-1"),
				InstanceName:  util.StringPtr("instance-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *admin20250312010.UpdateStreamProcessorApiParams) {
				require.NotNil(t, result)
				// WorkspaceName takes precedence over InstanceName
				assert.Equal(t, "workspace-1", result.TenantName)
			},
		},
		"neitherWorkspaceNorInstanceName": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			expectedError: true,
			validateFunc:  nil,
		},
		"withNilOptions": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
				Options:       nil,
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *admin20250312010.UpdateStreamProcessorApiParams) {
				require.NotNil(t, result)
				require.NotNil(t, result.StreamsModifyStreamProcessor)
				assert.Nil(t, result.StreamsModifyStreamProcessor.Options)
			},
		},
		"withOptionsButNilDlq": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
				Options: &resource.StreamsOptions{
					Dlq: nil,
				},
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *admin20250312010.UpdateStreamProcessorApiParams) {
				require.NotNil(t, result)
				require.NotNil(t, result.StreamsModifyStreamProcessor)
				assert.Nil(t, result.StreamsModifyStreamProcessor.Options)
			},
		},
		"nilProcessorName": {
			model: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: nil,
				WorkspaceName: util.StringPtr("workspace-1"),
				Pipeline:      util.StringPtr(`[{"$match": {"status": "active"}}]`),
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *admin20250312010.UpdateStreamProcessorApiParams) {
				require.NotNil(t, result)
				assert.Empty(t, result.ProcessorName)
			},
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
				if tc.validateFunc != nil {
					tc.validateFunc(t, result)
				}
			}
		})
	}
}

func TestGetStreamProcessorModel(t *testing.T) {
	testCases := map[string]struct {
		streamProcessor *admin20250312010.StreamsProcessorWithStats
		currentModel    *resource.Model
		validateFunc    func(t *testing.T, result *resource.Model)
		expectedError   bool
	}{
		"minimalConversion": {
			streamProcessor: &admin20250312010.StreamsProcessorWithStats{
				Name:  "test-processor",
				Id:    "507f1f77bcf86cd799439011",
				State: "CREATED",
			},
			currentModel:  nil,
			expectedError: false,
			validateFunc: func(t *testing.T, result *resource.Model) {
				require.NotNil(t, result)
				assert.Equal(t, "test-processor", util.SafeString(result.ProcessorName))
				assert.Equal(t, "507f1f77bcf86cd799439011", util.SafeString(result.Id))
				assert.Equal(t, "CREATED", util.SafeString(result.State))
			},
		},
		"withPipeline": {
			streamProcessor: func() *admin20250312010.StreamsProcessorWithStats {
				pipeline := []any{
					map[string]any{"$match": map[string]any{"status": "active"}},
				}
				return &admin20250312010.StreamsProcessorWithStats{
					Name:     "test-processor",
					Id:       "507f1f77bcf86cd799439011",
					State:    "CREATED",
					Pipeline: pipeline,
				}
			}(),
			currentModel:  nil,
			expectedError: false,
			validateFunc: func(t *testing.T, result *resource.Model) {
				require.NotNil(t, result)
				assert.NotNil(t, result.Pipeline)
				// Verify pipeline is valid JSON
				var pipelineJSON any
				err := json.Unmarshal([]byte(*result.Pipeline), &pipelineJSON)
				require.NoError(t, err)
			},
		},
		"withStats": {
			streamProcessor: &admin20250312010.StreamsProcessorWithStats{
				Name:  "test-processor",
				Id:    "507f1f77bcf86cd799439011",
				State: "CREATED",
				Stats: map[string]any{
					"bytesProcessed": 1000,
				},
			},
			currentModel:  nil,
			expectedError: false,
			validateFunc: func(t *testing.T, result *resource.Model) {
				require.NotNil(t, result)
				assert.NotNil(t, result.Stats)
				// Verify stats is valid JSON
				var statsJSON any
				err := json.Unmarshal([]byte(*result.Stats), &statsJSON)
				require.NoError(t, err)
			},
		},
		"withOptions": {
			streamProcessor: &admin20250312010.StreamsProcessorWithStats{
				Name:  "test-processor",
				Id:    "507f1f77bcf86cd799439011",
				State: "CREATED",
				Options: &admin20250312010.StreamsOptions{
					Dlq: &admin20250312010.StreamsDLQ{
						Coll:           admin20250312010.PtrString("dlq-collection"),
						ConnectionName: admin20250312010.PtrString("dlq-connection"),
						Db:             admin20250312010.PtrString("dlq-db"),
					},
				},
			},
			currentModel:  nil,
			expectedError: false,
			validateFunc: func(t *testing.T, result *resource.Model) {
				require.NotNil(t, result)
				require.NotNil(t, result.Options)
				require.NotNil(t, result.Options.Dlq)
				assert.Equal(t, "dlq-collection", util.SafeString(result.Options.Dlq.Coll))
				assert.Equal(t, "dlq-connection", util.SafeString(result.Options.Dlq.ConnectionName))
				assert.Equal(t, "dlq-db", util.SafeString(result.Options.Dlq.Db))
			},
		},
		"preserveCurrentModelOptions": {
			streamProcessor: &admin20250312010.StreamsProcessorWithStats{
				Name:  "test-processor",
				Id:    "507f1f77bcf86cd799439011",
				State: "CREATED",
				// No options in response
			},
			currentModel: &resource.Model{
				Options: &resource.StreamsOptions{
					Dlq: &resource.StreamsDLQ{
						Coll:           util.StringPtr("existing-dlq-collection"),
						ConnectionName: util.StringPtr("existing-dlq-connection"),
						Db:             util.StringPtr("existing-dlq-db"),
					},
				},
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *resource.Model) {
				require.NotNil(t, result)
				require.NotNil(t, result.Options)
				require.NotNil(t, result.Options.Dlq)
				// Should preserve current model's options
				assert.Equal(t, "existing-dlq-collection", util.SafeString(result.Options.Dlq.Coll))
				assert.Equal(t, "existing-dlq-connection", util.SafeString(result.Options.Dlq.ConnectionName))
				assert.Equal(t, "existing-dlq-db", util.SafeString(result.Options.Dlq.Db))
			},
		},
		"withCurrentModel": {
			streamProcessor: &admin20250312010.StreamsProcessorWithStats{
				Name:  "test-processor",
				Id:    "507f1f77bcf86cd799439011",
				State: "STARTED",
			},
			currentModel: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *resource.Model) {
				require.NotNil(t, result)
				// Should use currentModel as base
				assert.Equal(t, "507f1f77bcf86cd799439011", util.SafeString(result.ProjectId))
				assert.Equal(t, "test-processor", util.SafeString(result.ProcessorName))
				assert.Equal(t, "STARTED", util.SafeString(result.State))
			},
		},
		"withAllFields": {
			streamProcessor: func() *admin20250312010.StreamsProcessorWithStats {
				pipeline := []any{
					map[string]any{"$match": map[string]any{"status": "active"}},
					map[string]any{"$group": map[string]any{"_id": "$category"}},
				}
				return &admin20250312010.StreamsProcessorWithStats{
					Name:     "test-processor",
					Id:       "507f1f77bcf86cd799439011",
					State:    "STARTED",
					Pipeline: pipeline,
					Stats: map[string]any{
						"bytesProcessed":   5000,
						"recordsProcessed": 500,
					},
					Options: &admin20250312010.StreamsOptions{
						Dlq: &admin20250312010.StreamsDLQ{
							Coll:           admin20250312010.PtrString("dlq-collection"),
							ConnectionName: admin20250312010.PtrString("dlq-connection"),
							Db:             admin20250312010.PtrString("dlq-db"),
						},
					},
				}
			}(),
			currentModel: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				ProcessorName: util.StringPtr("test-processor"),
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *resource.Model) {
				require.NotNil(t, result)
				assert.Equal(t, "test-processor", util.SafeString(result.ProcessorName))
				assert.Equal(t, "507f1f77bcf86cd799439011", util.SafeString(result.Id))
				assert.Equal(t, "STARTED", util.SafeString(result.State))
				assert.NotNil(t, result.Pipeline)
				assert.NotNil(t, result.Stats)
				require.NotNil(t, result.Options)
				require.NotNil(t, result.Options.Dlq)
				assert.Equal(t, "dlq-collection", util.SafeString(result.Options.Dlq.Coll))
				assert.Equal(t, "dlq-connection", util.SafeString(result.Options.Dlq.ConnectionName))
				assert.Equal(t, "dlq-db", util.SafeString(result.Options.Dlq.Db))
			},
		},
		"withOptionsButNoDlq": {
			streamProcessor: &admin20250312010.StreamsProcessorWithStats{
				Name:    "test-processor",
				Id:      "507f1f77bcf86cd799439011",
				State:   "CREATED",
				Options: &admin20250312010.StreamsOptions{
					// Options exists but Dlq is nil
				},
			},
			currentModel: &resource.Model{
				Options: &resource.StreamsOptions{
					Dlq: &resource.StreamsDLQ{
						Coll:           util.StringPtr("existing-dlq-collection"),
						ConnectionName: util.StringPtr("existing-dlq-connection"),
						Db:             util.StringPtr("existing-dlq-db"),
					},
				},
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *resource.Model) {
				require.NotNil(t, result)
				// Should preserve current model's options when response has no Dlq
				require.NotNil(t, result.Options)
				require.NotNil(t, result.Options.Dlq)
				assert.Equal(t, "existing-dlq-collection", util.SafeString(result.Options.Dlq.Coll))
			},
		},
		"withNilPipeline": {
			streamProcessor: &admin20250312010.StreamsProcessorWithStats{
				Name:     "test-processor",
				Id:       "507f1f77bcf86cd799439011",
				State:    "CREATED",
				Pipeline: nil,
			},
			currentModel:  nil,
			expectedError: false,
			validateFunc: func(t *testing.T, result *resource.Model) {
				require.NotNil(t, result)
				assert.Nil(t, result.Pipeline)
			},
		},
		"withNilStats": {
			streamProcessor: &admin20250312010.StreamsProcessorWithStats{
				Name:  "test-processor",
				Id:    "507f1f77bcf86cd799439011",
				State: "CREATED",
				Stats: nil,
			},
			currentModel:  nil,
			expectedError: false,
			validateFunc: func(t *testing.T, result *resource.Model) {
				require.NotNil(t, result)
				assert.Nil(t, result.Stats)
			},
		},
		"withNilOptions": {
			streamProcessor: &admin20250312010.StreamsProcessorWithStats{
				Name:    "test-processor",
				Id:      "507f1f77bcf86cd799439011",
				State:   "CREATED",
				Options: nil,
			},
			currentModel: &resource.Model{
				Options: &resource.StreamsOptions{
					Dlq: &resource.StreamsDLQ{
						Coll:           util.StringPtr("existing-dlq-collection"),
						ConnectionName: util.StringPtr("existing-dlq-connection"),
						Db:             util.StringPtr("existing-dlq-db"),
					},
				},
			},
			expectedError: false,
			validateFunc: func(t *testing.T, result *resource.Model) {
				require.NotNil(t, result)
				// Should preserve current model's options when response has nil options
				require.NotNil(t, result.Options)
				require.NotNil(t, result.Options.Dlq)
				assert.Equal(t, "existing-dlq-collection", util.SafeString(result.Options.Dlq.Coll))
			},
		},
		"withEmptyPipeline": {
			streamProcessor: &admin20250312010.StreamsProcessorWithStats{
				Name:     "test-processor",
				Id:       "507f1f77bcf86cd799439011",
				State:    "CREATED",
				Pipeline: []any{},
			},
			currentModel:  nil,
			expectedError: false,
			validateFunc: func(t *testing.T, result *resource.Model) {
				require.NotNil(t, result)
				assert.NotNil(t, result.Pipeline)
				assert.Equal(t, "[]", *result.Pipeline)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := resource.GetStreamProcessorModel(tc.streamProcessor, tc.currentModel)
			if tc.expectedError {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				if tc.validateFunc != nil {
					tc.validateFunc(t, result)
				}
			}
		})
	}
}

// TestPipelineRoundTrip tests that converting pipeline to SDK and back preserves data
func TestPipelineRoundTrip(t *testing.T) {
	originalPipeline := `[{"$match": {"status": "active"}}, {"$group": {"_id": "$category", "count": {"$sum": 1}}}]`

	sdkPipeline, err := resource.ConvertPipelineToSdk(originalPipeline)
	require.NoError(t, err)

	convertedBack, err := resource.ConvertPipelineToString(sdkPipeline)
	require.NoError(t, err)

	// Parse both and compare
	var original, converted any
	require.NoError(t, json.Unmarshal([]byte(originalPipeline), &original))
	require.NoError(t, json.Unmarshal([]byte(convertedBack), &converted))
	assert.Equal(t, original, converted)
}

// TestStatsRoundTrip tests that converting stats to string and parsing back preserves data
func TestStatsRoundTrip(t *testing.T) {
	originalStats := map[string]any{
		"bytesProcessed":   1000,
		"recordsProcessed": 100,
		"nested": map[string]any{
			"value": 42,
		},
	}

	statsString, err := resource.ConvertStatsToString(originalStats)
	require.NoError(t, err)

	// Parse back and compare
	var parsedStats any
	require.NoError(t, json.Unmarshal([]byte(statsString), &parsedStats))

	// JSON unmarshaling converts numbers to float64, so we need to compare values
	parsedMap, ok := parsedStats.(map[string]any)
	require.True(t, ok)

	// Compare values accounting for int->float64 conversion
	assert.Equal(t, float64(1000), parsedMap["bytesProcessed"])
	assert.Equal(t, float64(100), parsedMap["recordsProcessed"])
	nested, ok := parsedMap["nested"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, float64(42), nested["value"])
}
