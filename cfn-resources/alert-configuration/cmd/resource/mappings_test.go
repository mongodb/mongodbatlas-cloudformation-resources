package resource_test

import (
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/alert-configuration/cmd/resource"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"
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
