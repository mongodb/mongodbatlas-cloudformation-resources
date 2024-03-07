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

package utility

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

const stackStatusWait = 2 * time.Second

func NewCFNClient() (client *cfn.Client, err error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error loading AWS configuration: %w", err)
	}
	return cfn.NewFromConfig(cfg), nil
}

func CreateStack(t *testing.T, client *cfn.Client, stackName string, fileContent string) *cfn.DescribeStacksOutput {
	t.Helper()
	output, err := createStackAndWait(client, stackName, fileContent)
	FailNowIfError(t, "Error during stack creation: %v", err)

	return output
}

func createStackAndWait(client *cfn.Client, name, stackBody string) (*cfn.DescribeStacksOutput, error) {
	input := &cfn.CreateStackInput{
		StackName:    aws.String(name),
		TemplateBody: aws.String(stackBody),
	}

	resp, err := client.CreateStack(context.Background(), input)
	if err != nil {
		return nil, err
	}

	describeStackOutput, err := waitForStackCreateComplete(client, *resp.StackId)
	if err != nil {
		return nil, err
	}

	return describeStackOutput, nil
}

func waitForStackCreateComplete(svc *cfn.Client, stackID string) (*cfn.DescribeStacksOutput, error) {
	req := cfn.DescribeStacksInput{
		StackName: aws.String(stackID),
	}
	for {
		resp, err := svc.DescribeStacks(context.Background(), &req)
		if err != nil {
			return nil, err
		}
		if len(resp.Stacks) == 0 {
			return nil, fmt.Errorf("stack not found")
		}
		statusStr := string(resp.Stacks[0].StackStatus)
		switch statusStr {
		case "CREATE_COMPLETE":
			return resp, nil
		case "CREATE_FAILED", "ROLLBACK_COMPLETE":
			return nil, fmt.Errorf("stack status: %s : %s", statusStr, util.SafeString(resp.Stacks[0].StackStatusReason))
		}
		time.Sleep(stackStatusWait)
	}
}

func DeleteStack(t *testing.T, client *cfn.Client, stackName string) *cfn.DescribeStacksOutput {
	t.Helper()
	output, err := deleteStackAndWait(client, stackName)
	FailNowIfError(t, "Error during stack deletion: %v", err)

	return output
}

func deleteStackAndWait(svc *cfn.Client, stackName string) (*cfn.DescribeStacksOutput, error) {
	input := &cfn.DeleteStackInput{
		StackName: aws.String(stackName),
	}
	_, err := svc.DeleteStack(context.Background(), input)
	if err != nil {
		return nil, err
	}

	describeStackOutput, err := waitForStackDeleteComplete(svc, stackName)
	if err != nil {
		return nil, err
	}

	return describeStackOutput, nil
}

func waitForStackDeleteComplete(svc *cfn.Client, stackID string) (*cfn.DescribeStacksOutput, error) {
	req := cfn.DescribeStacksInput{
		StackName: aws.String(stackID),
	}
	for {
		resp, err := svc.DescribeStacks(context.Background(), &req)
		if err != nil {
			return resp, nil
		}

		statusStr := string(resp.Stacks[0].StackStatus)
		switch statusStr {
		case "DELETE_COMPLETE":
			return resp, nil
		case "DELETE_FAILED", "ROLLBACK_COMPLETE":
			return nil, fmt.Errorf("stack status: %s : %s", statusStr, util.SafeString(resp.Stacks[0].StackStatusReason))
		}
		time.Sleep(stackStatusWait)
	}
}

func UpdateStack(t *testing.T, client *cfn.Client, stackName string, templateBody string) *cfn.DescribeStacksOutput {
	t.Helper()
	output, err := updateStackAndWait(client, stackName, templateBody)
	FailNowIfError(t, "Error during stack update: %v", err)

	return output
}

func updateStackAndWait(svc *cfn.Client, stackName, stackBody string) (*cfn.DescribeStacksOutput, error) {
	input := &cfn.UpdateStackInput{
		StackName:    aws.String(stackName),
		TemplateBody: aws.String(stackBody),
	}

	updateOutput, err := svc.UpdateStack(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("error updating cloudformation stack: %w", err)
	}

	return waitForStackUpdateComplete(svc, *updateOutput.StackId)
}

func waitForStackUpdateComplete(svc *cfn.Client, stackID string) (*cfn.DescribeStacksOutput, error) {
	req := cfn.DescribeStacksInput{
		StackName: aws.String(stackID),
	}
	for {
		resp, err := svc.DescribeStacks(context.Background(), &req)
		if err != nil {
			return nil, err
		}
		if len(resp.Stacks) == 0 {
			return nil, fmt.Errorf("stack not found")
		}
		statusStr := string(resp.Stacks[0].StackStatus)
		switch statusStr {
		case "UPDATE_COMPLETE":
			return resp, nil
		case "UPDATE_FAILED", "UPDATE_ROLLBACK_COMPLETE", "UPDATE_ROLLBACK_FAILED", "ROLLBACK_COMPLETE":
			return nil, fmt.Errorf("stack status: %s : %s", statusStr, util.SafeString(resp.Stacks[0].StackStatusReason))
		}
		time.Sleep(stackStatusWait)
	}
}

func TestIsTemplateValid(t *testing.T, svc *cfn.Client, template string) {
	t.Helper()
	input := &cfn.ValidateTemplateInput{
		TemplateBody: aws.String(template),
	}

	_, err := svc.ValidateTemplate(context.Background(), input)
	FailNowIfError(t, "invalid cloudformation stack: %v", err)
}

func DeleteStackForCleanup(t *testing.T, c *cfn.Client, stackName string) {
	t.Helper()
	input := &cfn.DeleteStackInput{
		StackName: aws.String(stackName),
	}
	_, err := c.DeleteStack(context.Background(), input)
	t.Logf("error response when deleting stack for cleanup: %v", err)
}
