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
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
	"go.mongodb.org/atlas/mongodbatlas"
)

func setup() {
	util.SetupLogger("mongodb-atlas-project-ip-access-list")
}

var CreateRequiredFields = []string{constants.ProjectID, constants.AccessList}
var ReadRequiredFields = []string{constants.ProjectID}
var UpdateRequiredFields = []string{constants.ProjectID, constants.AccessList}
var DeleteRequiredFields = []string{constants.ProjectID}
var ListRequiredFields = []string{constants.ProjectID}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return CreateOp(req, prevModel, currentModel)
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return ReadOp(req, prevModel, currentModel)
}

// Update handles the Update event from the Cloudformation service.
// Logic: Atlas does not provide an endpoint to update a single entry in the accesslist.
// As a result, we delete all the entries in the current model + previous model
// and then create all the entries in the current model.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return UpdateOp(req, prevModel, currentModel)
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return DeleteOp(req, prevModel, currentModel)
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	// Create atlas client
	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	var pageNum, itemsPerPage int
	var includeCount bool

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

	listOptions := &mongodbatlas.ListOptions{
		PageNum:      pageNum,
		IncludeCount: includeCount,
		ItemsPerPage: itemsPerPage,
	}

	result, resp, err := client.ProjectIPAccessList.List(context.Background(), *currentModel.ProjectId, listOptions)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			resp.Response), nil
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

func newPaginatedNetworkAccess(model *Model) (*admin.PaginatedNetworkAccess, error) {
	var accesslist []admin.NetworkPermissionEntry
	for i := range model.AccessList {
		modelAccessList := model.AccessList[i]
		projectIPAccessList := admin.NetworkPermissionEntry{}

		if modelAccessList.DeleteAfterDate != nil {
			deleteAfterDate, err := util.StringToTime(*modelAccessList.DeleteAfterDate)
			if err != nil {
				return nil, err
			}
			projectIPAccessList.DeleteAfterDate = &deleteAfterDate
		}

		if modelAccessList.Comment != nil {
			projectIPAccessList.Comment = modelAccessList.Comment
		}

		if modelAccessList.CIDRBlock != nil {
			projectIPAccessList.CidrBlock = modelAccessList.CIDRBlock
		}

		if modelAccessList.IPAddress != nil {
			projectIPAccessList.IpAddress = modelAccessList.IPAddress
		}

		if modelAccessList.AwsSecurityGroup != nil {
			projectIPAccessList.AwsSecurityGroup = modelAccessList.AwsSecurityGroup
		}

		accesslist = append(accesslist, projectIPAccessList)
	}

	return &admin.PaginatedNetworkAccess{
		Results: accesslist,
	}, nil
}

func getProjectIPAccessListRequest(model *Model) []*mongodbatlas.ProjectIPAccessList {
	var accesslist []*mongodbatlas.ProjectIPAccessList
	for i := range model.AccessList {
		modelAccessList := model.AccessList[i]
		projectIPAccessList := &mongodbatlas.ProjectIPAccessList{}

		if modelAccessList.DeleteAfterDate != nil {
			projectIPAccessList.DeleteAfterDate = *modelAccessList.DeleteAfterDate
		}

		if modelAccessList.Comment != nil {
			projectIPAccessList.Comment = *modelAccessList.Comment
		}

		if modelAccessList.CIDRBlock != nil {
			projectIPAccessList.CIDRBlock = *modelAccessList.CIDRBlock
		}

		if modelAccessList.IPAddress != nil {
			projectIPAccessList.IPAddress = *modelAccessList.IPAddress
		}

		if modelAccessList.AwsSecurityGroup != nil {
			projectIPAccessList.AwsSecurityGroup = *modelAccessList.AwsSecurityGroup
		}

		accesslist = append(accesslist, projectIPAccessList)
	}

	_, _ = logger.Debugf("getProjectIPAccessListRequest accesslist:%v", accesslist)
	return accesslist
}

func getEntry(entry AccessListDefinition) (string, error) {
	if entry.CIDRBlock != nil && *entry.CIDRBlock != "" {
		return *entry.CIDRBlock, nil
	}

	if entry.AwsSecurityGroup != nil && *entry.AwsSecurityGroup != "" {
		return *entry.AwsSecurityGroup, nil
	}

	if entry.IPAddress != nil && *entry.IPAddress != "" {
		return *entry.IPAddress, nil
	}

	return "", fmt.Errorf("AccessList entry must have one of the following fields: cidrBlock, awsSecurityGroup, ipAddress")
}

