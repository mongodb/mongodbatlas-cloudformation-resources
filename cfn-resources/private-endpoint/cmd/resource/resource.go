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
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
)

const (
	providerName = "AWS"
)

func setup() {
	util.SetupLogger("mongodb-atlas-private-endpoint")
}

var CreateRequiredFields = []string{constants.ProjectID}
var ReadRequiredFields = []string{constants.GroupID, constants.ID, constants.Region}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.GroupID, constants.ID}
var ListRequiredFields = []string{constants.GroupID}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	endpointRequest := admin.CreateEndpointRequest{
		Id: currentModel.InterfaceEndpointId,
	}

	privateEndpointRequest := client.AtlasV2.PrivateEndpointServicesApi.CreatePrivateEndpoint(context.Background(), *currentModel.ProjectId,
		*currentModel.CloudProvider, *currentModel.EndpointServiceId, &endpointRequest)
	privateEndpoint, response, err := privateEndpointRequest.Execute()
	defer response.Body.Close()
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("error creating Serverless Private Endpoint %s",
				err.Error()), response),
			nil
	}

	currentModel.EndpointId = privateEndpoint.InterfaceEndpointId
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create completed",
		ResourceModel:   currentModel}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Get successful",
		ResourceModel:   currentModel}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	setup()

	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	privateEndpointRequest := client.AtlasV2.PrivateEndpointServicesApi.DeletePrivateEndpoint(context.Background(), *currentModel.ProjectId,
		*currentModel.CloudProvider, *currentModel.EndpointId, *currentModel.EndpointServiceId)
	_, response, err := privateEndpointRequest.Execute()
	defer response.Body.Close()
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("error creating Serverless Private Endpoint %s",
				err.Error()), response),
			nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete completed"}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{
		OperationStatus: handler.Failed,
		Message:         "List successful"}, nil
}
