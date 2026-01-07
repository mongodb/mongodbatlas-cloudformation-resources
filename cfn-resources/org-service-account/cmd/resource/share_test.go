// Copyright 2025 MongoDB Inc
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
	"time"

	"go.mongodb.org/atlas-sdk/v20250312010/admin"
	"go.mongodb.org/atlas-sdk/v20250312010/mockadmin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/ptr"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/org-service-account/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func createTestOrgServiceAccountModel() *resource.Model {
	orgID := "63350255419cf25e3d511c95"
	name := "test-service-account"
	description := "Test description"
	roles := []string{"ORG_MEMBER"}
	secretExpiresAfterHours := 720
	profile := "default"

	return &resource.Model{
		Profile:                 &profile,
		OrgId:                   &orgID,
		Name:                    &name,
		Description:             &description,
		Roles:                   roles,
		SecretExpiresAfterHours: &secretExpiresAfterHours,
	}
}

func createTestOrgServiceAccountResponse() *admin.OrgServiceAccount {
	now := time.Now()
	clientID := "mdb_sa_id_123456789"
	name := "test-service-account"
	description := "Test description"
	roles := []string{"ORG_MEMBER"}
	secretID := "secret-id-123"

	return &admin.OrgServiceAccount{
		ClientId:    &clientID,
		Name:        &name,
		Description: &description,
		Roles:       &roles,
		CreatedAt:   &now,
		Secrets: &[]admin.ServiceAccountSecret{
			{
				Id:                secretID,
				CreatedAt:         now,
				ExpiresAt:         now.Add(720 * time.Hour),
				MaskedSecretValue: ptr.String("****"),
				Secret:            ptr.String("secret-value"),
			},
		},
	}
}

func TestConstants(t *testing.T) {
	assert.Equal(t, []string{"OrgId", "Name", "Description", "Roles"}, resource.CreateRequiredFields)
	assert.Equal(t, []string{"OrgId", "ClientId"}, resource.ReadRequiredFields)
	assert.Equal(t, []string{"OrgId", "ClientId"}, resource.UpdateRequiredFields)
	assert.Equal(t, []string{"OrgId", "ClientId"}, resource.DeleteRequiredFields)
	assert.Equal(t, []string{"OrgId"}, resource.ListRequiredFields)
}

