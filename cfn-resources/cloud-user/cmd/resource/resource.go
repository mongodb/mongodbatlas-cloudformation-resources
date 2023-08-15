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
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	atlasSDK "go.mongodb.org/atlas-sdk/v20230201002/admin"
	"net/http"
	"sort"
)

var CreateRequiredFields = []string{}
var ReadRequiredFields = []string{}
var DeleteRequiredFields = []string{}

const (
	CREATE = "CREATE"
	READ   = "READ"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
	LIST   = "LIST"
)

func setup() {
	util.SetupLogger("mongodb-atlas-cloud-user")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Set the roles from model
	createCloudUser := setCloudUser(currentModel)
	createRequest := atlas.AtlasV2.MongoDBCloudUsersApi.CreateUser(
		context.Background(),
		&createCloudUser,
	)
	userDetails, response, err := createRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		return handleError(response, CREATE, err)
	}

	currentModel.Id = userDetails.Id

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   currentModel}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	userDetails, response, err := getUserDetails(atlas, currentModel)

	defer closeResponse(response)
	if err != nil {
		return handleError(response, READ, err)
	}

	currentModel.readAPIKeyDetails(*userDetails)
	_, _ = logger.Debugf("Read Response: %+v", currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
		ResourceModel:   currentModel}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	deleteRequest := atlas.AtlasV2.OrganizationsApi.RemoveOrganizationUser(
		context.Background(),
		*currentModel.OrgId,
		*currentModel.Id,
	)
	_, response, err := deleteRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		return handleError(response, DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
		ResourceModel:   nil}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func closeResponse(response *http.Response) {
	if response != nil {
		response.Body.Close()
	}
}

func handleError(response *http.Response, method string, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())
	_, _ = logger.Warn(errMsg)
	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}
	if response.StatusCode == http.StatusBadRequest {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}

func getUserDetails(atlas *util.MongoDBClient, currentModel *Model) (*atlasSDK.CloudAppUser, *http.Response, error) {
	getUserRequest := atlas.AtlasV2.MongoDBCloudUsersApi.GetUser(context.Background(), *currentModel.Id)
	return getUserRequest.Execute()
}

func areStringArraysEqualIgnoreOrder(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	sortedArr1 := make([]string, len(arr1))
	copy(sortedArr1, arr1)
	sort.Strings(sortedArr1)

	sortedArr2 := make([]string, len(arr2))
	copy(sortedArr2, arr2)
	sort.Strings(sortedArr2)

	for i := 0; i < len(sortedArr1); i++ {
		if sortedArr1[i] != sortedArr2[i] {
			return false
		}
	}

	return true
}

func (model *Model) readAPIKeyDetails(user atlasSDK.CloudAppUser) Model {
	model.Country = util.Pointer(user.Country)
	model.EmailAddress = util.Pointer(user.EmailAddress)
	model.FirstName = util.Pointer(user.FirstName)
	model.Id = user.Id
	model.LastAuth = util.Pointer(user.LastAuth.String())
	model.LastName = util.Pointer(user.LastName)
	model.MobileNumber = util.Pointer(user.MobileNumber)
	//model.Password = util.Pointer(user.Password)
	model.TeamIds = user.TeamIds
	model.Username = util.Pointer(user.Username)
	roles := make([]ProjectAssignment, len(user.Roles))
	var orgRoles []string
	for ind := range user.Roles {
		var role ProjectAssignment
		if user.Roles[ind].GroupId != nil {
			role.ProjectId = user.Roles[ind].GroupId
			role.RoleName = user.Roles[ind].RoleName
		}
		if user.Roles[ind].OrgId != nil {
			model.OrgId = user.Roles[ind].OrgId
			orgRoles = append(orgRoles, *user.Roles[ind].RoleName)
		}
	}
	model.ProjectAssignments = roles
	model.Roles = orgRoles

	links := make([]Link, len(user.Links))
	for ind := range user.Links {
		var link Link
		link.Href = user.Links[ind].Href
		link.Rel = user.Links[ind].Rel
	}
	model.Links = links
	return *model
}

func setCloudUser(currentModel *Model) atlasSDK.CloudAppUser {
	cloudUserInput := atlasSDK.CloudAppUser{}

	cloudUserInput.Country = *currentModel.Country

	cloudUserInput.EmailAddress = *currentModel.EmailAddress

	cloudUserInput.FirstName = *currentModel.FirstName

	cloudUserInput.LastName = *currentModel.LastName

	cloudUserInput.MobileNumber = *currentModel.MobileNumber

	cloudUserInput.Password = *currentModel.Password

	cloudUserInput.Username = *currentModel.Username

	var projectAssignments []atlasSDK.CloudAccessRoleAssignment
	for ind := range currentModel.Roles {
		var role atlasSDK.CloudAccessRoleAssignment
		role.RoleName = currentModel.ProjectAssignments[ind].RoleName
		role.GroupId = currentModel.ProjectAssignments[ind].ProjectId
		projectAssignments = append(projectAssignments, role)
	}

	cloudUserInput.Roles = projectAssignments
	for ind := range currentModel.Roles {
		var orgRole atlasSDK.CloudAccessRoleAssignment
		orgRole.OrgId = currentModel.OrgId
		orgRole.RoleName = util.Pointer(currentModel.Roles[ind])
		cloudUserInput.Roles = append(cloudUserInput.Roles, orgRole)
	}

	return cloudUserInput
}
