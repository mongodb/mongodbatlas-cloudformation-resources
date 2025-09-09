// Copyright 2024 MongoDB Inc
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
package cluster_test

import (
	ctx "context"
	"net/http"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/cluster/cmd/resource"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/test/e2e/utility"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/stretchr/testify/assert"
	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

type localTestContext struct {
	err                 error
	cfnClient           *cloudformation.Client
	atlasClient         *admin20231115014.APIClient
	resourceCtx         utility.ResourceContext
	template            string
	replicationIDCreate string
	clusterTmplObj      testCluster
}

type testCluster struct {
	ResourceTypeName string
	Name             string
	FlexName         string
	Profile          string
	ProjectID        string
	ReplicationSpecs []resource.AdvancedReplicationSpec
	NodeCount        int
}

const (
	resourceTypeName  = "MongoDB::Atlas::Cluster"
	resourceDirectory = "cluster"
	cfnTemplatePath   = "cluster.json.template"
)

func replicationSpec(nodeCount int, region, zoneName string) resource.AdvancedReplicationSpec {
	return resource.AdvancedReplicationSpec{
		NumShards: util.IntPtr(1),
		ZoneName:  &zoneName,
		AdvancedRegionConfigs: []resource.AdvancedRegionConfig{{
			RegionName:   &region,
			Priority:     util.IntPtr(7),
			ProviderName: util.StringPtr("AWS"),
			ElectableSpecs: &resource.Specs{
				EbsVolumeType: util.StringPtr("STANDARD"),
				InstanceSize:  util.StringPtr("M10"),
				NodeCount:     &nodeCount,
			},
		}},
	}
}

var (
	profile                    = os.Getenv("MONGODB_ATLAS_SECRET_PROFILE")
	orgID                      = os.Getenv("MONGODB_ATLAS_ORG_ID")
	e2eRandSuffix              = utility.GetRandNum().String()
	testProjectName            = "cfn-e2e-cluster" + e2eRandSuffix
	testClusterName            = "cfn-e2e-cluster" + e2eRandSuffix
	testFlexClusterName        = "cfn-e2e-cluster-flex" + e2eRandSuffix
	stackName                  = "stack-cluster-e2e-" + e2eRandSuffix
	nodeCountCreate            = 3
	nodeCountUpdate            = 5
	zone1                      = "zone1"
	zone2                      = "zone2"
	replicationSpecZone1Create = replicationSpec(nodeCountCreate, "US_EAST_1", zone1)
	replicationSpecZone1Update = replicationSpec(nodeCountUpdate, "US_EAST_1", zone1)
	replicationSpecZone2       = replicationSpec(nodeCountUpdate, "EU_WEST_1", zone2)
	replicationSpecsCreate     = []resource.AdvancedReplicationSpec{replicationSpecZone1Create}
	replicationSpecsUpdate     = []resource.AdvancedReplicationSpec{replicationSpecZone1Update, replicationSpecZone2}
)

func TestClusterCFN(t *testing.T) {
	testCtx := setupSuite(t)

	t.Run("Validate Template", func(t *testing.T) {
		utility.TestIsTemplateValid(t, testCtx.cfnClient, testCtx.template)
	})

	t.Run("Create Stack", func(t *testing.T) {
		testCreateStack(t, testCtx)
	})

	t.Run("Update Stack", func(t *testing.T) {
		testUpdateStack(t, testCtx)
	})

	t.Run("Delete Stack", func(t *testing.T) {
		testDeleteStack(t, testCtx)
	})
}

func setupSuite(t *testing.T) *localTestContext {
	t.Helper()
	t.Log("Setting up suite")
	testCtx := new(localTestContext)
	testCtx.setUp(t)

	return testCtx
}

func (c *localTestContext) setUp(t *testing.T) {
	t.Helper()
	c.resourceCtx = utility.InitResourceCtx(stackName, e2eRandSuffix, resourceTypeName, resourceDirectory)
	c.cfnClient, c.atlasClient = utility.NewClients20231115014(t)
	utility.PublishToPrivateRegistry(t, c.resourceCtx)
	c.setupPrerequisites(t)
}

func testCreateStack(t *testing.T, c *localTestContext) {
	t.Helper()

	output := utility.CreateStack(t, c.cfnClient, stackName, c.template)
	clusterID := getClusterIDFromStack(output)
	flexClusterID := getFlexClusterIDFromStack(output)

	project, cluster, flexCluster := readFromAtlas(t, c)

	a := assert.New(t)
	a.Equal(int64(2), project.ClusterCount)
	a.Equal(cluster.GetId(), clusterID)
	a.Equal(flexCluster.GetId(), flexClusterID)

	replicationSpecs := cluster.GetReplicationSpecs()
	checkReplicationSpecs(a, replicationSpecs, nodeCountCreate, 1)
	a.NotEmpty(replicationSpecs[0].GetId())
	c.replicationIDCreate = replicationSpecs[0].GetId()

	flexReplicationSpecs := flexCluster.GetReplicationSpecs()
	a.Len(flexReplicationSpecs, 1)
	a.Len(flexReplicationSpecs[0].GetRegionConfigs(), 1)
	flexRegionConfig := flexReplicationSpecs[0].GetRegionConfigs()[0]
	a.Equal("FLEX", flexRegionConfig.GetProviderName())
	a.Equal("US_EAST_1", flexRegionConfig.GetRegionName())
}

func testUpdateStack(t *testing.T, c *localTestContext) {
	t.Helper()

	c.clusterTmplObj.ReplicationSpecs = replicationSpecsUpdate
	c.template, c.err = newCFNTemplate(c.clusterTmplObj)

	output := utility.UpdateStack(t, c.cfnClient, stackName, c.template)
	clusterID := getClusterIDFromStack(output)
	flexClusterID := getFlexClusterIDFromStack(output)

	project, cluster, flexCluster := readFromAtlas(t, c)

	a := assert.New(t)
	a.Equal(int64(2), project.ClusterCount)
	a.Equal(cluster.GetId(), clusterID)
	a.Equal(flexCluster.GetId(), flexClusterID)

	replicationSpecs := cluster.GetReplicationSpecs()
	checkReplicationSpecs(a, replicationSpecs, nodeCountUpdate, 2)
	a.NotEmpty(replicationSpecs[0].GetId())
	a.Equal(replicationSpecs[0].GetId(), c.replicationIDCreate)
}

func testDeleteStack(t *testing.T, c *localTestContext) {
	t.Helper()
	utility.DeleteStack(t, c.cfnClient, stackName)
	a := assert.New(t)
	_, resp, _ := c.atlasClient.ClustersApi.GetCluster(ctx.Background(), c.clusterTmplObj.ProjectID, c.clusterTmplObj.Name).Execute()
	a.Equal(404, resp.StatusCode)
	_, flexResp, _ := c.atlasClient.ClustersApi.GetCluster(ctx.Background(), c.clusterTmplObj.ProjectID, c.clusterTmplObj.FlexName).Execute()
	a.Equal(404, flexResp.StatusCode)
}

func cleanupResources(t *testing.T, c *localTestContext) {
	t.Helper()
	utility.DeleteStackForCleanup(t, c.cfnClient, stackName)
}

func cleanupPrerequisites(t *testing.T, c *localTestContext) {
	t.Helper()
	t.Log("Cleaning up prerequisites")
	if os.Getenv("MONGODB_ATLAS_PROJECT_ID") == "" {
		utility.DeleteProject(t, c.atlasClient, c.clusterTmplObj.ProjectID)
	} else {
		t.Log("skipping project deletion (project managed outside of test)")
	}
}

func (c *localTestContext) setupPrerequisites(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		cleanupPrerequisites(t, c)
		cleanupResources(t, c)
	})
	t.Log("Setting up prerequisites")
	var projectID string
	if projectIDEnvVar := os.Getenv("MONGODB_ATLAS_PROJECT_ID"); projectIDEnvVar != "" {
		t.Logf("using projectID from env var %s", projectIDEnvVar)
		projectID = projectIDEnvVar
	} else {
		projectID = utility.CreateProject(t, c.atlasClient, orgID, testProjectName)
	}

	c.clusterTmplObj = testCluster{
		Name:             testClusterName,
		FlexName:         testFlexClusterName,
		ProjectID:        projectID,
		Profile:          profile,
		NodeCount:        nodeCountCreate,
		ResourceTypeName: os.Getenv("RESOURCE_TYPE_NAME_FOR_E2E"),
		ReplicationSpecs: replicationSpecsCreate,
	}

	c.template, c.err = newCFNTemplate(c.clusterTmplObj)
	utility.FailNowIfError(t, "Error while reading CFN Template: %v", c.err)
}