func TestValidationErrors(t *testing.T) {
	testCases := map[string]struct {
		operation    func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error)
		currentModel *resource.Model
		expectedMsg  string
	}{
		"Create_missingOrgId": {
			operation: resource.Create,
			currentModel: &resource.Model{
				Name:        util.StringPtr("test-service-account"),
				Description: util.StringPtr("Test description"),
				Roles:       []string{"ORG_MEMBER"},
			},
			expectedMsg: "required",
		},
		"Create_missingName": {
			operation: resource.Create,
			currentModel: &resource.Model{
				OrgId:       util.StringPtr("63350255419cf25e3d511c95"),
				Description: util.StringPtr("Test description"),
				Roles:       []string{"ORG_MEMBER"},
			},
			expectedMsg: "required",
		},
		"Create_missingDescription": {
			operation: resource.Create,
			currentModel: &resource.Model{
				OrgId: util.StringPtr("63350255419cf25e3d511c95"),
				Name:  util.StringPtr("test-service-account"),
				Roles: []string{"ORG_MEMBER"},
			},
			expectedMsg: "required",
		},
		"Create_missingRoles": {
			operation: resource.Create,
			currentModel: &resource.Model{
				OrgId:       util.StringPtr("63350255419cf25e3d511c95"),
				Name:        util.StringPtr("test-service-account"),
				Description: util.StringPtr("Test description"),
			},
			expectedMsg: "required",
		},
		"Read_missingOrgId": {
			operation:    resource.Read,
			currentModel: &resource.Model{ClientId: util.StringPtr("mdb_sa_id_123456789")},
			expectedMsg:  "required",
		},
		"Read_missingClientId": {
			operation:    resource.Read,
			currentModel: &resource.Model{OrgId: util.StringPtr("63350255419cf25e3d511c95")},
			expectedMsg:  "required",
		},
		"Update_missingOrgId": {
			operation:    resource.Update,
			currentModel: &resource.Model{ClientId: util.StringPtr("mdb_sa_id_123456789")},
			expectedMsg:  "required",
		},
		"Update_missingClientId": {
			operation:    resource.Update,
			currentModel: &resource.Model{OrgId: util.StringPtr("63350255419cf25e3d511c95")},
			expectedMsg:  "required",
		},
		"Delete_missingOrgId": {
			operation:    resource.Delete,
			currentModel: &resource.Model{ClientId: util.StringPtr("mdb_sa_id_123456789")},
			expectedMsg:  "required",
		},
		"Delete_missingClientId": {
			operation:    resource.Delete,
			currentModel: &resource.Model{OrgId: util.StringPtr("63350255419cf25e3d511c95")},
			expectedMsg:  "required",
		},
		"List_missingOrgId": {
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

func TestCreateWithMocks(t *testing.T) {
	originalSetupRequest := resource.SetupRequest
	defer func() { resource.SetupRequest = originalSetupRequest }()

	testCases := map[string]struct {
		mockSetup      func(*mockadmin.ServiceAccountsApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
	}{
		"successfulCreate": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				resp := createTestOrgServiceAccountResponse()
				m.EXPECT().CreateOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.CreateOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().CreateOrgServiceAccountExecute(mock.Anything).
					Return(resp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, constants.Complete, event.Message)
				require.NotNil(t, event.ResourceModel)
				model := event.ResourceModel.(*resource.Model)
				assert.NotNil(t, model.ClientId)
				assert.NotNil(t, model.Secrets)
				if len(model.Secrets) > 0 {
					assert.NotNil(t, model.Secrets[0].Secret, "Secret should be present on create")
				}
			},
		},
		"createWithError": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				m.EXPECT().CreateOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.CreateOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().CreateOrgServiceAccountExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("API error"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockServiceAccountsAPI := mockadmin.NewServiceAccountsApi(t)
			tc.mockSetup(mockServiceAccountsAPI)

			mockClient := &admin.APIClient{ServiceAccountsApi: mockServiceAccountsAPI}
			mongoClient := &util.MongoDBClient{AtlasSDK: mockClient}

			resource.SetupRequest = func(req handler.Request, model *resource.Model, requiredFields []string) (*util.MongoDBClient, *handler.ProgressEvent) {
				return mongoClient, nil
			}

			event, err := resource.Create(handler.Request{}, nil, createTestOrgServiceAccountModel())
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)

			if tc.validateResult != nil {
				tc.validateResult(t, event)
			}
		})
	}
}

func TestReadWithMocks(t *testing.T) {
	originalSetupRequest := resource.SetupRequest
	defer func() { resource.SetupRequest = originalSetupRequest }()

	testCases := map[string]struct {
		mockSetup      func(*mockadmin.ServiceAccountsApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
	}{
		"successfulRead": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				resp := createTestOrgServiceAccountResponse()
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(resp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, constants.ReadComplete, event.Message)
				require.NotNil(t, event.ResourceModel)
				model := event.ResourceModel.(*resource.Model)
				assert.NotNil(t, model.ClientId)
				// Verify secrets are masked on read
				if model.Secrets != nil {
					for _, secret := range model.Secrets {
						assert.Nil(t, secret.Secret, "Secret should be masked on read")
						assert.NotNil(t, secret.MaskedSecretValue)
					}
				}
			},
		},
		"readNotFound": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
		"readWithError": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockServiceAccountsAPI := mockadmin.NewServiceAccountsApi(t)
			tc.mockSetup(mockServiceAccountsAPI)

			mockClient := &admin.APIClient{ServiceAccountsApi: mockServiceAccountsAPI}
			mongoClient := &util.MongoDBClient{AtlasSDK: mockClient}

			resource.SetupRequest = func(req handler.Request, model *resource.Model, requiredFields []string) (*util.MongoDBClient, *handler.ProgressEvent) {
				return mongoClient, nil
			}

			model := createTestOrgServiceAccountModel()
			clientID := "mdb_sa_id_123456789"
			model.ClientId = &clientID

			event, err := resource.Read(handler.Request{}, nil, model)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)

			if tc.validateResult != nil {
				tc.validateResult(t, event)
			}
		})
	}
}

