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
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if len(currentModel.AccessList) == 0 {
		return progressevents.GetFailedEventByCode("AccessList must not be empty", string(types.HandlerErrorCodeInvalidRequest)), nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	// Create atlas client
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	event, err := createEntries(currentModel, client)
	if event.OperationStatus == handler.Failed || err != nil {
		return event, nil
	}

	_, _ = logger.Debugf("Create --- currentModel:%+v", currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

func createEntries(model *Model, client *util.MongoDBClient) (handler.ProgressEvent, error) {
	request, err := newPaginatedNetworkAccess(model)
	if err != nil {
		return handler.ProgressEvent{
			Message:          "Error in parsing the resource schema",
			OperationStatus:  handler.Failed,
			HandlerErrorCode: string(types.HandlerErrorCodeAlreadyExists)}, err
	}

	projectID := *model.ProjectId

	if isEntryAlreadyInAccessList, err := isEntryAlreadyInAccessList(client, model); isEntryAlreadyInAccessList || err != nil {
		if err != nil {
			return handler.ProgressEvent{
				Message:          fmt.Sprintf("Error validating entries: %s", err.Error()),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: string(types.HandlerErrorCodeInternalFailure)}, err
		}
		return handler.ProgressEvent{
			Message:          "Entry already exists in the access list",
			OperationStatus:  handler.Failed,
			HandlerErrorCode: string(types.HandlerErrorCodeAlreadyExists)}, err
	}

	if _, _, err := client.Atlas20231115002.ProjectIPAccessListApi.CreateProjectIpAccessList(context.Background(), projectID, &request.Results).Execute(); err != nil {
		_, _ = logger.Warnf("Error createEntries projectId:%s, err:%+v", projectID, err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest)}, err
	}

	return handler.ProgressEvent{}, nil
}
