package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	matlasClient "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

const (
	defaultProviderName = "AWS"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := currentModel.ProjectId.Value()
	providerName := currentModel.ProviderName.Value()
	containerRequest := &matlasClient.Container{}

	if projectID == nil || *projectID == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network container: `project_id` must be set")
	}
	if providerName == nil || *providerName == "" {
		aws := defaultProviderName
		providerName = &aws
	}
	regionName := currentModel.RegionName.Value()
	if regionName == nil || *regionName == "" {
		return handler.ProgressEvent{}, fmt.Errorf("`error creating network container: region_name` must be set")
	}
	containerRequest.RegionName = *regionName
	containerRequest.ProviderName = *providerName
	CIDR := currentModel.AtlasCidrBlock.Value()
	if CIDR == nil || *CIDR == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network container: `atlasCidrBlock` must be set")
	}
	containerRequest.AtlasCIDRBlock = *CIDR
	containerResponse, _, err := client.Containers.Create(context.Background(), *projectID, containerRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network container: %s", err)
	}

	currentModel.Id = encoding.NewString(containerResponse.ID)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId.Value()
	containerID := *currentModel.Id.Value()

	containerResponse, _, err := client.Containers.Get(context.Background(), projectID, containerID)

	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading container with id(project: %s, container: %s): %s", projectID, containerID, err)
	}

	currentModel.RegionName = encoding.NewString(containerResponse.RegionName)
	currentModel.Provisioned = encoding.NewBool(*containerResponse.Provisioned)
	currentModel.VpcId = encoding.NewString(containerResponse.VPCID)
	currentModel.AtlasCidrBlock = encoding.NewString(containerResponse.AtlasCIDRBlock)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectId := *currentModel.ProjectId.Value()
	containerId := *currentModel.Id.Value()
	containerRequest := &matlasClient.Container{}
	providerName := currentModel.ProviderName.Value()
	if providerName == nil || *providerName == "" {
		aws := defaultProviderName
		providerName = &aws
	}
	CIDR := currentModel.AtlasCidrBlock.Value()
	if CIDR != nil {
		containerRequest.AtlasCIDRBlock = *CIDR
	}
	containerRequest.ProviderName = *providerName
	containerRequest.RegionName = *currentModel.RegionName.Value()
	containerResponse, _, err := client.Containers.Update(context.Background(), projectId, containerId, containerRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error updating container with id(project: %s, container: %s): %s", projectId, containerRequest, err)
	}

	currentModel.Id = encoding.NewString(containerResponse.ID)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectId := *currentModel.ProjectId.Value()
	containerId := *currentModel.Id.Value()

	_, err = client.Containers.Delete(context.Background(), projectId, containerId)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting container with id(project: %s, container: %s): %s", projectId, containerId, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectId := *currentModel.ProjectId.Value()
	containerRequest := &matlasClient.ContainersListOptions{
		ProviderName: *currentModel.ProviderName.Value(),
		ListOptions:  matlasClient.ListOptions{},
	}
	containerResponse, _, err := client.Containers.List(context.Background(), projectId, containerRequest)
	var models []Model
	for _, container := range containerResponse {
		var model Model
		model.RegionName = encoding.NewString(container.RegionName)
		model.Provisioned = encoding.NewBool(*container.Provisioned)
		model.VpcId = encoding.NewString(container.VPCID)
		model.AtlasCidrBlock = encoding.NewString(container.AtlasCIDRBlock)

		models = append(models, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   models,
	}, nil
}
