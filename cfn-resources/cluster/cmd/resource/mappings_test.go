// Copyright 2024 MongoDB Inc
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

	"github.com/mongodb/mongodbatlas-cloudformation-resources/cluster/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func TestAsProviderRegion(t *testing.T) {
	testCases := map[string]struct {
		from        []admin.ReplicationSpec
		to          []admin.ReplicationSpec
		expectedIds []string
	}{
		"emptyIsOk": {[]admin.ReplicationSpec{}, []admin.ReplicationSpec{}, []string{}},
		"zoneNameMatch": {
			[]admin.ReplicationSpec{{Id: util.StringPtr("id1"), ZoneName: util.StringPtr("z1")}},
			[]admin.ReplicationSpec{{ZoneName: util.StringPtr("z1")}},
			[]string{"id1"},
		},
		"providerRegionMatch": {
			[]admin.ReplicationSpec{{Id: util.StringPtr("id1"), RegionConfigs: regionConfig("AWS", "US_EAST_1")}},
			[]admin.ReplicationSpec{{RegionConfigs: regionConfig("AWS", "US_EAST_1")}},
			[]string{"id1"},
		},
		"noMatchRegion": {
			[]admin.ReplicationSpec{{Id: util.StringPtr("id1"), RegionConfigs: regionConfig("AWS", "US_EAST_1")}},
			[]admin.ReplicationSpec{{RegionConfigs: regionConfig("AWS", "US_EAST_2")}},
			[]string{""},
		},
		"noMatchZone": {
			[]admin.ReplicationSpec{{Id: util.StringPtr("id1"), ZoneName: util.StringPtr("z1")}},
			[]admin.ReplicationSpec{{ZoneName: util.StringPtr("z2")}},
			[]string{""},
		},
		"existingId": {
			[]admin.ReplicationSpec{{Id: util.StringPtr("id1"), ZoneName: util.StringPtr("z1")}},
			[]admin.ReplicationSpec{{Id: util.StringPtr("existing"), ZoneName: util.StringPtr("z1")}},
			[]string{"existing"},
		},
		"idMatchedOnlyOnce": {
			[]admin.ReplicationSpec{{Id: util.StringPtr("id1"), ZoneName: util.StringPtr("z1"), RegionConfigs: regionConfig("AWS", "US_EAST_1")}},
			[]admin.ReplicationSpec{{ZoneName: util.StringPtr("z1")}, {RegionConfigs: regionConfig("AWS", "US_EAST_1")}},
			[]string{"id1", ""},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			updated := resource.PopulateReplicationSpecIds(tc.from, tc.to)
			ids := []string{}
			for _, spec := range *updated {
				ids = append(ids, spec.GetId())
			}
			assert.Equal(t, tc.expectedIds, ids)
		})
	}
}

func regionConfig(provider, region string) *[]admin.CloudRegionConfig {
	return &[]admin.CloudRegionConfig{{
		RegionName:   &region,
		ProviderName: &provider,
	}}
}
