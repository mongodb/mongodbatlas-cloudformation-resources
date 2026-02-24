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

	"github.com/aws/smithy-go/ptr"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/service-account-secret/cmd/resource"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20250312014/admin"
)

func TestUpdateModelFromSecret(t *testing.T) {
	now := time.Now()
	secretID := "698eac1419a6b89540c8e7b5"
	secretValue := "mdb_sa_sk_test123" //nolint:gosec // test data
	maskedValue := "****7b5"
	orgID := "695feecd6de62c462a0f09cb"
	clientID := "mdb_sa_id_698eac1419a6b89540c8e7b6"

	model := &resource.Model{
		OrgId:    ptr.String(orgID),
		ClientId: ptr.String(clientID),
	}

	secret := &admin.ServiceAccountSecret{
		Id:                secretID,
		Secret:            ptr.String(secretValue),
		MaskedSecretValue: ptr.String(maskedValue),
		CreatedAt:         now,
		ExpiresAt:         now.Add(720 * time.Hour),
		LastUsedAt:        ptr.Time(now.Add(-1 * time.Hour)),
	}

	resource.UpdateModelFromSecret(model, secret)

	assert.Equal(t, secretID, *model.SecretId)
	assert.Equal(t, secretValue, *model.Secret)
	assert.Equal(t, maskedValue, *model.MaskedSecretValue)
	assert.NotNil(t, model.CreatedAt)
	assert.NotNil(t, model.ExpiresAt)
	assert.NotNil(t, model.LastUsedAt)
	assert.Equal(t, orgID, *model.OrgId)
	assert.Equal(t, clientID, *model.ClientId)
}
