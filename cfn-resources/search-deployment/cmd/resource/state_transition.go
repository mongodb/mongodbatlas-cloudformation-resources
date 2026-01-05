// Copyright 2024 MongoDB Inc
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
	"net/http"
	"strings"

	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

func ValidateProgress(connV2 admin20250312010.APIClient, currentModel *Model, isDelete bool) handler.ProgressEvent {
	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)

	apiResp, resp, err := connV2.AtlasSearchApi.GetClusterSearchDeployment(context.Background(), projectID, clusterName).Execute()

	notFound := resp != nil && resp.StatusCode == http.StatusNotFound
	doesNotExist := resp != nil && resp.StatusCode == http.StatusBadRequest && err != nil &&
		strings.Contains(err.Error(), "ATLAS_SEARCH_DEPLOYMENT_DOES_NOT_EXIST")

	if err != nil {
		if isDelete && (notFound || doesNotExist) {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         constants.Complete,
			}
		}
		return progressevent.GetFailedEventByResponse(err.Error(), resp)
	}

	state := constants.DeletedState
	if apiResp != nil && apiResp.StateName != nil {
		state = *apiResp.StateName
	}
	targetState := constants.IdleState
	if isDelete {
		targetState = constants.DeletedState
	}

	if state != targetState {
		return inProgressEvent(constants.Pending, currentModel, apiResp)
	}

	if isDelete {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         constants.Complete,
		}
	}

	newModel := NewCFNSearchDeployment(currentModel, apiResp)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
		ResourceModel:   &newModel,
	}
}
