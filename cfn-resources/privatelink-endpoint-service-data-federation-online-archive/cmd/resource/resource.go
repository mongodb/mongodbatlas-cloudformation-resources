// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	ctx "context"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.EndpointID}
var ReadRequiredFields = []string{constants.ProjectID, constants.EndpointID}
var DeleteRequiredFields = []string{constants.ProjectID, constants.EndpointID}
var ListRequiredFields = []string{constants.ProjectID}

const (
	AlreadyExists = "already exists"
)

func setup() {
	util.SetupLogger("mongodb-atlas-federated-query-limit")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	readModel := Model{ProjectId: currentModel.ProjectId, EndpointId: currentModel.EndpointId}
	_, err := readModel.getPrivateEndpoint(atlas)

	if err == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          AlreadyExists,
			HandlerErrorCode: string(types.HandlerErrorCodeAlreadyExists)}, nil
	}

	response, err := createOrUpdate(currentModel, atlas)

	if err != nil {
		return handleError(response, err)
	}

	// Read endpoint
	readResponse, err := currentModel.getPrivateEndpoint(atlas)

	if err != nil {
		return handleError(readResponse, err)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   currentModel}, nil
}

func createOrUpdate(currentModel *Model, client *util.MongoDBClient) (*http.Response, error) {
	provider := constants.AWS
	privateNetworkEndpointIDEntry := admin20231115014.PrivateNetworkEndpointIdEntry{
		EndpointId: *currentModel.EndpointId,
		Comment:    currentModel.Comment,
		Type:       currentModel.Type,
		Provider:   &provider,
	}
	createRequest := client.Atlas20231115014.DataFederationApi.CreateDataFederationPrivateEndpoint(
		ctx.Background(),
		*currentModel.ProjectId,
		&privateNetworkEndpointIDEntry,
	)
	_, response, err := createRequest.Execute()
	return response, err
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	response, err := currentModel.getPrivateEndpoint(atlas)

	if err != nil {
		return handleError(response, err)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
		ResourceModel:   currentModel}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	readModel := Model{ProjectId: currentModel.ProjectId, EndpointId: currentModel.EndpointId}
	readResponse, err := readModel.getPrivateEndpoint(atlas)

	if err != nil {
		return handleError(readResponse, err)
	}
	response, err := createOrUpdate(currentModel, atlas)

	if err != nil {
		return handleError(response, err)
	}

	// Read endpoint
	readResponse, err = currentModel.getPrivateEndpoint(atlas)

	if err != nil {
		return handleError(readResponse, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   currentModel}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	_, response, err := client.Atlas20231115014.DataFederationApi.DeleteDataFederationPrivateEndpoint(
		ctx.Background(),
		*currentModel.ProjectId,
		*currentModel.EndpointId,
	).Execute()

	if err != nil {
		return handleError(response, err)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
		ResourceModel:   nil}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	pe, response, err := client.Atlas20231115014.DataFederationApi.ListDataFederationPrivateEndpoints(
		ctx.Background(),
		*currentModel.ProjectId,
	).Execute()

	if err != nil {
		return handleError(response, err)
	}
	endpoints := make([]interface{}, len(pe.GetResults()))
	for i, e := range pe.GetResults() {
		eID := e.GetEndpointId()
		endpoints[i] = Model{
			ProjectId:  currentModel.ProjectId,
			Profile:    currentModel.Profile,
			Comment:    e.Comment,
			Type:       e.Type,
			EndpointId: &eID,
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Completed",
		ResourceModels:  endpoints}, nil
}

func (model *Model) getPrivateEndpoint(client *util.MongoDBClient) (*http.Response, error) {
	readRequest := client.Atlas20231115014.DataFederationApi.GetDataFederationPrivateEndpoint(
		ctx.Background(),
		*model.ProjectId,
		*model.EndpointId,
	)
	paginatedPrivateNetworkEndpointIDEntry, response, err := readRequest.Execute()
	if err != nil {
		return response, err
	}
	model.readPrivateEndpoint(paginatedPrivateNetworkEndpointIDEntry)
	return response, err
}

func (model *Model) readPrivateEndpoint(pe *admin20231115014.PrivateNetworkEndpointIdEntry) *Model {
	if pe == nil {
		return model
	}
	model.Comment = pe.Comment
	model.Type = pe.Type
	model.EndpointId = &pe.EndpointId
	return model
}

func handleError(response *http.Response, err error) (handler.ProgressEvent, error) {
	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: string(types.HandlerErrorCodeAlreadyExists)}, nil
	}
	return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error during execution : %s", err.Error()), response), nil
}
