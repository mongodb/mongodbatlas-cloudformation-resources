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
	"context"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

func HandleCreate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	s3LogIntegrationReq := NewLogIntegrationCreateRequest(model)
	logIntegrationResp, resp, err := client.AtlasSDK.PushBasedLogExportApi.CreateGroupLogIntegration(context.Background(), *model.ProjectId, s3LogIntegrationReq).Execute()
	if err != nil {
		return handleError(resp, err, "Error creating log integration")
	}

	UpdateLogIntegrationModel(model, logIntegrationResp)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create complete",
		ResourceModel:   model,
	}
}

func HandleRead(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	logIntegrationResp, resp, err := client.AtlasSDK.PushBasedLogExportApi.GetGroupLogIntegration(context.Background(), *model.ProjectId, *model.IntegrationId).Execute()
	if err != nil {
		return handleError(resp, err, "Error reading log integration")
	}

	UpdateLogIntegrationModel(model, logIntegrationResp)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   model,
	}
}

func HandleUpdate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	logIntegrationReq := NewLogIntegrationUpdateRequest(model)
	logIntegrationResp, resp, err := client.AtlasSDK.PushBasedLogExportApi.UpdateGroupLogIntegration(context.Background(), *model.ProjectId, *model.IntegrationId, logIntegrationReq).Execute()
	if err != nil {
		return handleError(resp, err, "Error updating log integration")
	}

	UpdateLogIntegrationModel(model, logIntegrationResp)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update complete",
		ResourceModel:   model,
	}
}

func HandleDelete(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	resp, err := client.AtlasSDK.PushBasedLogExportApi.DeleteGroupLogIntegration(context.Background(), *model.ProjectId, *model.IntegrationId).Execute()
	if err != nil {
		return handleError(resp, err, "Error deleting log integration")
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete complete",
	}
}

func HandleList(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	paginatedResp, resp, err := client.AtlasSDK.PushBasedLogExportApi.ListGroupLogIntegrations(context.Background(), *model.ProjectId).Execute()
	if err != nil {
		return handleError(resp, err, "Error listing log integrations")
	}

	var allModels []*Model
	results := paginatedResp.GetResults()
	for i := range results {
		modelItem := &Model{
			ProjectId: model.ProjectId,
			Profile:   model.Profile,
		}
		UpdateLogIntegrationModel(modelItem, &results[i])
		allModels = append(allModels, modelItem)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModel:   allModels,
	}
}

func handleError(resp *http.Response, err error, message string) handler.ProgressEvent {
	errMsg := fmt.Sprintf("%s: %v", message, err)
	return progressevent.GetFailedEventByResponse(errMsg, resp)
}
