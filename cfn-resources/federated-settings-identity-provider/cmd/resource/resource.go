// Copyright 2026 MongoDB Inc
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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
)

var (
	createRequiredFields = []string{constants.FederationSettingsID, constants.Name, constants.IssuerURI}
	readRequiredFields   = []string{constants.FederationSettingsID, constants.IdpID}
	updateRequiredFields = []string{constants.FederationSettingsID, constants.IdpID}
	deleteRequiredFields = []string{constants.FederationSettingsID, constants.IdpID}
	listRequiredFields   = []string{constants.FederationSettingsID}
)

func setupRequest(req handler.Request, currentModel *Model, requiredFields []string) (*util.MongoDBClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-federated-settings-identity-provider")
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	if errEvent := validator.ValidateModel(requiredFields, currentModel); errEvent != nil {
		return nil, errEvent
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return nil, pe
	}
	return client, nil
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, peErr := setupRequest(req, currentModel, createRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}
	return HandleCreate(client, currentModel), nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, peErr := setupRequest(req, currentModel, readRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}
	return HandleRead(client, currentModel), nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, peErr := setupRequest(req, currentModel, updateRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}
	return HandleUpdate(client, prevModel, currentModel), nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, peErr := setupRequest(req, currentModel, deleteRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}
	return HandleDelete(client, currentModel), nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, peErr := setupRequest(req, currentModel, listRequiredFields)
	if peErr != nil {
		return *peErr, nil
	}
	return HandleList(client, currentModel), nil
}
