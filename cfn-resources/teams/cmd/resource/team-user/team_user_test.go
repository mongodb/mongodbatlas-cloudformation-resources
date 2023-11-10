//go:build unit

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

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231001001/admin"

	mock_util "github.com/mongodb/mongodbatlas-cloudformation-resources/test/mocks"
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
	assert.True(t, usersSet[user1])
	assert.True(t, usersSet[user2])
	assert.True(t, usersSet[user3])

	// Check that the map does not contain any unexpected user IDs
	assert.False(t, usersSet["user4"])
	assert.False(t, usersSet["user5"])
	assert.False(t, usersSet["user6"])
}

func TestFilterOnlyValidUsernames(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAtlasV2Client := mock_util.NewMockTeamUsersAPI(mockCtrl)

	// Create a slice of usernames, including some valid and some invalid usernames
	validuser1 := "validuser1"
	validuser2 := "validuser2"
	usernames := []string{validuser1, validuser2}

	// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
	mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), validuser1).Return(&atlasv2.CloudAppUser{Id: &validuser1}, nil, nil)
	mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), validuser2).Return(&atlasv2.CloudAppUser{Id: &validuser2}, nil, nil)

	// Call the FilterOnlyValidUsernames function
	validUsers, _, err := FilterOnlyValidUsernames(mockAtlasV2Client, usernames)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check that the returned slice contains only the valid usernames
	assert.Equal(t, len(validUsers), 2)
	assert.Equal(t, validuser1, *validUsers[0].Id)
	assert.Equal(t, validuser2, *validUsers[1].Id)
}

func TestFilterOnlyValidUsernamesWithInvalidInput(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAtlasV2Client := mock_util.NewMockTeamUsersAPI(mockCtrl)

	// Create a slice of usernames, including some valid and some invalid usernames
	validuser1 := "validuser1"
	invaliduser1 := "invaliduser1"
	usernames := []string{validuser1, invaliduser1}

	// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
	mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), validuser1).Return(&atlasv2.CloudAppUser{Id: &validuser1}, nil, nil)
	mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), invaliduser1).Return(nil, nil, errors.New("invalid username"))

	// Call the FilterOnlyValidUsernames function
	_, _, err := FilterOnlyValidUsernames(mockAtlasV2Client, usernames)

	assert.NotNil(t, err)
}

