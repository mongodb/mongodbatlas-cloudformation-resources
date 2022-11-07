package resource

import (
	"context"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	defaultProviderName = "AWS"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Not Supported",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Not Supported",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Not Supported",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	_, _ = logger.Debugf("Delete currentModel:%+v", currentModel)
	projectID := *currentModel.ProjectId
	containerID := *currentModel.Id

	_, err = client.Containers.Delete(context.Background(), projectID, containerID)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting container with id(project: %s, container: %s): %s", projectID, containerID, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	_, _ = logger.Debugf("List currentModel:%+v", currentModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId
	providerName := currentModel.ProviderName
	if providerName == nil || *providerName == "" {
		aws := defaultProviderName
		providerName = &aws
	}
	_, _ = logger.Debugf("projectId:%v", projectID)
	_, _ = logger.Debugf("providerName:%v", providerName)
	containerRequest := &mongodbatlas.ContainersListOptions{
		ProviderName: *providerName,
		ListOptions:  mongodbatlas.ListOptions{},
	}
	_, _ = logger.Debugf("List projectId:%v, containerRequest:%v", projectID, containerRequest)
	containerResponse, _, err := client.Containers.List(context.TODO(), projectID, containerRequest)
	if err != nil {
		_, _ = logger.Warnf("Error %v", err)
		return handler.ProgressEvent{}, err
	}
	_, _ = logger.Debugf("containerResponse:%v", containerResponse)

	var models []Model
	for ind := range containerResponse {
		var model Model
		model.RegionName = &containerResponse[ind].RegionName
		model.Provisioned = containerResponse[ind].Provisioned
		model.Id = &containerResponse[ind].ID
		model.VpcId = &containerResponse[ind].VPCID
		model.AtlasCIDRBlock = &containerResponse[ind].AtlasCIDRBlock

		models = append(models, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   models,
	}, nil
}
