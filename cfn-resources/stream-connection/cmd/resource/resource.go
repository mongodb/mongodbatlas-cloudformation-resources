package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/atlas-sdk/v20231115007/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

const (
	ClusterConnectionType = "Cluster"
	KafkaConnectionType   = "Kafka"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ConnectionName, constants.Type}
var ReadRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ConnectionName}
var UpdateRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ConnectionName, constants.Type}
var DeleteRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ConnectionName}
var ListRequiredFields = []string{constants.ProjectID, constants.InstanceName}

// TODO - remove logs, extract chores, handle errors, add dependency for cluster/kafka conn types
func setup(cfnFunc constants.CfnFunctions, req handler.Request, currentModel *Model) (*admin.APIClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-stream-connection")

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	var requiredFields []string

	switch cfnFunc {
	case constants.CREATE:
		requiredFields = CreateRequiredFields
	case constants.READ:
		requiredFields = ReadRequiredFields
	case constants.UPDATE:
		requiredFields = UpdateRequiredFields
	case constants.DELETE:
		requiredFields = DeleteRequiredFields
	case constants.LIST:
		requiredFields = ListRequiredFields
	}

	if errEvent := validator.ValidateModel(requiredFields, currentModel); errEvent != nil {
		return nil, errEvent
	}

	client, peErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	if peErr != nil {
		return nil, peErr
	}
	return client.AtlasSDK, nil

}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := setup(constants.CREATE, req, currentModel)
	if peErr != nil {
		return *peErr, nil
	}
	// setup()
	// log.Debugf("...............1. Create() currentModel:%+v", currentModel)

	// util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	// if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
	// 	return *errEvent, nil
	// }
	// log.Debugf("\n.............2. Create() currentModel:%+v", currentModel)

	// client, peErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	// if peErr != nil {
	// 	return *peErr, nil
	// }
	// conn := client.AtlasSDK
	ctx := context.Background()

	log.Debugf("\n3. Create() currentModel:%+v", currentModel)

	//handler logic main
	projectID := currentModel.ProjectId
	instanceName := currentModel.InstanceName
	streamConnectionReq := NewStreamConnectionReq(ctx, currentModel)

	jsonBytes, err := json.MarshalIndent(streamConnectionReq, "", "\t")
	log.Debugf("\n4. Create() streamConnectionReq: %s", string(jsonBytes))

	streamConnResp, apiResp, err := conn.StreamsApi.CreateStreamConnection(ctx, *projectID, *instanceName, streamConnectionReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.CREATE, err)
	}

	jsonBytes, err = json.MarshalIndent(streamConnResp, "", "\t")
	log.Debugf("\nstreamConnResp from API:  %s", string(jsonBytes))
	log.Debugf("\nstreamConnResp from API:  %+v", streamConnResp)

	// readModel := newStreamConnectionModel(streamConnResp, projectID, instanceName, currentModel.Profile)
	resourceModel := GetStreamConnectionModel(streamConnResp, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   resourceModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := setup(constants.READ, req, currentModel)
	if peErr != nil {
		return *peErr, nil
	}
	// setup()
	// log.Debugf("Read() currentModel:%+v", currentModel)
	// util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	// if errEvent := validator.ValidateModel(ReadRequiredFields, currentModel); errEvent != nil {
	// 	return *errEvent, nil
	// }

	// // Create atlas client
	// client, peErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	// if peErr != nil {
	// 	return *peErr, nil
	// }
	// conn := client.AtlasSDK

	projectID := currentModel.ProjectId
	instanceName := currentModel.InstanceName
	connectionName := currentModel.ConnectionName
	streamConnResp, apiResp, err := conn.StreamsApi.GetStreamConnection(context.Background(), *projectID, *instanceName, *connectionName).Execute()
	if err != nil {
		log.Debugf("Read - error: %+v", err)
		return handleError(apiResp, constants.READ, err)
	}

	// readModel := newStreamConnectionModel(streamConnResp, projectID, instanceName, currentModel.Profile)
	resourceModel := GetStreamConnectionModel(streamConnResp, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   resourceModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := setup(constants.UPDATE, req, currentModel)
	if peErr != nil {
		return *peErr, nil
	}
	// setup()
	// log.Debugf("Update() currentModel:%+v", currentModel)
	// util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	// if errEvent := validator.ValidateModel(UpdateRequiredFields, currentModel); errEvent != nil {
	// 	return *errEvent, nil
	// }

	// // Create atlas client
	// client, peErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	// if peErr != nil {
	// 	return *peErr, nil
	// }
	// conn := client.AtlasSDK
	ctx := context.Background()

	projectID := currentModel.ProjectId
	instanceName := currentModel.InstanceName
	connectionName := currentModel.ConnectionName
	streamConnectionReq := NewStreamConnectionReq(ctx, currentModel)
	streamConnResp, apiResp, err := conn.StreamsApi.UpdateStreamConnection(ctx, *projectID, *instanceName, *connectionName, streamConnectionReq).Execute()
	if err != nil {
		return handleError(apiResp, constants.UPDATE, err)
	}

	// readModel := newStreamConnectionModel(streamConnResp, projectID, instanceName, currentModel.Profile)
	resourceModel := GetStreamConnectionModel(streamConnResp, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   resourceModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	conn, peErr := setup(constants.DELETE, req, currentModel)
	if peErr != nil {
		return *peErr, nil
	}
	// setup()
	// log.Debugf("...................1. Delete() currentModel:%+v", currentModel)
	// util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	// if errEvent := validator.ValidateModel(DeleteRequiredFields, currentModel); errEvent != nil {
	// 	return *errEvent, nil
	// }

	// client, peErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	// if peErr != nil {
	// 	return *peErr, nil
	// }
	// conn := client.AtlasSDK
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
	conn, peErr := setup(constants.LIST, req, currentModel)
	if peErr != nil {
		return *peErr, nil
	}
	// setup()
	// log.Debugf("List() currentModel:%+v", currentModel)

	// util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	// if errEvent := validator.ValidateModel(ListRequiredFields, currentModel); errEvent != nil {
	// 	return *errEvent, nil
	// }

	// // Create atlas client
	// client, peErr := util.NewAtlasV2OnlyClientLatest(&req, currentModel.Profile, true)
	// if peErr != nil {
	// 	return *peErr, nil
	// }
	// conn := client.AtlasSDK
	ctx := context.Background()

	projectID := currentModel.ProjectId
	instanceName := currentModel.InstanceName

	// pageNum := 1
	// accumulatedStreamConns := make([]admin.StreamsConnection, 0)
	// for ok := true; ok; {
	// 	streamConns, apiResp, err := conn.StreamsApi.ListStreamConnectionsWithParams(ctx, &admin.ListStreamConnectionsApiParams{
	// 		GroupId:      *projectID,
	// 		TenantName:   *instanceName,
	// 		ItemsPerPage: util.Pointer(100),
	// 		PageNum:      util.Pointer(pageNum),
	// 	}).Execute()

	// 	if err != nil {
	// 		return handleError(apiResp, constants.LIST, err)
	// 	}
	// 	accumulatedStreamConns = append(accumulatedStreamConns, streamConns.GetResults()...)
	// 	ok = streamConns.GetTotalCount() > len(accumulatedStreamConns)
	// 	pageNum++
	// }
	accumulatedStreamConns, apiResp, err := getAllStreamConnections(ctx, conn, *projectID, *instanceName)
	if err != nil {
		return handleError(apiResp, constants.LIST, err)
	}

	response := make([]interface{}, 0)
	for i := range accumulatedStreamConns {
		model := GetStreamConnectionModel(&accumulatedStreamConns[i], nil)
		model.ProjectId = currentModel.ProjectId
		model.InstanceName = currentModel.InstanceName
		model.Profile = currentModel.Profile

		response = append(response, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModels:  response,
	}, nil
}

func getAllStreamConnections(ctx context.Context, conn *admin.APIClient, projectID, instanceName string) ([]admin.StreamsConnection, *http.Response, error) {
	pageNum := 1
	accumulatedStreamConns := make([]admin.StreamsConnection, 0)
	for ok := true; ok; {
		streamConns, apiResp, err := conn.StreamsApi.ListStreamConnectionsWithParams(ctx, &admin.ListStreamConnectionsApiParams{
			GroupId:      projectID,
			TenantName:   instanceName,
			ItemsPerPage: util.Pointer(100),
			PageNum:      util.Pointer(pageNum),
		}).Execute()

		if err != nil {
			return nil, apiResp, err
		}
		accumulatedStreamConns = append(accumulatedStreamConns, streamConns.GetResults()...)
		ok = streamConns.GetTotalCount() > len(accumulatedStreamConns)
		pageNum++
	}
	return accumulatedStreamConns, nil, nil
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())

	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}
