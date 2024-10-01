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
	"context"
	"fmt"
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
	"go.mongodb.org/atlas-sdk/v20240805004/admin"
)

var CreateOrUpdateRequiredFields = []string{constants.ProjectID, constants.TenantName, constants.LimitName, constants.Value}
var ReadRequiredFields = []string{constants.ProjectID, constants.TenantName, constants.LimitName}
var DeleteRequiredFields = []string{constants.ProjectID, constants.TenantName, constants.LimitName}
var ListRequiredFields = []string{constants.ProjectID, constants.TenantName}

const (
	AlreadyExists = "already exists"
	DoesntExists  = "does not exist"
	CREATE        = "CREATE"
	READ          = "READ"
	UPDATE        = "UPDATE"
	DELETE        = "DELETE"
	LIST          = "LIST"
)

func setup() {
	util.SetupLogger("mongodb-atlas-federated-query-limit")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateOrUpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	_, _, err := getFederatedQueryLimit(atlas, currentModel)

	if err == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          AlreadyExists,
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}
	// create and update uses same PATCH API
	return createOrUpdateQueryLimit(currentModel, atlas, CREATE)
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	_, _ = logger.Debugf("Initiating Read Execute: %+v", currentModel)

	queryLimit, response, err := getFederatedQueryLimit(atlas, currentModel)

	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, READ, err)
	}
	currentModel.getQueryLimit(queryLimit)
	_, _ = logger.Debugf("Read Response: %+v", currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
		ResourceModel:   currentModel}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateOrUpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	// Check already exists or not
	_, response, err := getFederatedQueryLimit(atlas, currentModel)

	if err != nil && response.StatusCode == http.StatusNotFound {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          DoesntExists,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	return createOrUpdateQueryLimit(currentModel, atlas, UPDATE)
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	_, response, err := client.Atlas20231115014.DataFederationApi.DeleteOneDataFederationInstanceQueryLimit(
		context.Background(),
		*currentModel.ProjectId,
		*currentModel.TenantName,
		*currentModel.LimitName,
	).Execute()

	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
		ResourceModel:   nil}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	listQueryLimitsAPIResult, response, err := client.Atlas20231115014.DataFederationApi.ReturnFederatedDatabaseQueryLimits(
		context.Background(),
		*currentModel.ProjectId,
		*currentModel.TenantName,
	).Execute()

	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, LIST, err)
	}
	queryLimits := make([]interface{}, 0)
	for i := range listQueryLimitsAPIResult {
		queryLimit := Model{
			ProjectId: currentModel.ProjectId,
			Profile:   currentModel.Profile,
		}
		queryLimit.getQueryLimit(&listQueryLimitsAPIResult[i])
		queryLimits = append(queryLimits, queryLimit)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Completed",
		ResourceModels:  queryLimits}, nil
}

func handleError(response *http.Response, method string, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())
	_, _ = logger.Warn(errMsg)
	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}
	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}

func getFederatedQueryLimit(client *util.MongoDBClient, currentModel *Model) (*admin.DataFederationTenantQueryLimit, *http.Response, error) {
	getQueryLimitAPIRequest := client.Atlas20231115014.DataFederationApi.ReturnFederatedDatabaseQueryLimit(
		context.Background(),
		*currentModel.ProjectId,
		*currentModel.TenantName,
		*currentModel.LimitName,
	)
	queryLimit, response, err := getQueryLimitAPIRequest.Execute()
	return queryLimit, response, err
}

func createOrUpdateQueryLimit(currentModel *Model, client *util.MongoDBClient, method string) (handler.ProgressEvent, error) {
	queryLimitInput := currentModel.setQueryLimit()
	queryLimit, response, err := client.Atlas20231115014.DataFederationApi.CreateOneDataFederationQueryLimit(
		context.Background(),
		*currentModel.ProjectId,
		*currentModel.TenantName,
		*currentModel.LimitName,
		queryLimitInput,
	).Execute()

	if err != nil {
		return handleError(response, method, err)
	}

	currentModel.getQueryLimit(queryLimit)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         fmt.Sprintf("%s completed", method),
		ResourceModel:   currentModel}, nil
}

func (model *Model) setQueryLimit() *admin.DataFederationTenantQueryLimit {
	queryLimit := &admin.DataFederationTenantQueryLimit{
		OverrunPolicy: model.OverrunPolicy,
	}
	if model.Value != nil {
		queryLimit.Value = cast.ToInt64(*model.Value)
	}
	return queryLimit
}
func (model *Model) getQueryLimit(atlasQueryLimit *admin.DataFederationTenantQueryLimit) *Model {
	if atlasQueryLimit == nil {
		return nil
	}

	model.TenantName = atlasQueryLimit.TenantName
	model.OverrunPolicy = atlasQueryLimit.OverrunPolicy
	model.LimitName = &atlasQueryLimit.Name
	if atlasQueryLimit.CurrentUsage != nil {
		currentUsage := cast.ToString(atlasQueryLimit.CurrentUsage)
		model.CurrentUsage = &currentUsage
	}
	if atlasQueryLimit.DefaultLimit != nil {
		defaultLimit := cast.ToString(atlasQueryLimit.DefaultLimit)
		model.DefaultLimit = &defaultLimit
	}
	if atlasQueryLimit.LastModifiedDate != nil {
		lastModifiedDate := cast.ToString(atlasQueryLimit.LastModifiedDate)
		model.LastModifiedDate = &lastModifiedDate
	}
	if atlasQueryLimit.MaximumLimit != nil {
		maximumLimit := cast.ToString(atlasQueryLimit.MaximumLimit)
		model.MaximumLimit = &maximumLimit
	}

	value := cast.ToString(atlasQueryLimit.Value)
	model.Value = &value

	return model
}
