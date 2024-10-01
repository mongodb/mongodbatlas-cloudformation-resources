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
	"go.mongodb.org/atlas-sdk/v20240805004/admin"
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

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	certificate, resp, err := client.Atlas20231115002.LDAPConfigurationApi.GetLDAPConfiguration(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if isEnabled(certificate) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "resource already exists",
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}

	if expirationMonths := aws.IntValue(currentModel.MonthsUntilExpiration); expirationMonths > 0 {
		cert := admin.NewUserCert()
		cert.MonthsUntilExpiration = &expirationMonths
		res, _, err := client.Atlas20231115002.X509AuthenticationApi.CreateDatabaseUserCertificate(context.Background(), *currentModel.ProjectId, *currentModel.UserName, cert).Execute()
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		if res != "" {
			currentModel.CustomerX509 = &CustomerX509{
				Cas: &res,
			}
		}
	} else {
		customerX509 := &admin.DBUserTLSX509Settings{Cas: currentModel.CustomerX509.Cas}
		_, _, err := client.Atlas20231115002.LDAPConfigurationApi.SaveLDAPConfiguration(context.Background(), *currentModel.ProjectId, &admin.UserSecurity{CustomerX509: customerX509}).Execute()
		if err != nil {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
	}

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

	certificate, resp, err := client.Atlas20231115002.LDAPConfigurationApi.GetLDAPConfiguration(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if !isEnabled(certificate) {
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

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(CreateRequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	certificate, resp, err := client.Atlas20231115002.LDAPConfigurationApi.GetLDAPConfiguration(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if !isEnabled(certificate) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "config is not available",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	_, _, err = client.Atlas20231115002.X509AuthenticationApi.DisableCustomerManagedX509(context.Background(), *currentModel.ProjectId).Execute()
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

func isEnabled(certificate *admin.UserSecurity) bool {
	return certificate != nil && certificate.CustomerX509 != nil && util.IsStringPresent(certificate.CustomerX509.Cas)
}
