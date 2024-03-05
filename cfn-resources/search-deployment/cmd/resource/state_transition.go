package resource

import (
	"context"
	"net/http"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func handleStateTransition(connV2 admin.APIClient, currentModel *Model, targetState string) handler.ProgressEvent {
	projectID := util.SafeString(currentModel.ProjectId)
	clusterName := util.SafeString(currentModel.ClusterName)
	apiResp, resp, err := connV2.AtlasSearchApi.GetAtlasSearchDeployment(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		if targetState == constants.DeletedState && resp.StatusCode == http.StatusBadRequest && strings.Contains(err.Error(), searchDeploymentDoesNotExistsError) {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				ResourceModel:   nil,
				Message:         constants.Complete,
			}
		}
		return progressevent.GetFailedEventByResponse(err.Error(), resp)
	}

	newModel := newCFNSearchDeployment(currentModel, apiResp)
	if util.SafeString(newModel.StateName) == targetState {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			ResourceModel:   newModel,
			Message:         constants.Complete,
		}
	}

	return inProgressEvent(constants.Pending, &newModel)
}
