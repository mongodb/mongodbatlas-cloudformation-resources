// Copyright 2026 MongoDB Inc
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
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/test/e2e/utility"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

type skipWaitTestContext struct {
	cfnClient              *cloudformation.Client
	atlasClient20231115014 *admin20231115014.APIClient
	resourceCtx            utility.ResourceContext
	template               string
	clusterTmplObj         skipWaitTestCluster
}

type skipWaitTestCluster struct {
	ResourceTypeName     string
	Name                 string
	Profile              string
	ProjectID            string
	DefaultReadConcern   string
	NodeCount            int
	SkipWaitForIdle      bool
	WithAdvancedSettings bool
}

const (
	skipWaitCfnTemplatePath = "cluster_skip_wait.json.template"
	skipWaitIdleStateName   = "IDLE"
	skipWaitPollInterval    = 30 * time.Second
	skipWaitPollAttempts    = 60
	// cloud-dev returns transient rejections, so reads are retried rather than failing the run.
	skipWaitReadAttempts      = 5
	skipWaitReadRetryInterval = 5 * time.Second
)

var (
	skipWaitProfile     = os.Getenv("MONGODB_ATLAS_SECRET_PROFILE")
	skipWaitOrgID       = os.Getenv("MONGODB_ATLAS_ORG_ID")
	skipWaitRandSuffix  = utility.GetRandNum().String()
	skipWaitProjectName = "cfn-e2e-cluster-skip-wait" + skipWaitRandSuffix
	skipWaitClusterName = "cfn-e2e-cluster-skip-wait" + skipWaitRandSuffix
	skipWaitStackName   = "stack-cluster-skip-wait-e2e-" + skipWaitRandSuffix
)

func TestClusterSkipWaitForIdleCFN(t *testing.T) {
	testCtx := setupSkipWaitSuite(t)

	t.Run("Validate Template", func(t *testing.T) {
		utility.TestIsTemplateValid(t, testCtx.cfnClient, testCtx.template)
	})

	t.Run("Create Stack Without Waiting For IDLE", func(t *testing.T) {
		testCreateSkipWaitStack(t, testCtx)
	})

	t.Run("Update Stack Without Waiting For IDLE", func(t *testing.T) {
		testUpdateSkipWaitStack(t, testCtx)
	})

	t.Run("Update Stack With The Default Wait", func(t *testing.T) {
		testUpdateWithTheDefaultWait(t, testCtx)
	})

	t.Run("Delete Stack", func(t *testing.T) {
		testDeleteSkipWaitStack(t, testCtx)
	})
}

func setupSkipWaitSuite(t *testing.T) *skipWaitTestContext {
	t.Helper()
	t.Log("Setting up skip wait suite")
	testCtx := new(skipWaitTestContext)
	testCtx.setUp(t)
	return testCtx
}

func (c *skipWaitTestContext) setUp(t *testing.T) {
	t.Helper()
	c.resourceCtx = utility.InitResourceCtx(skipWaitStackName, skipWaitRandSuffix, resourceTypeName, resourceDirectory)
	c.cfnClient, _ = utility.NewClients(t)
	_, c.atlasClient20231115014 = utility.NewClients20231115014(t)

	utility.PublishToPrivateRegistry(t, c.resourceCtx)
	c.setupPrerequisites(t)
}

func (c *skipWaitTestContext) setupPrerequisites(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		// The stack has to go first, so the cluster is gone before the project is deleted.
		cleanupSkipWaitResources(t, c)
		cleanupSkipWaitPrerequisites(t, c)
	})
	t.Log("Setting up prerequisites for skip wait test")

	var projectID string
	if projectIDEnvVar := os.Getenv("MONGODB_ATLAS_PROJECT_ID"); projectIDEnvVar != "" {
		t.Logf("using projectID from env var %s", projectIDEnvVar)
		projectID = projectIDEnvVar
	} else {
		projectID = utility.CreateProject(t, c.atlasClient20231115014, skipWaitOrgID, skipWaitProjectName)
	}

	// The create and update steps use different read concerns, and cloud-dev accepts only available
	// and local, so the update assertion proves the fast path applied the change.
	c.clusterTmplObj = skipWaitTestCluster{
		Name:                 skipWaitClusterName,
		ProjectID:            projectID,
		Profile:              skipWaitProfile,
		DefaultReadConcern:   "local",
		NodeCount:            3,
		SkipWaitForIdle:      true,
		WithAdvancedSettings: true,
		ResourceTypeName:     os.Getenv("RESOURCE_TYPE_NAME_FOR_E2E"),
	}

	var err error
	c.template, err = newSkipWaitCFNTemplate(c.clusterTmplObj)
	utility.FailNowIfError(t, "Error while reading skip wait CFN Template: %v", err)
	t.Logf("Skip wait test setup complete. ProjectID: %s, ClusterName: %s", c.clusterTmplObj.ProjectID, c.clusterTmplObj.Name)
}

