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
	"encoding/json"
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/cluster/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func TestAddReplicationSpecIDs(t *testing.T) {
	testCases := map[string]struct {
		from        []admin20231115014.ReplicationSpec
		to          []admin20231115014.ReplicationSpec
		expectedIDs []string
	}{
		"emptyIsOk": {[]admin20231115014.ReplicationSpec{}, []admin20231115014.ReplicationSpec{}, []string{}},
		"zoneNameMatch": {
			[]admin20231115014.ReplicationSpec{{Id: util.StringPtr("id1"), ZoneName: util.StringPtr("z1")}},
			[]admin20231115014.ReplicationSpec{{ZoneName: util.StringPtr("z1")}},
			[]string{"id1"},
		},
		"providerRegionMatch": {
			[]admin20231115014.ReplicationSpec{{Id: util.StringPtr("id1"), RegionConfigs: regionConfig("AWS", "US_EAST_1")}},
			[]admin20231115014.ReplicationSpec{{RegionConfigs: regionConfig("AWS", "US_EAST_1")}},
			[]string{"id1"},
		},
		"noMatchRegion": {
			[]admin20231115014.ReplicationSpec{{Id: util.StringPtr("id1"), RegionConfigs: regionConfig("AWS", "US_EAST_1")}},
			[]admin20231115014.ReplicationSpec{{RegionConfigs: regionConfig("AWS", "US_EAST_2")}},
			[]string{""},
		},
		"noMatchZone": {
			[]admin20231115014.ReplicationSpec{{Id: util.StringPtr("id1"), ZoneName: util.StringPtr("z1")}},
			[]admin20231115014.ReplicationSpec{{ZoneName: util.StringPtr("z2")}},
			[]string{""},
		},
		"existingId": {
			[]admin20231115014.ReplicationSpec{{Id: util.StringPtr("id1"), ZoneName: util.StringPtr("z1")}},
			[]admin20231115014.ReplicationSpec{{Id: util.StringPtr("existing"), ZoneName: util.StringPtr("z1")}},
			[]string{"existing"},
		},
		"idMatchedOnlyOnce": {
			[]admin20231115014.ReplicationSpec{{Id: util.StringPtr("id1"), ZoneName: util.StringPtr("z1"), RegionConfigs: regionConfig("AWS", "US_EAST_1")}},
			[]admin20231115014.ReplicationSpec{{ZoneName: util.StringPtr("z1")}, {RegionConfigs: regionConfig("AWS", "US_EAST_1")}},
			[]string{"id1", ""},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			updated := resource.AddReplicationSpecIDs(tc.from, tc.to)
			ids := []string{}
			for _, spec := range *updated {
				ids = append(ids, spec.GetId())
			}
			assert.Equal(t, tc.expectedIDs, ids)
		})
	}
}

func regionConfig(provider, region string) *[]admin20231115014.CloudRegionConfig {
	return &[]admin20231115014.CloudRegionConfig{{
		RegionName:   &region,
		ProviderName: &provider,
	}}
}

func TestNewHardwareSpec(t *testing.T) {
	testCases := map[string]struct {
		spec     resource.Specs
		expected string
	}{
		"empty": {
			expected: `{}`,
			spec:     resource.Specs{},
		},
		"instanceSizeAndCount": {
			expected: `{"instanceSize":"M10","nodeCount":3}`,
			spec: resource.Specs{
				InstanceSize: util.StringPtr("M10"),
				NodeCount:    util.IntPtr(3),
			},
		},
		"allAttributes": {
			expected: `{"diskIOPS":100,"ebsVolumeType":"STANDARD","instanceSize":"M10","nodeCount":3}`,
			spec: resource.Specs{
				DiskIOPS:      util.StringPtr("100"),
				EbsVolumeType: util.StringPtr("STANDARD"),
				InstanceSize:  util.StringPtr("M10"),
				NodeCount:     util.IntPtr(3),
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			hardwareSpec := resource.NewHardwareSpec(&tc.spec)
			hwSpecJSON, err := json.Marshal(hardwareSpec)
			require.NoError(t, err)
			assert.JSONEq(t, tc.expected, string(hwSpecJSON))
		})
	}
}