func TestUpdateWithMocks(t *testing.T) {
	originalSetupRequest := resource.SetupRequest
	defer func() { resource.SetupRequest = originalSetupRequest }()

	testCases := map[string]struct {
		mockSetup      func(*mockadmin.ServiceAccountsApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
	}{
		"successfulUpdate": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				resp := createTestOrgServiceAccountResponse()
				updatedName := "updated-service-account"
				resp.Name = &updatedName
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(resp, &http.Response{StatusCode: 200}, nil)
				m.EXPECT().UpdateOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(admin.UpdateOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().UpdateOrgServiceAccountExecute(mock.Anything).
					Return(resp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, constants.Complete, event.Message)
				require.NotNil(t, event.ResourceModel)
				model := event.ResourceModel.(*resource.Model)
				assert.NotNil(t, model.ClientId)
				// Verify secrets are masked on update
				if model.Secrets != nil {
					for _, secret := range model.Secrets {
						assert.Nil(t, secret.Secret, "Secret should be masked on update")
					}
				}
			},
		},
		"updateNotFound": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
				assert.Equal(t, "Resource not found", event.Message)
			},
		},
		"updateWithError": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("internal server error"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockServiceAccountsAPI := mockadmin.NewServiceAccountsApi(t)
			tc.mockSetup(mockServiceAccountsAPI)

			mockClient := &admin.APIClient{ServiceAccountsApi: mockServiceAccountsAPI}
			mongoClient := &util.MongoDBClient{AtlasSDK: mockClient}

			resource.SetupRequest = func(req handler.Request, model *resource.Model, requiredFields []string) (*util.MongoDBClient, *handler.ProgressEvent) {
				return mongoClient, nil
			}

			model := createTestOrgServiceAccountModel()
			clientID := "mdb_sa_id_123456789"
			model.ClientId = &clientID

			event, err := resource.Update(handler.Request{}, nil, model)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)

			if tc.validateResult != nil {
				tc.validateResult(t, event)
			}
		})
	}
}

