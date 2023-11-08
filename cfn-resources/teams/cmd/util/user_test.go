package util

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231001001/admin"

	mock_admin "github.com/mongodb/mongodbatlas-cloudformation-resources/test/mocks"
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

	mockAtlasV2Client := mock_admin.NewMockMongoDBCloudUsersApi(mockCtrl)

	// Create a slice of usernames, including some valid and some invalid usernames
	usernames := []string{"validuser1", "invaliduser1", "validuser2", "invaliduser2"}

	// Set up mock expectations for successful and unsuccessful calls to GetUserByUsername
	mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), "validuser1").Return(&atlasv2.CloudAppUser{}, nil)
	mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), "invaliduser1").Return(nil, errors.New("invalid username"))
	mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), "validuser2").Return(&atlasv2.CloudAppUser{}, nil)
	mockAtlasV2Client.EXPECT().GetUserByUsername(context.Background(), "invaliduser2").Return(nil, errors.New("invalid username"))

	// Call the FilterOnlyValidUsernames function
	validUsers := FilterOnlyValidUsernames(mockAtlasV2Client, usernames)

	// Check that the returned slice contains only the valid usernames
	assert.Equal(t, len(validUsers), 2)
	assert.Equal(t, validUsers[0].Username, "validuser1")
	assert.Equal(t, validUsers[1].Username, "validuser2")
}