// deleteEntriesForUpdate deletes entries in the atlas access list without failing if the entry is NOT_FOUND.
// This function is used in the update handler where we don't want to fail if the entry is not found.
// Note: The delete handler MUST fail if the entry is not found otherwise "cfn test" will fail.
func deleteEntriesForUpdate(list []AccessListDefinition, projectID string, client *util.MongoDBClient) handler.ProgressEvent {
	for _, accessListEntry := range list {
		entry, err := getEntry(accessListEntry)
		if err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting the resource: %s", err.Error()),
				nil)
		}

		if _, resp, err := client.AtlasV2.ProjectIPAccessListApi.DeleteProjectIpAccessList(context.Background(), projectID, entry).Execute(); err != nil {
			if resp.StatusCode == http.StatusNotFound {
				_, _ = logger.Warnf("Accesslist entry Not Found: %s, err:%+v", entry, err)
				continue
			}

			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error deleting the resource: %s", err.Error()),
				resp)
		}
	}

	return handler.ProgressEvent{}
}

func deleteEntries(model *Model, client *util.MongoDBClient) handler.ProgressEvent {
	for _, accessListEntry := range model.AccessList {
		entry, err := getEntry(accessListEntry)
		if err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error getting the resource: %s", err.Error()),
				nil)
		}

		if _, resp, err := client.AtlasV2.ProjectIPAccessListApi.DeleteProjectIpAccessList(context.Background(), *model.ProjectId, entry).Execute(); err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error deleting the resource: %s", err.Error()),
				resp)
		}
	}

	return handler.ProgressEvent{}
}

func getAllEntries(client *util.MongoDBClient, projectID string) (*admin.PaginatedNetworkAccess, error) {
	includeCount := true
	itemPerPage := 500

	listOptions := &admin.ListProjectIpAccessListsApiParams{
		GroupId:      projectID,
		IncludeCount: &includeCount,
		ItemsPerPage: &itemPerPage,
	}
	accessList, _, err := client.AtlasV2.ProjectIPAccessListApi.ListProjectIpAccessListsWithParams(context.Background(), listOptions).Execute()
	if err != nil {
		return nil, err
	}

	return accessList, nil
}

// isEntryAlreadyInAccessList checks if the entry already exists in the atlas access list
func isEntryAlreadyInAccessList(client *util.MongoDBClient, model *Model) (bool, error) {
	existingEntries, err := getAllEntries(client, *model.ProjectId)
	if err != nil {
		return false, err
	}

	existingEntriesMap := newAccessListMap(existingEntries.Results)
	for _, entry := range model.AccessList {
		if isEntryInMap(entry, existingEntriesMap) {
			return true, nil
		}
	}

	return false, nil
}

func isEntryInMap(entry AccessListDefinition, accessListMap map[string]bool) bool {
	if entry.CIDRBlock != nil && accessListMap[*entry.CIDRBlock] {
		return true
	}

	if entry.IPAddress != nil && accessListMap[*entry.IPAddress] {
		return true
	}

	if entry.AwsSecurityGroup != nil && accessListMap[*entry.AwsSecurityGroup] {
		return true
	}

	return false
}

func newAccessListMap(accessList []mongodbatlas.ProjectIPAccessList) map[string]bool {
	m := make(map[string]bool)
	for _, entry := range accessList {
		if entry.CIDRBlock != "" {
			m[entry.CIDRBlock] = true
			continue
		}

		if entry.IPAddress != "" {
			m[entry.IPAddress] = true
			continue
		}

		if entry.AwsSecurityGroup != "" {
			m[entry.AwsSecurityGroup] = true
			continue
		}
	}
	return m
}

func (m *AccessListDefinition) completeByConnection(c mongodbatlas.ProjectIPAccessList) {
	m.IPAddress = &c.IPAddress
	m.CIDRBlock = &c.CIDRBlock
	m.Comment = &c.Comment
	m.AwsSecurityGroup = &c.AwsSecurityGroup
}
