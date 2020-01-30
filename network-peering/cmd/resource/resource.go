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

	defaultProviderName := "AWS"
	projectID := *currentModel.ProjectId.Value()
	peerRequest := matlasClient.Peer{
		ContainerID: *currentModel.ContainerId.Value(),
	}

	region := currentModel.AccepterRegionName.Value()
	if region == nil || *region == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: `accepter_region_name` must be set")
	}
	awsAccountId := currentModel.AwsAccountId.Value()
	if awsAccountId == nil || *awsAccountId == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: `aws_account_id` must be set")
	}
	rtCIDR := currentModel.RouteTableCidrBlock.Value()
	if rtCIDR == nil || *rtCIDR == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: `route_table_cidr_block` must be set")
	}
	vpcID := currentModel.VpcId.Value()
	if vpcID == nil || *vpcID == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: `vpc_id` must be set")
	}
	providerName := currentModel.ProviderName.Value()
	if providerName == nil || *providerName == "" {
		providerName = &defaultProviderName
	}
	peerRequest.AccepterRegionName = *region
	peerRequest.AWSAccountId = *awsAccountId
	peerRequest.RouteTableCIDRBlock = *rtCIDR
	peerRequest.VpcID = *vpcID
	peerRequest.ProviderName = *providerName

	peerResponse, _, err := client.Peers.Create(context.Background(), projectID, &peerRequest)

	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: %s", err)
	}

	currentModel.Id = encoding.NewString(peerResponse.ID)

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
	peerID := *currentModel.Id.Value()

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
	currentModel.StatusName = encoding.NewString(peerResponse.StatusName)
	currentModel.ProviderName = encoding.NewString(peerResponse.ProviderName)

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
	peerID := *currentModel.Id.Value()
	peerRequest := matlasClient.Peer{}

	region := currentModel.AccepterRegionName.Value()
	if region != nil {
		peerRequest.AccepterRegionName = *region
	}
	accountID := currentModel.AwsAccountId.Value()
	if accountID != nil {
		peerRequest.AWSAccountId = *accountID
	}
	peerRequest.ProviderName = "AWS"
	rtTableBlock := currentModel.RouteTableCidrBlock.Value()
	if rtTableBlock != nil {
		peerRequest.RouteTableCIDRBlock = *rtTableBlock
	}
	vpcId := currentModel.VpcId.Value()
	if vpcId != nil {
		peerRequest.VpcID = *vpcId
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

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateProgress(client, currentModel, "DELETED")
	}

	projectId := *currentModel.ProjectId.Value()
	peerId := *currentModel.Id.Value()
	_, err = client.Peers.Delete(context.Background(), projectId, peerId)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting peer with id(project: %s, peer: %s): %s", projectId, peerId, err)
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              "Delete Complete",
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 10,
		CallbackContext: map[string]interface{}{
			"stateName": "DELETING",
		},
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId.Value()
	peerResponse, _, err := client.Peers.List(context.Background(), projectID, &matlasClient.ListOptions{})
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
		model.StatusName = encoding.NewString(peer.StatusName)
		model.ProviderName = encoding.NewString(peer.ProviderName)

		models = append(models, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   models,
	}, nil
}

func validateProgress(client *matlasClient.Client, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	isReady, state, err := networkPeeringIsReady(client, *currentModel.ProjectId.Value(), *currentModel.Id.Value(), targetState)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	if !isReady {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = handler.InProgress
		p.CallbackDelaySeconds = 15
		p.Message = "Pending"
		p.CallbackContext = map[string]interface{}{
			"stateName" : state,
		}
		return p, nil
	}

	p := handler.NewProgressEvent()
	p.ResourceModel = currentModel
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	return p, nil
}

func networkPeeringIsReady(client *matlasClient.Client, projectId, peerId, targetState string)(bool, string, error){
	peerResponse, resp, err := client.Peers.Get(context.Background(), projectId, peerId)
	if err != nil {
		if resp != nil && resp.StatusCode == 404{
			return true, "DELETED", nil
		}
	}
	return peerResponse.StatusName == targetState, peerResponse.StatusName, nil
}