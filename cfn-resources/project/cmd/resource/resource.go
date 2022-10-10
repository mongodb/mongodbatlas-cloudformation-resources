package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/atlas/mongodbatlas"
)

type UpdateApiKey struct {
	Key     string
	ApiKeys *mongodbatlas.AssignAPIKey
}

func setup() {
	util.SetupLogger("mongodb-atlas-project")

}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("Create log.Debugf-- currentModel: %+v", *currentModel)
	log.Debug("Create Debug whwoooo hoooo!")
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	var projectOwnerId string
	if currentModel.ProjectOwnerId != nil {
		projectOwnerId = *currentModel.ProjectOwnerId
	}
	project, _, err := client.Projects.Create(context.Background(), &mongodbatlas.Project{
		Name:                      *currentModel.Name,
		OrgID:                     *currentModel.OrgId,
		WithDefaultAlertsSettings: currentModel.WithDefaultAlertsSettings,
	}, &mongodbatlas.CreateProjectOptions{ProjectOwnerID: projectOwnerId})
	if err != nil {
		//return handler.ProgressEvent{}, fmt.Errorf("error creating project: %s", err)
		log.Debugf("Create - error: %+v", err)
		// TODO- Should detect and return HandlerErrorCodeAlreadyExists
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil

	}

	//Update project settings
	projectSettings := mongodbatlas.ProjectSettings{
		IsCollectDatabaseSpecificsStatisticsEnabled: currentModel.ProjectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
		IsRealtimePerformancePanelEnabled:           currentModel.ProjectSettings.IsRealtimePerformancePanelEnabled,
		IsDataExplorerEnabled:                       currentModel.ProjectSettings.IsDataExplorerEnabled,
		IsPerformanceAdvisorEnabled:                 currentModel.ProjectSettings.IsPerformanceAdvisorEnabled,
		IsSchemaAdvisorEnabled:                      currentModel.ProjectSettings.IsSchemaAdvisorEnabled,
	}

	client.Projects.UpdateProjectSettings(context.Background(), project.ID, &projectSettings)

	if err != nil {
		log.Debugf("Update - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Failed to update Project settings",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil

	}
	currentModel.Id = &project.ID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	//spew.Dump(currentModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		//return handler.ProgressEvent{}, err
		return handler.ProgressEvent{
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, err
	}

	name := *currentModel.Name

	if len(name) > 0 {
		project, _, err := client.Projects.GetOneProjectByName(context.Background(), name)
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Resource Not Found",
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		currentModel.Name = &project.Name
		currentModel.OrgId = &project.OrgID
		currentModel.Created = &project.Created
		currentModel.ClusterCount = &project.ClusterCount

		if err == nil {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Read Complete",
				ResourceModel:   currentModel,
			}, nil
		}
	}

	id := *currentModel.Id
	log.Debugf("Looking for project: %s", id)
	project, _, err := client.Projects.GetOneProject(context.Background(), id)
	if err != nil {
		//return handler.ProgressEvent{}, fmt.Errorf(
		log.Debugf("Read - error reading project with id(%s): %s", id, err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	currentModel.Name = &project.Name
	currentModel.OrgId = &project.OrgID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
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

	//Get teams from project
	teamsAssigned, _, err := client.Projects.GetProjectTeamsAssigned(context.Background(), projectId)
	if err != nil {
		log.Infof("ProjectId : %s, Error: %s", projectId, err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Error while finding teams in project",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	// Find the change in teams
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
			log.Debug("Error: %s", err)
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

	//Get APIKeys from project
	projectApiKeys, _, err := client.ProjectAPIKeys.List(context.Background(), projectId, &mongodbatlas.ListOptions{ItemsPerPage: 1000, IncludeCount: true})
	if err != nil {
		log.Debug("Error: %s", projectId, err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Error while finding api keys in project",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	//log.Infof("keys: %+v", currentModel.ProjectApiKeys)
	//Get Change in ApiKeys
	newApiKeys, changedKeys, removeKeys := getChangeInApiKeys(currentModel.ProjectApiKeys, projectApiKeys)

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
		//log.Infof("Removed: %s", key)

	}

	//Add Keys
	for _, key := range newApiKeys {
		_, err = client.ProjectAPIKeys.Assign(context.Background(), projectId, key.Key, key.ApiKeys)
		if err != nil {
			log.Infof("Error: %s", err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Error while Un-assigning Key to project",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		//log.Infof("Added: %s", key)

	}

	//Update Key Roles
	for _, key := range changedKeys {
		_, err = client.ProjectAPIKeys.Assign(context.Background(), projectId, key.Key, key.ApiKeys)
		if err != nil {
			log.Infof("Error: %s", err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          "Error while Un-assigning Key to project",
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		//log.Infof("Updated: %s", key)
	}

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
		log.Infof("Update - error: %+v", err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Failed to update Project settings",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil

	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil

	//return handler.ProgressEvent{
	//	OperationStatus:  handler.Failed,
	//	Message:          "Update Not Supported",
	//	ResourceModel:    nil,
	//	HandlerErrorCode: cloudformation.HandlerErrorCodeNotUpdatable,
	//}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err

	}
	log.Debug("Delete Debug whwoooo hoooo!")
	log.Debugf("Delete Project prevModel:%+v currentModel:%+v", *prevModel, *currentModel)

	var id string
	if currentModel.Id != nil {
		id = *currentModel.Id
	}

	if len(id) == 0 {
		name := *currentModel.Name
		if len(name) > 0 {
			log.Debugf("Project id was nil, try lookup name:%s", name)
			project, _, err := client.Projects.GetOneProjectByName(context.Background(), name)
			if err != nil {
				return handler.ProgressEvent{
					Message:          err.Error(),
					HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, err
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

	_, err = client.Projects.Delete(context.Background(), id)
	if err != nil {
		//return handler.ProgressEvent{}, fmt.Errorf("####error deleting project with id(%s): %s", id, err)
		log.Warnf("####error deleting project with id(%s): %s", id, err)

		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Resource Not Found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
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
	mm := []interface{}{}
	for _, project := range projects.Results {
		var m Model
		m.Name = &project.Name
		m.OrgId = &project.OrgID
		m.Created = &project.Created
		m.ClusterCount = &project.ClusterCount
		m.Id = &project.ID
		mm = append(mm, m)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  mm,
	}, nil
}

func getChangeInTeams(currentTeams []ProjectTeam, oTeams []*mongodbatlas.Result) ([]*mongodbatlas.ProjectTeam, []*mongodbatlas.ProjectTeam, []*mongodbatlas.ProjectTeam) {
	var newTeams []*mongodbatlas.ProjectTeam
	var changedTeams []*mongodbatlas.ProjectTeam
	var removeTeams []*mongodbatlas.ProjectTeam

	for _, nTeam := range currentTeams {
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

	for _, oTeam := range oTeams {
		matched := false
		for _, nTeam := range currentTeams {
			if nTeam.TeamId != nil && *nTeam.TeamId == oTeam.TeamID {
				matched = true
				break
			}
		}
		if !matched && len(currentTeams) > 0 {
			removeTeams = append(removeTeams, &mongodbatlas.ProjectTeam{TeamID: oTeam.TeamID, RoleNames: oTeam.RoleNames})
		}
	}
	return newTeams, changedTeams, removeTeams
}

func getChangeInApiKeys(currentKeys []ProjectApiKey, oKeys []mongodbatlas.APIKey) ([]UpdateApiKey, []UpdateApiKey, []UpdateApiKey) {
	var newKeys []UpdateApiKey
	var changedKeys []UpdateApiKey
	var removeKeys []UpdateApiKey

	for _, nKey := range currentKeys {
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

	for _, oKey := range oKeys {
		matched := false
		for _, nKey := range currentKeys {
			if nKey.Key != nil && *nKey.Key == oKey.ID {
				matched = true
				break
			}
		}
		if !matched && len(currentKeys) > 0 {
			removeKeys = append(removeKeys, UpdateApiKey{Key: oKey.ID})
		}
	}
	return newKeys, changedKeys, removeKeys
}
