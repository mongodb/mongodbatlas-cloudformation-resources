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
	"time"

	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/aws/smithy-go/ptr"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/project-service-account/cmd/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGroupServiceAccountCreateReq(t *testing.T) {
	tests := []struct {
		input     *resource.Model
		expected  *admin.GroupServiceAccountRequest
		name      string
		expectErr bool
	}{
		{
			name:      "Nil Input",
			input:     nil,
			expected:  nil,
			expectErr: true,
		},
		{
			name: "Valid Input - Roles Sorted",
			input: &resource.Model{
				Name:                    ptr.String("test"),
				Description:             ptr.String("desc"),
				Roles:                   []string{"GROUP_OWNER", "GROUP_READ_ONLY"},
				SecretExpiresAfterHours: ptr.Int(720),
			},
			expected: &admin.GroupServiceAccountRequest{
				Name:                    "test",
				Description:             "desc",
				Roles:                   []string{"GROUP_OWNER", "GROUP_READ_ONLY"},
				SecretExpiresAfterHours: 720,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := resource.NewGroupServiceAccountCreateReq(tt.input)
			if tt.expectErr {
				require.Error(t, err)
				assert.Nil(t, actual)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.expected.Name, actual.Name)
			assert.Equal(t, tt.expected.Description, actual.Description)
			assert.Equal(t, tt.expected.Roles, actual.Roles)
			assert.Equal(t, tt.expected.SecretExpiresAfterHours, actual.SecretExpiresAfterHours)
		})
	}
}

func TestNewGroupServiceAccountUpdateReq(t *testing.T) {
	tests := []struct {
		input     *resource.Model
		expected  *admin.GroupServiceAccountUpdateRequest
		name      string
		expectErr bool
	}{
		{
			name:      "Nil Input",
			input:     nil,
			expected:  nil,
			expectErr: true,
		},
		{
			name: "Valid Input - Roles Sorted",
			input: &resource.Model{
				Name:        ptr.String("test"),
				Description: ptr.String("desc"),
				Roles:       []string{"GROUP_OWNER", "GROUP_READ_ONLY"},
			},
			expected: &admin.GroupServiceAccountUpdateRequest{
				Name:        ptr.String("test"),
				Description: ptr.String("desc"),
				Roles:       &[]string{"GROUP_OWNER", "GROUP_READ_ONLY"},
			},
		},
		{
			name: "Empty Roles",
			input: &resource.Model{
				Name:        ptr.String("test"),
				Description: ptr.String("desc"),
				Roles:       []string{},
			},
			expected: &admin.GroupServiceAccountUpdateRequest{
				Name:        ptr.String("test"),
				Description: ptr.String("desc"),
				Roles:       nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := resource.NewGroupServiceAccountUpdateReq(tt.input)
			if tt.expectErr {
				require.Error(t, err)
				assert.Nil(t, actual)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.expected.Name, actual.Name)
			assert.Equal(t, tt.expected.Description, actual.Description)
			if tt.expected.Roles == nil {
				assert.Nil(t, actual.Roles)
			} else {
				assert.Equal(t, *tt.expected.Roles, *actual.Roles)
			}
		})
	}
}

func TestGetGroupServiceAccountModel(t *testing.T) {
	now := time.Now()
	clientID := "mdb_sa_id_123"
	projectID := "63350255419cf25e3d511c95"
	secretID := "secret-123"

	tests := []struct {
		inputSDK   *admin.GroupServiceAccount
		inputModel *resource.Model
		validate   func(*testing.T, *resource.Model)
		name       string
	}{
		{
			name:       "Nil SDK Input",
			inputSDK:   nil,
			inputModel: nil,
			validate: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.NotNil(t, result)
			},
		},
		{
			name: "Valid SDK Input - Preserve ProjectId/Profile",
			inputSDK: &admin.GroupServiceAccount{
				ClientId:    ptr.String(clientID),
				Name:        ptr.String("test"),
				Description: ptr.String("desc"),
				Roles:       &[]string{"GROUP_OWNER"},
				CreatedAt:   &now,
			},
			inputModel: &resource.Model{
				ProjectId: ptr.String(projectID),
				Profile:   ptr.String("default"),
			},
			validate: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.Equal(t, projectID, *result.ProjectId)
				assert.Equal(t, "default", *result.Profile)
				assert.Equal(t, clientID, *result.ClientId)
			},
		},
		{
			name: "Valid SDK Input - Preserve Roles Order",
			inputSDK: &admin.GroupServiceAccount{
				ClientId:  ptr.String(clientID),
				Name:      ptr.String("test"),
				Roles:     &[]string{"GROUP_OWNER", "GROUP_READ_ONLY"},
				CreatedAt: &now,
			},
			inputModel: &resource.Model{
				Roles: []string{"GROUP_READ_ONLY", "GROUP_OWNER"},
			},
			validate: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.Equal(t, []string{"GROUP_READ_ONLY", "GROUP_OWNER"}, result.Roles)
			},
		},
		{
			name: "Valid SDK Input - With Secrets",
			inputSDK: &admin.GroupServiceAccount{
				ClientId:  ptr.String(clientID),
				Name:      ptr.String("test"),
				CreatedAt: &now,
				Secrets: &[]admin.ServiceAccountSecret{
					{
						Id:                secretID,
						CreatedAt:         now,
						ExpiresAt:         now.Add(720 * time.Hour),
						MaskedSecretValue: ptr.String("****"),
						Secret:            ptr.String("secret-value"),
					},
				},
			},
			inputModel: nil,
			validate: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.NotNil(t, result.Secrets)
				assert.Len(t, result.Secrets, 1)
				assert.Equal(t, secretID, *result.Secrets[0].Id)
				assert.Equal(t, "secret-value", *result.Secrets[0].Secret)
			},
		},
		{
			name: "Nil Secrets",
			inputSDK: &admin.GroupServiceAccount{
				ClientId:  ptr.String(clientID),
				Name:      ptr.String("test"),
				CreatedAt: &now,
				Secrets:   nil,
			},
			inputModel: nil,
			validate: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.Nil(t, result.Secrets)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resource.GetGroupServiceAccountModel(tt.inputSDK, tt.inputModel)
			tt.validate(t, result)
		})
	}
}
