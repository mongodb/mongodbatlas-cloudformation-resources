package resource

import (
	"context"
	"fmt"
	"log"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	var roles []mongodbatlas.Role
	for _, r := range currentModel.Roles {

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

	groupID := *currentModel.ProjectId

	user := &mongodbatlas.DatabaseUser{
		Roles:        roles,
		GroupID:      groupID,
		Username:     *currentModel.Username,
		Password:     *currentModel.Password,
		DatabaseName: *currentModel.DatabaseName,
	}

	if currentModel.LdapAuthType != nil {
		user.LDAPAuthType = *currentModel.LdapAuthType
	}

	log.Printf("Arguments: Project ID: %s, Request %#+v", groupID, user)

	_, _, err = client.DatabaseUsers.Create(context.Background(), groupID, user)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating database user: %s", err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	groupID := *currentModel.ProjectId
	username := *currentModel.Username
	dbName := *currentModel.DatabaseName
	databaseUser, _, err := client.DatabaseUsers.Get(context.Background(), dbName, groupID, username)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error fetching database user (%s): %s", username, err)
	}

	currentModel.DatabaseName = &databaseUser.Username
	currentModel.LdapAuthType = &databaseUser.LDAPAuthType

	var roles []RoleDefinition

	for _, r := range databaseUser.Roles {
		role := RoleDefinition{
			CollectionName: &r.CollectionName,
			DatabaseName:   &r.DatabaseName,
			RoleName:       &r.RoleName,
		}

		roles = append(roles, role)
	}
	currentModel.Roles = roles

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	var roles []mongodbatlas.Role
	for _, r := range currentModel.Roles {
		role := mongodbatlas.Role{
			CollectionName: *r.CollectionName,
			DatabaseName:   *r.DatabaseName,
			RoleName:       *r.RoleName,
		}

		roles = append(roles, role)
	}

	groupID := *currentModel.ProjectId
	username := *currentModel.Username

	_, _, err = client.DatabaseUsers.Update(context.Background(), groupID, username,
		&mongodbatlas.DatabaseUser{
			Roles:        roles,
			GroupID:      groupID,
			Username:     username,
			Password:     *currentModel.Password,
			DatabaseName: *currentModel.Password,
			LDAPAuthType: *currentModel.LdapAuthType,
		})
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error updating database user (%s): %s", username, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	groupID := *currentModel.ProjectId
	username := *currentModel.Username
	dbName := *currentModel.DatabaseName

	_, err = client.DatabaseUsers.Delete(context.Background(), dbName, groupID, username)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting database user (%s): %s", username, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   currentModel,
	}, nil
}

// List NOOP
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   currentModel,
	}, nil
}
