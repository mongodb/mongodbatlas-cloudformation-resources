package resource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20240805004/admin"
)

var CreateRequiredFields = []string{"Name", "Policies", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var ReadRequiredFields = []string{"OrgId", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var UpdateRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var DeleteRequiredFields = []string{"OrgId", "ResourcePolicyId", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var ListRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey"}

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
	resourcePolicyResp, apiResp, err := conn.AtlasResourcePoliciesApi.CreateAtlasResourcePolicy(ctx, *orgID, resourcePolicyReq).Execute()
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
	conn, peErr := initEnvWithLatestClient(req, currentModel, CreateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	orgID := currentModel.OrgId
	resourcePolicyID := currentModel.ResourcePolicyId
	resourcePolicyResp, apiResp, err := conn.AtlasResourcePoliciesApi.GetAtlasResourcePolicy(ctx, *orgID, *resourcePolicyID).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}

	resourceModel := GetResourcePolicyModel(resourcePolicyResp, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   resourceModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := initEnvWithLatestClient(req, currentModel, CreateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	ctx := context.Background()

	orgID := currentModel.OrgId
	resourcePolicyID := currentModel.ResourcePolicyId
	resourcePolicyReq := NewResourcePolicyUpdateReq(currentModel)
	resourcePolicyResp, apiResp, err := conn.AtlasResourcePoliciesApi.UpdateAtlasResourcePolicy(ctx, *orgID, *resourcePolicyID, resourcePolicyReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
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
	resourcePolicyID := currentModel.ResourcePolicyId
	if _, apiResp, err := conn.AtlasResourcePoliciesApi.DeleteAtlasResourcePolicy(ctx, *orgID, *resourcePolicyID).Execute(); err != nil {
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

	resourcePolicies, apiResp, err := conn.AtlasResourcePoliciesApi.GetAtlasResourcePolicies(ctx, *orgID).Execute()
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
