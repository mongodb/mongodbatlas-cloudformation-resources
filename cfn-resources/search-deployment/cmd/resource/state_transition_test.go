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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/testutil/mocksvc"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/atlas-sdk/v20231115014/admin"
)

type stateTransitionTestCase struct {
	name                string
	respModel           *admin.ApiSearchDeploymentResponse
	respHTTP            *http.Response
	respError           error
	targetState         string
	expectedEventStatus handler.Status
}

var prevModel = resource.Model{
	Profile:     admin.PtrString(profile),
	ClusterName: admin.PtrString(clusterName),
	ProjectId:   admin.PtrString(dummyProjectID),
}

func TestStateTransitionProgressEvents(t *testing.T) {
	testCases := []stateTransitionTestCase{
		{
			name: "State in WORKING with target IDLE should return in progress event",
			respModel: &admin.ApiSearchDeploymentResponse{
				StateName: admin.PtrString("UPDATING"),
			},
			respHTTP: &http.Response{
				StatusCode: 200,
			},
			respError:           nil,
			targetState:         constants.IdleState,
			expectedEventStatus: handler.InProgress,
		},
		{
			name: "State in IDLE with target IDLE should return success event",
			respModel: &admin.ApiSearchDeploymentResponse{
				StateName: admin.PtrString("IDLE"),
			},
			respHTTP: &http.Response{
				StatusCode: 200,
			},
			respError:           nil,
			targetState:         constants.IdleState,
			expectedEventStatus: handler.Success,
		},
		{
			name:      "400 response with target DELETED should return success event",
			respModel: nil,
			respHTTP: &http.Response{
				StatusCode: 400,
			},
			respError:           errors.New(resource.SearchDeploymentDoesNotExistsError),
			targetState:         constants.DeletedState,
			expectedEventStatus: handler.Success,
		},
		{
			name: "State in WORKING with target DELETED should return in progress event",
			respModel: &admin.ApiSearchDeploymentResponse{
				StateName: admin.PtrString("UPDATING"),
			},
			respHTTP: &http.Response{
				StatusCode: 200,
			},
			respError:           nil,
			targetState:         constants.DeletedState,
			expectedEventStatus: handler.InProgress,
		},
		{
			name: "State in IDLE with target DELETED should return in progress event",
			respModel: &admin.ApiSearchDeploymentResponse{
				StateName: admin.PtrString("IDLE"),
			},
			respHTTP: &http.Response{
				StatusCode: 200,
			},
			respError:           nil,
			targetState:         constants.DeletedState,
			expectedEventStatus: handler.InProgress,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := mocksvc.NewAtlasSearchApi(t)
			m.EXPECT().GetAtlasSearchDeployment(mock.Anything, mock.Anything, mock.Anything).Return(admin.GetAtlasSearchDeploymentApiRequest{ApiService: m}).Once()
			m.EXPECT().GetAtlasSearchDeploymentExecute(mock.Anything).Return(tc.respModel, tc.respHTTP, tc.respError).Once()

			client := admin.APIClient{AtlasSearchApi: m}
			eventResult := resource.HandleStateTransition(client, &prevModel, tc.targetState)
			assert.Equal(t, tc.expectedEventStatus, eventResult.OperationStatus)
		})
	}
}
