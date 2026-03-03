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

	"go.mongodb.org/atlas-sdk/v20250312014/admin"

	"github.com/aws/smithy-go/ptr"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/service-account/cmd/resource"
	"github.com/stretchr/testify/assert"
)

func TestNewOrgServiceAccountCreateReq(t *testing.T) {
	tests := []struct {
		input    *resource.Model
		expected *admin.OrgServiceAccountRequest
		name     string
	}{
		{
			name:     "Nil Input",
			input:    nil,
			expected: nil,
		},
		{
			name: "Valid Input - Roles Sorted",
			input: &resource.Model{
				Name:                    ptr.String("test"),
				Description:             ptr.String("desc"),
				Roles:                   []string{"ORG_MEMBER", "ORG_GROUP_CREATOR"},
				SecretExpiresAfterHours: ptr.Int(720),
			},
			expected: &admin.OrgServiceAccountRequest{
				Name:                    "test",
				Description:             "desc",
				Roles:                   []string{"ORG_GROUP_CREATOR", "ORG_MEMBER"},
				SecretExpiresAfterHours: 720,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := resource.NewOrgServiceAccountCreateReq(tt.input)
			if tt.expected == nil {
				assert.Nil(t, actual)
				return
			}
			assert.Equal(t, tt.expected.Name, actual.Name)
			assert.Equal(t, tt.expected.Description, actual.Description)
			assert.Equal(t, tt.expected.Roles, actual.Roles)
			assert.Equal(t, tt.expected.SecretExpiresAfterHours, actual.SecretExpiresAfterHours)
		})
	}
}

func TestNewOrgServiceAccountUpdateReq(t *testing.T) {
	tests := []struct {
		input    *resource.Model
		expected *admin.OrgServiceAccountUpdateRequest
		name     string
	}{
		{
			name:     "Nil Input",
			input:    nil,
			expected: nil,
		},
		{
			name: "Valid Input - Roles Sorted",
			input: &resource.Model{
				Name:        ptr.String("test"),
				Description: ptr.String("desc"),
				Roles:       []string{"ORG_MEMBER", "ORG_GROUP_CREATOR"},
			},
			expected: &admin.OrgServiceAccountUpdateRequest{
				Name:        ptr.String("test"),
				Description: ptr.String("desc"),
				Roles:       &[]string{"ORG_GROUP_CREATOR", "ORG_MEMBER"},
			},
		},
		{
			name: "Empty Roles",
			input: &resource.Model{
				Name:        ptr.String("test"),
				Description: ptr.String("desc"),
				Roles:       []string{},
			},
			expected: &admin.OrgServiceAccountUpdateRequest{
				Name:        ptr.String("test"),
				Description: ptr.String("desc"),
				Roles:       nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := resource.NewOrgServiceAccountUpdateReq(tt.input)
			if tt.expected == nil {
				assert.Nil(t, actual)
				return
			}
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

func TestGetOrgServiceAccountModel(t *testing.T) {
	now := time.Now()
	clientID := "mdb_sa_id_123"
	orgID := "63350255419cf25e3d511c95"
	secretID := "secret-123"

	tests := []struct {
		inputSDK   *admin.OrgServiceAccount
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
			name: "Valid SDK Input - Preserve OrgId/Profile",
			inputSDK: &admin.OrgServiceAccount{
				ClientId:    ptr.String(clientID),
				Name:        ptr.String("test"),
				Description: ptr.String("desc"),
				Roles:       &[]string{"ORG_MEMBER"},
				CreatedAt:   &now,
			},
			inputModel: &resource.Model{
				OrgId:   ptr.String(orgID),
				Profile: ptr.String("default"),
			},
			validate: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.Equal(t, orgID, *result.OrgId)
				assert.Equal(t, "default", *result.Profile)
				assert.Equal(t, clientID, *result.ClientId)
			},
		},
		{
			name: "Valid SDK Input - Preserve Roles Order",
			inputSDK: &admin.OrgServiceAccount{
				ClientId:  ptr.String(clientID),
				Name:      ptr.String("test"),
				Roles:     &[]string{"ORG_MEMBER", "ORG_GROUP_CREATOR"},
				CreatedAt: &now,
			},
			inputModel: &resource.Model{
				Roles: []string{"ORG_GROUP_CREATOR", "ORG_MEMBER"},
			},
			validate: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.Equal(t, []string{"ORG_GROUP_CREATOR", "ORG_MEMBER"}, result.Roles)
			},
		},
		{
			name: "Valid SDK Input - With Secrets",
			inputSDK: &admin.OrgServiceAccount{
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
			inputSDK: &admin.OrgServiceAccount{
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
			result := resource.GetOrgServiceAccountModel(tt.inputSDK, tt.inputModel)
			tt.validate(t, result)
		})
	}
}
