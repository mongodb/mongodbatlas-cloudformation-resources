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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"

	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"
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
	ldapConf, resp, err := client.Atlas20231115002.LDAPConfigurationApi.GetLDAPConfiguration(ctx, *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if isResourceEnabled(ldapConf) {
		return progressevent.GetFailedEventByCode("Authentication is already enabled for the selected project", string(types.HandlerErrorCodeAlreadyExists)), nil
	}

	currentModel.AuthenticationEnabled = aws.Bool(true)

	ldapReq := currentModel.GetAtlasModel()

	LDAPConfigResponse, resp, err := client.Atlas20231115002.LDAPConfigurationApi.SaveLDAPConfiguration(ctx, *currentModel.ProjectId, ldapReq).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.CompleteByResponse(*LDAPConfigResponse)

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

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	ldapConf, pe := get(client, *currentModel.ProjectId)
	if pe != nil {
		return *pe, nil
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

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	if _, pe := get(client, *currentModel.ProjectId); pe != nil {
		return *pe, nil
	}

	ldapReq := currentModel.GetAtlasModel()

	ctx := context.Background()
	LDAPConfigResponse, resp, err := client.Atlas20231115002.LDAPConfigurationApi.SaveLDAPConfiguration(ctx, *currentModel.ProjectId, ldapReq).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.CompleteByResponse(*LDAPConfigResponse)

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

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	if _, pe := get(client, *currentModel.ProjectId); pe != nil {
		return *pe, nil
	}

	ldapReq := currentModel.GetAtlasModel()
	ldapReq.Ldap.AuthorizationEnabled = aws.Bool(false)
	ldapReq.Ldap.AuthenticationEnabled = aws.Bool(false)

	ctx := context.Background()
	_, resp, err := client.Atlas20231115002.LDAPConfigurationApi.SaveLDAPConfiguration(ctx, *currentModel.ProjectId, ldapReq).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete success",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func (m *Model) CompleteByResponse(resp admin20231115002.UserSecurity) {
	m.AuthenticationEnabled = resp.Ldap.AuthenticationEnabled
	m.AuthorizationEnabled = resp.Ldap.AuthorizationEnabled

	mappings := make([]ApiAtlasNDSUserToDNMappingView, len(resp.Ldap.UserToDNMapping))

	for i := range resp.Ldap.UserToDNMapping {
		mappings[i] = ApiAtlasNDSUserToDNMappingView{
			Match:        &resp.Ldap.UserToDNMapping[i].Match,
			Substitution: resp.Ldap.UserToDNMapping[i].Substitution,
			LdapQuery:    resp.Ldap.UserToDNMapping[i].LdapQuery,
		}
	}
	m.UserToDNMapping = mappings
}

func get(client *util.MongoDBClient, groupID string) (*admin20231115002.UserSecurity, *handler.ProgressEvent) {
	ctx := context.Background()
	ldapConf, resp, err := client.Atlas20231115002.LDAPConfigurationApi.GetLDAPConfiguration(ctx, groupID).Execute()
	if err != nil {
		errPe := progressevent.GetFailedEventByResponse(err.Error(), resp)
		return nil, &errPe
	}

	if !isResourceEnabled(ldapConf) {
		errPe := progressevent.GetFailedEventByCode("LDAP Authentication is disabled for the selected project", string(types.HandlerErrorCodeNotFound))
		return nil, &errPe
	}

	return ldapConf, nil
}

func (m *Model) GetAtlasModel() *admin20231115002.UserSecurity {
	DNMapping := getUserToDNMapping(m.UserToDNMapping)

	ldap := &admin20231115002.LDAPSecuritySettings{
		AuthenticationEnabled: aws.Bool(true),
		Hostname:              m.Hostname,
		Port:                  m.Port,
		BindUsername:          m.BindUsername,
		UserToDNMapping:       DNMapping,
		BindPassword:          m.BindPassword,
	}

	ldapReq := &admin20231115002.UserSecurity{
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

func getUserToDNMapping(ndsUserMapping []ApiAtlasNDSUserToDNMappingView) []admin20231115002.UserToDNMapping {
	mappings := make([]admin20231115002.UserToDNMapping, len(ndsUserMapping))
	for i := range ndsUserMapping {
		mappings[i] = admin20231115002.UserToDNMapping{
			Match:        *ndsUserMapping[i].Match,
			Substitution: ndsUserMapping[i].Substitution,
			LdapQuery:    ndsUserMapping[i].LdapQuery,
		}
	}
	return mappings
}

func isResourceEnabled(ldapConf *admin20231115002.UserSecurity) bool {
	return ldapConf.Ldap.AuthenticationEnabled != nil && *ldapConf.Ldap.AuthenticationEnabled
}
