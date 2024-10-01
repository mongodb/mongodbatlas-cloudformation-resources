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
	"encoding/json"
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
	"go.mongodb.org/atlas-sdk/v20240805004/admin"
)

func setup() {
	util.SetupLogger("mongodb-atlas-search-index")
}

var CreateRequiredFields = []string{constants.ProjectID, constants.ClusterName}
var ReadRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.IndexID}
var UpdateRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.IndexID}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.IndexID}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.Atlas20231115002

	ctx := context.Background()
	indexID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(indexID)
		currentModel.IndexId = &id
		return validateProgress(ctx, atlasV2, currentModel, string(handler.InProgress))
	}

	searchIndex, err := newSearchIndex(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			ResourceModel:    currentModel,
		}, nil
	}

	newSearchIndex, _, err := atlasV2.AtlasSearchApi.CreateAtlasSearchIndex(ctx, *currentModel.ProjectId, *currentModel.ClusterName, searchIndex).Execute()
	if err != nil {
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  cloudformation.OperationStatusFailed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	currentModel.Status = newSearchIndex.Status
	currentModel.IndexId = newSearchIndex.IndexID
	return handler.ProgressEvent{
		OperationStatus: status(currentModel),
		Message:         "Create Complete",
		ResourceModel:   currentModel,
		CallbackContext: map[string]any{
			"stateName": newSearchIndex.Status,
			"id":        currentModel.IndexId,
		},
		CallbackDelaySeconds: 120,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if currentModel.IndexId == nil {
		err := errors.New("no Id found in currentModel")
		return handler.ProgressEvent{
			OperationStatus:  cloudformation.OperationStatusFailed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	if errEvent := validator.ValidateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.Atlas20231115002

	searchIndex, resp, err := atlasV2.AtlasSearchApi.GetAtlasSearchIndex(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.IndexId).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
	}
	currentModel.Status = searchIndex.Status
	currentModel.Type = searchIndex.Type
	return handler.ProgressEvent{
		OperationStatus: cloudformation.OperationStatusSuccess,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if currentModel.IndexId == nil {
		err := errors.New("no Id found in currentModel")
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	if errEvent := validator.ValidateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.Atlas20231115002

	ctx := context.Background()
	indexID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(indexID)
		currentModel.IndexId = &id
		return validateProgress(ctx, atlasV2, currentModel, string(handler.InProgress))
	}
	searchIndex, err := newSearchIndex(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			ResourceModel:    currentModel,
		}, nil
	}

	updatedSearchIndex, res, err := atlasV2.AtlasSearchApi.UpdateAtlasSearchIndex(
		context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.IndexId, searchIndex).Execute()
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
	currentModel.Status = updatedSearchIndex.Status
	return handler.ProgressEvent{
		OperationStatus: status(currentModel),
		Message:         "Update Complete",
		ResourceModel:   currentModel,
		CallbackContext: map[string]any{
			"stateName": updatedSearchIndex.Status,
			"id":        currentModel.IndexId,
		},
		CallbackDelaySeconds: 120,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if currentModel.IndexId == nil {
		err := errors.New("no Id found in currentModel")
		return handler.ProgressEvent{
			OperationStatus:  cloudformation.OperationStatusFailed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	if errEvent := validator.ValidateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.Atlas20231115002

	ctx := context.Background()

	indexID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(indexID)
		currentModel.IndexId = &id
		return validateProgress(ctx, atlasV2, currentModel, string(handler.InProgress))
	}

	_, resp, err := atlasV2.AtlasSearchApi.DeleteAtlasSearchIndex(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.IndexId).Execute()
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
		CallbackContext: map[string]any{
			"stateName": handler.InProgress,
			"id":        currentModel.IndexId,
		},
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 120,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validator.ValidateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.Atlas20231115002

	ctx := context.Background()

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateProgress(ctx, atlasV2, currentModel, "IDLE")
	}

	indices, _, err := atlasV2.AtlasSearchApi.ListAtlasSearchIndexes(
		context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.CollectionName, *currentModel.Database).Execute()
	if err != nil {
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	response := make([]any, 0, len(indices))
	for i := range indices {
		response = append(response, indices[i])
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  response,
	}, nil
}

func newSearchIndex(currentModel *Model) (*admin.ClusterSearchIndex, error) {
	searchIndex := &admin.ClusterSearchIndex{
		Analyzer:       currentModel.Analyzer,
		CollectionName: aws.StringValue(currentModel.CollectionName),
		Database:       aws.StringValue(currentModel.Database),
		IndexID:        currentModel.IndexId,
		Name:           aws.StringValue(currentModel.Name),
		SearchAnalyzer: currentModel.SearchAnalyzer,
		Status:         currentModel.Status,
		Type:           currentModel.Type,
	}

	if fields, err := convertStringToInterfaceMap(currentModel.Fields); err == nil {
		searchIndex.Fields = fields
	}

	if currentModel.Mappings != nil {
		mapping, err := newMappings(currentModel)
		if err != nil {
			return nil, err
		}
		searchIndex.Mappings = mapping
	}
	analyzers := make([]admin.ApiAtlasFTSAnalyzers, 0, len(currentModel.Analyzers))
	for i := range currentModel.Analyzers {
		charFilters, err := ConvertToAnySlice(currentModel.Analyzers[i].CharFilters)
		if err != nil {
			return nil, err
		}

		tokenFilters, err := ConvertToAnySlice(currentModel.Analyzers[i].TokenFilters)
		if err != nil {
			return nil, err
		}

		s := admin.ApiAtlasFTSAnalyzers{
			CharFilters:  charFilters,
			Name:         *currentModel.Analyzers[i].Name,
			TokenFilters: tokenFilters,
			Tokenizer:    NewTokenizerModel(currentModel.Analyzers[i].Tokenizer),
		}
		analyzers = append(analyzers, s)
	}
	if len(analyzers) > 0 {
		searchIndex.Analyzers = analyzers
	}

	synonyms := make([]admin.SearchSynonymMappingDefinition, 0, len(currentModel.Synonyms))
	for i := range currentModel.Synonyms {
		s := admin.SearchSynonymMappingDefinition{
			Analyzer: *currentModel.Synonyms[i].Analyzer,
			Name:     *currentModel.Synonyms[i].Name,
			Source: admin.SynonymSource{
				Collection: *currentModel.Synonyms[i].Source.Collection,
			},
		}
		synonyms = append(synonyms, s)
	}
	if len(synonyms) > 0 {
		searchIndex.Synonyms = synonyms
	}
	return searchIndex, nil
}

func ConvertToAnySlice(input []string) ([]any, error) {
	var result []any

	for _, jsonStr := range input {
		var data any
		err := json.Unmarshal([]byte(jsonStr), &data)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}

func NewTokenizerModel(tokenizer *ApiAtlasFTSAnalyzersTokenizer) admin.ApiAtlasFTSAnalyzersTokenizer {
	if tokenizer == nil {
		return admin.ApiAtlasFTSAnalyzersTokenizer{}
	}
	return admin.ApiAtlasFTSAnalyzersTokenizer{
		MaxGram:        tokenizer.MaxGram,
		MinGram:        tokenizer.MinGram,
		Type:           tokenizer.Type,
		Group:          tokenizer.Group,
		Pattern:        tokenizer.Pattern,
		MaxTokenLength: tokenizer.MaxTokenLength,
	}
}

func newMappings(currentModel *Model) (*admin.ApiAtlasFTSMappings, error) {
	if currentModel.Mappings == nil {
		return nil, nil
	}

	fields, err := convertStringToInterface(currentModel.Mappings.Fields)
	if err != nil {
		return nil, err
	}
	return &admin.ApiAtlasFTSMappings{
		Dynamic: currentModel.Mappings.Dynamic,
		Fields:  fields,
	}, nil
}

func convertStringToInterface(fields *string) (map[string]any, error) {
	if !util.IsStringPresent(fields) {
		return nil, nil
	}
	var data map[string]any
	if err := json.Unmarshal([]byte(*fields), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func convertStringToInterfaceMap(fields *string) ([]map[string]any, error) {
	if !util.IsStringPresent(fields) {
		return nil, nil
	}
	var data []map[string]any
	if err := json.Unmarshal([]byte(*fields), &data); err != nil {
		return nil, err
	}
	return data, nil
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

func validateProgress(ctx context.Context, client *admin.APIClient, currentModel *Model, targetState string) (event handler.ProgressEvent, err error) {
	index, err := SearchIndexExists(ctx, client, currentModel)
	if err != nil {
		_, _ = logger.Debugf("Error Cluster validate progress() err: %+v", err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}
	if util.AreStringPtrEqual(index.Status, &targetState) {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = cloudformation.OperationStatusInProgress
		p.CallbackDelaySeconds = 120
		p.Message = "Pending"
		p.CallbackContext = map[string]any{
			"stateName": index.Status,
			"id":        currentModel.IndexId,
		}
		return p, nil
	}
	p := handler.NewProgressEvent()
	if util.AreStringPtrEqual(index.Status, admin.PtrString(cloudformation.OperationStatusFailed)) {
		p.OperationStatus = cloudformation.OperationStatusFailed
		p.Message = "Failed"
		p.HandlerErrorCode = cloudformation.HandlerErrorCodeInvalidRequest
		p.ResourceModel = currentModel
		return p, nil
	}
	p.OperationStatus = cloudformation.OperationStatusSuccess
	p.Message = "Complete"
	if util.IsStringPresent(index.Status) && *index.Status != "DELETED" {
		p.ResourceModel = currentModel
	}
	return p, nil
}

func SearchIndexExists(ctx context.Context, atlasV2 *admin.APIClient, currentModel *Model) (*admin.ClusterSearchIndex, error) {
	index, resp, err := atlasV2.AtlasSearchApi.GetAtlasSearchIndex(ctx, *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.IndexId).Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return &admin.ClusterSearchIndex{Status: admin.PtrString("DELETED")}, nil
		}
	}
	return index, err
}