func newCFNTemplate(tmpl testCluster) (string, error) {
	return utility.ExecuteGoTemplate(cfnTemplatePath, tmpl)
}

func checkReplicationSpecs(a *assert.Assertions, replicationSpecs []admin20231115014.ReplicationSpec, nodeCount, length int) {
	a.Len(replicationSpecs, length)
	for i, spec := range replicationSpecs {
		for _, config := range spec.GetRegionConfigs() {
			hwSpec := config.GetElectableSpecs()
			a.Equal(nodeCount, hwSpec.GetNodeCount())
		}
		switch i {
		case 0:
			a.Equal(zone1, spec.GetZoneName())
		case 1:
			a.Equal(zone2, spec.GetZoneName())
		}
	}
}

func readFromAtlas(t *testing.T, c *localTestContext) (project *admin20231115014.Group, cluster *admin20231115014.AdvancedClusterDescription, flexCluster *admin20231115014.AdvancedClusterDescription) {
	t.Helper()
	context := ctx.Background()
	projectID := c.clusterTmplObj.ProjectID
	var err error
	var getProjectResponse, getClusterResponse, getFlexClusterResponse *http.Response
	project, getProjectResponse, err = c.atlasClient.ProjectsApi.GetProject(context, projectID).Execute()
	utility.FailNowIfError(t, "Error while retrieving Project from Atlas: %v", err)
	cluster, getClusterResponse, err = c.atlasClient.ClustersApi.GetCluster(context, projectID, c.clusterTmplObj.Name).Execute()
	utility.FailNowIfError(t, "Err while retrieving Cluster from Atlas: %v", err)
	flexCluster, getFlexClusterResponse, err = c.atlasClient.ClustersApi.GetCluster(context, projectID, c.clusterTmplObj.FlexName).Execute()
	utility.FailNowIfError(t, "Err while retrieving Flex Cluster from Atlas: %v", err)
	assert.Equal(t, 200, getProjectResponse.StatusCode)
	assert.Equal(t, 200, getClusterResponse.StatusCode)
	assert.Equal(t, 200, getFlexClusterResponse.StatusCode)
	return
}

func getClusterIDFromStack(output *cloudformation.DescribeStacksOutput) string {
	stackOutputs := output.Stacks[0].Outputs
	for i := 0; i < len(stackOutputs); i++ {
		if *aws.String(*stackOutputs[i].OutputKey) == "MongoDBAtlasClusterID" {
			return *aws.String(*stackOutputs[i].OutputValue)
		}
	}
	return ""
}

func getFlexClusterIDFromStack(output *cloudformation.DescribeStacksOutput) string {
	stackOutputs := output.Stacks[0].Outputs
	for i := 0; i < len(stackOutputs); i++ {
		if *aws.String(*stackOutputs[i].OutputKey) == "MongoDBAtlasFlexClusterID" {
			return *aws.String(*stackOutputs[i].OutputValue)
		}
	}
	return ""
}
