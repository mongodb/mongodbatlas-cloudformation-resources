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

	"github.com/mongodb/mongodbatlas-cloudformation-resources/log-integration/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"
)

func TestNewLogIntegrationCreateRequest(t *testing.T) {
	model := &resource.Model{
		Type:       util.StringPtr("S3"),
		BucketName: util.StringPtr("test-bucket"),
		IamRoleId:  util.StringPtr("arn:aws:iam::123456789012:role/test-role"),
		PrefixPath: util.StringPtr("/logs"),
		LogTypes:   []string{"AUDIT", "FTDC"},
		KmsKey:     util.StringPtr("arn:aws:kms:us-east-1:123456789012:key/test-key"),
	}

	result := resource.NewLogIntegrationCreateRequest(model)

	require.NotNil(t, result)
	assert.Equal(t, util.SafeString(model.Type), result.Type)
	assert.Equal(t, util.SafeString(model.BucketName), result.BucketName)
	assert.Equal(t, util.SafeString(model.IamRoleId), result.IamRoleId)
	assert.Equal(t, util.SafeString(model.PrefixPath), result.PrefixPath)
	assert.Equal(t, model.LogTypes, result.LogTypes)
	assert.Equal(t, util.StringPtr("arn:aws:kms:us-east-1:123456789012:key/test-key"), result.KmsKey)
}

func TestNewLogIntegrationUpdateRequest(t *testing.T) {
	model := &resource.Model{
		Type:       util.StringPtr("S3"),
		BucketName: util.StringPtr("test-bucket"),
		IamRoleId:  util.StringPtr("arn:aws:iam::123456789012:role/test-role"),
		PrefixPath: util.StringPtr("/logs"),
		LogTypes:   []string{"AUDIT", "PROFILER"},
		KmsKey:     util.StringPtr("arn:aws:kms:us-east-1:123456789012:key/test-key"),
	}

	result := resource.NewLogIntegrationUpdateRequest(model)

	require.NotNil(t, result)
	assert.Equal(t, util.SafeString(model.Type), result.Type)
	assert.Equal(t, model.BucketName, result.BucketName)
	assert.Equal(t, model.IamRoleId, result.IamRoleId)
	assert.Equal(t, model.PrefixPath, result.PrefixPath)
	require.NotNil(t, result.LogTypes)
	assert.Equal(t, model.LogTypes, *result.LogTypes)
	assert.Equal(t, util.StringPtr("arn:aws:kms:us-east-1:123456789012:key/test-key"), result.KmsKey)
}

func TestUpdateLogIntegrationModel(t *testing.T) {
	model := &resource.Model{}
	response := &admin.LogIntegrationResponse{
		Id:         "integration-123",
		Type:       "S3",
		BucketName: admin.PtrString("test-bucket"),
		IamRoleId:  admin.PtrString("arn:aws:iam::123456789012:role/test-role"),
		PrefixPath: admin.PtrString("/logs"),
		LogTypes:   &[]string{"AUDIT", "FTDC"},
		KmsKey:     admin.PtrString("arn:aws:kms:us-east-1:123456789012:key/test-key"),
	}

	resource.UpdateLogIntegrationModel(model, response)

	assert.Equal(t, &response.Id, model.IntegrationId)
	assert.Equal(t, &response.Type, model.Type)
	assert.Equal(t, response.BucketName, model.BucketName)
	assert.Equal(t, response.IamRoleId, model.IamRoleId)
	assert.Equal(t, response.PrefixPath, model.PrefixPath)
	assert.Equal(t, response.KmsKey, model.KmsKey)
	assert.Equal(t, *response.LogTypes, model.LogTypes)
}
