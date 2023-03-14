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

package project

import (
	ctx "context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/integration-tests/util"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

type TestContext struct {
	cfnClient   *cfn.Client
	atlasClient *mongodbatlas.Client
	template    []byte
	stackName   string
	projectName string
	projectId   string
	err         error
}

func TestProjectCFN(t *testing.T) {
	ctx := new(TestContext)
	ctx.setUp()
	testIsTemplateValid(t)
	testCreateStack(t, ctx)
	testUpdateStack(t, ctx)
	testDeleteStack(t, ctx)
}

func testIsTemplateValid(t *testing.T) {
	// TODO
}

func (s *TestContext) setUp() error {
	s.atlasClient, s.err = util.NewMongoDBClient()
	if s.err != nil {
		fmt.Println("Unable to create atlas client, please check env variables")
		return s.err
	}
	s.cfnClient, s.err = util.GetAWSClient()
	if s.err != nil {
		fmt.Println("Unable to create AWS client, please check AWS config is correctly setup")
		return s.err
	}
	s.stackName = "stack-project-int-test"

	// Read required data from resource CFN template
	s.readFromJsonTemplate("project_template.json")
	return nil
}

func testCreateStack(t *testing.T, c *TestContext) {
	output, _ := util.CreateStack(c.cfnClient, c.stackName, c.template)
	c.projectId = getProjectIdFromStack(output)

	ctx := ctx.Background()
	_, resp, _ := c.atlasClient.Projects.GetOneProject(ctx, c.projectId)

	assert.Equal(t, resp.StatusCode, 200)
}

func testUpdateStack(t *testing.T, c *TestContext) {
	// TODO
}

func testDeleteStack(t *testing.T, c *TestContext) {
	util.DeleteStack(c.cfnClient, c.stackName)

	ctx := ctx.Background()
	_, resp, _ := c.atlasClient.Projects.GetOneProject(ctx, c.projectId)

	assert.Equal(t, resp.StatusCode, 404)
}

func (s *TestContext) readFromJsonTemplate(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		panic(err)
	}

	resources := jsonData["Resources"].(map[string]interface{})
	project := resources["Project"].(map[string]interface{})
	properties := project["Properties"].(map[string]interface{})

	s.projectName = properties["Name"].(string)
	s.template = data
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
