// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package teamuser

import (
	"context"
	"errors"
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/testutil/mocksvc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func TestInitUserSet(t *testing.T) {
	user1 := "user1"
	user2 := "user2"
	user3 := "user3"
	// Create a slice of CloudAppUser objects
	users := []atlasv2.CloudAppUser{
		{Id: &user1},
		{Id: &user2},
		{Id: &user3},
	}

	// Call the initUserSet function to create a map of user IDs
	usersSet := initUserSet(users)

	// Check that the map contains the expected user IDs
	assert.Contains(t, usersSet, user1)
	assert.Contains(t, usersSet, user2)
	assert.Contains(t, usersSet, user3)

	// Check that the map does not contain any unexpected user IDs
	assert.NotContains(t, usersSet, "user4")
	assert.NotContains(t, usersSet, "user5")
	assert.NotContains(t, usersSet, "user6")
}

func TestValidateUsernames(t *testing.T) {
	mockClient := mocksvc.NewTeamUsersAPI(t)

	// Create a slice of usernames, including some valid and some invalid usernames
	validuser1 := "validuser1"
	validuser2 := "validuser2"
	usernames := []string{validuser1, validuser2}

	// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
	mockClient.EXPECT().GetUserByUsername(context.Background(), validuser1).Return(&atlasv2.CloudAppUser{Id: &validuser1}, nil, nil)
	mockClient.EXPECT().GetUserByUsername(context.Background(), validuser2).Return(&atlasv2.CloudAppUser{Id: &validuser2}, nil, nil)

	// Call the ValidateUsernames function
	validUsers, _, err := validateUsernames(mockClient, usernames)

	require.NoError(t, err)

	// Check that the returned slice contains only the valid usernames
	assert.Len(t, validUsers, 2)
	assert.Equal(t, validuser1, validUsers[0].GetId())
	assert.Equal(t, validuser2, validUsers[1].GetId())
}

func TestValidateUsernamesWithInvalidInput(t *testing.T) {
	mockAtlasV2Client := mocksvc.NewTeamUsersAPI(t)

	// Create a slice of usernames, including some valid and some invalid usernames
	validuser1 := "validuser1"
	invaliduser1 := "invaliduser1"
	usernames := []string{validuser1, invaliduser1}

	// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
	mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), validuser1).Return(&atlasv2.CloudAppUser{Id: &validuser1}, nil, nil)
	mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), invaliduser1).Return(nil, nil, errors.New("invalid username"))

	// Call the ValidateUsernames function
	_, _, err := validateUsernames(mockAtlasV2Client, usernames)

	require.Error(t, err)
}

func TestGetChangesForTeamUsers(t *testing.T) {
	user1 := "user1"
	user2 := "user2"
	user3 := "user3"

	// Test cases
	testCases := []struct {
		testName         string
		currentUsers     []atlasv2.CloudAppUser
		newUsers         []atlasv2.CloudAppUser
		expectedToAdd    []string
		expectedToDelete []string
	}{
		{
			testName: "succeeds adding a new user and removing an existing one",
			currentUsers: []atlasv2.CloudAppUser{
				{Id: &user1},
				{Id: &user2},
			},
			newUsers: []atlasv2.CloudAppUser{
				{Id: &user1},
				{Id: &user3},
			},
			expectedToAdd:    []string{user3},
			expectedToDelete: []string{user2},
		},
		{
			testName:     "succeeds adding all users",
			currentUsers: []atlasv2.CloudAppUser{},
			newUsers: []atlasv2.CloudAppUser{
				{Id: &user1},
				{Id: &user2},
			},
			expectedToAdd:    []string{user1, user2},
			expectedToDelete: []string{},
		},
		{
			testName: "succeeds removing both users",
			currentUsers: []atlasv2.CloudAppUser{
				{Id: &user1},
				{Id: &user2},
			},
			newUsers:         []atlasv2.CloudAppUser{},
			expectedToAdd:    []string{},
			expectedToDelete: []string{user1, user2},
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			toAdd, toDelete, err := getChangesForTeamUsers(testCase.currentUsers, testCase.newUsers)
			require.NoError(t, err)
			assert.ElementsMatch(t, testCase.expectedToAdd, toAdd)
			assert.ElementsMatch(t, testCase.expectedToDelete, toDelete)
		})
	}
}

