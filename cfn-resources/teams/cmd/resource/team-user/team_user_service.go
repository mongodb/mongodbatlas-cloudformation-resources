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

//go:generate mockgen -source=team_user_service.go -destination=./mocks/mocks.go -package=mocks
package teamuser

import (
	"context"
	"net/http"

	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

// Service interface was defined over mocking direct SDK as specific mapping of input parameters to responses is needed
type TeamUsersAPI interface {
	GetUserByUsername(ctx context.Context, userName string) (*admin20231115002.CloudAppUser, *http.Response, error)
	AddTeamUser(ctx context.Context, orgID string, teamID string, addUserToTeam *[]admin20231115002.AddUserToTeam) (*admin20231115002.PaginatedApiAppUser, *http.Response, error)
	RemoveTeamUser(ctx context.Context, orgID string, teamID string, userID string) (*http.Response, error)
}

type TeamUsersAPIService struct {
	mongoDBCloudUsersAPI admin20231115002.MongoDBCloudUsersApi
	teamsAPI             admin20231115002.TeamsApi
}

func NewTeamUsersAPIService(client *admin20231115002.APIClient) *TeamUsersAPIService {
	return &TeamUsersAPIService{
		mongoDBCloudUsersAPI: client.MongoDBCloudUsersApi,
		teamsAPI:             client.TeamsApi,
	}
}

func (s *TeamUsersAPIService) GetUserByUsername(ctx context.Context, userName string) (*admin20231115002.CloudAppUser, *http.Response, error) {
	return s.mongoDBCloudUsersAPI.GetUserByUsername(context.Background(), userName).Execute()
}

func (s *TeamUsersAPIService) AddTeamUser(ctx context.Context, orgID string, teamID string, addUserToTeam *[]admin20231115002.AddUserToTeam) (*admin20231115002.PaginatedApiAppUser, *http.Response, error) {
	return s.teamsAPI.AddTeamUser(ctx, orgID, teamID, addUserToTeam).Execute()
}

func (s *TeamUsersAPIService) RemoveTeamUser(ctx context.Context, orgID string, teamID string, userID string) (*http.Response, error) {
	return s.teamsAPI.RemoveTeamUser(ctx, orgID, teamID, userID).Execute()
}
