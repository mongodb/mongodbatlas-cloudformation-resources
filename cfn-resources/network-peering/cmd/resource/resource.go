package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"go.mongodb.org/atlas/mongodbatlas"
	"os"
    "github.com/aws/aws-sdk-go/service/cloudformation"
    log "github.com/sirupsen/logrus"
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
    DefaultRouteTableCIDRBlock = "10.0.0.0/24"
)

func init() {
    util.InitLogger()
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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return &container, err
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
	opt := &mongodbatlas.ContainersListOptions{ProviderName: "AWS"}
	log.Debugf("Looking for any AWS containers for this project:%s. opt:%+v", projectId, opt)
	cr, _, err := client.Containers.List(context.TODO(), projectId, opt)
	if err != nil {
		return &container, err
	}
	log.Debugf("found AWS containers for project:%+v", cr)
	// cr is a list, need filter on our region?
	for i := range cr {
		log.Debugf("RegionName:%s, region:%s", cr[i].RegionName, *region)
		if cr[i].RegionName == *region {
			log.Debugf("Found AWS container for region:%v, %v", region, cr[i])
			if currentModel.ContainerId != nil {
				if cr[i].ID != *currentModel.ContainerId {
					log.Debugf("Error: resource has ContainerId set to %v, however there is already an AWS Network Container for this Atlas Project:%+v. Remove the ContainerId property from your template and rety.", *currentModel.ContainerId, cr[i])
				}
			}
			return &cr[i], nil
		}
	}
	// Didn't find one for this AWS region, need to create
	log.Debugf("projectId:%v, region:%v, cidr:%+v", projectId, region, &DefaultAWSCIDR)
	containerRequest := &mongodbatlas.Container{}
	containerRequest.RegionName = *region
	containerRequest.ProviderName = "AWS"
	containerRequest.AtlasCIDRBlock = DefaultAWSCIDR
	log.Debugf("containerRequest:%+v", containerRequest)
	containerResponse, _, err := client.Containers.Create(context.TODO(), *currentModel.ProjectId, containerRequest)
	if err != nil {
		return &container, err
	}
	log.Debugf("created container response:%v", containerResponse)
	return containerResponse, nil
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
        log.WithFields(log.Fields{"err":err,}).Error("Create - error")
		return handler.ProgressEvent{
            OperationStatus: handler.Failed,
            Message: err.Error(),
            HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

    log.WithFields(log.Fields{"currentModel":currentModel,}).Debug("Create")
	projectID := *currentModel.ProjectId
	container, err := validateOrCreateNetworkContainer(&req, prevModel, currentModel)

	if err != nil {
		log.Errorf("error network container mgmt: %v", err)
		return handler.ProgressEvent{
            OperationStatus: handler.Failed,
            Message: err.Error(),
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
		log.Debugf("AccepterRegionName was not set, default to req.RequestContext.Region:%v", region)
	}
	awsAccountId := currentModel.AwsAccountId
	if awsAccountId == nil || *awsAccountId == "" {
		awsAccountId = &req.RequestContext.AccountID
		log.Debugf("AwsAccountIdwas not set, default to req.RequestContext.AccountID:%v", awsAccountId)
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
		log.Errorf("error creating network peering: %s", err)
		return handler.ProgressEvent{
            OperationStatus: handler.Failed,
            Message: err.Error(),
            HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

    log.WithFields(log.Fields{"peerRequest":peerRequest,}).Debug("CREATE ---> peerResponse")
	currentModel.Id = &peerResponse.ID
    log.WithFields(log.Fields{"currentModel":currentModel,}).Debug("CREATE ---> currentModel")

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
        log.WithFields(log.Fields{"err":err,}).Error("Read - error")
		return handler.ProgressEvent{
            OperationStatus: handler.Failed,
            Message: err.Error(),
            HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

    log.WithFields(log.Fields{"currentModel":currentModel,}).Debug("Read")
	projectID := *currentModel.ProjectId
    if currentModel.Id == nil {
        return handler.ProgressEvent{
            Message: fmt.Sprintf("No Id found in model:%+v for Update", currentModel),
            OperationStatus: handler.Failed,
            HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil

    }
	peerID := *currentModel.Id

	peerResponse, resp, err := client.Peers.Get(context.Background(), projectID, peerID)
	if err != nil {
        if resp != nil && resp.StatusCode == 404 {
            log.Warnf("Resource Not Found 404 for READ projectId:%s, peerID:%+v, err:%+v", projectID, peerID, err)
            return handler.ProgressEvent{
                Message: err.Error(),
                OperationStatus: handler.Failed,
                HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
        } else {
            log.Errorf("Error READ projectId:%s, err:%+v", projectID, err)
            return handler.ProgressEvent{
                Message: err.Error(),
                OperationStatus: handler.Failed,
                HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
        }
	}
    log.WithFields(log.Fields{"peerResponse":peerResponse,}).Debug(" READ--> peerResponse")

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
    log.WithFields(log.Fields{"currentModel":currentModel,}).Debug("Update")
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
        log.WithFields(log.Fields{"err":err,}).Error("Update - error")
		return handler.ProgressEvent{
            OperationStatus: handler.Failed,
            Message: err.Error(),
            HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	projectID := *currentModel.ProjectId
    if currentModel.Id == nil {
        return handler.ProgressEvent{
            Message: fmt.Sprintf("No Id found in model:%+v for Update", currentModel),
            OperationStatus: handler.Failed,
            HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil

    }

	peerID := *currentModel.Id
	peerRequest := mongodbatlas.Peer{}
    log.WithFields(log.Fields{"projectID":projectID,"peerID":peerID}).Debug("peer1")

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
    log.WithFields(log.Fields{"peerRequest":peerRequest,}).Debug("peer2")
	peerResponse, resp, err := client.Peers.Update(context.Background(), projectID, peerID, &peerRequest)
	if err != nil {
        if resp != nil && resp.StatusCode == 404 {
            log.Warnf("Resource Not Found 404 for READ projectId:%s, peerID:%+v, err:%+v", projectID, peerID, err)
            return handler.ProgressEvent{
                Message: err.Error(),
                OperationStatus: handler.Failed,
                HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
        } else {
            log.Errorf("Error READ projectId:%s, err:%+v", projectID, err)
            return handler.ProgressEvent{
                Message: err.Error(),
                OperationStatus: handler.Failed,
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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
        log.WithFields(log.Fields{"err":err,}).Error("Delete - error")
		return handler.ProgressEvent{
            OperationStatus: handler.Failed,
            Message: err.Error(),
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
                Message: err.Error(),
                OperationStatus: handler.Failed,
                HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
        } else {
            log.Errorf("Error DELETE projectId:%s, err:%+v", projectId, err)
            return handler.ProgressEvent{
                Message: err.Error(),
                OperationStatus: handler.Failed,
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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
        log.WithFields(log.Fields{"err":err,}).Error("List - error")
		return handler.ProgressEvent{
            OperationStatus: handler.Failed,
            Message: err.Error(),
            HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	projectID := *currentModel.ProjectId
	peerResponse, resp, err := client.Peers.List(context.Background(), projectID, &mongodbatlas.ContainersListOptions{})
	if err != nil {
        if resp != nil && resp.StatusCode == 404 {
            log.Warnf("Resource Not Found 404 for READ projectId:%s, err:%+v", projectID, err)
            return handler.ProgressEvent{
                Message: err.Error(),
                OperationStatus: handler.Failed,
                HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
        } else {
            log.Errorf("Error READ projectId:%s, err:%+v", projectID, err)
            return handler.ProgressEvent{
                Message: err.Error(),
                OperationStatus: handler.Failed,
                HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
        }
	}

    models := []interface{} {}
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
		ResourceModels:   models,
	}, nil
}

func validateProgress(client *mongodbatlas.Client, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	isReady, state, err := networkPeeringIsReady(client, *currentModel.ProjectId, *currentModel.Id, targetState)
	if err != nil {
        log.WithFields(log.Fields{"err":err,}).Error("validateProgress - error")
		return handler.ProgressEvent{
            OperationStatus: handler.Failed,
            Message: err.Error(),
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
    if state=="DELETED" {
        log.Debugf("Do not set ResourceModel property for DELETED resources")
    } else {
        log.Warningf("validateProgress isReady was true but state not DELETED?")
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
