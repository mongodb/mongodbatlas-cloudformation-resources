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
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"go.mongodb.org/atlas/mongodbatlas"
)

var deleteRequiredFields = []string{constants.ProjectID, constants.ID}

func deleteOperation(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	_, _ = logger.Debugf("Delete currentModel:%+v", currentModel)

	if errEvent := validateModel(deleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	_, _ = logger.Debugf("Delete currentModel:%+v", currentModel)
	projectID := *currentModel.ProjectId
	containerID := *currentModel.Id

	containerResponse, response, err := client.Containers.Get(context.Background(), projectID, containerID)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource: %s", err.Error()),
			response.Response), nil
	}
	if containerResponse != nil && containerResponse.Provisioned != nil && *containerResponse.Provisioned {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message: `You are trying to delete a container that is in use. (container.provisioned = true)
			Please, make sure to delete the network peering and the atlas cluster before deleting the container`,
			HandlerErrorCode: cloudformation.HandlerErrorCodeResourceConflict,
		}, nil
	}

	if response, err := client.Containers.Delete(context.Background(), projectID, containerID); err != nil {
		return retryDeleteIfRequired(client, response, err, projectID, containerID)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func retryDeleteIfRequired(client *mongodbatlas.Client, response *mongodbatlas.Response, err error, projectID, containerID string) (handler.ProgressEvent, error) {
	if response.StatusCode != 409 {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource: %s", err.Error()),
			response.Response), err
	}

	// handling "CANNOT_DELETE_RECENTLY_CREATED_CONTAINER" error during release process:
	// During the release process, the container is created and deleted in a short period of time which cause
	// the deletion to fail with the error "CANNOT_DELETE_RECENTLY_CREATED_CONTAINER".
	time.Sleep(time.Second * 5)
	responseSecondCall, errSecondCall := client.Containers.Delete(context.Background(), projectID, containerID)
	if errSecondCall == nil {
		return handler.ProgressEvent{OperationStatus: handler.Success, Message: "Delete Complete"}, nil
	}

	// A second reason why the deletion can fail with 409 is because the container is in use.
	if responseSecondCall.StatusCode == 409 {
		// The deletion will fail if the there is an atlas cluster or network peering
		// available in the same region of the container
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          fmt.Sprintf("Please, make sure to delete the network peering and the atlas cluster before deleting the container: %s", errSecondCall.Error()),
			HandlerErrorCode: cloudformation.HandlerErrorCodeResourceConflict,
		}, nil
	}

	return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource: %s", errSecondCall.Error()),
		response.Response), errSecondCall
}
