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
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
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

func TestIsAWSRole(t *testing.T) {
	assert.False(t, cloudprovideraccess.IsAWSRole(nil))
	assert.True(t, cloudprovideraccess.IsAWSRole(&admin.CloudProviderAccessRole{ProviderName: "AWS"}))
	assert.False(t, cloudprovideraccess.IsAWSRole(&admin.CloudProviderAccessRole{ProviderName: "AZURE"}))
}

func TestDeauthorizeRole(t *testing.T) {
	const (
		projectID = "66d979971ec97b7de1ef8777"
		roleID    = "6ad8f2f0202367917eff2e87"
	)

	testCases := []struct {
		expectedCode string
		name         string
		body         string
		status       int
		wantEvent    bool
		transportErr bool
	}{
		{
			name:   "A removed registration is a success",
			status: http.StatusNoContent,
		},
		{
			name:   "A registration that is already gone is a success",
			status: http.StatusNotFound,
			body:   `{"errorCode":"CLOUD_PROVIDER_ACCESS_ROLE_NOT_FOUND","detail":"not found"}`,
		},
		{
			name:         "An in-use registration is a conflict",
			status:       http.StatusBadRequest,
			body:         `{"errorCode":"AWS_IAM_ROLE_IN_USE","detail":"still in use"}`,
			wantEvent:    true,
			expectedCode: string(types.HandlerErrorCodeResourceConflict),
		},
		{
			name:         "An unrelated failure keeps its mapped code",
			status:       http.StatusInternalServerError,
			body:         `{"errorCode":"TEST_ERROR","detail":"failed"}`,
			wantEvent:    true,
			expectedCode: string(types.HandlerErrorCodeServiceInternalError),
		},
		{
			name:         "A transport failure is an internal error",
			transportErr: true,
			wantEvent:    true,
			expectedCode: string(types.HandlerErrorCodeInternalFailure),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := cloudProviderAccessClient(t, func(req *http.Request) (*http.Response, error) {
				require.Equal(t, http.MethodDelete, req.Method)
				require.Equal(t, "/api/atlas/v2/groups/"+projectID+"/cloudProviderAccess/AWS/"+roleID, req.URL.Path)
				if tc.transportErr {
					return nil, errors.New("connection refused")
				}
				return atlasResponse(tc.status, tc.body), nil
			})

			event := cloudprovideraccess.DeauthorizeRole(context.Background(), client, projectID, roleID)

			if !tc.wantEvent {
				assert.Nil(t, event)
				return
			}
			require.NotNil(t, event)
			assert.Equal(t, tc.expectedCode, event.HandlerErrorCode)
		})
	}
}

func cloudProviderAccessClient(t *testing.T, transport roundTripFunc) *util.MongoDBClient {
	t.Helper()
	client, err := admin.NewClient(
		admin.UseBaseURL("https://atlas.example.test/"),
		admin.UseHTTPClient(&http.Client{Transport: transport}),
	)
	require.NoError(t, err)
	return &util.MongoDBClient{AtlasSDK: client}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func atlasResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}