func TestUpdateTeamUsers(t *testing.T) {
	validuser1 := "validuser1"
	validuser2 := "validuser2"
	invaliduser1 := "invaliduser1"

	testCases := []struct {
		mockFuncExpectations func(*mocksvc.TeamUsersAPI)
		existingTeamUsers    *atlasv2.PaginatedApiAppUser
		expectError          require.ErrorAssertionFunc
		testName             string
		usernames            []string
	}{
		{
			testName: "succeeds but no changes are required",
			mockFuncExpectations: func(mockClient *mocksvc.TeamUsersAPI) {
				// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
				mockClient.EXPECT().GetUserByUsername(context.Background(), validuser1).Return(&atlasv2.CloudAppUser{Id: &validuser1}, nil, nil)
				mockClient.EXPECT().GetUserByUsername(context.Background(), validuser2).Return(&atlasv2.CloudAppUser{Id: &validuser2}, nil, nil)
			},
			existingTeamUsers: &atlasv2.PaginatedApiAppUser{Results: []atlasv2.CloudAppUser{{Id: &validuser1}, {Id: &validuser2}}},
			usernames:         []string{validuser1, validuser2},
			expectError:       require.NoError,
		},
		{
			testName: "fails because one user is invalid",
			mockFuncExpectations: func(mockClient *mocksvc.TeamUsersAPI) {
				// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
				mockClient.EXPECT().GetUserByUsername(context.Background(), validuser1).Return(&atlasv2.CloudAppUser{Id: &validuser1}, nil, nil)
				mockClient.EXPECT().GetUserByUsername(context.Background(), invaliduser1).Return(nil, nil, errors.New("invalid username"))
			},
			existingTeamUsers: &atlasv2.PaginatedApiAppUser{Results: []atlasv2.CloudAppUser{{Id: &validuser1}, {Id: &invaliduser1}}},
			usernames:         []string{validuser1, invaliduser1},
			expectError:       require.Error,
		},
		{
			testName: "succeeds with one user to be added",
			mockFuncExpectations: func(mockClient *mocksvc.TeamUsersAPI) {
				// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
				mockClient.EXPECT().GetUserByUsername(context.Background(), validuser1).Return(&atlasv2.CloudAppUser{Id: &validuser1}, nil, nil)
				mockClient.EXPECT().GetUserByUsername(context.Background(), validuser2).Return(&atlasv2.CloudAppUser{Id: &validuser2}, nil, nil)

				var newUsers []atlasv2.AddUserToTeam
				newUsers = append(newUsers, atlasv2.AddUserToTeam{Id: validuser2})
				mockClient.EXPECT().AddTeamUser(context.Background(), "orgID", "teamID", &newUsers).Return(nil, nil, nil)
			},
			existingTeamUsers: &atlasv2.PaginatedApiAppUser{Results: []atlasv2.CloudAppUser{{Id: &validuser1}}},
			usernames:         []string{validuser1, validuser2},
			expectError:       require.NoError,
		},
		{
			testName: "succeeds with one user to be removed",
			mockFuncExpectations: func(mockClient *mocksvc.TeamUsersAPI) {
				// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
				mockClient.EXPECT().GetUserByUsername(context.Background(), validuser2).Return(&atlasv2.CloudAppUser{Id: &validuser2}, nil, nil)
				mockClient.EXPECT().RemoveTeamUser(context.Background(), "orgID", "teamID", validuser1).Return(nil, nil)
			},
			existingTeamUsers: &atlasv2.PaginatedApiAppUser{Results: []atlasv2.CloudAppUser{{Id: &validuser1}, {Id: &validuser2}}},
			usernames:         []string{validuser2},
			expectError:       require.NoError,
		},
		{
			testName: "succeeds with one user to be added and the other removed",
			mockFuncExpectations: func(mockClient *mocksvc.TeamUsersAPI) {
				// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
				mockClient.EXPECT().GetUserByUsername(context.Background(), validuser1).Return(&atlasv2.CloudAppUser{Id: &validuser1}, nil, nil)
				mockClient.EXPECT().RemoveTeamUser(context.Background(), "orgID", "teamID", validuser2).Return(nil, nil)
				var newUsers []atlasv2.AddUserToTeam
				newUsers = append(newUsers, atlasv2.AddUserToTeam{Id: validuser1})
				mockClient.EXPECT().AddTeamUser(context.Background(), "orgID", "teamID", &newUsers).Return(nil, nil, nil)
			},
			existingTeamUsers: &atlasv2.PaginatedApiAppUser{Results: []atlasv2.CloudAppUser{{Id: &validuser2}}},
			usernames:         []string{validuser1},
			expectError:       require.NoError,
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			mockAtlasV2Client := mocksvc.NewTeamUsersAPI(t)
			testCase.mockFuncExpectations(mockAtlasV2Client)
			testCase.expectError(t, UpdateTeamUsers(mockAtlasV2Client, testCase.existingTeamUsers, testCase.usernames, "orgID", "teamID"))
		})
	}
}
