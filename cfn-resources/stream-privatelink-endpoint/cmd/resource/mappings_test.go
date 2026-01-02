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

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-privatelink-endpoint/cmd/resource"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20250312010/admin"
)

func TestGetStreamPrivatelinkEndpointModel(t *testing.T) {
	projectID := "507f1f77bcf86cd799439011"
	connectionID := "507f1f77bcf86cd799439012"
	providerName := "AWS"
	vendor := "MSK"
	region := "us-east-1"
	arn := "arn:aws:kafka:us-east-1:123456789012:cluster/test-cluster/12345678-1234-1234-1234-123456789012-1"
	state := "DONE"
	dnsDomain := "test.example.com"
	dnsSubDomain := []string{"az1", "az2"}

	testCases := map[string]struct {
		apiResp      *admin.StreamsPrivateLinkConnection
		currentModel *resource.Model
		validateFunc func(t *testing.T, result *resource.Model)
	}{
		"withCurrentModel": {
			apiResp: &admin.StreamsPrivateLinkConnection{
				Id:       &connectionID,
				Provider: providerName,
				Vendor:   &vendor,
				Region:   &region,
				Arn:      &arn,
				State:    &state,
			},
			currentModel: &resource.Model{
				ProjectId:    &projectID,
				Profile:      util.StringPtr("default"),
				ProviderName: &providerName,
			},
			validateFunc: func(t *testing.T, result *resource.Model) {
				assert.Equal(t, connectionID, *result.Id)
				assert.Equal(t, providerName, *result.ProviderName)
				assert.Equal(t, vendor, *result.Vendor)
				assert.Equal(t, region, *result.Region)
				assert.Equal(t, arn, *result.Arn)
				assert.Equal(t, state, *result.State)
				assert.Equal(t, projectID, *result.ProjectId)
				assert.Equal(t, "default", *result.Profile)
			},
		},
		"withoutCurrentModel": {
			apiResp: &admin.StreamsPrivateLinkConnection{
				Id:       &connectionID,
				Provider: providerName,
				Vendor:   &vendor,
				Region:   &region,
			},
			currentModel: nil,
			validateFunc: func(t *testing.T, result *resource.Model) {
				assert.Equal(t, connectionID, *result.Id)
				assert.Equal(t, providerName, *result.ProviderName)
				assert.Equal(t, vendor, *result.Vendor)
				assert.Equal(t, region, *result.Region)
				assert.Nil(t, result.ProjectId)
			},
		},
		"withDnsFields": {
			apiResp: &admin.StreamsPrivateLinkConnection{
				Id:           &connectionID,
				Provider:     providerName,
				Vendor:       util.StringPtr("CONFLUENT"),
				DnsDomain:    &dnsDomain,
				DnsSubDomain: &dnsSubDomain,
			},
			currentModel: &resource.Model{
				ProjectId: &projectID,
			},
			validateFunc: func(t *testing.T, result *resource.Model) {
				assert.Equal(t, dnsDomain, *result.DnsDomain)
				assert.Equal(t, dnsSubDomain, result.DnsSubDomain)
			},
		},
		"withReadOnlyFields": {
			apiResp: &admin.StreamsPrivateLinkConnection{
				Id:                    &connectionID,
				Provider:              providerName,
				InterfaceEndpointId:   util.StringPtr("vpce-12345678"),
				InterfaceEndpointName: util.StringPtr("test-endpoint"),
				ProviderAccountId:     util.StringPtr("123456789012"),
				State:                 &state,
				ErrorMessage:          util.StringPtr("Test error"),
			},
			currentModel: &resource.Model{
				ProjectId: &projectID,
			},
			validateFunc: func(t *testing.T, result *resource.Model) {
				assert.Equal(t, "vpce-12345678", *result.InterfaceEndpointId)
				assert.Equal(t, "test-endpoint", *result.InterfaceEndpointName)
				assert.Equal(t, "123456789012", *result.ProviderAccountId)
				assert.Equal(t, state, *result.State)
				assert.Equal(t, "Test error", *result.ErrorMessage)
			},
		},
		"preservesCurrentModelFields": {
			apiResp: &admin.StreamsPrivateLinkConnection{
				Id:       &connectionID,
				Provider: providerName,
			},
			currentModel: &resource.Model{
				ProjectId:    &projectID,
				Profile:      util.StringPtr("custom-profile"),
				ProviderName: &providerName,
				Vendor:       &vendor,
			},
			validateFunc: func(t *testing.T, result *resource.Model) {
				// Should preserve ProjectId and Profile from currentModel
				assert.Equal(t, projectID, *result.ProjectId)
				assert.Equal(t, "custom-profile", *result.Profile)
				// Should update from API response
				assert.Equal(t, connectionID, *result.Id)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := resource.GetStreamPrivatelinkEndpointModel(tc.apiResp, tc.currentModel)
			if tc.validateFunc != nil {
				tc.validateFunc(t, result)
			}
		})
	}
}

