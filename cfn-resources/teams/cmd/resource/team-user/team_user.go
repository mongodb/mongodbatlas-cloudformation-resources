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

package teamuser

import (
	"context"
	"net/http"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240805001/admin"
)

func UpdateTeamUsers(c TeamUsersAPI, existingTeamUsers *atlasv2.PaginatedApiAppUser, usernames []string, orgID, teamID string) error {
	var newUsers []atlasv2.AddUserToTeam

	validUsernames, _, err := validateUsernames(c, usernames)
	if err != nil {
		return err
	}
	usersToAdd, usersToRemove, err := getChangesForTeamUsers(existingTeamUsers.Results, validUsernames)
	if err != nil {
		return err
	}

	for ind := range usersToRemove {
		// remove user from team
		_, err := c.RemoveTeamUser(context.Background(), orgID, teamID, util.SafeString(&usersToRemove[ind]))
		if err != nil {
			return err
		}
	}

	for ind := range usersToAdd {
		// add user to team
		newUsers = append(newUsers, atlasv2.AddUserToTeam{Id: util.SafeString(&usersToAdd[ind])})
	}
	// save all new users
	if len(newUsers) > 0 {
		_, _, err = c.AddTeamUser(context.Background(), orgID, teamID, &newUsers)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateUsernames(c TeamUsersAPI, usernames []string) ([]atlasv2.CloudAppUser, *http.Response, error) {
	var validUsers []atlasv2.CloudAppUser
	for _, elem := range usernames {
		userToAdd, httpResp, err := c.GetUserByUsername(context.Background(), elem)
		if err != nil {
			_, _ = logger.Warnf("Error while getting the user %s: (%+v) \n", elem, err)
			return nil, httpResp, err
		}
		validUsers = append(validUsers, *userToAdd)
	}
	return validUsers, nil, nil
}

func initUserSet(users []atlasv2.CloudAppUser) map[string]interface{} {
	usersSet := make(map[string]interface{}, len(users))
	for i := 0; i < len(users); i++ {
		usersSet[users[i].GetId()] = true
	}
	return usersSet
}

func getChangesForTeamUsers(currentUsers []atlasv2.CloudAppUser, newUsers []atlasv2.CloudAppUser) (toAdd []string, toDelete []string, err error) {
	// Create two sets to store the elements in A and B
	currentUsersSet := initUserSet(currentUsers)
	newUsersSet := initUserSet(newUsers)

	// Create two arrays to store the elements to be added and deleted
	toAdd = []string{}
	toDelete = []string{}

	// Iterate over the elements in B and add them to the toAdd array if they are not in A
	for elem := range newUsersSet {
		if _, ok := currentUsersSet[elem]; !ok {
			toAdd = append(toAdd, elem)
		}
	}

	// Iterate over the elements in A and add them to the toDelete array if they are not in B
	for elem := range currentUsersSet {
		if _, ok := newUsersSet[elem]; !ok {
			toDelete = append(toDelete, elem)
		}
	}

	// Return the two arrays
	return toAdd, toDelete, nil
}
