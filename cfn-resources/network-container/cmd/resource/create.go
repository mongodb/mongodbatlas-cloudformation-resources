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

	"go.mongodb.org/atlas-sdk/v20240805001/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

var createRequiredFields = []string{constants.ProjectID, constants.RegionName, constants.AtlasCIDRBlock}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if err := validateCreateModel(createRequiredFields, currentModel); err != nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         err.Error(),
			ResourceModel:   currentModel,
		}, err
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	containerRequest := &admin.CloudProviderContainer{
		ProviderName:   admin.PtrString(constants.AWS),
		RegionName:     currentModel.RegionName,
		AtlasCidrBlock: currentModel.AtlasCidrBlock,
	}

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

func createContainer(client *util.MongoDBClient, projectID string, request *admin.CloudProviderContainer) (string, error) {
	container, httpResponse, err := client.Atlas20231115002.NetworkPeeringApi.CreatePeeringContainer(context.Background(), projectID, request).Execute()
	if err == nil {
		return *container.Id, nil
	}

	if httpResponse.StatusCode != http.StatusConflict {
		return "", fmt.Errorf("error creating network container: %w", err)
	}

	_, _ = logger.Debugf("Container already exists for this group. Try return existing container. err: %v", err)

	args := admin.ListPeeringContainerByCloudProviderApiParams{
		GroupId:      projectID,
		ProviderName: request.ProviderName,
	}

	containers, _, err := client.Atlas20231115002.NetworkPeeringApi.ListPeeringContainerByCloudProviderWithParams(context.Background(), &args).Execute()
	if err != nil {
		return "", fmt.Errorf("error Containers.ListAll err:%v", err)
	}

	for i := range containers.Results {
		if containers.Results[i].RegionName == request.RegionName {
			return *containers.Results[i].Id, nil
		}
	}

	return "", errors.New("error creating network container")
}

func validateCreateModel(fields []string, model *Model) error {
	if !util.IsStringPresent(model.ProjectId) {
		return fmt.Errorf("error creating network container: `%s` must be set", constants.ProjectID)
	}

	if !util.IsStringPresent(model.RegionName) {
		return fmt.Errorf("`error creating network container: `%s` must be set", constants.RegionName)
	}

	if !util.IsStringPresent(model.AtlasCidrBlock) {
		return fmt.Errorf("error creating network container: `%s` must be set", constants.AtlasCIDRBlock)
	}

	if event := validator.ValidateModel(fields, model); event != nil {
		return errors.New(event.Message)
	}

	return nil
}
