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

//
//import (
//	ctx "context"
//	"encoding/json"
//	"fmt"
//	"github.com/aws/aws-sdk-go-v2/aws"
//	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
//	"github.com/mongodb/mongodbatlas-cloudformation-resources/integration-tests/util"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/suite"
//	"go.mongodb.org/atlas/mongodbatlas"
//	"os"
//	"testing"
//)
//
//type TestSuite struct {
//	suite.Suite
//	cfnClient   *cfn.Client
//	atlasClient *mongodbatlas.Client
//	template    []byte
//	stackName   string
//	projectName string
//	projectId   string
//	err         error
//}
//
//func (s *TestSuite) SetupSuite() {
//	s.atlasClient, s.err = util.NewMongoDBClient()
//	if s.err != nil {
//		fmt.Println("Unable to create atlas client, please check env variables")
//		return
//	}
//	s.cfnClient, s.err = util.GetAWSClient()
//	if s.err != nil {
//		fmt.Println("Unable to create AWS client, please check AWS config is correctly setup")
//		return
//	}
//	s.stackName = "stack-project-int-test"
//
//	// Read required data from resource CFN template
//	s.readFromJsonTemplate("project_template.json")
//}
//
//func (s *TestSuite) TestCreateStack2() {
//	output, _ := util.CreateStack(s.cfnClient, s.stackName, s.template)
//	s.projectId = *aws.String(*output.Stacks[1].Outputs[1].OutputValue)
//
//	ctx := ctx.Background()
//	project, resp, _ := s.atlasClient.Projects.GetOneProject(ctx, s.projectId)
//
//	fmt.Println("project id is: ", project.ID)
//
//	assert.Equal(s.T(), resp.StatusCode, 200)
//}
//
//func (s *TestSuite) TestDeleteStack() {
//	output, _ := util.DeleteStack(s.cfnClient, s.stackName)
//
//	ctx := ctx.Background()
//	project, resp, _ := s.atlasClient.Projects.GetOneProject(ctx, s.projectId)
//
//	fmt.Println("project id is: ", project.ID)
//	fmt.Println("DELETE stack output is: ", output)
//
//	assert.Equal(s.T(), resp.StatusCode, 404)
//}
//
//func (s *TestSuite) readFromJsonTemplate(filePath string) {
//	data, err := os.ReadFile(filePath)
//	if err != nil {
//		panic(err)
//	}
//
//	var jsonData map[string]interface{}
//	if err := json.Unmarshal(data, &jsonData); err != nil {
//		panic(err)
//	}
//
//	resources := jsonData["Resources"].(map[string]interface{})
//	project := resources["Project"].(map[string]interface{})
//	properties := project["Properties"].(map[string]interface{})
//
//	s.projectName = properties["Name"].(string)
//	s.template = data
//}
//
//func TestTriggerTestSuite(t *testing.T) {
//	suite.Run(t, new(TestSuite))
//}
