package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
    "go.mongodb.org/atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	defaultProviderName := "AWS"
	projectID := *currentModel.ProjectId
	peerRequest := mongodbatlas.Peer{
		ContainerID: *currentModel.ContainerId,
	}

	region := currentModel.AccepterRegionName
	if region == nil || *region == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: `accepter_region_name` must be set")
	}
	awsAccountId := currentModel.AwsAccountId
	if awsAccountId == nil || *awsAccountId == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: `aws_account_id` must be set")
	}
	rtCIDR := currentModel.RouteTableCidrBlock
	if rtCIDR == nil || *rtCIDR == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: `route_table_cidr_block` must be set")
	}
	vpcID := currentModel.VpcId
	if vpcID == nil || *vpcID == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: `vpc_id` must be set")
	}
	providerName := currentModel.ProviderName
	if providerName == nil || *providerName == "" {
		providerName = &defaultProviderName
	}
	peerRequest.AccepterRegionName = *region
	peerRequest.AWSAccountID = *awsAccountId
	peerRequest.RouteTableCIDRBlock = *rtCIDR
	peerRequest.VpcID = *vpcID
	peerRequest.ProviderName = *providerName

	peerResponse, _, err := client.Peers.Create(context.Background(), projectID, &peerRequest)

	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: %s", err)
	}

	currentModel.Id = &peerResponse.ID

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId
	peerID := *currentModel.Id

	peerResponse, _, err := client.Peers.Get(context.Background(), projectID, peerID)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading peer with id(project: %s, peer: %s): %s", projectID, peerID, err)
	}

	currentModel.AccepterRegionName = &peerResponse.AccepterRegionName
	currentModel.AwsAccountId = &peerResponse.AWSAccountID
	currentModel.RouteTableCidrBlock = &peerResponse.RouteTableCIDRBlock
	currentModel.VpcId = &peerResponse.VpcID
	currentModel.ConnectionId = &peerResponse.ConnectionID
	currentModel.ErrorStateName = &peerResponse.ErrorStateName
	currentModel.StatusName = &peerResponse.StatusName
	currentModel.ProviderName = &peerResponse.ProviderName

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId
	peerID := *currentModel.Id
	peerRequest := mongodbatlas.Peer{}

	region := currentModel.AccepterRegionName
	if region != nil {
		peerRequest.AccepterRegionName = *region
	}
	accountID := currentModel.AwsAccountId
	if accountID != nil {
		peerRequest.AWSAccountID = *accountID
	}
	peerRequest.ProviderName = "AWS"
	rtTableBlock := currentModel.RouteTableCidrBlock
	if rtTableBlock != nil {
		peerRequest.RouteTableCIDRBlock = *rtTableBlock
	}
	vpcId := currentModel.VpcId
	if vpcId != nil {
		peerRequest.VpcID = *vpcId
	}
	peerResponse, _, err := client.Peers.Update(context.Background(), projectID, peerID, &peerRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error updating peer with id(project: %s, peer: %s): %s", projectID, peerID, err)
	}

	currentModel.Id = &peerResponse.ID

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateProgress(client, currentModel, "DELETED")
	}

	projectId := *currentModel.ProjectId
	peerId := *currentModel.Id
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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId
	peerResponse, _, err := client.Peers.List(context.Background(), projectID, &mongodbatlas.ContainersListOptions{})
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading pf list peer with id(project: %s): %s", projectID, err)
	}

	var models []Model
	for _, peer := range peerResponse {
		var model Model
		model.AccepterRegionName = &peer.AccepterRegionName
		model.AwsAccountId = &peer.AWSAccountID
		model.RouteTableCidrBlock = &peer.RouteTableCIDRBlock
		model.VpcId = &peer.VpcID
		model.ConnectionId = &peer.ConnectionID
		model.ErrorStateName = &peer.ErrorStateName
		model.StatusName = &peer.StatusName
		model.ProviderName = &peer.ProviderName

		models = append(models, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   models,
	}, nil
}

func validateProgress(client *mongodbatlas.Client, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	isReady, state, err := networkPeeringIsReady(client, *currentModel.ProjectId, *currentModel.Id, targetState)
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

func networkPeeringIsReady(client *mongodbatlas.Client, projectId, peerId, targetState string)(bool, string, error){
	peerResponse, resp, err := client.Peers.Get(context.Background(), projectId, peerId)
	if err != nil {
		if resp != nil && resp.StatusCode == 404{
			return true, "DELETED", nil
		}
	}
	return peerResponse.StatusName == targetState, peerResponse.StatusName, nil
}
