// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
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
	"net/http"
	"sort"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/secrets"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"

	"go.mongodb.org/atlas-sdk/v20240805001/admin"
)

var CreateRequiredFields = []string{constants.OrgID, constants.Description, constants.AwsSecretName}
var UpdateRequiredFields = []string{constants.OrgID, constants.APIUserID, constants.Description}
var ReadRequiredFields = []string{constants.OrgID, constants.APIUserID}
var DeleteRequiredFields = []string{constants.OrgID, constants.APIUserID}
var ListRequiredFields = []string{constants.OrgID}

type APIKeySecret struct {
	APIUserID  string
	PublicKey  string
	PrivateKey string
}

func setup() {
	util.SetupLogger("mongodb-atlas-api-key")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	apiKeyInput := admin.CreateAtlasOrganizationApiKey{
		Desc:  util.SafeString(currentModel.Description),
		Roles: currentModel.Roles,
	}
	apiKeyUserDetails, response, err := client.AtlasSDK.ProgrammaticAPIKeysApi.CreateApiKey(
		context.Background(),
		*currentModel.OrgId,
		&apiKeyInput,
	).Execute()

	if err != nil {
		return handleError(response, constants.CREATE, err)
	}

	currentModel.APIUserId = apiKeyUserDetails.Id

	// Save PrivateKey in AWS SecretManager
	secret := APIKeySecret{APIUserID: *currentModel.APIUserId, PublicKey: *apiKeyUserDetails.PublicKey, PrivateKey: *apiKeyUserDetails.PrivateKey}

	_, _, err = secrets.PutSecret(&req, *currentModel.AwsSecretName, secret, currentModel.Description)
	if err != nil {
		// Delete the APIKey from Atlas
		_, _ = Delete(req, prevModel, currentModel)
		response = &http.Response{StatusCode: http.StatusInternalServerError}
		return handleError(response, constants.CREATE, err)
	}
	// Assign Org APIKey to given projects i.e. projectAssignments
	if len(currentModel.ProjectAssignments) > 0 {
		for i := range currentModel.ProjectAssignments {
			handlerEvent, err := assignProjects(client, currentModel.ProjectAssignments[i], currentModel.APIUserId)
			if err != nil {
				return handlerEvent, nil
			}
		}
	}
	// writeOnly property not supposed to be in the response
	currentModel.AwsSecretName = nil

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   currentModel}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	apiKeyUserDetails, arn, response, err := getAPIkeyDetails(&req, atlas, currentModel)

	if err != nil {
		return handleError(response, constants.READ, err)
	}
	currentModel.AwsSecretArn = arn
	currentModel.readAPIKeyDetails(*apiKeyUserDetails)
	_, _ = logger.Debugf("Read Response: %+v", currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
		ResourceModel:   currentModel}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	apiKeyInput := admin.UpdateAtlasOrganizationApiKey{
		Desc:  currentModel.Description,
		Roles: &currentModel.Roles,
	}

	apiKeyUserDetails, response, err := client.AtlasSDK.ProgrammaticAPIKeysApi.UpdateApiKey(
		context.Background(),
		*currentModel.OrgId,
		*currentModel.APIUserId,
		&apiKeyInput,
	).Execute()

	if err != nil {
		return handleError(response, constants.UPDATE, err)
	}

	existingModel := Model{APIUserId: currentModel.APIUserId, OrgId: currentModel.OrgId}
	existingModel.readAPIKeyDetails(*apiKeyUserDetails)

	_, response, err = updateProjectAssignments(client, currentModel, &existingModel)

	if err != nil {
		return handleError(response, constants.UPDATE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   currentModel}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	_, response, err := client.AtlasSDK.ProgrammaticAPIKeysApi.DeleteApiKey(
		context.Background(),
		*currentModel.OrgId,
		*currentModel.APIUserId,
	).Execute()

	if err != nil {
		return handleError(response, constants.DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
		ResourceModel:   nil}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	apiKeyRequest := client.AtlasSDK.ProgrammaticAPIKeysApi.ListApiKeys(
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

	if err != nil {
		return handleError(response, constants.LIST, err)
	}

	apiKeyList := pagedAPIKeysList.GetResults()
	apiKeys := make([]interface{}, len(apiKeyList))
	for i := range apiKeyList {
		var model Model
		model.readAPIKeyDetails(apiKeyList[i])
		model.Profile = currentModel.Profile
		model.OrgId = currentModel.OrgId
		model.ListOptions = currentModel.ListOptions
		apiKeys[i] = model
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Completed",
		ResourceModels:  apiKeys}, nil
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) (handler.ProgressEvent, error) {
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

func assignProjects(client *util.MongoDBClient, project ProjectAssignment, apiUserID *string) (handler.ProgressEvent, error) {
	_, updateResponse, err := updateOrgKeyProjectRoles(project, client, apiUserID)
	if err != nil {
		return handleError(updateResponse, constants.CREATE, err)
	}
	return handler.ProgressEvent{}, err
}

func getAPIkeyDetails(req *handler.Request, client *util.MongoDBClient, currentModel *Model) (*admin.ApiKeyUserDetails, *string, *http.Response, error) {
	apiKeyRequest := client.AtlasSDK.ProgrammaticAPIKeysApi.GetApiKey(
		context.Background(),
		*currentModel.OrgId,
		*currentModel.APIUserId,
	)
	apiKeyUserDetails, response, err := apiKeyRequest.Execute()

	if err != nil {
		return apiKeyUserDetails, nil, response, err
	}
	var arn *string
	return apiKeyUserDetails, arn, response, err
}

func updateOrgKeyProjectRoles(projectAssignment ProjectAssignment, client *util.MongoDBClient, orgKeyID *string) (*admin.ApiKeyUserDetails, *http.Response, error) {
	projectAPIKeyInput := admin.UpdateAtlasProjectApiKey{
		Roles: &projectAssignment.Roles,
	}
	assignAPIRequest := client.AtlasSDK.ProgrammaticAPIKeysApi.UpdateApiKeyRoles(
		context.Background(),
		*projectAssignment.ProjectId,
		*orgKeyID,
		&projectAPIKeyInput,
	)

	return assignAPIRequest.Execute()
}

func unAssignProjectFromOrgKey(projectAssignment ProjectAssignment, client *util.MongoDBClient, orgKeyID *string) (map[string]interface{}, *http.Response, error) {
	unAssignAPIRequest := client.AtlasSDK.ProgrammaticAPIKeysApi.RemoveProjectApiKey(
		context.Background(),
		*projectAssignment.ProjectId,
		*orgKeyID,
	)

	return unAssignAPIRequest.Execute()
}

func updateProjectAssignments(atlasClient *util.MongoDBClient, currentModel *Model, existingModel *Model) (result interface{}, response *http.Response, err error) {
	// update projectAssignments
	newAssignments, updateAssignments, removeAssignments := getChangesInProjectAssignments(currentModel.ProjectAssignments, existingModel.ProjectAssignments)

	// Assignment
	for i := range newAssignments {
		result, response, err = updateOrgKeyProjectRoles(newAssignments[i], atlasClient, currentModel.APIUserId)
		if err != nil {
			break
		}
	}

	if err != nil {
		return result, response, err
	}

	// Update Project Roles
	for i := range updateAssignments {
		result, response, err = updateOrgKeyProjectRoles(updateAssignments[i], atlasClient, currentModel.APIUserId)
		if err != nil {
			break
		}
	}

	if err != nil {
		return result, response, err
	}

	// Remove Assignment
	for i := range removeAssignments {
		result, response, err = unAssignProjectFromOrgKey(removeAssignments[i], atlasClient, currentModel.APIUserId)
		if err != nil {
			break
		}
	}

	if err != nil {
		return result, response, err
	}

	return result, response, err
}

func getChangesInProjectAssignments(
	inputProjectAssignments []ProjectAssignment, existingProjectAssignments []ProjectAssignment,
) (newAssignments, updateAssignments, removeAssignments []ProjectAssignment) {
	for i := range inputProjectAssignments {
		isExistingProject := false
		for e := range existingProjectAssignments {
			// Update : Matched with existing Project
			if *inputProjectAssignments[i].ProjectId == *existingProjectAssignments[e].ProjectId {
				isExistingProject = true
				// if Roles are not matching, then consider for update ProjectAssignment
				if !areStringArraysEqualIgnoreOrder(inputProjectAssignments[i].Roles, existingProjectAssignments[i].Roles) {
					updateAssignments = append(updateAssignments, inputProjectAssignments[i])
				}
				break
			}
		}

		// New Project Assignment
		if !isExistingProject {
			newAssignments = append(newAssignments, inputProjectAssignments[i])
		}
	}
	for e := range existingProjectAssignments {
		matchedWithExisting := false
		for i := range inputProjectAssignments {
			if *inputProjectAssignments[i].ProjectId == *existingProjectAssignments[e].ProjectId {
				matchedWithExisting = true
				break
			}
		}
		// Removable Assignments
		if !matchedWithExisting {
			removeAssignments = append(removeAssignments, existingProjectAssignments[e])
		}
	}
	return newAssignments, updateAssignments, removeAssignments
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

func (model *Model) readAPIKeyDetails(apikey admin.ApiKeyUserDetails) Model {
	model.APIUserId = apikey.Id
	model.Description = apikey.Desc
	model.PublicKey = apikey.PublicKey
	model.PrivateKey = apikey.PrivateKey
	var roles []string
	var projectAssignments []ProjectAssignment
	var projectRolesMap = map[string][]string{}
	apiKeyRoles := apikey.GetRoles()
	for i := range apiKeyRoles {
		// org roles
		if apiKeyRoles[i].OrgId != nil && apiKeyRoles[i].RoleName != nil {
			roles = append(roles, *apiKeyRoles[i].RoleName)
		}
		// project roles
		if apiKeyRoles[i].GroupId != nil {
			if apiKeyRoles[i].RoleName != nil {
				projectRolesMap[*apiKeyRoles[i].GroupId] = append(projectRolesMap[*apiKeyRoles[i].GroupId], *apiKeyRoles[i].RoleName)
			}
		}
	}
	for projectID, roles := range projectRolesMap {
		projectAssignment := new(ProjectAssignment)
		projectAssignment.Roles = roles
		ID := projectID
		projectAssignment.ProjectId = &ID
		projectAssignments = append(projectAssignments, *projectAssignment)
	}
	model.Roles = roles
	model.ProjectAssignments = projectAssignments

	return *model
}
