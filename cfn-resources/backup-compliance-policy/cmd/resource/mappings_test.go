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

	"github.com/mongodb/mongodbatlas-cloudformation-resources/backup-compliance-policy/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20250312014/admin"
)

func TestPolicyItemConversions(t *testing.T) {
	t.Run("getOnDemandPolicyItem", func(t *testing.T) {
		assert.Nil(t, resource.GetOnDemandPolicyItem(nil))

		item := &admin.BackupComplianceOnDemandPolicyItem{
			Id:                util.StringPtr("id-1"),
			FrequencyType:     "ondemand",
			FrequencyInterval: 1,
			RetentionUnit:     "days",
			RetentionValue:    7,
		}
		result := resource.GetOnDemandPolicyItem(item)
		assert.Equal(t, "id-1", *result.Id)
		assert.Equal(t, 1, *result.FrequencyInterval)
		assert.Equal(t, 7, *result.RetentionValue)
	})

	t.Run("expandOnDemandPolicyItem", func(t *testing.T) {
		assert.Nil(t, resource.ExpandOnDemandPolicyItem(nil))

		freqInterval := 2
		retentionValue := 4
		model := &resource.OnDemandPolicyItem{
			FrequencyInterval: &freqInterval,
			RetentionUnit:     util.StringPtr("weeks"),
			RetentionValue:    &retentionValue,
		}
		result := resource.ExpandOnDemandPolicyItem(model)
		assert.Equal(t, "ondemand", result.FrequencyType)
		assert.Equal(t, 2, result.FrequencyInterval)
		assert.Equal(t, 4, result.RetentionValue)
	})

	t.Run("getScheduledPolicyItem", func(t *testing.T) {
		assert.Nil(t, resource.GetScheduledPolicyItem(nil))

		item := &admin.BackupComplianceScheduledPolicyItem{
			FrequencyType:     "hourly",
			FrequencyInterval: 6,
			RetentionUnit:     "days",
			RetentionValue:    3,
		}
		result := resource.GetScheduledPolicyItem(item)
		assert.Equal(t, "hourly", *result.FrequencyType)
		assert.Equal(t, 6, *result.FrequencyInterval)
	})

	t.Run("expandScheduledPolicyItem", func(t *testing.T) {
		freqInterval := 1
		retentionValue := 4
		model := &resource.ScheduledPolicyItem{
			FrequencyInterval: &freqInterval,
			RetentionUnit:     util.StringPtr("weeks"),
			RetentionValue:    &retentionValue,
		}
		result := resource.ExpandScheduledPolicyItem(model, "daily")
		assert.Equal(t, "daily", result.FrequencyType)
		assert.Equal(t, 1, result.FrequencyInterval)
	})
}

func TestSetBackupCompliancePolicyData(t *testing.T) {
	tests := map[string]struct {
		policy *admin.DataProtectionSettings20231001
		check  func(*testing.T, *resource.Model)
	}{
		"nil policy": {
			policy: nil,
			check: func(t *testing.T, m *resource.Model) {
				t.Helper()
				assert.Nil(t, m.ProjectId)
			},
		},
		"basic fields": {
			policy: &admin.DataProtectionSettings20231001{
				ProjectId:             util.StringPtr("proj-123"),
				AuthorizedEmail:       "admin@example.com",
				CopyProtectionEnabled: util.Pointer(true),
				RestoreWindowDays:     util.IntPtr(7),
				State:                 util.StringPtr("ACTIVE"),
			},
			check: func(t *testing.T, m *resource.Model) {
				t.Helper()
				assert.Equal(t, "proj-123", *m.ProjectId)
				assert.Equal(t, "admin@example.com", *m.AuthorizedEmail)
				assert.True(t, *m.CopyProtectionEnabled)
				assert.Equal(t, 7, *m.RestoreWindowDays)
				assert.Equal(t, "ACTIVE", *m.State)
			},
		},
		"with scheduled items": {
			policy: &admin.DataProtectionSettings20231001{
				ProjectId:       util.StringPtr("proj-456"),
				AuthorizedEmail: "test@example.com",
				ScheduledPolicyItems: &[]admin.BackupComplianceScheduledPolicyItem{
					{FrequencyType: "hourly", FrequencyInterval: 6, RetentionUnit: "days", RetentionValue: 3},
					{FrequencyType: "daily", FrequencyInterval: 1, RetentionUnit: "weeks", RetentionValue: 1},
					{FrequencyType: "weekly", FrequencyInterval: 1, RetentionUnit: "months", RetentionValue: 1},
				},
				State: util.StringPtr("ACTIVE"),
			},
			check: func(t *testing.T, m *resource.Model) {
				t.Helper()
				assert.NotNil(t, m.PolicyItemHourly)
				assert.NotNil(t, m.PolicyItemDaily)
				assert.Len(t, m.PolicyItemWeekly, 1)
				assert.Equal(t, "hourly", *m.PolicyItemHourly.FrequencyType)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			model := &resource.Model{}
			resource.SetBackupCompliancePolicyData(model, tc.policy)
			tc.check(t, model)
		})
	}
}

