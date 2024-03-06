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
	"testing"
	"time"

	"github.com/mongodb-forks/digest"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
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
		Usernames: &[]string{orgUser.Username},
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
	sdkV2Client, err := c.NewSDKV2LatestClient(client)
	if err != nil {
		return nil, fmt.Errorf("unable to create Atlas client")
	}

	return sdkV2Client, nil
}

func CreateProject(t *testing.T, atlasClient *admin.APIClient, orgID, projectName string) string {
	t.Helper()
	projectParams := &admin.Group{
		Name:  projectName,
		OrgId: orgID,
	}
	resp, _, err := atlasClient.ProjectsApi.CreateProject(context.Background(), projectParams).Execute()
	FailNowIfError(t, "Error creating Atlas Project: %v", err)
	return resp.GetId()
}

func DeleteProject(t *testing.T, atlasClient *admin.APIClient, projectID string) {
	t.Helper()
	_, _, err := atlasClient.ProjectsApi.DeleteProject(context.Background(), projectID).Execute()
	if err != nil {
		t.Logf("Atlas Project could not be deleted during cleanup: %s\n", err.Error())
	}
}

func CreateCluster(t *testing.T, atlasClient *admin.APIClient, projectID, clusterName string) {
	t.Helper()
	clusterParams := &admin.AdvancedClusterDescription{
		Name:        &clusterName,
		ClusterType: admin.PtrString("REPLICASET"),
		ReplicationSpecs: &[]admin.ReplicationSpec{
			{
				NumShards: admin.PtrInt(1),
				RegionConfigs: &[]admin.CloudRegionConfig{
					{
						ProviderName: admin.PtrString("AWS"),
						RegionName:   admin.PtrString("US_EAST_1"),
						Priority:     admin.PtrInt(7),
						ElectableSpecs: &admin.HardwareSpec{
							InstanceSize: admin.PtrString("M10"),
							NodeCount:    admin.PtrInt(3),
						},
					},
				},
			},
		},
	}
	_, _, err := atlasClient.ClustersApi.CreateCluster(context.Background(), projectID, clusterParams).Execute()
	FailNowIfError(t, "Error creating Atlas Cluster: %v", err)
	waitCluster(t, atlasClient, projectID, clusterName)
}

func DeleteCluster(t *testing.T, atlasClient *admin.APIClient, projectID, clusterName string) {
	t.Helper()
	_, err := atlasClient.ClustersApi.DeleteCluster(context.Background(), projectID, clusterName).Execute()
	if err != nil {
		t.Logf("Atlas Cluster could not be deleted during cleanup: %s\n", err.Error())
		return
	}
	waitCluster(t, atlasClient, projectID, clusterName)
}

func waitCluster(t *testing.T, atlasClient *admin.APIClient, projectID, clusterName string) {
	t.Helper()
	for {
		time.Sleep(time.Second * 30)
		read, httpResp, err := atlasClient.ClustersApi.GetCluster(context.Background(), projectID, clusterName).Execute()
		if read.GetStateName() == "IDLE" || httpResp.StatusCode == 404 {
			return
		}
		if err != nil {
			FailNowIfError(t, "Error watching Atlas Cluster: %v", err)
			return
		}
	}
}

func getAtlasEnv() (atlasEnvOpts *AtlasEnvOptions, err error) {
	orgID := os.Getenv("MONGODB_ATLAS_ORG_ID")
	publicKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	privateKey := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")
	baseURL := os.Getenv("MONGODB_ATLAS_BASE_URL")

	if orgID == "" || publicKey == "" || privateKey == "" || baseURL == "" {
		return nil, fmt.Errorf("please ensure following env variables are set: " +
			"MONGODB_ATLAS_ORG_ID, MONGODB_ATLAS_PUBLIC_KEY, MONGODB_ATLAS_PRIVATE_KEY, MONGODB_ATLAS_BASE_URL, MONGODB_ATLAS_SECRET_PROFILE")
	}

	return &AtlasEnvOptions{orgID, privateKey, publicKey, baseURL}, nil
}

func getExistingOrgUser(ctx context.Context, client *admin.APIClient, orgID string) (*admin.CloudAppUser, error) {
	usersResponse, _, err := client.OrganizationsApi.ListOrganizationUsers(ctx, orgID).Execute()
	if err != nil {
		return nil, err
	}
	return &usersResponse.GetResults()[0], nil
}
