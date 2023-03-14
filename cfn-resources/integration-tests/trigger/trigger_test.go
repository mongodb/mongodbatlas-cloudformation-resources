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

package trigger

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
	"go.mongodb.org/realm/realm"
)

type TestContext struct {
	cfnClient   *cfn.Client
	realmClient *realm.Client
	template    []byte
	stackName   string
	triggerId   string
	appId       string
	projectId   string
	err         error
}

func TestTriggerCFN(t *testing.T) {
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

func (c *TestContext) setUp() error {
	ctx := ctx.Background()
	c.realmClient, c.err = util.GetRealmClient(ctx)
	if c.err != nil {
		fmt.Println("Unable to create realm client, please check env variables")
		return c.err
	}

	c.cfnClient, c.err = util.GetAWSClient()
	if c.err != nil {
		fmt.Println("Unable to create AWS client, please check AWS config is correctly setup")
		return c.err
	}
	c.stackName = "stack-trigger-int-test"

	// Read required data from resource CFN template
	c.readFromJsonTemplate("trigger_template.json")
	return nil
}

func testCreateStack(t *testing.T, c *TestContext) {
	output, err := util.CreateStack(c.cfnClient, c.stackName, c.template)
	if err != nil {
		fmt.Println("Error creating stack:", err)
		return
	}
	c.triggerId = *aws.String(*output.Stacks[0].Outputs[0].OutputValue)

	ctx := ctx.Background()
	_, resp, err := c.realmClient.EventTriggers.Get(ctx, c.projectId, c.appId, c.triggerId)

	assert.Equal(t, resp.StatusCode, 200)
}

func testUpdateStack(t *testing.T, c *TestContext) {
	// TODO
}

func testDeleteStack(t *testing.T, c *TestContext) {
	util.DeleteStack(c.cfnClient, c.stackName)

	ctx := ctx.Background()
	_, resp, _ := c.realmClient.EventTriggers.Get(ctx, c.projectId, c.appId, c.triggerId)

	assert.Equal(t, resp.StatusCode, 404)
}

func (c *TestContext) readFromJsonTemplate(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		panic(err)
	}

	resources := jsonData["Resources"].(map[string]interface{})
	eventTrigger := resources["EventTrigger"].(map[string]interface{})
	properties := eventTrigger["Properties"].(map[string]interface{})

	c.appId = properties["AppId"].(string)
	c.projectId = properties["ProjectId"].(string)
	c.template = data
}