func TestNewStreamPrivatelinkEndpointReq(t *testing.T) {
	providerName := "AWS"
	vendorMSK := "MSK"
	vendorConfluent := "CONFLUENT"
	vendorS3 := "S3"
	region := "us-east-1"
	arn := "arn:aws:kafka:us-east-1:123456789012:cluster/test-cluster/12345678-1234-1234-1234-123456789012-1"
	serviceEndpointId := "com.amazonaws.vpce.us-east-1.vpce-svc-12345678"
	dnsDomain := "test.example.com"
	dnsSubDomain := []string{"az1", "az2"}

	testCases := map[string]struct {
		model        *resource.Model
		validateFunc func(t *testing.T, result *admin.StreamsPrivateLinkConnection)
	}{
		"MSKVendor": {
			model: &resource.Model{
				ProviderName: &providerName,
				Vendor:       &vendorMSK,
				Arn:          &arn,
			},
			validateFunc: func(t *testing.T, result *admin.StreamsPrivateLinkConnection) {
				assert.Equal(t, providerName, result.Provider)
				assert.Equal(t, vendorMSK, *result.Vendor)
				assert.Equal(t, arn, *result.Arn)
				assert.Nil(t, result.Region)
				assert.Nil(t, result.ServiceEndpointId)
			},
		},
		"ConfluentVendor": {
			model: &resource.Model{
				ProviderName:      &providerName,
				Vendor:            &vendorConfluent,
				Region:            &region,
				ServiceEndpointId: &serviceEndpointId,
				DnsDomain:         &dnsDomain,
				DnsSubDomain:      dnsSubDomain,
			},
			validateFunc: func(t *testing.T, result *admin.StreamsPrivateLinkConnection) {
				assert.Equal(t, providerName, result.Provider)
				assert.Equal(t, vendorConfluent, *result.Vendor)
				assert.Equal(t, region, *result.Region)
				assert.Equal(t, serviceEndpointId, *result.ServiceEndpointId)
				assert.Equal(t, dnsDomain, *result.DnsDomain)
				assert.Equal(t, dnsSubDomain, *result.DnsSubDomain)
			},
		},
		"S3Vendor": {
			model: &resource.Model{
				ProviderName:      &providerName,
				Vendor:            &vendorS3,
				Region:            &region,
				ServiceEndpointId: &serviceEndpointId,
			},
			validateFunc: func(t *testing.T, result *admin.StreamsPrivateLinkConnection) {
				assert.Equal(t, providerName, result.Provider)
				assert.Equal(t, vendorS3, *result.Vendor)
				assert.Equal(t, region, *result.Region)
				assert.Equal(t, serviceEndpointId, *result.ServiceEndpointId)
			},
		},
		"nilDnsSubDomain": {
			model: &resource.Model{
				ProviderName: &providerName,
				Vendor:       &vendorConfluent,
				DnsSubDomain: nil,
			},
			validateFunc: func(t *testing.T, result *admin.StreamsPrivateLinkConnection) {
				assert.Nil(t, result.DnsSubDomain)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := resource.NewStreamPrivatelinkEndpointReq(tc.model)
			if tc.validateFunc != nil {
				tc.validateFunc(t, result)
			}
		})
	}
}
