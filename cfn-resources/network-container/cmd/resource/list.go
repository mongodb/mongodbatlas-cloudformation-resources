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
	"log"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"go.mongodb.org/atlas-sdk/v20240805004/admin"
)

var listRequiredFields = []string{constants.ProjectID}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	_, _ = logger.Debugf("List currentModel:%+v", currentModel)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	_, _ = logger.Debugf("List currentModel:%+v", currentModel)

	if errEvent := validateModel(listRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	aws := constants.AWS
	containerRequest := &admin.ListPeeringContainerByCloudProviderApiParams{
		ProviderName: &aws,
		GroupId:      *currentModel.ProjectId,
	}
	_, _ = logger.Debugf("List - containerRequest:%v", containerRequest)
	containerResponse, _, err := client.Atlas20231115002.NetworkPeeringApi.ListPeeringContainerByCloudProviderWithParams(context.Background(), containerRequest).Execute()
	if err != nil {
		_, _ = logger.Warnf("Error %v", err)
		return handler.ProgressEvent{}, err
	}

	_, _ = logger.Debugf("containerResponse:%v", containerResponse)

	containers := containerResponse.Results
	mm := make([]interface{}, 0)
	for i := range containers {
		mm = append(mm, completeByConnection(&containers[i], *currentModel.ProjectId, *currentModel.Profile))
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  mm,
	}, nil
}

func completeByConnection(c *admin.CloudProviderContainer, projectID, profileName string) Model {
	return Model{
		RegionName:     c.RegionName,
		Provisioned:    c.Provisioned,
		Id:             c.Id,
		VpcId:          c.VpcId,
		AtlasCidrBlock: c.AtlasCidrBlock,
		ProjectId:      &projectID,
		Profile:        &profileName,
	}
}
