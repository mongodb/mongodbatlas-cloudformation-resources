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

func createTestModel() *resource.Model {
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

func createTestResponse() *admin.OrgServiceAccount {
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

func TestCRUDOperations(t *testing.T) {
	originalSetupRequest := resource.SetupRequest
	defer func() { resource.SetupRequest = originalSetupRequest }()

	tests := []struct {
		name           string
		operation      func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error)
		setupModel     func() *resource.Model
		mockSetup      func(*mockadmin.ServiceAccountsApi)
		validate       func(*testing.T, handler.ProgressEvent)
		expectedStatus handler.Status
	}{
		{
			name:       "Create_Success",
			operation:  resource.Create,
			setupModel: createTestModel,
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				resp := createTestResponse()
				m.EXPECT().CreateOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.CreateOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().CreateOrgServiceAccountExecute(mock.Anything).
					Return(resp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validate: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, constants.Complete, event.Message)
				require.NotNil(t, event.ResourceModel)
				model := event.ResourceModel.(*resource.Model)
				assert.NotNil(t, model.ClientId)
				if len(model.Secrets) > 0 {
					assert.NotNil(t, model.Secrets[0].Secret, "Secret should be present on create")
				}
			},
		},
		{
			name:       "Create_Error",
			operation:  resource.Create,
			setupModel: createTestModel,
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				m.EXPECT().CreateOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.CreateOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().CreateOrgServiceAccountExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("API error"))
			},
			expectedStatus: handler.Failed,
		},
		{
			name:      "Read_Success",
			operation: resource.Read,
			setupModel: func() *resource.Model {
				model := createTestModel()
				clientID := "mdb_sa_id_123456789"
				model.ClientId = &clientID
				return model
			},
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				resp := createTestResponse()
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(resp, &http.Response{StatusCode: 200}, nil)
			},
			expectedStatus: handler.Success,
			validate: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, constants.ReadComplete, event.Message)
				model := event.ResourceModel.(*resource.Model)
				if model.Secrets != nil {
					for _, secret := range model.Secrets {
						assert.Nil(t, secret.Secret, "Secret should be masked on read")
					}
				}
			},
		},
		{
			name:      "Read_NotFound",
			operation: resource.Read,
			setupModel: func() *resource.Model {
				model := createTestModel()
				clientID := "mdb_sa_id_123456789"
				model.ClientId = &clientID
				return model
			},
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
		{
			name:      "Update_Success",
			operation: resource.Update,
			setupModel: func() *resource.Model {
				model := createTestModel()
				clientID := "mdb_sa_id_123456789"
				model.ClientId = &clientID
				return model
			},
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				resp := createTestResponse()
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
			validate: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				model := event.ResourceModel.(*resource.Model)
				if model.Secrets != nil {
					for _, secret := range model.Secrets {
						assert.Nil(t, secret.Secret, "Secret should be masked on update")
					}
				}
			},
		},
		{
			name:      "Update_NotFound",
			operation: resource.Update,
			setupModel: func() *resource.Model {
				model := createTestModel()
				clientID := "mdb_sa_id_123456789"
				model.ClientId = &clientID
				return model
			},
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
			validate: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Equal(t, string(types.HandlerErrorCodeNotFound), event.HandlerErrorCode)
			},
		},
		{
			name:      "Delete_Success",
			operation: resource.Delete,
			setupModel: func() *resource.Model {
				model := createTestModel()
				clientID := "mdb_sa_id_123456789"
				model.ClientId = &clientID
				return model
			},
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				resp := createTestResponse()
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
		{
			name:      "Delete_NotFound",
			operation: resource.Delete,
			setupModel: func() *resource.Model {
				model := createTestModel()
				clientID := "mdb_sa_id_123456789"
				model.ClientId = &clientID
				return model
			},
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				m.EXPECT().GetOrgServiceAccount(mock.Anything, mock.Anything, mock.Anything).
					Return(admin.GetOrgServiceAccountApiRequest{ApiService: m})
				m.EXPECT().GetOrgServiceAccountExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 404}, fmt.Errorf("not found"))
			},
			expectedStatus: handler.Failed,
		},
		{
			name:       "List_Success",
			operation:  resource.List,
			setupModel: createTestModel,
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				account1 := createTestResponse()
				account2 := createTestResponse()
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
			validate: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				require.NotNil(t, event.ResourceModels)
				assert.GreaterOrEqual(t, len(event.ResourceModels), 1)
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
		{
			name:       "List_Empty",
			operation:  resource.List,
			setupModel: createTestModel,
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
			validate: func(t *testing.T, event handler.ProgressEvent) {
				t.Helper()
				assert.Empty(t, event.ResourceModels)
			},
		},
		{
			name:       "List_Error",
			operation:  resource.List,
			setupModel: createTestModel,
			mockSetup: func(m *mockadmin.ServiceAccountsApi) {
				m.EXPECT().ListOrgServiceAccounts(mock.Anything, mock.Anything).
					Return(admin.ListOrgServiceAccountsApiRequest{ApiService: m})
				m.EXPECT().ListOrgServiceAccountsExecute(mock.Anything).
					Return(nil, &http.Response{StatusCode: 500}, fmt.Errorf("list failed"))
			},
			expectedStatus: handler.Failed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServiceAccountsAPI := mockadmin.NewServiceAccountsApi(t)
			tt.mockSetup(mockServiceAccountsAPI)

			mockClient := &admin.APIClient{ServiceAccountsApi: mockServiceAccountsAPI}
			mongoClient := &util.MongoDBClient{AtlasSDK: mockClient}

			resource.SetupRequest = func(req handler.Request, model *resource.Model, requiredFields []string) (*util.MongoDBClient, *handler.ProgressEvent) {
				return mongoClient, nil
			}

			event, err := tt.operation(handler.Request{}, nil, tt.setupModel())
			require.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, event.OperationStatus)

			if tt.validate != nil {
				tt.validate(t, event)
			}
		})
	}
}

func TestHandleError(t *testing.T) {
	tests := []struct {
		name           string
		response       *http.Response
		err            error
		expectedStatus handler.Status
	}{
		{"NotFound", &http.Response{StatusCode: http.StatusNotFound}, fmt.Errorf("not found"), handler.Failed},
		{"InternalServerError", &http.Response{StatusCode: http.StatusInternalServerError}, fmt.Errorf("server error"), handler.Failed},
		{"BadRequest", &http.Response{StatusCode: http.StatusBadRequest}, fmt.Errorf("bad request"), handler.Failed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := resource.HandleError(tt.response, constants.CREATE, tt.err)
			assert.Equal(t, tt.expectedStatus, event.OperationStatus)
		})
	}
}
