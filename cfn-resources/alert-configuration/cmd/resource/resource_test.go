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

	"github.com/mongodb/mongodbatlas-cloudformation-resources/alert-configuration/cmd/resource"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20250312021/admin"
)

func TestConvertToMongoModel_ClearsAtlasReturnedThreshold(t *testing.T) {
	metricThresholdName := "NORMALIZED_SYSTEM_CPU_USER"

	t.Run("only MetricThreshold in model clears Threshold", func(t *testing.T) {
		atlasReq := &admin.GroupAlertsConfig{
			MetricThreshold: &admin.FlexClusterMetricThreshold{MetricName: metricThresholdName},
			Threshold:       &admin.StreamProcessorMetricThreshold{},
		}
		model := &resource.Model{
			MetricThreshold: &resource.MetricThresholdView{MetricName: &metricThresholdName},
		}
		result := resource.ConvertToMongoModel(atlasReq, model)
		assert.NotNil(t, result.MetricThreshold)
		assert.Nil(t, result.Threshold)
	})

	t.Run("only Threshold in model clears MetricThreshold", func(t *testing.T) {
		thresholdMetricName := "STREAM_PROCESSOR_KAFKA_LAG"
		atlasReq := &admin.GroupAlertsConfig{
			MetricThreshold: &admin.FlexClusterMetricThreshold{MetricName: metricThresholdName},
			Threshold:       &admin.StreamProcessorMetricThreshold{},
		}
		model := &resource.Model{
			Threshold: &resource.IntegerThresholdView{MetricName: &thresholdMetricName},
		}
		result := resource.ConvertToMongoModel(atlasReq, model)
		assert.Nil(t, result.MetricThreshold)
		assert.NotNil(t, result.Threshold)
	})
}
