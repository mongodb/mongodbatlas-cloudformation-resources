package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	matlasClient "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)


// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId.Value()
	providerName := *currentModel.ProviderName.Value()
	containerRequest := &matlasClient.Container{}

	if providerName == "AWS"{
		regionName := *currentModel.RegionName.Value()
		if regionName == ""{
			return handler.ProgressEvent{}, fmt.Errorf("`region_name` must be set when `provider_name` is AWS")
		}
		containerRequest.RegionName = regionName
	}

	if providerName == "AZURE"{
		region := *currentModel.RegionName.Value()
		if region == ""{
			return handler.ProgressEvent{}, fmt.Errorf("`region` must be set when `provider_name` is AZURE")
		}
		containerRequest.Region = region
	}

	containerResponse, _, err := client.Containers.Create(context.Background(), projectID, containerRequest)

	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network container: %s", err)
	}

	currentModel.Id = encoding.NewString(containerResponse.ID)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message: "Create complete",
		ResourceModel: currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId.Value()
	containerID := *currentModel.ContainerId.Value()

	containerResponse, _, err := client.Containers.Get(context.Background(), projectID, containerID)

	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading container with id(project: %s, container: %s): %s", projectID, containerID, err)
	}

	currentModel.RegionName = encoding.NewString(containerResponse.RegionName)
	currentModel.Region = encoding.NewString(containerResponse.Region)
	currentModel.AzureSubscriptionId = encoding.NewString(containerResponse.AzureSubscriptionID)
	currentModel.Provisioned = encoding.NewBool(*containerResponse.Provisioned)
	currentModel.GcpProjectId = encoding.NewString(containerResponse.GCPProjectID)
	currentModel.NetworkName = encoding.NewString(containerResponse.NetworkName)
	currentModel.VpcId = encoding.NewString(containerResponse.VPCID)
	currentModel.VnetName = encoding.NewString(containerResponse.VNetName)
	currentModel.ContainerId = encoding.NewString(containerResponse.ID)
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

	projectID := *currentModel.ProjectId.Value()
	containerID := *currentModel.ContainerId.Value()
	containerRequest := &matlasClient.Container{}

	if prevModel.AtlasCidrBlock.Value() != currentModel.AtlasCidrBlock.Value(){
		containerRequest.AtlasCIDRBlock = *currentModel.AtlasCidrBlock.Value()
		containerRequest.ProviderName = *currentModel.ProviderName.Value()
	}

	if prevModel.ProviderName.Value() != currentModel.ProviderName.Value(){
		containerRequest.ProviderName = *currentModel.ProviderName.Value()
	}

	if prevModel.RegionName.Value() != currentModel.RegionName.Value(){
		containerRequest.RegionName = *currentModel.RegionName.Value()
	}

	if prevModel.Region.Value() != currentModel.Region.Value(){
		containerRequest.Region = *currentModel.Region.Value()
	}

	containerResponse, _, _ := client.Containers.Update(context.Background(), projectID, containerID, containerRequest)

	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error updating container with id(project: %s, container: %s): %s", projectID, containerRequest, err)
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

	projectID := *currentModel.ProjectId.Value()
	containerID := *currentModel.ContainerId.Value()


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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId.Value()
	containerRequest := &matlasClient.ListOptions{
		PageNum: int(*currentModel.PageNum.Value()),
		ItemsPerPage: int(*currentModel.ItemsPerPage.Value()),
	}

	containerResponse, _, err := client.Containers.List(context.Background(), projectID, containerRequest)
	var models []Model

	for _, container := range containerResponse{
		var model Model
		model.RegionName = encoding.NewString(container.RegionName)
		model.Region = encoding.NewString(container.Region)
		model.AzureSubscriptionId = encoding.NewString(container.AzureSubscriptionID)
		model.Provisioned = encoding.NewBool(*container.Provisioned)
		model.GcpProjectId = encoding.NewString(container.GCPProjectID)
		model.NetworkName = encoding.NewString(container.NetworkName)
		model.VpcId = encoding.NewString(container.VPCID)
		model.VnetName = encoding.NewString(container.VNetName)
		model.ContainerId = encoding.NewString(container.ID)
		model.AtlasCidrBlock = encoding.NewString(container.AtlasCIDRBlock)

	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   models,
	}, nil
}
