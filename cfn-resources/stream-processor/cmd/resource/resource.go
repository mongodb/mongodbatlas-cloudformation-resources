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
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

const (
	InitiatingState = "INIT"
	CreatingState   = "CREATING"
	CreatedState    = "CREATED"
	StartedState    = "STARTED"
	StoppedState    = "STOPPED"
	DroppedState    = "DROPPED"
	FailedState     = "FAILED"
)

const (
	defaultCallbackDelaySeconds = 3
	DefaultCreateTimeout        = 20 * time.Minute
)

var (
	createRequiredFields           = []string{constants.ProjectID, constants.WorkspaceName, constants.ProcessorName, constants.Pipeline}
	readUpdateDeleteRequiredFields = []string{constants.ProjectID, constants.WorkspaceName, constants.ProcessorName}
	listRequiredFields             = []string{constants.ProjectID, constants.WorkspaceName}
)

func Create(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, createRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return HandleCreate(&req, client, model), nil
}

func Read(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, readUpdateDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return HandleRead(&req, client, model), nil
}

func Update(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, readUpdateDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return HandleUpdate(&req, client, prevModel, model), nil
}

func Delete(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, readUpdateDeleteRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return HandleDelete(&req, client, model), nil
}

func List(req handler.Request, prevModel *Model, model *Model) (handler.ProgressEvent, error) {
	client, setupErr := setupRequest(req, model, listRequiredFields)
	if setupErr != nil {
		return *setupErr, nil
	}
	return HandleList(&req, client, model), nil
}

func setupRequest(req handler.Request, model *Model, requiredFields []string) (*util.MongoDBClient, *handler.ProgressEvent) {
	util.SetupLogger("mongodb-atlas-stream-processor")
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
