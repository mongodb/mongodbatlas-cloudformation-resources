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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

var updateRequiredFields = []string{constants.ProjectID, constants.ID}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if errEvent := validateModel(updateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	projectID := *currentModel.ProjectId
	containerID := *currentModel.Id
	containerRequest := &admin20231115002.CloudProviderContainer{}

	CIDR := currentModel.AtlasCidrBlock
	if CIDR != nil {
		containerRequest.AtlasCidrBlock = CIDR
	}

	containerRequest.ProviderName = admin20231115002.PtrString(constants.AWS)
	containerRequest.RegionName = currentModel.RegionName
	containerResponse, resp, err := client.Atlas20231115002.NetworkPeeringApi.UpdatePeeringContainer(context.Background(), projectID, containerID, containerRequest).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp), nil
	}

	currentModel.Id = containerResponse.Id
	_, _ = logger.Debugf("Create network container - Id: %v", currentModel.Id)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}
