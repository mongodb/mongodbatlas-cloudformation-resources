// Copyright 2025 MongoDB Inc
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
package flexcluster_test

import (
	ctx "context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/test/e2e/utility"
	"github.com/stretchr/testify/assert"
	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"
	"go.mongodb.org/atlas-sdk/v20250312012/admin"
)

type localTestContext struct {
	cfnClient              *cloudformation.Client
	atlasClient            *admin.APIClient
	atlasClient20231115014 *admin20231115014.APIClient
	resourceCtx            utility.ResourceContext
	template               string
	templateData           testData
}

type testData struct {
	ResourceTypeName string
	Profile          string
	ProjectID        string
	ClusterName      string
	InstanceSize     string
	Tags             []Tag
	NodeCount        int
}

type Tag struct {
	Key   string
	Value string
}

const (
	resourceTypeName  = "MongoDB::Atlas::FlexCluster"
	resourceDirectory = "flex-cluster"
	cfnTemplatePath   = "flexcluster.json.template"
)

var (
	profile       = os.Getenv("MONGODB_ATLAS_SECRET_PROFILE")
	orgID         = os.Getenv("MONGODB_ATLAS_ORG_ID")
	e2eRandSuffix = utility.GetRandNum().String()
	stackName     = "stack-flexcluster-e2e-" + e2eRandSuffix
	clusterName   = "cfn-test-bot-" + e2eRandSuffix
	projectName   = "cfn-test-bot-" + e2eRandSuffix
)

func TestFlexClusterCFN(t *testing.T) {
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
	_, c.atlasClient20231115014 = utility.NewClients20231115014(t)
	utility.PublishToPrivateRegistry(t, c.resourceCtx)
	c.setupPrerequisites(t)
}

func testCreateStack(t *testing.T, c *localTestContext) {
	t.Helper()
	utility.CreateStack(t, c.cfnClient, stackName, c.template)
	resp, httpResp, err := c.atlasClient.FlexClustersApi.GetFlexCluster(ctx.Background(), c.templateData.ProjectID, c.templateData.ClusterName).Execute()
	utility.FailNowIfError(t, "Error while retrieving Flex Cluster from Atlas: %v", err)
	assert.Equal(t, 200, httpResp.StatusCode)
	assert.NotNil(t, resp)
	assert.True(t, resp.Tags == nil || len(*resp.Tags) == 0)
}

func testUpdateStack(t *testing.T, c *localTestContext) {
	t.Helper()
	c.templateData.Tags = []Tag{
		{Key: "environment", Value: "development"},
		{Key: "team", Value: "platform"}}
	var err error
	c.template, err = newCFNTemplate(c.templateData)
	utility.FailNowIfError(t, "Error while reading updated CFN Template: %v", err)
	utility.UpdateStack(t, c.cfnClient, stackName, c.template)
	resp, httpResp, err := c.atlasClient.FlexClustersApi.GetFlexCluster(ctx.Background(), c.templateData.ProjectID, c.templateData.ClusterName).Execute()
	utility.FailNowIfError(t, "Error while retrieving Flex Cluster from Atlas: %v", err)
	assert.Equal(t, 200, httpResp.StatusCode)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Tags)
	assert.Len(t, *resp.Tags, 2)
	assert.Contains(t, *resp.Tags, admin.ResourceTag{Key: "environment", Value: "development"})
	assert.Contains(t, *resp.Tags, admin.ResourceTag{Key: "team", Value: "platform"})
}

func testDeleteStack(t *testing.T, c *localTestContext) {
	t.Helper()
	utility.DeleteStack(t, c.cfnClient, stackName)
	resp, _, _ := c.atlasClient.FlexClustersApi.GetFlexCluster(ctx.Background(), c.templateData.ProjectID, c.templateData.ClusterName).Execute()
	assert.Nil(t, resp)
}

func cleanupResources(t *testing.T, c *localTestContext) {
	t.Helper()
	utility.DeleteStackForCleanup(t, c.cfnClient, stackName)
	t.Log("Cleaning up resources")
	utility.DeleteProject(t, c.atlasClient20231115014, c.templateData.ProjectID)
}

func (c *localTestContext) setupPrerequisites(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		cleanupResources(t, c)
	})
	t.Log("Setting up prerequisites")
	projectID := utility.CreateProject(t, c.atlasClient20231115014, orgID, projectName)
	c.templateData = testData{
		ResourceTypeName: os.Getenv("RESOURCE_TYPE_NAME_FOR_E2E"),
		Profile:          profile,
		ProjectID:        projectID,
		ClusterName:      clusterName,
	}
	var err error
	c.template, err = newCFNTemplate(c.templateData)
	utility.FailNowIfError(t, "Error while reading CFN Template: %v", err)
}

func newCFNTemplate(data testData) (string, error) {
	return utility.ExecuteGoTemplate(cfnTemplatePath, data)
}
