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
	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

type sdkToCFNModelTestCase struct {
	prevModel     resource.Model
	expectedModel resource.Model
	SDKResp       admin20231115014.ApiSearchDeploymentResponse
	name          string
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
				Profile:     admin20231115014.PtrString(profile),
				ClusterName: admin20231115014.PtrString(clusterName),
				ProjectId:   admin20231115014.PtrString(dummyProjectID),
			},
			SDKResp: admin20231115014.ApiSearchDeploymentResponse{
				Id:        admin20231115014.PtrString(dummyDeploymentID),
				GroupId:   admin20231115014.PtrString(dummyProjectID),
				StateName: admin20231115014.PtrString(stateName),
				Specs: &[]admin20231115014.ApiSearchDeploymentSpec{
					{
						InstanceSize: instanceSize,
						NodeCount:    nodeCount,
					},
				},
			},
			expectedModel: resource.Model{
				Profile:     admin20231115014.PtrString(profile),
				ClusterName: admin20231115014.PtrString(clusterName),
				ProjectId:   admin20231115014.PtrString(dummyProjectID),
				Id:          admin20231115014.PtrString(dummyDeploymentID),
				StateName:   admin20231115014.PtrString(stateName),
				Specs: []resource.ApiSearchDeploymentSpec{
					{
						InstanceSize: admin20231115014.PtrString(instanceSize),
						NodeCount:    admin20231115014.PtrInt(nodeCount),
					},
				},
			},
		},
		{
			name: "Empty specs array",
			prevModel: resource.Model{
				Profile:     admin20231115014.PtrString(profile),
				ClusterName: admin20231115014.PtrString(clusterName),
				ProjectId:   admin20231115014.PtrString(dummyProjectID),
			},
			SDKResp: admin20231115014.ApiSearchDeploymentResponse{
				Id:        admin20231115014.PtrString(dummyDeploymentID),
				GroupId:   admin20231115014.PtrString(dummyProjectID),
				StateName: admin20231115014.PtrString(stateName),
				Specs:     &[]admin20231115014.ApiSearchDeploymentSpec{},
			},
			expectedModel: resource.Model{
				Profile:     admin20231115014.PtrString(profile),
				ClusterName: admin20231115014.PtrString(clusterName),
				ProjectId:   admin20231115014.PtrString(dummyProjectID),
				Id:          admin20231115014.PtrString(dummyDeploymentID),
				StateName:   admin20231115014.PtrString(stateName),
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
		name           string
		expectedSDKReq admin20231115014.ApiSearchDeploymentRequest
	}{
		{
			name: "Complete CFN model",
			model: resource.Model{
				Profile:     admin20231115014.PtrString(profile),
				ClusterName: admin20231115014.PtrString(clusterName),
				ProjectId:   admin20231115014.PtrString(dummyProjectID),
				Id:          admin20231115014.PtrString(dummyDeploymentID),
				StateName:   admin20231115014.PtrString(stateName),
				Specs: []resource.ApiSearchDeploymentSpec{
					{
						InstanceSize: admin20231115014.PtrString(instanceSize),
						NodeCount:    admin20231115014.PtrInt(nodeCount),
					},
				},
			},
			expectedSDKReq: admin20231115014.ApiSearchDeploymentRequest{
				Specs: []admin20231115014.ApiSearchDeploymentSpec{
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
				Profile:     admin20231115014.PtrString(profile),
				ClusterName: admin20231115014.PtrString(clusterName),
				ProjectId:   admin20231115014.PtrString(dummyProjectID),
				Id:          admin20231115014.PtrString(dummyDeploymentID),
				StateName:   admin20231115014.PtrString(stateName),
				Specs:       []resource.ApiSearchDeploymentSpec{},
			},
			expectedSDKReq: admin20231115014.ApiSearchDeploymentRequest{
				Specs: []admin20231115014.ApiSearchDeploymentSpec{},
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
