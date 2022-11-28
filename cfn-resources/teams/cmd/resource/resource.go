package resource

import (
	"context"
	"errors"
	"fmt"

	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.OrgID}
var ReadRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.OrgID, constants.TeamID}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("Create Team - Request:%+v", currentModel)

	// Validate required fields in the request
	event, client, err := ValidateRequest(CreateRequiredFields, currentModel)
	if err != nil {
		return event, err
	}

	// Create Atlas API Request Object
	teamRequest := convertToMongoClientModel(currentModel, nil)

	// call API to create
	team, teamsResp, err := client.Teams.Create(context.Background(), *currentModel.OrgId,
		teamRequest)
	if err != nil && err.Error() != "<nil>" {
		if err != nil {
			return handler.ProgressEvent{}, fmt.Errorf("error creating Team information: %s", err)
		}
	}
	// create response model
	currentModel = convertToModel(team, currentModel)
	_, _ = logger.Debugf("Created Successfully - (%s)", teamsResp.Body)

	event = handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("Read Team - Request: %+v", currentModel)
	// Validate required fields in the request
	event, client, err := ValidateRequest(ReadRequiredFields, currentModel)
	if err != nil {
		return event, err
	}
	// API call to read snapshot to read using ID field
	orgID := *currentModel.OrgId
	teamID := *currentModel.TeamId
	var team *mongodbatlas.Team
	// get team by id or name
	if *currentModel.TeamId != "" {
		team, _, err = client.Teams.Get(context.Background(), orgID, teamID)
	} else if *currentModel.Name != "" {
		team, _, err = client.Teams.GetOneTeamByName(context.Background(), *currentModel.OrgId, *currentModel.Name)
	}
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	if team != nil {
		currentModel = convertToModel(team, currentModel)
	}
	// API call to get all users assigned
	users, _, err := client.Teams.GetTeamUsersAssigned(context.Background(), orgID, *currentModel.TeamId)
	if err != nil {
		_, _ = logger.Debugf("error getting Team user information: %s", err)
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
			RoleName: &role[ind].RoleName,
			GroupId:  &role[ind].GroupID,
			OrgId:    &role[ind].OrgID,
		}
		modelRole = append(modelRole, pe)
	}
	return modelRole
}
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("Update Team Request :%+v", currentModel)
	// Validate required fields in the request
	event, client, err := ValidateRequest(ReadRequiredFields, currentModel)
	if err != nil {
		return event, err
	}
	isExist := isExist(client, currentModel)
	if !isExist {
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
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	_, _ = logger.Debugf("Read -Team information status (%d)", res.StatusCode)

	// rename the team
	if team.Name != *currentModel.Name {
		_, _, err := client.Teams.Rename(context.Background(), *currentModel.OrgId, *currentModel.TeamId, *currentModel.Name)
		if err != nil {
			_, _ = logger.Debugf("error updating Team information: %v", err)
		}
	}

	// add/remove user to/from teams
	if currentModel.Usernames != nil {
		// get the current team's users
		users, _, err := client.Teams.GetTeamUsersAssigned(context.Background(), *currentModel.OrgId, *currentModel.TeamId)
		if err != nil {
			_, _ = logger.Debugf("Read -error (%v)", err)
		}
		usernames := currentModel.Usernames
		var newUsers []string
		for ind := range usernames {
			currentUser, isExistingUser := isUserExist(users, usernames[ind])
			_, _ = logger.Debugf("team -current user %v \n", currentUser)
			if isExistingUser {
				_, _ = logger.Debugf("team -remove user starts \n")
				_, err := client.Teams.RemoveUserToTeam(context.Background(), *currentModel.OrgId, *currentModel.TeamId, currentUser.ID)
				if err != nil {
					_, _ = logger.Debugf("Remove Team -error (%v) \n", err)
				}
			} else {
				_, _ = logger.Debugf("team -append user starts \n")
				user, _, err := client.AtlasUsers.GetByName(context.Background(), usernames[ind])
				if err != nil {
					_, _ = logger.Debugf("Error reading user (%s)  with error (%v) \n", usernames[ind], err)
				}
				// if the user exists, we will store its ID
				if user != nil {
					newUsers = append(newUsers, user.ID)
				}
			}
		}
		_, _ = logger.Debugf("team -Add user starts \n")
		_, _, err = client.Teams.AddUsersToTeam(context.Background(), *currentModel.OrgId, *currentModel.TeamId, newUsers)
		if err != nil {
			_, _ = logger.Debugf("team -Add user response status (%d)", res.StatusCode)
		}
	}

	// update roles to team
	roleNames := currentModel.RoleNames
	if len(roleNames) > 0 {
		_, _ = logger.Debugf("update Role starts \n")
		teamRequest := &mongodbatlas.TeamUpdateRoles{RoleNames: roleNames}
		_, _, err = client.Teams.UpdateTeamRoles(context.Background(), *currentModel.OrgId, *currentModel.TeamId, teamRequest)
		if err != nil {
			_, _ = logger.Debugf("update Role response status (%d) \n", res.StatusCode)
		}
	}
	event = handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}
	return event, nil
}
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("List Teams  Request :%+v", currentModel)

	// Validate required fields in the request
	event, client, err := ValidateRequest(CreateRequiredFields, currentModel)
	if err != nil {
		return event, err
	}

	// Create Atlas API Request Object
	orgID := *currentModel.OrgId
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}
	// API call
	teams, _, err := client.Teams.List(context.Background(), orgID, params)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading teamlist with  id(Organization: %s): %s", orgID, err)
	}
	var models []interface{}

	for i := 0; i < len(teams); i++ {
		var model Model
		model.TeamId = &teams[i].ID
		model.Name = &teams[i].Name
		model.Usernames = teams[i].Usernames
		fmt.Println(teams[i])
		models = append(models, model)
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

	// Validate required fields in the request
	event, client, err := ValidateRequest(CreateRequiredFields, currentModel)
	if err != nil {
		return event, err
	}

	isExist := isExist(client, currentModel)
	if !isExist {
		_, _ = logger.Debugf("error deleting Team : %s", *currentModel.TeamId)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
		}, nil
	}
	_, err = client.Teams.RemoveTeamFromOrganization(context.Background(), *currentModel.OrgId, *currentModel.TeamId)
	if err != nil {
		var target *mongodbatlas.ErrorResponse
		if errors.As(err, &target) && target.ErrorCode == "CANNOT_DELETE_TEAM_ASSIGNED_TO_PROJECT" {
			projectID, err := getProjectIDByTeamID(context.Background(), client, *currentModel.TeamId)
			if err != nil {
				_, _ = logger.Debugf("error to get assigned project details for Team: %s", *currentModel.TeamId)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Unable to Delete",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInternalFailure,
				}, nil
			}
			_, err = client.Teams.RemoveTeamFromProject(context.Background(), projectID, *currentModel.TeamId)
			if err != nil {
				_, _ = logger.Debugf("error deleting Team from project: %s", *currentModel.TeamId)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Unable to Delete",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInternalFailure,
				}, nil
			}
			_, err = client.Teams.RemoveTeamFromOrganization(context.Background(), *currentModel.OrgId, *currentModel.TeamId)
			if err != nil {
				_, _ = logger.Debugf("error deleting team from organization in retry : %s", *currentModel.TeamId)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Unable to Delete",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInternalFailure,
				}, nil
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

