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
	"fmt"
	"net/http"
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/search-deployment/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"
	"go.mongodb.org/atlas-sdk/v20250312010/mockadmin"
)

func createTestSearchDeploymentModel() *resource.Model {
	projectID := "507f1f77bcf86cd799439011"
	clusterName := "test-cluster"
	profile := "default"
	instanceSize := "S20_HIGHCPU_NVME"
	nodeCount := 2

	return &resource.Model{
		Profile:     &profile,
		ProjectId:   &projectID,
		ClusterName: &clusterName,
		Specs: []resource.ApiSearchDeploymentSpec{
			{
				InstanceSize: &instanceSize,
				NodeCount:    &nodeCount,
			},
		},
	}
}

func createTestSearchDeploymentResponse() *admin20250312010.ApiSearchDeploymentResponse {
	id := "test-id-123"
	stateName := "IDLE"
	return &admin20250312010.ApiSearchDeploymentResponse{
		Id:        &id,
		StateName: &stateName,
		Specs: &[]admin20250312010.ApiSearchDeploymentSpec{
			{
				InstanceSize: "S20_HIGHCPU_NVME",
				NodeCount:    2,
			},
		},
	}
}

func TestIsCallback(t *testing.T) {
	testCases := map[string]struct {
		req      handler.Request
		expected bool
	}{
		"withCallback": {
			req:      handler.Request{CallbackContext: map[string]interface{}{"callbackSearchDeployment": true}},
			expected: true,
		},
		"withoutCallback": {
			req:      handler.Request{CallbackContext: map[string]interface{}{}},
			expected: false,
		},
		"nilCallbackContext": {
			req:      handler.Request{},
			expected: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := resource.IsCallback(&tc.req)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestConstants(t *testing.T) {
	assert.Equal(t, 40, resource.CallBackSeconds)
	assert.Equal(t, "ATLAS_SEARCH_DEPLOYMENT_DOES_NOT_EXIST", resource.SearchDeploymentDoesNotExistsErrorAPI)
	assert.Equal(t, "ATLAS_SEARCH_DEPLOYMENT_ALREADY_EXISTS", resource.SearchDeploymentAlreadyExistsErrorAPI)
}

func TestRequiredFields(t *testing.T) {
	assert.Equal(t, []string{constants.ProjectID, constants.ClusterName, constants.Specs}, resource.CreateRequiredFields)
	assert.Equal(t, []string{constants.ProjectID, constants.ClusterName}, resource.ReadRequiredFields)
	assert.Equal(t, []string{constants.ProjectID, constants.ClusterName, constants.Specs}, resource.UpdateRequiredFields)
	assert.Equal(t, []string{constants.ProjectID, constants.ClusterName}, resource.DeleteRequiredFields)
}

func TestList(t *testing.T) {
	req := handler.Request{}
	event, err := resource.List(req, nil, nil)

	require.Error(t, err)
	assert.Equal(t, "not implemented: List", err.Error())
	assert.Equal(t, handler.ProgressEvent{}, event)
}

func TestGetModelWithID(t *testing.T) {
	id := "test-id-123"
	projectID := "507f1f77bcf86cd799439011"
	clusterName := "test-cluster"

	testCases := map[string]struct {
		currentModel *resource.Model
		prevModel    *resource.Model
		checkResp    *admin20250312010.ApiSearchDeploymentResponse
		expectedID   *string
	}{
		"currentModelHasID": {
			currentModel: &resource.Model{Id: &id, ProjectId: &projectID, ClusterName: &clusterName},
			prevModel:    nil,
			checkResp:    nil,
			expectedID:   &id,
		},
		"prevModelHasID": {
			currentModel: &resource.Model{ProjectId: &projectID, ClusterName: &clusterName},
			prevModel:    &resource.Model{Id: &id, ProjectId: &projectID, ClusterName: &clusterName},
			checkResp:    nil,
			expectedID:   &id,
		},
		"checkRespHasID": {
			currentModel: &resource.Model{ProjectId: &projectID, ClusterName: &clusterName},
			prevModel:    nil,
			checkResp:    &admin20250312010.ApiSearchDeploymentResponse{Id: &id},
			expectedID:   &id,
		},
		"noIDAvailable": {
			currentModel: &resource.Model{ProjectId: &projectID, ClusterName: &clusterName},
			prevModel:    nil,
			checkResp:    nil,
			expectedID:   nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := resource.GetModelWithID(tc.currentModel, tc.prevModel, tc.checkResp)
			if tc.expectedID == nil {
				assert.Nil(t, result)
			} else {
				require.NotNil(t, result)
				assert.Equal(t, *tc.expectedID, *result.Id)
			}
		})
	}
}

func TestHandleError(t *testing.T) {
	createSDKError := func(errorCode string, statusCode int) *admin20250312010.GenericOpenAPIError {
		apiErr := admin20250312010.ApiError{
			Error:     statusCode,
			ErrorCode: errorCode,
		}
		sdkErr := &admin20250312010.GenericOpenAPIError{}
		sdkErr.SetModel(apiErr)
		return sdkErr
	}

	testCases := map[string]struct {
		response          *http.Response
		err               error
		expectedStatus    handler.Status
		expectedErrorCode string
	}{
		"AlreadyExistsError": {
			response:          &http.Response{StatusCode: http.StatusBadRequest},
			err:               createSDKError(resource.SearchDeploymentAlreadyExistsErrorAPI, http.StatusBadRequest),
			expectedStatus:    handler.Failed,
			expectedErrorCode: "AlreadyExists",
		},
		"DoesNotExistError": {
			response:          &http.Response{StatusCode: http.StatusBadRequest},
			err:               createSDKError(resource.SearchDeploymentDoesNotExistsErrorAPI, http.StatusBadRequest),
			expectedStatus:    handler.Failed,
			expectedErrorCode: "NotFound",
		},
		"NotFoundError": {
			response:          &http.Response{StatusCode: http.StatusNotFound},
			err:               fmt.Errorf("resource not found"),
			expectedStatus:    handler.Failed,
			expectedErrorCode: "NotFound",
		},
		"ErrorContainsNotExist": {
			response:          &http.Response{StatusCode: http.StatusInternalServerError},
			err:               fmt.Errorf("resource does not exist"),
			expectedStatus:    handler.Failed,
			expectedErrorCode: "NotFound",
		},
		"ErrorContainsBeingDeleted": {
			response:          &http.Response{StatusCode: http.StatusInternalServerError},
			err:               fmt.Errorf("resource is being deleted"),
			expectedStatus:    handler.Failed,
			expectedErrorCode: "NotFound",
		},
		"GenericError": {
			response:       &http.Response{StatusCode: http.StatusInternalServerError},
			err:            fmt.Errorf("internal server error"),
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.HandleError(tc.response, tc.err)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.expectedErrorCode != "" {
				assert.Equal(t, tc.expectedErrorCode, event.HandlerErrorCode)
			}
		})
	}
}

func TestCreateValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel *resource.Model
		expectedMsg  string
	}{
		"missingProjectId":   {&resource.Model{ClusterName: util.StringPtr("test-cluster"), Specs: []resource.ApiSearchDeploymentSpec{{InstanceSize: util.StringPtr("S20_HIGHCPU_NVME"), NodeCount: util.IntPtr(2)}}}, "required"},
		"missingClusterName": {&resource.Model{ProjectId: util.StringPtr("507f1f77bcf86cd799439011"), Specs: []resource.ApiSearchDeploymentSpec{{InstanceSize: util.StringPtr("S20_HIGHCPU_NVME"), NodeCount: util.IntPtr(2)}}}, "required"},
		"missingSpecs":       {&resource.Model{ProjectId: util.StringPtr("507f1f77bcf86cd799439011"), ClusterName: util.StringPtr("test-cluster")}, "required"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.Create(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestReadValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel *resource.Model
		expectedMsg  string
	}{
		"missingProjectId":   {&resource.Model{ClusterName: util.StringPtr("test-cluster")}, "required"},
		"missingClusterName": {&resource.Model{ProjectId: util.StringPtr("507f1f77bcf86cd799439011")}, "required"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.Read(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestUpdateValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel *resource.Model
		expectedMsg  string
	}{
		"missingProjectId":   {&resource.Model{ClusterName: util.StringPtr("test-cluster"), Specs: []resource.ApiSearchDeploymentSpec{{InstanceSize: util.StringPtr("S20_HIGHCPU_NVME"), NodeCount: util.IntPtr(2)}}}, "required"},
		"missingClusterName": {&resource.Model{ProjectId: util.StringPtr("507f1f77bcf86cd799439011"), Specs: []resource.ApiSearchDeploymentSpec{{InstanceSize: util.StringPtr("S20_HIGHCPU_NVME"), NodeCount: util.IntPtr(2)}}}, "required"},
		"missingSpecs":       {&resource.Model{ProjectId: util.StringPtr("507f1f77bcf86cd799439011"), ClusterName: util.StringPtr("test-cluster")}, "required"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.Update(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestDeleteValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		currentModel *resource.Model
		expectedMsg  string
	}{
		"missingProjectId":   {&resource.Model{ClusterName: util.StringPtr("test-cluster")}, "required"},
		"missingClusterName": {&resource.Model{ProjectId: util.StringPtr("507f1f77bcf86cd799439011")}, "required"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.Delete(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

func TestCreateWithMocks(t *testing.T) {
	originalInitEnv := resource.InitEnvWithClient
	defer func() { resource.InitEnvWithClient = originalInitEnv }()

	testCases := map[string]struct {
		mockSetup      func(*mockadmin.AtlasSearchApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
		req            handler.Request
	}{
		"successfulCreate": {
			req: handler.Request{},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				idleResp := createTestSearchDeploymentResponse()
				stateName := "IDLE"
				idleResp.StateName = &stateName
				m.EXPECT().CreateClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().CreateClusterSearchDeploymentExecute(mock.Anything).
					Return(idleResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, constants.Complete, event.Message)
				assert.NotNil(t, event.ResourceModel)
			},
		},
		"createWithCallback": {
			req: handler.Request{CallbackContext: map[string]interface{}{"callbackSearchDeployment": true}},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(createTestSearchDeploymentResponse(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"createWithError": {
			req: handler.Request{},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().CreateClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().CreateClusterSearchDeploymentExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("API error"))
			},
			expectedStatus: handler.Failed,
		},
		"createAlreadyExists": {
			req: handler.Request{},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				createErr := fmt.Errorf("error creating: %s", resource.SearchDeploymentAlreadyExistsErrorAPI)

				m.EXPECT().CreateClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().CreateClusterSearchDeploymentExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: http.StatusConflict}, createErr)

				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(createTestSearchDeploymentResponse(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, constants.Complete, event.Message)
			},
		},
		"createNotFoundThenSuccess": {
			req: handler.Request{},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().CreateClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().CreateClusterSearchDeploymentExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: http.StatusNotFound}, fmt.Errorf("not found"))

				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(createTestSearchDeploymentResponse(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"createNoIDInResponse": {
			req: handler.Request{},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				respWithNoID := &admin20250312010.ApiSearchDeploymentResponse{
					StateName: admin20250312010.PtrString("UPDATING"),
				}
				m.EXPECT().CreateClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().CreateClusterSearchDeploymentExecute(mock.Anything).
					Return(respWithNoID, &http.Response{StatusCode: 200}, nil)

				respWithNoID2 := &admin20250312010.ApiSearchDeploymentResponse{
					StateName: admin20250312010.PtrString("UPDATING"),
				}
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(respWithNoID2, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"createInProgressState": {
			req: handler.Request{},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				updatingResp := createTestSearchDeploymentResponse()
				stateName := "UPDATING"
				updatingResp.StateName = &stateName

				m.EXPECT().CreateClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.CreateClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().CreateClusterSearchDeploymentExecute(mock.Anything).
					Return(updatingResp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockSearchAPI := mockadmin.NewAtlasSearchApi(t)
			tc.mockSetup(mockSearchAPI)

			mockClient := &admin20250312010.APIClient{AtlasSearchApi: mockSearchAPI}
			resource.InitEnvWithClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := resource.Create(tc.req, nil, createTestSearchDeploymentModel())
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.validateResult != nil {
				tc.validateResult(t, event)
			}
		})
	}
}

func TestReadWithMocks(t *testing.T) {
	originalInitEnv := resource.InitEnvWithClient
	defer func() { resource.InitEnvWithClient = originalInitEnv }()

	testCases := map[string]struct {
		mockSetup      func(*mockadmin.AtlasSearchApi)
		expectedStatus handler.Status
	}{
		"successfulRead": {
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(createTestSearchDeploymentResponse(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"readNotFound": {
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockSearchAPI := mockadmin.NewAtlasSearchApi(t)
			tc.mockSetup(mockSearchAPI)

			mockClient := &admin20250312010.APIClient{AtlasSearchApi: mockSearchAPI}
			resource.InitEnvWithClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := resource.Read(handler.Request{}, nil, createTestSearchDeploymentModel())
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

func TestUpdateWithMocks(t *testing.T) {
	originalInitEnv := resource.InitEnvWithClient
	defer func() { resource.InitEnvWithClient = originalInitEnv }()

	testCases := map[string]struct {
		mockSetup      func(*mockadmin.AtlasSearchApi)
		expectedStatus handler.Status
		req            handler.Request
	}{
		"successfulUpdate": {
			req: handler.Request{},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(createTestSearchDeploymentResponse(), &http.Response{StatusCode: 200}, nil)

				m.EXPECT().UpdateClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.UpdateClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().UpdateClusterSearchDeploymentExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 200}, nil)

				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(createTestSearchDeploymentResponse(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"updateWithCallback": {
			req: handler.Request{CallbackContext: map[string]interface{}{"callbackSearchDeployment": true}},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(createTestSearchDeploymentResponse(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
		},
		"updateWithError": {
			req: handler.Request{},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(createTestSearchDeploymentResponse(), &http.Response{StatusCode: 200}, nil)
				m.EXPECT().UpdateClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.UpdateClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().UpdateClusterSearchDeploymentExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("update failed"))
			},
			expectedStatus: handler.Failed,
		},
		"updateResourceNotFound": {
			req: handler.Request{},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: http.StatusNotFound}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
		"updateResourceNilResponse": {
			req: handler.Request{},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				respWithNoID := &admin20250312010.ApiSearchDeploymentResponse{}
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(respWithNoID, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Failed,
		},
		"updateGetErrorAfterUpdate": {
			req: handler.Request{},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				checkResp := createTestSearchDeploymentResponse()
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(checkResp, &http.Response{StatusCode: 200}, nil).Once()

				m.EXPECT().UpdateClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.UpdateClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().UpdateClusterSearchDeploymentExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 200}, nil)

				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("get failed"))
			},
			expectedStatus: handler.InProgress,
		},
		"updateGetNilResponseAfterUpdate": {
			req: handler.Request{},
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				checkResp := createTestSearchDeploymentResponse()
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(checkResp, &http.Response{StatusCode: 200}, nil).Once()

				m.EXPECT().UpdateClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.UpdateClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().UpdateClusterSearchDeploymentExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 200}, nil)

				respWithNoID := &admin20250312010.ApiSearchDeploymentResponse{}
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(respWithNoID, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockSearchAPI := mockadmin.NewAtlasSearchApi(t)
			tc.mockSetup(mockSearchAPI)

			mockClient := &admin20250312010.APIClient{AtlasSearchApi: mockSearchAPI}
			resource.InitEnvWithClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
				return mockClient, nil
			}

			event, err := resource.Update(tc.req, nil, createTestSearchDeploymentModel())
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

func TestDeleteWithMocks(t *testing.T) {
	originalInitEnv := resource.InitEnvWithClient
	defer func() { resource.InitEnvWithClient = originalInitEnv }()

	testCases := map[string]struct {
		mockSetup      func(*mockadmin.AtlasSearchApi)
		currentModel   *resource.Model
		prevModel      *resource.Model
		expectedStatus handler.Status
		req            handler.Request
	}{
		"successfulDelete": {
			req:          handler.Request{},
			currentModel: createTestSearchDeploymentModel(),
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().DeleteClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.DeleteClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().DeleteClusterSearchDeploymentExecute(mock.Anything).
					Return(&http.Response{StatusCode: 200}, nil)
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(createTestSearchDeploymentResponse(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"deleteWithCallback": {
			req:          handler.Request{CallbackContext: map[string]interface{}{"callbackSearchDeployment": true}},
			currentModel: createTestSearchDeploymentModel(),
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Success,
		},
		"deleteWithError": {
			req:          handler.Request{},
			currentModel: createTestSearchDeploymentModel(),
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().DeleteClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.DeleteClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().DeleteClusterSearchDeploymentExecute(mock.Anything).
					Return(&http.Response{StatusCode: 500}, fmt.Errorf("delete failed"))
			},
			expectedStatus: handler.Failed,
		},
		"deleteResourceNotFoundInGet": {
			req:          handler.Request{},
			currentModel: createTestSearchDeploymentModel(),
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().DeleteClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.DeleteClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().DeleteClusterSearchDeploymentExecute(mock.Anything).
					Return(&http.Response{StatusCode: 200}, nil)
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.InProgress,
		},
		"deleteWithNilCurrentModelUsesPrevModel": {
			req:          handler.Request{},
			currentModel: nil,
			prevModel:    createTestSearchDeploymentModel(),
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
				m.EXPECT().DeleteClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.DeleteClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().DeleteClusterSearchDeploymentExecute(mock.Anything).
					Return(&http.Response{StatusCode: 200}, nil)
				m.EXPECT().GetClusterSearchDeployment(mock.Anything, mock.Anything, mock.Anything).
					Return(admin20250312010.GetClusterSearchDeploymentApiRequest{ApiService: m})
				m.EXPECT().GetClusterSearchDeploymentExecute(mock.Anything).
					Return(createTestSearchDeploymentResponse(), &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.InProgress,
		},
		"deleteWithNilCurrentModelNoProjectID": {
			req: handler.Request{},
			currentModel: &resource.Model{
				ClusterName: util.StringPtr("test-cluster"),
			},
			prevModel: nil,
			mockSetup: func(m *mockadmin.AtlasSearchApi) {
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockSearchAPI := mockadmin.NewAtlasSearchApi(t)
			tc.mockSetup(mockSearchAPI)

			mockClient := &admin20250312010.APIClient{AtlasSearchApi: mockSearchAPI}

			if tc.expectedStatus != handler.Failed || name != "deleteWithNilCurrentModelNoProjectID" {
				resource.InitEnvWithClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin20250312010.APIClient, *handler.ProgressEvent) {
					return mockClient, nil
				}
			} else {
				resource.InitEnvWithClient = originalInitEnv
			}

			currentModel := tc.currentModel
			if currentModel == nil {
				currentModel = createTestSearchDeploymentModel()
			}
			event, err := resource.Delete(tc.req, tc.prevModel, currentModel)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}
