package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progress_event"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/atlas/mongodbatlas"
)

func setup() {
	util.SetupLogger("mongodb-atlas-database-user")
}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

var CreateRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "DatabaseName", "ProjectId", "Roles", "Username"}
var ReadRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId", "DatabaseName", "Username"}
var UpdateRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "DatabaseName", "ProjectId", "Roles", "Username"}
var DeleteRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId", "DatabaseName", "Username"}
var ListRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId"}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf(" currentModel: %#+v, prevModel: %#+v", currentModel, prevModel)

	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	var roles []mongodbatlas.Role
	for i, _ := range currentModel.Roles {
		r := currentModel.Roles[i]
		role := mongodbatlas.Role{}
		if r.CollectionName != nil {
			role.CollectionName = *r.CollectionName
		}

		if r.DatabaseName != nil {
			role.DatabaseName = *r.DatabaseName
		}

		if r.RoleName != nil {
			role.RoleName = *r.RoleName
		}

		roles = append(roles, role)
	}
	log.Debugf("roles: %#+v", roles)

	var labels []mongodbatlas.Label
	for i, _ := range currentModel.Labels {
		l := currentModel.Labels[i]
		label := mongodbatlas.Label{
			Key:   *l.Key,
			Value: *l.Value,
		}
		labels = append(labels, label)
	}
	log.Debugf("labels: %#+v", labels)

	var scopes []mongodbatlas.Scope
	for i, _ := range currentModel.Scopes {
		s := currentModel.Scopes[i]
		scope := mongodbatlas.Scope{
			Name: *s.Name,
			Type: *s.Type,
		}
		scopes = append(scopes, scope)
	}
	log.Debugf("scopes: %#+v", scopes)

	groupID := *currentModel.ProjectId
	log.Debugf("groupID: %#+v", groupID)

	none := "NONE"
	if currentModel.LdapAuthType == nil {
		currentModel.LdapAuthType = &none
	}

	if currentModel.AWSIAMType == nil {
		currentModel.AWSIAMType = &none
	}

	if currentModel.X509Type == nil {
		currentModel.X509Type = &none
	}

	if currentModel.Password == nil {
		if (currentModel.LdapAuthType == &none) && (currentModel.AWSIAMType == &none) && (currentModel.X509Type == &none) {
			err := fmt.Errorf("Password cannot be empty if not LDAP or IAM or X509: %v", currentModel)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		s := ""
		currentModel.Password = &s
	}

	if (currentModel.X509Type != &none) || (currentModel.DeleteAfterDate == nil) {
		s := ""
		currentModel.DeleteAfterDate = &s
	}

	log.Printf("Check Delete after date here::???????")
	spew.Dump(currentModel)

	user := getDbUser(roles, groupID, currentModel, labels, scopes)

	log.Debugf("Arguments: Project ID: %s, Request %#+v", groupID, user)

	newUser, res, err := client.DatabaseUsers.Create(context.Background(), groupID, user)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			res.Response), nil
	}
	log.Debugf("newUser: %s", newUser)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	groupID := *currentModel.ProjectId
	username := *currentModel.Username
	dbName := *currentModel.DatabaseName
	databaseUser, resp, err := client.DatabaseUsers.Get(context.Background(), dbName, groupID, username)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	currentModel.DatabaseName = &databaseUser.DatabaseName

	//currentModel.LdapAuthType = &databaseUser.LDAPAuthType
	if currentModel.LdapAuthType != nil {
		currentModel.LdapAuthType = &databaseUser.LDAPAuthType
	}

	if currentModel.AWSIAMType != nil {
		currentModel.AWSIAMType = &databaseUser.AWSIAMType
	}

	if currentModel.X509Type != nil {
		currentModel.X509Type = &databaseUser.X509Type
	}

	currentModel.Username = &databaseUser.Username

	log.Debugf("databaseUser:%+v", databaseUser)
	var roles []RoleDefinition

	for i, _ := range databaseUser.Roles {
		r := databaseUser.Roles[i]
		role := RoleDefinition{
			CollectionName: &r.CollectionName,
			DatabaseName:   &r.DatabaseName,
			RoleName:       &r.RoleName,
		}

		roles = append(roles, role)
	}
	currentModel.Roles = roles
	log.Debugf("currentModel.Roles:%+v", roles)
	var labels []LabelDefinition

	for i, _ := range databaseUser.Labels {
		l := databaseUser.Labels[i]
		label := LabelDefinition{
			Key:   &l.Key,
			Value: &l.Value,
		}

		labels = append(labels, label)
	}
	currentModel.Labels = labels

	cfnid := fmt.Sprintf("%s-%s", *currentModel.Username, groupID)
	currentModel.UserCFNIdentifier = &cfnid
	log.Debugf("READ----> currentModel:%s", spew.Sdump(currentModel))
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)

	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	log.Debugf("Update currentModel:%+v", currentModel)
	roles := []mongodbatlas.Role{}
	for i, _ := range currentModel.Roles {
		r := currentModel.Roles[i]
		role := mongodbatlas.Role{}
		if r.CollectionName != nil {
			role.CollectionName = *r.CollectionName
		} else {
			role.CollectionName = ""
		}
		role.DatabaseName = *r.DatabaseName
		role.RoleName = *r.RoleName
		roles = append(roles, role)
	}

	log.Debugf("Update roles:%+v", roles)
	labels := []mongodbatlas.Label{}
	for i, _ := range currentModel.Labels {
		l := currentModel.Labels[i]
		label := mongodbatlas.Label{
			Key:   *l.Key,
			Value: *l.Value,
		}
		labels = append(labels, label)
	}
	log.Debugf("Update labels: %#+v", labels)

	var scopes []mongodbatlas.Scope
	for i, _ := range currentModel.Scopes {
		s := currentModel.Scopes[i]
		scope := mongodbatlas.Scope{
			Name: *s.Name,
			Type: *s.Type,
		}
		scopes = append(scopes, scope)
	}
	log.Debugf("Update scopes: %#+v", scopes)

	groupID := *currentModel.ProjectId
	log.Debugf("groupID: %#+v", groupID)

	none := "NONE"
	if currentModel.LdapAuthType == nil {
		currentModel.LdapAuthType = &none
	}

	if currentModel.AWSIAMType == nil {
		currentModel.AWSIAMType = &none
	}

	if currentModel.X509Type == nil {
		currentModel.X509Type = &none
	}

	if currentModel.Password == nil {
		if (currentModel.LdapAuthType == &none) && (currentModel.AWSIAMType == &none) && (currentModel.X509Type == &none) {
			err := fmt.Errorf("Password cannot be empty if not LDAP or IAM or X509: %v", currentModel)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		s := ""
		currentModel.Password = &s
	}

	if (currentModel.X509Type != &none) || (currentModel.DeleteAfterDate == nil) {
		s := ""
		currentModel.DeleteAfterDate = &s
	}

	log.Printf("Check Delete after date here::???????")
	spew.Dump(currentModel)

	dbu := getDbUser(roles, groupID, currentModel, labels, scopes)
	log.Debugf("dbu:%+v", dbu)
	_, resp, err := client.DatabaseUsers.Update(context.Background(), groupID, *currentModel.Username, dbu)

	log.Debugf("Update resp:%+v", resp)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

func getDbUser(roles []mongodbatlas.Role, groupID string, currentModel *Model, labels []mongodbatlas.Label, scopes []mongodbatlas.Scope) *mongodbatlas.DatabaseUser {
	return &mongodbatlas.DatabaseUser{
		Roles:           roles,
		GroupID:         groupID,
		Username:        *currentModel.Username,
		Password:        *currentModel.Password,
		DatabaseName:    *currentModel.DatabaseName,
		Labels:          labels,
		Scopes:          scopes,
		LDAPAuthType:    *currentModel.LdapAuthType,
		AWSIAMType:      *currentModel.AWSIAMType,
		X509Type:        *currentModel.X509Type,
		DeleteAfterDate: *currentModel.DeleteAfterDate,
	}
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	log.Debugf("Create req:%+v, prevModel:%s, currentModel:%s", req, spew.Sdump(prevModel), spew.Sdump(currentModel))

	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		//return handler.ProgressEvent{}, err
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	groupID := *currentModel.ProjectId
	username := *currentModel.Username
	dbName := *currentModel.DatabaseName

	resp, err := client.DatabaseUsers.Delete(context.Background(), dbName, groupID, username)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

// List NOOP
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	groupID := *currentModel.ProjectId
	dbUserModels := []interface{}{}

	databaseUsers, resp, err := client.DatabaseUsers.List(context.Background(), groupID, nil)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
	}

	for i, _ := range databaseUsers {
		var model Model
		databaseUser := databaseUsers[i]
		model.DatabaseName = &databaseUser.DatabaseName
		model.LdapAuthType = &databaseUser.LDAPAuthType
		model.X509Type = &databaseUser.X509Type
		model.Username = &databaseUser.Username
		var roles []RoleDefinition

		for i, _ := range databaseUser.Roles {
			r := databaseUser.Roles[i]
			role := RoleDefinition{
				CollectionName: &r.CollectionName,
				DatabaseName:   &r.DatabaseName,
				RoleName:       &r.RoleName,
			}

			roles = append(roles, role)
		}
		model.Roles = roles

		var labels []LabelDefinition

		for i, _ := range databaseUser.Labels {
			l := databaseUser.Labels[i]
			label := LabelDefinition{
				Key:   &l.Key,
				Value: &l.Value,
			}

			labels = append(labels, label)
		}
		model.Labels = labels

		cfnid := fmt.Sprintf("%s-%s", databaseUser.Username, databaseUser.GroupID)
		model.UserCFNIdentifier = &cfnid

		dbUserModels = append(dbUserModels, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  dbUserModels,
	}, nil
}
