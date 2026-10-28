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

package cloudprovideraccess_test

import (
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/cloudprovideraccess"
)

func TestFeatureUsages(t *testing.T) {
	const (
		projectID  = "66d979971ec97b7de1ef8777"
		bucketName = "mongodb-atlas-log-export"
	)

	testCases := []struct {
		name     string
		usages   *[]admin.CloudProviderAccessFeatureUsage
		expected []cloudprovideraccess.FeatureUsage
	}{
		{
			name:     "No usages are omitted",
			expected: nil,
		},
		{
			name:     "An empty list is omitted",
			usages:   &[]admin.CloudProviderAccessFeatureUsage{},
			expected: nil,
		},
		{
			name: "A feature that Atlas does not identify has no identifiers",
			usages: &[]admin.CloudProviderAccessFeatureUsage{
				{FeatureType: aws.String("ENCRYPTION_AT_REST")},
			},
			expected: []cloudprovideraccess.FeatureUsage{
				{FeatureType: aws.String("ENCRYPTION_AT_REST")},
			},
		},
		{
			name: "A feature with identifiers keeps them",
			usages: &[]admin.CloudProviderAccessFeatureUsage{
				{
					FeatureType: aws.String("PUSH_BASED_LOG_EXPORT"),
					FeatureId: &admin.CloudProviderAccessFeatureUsagePushBasedLogExportFeatureId{
						BucketName: aws.String(bucketName),
						GroupId:    aws.String(projectID),
					},
				},
			},
			expected: []cloudprovideraccess.FeatureUsage{
				{
					FeatureType: aws.String("PUSH_BASED_LOG_EXPORT"),
					ProjectID:   aws.String(projectID),
					BucketName:  aws.String(bucketName),
				},
			},
		},
		{
			name: "An empty feature id is not mapped",
			usages: &[]admin.CloudProviderAccessFeatureUsage{
				{
					FeatureType: aws.String("ENCRYPTION_AT_REST"),
					FeatureId:   &admin.CloudProviderAccessFeatureUsagePushBasedLogExportFeatureId{},
				},
			},
			expected: []cloudprovideraccess.FeatureUsage{
				{FeatureType: aws.String("ENCRYPTION_AT_REST")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, cloudprovideraccess.FeatureUsages(tc.usages))
		})
	}
}

func TestIsRoleNotFound(t *testing.T) {
	assert.True(t, cloudprovideraccess.IsRoleNotFound(nil, &http.Response{StatusCode: http.StatusNotFound}))
	assert.False(t, cloudprovideraccess.IsRoleNotFound(nil, &http.Response{StatusCode: http.StatusInternalServerError}))
	assert.False(t, cloudprovideraccess.IsRoleNotFound(nil, nil))
}
