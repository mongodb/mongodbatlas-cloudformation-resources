package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

type sdkToCFNModelTestCase struct {
	name          string
	SDKResp       admin.ApiSearchDeploymentResponse
	prevModel     Model
	expectedModel Model
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
			prevModel: Model{
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
			expectedModel: Model{
				Profile:     admin.PtrString(profile),
				ClusterName: admin.PtrString(clusterName),
				ProjectId:   admin.PtrString(dummyProjectID),
				Id:          admin.PtrString(dummyDeploymentID),
				StateName:   admin.PtrString(stateName),
				Specs: []ApiSearchDeploymentSpec{
					{
						InstanceSize: admin.PtrString(instanceSize),
						NodeCount:    admin.PtrInt(nodeCount),
					},
				},
			},
		},
		{
			name: "Empty specs array",
			prevModel: Model{
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
			expectedModel: Model{
				Profile:     admin.PtrString(profile),
				ClusterName: admin.PtrString(clusterName),
				ProjectId:   admin.PtrString(dummyProjectID),
				Id:          admin.PtrString(dummyDeploymentID),
				StateName:   admin.PtrString(stateName),
				Specs:       []ApiSearchDeploymentSpec{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultModel := NewCFNSearchDeployment(&tc.prevModel, &tc.SDKResp)
			assert.Equal(t, tc.expectedModel, resultModel)
		})
	}
}

func TestCFNModelToSDK(t *testing.T) {
	testCases := []struct {
		model          Model
		expectedSDKReq admin.ApiSearchDeploymentRequest
		name           string
	}{
		{
			name: "Complete CFN model",
			model: Model{
				Profile:     admin.PtrString(profile),
				ClusterName: admin.PtrString(clusterName),
				ProjectId:   admin.PtrString(dummyProjectID),
				Id:          admin.PtrString(dummyDeploymentID),
				StateName:   admin.PtrString(stateName),
				Specs: []ApiSearchDeploymentSpec{
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
			model: Model{
				Profile:     admin.PtrString(profile),
				ClusterName: admin.PtrString(clusterName),
				ProjectId:   admin.PtrString(dummyProjectID),
				Id:          admin.PtrString(dummyDeploymentID),
				StateName:   admin.PtrString(stateName),
				Specs:       []ApiSearchDeploymentSpec{},
			},
			expectedSDKReq: admin.ApiSearchDeploymentRequest{
				Specs: &[]admin.ApiSearchDeploymentSpec{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultReq := NewSearchDeploymentReq(&tc.model)
			assert.Equal(t, tc.expectedSDKReq, resultReq)
		})
	}
}
