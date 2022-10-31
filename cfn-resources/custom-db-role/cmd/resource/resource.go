package resource

import (
	"context"
	"fmt"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progress_event"
	"net/http"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/custom-db-role/cmd/validator_def"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
)

func setup() {
	util.SetupLogger("mongodb-atlas-custom-db-role")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	modelValidation := validator.ValidateModel(constants.Create, validator_def.ModelValidator{}, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	mongodbClient, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	atlasCustomDBRole := currentModel.ToCustomDBRole()

	customDBRole, response, err := mongodbClient.CustomDBRoles.Create(context.Background(), *currentModel.GroupId, &atlasCustomDBRole)
	if err != nil {
		if response.Response.StatusCode == http.StatusConflict {
			return progress_events.GetFailedEventByCode("Resource already exists",
				cloudformation.HandlerErrorCodeAlreadyExists), nil
		}

		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error creating resource : %s", err.Error()),
			response.Response), nil
	}

	currentModel.completeByAtlasRole(*customDBRole)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   currentModel}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	modelValidation := validator.ValidateModel(constants.Read, validator_def.ModelValidator{}, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	mongodbClient, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	atlasCustomDdRole, response, err := mongodbClient.CustomDBRoles.Get(context.Background(), *currentModel.GroupId, *currentModel.RoleName)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response), nil
	}

	currentModel.completeByAtlasRole(*atlasCustomDdRole)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Get successful",
		ResourceModel:   currentModel}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	modelValidation := validator.ValidateModel(constants.Update, validator_def.ModelValidator{}, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	mongodbClient, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	var actions []mongodbatlas.Action
	for _, a := range currentModel.Actions {
		actions = append(actions, a.toAtlasAction())
	}

	var inheritedRoles []mongodbatlas.InheritedRole
	for _, ir := range currentModel.InheritedRoles {
		inheritedRoles = append(inheritedRoles, ir.toAtlasInheritedRole())
	}

	inputCustomDBRole := mongodbatlas.CustomDBRole{
		Actions:        actions,
		InheritedRoles: inheritedRoles,
	}

	atlasCustomDdRole, response, err := mongodbClient.CustomDBRoles.Update(context.Background(), *currentModel.GroupId,
		*currentModel.RoleName, &inputCustomDBRole)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response), nil
	}

	currentModel.completeByAtlasRole(*atlasCustomDdRole)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update successful",
		ResourceModel:   currentModel}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	modelValidation := validator.ValidateModel(constants.Delete, validator_def.ModelValidator{}, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	mongodbClient, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	response, err := mongodbClient.CustomDBRoles.Delete(context.Background(), *currentModel.GroupId, *currentModel.RoleName)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error deleting resource : %s", err.Error()),
			response.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete success"}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	modelValidation := validator.ValidateModel(constants.List, validator_def.ModelValidator{}, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	mongodbClient, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}

	customDBRoleResponse, response, err := mongodbClient.CustomDBRoles.List(context.Background(),
		*currentModel.GroupId,
		params)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error listing resource : %s", err.Error()),
			response.Response), nil
	}

	mm := make([]interface{}, 0)
	for _, customDBRole := range *customDBRoleResponse {
		var m Model
		m.completeByAtlasRole(customDBRole)
		mm = append(mm, m)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List successful",
		ResourceModels:  mm}, nil
}

func (m *Model) ToCustomDBRole() mongodbatlas.CustomDBRole {
	var actions []mongodbatlas.Action
	for _, a := range m.Actions {
		actions = append(actions, a.toAtlasAction())
	}

	var inheritedRoles []mongodbatlas.InheritedRole
	for _, ir := range m.InheritedRoles {
		inheritedRoles = append(inheritedRoles, ir.toAtlasInheritedRole())
	}

	return mongodbatlas.CustomDBRole{
		Actions:        actions,
		InheritedRoles: inheritedRoles,
		RoleName:       *m.RoleName,
	}
}

func (a InheritedRole) toAtlasInheritedRole() mongodbatlas.InheritedRole {
	return mongodbatlas.InheritedRole{
		Db:   *a.Db,
		Role: *a.Role,
	}
}

func (a Action) toAtlasAction() mongodbatlas.Action {
	var resources []mongodbatlas.Resource
	for _, r := range a.Resources {
		resources = append(resources, r.toAtlasResource())
	}

	return mongodbatlas.Action{
		Action:    *a.Action,
		Resources: resources,
	}
}

func (r Resource) toAtlasResource() mongodbatlas.Resource {
	return mongodbatlas.Resource{
		Collection: r.Collection,
		DB:         r.DB,
		Cluster:    r.Cluster,
	}
}

func (m *Model) completeByAtlasRole(role mongodbatlas.CustomDBRole) {
	var actions []Action
	for _, a := range role.Actions {
		actions = append(actions, atlasActionToModel(a))
	}

	var inheritedRoles []InheritedRole
	for _, ir := range role.InheritedRoles {
		inheritedRoles = append(inheritedRoles, atlasInheritedRoleToModel(ir))
	}

	m.Actions = actions
	m.InheritedRoles = inheritedRoles
	m.RoleName = &role.RoleName
}

func atlasActionToModel(action mongodbatlas.Action) Action {
	var resources []Resource
	for _, r := range action.Resources {
		resources = append(resources, atlasResourceToModel(r))
	}

	return Action{
		Action:    &action.Action,
		Resources: resources,
	}
}

func atlasResourceToModel(resource mongodbatlas.Resource) Resource {
	return Resource{
		Collection: resource.Collection,
		DB:         resource.DB,
		Cluster:    resource.Cluster,
	}
}

func atlasInheritedRoleToModel(inheritedRole mongodbatlas.InheritedRole) InheritedRole {
	return InheritedRole{
		Db:   &inheritedRole.Db,
		Role: &inheritedRole.Role,
	}
}
