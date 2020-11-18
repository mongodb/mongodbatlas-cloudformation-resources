package resource

import (
	"context"
	"fmt"
	"log"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
    "go.mongodb.org/atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
    log.Printf("19 currentModel: %s, prevModel: %#+v", currentModel, prevModel)

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

    var labels []mongodbatlas.Label
	for _, l := range currentModel.Labels {

		label := mongodbatlas.Label{
            Key: *l.Key,
            Value: *l.Value,
        }
		labels = append(labels, label)
    }

    var scopes []mongodbatlas.Scope
	for _, l := range currentModel.Scopes {

		scope := mongodbatlas.Scope{
            Name: *l.Name,
            Type: *l.Type,
        }
		scopes = append(scopes, scope)
    }

    groupID := *currentModel.ProjectId

	user := &mongodbatlas.DatabaseUser{
		Roles:        roles,
		GroupID:      groupID,
		Username:     *currentModel.Username,
		Password:     *currentModel.Password,
		DatabaseName: *currentModel.DatabaseName,
        Labels:       labels,
        Scopes:       scopes,
	}

	if currentModel.LdapAuthType != nil {
		user.LDAPAuthType = *currentModel.LdapAuthType
	}
	if currentModel.AWSIAMType != nil {
		user.AWSIAMType = *currentModel.AWSIAMType
	}

	log.Printf("Arguments: Project ID: %s, Request %#+v", groupID, user)
    pid := currentModel.ProjectId
    cfnid := fmt.Sprintf("%v-%v",pid,currentModel.Username)
    currentModel.UserCNFIdentifier = &cfnid
    log.Printf("UserCFNIdentifier: %s",cfnid)

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

	var labels []LabelDefinition

	for _, l := range databaseUser.Labels {
		label := LabelDefinition{
			Key: &l.Key,
			Value:   &l.Value,
		}

		labels = append(labels, label)
	}
	currentModel.Labels = labels

    cfnid := fmt.Sprintf("%v-%v",&currentModel.ProjectId,currentModel.Username)
    currentModel.UserCNFIdentifier = &cfnid

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
    var labels []mongodbatlas.Label
	for _, l := range currentModel.Labels {

		label := mongodbatlas.Label{
            Key: *l.Key,
            Value: *l.Value,
        }
		labels = append(labels, label)
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
			Labels:       labels,
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
