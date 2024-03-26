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
package project_test

import (
	ctx "context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/test/e2e/utility"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20231115008/admin"
)

type localTestContext struct {
	cfnClient      *cloudformation.Client
	atlasClient    *admin.APIClient
	projectTmplObj testProject
	resourceCtx    utility.ResourceContext

	template string
	err      error
}

type testProject struct {
	ResourceTypeName string
	Name             string
	OrgID            string
	Profile          string
	TeamID           string
	ProjectName      string
	ProjectID        string
}

const (
	resourceTypeName  = "MongoDB::Atlas::Project"
	resourceDirectory = "project"
	cfnTemplatePath   = "project.json.template"
)

var (
	profile         = os.Getenv("MONGODB_ATLAS_SECRET_PROFILE")
	orgID           = os.Getenv("MONGODB_ATLAS_ORG_ID")
	e2eRandSuffix   = utility.GetRandNum().String()
	testProjectName = "cfn-e2e-project" + e2eRandSuffix
	testTeamName    = "cfn-e2e-team" + e2eRandSuffix
	stackName       = "stack-project-e2e-" + e2eRandSuffix
)

func TestProjectCFN(t *testing.T) {
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

	output := utility.CreateStack(t, c.cfnClient, stackName, c.template)
	c.projectTmplObj.ProjectID = getProjectIDFromStack(output)

	project, getProjectResponse, err := c.atlasClient.ProjectsApi.GetProject(ctx.Background(), c.projectTmplObj.ProjectID).Execute()
	utility.FailNowIfError(t, "Error while retrieving Project from Atlas: %v", err)

	teamsAssigned, _, _ := c.atlasClient.TeamsApi.ListProjectTeams(ctx.Background(), *project.Id).Execute()
	utility.FailNowIfError(t, "Error when retrieving Project TeamsAssigned: %v", err)

	a := assert.New(t)
	a.Equal(c.projectTmplObj.TeamID, *teamsAssigned.GetResults()[0].TeamId)
	a.Equal(200, getProjectResponse.StatusCode)
}

func testUpdateStack(t *testing.T, c *localTestContext) {
	t.Helper()

	// create CFN template with updated project name
	c.projectTmplObj.Name += "-updated"
	c.template, c.err = newCFNTemplate(c.projectTmplObj)

	output := utility.UpdateStack(t, c.cfnClient, stackName, c.template)
	c.projectTmplObj.ProjectID = getProjectIDFromStack(output)

	project, _, err := c.atlasClient.ProjectsApi.GetProject(ctx.Background(), c.projectTmplObj.ProjectID).Execute()
	utility.FailNowIfError(t, "Error while retrieving Project from Atlas: %v", err)

	teamsAssigned, _, err := c.atlasClient.TeamsApi.ListProjectTeams(ctx.Background(), *project.Id).Execute()
	utility.FailNowIfError(t, "Error when retrieving Project TeamsAssigned: %v", err)

	a := assert.New(t)
	a.Equal(c.projectTmplObj.TeamID, *teamsAssigned.GetResults()[0].TeamId)
	a.Equal(project.Name, c.projectTmplObj.Name)
}

func testDeleteStack(t *testing.T, c *localTestContext) {
	t.Helper()

	utility.DeleteStack(t, c.cfnClient, stackName)
	_, resp, _ := c.atlasClient.ProjectsApi.GetProject(ctx.Background(), c.projectTmplObj.ProjectID).Execute()

	a := assert.New(t)
	a.Equal(resp.StatusCode, 404)
}

func getProjectIDFromStack(output *cloudformation.DescribeStacksOutput) string {
	stackOutputs := output.Stacks[0].Outputs
	for i := 0; i < len(stackOutputs); i++ {
		if *aws.String(*stackOutputs[i].OutputKey) == "ProjectId" {
			return *aws.String(*stackOutputs[1].OutputValue)
		}
	}
	return ""
}

func cleanupResources(t *testing.T, c *localTestContext) {
	t.Helper()
	utility.DeleteStackForCleanup(t, c.cfnClient, stackName)

	_, _, err := c.atlasClient.ProjectsApi.GetProject(ctx.Background(), c.projectTmplObj.ProjectID).Execute()
	if err == nil {
		_, _, err = c.atlasClient.ProjectsApi.DeleteProject(ctx.Background(), c.projectTmplObj.ProjectID).Execute()
		if err != nil {
			t.Logf("Atlas Project could not be deleted during cleanup: %v", err)
		} else {
			t.Logf("Atlas Project successfully deleted during cleanup: %v", err)
		}
	}
}

func cleanupPrerequisites(t *testing.T, c *localTestContext) {
	t.Helper()
	t.Log("Cleaning up prerequisites")
	_, _, err := c.atlasClient.TeamsApi.DeleteTeam(ctx.Background(), orgID, c.projectTmplObj.TeamID).Execute()
	if err != nil {
		t.Logf("Atlas Team could not be deleted during cleanup: %s\n", err.Error())
	}
}

func (c *localTestContext) setupPrerequisites(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		cleanupPrerequisites(t, c)
		cleanupResources(t, c)
	})

	t.Log("Setting up prerequisites")
	team, _ := utility.NewAtlasTeam(ctx.Background(), c.atlasClient, testTeamName, orgID)

	c.projectTmplObj = testProject{
		Name:             testProjectName,
		OrgID:            orgID,
		Profile:          profile,
		TeamID:           *team.Id,
		ResourceTypeName: os.Getenv("RESOURCE_TYPE_NAME_FOR_E2E"),
	}

	// Read required data from resource CFN template
	c.template, c.err = newCFNTemplate(c.projectTmplObj)
	utility.FailNowIfError(t, "Error while reading CFN Template: %v", c.err)
}

func newCFNTemplate(tmpl testProject) (string, error) {
	return utility.ExecuteGoTemplate(cfnTemplatePath, tmpl)
}
