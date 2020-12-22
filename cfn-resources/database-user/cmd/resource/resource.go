package resource

import (
	"context"
	"fmt"
	"log"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"go.mongodb.org/atlas/mongodbatlas"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.Printf("19 currentModel: %#+v, prevModel: %#+v", currentModel, prevModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	log.Printf("Back from clieint:  %#+v", client)
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
	log.Printf("roles: %#+v", roles)

	var labels []mongodbatlas.Label
	for _, l := range currentModel.Labels {

		label := mongodbatlas.Label{
			Key:   *l.Key,
			Value: *l.Value,
		}
		labels = append(labels, label)
	}
	log.Printf("labels: %#+v", labels)

	var scopes []mongodbatlas.Scope
	for _, l := range currentModel.Scopes {

		scope := mongodbatlas.Scope{
			Name: *l.Name,
			Type: *l.Type,
		}
		scopes = append(scopes, scope)
	}
	log.Printf("scopes: %#+v", scopes)

	groupID := *currentModel.ProjectId
	log.Printf("groupID: %#+v", groupID)

	none := "NONE"
	if currentModel.LdapAuthType == nil {
		currentModel.LdapAuthType = &none
	}

	if currentModel.AWSIAMType == nil {
		currentModel.AWSIAMType = &none
	}

	if currentModel.Password == nil {
		if (currentModel.LdapAuthType == &none) && (currentModel.AWSIAMType == &none) {
			return handler.ProgressEvent{}, fmt.Errorf("Password cannot be empty if not LDAP or IAM: %v", currentModel)
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

	projectResID := &util.ResourceIdentifier{
		ResourceType: "Project",
		ResourceID:   groupID,
	}
	resourceID := util.NewResourceIdentifier("DBUser", user.Username, projectResID)
	log.Printf("Created resourceID:%s", resourceID)

	//cfnid := fmt.Sprintf("%s-%s",pid,*currentModel.Username)
	//currentModel.UserCNFIdentifier = &cfnid
	cfnid := resourceID.String()
	currentModel.UserCNFIdentifier = &cfnid
	log.Printf("UserCFNIdentifier: %s", cfnid)

	log.Printf("Arguments: Project ID: %s, Request %#+v", groupID, user)

	newUser, _, err := client.DatabaseUsers.Create(context.Background(), groupID, user)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating database user: %s", err)
	}
	log.Printf("newUser: %s", newUser)

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
			Key:   &l.Key,
			Value: &l.Value,
		}

		labels = append(labels, label)
	}
	currentModel.Labels = labels

	cfnid := fmt.Sprintf("%v-%v", &currentModel.ProjectId, currentModel.Username)
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
			Key:   *l.Key,
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
	log.Printf("Create req:%+v, prevModel:%s, currentModel:%s", req, spew.Sdump(prevModel), spew.Sdump(currentModel))
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	groupID := *currentModel.ProjectId
	username := *currentModel.Username
	dbName := *currentModel.DatabaseName

	resp, err := client.DatabaseUsers.Delete(context.Background(), dbName, groupID, username)
	if err != nil {
		// Log and handle 404 ok
		if resp != nil && resp.StatusCode == 404 {
			log.Printf("Resource not found for Delete. Returning SUCCESS to clean orphan AWS resource. resp:%+v, error:%+v", resp, err)
		} else {
			return handler.ProgressEvent{}, fmt.Errorf("error deleting database user:%s, err:%+v, resp:%+v", username, err, resp)
		}
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
