package resource

import (
	"context"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	localConstants "github.com/mongodb/mongodbatlas-cloudformation-resources/ldap-configuration/cmd/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressEvents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.GroupID, localConstants.AuthenticationEnabled,
	localConstants.BindUsername, localConstants.Hostname, localConstants.Port, localConstants.UserToDNMapping}
var ReadRequiredFields = []string{constants.GroupID, constants.PubKey, constants.PvtKey}
var UpdateRequiredFields = []string{constants.GroupID, constants.PubKey, constants.PvtKey}
var DeleteRequiredFields = []string{constants.GroupID, constants.PubKey, constants.PvtKey}
var ListRequiredFields = []string{constants.PubKey, constants.PvtKey}

func setup() {
	util.SetupLogger("mongodb-atlas-ldap-configuration")
}

func (m *Model) CompleteByResponse(resp mongodbatlas.LDAPConfiguration) {
	m.BindUsername = &resp.LDAP.BindUsername
	m.Hostname = &resp.LDAP.Hostname
	m.AuthenticationEnabled = &resp.LDAP.AuthenticationEnabled
	m.AuthorizationEnabled = &resp.LDAP.AuthorizationEnabled
	m.CaCertificate = &resp.LDAP.CaCertificate
	m.AuthzQueryTemplate = &resp.LDAP.AuthzQueryTemplate
	m.BindPassword = &resp.LDAP.BindPassword

	mapping := make([]ApiAtlasNDSUserToDNMappingView, len(resp.LDAP.UserToDNMapping))

	for i := range resp.LDAP.UserToDNMapping {
		ndsMap := ApiAtlasNDSUserToDNMappingView{
			Match:        &resp.LDAP.UserToDNMapping[i].Match,
			Substitution: &resp.LDAP.UserToDNMapping[i].Substitution,
			LdapQuery:    &resp.LDAP.UserToDNMapping[i].LDAPQuery,
		}
		mapping = append(mapping, ndsMap)
	}
	m.UserToDNMapping = mapping
}

func getUserToDNMapping(ndsUserMapping []ApiAtlasNDSUserToDNMappingView) []*mongodbatlas.UserToDNMapping {
	mapping := make([]*mongodbatlas.UserToDNMapping, len(ndsUserMapping))

	for i := range ndsUserMapping {
		ndsMap := mongodbatlas.UserToDNMapping{
			Match:        *ndsUserMapping[i].Match,
			Substitution: *ndsUserMapping[i].Substitution,
			LDAPQuery:    *ndsUserMapping[i].LdapQuery,
		}
		mapping = append(mapping, &ndsMap)
	}

	return mapping
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	DNMapping := getUserToDNMapping(currentModel.UserToDNMapping)

	ldap := &mongodbatlas.LDAP{
		AuthenticationEnabled: *currentModel.AuthenticationEnabled,
		Hostname:              *currentModel.Hostname,
		Port:                  *currentModel.Port,
		BindUsername:          *currentModel.BindUsername,
		UserToDNMapping:       DNMapping,
		BindPassword:          "",
	}

	ldapReq := &mongodbatlas.LDAPConfiguration{
		LDAP: ldap,
	}

	LDAPConfigResponse, res, err := client.LDAPConfigurations.Save(context.Background(), *currentModel.GroupId, ldapReq)
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progressEvents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.CompleteByResponse(*LDAPConfigResponse)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Read() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Read - error: %+v", err)
		return progressEvents.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	var res *mongodbatlas.Response

	if err != nil {
		log.Debugf("Read - error: %+v", err)
		return progressEvents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Update() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Update - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	/*
	   Considerable params from currentModel:
	   ApiKeys, GroupId, Links, Links, CustomerX509, Links, Ldap, ...
	*/
	/*
	    // Pseudocode:
	    res , resModel, err := client.Ldapconfiguration.Update(context.Background(),&mongodbatlas.Ldapconfiguration{
	   	ApiKeys:currentModel.ApiKeys,
	   	GroupId:currentModel.GroupId,
	   	Links:currentModel.Links,
	   	Links:currentModel.Links,
	   	CustomerX509:currentModel.CustomerX509,
	   	Links:currentModel.Links,
	   	Ldap:currentModel.Ldap,
	   })

	*/

	if err != nil {
		log.Debugf("Update - error: %+v", err)
		return progressEvents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Delete() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Delete - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	//
	/*
	    // Pseudocode:
	    res , resModel, err := client.Ldapconfiguration.Delete(context.Background(),&mongodbatlas.Ldapconfiguration{
	   })

	*/

	if err != nil {
		log.Debugf("Delete - error: %+v", err)
		return progressEvents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("List - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	//
	/*
	    // Pseudocode:
	    res , resModel, err := client.Ldapconfiguration.List(context.Background(),&mongodbatlas.Ldapconfiguration{
	   })

	*/

	if err != nil {
		log.Debugf("List - error: %+v", err)
		return progressEvents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		Message:         "Delete success",
		OperationStatus: handler.Success,
	}, nil
}
