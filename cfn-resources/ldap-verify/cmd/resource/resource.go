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
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20240805004/admin"
)

const (
	BindUsername = "BindUsername"
	BindPassword = "BindPassword"
	RequestID    = "RequestId"
)

var CreateRequiredFields = []string{constants.ProjectID, BindUsername, BindPassword, constants.HostName, constants.Port}
var ReadRequiredFields = []string{constants.ProjectID, RequestID}
var DeleteRequiredFields = []string{constants.ProjectID, RequestID}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-ldap-verify")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if err := validateModel(CreateRequiredFields, currentModel); err != nil {
		return *err, nil
	}
	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}
	if req.CallbackContext != nil {
		return validateProgress(client, currentModel, req), nil
	}

	params := currentModel.GetAtlasParams()
	ldapResponse, resp, err := client.Atlas20231115002.LDAPConfigurationApi.VerifyLDAPConfiguration(context.Background(), *currentModel.ProjectId, params).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.CompleteByResponse(ldapResponse)

	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              "Create in progress",
		ResourceModel:        currentModel,
		CallbackDelaySeconds: 10,
		CallbackContext: map[string]interface{}{
			"RequestId": currentModel.RequestId,
		},
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if err := validateModel(ReadRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	ldapResponse, resp, err := client.Atlas20231115002.LDAPConfigurationApi.GetLDAPConfigurationStatus(context.Background(), *currentModel.ProjectId, *currentModel.RequestId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	currentModel.CompleteByResponse(ldapResponse)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if err := validateModel(DeleteRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	_, resp, err := client.Atlas20231115002.LDAPConfigurationApi.GetLDAPConfigurationStatus(context.Background(), *currentModel.ProjectId, *currentModel.RequestId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(fmt.Sprintf("Error deleting resource : %s", err.Error()), resp), nil
	}

	params := &admin.LDAPVerifyConnectivityJobRequestParams{
		Hostname:     "-",
		Port:         1111,
		BindPassword: "-",
		BindUsername: "-",
	}
	if _, resp, err := client.Atlas20231115002.LDAPConfigurationApi.VerifyLDAPConfiguration(context.Background(), *currentModel.ProjectId, params).Execute(); err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete completed successfully",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func (m *Model) GetAtlasParams() *admin.LDAPVerifyConnectivityJobRequestParams {
	ldap := admin.LDAPVerifyConnectivityJobRequestParams{
		Hostname:     *m.HostName,
		Port:         *m.Port,
		BindPassword: *m.BindPassword,
		BindUsername: *m.BindUsername,
	}

	if m.AuthzQueryTemplate != nil {
		ldap.AuthzQueryTemplate = m.AuthzQueryTemplate
	}

	if m.CaCertificate != nil {
		ldap.CaCertificate = m.CaCertificate
	}

	return &ldap
}

func (m *Model) CompleteByResponse(resp *admin.LDAPVerifyConnectivityJobRequest) {
	m.RequestId = resp.RequestId

	mapping := make([]Validation, len(resp.Validations))

	for i := range resp.Validations {
		validation := Validation{
			Status:         resp.Validations[i].Status,
			ValidationType: resp.Validations[i].ValidationType,
		}
		mapping[i] = validation
	}

	m.Validations = mapping
	m.Status = resp.Status
}

func validateProgress(client *util.MongoDBClient, model *Model, req handler.Request) handler.ProgressEvent {
	requestID := req.CallbackContext["RequestId"].(string)

	ldapResponse, resp, err := client.Atlas20231115002.LDAPConfigurationApi.GetLDAPConfigurationStatus(context.Background(), *model.ProjectId, requestID).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp)
	}

	switch *ldapResponse.Status {
	case "PENDING":
		return handler.ProgressEvent{
			OperationStatus: handler.InProgress,
			Message:         "Create in progress",
			ResourceModel:   model,
			CallbackContext: map[string]interface{}{
				"RequestId": requestID,
			},
		}
	case "SUCCESS":
		model.CompleteByResponse(ldapResponse)
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Create successfully",
			ResourceModel:   model,
		}
	default:
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         getFailedMessage(ldapResponse),
		}
	}
}

func getFailedMessage(configuration *admin.LDAPVerifyConnectivityJobRequest) string {
	for _, i := range configuration.Validations {
		if *i.Status == "FAIL" {
			return fmt.Sprintf("Validation fail: %s", *i.ValidationType)
		}
	}
	return ""
}