func TestGetUsersToAddAndRemove(t *testing.T) {
	user1 := "user1"
	user2 := "user2"
	user3 := "user3"

	// Test cases
	testCases := []struct {
		currentUsers     []atlasv2.CloudAppUser
		newUsers         []atlasv2.CloudAppUser
		expectedToAdd    []string
		expectedToDelete []string
	}{
		{
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
			currentUsers: []atlasv2.CloudAppUser{},
			newUsers: []atlasv2.CloudAppUser{
				{Id: &user1},
				{Id: &user2},
			},
			expectedToAdd:    []string{user1, user2},
			expectedToDelete: []string{},
		},
		{
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
		toAdd, toDelete, err := GetUsersToAddAndRemove(testCase.currentUsers, testCase.newUsers)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		assert.Equal(t, testCase.expectedToAdd, toAdd)
		assert.Equal(t, testCase.expectedToDelete, toDelete)
	}
}

func TestUpdateTeamUsers(t *testing.T) {
	validuser1 := "validuser1"
	validuser2 := "validuser2"
	invaliduser1 := "invaliduser1"

	testCases := []struct {
		mockFuncExpectations func(*gomock.Controller) *mock_util.MockTeamUsersAPI
		existingTeamUsers    *atlasv2.PaginatedApiAppUser
		usernames            []string
		expectError          bool
	}{
		{ // Succeeds but no changes are required
			mockFuncExpectations: func(mockCtrl *gomock.Controller) *mock_util.MockTeamUsersAPI {
				mockAtlasV2Client := mock_util.NewMockTeamUsersAPI(mockCtrl)

				// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
				mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), validuser1).Return(&atlasv2.CloudAppUser{Id: &validuser1}, nil, nil)
				mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), validuser2).Return(&atlasv2.CloudAppUser{Id: &validuser2}, nil, nil)

				return mockAtlasV2Client
			},
			existingTeamUsers: &atlasv2.PaginatedApiAppUser{Results: []atlasv2.CloudAppUser{{Id: &validuser1}, {Id: &validuser2}}},
			usernames:         []string{validuser1, validuser2},
			expectError:       false,
		},
		{ // Fails because one user is invalid
			mockFuncExpectations: func(mockCtrl *gomock.Controller) *mock_util.MockTeamUsersAPI {
				mockAtlasV2Client := mock_util.NewMockTeamUsersAPI(mockCtrl)

				// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
				mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), validuser1).Return(&atlasv2.CloudAppUser{Id: &validuser1}, nil, nil)
				mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), invaliduser1).Return(nil, nil, errors.New("invalid username"))

				return mockAtlasV2Client
			},
			existingTeamUsers: &atlasv2.PaginatedApiAppUser{Results: []atlasv2.CloudAppUser{{Id: &validuser1}, {Id: &invaliduser1}}},
			usernames:         []string{validuser1, invaliduser1},
			expectError:       true,
		},
		{ // Succeeds and one user has to be added
			mockFuncExpectations: func(mockCtrl *gomock.Controller) *mock_util.MockTeamUsersAPI {
				mockAtlasV2Client := mock_util.NewMockTeamUsersAPI(mockCtrl)

				// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
				mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), validuser1).Return(&atlasv2.CloudAppUser{Id: &validuser1}, nil, nil)
				mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), validuser2).Return(&atlasv2.CloudAppUser{Id: &validuser2}, nil, nil)

				var newUsers []atlasv2.AddUserToTeam
				newUsers = append(newUsers, atlasv2.AddUserToTeam{Id: validuser2})
				mockAtlasV2Client.EXPECT().AddTeamUser(context.Background(), "orgID", "teamID", &newUsers)

				return mockAtlasV2Client
			},
			existingTeamUsers: &atlasv2.PaginatedApiAppUser{Results: []atlasv2.CloudAppUser{{Id: &validuser1}}},
			usernames:         []string{validuser1, validuser2},
			expectError:       false,
		},
		{ // Succeeds and one user has to be removed
			mockFuncExpectations: func(mockCtrl *gomock.Controller) *mock_util.MockTeamUsersAPI {
				mockAtlasV2Client := mock_util.NewMockTeamUsersAPI(mockCtrl)

				// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
				mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), validuser2).Return(&atlasv2.CloudAppUser{Id: &validuser2}, nil, nil)
				mockAtlasV2Client.EXPECT().RemoveTeamUser(context.Background(), "orgID", "teamID", validuser1)

				return mockAtlasV2Client
			},
			existingTeamUsers: &atlasv2.PaginatedApiAppUser{Results: []atlasv2.CloudAppUser{{Id: &validuser1}, {Id: &validuser2}}},
			usernames:         []string{validuser2},
			expectError:       false,
		},
		{ // Succeeds and one user has to be added and the other removed
			mockFuncExpectations: func(mockCtrl *gomock.Controller) *mock_util.MockTeamUsersAPI {
				mockAtlasV2Client := mock_util.NewMockTeamUsersAPI(mockCtrl)

				// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
				mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), validuser1).Return(&atlasv2.CloudAppUser{Id: &validuser1}, nil, nil)
				mockAtlasV2Client.EXPECT().RemoveTeamUser(context.Background(), "orgID", "teamID", validuser2)
				var newUsers []atlasv2.AddUserToTeam
				newUsers = append(newUsers, atlasv2.AddUserToTeam{Id: validuser1})
				mockAtlasV2Client.EXPECT().AddTeamUser(context.Background(), "orgID", "teamID", &newUsers)

				return mockAtlasV2Client
			},
			existingTeamUsers: &atlasv2.PaginatedApiAppUser{Results: []atlasv2.CloudAppUser{{Id: &validuser2}}},
			usernames:         []string{validuser1},
			expectError:       false,
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		mockCtrl := gomock.NewController(t)
		client := testCase.mockFuncExpectations(mockCtrl)

		err := UpdateTeamUsers(client, testCase.existingTeamUsers, testCase.usernames, "orgID", "teamID")

		assert.Equal(t, testCase.expectError, err != nil)

		mockCtrl.Finish()
	}
}
