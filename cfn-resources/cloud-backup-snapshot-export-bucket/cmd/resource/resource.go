package resource

import (
	"context"
	"errors"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{"GroupId", "BucketName", "IamRoleId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var ReadRequiredFields = []string{"GroupId", "Id", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{"GroupId", "Id", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}
var ListRequiredFields = []string{"GroupId", "ApiKeys.PrivateKey", "ApiKeys.PublicKey"}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-CloudBackupSnapshotExportBucket")
}

func (m *Model) completeByAtlasModel(bucket *mongodbatlas.CloudProviderSnapshotExportBucket) {
	m.Id = &bucket.ID
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	input := &mongodbatlas.CloudProviderSnapshotExportBucket{
		BucketName:    *currentModel.BucketName,
		CloudProvider: constants.AWS,
		IAMRoleID:     *currentModel.IamRoleId,
	}

	output, res, err := client.CloudProviderSnapshotExportBuckets.Create(context.Background(), *currentModel.GroupId, input)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.completeByAtlasModel(output)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create successful",
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	output, res, err := client.CloudProviderSnapshotExportBuckets.Get(context.Background(), *currentModel.GroupId, *currentModel.Id)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.completeByAtlasModel(output)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Get successful",
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	res, err = client.CloudProviderSnapshotExportBuckets.Delete(context.Background(), *currentModel.GroupId, *currentModel.Id)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete successful",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}
	output, res, err := client.CloudProviderSnapshotExportBuckets.List(context.Background(), *currentModel.GroupId, params)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	resultList := make([]interface{}, 0)

	for i, _ := range output.Results {
		model := Model{}
		model.completeByAtlasModel(output.Results[i])
		resultList = append(resultList, model)
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List successful",
		ResourceModels:  resultList,
	}, nil
}
