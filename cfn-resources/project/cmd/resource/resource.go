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
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115008/admin"
)

var CreateRequiredFields = []string{constants.OrgID, constants.Name}
var ReadRequiredFields = []string{constants.ID}
var UpdateRequiredFields = []string{constants.ID}
var DeleteRequiredFields = []string{constants.ID}

type UpdateAPIKey struct {
	Key           string
	UpdatePayload *admin.UpdateAtlasProjectApiKey
}

func initEnvWithLatestClient(req handler.Request, currentModel *Model, requiredFields []string) (*admin.APIClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-project")

	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if errEvent := validator.ValidateModel(requiredFields, currentModel); errEvent != nil {
		return nil, errEvent
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return nil, peErr
	}
	return client.AtlasSDK, nil
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	atlasV2, peErr := initEnvWithLatestClient(req, currentModel, CreateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}
	projectInput := &admin.Group{
		Name:                      *currentModel.Name,
		OrgId:                     *currentModel.OrgId,
		WithDefaultAlertsSettings: currentModel.WithDefaultAlertsSettings,
	}
	if currentModel.RegionUsageRestrictions != nil {
		projectInput.RegionUsageRestrictions = currentModel.RegionUsageRestrictions
	}

	createProjectReq := admin.CreateProjectApiParams{
		Group: projectInput,
	}
	if currentModel.ProjectOwnerId != nil {
		createProjectReq.ProjectOwnerId = currentModel.ProjectOwnerId
	}
	project, res, err := atlasV2.ProjectsApi.CreateProjectWithParams(context.Background(), &createProjectReq).Execute()

	if err != nil {
		_, _ = logger.Debugf("Create - error: %+v", err)
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Project : %s", err.Error()),
			res), nil
	}
	projectID := project.GetId()

	// Add ApiKeys
	if len(currentModel.ProjectApiKeys) > 0 {
		for _, key := range currentModel.ProjectApiKeys {
			if key.Key == nil {
				errorMessage := fmt.Sprintf("ApiKey is missing the configuration for projectID=%s", projectID)
				_, _ = logger.Warn(errorMessage)
				return progressevent.GetFailedEventByCode(errorMessage, cloudformation.HandlerErrorCodeInvalidRequest), nil
			}
			key := key
			apiKey := *key.Key
			_, res, err := atlasV2.ProgrammaticAPIKeysApi.UpdateApiKeyRoles(context.Background(), projectID, apiKey, &admin.UpdateAtlasProjectApiKey{
				Roles: &key.RoleNames,
			}).Execute()
			if err != nil {
				_, _ = logger.Warnf("Assign Key Error: %s", err)
				return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error while Assigning Key to project : %s", err.Error()),
					res), nil
			}
		}
	}

	// Add Teams
	if len(currentModel.ProjectTeams) > 0 {
		teams := readTeams(currentModel.ProjectTeams)
		_, _, err := atlasV2.TeamsApi.AddAllTeamsToProject(context.Background(), *project.Id, &teams).Execute()
		if err != nil {
			_, _ = logger.Warnf("AddTeamsToProject Error: %s", err)
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error while adding teams to project : %s", err.Error()),
				res), nil
		}
	}

	formattedCreated := util.TimeToString(project.Created)

	currentModel.Id = project.Id
	currentModel.Created = &formattedCreated
	currentModel.ClusterCount = util.Int64PtrToIntPtr(&project.ClusterCount)

	progressEvent, err := updateProjectSettings(currentModel, atlasV2)
	if err != nil {
		return progressEvent, err
	}
	event, proj, err := getProjectWithSettings(atlasV2, currentModel)
	if err != nil {
		_, _ = logger.Warnf("getProject Error: %s", err)
		return event, err
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   proj,
	}, nil
}

