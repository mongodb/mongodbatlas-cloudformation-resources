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
	"errors"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115008/admin"
)

var CreateRequiredFields = []string{constants.OrgID, constants.APIUserID}
var ReadRequiredFields = []string{constants.OrgID, constants.APIUserID}
var DeleteRequiredFields = []string{constants.OrgID, constants.APIUserID}
var ListRequiredFields = []string{constants.OrgID, constants.APIUserID}

const (
	CREATE                 = "CREATE"
	READ                   = "READ"
	DELETE                 = "DELETE"
	LIST                   = "LIST"
	MutualExclusiveMessage = "Only one of IpAddress or CidrBlock is required"
	EitherOrMessage        = "Either IpAddress or CidrBlock is required"
)

// validateModel to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-access-list-api-key")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	orgID := *currentModel.OrgId
	apiKeyID := *currentModel.APIUserId

	if currentModel.CidrBlock == nil && currentModel.IpAddress == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          EitherOrMessage,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if currentModel.CidrBlock != nil && currentModel.IpAddress != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          MutualExclusiveMessage,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	// createReq.ApiService.
	entryList := make([]admin.UserAccessListRequest, 0)
	var access admin.UserAccessListRequest
	if currentModel.CidrBlock != nil {
		access.CidrBlock = currentModel.CidrBlock
	}
	if currentModel.IpAddress != nil {
		access.IpAddress = currentModel.IpAddress
	}
	entryList = append(entryList, access)

	createAccessListAPIKey := client.AtlasSDK.ProgrammaticAPIKeysApi.CreateApiKeyAccessList(context.Background(), orgID, apiKeyID, &entryList)
	_, response, err := createAccessListAPIKey.Execute()

	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, CREATE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	orgID := *currentModel.OrgId
	apiKeyID := *currentModel.APIUserId
	if currentModel.CidrBlock == nil && currentModel.IpAddress == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          EitherOrMessage,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if currentModel.CidrBlock != nil && currentModel.IpAddress != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          MutualExclusiveMessage,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	entry := getEntryAddress(currentModel)
	readAccessListAPIKey := client.Atlas20231115002.ProgrammaticAPIKeysApi.GetApiKeyAccessList(context.Background(), orgID, entry, apiKeyID)
	_, response, err := readAccessListAPIKey.Execute()
	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, READ, err)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
		ResourceModel:   currentModel}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	orgID := *currentModel.OrgId
	apiKeyID := *currentModel.APIUserId
	if currentModel.CidrBlock == nil && currentModel.IpAddress == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          EitherOrMessage,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if currentModel.CidrBlock != nil && currentModel.IpAddress != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          MutualExclusiveMessage,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	entry := getEntryAddress(currentModel)
	deleteAccessListAPIKey := client.AtlasSDK.ProgrammaticAPIKeysApi.DeleteApiKeyAccessListEntry(context.Background(), orgID, apiKeyID, entry)
	_, response, err := deleteAccessListAPIKey.Execute()

	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
		ResourceModel:   nil}, nil
}

func getEntryAddress(currentModel *Model) string {
	var entry string
	if currentModel.CidrBlock != nil {
		entry = *currentModel.CidrBlock
	}
	if currentModel.IpAddress != nil {
		entry = *currentModel.IpAddress
	}
	return entry
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	orgID := *currentModel.OrgId
	apiKeyID := *currentModel.APIUserId

	accessListResponse, response, err := client.AtlasSDK.ProgrammaticAPIKeysApi.ListApiKeyAccessListsEntries(context.Background(), orgID, apiKeyID).Execute()

	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, LIST, err)
	}

	accessListModels := make([]interface{}, 0)
	apiResults := accessListResponse.GetResults()
	for i := range apiResults {
		l := apiResults[i]
		label := Model{
			CidrBlock: l.CidrBlock,
			APIUserId: currentModel.APIUserId,
			OrgId:     currentModel.OrgId,
			Profile:   currentModel.Profile,
			IpAddress: l.IpAddress,
		}
		accessListModels = append(accessListModels, label)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  accessListModels,
	}, nil
}

func handleError(response *http.Response, method string, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())
	_, _ = logger.Warn(errMsg)
	if response == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInternalFailure}, nil
	}

	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}
	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}
