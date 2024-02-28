package resource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ConnectionName, constants.Type}
var ReadRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ConnectionName}
var UpdateRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ConnectionName, constants.Type}
var DeleteRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ConnectionName}
var ListRequiredFields = []string{constants.ProjectID, constants.InstanceName}

// TODO - remove logs, extract chores, handle errors, add dependency for cluster/kafka conn types
func setup() {
	util.SetupLogger("mongodb-atlas-stream-connection")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	setup()
	log.Debugf("Create() currentModel:%+v", currentModel)

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, peErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if peErr != nil {
		return *peErr, nil
	}
	conn := client.AtlasSDK
	ctx := context.Background()

	//handler logic main
	projectID := currentModel.ProjectId
	instanceName := currentModel.InstanceName
	streamConnectionReq := newStreamConnectionReq(ctx, currentModel)

	streamConnResp, apiResp, err := conn.StreamsApi.CreateStreamConnection(ctx, *projectID, *instanceName, streamConnectionReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}

	readModel := newStreamConnectionModel(streamConnResp, projectID, instanceName, currentModel.Profile)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   readModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("Read() currentModel:%+v", currentModel)
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	client, peErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if peErr != nil {
		return *peErr, nil
	}
	conn := client.AtlasSDK

	projectID := currentModel.ProjectId
	instanceName := currentModel.InstanceName
	connectionName := currentModel.ConnectionName
	streamConnResp, apiResp, err := conn.StreamsApi.GetStreamConnection(context.Background(), *projectID, *instanceName, *connectionName).Execute()
	if err != nil {
		log.Debugf("Read - error: %+v", err)
		return handleError(apiResp, constants.READ, err)
	}

	readModel := newStreamConnectionModel(streamConnResp, projectID, instanceName, currentModel.Profile)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   readModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("Update() currentModel:%+v", currentModel)
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	client, peErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if peErr != nil {
		return *peErr, nil
	}
	conn := client.AtlasSDK
	ctx := context.Background()

	projectID := currentModel.ProjectId
	instanceName := currentModel.InstanceName
	connectionName := currentModel.ConnectionName
	streamConnectionReq := newStreamConnectionReq(ctx, currentModel)
	streamConnResp, apiResp, err := conn.StreamsApi.UpdateStreamConnection(ctx, *projectID, *instanceName, *connectionName, streamConnectionReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.UPDATE, err)
	}

	readModel := newStreamConnectionModel(streamConnResp, projectID, instanceName, currentModel.Profile)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   readModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("Delete() currentModel:%+v", currentModel)
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, peErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if peErr != nil {
		return *peErr, nil
	}
	conn := client.AtlasSDK
	ctx := context.Background()

	projectID := currentModel.ProjectId
	instanceName := currentModel.InstanceName
	connectionName := currentModel.ConnectionName
	if _, apiResp, err := conn.StreamsApi.DeleteStreamConnection(ctx, *projectID, *instanceName, *connectionName).Execute(); err != nil {
		return handleError(apiResp, constants.DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
		ResourceModel:   nil}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("List() currentModel:%+v", currentModel)

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	client, peErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if peErr != nil {
		return *peErr, nil
	}
	conn := client.AtlasSDK
	ctx := context.Background()

	projectID := currentModel.ProjectId
	instanceName := currentModel.InstanceName

	pageNum := 1
	accumulatedStreamConns := make([]admin.StreamsConnection, 0)
	for ok := true; ok; {
		streamConns, apiResp, err := conn.StreamsApi.ListStreamConnectionsWithParams(ctx, &admin.ListStreamConnectionsApiParams{
			GroupId:      *projectID,
			TenantName:   *instanceName,
			ItemsPerPage: util.Pointer(100),
			PageNum:      util.Pointer(pageNum),
		}).Execute()

		if err != nil {
			return handleError(apiResp, constants.LIST, err)
		}
		accumulatedStreamConns = append(accumulatedStreamConns, streamConns.GetResults()...)
		ok = streamConns.GetTotalCount() > len(accumulatedStreamConns)
		pageNum++
	}

	response := make([]interface{}, 0)
	for _, streamConn := range accumulatedStreamConns {
		model := newStreamConnectionModel(&streamConn, projectID, instanceName, currentModel.Profile)
		response = append(response, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModels:  response,
	}, nil
}

func newStreamConnectionModel(streamsConn *admin.StreamsConnection, projectID, instanceName, profile *string) *Model {
	if streamsConn == nil {
		return nil
	}

	model := &Model{
		ProjectId:        projectID,
		InstanceName:     instanceName,
		Profile:          profile,
		ConnectionName:   streamsConn.Name,
		Type:             streamsConn.Type,
		ClusterName:      streamsConn.ClusterName,
		BootstrapServers: streamsConn.BootstrapServers,
	}

	if streamsConn.DbRoleToExecute != nil {
		model.DbRoleToExecute = &DBRoleToExecute{
			Role: streamsConn.DbRoleToExecute.Role,
			Type: streamsConn.DbRoleToExecute.Type,
		}
	}

	if streamsConn.Authentication != nil {
		model.Authentication = &StreamsKafkaAuthentication{
			Mechanism: streamsConn.Authentication.Mechanism,
			Username:  streamsConn.Authentication.Username,
			Password:  streamsConn.Authentication.Password,
		}
	}

	if streamsConn.Security != nil {
		model.Security = &StreamsKafkaSecurity{
			BrokerPublicCertificate: streamsConn.Security.BrokerPublicCertificate,
			Protocol:                streamsConn.Security.Protocol,
		}
	}

	if streamsConn.Config != nil {
		model.Config = *streamsConn.Config
	}
	return model
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())
	_, _ = log.Warn(errMsg)
	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}

	if response.StatusCode == http.StatusUnauthorized {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Not found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	if response.StatusCode == http.StatusBadRequest {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}

func newStreamConnectionReq(ctx context.Context, model *Model) *admin.StreamsConnection {
	streamConnection := admin.StreamsConnection{
		Name:             model.ConnectionName,
		Type:             model.Type,
		ClusterName:      model.ClusterName,
		BootstrapServers: model.BootstrapServers,
	}
	if model.Authentication != nil {
		authenticationModel := model.Authentication
		streamConnection.Authentication = &admin.StreamsKafkaAuthentication{
			Mechanism: authenticationModel.Mechanism,
			Password:  authenticationModel.Password,
			Username:  authenticationModel.Username,
		}
	}
	if model.Security != nil {
		securityModel := model.Security
		streamConnection.Security = &admin.StreamsKafkaSecurity{
			BrokerPublicCertificate: securityModel.BrokerPublicCertificate,
			Protocol:                securityModel.Protocol,
		}
	}

	if model.Config != nil {
		streamConnection.Config = &model.Config
	}

	return &streamConnection
}