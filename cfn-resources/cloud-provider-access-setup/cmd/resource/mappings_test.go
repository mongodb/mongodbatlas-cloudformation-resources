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

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/cloud-provider-access-setup/cmd/resource"
)

const (
	profile         = "default"
	projectID       = "66d979971ec97b7de1ef8777"
	roleID          = "6ad8f2f0202367917eff2e87"
	atlasAccountArn = "arn:aws:iam::111111111111:root"
	externalID      = "2466a342-0000-0000-0000-000000000000"
)

func createdDate() *time.Time {
	date := time.Date(2026, 10, 21, 12, 0, 0, 0, time.UTC)
	return &date
}

func TestNewCFNCloudProviderAccessSetup(t *testing.T) {
	testCases := []struct {
		name          string
		model         resource.Model
		role          *admin.CloudProviderAccessRole
		expectedModel resource.Model
	}{
		{
			name: "Maps the registration and its feature usages",
			model: resource.Model{
				Profile:      aws.String(profile),
				ProjectId:    aws.String(projectID),
				ProviderName: aws.String("AWS"),
			},
			role: &admin.CloudProviderAccessRole{
				RoleId:                     aws.String(roleID),
				AtlasAWSAccountArn:         aws.String(atlasAccountArn),
				AtlasAssumedRoleExternalId: aws.String(externalID),
				CreatedDate:                createdDate(),
				FeatureUsages: &[]admin.CloudProviderAccessFeatureUsage{
					{FeatureType: aws.String("ENCRYPTION_AT_REST")},
				},
			},
			expectedModel: resource.Model{
				Profile:                    aws.String(profile),
				ProjectId:                  aws.String(projectID),
				ProviderName:               aws.String("AWS"),
				RoleId:                     aws.String(roleID),
				AtlasAWSAccountArn:         aws.String(atlasAccountArn),
				AtlasAssumedRoleExternalId: aws.String(externalID),
				CreatedDate:                aws.String("2026-10-21T12:00:00Z"),
				FeatureUsages: []resource.FeatureUsage{
					{FeatureType: aws.String("ENCRYPTION_AT_REST")},
				},
			},
		},
		{
			name: "A role that no feature uses has no feature usages",
			model: resource.Model{
				Profile:      aws.String(profile),
				ProjectId:    aws.String(projectID),
				ProviderName: aws.String("AWS"),
			},
			role: &admin.CloudProviderAccessRole{
				RoleId:                     aws.String(roleID),
				AtlasAWSAccountArn:         aws.String(atlasAccountArn),
				AtlasAssumedRoleExternalId: aws.String(externalID),
				CreatedDate:                createdDate(),
			},
			expectedModel: resource.Model{
				Profile:                    aws.String(profile),
				ProjectId:                  aws.String(projectID),
				ProviderName:               aws.String("AWS"),
				RoleId:                     aws.String(roleID),
				AtlasAWSAccountArn:         aws.String(atlasAccountArn),
				AtlasAssumedRoleExternalId: aws.String(externalID),
				CreatedDate:                aws.String("2026-10-21T12:00:00Z"),
				FeatureUsages:              nil,
			},
		},
		{
			name: "An absent ProviderName defaults to the only supported provider",
			model: resource.Model{
				Profile:   aws.String(profile),
				ProjectId: aws.String(projectID),
			},
			role: &admin.CloudProviderAccessRole{
				RoleId: aws.String(roleID),
			},
			expectedModel: resource.Model{
				Profile:      aws.String(profile),
				ProjectId:    aws.String(projectID),
				ProviderName: aws.String("AWS"),
				RoleId:       aws.String(roleID),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := resource.NewCFNCloudProviderAccessSetup(&tc.model, tc.role)
			assert.Equal(t, tc.expectedModel, result)
		})
	}
}

func TestNewCFNCloudProviderAccessSetupFromAWSIAMRole(t *testing.T) {
	model := resource.Model{
		Profile:      aws.String(profile),
		ProjectId:    aws.String(projectID),
		ProviderName: aws.String("AWS"),
	}
	role := &admin.CloudProviderAccessAWSIAMRole{
		RoleId:                     aws.String(roleID),
		AtlasAWSAccountArn:         aws.String(atlasAccountArn),
		AtlasAssumedRoleExternalId: aws.String(externalID),
		CreatedDate:                createdDate(),
		FeatureUsages: &[]admin.CloudProviderAccessFeatureUsage{
			{FeatureType: aws.String("PUSH_BASED_LOG_EXPORT")},
		},
	}

	result := resource.NewCFNCloudProviderAccessSetupFromAWSIAMRole(&model, role)

	assert.Equal(t, resource.Model{
		Profile:                    aws.String(profile),
		ProjectId:                  aws.String(projectID),
		ProviderName:               aws.String("AWS"),
		RoleId:                     aws.String(roleID),
		AtlasAWSAccountArn:         aws.String(atlasAccountArn),
		AtlasAssumedRoleExternalId: aws.String(externalID),
		CreatedDate:                aws.String("2026-10-21T12:00:00Z"),
		FeatureUsages: []resource.FeatureUsage{
			{FeatureType: aws.String("PUSH_BASED_LOG_EXPORT")},
		},
	}, result)
}

func TestFindAWSRole(t *testing.T) {
	const otherRoleID = "111111111111111111111111"

	testCases := []struct {
		roles    *admin.CloudProviderAccessRoles
		expected *admin.CloudProviderAccessAWSIAMRole
		name     string
		roleID   string
	}{
		{
			name: "Finds a role among several",
			roles: &admin.CloudProviderAccessRoles{
				AwsIamRoles: &[]admin.CloudProviderAccessAWSIAMRole{
					{RoleId: aws.String(otherRoleID)},
					{RoleId: aws.String(roleID)},
				},
			},
			roleID:   roleID,
			expected: &admin.CloudProviderAccessAWSIAMRole{RoleId: aws.String(roleID)},
		},
		{
			name: "A role that is not registered is not found",
			roles: &admin.CloudProviderAccessRoles{
				AwsIamRoles: &[]admin.CloudProviderAccessAWSIAMRole{
					{RoleId: aws.String(otherRoleID)},
				},
			},
			roleID:   roleID,
			expected: nil,
		},
		{
			name:     "A response without AWS roles has no match",
			roles:    &admin.CloudProviderAccessRoles{},
			roleID:   roleID,
			expected: nil,
		},
		{
			name:     "A nil response has no match",
			roles:    nil,
			roleID:   roleID,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, resource.FindAWSRole(tc.roles, tc.roleID))
		})
	}
}
