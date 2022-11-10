package resource

import (
	"context"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"go.mongodb.org/atlas/mongodbatlas"
)

func setup() {
	util.SetupLogger("mongodb-atlas-database-user")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	_, _ = logger.Debugf(" currentModel: %#+v, prevModel: %#+v", currentModel, prevModel)

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	var roles []mongodbatlas.Role
	for i := range currentModel.Roles {
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
	_, _ = logger.Debugf("roles: %#+v", roles)

	var labels []mongodbatlas.Label
	for i := range currentModel.Labels {
		l := currentModel.Labels[i]
		label := mongodbatlas.Label{
			Key:   *l.Key,
			Value: *l.Value,
		}
		labels = append(labels, label)
	}
	_, _ = logger.Debugf("labels: %#+v", labels)

	var scopes []mongodbatlas.Scope
	for i := range currentModel.Scopes {
		s := currentModel.Scopes[i]
		scope := mongodbatlas.Scope{
			Name: *s.Name,
			Type: *s.Type,
		}
		scopes = append(scopes, scope)
	}
	_, _ = logger.Debugf("scopes: %#+v", scopes)

	groupID := *currentModel.ProjectId
	_, _ = logger.Debugf("groupID: %#+v", groupID)

	none := "NONE"
	if currentModel.LdapAuthType == nil {
		currentModel.LdapAuthType = &none
	}

	if currentModel.AWSIAMType == nil {
		currentModel.AWSIAMType = &none
	}

	if currentModel.Password == nil {
		if (currentModel.LdapAuthType == &none) && (currentModel.AWSIAMType == &none) {
			err = fmt.Errorf("password cannot be empty if not LDAP or IAM: %v", currentModel)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		s := ""
		currentModel.Password = &s
	}

	user := &mongodbatlas.DatabaseUser{
		Roles:        roles,
		GroupID:      groupID,
		Username:     *currentModel.Username,
		Password:     *currentModel.Password,
		DatabaseName: *currentModel.DatabaseName,
		Labels:       labels,
		Scopes:       scopes,
		LDAPAuthType: *currentModel.LdapAuthType,
		AWSIAMType:   *currentModel.AWSIAMType,
	}

	cfnid := fmt.Sprintf("%s-%s", user.Username, groupID)
	currentModel.UserCFNIdentifier = &cfnid
	_, _ = logger.Debugf("Created UserCFNIdentifier: %s", cfnid)

	_, _ = logger.Debugf("Arguments: Project ID: %s, Request %#+v", groupID, user)

	newUser, res, err := client.DatabaseUsers.Create(context.Background(), groupID, user)
	if err != nil {
		_, _ = logger.Debugf("Error creating new db user: res:%+v, err:%+v", res, err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	_, _ = logger.Debugf("newUser: %s", newUser)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
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
		_, _ = logger.Debugf("error fetching database groupID:%s, dbName:%s user:%s, error: %s", groupID, dbName, username, err)
		if resp != nil && resp.StatusCode == 404 {
			_, _ = logger.Debugf("Resource Not Found 404 for READ groupId:%s, dbName:%s, database user:%s, err:%+v, resp:%+v", groupID, dbName, username, err, resp)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}

		_, _ = logger.Debugf("Error READ groupId:%s, dbName:%s, database user:%s, err:%+v, resp:%+v", groupID, dbName, username, err, resp)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}

	currentModel.DatabaseName = &databaseUser.DatabaseName

	if currentModel.LdapAuthType != nil {
		currentModel.LdapAuthType = &databaseUser.LDAPAuthType
	}

	if currentModel.AWSIAMType != nil {
		currentModel.AWSIAMType = &databaseUser.AWSIAMType
	}

	currentModel.Username = &databaseUser.Username

	_, _ = logger.Debugf("databaseUser:%+v", databaseUser)
	var roles []RoleDefinition

	for i := range databaseUser.Roles {
		r := databaseUser.Roles[i]
		role := RoleDefinition{
			CollectionName: &r.CollectionName,
			DatabaseName:   &r.DatabaseName,
			RoleName:       &r.RoleName,
		}

		roles = append(roles, role)
	}
	currentModel.Roles = roles
	_, _ = logger.Debugf("currentModel.Roles:%+v", roles)
	var labels []LabelDefinition

	for i := range databaseUser.Labels {
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
	_, _ = logger.Debugf("READ----> currentModel:%s", spew.Sdump(currentModel))
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

	_, _ = logger.Debugf("Update currentModel:%+v", currentModel)
	roles := []mongodbatlas.Role{}
	for i := range currentModel.Roles {
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

	_, _ = logger.Debugf("Update roles:%+v", roles)
	labels := []mongodbatlas.Label{}
	for i := range currentModel.Labels {
		l := currentModel.Labels[i]
		label := mongodbatlas.Label{
			Key:   *l.Key,
			Value: *l.Value,
		}
		labels = append(labels, label)
	}

	if currentModel.LdapAuthType == nil {
		currentModel.LdapAuthType = new(string)
	}
	groupID := *currentModel.ProjectId
	username := *currentModel.Username
	_, _ = logger.Debugf("groupID:%s, username:%s", groupID, username)
	dbu := &mongodbatlas.DatabaseUser{
		Roles:        roles,
		GroupID:      groupID,
		Username:     username,
		Password:     *currentModel.Password,
		DatabaseName: *currentModel.DatabaseName,
		LDAPAuthType: *currentModel.LdapAuthType,
		Labels:       labels,
	}
	_, _ = logger.Debugf("dbu:%+v", dbu)
	_, resp, err := client.DatabaseUsers.Update(context.Background(), groupID, username, dbu)

	_, _ = logger.Debugf("Update resp:%+v", resp)
	if err != nil {
		_, _ = logger.Debugf("Error Update database user:%s, error: %s", username, err)
		if resp != nil && resp.StatusCode == 404 {
			_, _ = logger.Warnf("Resource Not Found 404 for UPDATE groupId:%s, database user:%s, err:%+v, resp:%+v", groupID, username, err, resp)
			return handler.ProgressEvent{
				Message:          err.Error(),
				OperationStatus:  handler.Failed,
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = logger.Warnf("Error UPDATE groupId:%s, database user:%s, err:%+v, resp:%+v", groupID, username, err, resp)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	_, _ = logger.Debugf("Create req:%+v, prevModel:%s, currentModel:%s", req, spew.Sdump(prevModel), spew.Sdump(currentModel))
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

	resp, err := client.DatabaseUsers.Delete(context.Background(), dbName, groupID, username)
	if err != nil {
		// Log and handle 404 ok
		if resp != nil && resp.StatusCode == 404 {
			_, _ = logger.Warnf("Resource not found for Delete. resp:%+v, error:%+v", resp, err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
		}
		_, _ = logger.Warnf("Error deleting database user:%s, err:%+v, resp:%+v", username, err, resp)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

// List NOOP
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	groupID := *currentModel.ProjectId
	var dbUserModels []interface{}

	databaseUsers, _, err := client.DatabaseUsers.List(context.Background(), groupID, nil)
	if err != nil {
		_, _ = logger.Debugf("error fetching database users groupId%s, error: %s", groupID, err)
		return handler.ProgressEvent{
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeServiceInternalError}, nil
	}

	for i := range databaseUsers {
		var model Model
		databaseUser := databaseUsers[i]
		model.DatabaseName = &databaseUser.DatabaseName
		model.LdapAuthType = &databaseUser.LDAPAuthType
		model.Username = &databaseUser.Username
		var roles []RoleDefinition

		for i := range databaseUser.Roles {
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

		for i := range databaseUser.Labels {
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