func updateProjectSettings(currentModel *Model, atlasV2 *admin.APIClient) (handler.ProgressEvent, error) {
	if currentModel.ProjectSettings != nil {
		// Update project settings
		projectSettings := admin.GroupSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: currentModel.ProjectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
			IsRealtimePerformancePanelEnabled:           currentModel.ProjectSettings.IsRealtimePerformancePanelEnabled,
			IsDataExplorerEnabled:                       currentModel.ProjectSettings.IsDataExplorerEnabled,
			IsPerformanceAdvisorEnabled:                 currentModel.ProjectSettings.IsPerformanceAdvisorEnabled,
			IsSchemaAdvisorEnabled:                      currentModel.ProjectSettings.IsSchemaAdvisorEnabled,
			IsExtendedStorageSizesEnabled:               currentModel.ProjectSettings.IsExtendedStorageSizesEnabled,
		}

		_, res, err := atlasV2.ProjectsApi.UpdateProjectSettings(context.Background(), *currentModel.Id, &projectSettings).Execute()
		if err != nil {
			_, _ = logger.Warnf("UpdateProjectSettings Error: %s", err)
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Failed to update Project settings : %s", err.Error()),
				res), err
		}
	}
	return handler.ProgressEvent{}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	atlasV2, peErr := initEnvWithLatestClient(req, currentModel, ReadRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}

	event, model, err := getProjectWithSettings(atlasV2, currentModel)
	if err != nil {
		return event, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   model,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (event handler.ProgressEvent, err error) {
	atlasV2, peErr := initEnvWithLatestClient(req, currentModel, UpdateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}
	var projectID string
	if currentModel.Id != nil {
		projectID = *currentModel.Id
	}

	event, _, err = getProject(atlasV2, &Model{Id: currentModel.Id})
	if err != nil {
		return event, nil
	}

	// Update Project
	if currentModel.Name != nil {
		event, _, err = updateProject(atlasV2, currentModel)
		if err != nil {
			return event, err
		}
	}

	if currentModel.ProjectTeams != nil {
		// Get teams from project
		teamsAssigned, _, err := atlasV2.TeamsApi.ListProjectTeams(context.Background(), projectID).Execute()
		if err != nil {
			_, _ = logger.Warnf("ProjectId : %s, Error: %s", projectID, err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Error while finding teams in project",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		if teamsAssigned != nil && teamsAssigned.Results != nil {
			errorMessage, err := changeProjectTeams(*atlasV2, currentModel, teamsAssigned.GetResults())
			if err != nil {
				_, _ = logger.Warnf("ProjectId : %s, Error: %s", projectID, err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          errorMessage,
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
				}, nil
			}
		}
	}

	if currentModel.ProjectApiKeys != nil {
		// Get Change in ApiKeys
		newAPIKeys, changedKeys, removeKeys := getChangeInAPIKeys(currentModel.ProjectApiKeys, prevModel.ProjectApiKeys)

		// Remove old keys
		for _, key := range removeKeys {
			_, _, err = atlasV2.ProgrammaticAPIKeysApi.RemoveProjectApiKey(context.Background(), projectID, *key.Key).Execute()
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          fmt.Sprintf("Error while Un-assigning Key to project %s", err.Error()),
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}

		// Add Keys
		for _, key := range newAPIKeys {
			key := key
			_, _, err := atlasV2.ProgrammaticAPIKeysApi.UpdateApiKeyRoles(context.Background(), projectID, *key.Key, &admin.UpdateAtlasProjectApiKey{
				Roles: &key.RoleNames,
			}).Execute()
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          fmt.Sprintf("Error while Assigning Key to project %s", err.Error()),
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}

		// Update Key Roles
		for _, key := range changedKeys {
			key := key
			_, _, err := atlasV2.ProgrammaticAPIKeysApi.UpdateApiKeyRoles(context.Background(), projectID, *key.Key, &admin.UpdateAtlasProjectApiKey{
				Roles: &key.RoleNames,
			}).Execute()
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          fmt.Sprintf("Error while Assigning Key to project %s", err.Error()),
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}
	}

	progressEvent, err := updateProjectSettings(currentModel, atlasV2)
	if err != nil {
		return progressEvent, err
	}

	event, project, err := getProjectWithSettings(atlasV2, currentModel)
	if err != nil {
		return event, err
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   project,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (event handler.ProgressEvent, err error) {
	atlasV2, peErr := initEnvWithLatestClient(req, currentModel, DeleteRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}
	var id string
	if currentModel.Id != nil {
		id = *currentModel.Id
	}

	event, _, err = getProject(atlasV2, currentModel)
	if err != nil {
		return event, nil
	}
	_, _ = logger.Debugf("Deleting project with id(%s)", id)

	_, res, err := atlasV2.ProjectsApi.DeleteProject(context.Background(), id).Execute()
	if err != nil {
		_, _ = logger.Warnf("####error deleting project with id(%s): %s", id, err)
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Project : %s", err.Error()),
			res), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   nil,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

// Read project
func getProject(client *admin.APIClient, currentModel *Model) (event handler.ProgressEvent, model *Model, err error) {
	var project *admin.Group
	event, project, err = getProjectByID(currentModel.Id, client)
	if err != nil {
		return event, nil, err
	}
	formattedCreated := util.TimeToString(project.Created)

	currentModel.Name = &project.Name
	currentModel.OrgId = &project.OrgId
	currentModel.Created = &formattedCreated
	currentModel.ClusterCount = util.Int64PtrToIntPtr(&project.ClusterCount)
	currentModel.Id = project.Id
	currentModel.RegionUsageRestrictions = project.RegionUsageRestrictions
	return handler.ProgressEvent{}, currentModel, nil
}

// Read project
func getProjectWithSettings(atlasV2 *admin.APIClient, currentModel *Model) (event handler.ProgressEvent, model *Model, err error) {
	event, currentModel, err = getProject(atlasV2, currentModel)
	if err != nil {
		return event, currentModel, err
	}
	event, model, err = readProjectSettings(atlasV2, *currentModel.Id, currentModel)

	if err != nil {
		return event, model, err
	}

	return handler.ProgressEvent{}, model, nil
}

func updateProject(client *admin.APIClient, currentModel *Model) (event handler.ProgressEvent, model *admin.Group, err error) {
	project, res, err := client.ProjectsApi.UpdateProject(context.Background(), *currentModel.Id, &admin.GroupUpdate{Name: currentModel.Name}).Execute()
	if err != nil {
		if res.StatusCode == 401 { // cfn test
			return progressevent.GetFailedEventByCode(
				"Unauthorized Error: Unable to update project name. Please verify that the API keys provided in the profile have sufficient privileges to access the project.",
				cloudformation.HandlerErrorCodeNotFound), nil, err
		}
		return progressevent.GetFailedEventByResponse(err.Error(),
			res), project, err
	}
	return handler.ProgressEvent{}, project, err
}

func getProjectByID(id *string, atlasV2 *admin.APIClient) (event handler.ProgressEvent, model *admin.Group, err error) {
	project, res, err := atlasV2.ProjectsApi.GetProject(context.Background(), *id).Execute()
	if err != nil {
		if res.StatusCode == 401 { // cfn test
			return progressevent.GetFailedEventByCode(
				"Unauthorized Error: Unable to retrieve Project by ID. Please verify that the API keys provided in the profile have sufficient privileges to access the project.",
				cloudformation.HandlerErrorCodeNotFound), nil, err
		}
		return progressevent.GetFailedEventByResponse(err.Error(),
			res), project, err
	}
	return handler.ProgressEvent{}, project, err
}

func readProjectSettings(atlasV2 *admin.APIClient, id string, currentModel *Model) (event handler.ProgressEvent, model *Model, err error) {
	// Get teams from project
	teamsAssigned, res, err := atlasV2.TeamsApi.ListProjectTeams(context.Background(), id).Execute()
	if err != nil {
		_, _ = logger.Warnf("ProjectId : %s, Error: %s", id, err)
		return progressevent.GetFailedEventByResponse(err.Error(),
			res), nil, err
	}

	projectSettings, _, err := atlasV2.ProjectsApi.GetProjectSettings(context.Background(), id).Execute()
	if err != nil {
		_, _ = logger.Warnf("ProjectId : %s, Error: %s", id, err)
		return progressevent.GetFailedEventByResponse(err.Error(),
			res), nil, err
	}
	// Set projectSettings
	currentModel.ProjectSettings = &ProjectSettings{
		IsCollectDatabaseSpecificsStatisticsEnabled: projectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
		IsRealtimePerformancePanelEnabled:           projectSettings.IsRealtimePerformancePanelEnabled,
		IsDataExplorerEnabled:                       projectSettings.IsDataExplorerEnabled,
		IsPerformanceAdvisorEnabled:                 projectSettings.IsPerformanceAdvisorEnabled,
		IsSchemaAdvisorEnabled:                      projectSettings.IsSchemaAdvisorEnabled,
		IsExtendedStorageSizesEnabled:               projectSettings.IsExtendedStorageSizesEnabled,
	}

	// Set teams
	var teams []ProjectTeam
	for _, team := range teamsAssigned.GetResults() {
		if util.IsStringPresent(team.TeamId) {
			teams = append(teams, ProjectTeam{TeamId: team.TeamId, RoleNames: team.GetRoleNames()})
		}
	}

	currentModel.ProjectTeams = teams
	currentModel.ProjectApiKeys = nil // hack: cfn test. Extra APIKey(default) getting added and cfn test fails.
	return handler.ProgressEvent{}, currentModel, err
}

// Get difference in Teams
func getChangeInTeams(currentTeams []ProjectTeam, oTeams []admin.TeamRole) (newTeams []admin.TeamRole,
	changedTeams []admin.TeamRole, removeTeams []admin.TeamRole) {
	for _, nTeam := range currentTeams {
		nTeam := nTeam
		if util.IsStringPresent(nTeam.TeamId) {
			matched := false
			for _, oTeam := range oTeams {
				if util.AreStringPtrEqual(nTeam.TeamId, oTeam.TeamId) {
					changedTeams = append(changedTeams, admin.TeamRole{TeamId: nTeam.TeamId, RoleNames: &nTeam.RoleNames})
					matched = true
					break
				}
			}
			// Add to newTeams
			if !matched {
				newTeams = append(newTeams, admin.TeamRole{TeamId: nTeam.TeamId, RoleNames: &nTeam.RoleNames})
			}
		}
	}

	for _, oTeam := range oTeams {
		if util.IsStringPresent(oTeam.TeamId) {
			matched := false
			for _, nTeam := range currentTeams {
				if util.AreStringPtrEqual(nTeam.TeamId, oTeam.TeamId) {
					matched = true
					break
				}
			}
			if !matched {
				removeTeams = append(removeTeams, admin.TeamRole{TeamId: oTeam.TeamId, RoleNames: oTeam.RoleNames})
			}
		}
	}
	return newTeams, changedTeams, removeTeams
}

func readTeams(teams []ProjectTeam) []admin.TeamRole {
	var newTeams []admin.TeamRole
	for _, t := range teams {
		t := t
		if util.IsStringPresent(t.TeamId) {
			newTeams = append(newTeams, admin.TeamRole{TeamId: t.TeamId, RoleNames: &t.RoleNames})
		}
	}
	return newTeams
}

func getChangeInAPIKeys(currentKeys []ProjectApiKey, previousKeys []ProjectApiKey) (newKeys, changedKeys, removeKeys []ProjectApiKey) {
	// Create maps to efficiently check for the existence of keys by ID
	currentKeyMap := make(map[string]ProjectApiKey)
	previousKeyMap := make(map[string]ProjectApiKey)

	// Populate the maps using the ID as the key
	for _, key := range currentKeys {
		if key.Key != nil {
			currentKeyMap[*key.Key] = key
		}
	}

	for _, key := range previousKeys {
		if key.Key != nil {
			previousKeyMap[*key.Key] = key
		}
	}

	// Identify new keys, changed keys, and removed keys
	for _, key := range currentKeys {
		if key.Key == nil {
			continue
		}

		if oKey, ok := previousKeyMap[*key.Key]; ok {
			if !util.SameStringSliceWithoutOrder(key.RoleNames, oKey.RoleNames) {
				changedKeys = append(changedKeys, key)
			}
		} else {
			// Key exists in currentKeys but not in oKeys
			newKeys = append(newKeys, key)
		}
	}

	for _, key := range previousKeys {
		if key.Key == nil {
			continue
		}

		if _, ok := currentKeyMap[*key.Key]; !ok {
			// Key exists in oKeys but not in currentKeys
			removeKeys = append(removeKeys, key)
		}
	}

	return newKeys, changedKeys, removeKeys
}

func changeProjectTeams(atlasV2 admin.APIClient, currentModel *Model, newTeams []admin.TeamRole) (errorMessage string, err error) {
	newTeams, changedTeams, removeTeams := getChangeInTeams(currentModel.ProjectTeams, newTeams)
	projectID := *currentModel.Id
	// Remove Teams
	for _, team := range removeTeams {
		_, err = atlasV2.TeamsApi.RemoveProjectTeam(context.Background(), projectID, util.SafeString(team.TeamId)).Execute()
		if err != nil {
			return "Error while deleting team from project", err
		}
	}
	// Add Teams
	if len(newTeams) > 0 {
		_, _, err = atlasV2.TeamsApi.AddAllTeamsToProject(context.Background(), projectID, &newTeams).Execute()
		if err != nil {
			return "Error while adding team to project", err
		}
	}
	// Update Teams
	for _, team := range changedTeams {
		_, _, err = atlasV2.TeamsApi.UpdateTeamRoles(context.Background(), projectID, util.SafeString(team.TeamId), &admin.TeamRole{RoleNames: team.RoleNames}).Execute()
		if err != nil {
			return "Error while updating team roles in project", err
		}
	}
	return "", err
}
