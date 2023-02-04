// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"context"
	"errors"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressEvents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/openlyinc/pointy"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	BindPassword = "BindPassword"
	BindUsername = "BindUsername"
	Hostname     = "Hostname"
	Port         = "Port"
)

var CreateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.GroupID,
	BindUsername, BindPassword, Hostname, Port}
var ReadRequiredFields = []string{constants.GroupID, constants.PubKey, constants.PvtKey}
var UpdateRequiredFields = []string{constants.GroupID, constants.PubKey, constants.PvtKey}
var DeleteRequiredFields = []string{constants.GroupID, constants.PubKey, constants.PvtKey}
var ListRequiredFields = []string{constants.PubKey, constants.PvtKey}

func setup() {
	util.SetupLogger("mongodb-atlas-ldap-configuration")
}

func (m *Model) CompleteByResponse(resp mongodbatlas.LDAPConfiguration) {
	m.AuthenticationEnabled = resp.LDAP.AuthenticationEnabled
	m.AuthorizationEnabled = resp.LDAP.AuthorizationEnabled

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

func (m *Model) GetAtlasModel() *mongodbatlas.LDAPConfiguration {
	DNMapping := getUserToDNMapping(m.UserToDNMapping)

	ldap := &mongodbatlas.LDAP{
		AuthenticationEnabled: pointy.Bool(true),
		Hostname:              m.Hostname,
		Port:                  m.Port,
		BindUsername:          m.BindUsername,
		UserToDNMapping:       DNMapping,
		BindPassword:          m.BindPassword,
	}

	ldapReq := &mongodbatlas.LDAPConfiguration{
		LDAP: ldap,
	}

	if m.AuthzQueryTemplate != nil {
		ldapReq.LDAP.AuthzQueryTemplate = m.AuthzQueryTemplate
	}

	if m.CaCertificate != nil {
		ldapReq.LDAP.CaCertificate = m.CaCertificate
	}

	if m.AuthorizationEnabled != nil {
		ldapReq.LDAP.AuthorizationEnabled = m.AuthorizationEnabled
	}

	return ldapReq
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
	if modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progressEvents.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	enabled := true

	currentModel.AuthenticationEnabled = &enabled

	ldapReq := currentModel.GetAtlasModel()

	LDAPConfigResponse, res, err := client.LDAPConfigurations.Save(context.Background(), *currentModel.GroupId, ldapReq)
	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progressEvents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.CompleteByResponse(*LDAPConfigResponse)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create successfully",
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Read() currentModel:%+v", currentModel)

	// Validation
	if modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Read - error: %+v", err)
		return progressEvents.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	ldapConf, errPe := Get(client, *currentModel.GroupId)
	if errPe != nil {
		return *errPe, nil
	}

	currentModel.CompleteByResponse(*ldapConf)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func isResourceEnabled(ldapConf mongodbatlas.LDAPConfiguration) bool {
	if ldapConf.LDAP.AuthenticationEnabled != nil {
		return *ldapConf.LDAP.AuthenticationEnabled
	}
	return false
}

func Get(client *mongodbatlas.Client, groupID string) (*mongodbatlas.LDAPConfiguration, *handler.ProgressEvent) {
	ldapConf, res, err := client.LDAPConfigurations.Get(context.Background(), groupID)
	if err != nil {
		errPe := progressEvents.GetFailedEventByResponse(err.Error(), res.Response)
		return nil, &errPe
	}

	if !isResourceEnabled(*ldapConf) {
		errPe := progressEvents.GetFailedEventByCode("Authentication is disabled for the selected project", cloudformation.HandlerErrorCodeNotFound)
		return nil, &errPe
	}

	return ldapConf, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	if modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Update - error: %+v", err)
		return progressEvents.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// Validate if resource exists
	_, errPe := Get(client, *currentModel.GroupId)
	if errPe != nil {
		return *errPe, nil
	}

	ldapReq := currentModel.GetAtlasModel()

	LDAPConfigResponse, res, err := client.LDAPConfigurations.Save(context.Background(), *currentModel.GroupId, ldapReq)
	if err != nil {
		_, _ = log.Debugf("Update - error: %+v", err)
		return progressEvents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.CompleteByResponse(*LDAPConfigResponse)

	if err != nil {
		_, _ = log.Debugf("Update - error: %+v", err)
		return progressEvents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	if modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Debugf("Delete - error: %+v", err)
		return progressEvents.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	// Validate if resource exists
	_, errPe := Get(client, *currentModel.GroupId)
	if errPe != nil {
		return *errPe, nil
	}

	ldapReq := currentModel.GetAtlasModel()
	ldapReq.LDAP.AuthorizationEnabled = pointy.Bool(false)
	ldapReq.LDAP.AuthenticationEnabled = pointy.Bool(false)

	_, res, err := client.LDAPConfigurations.Save(context.Background(), *currentModel.GroupId, ldapReq)
	if err != nil {
		_, _ = log.Debugf("Update - error: %+v", err)
		return progressEvents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete success",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}
