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
	"fmt"
	"log"
	"os"
	"path"
	"testing"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/aws"
	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/integration-tests/util"
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
	OrgId            string
	Profile          string
	TeamId           string
	ProjectName      string
	ProjectId        string
}

const (
	resourceTypeName  = "MongoDB::Atlas::Project"
	resourceDirectory = "project"
)

var (
	profile         = os.Getenv("ATLAS_SECRET_PROFILE")
	orgId           = os.Getenv("ATLAS_ORG_ID")
	e2eRandSuffix   = fmt.Sprint(util.GetRandNum())
	testProjectName = "cfn-e2e-project" + e2eRandSuffix
	testTeamName    = "cfn-e2e-team" + e2eRandSuffix
	stackName       = "stack-project-e2e-" + e2eRandSuffix
)

func TestProjectCFN(t *testing.T) {
	teardownSuite, testCtx := setupSuite(t)
	defer teardownSuite(t)

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

func setupSuite(t *testing.T) (func(t *testing.T), *LocalTestContext) {
	log.Println("Setting up suite")
	testCtx := new(LocalTestContext)
	testCtx.setUp()

	// Return function to teardown the test suite
	return func(t *testing.T) {
		log.Println("Tearing down suite")
		util.RunCleanupScript(testCtx.resourceCtx)
		deleteProjectIfExists(testCtx)
		cleanupPrerequisites(testCtx)
	}, testCtx
}

func (c *LocalTestContext) setUp() {
	c.resourceCtx = util.InitResourceCtx(e2eRandSuffix, resourceTypeName, resourceDirectory)
	c.setClients()
	util.PublishToPrivateRegistry(c.resourceCtx)
	c.setupPrerequisites()
}

func testIsTemplateValid(t *testing.T, c *LocalTestContext) {
	t.Helper()
	isValid, _ := util.ValidateTemplate(c.cfnClient, c.template)
	a := assert.New(t)
	a.True(isValid)
}

func testCreateStack(t *testing.T, c *LocalTestContext) {
	t.Helper()
	ctx := ctx.Background()

	output, _ := util.CreateStack(c.cfnClient, stackName, c.template)
	c.projectTmplObj.ProjectId = getProjectIdFromStack(output)

	project, getProjectResponse, _ := c.atlasClient.Projects.GetOneProject(ctx, c.projectTmplObj.ProjectId)
	teamsAssigned, _, err := c.atlasClient.Projects.GetProjectTeamsAssigned(ctx, project.ID)
	if err != nil {
		log.Printf("Error when validating project teamsAssigned: %s", err.Error())
		return
	}

	a := assert.New(t)
	a.Equal(c.projectTmplObj.TeamId, teamsAssigned.Results[0].TeamID)
	a.Equal(200, getProjectResponse.StatusCode)
}

func testUpdateStack(t *testing.T, c *LocalTestContext) {
	t.Helper()
	ctx := ctx.Background()

	// create CFN template with updated project name
	c.projectTmplObj.Name += "-updated"
	c.template, c.err = getCFNTemplate(c.projectTmplObj)

	output, _ := util.UpdateStack(c.cfnClient, stackName, c.template)
	c.projectTmplObj.ProjectId = getProjectIdFromStack(output)

	project, _, _ := c.atlasClient.Projects.GetOneProject(ctx, c.projectTmplObj.ProjectId)
	teamsAssigned, _, err := c.atlasClient.Projects.GetProjectTeamsAssigned(ctx, project.ID)
	if err != nil {
		log.Printf("Error when validating project teamsAssigned: %s", err.Error())
		return
	}

	a := assert.New(t)
	a.Equal(c.projectTmplObj.TeamId, teamsAssigned.Results[0].TeamID)
	a.Equal(project.Name, c.projectTmplObj.Name)
}

func testDeleteStack(t *testing.T, c *LocalTestContext) {
	t.Helper()
	util.DeleteStack(c.cfnClient, stackName)

	ctx := ctx.Background()
	_, resp, _ := c.atlasClient.Projects.GetOneProject(ctx, c.projectTmplObj.ProjectId)

	a := assert.New(t)
	a.Equal(resp.StatusCode, 404)
}

func (c *LocalTestContext) setClients() error {
	c.atlasClient, c.err = util.NewMongoDBClient()
	if c.err != nil {
		fmt.Println("Unable to create atlas client, please check env variables")
		return c.err
	}
	c.cfnClient, c.err = util.GetAWSClient()
	if c.err != nil {
		fmt.Println("Unable to create AWS client, please check AWS config is correctly setup")
		return c.err
	}
	return nil
}

func getProjectIdFromStack(output *cfn.DescribeStacksOutput) string {
	stackOutputs := output.Stacks[0].Outputs
	for i := 0; i < len(stackOutputs); i++ {
		if *aws.String(*stackOutputs[i].OutputKey) == "ProjectId" {
			return *aws.String(*stackOutputs[1].OutputValue)
		}
	}
	return ""
}

func deleteProjectIfExists(c *LocalTestContext) {
	_, _, err := c.atlasClient.Projects.GetOneProject(ctx.Background(), c.projectTmplObj.ProjectId)
	if err == nil {
		_, err = c.atlasClient.Projects.Delete(ctx.Background(), c.projectTmplObj.ProjectId)
		if err != nil {
			log.Println("Atlas Project could not be deleted during cleanup")
		} else {
			log.Println("Atlas Project successfully deleted during cleanup")
		}
	}
}

func cleanupPrerequisites(c *LocalTestContext) {
	_, err := c.atlasClient.Teams.RemoveTeamFromOrganization(ctx.Background(), orgId, c.projectTmplObj.TeamId)
	if err != nil {
		log.Printf("Atlas Team could not be deleted during cleanup: %s\n", err.Error())
	}
}

func (c *LocalTestContext) setupPrerequisites() error {
	team, _ := util.GetNewAtlasTeam(ctx.Background(), c.atlasClient, testTeamName, orgId)

	c.projectTmplObj = TestProject{
		Name:             testProjectName,
		OrgId:            orgId,
		Profile:          profile,
		TeamId:           team.ID,
		ResourceTypeName: c.resourceCtx.ResourceTypeNameForE2e,
	}

	// Read required data from resource CFN template
	c.template, c.err = getCFNTemplate(c.projectTmplObj)
	if c.err != nil {
		panic(c.err)
	}
	return nil
}

func getCFNTemplate(tmpl TestProject) (string, error) {
	return executeGoTemplate(tmpl)
}

func executeGoTemplate(projectTmpl TestProject) (string, error) {
	var cfnGoTemplateStr bytes.Buffer
	name := path.Base("template/cfnTemplate.json")
	cfnGoTemplate, err := template.New(name).ParseFiles("template/cfnTemplate.json")
	if err != nil {
		panic(err)
	}
	err = cfnGoTemplate.Execute(&cfnGoTemplateStr, projectTmpl)
	if err != nil {
		panic(err)
	}
	return cfnGoTemplateStr.String(), nil
}
