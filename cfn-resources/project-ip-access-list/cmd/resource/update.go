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
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if !util.IsStringPresent(currentModel.Profile) {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// "cfn test" will fail without these validations
	if len(prevModel.AccessList) == 0 {
		return handler.ProgressEvent{
			Message:          "The previous model does not have entry. You should use CREATE instead of UPDATE",
			OperationStatus:  handler.Failed,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
	}

	existingEntries, err := getAllEntries(client, *currentModel.ProjectId)
	if err != nil {
		return handler.ProgressEvent{
			Message:          "Error in retrieving the existing entries",
			OperationStatus:  handler.Failed,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, err
	}

	if *existingEntries.TotalCount == 0 {
		return handler.ProgressEvent{
			Message:          "You have no entry in the accesslist. You should use CREATE instead of UPDATE",
			OperationStatus:  handler.Failed,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
	}

	if len(currentModel.AccessList) == 0 {
		return handler.ProgressEvent{
			Message:          "You cannot have an empty accesslist. You shoud use DELETE instead of UPDATE",
			OperationStatus:  handler.Failed,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
	}

	// We need to make sure that the entries in the previous and current model are not in the accesslist.
	// Why do we need to do delete entries in the previous model?
	// Scenario: If the user updates the current model by removing one of the entries in the accesslist,
	// CFN will call the UPDATE operation which won't delete the removed entry bacause it is no longer in the current model.
	entriesToDelete := currentModel.AccessList
	entriesToDelete = append(entriesToDelete, prevModel.AccessList...)

	progressEvent := deleteEntriesForUpdate(entriesToDelete, *currentModel.ProjectId, client)
	if progressEvent.OperationStatus == handler.Failed {
		return progressEvent, nil
	}

	progressEvent, err = createEntries(currentModel, client)
	if progressEvent.OperationStatus == handler.Failed || err != nil {
		return progressEvent, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}
