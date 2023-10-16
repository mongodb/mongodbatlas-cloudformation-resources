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
	"runtime"

	"github.com/mongodb-forks/digest"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/version"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type AtlasEnvOptions struct {
	OrgID      string
	PrivateKey string
	PublicKey  string
	BaseURL    string
}

const (
	cfnTool = "mongodbatlas-cloudformation-resources"
)

var (
	toolName  = cfnTool
	userAgent = fmt.Sprintf("%s/%s (%s;%s)", toolName, version.Version, runtime.GOOS, runtime.GOARCH)
)

func NewAtlasTeam(ctx context.Context, client *atlas.Client, name string, orgID string) (*atlas.Team, error) {
	orgUser, _ := getExistingOrgUser(ctx, client, orgID)
	teamRequest := atlas.Team{
		Name:      name,
		Usernames: []string{orgUser.Username},
	}
	team, _, err := client.Teams.Create(ctx, orgID, &teamRequest)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func NewMongoDBClient() (atlasClient *atlas.Client, err error) {
	atlasEnv, err := getAtlasEnv()
	if err != nil {
		return nil, err
	}
	t := digest.NewTransport(atlasEnv.PublicKey, atlasEnv.PrivateKey)
	client, _ := t.Client()

	opts := []atlas.ClientOpt{atlas.SetUserAgent(userAgent)}
	if baseURL := atlasEnv.BaseURL; baseURL != "" {
		opts = append(opts, atlas.SetBaseURL(baseURL))
	}

	mongodbClient, err := atlas.New(client, opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to create Atlas client")
	}

	return mongodbClient, nil
}

func getAtlasEnv() (atlasEnvOpts *AtlasEnvOptions, err error) {
	orgID := os.Getenv("MONGODB_ATLAS_ORG_ID")
	publicKey := os.Getenv("ATLAS_PUBLIC_KEY")
	privateKey := os.Getenv("ATLAS_PRIVATE_KEY")
	baseURL := os.Getenv("MONGODB_ATLAS_BASE_URL")

	if orgID == "" || publicKey == "" || privateKey == "" || baseURL == "" {
		return nil, fmt.Errorf("please ensure following env variables are set: " +
			"MONGODB_ATLAS_ORG_ID, ATLAS_PUBLIC_KEY, ATLAS_PRIVATE_KEY, MONGODB_ATLAS_BASE_URL, MONGODB_ATLAS_SECRET_PROFILE")
	}

	return &AtlasEnvOptions{orgID, privateKey, publicKey, baseURL}, nil
}

func getExistingOrgUser(ctx context.Context, client *atlas.Client, orgID string) (*atlas.AtlasUser, error) {
	usersResponse, _, err := client.Organizations.Users(ctx, orgID, &atlas.ListOptions{})
	if err != nil {
		return nil, err
	}
	return &usersResponse.Results[0], nil
}
