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
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/cloud-provider-access-authorization/cmd/resource"
)

const (
	cannotAssumeRole = "CANNOT_ASSUME_ROLE"
	roleInUse        = "AWS_IAM_ROLE_IN_USE"
	roleNotFound     = "CLOUD_PROVIDER_ACCESS_ROLE_NOT_FOUND"
)

// apiError builds the error the Atlas SDK returns for a failed request: a
// GenericOpenAPIError carrying the decoded ApiError and the formatted message.
func apiError(code string) error {
	e := &admin20231115014.GenericOpenAPIError{}
	e.SetError(fmt.Sprintf("PATCH /cloudProviderAccess: HTTP 400 (Error code: %q) Reason: Bad Request", code))
	e.SetModel(admin20231115014.ApiError{ErrorCode: &code})
	return e
}

func TestAuthorizeWithRetry_RetriesCannotAssumeRoleThenSucceeds(t *testing.T) {
	var (
		calls  int
		slept  []time.Duration
		delay  = 10 * time.Second
		role   = &admin20231115014.CloudProviderAccessRole{ProviderName: "AWS"}
		okResp = &http.Response{StatusCode: http.StatusOK}
	)

	got, resp, err := resource.AuthorizeWithRetry(3, delay, func(d time.Duration) { slept = append(slept, d) }, func() (*admin20231115014.CloudProviderAccessRole, *http.Response, error) {
		calls++
		if calls < 3 {
			return nil, nil, apiError(cannotAssumeRole)
		}
		return role, okResp, nil
	})

	require.NoError(t, err)
	assert.Same(t, role, got)
	assert.Same(t, okResp, resp)
	assert.Equal(t, 3, calls)
	assert.Equal(t, []time.Duration{delay, delay}, slept)
}

func TestAuthorizeWithRetry_StopsOnNonRetryableError(t *testing.T) {
	var (
		calls int
		slept int
	)

	got, _, err := resource.AuthorizeWithRetry(3, time.Second, func(time.Duration) { slept++ }, func() (*admin20231115014.CloudProviderAccessRole, *http.Response, error) {
		calls++
		return nil, &http.Response{StatusCode: http.StatusBadRequest}, apiError(roleInUse)
	})

	require.Error(t, err)
	assert.Nil(t, got)
	assert.Equal(t, 1, calls)
	assert.Equal(t, 0, slept)
}

func TestAuthorizeWithRetry_ExhaustsAttempts(t *testing.T) {
	var (
		calls int
		slept int
	)

	got, _, err := resource.AuthorizeWithRetry(3, time.Second, func(time.Duration) { slept++ }, func() (*admin20231115014.CloudProviderAccessRole, *http.Response, error) {
		calls++
		return nil, nil, apiError(cannotAssumeRole)
	})

	require.Error(t, err)
	assert.Nil(t, got)
	assert.Equal(t, 3, calls)
	assert.Equal(t, 2, slept)
}

func TestAuthorizeWithRetry_SucceedsOnFirstAttempt(t *testing.T) {
	var (
		calls int
		slept int
	)

	_, _, err := resource.AuthorizeWithRetry(3, time.Second, func(time.Duration) { slept++ }, func() (*admin20231115014.CloudProviderAccessRole, *http.Response, error) {
		calls++
		return &admin20231115014.CloudProviderAccessRole{}, nil, nil
	})

	require.NoError(t, err)
	assert.Equal(t, 1, calls)
	assert.Equal(t, 0, slept)
}

func TestHasErrorCode(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		code     string
		expected bool
	}{
		{name: "nil error", err: nil, code: cannotAssumeRole, expected: false},
		{name: "decoded model code", err: apiError(cannotAssumeRole), code: cannotAssumeRole, expected: true},
		{name: "decoded model code mismatch", err: apiError(cannotAssumeRole), code: roleInUse, expected: false},
		{name: "message fallback", err: errors.New("HTTP 400 (Error code: \"AWS_IAM_ROLE_IN_USE\")"), code: roleInUse, expected: true},
		{name: "unrelated error", err: errors.New("connection reset"), code: cannotAssumeRole, expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, resource.HasErrorCode(tc.err, tc.code))
		})
	}
}

func TestIsAuthorized(t *testing.T) {
	authorizedDate := time.Date(2026, 10, 21, 12, 0, 0, 0, time.UTC)

	assert.False(t, resource.IsAuthorized(nil))
	assert.False(t, resource.IsAuthorized(&admin20231115014.CloudProviderAccessRole{}))
	assert.True(t, resource.IsAuthorized(&admin20231115014.CloudProviderAccessRole{AuthorizedDate: &authorizedDate}))
}

func TestFeatureUsagesFromRole_AlwaysNonNil(t *testing.T) {
	usages := resource.FeatureUsagesFromRole(&admin20231115014.CloudProviderAccessRole{})

	require.NotNil(t, usages)
	assert.Empty(t, usages)
}

func TestFeatureUsagesFromRole_MapsFeatureId(t *testing.T) {
	featureType := "PUSH_BASED_LOG_EXPORT"
	groupID := "66d979971ec97b7de1ef8777"
	bucketName := "atlas-logs-export"
	usages := resource.FeatureUsagesFromRole(&admin20231115014.CloudProviderAccessRole{
		FeatureUsages: &[]admin20231115014.CloudProviderAccessFeatureUsage{
			{
				FeatureType: &featureType,
				FeatureId: &admin20231115014.CloudProviderAccessFeatureUsagePushBasedLogExportFeatureId{
					GroupId:    &groupID,
					BucketName: &bucketName,
				},
			},
		},
	})

	require.Len(t, usages, 1)
	require.NotNil(t, usages[0].FeatureType)
	assert.Equal(t, featureType, *usages[0].FeatureType)
	require.NotNil(t, usages[0].FeatureId)
	require.NotNil(t, usages[0].FeatureId.GroupId)
	assert.Equal(t, groupID, *usages[0].FeatureId.GroupId)
	require.NotNil(t, usages[0].FeatureId.BucketName)
	assert.Equal(t, bucketName, *usages[0].FeatureId.BucketName)
}

func TestSetModelFromRole_KeepsIdentifierAndSetsReadOnly(t *testing.T) {
	projectID := "66d979971ec97b7de1ef8777"
	roleID := "6ad8f2f0202367917eff2e87"
	arn := "arn:aws:iam::123456789012:role/atlas-role"
	authorizedDate := time.Date(2026, 10, 21, 12, 0, 0, 0, time.UTC)

	model := &resource.Model{ProjectId: &projectID, RoleId: &roleID, IamAssumedRoleArn: &arn}
	role := &admin20231115014.CloudProviderAccessRole{
		ProviderName:      "AWS",
		RoleId:            &roleID,
		IamAssumedRoleArn: &arn,
		AuthorizedDate:    &authorizedDate,
	}

	resource.SetModelFromRole(model, role)

	require.NotNil(t, model.ProjectId)
	assert.Equal(t, projectID, *model.ProjectId)
	require.NotNil(t, model.RoleId)
	assert.Equal(t, roleID, *model.RoleId)
	require.NotNil(t, model.AuthorizedDate)
	assert.Equal(t, authorizedDate.UTC().Format(time.RFC3339Nano), *model.AuthorizedDate)
	assert.NotNil(t, model.FeatureUsages)
}
