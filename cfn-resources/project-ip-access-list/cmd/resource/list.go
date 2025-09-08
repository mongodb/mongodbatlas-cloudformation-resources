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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if !util.IsStringPresent(currentModel.Profile) {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	pageNum := 0
	itemsPerPage := 500
	includeCount := true

	if currentModel.ListOptions != nil {
		if currentModel.ListOptions.PageNum != nil {
			pageNum = *currentModel.ListOptions.PageNum
		}

		if currentModel.ListOptions.IncludeCount != nil {
			includeCount = *currentModel.ListOptions.IncludeCount
		}

		if currentModel.ListOptions.ItemsPerPage != nil {
			itemsPerPage = *currentModel.ListOptions.ItemsPerPage
		}
	}

	listOptions := &admin20231115002.ListProjectIpAccessListsApiParams{
		GroupId:      *currentModel.ProjectId,
		IncludeCount: &includeCount,
		ItemsPerPage: &itemsPerPage,
		PageNum:      &pageNum,
	}

	result, resp, err := client.Atlas20231115002.ProjectIPAccessListApi.ListProjectIpAccessListsWithParams(context.Background(), listOptions).Execute()
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp), nil
	}

	mm := make([]AccessListDefinition, 0)
	for i := range result.Results {
		var m AccessListDefinition
		m.completeByConnection(result.Results[i])
		mm = append(mm, m)
	}
	currentModel.AccessList = mm
	// create list with 1
	models := []interface{}{}
	models = append(models, currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  models,
	}, nil
}
