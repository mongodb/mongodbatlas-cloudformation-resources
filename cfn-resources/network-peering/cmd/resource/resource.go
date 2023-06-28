// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

func setup() {
	util.SetupLogger("mongodb-atlas-network-peering")
}

const (
	StatusPendingAcceptance string = "PENDING_ACCEPTANCE"
	StatusFailed                   = "FAILED"
	StatusAvailable                = "AVAILABLE"
	StatusDeleted                  = "DELETED"
	StatusInitiating               = "INITIATING"
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
	DefaultAWSCIDR             = "172.31.0.0/21"
	DefaultRouteTableCIDRBlock = "10.0.0.0/24"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.ContainerID,
	constants.AccepterRegionName, constants.AwsAccountID, constants.RouteTableCIDRBlock, constants.VPCID}

var ReadRequiredFields = []string{constants.ProjectID, constants.ID}
var UpdateRequiredFields = []string{constants.ProjectID, constants.AccepterRegionName,
	constants.AwsAccountID, constants.RouteTableCIDRBlock}

var DeleteRequiredFields = []string{constants.ProjectID, constants.ID}
var ListRequiredFields = []string{constants.ProjectID}

// validateModel to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if _, ok := req.CallbackContext["stateName"]; ok {
		currentModel.Id = aws.String(req.CallbackContext["id"].(string))
		return validateCreationProcess(client, currentModel), nil
	}

	projectID := *currentModel.ProjectId
	awsAccountID := currentModel.AwsAccountId
	if awsAccountID == nil || *awsAccountID == "" {
		awsAccountID = &req.RequestContext.AccountID
		_, _ = logger.Debugf("AwsAccountIdwas not set, default to req.RequestContext.AccountID:%v", awsAccountID)
	}

	peerRequest := mongodbatlas.Peer{
		ContainerID:         *currentModel.ContainerId,
		VpcID:               *currentModel.VpcId,
		AccepterRegionName:  *currentModel.AccepterRegionName,
		AWSAccountID:        *awsAccountID,
		RouteTableCIDRBlock: *currentModel.RouteTableCIDRBlock,
		ProviderName:        constants.AWS,
	}

	peerResponse, resp, err := client.Peers.Create(context.Background(), projectID, &peerRequest)
	if err != nil {
		_, _ = logger.Warnf("error creating network peering: %s", err)
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	currentModel.Id = &peerResponse.ID

	return progressevents.GetInProgressProgressEvent("Creating",
		map[string]interface{}{
			"stateName": StatusInitiating,
			"id":        &peerResponse.ID,
		},
		currentModel,
		5,
	), nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	projectID := *currentModel.ProjectId
	peerID := *currentModel.Id

	peerResponse, resp, err := client.Peers.Get(context.Background(), projectID, peerID)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

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

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	_, _ = logger.Debugf("Update currentModel:%+v", currentModel)

	if currentModel.Profile == nil {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	projectID := *currentModel.ProjectId
	if currentModel.Id == nil || *currentModel.Profile == "" {
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

	peerRequest.ProviderName = constants.AWS
	rtTableBlock := currentModel.RouteTableCIDRBlock
	if rtTableBlock != nil {
		peerRequest.RouteTableCIDRBlock = *rtTableBlock
	}
	vpcID := currentModel.VpcId
	if vpcID != nil {
		peerRequest.VpcID = *vpcID
	}
	peerResponse, resp, err := client.Peers.Update(context.Background(), projectID, peerID, &peerRequest)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error updating resource : %s", err.Error()),
			resp.Response), nil
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
	if errEvent := validateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateDeletionProcess(client, currentModel), nil
	}

	projectID := *currentModel.ProjectId
	peerID := *currentModel.Id
	resp, err := client.Peers.Delete(context.Background(), projectID, peerID)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	return progressevents.GetInProgressProgressEvent("Deleting",
		map[string]interface{}{
			"stateName": StatusDeleted,
		},
		currentModel,
		5,
	), nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	projectID := *currentModel.ProjectId
	peerResponse, resp, err := client.Peers.List(context.Background(), projectID, &mongodbatlas.ContainersListOptions{})
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	models := make([]interface{}, 0)
	for i := range peerResponse {
		var model Model
		model.AccepterRegionName = &peerResponse[i].AccepterRegionName
		model.AwsAccountId = &peerResponse[i].AWSAccountID
		model.RouteTableCIDRBlock = &peerResponse[i].RouteTableCIDRBlock
		model.VpcId = &peerResponse[i].VpcID
		model.Id = &peerResponse[i].ID
		model.ConnectionId = &peerResponse[i].ConnectionID
		model.ErrorStateName = &peerResponse[i].ErrorStateName
		model.StatusName = &peerResponse[i].StatusName

		models = append(models, model)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}

func validateDeletionProcess(client *mongodbatlas.Client, currentModel *Model) handler.ProgressEvent {
	state, _, err := getStatus(client, *currentModel.ProjectId, *currentModel.Id)
	if err != nil {
		return progressevents.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest)
	}

	if state == StatusDeleted {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Complete",
		}
	}

	return progressevents.GetInProgressProgressEvent("Deleting",
		map[string]interface{}{
			"stateName": state,
		},
		currentModel,
		5,
	)
}

func validateCreationProcess(client *mongodbatlas.Client, currentModel *Model) handler.ProgressEvent {
	state, errorMessage, err := getStatus(client, *currentModel.ProjectId, *currentModel.Id)
	if err != nil {
		return progressevents.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest)
	}

	if state == StatusPendingAcceptance || state == StatusAvailable {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Complete",
			ResourceModel:   currentModel,
		}
	}

	if state == StatusFailed {
		return progressevents.GetFailedEventByCode(errorMessage, cloudformation.HandlerErrorCodeInternalFailure)
	}

	return progressevents.GetInProgressProgressEvent("Creating",
		map[string]interface{}{
			"stateName": state,
		},
		currentModel,
		5,
	)
}

func getStatus(client *mongodbatlas.Client, projectID, peerID string) (string, string, error) {
	peerResponse, resp, err := client.Peers.Get(context.Background(), projectID, peerID)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return StatusDeleted, "", nil
		} else {
			return "", "", err
		}
	}

	return peerResponse.StatusName, peerResponse.ErrorStateName, nil
}
