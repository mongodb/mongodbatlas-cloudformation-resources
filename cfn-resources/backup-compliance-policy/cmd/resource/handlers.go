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

package resource

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"go.mongodb.org/atlas-sdk/v20250312012/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

const (
	policyStateActive                        = "ACTIVE"
	errorCannotUpdateWithPendingAction       = "CANNOT_UPDATE_BACKUP_COMPLIANCE_POLICY_SETTINGS_WITH_PENDING_ACTION"
	errorCannotDisableBackupCompliancePolicy = "CANNOT_DISABLE_BACKUP_COMPLIANCE_POLICY"
)

var callbackContext = map[string]any{"callbackBackupCompliancePolicy": true}

func IsCallback(req *handler.Request) bool {
	_, found := req.CallbackContext["callbackBackupCompliancePolicy"]
	return found
}

func isPendingActionError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), errorCannotUpdateWithPendingAction)
}

func handlePendingAction(ctx context.Context, client *util.MongoDBClient, model *Model, projectID string) handler.ProgressEvent {
	policy, _, getErr := client.AtlasSDK.CloudBackupsApi.GetCompliancePolicy(ctx, projectID).Execute()
	if getErr == nil && policy != nil {
		SetBackupCompliancePolicyData(model, policy)
	}
	return inProgressEvent(model, policy)
}

func checkPolicyNotFound(policy *admin.DataProtectionSettings20231001, apiResp *http.Response, err error, projectID string, operation constants.CfnFunctions) *handler.ProgressEvent {
	if err != nil {
		if util.StatusNotFound(apiResp) {
			pe := progress_events.GetFailedEventByCode(
				"Backup Compliance Policy not found for project: "+projectID,
				string(types.HandlerErrorCodeNotFound))
			return &pe
		}
		pe := handleError(apiResp, operation, err)
		return &pe
	}

	if policy != nil && policy.GetProjectId() == "" {
		pe := progress_events.GetFailedEventByCode(
			"Backup Compliance Policy not found for project: "+projectID,
			string(types.HandlerErrorCodeNotFound))
		return &pe
	}

	return nil
}

func HandleCreate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	if IsCallback(req) {
		return validateProgress(client, model, false, constants.CREATE)
	}

	ctx := context.Background()
	projectID := *model.ProjectId

	existingPolicy, apiResp, err := client.AtlasSDK.CloudBackupsApi.GetCompliancePolicy(ctx, projectID).Execute()
	if err != nil {
		if !util.StatusNotFound(apiResp) {
			return handleError(apiResp, constants.CREATE, err)
		}
	} else if existingPolicy != nil && existingPolicy.GetState() == policyStateActive {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Backup Compliance Policy already exists for project: " + projectID,
			HandlerErrorCode: string(types.HandlerErrorCodeAlreadyExists),
		}
	}

	dataProtectionSettings := ExpandDataProtectionSettings(model, projectID)

	params := admin.UpdateCompliancePolicyApiParams{
		GroupId:                        projectID,
		DataProtectionSettings20231001: dataProtectionSettings,
		OverwriteBackupPolicies:        util.Pointer(false),
	}

	_, apiResp, err = client.AtlasSDK.CloudBackupsApi.UpdateCompliancePolicyWithParams(ctx, &params).Execute()
	if err != nil {
		if isPendingActionError(err) {
			return handlePendingAction(ctx, client, model, projectID)
		}
		return handleError(apiResp, constants.CREATE, err)
	}

	policy, apiResp, err := client.AtlasSDK.CloudBackupsApi.GetCompliancePolicy(ctx, projectID).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}

	return inProgressEvent(model, policy)
}

func HandleRead(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	projectID := *model.ProjectId

	policy, apiResp, err := client.AtlasSDK.CloudBackupsApi.GetCompliancePolicy(ctx, projectID).Execute()
	if pe := checkPolicyNotFound(policy, apiResp, err, projectID, constants.READ); pe != nil {
		return *pe
	}

	SetBackupCompliancePolicyData(model, policy)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   model,
	}
}

