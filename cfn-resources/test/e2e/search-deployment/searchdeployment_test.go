// Copyright 2023 MongoDB Inc
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
	"bytes"
	ctx "context"
	"os"
	"path"
	"testing"
	"text/template"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/test/e2e/utility"

	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

type localTestContext struct {
	cfnClient      *cfn.Client
	atlasClient    *admin.APIClient
	projectTmplObj testProject
	resourceCtx    utility.ResourceContext

	template string
	err      error
}

type testProject struct {
	ResourceTypeName string
	Profile          string
	ProjectID        string
	ClusterName      string
}

const (
	resourceTypeName  = "MongoDB::Atlas::SearchDeployment"
	resourceDirectory = "search-deployment"
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
	deployment, resp, err := c.atlasClient.AtlasSearchApi.GetAtlasSearchDeployment(ctx.Background(), c.projectTmplObj.ProjectID, c.projectTmplObj.ClusterName).Execute()
	utility.FailNowIfError(t, "Error while retrieving Project from Atlas: %v", err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, c.projectTmplObj.ProjectID, deployment.GetGroupId())
}

func testUpdateStack(t *testing.T, c *localTestContext) {
	t.Helper()
}

func testDeleteStack(t *testing.T, c *localTestContext) {
	t.Helper()

	utility.DeleteStack(t, c.cfnClient, stackName)
	_, resp, _ := c.atlasClient.AtlasSearchApi.GetAtlasSearchDeployment(ctx.Background(), c.projectTmplObj.ProjectID, c.projectTmplObj.ClusterName).Execute()
	assert.Equal(t, 404, resp.StatusCode)
}

func cleanupResources(t *testing.T, c *localTestContext) {
	t.Helper()
	utility.DeleteStackForCleanup(t, c.cfnClient, stackName)

	t.Log("Cleaning up resources")
	_, _, err := c.atlasClient.ProjectsApi.DeleteProject(ctx.Background(), c.projectTmplObj.ProjectID).Execute()
	if err != nil {
		t.Logf("Atlas Project could not be deleted during cleanup: %s\n", err.Error())
	}
}

func (c *localTestContext) setupPrerequisites(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		cleanupResources(t, c)
	})

	t.Log("Setting up prerequisites")
	params := &admin.Group{
		Name:  projectName,
		OrgId: orgID,
	}
	resp, _, err := c.atlasClient.ProjectsApi.CreateProject(ctx.Background(), params).Execute()
	utility.FailNowIfError(t, "Error creating Atlas Project: %v", err)

	c.projectTmplObj = testProject{
		ResourceTypeName: os.Getenv("RESOURCE_TYPE_NAME_FOR_E2E"),
		Profile:          profile,
		ProjectID:        resp.GetId(),
		ClusterName:      clusterName,
	}

	// Read required data from resource CFN template
	c.template, c.err = newCFNTemplate(c.projectTmplObj)
	utility.FailNowIfError(t, "Error while reading CFN Template: %v", c.err)
}

func newCFNTemplate(tmpl testProject) (string, error) {
	return executeGoTemplate(tmpl)
}

func executeGoTemplate(projectTmpl testProject) (string, error) {
	var cfnGoTemplateStr bytes.Buffer
	cfnTemplatePath := "searchdeployment_template.json"

	name := path.Base(cfnTemplatePath)
	cfnGoTemplate, err := template.New(name).ParseFiles(cfnTemplatePath)
	if err != nil {
		return "", err
	}
	err = cfnGoTemplate.Execute(&cfnGoTemplateStr, projectTmpl)
	if err != nil {
		return "", err
	}
	return cfnGoTemplateStr.String(), nil
}
