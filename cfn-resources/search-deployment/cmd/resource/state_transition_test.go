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
	"errors"
	"net/http"
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/search-deployment/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/atlas-sdk/v20250312014/admin"
	"go.mongodb.org/atlas-sdk/v20250312014/mockadmin"
)

type stateTransitionTestCase struct {
	name                string
	respModel           *admin.ApiSearchDeploymentResponse
	respHTTP            *http.Response
	respError           error
	targetState         string
	expectedEventStatus handler.Status
}

var prevModelStateTransition = resource.Model{
	Profile:     admin.PtrString(profile),
	ClusterName: admin.PtrString(clusterName),
	ProjectId:   admin.PtrString(dummyProjectID),
}

func TestStateTransitionProgressEvents(t *testing.T) {
	testCases := []stateTransitionTestCase{
		{
			name: "State in WORKING with target IDLE should return in progress event",
			respModel: &admin.ApiSearchDeploymentResponse{
				Id:        admin.PtrString(dummyDeploymentID),
				StateName: admin.PtrString("UPDATING"),
				Specs:     &[]admin.ApiSearchDeploymentSpec{{InstanceSize: instanceSize, NodeCount: nodeCount}},
			},
			respHTTP:            &http.Response{StatusCode: 200},
			targetState:         constants.IdleState,
			expectedEventStatus: handler.InProgress,
		},
		{
			name: "State in IDLE with target IDLE should return success event",
			respModel: &admin.ApiSearchDeploymentResponse{
				Id:        admin.PtrString(dummyDeploymentID),
				StateName: admin.PtrString("IDLE"),
				Specs:     &[]admin.ApiSearchDeploymentSpec{{InstanceSize: instanceSize, NodeCount: nodeCount}},
			},
			respHTTP:            &http.Response{StatusCode: 200},
			targetState:         constants.IdleState,
			expectedEventStatus: handler.Success,
		},
		{
			name:                "400 response with target DELETED should return success event",
			respHTTP:            &http.Response{StatusCode: 400},
			respError:           errors.New(resource.SearchDeploymentDoesNotExistsError),
			targetState:         constants.DeletedState,
			expectedEventStatus: handler.Success,
		},
		{
			name: "State in WORKING with target DELETED should return in progress event",
			respModel: &admin.ApiSearchDeploymentResponse{
				Id:        admin.PtrString(dummyDeploymentID),
				StateName: admin.PtrString("UPDATING"),
				Specs:     &[]admin.ApiSearchDeploymentSpec{{InstanceSize: instanceSize, NodeCount: nodeCount}},
			},
			respHTTP:            &http.Response{StatusCode: 200},
			targetState:         constants.DeletedState,
			expectedEventStatus: handler.InProgress,
		},
		{
			name: "State in IDLE with target DELETED should return in progress event",
			respModel: &admin.ApiSearchDeploymentResponse{
				Id:        admin.PtrString(dummyDeploymentID),
				StateName: admin.PtrString("IDLE"),
				Specs:     &[]admin.ApiSearchDeploymentSpec{{InstanceSize: instanceSize, NodeCount: nodeCount}},
			},
			respHTTP:            &http.Response{StatusCode: 200},
			targetState:         constants.DeletedState,
			expectedEventStatus: handler.InProgress,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockSearchAPI := mockadmin.NewAtlasSearchApi(t)
			req := admin.GetClusterSearchDeploymentApiRequest{ApiService: mockSearchAPI}
			mockSearchAPI.EXPECT().GetClusterSearchDeployment(mock.Anything, dummyProjectID, clusterName).Return(req).Once()
			mockSearchAPI.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).Return(tc.respModel, tc.respHTTP, tc.respError).Once()

			client := admin.APIClient{AtlasSearchApi: mockSearchAPI}
			eventResult := resource.HandleStateTransition(client, &prevModelStateTransition, tc.targetState)

			assert.Equal(t, tc.expectedEventStatus, eventResult.OperationStatus)
		})
	}
}
