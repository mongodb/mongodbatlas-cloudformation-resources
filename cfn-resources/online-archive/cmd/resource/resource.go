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

	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.Criteria, constants.CriteriaType}
var ReadRequiredFields = []string{constants.ProjectID, constants.ArchiveID, constants.ClusterName}
var UpdateRequiredFields = []string{constants.ProjectID, constants.ArchiveID, constants.ClusterName, constants.Criteria}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ArchiveID, constants.ClusterName}
var ListRequiredFields = []string{constants.ProjectID}

var requiredCriteriaType = map[string][]string{
	"DATE":   {"DateField", "ExpireAfterDays"},
	"CUSTOM": {"Query"},
}

func setup() {
	util.SetupLogger("mongodb-atlas-online-archive")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if err := validator.ValidateModel(CreateRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	ctx := context.Background()
	archiveID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(archiveID)
		currentModel.ArchiveId = &id
		return validateProgress2(ctx, client, currentModel, "PENDING")
	}

	inputRequest, errHandler := mapToArchivePayload2(currentModel)
	if errHandler != nil {
		return *errHandler, nil
	}
	outputRequest, resp, err := client.AtlasV2.OnlineArchiveApi.CreateOnlineArchive(ctx, *currentModel.ProjectId, *currentModel.ClusterName, &inputRequest).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(err.Error(), resp), nil
	}
	currentModel.ArchiveId = outputRequest.Id
	currentModel.Criteria.ExpireAfterDays = outputRequest.Criteria.ExpireAfterDays
	currentModel.State = outputRequest.State
	currentModel.TotalCount = aws.Float64(1)
	return handler.ProgressEvent{
		OperationStatus: handler.InProgress,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
		CallbackContext: map[string]interface{}{
			"stateName": currentModel.State,
			"id":        currentModel.ArchiveId,
		},
		CallbackDelaySeconds: 20,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if currentModel.ArchiveId == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "no Id found in currentModel",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	if err := validator.ValidateModel(ReadRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	olArchive, resp, err := client.AtlasV2.OnlineArchiveApi.GetOnlineArchive(context.Background(), *currentModel.ProjectId, *currentModel.ArchiveId, *currentModel.ClusterName).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(err.Error(), resp), nil
	}
	currentModel.ArchiveId = olArchive.Id
	currentModel.State = olArchive.State
	currentModel.ProjectId = olArchive.GroupId
	currentModel.TotalCount = aws.Float64(1)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
		Message:         "read online archiving",
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if currentModel.ArchiveId == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "no Id found in currentModel",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	if err := validator.ValidateModel(UpdateRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	inputRequest, errHandler := mapToArchivePayload3(currentModel)
	if errHandler != nil {
		return *errHandler, nil
	}
	ctx := context.Background()
	if a, _ := ArchiveExists2(ctx, client, currentModel); *a.State == "DELETED" {
		return progress_events.GetFailedEventByResponse("Archive not found", &http.Response{StatusCode: 404}), nil
	}
	outputRequest, resp, err := client.AtlasV2.OnlineArchiveApi.UpdateOnlineArchive(ctx, *currentModel.ProjectId, *currentModel.ArchiveId, *currentModel.ClusterName, &inputRequest).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.ArchiveId = outputRequest.Id
	currentModel.Criteria.ExpireAfterDays = outputRequest.Criteria.ExpireAfterDays
	currentModel.State = outputRequest.State
	currentModel.TotalCount = aws.Float64(1)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
		Message:         "updated online archiving",
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if err := validator.ValidateModel(DeleteRequiredFields, currentModel); err != nil {
		return *err, nil
	}
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}
	ctx := context.Background()
	archiveID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(archiveID)
		currentModel.ArchiveId = &id
		return validateProgress2(ctx, client, currentModel, "PENDING")
	}

	if a, _ := ArchiveExists2(ctx, client, currentModel); *a.State == "DELETED" {
		return progress_events.GetFailedEventByResponse("Archive not found", &http.Response{StatusCode: 404}), nil
	}

	_, resp, err := client.AtlasV2.OnlineArchiveApi.DeleteOnlineArchive(ctx, *currentModel.ProjectId, *currentModel.ArchiveId, *currentModel.ClusterName).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.InProgress,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
		CallbackContext: map[string]interface{}{
			"stateName": currentModel.State,
			"id":        currentModel.ArchiveId,
		},
		CallbackDelaySeconds: 10,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if err := validator.ValidateModel(ListRequiredFields, currentModel); err != nil {
		return *err, nil
	}
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}
	params := admin.ListOnlineArchivesApiParams{
		GroupId:      *currentModel.ProjectId,
		ClusterName:  *currentModel.ClusterName,
		IncludeCount: currentModel.IncludeCount,
		ItemsPerPage: currentModel.ItemsPerPage,
		PageNum:      currentModel.PageNum,
	}
	archives, resp, err := client.AtlasV2.OnlineArchiveApi.ListOnlineArchivesWithParams(context.Background(), &params).Execute()
	if err != nil {
		return progress_events.GetFailedEventByResponse(err.Error(), resp), nil
	}

	resources := make([]any, 0, len(archives.Results))
	for _, v := range archives.Results {
		model := Model{
			ArchiveId: v.Id,
			ProjectId: currentModel.ProjectId,
			State:     v.State,
		}
		if archives.TotalCount != nil {
			model.TotalCount = aws.Float64(float64(*archives.TotalCount))
		}
		resources = append(resources, model)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModels:  resources,
	}, nil
}

