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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"go.mongodb.org/atlas-sdk/v20250312014/admin"
)

func HandleStateTransition(connV2 admin.APIClient, currentModel *Model, targetState string) handler.ProgressEvent {
	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	apiResp, resp, err := connV2.AtlasSearchApi.GetClusterSearchDeployment(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if targetState == constants.DeletedState && resp != nil && resp.StatusCode == http.StatusBadRequest && strings.Contains(err.Error(), SearchDeploymentDoesNotExistsError) {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				ResourceModel:   nil,
				Message:         constants.Complete,
			}
		}
		return progressevent.GetFailedEventByResponse(err.Error(), resp)
	}

	newModel := NewCFNSearchDeployment(currentModel, apiResp)

	// For delete operations, check if Specs is empty - this indicates deletion is complete
	// The Atlas API returns 200 with only basic fields (no Specs) when deployment is deleted
	if targetState == constants.DeletedState && len(newModel.Specs) == 0 {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			ResourceModel:   nil,
			Message:         constants.Complete,
		}
	}

	if util.SafeString(newModel.StateName) == targetState {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			ResourceModel:   newModel,
			Message:         constants.Complete,
		}
	}

	return inProgressEvent(constants.Pending, &newModel)
}
