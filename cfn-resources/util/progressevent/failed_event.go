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
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

func getHandlerErrorCode(response *http.Response) string {
	if response == nil {
		return string(types.HandlerErrorCodeInternalFailure)
	}
	switch response.StatusCode {
	case http.StatusBadRequest:
		return string(types.HandlerErrorCodeInvalidRequest)
	case http.StatusNotFound:
		return string(types.HandlerErrorCodeNotFound)
	case http.StatusConflict:
		return string(types.HandlerErrorCodeAlreadyExists)
	case http.StatusInternalServerError:
		return string(types.HandlerErrorCodeServiceInternalError)
	case http.StatusPaymentRequired, http.StatusUnauthorized:
		return string(types.HandlerErrorCodeAccessDenied)
	default:
		return string(types.HandlerErrorCodeInternalFailure)
	}
}

func GetFailedEventByResponse(message string, response *http.Response) handler.ProgressEvent {
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
