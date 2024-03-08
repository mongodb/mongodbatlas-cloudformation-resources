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

	"github.com/mongodb/mongodbatlas-cloudformation-resources/search-deployment/cmd/resource"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

type sdkToCFNModelTestCase struct {
	name          string
	SDKResp       admin.ApiSearchDeploymentResponse
	prevModel     resource.Model
	expectedModel resource.Model
}

const (
	profile           = "customProfile"
	dummyDeploymentID = "111111111111111111111111"
	dummyProjectID    = "222222222222222222222222"
	stateName         = "IDLE"
	clusterName       = "Cluster0"
	instanceSize      = "S20_HIGHCPU_NVME"
	nodeCount         = 2
)

func TestSDKToCFNModel(t *testing.T) {
	testCases := []sdkToCFNModelTestCase{
		{
			name: "Complete SDK response",
			prevModel: resource.Model{
				Profile:     admin.PtrString(profile),
				ClusterName: admin.PtrString(clusterName),
				ProjectId:   admin.PtrString(dummyProjectID),
			},
			SDKResp: admin.ApiSearchDeploymentResponse{
				Id:        admin.PtrString(dummyDeploymentID),
				GroupId:   admin.PtrString(dummyProjectID),
				StateName: admin.PtrString(stateName),
				Specs: &[]admin.ApiSearchDeploymentSpec{
					{
						InstanceSize: instanceSize,
						NodeCount:    nodeCount,
					},
				},
			},
			expectedModel: resource.Model{
				Profile:     admin.PtrString(profile),
				ClusterName: admin.PtrString(clusterName),
				ProjectId:   admin.PtrString(dummyProjectID),
				Id:          admin.PtrString(dummyDeploymentID),
				StateName:   admin.PtrString(stateName),
				Specs: []resource.ApiSearchDeploymentSpec{
					{
						InstanceSize: admin.PtrString(instanceSize),
						NodeCount:    admin.PtrInt(nodeCount),
					},
				},
			},
		},
		{
			name: "Empty specs array",
			prevModel: resource.Model{
				Profile:     admin.PtrString(profile),
				ClusterName: admin.PtrString(clusterName),
				ProjectId:   admin.PtrString(dummyProjectID),
			},
			SDKResp: admin.ApiSearchDeploymentResponse{
				Id:        admin.PtrString(dummyDeploymentID),
				GroupId:   admin.PtrString(dummyProjectID),
				StateName: admin.PtrString(stateName),
				Specs:     &[]admin.ApiSearchDeploymentSpec{},
			},
			expectedModel: resource.Model{
				Profile:     admin.PtrString(profile),
				ClusterName: admin.PtrString(clusterName),
				ProjectId:   admin.PtrString(dummyProjectID),
				Id:          admin.PtrString(dummyDeploymentID),
				StateName:   admin.PtrString(stateName),
				Specs:       []resource.ApiSearchDeploymentSpec{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultModel := resource.NewCFNSearchDeployment(&tc.prevModel, &tc.SDKResp)
			assert.Equal(t, tc.expectedModel, resultModel)
		})
	}
}

func TestCFNModelToSDK(t *testing.T) {
	testCases := []struct {
		model          resource.Model
		expectedSDKReq admin.ApiSearchDeploymentRequest
		name           string
	}{
		{
			name: "Complete CFN model",
			model: resource.Model{
				Profile:     admin.PtrString(profile),
				ClusterName: admin.PtrString(clusterName),
				ProjectId:   admin.PtrString(dummyProjectID),
				Id:          admin.PtrString(dummyDeploymentID),
				StateName:   admin.PtrString(stateName),
				Specs: []resource.ApiSearchDeploymentSpec{
					{
						InstanceSize: admin.PtrString(instanceSize),
						NodeCount:    admin.PtrInt(nodeCount),
					},
				},
			},
			expectedSDKReq: admin.ApiSearchDeploymentRequest{
				Specs: &[]admin.ApiSearchDeploymentSpec{
					{
						InstanceSize: instanceSize,
						NodeCount:    nodeCount,
					},
				},
			},
		},
		{
			name: "Empty specs array",
			model: resource.Model{
				Profile:     admin.PtrString(profile),
				ClusterName: admin.PtrString(clusterName),
				ProjectId:   admin.PtrString(dummyProjectID),
				Id:          admin.PtrString(dummyDeploymentID),
				StateName:   admin.PtrString(stateName),
				Specs:       []resource.ApiSearchDeploymentSpec{},
			},
			expectedSDKReq: admin.ApiSearchDeploymentRequest{
				Specs: &[]admin.ApiSearchDeploymentSpec{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultReq := resource.NewSearchDeploymentReq(&tc.model)
			assert.Equal(t, tc.expectedSDKReq, resultReq)
		})
	}
}
