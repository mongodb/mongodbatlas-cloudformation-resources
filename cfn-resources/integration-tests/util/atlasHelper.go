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

package util

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/mongodb-forks/digest"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/version"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	realmAuth "go.mongodb.org/realm/auth"
	"go.mongodb.org/realm/realm"
)

type AtlasEnvOptions struct {
	OrgId      string
	PrivateKey string
	PublicKey  string
	BaseURL    string
}

const (
	cfnTool = "mongodbatlas-cloudformation-resources"
)

var (
	toolName           = cfnTool
	defaultLogLevel    = "warning"
	userAgent          = fmt.Sprintf("%s/%s (%s;%s)", toolName, version.Version, runtime.GOOS, runtime.GOARCH)
	terraformUserAgent = "terraform-provider-mongodbatlas"
)

func GetNewAtlasTeam(ctx context.Context, client *atlas.Client, name string, orgId string) (*atlas.Team, error) {
	orgUser, err := getExistingOrgUser(ctx, client, orgId)
	teamRequest := atlas.Team{
		Name:      name,
		Usernames: []string{orgUser.Username},
	}
	team, _, err := client.Teams.Create(ctx, orgId, &teamRequest)
	if err != nil {
		log.Println("Error when creating Atlas Team")
		return nil, err
	}
	return team, nil
}

func GetNewAtlasProject(ctx context.Context, client *atlas.Client, name string, orgId string) (*atlas.Project, error) {
	project, _, err := client.Projects.Create(ctx, &atlas.Project{
		Name:  name,
		OrgID: orgId,
	}, &atlas.CreateProjectOptions{})

	if err != nil {
		log.Println("Unable to create AWS client, please check AWS config is correctly setup")
		return nil, err
	}

	return project, nil
}

func GetRealmClient(ctx context.Context) (*realm.Client, error) {
	atlasEnv, err := getAtlasEnv()
	if err != nil {
		return nil, err
	}
	optsRealm := []realm.ClientOpt{realm.SetUserAgent(terraformUserAgent)}
	authConfig := realmAuth.NewConfig(nil)
	token, err := authConfig.NewTokenFromCredentials(ctx, atlasEnv.PublicKey, atlasEnv.PrivateKey)
	if err != nil {
		return nil, err
	}

	clientRealm := realmAuth.NewClient(realmAuth.BasicTokenSource(token))
	realmClient, err := realm.New(clientRealm, optsRealm...)
	if err != nil {
		return nil, err
	}

	return realmClient, nil
}

func NewMongoDBClient() (atlasClient *atlas.Client, err error) {
	atlasEnv, err := getAtlasEnv()
	if err != nil {
		return nil, err
	}
	t := digest.NewTransport(atlasEnv.PublicKey, atlasEnv.PrivateKey)
	client, err := t.Client()

	opts := []atlas.ClientOpt{atlas.SetUserAgent(userAgent)}
	if baseURL := atlasEnv.BaseURL; baseURL != "" {
		opts = append(opts, atlas.SetBaseURL(baseURL))
	}

	mongodbClient, err := atlas.New(client, opts...)
	if err != nil {
		return nil, errors.New("unable to create Atlas client")
	}

	return mongodbClient, nil
}

func GetNewBasicClusterWithSampleData(ctx context.Context, client *atlas.Client, projectId string) (*atlas.AdvancedCluster, error) {
	// TODO
	return nil, nil
}

func getAtlasEnv() (atlasEnvOpts *AtlasEnvOptions, err error) {
	orgId1 := os.Getenv("ATLAS_ORG_ID")
	fmt.Println(orgId1)
	orgId, OrgIdOk := os.LookupEnv("ATLAS_ORG_ID")
	publicKey, publicKeyOk := os.LookupEnv("ATLAS_PUBLIC_KEY")
	privateKey, privateKeyOk := os.LookupEnv("ATLAS_PRIVATE_KEY")
	baseUrl, baseUrlOk := os.LookupEnv("ATLAS_BASE_URL")

	if !OrgIdOk || !privateKeyOk || !publicKeyOk || !baseUrlOk {
		return nil, errors.New("please ensure following env variables are set: " +
			"ATLAS_ORG_ID, ATLAS_PUBLIC_KEY, ATLAS_PRIVATE_KEY, ATLAS_BASE_URL, ATLAS_SECRET_PROFILE")
	}

	return &AtlasEnvOptions{orgId, privateKey, publicKey, baseUrl}, nil
}

func getExistingOrgUser(ctx context.Context, client *atlas.Client, orgId string) (*atlas.AtlasUser, error) {
	usersResponse, _, err := client.Organizations.Users(ctx, orgId, &atlas.ListOptions{})
	if err != nil {
		return nil, err
	}
	return &usersResponse.Results[0], nil
}