func newSkipWaitCFNTemplate(tmpl skipWaitTestCluster) (string, error) {
	return utility.ExecuteGoTemplate(skipWaitCfnTemplatePath, tmpl)
}

func testCreateSkipWaitStack(t *testing.T, c *skipWaitTestContext) {
	t.Helper()
	t.Logf("Creating skip wait stack with template:\n%s", c.template)

	output := utility.CreateStack(t, c.cfnClient, skipWaitStackName, c.template)

	// The stack completes as soon as Atlas accepts the create, so the cluster is still converging.
	cluster := readSkipWaitClusterFromAtlas(t, c)
	assert.NotEqual(t, skipWaitIdleStateName, cluster.GetStateName())

	// The fast path reads the cluster once, so Id is populated without waiting for IDLE.
	assert.NotEmpty(t, getSkipWaitClusterIDFromStack(output))

	// AdvancedSettings are sent right after the create call and survive convergence.
	waitForSkipWaitClusterIdle(t, c)
	assert.Equal(t, c.clusterTmplObj.DefaultReadConcern, readSkipWaitProcessArgs(t, c).GetDefaultReadConcern())
}

func testUpdateSkipWaitStack(t *testing.T, c *skipWaitTestContext) {
	t.Helper()

	c.clusterTmplObj.DefaultReadConcern = "available"
	c.clusterTmplObj.NodeCount = 5

	var err error
	c.template, err = newSkipWaitCFNTemplate(c.clusterTmplObj)
	utility.FailNowIfError(t, "Error while reading skip wait CFN Template: %v", err)

	utility.UpdateStack(t, c.cfnClient, skipWaitStackName, c.template)

	// The node count change is still in flight when the stack update completes.
	cluster := readSkipWaitClusterFromAtlas(t, c)
	assert.NotEqual(t, skipWaitIdleStateName, cluster.GetStateName())

	// AdvancedSettings are sent before the change and survive convergence.
	waitForSkipWaitClusterIdle(t, c)
	assert.Equal(t, c.clusterTmplObj.DefaultReadConcern, readSkipWaitProcessArgs(t, c).GetDefaultReadConcern())
	assert.Equal(t, c.clusterTmplObj.NodeCount, readSkipWaitNodeCount(t, c))
}

func testUpdateWithTheDefaultWait(t *testing.T, c *skipWaitTestContext) {
	t.Helper()

	// Removing the property reverts to the default behavior, so the handler waits for the change to converge.
	// AdvancedSettings are dropped too, because the default path applies them after convergence and that
	// call starts a second Atlas update, which would leave the cluster UPDATING at this point.
	c.clusterTmplObj.SkipWaitForIdle = false
	c.clusterTmplObj.NodeCount = 7
	c.clusterTmplObj.WithAdvancedSettings = false

	var err error
	c.template, err = newSkipWaitCFNTemplate(c.clusterTmplObj)
	utility.FailNowIfError(t, "Error while reading skip wait CFN Template: %v", err)

	utility.UpdateStack(t, c.cfnClient, skipWaitStackName, c.template)

	// The fast path above returns while Atlas is still converging. This one only completes once the
	// change has converged, so the cluster is already IDLE when the stack update returns.
	cluster := readSkipWaitClusterFromAtlas(t, c)
	assert.Equal(t, skipWaitIdleStateName, cluster.GetStateName())
	assert.Equal(t, c.clusterTmplObj.NodeCount, readSkipWaitNodeCount(t, c))
}

