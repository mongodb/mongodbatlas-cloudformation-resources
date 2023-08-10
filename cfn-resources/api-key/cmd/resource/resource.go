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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	atlasSDK "go.mongodb.org/atlas-sdk/v20230201002/admin"
	"net/http"
	"sort"
)

var CreateRequiredFields = []string{constants.OrgID, constants.Description}
var UpdateRequiredFields = []string{constants.OrgID, constants.ID}
var ReadRequiredFields = []string{constants.OrgID, constants.ID}
var DeleteRequiredFields = []string{constants.OrgID, constants.ID}
var ListRequiredFields = []string{constants.OrgID}

const (
	CREATE = "CREATE"
	READ   = "READ"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
	LIST   = "LIST"
)

func setup() {
	util.SetupLogger("mongodb-atlas-api-key")
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
	apiKeyInput := atlasSDK.CreateAtlasOrganizationApiKey{
		Desc:  currentModel.Description,
		Roles: currentModel.Roles,
	}
	apiKeyRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.CreateApiKey(
		context.Background(),
		*currentModel.OrgId,
		&apiKeyInput,
	)
	apiKeyUserDetails, response, err := apiKeyRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		return handleError(response, CREATE, err)
	}

	// Read response
	currentModel.Id = apiKeyUserDetails.Id

	// Assign Org APIKey to given projects i.e. projectAssignments
	if len(currentModel.ProjectAssignments) > 0 {
		for i := range currentModel.ProjectAssignments {
			_, response, err = assignOrgKeyToProject(currentModel.ProjectAssignments[i], atlas, currentModel.Id)
			if err != nil {
				break
			}
		}
		defer closeResponse(response)
		if err != nil {
			return handleError(response, CREATE, err)
		}
	}

	apiKeyUserDetails, response, err = getAPIkeyDetails(atlas, currentModel)
	model := currentModel.readAPIKeyDetails(apiKeyUserDetails)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   model}, nil
}

func updateOrgKeyToProject(projectAssignment ProjectAssignment, atlas *util.MongoDBClient, orgKeyId *string) (*atlasSDK.ApiKeyUserDetails, *http.Response, error) {
	// Set the roles from model
	projectAPIKeyInput := atlasSDK.CreateAtlasProjectApiKey{
		Roles: projectAssignment.Roles,
	}
	assignAPIRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.UpdateApiKeyRoles(
		context.Background(),
		*projectAssignment.ProjectId,
		*orgKeyId,
		&projectAPIKeyInput,
	)

	return assignAPIRequest.Execute()
}

func assignOrgKeyToProject(projectAssignment ProjectAssignment, atlas *util.MongoDBClient, orgKeyId *string) (*atlasSDK.ApiKeyUserDetails, *http.Response, error) {
	// Set the roles from model
	accessRoleAssignments := make([]atlasSDK.UserAccessRoleAssignment, 1)
	projectAPIKeyInput := atlasSDK.UserAccessRoleAssignment{
		ApiUserId: orgKeyId,
		Roles:     projectAssignment.Roles,
	}
	accessRoleAssignments[0] = projectAPIKeyInput
	assignAPIRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.AddProjectApiKey(
		context.Background(),
		*projectAssignment.ProjectId,
		*orgKeyId,
		&accessRoleAssignments,
	)

	return assignAPIRequest.Execute()
}

func unAssignOrgKeyToProject(projectAssignment ProjectAssignment, atlas *util.MongoDBClient, orgKeyId *string) (map[string]interface{}, *http.Response, error) {
	unAssignAPIRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.RemoveProjectApiKey(
		context.Background(),
		*projectAssignment.ProjectId,
		*orgKeyId,
	)

	return unAssignAPIRequest.Execute()
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

	apiKeyUserDetails, response, err := getAPIkeyDetails(atlas, currentModel)

	defer closeResponse(response)
	if err != nil {
		return handleError(response, READ, err)
	}

	model := currentModel.readAPIKeyDetails(apiKeyUserDetails)
	_, _ = logger.Debugf("Read Response: %+v", currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
		ResourceModel:   model}, nil
}

func getAPIkeyDetails(atlas *util.MongoDBClient, currentModel *Model) (*atlasSDK.ApiKeyUserDetails, *http.Response, error) {
	apiKeyRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.GetApiKey(
		context.Background(),
		*currentModel.OrgId,
		*currentModel.Id,
	)
	apiKeyUserDetails, response, err := apiKeyRequest.Execute()
	return apiKeyUserDetails, response, err
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel)
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
	apiKeyInput := atlasSDK.CreateAtlasOrganizationApiKey{
		Desc:  currentModel.Description,
		Roles: currentModel.Roles,
	}
	updateRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.UpdateApiKey(
		context.Background(),
		*currentModel.OrgId,
		*currentModel.Id,
		&apiKeyInput,
	)
	apiKeyUserDetails, response, err := updateRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		return handleError(response, UPDATE, err)
	}

	// Read response
	model := currentModel.readAPIKeyDetails(apiKeyUserDetails)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   model}, nil
}

