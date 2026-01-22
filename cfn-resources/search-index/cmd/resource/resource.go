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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20250312012/admin"
)

func setup() {
	util.SetupLogger("mongodb-atlas-search-index")
}

var CreateRequiredFields = []string{constants.ProjectID, constants.ClusterName}
var ReadRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.IndexID}
var UpdateRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.IndexID}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.IndexID}
var ListRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.CollectionName, constants.Database}

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
	atlasV2 := client.AtlasSDK

	ctx := context.Background()
	indexID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(indexID)
		currentModel.IndexId = &id
		return validateProgress(ctx, atlasV2, currentModel, string(handler.InProgress))
	}

	searchIndexRequest, err := newSearchIndexCreateRequest(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
			ResourceModel:    currentModel,
		}, nil
	}

	newSearchIndex, _, err := atlasV2.AtlasSearchApi.CreateClusterSearchIndex(ctx, *currentModel.ProjectId, *currentModel.ClusterName, searchIndexRequest).Execute()
	if err != nil {
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest)}, nil
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
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
	}
	if errEvent := validator.ValidateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.AtlasSDK

	searchIndex, resp, err := atlasV2.AtlasSearchApi.GetClusterSearchIndex(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.IndexId).Execute()
	if err != nil {
		if util.StatusNotFound(resp) {
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
		}
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: string(types.HandlerErrorCodeServiceInternalError)}, nil
	}
	currentModel.Status = searchIndex.Status
	currentModel.Type = searchIndex.Type
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
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
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
	}

	if errEvent := validator.ValidateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.AtlasSDK

	ctx := context.Background()
	indexID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(indexID)
		currentModel.IndexId = &id
		return validateProgress(ctx, atlasV2, currentModel, string(handler.InProgress))
	}
	searchIndexRequest, err := newSearchIndexUpdateRequest(currentModel)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
			ResourceModel:    currentModel,
		}, nil
	}

	updatedSearchIndex, res, err := atlasV2.AtlasSearchApi.UpdateClusterSearchIndex(
		context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.IndexId, searchIndexRequest).Execute()
	if err != nil {
		if util.StatusNotFound(res) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
		}
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: string(types.HandlerErrorCodeServiceInternalError)}, nil
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
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
	}

	if errEvent := validator.ValidateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.AtlasSDK

	ctx := context.Background()

	indexID, iOK := req.CallbackContext["id"]
	if _, ok := req.CallbackContext["stateName"]; ok && iOK {
		id := cast.ToString(indexID)
		currentModel.IndexId = &id
		return validateProgress(ctx, atlasV2, currentModel, string(handler.InProgress))
	}

	resp, err := atlasV2.AtlasSearchApi.DeleteClusterSearchIndex(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.IndexId).Execute()
	if err != nil {
		if util.StatusInternalServerError(resp) || util.StatusNotFound(resp) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          string(types.HandlerErrorCodeNotFound),
				HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
		}
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
	}
	return handler.ProgressEvent{
		OperationStatus: handler.InProgress,
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
	if errEvent := validator.ValidateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}
	atlasV2 := client.AtlasSDK

	ctx := context.Background()

	if _, ok := req.CallbackContext["stateName"]; ok {
		return validateProgress(ctx, atlasV2, currentModel, "IDLE")
	}

	indices, _, err := atlasV2.AtlasSearchApi.ListSearchIndex(
		context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.CollectionName, *currentModel.Database).Execute()
	if err != nil {
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: string(types.HandlerErrorCodeServiceInternalError)}, nil
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

// buildSearchIndexDefinition builds the common definition structure used by both create and update operations
type searchIndexDefinition struct {
	Analyzer       *string
	SearchAnalyzer *string
	NumPartitions  *int
	Fields         *[]any
	Mappings       *admin.SearchMappings
	Analyzers      *[]admin.AtlasSearchAnalyzer
	Synonyms       *[]admin.SearchSynonymMappingDefinition
	StoredSource   any
	TypeSets       *[]admin.SearchTypeSets
}

