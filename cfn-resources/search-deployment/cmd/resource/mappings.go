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

import "go.mongodb.org/atlas-sdk/v20240805001/admin"

func NewCFNSearchDeployment(prevModel *Model, apiResp *admin.ApiSearchDeploymentResponse) Model {
	respSpecs := apiResp.GetSpecs()
	resultSpecs := make([]ApiSearchDeploymentSpec, len(respSpecs))
	for i := range respSpecs {
		resultSpecs[i] = ApiSearchDeploymentSpec{
			InstanceSize: &respSpecs[i].InstanceSize,
			NodeCount:    &respSpecs[i].NodeCount,
		}
	}
	return Model{
		Profile:     prevModel.Profile,
		ClusterName: prevModel.ClusterName,
		ProjectId:   prevModel.ProjectId,
		Id:          apiResp.Id,
		Specs:       resultSpecs,
		StateName:   apiResp.StateName,
	}
}

func NewSearchDeploymentReq(model *Model) admin.ApiSearchDeploymentRequest {
	modelSpecs := model.Specs
	requestSpecs := make([]admin.ApiSearchDeploymentSpec, len(modelSpecs))
	for i, spec := range modelSpecs {
		// Both spec fields are required in CFN model and will be defined
		requestSpecs[i] = admin.ApiSearchDeploymentSpec{
			InstanceSize: *spec.InstanceSize,
			NodeCount:    *spec.NodeCount,
		}
	}
	return admin.ApiSearchDeploymentRequest{Specs: requestSpecs}
}