func HandleUpdate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	if IsCallback(req) {
		return validateProgress(client, model, false, constants.UPDATE)
	}

	ctx := context.Background()
	projectID := *model.ProjectId

	policy, apiResp, err := client.AtlasSDK.CloudBackupsApi.GetCompliancePolicy(ctx, projectID).Execute()
	if pe := checkPolicyNotFound(policy, apiResp, err, projectID, constants.UPDATE); pe != nil {
		return *pe
	}

	dataProtectionSettings := ExpandDataProtectionSettings(model, projectID)

	params := admin.UpdateCompliancePolicyApiParams{
		GroupId:                        projectID,
		DataProtectionSettings20231001: dataProtectionSettings,
		OverwriteBackupPolicies:        util.Pointer(false),
	}

	_, apiResp, err = client.AtlasSDK.CloudBackupsApi.UpdateCompliancePolicyWithParams(ctx, &params).Execute()
	if err != nil {
		if isPendingActionError(err) {
			return handlePendingAction(ctx, client, model, projectID)
		}
		return handleError(apiResp, constants.UPDATE, err)
	}

	policy, apiResp, err = client.AtlasSDK.CloudBackupsApi.GetCompliancePolicy(ctx, projectID).Execute()
	if err != nil {
		return handleError(apiResp, constants.UPDATE, err)
	}

	return inProgressEvent(model, policy)
}

func HandleDelete(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	if IsCallback(req) {
		return validateProgress(client, model, true, constants.DELETE)
	}

	ctx := context.Background()
	projectID := *model.ProjectId

	policy, apiResp, err := client.AtlasSDK.CloudBackupsApi.GetCompliancePolicy(ctx, projectID).Execute()
	if pe := checkPolicyNotFound(policy, apiResp, err, projectID, constants.DELETE); pe != nil {
		return *pe
	}

	apiResp, err = client.AtlasSDK.CloudBackupsApi.DisableCompliancePolicy(ctx, projectID).Execute()
	if err == nil {
		return inProgressEvent(model, nil)
	}

	errorMessage := err.Error()

	if strings.Contains(errorMessage, errorCannotDisableBackupCompliancePolicy) {
		_, _ = logger.Warnf("Cannot disable backup compliance policy for project %s: Policy requires MongoDB support or removing all clusters & retained snapshots", projectID)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          fmt.Sprintf("Cannot disable Backup Compliance Policy for project %s. Disabling BCP requires MongoDB support or removing all clusters & retained snapshots. This can only be done by filing a support ticket.", projectID),
			HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
		}
	}

	if isPendingActionError(err) {
		return handlePendingAction(ctx, client, model, projectID)
	}

	if util.StatusNotFound(apiResp) {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         constants.Complete,
		}
	}

	_, _ = logger.Warnf("Unexpected delete error for project %s: %s", projectID, errorMessage)
	return handleError(apiResp, constants.DELETE, err)
}

func HandleList(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()
	projectID := *model.ProjectId

	policy, apiResp, err := client.AtlasSDK.CloudBackupsApi.GetCompliancePolicy(ctx, projectID).Execute()
	if err != nil {
		if !util.StatusNotFound(apiResp) {
			return handleError(apiResp, constants.LIST, err)
		}
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         constants.Complete,
			ResourceModels:  []any{},
		}
	}

	if policy != nil && policy.GetProjectId() == "" {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         constants.Complete,
			ResourceModels:  []any{},
		}
	}

	SetBackupCompliancePolicyData(model, policy)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModels:  []any{model},
	}
}

func inProgressEvent(model *Model, policy *admin.DataProtectionSettings20231001) handler.ProgressEvent {
	if policy != nil {
		SetBackupCompliancePolicyData(model, policy)
	}
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              constants.Pending,
		ResourceModel:        model,
		CallbackDelaySeconds: callBackSeconds,
		CallbackContext:      callbackContext,
	}
}

func validateProgress(client *util.MongoDBClient, model *Model, isDelete bool, operation constants.CfnFunctions) handler.ProgressEvent {
	ctx := context.Background()
	projectID := *model.ProjectId

	policy, resp, err := client.AtlasSDK.CloudBackupsApi.GetCompliancePolicy(ctx, projectID).Execute()
	notFound := util.StatusNotFound(resp)
	policyDeleted := notFound || (policy != nil && policy.GetProjectId() == "")

	if err != nil && !notFound {
		if isPendingActionError(err) {
			return inProgressEvent(model, policy)
		}
		return handleError(resp, operation, err)
	}

	if isDelete && policyDeleted {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         constants.Complete,
		}
	}

	if isDelete {
		// Don't call DisableCompliancePolicy again - it always returns 204
		return inProgressEvent(model, nil)
	}

	if policy != nil {
		SetBackupCompliancePolicyData(model, policy)
		if policy.GetState() == policyStateActive {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         constants.Complete,
				ResourceModel:   model,
			}
		}
		return inProgressEvent(model, policy)
	}

	if notFound {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Backup Compliance Policy was deleted during operation",
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
		}
	}

	return inProgressEvent(model, policy)
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) handler.ProgressEvent {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())
	return progress_events.GetFailedEventByResponse(errMsg, response)
}
