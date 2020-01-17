package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"
	matlasClient "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
)


// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId.Value()
	providerName := *currentModel.ProviderName.Value()
	peerRequest := matlasClient.Peer{}

	if providerName == "AWS"{
		region := *currentModel.AccepterRegionName.Value()
		if region == ""{
			return handler.ProgressEvent{}, fmt.Errorf("`accepter_region_name` must be set when `provider_name` is `AWS`")
		}
		awsAccountId := *currentModel.AwsAccountId.Value()
		if awsAccountId == ""{
			return handler.ProgressEvent{}, fmt.Errorf("`aws_account_id` must be set when `provider_name` is `AWS`")
		}
		rtCIDR := *currentModel.RouteTableCidrBlock.Value()
		if rtCIDR == ""{
			return handler.ProgressEvent{}, fmt.Errorf("``route_table_cidr_block` must be set when `provider_name` is `AWS`")
		}
		vpcID := *currentModel.VpcId.Value()
		if vpcID == ""{
			return handler.ProgressEvent{}, fmt.Errorf("`vpc_id` must be set when `provider_name` is `AWS`")
		}
		peerRequest.AccepterRegionName = region
		peerRequest.AWSAccountId = awsAccountId
		peerRequest.RouteTableCIDRBlock = rtCIDR
		peerRequest.VpcID = vpcID
	}

	if providerName == "GCPP"{
		gcpProjectID := *currentModel.GcpProjectId.Value()
		if gcpProjectID == ""{
			return handler.ProgressEvent{}, fmt.Errorf("`gcp_project_id` must be set when `provider_name` is `GCP`")
		}
		networkName := *currentModel.NetworkName.Value()
		if networkName == ""{
			return handler.ProgressEvent{}, fmt.Errorf("`network_name` must be set when `provider_name` is `GCP`")
		}

		peerRequest.GCPProjectID = gcpProjectID
		peerRequest.NetworkName = networkName
	}

	if providerName == "AZURE"{
		atlasCidrBlock := *currentModel.AtlasCidrBlock.Value()
		if atlasCidrBlock == ""{
			return handler.ProgressEvent{}, fmt.Errorf("`atlas_cidr_block` must be set when `provider_name` is `AZURE`")
		}
		azureDirectoryID := *currentModel.AzureDirectoryId.Value()
		if azureDirectoryID == ""{
			return handler.ProgressEvent{}, fmt.Errorf("`azure_directory_id` must be set when `provider_name` is `AZURE`")
		}
		azureSubscriptionID := *currentModel.AzureSubscriptionId.Value()
		if azureSubscriptionID == ""{
			return handler.ProgressEvent{}, fmt.Errorf("``azure_subscription_id` must be set when `provider_name` is `AZURE`")
		}
		resourceGroupName := *currentModel.ResourceGroupName.Value()
		if resourceGroupName == ""{
			return handler.ProgressEvent{}, fmt.Errorf("`resource_group_name` must be set when `provider_name` is `AZURE`")
		}
		vnetName := *currentModel.VnetName.Value()
		if vnetName == ""{
			return handler.ProgressEvent{}, fmt.Errorf("`vnet_name` must be set when `provider_name` is `AZURE`")
		}
		peerRequest.AtlasCIDRBlock = atlasCidrBlock
		peerRequest.AzureDirectoryID = azureDirectoryID
		peerRequest.AzureSubscriptionId = azureSubscriptionID
		peerRequest.ResourceGroupName = resourceGroupName
		peerRequest.VNetName = vnetName
	}

	peerResponse, _, err := client.Peers.Create(context.Background(),projectID, &peerRequest)

	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: %s", err)
	}

	currentModel.Id = encoding.NewString(peerResponse.ID)

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
	peerID := *currentModel.PeerId.Value()

	peerResponse, _, err := client.Peers.Get(context.Background(), projectID, peerID)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading peer with id(project: %s, peer: %s): %s", projectID, peerID, err)
	}

	currentModel.AccepterRegionName = encoding.NewString(peerResponse.AccepterRegionName)
	currentModel.AwsAccountId = encoding.NewString(peerResponse.AWSAccountId)
	currentModel.RouteTableCidrBlock = encoding.NewString(peerResponse.RouteTableCIDRBlock)
	currentModel.VpcId = encoding.NewString(peerResponse.VpcID)
	currentModel.ConnectionId = encoding.NewString(peerResponse.ConnectionID)
	currentModel.ErrorStateName = encoding.NewString(peerResponse.ErrorStateName)
	currentModel.AtlasId = encoding.NewString(peerResponse.AtlasCIDRBlock)
	currentModel.StatusName = encoding.NewString(peerResponse.StatusName)
	currentModel.AzureDirectoryId = encoding.NewString(peerResponse.AzureDirectoryID)
	currentModel.AzureSubscriptionId = encoding.NewString(peerResponse.AzureSubscriptionId)
	currentModel.ResourceGroupName = encoding.NewString(peerResponse.ResourceGroupName)
	currentModel.VnetName = encoding.NewString(peerResponse.VNetName)
	currentModel.ErrorState = encoding.NewString(peerResponse.ErrorState)
	currentModel.Status = encoding.NewString(peerResponse.Status)
	currentModel.GcpProjectId = encoding.NewString(peerResponse.GCPProjectID)
	currentModel.NetworkName = encoding.NewString(peerResponse.NetworkName)
	currentModel.ErrorMessage = encoding.NewString(peerResponse.ErrorMessage)
	currentModel.Id = encoding.NewString(peerResponse.ID)
	currentModel.ProviderName = encoding.NewString(peerResponse.ProviderName)
	currentModel.AtlasCidrBlock = encoding.NewString(peerResponse.AtlasCIDRBlock)


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
	peerID := *currentModel.PeerId.Value()
	peerRequest := matlasClient.Peer{}

	if prevModel.AccepterRegionName.Value() != currentModel.AccepterRegionName.Value(){
		peerRequest.AccepterRegionName = *currentModel.AccepterRegionName.Value()
	}
	if prevModel.AwsAccountId.Value() != currentModel.AwsAccountId.Value(){
		peerRequest.AWSAccountId = *currentModel.AwsAccountId.Value()
	}
	if prevModel.ProviderName.Value() != currentModel.ProviderName.Value(){
		peerRequest.ProviderName = *currentModel.ProviderName.Value()
	}
	if prevModel.RouteTableCidrBlock.Value() != currentModel.RouteTableCidrBlock.Value(){
		peerRequest.RouteTableCIDRBlock = *currentModel.RouteTableCidrBlock.Value()
	}
	if prevModel.VpcId.Value() != currentModel.VpcId.Value(){
		peerRequest.VpcID = *currentModel.VpcId.Value()
	}
	if prevModel.AzureDirectoryId.Value() != currentModel.AzureDirectoryId.Value(){
		peerRequest.AzureDirectoryID = *currentModel.AzureDirectoryId.Value()
	}
	if prevModel.AzureSubscriptionId.Value() != currentModel.AzureSubscriptionId.Value(){
		peerRequest.AzureSubscriptionId = *currentModel.AzureSubscriptionId.Value()
	}
	if prevModel.ResourceGroupName.Value() != currentModel.ResourceGroupName.Value(){
		peerRequest.ResourceGroupName = *currentModel.ResourceGroupName.Value()
	}
	if prevModel.VnetName.Value() != currentModel.VnetName.Value(){
		peerRequest.VNetName = *currentModel.VnetName.Value()
	}
	if prevModel.GcpProjectId.Value() != currentModel.GcpProjectId.Value(){
		peerRequest.GCPProjectID = *currentModel.GcpProjectId.Value()
	}
	if prevModel.NetworkName.Value() != currentModel.NetworkName.Value(){
		peerRequest.NetworkName = *currentModel.NetworkName.Value()
	}

	peerResponse, _, err := client.Peers.Update(context.Background(), projectID, peerID, &peerRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error updating peer with id(project: %s, peer: %s): %s", projectID, peerID, err)
	}

	currentModel.Id = encoding.NewString(peerResponse.ID)


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
	peerID := *currentModel.PeerId.Value()


	_, err = client.Peers.Delete(context.Background(), projectID, peerID)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting peer with id(project: %s, peer: %s): %s", projectID, peerID, err)
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
	peerRequest := matlasClient.ListOptions{
		PageNum: int(*currentModel.PageNum.Value()),
		ItemsPerPage: int(*currentModel.ItemsPerPage.Value()),
	}

	peerResponse, _, err := client.Peers.List(context.Background(), projectID, &peerRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading pf list peer with id(project: %s): %s", projectID, err)
	}

	var models []Model
	for _, peer := range peerResponse {
		var model Model
		model.AccepterRegionName = encoding.NewString(peer.AccepterRegionName)
		model.AwsAccountId = encoding.NewString(peer.AWSAccountId)
		model.RouteTableCidrBlock = encoding.NewString(peer.RouteTableCIDRBlock)
		model.VpcId = encoding.NewString(peer.VpcID)
		model.ConnectionId = encoding.NewString(peer.ConnectionID)
		model.ErrorStateName = encoding.NewString(peer.ErrorStateName)
		model.AtlasId = encoding.NewString(peer.AtlasCIDRBlock)
		model.StatusName = encoding.NewString(peer.StatusName)
		model.AzureDirectoryId = encoding.NewString(peer.AzureDirectoryID)
		model.AzureSubscriptionId = encoding.NewString(peer.AzureSubscriptionId)
		model.ResourceGroupName = encoding.NewString(peer.ResourceGroupName)
		model.VnetName = encoding.NewString(peer.VNetName)
		model.ErrorState = encoding.NewString(peer.ErrorState)
		model.Status = encoding.NewString(peer.Status)
		model.GcpProjectId = encoding.NewString(peer.GCPProjectID)
		model.NetworkName = encoding.NewString(peer.NetworkName)
		model.ErrorMessage = encoding.NewString(peer.ErrorMessage)
		model.Id = encoding.NewString(peer.ID)
		model.ProviderName = encoding.NewString(peer.ProviderName)
		model.AtlasCidrBlock = encoding.NewString(peer.AtlasCIDRBlock)

		models = append(models, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   models,
	}, nil
}
