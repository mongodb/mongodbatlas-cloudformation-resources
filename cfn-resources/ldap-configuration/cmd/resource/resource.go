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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"

	"go.mongodb.org/atlas-sdk/v20230201008/admin"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	BindPassword = "BindPassword"
	BindUsername = "BindUsername"
	Hostname     = "Hostname"
	Port         = "Port"
)

var CreateRequiredFields = []string{constants.ProjectID,
	BindUsername, BindPassword, Hostname, Port}
var ReadRequiredFields = []string{constants.ProjectID}
var UpdateRequiredFields = []string{constants.ProjectID}
var DeleteRequiredFields = []string{constants.ProjectID}
var ListRequiredFields []string

func setup() {
	util.SetupLogger("mongodb-atlas-ldap-configuration")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(CreateRequiredFields, currentModel); err != nil {
		return *err, nil
	}
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}
	ctx := context.Background()
	ldapConf, resp, err := client.AtlasV2.LDAPConfigurationApi.GetLDAPConfiguration(ctx, *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if isResourceEnabled2(*ldapConf) {
		return progressevent.GetFailedEventByCode("Authentication is already enabled for the selected project", cloudformation.HandlerErrorCodeAlreadyExists), nil
	}

	enabled := true

	currentModel.AuthenticationEnabled = &enabled

	ldapReq := currentModel.GetAtlasModel2()

	LDAPConfigResponse, resp, err := client.AtlasV2.LDAPConfigurationApi.SaveLDAPConfiguration(ctx, *currentModel.ProjectId, ldapReq).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.CompleteByResponse2(*LDAPConfigResponse)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create successfully",
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if err := validator.ValidateModel(ReadRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	ldapConf, errPe := get(client, *currentModel.ProjectId)
	if errPe != nil {
		return *errPe, nil
	}

	currentModel.CompleteByResponse(*ldapConf)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if err := validator.ValidateModel(UpdateRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Validate if resource exists
	_, errPe := get(client, *currentModel.ProjectId)
	if errPe != nil {
		return *errPe, nil
	}

	ldapReq := currentModel.GetAtlasModel()

	LDAPConfigResponse, res, err := client.LDAPConfigurations.Save(context.Background(), *currentModel.ProjectId, ldapReq)
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.CompleteByResponse(*LDAPConfigResponse)

	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if err := validator.ValidateModel(DeleteRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Validate if resource exists
	_, errPe := get(client, *currentModel.ProjectId)
	if errPe != nil {
		return *errPe, nil
	}

	ldapReq := currentModel.GetAtlasModel()
	ldapReq.LDAP.AuthorizationEnabled = aws.Bool(false)
	ldapReq.LDAP.AuthenticationEnabled = aws.Bool(false)

	_, res, err := client.LDAPConfigurations.Save(context.Background(), *currentModel.ProjectId, ldapReq)
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete success",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
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

func (m *Model) CompleteByResponse2(resp admin.UserSecurity) {
	m.AuthenticationEnabled = resp.Ldap.AuthenticationEnabled
	m.AuthorizationEnabled = resp.Ldap.AuthorizationEnabled

	mapping := make([]ApiAtlasNDSUserToDNMappingView, len(resp.Ldap.UserToDNMapping))

	for i := range resp.Ldap.UserToDNMapping {
		ndsMap := ApiAtlasNDSUserToDNMappingView{
			Match:        &resp.Ldap.UserToDNMapping[i].Match,
			Substitution: resp.Ldap.UserToDNMapping[i].Substitution,
			LdapQuery:    resp.Ldap.UserToDNMapping[i].LdapQuery,
		}
		mapping = append(mapping, ndsMap)
	}
	m.UserToDNMapping = mapping
}

func get(client *mongodbatlas.Client, groupID string) (*mongodbatlas.LDAPConfiguration, *handler.ProgressEvent) {
	ldapConf, res, err := client.LDAPConfigurations.Get(context.Background(), groupID)
	if err != nil {
		errPe := progressevent.GetFailedEventByResponse(err.Error(), res.Response)
		return nil, &errPe
	}

	if !isResourceEnabled(*ldapConf) {
		errPe := progressevent.GetFailedEventByCode("Authentication is disabled for the selected project", cloudformation.HandlerErrorCodeNotFound)
		return nil, &errPe
	}

	return ldapConf, nil
}

func (m *Model) GetAtlasModel() *mongodbatlas.LDAPConfiguration {
	DNMapping := getUserToDNMapping(m.UserToDNMapping)

	ldap := &mongodbatlas.LDAP{
		AuthenticationEnabled: aws.Bool(true),
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

func (m *Model) GetAtlasModel2() *admin.UserSecurity {
	DNMapping := getUserToDNMapping2(m.UserToDNMapping)

	ldap := &admin.LDAPSecuritySettings{
		AuthenticationEnabled: aws.Bool(true),
		Hostname:              m.Hostname,
		Port:                  m.Port,
		BindUsername:          m.BindUsername,
		UserToDNMapping:       DNMapping,
		BindPassword:          m.BindPassword,
	}

	ldapReq := &admin.UserSecurity{
		Ldap: ldap,
	}

	if m.AuthzQueryTemplate != nil {
		ldapReq.Ldap.AuthzQueryTemplate = m.AuthzQueryTemplate
	}

	if m.CaCertificate != nil {
		ldapReq.Ldap.CaCertificate = m.CaCertificate
	}

	if m.AuthorizationEnabled != nil {
		ldapReq.Ldap.AuthorizationEnabled = m.AuthorizationEnabled
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

func getUserToDNMapping2(ndsUserMapping []ApiAtlasNDSUserToDNMappingView) []admin.UserToDNMapping {
	mapping := make([]admin.UserToDNMapping, len(ndsUserMapping))

	for i := range ndsUserMapping {
		ndsMap := admin.UserToDNMapping{
			Match:        *ndsUserMapping[i].Match,
			Substitution: ndsUserMapping[i].Substitution,
			LdapQuery:    ndsUserMapping[i].LdapQuery,
		}
		mapping = append(mapping, ndsMap)
	}

	return mapping
}
func isResourceEnabled(ldapConf mongodbatlas.LDAPConfiguration) bool {
	if ldapConf.LDAP.AuthenticationEnabled != nil {
		return *ldapConf.LDAP.AuthenticationEnabled
	}
	return false
}

func isResourceEnabled2(ldapConf admin.UserSecurity) bool {
	return ldapConf.Ldap.AuthenticationEnabled != nil && *ldapConf.Ldap.AuthenticationEnabled
}
