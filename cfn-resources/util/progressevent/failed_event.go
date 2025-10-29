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
	cloudformationtypes "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

func getHandlerErrorCode(response *http.Response) string {
	switch response.StatusCode {
	case http.StatusBadRequest:
		return string(cloudformationtypes.HandlerErrorCodeInvalidRequest)
	case http.StatusNotFound:
		return string(cloudformationtypes.HandlerErrorCodeNotFound)
	case http.StatusInternalServerError:
		return string(cloudformationtypes.HandlerErrorCodeServiceInternalError)
	case http.StatusPaymentRequired, http.StatusUnauthorized:
		return string(cloudformationtypes.HandlerErrorCodeAccessDenied)
	default:
		return string(cloudformationtypes.HandlerErrorCodeInternalFailure)
	}
}

func GetFailedEventByResponse(message string, response *http.Response) handler.ProgressEvent {
	if response == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          message,
			HandlerErrorCode: string(cloudformationtypes.HandlerErrorCodeHandlerInternalFailure)}
	}

	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          message,
			HandlerErrorCode: string(cloudformationtypes.HandlerErrorCodeAlreadyExists)}
	}

	if response.StatusCode == http.StatusUnauthorized {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Not found",
			HandlerErrorCode: string(cloudformationtypes.HandlerErrorCodeNotFound)}
	}

	if response.StatusCode == http.StatusBadRequest {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          message,
			HandlerErrorCode: string(cloudformationtypes.HandlerErrorCodeNotFound)}
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
