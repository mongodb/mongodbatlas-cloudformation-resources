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
	"fmt"
	"net/http"
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/search-deployment/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"
	"go.mongodb.org/atlas-sdk/v20250312010/mockadmin"
)

const (
	stProfile        = "customProfile"
	stDummyProjectID = "222222222222222222222222"
	stClusterName    = "Cluster0"
)

type stateTransitionTestCase struct {
	respModel           *admin20250312010.ApiSearchDeploymentResponse
	respHTTP            *http.Response
	respError           error
	validateResult      func(t *testing.T, event handler.ProgressEvent)
	name                string
	expectedEventStatus handler.Status
	isDelete            bool
}

func createTestModel(projectID, clusterName, profile string) resource.Model {
	return resource.Model{
		Profile:     admin20250312010.PtrString(profile),
		ClusterName: admin20250312010.PtrString(clusterName),
		ProjectId:   admin20250312010.PtrString(projectID),
	}
}

func TestStateTransitionProgressEvents(t *testing.T) {
	testCases := []stateTransitionTestCase{
		{
			name: "State UPDATING with target IDLE returns in progress",
			respModel: &admin20250312010.ApiSearchDeploymentResponse{
				Id:        admin20250312010.PtrString("test-id-123"),
				StateName: admin20250312010.PtrString("UPDATING"),
				Specs: &[]admin20250312010.ApiSearchDeploymentSpec{
					{InstanceSize: "S20_HIGHCPU_NVME", NodeCount: 2},
				},
			},
			respHTTP:            &http.Response{StatusCode: 200},
			isDelete:            false,
			expectedEventStatus: handler.InProgress,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, constants.Pending, event.Message)
				assert.NotNil(t, event.ResourceModel)
			},
		},
		{
			name: "State IDLE with target IDLE returns success",
			respModel: &admin20250312010.ApiSearchDeploymentResponse{
				Id:        admin20250312010.PtrString("test-id-456"),
				StateName: admin20250312010.PtrString("IDLE"),
				Specs: &[]admin20250312010.ApiSearchDeploymentSpec{
					{InstanceSize: "S20_HIGHCPU_NVME", NodeCount: 2},
				},
			},
			respHTTP:            &http.Response{StatusCode: 200},
			isDelete:            false,
			expectedEventStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, constants.Complete, event.Message)
				model := event.ResourceModel.(*resource.Model)
				assert.Equal(t, "IDLE", *model.StateName)
				assert.Equal(t, stProfile, *model.Profile)
				assert.Equal(t, stClusterName, *model.ClusterName)
			},
		},
		{
			name:                "400 with DoesNotExist and target DELETED returns success",
			respHTTP:            &http.Response{StatusCode: 400},
			respError:           errors.New(resource.SearchDeploymentDoesNotExistsErrorAPI),
			isDelete:            true,
			expectedEventStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, constants.Complete, event.Message)
				assert.Nil(t, event.ResourceModel)
			},
		},
		{
			name: "State IDLE with target DELETED returns in progress",
			respModel: &admin20250312010.ApiSearchDeploymentResponse{
				Id:        admin20250312010.PtrString("test-id-101"),
				StateName: admin20250312010.PtrString("IDLE"),
				Specs:     &[]admin20250312010.ApiSearchDeploymentSpec{{InstanceSize: "S20_HIGHCPU_NVME", NodeCount: 2}},
			},
			respHTTP:            &http.Response{StatusCode: 200},
			isDelete:            true,
			expectedEventStatus: handler.InProgress,
		},
		{
			name:                "500 error returns failed",
			respHTTP:            &http.Response{StatusCode: 500},
			respError:           errors.New("internal server error"),
			isDelete:            false,
			expectedEventStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Contains(t, event.Message, "internal server error")
			},
		},
		{
			name:                "404 error returns failed",
			respHTTP:            &http.Response{StatusCode: 404},
			respError:           errors.New("resource not found"),
			isDelete:            false,
			expectedEventStatus: handler.Failed,
		},
		{
			name:                "400 without specific error code returns failed",
			respHTTP:            &http.Response{StatusCode: 400},
			respError:           errors.New("bad request"),
			isDelete:            true,
			expectedEventStatus: handler.Failed,
		},
		{
			name: "Response with EncryptionAtRestProvider includes it in model",
			respModel: &admin20250312010.ApiSearchDeploymentResponse{
				Id:                       admin20250312010.PtrString("test-id-404"),
				StateName:                admin20250312010.PtrString("IDLE"),
				EncryptionAtRestProvider: admin20250312010.PtrString("AWS"),
				Specs:                    &[]admin20250312010.ApiSearchDeploymentSpec{{InstanceSize: "S20_HIGHCPU_NVME", NodeCount: 2}},
			},
			respHTTP:            &http.Response{StatusCode: 200},
			isDelete:            false,
			expectedEventStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				model := event.ResourceModel.(*resource.Model)
				require.NotNil(t, model.EncryptionAtRestProvider)
				assert.Equal(t, "AWS", *model.EncryptionAtRestProvider)
			},
		},
		{
			name: "Empty specs array handled correctly",
			respModel: &admin20250312010.ApiSearchDeploymentResponse{
				Id:        admin20250312010.PtrString("test-id-505"),
				StateName: admin20250312010.PtrString("IDLE"),
				Specs:     &[]admin20250312010.ApiSearchDeploymentSpec{},
			},
			respHTTP:            &http.Response{StatusCode: 200},
			isDelete:            false,
			expectedEventStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				model := event.ResourceModel.(*resource.Model)
				assert.Empty(t, model.Specs)
			},
		},
		{
			name:                "503 service unavailable returns failed",
			respHTTP:            &http.Response{StatusCode: 503},
			respError:           fmt.Errorf("service unavailable"),
			isDelete:            false,
			expectedEventStatus: handler.Failed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockSearchAPI := mockadmin.NewAtlasSearchApi(t)

			req := admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: mockSearchAPI}
			mockSearchAPI.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
				Return(req).Once()
			mockSearchAPI.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
				Return(tc.respModel, tc.respHTTP, tc.respError).Once()

			client := admin20250312010.APIClient{AtlasSearchApi: mockSearchAPI}
			testModel := createTestModel(stDummyProjectID, stClusterName, stProfile)

			eventResult := resource.ValidateProgress(client, &testModel, tc.isDelete)

			assert.Equal(t, tc.expectedEventStatus, eventResult.OperationStatus)
			if tc.validateResult != nil {
				tc.validateResult(t, eventResult)
			}
		})
	}
}
