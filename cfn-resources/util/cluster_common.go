// Copyright 2025 MongoDB Inc
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

package util

import (
	"net/http"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

// HandleClusterError provides common error handling for cluster operations.
func HandleClusterError(err error, resp *http.Response) *handler.ProgressEvent {
	if err == nil {
		return nil
	}
	pe := progressevent.GetFailedEventByResponse(err.Error(), resp)
	if resp != nil && resp.StatusCode == http.StatusBadRequest && strings.Contains(err.Error(), constants.Duplicate) {
		pe.HandlerErrorCode = cloudformation.HandlerErrorCodeAlreadyExists
	}
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		pe.HandlerErrorCode = cloudformation.HandlerErrorCodeNotFound
	}
	if strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "being deleted") {
		pe.HandlerErrorCode = cloudformation.HandlerErrorCodeNotFound
	}
	return &pe
}