func updateProjectAssignments(atlasClient *util.MongoDBClient, currentModel *Model, model *Model) (result interface{}, response *http.Response, err error) {
	//update projectAssignments
	newAssignments, updateAssignments, removeAssignments := getChangesInProjectAssignments(currentModel.ProjectAssignments, model.ProjectAssignments)

	// Assignment
	for i := range newAssignments {
		result, response, err = assignOrgKeyToProject(newAssignments[i], atlasClient, model.Id)
		if err != nil {
			break
		}
	}
	defer closeResponse(response)
	if err != nil {
		return
	}

	// Update
	for i := range updateAssignments {
		result, response, err = updateOrgKeyToProject(newAssignments[i], atlasClient, model.Id)
		if err != nil {
			break
		}
	}
	defer closeResponse(response)
	if err != nil {
		return
	}

	// Remove Assignment
	for i := range removeAssignments {
		result, response, err = unAssignOrgKeyToProject(newAssignments[i], atlasClient, model.Id)
		if err != nil {
			break
		}
	}
	defer closeResponse(response)
	if err != nil {
		return
	}

	return
}

func getChangesInProjectAssignments(
	inputProjectAssignments []ProjectAssignment, existingProjectAssignments []ProjectAssignment,
) (newAssignments, updateAssignments, removeAssignments []ProjectAssignment) {
	for i := range inputProjectAssignments {
		matched := false
		for e := range existingProjectAssignments {
			//Matched with existing Project
			if inputProjectAssignments[i].ProjectId == existingProjectAssignments[e].ProjectId {

				// if Roles are not matching consider for update ProjectAssignment
				if !areStringArraysEqualIgnoreOrder(inputProjectAssignments[i].Roles, existingProjectAssignments[i].Roles) {
					updateAssignments = append(updateAssignments, inputProjectAssignments[i])
				}
				matched = true
			}
		}
		// New Project Assignment
		if !matched {
			newAssignments = append(newAssignments, inputProjectAssignments[i])
		}
	}
	for e := range existingProjectAssignments {
		matched := false
		for i := range inputProjectAssignments {
			if inputProjectAssignments[i].ProjectId == existingProjectAssignments[e].ProjectId {
				matched = true
			}
		}
		if !matched {
			removeAssignments = append(removeAssignments, existingProjectAssignments[e])
		}
	}
	return
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
	deleteRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.DeleteApiKey(
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
	setup()

	modelValidation := validator.ValidateModel(ListRequiredFields, currentModel)
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
	apiKeyRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.ListApiKeys(
		context.Background(),
		*currentModel.OrgId,
	)

	if currentModel.ListOptions.PageNum != nil {
		apiKeyRequest = apiKeyRequest.PageNum(*currentModel.ListOptions.PageNum)
	}
	// For CFN Test if the no.of keys are more we have to increase the ItemsPerPage value and test
	// So that it fetches all the keys and passes create_list test case.
	if currentModel.ListOptions.ItemsPerPage != nil {
		apiKeyRequest = apiKeyRequest.ItemsPerPage(*currentModel.ListOptions.ItemsPerPage)
	}
	if currentModel.ListOptions.IncludeCount != nil {
		apiKeyRequest = apiKeyRequest.IncludeCount(*currentModel.ListOptions.IncludeCount)
	}
	pagedAPIKeysList, response, err := apiKeyRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		return handleError(response, LIST, err)
	}
	apiKeyList := pagedAPIKeysList.Results

	apiKeys := make([]interface{}, len(apiKeyList))
	for i := range apiKeyList {
		model := new(Model)
		model = model.readAPIKeyDetails(&apiKeyList[i])
		model.Profile = currentModel.Profile
		model.OrgId = currentModel.OrgId
		apiKeys[i] = model
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Completed",
		ResourceModels:  apiKeys}, nil
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

func (m *Model) readAPIKeyDetails(apikey *atlasSDK.ApiKeyUserDetails) *Model {
	model := new(Model)
	model.Profile = m.Profile
	model.OrgId = m.OrgId
	model.Id = apikey.Id
	model.Description = apikey.Desc
	model.PublicKey = apikey.PublicKey
	model.PrivateKey = apikey.PrivateKey
	var roles []string
	var projectRolesMap = map[string][]string{}
	for i := range apikey.Roles {
		// org roles
		if apikey.Roles[i].OrgId != nil && apikey.Roles[i].RoleName != nil {
			roles = append(roles, *apikey.Roles[i].RoleName)
		}
		// project roles
		if apikey.Roles[i].GroupId != nil {
			if apikey.Roles[i].RoleName != nil {
				projectRolesMap[*apikey.Roles[i].GroupId] = append(projectRolesMap[*apikey.Roles[i].GroupId], *apikey.Roles[i].RoleName)
			}
		}
	}
	for projectId, roles := range projectRolesMap {
		projectAssignment := new(ProjectAssignment)
		projectAssignment.Roles = roles
		ProjId := projectId
		projectAssignment.ProjectId = &ProjId
		model.ProjectAssignments = append(model.ProjectAssignments, *projectAssignment)
	}
	model.Roles = roles

	var links []Link
	for i := range apikey.Links {
		link := Link{
			Href: apikey.Links[i].Href,
			Rel:  apikey.Links[i].Rel,
		}
		links = append(links, link)
	}
	model.Links = links
	return model
}
