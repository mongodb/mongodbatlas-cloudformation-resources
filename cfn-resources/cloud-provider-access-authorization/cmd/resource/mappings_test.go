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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/cloud-provider-access-authorization/cmd/resource"
)

func TestFeatureUsagesFromRole(t *testing.T) {
	projectID := "66d979971ec97b7de1ef8777"
	bucketName := "mongodb-atlas-log-export"
	testCases := []struct {
		wantIDs   []*resource.FeatureId
		name      string
		wantTypes []string
		usages    []admin.CloudProviderAccessFeatureUsage
	}{
		{
			name:      "Maps feature types in order",
			usages:    []admin.CloudProviderAccessFeatureUsage{{FeatureType: aws.String("ENCRYPTION_AT_REST")}, {FeatureType: aws.String("PUSH_BASED_LOG_EXPORT")}},
			wantTypes: []string{"ENCRYPTION_AT_REST", "PUSH_BASED_LOG_EXPORT"},
			wantIDs:   []*resource.FeatureId{nil, nil},
		},
		{
			name: "Maps the project and bucket of an identified feature",
			usages: []admin.CloudProviderAccessFeatureUsage{{
				FeatureType: aws.String("PUSH_BASED_LOG_EXPORT"),
				FeatureId: &admin.CloudProviderAccessFeatureUsagePushBasedLogExportFeatureId{
					BucketName: aws.String(bucketName),
					GroupId:    aws.String(projectID),
				},
			}},
			wantTypes: []string{"PUSH_BASED_LOG_EXPORT"},
			wantIDs:   []*resource.FeatureId{{ProjectId: aws.String(projectID), BucketName: aws.String(bucketName)}},
		},
		{
			name: "Omits a feature id without identifiers",
			usages: []admin.CloudProviderAccessFeatureUsage{{
				FeatureType: aws.String("ENCRYPTION_AT_REST"),
				FeatureId:   &admin.CloudProviderAccessFeatureUsagePushBasedLogExportFeatureId{},
			}},
			wantTypes: []string{"ENCRYPTION_AT_REST"},
			wantIDs:   []*resource.FeatureId{nil},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			usages := resource.FeatureUsagesFromRole(&admin.CloudProviderAccessRole{FeatureUsages: &tc.usages})

			require.Len(t, usages, len(tc.wantTypes))
			for i, wantType := range tc.wantTypes {
				require.NotNil(t, usages[i].FeatureType)
				assert.Equal(t, wantType, *usages[i].FeatureType)
				assert.Equal(t, tc.wantIDs[i], usages[i].FeatureId)
			}
		})
	}
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
