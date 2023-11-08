package util

import (
	"context"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231001001/admin"
)

func FilterOnlyValidUsernames(mongoDBCloudUsersApiClient atlasv2.MongoDBCloudUsersApi, usernames []string) []atlasv2.CloudAppUser {
	var validUsers []atlasv2.CloudAppUser
	for _, elem := range usernames {
		userToAdd, _, err := mongoDBCloudUsersApiClient.GetUserByUsername(context.Background(), elem).Execute()
		if err != nil {
			_, _ = logger.Warnf("Error while getting the user by username %s: (%+v) \n", elem, err)
		} else {
			validUsers = append(validUsers, *userToAdd)
		}
	}
	return validUsers
}

func initUserSet(users []atlasv2.CloudAppUser) map[string]bool {
	usersSet := make(map[string]bool)
	for _, elem := range users {
		usersSet[*elem.Id] = true
	}
	return usersSet
}

func GetUserDeltas(currentUsers []atlasv2.CloudAppUser, newUsers []atlasv2.CloudAppUser) ([]string, []string, error) {
	// Create two sets to store the elements in A and B
	currentUsersSet := initUserSet(currentUsers)
	newUsersSet := initUserSet(newUsers)

	// Create two arrays to store the elements to be added and deleted
	toAdd := []string{}
	toDelete := []string{}

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