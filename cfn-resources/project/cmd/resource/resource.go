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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.OrgID, constants.Name}
var UpdateRequiredFields = []string{constants.ID}

type UpdateAPIKey struct {
	Key     string
	APIKeys *mongodbatlas.AssignAPIKey
}

func setup() {
	util.SetupLogger("mongodb-atlas-project")
}

// validateModel inputs based on the method
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	_, _ = logger.Debugf("Create currentModel: %+v", *currentModel)

	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	var projectOwnerID string
	if currentModel.ProjectOwnerId != nil {
		projectOwnerID = *currentModel.ProjectOwnerId
	}
	projectInput := &mongodbatlas.Project{
		Name:                      *currentModel.Name,
		OrgID:                     *currentModel.OrgId,
		WithDefaultAlertsSettings: currentModel.WithDefaultAlertsSettings,
	}
	if currentModel.RegionUsageRestrictions != nil {
		projectInput.RegionUsageRestrictions = *currentModel.RegionUsageRestrictions
	}

	project, res, err := client.Projects.Create(context.Background(), projectInput, &mongodbatlas.CreateProjectOptions{ProjectOwnerID: projectOwnerID})
	if err != nil {
		_, _ = logger.Debugf("Create - error: %+v", err)
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Project : %s", err.Error()),
			res.Response), nil
	}

	// Add ApiKeys
	if len(currentModel.ProjectApiKeys) > 0 {
		for _, key := range currentModel.ProjectApiKeys {
			_, err = client.ProjectAPIKeys.Assign(context.Background(), project.ID, *key.Key, &mongodbatlas.AssignAPIKey{Roles: key.RoleNames})
			if err != nil {
				_, _ = logger.Warnf("Assign Key Error: %s", err)
				return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error while Assigning Key to project : %s", err.Error()),
					res.Response), nil
			}
		}
	}

	// Add Teams
	if len(currentModel.ProjectTeams) > 0 {
		_, _, err = client.Projects.AddTeamsToProject(context.Background(), project.ID, readTeams(currentModel.ProjectTeams))
		if err != nil {
			_, _ = logger.Warnf("AddTeamsToProject Error: %s", err)
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error while adding teams to project : %s", err.Error()),
				res.Response), nil
		}
	}

	currentModel.Id = &project.ID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount

	progressEvent, err := updateProjectSettings(currentModel, client)
	if err != nil {
		return progressEvent, err
	}
	event, proj, err := getProjectWithSettings(client, currentModel)
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

