package resource

import (
	"context"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.ProjID, constants.PvtKey, constants.PubKey, constants.Name}
var ReadRequiredFields = []string{constants.ProjID, constants.Name, constants.PvtKey, constants.PubKey}
var UpdateRequiredFields = []string{constants.PvtKey, constants.PubKey}
var DeleteRequiredFields = []string{constants.ProjID, constants.Name, constants.PvtKey, constants.PubKey}
var ListRequiredFields = []string{constants.PvtKey, constants.PubKey}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-ServerlessInstance")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Create() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Warnf("Create - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	serverlessInstanceRequest := &mongodbatlas.ServerlessCreateRequestParams{
		Name:                         *currentModel.Name,
		ProviderSettings:             setProviderSettings(currentModel),
		ServerlessBackupOptions:      setBackupOptions(currentModel),
		TerminationProtectionEnabled: currentModel.TerminationProtectionEnabled,
	}

	_, res, err := client.ServerlessInstances.Create(context.Background(), *currentModel.ProjectID, serverlessInstanceRequest)
	if err != nil {
		_, _ = log.Warnf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Response
	return Read(req, prevModel, currentModel)
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Read() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Warnf("Read - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	cluster, res, err := client.ServerlessInstances.Get(context.Background(), *currentModel.ProjectID, *currentModel.Name)
	if err != nil {
		_, _ = log.Warnf("Read - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	// Read Instance
	model := readServerlessInstance(cluster)
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   model,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Update() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Warnf("Update - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	serverlessInstanceRequest := &mongodbatlas.ServerlessUpdateRequestParams{
		ServerlessBackupOptions:      setBackupOptions(currentModel),
		TerminationProtectionEnabled: currentModel.TerminationProtectionEnabled,
	}

	_, res, err := client.ServerlessInstances.Update(context.Background(), *currentModel.ProjectID, *currentModel.Name, serverlessInstanceRequest)
	if err != nil {
		_, _ = log.Warnf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	// Response
	return Read(req, prevModel, currentModel)
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Delete() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Warnf("Delete - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	res, err := client.ServerlessInstances.Delete(context.Background(), *currentModel.ProjectID, *currentModel.Name)
	if err != nil {
		_, _ = log.Warnf("Delete - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   nil,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("List() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Warnf("List - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	listOptions := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 1000,
	}
	clustersResp, res, err := client.ServerlessInstances.List(context.Background(), *currentModel.ProjectID, listOptions)
	if err != nil {
		_, _ = log.Warnf("List - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	instances := []interface{}{}
	for i := range clustersResp.Results {
		var cluster = &Model{}
		cluster = readServerlessInstance(clustersResp.Results[i])
		instances = append(instances, cluster)
	}
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   instances,
	}, nil
}

func setBackupOptions(currentModel *Model) *mongodbatlas.ServerlessBackupOptions {
	var serverlessBackupOptions = &mongodbatlas.ServerlessBackupOptions{}
	if currentModel.ContinuousBackupEnabled != nil {
		serverlessBackupOptions = &mongodbatlas.ServerlessBackupOptions{
			ServerlessContinuousBackupEnabled: currentModel.ContinuousBackupEnabled,
		}
	}
	return serverlessBackupOptions
}

func setProviderSettings(currentModel *Model) *mongodbatlas.ServerlessProviderSettings {
	var serverlessProviderSettings = &mongodbatlas.ServerlessProviderSettings{}
	if currentModel.ProviderSettings != nil {
		if currentModel.ProviderSettings.BackingProviderName != nil {
			serverlessProviderSettings.BackingProviderName = *currentModel.ProviderSettings.BackingProviderName
		}
		if currentModel.ProviderSettings.ProviderName != nil {
			serverlessProviderSettings.ProviderName = *currentModel.ProviderSettings.ProviderName
		}
		if currentModel.ProviderSettings.RegionName != nil {
			serverlessProviderSettings.RegionName = *currentModel.ProviderSettings.RegionName
		}
	}
	return serverlessProviderSettings
}

func readServerlessInstance(cluster *mongodbatlas.Cluster) (model *Model) {
	var serverless = &Model{}
	serverless.Name = &cluster.Name

	if cluster.ProviderSettings != nil {
		serverless.ProviderSettings = &ServerlessInstanceProviderSettings{
			BackingProviderName: &cluster.ProviderSettings.BackingProviderName,
			ProviderName:        &cluster.ProviderSettings.ProviderName,
			RegionName:          &cluster.ProviderSettings.RegionName,
		}
	}

	if cluster.ServerlessBackupOptions != nil {
		serverless.ContinuousBackupEnabled = cluster.ServerlessBackupOptions.ServerlessContinuousBackupEnabled
	}

	if cluster.ConnectionStrings != nil {
		serverless.ConnectionStrings = &ServerlessInstanceConnectionStrings{
			StandardSrv:     &cluster.ConnectionStrings.StandardSrv,
			PrivateEndpoint: readPrivateEndpoint(cluster.ConnectionStrings.PrivateEndpoint),
		}
	}
	serverless.CreateDate = &cluster.CreateDate
	serverless.MongoDBVersion = &cluster.MongoDBVersion
	serverless.Links = readLinks(cluster.Links)
	serverless.TerminationProtectionEnabled = cluster.TerminationProtectionEnabled
	serverless.StateName = &cluster.StateName
	return serverless
}

func readPrivateEndpoint(privateEPs []mongodbatlas.PrivateEndpoint) (endPoints []ServerlessInstancePrivateEndpoint) {
	for i := range privateEPs {
		var pep = ServerlessInstancePrivateEndpoint{}
		pep.Endpoints = readPrivateEndpointEndpoints(privateEPs[i].Endpoints)
		pep.Type = &privateEPs[i].Type
		pep.SrvConnectionString = &privateEPs[i].SRVConnectionString
		endPoints = append(endPoints, pep)
	}
	return
}

func readPrivateEndpointEndpoints(peEndpoints []mongodbatlas.Endpoint) (epEndpoints []ServerlessInstancePrivateEndpointEndpoint) {
	for i := range peEndpoints {
		epEndpoints = append(epEndpoints, ServerlessInstancePrivateEndpointEndpoint{
			EndpointId:   &peEndpoints[i].EndpointID,
			ProviderName: &peEndpoints[i].ProviderName,
			Region:       &peEndpoints[i].Region,
		})
	}
	return
}

func readLinks(clsLinks []*mongodbatlas.Link) (links []Link) {
	for i := range clsLinks {
		links = append(links, Link{
			Href: &clsLinks[i].Href,
			Rel:  &clsLinks[i].Rel,
		})
	}
	return
}
