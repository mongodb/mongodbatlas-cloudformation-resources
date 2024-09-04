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
package searchdeployment_test

import (
	ctx "context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/test/e2e/utility"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20231115014/admin"
)

type localTestContext struct {
	cfnClient    *cloudformation.Client
	atlasClient  *admin.APIClient
	resourceCtx  utility.ResourceContext
	template     string
	templateData testData
}

type testData struct {
	ResourceTypeName string
	Profile          string
	ProjectID        string
	ClusterName      string
	InstanceSize     string
	NodeCount        int
}

const (
	resourceTypeName    = "MongoDB::Atlas::SearchDeployment"
	resourceDirectory   = "search-deployment"
	cfnTemplatePath     = "searchdeployment.json.template"
	instanceSize        = "S20_HIGHCPU_NVME"
	nodeCount           = 3
	instanceSizeUpdated = "S30_HIGHCPU_NVME"
	nodeCountUpdated    = 2
)

var (
	profile       = os.Getenv("MONGODB_ATLAS_SECRET_PROFILE")
	orgID         = os.Getenv("MONGODB_ATLAS_ORG_ID")
	e2eRandSuffix = utility.GetRandNum().String()
	stackName     = "stack-searchdeployment-e2e-" + e2eRandSuffix
	clusterName   = "cfn-e2e-searchdeployment" + e2eRandSuffix
	projectName   = "cfn-e2e-searchdeployment" + e2eRandSuffix
)

func TestSearchDeploymentCFN(t *testing.T) {
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
	c.cfnClient, c.atlasClient = utility.NewClients(t)
	utility.PublishToPrivateRegistry(t, c.resourceCtx)
	c.setupPrerequisites(t)
}

func testCreateStack(t *testing.T, c *localTestContext) {
	t.Helper()
	utility.CreateStack(t, c.cfnClient, stackName, c.template)
	resp, httpResp, err := c.atlasClient.AtlasSearchApi.GetAtlasSearchDeployment(ctx.Background(), c.templateData.ProjectID, c.templateData.ClusterName).Execute()
	utility.FailNowIfError(t, "Error while retrieving Search Deployment from Atlas: %v", err)
	assert.Equal(t, 200, httpResp.StatusCode)
	assert.Equal(t, instanceSize, resp.GetSpecs()[0].GetInstanceSize())
	assert.Equal(t, nodeCount, resp.GetSpecs()[0].GetNodeCount())
}

func testUpdateStack(t *testing.T, c *localTestContext) {
	t.Helper()
	c.templateData.InstanceSize = instanceSizeUpdated
	c.templateData.NodeCount = nodeCountUpdated
	var err error
	c.template, err = newCFNTemplate(c.templateData)
	utility.FailNowIfError(t, "Error while reading updated CFN Template: %v", err)
	utility.UpdateStack(t, c.cfnClient, stackName, c.template)
	resp, httpResp, err := c.atlasClient.AtlasSearchApi.GetAtlasSearchDeployment(ctx.Background(), c.templateData.ProjectID, c.templateData.ClusterName).Execute()
	utility.FailNowIfError(t, "Error while retrieving Search Deployment from Atlas: %v", err)
	assert.Equal(t, 200, httpResp.StatusCode)
	assert.Equal(t, instanceSizeUpdated, resp.GetSpecs()[0].GetInstanceSize())
	assert.Equal(t, nodeCountUpdated, resp.GetSpecs()[0].GetNodeCount())
}

func testDeleteStack(t *testing.T, c *localTestContext) {
	t.Helper()
	utility.DeleteStack(t, c.cfnClient, stackName)
	resp, _, _ := c.atlasClient.AtlasSearchApi.GetAtlasSearchDeployment(ctx.Background(), c.templateData.ProjectID, c.templateData.ClusterName).Execute()
	assert.Nil(t, resp)
}

func cleanupResources(t *testing.T, c *localTestContext) {
	t.Helper()
	utility.DeleteStackForCleanup(t, c.cfnClient, stackName)
	t.Log("Cleaning up resources")
	utility.DeleteCluster(t, c.atlasClient, c.templateData.ProjectID, c.templateData.ClusterName)
	utility.DeleteProject(t, c.atlasClient, c.templateData.ProjectID)
}

func (c *localTestContext) setupPrerequisites(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		cleanupResources(t, c)
	})
	t.Log("Setting up prerequisites")

	projectID := utility.CreateProject(t, c.atlasClient, orgID, projectName)
	utility.CreateCluster(t, c.atlasClient, projectID, clusterName)

	c.templateData = testData{
		ResourceTypeName: os.Getenv("RESOURCE_TYPE_NAME_FOR_E2E"),
		Profile:          profile,
		ProjectID:        projectID,
		ClusterName:      clusterName,
		InstanceSize:     instanceSize,
		NodeCount:        nodeCount,
	}

	var err error
	c.template, err = newCFNTemplate(c.templateData)
	utility.FailNowIfError(t, "Error while reading CFN Template: %v", err)
}

func newCFNTemplate(data testData) (string, error) {
	return utility.ExecuteGoTemplate(cfnTemplatePath, data)
}