func buildSearchIndexDefinition(currentModel *Model) (*searchIndexDefinition, error) {
	def := &searchIndexDefinition{
		Analyzer:       currentModel.Analyzer,
		SearchAnalyzer: currentModel.SearchAnalyzer,
		NumPartitions:  currentModel.NumPartitions,
	}

	// Add Fields support
	if fields, err := convertStringToInterfaceMap(currentModel.Fields); err == nil && fields != nil {
		fieldsAny := make([]any, 0, len(fields))
		for _, f := range fields {
			fieldsAny = append(fieldsAny, f)
		}
		def.Fields = &fieldsAny
	}

	// Add Mappings support
	if currentModel.Mappings != nil {
		mapping, err := newMappings(currentModel)
		if err != nil {
			return nil, err
		}
		def.Mappings = mapping
	}

	// Add Analyzers support
	if len(currentModel.Analyzers) > 0 {
		analyzers := make([]admin.AtlasSearchAnalyzer, 0, len(currentModel.Analyzers))
		for i := range currentModel.Analyzers {
			charFilters, err := ConvertToAnySlice(currentModel.Analyzers[i].CharFilters)
			if err != nil {
				return nil, err
			}
			tokenFilters, err := ConvertToAnySlice(currentModel.Analyzers[i].TokenFilters)
			if err != nil {
				return nil, err
			}
			tokenizer := NewTokenizerModel(currentModel.Analyzers[i].Tokenizer)
			analyzers = append(analyzers, admin.AtlasSearchAnalyzer{
				CharFilters:  &charFilters,
				Name:         *currentModel.Analyzers[i].Name,
				TokenFilters: &tokenFilters,
				Tokenizer:    &tokenizer,
			})
		}
		def.Analyzers = &analyzers
	}

	// Add Synonyms support
	if len(currentModel.Synonyms) > 0 {
		synonyms := make([]admin.SearchSynonymMappingDefinition, 0, len(currentModel.Synonyms))
		for i := range currentModel.Synonyms {
			synonyms = append(synonyms, admin.SearchSynonymMappingDefinition{
				Analyzer: *currentModel.Synonyms[i].Analyzer,
				Name:     *currentModel.Synonyms[i].Name,
				Source: admin.SynonymSource{
					Collection: *currentModel.Synonyms[i].Source.Collection,
				},
			})
		}
		def.Synonyms = &synonyms
	}

	// Add StoredSource support
	if storedSource, err := ConvertStringToStoredSource(currentModel.StoredSource); err == nil && storedSource != nil {
		def.StoredSource = storedSource
	}

	// Add TypeSets support
	if len(currentModel.TypeSets) > 0 {
		typeSets := make([]admin.SearchTypeSets, 0, len(currentModel.TypeSets))
		for i := range currentModel.TypeSets {
			ts := admin.SearchTypeSets{
				Name: *currentModel.TypeSets[i].Name,
			}
			if typesList, err := convertStringToInterfaceMap(currentModel.TypeSets[i].Types); err == nil && typesList != nil {
				typesAny := make([]any, 0, len(typesList))
				for _, t := range typesList {
					typesAny = append(typesAny, t)
				}
				ts.Types = &typesAny
			}
			typeSets = append(typeSets, ts)
		}
		def.TypeSets = &typeSets
	}

	return def, nil
}

func newSearchIndexCreateRequest(currentModel *Model) (*admin.SearchIndexCreateRequest, error) {
	def, err := buildSearchIndexDefinition(currentModel)
	if err != nil {
		return nil, err
	}

	return &admin.SearchIndexCreateRequest{
		CollectionName: aws.ToString(currentModel.CollectionName),
		Database:       aws.ToString(currentModel.Database),
		Name:           aws.ToString(currentModel.Name),
		Type:           currentModel.Type,
		Definition: &admin.BaseSearchIndexCreateRequestDefinition{
			Analyzer:       def.Analyzer,
			SearchAnalyzer: def.SearchAnalyzer,
			NumPartitions:  def.NumPartitions,
			Fields:         def.Fields,
			Mappings:       def.Mappings,
			Analyzers:      def.Analyzers,
			Synonyms:       def.Synonyms,
			StoredSource:   def.StoredSource,
			TypeSets:       def.TypeSets,
		},
	}, nil
}

