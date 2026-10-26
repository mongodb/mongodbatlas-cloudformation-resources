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
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/cloud-provider-access-authorization/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

const (
	testProjectID = "66d979971ec97b7de1ef8777"
	testRoleID    = "6ad8f2f0202367917eff2e87"
	testARN       = "arn:aws:iam::123456789012:role/atlas-role"
	otherARN      = "arn:aws:iam::123456789012:role/other-role"
)

func TestCreateCheckpointsBeforeAuthorizingAndRetriesCallbacks(t *testing.T) {
	var (
		patches    int
		authorized bool
	)
	client := authorizationClient(t, func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "/api/atlas/v2/groups/"+testProjectID+"/cloudProviderAccess/"+testRoleID, req.URL.Path)
		switch req.Method {
		case http.MethodGet:
			if authorized {
				return roleResponse(testARN), nil
			}
			return roleResponse(""), nil
		case http.MethodPatch:
			patches++
			var body map[string]any
			require.NoError(t, json.NewDecoder(req.Body).Decode(&body))
			assert.Equal(t, testARN, body["iamAssumedRoleArn"])
			assert.Equal(t, "AWS", body["providerName"])
			if patches == 1 {
				return errorResponse(http.StatusBadRequest, cannotAssumeRole), nil
			}
			authorized = true
			return roleResponse(testARN), nil
		default:
			t.Fatalf("unexpected method: %s", req.Method)
			return nil, nil
		}
	})
	model := authorizationModel()
	first := resource.HandleCreate(handler.Request{}, client, model)
	require.Equal(t, handler.InProgress, first.OperationStatus)
	assertPendingIdentifiers(t, first)
	assert.Zero(t, patches, "checkpoint before changing Atlas")

	retry := resource.HandleCreate(callbackRequest(t, first), client, model)
	require.Equal(t, handler.InProgress, retry.OperationStatus)
	assertPendingIdentifiers(t, retry)
	assert.Equal(t, 1, patches, "only one PATCH per invocation")
	assert.Positive(t, retry.CallbackDelaySeconds)
	assert.Equal(t, first.CallbackContext["startedAt"], retry.CallbackContext["startedAt"])

	req := callbackRequest(t, retry)
	complete := resource.HandleCreate(req, client, model)
	require.Equal(t, handler.Success, complete.OperationStatus)
	assert.NotNil(t, model.AuthorizedDate)
	assert.Equal(t, 2, patches)

	// Replaying the callback after its successful response is lost must succeed,
	// whereas a different create for the same resource must still fail.
	replayed := resource.HandleCreate(req, client, model)
	assert.Equal(t, handler.Success, replayed.OperationStatus)
	assert.Equal(t, 2, patches)
	duplicate := resource.HandleCreate(handler.Request{}, client, authorizationModel())
	assert.Equal(t, string(types.HandlerErrorCodeAlreadyExists), duplicate.HandlerErrorCode)
}

func TestCreateRejectsAnyExistingAuthorization(t *testing.T) {
	for _, arn := range []string{testARN, otherARN} {
		t.Run(arn, func(t *testing.T) {
			var methods []string
			client := authorizationClient(t, func(req *http.Request) (*http.Response, error) {
				methods = append(methods, req.Method)
				return roleResponse(arn), nil
			})

			event := resource.HandleCreate(handler.Request{}, client, authorizationModel())

			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Equal(t, string(types.HandlerErrorCodeAlreadyExists), event.HandlerErrorCode)
			assert.Contains(t, event.Message, arn)
			assert.Equal(t, []string{http.MethodGet}, methods)
		})
	}
}

func TestCreateDoesNotReuseCallbackForAnotherTarget(t *testing.T) {
	for _, field := range []string{"operation", "projectId", "roleId", "profile", "arn"} {
		t.Run(field, func(t *testing.T) {
			authorized := false
			client := authorizationClient(t, func(req *http.Request) (*http.Response, error) {
				require.Equal(t, http.MethodGet, req.Method)
				if authorized {
					return roleResponse(testARN), nil
				}
				return roleResponse(""), nil
			})
			model := authorizationModel()
			pending := resource.HandleCreate(handler.Request{}, client, model)
			req := callbackRequest(t, pending)
			req.CallbackContext[field] = "another-target"
			authorized = true

			event := resource.HandleCreate(req, client, model)

			assert.Equal(t, string(types.HandlerErrorCodeAlreadyExists), event.HandlerErrorCode)
		})
	}
}

