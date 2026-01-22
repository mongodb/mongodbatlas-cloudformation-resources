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

	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-privatelink-endpoint/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"
)

func TestUpdateModel(t *testing.T) {
	testCases := map[string]struct {
		initialModel *resource.Model
		apiResp      *admin.StreamsPrivateLinkConnection
		expected     *resource.Model
	}{
		"nilApiResp": {
			initialModel: &resource.Model{
				ProjectId: util.StringPtr("project-123"),
			},
			apiResp: nil,
			expected: &resource.Model{
				ProjectId: util.StringPtr("project-123"),
			},
		},
		"fullApiResp": {
			initialModel: &resource.Model{
				ProjectId: util.StringPtr("project-123"),
			},
			apiResp: &admin.StreamsPrivateLinkConnection{
				Id:                    util.StringPtr("connection-123"),
				Provider:              "AWS",
				Vendor:                util.StringPtr("MSK"),
				Region:                util.StringPtr("us-east-1"),
				ServiceEndpointId:     util.StringPtr("vpce-123"),
				Arn:                   util.StringPtr("arn:aws:kafka:us-east-1:123456789012:cluster/msk-cluster/uuid"),
				DnsDomain:             util.StringPtr("example.com"),
				DnsSubDomain:          &[]string{"sub1", "sub2"},
				InterfaceEndpointId:   util.StringPtr("interface-123"),
				InterfaceEndpointName: util.StringPtr("interface-name"),
				ProviderAccountId:     util.StringPtr("123456789012"),
				State:                 util.StringPtr("DONE"),
				ErrorMessage:          util.StringPtr(""),
			},
			expected: &resource.Model{
				ProjectId:             util.StringPtr("project-123"),
				Id:                    util.StringPtr("connection-123"),
				ProviderName:          util.StringPtr("AWS"),
				Vendor:                util.StringPtr("MSK"),
				Region:                util.StringPtr("us-east-1"),
				ServiceEndpointId:     util.StringPtr("vpce-123"),
				Arn:                   util.StringPtr("arn:aws:kafka:us-east-1:123456789012:cluster/msk-cluster/uuid"),
				DnsDomain:             util.StringPtr("example.com"),
				DnsSubDomain:          []string{"sub1", "sub2"},
				InterfaceEndpointId:   util.StringPtr("interface-123"),
				InterfaceEndpointName: util.StringPtr("interface-name"),
				ProviderAccountId:     util.StringPtr("123456789012"),
				State:                 util.StringPtr("DONE"),
				ErrorMessage:          util.StringPtr(""),
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			resource.UpdateModel(tc.initialModel, tc.apiResp)
			assert.Equal(t, tc.expected.Id, tc.initialModel.Id)
			assert.Equal(t, tc.expected.ProviderName, tc.initialModel.ProviderName)
			assert.Equal(t, tc.expected.Vendor, tc.initialModel.Vendor)
			assert.Equal(t, tc.expected.Region, tc.initialModel.Region)
			assert.Equal(t, tc.expected.ServiceEndpointId, tc.initialModel.ServiceEndpointId)
			assert.Equal(t, tc.expected.Arn, tc.initialModel.Arn)
			assert.Equal(t, tc.expected.DnsDomain, tc.initialModel.DnsDomain)
			assert.Equal(t, tc.expected.DnsSubDomain, tc.initialModel.DnsSubDomain)
			assert.Equal(t, tc.expected.InterfaceEndpointId, tc.initialModel.InterfaceEndpointId)
			assert.Equal(t, tc.expected.InterfaceEndpointName, tc.initialModel.InterfaceEndpointName)
			assert.Equal(t, tc.expected.ProviderAccountId, tc.initialModel.ProviderAccountId)
			assert.Equal(t, tc.expected.State, tc.initialModel.State)
			assert.Equal(t, tc.expected.ErrorMessage, tc.initialModel.ErrorMessage)
		})
	}
}

func TestNewStreamPrivatelinkEndpointReq(t *testing.T) {
	testCases := map[string]struct {
		model    *resource.Model
		validate func(t *testing.T, req *admin.StreamsPrivateLinkConnection)
	}{
		"minimalRequest": {
			model: &resource.Model{
				ProviderName: util.StringPtr("AWS"),
				Vendor:       util.StringPtr("MSK"),
			},
			validate: func(t *testing.T, req *admin.StreamsPrivateLinkConnection) {
				t.Helper()
				assert.Equal(t, "AWS", req.Provider)
				assert.Equal(t, "MSK", *req.Vendor)
				assert.Nil(t, req.Region)
				assert.Nil(t, req.ServiceEndpointId)
				assert.Nil(t, req.Arn)
				assert.Nil(t, req.DnsDomain)
				assert.Nil(t, req.DnsSubDomain)
			},
		},
		"mskVendor": {
			model: &resource.Model{
				ProviderName: util.StringPtr("AWS"),
				Vendor:       util.StringPtr("MSK"),
				Arn:          util.StringPtr("arn:aws:kafka:us-east-1:123456789012:cluster/msk-cluster/uuid"),
			},
			validate: func(t *testing.T, req *admin.StreamsPrivateLinkConnection) {
				t.Helper()
				assert.Equal(t, "AWS", req.Provider)
				assert.Equal(t, "MSK", *req.Vendor)
				assert.Equal(t, "arn:aws:kafka:us-east-1:123456789012:cluster/msk-cluster/uuid", *req.Arn)
			},
		},
		"confluentVendor": {
			model: &resource.Model{
				ProviderName:      util.StringPtr("AWS"),
				Vendor:            util.StringPtr("CONFLUENT"),
				Region:            util.StringPtr("us-west-2"),
				ServiceEndpointId: util.StringPtr("vpce-456"),
				DnsDomain:         util.StringPtr("confluent.cloud"),
				DnsSubDomain:      []string{"sub1"},
			},
			validate: func(t *testing.T, req *admin.StreamsPrivateLinkConnection) {
				t.Helper()
				assert.Equal(t, "AWS", req.Provider)
				assert.Equal(t, "CONFLUENT", *req.Vendor)
				assert.Equal(t, "us-west-2", *req.Region)
				assert.Equal(t, "vpce-456", *req.ServiceEndpointId)
				assert.Equal(t, "confluent.cloud", *req.DnsDomain)
				require.NotNil(t, req.DnsSubDomain)
				assert.Equal(t, []string{"sub1"}, *req.DnsSubDomain)
			},
		},
		"s3Vendor": {
			model: &resource.Model{
				ProviderName:      util.StringPtr("AWS"),
				Vendor:            util.StringPtr("S3"),
				Region:            util.StringPtr("eu-west-1"),
				ServiceEndpointId: util.StringPtr("com.amazonaws.eu-west-1.s3"),
			},
			validate: func(t *testing.T, req *admin.StreamsPrivateLinkConnection) {
				t.Helper()
				assert.Equal(t, "AWS", req.Provider)
				assert.Equal(t, "S3", *req.Vendor)
				assert.Equal(t, "eu-west-1", *req.Region)
				assert.Equal(t, "com.amazonaws.eu-west-1.s3", *req.ServiceEndpointId)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := resource.NewStreamPrivatelinkEndpointReq(tc.model)
			require.NotNil(t, result)
			tc.validate(t, result)
		})
	}
}
