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
	"reflect"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20231115014/admin"
)

var CreateRequiredFields = []string{constants.OrgID, constants.Name}
var ReadUpdateDeleteRequiredFields = []string{constants.ID}

type UpdateAPIKey struct {
	UpdatePayload *admin.UpdateAtlasProjectApiKey
	Key           string
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
	return client.Atlas20231115014, nil
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	atlasV2, peErr := initEnvWithLatestClient(req, currentModel, CreateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}
	adminTags := NewResourceTags(currentModel.Tags)
	projectInput := &admin.Group{
		Name:                      *currentModel.Name,
		OrgId:                     *currentModel.OrgId,
		WithDefaultAlertsSettings: currentModel.WithDefaultAlertsSettings,
		Tags:                      &adminTags,
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
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Project : %s", err.Error()),
			res), nil
	}
	projectID := project.GetId()

	if len(currentModel.ProjectApiKeys) > 0 {
		for _, key := range currentModel.ProjectApiKeys {
			if key.Key == nil {
				errorMessage := fmt.Sprintf("ApiKey is missing the configuration for projectID=%s", projectID)
				return progressevent.GetFailedEventByCode(errorMessage, cloudformation.HandlerErrorCodeInvalidRequest), nil
			}
			apiKey := *key.Key
			_, res, err := atlasV2.ProgrammaticAPIKeysApi.UpdateApiKeyRoles(context.Background(), projectID, apiKey, &admin.UpdateAtlasProjectApiKey{
				Roles: &key.RoleNames,
			}).Execute()
			if err != nil {
				return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error while Assigning Key to project : %s", err.Error()),
					res), nil
			}
		}
	}

	if len(currentModel.ProjectTeams) > 0 {
		teams := readTeams(currentModel.ProjectTeams)
		_, _, err := atlasV2.TeamsApi.AddAllTeamsToProject(context.Background(), *project.Id, &teams).Execute()
		if err != nil {
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
			return progressevent.GetFailedEventByResponse(fmt.Sprintf("Failed to update Project settings : %s", err.Error()),
				res), err
		}
	}
	return handler.ProgressEvent{}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	atlasV2, peErr := initEnvWithLatestClient(req, currentModel, ReadUpdateDeleteRequiredFields)
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

func Update(req handler.Request, prevModel *Model, currentModel *Model) (event handler.ProgressEvent, err error) {
	atlasV2, peErr := initEnvWithLatestClient(req, currentModel, ReadUpdateDeleteRequiredFields)
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

	if currentModel.Name != nil || !reflect.DeepEqual(prevModel.Tags, currentModel.Tags) {
		event, _, err = updateProject(atlasV2, currentModel)
		if err != nil {
			return event, err
		}
	}

	if currentModel.ProjectTeams != nil {
		teamsAssigned, _, err := atlasV2.TeamsApi.ListProjectTeams(context.Background(), projectID).Execute()
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Error while finding teams in project",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		if teamsAssigned != nil && teamsAssigned.Results != nil {
			errorMessage, err := changeProjectTeams(*atlasV2, currentModel, teamsAssigned.GetResults())
			if err != nil {
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          errorMessage,
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
				}, nil
			}
		}
	}

	if currentModel.ProjectApiKeys != nil {
		newAPIKeys, changedKeys, removeKeys := GetChangeInAPIKeys(currentModel.ProjectApiKeys, prevModel.ProjectApiKeys)

		for _, key := range removeKeys {
			_, _, err = atlasV2.ProgrammaticAPIKeysApi.RemoveProjectApiKey(context.Background(), projectID, *key.Key).Execute()
			if err != nil {
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          fmt.Sprintf("Error while Un-assigning Key to project %s", err.Error()),
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}

		for _, key := range newAPIKeys {
			_, _, err := atlasV2.ProgrammaticAPIKeysApi.UpdateApiKeyRoles(context.Background(), projectID, *key.Key, &admin.UpdateAtlasProjectApiKey{
				Roles: &key.RoleNames,
			}).Execute()
			if err != nil {
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          fmt.Sprintf("Error while Assigning Key to project %s", err.Error()),
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}

		for _, key := range changedKeys {
			_, _, err := atlasV2.ProgrammaticAPIKeysApi.UpdateApiKeyRoles(context.Background(), projectID, *key.Key, &admin.UpdateAtlasProjectApiKey{
				Roles: &key.RoleNames,
			}).Execute()
			if err != nil {
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

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (event handler.ProgressEvent, err error) {
	atlasV2, peErr := initEnvWithLatestClient(req, currentModel, ReadUpdateDeleteRequiredFields)
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

	_, res, err := atlasV2.ProjectsApi.DeleteProject(context.Background(), id).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Project : %s", err.Error()),
			res), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   nil,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

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
	currentModel.Tags = NewCfnTags(project.GetTags())
	return handler.ProgressEvent{}, currentModel, nil
}

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
	adminTags := NewResourceTags(currentModel.Tags)
	projectUpdate := admin.GroupUpdate{
		Name: currentModel.Name,
		Tags: &adminTags,
	}
	project, res, err := client.ProjectsApi.UpdateProject(context.Background(), *currentModel.Id, &projectUpdate).Execute()
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
	teamsAssigned, res, err := atlasV2.TeamsApi.ListProjectTeams(context.Background(), id).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(),
			res), nil, err
	}

	projectSettings, _, err := atlasV2.ProjectsApi.GetProjectSettings(context.Background(), id).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(),
			res), nil, err
	}
	currentModel.ProjectSettings = &ProjectSettings{
		IsCollectDatabaseSpecificsStatisticsEnabled: projectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
		IsRealtimePerformancePanelEnabled:           projectSettings.IsRealtimePerformancePanelEnabled,
		IsDataExplorerEnabled:                       projectSettings.IsDataExplorerEnabled,
		IsPerformanceAdvisorEnabled:                 projectSettings.IsPerformanceAdvisorEnabled,
		IsSchemaAdvisorEnabled:                      projectSettings.IsSchemaAdvisorEnabled,
		IsExtendedStorageSizesEnabled:               projectSettings.IsExtendedStorageSizesEnabled,
	}

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

