// Copyright 2025 MongoDB Inc
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

	"go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/aws/smithy-go/ptr"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/org-service-account/cmd/resource"
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
			name: "Valid Input with Single Role",
			input: &resource.Model{
				Name:                    ptr.String("test-service-account"),
				Description:             ptr.String("Test description"),
				Roles:                   []string{"ORG_MEMBER"},
				SecretExpiresAfterHours: ptr.Int(720),
			},
			expected: &admin.OrgServiceAccountRequest{
				Name:                    "test-service-account",
				Description:             "Test description",
				Roles:                   []string{"ORG_MEMBER"},
				SecretExpiresAfterHours: 720,
			},
		},
		{
			name: "Valid Input with Multiple Roles - Should Sort",
			input: &resource.Model{
				Name:                    ptr.String("test-service-account"),
				Description:             ptr.String("Test description"),
				Roles:                   []string{"ORG_GROUP_CREATOR", "ORG_MEMBER"},
				SecretExpiresAfterHours: ptr.Int(360),
			},
			expected: &admin.OrgServiceAccountRequest{
				Name:                    "test-service-account",
				Description:             "Test description",
				Roles:                   []string{"ORG_GROUP_CREATOR", "ORG_MEMBER"}, // Sorted
				SecretExpiresAfterHours: 360,
			},
		},
		{
			name: "Valid Input with Unsorted Roles - Should Sort",
			input: &resource.Model{
				Name:                    ptr.String("test-service-account"),
				Description:             ptr.String("Test description"),
				Roles:                   []string{"ORG_MEMBER", "ORG_GROUP_CREATOR"},
				SecretExpiresAfterHours: ptr.Int(720),
			},
			expected: &admin.OrgServiceAccountRequest{
				Name:                    "test-service-account",
				Description:             "Test description",
				Roles:                   []string{"ORG_GROUP_CREATOR", "ORG_MEMBER"}, // Sorted
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
			assert.NotNil(t, actual)
			assert.Equal(t, tt.expected.Name, actual.Name)
			assert.Equal(t, tt.expected.Description, actual.Description)
			assert.Equal(t, tt.expected.SecretExpiresAfterHours, actual.SecretExpiresAfterHours)
			// Roles should be sorted
			assert.Equal(t, tt.expected.Roles, actual.Roles)
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
			name: "Valid Input with Single Role",
			input: &resource.Model{
				Name:        ptr.String("test-service-account"),
				Description: ptr.String("Updated description"),
				Roles:       []string{"ORG_MEMBER"},
			},
			expected: &admin.OrgServiceAccountUpdateRequest{
				Name:        ptr.String("test-service-account"),
				Description: ptr.String("Updated description"),
				Roles:       &[]string{"ORG_MEMBER"},
			},
		},
		{
			name: "Valid Input with Multiple Roles - Should Sort",
			input: &resource.Model{
				Name:        ptr.String("test-service-account"),
				Description: ptr.String("Updated description"),
				Roles:       []string{"ORG_GROUP_CREATOR", "ORG_MEMBER"},
			},
			expected: &admin.OrgServiceAccountUpdateRequest{
				Name:        ptr.String("test-service-account"),
				Description: ptr.String("Updated description"),
				Roles:       &[]string{"ORG_GROUP_CREATOR", "ORG_MEMBER"}, // Sorted
			},
		},
		{
			name: "Valid Input with Unsorted Roles - Should Sort",
			input: &resource.Model{
				Name:        ptr.String("test-service-account"),
				Description: ptr.String("Updated description"),
				Roles:       []string{"ORG_MEMBER", "ORG_GROUP_CREATOR"},
			},
			expected: &admin.OrgServiceAccountUpdateRequest{
				Name:        ptr.String("test-service-account"),
				Description: ptr.String("Updated description"),
				Roles:       &[]string{"ORG_GROUP_CREATOR", "ORG_MEMBER"}, // Sorted
			},
		},
		{
			name: "Empty Roles Array",
			input: &resource.Model{
				Name:        ptr.String("test-service-account"),
				Description: ptr.String("Updated description"),
				Roles:       []string{},
			},
			expected: &admin.OrgServiceAccountUpdateRequest{
				Name:        ptr.String("test-service-account"),
				Description: ptr.String("Updated description"),
				Roles:       nil, // Empty array should result in nil pointer
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
			assert.NotNil(t, actual)
			assert.Equal(t, tt.expected.Name, actual.Name)
			assert.Equal(t, tt.expected.Description, actual.Description)
			if tt.expected.Roles == nil {
				assert.Nil(t, actual.Roles)
			} else {
				assert.NotNil(t, actual.Roles)
				assert.Equal(t, *tt.expected.Roles, *actual.Roles)
			}
		})
	}
}