func TestExpandDataProtectionSettings(t *testing.T) {
	tests := map[string]struct {
		model *resource.Model
		check func(*testing.T, *admin.DataProtectionSettings20231001)
	}{
		"with empty policy items": {
			model: &resource.Model{
				AuthorizedEmail:    util.StringPtr("admin@example.com"),
				OnDemandPolicyItem: &resource.OnDemandPolicyItem{},
				PolicyItemHourly:   &resource.ScheduledPolicyItem{},
				PolicyItemDaily:    &resource.ScheduledPolicyItem{},
			},
			check: func(t *testing.T, s *admin.DataProtectionSettings20231001) {
				t.Helper()
				assert.Nil(t, s.OnDemandPolicyItem, "empty OnDemandPolicyItem should be omitted")
				assert.Nil(t, s.ScheduledPolicyItems, "empty policy items should not create scheduled items")
			},
		},
		"minimal model": {
			model: &resource.Model{
				AuthorizedEmail:         util.StringPtr("admin@example.com"),
				AuthorizedUserFirstName: util.StringPtr("Jane"),
				AuthorizedUserLastName:  util.StringPtr("Smith"),
			},
			check: func(t *testing.T, s *admin.DataProtectionSettings20231001) {
				t.Helper()
				assert.Equal(t, "admin@example.com", s.AuthorizedEmail)
				assert.Equal(t, "Jane", s.AuthorizedUserFirstName)
				assert.False(t, *s.CopyProtectionEnabled)
			},
		},
		"with booleans and integers": {
			model: &resource.Model{
				AuthorizedEmail:         util.StringPtr("test@example.com"),
				CopyProtectionEnabled:   util.Pointer(true),
				EncryptionAtRestEnabled: util.Pointer(true),
				PitEnabled:              util.Pointer(false),
				RestoreWindowDays:       util.IntPtr(14),
			},
			check: func(t *testing.T, s *admin.DataProtectionSettings20231001) {
				t.Helper()
				assert.True(t, *s.CopyProtectionEnabled)
				assert.True(t, *s.EncryptionAtRestEnabled)
				assert.False(t, *s.PitEnabled)
				assert.Equal(t, 14, *s.RestoreWindowDays)
			},
		},
		"with all policy items": {
			model: &resource.Model{
				AuthorizedEmail: util.StringPtr("admin@example.com"),
				PolicyItemHourly: &resource.ScheduledPolicyItem{
					FrequencyInterval: util.IntPtr(6),
					RetentionUnit:     util.StringPtr("days"),
					RetentionValue:    util.IntPtr(3),
				},
				PolicyItemDaily: &resource.ScheduledPolicyItem{
					FrequencyInterval: util.IntPtr(1),
					RetentionUnit:     util.StringPtr("weeks"),
					RetentionValue:    util.IntPtr(1),
				},
				PolicyItemWeekly: []resource.ScheduledPolicyItem{
					{
						FrequencyInterval: util.IntPtr(1),
						RetentionUnit:     util.StringPtr("months"),
						RetentionValue:    util.IntPtr(2),
					},
				},
			},
			check: func(t *testing.T, s *admin.DataProtectionSettings20231001) {
				t.Helper()
				assert.NotNil(t, s.ScheduledPolicyItems)
				items := *s.ScheduledPolicyItems
				assert.Len(t, items, 3)
				assert.Equal(t, "hourly", items[0].FrequencyType)
				assert.Equal(t, "daily", items[1].FrequencyType)
				assert.Equal(t, "weekly", items[2].FrequencyType)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := resource.ExpandDataProtectionSettings(tc.model, "test-project")
			assert.Equal(t, "test-project", *result.ProjectId)
			tc.check(t, result)
		})
	}
}

func TestRoundTripConversion(t *testing.T) {
	t.Run("policy survives round trip", func(t *testing.T) {
		original := &admin.DataProtectionSettings20231001{
			ProjectId:               util.StringPtr("proj-999"),
			AuthorizedEmail:         "security@example.com",
			AuthorizedUserFirstName: "Test",
			AuthorizedUserLastName:  "User",
			CopyProtectionEnabled:   util.Pointer(true),
			PitEnabled:              util.Pointer(false),
			RestoreWindowDays:       util.IntPtr(7),
			State:                   util.StringPtr("ACTIVE"),
			UpdatedDate:             &time.Time{},
			UpdatedUser:             util.StringPtr("admin"),
		}

		model := &resource.Model{}
		resource.SetBackupCompliancePolicyData(model, original)

		result := resource.ExpandDataProtectionSettings(model, "proj-999")

		assert.Equal(t, *original.ProjectId, *result.ProjectId)
		assert.Equal(t, original.AuthorizedEmail, result.AuthorizedEmail)
		assert.Equal(t, *original.CopyProtectionEnabled, *result.CopyProtectionEnabled)
		assert.Equal(t, *original.RestoreWindowDays, *result.RestoreWindowDays)
	})
}
