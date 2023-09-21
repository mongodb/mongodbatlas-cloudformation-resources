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
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
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

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	atlasV2 := client.AtlasV2
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
		teamResponse, resp, err := atlasV2.TeamsApi.CreateTeam(context.Background(), orgID, &admin.Team{
			Name:      cast.ToString(currentModel.Name),
			Usernames: currentModel.Usernames,
		}).Execute()

		if err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("unable to create team %v", err), resp), nil
		}
		teamID = util.SafeString(teamResponse.Id)
		currentModel = convertTeamToModel(teamResponse, currentModel)
	}

	// add existing team or newly created team to project if project id exist in the request
	if projectID != "" && len(currentModel.RoleNames) > 0 {
		createRequest := []admin.TeamRole{{
			TeamId:    &teamID,
			RoleNames: currentModel.RoleNames,
		}}
		_, _, err := atlasV2.TeamsApi.AddAllTeamsToProject(context.Background(), projectID, &createRequest).Execute()
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

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	atlasV2 := client.AtlasV2
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	// API call to read snapshot to read using ID field
	teamID := cast.ToString(currentModel.TeamId)
	orgID := cast.ToString(currentModel.OrgId)
	teamName := cast.ToString(currentModel.Name)
	var team *admin.TeamResponse
	var resp *http.Response
	var err error
	// get team by id or name
	if teamID != "" {
		team, resp, err = atlasV2.TeamsApi.GetTeamById(context.Background(), orgID, teamID).Execute()
	} else if teamName != "" {
		// get team by name
		team, resp, err = atlasV2.TeamsApi.GetTeamByName(context.Background(), orgID, teamName).Execute()
	}

	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel = convertTeamResponseToModel(team, currentModel)

	paginatedResp, _, err := atlasV2.TeamsApi.ListTeamUsers(context.Background(), orgID, *currentModel.TeamId).Execute()
	if err != nil {
		_, _ = logger.Warnf("error getting Team user information: %v", err)
	}
	if paginatedResp != nil {
		usersRespList := paginatedResp.Results
		var userNames []string
		var userList []AtlasUser
		for ind := range usersRespList {
			userNames = append(userNames, usersRespList[ind].Username)
			userList = append(userList, newAtlasUser(usersRespList[ind]))
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

func newAtlasUser(user admin.CloudAppUser) AtlasUser {
	return AtlasUser{
		Country:      &user.Country,
		EmailAddress: &user.EmailAddress,
		FirstName:    &user.FirstName,
		Id:           user.Id,
		LastName:     &user.LastName,
		MobileNumber: &user.MobileNumber,
		Password:     &user.Password,
		Roles:        newAtlasRoles(user.Roles),
		TeamIds:      user.TeamIds,
		Username:     &user.Username,
	}
}
func newAtlasRoles(roles []admin.CloudAccessRoleAssignment) []AtlasRole {
	var modelRole []AtlasRole
	if roles == nil {
		return modelRole
	}
	for ind := range roles {
		pe := AtlasRole{
			RoleName:  roles[ind].RoleName,
			ProjectId: roles[ind].GroupId,
			OrgId:     roles[ind].OrgId,
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

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	atlasV2 := client.AtlasV2
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	team, res, err := getTeam(atlasV2, currentModel)
	if err != nil && res != nil {
		_, _ = logger.Debugf("error getting Team information: %s", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res), nil
	} else if err != nil {
		_, _ = logger.Debugf("error getting Team information: %s", *currentModel.TeamId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	teamID := cast.ToString(currentModel.TeamId)
	orgID := cast.ToString(currentModel.OrgId)
	teamName := cast.ToString(currentModel.Name)
	projectID := cast.ToString(currentModel.ProjectId)

	// add existing team or newly created team to project if project id exist in the request
	if projectID != "" && len(currentModel.RoleNames) > 0 {
		createRequest := []admin.TeamRole{{
			TeamId:    &teamID,
			RoleNames: currentModel.RoleNames,
		}}
		_, _, err := atlasV2.TeamsApi.AddAllTeamsToProject(context.Background(), projectID, &createRequest).Execute()
		if err != nil {
			_, _ = logger.Warnf("error adding Team(%s) to project(%s): reason : %v", teamID, projectID, err)
		}
	}

	// rename the team
	if !util.AreStringPtrEqual(team.Name, &teamName) {
		_, _, err := atlasV2.TeamsApi.RenameTeam(context.Background(), orgID, teamID, &admin.Team{
			Name: teamName,
		}).Execute()
		if err != nil {
			_, _ = logger.Warnf("error updating Team information: %v", err)
		}
	}

	// add/remove user to/from teams
	if currentModel.Usernames != nil {
		// get the current users list for the team
		paginatedResp, _, err := atlasV2.TeamsApi.ListTeamUsers(context.Background(), orgID, teamID).Execute()
		if err != nil {
			_, _ = logger.Warnf("get assigned user to team -error (%v)", err)
		}
		usernames := currentModel.Usernames
		var newUsers []admin.AddUserToTeam
		for ind := range usernames {
			currentUser, isExistingUser := isUserExist(paginatedResp.Results, usernames[ind])

			if isExistingUser {
				// remove user from team
				_, err := atlasV2.TeamsApi.RemoveTeamUser(context.Background(), orgID, teamID, util.SafeString(currentUser.Id)).Execute()
				if err != nil {
					_, _ = logger.Warnf("remove user(%s) from Team(%s) -error (%v) \n", util.SafeString(currentUser.Id), teamID, err)
				}
			} else {
				// add user to team
				user, _, err := atlasV2.MongoDBCloudUsersApi.GetUserByUsername(context.Background(), usernames[ind]).Execute()
				if err != nil {
					_, _ = logger.Warnf("Error reading user (%s)  with error (%v) \n", usernames[ind], err)
				}
				// if the user exists, we will store its ID so that we can save as user list later
				if user != nil {
					newUsers = append(newUsers, admin.AddUserToTeam{Id: util.SafeString(user.Id)})
				}
			}
		}
		// save all new users
		if len(newUsers) > 0 {
			atlasV2.TeamsApi.AddTeamUser(context.Background(), orgID, teamID, &newUsers)
			if err != nil {
				_, _ = logger.Warnf("team -Add users error (%+v) \n", err)
			}
		}
	}

	// update roles to team
	roleNames := currentModel.RoleNames
	if len(roleNames) > 0 && currentModel.ProjectId != nil {
		teamRequest := &admin.TeamRole{RoleNames: roleNames}
		_, _, err = atlasV2.TeamsApi.UpdateTeamRoles(context.Background(), projectID, teamID, teamRequest).Execute()
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

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	atlasV2 := client.AtlasV2
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	// Create Atlas API Request Object
	orgID := cast.ToString(currentModel.OrgId)
	projectID := cast.ToString(currentModel.ProjectId)
	var models []interface{}
	var resp *http.Response
	var err error
	// API call to get teams for project id
	if projectID != "" {
		var teamsAssigned *admin.PaginatedTeamRole
		teamsAssigned, resp, err = atlasV2.TeamsApi.ListProjectTeams(context.Background(), projectID).Execute()

		if err != nil {
			return progressevents.GetFailedEventByResponse(err.Error(), resp), nil
		}

		teamsProjectList := teamsAssigned.Results
		for i := 0; i < len(teamsProjectList); i++ {
			models = append(models, convertProjectTeamToModel(teamsProjectList[i]))
		}
	} else {
		// API call to get teams from organization
		var paginatedResp *admin.PaginatedTeam
		paginatedResp, resp, err = atlasV2.TeamsApi.ListOrganizationTeams(context.Background(), orgID).Execute()

		if err != nil {
			return progressevents.GetFailedEventByResponse(err.Error(), resp), nil
		}
		teams := paginatedResp.Results
		for i := 0; i < len(teams); i++ {
			models = append(models, convertTeamResponseToModel(&teams[i], nil))
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

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, handlerError := util.NewAtlasClient(&req, currentModel.Profile)
	atlasV2 := client.AtlasV2
	if handlerError != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", handlerError)
		return *handlerError, errors.New(handlerError.Message)
	}

	team, _, _ := getTeam(atlasV2, currentModel)
	if team == nil {
		_, _ = logger.Debugf("error getting Team information: %s", *currentModel.TeamId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	if currentModel.ProjectId != nil {
		if err := removeFromProject(atlasV2, currentModel); err != nil {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Unable to Delete",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInternalFailure,
			}, nil
		}
	} else {
		// remove from organization
		err := removeFromOrganization(atlasV2, currentModel)
		if err != nil {
			var target *mongodbatlas.ErrorResponse
			// if team is assigned to project then first delete from project
			if errors.As(err, &target) && target.ErrorCode == "CANNOT_DELETE_TEAM_ASSIGNED_TO_PROJECT" {
				if err := removeFromProject(atlasV2, currentModel); err != nil {
					return handler.ProgressEvent{
						OperationStatus:  handler.Failed,
						Message:          "Unable to Delete",
						HandlerErrorCode: cloudformation.HandlerErrorCodeInternalFailure,
					}, nil
				}

				// remove from organization if successfully deleted from project
				if err := removeFromOrganization(atlasV2, currentModel); err != nil {
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
func removeFromProject(atlasV2 *admin.APIClient, currentModel *Model) error {
	teamID := cast.ToString(currentModel.TeamId)
	projectID, err := getProjectIDByTeamID(context.Background(), atlasV2, teamID)
	if err != nil {
		_, _ = logger.Debugf("error to get assigned project details for Team: %s", teamID)
		return err
	}
	_, err = atlasV2.TeamsApi.RemoveProjectTeam(context.Background(), projectID, teamID).Execute()
	if err != nil {
		_, _ = logger.Debugf("error deleting Team from project: %s", teamID)
		return err
	}
	return nil
}

func removeFromOrganization(atlasV2 *admin.APIClient, currentModel *Model) error {
	teamID := cast.ToString(currentModel.TeamId)
	orgID := cast.ToString(currentModel.OrgId)

	_, _, err := atlasV2.TeamsApi.DeleteTeam(context.Background(), orgID, teamID).Execute()
	if err != nil {
		_, _ = logger.Debugf("error deleting team from organization in retry : %s", teamID)
		return err
	}
	return nil
}

func getTeam(atlasV2 *admin.APIClient, currentModel *Model) (*admin.TeamResponse, *http.Response, error) {
	teamID := cast.ToString(currentModel.TeamId)
	orgID := cast.ToString(currentModel.OrgId)
	teamName := cast.ToString(currentModel.Name)
	if *currentModel.TeamId != "" {
		team, res, err := atlasV2.TeamsApi.GetTeamById(context.Background(), orgID, teamID).Execute()
		return team, res, err
	} else if *currentModel.Name != "" {
		team, res, err := atlasV2.TeamsApi.GetTeamByName(context.Background(), orgID, teamName).Execute()
		return team, res, err
	}
	return nil, nil, errors.New("could not fetch Team as neither TeamId or Name were defined in model")
}

func isUserExist(users []admin.CloudAppUser, username string) (admin.CloudAppUser, bool) {
	endLoop := len(users)
	for ind := 0; ind < endLoop; ind++ {
		_, _ = logger.Debugf("atlas user : %s,target User %s", users[ind].Username, username)
		if users[ind].Username == username {
			return users[ind], true
		}
	}
	return admin.CloudAppUser{}, false
}

func getProjectIDByTeamID(ctx context.Context, atlasV2 *admin.APIClient, teamID string) (string, error) {
	paginatedResp, _, err := atlasV2.ProjectsApi.ListProjects(context.Background()).Execute()
	if err != nil {
		return "", fmt.Errorf("error getting projects information: %s", err)
	}

	for _, project := range paginatedResp.Results {
		teams, _, err := atlasV2.TeamsApi.ListProjectTeams(ctx, util.SafeString(project.Id)).Execute()
		if err != nil {
			return "", fmt.Errorf("error getting teams from project information: %s", err)
		}

		for _, team := range teams.Results {
			if util.AreStringPtrEqual(team.TeamId, &teamID) {
				return util.SafeString(project.Id), nil
			}
		}
	}
	return "", nil
}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func convertProjectTeamToModel(team admin.TeamRole) *Model {
	return &Model{
		RoleNames: team.RoleNames,
		TeamId:    team.TeamId,
	}
}
func convertTeamToModel(team *admin.Team, result *Model) *Model {
	if result == nil {
		result = new(Model)
	}

	if team.Id != nil && *team.Id != "" {
		result.TeamId = team.Id
	}
	if team.Name != "" {
		result.Name = &team.Name
	}
	if team.Usernames != nil {
		result.Usernames = team.Usernames
	}
	return result
}

func convertTeamResponseToModel(team *admin.TeamResponse, result *Model) *Model {
	if result == nil {
		result = new(Model)
	}

	if team.Id != nil && *team.Id != "" {
		result.TeamId = team.Id
	}
	if team.Id != nil && *team.Name != "" {
		result.Name = team.Name
	}
	return result
}
