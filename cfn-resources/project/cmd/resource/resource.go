package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/project/cmd/validation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progress_event"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/atlas/mongodbatlas"
	"strings"
)

type UpdateApiKey struct {
	Key     string
	ApiKeys *mongodbatlas.AssignAPIKey
}

func setup() {
	util.SetupLogger("mongodb-atlas-project")
}

// validateModel inputs based on the method
func validateModel(event constants.Event, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(event, validation.ModelValidator{}, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("Create log.Debugf-- currentModel: %+v", *currentModel)

	modelValidation := validateModel(constants.Create, currentModel)
	if modelValidation != nil {
		log.Debugf("Validation Error")
		return *modelValidation, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("CreateMongoDBClient error: %s", err)
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	var projectOwnerId string
	if currentModel.ProjectOwnerId != nil {
		projectOwnerId = *currentModel.ProjectOwnerId
	}
	project, res, err := client.Projects.Create(context.Background(), &mongodbatlas.Project{
		Name:                      *currentModel.Name,
		OrgID:                     *currentModel.OrgId,
		WithDefaultAlertsSettings: currentModel.WithDefaultAlertsSettings,
	}, &mongodbatlas.CreateProjectOptions{ProjectOwnerID: projectOwnerId})
	if err != nil {
		//return handler.ProgressEvent{}, fmt.Errorf("error creating project: %s", err)
		log.Debugf("Create - error: %+v", err)
		// TODO- Should detect and return HandlerErrorCodeAlreadyExists
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Project : %s", err.Error()),
			res.Response), nil

	}

	//Add ApiKeys
	if len(currentModel.ProjectApiKeys) > 0 {
		for _, key := range currentModel.ProjectApiKeys {
			_, err = client.ProjectAPIKeys.Assign(context.Background(), project.ID, *key.Key, &mongodbatlas.AssignAPIKey{Roles: key.RoleNames})
			if err != nil {
				log.Debugf("Assign Key Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while Assigning Key to project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}
	}

	//Add Teams
	if len(currentModel.ProjectTeams) > 0 {
		_, _, err = client.Projects.AddTeamsToProject(context.Background(), project.ID, readTeams(currentModel.ProjectTeams))
		if err != nil {
			log.Debugf("AddTeamsToProject Error: %s", err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Error while adding teams to project",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
	}

	if currentModel.ProjectSettings != nil {
		//Update project settings
		projectSettings := mongodbatlas.ProjectSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: currentModel.ProjectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
			IsRealtimePerformancePanelEnabled:           currentModel.ProjectSettings.IsRealtimePerformancePanelEnabled,
			IsDataExplorerEnabled:                       currentModel.ProjectSettings.IsDataExplorerEnabled,
			IsPerformanceAdvisorEnabled:                 currentModel.ProjectSettings.IsPerformanceAdvisorEnabled,
			IsSchemaAdvisorEnabled:                      currentModel.ProjectSettings.IsSchemaAdvisorEnabled,
		}

		_, res, err = client.Projects.UpdateProjectSettings(context.Background(), project.ID, &projectSettings)
		if err != nil {
			log.Debugf("UpdateProjectSettings Error: %s", err)
			return progress_events.GetFailedEventByResponse(fmt.Sprintf("Failed to update Project settings : %s", err.Error()),
				res.Response), nil
		}
	}

	currentModel.Id = &project.ID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount

	event, errr, failed, proj := getProject(*currentModel.Name, client, currentModel, err)
	if failed {
		log.Debugf("getProject Error: %s", err)
		return event, errr
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
	modelValidation := validateModel(constants.Read, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	var name string
	if currentModel.Name != nil {
		name = *currentModel.Name
	}
	event, errr, failed, model := getProject(name, client, currentModel, err)
	if failed {
		return event, errr
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   model,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	modelValidation := validateModel(constants.Update, currentModel)
	if modelValidation != nil {
		log.Debugf("Validation Error")
		return *modelValidation, nil
	}
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	var projectId string
	if currentModel.Id != nil {
		projectId = *currentModel.Id
	}

	if currentModel.ProjectTeams != nil {
		//Get teams from project
		teamsAssigned, _, err := client.Projects.GetProjectTeamsAssigned(context.Background(), projectId)
		if err != nil {
			log.Debugf("ProjectId : %s, Error: %s", projectId, err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Error while finding teams in project",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		newTeams, changedTeams, removeTeams := getChangeInTeams(currentModel.ProjectTeams, teamsAssigned.Results)

		//Remove Teams
		for _, team := range removeTeams {
			_, err := client.Teams.RemoveTeamFromProject(context.Background(), projectId, team.TeamID)
			if err != nil {
				log.Debug("Error: %s", projectId, err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while deleting team from project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}
		// Add Teams
		if len(newTeams) > 0 {
			_, _, err = client.Projects.AddTeamsToProject(context.Background(), projectId, newTeams)
			if err != nil {
				log.Debugf("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while adding team to project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}
		// Update Teams
		for _, team := range changedTeams {
			_, _, err = client.Teams.UpdateTeamRoles(context.Background(), projectId, team.TeamID, &mongodbatlas.TeamUpdateRoles{RoleNames: team.RoleNames})
			if err != nil {
				log.Debug("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while updating team roles in project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
		}
	}

	if currentModel.ProjectApiKeys != nil {
		//Get APIKeys from project
		projectApiKeys, _, err := client.ProjectAPIKeys.List(context.Background(), projectId, &mongodbatlas.ListOptions{ItemsPerPage: 1000, IncludeCount: true})
		if err != nil {
			log.Debug("Error: %s", projectId, err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Error while finding api keys in project",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}

		//log.Debugf("keys: %+v", currentModel.ProjectApiKeys)
		//Get Change in ApiKeys
		newApiKeys, changedKeys, removeKeys := getChangeInApiKeys(*currentModel.Id, currentModel.ProjectApiKeys, projectApiKeys)

		//Remove old keys
		for _, key := range removeKeys {
			_, err = client.ProjectAPIKeys.Unassign(context.Background(), projectId, key.Key)
			if err != nil {
				log.Debug("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while Un-assigning Key to project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
			//log.Debugf("Removed: %s", key)

		}

		//Add Keys
		for _, key := range newApiKeys {
			_, err = client.ProjectAPIKeys.Assign(context.Background(), projectId, key.Key, key.ApiKeys)
			if err != nil {
				log.Debugf("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while Assigning Key to project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
			//log.Debugf("Added: %s", key)

		}

		//Update Key Roles
		for _, key := range changedKeys {
			_, err = client.ProjectAPIKeys.Assign(context.Background(), projectId, key.Key, key.ApiKeys)
			if err != nil {
				log.Debugf("Error: %s", err)
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          "Error while Un-assigning Key to project",
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
			}
			//log.Debugf("Updated: %s", key)
		}
	}

	if currentModel.ProjectSettings != nil {
		//Update project settings
		projectSettings := mongodbatlas.ProjectSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: currentModel.ProjectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
			IsRealtimePerformancePanelEnabled:           currentModel.ProjectSettings.IsRealtimePerformancePanelEnabled,
			IsDataExplorerEnabled:                       currentModel.ProjectSettings.IsDataExplorerEnabled,
			IsPerformanceAdvisorEnabled:                 currentModel.ProjectSettings.IsPerformanceAdvisorEnabled,
			IsSchemaAdvisorEnabled:                      currentModel.ProjectSettings.IsSchemaAdvisorEnabled,
		}
		_, _, err = client.Projects.UpdateProjectSettings(context.Background(), projectId, &projectSettings)
		if err != nil {
			log.Debugf("Update - error: %+v", err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Failed to update Project settings",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
	}

	event, err2, problem, project := getProject("", client, currentModel, err)
	if problem {
		return event, err2
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   project,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validateModel(constants.Delete, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err

	}
	log.Debugf("Delete Project prevModel:%+v currentModel:%+v", *prevModel, *currentModel)

	var id string
	if currentModel.Id != nil {
		id = *currentModel.Id
	}

	if len(id) == 0 {
		var name string
		if currentModel.Name != nil {
			name = *currentModel.Name
		}
		if len(name) > 0 {
			log.Debugf("Project id was nil, try lookup name:%s", name)
			project, res, err := client.Projects.GetOneProjectByName(context.Background(), name)
			if err != nil {
				return progress_events.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Project : %s", err.Error()),
					res.Response), nil
			}
			log.Debugf("Looked up project:%+v", project)
			id = project.ID
		} else {
			err := fmt.Errorf("@@@@Error deleting project. No Id or Name found currentModel:%+v)", currentModel)
			return handler.ProgressEvent{
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, err
		}
	}
	log.Debugf("Deleting project with id(%s)", id)

	res, err := client.Projects.Delete(context.Background(), id)
	if err != nil {
		//return handler.ProgressEvent{}, fmt.Errorf("####error deleting project with id(%s): %s", id, err)
		log.Warnf("####error deleting project with id(%s): %s", id, err)
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Failed to Create Project : %s", err.Error()),
			res.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		//ResourceModel:   currentModel,
		ResourceModel: nil,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("List.Project prevModel:%+v currentModel:%+v", prevModel, currentModel)

	modelValidation := validateModel(constants.List, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
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
	projectModels := []interface{}{}
	for _, project := range projects.Results {
		var m Model
		m.Name = &project.Name
		m.Id = &project.ID
		m.ApiKeys = currentModel.ApiKeys
		event, err2, model, failed := readProjectSettings(err, client, project.ID, &m)
		if failed {
			return event, err2
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
func getProject(name string, client *mongodbatlas.Client, currentModel *Model, err error) (handler.ProgressEvent, error, bool, *Model) {
	var id string
	if len(name) > 0 {
		project, _, err := client.Projects.GetOneProjectByName(context.Background(), name)
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil, true, nil
		}
		currentModel.Name = &project.Name
		currentModel.OrgId = &project.OrgID
		currentModel.Created = &project.Created
		currentModel.ClusterCount = &project.ClusterCount
		id = project.ID
		currentModel.Id = &project.ID
	} else {
		id := *currentModel.Id
		log.Debugf("Looking for project: %s", id)
		project, _, err := client.Projects.GetOneProject(context.Background(), id)
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil, true, nil
		}
		currentModel.Name = &project.Name
		currentModel.OrgId = &project.OrgID
		currentModel.Created = &project.Created
		currentModel.ClusterCount = &project.ClusterCount
		currentModel.Id = &project.ID
	}

	event, err2, model, failed := readProjectSettings(err, client, id, currentModel)

	if failed {
		return event, err2, true, model
	}

	return handler.ProgressEvent{}, nil, false, model
}

func readProjectSettings(err error, client *mongodbatlas.Client, id string, currentModel *Model) (handler.ProgressEvent, error, *Model, bool) {
	//Get teams from project
	teamsAssigned, res, err := client.Projects.GetProjectTeamsAssigned(context.Background(), id)
	if err != nil {
		log.Debug("ProjectId : %s, Error: %s", id, err)
		return progress_events.GetFailedEventByResponse(err.Error(),
			res.Response), nil, nil, true
	}

	//Get APIKeys from project
	projectApiKeys, res, err := client.ProjectAPIKeys.List(context.Background(), id, &mongodbatlas.ListOptions{ItemsPerPage: 1000, IncludeCount: true})
	if err != nil {
		log.Debug("Error: %s", id, err)
		return progress_events.GetFailedEventByResponse(err.Error(),
			res.Response), nil, nil, true
	}

	projectSettings, _, err := client.Projects.GetProjectSettings(context.Background(), id)
	if err != nil {
		log.Debug("Error: %s", id, err)
		return progress_events.GetFailedEventByResponse(err.Error(),
			res.Response), nil, nil, true
	}
	//Set projectSettings
	currentModel.ProjectSettings = &ProjectSettings{
		IsCollectDatabaseSpecificsStatisticsEnabled: projectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
		IsRealtimePerformancePanelEnabled:           projectSettings.IsRealtimePerformancePanelEnabled,
		IsDataExplorerEnabled:                       projectSettings.IsDataExplorerEnabled,
		IsPerformanceAdvisorEnabled:                 projectSettings.IsPerformanceAdvisorEnabled,
		IsSchemaAdvisorEnabled:                      projectSettings.IsSchemaAdvisorEnabled,
	}

	//Set teams
	teams := []ProjectTeam{}
	for _, team := range teamsAssigned.Results {
		if len(team.TeamID) > 0 {
			teams = append(teams, ProjectTeam{TeamId: &team.TeamID, RoleNames: team.RoleNames})
		}
	}

	//Set api-keys
	apiKeys := readKeys(*currentModel.Id, *currentModel.ApiKeys.PublicKey, projectApiKeys)
	currentModel.ProjectTeams = teams
	currentModel.ProjectApiKeys = apiKeys
	return handler.ProgressEvent{}, nil, currentModel, false
}

// Get difference in Teams
func getChangeInTeams(currentTeams []ProjectTeam, oTeams []*mongodbatlas.Result) ([]*mongodbatlas.ProjectTeam, []*mongodbatlas.ProjectTeam, []*mongodbatlas.ProjectTeam) {
	var newTeams []*mongodbatlas.ProjectTeam
	var changedTeams []*mongodbatlas.ProjectTeam
	var removeTeams []*mongodbatlas.ProjectTeam

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
			//Add to newTeams
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
func getChangeInApiKeys(groupId string, currentKeys []ProjectApiKey, oKeys []mongodbatlas.APIKey) ([]UpdateApiKey, []UpdateApiKey, []UpdateApiKey) {
	var newKeys []UpdateApiKey
	var changedKeys []UpdateApiKey
	var removeKeys []UpdateApiKey

	for _, nKey := range currentKeys {
		if nKey.Key != nil && len(*nKey.Key) > 0 {
			matched := false
			for _, oKey := range oKeys {
				if nKey.Key != nil && *nKey.Key == oKey.ID {
					changedKeys = append(changedKeys, UpdateApiKey{Key: *nKey.Key, ApiKeys: &mongodbatlas.AssignAPIKey{Roles: nKey.RoleNames}})
					matched = true
					break
				}
			}
			//Add to newKeys
			if !matched {
				newKeys = append(newKeys, UpdateApiKey{Key: *nKey.Key, ApiKeys: &mongodbatlas.AssignAPIKey{Roles: nKey.RoleNames}})
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
					//Consider only current ProjectRoles
					if role.GroupID == groupId {
						removeKeys = append(removeKeys, UpdateApiKey{Key: oKey.ID})
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

func readKeys(groupId, publicKey string, keys []mongodbatlas.APIKey) []ProjectApiKey {
	var apiKeys []ProjectApiKey
	for _, key := range keys {
		var roles []string
		for _, role := range key.Roles {
			if role.GroupID == groupId && !strings.HasPrefix(role.RoleName, "ORG_") {
				roles = append(roles, role.RoleName)
			}
		}
		apiKeys = append(apiKeys, ProjectApiKey{Key: &key.ID, RoleNames: roles})
	}
	return apiKeys
}
