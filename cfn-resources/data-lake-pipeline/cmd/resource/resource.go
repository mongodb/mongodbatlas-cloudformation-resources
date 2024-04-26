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
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115002/admin"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.Name, constants.Sink}
var ReadRequiredFields = []string{constants.ProjectID, constants.Name}
var UpdateRequiredFields = []string{constants.ProjectID, constants.Name, constants.Sink}
var DeleteRequiredFields = []string{constants.ProjectID, constants.Name}
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

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	groupID := *currentModel.ProjectId
	dataLakeIntegrationPipeline := generateDataLakeIntegrationPipeline(currentModel)

	pe, response, err := client.Atlas20231115002.DataLakePipelinesApi.CreatePipeline(ctx.Background(), groupID, dataLakeIntegrationPipeline).Execute()

	if err != nil {
		if response.StatusCode == http.StatusBadRequest {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
		}
		return handleError(response, err)
	}

	model := ReadResponseModelGeneration(pe)
	if model == nil {
		errorMsg := "Response model from the API is empty or nil "
		return progress_events.GetFailedEventByResponse(errorMsg, response), nil
	}
	model.Profile = currentModel.Profile
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   model}, nil
}

func generateDataLakeIntegrationPipeline(currentModel *Model) *admin.DataLakeIngestionPipeline {
	dataLakeIntegrationPipeline := admin.DataLakeIngestionPipeline{
		GroupId: currentModel.ProjectId,
		Name:    currentModel.Name,
		Sink: &admin.IngestionSink{
			MetadataProvider: currentModel.Sink.MetadataProvider,
			MetadataRegion:   currentModel.Sink.MetadataRegion,
			PartitionFields:  make([]admin.DataLakePipelinesPartitionField, len(currentModel.Sink.PartitionFields)),
		},
		Source: &admin.IngestionSource{
			Type:           currentModel.Source.Type,
			ClusterName:    currentModel.Source.ClusterName,
			CollectionName: currentModel.Source.CollectionName,
			DatabaseName:   currentModel.Source.DatabaseName,
			GroupId:        currentModel.Source.GroupId,
		},
		Transformations: make([]admin.FieldTransformation, len(currentModel.Transformations)),
	}

	for i, partitionField := range currentModel.Sink.PartitionFields {
		dataLakeIntegrationPipeline.Sink.PartitionFields[i] = admin.DataLakePipelinesPartitionField{
			FieldName: *partitionField.FieldName,
			Order:     *partitionField.Order,
		}
	}

	for i, transformation := range currentModel.Transformations {
		dataLakeIntegrationPipeline.Transformations[i] = admin.FieldTransformation{
			Field: transformation.Field,
			Type:  transformation.Type,
		}
	}

	return &dataLakeIntegrationPipeline
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

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	groupID := *currentModel.ProjectId
	pipelineName := *currentModel.Name

	pe, response, err := client.Atlas20231115002.DataLakePipelinesApi.GetPipeline(ctx.Background(), groupID, pipelineName).Execute()

	if err != nil {
		return handleError(response, err)
	}

	model := ReadResponseModelGeneration(pe)
	if model == nil {
		errorMsg := "Response model from the API is empty or nil "
		return progress_events.GetFailedEventByResponse(errorMsg, response), nil
	}
	model.Profile = currentModel.Profile
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
		ResourceModel:   model}, nil
}

func ReadResponseModelGeneration(pe *admin.DataLakeIngestionPipeline) (model *Model) {
	if pe != nil {
		source := Source{
			Type:           pe.Source.Type,
			ClusterName:    pe.Source.ClusterName,
			CollectionName: pe.Source.CollectionName,
			DatabaseName:   pe.Source.DatabaseName,
			GroupId:        pe.Source.GroupId,
		}

		partitionArr := []PartitionFields{}
		sink := Sink{}

		if pe.Sink.PartitionFields != nil {
			for i := range pe.Sink.PartitionFields {
				partitionField := PartitionFields{
					FieldName: &pe.Sink.PartitionFields[i].FieldName,
					Order:     &pe.Sink.PartitionFields[i].Order,
				}
				partitionArr = append(partitionArr, partitionField)
			}
			sink.Type = pe.Sink.Type
			sink.MetadataProvider = pe.Sink.MetadataProvider
			sink.MetadataRegion = pe.Sink.MetadataRegion
			sink.PartitionFields = partitionArr
		}
		transformationsArr := []Transformations{}
		if pe.Transformations != nil {
			for i := range pe.Transformations {
				transformations := Transformations{
					Field: pe.Transformations[i].Field,
					Type:  pe.Transformations[i].Type,
				}
				transformationsArr = append(transformationsArr, transformations)
			}
		}

		createdStr := pe.CreatedDate.Format(time.RFC3339)
		lastUpdatedStr := pe.LastUpdatedDate.Format(time.RFC3339)

		models := Model{
			ProjectId:       pe.GroupId,
			Name:            pe.Name,
			Id:              pe.Id,
			CreatedDate:     &createdStr,
			LastUpdatedDate: &lastUpdatedStr,
			Sink:            &sink,
			Source:          &source,
			State:           pe.State,
			Transformations: transformationsArr,
		}
		return &models
	}
	return nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel)
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

	groupID := *currentModel.ProjectId
	pipelineName := *currentModel.Name
	dataLakeIntegrationPipeline := generateDataLakeIntegrationPipeline(currentModel)

	pe, response, err := client.Atlas20231115002.DataLakePipelinesApi.UpdatePipeline(ctx.Background(), groupID, pipelineName, dataLakeIntegrationPipeline).Execute()

	if err != nil {
		if response.StatusCode == http.StatusBadRequest {
			return progress_events.GetFailedEventByCode(fmt.Sprintf("Error during execution : %s", err.Error()),
				cloudformation.HandlerErrorCodeAlreadyExists), nil
		}
		return handleError(response, err)
	}

	model := ReadResponseModelGeneration(pe)
	if model == nil {
		errorMsg := "Response model from the API is empty or nil "
		return progress_events.GetFailedEventByResponse(errorMsg, response), nil
	}
	model.Profile = currentModel.Profile
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   model}, nil
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

	groupID := *currentModel.ProjectId
	pipelineName := *currentModel.Name

	_, response, err := client.Atlas20231115002.DataLakePipelinesApi.DeletePipeline(ctx.Background(), groupID, pipelineName).Execute()

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

	groupID := *currentModel.ProjectId
	pe, response, err := client.Atlas20231115002.DataLakePipelinesApi.ListPipelines(ctx.Background(), groupID).Execute()

	if err != nil {
		return handleError(response, err)
	}

	var list = make([]interface{}, 0)
	for ind := range pe {
		model := ReadResponseModelGeneration(&pe[ind])
		if model == nil {
			errorMsg := "Response model from the API is empty or nil "
			return progress_events.GetFailedEventByResponse(errorMsg, response), nil
		}
		model.Profile = currentModel.Profile
		list = append(list, *model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  list}, nil
}

func handleError(response *http.Response, err error) (handler.ProgressEvent, error) {
	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}
	return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error during execution : %s", err.Error()), response), nil
}
