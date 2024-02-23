package resource

import (
	"context"
	"errors"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"

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

	atlasV2 := client.AtlasSDK

	streamInstance, resp, err := atlasV2.StreamsApi.GetStreamInstance(context.Background(), *currentModel.GroupId, *currentModel.Name).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
	}

	currentModel.Id = streamInstance.Id
	currentModel.Hostnames = streamInstance.GetHostnames()
	currentModel.DataProcessRegion = newModelDataRegion(streamInstance.DataProcessRegion)
	currentModel.StreamConfig = newModelStreamConfig(streamInstance.StreamConfig)
	currentModel.Connections = newModelConnections(streamInstance.GetConnections())

	return handler.ProgressEvent{
		OperationStatus: cloudformation.OperationStatusSuccess,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
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

	updateRequest := &admin.StreamsDataProcessRegion{
		CloudProvider: *currentModel.DataProcessRegion.CloudProvider,
		Region:        *currentModel.DataProcessRegion.Region,
	}

	atlasV2 := client.AtlasSDK

	_, resp, err := atlasV2.StreamsApi.UpdateStreamInstance(context.Background(), *currentModel.GroupId, *currentModel.Name, updateRequest).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update success",
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
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

	atlasV2 := client.AtlasSDK

	_, resp, err := atlasV2.StreamsApi.DeleteStreamInstance(context.Background(), *currentModel.GroupId, *currentModel.Name).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete success",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
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

	atlasV2 := client.AtlasSDK

	accumulatedStreamInstances := make([]admin.StreamsTenant, 0)
	for ok := true; ok; {
		streamInstances, resp, err := atlasV2.StreamsApi.ListStreamInstances(context.Background(), *currentModel.GroupId).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		}
		accumulatedStreamInstances = append(accumulatedStreamInstances, streamInstances.GetResults()...)
		ok = streamInstances.GetTotalCount() > len(accumulatedStreamInstances)
	}

	response := make([]any, 0, len(accumulatedStreamInstances))
	for i := range accumulatedStreamInstances {
		response = append(response, accumulatedStreamInstances[i])
	}

	return handler.ProgressEvent{
		OperationStatus: cloudformation.OperationStatusSuccess,
		Message:         "List Complete",
		ResourceModel:   response,
	}, nil
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

func newModelDataRegion(dataProcessRegion *admin.StreamsDataProcessRegion) *StreamsDataProcessRegion {
	return &StreamsDataProcessRegion{
		CloudProvider: &dataProcessRegion.CloudProvider,
		Region:        &dataProcessRegion.Region,
	}
}

func newModelStreamConfig(streamConfig *admin.StreamConfig) *StreamConfig {
	return &StreamConfig{
		Tier: streamConfig.Tier,
	}
}

func newModelDbRoleToExecute(dbRole *admin.DBRoleToExecute) *DBRoleToExecute {
	return &DBRoleToExecute{
		Role: dbRole.Role,
		Type: dbRole.Type,
	}
}

func newModelAuthentication(authentication *admin.StreamsKafkaAuthentication) *StreamsKafkaAuthentication {
	return &StreamsKafkaAuthentication{
		Mechanism: authentication.Mechanism,
		Password:  authentication.Password,
		Username:  authentication.Username,
	}
}

func newModelSecurity(security *admin.StreamsKafkaSecurity) *StreamsKafkaSecurity {
	return &StreamsKafkaSecurity{
		BrokerPublicCertificate: security.BrokerPublicCertificate,
		Protocol:                security.Protocol,
	}
}

func newModelConnections(streamConfig []admin.StreamsConnection) []StreamsConnection {
	if len(streamConfig) == 0 {
		return nil
	}

	connections := make([]StreamsConnection, 0)
	for _, connection := range streamConfig {
		modelConnection := StreamsConnection{
			Name: connection.Name,
			Type: connection.Type,
		}
		if connection.GetType() == Cluster {
			modelConnection.ClusterName = connection.ClusterName
			modelConnection.DbRoleToExecute = newModelDbRoleToExecute(connection.DbRoleToExecute)
		} else if connection.GetType() == Kafka {
			modelConnection.BootstrapServers = connection.BootstrapServers
			modelConnection.Authentication = newModelAuthentication(connection.Authentication)
			modelConnection.Security = newModelSecurity(connection.Security)
		}
		connections = append(connections, modelConnection)
	}
	return connections
}
