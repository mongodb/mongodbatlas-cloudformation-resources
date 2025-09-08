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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"
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

func newPaginatedNetworkAccess(model *Model) (*admin20231115002.PaginatedNetworkAccess, error) {
	var accesslist []admin20231115002.NetworkPermissionEntry
	for i := range model.AccessList {
		modelAccessList := model.AccessList[i]
		projectIPAccessList := admin20231115002.NetworkPermissionEntry{}

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

	return &admin20231115002.PaginatedNetworkAccess{
		Results: accesslist,
	}, nil
}

func getEntry(entry AccessListDefinition) (string, error) {
	if util.IsStringPresent(entry.CIDRBlock) {
		return *entry.CIDRBlock, nil
	}

	if util.IsStringPresent(entry.AwsSecurityGroup) {
		return *entry.AwsSecurityGroup, nil
	}

	if util.IsStringPresent(entry.IPAddress) {
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

		if _, resp, err := client.Atlas20231115002.ProjectIPAccessListApi.DeleteProjectIpAccessList(context.Background(), projectID, entry).Execute(); err != nil {
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

		if _, resp, err := client.Atlas20231115002.ProjectIPAccessListApi.DeleteProjectIpAccessList(context.Background(), *model.ProjectId, entry).Execute(); err != nil {
			return progressevents.GetFailedEventByResponse(fmt.Sprintf("Error deleting the resource: %s", err.Error()),
				resp)
		}
	}

	return handler.ProgressEvent{}
}

func getAllEntries(client *util.MongoDBClient, projectID string) (*admin20231115002.PaginatedNetworkAccess, error) {
	includeCount := true
	itemPerPage := 500

	listOptions := &admin20231115002.ListProjectIpAccessListsApiParams{
		GroupId:      projectID,
		IncludeCount: &includeCount,
		ItemsPerPage: &itemPerPage,
	}
	accessList, _, err := client.Atlas20231115002.ProjectIPAccessListApi.ListProjectIpAccessListsWithParams(context.Background(), listOptions).Execute()
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

func newAccessListMap(accessList []admin20231115002.NetworkPermissionEntry) map[string]bool {
	m := make(map[string]bool)
	for _, entry := range accessList {
		if util.IsStringPresent(entry.CidrBlock) {
			m[*entry.CidrBlock] = true
			continue
		}

		if util.IsStringPresent(entry.IpAddress) {
			m[*entry.IpAddress] = true
			continue
		}

		if util.IsStringPresent(entry.AwsSecurityGroup) {
			m[*entry.AwsSecurityGroup] = true
			continue
		}
	}
	return m
}

func (m *AccessListDefinition) completeByConnection(c admin20231115002.NetworkPermissionEntry) {
	m.IPAddress = c.IpAddress
	m.CIDRBlock = c.CidrBlock
	m.Comment = c.Comment
	m.AwsSecurityGroup = c.AwsSecurityGroup
}
