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
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

func setup() {
	util.SetupLogger("search-index")
}

var CreateRequiredFields = []string{constants.GroupID, constants.ClusterName, constants.PubKey, constants.PvtKey}
var ReadRequiredFields = []string{constants.GroupID, constants.ClusterName, constants.IndexID}
var UpdateRequiredFields = []string{constants.GroupID, constants.ClusterName, constants.IndexID}
var DeleteRequiredFields = []string{constants.GroupID, constants.ClusterName, constants.IndexID}

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
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, handlerError := util.NewMongoDBClient(req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	ctx := context.Background()
	indexID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(indexID)
		currentModel.IndexId = &id
		return validateProgress(ctx, client, currentModel, string(handler.InProgress))
	}
	searchIndex, err := ToSearchIndex(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			ResourceModel:    currentModel,
		}, nil
	}
	newSearchIndex, _, err := client.Search.CreateIndex(ctx, *currentModel.ProjectId, *currentModel.ClusterName, &searchIndex)
	if err != nil {
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  cloudformation.OperationStatusFailed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	currentModel.Status = &newSearchIndex.Status
	currentModel.IndexId = &newSearchIndex.IndexID
	return handler.ProgressEvent{
		OperationStatus: status(currentModel),
		Message:         "Create Complete",
		ResourceModel:   currentModel,
		CallbackContext: map[string]interface{}{
			"stateName": newSearchIndex.Status,
			"id":        currentModel.IndexId,
		},
		CallbackDelaySeconds: 120,
	}, nil
}

func ToSearchIndex(currentModel *Model) (mongodbatlas.SearchIndex, error) {
	return newSearchIndex(currentModel)
}

func newSearchIndex(currentModel *Model) (mongodbatlas.SearchIndex, error) {
	searchIndex := mongodbatlas.SearchIndex{
		Analyzer:       aws.StringValue(currentModel.Analyzer),
		CollectionName: aws.StringValue(currentModel.CollectionName),
		Database:       aws.StringValue(currentModel.Database),
		IndexID:        aws.StringValue(currentModel.IndexId),
		Name:           aws.StringValue(currentModel.Name),
		SearchAnalyzer: aws.StringValue(currentModel.SearchAnalyzer),
		Status:         aws.StringValue(currentModel.Status),
	}
	if currentModel.Mappings != nil {
		var sec map[string]interface{}
		var err error
		if currentModel.Mappings.Fields != nil {
			sec, err = cast.ToStringMapE(currentModel.Mappings.Fields)
			if err != nil {
				return mongodbatlas.SearchIndex{}, err
			}
		}
		searchIndex.Mappings = &mongodbatlas.IndexMapping{
			Dynamic: func() bool {
				if currentModel.Mappings.Dynamic != nil {
					return *currentModel.Mappings.Dynamic
				}
				return false
			}(),
			Fields: &sec,
		}
	}
	analyzers := make([]map[string]interface{}, 0, len(currentModel.Analyzers))
	for _, v := range currentModel.Analyzers {
		s, err := util.ToStringMapE(v)
		if err != nil {
			return mongodbatlas.SearchIndex{}, err
		}
		analyzers = append(analyzers, s)
	}
	if len(analyzers) > 0 {
		searchIndex.Analyzers = analyzers
	}

	synonyms := make([]map[string]interface{}, 0, len(currentModel.Synonyms))
	for _, v := range currentModel.Synonyms {
		s, err := util.ToStringMapE(v)
		if err != nil {
			return mongodbatlas.SearchIndex{}, err
		}
		synonyms = append(synonyms, s)
	}
	if len(synonyms) > 0 {
		searchIndex.Synonyms = synonyms
	}
	return searchIndex, nil
}

