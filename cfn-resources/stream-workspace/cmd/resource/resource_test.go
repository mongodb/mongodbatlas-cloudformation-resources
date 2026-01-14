// Copyright 2026 MongoDB Inc
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
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/stream-workspace/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312012/admin"
	"go.mongodb.org/atlas-sdk/v20250312012/mockadmin"
)

func createTestStreamWorkspaceModel() *resource.Model {
	projectID := "507f1f77bcf86cd799439011"
	workspaceName := "test-workspace"
	profile := "default"
	cloudProvider := "AWS"
	region := "VIRGINIA_USA"
	tier := "SP30"
	maxTierSize := "SP50"

	return &resource.Model{
		Profile:       &profile,
		ProjectId:     &projectID,
		WorkspaceName: &workspaceName,
		DataProcessRegion: &resource.StreamsDataProcessRegion{
			CloudProvider: &cloudProvider,
			Region:        &region,
		},
		StreamConfig: &resource.StreamConfig{
			Tier:        &tier,
			MaxTierSize: &maxTierSize,
		},
	}
}

func createTestStreamWorkspaceResponse() *admin.StreamsTenant {
	id := "test-workspace-id-123"
	workspaceName := "test-workspace"
	projectID := "507f1f77bcf86cd799439011"
	cloudProvider := "AWS"
	region := "VIRGINIA_USA"
	tier := "SP30"
	maxTierSize := "SP50"
	hostname1 := "hostname1.example.com"
	hostname2 := "hostname2.example.com"

	return &admin.StreamsTenant{
		Id:        &id,
		Name:      &workspaceName,
		GroupId:   &projectID,
		Hostnames: &[]string{hostname1, hostname2},
		DataProcessRegion: &admin.StreamsDataProcessRegion{
			CloudProvider: cloudProvider,
			Region:        region,
		},
		StreamConfig: &admin.StreamConfig{
			Tier:        &tier,
			MaxTierSize: &maxTierSize,
		},
	}
}

func TestConstants(t *testing.T) {
	assert.Equal(t, []string{"WorkspaceName", constants.ProjectID, constants.DataProcessRegion}, resource.CreateRequiredFields)
	assert.Equal(t, []string{"WorkspaceName", constants.ProjectID}, resource.ReadRequiredFields)
	assert.Equal(t, []string{"WorkspaceName", constants.ProjectID, constants.DataProcessRegion}, resource.UpdateRequiredFields)
	assert.Equal(t, []string{"WorkspaceName", constants.ProjectID}, resource.DeleteRequiredFields)
	assert.Equal(t, []string{constants.ProjectID}, resource.ListRequiredFields)
	assert.Equal(t, "Kafka", resource.Kafka)
	assert.Equal(t, "Cluster", resource.Cluster)
	assert.Equal(t, 100, resource.DefaultItemsPerPage)
}

func TestValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		operation    func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error)
		currentModel *resource.Model
		expectedMsg  string
	}{
		"Create_missingWorkspaceName": {
			operation: resource.Create,
			currentModel: &resource.Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
				DataProcessRegion: &resource.StreamsDataProcessRegion{
					CloudProvider: util.StringPtr("AWS"),
					Region:        util.StringPtr("VIRGINIA_USA"),
				},
			},
			expectedMsg: "required",
		},
		"Create_missingProjectId": {
			operation: resource.Create,
			currentModel: &resource.Model{
				WorkspaceName: util.StringPtr("test-workspace"),
				DataProcessRegion: &resource.StreamsDataProcessRegion{
					CloudProvider: util.StringPtr("AWS"),
					Region:        util.StringPtr("VIRGINIA_USA"),
				},
			},
			expectedMsg: "required",
		},
		"Create_missingDataProcessRegion": {
			operation: resource.Create,
			currentModel: &resource.Model{
				ProjectId:     util.StringPtr("507f1f77bcf86cd799439011"),
				WorkspaceName: util.StringPtr("test-workspace"),
			},
			expectedMsg: "required",
		},
		"Read_missingWorkspaceName": {
			operation:    resource.Read,
			currentModel: &resource.Model{ProjectId: util.StringPtr("507f1f77bcf86cd799439011")},
			expectedMsg:  "required",
		},
		"Read_missingProjectId": {
			operation:    resource.Read,
			currentModel: &resource.Model{WorkspaceName: util.StringPtr("test-workspace")},
			expectedMsg:  "required",
		},
		"Update_missingWorkspaceName": {
			operation: resource.Update,
			currentModel: &resource.Model{
				ProjectId: util.StringPtr("507f1f77bcf86cd799439011"),
				DataProcessRegion: &resource.StreamsDataProcessRegion{
					CloudProvider: util.StringPtr("AWS"),
					Region:        util.StringPtr("VIRGINIA_USA"),
				},
			},
			expectedMsg: "required",
		},
		"Delete_missingWorkspaceName": {
			operation:    resource.Delete,
			currentModel: &resource.Model{ProjectId: util.StringPtr("507f1f77bcf86cd799439011")},
			expectedMsg:  "required",
		},
		"List_missingProjectId": {
			operation:    resource.List,
			currentModel: &resource.Model{},
			expectedMsg:  "required",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := tc.operation(handler.Request{}, nil, tc.currentModel)
			require.NoError(t, err)
			assert.Equal(t, handler.Failed, event.OperationStatus)
			assert.Contains(t, event.Message, tc.expectedMsg)
		})
	}
}

// setupMockClient creates a mock API client and sets up the InitEnvWithLatestClient function.
// It returns a cleanup function that should be deferred.
func setupMockClient(t *testing.T, mockSetup func(*mockadmin.StreamsApi)) func() {
	t.Helper()
	originalInitEnv := resource.InitEnvWithLatestClient
	mockStreamsAPI := mockadmin.NewStreamsApi(t)
	mockSetup(mockStreamsAPI)

	mockClient := &admin.APIClient{StreamsApi: mockStreamsAPI}
	resource.InitEnvWithLatestClient = func(req handler.Request, currentModel *resource.Model, requiredFields []string) (*admin.APIClient, *handler.ProgressEvent) {
		return mockClient, nil
	}

	return func() { resource.InitEnvWithLatestClient = originalInitEnv }
}

