package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"go.mongodb.org/atlas/mongodbatlas"
	"log"
	"os"
	//"github.com/davecgh/go-spew/spew"
)

// Helper to check container id or create one for the AWS region for
// the given project. This is patterned off the mongocli logic:
// https://github.com/mongodb/mongocli/blob/master/internal/cli/atlas/networking/peering/create/aws.go
// Expects the currentModel to have, (w/ ApiKeys):
// ProjectId,ContainerId ---> Try to lookup the container id, just check it's valid.
// or
// ProjectId,AccepterRegionName ---> new AWS container for that region, if you omit the region then will attempt to use env AWS_REGION
// and allows
// AtlasCIDRBlock  - defaults to: "172.31.0.0/21"

var (
	DefaultAWSCIDR = "172.31.0.0/21"
)

func validateOrCreateNetworkContainer(req *handler.Request, prevModel *Model, currentModel *Model) (*mongodbatlas.Container, error) {
	log.Printf("validateOrCreateNetworkContainer prevModel:%+v, currentModel:%+v", prevModel, currentModel)
	var container mongodbatlas.Container
	if currentModel.ApiKeys == nil {
		return &container, fmt.Errorf("No ApiKeys found in currentModel:%+v", currentModel)
	}
	if currentModel.ProjectId == nil {
		return &container, fmt.Errorf("ProjectId was not set! currentModel:%+v", currentModel)
	}
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return &container, err
	}

	var ar string
	if currentModel.AccepterRegionName == nil { // use lambda default
		r := os.Getenv("AWS_REGION")
		log.Printf("AccepterRegionName was nil, found AWS_REGION region:%v", r)
		ar = util.EnsureAtlasRegion(r)
	} else {
		r := *currentModel.AccepterRegionName
		log.Printf("AccepterRegionName was SET to:%v", r)
		ar = util.EnsureAtlasRegion(r)
	}
	log.Printf("converted to atlas region :%v", ar)
	CIDR := currentModel.RouteTableCIDRBlock
	if CIDR == nil {
		CIDR := &DefaultAWSCIDR
		log.Printf("CIDR was not set, default to:%v", *CIDR)
	}

	projectId := *currentModel.ProjectId
	region := &ar
	cidr := CIDR

	// Check if have AWS container for this group,
	// if so return it -
	// if passed a ContainerId and it does not match, then
	// return an ERROR, explain to remove the ContainerId parameter
	opt := &mongodbatlas.ContainersListOptions{ProviderName: "AWS"}
	log.Printf("Looking for any AWS containers for this project:%s. opt:%+v", projectId, opt)
	cr, _, err := client.Containers.List(context.TODO(), projectId, opt)
	if err != nil {
		return &container, err
	}
	log.Printf("found AWS containers for project:%+v", cr)
	// cr is a list, need filter on our region?
	for i := range cr {
		log.Printf("RegionName:%s, region:%s", cr[i].RegionName, *region)
		if cr[i].RegionName == *region {
			log.Printf("Found AWS container for region:%v, %v", region, cr[i])
			if currentModel.ContainerId != nil {
				if cr[i].ID != *currentModel.ContainerId {
					log.Printf("Error: resource has ContainerId set to %v, however there is already an AWS Network Container for this Atlas Project:%+v. Remove the ContainerId property from your template and rety.", *currentModel.ContainerId, cr[i])
				}
			}
			return &cr[i], nil
		}
	}
	// Didn't find one for this AWS region, need to create
	log.Printf("projectId:%v, region:%v, cidr:%+v", projectId, region, cidr)
	containerRequest := &mongodbatlas.Container{}
	containerRequest.RegionName = *region
	containerRequest.ProviderName = "AWS"
	containerRequest.AtlasCIDRBlock = *cidr
	log.Printf("containerRequest:%+v", containerRequest)
	containerResponse, _, err := client.Containers.Create(context.TODO(), *currentModel.ProjectId, containerRequest)
	if err != nil {
		return &container, err
	}
	log.Printf("created container response:%v", containerResponse)
	return containerResponse, nil
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	//log.Printf("Create req:%+v, prevModel:%s, currentModel:%s",req,spew.Sdump(prevModel),spew.Sdump(currentModel))

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId
	container, err := validateOrCreateNetworkContainer(&req, prevModel, currentModel)

	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error network container mgmt: %v", err)
	}
	log.Printf("Found valid container:%+v", container)

	peerRequest := mongodbatlas.Peer{
		ContainerID:  container.ID,
		VpcID:        *currentModel.VpcId,
		ProviderName: container.ProviderName,
	}

	region := currentModel.AccepterRegionName
	log.Printf("Create region=%v ~~~~~~~~~~~~~~~~~~~~~~~~", *region)
	if region == nil || *region == "" {
		region = &req.RequestContext.Region
		log.Printf("AccepterRegionName was not set, default to req.RequestContext.Region:%v", region)
	}
	awsAccountId := currentModel.AwsAccountId
	if awsAccountId == nil || *awsAccountId == "" {
		awsAccountId = &req.RequestContext.AccountID
		log.Printf("AwsAccountIdwas not set, default to req.RequestContext.AccountID:%v", awsAccountId)
	}
	rtCIDR := currentModel.RouteTableCIDRBlock
	if rtCIDR == nil || *rtCIDR == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: `RouteTableCIDRBlock` must be set")
	}

	peerRequest.AccepterRegionName = *region
	peerRequest.AWSAccountID = *awsAccountId
	peerRequest.RouteTableCIDRBlock = *rtCIDR
	log.Printf("peerRequest:%+v", peerRequest)
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
	currentModel.RouteTableCIDRBlock = &peerResponse.RouteTableCIDRBlock
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
	rtTableBlock := currentModel.RouteTableCIDRBlock
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
		model.RouteTableCIDRBlock = &peer.RouteTableCIDRBlock
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
			"stateName": state,
		}
		return p, nil
	}

	p := handler.NewProgressEvent()
	p.ResourceModel = currentModel
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	return p, nil
}

func networkPeeringIsReady(client *mongodbatlas.Client, projectId, peerId, targetState string) (bool, string, error) {
	peerResponse, resp, err := client.Peers.Get(context.Background(), projectId, peerId)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return true, "DELETED", nil
		}
	}
	return peerResponse.StatusName == targetState, peerResponse.StatusName, nil
}
