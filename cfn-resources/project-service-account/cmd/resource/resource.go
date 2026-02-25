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
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

var (
	CreateRequiredFields = []string{constants.ProjectID, constants.Name, constants.Description, constants.Roles, constants.SecretExpiresAfterHours}
	ReadRequiredFields   = []string{constants.ProjectID, constants.ClientID}
	UpdateRequiredFields = []string{constants.ProjectID, constants.ClientID}
	DeleteRequiredFields = []string{constants.ProjectID, constants.ClientID}
	ListRequiredFields   = []string{constants.ProjectID}
)

func setupRequest(req handler.Request, model *Model, requiredFields []string) (*util.MongoDBClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-project-service-account")
	if modelValidation := validator.ValidateModel(requiredFields, model); modelValidation != nil {
		return nil, modelValidation
	}
	util.SetDefaultProfileIfNotDefined(&model.Profile)
	client, peErr := util.NewAtlasClient(&req, model.Profile)
	if peErr != nil {
		return nil, peErr
	}
	return client, nil
}

func Create(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, CreateRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return handleCreate(client, model), nil
}

func Read(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, ReadRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return handleRead(client, model), nil
}

func Update(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, UpdateRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return handleUpdate(client, model), nil
}

func Delete(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, DeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return handleDelete(client, model), nil
}

func List(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, ListRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return handleList(client, model), nil
}