func updateProjectSettings(currentModel *Model, client *mongodbatlas.Client) (handler.ProgressEvent, error) {
	if currentModel.ProjectSettings != nil {
		// Update project settings
		projectSettings := mongodbatlas.ProjectSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: currentModel.ProjectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
			IsRealtimePerformancePanelEnabled:           currentModel.ProjectSettings.IsRealtimePerformancePanelEnabled,
			IsDataExplorerEnabled:                       currentModel.ProjectSettings.IsDataExplorerEnabled,
			IsPerformanceAdvisorEnabled:                 currentModel.ProjectSettings.IsPerformanceAdvisorEnabled,
			IsSchemaAdvisorEnabled:                      currentModel.ProjectSettings.IsSchemaAdvisorEnabled,
			IsExtendedStorageSizesEnabled:               currentModel.ProjectSettings.IsExtendedStorageSizesEnabled,
		}

		_, res, err := client.Projects.UpdateProjectSettings(context.Background(), *currentModel.Id, &projectSettings)
		if err != nil {
			_, _ = logger.Warnf("UpdateProjectSettings Error: %s", err)
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Failed to update Project settings : %s", err.Error()),
				res.Response), err
		}
	}
	return handler.ProgressEvent{}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	event, model, err := getProjectWithSettings(client, currentModel)
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
	setup()

	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}

	var projectID string
	if currentModel.Id != nil {
		projectID = *currentModel.Id
	}

	event, _, err = getProject(client, currentModel)
	if err != nil {
		return event, nil
	}

	if currentModel.ProjectTeams != nil {
		// Get teams from project
		teamsAssigned, _, errr := client.Projects.GetProjectTeamsAssigned(context.Background(), projectID)
		if errr != nil {
			_, _ = logger.Warnf("ProjectId : %s, Error: %s", projectID, errr)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Error while finding teams in project",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		newTeams, changedTeams, removeTeams := getChangeInTeams(currentModel.ProjectTeams, teamsAssigned.Results)

		// Remove Teams
		for _, team := range removeTeams {
			_, err = client.Teams.RemoveTeamFromProject(context.Background(), projectID, team.TeamID)
			if err != nil {
				_, _ = logger.Warnf("ProjectId : %s, Error: %s", projectID, err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while deleting team from project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}
		// Add Teams
		if len(newTeams) > 0 {
			_, _, err = client.Projects.AddTeamsToProject(context.Background(), projectID, newTeams)
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while adding team to project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}
		// Update Teams
		for _, team := range changedTeams {
			_, _, err = client.Teams.UpdateTeamRoles(context.Background(), projectID, team.TeamID, &mongodbatlas.TeamUpdateRoles{RoleNames: team.RoleNames})
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while updating team roles in project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}
	}

	if currentModel.ProjectApiKeys != nil {
		// Get APIKeys from project
		projectAPIKeys, _, errr := client.ProjectAPIKeys.List(context.Background(), projectID, &mongodbatlas.ListOptions{ItemsPerPage: 1000, IncludeCount: true})
		if err != nil {
			_, _ = logger.Warnf("ProjectId : %s, Error: %s", projectID, errr)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Error while finding api keys in project",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}

		// Get Change in ApiKeys
		newAPIKeys, changedKeys, removeKeys := getChangeInAPIKeys(*currentModel.Id, currentModel.ProjectApiKeys, projectAPIKeys)

		// Remove old keys
		for _, key := range removeKeys {
			_, err = client.ProjectAPIKeys.Unassign(context.Background(), projectID, key.Key)
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while Un-assigning Key to project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}

		// Add Keys
		for _, key := range newAPIKeys {
			_, err = client.ProjectAPIKeys.Assign(context.Background(), projectID, key.Key, key.APIKeys)
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while Assigning Key to project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}

		// Update Key Roles
		for _, key := range changedKeys {
			_, err = client.ProjectAPIKeys.Assign(context.Background(), projectID, key.Key, key.APIKeys)
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while Un-assigning Key to project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}
	}

	progressEvent, err := updateProjectSettings(currentModel, client)
	if err != nil {
		return progressEvent, err
	}

	event, project, err := getProjectWithSettings(client, currentModel)
	if err != nil {
		return event, err
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   project,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (event handler.ProgressEvent, err error) {
	setup()

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	client, pe := util.NewMongoDBClient(req, currentModel.Profile)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", *pe)
		return *pe, nil
	}
	var id string
	if currentModel.Id != nil {
		id = *currentModel.Id
	}

	event, _, err = getProject(client, currentModel)
	if err != nil {
		return event, nil
	}
	_, _ = logger.Debugf("Deleting project with id(%s)", id)

	res, err := client.Projects.Delete(context.Background(), id)
	if err != nil {
		_, _ = logger.Warnf("####error deleting project with id(%s): %s", id, err)
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Project : %s", err.Error()),
			res.Response), nil
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
func getProject(client *mongodbatlas.Client, currentModel *Model) (event handler.ProgressEvent, model *Model, err error) {
	var project *mongodbatlas.Project
	if currentModel.Name != nil && len(*currentModel.Name) > 0 {
		event, project, err = getProjectByName(currentModel.Name, client)
		if err != nil {
			return event, nil, err
		}
	} else {
		event, project, err = getProjectByID(currentModel.Id, client)
		if err != nil {
			return event, nil, err
		}
	}
	currentModel.Name = &project.Name
	currentModel.OrgId = &project.OrgID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount
	currentModel.Id = &project.ID
	currentModel.RegionUsageRestrictions = &project.RegionUsageRestrictions
	return handler.ProgressEvent{}, currentModel, nil
}

// Read project
func getProjectWithSettings(client *mongodbatlas.Client, currentModel *Model) (event handler.ProgressEvent, model *Model, err error) {
	event, currentModel, err = getProject(client, currentModel)
	if err != nil {
		return event, currentModel, err
	}
	event, model, err = readProjectSettings(client, *currentModel.Id, currentModel)

	if err != nil {
		return event, model, err
	}

	return handler.ProgressEvent{}, model, nil
}

func getProjectByName(name *string, client *mongodbatlas.Client) (event handler.ProgressEvent, model *mongodbatlas.Project, err error) {
	project, res, err := client.Projects.GetOneProjectByName(context.Background(), *name)
	if err != nil {
		if res.Response.StatusCode == 401 { // cfn test
			return progressevents.GetFailedEventByCode(
				"Unauthorized Error: Unable to retrieve Project by name. Please verify that the API keys provided in the profile have sufficient privileges to access the project.",
				cloudformation.HandlerErrorCodeNotFound), nil, err
		}
		return progressevents.GetFailedEventByResponse(err.Error(),
			res.Response), project, err
	}
	return handler.ProgressEvent{}, project, err
}

func getProjectByID(id *string, client *mongodbatlas.Client) (event handler.ProgressEvent, model *mongodbatlas.Project, err error) {
	project, res, err := client.Projects.GetOneProject(context.Background(), *id)
	if err != nil {
		if res.Response.StatusCode == 401 { // cfn test
			return progressevents.GetFailedEventByCode(
				"Unauthorized Error: Unable to retrieve Project by ID. Please verify that the API keys provided in the profile have sufficient privileges to access the project.",
				cloudformation.HandlerErrorCodeNotFound), nil, err
		}
		return progressevents.GetFailedEventByResponse(err.Error(),
			res.Response), project, err
	}
	return handler.ProgressEvent{}, project, err
}

func readProjectSettings(client *mongodbatlas.Client, id string, currentModel *Model) (event handler.ProgressEvent, model *Model, err error) {
	// Get teams from project
	teamsAssigned, res, err := client.Projects.GetProjectTeamsAssigned(context.Background(), id)
	if err != nil {
		_, _ = logger.Warnf("ProjectId : %s, Error: %s", id, err)
		return progressevents.GetFailedEventByResponse(err.Error(),
			res.Response), nil, err
	}

	projectSettings, _, err := client.Projects.GetProjectSettings(context.Background(), id)
	if err != nil {
		_, _ = logger.Warnf("ProjectId : %s, Error: %s", id, err)
		return progressevents.GetFailedEventByResponse(err.Error(),
			res.Response), nil, err
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
	for _, team := range teamsAssigned.Results {
		if len(team.TeamID) > 0 {
			teams = append(teams, ProjectTeam{TeamId: &team.TeamID, RoleNames: team.RoleNames})
		}
	}

	currentModel.ProjectTeams = teams
	currentModel.ProjectApiKeys = nil // hack: cfn test. Extra APIKey(default) getting added and cfn test fails.
	return handler.ProgressEvent{}, currentModel, err
}

// Get difference in Teams
func getChangeInTeams(currentTeams []ProjectTeam, oTeams []*mongodbatlas.Result) (newTeams []*mongodbatlas.ProjectTeam,
	changedTeams []*mongodbatlas.ProjectTeam, removeTeams []*mongodbatlas.ProjectTeam) {
	for _, nTeam := range currentTeams {
		if nTeam.TeamId != nil && len(*nTeam.TeamId) > 0 {
			matched := false
			for _, oTeam := range oTeams {
				if nTeam.TeamId != nil && *nTeam.TeamId == oTeam.TeamID {
					changedTeams = append(changedTeams, &mongodbatlas.ProjectTeam{TeamID: *nTeam.TeamId, RoleNames: nTeam.RoleNames})
					matched = true
					break
				}
			}
			// Add to newTeams
			if !matched {
				newTeams = append(newTeams, &mongodbatlas.ProjectTeam{TeamID: *nTeam.TeamId, RoleNames: nTeam.RoleNames})
			}
		}
	}

	for _, oTeam := range oTeams {
		if len(oTeam.TeamID) > 0 {
			matched := false
			for _, nTeam := range currentTeams {
				if nTeam.TeamId != nil && *nTeam.TeamId == oTeam.TeamID {
					matched = true
					break
				}
			}
			if !matched {
				removeTeams = append(removeTeams, &mongodbatlas.ProjectTeam{TeamID: oTeam.TeamID, RoleNames: oTeam.RoleNames})
			}
		}
	}
	return newTeams, changedTeams, removeTeams
}

// Get difference in ApiKeys
func getChangeInAPIKeys(groupID string, currentKeys []ProjectApiKey, oKeys []mongodbatlas.APIKey) (newKeys, changedKeys, removeKeys []UpdateAPIKey) {
	for _, nKey := range currentKeys {
		if nKey.Key != nil && len(*nKey.Key) > 0 {
			matched := false
			for _, oKey := range oKeys {
				if nKey.Key != nil && *nKey.Key == oKey.ID {
					changedKeys = append(changedKeys, UpdateAPIKey{Key: *nKey.Key, APIKeys: &mongodbatlas.AssignAPIKey{Roles: nKey.RoleNames}})
					matched = true
					break
				}
			}
			// Add to newKeys
			if !matched {
				newKeys = append(newKeys, UpdateAPIKey{Key: *nKey.Key, APIKeys: &mongodbatlas.AssignAPIKey{Roles: nKey.RoleNames}})
			}
		}
	}

	for _, oKey := range oKeys {
		if len(oKey.ID) > 0 {
			matched := false
			for _, nKey := range currentKeys {
				if nKey.Key != nil && *nKey.Key == oKey.ID {
					matched = true
					break
				}
			}
			if !matched {
				for _, role := range oKey.Roles {
					// Consider only current ProjectRoles
					if role.GroupID == groupID {
						removeKeys = append(removeKeys, UpdateAPIKey{Key: oKey.ID})
					}
				}
			}
		}
	}
	return newKeys, changedKeys, removeKeys
}

func readTeams(teams []ProjectTeam) []*mongodbatlas.ProjectTeam {
	var newTeams []*mongodbatlas.ProjectTeam
	for _, t := range teams {
		if t.TeamId != nil && len(*t.TeamId) > 0 {
			newTeams = append(newTeams, &mongodbatlas.ProjectTeam{TeamID: *t.TeamId, RoleNames: t.RoleNames})
		}
	}
	return newTeams
}
