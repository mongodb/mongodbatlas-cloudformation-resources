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

	"go.mongodb.org/atlas-sdk/v20250312014/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
)

const (
	stateDone            = "DONE"
	stateFailed          = "FAILED"
	stateDeleteRequested = "DELETE_REQUESTED"
	stateDeleting        = "DELETING"
	stateDeleted         = "DELETED"

	callbackDelaySeconds = 3
	callbackKey          = "callbackStreamPrivatelinkEndpoint"
)

var callbackContext = map[string]any{callbackKey: true}

func isCallback(req *handler.Request) bool {
	_, found := req.CallbackContext[callbackKey]
	return found
}

func validateProgress(client *util.MongoDBClient, model *Model, isDelete bool) handler.ProgressEvent {
	ctx := context.Background()
	projectID := util.SafeString(model.ProjectId)
	connectionID := util.SafeString(model.Id)

	streamsPrivateLinkConnection, apiResp, err := client.AtlasSDK.StreamsApi.GetPrivateLinkConnection(ctx, projectID, connectionID).Execute()
	notFound := util.StatusNotFound(apiResp)

	if err != nil && !notFound {
		return handleError(apiResp, constants.READ, err)
	}

	state := stateDeleted
	if streamsPrivateLinkConnection != nil {
		state = streamsPrivateLinkConnection.GetState()
	}

	targetState := stateDone
	if isDelete {
		targetState = stateDeleted
	}

	if state == stateFailed {
		return handleFailedState(streamsPrivateLinkConnection)
	}

	if state != targetState {
		return inProgressEvent(model, streamsPrivateLinkConnection)
	}

	if isDelete {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         constants.Complete,
		}
	}

	UpdateModel(model, streamsPrivateLinkConnection)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModel:   model,
	}
}

func inProgressEvent(model *Model, apiResp *admin.StreamsPrivateLinkConnection) handler.ProgressEvent {
	UpdateModel(model, apiResp)
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              constants.Pending,
		ResourceModel:        model,
		CallbackDelaySeconds: callbackDelaySeconds,
		CallbackContext:      callbackContext,
	}
}