func newSearchIndexUpdateRequest(currentModel *Model) (*admin.SearchIndexUpdateRequest, error) {
	def, err := buildSearchIndexDefinition(currentModel)
	if err != nil {
		return nil, err
	}

	return &admin.SearchIndexUpdateRequest{
		Definition: admin.SearchIndexUpdateRequestDefinition{
			Analyzer:       def.Analyzer,
			SearchAnalyzer: def.SearchAnalyzer,
			NumPartitions:  def.NumPartitions,
			Fields:         def.Fields,
			Mappings:       def.Mappings,
			Analyzers:      def.Analyzers,
			Synonyms:       def.Synonyms,
			StoredSource:   def.StoredSource,
			TypeSets:       def.TypeSets,
		},
	}, nil
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

func NewTokenizerModel(tokenizer *ApiAtlasFTSAnalyzersTokenizer) map[string]any {
	if tokenizer == nil {
		return nil
	}
	result := make(map[string]any)
	if tokenizer.MaxGram != nil {
		result["maxGram"] = *tokenizer.MaxGram
	}
	if tokenizer.MinGram != nil {
		result["minGram"] = *tokenizer.MinGram
	}
	if tokenizer.Type != nil {
		result["type"] = *tokenizer.Type
	}
	if tokenizer.Group != nil {
		result["group"] = *tokenizer.Group
	}
	if tokenizer.Pattern != nil {
		result["pattern"] = *tokenizer.Pattern
	}
	if tokenizer.MaxTokenLength != nil {
		result["maxTokenLength"] = *tokenizer.MaxTokenLength
	}
	return result
}

func newMappings(currentModel *Model) (*admin.SearchMappings, error) {
	if currentModel.Mappings == nil {
		return nil, nil
	}

	fields, err := convertStringToInterface(currentModel.Mappings.Fields)
	if err != nil {
		return nil, err
	}

	mapping := &admin.SearchMappings{Fields: &fields}

	// DynamicConfig takes precedence over Dynamic
	if util.IsStringPresent(currentModel.Mappings.DynamicConfig) {
		var dynamicConfig any
		if err := json.Unmarshal([]byte(*currentModel.Mappings.DynamicConfig), &dynamicConfig); err != nil {
			return nil, err
		}
		mapping.Dynamic = dynamicConfig
	} else if currentModel.Mappings.Dynamic != nil {
		mapping.Dynamic = currentModel.Mappings.Dynamic
	}

	return mapping, nil
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

func ConvertStringToStoredSource(storedSource *string) (any, error) {
	if !util.IsStringPresent(storedSource) {
		return nil, nil
	}

	// Try to parse as boolean first
	switch *storedSource {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	// Otherwise parse as JSON object
	var data any
	if err := json.Unmarshal([]byte(*storedSource), &data); err != nil {
		return nil, err
	}
	return data, nil
}

func status(currentModel *Model) handler.Status {
	switch *currentModel.Status {
	case string(handler.Success):
		return handler.Success
	case string(handler.Failed):
		return handler.Failed
	case string(handler.InProgress):
		return handler.InProgress
	}
	return handler.InProgress
}

func validateProgress(ctx context.Context, client *admin.APIClient, currentModel *Model, targetState string) (handler.ProgressEvent, error) {
	index, err := SearchIndexExists(ctx, client, currentModel)
	if err != nil {
		_, _ = logger.Debugf("Error Cluster validate progress() err: %+v", err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: string(types.HandlerErrorCodeServiceInternalError)}, nil
	}
	if util.AreStringPtrEqual(index.Status, &targetState) {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = handler.InProgress
		p.CallbackDelaySeconds = 120
		p.Message = "Pending"
		p.CallbackContext = map[string]any{
			"stateName": index.Status,
			"id":        currentModel.IndexId,
		}
		return p, nil
	}
	p := handler.NewProgressEvent()
	if util.AreStringPtrEqual(index.Status, admin.PtrString(string(handler.Failed))) {
		p.OperationStatus = handler.Failed
		p.Message = "Failed"
		p.HandlerErrorCode = string(types.HandlerErrorCodeInvalidRequest)
		p.ResourceModel = currentModel
		return p, nil
	}
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	if util.IsStringPresent(index.Status) && *index.Status != "DELETED" {
		p.ResourceModel = currentModel
	}
	return p, nil
}

func SearchIndexExists(ctx context.Context, atlasV2 *admin.APIClient, currentModel *Model) (*admin.SearchIndexResponse, error) {
	index, resp, err := atlasV2.AtlasSearchApi.GetClusterSearchIndex(ctx, *currentModel.ProjectId, *currentModel.ClusterName, *currentModel.IndexId).Execute()
	if err != nil {
		if util.StatusNotFound(resp) {
			return &admin.SearchIndexResponse{Status: admin.PtrString("DELETED")}, nil
		}
	}
	return index, err
}
