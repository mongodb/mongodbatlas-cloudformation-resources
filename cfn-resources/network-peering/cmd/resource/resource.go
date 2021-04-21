package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/atlas/mongodbatlas"
	"os"
)

func setup() {
	util.SetupLogger("mongodb-atlas-network-peering")
}

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
	DefaultAWSCIDR             = "172.31.0.0/21"
	DefaultRouteTableCIDRBlock = "10.0.0.0/24"
)

func findContainer(projectId string, region string, currentModel *Model) (bool, *mongodbatlas.Container, error) {
	var container mongodbatlas.Container
	log.Debugf("findContainer projectId:%+v, region:%+v", projectId, region)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return false, &container, err
	}
	opt := &mongodbatlas.ContainersListOptions{ProviderName: "AWS"}
	log.Debugf("Looking for any AWS containers for this project:%s. opt:%+v", projectId, opt)
	containers, _, err := client.Containers.List(context.TODO(), projectId, opt)
	if err != nil {
		return false, &container, err
	}
	log.Debugf("found AWS containers for project:%+v", containers)
	for i := range containers {
		log.Debugf("RegionName:%s, region:%s", containers[i].RegionName, region)
		if containers[i].RegionName == region {
			log.Debugf("Found AWS container for region:%v, %v", region, containers[i])
			return true, &containers[i], nil
		}
	}
	return false, &container, nil
}
func validateOrCreateNetworkContainer(req *handler.Request, prevModel *Model, currentModel *Model) (*mongodbatlas.Container, error) {
	log.Debugf("validateOrCreateNetworkContainer prevModel:%+v, currentModel:%+v", prevModel, currentModel)
	var container mongodbatlas.Container
	if currentModel.ApiKeys == nil {
		return &container, fmt.Errorf("No ApiKeys found in currentModel:%+v", currentModel)
	}
	if currentModel.ProjectId == nil {
		return &container, fmt.Errorf("ProjectId was not set! currentModel:%+v", currentModel)
	}

	var ar string
	if currentModel.AccepterRegionName == nil { // use lambda default
		r := os.Getenv("AWS_REGION")
		log.Debugf("AccepterRegionName was nil, found AWS_REGION region:%v", r)
		ar = util.EnsureAtlasRegion(r)
	} else {
		r := *currentModel.AccepterRegionName
		log.Debugf("AccepterRegionName was SET to:%v", r)
		ar = util.EnsureAtlasRegion(r)
	}
	log.Debugf("converted to atlas region :%v", ar)

	projectId := *currentModel.ProjectId
	region := &ar
	// Check if have AWS container for this group,
	// if so return it -
	// if passed a ContainerId and it does not match, then
	// return an ERROR, explain to remove the ContainerId parameter
	found, c, err := findContainer(projectId, *region, currentModel)
	if err != nil {
		return &container, err
	}
	if found {
		return c, nil
	}
	// Didn't find one for this AWS region, need to create
	log.Debugf("projectId:%v, region:%v, cidr:%+v", projectId, region, &DefaultAWSCIDR)
	containerRequest := &mongodbatlas.Container{}
	containerRequest.RegionName = *region
	containerRequest.ProviderName = "AWS"
	containerRequest.AtlasCIDRBlock = DefaultAWSCIDR
	log.Debugf("containerRequest:%+v", containerRequest)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return &container, err
	}
	containerResponse, resp, err := client.Containers.Create(context.TODO(), *currentModel.ProjectId, containerRequest)
	// TODO add logging here
	if resp != nil && resp.StatusCode == 409 {
		log.Warnf("Container already exists, looking for it: resp:%+v", resp)
		found, c, err := findContainer(projectId, *region, currentModel)
		if err != nil {
			return c, err
		}
		if found {
			return c, nil
		}
	} else if err != nil {
		return &container, err
	}
	log.Debugf("created container response:%v", containerResponse)
	return containerResponse, nil
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Warnf("Create - error err:%+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	log.Debugf("Create - currentModel:%+v", currentModel)
	projectID := *currentModel.ProjectId
	container, err := validateOrCreateNetworkContainer(&req, prevModel, currentModel)

	if err != nil {
		log.Warnf("error network container mgmt: %v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	log.Debugf("Found valid container:%+v", container)

	peerRequest := mongodbatlas.Peer{
		ContainerID:  container.ID,
		VpcID:        *currentModel.VpcId,
		ProviderName: container.ProviderName,
	}

	region := currentModel.AccepterRegionName
	log.Debugf("Create region=%v ~~~~~~~~~~~~~~~~~~~~~~~~", *region)
	if region == nil || *region == "" {
		region = &req.RequestContext.Region
		log.Infof("AccepterRegionName was not set, default to req.RequestContext.Region:%v", region)
	}
	awsAccountId := currentModel.AwsAccountId
	if awsAccountId == nil || *awsAccountId == "" {
		awsAccountId = &req.RequestContext.AccountID
		log.Infof("AwsAccountIdwas not set, default to req.RequestContext.AccountID:%v", awsAccountId)
	}
	rtCIDR := currentModel.RouteTableCIDRBlock
	if rtCIDR == nil || *rtCIDR == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network peering: `RouteTableCIDRBlock` must be set")
	}

	peerRequest.AccepterRegionName = *region
	peerRequest.AWSAccountID = *awsAccountId
	peerRequest.RouteTableCIDRBlock = *rtCIDR
	log.Debugf("peerRequest:%+v", peerRequest)
	peerResponse, _, err := client.Peers.Create(context.Background(), projectID, &peerRequest)

	if err != nil {
		log.Warnf("error creating network peering: %s", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	log.Debugf("Create peerResponse:%+v", peerResponse)
	currentModel.Id = &peerResponse.ID

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Warnf("Delete err:%+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	log.Debugf("Read - currentModel:%+v", currentModel)
	projectID := *currentModel.ProjectId
	if currentModel.Id == nil {
		return handler.ProgressEvent{
			Message:          fmt.Sprintf("No Id found in model:%+v for Update", currentModel),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil

	}
	peerID := *currentModel.Id

	peerResponse, resp, err := client.Peers.Get(context.Background(), projectID, peerID)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Warnf("Resource Not Found 404 for READ projectId:%s, peerID:%+v, err:%+v", projectID, peerID, err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		} else {
			log.Warnf("Error READ projectId:%s, err:%+v", projectID, err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
		}
	}
	log.Debugf("Read: peerResponse:%+v", peerResponse)

	currentModel.AwsAccountId = &peerResponse.AWSAccountID
	currentModel.RouteTableCIDRBlock = &peerResponse.RouteTableCIDRBlock
	currentModel.VpcId = &peerResponse.VpcID
	currentModel.Id = &peerResponse.ID
	currentModel.ConnectionId = &peerResponse.ConnectionID
	currentModel.ErrorStateName = &peerResponse.ErrorStateName
	if currentModel.ErrorStateName != nil {
		currentModel.ErrorStateName = &peerResponse.ErrorStateName
	}
	if currentModel.StatusName != nil {
		currentModel.StatusName = &peerResponse.StatusName
	}
	currentModel.ProviderName = &peerResponse.ProviderName

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("Update currentModel:%+v", currentModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Warnf("Update - error err:%+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	projectID := *currentModel.ProjectId
	if currentModel.Id == nil {
		return handler.ProgressEvent{
			Message:          fmt.Sprintf("No Id found in model:%+v for Update", currentModel),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil

	}

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
	peerResponse, resp, err := client.Peers.Update(context.Background(), projectID, peerID, &peerRequest)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Warnf("Resource Not Found 404 for READ projectId:%s, peerID:%+v, err:%+v", projectID, peerID, err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		} else {
			log.Warnf("Error READ projectId:%s, err:%+v", projectID, err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
		}
		//return handler.ProgressEvent{}, fmt.Errorf("error updating peer with id(project: %s, peer: %s): %s", projectID, peerID, err)
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
	setup()
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Delete - error err:%+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateProgress(client, currentModel, "DELETED")
	}

	projectId := *currentModel.ProjectId
	peerId := *currentModel.Id
	resp, err := client.Peers.Delete(context.Background(), projectId, peerId)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Warnf("Resource Not Found 404 for DELETE projectId:%s, peerID:%+v, err:%+v", projectId, peerId, err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		} else {
			log.Warnf("Error DELETE projectId:%s, err:%+v", projectId, err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
		}
	}

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              "Delete Inprogess",
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 10,
		CallbackContext: map[string]interface{}{
			"stateName": "DELETING",
		},
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Warnf("List - error err:%+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	projectID := *currentModel.ProjectId
	peerResponse, resp, err := client.Peers.List(context.Background(), projectID, &mongodbatlas.ContainersListOptions{})
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Warnf("Resource Not Found 404 for READ projectId:%s, err:%+v", projectID, err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		} else {
			log.Warnf("Error READ projectId:%s, err:%+v", projectID, err)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
		}
	}

	models := []interface{}{}
	for _, peer := range peerResponse {
		var model Model
		model.AccepterRegionName = &peer.AccepterRegionName
		model.AwsAccountId = &peer.AWSAccountID
		model.RouteTableCIDRBlock = &peer.RouteTableCIDRBlock
		model.VpcId = &peer.VpcID
		model.Id = &peer.ID
		model.ConnectionId = &peer.ConnectionID
		model.ErrorStateName = &peer.ErrorStateName
		model.StatusName = &peer.StatusName
		model.ProviderName = &peer.ProviderName

		models = append(models, model)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}

func validateProgress(client *mongodbatlas.Client, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	isReady, state, err := networkPeeringIsReady(client, *currentModel.ProjectId, *currentModel.Id, targetState)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		//return handler.ProgressEvent{}, err
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
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	if state == "DELETED" {
		log.Print("Do not set ResourceModel property for DELETED resources")
	} else {
		log.Print("validateProgress isReady was true but state not DELETED?")
		p.ResourceModel = currentModel
	}
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
