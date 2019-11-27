package resource

import (
	"context"
	"fmt"

	"github.com/aws-cloudformation/aws-cloudformation-rpdk-go-plugin/cfn/encoding"
	"github.com/aws-cloudformation/aws-cloudformation-rpdk-go-plugin/cfn/handler"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	var roles []mongodbatlas.Role
	for _, r := range currentModel.Roles {
		role := mongodbatlas.Role{
			CollectionName: *r.CollectionName.Value(),
			DatabaseName:   *r.DatabaseName.Value(),
			RoleName:       *r.RoleName.Value(),
		}

		roles = append(roles, role)
	}

	groupID := *currentModel.GroupId.Value()

	_, _, err = client.DatabaseUsers.Create(context.Background(), groupID,
		&mongodbatlas.DatabaseUser{
			Roles:        roles,
			GroupID:      groupID,
			Username:     *currentModel.Username.Value(),
			Password:     *currentModel.Password.Value(),
			DatabaseName: *currentModel.Password.Value(),
			LDAPAuthType: *currentModel.LdapAuthType.Value(),
		})
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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	groupID := *currentModel.GroupId.Value()
	username := *currentModel.Username.Value()
	databaseUser, _, err := client.DatabaseUsers.Get(context.Background(), groupID, username)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error fetching database user (%s): %s", username, err)
	}

	currentModel.DatabaseName = encoding.NewString(databaseUser.Username)
	currentModel.LdapAuthType = encoding.NewString(databaseUser.LDAPAuthType)

	var roles []RoleDefinition

	for _, r := range databaseUser.Roles {
		role := RoleDefinition{
			CollectionName: encoding.NewString(r.CollectionName),
			DatabaseName:   encoding.NewString(r.DatabaseName),
			RoleName:       encoding.NewString(r.RoleName),
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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	var roles []mongodbatlas.Role
	for _, r := range currentModel.Roles {
		role := mongodbatlas.Role{
			CollectionName: *r.CollectionName.Value(),
			DatabaseName:   *r.DatabaseName.Value(),
			RoleName:       *r.RoleName.Value(),
		}

		roles = append(roles, role)
	}

	groupID := *currentModel.GroupId.Value()
	username := *currentModel.Username.Value()

	_, _, err = client.DatabaseUsers.Update(context.Background(), groupID, username,
		&mongodbatlas.DatabaseUser{
			Roles:        roles,
			GroupID:      groupID,
			Username:     username,
			Password:     *currentModel.Password.Value(),
			DatabaseName: *currentModel.Password.Value(),
			LDAPAuthType: *currentModel.LdapAuthType.Value(),
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
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	groupID := *currentModel.GroupId.Value()
	username := *currentModel.Username.Value()

	_, err = client.DatabaseUsers.Delete(context.Background(), groupID, username)
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
