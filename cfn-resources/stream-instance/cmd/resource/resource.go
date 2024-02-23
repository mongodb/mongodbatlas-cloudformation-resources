package resource

import (
	"context"
	"errors"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func setup() {
	util.SetupLogger("mongodb-atlas-stream-instance")
}

var CreateRequiredFields = []string{constants.Name, constants.StreamConfig, constants.StreamConfig}

const Kafka = "Kafka"
const Cluster = "Cluster"

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	streamInstanceCreateReq := newStreamsTenant(currentModel)

	atlasV2 := client.AtlasSDK

	createdStreamInstance, resp, err := atlasV2.StreamsApi.CreateStreamInstance(context.Background(), *currentModel.GroupId, streamInstanceCreateReq).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.Id = createdStreamInstance.Id

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	return handler.ProgressEvent{}, errors.New("Not implemented: Read")
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	return handler.ProgressEvent{}, errors.New("Not implemented: Update")
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	return handler.ProgressEvent{}, errors.New("Not implemented: Delete")
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	return handler.ProgressEvent{}, errors.New("Not implemented: List")
}

func newStreamsTenant(model *Model) *admin.StreamsTenant {
	return &admin.StreamsTenant{
		Name:    model.Name,
		GroupId: model.GroupId,
		DataProcessRegion: &admin.StreamsDataProcessRegion{
			CloudProvider: *model.DataProcessRegion.CloudProvider,
			Region:        *model.DataProcessRegion.Region,
		},
		StreamConfig: &admin.StreamConfig{
			Tier: model.StreamConfig.Tier,
		},
	}
}