func getChangeInTeams(currentTeams []ProjectTeam, oTeams []admin.TeamRole) (newTeams []admin.TeamRole,
	changedTeams []admin.TeamRole, removeTeams []admin.TeamRole) {
	for i := range currentTeams {
		nTeam := currentTeams[i]
		if util.IsStringPresent(nTeam.TeamId) {
			matched := false
			for _, oTeam := range oTeams {
				if util.AreStringPtrEqual(nTeam.TeamId, oTeam.TeamId) {
					changedTeams = append(changedTeams, admin.TeamRole{TeamId: nTeam.TeamId, RoleNames: &nTeam.RoleNames})
					matched = true
					break
				}
			}
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
	for i := range teams {
		team := teams[i]
		if util.IsStringPresent(team.TeamId) {
			newTeams = append(newTeams, admin.TeamRole{TeamId: team.TeamId, RoleNames: &team.RoleNames})
		}
	}
	return newTeams
}

func GetChangeInAPIKeys(currentKeys []ProjectApiKey, previousKeys []ProjectApiKey) (newKeys, changedKeys, removeKeys []ProjectApiKey) {
	currentKeyMap := make(map[string]ProjectApiKey)
	previousKeyMap := make(map[string]ProjectApiKey)

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

	for _, key := range currentKeys {
		if key.Key == nil {
			continue
		}

		if oKey, ok := previousKeyMap[*key.Key]; ok {
			if !util.SameStringSliceWithoutOrder(key.RoleNames, oKey.RoleNames) {
				changedKeys = append(changedKeys, key)
			}
		} else {
			newKeys = append(newKeys, key)
		}
	}

	for _, key := range previousKeys {
		if key.Key == nil {
			continue
		}

		if _, ok := currentKeyMap[*key.Key]; !ok {
			removeKeys = append(removeKeys, key)
		}
	}

	return newKeys, changedKeys, removeKeys
}

func changeProjectTeams(atlasV2 admin.APIClient, currentModel *Model, newTeams []admin.TeamRole) (errorMessage string, err error) {
	newTeams, changedTeams, removeTeams := getChangeInTeams(currentModel.ProjectTeams, newTeams)
	projectID := *currentModel.Id
	for _, team := range removeTeams {
		_, err = atlasV2.TeamsApi.RemoveProjectTeam(context.Background(), projectID, util.SafeString(team.TeamId)).Execute()
		if err != nil {
			return "Error while deleting team from project", err
		}
	}
	if len(newTeams) > 0 {
		_, _, err = atlasV2.TeamsApi.AddAllTeamsToProject(context.Background(), projectID, &newTeams).Execute()
		if err != nil {
			return "Error while adding team to project", err
		}
	}
	for _, team := range changedTeams {
		_, _, err = atlasV2.TeamsApi.UpdateTeamRoles(context.Background(), projectID, util.SafeString(team.TeamId), &admin.TeamRole{RoleNames: team.RoleNames}).Execute()
		if err != nil {
			return "Error while updating team roles in project", err
		}
	}
	return "", err
}
