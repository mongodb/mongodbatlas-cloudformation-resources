package resource

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

const (
	CallBackSeconds = 10 // TODO temporary, must remove
	// CallBackSeconds = 40
	SearchDeploymentDoesNotExistsError = "ATLAS_FTS_DEPLOYMENT_DOES_NOT_EXIST"
)

// TODO: complete required fields
var CreateRequiredFields = []string{}
var ReadRequiredFields = []string{}
var UpdateRequiredFields = []string{}
var DeleteRequiredFields = []string{}

func setup() {
	util.SetupLogger("mongodb-atlas-searchdeployment")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, progressErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if progressErr != nil {
		return *progressErr, nil
	}
	connV2 := client.AtlasSDK

	// handling of subsequent retry calls
	if _, ok := req.CallbackContext[constants.ID]; ok {
		return handleStateTransition(*connV2, currentModel, constants.IdleState), nil
	}

	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	apiReq := newSearchDeploymentReq(*currentModel)
	apiResp, res, err := connV2.AtlasSearchApi.CreateAtlasSearchDeployment(context.Background(), projectID, clusterName, &apiReq).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Search Deployment: %s", err.Error()),
			res), nil
	}

	newModel := newCFNSearchDeployment(currentModel, apiResp)
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		ResourceModel:        newModel,
		Message:              fmt.Sprintf("Create Search Deployment `%s`", *apiResp.StateName),
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.ID: newModel.Id,
		},
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	// Validation
	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, progressErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if progressErr != nil {
		return *progressErr, nil
	}
	connV2 := client.AtlasSDK

	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	apiResp, res, err := connV2.AtlasSearchApi.GetAtlasSearchDeployment(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   newCFNSearchDeployment(currentModel, apiResp),
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	// Validation
	modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	// Validation
	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, progressErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if progressErr != nil {
		return *progressErr, nil
	}
	connV2 := client.AtlasSDK

	// handling of subsequent retry calls
	if _, ok := req.CallbackContext[constants.ID]; ok {
		return handleStateTransition(*connV2, currentModel, constants.DeletedState), nil
	}

	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	if resp, err := connV2.AtlasSearchApi.DeleteAtlasSearchDeployment(context.Background(), projectID, clusterName).Execute(); err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              constants.DeleteInProgress,
		ResourceModel:        currentModel,
		CallbackDelaySeconds: CallBackSeconds,
		CallbackContext: map[string]interface{}{
			constants.ID: currentModel.Id,
		}}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func handleStateTransition(connV2 admin.APIClient, currentModel *Model, targetState string) handler.ProgressEvent {
	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	apiResp, resp, err := connV2.AtlasSearchApi.GetAtlasSearchDeployment(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if targetState == constants.DeletedState && resp.StatusCode == 400 && strings.Contains(err.Error(), SearchDeploymentDoesNotExistsError) {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				ResourceModel:   currentModel,
				Message:         constants.Complete,
			}
		}

		return progressevent.GetFailedEventByResponse(err.Error(), resp)
	}

	newModel := newCFNSearchDeployment(currentModel, apiResp)
	if util.SafeString(currentModel.StateName) == targetState {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			ResourceModel:   newModel,
			Message:         constants.Complete,
		}
	}
	return handler.ProgressEvent{
		OperationStatus: handler.InProgress,
		ResourceModel:   newModel,
		Message:         constants.Pending,
		CallbackContext: map[string]interface{}{
			constants.ID: currentModel.Id,
		},
	}
}

func newCFNSearchDeployment(prevModel *Model, apiResp *admin.ApiSearchDeploymentResponse) Model {
	respSpecs := apiResp.GetSpecs()
	resultSpecs := make([]ApiSearchDeploymentSpec, len(respSpecs))
	for i, respSpec := range respSpecs {
		resultSpecs[i] = ApiSearchDeploymentSpec{
			InstanceSize: &respSpec.InstanceSize,
			NodeCount:    &respSpec.NodeCount,
		}
	}
	return Model{
		Profile:     prevModel.Profile,
		ClusterName: prevModel.ClusterName,
		ProjectId:   prevModel.ProjectId,
		Id:          apiResp.Id,
		Specs:       resultSpecs,
		StateName:   apiResp.StateName,
	}
}

func newSearchDeploymentReq(model Model) admin.ApiSearchDeploymentRequest {
	modelSpecs := model.Specs
	requestSpecs := make([]admin.ApiSearchDeploymentSpec, len(modelSpecs))
	for i, spec := range modelSpecs {
		requestSpecs[i] = admin.ApiSearchDeploymentSpec{
			InstanceSize: *spec.InstanceSize,
			NodeCount:    *spec.NodeCount, // TODO verify what happens if one of the properties is not defined
		}
	}
	return admin.ApiSearchDeploymentRequest{Specs: &requestSpecs}
}
