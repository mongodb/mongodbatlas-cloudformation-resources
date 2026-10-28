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
	assert.False(t, resource.IsAuthorized(&admin.CloudProviderAccessRole{ProviderName: "AWS"}))
	assert.False(t, resource.IsAuthorized(&admin.CloudProviderAccessRole{AuthorizedDate: &authorizedDate}))
	assert.False(t, resource.IsAuthorized(&admin.CloudProviderAccessRole{ProviderName: "AZURE", AuthorizedDate: &authorizedDate}))
	assert.True(t, resource.IsAuthorized(&admin.CloudProviderAccessRole{ProviderName: "AWS", AuthorizedDate: &authorizedDate}))
}
