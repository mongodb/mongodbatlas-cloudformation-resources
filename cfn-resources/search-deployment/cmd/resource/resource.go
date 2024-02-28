package resource

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
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

	// Validation
	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, err := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if err != nil {
		return *err, nil
	}
	connV2 := client.AtlasSDK

	projectID := *currentModel.ProjectId
	clusterName := *currentModel.ClusterName // TODO remove pointer access
	// TODO: searchDeploymentReq := NewSearchDeploymentReq(ctx, &searchDeploymentPlan)
	if _, res, err := connV2.AtlasSearchApi.CreateAtlasSearchDeployment(context.Background(), projectID, clusterName, &admin.ApiSearchDeploymentRequest{
		Specs: &[]admin.ApiSearchDeploymentSpec{
			{
				InstanceSize: *currentModel.Specs[0].InstanceSize,
				NodeCount:    *currentModel.Specs[0].NodeCount,
			},
		},
	}).Execute(); err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Search Deployment: %s", err.Error()),
			res), nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
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

	fmt.Printf("running read: %v", *currentModel)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
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

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}