func TestDeleteWithMocks(t *testing.T) {
	originalSetupRequest := resource.SetupRequest
	defer func() { resource.SetupRequest = originalSetupRequest }()

	testCases := map[string]struct {
		mockSetup      func(*mockadmin.ServiceAccountsApi)
		expectedStatus handler.Status
	}{
		"successfulDelete": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				resp := createTestOrgServiceAccountResponse()
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(resp, &http.Response{StatusCode: 200}, nil)
				m.EXPECT().DeleteOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.DeleteOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().DeleteOrgServiceAccountExecute(mock.Anything).
					Return(&http.Response{StatusCode: 204}, nil)
			},
			expectedStatus: handler.Success,
		},
		"deleteNotFound": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
		"deleteWithError": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				resp := createTestOrgServiceAccountResponse()
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(resp, &http.Response{StatusCode: 200}, nil)
				m.EXPECT().DeleteOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.DeleteOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().DeleteOrgServiceAccountExecute(mock.Anything).
					Return(&http.Response{StatusCode: 500}, fmt.Errorf("delete failed"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockServiceAccountsAPI := mockadmin.NewServiceAccountsApi(t)
			tc.mockSetup(mockServiceAccountsAPI)

			mockClient := &admin.APIClient{ServiceAccountsApi: mockServiceAccountsAPI}
			mongoClient := &util.MongoDBClient{AtlasSDK: mockClient}

			resource.SetupRequest = func(req handler.Request, model *resource.Model, requiredFields []string) (*util.MongoDBClient, *handler.ProgressEvent) {
				return mongoClient, nil
			}

			model := createTestOrgServiceAccountModel()
			clientID := "mdb_sa_id_123456789"
			model.ClientId = &clientID

			event, err := resource.Delete(handler.Request{}, nil, model)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}

func TestListWithMocks(t *testing.T) {
	originalSetupRequest := resource.SetupRequest
	defer func() { resource.SetupRequest = originalSetupRequest }()

	testCases := map[string]struct {
		mockSetup      func(*mockadmin.ServiceAccountsApi)
		validateResult func(t *testing.T, event handler.ProgressEvent)
		expectedStatus handler.Status
	}{
		"successfulList": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				account1 := createTestOrgServiceAccountResponse()
				account2 := createTestOrgServiceAccountResponse()
				account2Name := "test-service-account-2"
				account2ClientID := "mdb_sa_id_987654321"
				account2.Name = &account2Name
				account2.ClientId = &account2ClientID

				results := []admin.OrgServiceAccount{*account1, *account2}
				totalCount := 2

				m.EXPECT().ListOrgServiceAccounts(mock.Anything, mock.Anything).
					Return(admin.ListOrgServiceAccountsApiRequest{ApiService: m})
				m.EXPECT().ListOrgServiceAccountsExecute(mock.Anything).
					Return(&admin.PaginatedOrgServiceAccounts{
						Results:    &results,
						TotalCount: &totalCount,
					}, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				require.NotNil(t, event.ResourceModels)
				assert.GreaterOrEqual(t, len(event.ResourceModels), 1)
				// Verify secrets are masked in list
				for _, rm := range event.ResourceModels {
					model := rm.(*resource.Model)
					if model.Secrets != nil {
						for _, secret := range model.Secrets {
							assert.Nil(t, secret.Secret, "Secret should be masked in list")
						}
					}
				}
			},
		},
		"listEmpty": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				results := []admin.OrgServiceAccount{}
				totalCount := 0

				m.EXPECT().ListOrgServiceAccounts(mock.Anything, mock.Anything).
					Return(admin.ListOrgServiceAccountsApiRequest{ApiService: m})
				m.EXPECT().ListOrgServiceAccountsExecute(mock.Anything).
					Return(&admin.PaginatedOrgServiceAccounts{
						Results:    &results,
						TotalCount: &totalCount,
					}, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validateResult: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				require.NotNil(t, event.ResourceModels)
				assert.Empty(t, event.ResourceModels)
			},
		},
		"listWithError": {
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				m.EXPECT().ListOrgServiceAccounts(mock.Anything, mock.Anything).
					Return(admin.ListOrgServiceAccountsApiRequest{ApiService: m})
				m.EXPECT().ListOrgServiceAccountsExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("list failed"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockServiceAccountsAPI := mockadmin.NewServiceAccountsApi(t)
			tc.mockSetup(mockServiceAccountsAPI)

			mockClient := &admin.APIClient{ServiceAccountsApi: mockServiceAccountsAPI}
			mongoClient := &util.MongoDBClient{AtlasSDK: mockClient}

			resource.SetupRequest = func(req handler.Request, model *resource.Model, requiredFields []string) (*util.MongoDBClient, *handler.ProgressEvent) {
				return mongoClient, nil
			}

			event, err := resource.List(handler.Request{}, nil, createTestOrgServiceAccountModel())
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
			event := resource.HandleError(tc.response, constants.CREATE, tc.err)
			assert.Equal(t, tc.expectedStatus, event.OperationStatus)
		})
	}
}