func TestCRUDOperations(t *testing.T) {
	testCases := map[string]struct {
		operation      func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error)
		mockSetup      func(*mockadmin.StreamsApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
	}{
		"Read_notFound": {
			operation: resource.Read,
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().GetStreamWorkspace(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetStreamWorkspaceApiRequest{ApiService: m})
				m.EXPECT().GetStreamWorkspaceExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
			},
		},
		"Read_apiError": {
			operation: resource.Read,
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().GetStreamWorkspace(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetStreamWorkspaceApiRequest{ApiService: m})
				m.EXPECT().GetStreamWorkspaceExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
		"Update_notFound": {
			operation: resource.Update,
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().UpdateStreamWorkspace(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin.UpdateStreamWorkspaceApiRequest{ApiService: m})
				m.EXPECT().UpdateStreamWorkspaceExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
				assert.Equal(t, "StreamWorkspace not found", event.Message)
			},
		},
		"Update_apiError": {
			operation: resource.Update,
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().UpdateStreamWorkspace(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin.UpdateStreamWorkspaceApiRequest{ApiService: m})
				m.EXPECT().UpdateStreamWorkspaceExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
		"Create_apiError": {
			operation: resource.Create,
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().CreateStreamWorkspace(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.CreateStreamWorkspaceApiRequest{ApiService: m})
				m.EXPECT().CreateStreamWorkspaceExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("API error"))
			},
			expectedStatus: handler.Failed,
		},
		"Delete_apiError": {
			operation: resource.Delete,
			mockSetup: func(m *mockadmin.StreamsApi) {
				m.EXPECT().DeleteStreamWorkspace(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.DeleteStreamWorkspaceApiRequest{ApiService: m})
				m.EXPECT().DeleteStreamWorkspaceExecute(mock.Anything).
					Return(&http.Response{StatusCode: 500}, fmt.Errorf("delete failed"))
			},
			expectedStatus: handler.Failed,
		},
		"List_withPagination": {
			operation: resource.List,
			mockSetup: func(m *mockadmin.StreamsApi) {
				workspace1 := createTestStreamWorkspaceResponse()
				workspace2 := createTestStreamWorkspaceResponse()
				workspace2Name := "test-workspace-2"
				workspace2.Id = util.StringPtr("workspace-id-2")
				workspace2.Name = &workspace2Name

				page1Results := []admin.StreamsTenant{*workspace1}
				page2Results := []admin.StreamsTenant{*workspace2}
				totalCount := 2

				req1 := admin.ListStreamWorkspacesApiRequest{ApiService: m}
				m.EXPECT().ListStreamWorkspacesWithParams(mock.Anything, mock.MatchedBy(func(params *admin.ListStreamWorkspacesApiParams) bool {
					return params.PageNum != nil && *params.PageNum == 1
				})).Return(req1)
				m.EXPECT().ListStreamWorkspacesExecute(mock.MatchedBy(func(r admin.ListStreamWorkspacesApiRequest) bool { return true })).
					Return(&admin.PaginatedApiStreamsTenant{
						Results:    &page1Results,
						TotalCount: &totalCount,
					}, &http.Response{StatusCode: 200}, nil).Once()

				req2 := admin.ListStreamWorkspacesApiRequest{ApiService: m}
				m.EXPECT().ListStreamWorkspacesWithParams(mock.Anything, mock.MatchedBy(func(params *admin.ListStreamWorkspacesApiParams) bool {
					return params.PageNum != nil && *params.PageNum == 2
				})).Return(req2)
				m.EXPECT().ListStreamWorkspacesExecute(mock.MatchedBy(func(r admin.ListStreamWorkspacesApiRequest) bool { return true })).
					Return(&admin.PaginatedApiStreamsTenant{
						Results:    &page2Results,
						TotalCount: &totalCount,
					}, &http.Response{StatusCode: 200}, nil).Once()
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				require.NotNil(t, event.ResourceModels)
				assert.GreaterOrEqual(t, len(event.ResourceModels), 2)
			},
		},
		"List_apiError": {
			operation: resource.List,
			mockSetup: func(m *mockadmin.StreamsApi) {
				req := admin.ListStreamWorkspacesApiRequest{ApiService: m}
				m.EXPECT().ListStreamWorkspacesWithParams(mock.Anything, mock.Anything).Return(req)
				m.EXPECT().ListStreamWorkspacesExecute(mock.MatchedBy(func(r admin.ListStreamWorkspacesApiRequest) bool { return true })).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("list failed"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			cleanup := setupMockClient(t, tc.mockSetup)
			defer cleanup()

			event, err := tc.operation(handler.Request{}, nil, createTestStreamWorkspaceModel())
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
			if tc.validateResult != nil {
				tc.validateResult(t, event)
			}
		})
	}
}

func TestHandleError(t *testing.T) {
	testCases := map[string]struct {
		response       *http.Response
		err            error
		expectedStatus handler.Status
	}{
		"notFoundError": {
			response:       &http.Response{StatusCode: http.StatusNotFound},
			err:            fmt.Errorf("resource not found"),
			expectedStatus: handler.Failed,
		},
		"internalServerError": {
			response:       &http.Response{StatusCode: http.StatusInternalServerError},
			err:            fmt.Errorf("internal server error"),
			expectedStatus: handler.Failed,
		},
		"badRequestError": {
			response:       &http.Response{StatusCode: http.StatusBadRequest},
			err:            fmt.Errorf("bad request"),
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			event, err := resource.HandleError(tc.response, constants.CREATE, tc.err)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}
