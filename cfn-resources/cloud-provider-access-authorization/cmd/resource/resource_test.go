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
	"github.com/stretchr/testify/require"

	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/cloud-provider-access-authorization/cmd/resource"
)

const (
	cannotAssumeRole = "CANNOT_ASSUME_ROLE"
	roleInUse        = "AWS_IAM_ROLE_IN_USE"
	roleNotFound     = "CLOUD_PROVIDER_ACCESS_ROLE_NOT_FOUND"
)

func TestIsAuthorized(t *testing.T) {
	authorizedDate := time.Date(2026, 10, 21, 12, 0, 0, 0, time.UTC)

	assert.False(t, resource.IsAuthorized(nil))
	assert.False(t, resource.IsAuthorized(&admin.CloudProviderAccessRole{}))
	assert.True(t, resource.IsAuthorized(&admin.CloudProviderAccessRole{AuthorizedDate: &authorizedDate}))
}

func TestFeatureUsagesFromRole_MapsFeatureType(t *testing.T) {
	encryptionAtRest := "ENCRYPTION_AT_REST"
	logExport := "PUSH_BASED_LOG_EXPORT"
	usages := resource.FeatureUsagesFromRole(&admin.CloudProviderAccessRole{
		FeatureUsages: &[]admin.CloudProviderAccessFeatureUsage{
			{FeatureType: &encryptionAtRest},
			{FeatureType: &logExport},
		},
	})

	require.Len(t, usages, 2)
	require.NotNil(t, usages[0].FeatureType)
	assert.Equal(t, encryptionAtRest, *usages[0].FeatureType)
	require.NotNil(t, usages[1].FeatureType)
	assert.Equal(t, logExport, *usages[1].FeatureType)
}

func TestSetModelFromRole_KeepsIdentifierAndSetsReadOnly(t *testing.T) {
	projectID := "66d979971ec97b7de1ef8777"
	roleID := "6ad8f2f0202367917eff2e87"
	arn := "arn:aws:iam::123456789012:role/atlas-role"
	authorizedDate := time.Date(2026, 10, 21, 12, 0, 0, 0, time.UTC)

	model := &resource.Model{ProjectId: &projectID, RoleId: &roleID, IamAssumedRoleArn: &arn}
	featureType := "ENCRYPTION_AT_REST"
	role := &admin.CloudProviderAccessRole{
		ProviderName:      "AWS",
		RoleId:            &roleID,
		IamAssumedRoleArn: &arn,
		AuthorizedDate:    &authorizedDate,
		FeatureUsages:     &[]admin.CloudProviderAccessFeatureUsage{{FeatureType: &featureType}},
	}

	resource.SetModelFromRole(model, role)

	require.NotNil(t, model.ProjectId)
	assert.Equal(t, projectID, *model.ProjectId)
	require.NotNil(t, model.RoleId)
	assert.Equal(t, roleID, *model.RoleId)
	require.NotNil(t, model.AuthorizedDate)
	assert.Equal(t, authorizedDate.UTC().Format(time.RFC3339Nano), *model.AuthorizedDate)
	require.Len(t, model.FeatureUsages, 1)
	require.NotNil(t, model.FeatureUsages[0].FeatureType)
	assert.Equal(t, featureType, *model.FeatureUsages[0].FeatureType)
}