func status(currentModel *Model) handler.Status {
	switch *currentModel.Status {
	case string(handler.Success):
		return cloudformation.OperationStatusSuccess
	case string(handler.Failed):
		return cloudformation.OperationStatusFailed
	case string(handler.InProgress):
		return cloudformation.OperationStatusInProgress
	}
	return cloudformation.OperationStatusPending
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if currentModel.IndexId == nil {
		err := errors.New("no Id found in currentModel")
		return handler.ProgressEvent{
			OperationStatus:  cloudformation.OperationStatusFailed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, handlerError := util.NewMongoDBClient(req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	searchIndex, resp, err := client.Search.GetIndex(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.IndexId)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
	}
	currentModel.Name = &searchIndex.Name
	currentModel.Analyzer = &searchIndex.Analyzer
	currentModel.Database = &searchIndex.Database
	currentModel.CollectionName = &searchIndex.CollectionName
	currentModel.Mappings = &ApiAtlasFTSMappingsViewManual{}
	if searchIndex.Mappings != nil {
		currentModel.Mappings.Dynamic = &searchIndex.Mappings.Dynamic
	}
	return handler.ProgressEvent{
		OperationStatus: cloudformation.OperationStatusSuccess,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if currentModel.IndexId == nil {
		err := errors.New("no Id found in currentModel")
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, handlerError := util.NewMongoDBClient(req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	ctx := context.Background()
	indexID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(indexID)
		currentModel.IndexId = &id
		return validateProgress(ctx, client, currentModel, string(handler.InProgress))
	}
	searchIndex, err := ToSearchIndex(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			ResourceModel:    currentModel,
		}, nil
	}
	updatedSearchIndex, res, err := client.Search.UpdateIndex(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName,
		*currentModel.IndexId, &searchIndex)
	if err != nil {
		// Log and handle 404 ok
		if res != nil && res.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	currentModel.Status = &updatedSearchIndex.Status
	currentModel.IndexId = &updatedSearchIndex.IndexID
	currentModel.Name = &updatedSearchIndex.Name
	return handler.ProgressEvent{
		OperationStatus: status(currentModel),
		Message:         "Update Complete",
		ResourceModel:   currentModel,
		CallbackContext: map[string]interface{}{
			"stateName": updatedSearchIndex.Status,
			"id":        currentModel.IndexId,
		},
		CallbackDelaySeconds: 120,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if currentModel.IndexId == nil {
		err := errors.New("no Id found in currentModel")
		return handler.ProgressEvent{
			OperationStatus:  cloudformation.OperationStatusFailed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	if errEvent := validateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, handlerError := util.NewMongoDBClient(req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	ctx := context.Background()

	indexID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(indexID)
		currentModel.IndexId = &id
		return validateProgress(ctx, client, currentModel, string(handler.InProgress))
	}

	resp, err := client.Search.DeleteIndex(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.IndexId)
	if err != nil {
		if resp != nil && (resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusNotFound) {
			return handler.ProgressEvent{
				OperationStatus:  cloudformation.OperationStatusFailed,
				Message:          cloudformation.HandlerErrorCodeNotFound,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		return handler.ProgressEvent{
			OperationStatus:  cloudformation.OperationStatusFailed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	return handler.ProgressEvent{
		OperationStatus: cloudformation.OperationStatusInProgress,
		Message:         "Delete in progress",
		CallbackContext: map[string]interface{}{
			"stateName": handler.InProgress,
			"id":        currentModel.IndexId,
		},
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 120,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, handlerError := util.NewMongoDBClient(req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	ctx := context.Background()

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateProgress(ctx, client, currentModel, "IDLE")
	}

	indices, _, err := client.Search.ListIndexes(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName,
		*currentModel.Database, *currentModel.CollectionName, nil)
	if err != nil {
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	response := make([]interface{}, 0, len(indices))
	for _, v := range indices {
		response = append(response, v)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  response,
	}, nil
}

// Waits for the terminal stage from an intermediate stage
func validateProgress(ctx context.Context, client *mongodbatlas.Client, currentModel *Model, targetState string) (event handler.ProgressEvent, err error) {
	index, err := SearchIndexExists(ctx, client, currentModel)
	if err != nil {
		_, _ = logger.Debugf("Error Cluster validate progress() err: %+v", err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	if index.Status == targetState {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = cloudformation.OperationStatusInProgress
		p.CallbackDelaySeconds = 120
		p.Message = "Pending"
		p.CallbackContext = map[string]interface{}{
			"stateName": index.Status,
			"id":        currentModel.IndexId,
		}
		return p, nil
	}
	p := handler.NewProgressEvent()
	if index.Status == cloudformation.OperationStatusFailed {
		p.OperationStatus = cloudformation.OperationStatusFailed
		p.Message = "Failed"
		p.HandlerErrorCode = cloudformation.HandlerErrorCodeInvalidRequest
		p.ResourceModel = currentModel
		return p, nil
	}
	p.OperationStatus = cloudformation.OperationStatusSuccess
	p.Message = "Complete"
	if index.Status != "DELETED" {
		p.ResourceModel = currentModel
	}
	return p, nil
}

func SearchIndexExists(ctx context.Context, client *mongodbatlas.Client, currentModel *Model) (*mongodbatlas.SearchIndex, error) {
	index, resp, err := client.Search.GetIndex(ctx, *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.IndexId)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return &mongodbatlas.SearchIndex{Status: "DELETED"}, nil
		}
	}
	return index, err
}
