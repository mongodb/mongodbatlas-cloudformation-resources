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
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

var readRequiredFields = []string{constants.ProjectID, constants.ID}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if errEvent := validateModel(readRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	projectID := *currentModel.ProjectId
	containerID := *currentModel.Id

	containerResponse, response, err := client.AtlasV2.NetworkPeeringApi.GetPeeringContainer(context.Background(), projectID, containerID).Execute()
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response), nil
	}

	currentModel.RegionName = containerResponse.RegionName
	currentModel.Provisioned = containerResponse.Provisioned
	currentModel.VpcId = containerResponse.VpcId
	currentModel.AtlasCidrBlock = containerResponse.AtlasCidrBlock

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}
