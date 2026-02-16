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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/project-service-account-access-list-entry/cmd/resource"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"
)

func TestUpdateModelFromEntry(t *testing.T) {
	now := time.Now()
	projectID := "698e77e64b801087daabee06"
	clientID := "mdb_sa_id_698eac1419a6b89540c8e7b6"

	t.Run("CIDR Block entry", func(t *testing.T) {
		model := &resource.Model{
			ProjectId: ptr.String(projectID),
			ClientId:  ptr.String(clientID),
		}

		entry := &admin.ServiceAccountIPAccessListEntry{
			CidrBlock:       ptr.String("203.0.113.0/24"),
			CreatedAt:       ptr.Time(now),
			LastUsedAddress: ptr.String("203.0.113.42"),
			LastUsedAt:      ptr.Time(now.Add(-1 * time.Hour)),
			RequestCount:    ptr.Int(127),
		}

		resource.UpdateModelFromEntry(model, entry)

		assert.Equal(t, "203.0.113.0/24", *model.CIDRBlock)
		assert.Nil(t, model.IPAddress)
		assert.NotNil(t, model.CreatedAt)
		assert.Equal(t, "203.0.113.42", *model.LastUsedAddress)
		assert.NotNil(t, model.LastUsedAt)
		assert.Equal(t, 127, *model.RequestCount)
		assert.Equal(t, projectID, *model.ProjectId)
		assert.Equal(t, clientID, *model.ClientId)
	})

	t.Run("IP Address entry", func(t *testing.T) {
		model := &resource.Model{
			ProjectId: ptr.String(projectID),
			ClientId:  ptr.String(clientID),
		}

		entry := &admin.ServiceAccountIPAccessListEntry{
			IpAddress:    ptr.String("203.0.113.10"),
			CidrBlock:    ptr.String("203.0.113.10/32"),
			CreatedAt:    ptr.Time(now),
			RequestCount: ptr.Int(0),
		}

		resource.UpdateModelFromEntry(model, entry)

		assert.Equal(t, "203.0.113.10", *model.IPAddress)
		assert.Equal(t, "203.0.113.10/32", *model.CIDRBlock)
		assert.NotNil(t, model.CreatedAt)
		assert.Nil(t, model.LastUsedAddress)
		assert.Nil(t, model.LastUsedAt)
		assert.Equal(t, 0, *model.RequestCount)
		assert.Equal(t, projectID, *model.ProjectId)
		assert.Equal(t, clientID, *model.ClientId)
	})
}
