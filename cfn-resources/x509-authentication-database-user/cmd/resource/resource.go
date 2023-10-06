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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.UserID}
var ReadRequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-x509-authentication-database-user")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(CreateRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["ProjectId"].(string)
		currentModel.ProjectId = &sid
		return validateProgress(client, currentModel)
	}

	if isEnabled(client, currentModel) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "resource already exists",
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}

	projectID := *currentModel.ProjectId
	userName := *currentModel.UserName
	expirationMonths := *currentModel.MonthsUntilExpiration

	// create new user certificate
	if expirationMonths > 0 {
		res, _, err := client.X509AuthDBUsers.CreateUserCertificate(context.Background(), projectID, userName, expirationMonths)
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		if res != nil {
			currentModel.CustomerX509 = &CustomerX509{
				Cas: aws.String(res.Certificate),
			}
		}
	} else { // save customer provided certificate
		customerX509Cas := *currentModel.CustomerX509.Cas
		_, _, err := client.X509AuthDBUsers.SaveConfiguration(context.Background(), projectID, &mongodbatlas.CustomerX509{Cas: customerX509Cas})
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
	}
	// track progress
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Created  Certificate  for DB User ",
		ResourceModel:   currentModel,
	}
	return event, nil
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

	certificate, resp, err := client.AtlasV2.LDAPConfigurationApi.GetLDAPConfiguration(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if certificate == nil || certificate.CustomerX509 == nil || !util.IsStringPresent(certificate.CustomerX509.Cas) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "config is not available",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	currentModel.CustomerX509 = &CustomerX509{
		Cas: certificate.CustomerX509.Cas,
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete: ",
		ResourceModel:   currentModel,
	}, nil
}

func ReadUserX509Certificate(client *mongodbatlas.Client, currentModel *Model) (*Model, error) {
	projectID := *currentModel.ProjectId

	certificate, _, err := client.X509AuthDBUsers.GetCurrentX509Conf(context.Background(), projectID)
	if err != nil {
		return nil, fmt.Errorf("error reading MongoDB X509 Authentication for DB Users(%s) in the project(%s): %s", *currentModel.UserName, projectID, err)
	} else if certificate != nil {
		currentModel.CustomerX509 = &CustomerX509{
			Cas: &certificate.Cas,
		}
	}
	return currentModel, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(CreateRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if !isEnabled(client, currentModel) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "config is not available",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	projectID := *currentModel.ProjectId
	_, err := client.X509AuthDBUsers.DisableCustomerX509(context.Background(), projectID)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Unable to Delete",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInternalFailure,
		}, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func validateProgress(client *mongodbatlas.Client, currentModel *Model) (handler.ProgressEvent, error) {
	projectID := *currentModel.ProjectId
	isReady, state, err := certificateIsReady(client, projectID)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	if !isReady {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = handler.InProgress
		p.CallbackDelaySeconds = 10
		p.Message = "Pending"
		p.CallbackContext = map[string]interface{}{
			"status":    state,
			"ProjectId": *currentModel.ProjectId,
		}
		return p, nil
	}

	p := handler.NewProgressEvent()
	p.ResourceModel = currentModel
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	return p, nil
}

func isEnabled(client *mongodbatlas.Client, currentModel *Model) bool {
	projectID := *currentModel.ProjectId

	certificate, _, err := client.X509AuthDBUsers.GetCurrentX509Conf(context.Background(), projectID)
	if err != nil {
		return false
	} else if certificate != nil && certificate.Cas != "" {
		return true
	}

	return false
}

func certificateIsReady(client *mongodbatlas.Client, projectID string) (isExist bool, groupID string, errMsg error) {
	certificate, resp, err := client.X509AuthDBUsers.GetCurrentX509Conf(context.Background(), projectID)
	if err != nil {
		if certificate == nil && resp == nil {
			return false, "", err
		}
		if resp != nil && resp.StatusCode == 404 {
			return true, "deleted", nil
		}
		return false, "", err
	}
	return resp.StatusCode == constants.Success, "completed", nil
}