func TestAuthorizationCallbackTimesOutWithoutAnotherPatch(t *testing.T) {
	client := authorizationClient(t, func(req *http.Request) (*http.Response, error) {
		require.Equal(t, http.MethodGet, req.Method)
		return roleResponse(""), nil
	})
	model := authorizationModel()
	first := resource.HandleCreate(handler.Request{}, client, model)
	req := callbackRequest(t, first)
	req.CallbackContext["startedAt"] = time.Now().Add(-10 * time.Minute).UTC().Format(time.RFC3339Nano)

	event := resource.HandleCreate(req, client, model)

	assert.Equal(t, handler.Failed, event.OperationStatus)
	assert.Equal(t, string(types.HandlerErrorCodeServiceTimeout), event.HandlerErrorCode)
	assert.Same(t, model, event.ResourceModel)
}

func TestAuthorizationDoesNotRetryPermanentErrors(t *testing.T) {
	client := authorizationClient(t, func(req *http.Request) (*http.Response, error) {
		if req.Method == http.MethodGet {
			return roleResponse(""), nil
		}
		require.Equal(t, http.MethodPatch, req.Method)
		return errorResponse(http.StatusBadRequest, "INVALID_ATTRIBUTE"), nil
	})
	model := authorizationModel()
	first := resource.HandleCreate(handler.Request{}, client, model)

	event := resource.HandleCreate(callbackRequest(t, first), client, model)

	assert.Equal(t, handler.Failed, event.OperationStatus)
	assert.Equal(t, string(types.HandlerErrorCodeInvalidRequest), event.HandlerErrorCode)
	assert.Same(t, model, event.ResourceModel)
}

func TestCreateCallbackRetainsIdentifierWhenLookupFails(t *testing.T) {
	testCases := []struct {
		expectedCode types.HandlerErrorCode
		status       int
	}{
		{status: http.StatusInternalServerError, expectedCode: types.HandlerErrorCodeServiceInternalError},
		{status: http.StatusNotFound, expectedCode: types.HandlerErrorCodeNotFound},
	}
	for _, tc := range testCases {
		t.Run(http.StatusText(tc.status), func(t *testing.T) {
			lookupFails := false
			client := authorizationClient(t, func(req *http.Request) (*http.Response, error) {
				require.Equal(t, http.MethodGet, req.Method)
				if lookupFails {
					return errorResponse(tc.status, "UNEXPECTED_ERROR"), nil
				}
				return roleResponse(""), nil
			})
			model := authorizationModel()
			first := resource.HandleCreate(handler.Request{}, client, model)
			lookupFails = true

			event := resource.HandleCreate(callbackRequest(t, first), client, model)

			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Equal(t, string(tc.expectedCode), event.HandlerErrorCode)
			assert.Same(t, model, event.ResourceModel)
		})
	}
}

// assertPendingIdentifiers checks the contract the harness enforces while polling:
// an IN_PROGRESS response must carry a resourceModel with every primaryIdentifier
// property, so the next callback can address the same resource.
func assertPendingIdentifiers(t *testing.T, event handler.ProgressEvent) {
	t.Helper()
	model, ok := event.ResourceModel.(*resource.Model)
	require.True(t, ok, "IN_PROGRESS must carry the resource model")
	require.NotNil(t, model.ProjectId)
	require.NotNil(t, model.RoleId)
	require.NotNil(t, model.Profile)
	assert.Equal(t, testProjectID, *model.ProjectId)
	assert.Equal(t, testRoleID, *model.RoleId)
	assert.Equal(t, "custom-profile", *model.Profile)
}

func TestUpdateRequiresExistingAuthorization(t *testing.T) {
	for _, status := range []int{http.StatusOK, http.StatusNotFound} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			client := authorizationClient(t, func(req *http.Request) (*http.Response, error) {
				require.Equal(t, http.MethodGet, req.Method, "must not authorize a missing resource")
				if status == http.StatusNotFound {
					return errorResponse(status, roleNotFound), nil
				}
				return roleResponse(""), nil
			})

			event := resource.HandleUpdate(handler.Request{}, client, authorizationModel())

			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
		})
	}
}

