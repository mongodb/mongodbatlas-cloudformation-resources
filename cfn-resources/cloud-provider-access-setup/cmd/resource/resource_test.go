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

func TestCreateRegistersTheRole(t *testing.T) {
	var requests []string
	client := cloudProviderAccessClient(t, func(req *http.Request) (*http.Response, error) {
		requests = append(requests, req.Method+" "+req.URL.Path)
		require.Equal(t, http.MethodPost, req.Method)
		var body map[string]string
		require.NoError(t, json.NewDecoder(req.Body).Decode(&body))
		assert.Equal(t, map[string]string{"providerName": "AWS"}, body)
		return atlasResponse(http.StatusCreated, roleResponseJSON()), nil
	})
	model := resource.Model{
		Profile:   aws.String(profile),
		ProjectId: aws.String(projectID),
	}

	event := resource.HandleCreate(client, &model)

	require.Equal(t, handler.Success, event.OperationStatus)
	assert.Equal(t, "Create Complete", event.Message)
	registered, ok := event.ResourceModel.(resource.Model)
	require.True(t, ok)
	assert.Equal(t, aws.String(roleID), registered.RoleId)
	assert.Equal(t, aws.String(atlasAccountArn), registered.AtlasAWSAccountArn)
	assert.Equal(t, aws.String(externalID), registered.AtlasAssumedRoleExternalId)
	assert.Equal(t, []string{"POST /api/atlas/v2/groups/" + projectID + "/cloudProviderAccess"}, requests)
}

func TestCreateFailsWhenRegistrationHasNoRoleId(t *testing.T) {
	client := cloudProviderAccessClient(t, func(_ *http.Request) (*http.Response, error) {
		return atlasResponse(http.StatusCreated, `{"providerName":"AWS"}`), nil
	})
	model := resource.Model{
		Profile:   aws.String(profile),
		ProjectId: aws.String(projectID),
	}

	event := resource.HandleCreate(client, &model)

	assert.Equal(t, handler.Failed, event.OperationStatus)
	assert.Equal(t, string(types.HandlerErrorCodeInternalFailure), event.HandlerErrorCode)
}

func TestReadRejectsNonAWSRegistrations(t *testing.T) {
	for _, provider := range []string{"AZURE", "GCP"} {
		t.Run(provider, func(t *testing.T) {
			client := cloudProviderAccessClient(t, func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, http.MethodGet, req.Method)
				assert.Equal(t, "/api/atlas/v2/groups/"+projectID+"/cloudProviderAccess/"+roleID, req.URL.Path)
				idField := "roleId"
				if provider == "AZURE" {
					idField = "_id"
				}
				return atlasResponse(http.StatusOK, `{"providerName":"`+provider+`","`+idField+`":"`+roleID+`"}`), nil
			})
			model := resource.Model{
				Profile:   aws.String(profile),
				ProjectId: aws.String(projectID),
				RoleId:    aws.String(roleID),
			}

			event := resource.HandleRead(client, &model)

			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
			assert.Nil(t, event.ResourceModel)
		})
	}
}

func TestDeleteRegistration(t *testing.T) {
	const otherRoleID = "111111111111111111111111"

	listPath := "GET /api/atlas/v2/groups/" + projectID + "/cloudProviderAccess"
	deletePath := "DELETE /api/atlas/v2/groups/" + projectID + "/cloudProviderAccess/AWS/" + roleID

	testCases := []struct {
		name             string
		roleIDs          []string
		deleteBody       string
		expectedStatus   handler.Status
		expectedCode     string
		expectedContains string
		expectedRequests []string
		deleteStatus     int
	}{
		{
			name:             "A role that a feature still uses fails with ResourceConflict",
			roleIDs:          []string{roleID},
			deleteStatus:     http.StatusBadRequest,
			deleteBody:       `{"errorCode":"AWS_IAM_ROLE_IN_USE","detail":"Cannot remove the AWS IAM role because it is still in use."}`,
			expectedStatus:   handler.Failed,
			expectedCode:     string(types.HandlerErrorCodeResourceConflict),
			expectedContains: roleID,
			expectedRequests: []string{listPath, deletePath},
		},
		{
			name:             "A registration removed between the list and the delete is tolerated",
			roleIDs:          []string{roleID},
			deleteStatus:     http.StatusNotFound,
			deleteBody:       `{"errorCode":"CLOUD_PROVIDER_ACCESS_ROLE_NOT_FOUND","detail":"not found"}`,
			expectedStatus:   handler.Success,
			expectedRequests: []string{listPath, deletePath},
		},
		{
			name:             "A registration that is already gone is NotFound",
			roleIDs:          []string{otherRoleID},
			expectedStatus:   handler.Failed,
			expectedCode:     string(types.HandlerErrorCodeNotFound),
			expectedContains: roleID,
			expectedRequests: []string{listPath},
		},
		{
			name:             "A registration that is present is removed",
			roleIDs:          []string{otherRoleID, roleID},
			deleteStatus:     http.StatusNoContent,
			expectedStatus:   handler.Success,
			expectedRequests: []string{listPath, deletePath},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var requests []string
			client := cloudProviderAccessClient(t, func(req *http.Request) (*http.Response, error) {
				requests = append(requests, req.Method+" "+req.URL.Path)
				if req.Method == http.MethodGet {
					return atlasResponse(http.StatusOK, awsRolesJSON(tc.roleIDs...)), nil
				}
				return atlasResponse(tc.deleteStatus, tc.deleteBody), nil
			})
			model := resource.Model{
				Profile:   aws.String(profile),
				ProjectId: aws.String(projectID),
				RoleId:    aws.String(roleID),
			}

			event := resource.HandleDelete(client, &model)

			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			assert.Equal(t, tc.expectedCode, event.HandlerErrorCode)
			if tc.expectedContains != "" {
				assert.Contains(t, event.Message, tc.expectedContains)
			}
			assert.Equal(t, tc.expectedRequests, requests)
		})
	}
}

func TestListReturnsTheProjectRoles(t *testing.T) {
	const otherRoleID = "111111111111111111111111"

	client := cloudProviderAccessClient(t, func(req *http.Request) (*http.Response, error) {
		require.Equal(t, http.MethodGet, req.Method)
		return atlasResponse(http.StatusOK, awsRolesJSON(otherRoleID, roleID)), nil
	})
	model := resource.Model{
		Profile:   aws.String(profile),
		ProjectId: aws.String(projectID),
	}

	event := resource.HandleList(client, &model)

	require.Equal(t, handler.Success, event.OperationStatus)
	require.Len(t, event.ResourceModels, 2)

	first, ok := event.ResourceModels[0].(resource.Model)
	require.True(t, ok)
	assert.Equal(t, aws.String(otherRoleID), first.RoleId)
	assert.Equal(t, aws.String(profile), first.Profile)
	assert.Equal(t, aws.String(projectID), first.ProjectId)

	second, ok := event.ResourceModels[1].(resource.Model)
	require.True(t, ok)
	assert.Equal(t, aws.String(roleID), second.RoleId)
	assert.Equal(t, aws.String(externalID), second.AtlasAssumedRoleExternalId)
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

func awsRolesJSON(roleIDs ...string) string {
	roles := make([]string, 0, len(roleIDs))
	for _, id := range roleIDs {
		roles = append(roles, `{"providerName":"AWS","roleId":"`+id+`","atlasAWSAccountArn":"`+atlasAccountArn+
			`","atlasAssumedRoleExternalId":"`+externalID+`","createdDate":"2026-10-21T12:00:00Z"}`)
	}
	return `{"awsIamRoles":[` + strings.Join(roles, ",") + `]}`
}
