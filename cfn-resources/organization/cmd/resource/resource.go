// Copyright 2023 MongoDB Inc
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

package resource

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/atlas-sdk/v20250312012/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/secrets"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

var CreateRequiredFields = []string{constants.OrgOwnerID, constants.Name, constants.AwsSecretName, constants.OrgKeyDescription, constants.OrgKeyRoles}
var UpdateRequiredFields = []string{constants.OrgID, constants.Name, constants.AwsSecretName}
var ReadRequiredFields = []string{constants.OrgID, constants.AwsSecretName}
var DeleteRequiredFields = []string{constants.OrgID, constants.AwsSecretName}

const (
	CallBackSeconds  = 20
	DeletingState    = "Deleting"
	DeleteInProgress = "Delete Organization is in progress"
	DeleteCompleted  = "Delete Organization is completed"
)

type OrgProfile struct {
	OrgID      string
	PublicKey  string
	PrivateKey string
	BaseURL    string
}

type DeleteResponse struct {
	Error    error
	Response *http.Response
}

func setup() {
	util.SetupLogger("mongodb-atlas-organization")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	conn := client.AtlasSDK
	ctx := context.Background()

	_, _, err := secrets.Get(&req, *currentModel.AwsSecretName)
	if err != nil {
		// Delete the APIKey from Atlas
		_, _ = logger.Warnf("error : no Secret exists with %s", *currentModel.AwsSecretName)
		response := &http.Response{StatusCode: http.StatusBadRequest}
		return handleError(response, constants.CREATE, err)
	}

	apikeyInputs := setAPIkeyInputs(currentModel)
	setDefaultsIfNotDefined(currentModel)

	// Set the roles from model
	orgInput := &admin.CreateOrganizationRequest{
		ApiKey:                    apikeyInputs,
		OrgOwnerId:                currentModel.OrgOwnerId,
		Name:                      *currentModel.Name,
		SkipDefaultAlertsSettings: currentModel.SkipDefaultAlertsSettings,
	}
	if currentModel.FederatedSettingsId != nil {
		orgInput.FederationSettingsId = currentModel.FederatedSettingsId
	}
	org, response, err := conn.OrganizationsApi.CreateOrg(ctx, orgInput).Execute()
	if err != nil {
		return handleError(response, constants.CREATE, err)
	}

	orgID := org.Organization.GetId()

	// Read response
	currentModel.OrgId = &orgID

	// Save PrivateKey in AWS SecretManager
	secret := OrgProfile{OrgID: *currentModel.OrgId, PublicKey: *org.ApiKey.PublicKey, PrivateKey: *org.ApiKey.PrivateKey, BaseURL: client.Config.BaseURL}
	_, _, err = secrets.PutSecret(&req, *currentModel.AwsSecretName, secret, currentModel.APIKey.Description)
	if err != nil {
		// Delete the APIKey from Atlas
		response = &http.Response{StatusCode: http.StatusInternalServerError}
		return handleError(response, constants.CREATE, err)
	}

	newOrgClient, peErr := util.NewAtlasClientRemovingProfilePrefix(&req, currentModel.AwsSecretName)
	if peErr != nil {
		return *peErr, nil
	}
	conn = newOrgClient.AtlasSDK
	if _, _, errUpdate := conn.OrganizationsApi.UpdateOrgSettings(ctx, orgID, newOrganizationSettings(currentModel)).Execute(); errUpdate != nil {
		return handleError(response, constants.CREATE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   currentModel}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	newOrgClient, peErr := util.NewAtlasClientRemovingProfilePrefix(&req, currentModel.AwsSecretName)
	if peErr != nil {
		return *peErr, nil
	}

	model, response, err := currentModel.getOrgDetails(context.Background(), newOrgClient.AtlasSDK, currentModel)
	if err != nil {
		return handleError(response, constants.READ, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
		ResourceModel:   model}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	newOrgClient, peErr := util.NewAtlasClientRemovingProfilePrefix(&req, currentModel.AwsSecretName)
	if peErr != nil {
		return *peErr, nil
	}
	conn := newOrgClient.AtlasSDK
	ctx := context.Background()

	setDefaultsIfNotDefined(currentModel)
	atlasOrg := admin.AtlasOrganization{Id: currentModel.OrgId, Name: *currentModel.Name, SkipDefaultAlertsSettings: currentModel.SkipDefaultAlertsSettings}

	if _, response, err := conn.OrganizationsApi.UpdateOrg(ctx, *currentModel.OrgId, &atlasOrg).Execute(); err != nil {
		return handleError(response, constants.UPDATE, err)
	}

	if _, response, err := conn.OrganizationsApi.UpdateOrgSettings(ctx, *currentModel.OrgId, newOrganizationSettings(currentModel)).Execute(); err != nil {
		return handleError(response, constants.UPDATE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   currentModel}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	newOrgClient, peErr := util.NewAtlasClientRemovingProfilePrefix(&req, currentModel.AwsSecretName)
	if peErr != nil {
		return *peErr, nil
	}
	conn := newOrgClient.AtlasSDK
	ctx := context.Background()

	// Callback
	if _, idExists := req.CallbackContext[constants.StateName]; idExists {
		return deleteCallback(ctx, conn, currentModel)
	}

	// Read before delete
	_, response, err := currentModel.getOrgDetails(ctx, conn, currentModel)
	if err != nil {
		return handleError(response, constants.DELETE, err)
	}

	// If exists
	_, response, err = currentModel.getOrgDetails(ctx, conn, currentModel)
	if err != nil && response.StatusCode == http.StatusUnauthorized {
		return handleError(response, constants.DELETE, err)
	}

	deleteRequest := conn.OrganizationsApi.DeleteOrg(ctx, *currentModel.OrgId)

	// TODO delete

	// Since the Delete API is synchronous and takes more than 1 minute most of the time,
	// we need to make the call in a goroutine and return a progress event
	// after 10 Seconds. Reason for wait is that the Delete API
	// may throw error immediately if the resource is not found.

	responseChan := make(chan DeleteResponse, 1)
	go func() {
		response, err := deleteRequest.Execute()
		responseChan <- DeleteResponse{Error: err, Response: response}
	}()

	currentModel.IsDeleted = util.Pointer(false)
	select {
	case responseMsg := <-responseChan:
		if responseMsg.Error != nil {
			return handleError(responseMsg.Response, constants.DELETE, responseMsg.Error)
		}

	case <-time.After(30 * time.Second):
		// If the Delete is not completed in the above time,
		// we return a progress event with inProgress status and callback context
		return handler.ProgressEvent{
			OperationStatus:      handler.InProgress,
			Message:              DeleteInProgress,
			ResourceModel:        currentModel,
			CallbackDelaySeconds: CallBackSeconds,
			CallbackContext: map[string]interface{}{
				constants.StateName: DeletingState,
			},
		}, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         DeleteCompleted,
		ResourceModel:   nil}, nil
}

func deleteCallback(ctx context.Context, conn *admin.APIClient, currentModel *Model) (handler.ProgressEvent, error) {
	// Read before delete
	org, response, err := currentModel.getOrgDetails(ctx, conn, currentModel)
	if err != nil {
		if response.StatusCode == http.StatusUnauthorized {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         DeleteCompleted,
				ResourceModel:   nil}, nil
		}
		return handleError(response, constants.DELETE, err)
	}

	if *org.IsDeleted {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         DeleteCompleted,
			ResourceModel:   nil}, nil
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              DeleteInProgress,
		ResourceModel:        currentModel,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.StateName: DeletingState,
		},
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func (model *Model) getOrgDetails(ctx context.Context, conn *admin.APIClient, currentModel *Model) (responseModel *Model, response *http.Response, err error) {
	org, response, err := conn.OrganizationsApi.GetOrg(ctx, *currentModel.OrgId).Execute()
	if err != nil {
		return nil, response, err
	}
	model.Name = util.Pointer(org.Name)
	model.OrgId = org.Id
	model.IsDeleted = org.IsDeleted
	model.SkipDefaultAlertsSettings = org.SkipDefaultAlertsSettings

	settings, _, err := conn.OrganizationsApi.GetOrgSettings(ctx, org.GetId()).Execute()
	if err != nil {
		return nil, response, err
	}
	model.ApiAccessListRequired = settings.ApiAccessListRequired
	model.MultiFactorAuthRequired = settings.MultiFactorAuthRequired
	model.RestrictEmployeeAccess = settings.RestrictEmployeeAccess
	model.GenAIFeaturesEnabled = settings.GenAIFeaturesEnabled

	return model, response, nil
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())
	_, _ = logger.Warn(errMsg)
	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: string(types.HandlerErrorCodeAlreadyExists)}, nil
	}

	if response.StatusCode == http.StatusUnauthorized {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Not found",
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
	}

	if response.StatusCode == http.StatusBadRequest {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
	}
	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}

func setAPIkeyInputs(currentModel *Model) (apiKeyInput *admin.CreateAtlasOrganizationApiKey) {
	apiKeyInput = &admin.CreateAtlasOrganizationApiKey{
		Desc:  util.SafeString(currentModel.APIKey.Description),
		Roles: currentModel.APIKey.Roles,
	}
	return apiKeyInput
}

func newOrganizationSettings(model *Model) *admin.OrganizationSettings {
	return &admin.OrganizationSettings{
		ApiAccessListRequired:   model.ApiAccessListRequired,
		MultiFactorAuthRequired: model.MultiFactorAuthRequired,
		RestrictEmployeeAccess:  model.RestrictEmployeeAccess,
		GenAIFeaturesEnabled:    model.GenAIFeaturesEnabled,
	}
}

func setDefaultsIfNotDefined(m *Model) {
	if m == nil {
		return
	}
	if m.SkipDefaultAlertsSettings == nil {
		m.SkipDefaultAlertsSettings = util.Pointer(true)
	}
	if m.GenAIFeaturesEnabled == nil {
		m.GenAIFeaturesEnabled = util.Pointer(true)
	}
}
