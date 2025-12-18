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

package resource

import (
	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

// GetStreamPrivatelinkEndpointModel maps Atlas API response to CFN Model
// CRITICAL: Preserves currentModel if provided to maintain primary identifier fields
func GetStreamPrivatelinkEndpointModel(apiResp *admin20250312010.StreamsPrivateLinkConnection, currentModel *Model) *Model {
	var model *Model
	if currentModel != nil {
		model = currentModel
	} else {
		model = new(Model)
	}

	// Populate from API response
	if apiResp.Id != nil {
		model.Id = apiResp.Id
	}
	if apiResp.Provider != "" {
		model.ProviderName = util.Pointer(apiResp.Provider)
	}
	if apiResp.Vendor != nil {
		model.Vendor = apiResp.Vendor
	}
	if apiResp.Region != nil {
		model.Region = apiResp.Region
	}
	if apiResp.ServiceEndpointId != nil {
		model.ServiceEndpointId = apiResp.ServiceEndpointId
	}
	if apiResp.Arn != nil {
		model.Arn = apiResp.Arn
	}
	if apiResp.DnsDomain != nil {
		model.DnsDomain = apiResp.DnsDomain
	}
	if apiResp.DnsSubDomain != nil {
		model.DnsSubDomain = *apiResp.DnsSubDomain
	}
	if apiResp.InterfaceEndpointId != nil {
		model.InterfaceEndpointId = apiResp.InterfaceEndpointId
	}
	if apiResp.InterfaceEndpointName != nil {
		model.InterfaceEndpointName = apiResp.InterfaceEndpointName
	}
	if apiResp.ProviderAccountId != nil {
		model.ProviderAccountId = apiResp.ProviderAccountId
	}
	if apiResp.State != nil {
		model.State = apiResp.State
	}
	if apiResp.ErrorMessage != nil {
		model.ErrorMessage = apiResp.ErrorMessage
	}

	// Preserve ProjectId from currentModel if present, otherwise it should be set by caller
	if model.ProjectId == nil && currentModel != nil {
		model.ProjectId = currentModel.ProjectId
	}

	// Preserve Profile from currentModel if present
	if currentModel != nil && currentModel.Profile != nil {
		model.Profile = currentModel.Profile
	}

	return model
}

// NewStreamPrivatelinkEndpointReq maps CFN Model to Atlas API request
func NewStreamPrivatelinkEndpointReq(model *Model) *admin20250312010.StreamsPrivateLinkConnection {
	req := &admin20250312010.StreamsPrivateLinkConnection{
		Provider: util.SafeString(model.ProviderName),
	}

	if model.Vendor != nil {
		req.Vendor = model.Vendor
	}
	if model.Region != nil {
		req.Region = model.Region
	}
	if model.ServiceEndpointId != nil {
		req.ServiceEndpointId = model.ServiceEndpointId
	}
	if model.Arn != nil {
		req.Arn = model.Arn
	}
	if model.DnsDomain != nil {
		req.DnsDomain = model.DnsDomain
	}
	if model.DnsSubDomain != nil && len(model.DnsSubDomain) > 0 {
		req.DnsSubDomain = &model.DnsSubDomain
	}

	return req
}
