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
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/test/e2e/utility"
	"github.com/stretchr/testify/assert"
	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

type pauseTestContext struct {
	cfnClient              *cloudformation.Client
	atlasClient20231115014 *admin20231115014.APIClient
	resourceCtx            utility.ResourceContext
	template               string
	clusterTmplObj         pauseTestCluster
}

type pauseTestCluster struct {
	ResourceTypeName string
	Name             string
	Profile          string
	ProjectID        string
	Paused           string
}

const (
	pauseCfnTemplatePath = "cluster_pause.json.template"
)

// Replication specs are hardcoded in the CFN template for this test.

var (
	pauseProfile     = os.Getenv("MONGODB_ATLAS_SECRET_PROFILE")
	pauseOrgID       = os.Getenv("MONGODB_ATLAS_ORG_ID")
	pauseRandSuffix  = utility.GetRandNum().String()
	pauseProjectName = "cfn-e2e-cluster-pause" + pauseRandSuffix
	pauseClusterName = "cfn-e2e-cluster-pause" + pauseRandSuffix
	pauseStackName   = "stack-cluster-pause-e2e-" + pauseRandSuffix
)

func TestClusterPauseCFN(t *testing.T) {
	t.Parallel()
	testCtx := setupPauseSuite(t)

	t.Run("Validate Template", func(t *testing.T) {
		utility.TestIsTemplateValid(t, testCtx.cfnClient, testCtx.template)
	})

	t.Run("Create Stack", func(t *testing.T) {
		testCreatePauseStack(t, testCtx)
	})

	t.Run("Pause Cluster", func(t *testing.T) {
		testUpdatePauseState(t, testCtx, true)
	})

	t.Run("Unpause Cluster", func(t *testing.T) {
		testUpdatePauseState(t, testCtx, false)
	})

	t.Run("Delete Stack", func(t *testing.T) {
		testDeletePauseStack(t, testCtx)
	})
}

func setupPauseSuite(t *testing.T) *pauseTestContext {
	t.Helper()
	t.Log("Setting up pause suite")
	testCtx := new(pauseTestContext)
	testCtx.setUp(t)
	return testCtx
}

func (c *pauseTestContext) setUp(t *testing.T) {
	t.Helper()
	c.resourceCtx = utility.InitResourceCtx(pauseStackName, pauseRandSuffix, resourceTypeName, resourceDirectory)
	c.cfnClient, _ = utility.NewClients(t)
	_, c.atlasClient20231115014 = utility.NewClients20231115014(t)

	utility.PublishToPrivateRegistry(t, c.resourceCtx)
	c.setupPrerequisites(t)
}

func (c *pauseTestContext) setupPrerequisites(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		cleanupPausePrerequisites(t, c)
		cleanupPauseResources(t, c)
	})
	t.Log("Setting up prerequisites for pause test")

	var projectID string
	if projectIDEnvVar := os.Getenv("MONGODB_ATLAS_PROJECT_ID"); projectIDEnvVar != "" {
		t.Logf("using projectID from env var %s", projectIDEnvVar)
		projectID = projectIDEnvVar
	} else {
		projectID = utility.CreateProject(t, c.atlasClient20231115014, pauseOrgID, pauseProjectName)
	}

	c.clusterTmplObj = pauseTestCluster{
		Name:             pauseClusterName,
		ProjectID:        projectID,
		Profile:          pauseProfile,
		Paused:           "false",
		ResourceTypeName: os.Getenv("RESOURCE_TYPE_NAME_FOR_E2E"),
	}

	var err error
	c.template, err = newPauseCFNTemplate(c.clusterTmplObj)
	utility.FailNowIfError(t, "Error while reading pause CFN Template: %v", err)
	t.Logf("Pause test setup complete. ProjectID: %s, ClusterName: %s", c.clusterTmplObj.ProjectID, c.clusterTmplObj.Name)
}

func newPauseCFNTemplate(tmpl pauseTestCluster) (string, error) {
	return utility.ExecuteGoTemplate(pauseCfnTemplatePath, tmpl)
}

func testCreatePauseStack(t *testing.T, c *pauseTestContext) {
	t.Helper()
	t.Logf("Creating pause stack with template:\n%s", c.template)

	output := utility.CreateStack(t, c.cfnClient, pauseStackName, c.template)
	clusterID := getPauseClusterIDFromStack(output)

	cluster := readPauseClusterFromAtlas(t, c)

	a := assert.New(t)
	a.Equal(cluster.GetId(), clusterID)
	a.False(cluster.GetPaused())
}

func testUpdatePauseState(t *testing.T, c *pauseTestContext, pause bool) {
	t.Helper()

	if pause {
		c.clusterTmplObj.Paused = "true"
	} else {
		c.clusterTmplObj.Paused = "false"
	}

	var err error
	c.template, err = newPauseCFNTemplate(c.clusterTmplObj)
	utility.FailNowIfError(t, "Error while reading pause CFN Template: %v", err)

	output := utility.UpdateStack(t, c.cfnClient, pauseStackName, c.template)
	_ = getPauseClusterIDFromStack(output)

	cluster := readPauseClusterFromAtlas(t, c)

	a := assert.New(t)
	if pause {
		a.True(cluster.GetPaused())
	} else {
		a.False(cluster.GetPaused())
	}
}

func testDeletePauseStack(t *testing.T, c *pauseTestContext) {
	t.Helper()
	utility.DeleteStack(t, c.cfnClient, pauseStackName)
	a := assert.New(t)
	_, resp, _ := c.atlasClient20231115014.ClustersApi.GetCluster(ctx.Background(), c.clusterTmplObj.ProjectID, c.clusterTmplObj.Name).Execute()
	a.Equal(404, resp.StatusCode)
}

func cleanupPauseResources(t *testing.T, c *pauseTestContext) {
	t.Helper()
	utility.DeleteStackForCleanup(t, c.cfnClient, pauseStackName)
}

func cleanupPausePrerequisites(t *testing.T, c *pauseTestContext) {
	t.Helper()
	t.Log("Cleaning up pause test prerequisites")
	if os.Getenv("MONGODB_ATLAS_PROJECT_ID") == "" {
		utility.DeleteProject(t, c.atlasClient20231115014, c.clusterTmplObj.ProjectID)
	} else {
		t.Log("skipping project deletion (project managed outside of test)")
	}
}

func readPauseClusterFromAtlas(t *testing.T, c *pauseTestContext) *admin20231115014.AdvancedClusterDescription {
	t.Helper()
	context := ctx.Background()
	projectID := c.clusterTmplObj.ProjectID
	cluster, resp, err := c.atlasClient20231115014.ClustersApi.GetCluster(context, projectID, c.clusterTmplObj.Name).Execute()
	utility.FailNowIfError(t, "Err while retrieving Cluster from Atlas: %v", err)
	assert.Equal(t, 200, resp.StatusCode)
	return cluster
}

func getPauseClusterIDFromStack(output *cloudformation.DescribeStacksOutput) string {
	stackOutputs := output.Stacks[0].Outputs
	for i := 0; i < len(stackOutputs); i++ {
		if *stackOutputs[i].OutputKey == "MongoDBAtlasClusterID" {
			return *stackOutputs[i].OutputValue
		}
	}
	return ""
}