func isExist(client *mongodbatlas.Client, currentModel *Model) bool {
	if *currentModel.TeamId != "" {
		team, _, err := client.Teams.Get(context.Background(), *currentModel.OrgId, *currentModel.TeamId)
		if err != nil {
			return false
		}
		if team != nil {
			return true
		}
	} else if *currentModel.Name != "" {
		team, _, err := client.Teams.GetOneTeamByName(context.Background(), *currentModel.OrgId, *currentModel.Name)
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

// ValidateRequest function to validate the request
func ValidateRequest(requiredFields []string, currentModel *Model) (handler.ProgressEvent, *mongodbatlas.Client, error) {
	// Validate required fields are empty or nil
	if modelValidation := validateModel(requiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil, errors.New("required field not found")
	}
	// Validate API Keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil, err
	}

	return handler.ProgressEvent{}, client, nil
}

func convertToMongoClientModel(currentModel *Model, reqModel *mongodbatlas.Team) *mongodbatlas.Team {
	setup() // logger setup
	if reqModel == nil {
		reqModel = &mongodbatlas.Team{}
	}
	if currentModel.TeamId != nil {
		reqModel.ID = *currentModel.TeamId
	}
	if currentModel.Name != nil {
		reqModel.Name = *currentModel.Name
	}
	if currentModel.Usernames != nil {
		reqModel.Usernames = currentModel.Usernames
	}
	return reqModel
}
func convertToModel(team *mongodbatlas.Team, result *Model) *Model {
	setup() // logger setup
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
