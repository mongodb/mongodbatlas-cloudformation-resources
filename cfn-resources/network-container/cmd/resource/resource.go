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
	"log"
	"net/http"
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.RegionName, constants.AtlasCIDRBlock}
var ReadRequiredFields = []string{constants.ProjectID, constants.ID}
var UpdateRequiredFields = []string{constants.ProjectID, constants.ID}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ID}
var ListRequiredFields = []string{constants.ProjectID}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
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

	projectID := currentModel.ProjectId

	containerRequest := &mongodbatlas.Container{}

	if projectID == nil || *projectID == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network container: `%s` must be set", constants.ProjectID)
	}

	regionName := currentModel.RegionName
	if regionName == nil || *regionName == "" {
		return handler.ProgressEvent{}, fmt.Errorf("`error creating network container: `%s` must be set", constants.RegionName)
	}
	containerRequest.RegionName = *regionName
	containerRequest.ProviderName = constants.AWS
	CIDR := currentModel.AtlasCidrBlock
	if CIDR == nil || *CIDR == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network container: `%s` must be set", constants.AtlasCIDRBlock)
	}
	containerRequest.AtlasCIDRBlock = *CIDR
	containerResponse, res, err := client.Containers.Create(context.Background(), *projectID, containerRequest)
	if err != nil {
		if res.StatusCode == http.StatusConflict {
			_, _ = logger.Debugf("Container already exists for this group. Try return existing container. err: %v", err)
			containers, _, err2 := client.Containers.ListAll(context.Background(), *projectID, nil)
			if err2 != nil {
				_, _ = logger.Debugf("Error Containers.ListAll err:%v", err)
				return handler.ProgressEvent{}, fmt.Errorf("error Containers.ListAll err:%v", err)
			}
			_, _ = logger.Debugf("containers:%v", containers)
			if containers != nil {
				first := containers[0]
				_, _ = logger.Debugf("Will return reference to first container: first:%+v", first)
				currentModel.Id = &first.ID
			}
		} else {
			return handler.ProgressEvent{}, fmt.Errorf("error creating network container: %s", err)
		}
	} else {
		currentModel.Id = &containerResponse.ID
	}

	_, _ = logger.Debugf("Create about to return this --->> currentModel:%+v", currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
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
	containerID := *currentModel.Id

	containerResponse, response, err := client.Containers.Get(context.Background(), projectID, containerID)

	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response), nil
	}

	currentModel.RegionName = &containerResponse.RegionName
	currentModel.Provisioned = containerResponse.Provisioned
	currentModel.VpcId = &containerResponse.VPCID
	currentModel.AtlasCidrBlock = &containerResponse.AtlasCIDRBlock

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
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
	containerID := *currentModel.Id
	containerRequest := &mongodbatlas.Container{}

	CIDR := currentModel.AtlasCidrBlock
	if CIDR != nil {
		containerRequest.AtlasCIDRBlock = *CIDR
	}
	containerRequest.ProviderName = constants.AWS
	containerRequest.RegionName = *currentModel.RegionName
	containerResponse, resp, err := client.Containers.Update(context.Background(), projectID, containerID, containerRequest)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	currentModel.Id = &containerResponse.ID
	_, _ = logger.Debugf("Create network container - Id: %v", currentModel.Id)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	_, _ = logger.Debugf("Delete currentModel:%+v", currentModel)

	if errEvent := validateModel(DeleteRequiredFields, currentModel); errEvent != nil {
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

	response, err := deleteContainer(client, projectID, containerID)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func deleteContainer(client *mongodbatlas.Client, projectID string, containerID string) (*mongodbatlas.Response, error) {
	response, err := client.Containers.Delete(context.Background(), projectID, containerID)

	// handling "CANNOT_DELETE_RECENTLY_CREATED_CONTAINER" error
	if err != nil && response.StatusCode == 409 {
		time.Sleep(time.Second * 3)
		return client.Containers.Delete(context.Background(), projectID, containerID)
	}
	return response, err
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
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
		var m Model
		m.completeByConnection(containerResponse[i])
		m.ProjectId = currentModel.ProjectId
		m.Profile = currentModel.Profile
		mm = append(mm, m)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  mm,
	}, nil
}

func (m *Model) completeByConnection(c mongodbatlas.Container) {
	m.RegionName = &c.RegionName
	m.Provisioned = c.Provisioned
	m.Id = &c.ID
	m.VpcId = &c.VPCID
	m.AtlasCidrBlock = &c.AtlasCIDRBlock
}
