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
	"encoding/json"
	"fmt"

	"go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

func GetWorkspaceOrInstanceName(model *Model) (string, error) {
	if model.WorkspaceName != nil && *model.WorkspaceName != "" {
		return *model.WorkspaceName, nil
	}
	if model.InstanceName != nil && *model.InstanceName != "" {
		return *model.InstanceName, nil
	}
	return "", fmt.Errorf("either WorkspaceName or InstanceName must be provided")
}

func ConvertPipelineToSdk(pipeline string) ([]any, error) {
	var pipelineSliceOfMaps []any
	err := json.Unmarshal([]byte(pipeline), &pipelineSliceOfMaps)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal pipeline: %w", err)
	}
	return pipelineSliceOfMaps, nil
}

func ConvertPipelineToString(pipeline []any) (string, error) {
	pipelineJSON, err := json.Marshal(pipeline)
	if err != nil {
		return "", fmt.Errorf("failed to marshal pipeline: %w", err)
	}
	return string(pipelineJSON), nil
}

func ConvertStatsToString(stats any) (string, error) {
	if stats == nil {
		return "", nil
	}
	statsJSON, err := json.Marshal(stats)
	if err != nil {
		return "", fmt.Errorf("failed to marshal stats: %w", err)
	}
	return string(statsJSON), nil
}

func NewStreamProcessorReq(model *Model) (*admin.StreamsProcessor, error) {
	pipeline, err := ConvertPipelineToSdk(util.SafeString(model.Pipeline))
	if err != nil {
		return nil, err
	}

	streamProcessor := &admin.StreamsProcessor{
		Name:     model.ProcessorName,
		Pipeline: &pipeline,
	}

	if model.Options != nil && model.Options.Dlq != nil {
		dlq := model.Options.Dlq
		if dlq.Coll != nil && *dlq.Coll != "" &&
			dlq.ConnectionName != nil && *dlq.ConnectionName != "" &&
			dlq.Db != nil && *dlq.Db != "" {
			streamProcessor.Options = &admin.StreamsOptions{
				Dlq: &admin.StreamsDLQ{
					Coll:           dlq.Coll,
					ConnectionName: dlq.ConnectionName,
					Db:             dlq.Db,
				},
			}
		}
	}

	return streamProcessor, nil
}

func NewStreamProcessorUpdateReq(model *Model) (*admin.UpdateStreamProcessorApiParams, error) {
	pipeline, err := ConvertPipelineToSdk(util.SafeString(model.Pipeline))
	if err != nil {
		return nil, err
	}

	workspaceOrInstanceName, err := GetWorkspaceOrInstanceName(model)
	if err != nil {
		return nil, err
	}

	streamProcessorAPIParams := &admin.UpdateStreamProcessorApiParams{
		GroupId:       util.SafeString(model.ProjectId),
		TenantName:    workspaceOrInstanceName,
		ProcessorName: util.SafeString(model.ProcessorName),
		StreamsModifyStreamProcessor: &admin.StreamsModifyStreamProcessor{
			Name:     model.ProcessorName,
			Pipeline: &pipeline,
		},
	}

	if model.Options != nil && model.Options.Dlq != nil {
		dlq := model.Options.Dlq
		if dlq.Coll != nil && *dlq.Coll != "" &&
			dlq.ConnectionName != nil && *dlq.ConnectionName != "" &&
			dlq.Db != nil && *dlq.Db != "" {
			streamProcessorAPIParams.StreamsModifyStreamProcessor.Options = &admin.StreamsModifyStreamProcessorOptions{
				Dlq: &admin.StreamsDLQ{
					Coll:           dlq.Coll,
					ConnectionName: dlq.ConnectionName,
					Db:             dlq.Db,
				},
			}
		}
	}

	return streamProcessorAPIParams, nil
}

func GetStreamProcessorModel(streamProcessor *admin.StreamsProcessorWithStats, currentModel *Model) (*Model, error) {
	model := new(Model)

	if currentModel != nil {
		*model = *currentModel
		model.DeleteOnCreateTimeout = nil
	}

	model.ProcessorName = util.Pointer(streamProcessor.Name)
	model.Id = util.Pointer(streamProcessor.Id)
	model.State = util.Pointer(streamProcessor.State)

	if currentModel != nil && currentModel.Pipeline != nil {
		model.Pipeline = currentModel.Pipeline
	} else if streamProcessor.Pipeline != nil {
		pipelineStr, err := ConvertPipelineToString(streamProcessor.GetPipeline())
		if err != nil {
			return nil, err
		}
		model.Pipeline = &pipelineStr
	}

	if streamProcessor.Stats != nil {
		statsStr, err := ConvertStatsToString(streamProcessor.GetStats())
		if err != nil {
			return nil, err
		}
		model.Stats = &statsStr
	}

	if streamProcessor.Options != nil && streamProcessor.Options.Dlq != nil {
		apiDlq := streamProcessor.Options.Dlq
		if apiDlq.Coll != nil && *apiDlq.Coll != "" &&
			apiDlq.ConnectionName != nil && *apiDlq.ConnectionName != "" &&
			apiDlq.Db != nil && *apiDlq.Db != "" {
			model.Options = &StreamsOptions{
				Dlq: &StreamsDLQ{
					Coll:           apiDlq.Coll,
					ConnectionName: apiDlq.ConnectionName,
					Db:             apiDlq.Db,
				},
			}
		}
	} else {
		if currentModel != nil && currentModel.Options != nil && currentModel.Options.Dlq != nil {
			currentDlq := currentModel.Options.Dlq
			if currentDlq.Coll != nil && *currentDlq.Coll != "" &&
				currentDlq.ConnectionName != nil && *currentDlq.ConnectionName != "" &&
				currentDlq.Db != nil && *currentDlq.Db != "" {
				model.Options = currentModel.Options
			}
		}
	}

	return model, nil
}
