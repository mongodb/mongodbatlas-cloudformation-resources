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
	"go.mongodb.org/atlas/mongodbatlas"
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

func validateModel(fields []string, model interface{}) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	ctx := context.Background() // test
	archiveID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(archiveID)
		currentModel.ArchiveId = &id
		return validateProgress(ctx, client, currentModel, "PENDING")
	}

	inputRequest, errHandler := mapToArchivePayload(currentModel)
	if errHandler != nil {
		return *errHandler, nil
	}
	outputRequest, res, err := client.OnlineArchives.Create(context.Background(), *currentModel.ProjectId,
		*currentModel.ClusterName, &inputRequest)
	if err != nil {
		_, _ = logger.Debugf("Error creating archive: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	currentModel.ArchiveId = &outputRequest.ID
	currentModel.Criteria.ExpireAfterDays = aws.Int(int(aws.Float64Value(outputRequest.Criteria.ExpireAfterDays)))
	currentModel.State = &outputRequest.State
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

func mapToArchivePayload(currentModel *Model) (mongodbatlas.OnlineArchive, *handler.ProgressEvent) {
	requestInput := mongodbatlas.OnlineArchive{
		DBName:   *currentModel.DbName,
		CollName: *currentModel.CollName,
	}
	criteria, errHandler := mapCriteria(currentModel)
	if errHandler != nil {
		return requestInput, errHandler
	}
	requestInput.Criteria = criteria
	return requestInput, nil
}

func mapCriteria(currentModel *Model) (*mongodbatlas.OnlineArchiveCriteria, *handler.ProgressEvent) {
	criteriaModel := *currentModel.Criteria
	criteriaInput := &mongodbatlas.OnlineArchiveCriteria{
		Type:            aws.StringValue(criteriaModel.Type),
		DateField:       aws.StringValue(criteriaModel.DateField),
		ExpireAfterDays: aws.Float64(float64(aws.IntValue(criteriaModel.ExpireAfterDays))),
		DateFormat:      aws.StringValue(criteriaModel.DateFormat),
	}
	if criteriaInput.Type == "DATE" {
		requiredInputs := requiredCriteriaType[criteriaInput.Type]
		criteriaInputDate := validateModel(requiredInputs, criteriaModel)
		if criteriaInputDate != nil {
			return nil, criteriaInputDate
		}
	}
	if criteriaInput.Type == "CUSTOM" {
		criteriaInput.Query = aws.StringValue(criteriaModel.Query)
	}
	return criteriaInput, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if currentModel.ArchiveId == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "no Id found in currentModel",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	var res *mongodbatlas.Response
	olArchive, res, err := client.OnlineArchives.Get(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName,
		*currentModel.ArchiveId)
	if err != nil {
		_, _ = logger.Debugf("Error fetching archive: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	currentModel.ArchiveId = &olArchive.ID
	currentModel.State = &olArchive.State
	currentModel.ProjectId = &olArchive.GroupID
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
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}
	inputRequest, errHandler := mapToArchivePayload(currentModel)
	if errHandler != nil {
		return *errHandler, nil
	}
	a, _ := ArchiveExists(context.Background(), client, currentModel)
	if a.State == "DELETED" {
		return progress_events.GetFailedEventByResponse("Archive not found", &http.Response{StatusCode: 404}), nil
	}
	outputRequest, res, err := client.OnlineArchives.Update(context.Background(), *currentModel.ProjectId,
		*currentModel.ClusterName, *currentModel.ArchiveId, &inputRequest)
	if err != nil {
		_, _ = logger.Debugf("Error updating archive: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.ArchiveId = &outputRequest.ID
	currentModel.Criteria.ExpireAfterDays = aws.Int(int(aws.Float64Value(outputRequest.Criteria.ExpireAfterDays)))
	currentModel.State = &outputRequest.State
	currentModel.TotalCount = aws.Float64(1)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
		Message:         "updated online archiving",
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}
	ctx := context.Background()
	archiveID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(archiveID)
		currentModel.ArchiveId = &id
		return validateProgress(ctx, client, currentModel, "PENDING")
	}

	a, _ := ArchiveExists(context.Background(), client, currentModel)
	if a.State == "DELETED" {
		return progress_events.GetFailedEventByResponse("Archive not found", &http.Response{StatusCode: 404}), nil
	}

	res, err := client.OnlineArchives.Delete(ctx, *currentModel.ProjectId,
		*currentModel.ClusterName, *currentModel.ArchiveId)
	if err != nil {
		_, _ = logger.Debugf("Error deleting archive: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
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
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}
	archives, res, err := client.OnlineArchives.List(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName,
		&mongodbatlas.ListOptions{
			PageNum:      aws.IntValue(currentModel.PageNum),
			ItemsPerPage: aws.IntValue(currentModel.ItemsPerPage),
			IncludeCount: aws.BoolValue(currentModel.IncludeCount),
		})
	if err != nil {
		_, _ = logger.Debugf("Error listing archive: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	resources := make([]any, 0, len(archives.Results))
	for _, v := range archives.Results {
		resources = append(resources, Model{
			ArchiveId:  &v.ID,
			ProjectId:  currentModel.ProjectId,
			TotalCount: aws.Float64(float64(archives.TotalCount)),
			State:      &v.State,
		})
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModels:  resources,
	}, nil
}

// Waits for the terminal stage from an intermediate stage
func validateProgress(ctx context.Context, client *mongodbatlas.Client, currentModel *Model, targetState string) (event handler.ProgressEvent, err error) {
	archive, err := ArchiveExists(ctx, client, currentModel)
	if err != nil {
		_, _ = logger.Debugf("Error archive archive exists err: %+v", err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	if archive.State == targetState {
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
	if archive.State != "DELETED" {
		p.ResourceModel = currentModel
	}
	return p, nil
}

func ArchiveExists(ctx context.Context, client *mongodbatlas.Client, currentModel *Model) (*mongodbatlas.OnlineArchive, error) {
	archive, resp, err := client.OnlineArchives.Get(ctx, *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.ArchiveId)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return &mongodbatlas.OnlineArchive{State: "DELETED"}, nil
		}
	}
	return archive, nil
}
