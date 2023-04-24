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
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var createRequiredFields = []string{constants.ProjectID, constants.RegionName, constants.AtlasCIDRBlock}

func createOperation(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if err := validateCreateModel(createRequiredFields, currentModel); err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
			ResourceModel:   currentModel,
		}, err
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	containerRequest := &mongodbatlas.Container{}
	containerRequest.RegionName = *currentModel.RegionName
	containerRequest.ProviderName = constants.AWS
	containerRequest.AtlasCIDRBlock = *currentModel.AtlasCidrBlock

	containerID, err := createContainer(client, *currentModel.ProjectId, containerRequest)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
			ResourceModel:   currentModel,
		}, err
	}

	currentModel.Id = &containerID
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create complete",
		ResourceModel:   currentModel,
	}, nil
}

func createContainer(client *mongodbatlas.Client, projectID string, request *mongodbatlas.Container) (string, error) {
	container, httpResponse, err := client.Containers.Create(context.Background(), projectID, request)
	if err == nil {
		return container.ID, nil
	}

	if httpResponse.StatusCode != http.StatusConflict {
		return "", fmt.Errorf("error creating network container: %w", err)
	}

	_, _ = logger.Debugf("Container already exists for this group. Try return existing container. err: %v", err)
	containers, _, err := client.Containers.ListAll(context.Background(), projectID, nil)
	if err != nil {
		return "", fmt.Errorf("error Containers.ListAll err:%v", err)
	}

	for i := range containers {
		if containers[i].RegionName == request.RegionName {
			return containers[i].ID, nil
		}
	}

	return "", errors.New("error creating network container")
}

// function to validate inputs to all actions
func validateCreateModel(fields []string, model *Model) error {
	if model.ProjectId == nil || *model.ProjectId == "" {
		return fmt.Errorf("error creating network container: `%s` must be set", constants.ProjectID)
	}

	if model.RegionName == nil || *model.RegionName == "" {
		return fmt.Errorf("`error creating network container: `%s` must be set", constants.RegionName)
	}

	if model.AtlasCidrBlock == nil || *model.AtlasCidrBlock == "" {
		return fmt.Errorf("error creating network container: `%s` must be set", constants.AtlasCIDRBlock)
	}

	if event := validator.ValidateModel(fields, model); event != nil {
		return errors.New(event.Message)
	}

	return nil
}
