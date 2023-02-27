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
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.OrgID}
var ReadRequiredFields = []string{constants.OrgID, constants.TeamID}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(CreateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, errors.New("required field not found")
	}

	client, handlerError := util.NewMongoDBClient(req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	// API call to create team
	teamID := cast.ToString(currentModel.TeamId)
	orgID := cast.ToString(currentModel.OrgId)
	projectID := cast.ToString(currentModel.ProjectId)
	if teamID == "" {
		// create new team in organization
		teamRequest := mongodbatlas.Team{
			Name:      cast.ToString(currentModel.Name),
			Usernames: currentModel.Usernames,
		}
		teamResponse, resp, err := client.Teams.Create(context.Background(), orgID, &teamRequest)
		if err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("unable to create team %v", err), resp.Response), nil
		}
		teamID = teamResponse.ID
		currentModel = convertTeamToModel(teamResponse, currentModel)
	}

	// add existing team or newly created team to project if project id exist in the request
	if projectID != "" && len(currentModel.RoleNames) > 0 {
		createRequest := []*mongodbatlas.ProjectTeam{{
			TeamID:    teamID,
			RoleNames: currentModel.RoleNames,
		}}
		_, _, err := client.Projects.AddTeamsToProject(context.Background(), projectID, createRequest)
		if err != nil {
			_, _ = logger.Warnf("error adding Team(%s) to project(%s): reason : %v", teamID, projectID, err)
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(ReadRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	client, handlerError := util.NewMongoDBClient(req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	// API call to read snapshot to read using ID field
	teamID := cast.ToString(currentModel.TeamId)
	orgID := cast.ToString(currentModel.OrgId)
	teamName := cast.ToString(currentModel.Name)
	var team *mongodbatlas.Team
	var resp *mongodbatlas.Response
	var err error
	// get team by id or name
	if teamID != "" {
		team, resp, err = client.Teams.Get(context.Background(), orgID, teamID)
	} else if teamName != "" {
		// get team by name
		team, resp, err = client.Teams.GetOneTeamByName(context.Background(), orgID, teamName)
	}

	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	currentModel = convertTeamToModel(team, currentModel)

	// API call to get all users assigned
	users, _, err := client.Teams.GetTeamUsersAssigned(context.Background(), orgID, *currentModel.TeamId)
	if err != nil {
		_, _ = logger.Warnf("error getting Team user information: %v", err)
	}
	if users != nil {
		var userNames []string
		var userList []AtlasUser
		for ind := range users {
			userNames = append(userNames, users[ind].Username)
			userList = append(userList, flattenUser(users[ind]))
		}
		currentModel.Usernames = userNames
		currentModel.Users = userList
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

func flattenUser(user mongodbatlas.AtlasUser) AtlasUser {
	return AtlasUser{
		Country:      &user.Country,
		EmailAddress: &user.EmailAddress,
		FirstName:    &user.FirstName,
		Id:           &user.ID,
		LastName:     &user.LastName,
		MobileNumber: &user.MobileNumber,
		Password:     &user.Password,
		Roles:        flattenRole(user.Roles),
		TeamIds:      user.TeamIds,
		Username:     &user.Username,
	}
}
func flattenRole(role []mongodbatlas.AtlasRole) []AtlasRole {
	var modelRole []AtlasRole
	if role == nil {
		return modelRole
	}
	for ind := range role {
		pe := AtlasRole{
			RoleName:  &role[ind].RoleName,
			ProjectId: &role[ind].GroupID,
			OrgId:     &role[ind].OrgID,
		}
		modelRole = append(modelRole, pe)
	}
	return modelRole
}
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	// Validate required fields in the request
	if modelValidation := validateModel(ReadRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, errors.New("required field not found")
	}

	client, handlerError := util.NewMongoDBClient(req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	if !isExist(client, currentModel) {
		_, _ = logger.Debugf("error getting Team information: %s", *currentModel.TeamId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	// API call
	team, res, err := client.Teams.Get(context.Background(), *currentModel.OrgId, *currentModel.TeamId)
	if err != nil {
		_, _ = logger.Debugf("error getting Team information: %s", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	teamID := cast.ToString(currentModel.TeamId)
	orgID := cast.ToString(currentModel.OrgId)
	teamName := cast.ToString(currentModel.Name)
	projectID := cast.ToString(currentModel.ProjectId)

	// add existing team or newly created team to project if project id exist in the request
	if projectID != "" && len(currentModel.RoleNames) > 0 {
		createRequest := []*mongodbatlas.ProjectTeam{{
			TeamID:    teamID,
			RoleNames: currentModel.RoleNames,
		}}
		_, _, err := client.Projects.AddTeamsToProject(context.Background(), projectID, createRequest)
		if err != nil {
			_, _ = logger.Warnf("error adding Team(%s) to project(%s): reason : %v", teamID, projectID, err)
		}
	}

	// rename the team
	if team.Name != teamName {
		_, _, err := client.Teams.Rename(context.Background(), orgID, teamID, teamName)
		if err != nil {
			_, _ = logger.Warnf("error updating Team information: %v", err)
		}
	}

	// add/remove user to/from teams
	if currentModel.Usernames != nil {
		// get the current  users list for the team
		users, _, err := client.Teams.GetTeamUsersAssigned(context.Background(), orgID, teamID)
		if err != nil {
			_, _ = logger.Warnf("get assigned user to team -error (%v)", err)
		}
		usernames := currentModel.Usernames
		var newUsers []string
		for ind := range usernames {
			currentUser, isExistingUser := isUserExist(users, usernames[ind])

			if isExistingUser {
				// remove user from team
				_, err := client.Teams.RemoveUserToTeam(context.Background(), orgID, teamID, currentUser.ID)
				if err != nil {
					_, _ = logger.Warnf("remove user(%s) from Team(%s) -error (%v) \n", currentUser.ID, teamID, err)
				}
			} else {
				// add user to team
				user, _, err := client.AtlasUsers.GetByName(context.Background(), usernames[ind])
				if err != nil {
					_, _ = logger.Warnf("Error reading user (%s)  with error (%v) \n", usernames[ind], err)
				}
				// if the user exists, we will store its ID so that we can save as user list later
				if user != nil {
					newUsers = append(newUsers, user.ID)
				}
			}
		}
		// save all new users
		_, _, err = client.Teams.AddUsersToTeam(context.Background(), orgID, teamID, newUsers)
		if err != nil {
			_, _ = logger.Warnf("team -Add users error (%+v) \n", err)
		}
	}

	// update roles to team
	roleNames := currentModel.RoleNames
	if len(roleNames) > 0 && currentModel.ProjectId != nil {
		teamRequest := &mongodbatlas.TeamUpdateRoles{RoleNames: roleNames}
		_, _, err = client.Teams.UpdateTeamRoles(context.Background(), projectID, teamID, teamRequest)
		if err != nil {
			_, _ = logger.Warnf("update role to team  error (%+v) \n", err)
		}
	}
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("List Teams  Request :%+v", currentModel)

	client, handlerError := util.NewMongoDBClient(req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	// Create Atlas API Request Object
	orgID := cast.ToString(currentModel.OrgId)
	projectID := cast.ToString(currentModel.ProjectId)
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}
	var models []interface{}
	var resp *mongodbatlas.Response
	var err error
	// API call to get teams for project id
	if projectID != "" {
		var teamsAssigned *mongodbatlas.TeamsAssigned
		teamsAssigned, resp, err = client.Projects.GetProjectTeamsAssigned(context.Background(), projectID)

		if err != nil {
			return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
		}

		teamsProjectList := teamsAssigned.Results
		for i := 0; i < len(teamsProjectList); i++ {
			models = append(models, convertProjectTeamToModel(teamsProjectList[i]))
		}
	} else {
		// API call to get teams from organization
		var teams []mongodbatlas.Team
		teams, resp, err = client.Teams.List(context.Background(), orgID, params)

		if err != nil {
			return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
		}

		for i := 0; i < len(teams); i++ {
			models = append(models, convertTeamToModel(&teams[i], nil))
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("Delete Team  Request() :%+v", currentModel)

	client, handlerError := util.NewMongoDBClient(req, currentModel.Profile)
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	if !isExist(client, currentModel) {
		_, _ = logger.Debugf("error getting Team information: %s", *currentModel.TeamId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	if currentModel.ProjectId != nil {
		if err := removeFromProject(client, currentModel); err != nil {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Unable to Delete",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInternalFailure,
			}, nil
		}
	} else {
		// remove from organization
		err := removeFromOrganization(client, currentModel)
		if err != nil {
			var target *mongodbatlas.ErrorResponse
			// if team is assigned to project then first delete from project
			if errors.As(err, &target) && target.ErrorCode == "CANNOT_DELETE_TEAM_ASSIGNED_TO_PROJECT" {
				if err := removeFromProject(client, currentModel); err != nil {
					return handler.ProgressEvent{
						OperationStatus:  handler.Failed,
						Message:          "Unable to Delete",
						HandlerErrorCode: cloudformation.HandlerErrorCodeInternalFailure,
					}, nil
				}

				// remove from organization if successfully deleted from project
				if err := removeFromOrganization(client, currentModel); err != nil {
					return handler.ProgressEvent{
						OperationStatus:  handler.Failed,
						Message:          "Unable to Delete",
						HandlerErrorCode: cloudformation.HandlerErrorCodeInternalFailure,
					}, nil
				}
			}
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}
func setup() {
	util.SetupLogger("mongodb-atlas-teams")
}
func removeFromProject(client *mongodbatlas.Client, currentModel *Model) error {
	teamID := cast.ToString(currentModel.TeamId)
	projectID, err := getProjectIDByTeamID(context.Background(), client, teamID)
	if err != nil {
		_, _ = logger.Debugf("error to get assigned project details for Team: %s", teamID)
		return err
	}
	_, err = client.Teams.RemoveTeamFromProject(context.Background(), projectID, teamID)
	if err != nil {
		_, _ = logger.Debugf("error deleting Team from project: %s", teamID)
		return err
	}
	return nil
}
func removeFromOrganization(client *mongodbatlas.Client, currentModel *Model) error {
	teamID := cast.ToString(currentModel.TeamId)
	orgID := cast.ToString(currentModel.OrgId)

	_, err := client.Teams.RemoveTeamFromOrganization(context.Background(), orgID, teamID)
	if err != nil {
		_, _ = logger.Debugf("error deleting team from organization in retry : %s", teamID)
		return err
	}
	return nil
}
func isExist(client *mongodbatlas.Client, currentModel *Model) bool {
	teamID := cast.ToString(currentModel.TeamId)
	orgID := cast.ToString(currentModel.OrgId)
	teamName := cast.ToString(currentModel.Name)
	if *currentModel.TeamId != "" {
		team, _, err := client.Teams.Get(context.Background(), orgID, teamID)
		if err != nil {
			return false
		}
		if team != nil {
			return true
		}
	} else if *currentModel.Name != "" {
		team, _, err := client.Teams.GetOneTeamByName(context.Background(), orgID, teamName)
		if err != nil {
			return false
		}
		if team != nil {
			return true
		}
	}

	return false
}
func isUserExist(users []mongodbatlas.AtlasUser, username string) (mongodbatlas.AtlasUser, bool) {
	endLoop := len(users)
	for ind := 0; ind < endLoop; ind++ {
		_, _ = logger.Debugf("atlas user : %s,target User %s", users[ind].Username, username)
		if users[ind].Username == username {
			return users[ind], true
		}
	}
	return mongodbatlas.AtlasUser{}, false
}

func getProjectIDByTeamID(ctx context.Context, conn *mongodbatlas.Client, teamID string) (string, error) {
	options := &mongodbatlas.ListOptions{}
	projects, _, err := conn.Projects.GetAllProjects(ctx, options)
	if err != nil {
		return "", fmt.Errorf("error getting projects information: %s", err)
	}

	for _, project := range projects.Results {
		teams, _, err := conn.Projects.GetProjectTeamsAssigned(ctx, project.ID)
		if err != nil {
			return "", fmt.Errorf("error getting teams from project information: %s", err)
		}

		for _, team := range teams.Results {
			if team.TeamID == teamID {
				return project.ID, nil
			}
		}
	}
	return "", nil
}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func convertProjectTeamToModel(team *mongodbatlas.Result) *Model {
	if team == nil {
		return nil
	}
	return &Model{
		RoleNames: team.RoleNames,
		TeamId:    &team.TeamID,
	}
}
func convertTeamToModel(team *mongodbatlas.Team, result *Model) *Model {
	if result == nil {
		result = new(Model)
	}

	if team.ID != "" {
		result.TeamId = &team.ID
	}
	if team.Name != "" {
		result.Name = &team.Name
	}
	if team.Usernames != nil {
		result.Usernames = team.Usernames
	}
	return result
}
