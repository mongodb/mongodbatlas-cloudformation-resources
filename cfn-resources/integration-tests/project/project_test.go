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
package project

import (
	"bytes"
	ctx "context"
	"os"
	"path"
	"testing"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/aws"
	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	util "github.com/mongodb/mongodbatlas-cloudformation-resources/integration-tests/utility"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

type LocalTestContext struct {
	cfnClient      *cfn.Client
	atlasClient    *mongodbatlas.Client
	projectTmplObj TestProject
	resourceCtx    util.ResourceContext
	template       string
	err            error
}

type TestProject struct {
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
)

var (
	profile         = os.Getenv("ATLAS_SECRET_PROFILE")
	orgID           = os.Getenv("ATLAS_ORG_ID")
	e2eRandSuffix   = util.GetRandNum().String()
	testProjectName = "cfn-e2e-project" + e2eRandSuffix
	testTeamName    = "cfn-e2e-team" + e2eRandSuffix
	stackName       = "stack-project-e2e-" + e2eRandSuffix
)

func TestProjectCFN(t *testing.T) {
	testCtx := setupSuite(t)

	t.Run("Validate Template", func(t *testing.T) {
		testIsTemplateValid(t, testCtx)
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

func setupSuite(t *testing.T) *LocalTestContext {
	t.Helper()
	t.Log("Setting up suite")
	testCtx := new(LocalTestContext)
	testCtx.setUp(t)

	return testCtx
}

func (c *LocalTestContext) setUp(t *testing.T) {
	c.resourceCtx = util.InitResourceCtx(e2eRandSuffix, resourceTypeName, resourceDirectory)
	c.setClients(t)
	util.PublishToPrivateRegistry(t, c.resourceCtx)
	c.setupPrerequisites(t)
}

func testIsTemplateValid(t *testing.T, c *LocalTestContext) {
	t.Helper()

	isValid := util.ValidateTemplate(t, c.cfnClient, c.template)
	a := assert.New(t)
	a.True(isValid)
}

func testCreateStack(t *testing.T, c *LocalTestContext) {
	t.Helper()

	output := util.CreateStack(t, c.cfnClient, stackName, c.template)
	c.projectTmplObj.ProjectID = getProjectIDFromStack(output)

	project, getProjectResponse, err := c.atlasClient.Projects.GetOneProject(ctx.Background(), c.projectTmplObj.ProjectID)
	util.FailNowIfError(t, "Error while retrieving Project from Atlas: %v", err)

	teamsAssigned, _, _ := c.atlasClient.Projects.GetProjectTeamsAssigned(ctx.Background(), project.ID)
	util.FailNowIfError(t, "Error when retrieving Project TeamsAssigned: %v", err)

	a := assert.New(t)
	a.Equal(c.projectTmplObj.TeamID, teamsAssigned.Results[0].TeamID)
	a.Equal(200, getProjectResponse.StatusCode)
}

func testUpdateStack(t *testing.T, c *LocalTestContext) {
	t.Helper()

	// create CFN template with updated project name
	c.projectTmplObj.Name += "-updated"
	c.template, c.err = newCFNTemplate(c.projectTmplObj)

	output := util.UpdateStack(t, c.cfnClient, stackName, c.template)
	c.projectTmplObj.ProjectID = getProjectIDFromStack(output)

	project, _, err := c.atlasClient.Projects.GetOneProject(ctx.Background(), c.projectTmplObj.ProjectID)
	util.FailNowIfError(t, "Error while retrieving Project from Atlas: %v", err)

	teamsAssigned, _, err := c.atlasClient.Projects.GetProjectTeamsAssigned(ctx.Background(), project.ID)
	util.FailNowIfError(t, "Error when retrieving Project TeamsAssigned: %v", err)

	a := assert.New(t)
	a.Equal(c.projectTmplObj.TeamID, teamsAssigned.Results[0].TeamID)
	a.Equal(project.Name, c.projectTmplObj.Name)
}

func testDeleteStack(t *testing.T, c *LocalTestContext) {
	t.Helper()

	util.DeleteStack(t, c.cfnClient, stackName)

	_, resp, _ := c.atlasClient.Projects.GetOneProject(ctx.Background(), c.projectTmplObj.ProjectID)

	a := assert.New(t)
	a.Equal(resp.StatusCode, 404)
}

func (c *LocalTestContext) setClients(t *testing.T) {
	t.Helper()

	t.Log("Setting clients")
	c.atlasClient, c.err = util.NewMongoDBClient()
	util.FailNowIfError(t, "Unable to create atlas client: %v", c.err)

	c.cfnClient, c.err = util.NewCFNClient()
	util.FailNowIfError(t, "Unable to create AWS client, please check AWS config is correctly setup: %v", c.err)
}

func getProjectIDFromStack(output *cfn.DescribeStacksOutput) string {
	stackOutputs := output.Stacks[0].Outputs
	for i := 0; i < len(stackOutputs); i++ {
		if *aws.String(*stackOutputs[i].OutputKey) == "ProjectId" {
			return *aws.String(*stackOutputs[1].OutputValue)
		}
	}
	return ""
}

func deleteProjectIfExists(t *testing.T, c *LocalTestContext) {
	_, _, err := c.atlasClient.Projects.GetOneProject(ctx.Background(), c.projectTmplObj.ProjectID)
	if err == nil {
		_, err = c.atlasClient.Projects.Delete(ctx.Background(), c.projectTmplObj.ProjectID)
		if err != nil {
			t.Logf("Atlas Project could not be deleted during cleanup")
		} else {
			t.Logf("Atlas Project successfully deleted during cleanup")
		}
	}
}

func cleanupPrerequisites(t *testing.T, c *LocalTestContext) {
	t.Log("Cleaning up prerequisites")
	_, err := c.atlasClient.Teams.RemoveTeamFromOrganization(ctx.Background(), orgID, c.projectTmplObj.TeamID)
	if err != nil {
		t.Logf("Atlas Team could not be deleted during cleanup: %s\n", err.Error())
	}
}

func (c *LocalTestContext) setupPrerequisites(t *testing.T) {
	t.Log("Setting up prerequisites")
	team, _ := util.NewAtlasTeam(ctx.Background(), c.atlasClient, testTeamName, orgID)
	t.Cleanup(func() {
		cleanupPrerequisites(t, c)
		deleteProjectIfExists(t, c)
	})

	c.projectTmplObj = TestProject{
		Name:             testProjectName,
		OrgID:            orgID,
		Profile:          profile,
		TeamID:           team.ID,
		ResourceTypeName: c.resourceCtx.ResourceTypeNameForE2e,
	}

	// Read required data from resource CFN template
	c.template, c.err = newCFNTemplate(c.projectTmplObj)
	util.FailNowIfError(t, "Error while reading CFN Template: %v", c.err)
}

func newCFNTemplate(tmpl TestProject) (string, error) {
	return executeGoTemplate(tmpl)
}

func executeGoTemplate(projectTmpl TestProject) (string, error) {
	var cfnGoTemplateStr bytes.Buffer
	cfnTemplatePath := "template/cfnTemplate.json"

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
