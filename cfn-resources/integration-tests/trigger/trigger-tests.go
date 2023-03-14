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
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/integration-tests/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	cfnClient *cfn.Client
	template  []byte
	stackName string
	triggerId string
	appId     string
	projectId string
}

func (s *TestSuite) SetupAllSuite() {
	s.stackName = "stack-trigger-int-test"
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}
	fmt.Println("Loaded AWS configuration")
	s.cfnClient = cfn.NewFromConfig(cfg)

	// Read required data from resource CFN template
	s.readFromJsonTemplate("../../examples/trigger/trigger.json")
}

func (s *TestSuite) TestCreateStack(t *testing.T) {
	output, err := util.CreateStack(s.cfnClient, s.stackName, s.template)
	if err != nil {
		fmt.Println("Error creating stack:", err)
		return
	}
	s.triggerId = *aws.String(*output.Stacks[0].Outputs[0].OutputValue)

	ctx := context.Background()
	client, err := util.GetRealmClient(ctx)
	_, resp, err := client.EventTriggers.Get(ctx, s.projectId, s.appId, s.triggerId)

	assert.Equal(t, resp.StatusCode, 200)
}

func (s *TestSuite) TestUpdateStack(t *testing.T) {

}

func (s *TestSuite) TestDeleteStack(t *testing.T) {
	util.DeleteStack(s.cfnClient, s.stackName)
	// assert resource deleted
}

func (s *TestSuite) TearDownSuite() {

}

func (s *TestSuite) readFromJsonTemplate(filePath string) {
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

	s.appId = properties["AppId"].(string)
	s.projectId = properties["ProjectId"].(string)
	s.template = data
}

func TestTriggerTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
