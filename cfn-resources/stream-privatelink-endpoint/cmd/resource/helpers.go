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
	"fmt"
	"net/http"
	"strings"

	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

const (
	errorAlreadyExistsAPI         = "STREAM_PRIVATE_LINK_ALREADY_EXISTS"
	errorAlreadyExistsGeneric     = "already exists"
	errorMessageFailedStatus      = "Private endpoint is in a failed status"
	errorMessageUpdateUnsupported = "Updating the private endpoint for streams is not supported. To modify your infrastructure, please delete the existing resource and create a new one with the necessary updates"
)

func handleFailedState(connection *admin.StreamsPrivateLinkConnection) handler.ProgressEvent {
	errorMsg := errorMessageFailedStatus
	if connection != nil {
		if errMsg := connection.GetErrorMessage(); errMsg != "" {
			errorMsg = fmt.Sprintf("%s: %s", errorMsg, errMsg)
		}
	}
	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          errorMsg,
		HandlerErrorCode: string(types.HandlerErrorCodeInvalidRequest),
	}
}

func handleError(response *http.Response, method constants.CfnFunctions, err error) handler.ProgressEvent {
	errMsg := fmt.Sprintf("%s error: %s", method, err.Error())

	if util.StatusNotFound(response) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          constants.ResourceNotFound,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
		}
	}

	if util.StatusBadRequest(response) {
		errStr := err.Error()
		if strings.Contains(errStr, errorAlreadyExistsAPI) || strings.Contains(errStr, errorAlreadyExistsGeneric) {
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          errMsg,
				HandlerErrorCode: string(types.HandlerErrorCodeAlreadyExists),
			}
		}
	}

	return progress_events.GetFailedEventByResponse(errMsg, response)
}
