// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utility

import (
	"context"
	"fmt"
	"os"

	"github.com/mongodb-forks/digest"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"go.mongodb.org/atlas-sdk/v20231115002/admin"
)

type AtlasEnvOptions struct {
	OrgID      string
	PrivateKey string
	PublicKey  string
	BaseURL    string
}

func NewAtlasTeam(ctx context.Context, client *admin.APIClient, name string, orgID string) (*admin.Team, error) {
	orgUser, _ := getExistingOrgUser(ctx, client, orgID)
	teamRequest := admin.Team{
		Name:      name,
		Usernames: []string{orgUser.Username},
	}
	team, _, err := client.TeamsApi.CreateTeam(ctx, orgID, &teamRequest).Execute()
	if err != nil {
		return nil, err
	}
	return team, nil
}

func NewMongoDBClient() (atlasClient *admin.APIClient, err error) {
	atlasEnv, err := getAtlasEnv()
	if err != nil {
		return nil, err
	}
	t := digest.NewTransport(atlasEnv.PublicKey, atlasEnv.PrivateKey)
	client, _ := t.Client()

	c := util.Config{}
	if baseURL := atlasEnv.BaseURL; baseURL != "" {
		c.BaseURL = baseURL
	}
	// New SDK Client
	sdkV2Client, err := c.NewSDKV2Client(client)
	if err != nil {
		return nil, fmt.Errorf("unable to create Atlas client")
	}

	return sdkV2Client, nil
}

func getAtlasEnv() (atlasEnvOpts *AtlasEnvOptions, err error) {
	orgID := os.Getenv("ATLAS_ORG_ID")
	publicKey := os.Getenv("ATLAS_PUBLIC_KEY")
	privateKey := os.Getenv("ATLAS_PRIVATE_KEY")
	baseURL := os.Getenv("ATLAS_BASE_URL")

	if orgID == "" || publicKey == "" || privateKey == "" || baseURL == "" {
		return nil, fmt.Errorf("please ensure following env variables are set: " +
			"ATLAS_ORG_ID, ATLAS_PUBLIC_KEY, ATLAS_PRIVATE_KEY, ATLAS_BASE_URL, ATLAS_SECRET_PROFILE")
	}

	return &AtlasEnvOptions{orgID, privateKey, publicKey, baseURL}, nil
}

func getExistingOrgUser(ctx context.Context, client *admin.APIClient, orgID string) (*admin.CloudAppUser, error) {
	usersResponse, _, err := client.OrganizationsApi.ListOrganizationUsers(ctx, orgID).Execute()
	if err != nil {
		return nil, err
	}
	return &usersResponse.Results[0], nil
}
