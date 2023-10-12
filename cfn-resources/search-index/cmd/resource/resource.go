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
	"fmt"
	"net/http"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20231001001/admin"
)

// indexFieldParts index field should be fieldName:fieldType.
const indexFieldParts = 2

func setup() {
	util.SetupLogger("search-index")
}

var CreateRequiredFields = []string{constants.ProjectID, constants.ClusterName}
var ReadRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.IndexID}
var UpdateRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.IndexID}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.IndexID}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.AtlasV2

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
		CallbackContext: map[string]interface{}{
			"stateName": newSearchIndex.Status,
			"id":        currentModel.IndexId,
		},
		CallbackDelaySeconds: 120,
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
		s := admin.ApiAtlasFTSAnalyzers{
			CharFilters:  convertToAnySlice(currentModel.Analyzers[i].CharFilters),
			Name:         *currentModel.Analyzers[i].Name,
			TokenFilters: convertToAnySlice(currentModel.Analyzers[i].TokenFilters),
			Tokenizer:    newTokenizerModel(currentModel.Analyzers[i].Tokenizer),
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

// convertToAnySlice function converts a slice of map[string]interface{} to a slice of interface{}
func convertToAnySlice(input []map[string]interface{}) []interface{} {
	result := make([]interface{}, len(input))
	for i, v := range input {
		result[i] = v
	}
	return result
}

func newTokenizerModel(tokenizer map[string]interface{}) admin.ApiAtlasFTSAnalyzersTokenizer {
	result := admin.ApiAtlasFTSAnalyzersTokenizer{}
	if v, ok := tokenizer["maxGram"]; ok {
		result.MaxGram = admin.PtrInt(cast.ToInt(v))
	}
	if v, ok := tokenizer["minGram"]; ok {
		result.MinGram = admin.PtrInt(cast.ToInt(v))
	}
	if v, ok := tokenizer["type"]; ok {
		result.Type = admin.PtrString(cast.ToString(v))
	}
	if v, ok := tokenizer["group"]; ok {
		result.Group = admin.PtrInt(cast.ToInt(v))
	}
	if v, ok := tokenizer["pattern"]; ok {
		result.Pattern = admin.PtrString(cast.ToString(v))
	}
	if v, ok := tokenizer["maxTokenLength"]; ok {
		result.MaxTokenLength = admin.PtrInt(cast.ToInt(v))
	}
	return result
}

func newMappings(currentModel *Model) (*admin.ApiAtlasFTSMappings, error) {
	if currentModel.Mappings == nil {
		return nil, nil
	}

	sec, err := newMappingsFields(currentModel.Mappings.Fields)
	if err != nil {
		return nil, err
	}
	return &admin.ApiAtlasFTSMappings{
		Dynamic: currentModel.Mappings.Dynamic,
		Fields:  sec,
	}, nil
}

func newMappingsFields(fields []string) (map[string]interface{}, error) {
	if len(fields) == 0 {
		return nil, nil
	}

	fieldsMap := make(map[string]interface{})
	for _, p := range fields {
		f := strings.SplitN(p, ":", 2)
		if len(f) != indexFieldParts {
			return nil, fmt.Errorf("partition should be fieldName:fieldType, got: %s", p)
		}

		fieldType, err := parseFieldType(f[1])
		if err != nil {
			return nil, err
		}

		fieldsMap[f[0]] = fieldType
	}

	return fieldsMap, nil
}

func parseFieldType(jsonStr string) (interface{}, error) {
	var result interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
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
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
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

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.AtlasV2

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
	return handler.ProgressEvent{
		OperationStatus: cloudformation.OperationStatusSuccess,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
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

	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.AtlasV2

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
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
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

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.AtlasV2

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
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.AtlasV2

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
	response := make([]interface{}, 0, len(indices))
	for i := range indices {
		response = append(response, indices[i])
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  response,
	}, nil
}

// Waits for the terminal stage from an intermediate stage
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
		p.CallbackContext = map[string]interface{}{
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
