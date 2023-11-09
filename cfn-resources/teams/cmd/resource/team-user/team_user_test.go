package teamuser

import (
	"context"
	"errors"
	"reflect"
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

	mockAtlasV2Client := mock_util.NewMockUserFetcher(mockCtrl)

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
	assert.Equal(t, "validuser1", *validUsers[0].Id)
	assert.Equal(t, "validuser2", *validUsers[1].Id)
}

func TestFilterOnlyValidUsernamesWithInvalidInput(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAtlasV2Client := mock_util.NewMockUserFetcher(mockCtrl)

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

func TestGetUserDeltas1(t *testing.T) {
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
		toAdd, toDelete, err := GetUserDeltas(testCase.currentUsers, testCase.newUsers)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(toAdd, testCase.expectedToAdd) {
			t.Errorf("Expected toAdd to be %v, but got %v", testCase.expectedToAdd, toAdd)
		}

		if !reflect.DeepEqual(toDelete, testCase.expectedToDelete) {
			t.Errorf("Expected toDelete to be %v, but got %v", testCase.expectedToDelete, toDelete)
		}
	}
}
