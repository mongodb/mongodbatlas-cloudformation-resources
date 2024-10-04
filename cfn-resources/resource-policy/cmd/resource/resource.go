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

package resource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20240805004/admin"
)

var CreateRequiredFields = []string{"OrgId", "Name", "Policies"}
var ReadRequiredFields = []string{"OrgId", "Id"}
var UpdateRequiredFields = []string{"OrgId", "Id"}
var DeleteRequiredFields = []string{"OrgId", "Id"}
var ListRequiredFields = []string{"OrgId"}

func initEnvWithLatestClient(req handler.Request, currentModel *Model, requiredFields []string) (*admin.APIClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-resource-policy")

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if errEvent := validator.ValidateModel(requiredFields, currentModel); errEvent != nil {
		return nil, errEvent
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return nil, peErr
	}
	return client.AtlasSDK, nil
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := initEnvWithLatestClient(req, currentModel, CreateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	orgID := currentModel.OrgId
	resourcePolicyReq := NewResourcePolicyCreateReq(currentModel)
	resourcePolicyResp, apiResp, err := conn.ResourcePoliciesApi.CreateAtlasResourcePolicy(ctx, *orgID, resourcePolicyReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}
	resourceModel := GetResourcePolicyModel(resourcePolicyResp, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   resourceModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := initEnvWithLatestClient(req, currentModel, ReadRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	orgID := currentModel.OrgId
	resourcePolicyID := currentModel.Id
	resourcePolicyResp, apiResp, err := conn.ResourcePoliciesApi.GetAtlasResourcePolicy(ctx, *orgID, *resourcePolicyID).Execute()
	if err != nil {
		return handleError(apiResp, constants.READ, err)
	}

	resourceModel := GetResourcePolicyModel(resourcePolicyResp, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   resourceModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := initEnvWithLatestClient(req, currentModel, UpdateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	orgID := currentModel.OrgId
	resourcePolicyID := currentModel.Id
	_, _, err := conn.ResourcePoliciesApi.GetAtlasResourcePolicy(ctx, *orgID, *resourcePolicyID).Execute()
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Update failed",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
		}, nil
	}
	resourcePolicyReq := NewResourcePolicyUpdateReq(currentModel)
	resourcePolicyResp, apiResp, err := conn.ResourcePoliciesApi.UpdateAtlasResourcePolicy(ctx, *orgID, *resourcePolicyID, resourcePolicyReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.UPDATE, err)
	}

	resourceModel := GetResourcePolicyModel(resourcePolicyResp, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   resourceModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := initEnvWithLatestClient(req, currentModel, DeleteRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	orgID := currentModel.OrgId
	resourcePolicyID := currentModel.Id
	_, apiResp, err := conn.ResourcePoliciesApi.DeleteAtlasResourcePolicy(ctx, *orgID, *resourcePolicyID).Execute()
	if err != nil {
		return handleError(apiResp, constants.DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
		ResourceModel:   nil}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := initEnvWithLatestClient(req, currentModel, ListRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	orgID := currentModel.OrgId

	resourcePolicies, apiResp, err := conn.ResourcePoliciesApi.GetAtlasResourcePolicies(ctx, *orgID).Execute()
	if err != nil {
		return handleError(apiResp, constants.LIST, err)
	}

	response := make([]interface{}, 0)
	for i := range resourcePolicies {
		model := GetResourcePolicyModel(&resourcePolicies[i], nil)
		model.OrgId = currentModel.OrgId
		model.Profile = currentModel.Profile

		response = append(response, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModels:  response,
	}, nil
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())
	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}
