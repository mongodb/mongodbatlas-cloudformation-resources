package resource

import (
	"context"
	"errors"

	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	local_constants "github.com/mongodb/mongodbatlas-cloudformation-resources/cloud-backup-snapshot-export-bucket/cmd/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.GroupID, local_constants.BucketName, local_constants.IamRoleID, constants.PvtKey, constants.PubKey}
var ReadRequiredFields = []string{constants.GroupID, constants.ID, constants.PvtKey, constants.PubKey}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.GroupID, constants.ID, constants.PvtKey, constants.PubKey}
var ListRequiredFields = []string{constants.GroupID, constants.PvtKey, constants.PubKey}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-CloudBackupSnapshotExportBucket")
}

func (m *Model) completeByAtlasModel(bucket *mongodbatlas.CloudProviderSnapshotExportBucket) {
	m.BucketName = &bucket.BucketName
	m.IamRoleID = &bucket.IAMRoleID
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
		IAMRoleID:     *currentModel.IamRoleID,
	}

	output, res, err := client.CloudProviderSnapshotExportBuckets.Create(context.Background(), *currentModel.GroupId, input)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.Id = &output.ID

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

	_, _ = log.Warnf("Create cluster model : %+v", currentModel)

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

	output, res, err := client.CloudProviderSnapshotExportBuckets.List(context.Background(), *currentModel.GroupId, nil)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	resultList := make([]interface{}, 0)

	for i := range output.Results {
		model := Model{}
		model.completeByAtlasModel(output.Results[i])
		model.GroupId = currentModel.GroupId
		resultList = append(resultList, model)
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List successful",
		ResourceModels:  resultList,
	}, nil
}