func mapToArchivePayload2(currentModel *Model) (admin.BackupOnlineArchiveCreate, *handler.ProgressEvent) {
	requestInput := admin.BackupOnlineArchiveCreate{
		DbName:   *currentModel.DbName,
		CollName: *currentModel.CollName,
	}
	criteria, errHandler := mapCriteria2(currentModel)
	if errHandler != nil {
		return requestInput, errHandler
	}
	requestInput.Criteria = *criteria
	requestInput.PartitionFields = mapPartitionFields2(currentModel)
	return requestInput, nil
}

func mapToArchivePayload3(currentModel *Model) (admin.BackupOnlineArchive, *handler.ProgressEvent) {
	requestInput := admin.BackupOnlineArchive{
		DbName:   currentModel.DbName,
		CollName: currentModel.CollName,
	}
	criteria, errHandler := mapCriteria2(currentModel)
	if errHandler != nil {
		return requestInput, errHandler
	}
	requestInput.Criteria = criteria
	requestInput.PartitionFields = mapPartitionFields2(currentModel)
	return requestInput, nil
}

func mapCriteria2(currentModel *Model) (*admin.Criteria, *handler.ProgressEvent) {
	criteriaModel := *currentModel.Criteria
	criteriaInput := &admin.Criteria{
		Type:            criteriaModel.Type,
		DateField:       criteriaModel.DateField,
		ExpireAfterDays: criteriaModel.ExpireAfterDays,
		DateFormat:      criteriaModel.DateFormat,
	}
	if *criteriaInput.Type == "DATE" {
		requiredInputs := requiredCriteriaType[*criteriaInput.Type]
		criteriaInputDate := validator.ValidateModel(requiredInputs, criteriaModel)
		if criteriaInputDate != nil {
			return nil, criteriaInputDate
		}
	}
	if *criteriaInput.Type == "CUSTOM" {
		criteriaInput.Query = criteriaModel.Query
	}
	return criteriaInput, nil
}

func mapPartitionFields2(currentModel *Model) []admin.PartitionField {
	partitionFields := make([]admin.PartitionField, len(currentModel.PartitionFields))

	for i := range currentModel.PartitionFields {
		partitionField := &admin.PartitionField{}

		if currentModel.PartitionFields[i].FieldName != nil {
			partitionField.FieldName = *currentModel.PartitionFields[i].FieldName
		}

		if currentModel.PartitionFields[i].Order != nil {
			partitionField.Order = int(*currentModel.PartitionFields[i].Order)
		}

		partitionFields[i] = *partitionField
	}

	return partitionFields
}

// Waits for the terminal stage from an intermediate stage
func validateProgress2(ctx context.Context, client *util.MongoDBClient, currentModel *Model, targetState string) (event handler.ProgressEvent, err error) {
	archive, err := ArchiveExists2(ctx, client, currentModel)
	if err != nil {
		_, _ = logger.Debugf("Error archive archive exists err: %+v", err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	if *archive.State == targetState {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = cloudformation.OperationStatusInProgress
		p.CallbackDelaySeconds = 60
		p.Message = "Pending"
		p.CallbackContext = map[string]interface{}{
			"stateName": archive.State,
			"id":        currentModel.ArchiveId,
		}
		return p, nil
	}
	p := handler.NewProgressEvent()
	p.OperationStatus = cloudformation.OperationStatusSuccess
	p.Message = "Complete"
	if *archive.State != "DELETED" {
		p.ResourceModel = currentModel
	}
	return p, nil
}

func ArchiveExists2(ctx context.Context, client *util.MongoDBClient, currentModel *Model) (*admin.BackupOnlineArchive, error) {
	archive, resp, err := client.AtlasV2.OnlineArchiveApi.GetOnlineArchive(ctx, *currentModel.ProjectId, *currentModel.ArchiveId, *currentModel.ClusterName).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			state := "DELETED"
			return &admin.BackupOnlineArchive{State: &state}, nil
		}
	}
	return archive, nil
}
