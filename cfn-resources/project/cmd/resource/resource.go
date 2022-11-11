package resource

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.OrgID, constants.Name}
var ReadRequiredFields = []string{constants.PubKey, constants.PvtKey}
var UpdateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ID}
var DeleteRequiredFields = []string{constants.PubKey, constants.PvtKey}
var ListRequiredFields = []string{constants.PubKey, constants.PvtKey}

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

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %s", err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	var projectOwnerID string
	if currentModel.ProjectOwnerId != nil {
		projectOwnerID = *currentModel.ProjectOwnerId
	}
	project, res, err := client.Projects.Create(context.Background(), &mongodbatlas.Project{
		Name:                      *currentModel.Name,
		OrgID:                     *currentModel.OrgId,
		WithDefaultAlertsSettings: currentModel.WithDefaultAlertsSettings,
	}, &mongodbatlas.CreateProjectOptions{ProjectOwnerID: projectOwnerID})
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

	if currentModel.ProjectSettings != nil {
		// Update project settings
		projectSettings := mongodbatlas.ProjectSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: currentModel.ProjectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
			IsRealtimePerformancePanelEnabled:           currentModel.ProjectSettings.IsRealtimePerformancePanelEnabled,
			IsDataExplorerEnabled:                       currentModel.ProjectSettings.IsDataExplorerEnabled,
			IsPerformanceAdvisorEnabled:                 currentModel.ProjectSettings.IsPerformanceAdvisorEnabled,
			IsSchemaAdvisorEnabled:                      currentModel.ProjectSettings.IsSchemaAdvisorEnabled,
		}

		_, res, err = client.Projects.UpdateProjectSettings(context.Background(), project.ID, &projectSettings)
		if err != nil {
			_, _ = logger.Warnf("UpdateProjectSettings Error: %s", err)
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Failed to update Project settings : %s", err.Error()),
				res.Response), nil
		}
	}

	currentModel.Id = &project.ID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount

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

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
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

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	var projectID string
	if currentModel.Id != nil {
		projectID = *currentModel.Id
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

	if currentModel.ProjectSettings != nil {
		// Update project settings
		projectSettings := mongodbatlas.ProjectSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: currentModel.ProjectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
			IsRealtimePerformancePanelEnabled:           currentModel.ProjectSettings.IsRealtimePerformancePanelEnabled,
			IsDataExplorerEnabled:                       currentModel.ProjectSettings.IsDataExplorerEnabled,
			IsPerformanceAdvisorEnabled:                 currentModel.ProjectSettings.IsPerformanceAdvisorEnabled,
			IsSchemaAdvisorEnabled:                      currentModel.ProjectSettings.IsSchemaAdvisorEnabled,
		}
		_, _, err = client.Projects.UpdateProjectSettings(context.Background(), projectID, &projectSettings)
		if err != nil {
			_, _ = logger.Warnf("Update - error: %+v", err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Failed to update Project settings",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
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

	if errEvent := validateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	_, _ = logger.Debugf("Delete Project prevModel:%+v currentModel:%+v", *prevModel, *currentModel)

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
	setup()
	_, _ = logger.Debugf("List.Project prevModel:%+v currentModel:%+v", prevModel, currentModel)

	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return *errEvent, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	listOptions := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}
	projects, _, err := client.Projects.GetAllProjects(context.Background(), listOptions)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error retrieving projects: %s", err)
	}

	// Initialize like this in case no results will pass empty array
	var projectModels []interface{}
	for _, project := range projects.Results {
		var m Model
		m.Name = &project.Name
		m.Id = &project.ID
		m.ApiKeys = currentModel.ApiKeys
		event, model, err := readProjectSettings(client, project.ID, &m)
		if err != nil {
			return event, err
		}
		model.Name = &project.Name
		model.Id = &project.ID
		model.Created = &project.Created
		model.ClusterCount = &project.ClusterCount
		model.OrgId = &project.OrgID
		projectModels = append(projectModels, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  projectModels,
	}, nil
}

// Read project
func getProject(client *mongodbatlas.Client, currentModel *Model) (event handler.ProgressEvent, model *Model, err error) {
	var project *mongodbatlas.Project
	if len(*currentModel.Name) > 0 {
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
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Error while deleting project",
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil, err
		}
		return progressevents.GetFailedEventByResponse(err.Error(),
			res.Response), project, err
	}
	return handler.ProgressEvent{}, project, err
}

func getProjectByID(id *string, client *mongodbatlas.Client) (event handler.ProgressEvent, model *mongodbatlas.Project, err error) {
	project, res, err := client.Projects.GetOneProjectByName(context.Background(), *id)
	if err != nil {
		if res.Response.StatusCode == 401 { // cfn test
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Error while deleting project",
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil, err
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

	// Get APIKeys from project
	projectAPIKeys, res, err := client.ProjectAPIKeys.List(context.Background(), id, &mongodbatlas.ListOptions{ItemsPerPage: 1000, IncludeCount: true})
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
	}

	// Set teams
	var teams []ProjectTeam
	for _, team := range teamsAssigned.Results {
		if len(team.TeamID) > 0 {
			teams = append(teams, ProjectTeam{TeamId: &team.TeamID, RoleNames: team.RoleNames})
		}
	}

	// Set api-keys
	apiKeys := readKeys(*currentModel.Id, projectAPIKeys, currentModel)
	currentModel.ProjectTeams = teams
	currentModel.ProjectApiKeys = apiKeys
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

func readKeys(groupID string, keys []mongodbatlas.APIKey, currentModel *Model) []ProjectApiKey {
	var apiKeys []ProjectApiKey
	for i := range keys {
		// Don't include the org level key used for atlas authentication
		// cfn test doesn't allow extra keys in the response
		if keys[i].PublicKey == *currentModel.ApiKeys.PublicKey {
			continue
		}
		var roles []string
		for j := range keys[i].Roles {
			if keys[i].Roles[j].GroupID == groupID && !strings.HasPrefix(keys[i].Roles[j].RoleName, "ORG_") {
				roles = append(roles, keys[i].Roles[j].RoleName)
			}
		}
		apiKeys = append(apiKeys, ProjectApiKey{Key: &keys[i].ID, RoleNames: roles})
	}
	return apiKeys
}