func TestGetOrgServiceAccountModel(t *testing.T) {
	now := time.Now()
	secretID := "secret-id-123"
	clientID := "mdb_sa_id_123456789"
	orgID := "63350255419cf25e3d511c95"

	tests := []struct {
		name        string
		inputSDK    *admin.OrgServiceAccount
		inputModel  *resource.Model
		expected    func(*testing.T, *resource.Model)
		description string
	}{
		{
			name:        "Nil SDK Input",
			inputSDK:    nil,
			inputModel:  nil,
			description: "Should return empty model when SDK input is nil",
			expected: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.NotNil(t, result)
			},
		},
		{
			name: "Valid SDK Input with Nil Current Model",
			inputSDK: &admin.OrgServiceAccount{
				ClientId:    ptr.String(clientID),
				Name:        ptr.String("test-service-account"),
				Description: ptr.String("Test description"),
				Roles:       &[]string{"ORG_MEMBER", "ORG_GROUP_CREATOR"},
				CreatedAt:   &now,
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
			inputModel:  nil,
			description: "Should map SDK response to model and sort roles",
			expected: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.NotNil(t, result)
				assert.Equal(t, clientID, *result.ClientId)
				assert.Equal(t, "test-service-account", *result.Name)
				assert.Equal(t, "Test description", *result.Description)
				// Roles should be sorted
				assert.Equal(t, []string{"ORG_GROUP_CREATOR", "ORG_MEMBER"}, result.Roles)
				assert.NotNil(t, result.CreatedAt)
				assert.NotNil(t, result.Secrets)
				assert.Len(t, result.Secrets, 1)
				assert.Equal(t, secretID, *result.Secrets[0].Id)
				assert.Equal(t, "secret-value", *result.Secrets[0].Secret)
			},
		},
		{
			name: "Valid SDK Input with Current Model - Preserve OrgId and Profile",
			inputSDK: &admin.OrgServiceAccount{
				ClientId:    ptr.String(clientID),
				Name:        ptr.String("test-service-account"),
				Description: ptr.String("Test description"),
				Roles:       &[]string{"ORG_MEMBER", "ORG_GROUP_CREATOR"},
				CreatedAt:   &now,
			},
			inputModel: &resource.Model{
				OrgId:   ptr.String(orgID),
				Profile: ptr.String("default"),
			},
			description: "Should preserve OrgId and Profile from current model",
			expected: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.NotNil(t, result)
				assert.Equal(t, orgID, *result.OrgId)
				assert.Equal(t, "default", *result.Profile)
				assert.Equal(t, clientID, *result.ClientId)
			},
		},
		{
			name: "Valid SDK Input with Current Model Roles - Preserve Input Order",
			inputSDK: &admin.OrgServiceAccount{
				ClientId:    ptr.String(clientID),
				Name:        ptr.String("test-service-account"),
				Description: ptr.String("Test description"),
				Roles:       &[]string{"ORG_MEMBER", "ORG_GROUP_CREATOR"},
				CreatedAt:   &now,
			},
			inputModel: &resource.Model{
				OrgId:   ptr.String(orgID),
				Profile: ptr.String("default"),
				Roles:   []string{"ORG_GROUP_CREATOR", "ORG_MEMBER"}, // Different order
			},
			description: "Should preserve roles order from current model for update operations",
			expected: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.NotNil(t, result)
				// Should preserve input order, not sort
				assert.Equal(t, []string{"ORG_GROUP_CREATOR", "ORG_MEMBER"}, result.Roles)
			},
		},
		{
			name: "Valid SDK Input with Secrets - All Fields",
			inputSDK: &admin.OrgServiceAccount{
				ClientId:  ptr.String(clientID),
				Name:      ptr.String("test-service-account"),
				CreatedAt: &now,
				Secrets: &[]admin.ServiceAccountSecret{
					{
						Id:                secretID,
						CreatedAt:         now,
						ExpiresAt:         now.Add(720 * time.Hour),
						LastUsedAt:        &now,
						MaskedSecretValue: ptr.String("****"),
						Secret:            ptr.String("secret-value"),
					},
				},
			},
			inputModel:  nil,
			description: "Should map all secret fields correctly",
			expected: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.NotNil(t, result)
				assert.NotNil(t, result.Secrets)
				assert.Len(t, result.Secrets, 1)
				secret := result.Secrets[0]
				assert.Equal(t, secretID, *secret.Id)
				assert.NotNil(t, secret.CreatedAt)
				assert.NotNil(t, secret.ExpiresAt)
				assert.NotNil(t, secret.LastUsedAt)
				assert.Equal(t, "****", *secret.MaskedSecretValue)
				assert.Equal(t, "secret-value", *secret.Secret)
			},
		},
		{
			name: "Valid SDK Input with Nil Secrets",
			inputSDK: &admin.OrgServiceAccount{
				ClientId:  ptr.String(clientID),
				Name:      ptr.String("test-service-account"),
				CreatedAt: &now,
				Secrets:   nil,
			},
			inputModel:  nil,
			description: "Should handle nil secrets gracefully",
			expected: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.NotNil(t, result)
				assert.Nil(t, result.Secrets)
			},
		},
		{
			name: "Valid SDK Input with Empty Roles",
			inputSDK: &admin.OrgServiceAccount{
				ClientId:  ptr.String(clientID),
				Name:      ptr.String("test-service-account"),
				Roles:     &[]string{},
				CreatedAt: &now,
			},
			inputModel:  nil,
			description: "Should handle empty roles array",
			expected: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.NotNil(t, result)
				assert.Empty(t, result.Roles)
			},
		},
		{
			name: "Valid SDK Input with Nil Roles",
			inputSDK: &admin.OrgServiceAccount{
				ClientId:  ptr.String(clientID),
				Name:      ptr.String("test-service-account"),
				Roles:     nil,
				CreatedAt: &now,
			},
			inputModel:  nil,
			description: "Should handle nil roles gracefully",
			expected: func(t *testing.T, result *resource.Model) {
				t.Helper()
				assert.NotNil(t, result)
				// When roles is nil in SDK, GetOrgServiceAccountModel only sets Roles if account.Roles != nil
				// So if account.Roles is nil, result.Roles remains as initialized (empty slice from new(Model))
				// The test verifies the function doesn't panic with nil roles
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resource.GetOrgServiceAccountModel(tt.inputSDK, tt.inputModel)
			tt.expected(t, result)
		})
	}
}
