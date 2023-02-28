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
	userprofile "github.com/mongodb/mongodbatlas-cloudformation-resources/profile"

	"github.com/openlyinc/pointy"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	BindUsername = "BindUsername"
	BindPassword = "BindPassword"
	RequestID    = "RequestId"
)

var CreateRequiredFields = []string{constants.GroupID, BindUsername, BindPassword, constants.HostName, constants.Port}
var ReadRequiredFields = []string{constants.GroupID, RequestID}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.GroupID, RequestID}
var ListRequiredFields []string

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-LDAPVerify")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(userprofile.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	if req.CallbackContext != nil {
		return validateProgress(client, currentModel, req), nil
	}

	ldapReq := currentModel.GetAtlasModel()

	LDAPConfigResponse, res, err := client.LDAPConfigurations.Verify(context.Background(), *currentModel.GroupId, ldapReq)
	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.CompleteByResponse(*LDAPConfigResponse)

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

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client

	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(userprofile.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	var res *mongodbatlas.Response

	LDAPConfigResponse, res, err := client.LDAPConfigurations.GetStatus(context.Background(), *currentModel.GroupId, *currentModel.RequestId)
	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.CompleteByResponse(*LDAPConfigResponse)

	// Response
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

	// Validation
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil {
		currentModel.Profile = pointy.String(userprofile.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	_, res, err := client.LDAPConfigurations.GetStatus(context.Background(), *currentModel.GroupId, *currentModel.RequestId)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error deleting resource : %s", err.Error()),
			res.Response), nil
	}

	ldapReq := &mongodbatlas.LDAP{
		Hostname:     pointy.String("-"),
		Port:         pointy.Int(1111),
		BindPassword: pointy.String("-"),
		BindUsername: pointy.String("-"),
	}

	_, res, err = client.LDAPConfigurations.Verify(context.Background(), *currentModel.GroupId, ldapReq)
	if err != nil {
		_, _ = log.Debugf("Delete - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete completed successfully",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func (m *Model) GetAtlasModel() *mongodbatlas.LDAP {
	ldap := &mongodbatlas.LDAP{
		Hostname:     m.HostName,
		Port:         m.Port,
		BindPassword: m.BindPassword,
		BindUsername: m.BindUsername,
	}

	if m.AuthzQueryTemplate != nil {
		ldap.AuthzQueryTemplate = m.AuthzQueryTemplate
	}

	if m.CaCertificate != nil {
		ldap.CaCertificate = m.CaCertificate
	}

	return ldap
}

func (m *Model) CompleteByResponse(resp mongodbatlas.LDAPConfiguration) {
	m.RequestId = &resp.RequestID

	mapping := make([]Validation, len(resp.Validations))

	for i := range resp.Validations {
		validation := Validation{
			Status:         &resp.Validations[i].Status,
			ValidationType: &resp.Validations[i].ValidationType,
		}
		mapping[i] = validation
	}

	m.Validations = mapping
	m.Status = &resp.Status
}

func validateProgress(client *mongodbatlas.Client, model *Model, req handler.Request) handler.ProgressEvent {
	requestID := req.CallbackContext["RequestId"].(string)

	LDAPConfigResponse, res, err := client.LDAPConfigurations.GetStatus(context.Background(), *model.GroupId, requestID)
	if err != nil {
		_, _ = log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response)
	}

	switch LDAPConfigResponse.Status {
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
		model.CompleteByResponse(*LDAPConfigResponse)
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Create successfully",
			ResourceModel:   model,
		}
	default:
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         getFailedMessage(*LDAPConfigResponse),
		}
	}
}

func getFailedMessage(configuration mongodbatlas.LDAPConfiguration) string {
	for _, i := range configuration.Validations {
		if i.Status == "FAIL" {
			return fmt.Sprintf("Faild in validation %s", i.ValidationType)
		}
	}

	return ""
}