func TestUpdateRetriesAndKeepsTheRegistration(t *testing.T) {
	patches := 0
	currentARN := otherARN
	client := authorizationClient(t, func(req *http.Request) (*http.Response, error) {
		if req.Method == http.MethodGet {
			return roleResponse(currentARN), nil
		}
		require.Equal(t, http.MethodPatch, req.Method)
		patches++
		if patches == 1 {
			return errorResponse(http.StatusBadRequest, cannotAssumeRole), nil
		}
		currentARN = testARN
		return roleResponse(currentARN), nil
	})
	model := authorizationModel()
	first := resource.HandleUpdate(handler.Request{}, client, model)
	require.Equal(t, handler.InProgress, first.OperationStatus)
	assertPendingIdentifiers(t, first)

	event := resource.HandleUpdate(callbackRequest(t, first), client, model)
	require.Equal(t, handler.Success, event.OperationStatus)
	assert.Equal(t, testRoleID, *model.RoleId)
	assert.Equal(t, testARN, *model.IamAssumedRoleArn)
	assert.Equal(t, 2, patches)

	replayed := resource.HandleUpdate(handler.Request{}, client, model)
	assert.Equal(t, handler.Success, replayed.OperationStatus)
	assert.Equal(t, 2, patches, "no patch when already at the desired ARN")
}

func TestDeleteRoleInUseIsAResourceConflict(t *testing.T) {
	client := authorizationClient(t, func(req *http.Request) (*http.Response, error) {
		if req.Method == http.MethodGet {
			return roleResponse(testARN), nil
		}
		require.Equal(t, http.MethodDelete, req.Method)
		return errorResponse(http.StatusBadRequest, roleInUse), nil
	})

	event := resource.HandleDelete(client, authorizationModel())

	assert.Equal(t, handler.Failed, event.OperationStatus)
	assert.Equal(t, string(types.HandlerErrorCodeResourceConflict), event.HandlerErrorCode)
	assert.Contains(t, event.Message, testRoleID)
}

func TestDeleteRemovesTheRegistration(t *testing.T) {
	var methods []string
	client := authorizationClient(t, func(req *http.Request) (*http.Response, error) {
		methods = append(methods, req.Method)
		if req.Method == http.MethodGet {
			return roleResponse(testARN), nil
		}
		require.Equal(t, http.MethodDelete, req.Method)
		assert.Contains(t, req.URL.Path, "/cloudProviderAccess/AWS/"+testRoleID)
		return jsonResponse(http.StatusNoContent, ""), nil
	})

	event := resource.HandleDelete(client, authorizationModel())

	assert.Equal(t, handler.Success, event.OperationStatus)
	assert.Equal(t, "Delete Complete", event.Message)
	assert.Nil(t, event.ResourceModel, "a delete response must not carry a model")
	assert.Equal(t, []string{http.MethodGet, http.MethodDelete}, methods)
}

func TestDeleteTreatsAnAlreadyRemovedRegistrationAsNotFound(t *testing.T) {
	client := authorizationClient(t, func(req *http.Request) (*http.Response, error) {
		if req.Method == http.MethodGet {
			return roleResponse(testARN), nil
		}
		require.Equal(t, http.MethodDelete, req.Method)
		return errorResponse(http.StatusNotFound, roleNotFound), nil
	})

	event := resource.HandleDelete(client, authorizationModel())

	assert.Equal(t, handler.Failed, event.OperationStatus)
	assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
}

func TestReadAndDeleteTreatUnapprovedRegistrationAsNotFound(t *testing.T) {
	client := authorizationClient(t, func(req *http.Request) (*http.Response, error) {
		require.Equal(t, http.MethodGet, req.Method)
		return roleResponse(""), nil
	})
	for _, event := range []handler.ProgressEvent{
		resource.HandleRead(client, authorizationModel()),
		resource.HandleDelete(client, authorizationModel()),
	} {
		assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
	}
}

func authorizationModel() *resource.Model {
	return &resource.Model{
		ProjectId: new(testProjectID), RoleId: new(testRoleID),
		Profile: new("custom-profile"), IamAssumedRoleArn: new(testARN),
	}
}

func callbackRequest(t *testing.T, event handler.ProgressEvent) handler.Request {
	t.Helper()
	// Round-trip JSON to exercise the representation CloudFormation sends back.
	data, err := json.Marshal(event.CallbackContext)
	require.NoError(t, err)
	var callback map[string]any
	require.NoError(t, json.Unmarshal(data, &callback))
	return handler.Request{CallbackContext: callback}
}

func authorizationClient(t *testing.T, transport roundTripFunc) *util.MongoDBClient {
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

func jsonResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func roleResponse(arn string) *http.Response {
	body := `{"providerName":"AWS","roleId":"` + testRoleID + `"`
	if arn != "" {
		body += `,"iamAssumedRoleArn":"` + arn + `","authorizedDate":"2026-10-21T12:00:00Z"`
	}
	return jsonResponse(http.StatusOK, body+`}`)
}

func errorResponse(status int, code string) *http.Response {
	return jsonResponse(status, `{"errorCode":"`+code+`","detail":"test error"}`)
}
