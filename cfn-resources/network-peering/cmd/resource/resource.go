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
	"errors"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20240805001/admin"
)

func setup() {
	util.SetupLogger("mongodb-atlas-network-peering")
}

const (
	StatusPendingAcceptance string = "PENDING_ACCEPTANCE"
	StatusFailed            string = "FAILED"
	StatusAvailable         string = "AVAILABLE"
	StatusDeleted           string = "DELETED"
	StatusInitiating        string = "INITIATING"
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

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
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
	}

	peerRequest := admin.BaseNetworkPeeringConnectionSettings{
		ContainerId:         *currentModel.ContainerId,
		VpcId:               currentModel.VpcId,
		AccepterRegionName:  currentModel.AccepterRegionName,
		AwsAccountId:        awsAccountID,
		RouteTableCidrBlock: currentModel.RouteTableCIDRBlock,
		ProviderName:        admin.PtrString(constants.AWS),
	}

	peerResponse, resp, err := client.Atlas20231115002.NetworkPeeringApi.CreatePeeringConnection(context.Background(), projectID, &peerRequest).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(),
			resp), nil
	}

	currentModel.Id = peerResponse.Id
	return progressevent.GetInProgressProgressEvent("Creating",
		map[string]interface{}{
			"stateName": StatusInitiating,
			"id":        &peerResponse.Id,
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

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	projectID := *currentModel.ProjectId
	peerID := *currentModel.Id

	peerResponse, resp, err := client.Atlas20231115002.NetworkPeeringApi.GetPeeringConnection(context.Background(), projectID, peerID).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(),
			resp), nil
	}

	currentModel.AwsAccountId = peerResponse.AwsAccountId
	currentModel.RouteTableCIDRBlock = peerResponse.RouteTableCidrBlock
	currentModel.VpcId = peerResponse.VpcId
	currentModel.Id = peerResponse.Id
	currentModel.ConnectionId = peerResponse.ConnectionId
	currentModel.ErrorStateName = peerResponse.ErrorStateName
	if currentModel.ErrorStateName != nil {
		currentModel.ErrorStateName = peerResponse.ErrorStateName
	}
	if currentModel.StatusName != nil {
		currentModel.StatusName = peerResponse.StatusName
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
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
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
	peerRequest := admin.BaseNetworkPeeringConnectionSettings{}

	if region := currentModel.AccepterRegionName; region != nil {
		peerRequest.AccepterRegionName = region
	}

	if accountID := currentModel.AwsAccountId; accountID != nil {
		peerRequest.AwsAccountId = accountID
	}

	peerRequest.ProviderName = admin.PtrString(constants.AWS)
	if rtTableBlock := currentModel.RouteTableCIDRBlock; rtTableBlock != nil {
		peerRequest.RouteTableCidrBlock = rtTableBlock
	}

	if vpcID := currentModel.VpcId; vpcID != nil {
		peerRequest.VpcId = vpcID
	}

	peerRequest.ContainerId = *currentModel.ContainerId
	peerResponse, resp, err := client.Atlas20231115002.NetworkPeeringApi.UpdatePeeringConnection(context.Background(), projectID, peerID, &peerRequest).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.Id = peerResponse.Id
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

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateDeletionProcess(client, currentModel), nil
	}

	projectID := *currentModel.ProjectId
	peerID := *currentModel.Id
	_, resp, err := client.Atlas20231115002.NetworkPeeringApi.DeletePeeringConnection(context.Background(), projectID, peerID).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(),
			resp), nil
	}

	return progressevent.GetInProgressProgressEvent("Deleting",
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

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	projectID := *currentModel.ProjectId
	peerResponse, resp, err := client.Atlas20231115002.NetworkPeeringApi.ListPeeringConnections(context.Background(), projectID).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	models := make([]interface{}, 0)
	networkPeeringConnections := peerResponse.Results
	for i := range networkPeeringConnections {
		var model Model
		model.AccepterRegionName = networkPeeringConnections[i].AccepterRegionName
		model.AwsAccountId = networkPeeringConnections[i].AwsAccountId
		model.RouteTableCIDRBlock = networkPeeringConnections[i].RouteTableCidrBlock
		model.VpcId = networkPeeringConnections[i].VpcId
		model.Id = networkPeeringConnections[i].Id
		model.ConnectionId = networkPeeringConnections[i].ConnectionId
		model.ErrorStateName = networkPeeringConnections[i].ErrorStateName
		model.StatusName = networkPeeringConnections[i].StatusName

		models = append(models, model)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}

func validateDeletionProcess(client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	state, err := getStatus(client, *currentModel.ProjectId, *currentModel.Id)
	if err != nil {
		return progressevent.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest)
	}

	if state == StatusDeleted {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Complete",
		}
	}

	return progressevent.GetInProgressProgressEvent("Deleting",
		map[string]interface{}{
			"stateName": state,
		},
		currentModel,
		5,
	)
}

func validateCreationProcess(client *util.MongoDBClient, currentModel *Model) handler.ProgressEvent {
	state, err := getStatus(client, *currentModel.ProjectId, *currentModel.Id)
	if err != nil {
		return progressevent.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest)
	}

	if state == StatusPendingAcceptance || state == StatusAvailable {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Complete",
			ResourceModel:   currentModel,
		}
	}

	if state == StatusFailed {
		return progressevent.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInternalFailure)
	}

	return progressevent.GetInProgressProgressEvent("Creating",
		map[string]interface{}{
			"stateName": state,
			"id":        &currentModel.Id,
		},
		currentModel,
		5,
	)
}

func getStatus(client *util.MongoDBClient, projectID, peerID string) (statusName string, err error) {
	peerResponse, _, err := client.Atlas20231115002.NetworkPeeringApi.GetPeeringConnection(context.Background(), projectID, peerID).Execute()
	if err != nil {
		if apiError, ok := admin.AsError(err); ok && *apiError.Error == http.StatusNotFound {
			return StatusDeleted, nil
		}

		return "", err
	}

	if util.IsStringPresent(peerResponse.ErrorStateName) {
		err = errors.New(*peerResponse.ErrorStateName)
	}

	if util.IsStringPresent(peerResponse.StatusName) {
		statusName = *peerResponse.StatusName
	}

	return
}
