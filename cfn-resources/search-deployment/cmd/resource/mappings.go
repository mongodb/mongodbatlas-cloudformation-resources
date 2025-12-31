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

func NewCFNSearchDeployment(prevModel *Model, apiResp *admin20250312010.ApiSearchDeploymentResponse) Model {
	respSpecs := apiResp.GetSpecs()
	resultSpecs := make([]ApiSearchDeploymentSpec, len(respSpecs))
	for i := range respSpecs {
		instanceSize := respSpecs[i].InstanceSize
		// Follow cluster pattern: directly assign NodeCount from API response
		// Reference: mongodbatlas-cloudformation-resources/cfn-resources/cluster/cmd/resource/mappings.go:305,317
		// Note: API returns int, but CFN model expects *int, so convert to pointer
		nodeCount := respSpecs[i].NodeCount
		resultSpecs[i] = ApiSearchDeploymentSpec{
			InstanceSize: &instanceSize,
			NodeCount:    util.IntPtr(nodeCount),
		}
	}

	finalModel := Model{
		Profile:                  prevModel.Profile,
		ClusterName:              prevModel.ClusterName,
		ProjectId:                prevModel.ProjectId,
		Id:                       apiResp.Id,
		Specs:                    resultSpecs,
		StateName:                apiResp.StateName,
		EncryptionAtRestProvider: apiResp.EncryptionAtRestProvider,
	}

	return finalModel
}

func NewSearchDeploymentReq(model *Model) admin20250312010.ApiSearchDeploymentRequest {
	modelSpecs := model.Specs
	requestSpecs := make([]admin20250312010.ApiSearchDeploymentSpec, len(modelSpecs))
	for i, spec := range modelSpecs {
		// Both spec fields are required in CFN model and will be defined
		requestSpecs[i] = admin20250312010.ApiSearchDeploymentSpec{
			InstanceSize: *spec.InstanceSize,
			NodeCount:    *spec.NodeCount,
		}
	}
	return admin20250312010.ApiSearchDeploymentRequest{Specs: requestSpecs}
}
