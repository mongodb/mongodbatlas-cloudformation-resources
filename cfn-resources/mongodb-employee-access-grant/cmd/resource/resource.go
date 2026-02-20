// Copyright 2025 MongoDB Inc
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
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"
)

var (
	createUpdateRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.GrantType, constants.ExpirationTime}
	readDeleteRequiredFields   = []string{constants.ProjectID, constants.ClusterName}
	listRequiredFields         = []string{constants.ProjectID}
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, createUpdateRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return HandleCreate(&req, client, model), nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, readDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return HandleRead(&req, client, model), nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, createUpdateRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return HandleUpdate(&req, client, model), nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, readDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return HandleDelete(&req, client, model), nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, listRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return HandleList(&req, client, model), nil
}

func setupRequest(req handler.Request, model *Model, requiredFields []string) (*util.MongoDBClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-mongodbemployeeaccessgrant")
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

func updateModel(model *Model, apiResp *admin.EmployeeAccessGrant) {
	if apiResp == nil {
		return
	}
	model.GrantType = &apiResp.GrantType
	expirationTimeStr := apiResp.ExpirationTime.Format(time.RFC3339)
	model.ExpirationTime = &expirationTimeStr
}