func testDeleteSkipWaitStack(t *testing.T, c *skipWaitTestContext) {
	t.Helper()
	utility.DeleteStack(t, c.cfnClient, skipWaitStackName)
	_, resp, _ := c.atlasClient20231115014.ClustersApi.GetCluster(ctx.Background(), c.clusterTmplObj.ProjectID, c.clusterTmplObj.Name).Execute()
	assert.Equal(t, 404, resp.StatusCode)
}

func readSkipWaitClusterFromAtlas(t *testing.T, c *skipWaitTestContext) *admin20231115014.AdvancedClusterDescription {
	t.Helper()
	var lastErr error
	for attempt := range skipWaitReadAttempts {
		if attempt > 0 {
			time.Sleep(skipWaitReadRetryInterval)
		}
		cluster, resp, err := c.atlasClient20231115014.ClustersApi.GetCluster(ctx.Background(), c.clusterTmplObj.ProjectID, c.clusterTmplObj.Name).Execute()
		if err == nil {
			assert.Equal(t, 200, resp.StatusCode)
			return cluster
		}
		lastErr = err
		t.Logf("retrying read of cluster %s (attempt %d): %v", c.clusterTmplObj.Name, attempt+1, err)
	}
	t.Fatalf("Err while retrieving Cluster from Atlas after %d attempts: %v", skipWaitReadAttempts, lastErr)
	return nil
}

func readSkipWaitProcessArgs(t *testing.T, c *skipWaitTestContext) *admin20231115014.ClusterDescriptionProcessArgs {
	t.Helper()
	processArgs, resp, err := c.atlasClient20231115014.ClustersApi.GetClusterAdvancedConfiguration(ctx.Background(), c.clusterTmplObj.ProjectID, c.clusterTmplObj.Name).Execute()
	utility.FailNowIfError(t, "Err while retrieving AdvancedSettings from Atlas: %v", err)
	assert.Equal(t, 200, resp.StatusCode)
	return processArgs
}

func readSkipWaitNodeCount(t *testing.T, c *skipWaitTestContext) int {
	t.Helper()
	replicationSpecs := readSkipWaitClusterFromAtlas(t, c).GetReplicationSpecs()
	require.NotEmpty(t, replicationSpecs)
	regionConfigs := replicationSpecs[0].GetRegionConfigs()
	require.NotEmpty(t, regionConfigs)
	electableSpecs := regionConfigs[0].GetElectableSpecs()
	return electableSpecs.GetNodeCount()
}

func waitForSkipWaitClusterIdle(t *testing.T, c *skipWaitTestContext) {
	t.Helper()
	for range skipWaitPollAttempts {
		time.Sleep(skipWaitPollInterval)
		if readSkipWaitClusterFromAtlas(t, c).GetStateName() == skipWaitIdleStateName {
			return
		}
	}
	t.Fatalf("Cluster %s did not reach %s", c.clusterTmplObj.Name, skipWaitIdleStateName)
}

func getSkipWaitClusterIDFromStack(output *cloudformation.DescribeStacksOutput) string {
	stackOutputs := output.Stacks[0].Outputs
	for i := range stackOutputs {
		if *stackOutputs[i].OutputKey == "MongoDBAtlasClusterID" {
			return *stackOutputs[i].OutputValue
		}
	}
	return ""
}

func cleanupSkipWaitResources(t *testing.T, c *skipWaitTestContext) {
	t.Helper()
	utility.DeleteStackForCleanup(t, c.cfnClient, skipWaitStackName)
}

func cleanupSkipWaitPrerequisites(t *testing.T, c *skipWaitTestContext) {
	t.Helper()
	t.Log("Cleaning up skip wait test prerequisites")
	if os.Getenv("MONGODB_ATLAS_PROJECT_ID") == "" {
		utility.DeleteProject(t, c.atlasClient20231115014, c.clusterTmplObj.ProjectID)
	} else {
		t.Log("skipping project deletion (project managed outside of test)")
	}
}
