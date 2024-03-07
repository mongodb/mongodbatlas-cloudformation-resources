// Copyright 2023 MongoDB Inc
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

package progressevent

import (
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

func getHandlerErrorCode(response *http.Response) string {
	switch response.StatusCode {
	case http.StatusBadRequest:
		return cloudformation.HandlerErrorCodeInvalidRequest
	case http.StatusNotFound:
		return cloudformation.HandlerErrorCodeNotFound
	case http.StatusInternalServerError:
		return cloudformation.HandlerErrorCodeServiceInternalError
	case http.StatusPaymentRequired, http.StatusUnauthorized:
		return cloudformation.HandlerErrorCodeAccessDenied
	default:
		return cloudformation.HandlerErrorCodeInternalFailure
	}
}

func GetFailedEventByResponse(message string, response *http.Response) handler.ProgressEvent {
	if response == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          message,
			HandlerErrorCode: cloudformation.HandlerErrorCodeHandlerInternalFailure}
	}

	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          message,
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}
	}

	if response.StatusCode == http.StatusUnauthorized {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Not found",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}
	}

	if response.StatusCode == http.StatusBadRequest {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          message,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}
	}

	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          message,
		HandlerErrorCode: getHandlerErrorCode(response)}
}

func GetFailedEventByCode(message, handlerErrorCode string) handler.ProgressEvent {
	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          message,
		HandlerErrorCode: handlerErrorCode}
}
