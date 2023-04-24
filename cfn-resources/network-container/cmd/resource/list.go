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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"go.mongodb.org/atlas/mongodbatlas"
)

var ListRequiredFields = []string{constants.ProjectID}

// List handles the List event from the Cloudformation service.
func ListOperation(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	_, _ = logger.Debugf("List currentModel:%+v", currentModel)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	_, _ = logger.Debugf("List currentModel:%+v", currentModel)

	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
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

	projectID := *currentModel.ProjectId
	containerRequest := &mongodbatlas.ContainersListOptions{
		ProviderName: constants.AWS,
		ListOptions:  mongodbatlas.ListOptions{},
	}
	_, _ = logger.Debugf("List projectId:%v, containerRequest:%v", projectID, containerRequest)
	containerResponse, _, err := client.Containers.List(context.TODO(), projectID, containerRequest)
	if err != nil {
		_, _ = logger.Warnf("Error %v", err)
		return handler.ProgressEvent{}, err
	}

	_, _ = logger.Debugf("containerResponse:%v", containerResponse)

	mm := make([]interface{}, 0)
	for i := range containerResponse {
		mm = append(mm, completeByConnection(containerResponse[i]))
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  mm,
	}, nil
}

func completeByConnection(c mongodbatlas.Container) Model {
	return Model{
		RegionName:     &c.RegionName,
		Provisioned:    c.Provisioned,
		Id:             &c.ID,
		VpcId:          &c.VPCID,
		AtlasCidrBlock: &c.AtlasCIDRBlock,
	}
}
