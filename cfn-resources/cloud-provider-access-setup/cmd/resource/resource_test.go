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
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/cloud-provider-access-setup/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

func TestCreateCheckpointsRoleAndResumesWithoutRegisteringAgain(t *testing.T) {
	var requests []string
	client := cloudProviderAccessClient(t, func(req *http.Request) (*http.Response, error) {
		requests = append(requests, req.Method+" "+req.URL.Path)
		if req.Method == http.MethodPost {
			var body map[string]string
			require.NoError(t, json.NewDecoder(req.Body).Decode(&body))
			assert.Equal(t, map[string]string{"providerName": "AWS"}, body)
			return atlasResponse(http.StatusCreated, roleResponseJSON()), nil
		}
		require.Equal(t, http.MethodGet, req.Method)
		return atlasResponse(http.StatusOK, roleResponseJSON()), nil
	})
	model := resource.Model{
		Profile:      aws.String(profile),
		ProjectId:    aws.String(projectID),
		ProviderName: aws.String("AWS"),
	}

	created := resource.HandleCreate(nil, client, &model)
	require.Equal(t, handler.InProgress, created.OperationStatus)
	assert.Positive(t, created.CallbackDelaySeconds)
	assert.Equal(t, map[string]any{"roleId": roleID}, created.CallbackContext)
	checkpoint, ok := created.ResourceModel.(resource.Model)
	require.True(t, ok)
	assert.Equal(t, aws.String(roleID), checkpoint.RoleId)
	assert.Equal(t, aws.String(externalID), checkpoint.AtlasAssumedRoleExternalId)

	// Repeat the callback as well: every invocation must keep the checkpointed identity.
	for range 2 {
		resumed := resource.HandleCreate(created.CallbackContext, client, &checkpoint)
		require.Equal(t, handler.Success, resumed.OperationStatus)
		assert.Equal(t, "Create Complete", resumed.Message)
		assert.Equal(t, checkpoint, resumed.ResourceModel)
	}

	// CloudFormation may pass only the original properties plus the callback context, so the role id
	// must also travel in the callback context.
	resumed := resource.HandleCreate(created.CallbackContext, client, &resource.Model{
		Profile:      aws.String(profile),
		ProjectId:    aws.String(projectID),
		ProviderName: aws.String("AWS"),
	})
	require.Equal(t, handler.Success, resumed.OperationStatus)
	resumedModel, ok := resumed.ResourceModel.(resource.Model)
	require.True(t, ok)
	assert.Equal(t, aws.String(roleID), resumedModel.RoleId)

	assert.Equal(t, []string{
		"POST /api/atlas/v2/groups/" + projectID + "/cloudProviderAccess",
		"GET /api/atlas/v2/groups/" + projectID + "/cloudProviderAccess/" + roleID,
		"GET /api/atlas/v2/groups/" + projectID + "/cloudProviderAccess/" + roleID,
		"GET /api/atlas/v2/groups/" + projectID + "/cloudProviderAccess/" + roleID,
	}, requests)
}

func TestCreateWithKnownRoleDoesNotReplaceItWhenReadFails(t *testing.T) {
	testCases := []struct {
		name      string
		errorCode types.HandlerErrorCode
		status    int
	}{
		{"Missing registration", types.HandlerErrorCodeNotFound, http.StatusNotFound},
		{"Atlas failure", types.HandlerErrorCodeServiceInternalError, http.StatusInternalServerError},
		{"Access denied", types.HandlerErrorCodeAccessDenied, http.StatusUnauthorized},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var requests []string
			client := cloudProviderAccessClient(t, func(req *http.Request) (*http.Response, error) {
				requests = append(requests, req.Method+" "+req.URL.Path)
				return atlasResponse(tc.status, `{"errorCode":"TEST_ERROR","detail":"Read failed"}`), nil
			})
			model := resource.Model{
				Profile:      aws.String(profile),
				ProjectId:    aws.String(projectID),
				ProviderName: aws.String("AWS"),
				RoleId:       aws.String(roleID),
			}

			event := resource.HandleCreate(nil, client, &model)

			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Equal(t, string(tc.errorCode), event.HandlerErrorCode)
			assert.Equal(t, &model, event.ResourceModel, "retain the role ID for rollback")
			assert.Equal(t, []string{"GET /api/atlas/v2/groups/" + projectID + "/cloudProviderAccess/" + roleID}, requests)
		})
	}
}

func TestCreateFailsWhenRegistrationHasNoRoleId(t *testing.T) {
	client := cloudProviderAccessClient(t, func(_ *http.Request) (*http.Response, error) {
		return atlasResponse(http.StatusCreated, `{"providerName":"AWS"}`), nil
	})
	model := resource.Model{
		Profile:      aws.String(profile),
		ProjectId:    aws.String(projectID),
		ProviderName: aws.String("AWS"),
	}

	event := resource.HandleCreate(nil, client, &model)

	assert.Equal(t, handler.Failed, event.OperationStatus)
	assert.Equal(t, string(types.HandlerErrorCodeInternalFailure), event.HandlerErrorCode)
}

func TestCreateRejectsUnsupportedProviderBeforeAccessingCredentials(t *testing.T) {
	model := resource.Model{ProjectId: aws.String(projectID), ProviderName: aws.String("AZURE")}

	event, err := resource.Create(handler.Request{}, nil, &model)

	require.NoError(t, err)
	assert.Equal(t, handler.Failed, event.OperationStatus)
	assert.Equal(t, string(types.HandlerErrorCodeInvalidRequest), event.HandlerErrorCode)
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

func roleResponseJSON() string {
	return `{"providerName":"AWS","roleId":"` + roleID + `","atlasAWSAccountArn":"` + atlasAccountArn +
		`","atlasAssumedRoleExternalId":"` + externalID + `","createdDate":"2026-10-21T12:00:00Z"}`
}
