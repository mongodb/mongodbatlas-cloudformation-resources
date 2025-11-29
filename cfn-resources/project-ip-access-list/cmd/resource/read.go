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

package resource

import (
	"context"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if !util.IsStringPresent(currentModel.Profile) {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	result, resp, err := client.Atlas20231115002.ProjectIPAccessListApi.ListProjectIpAccessLists(context.Background(), *currentModel.ProjectId).Execute()
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp), nil
	}

	if *result.TotalCount == 0 {
		return handler.ProgressEvent{
			Message:          "The entry to read is not in the access list",
			OperationStatus:  handler.Failed,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound)}, nil
	}

	currentModel.TotalCount = result.TotalCount
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}
