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

package util

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func GetAWSClient() (client *cfn.Client, err error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, errors.New("Error loading AWS configuration: " + err.Error())
	}
	log.Println("Loaded AWS configuration")
	return cfn.NewFromConfig(cfg), nil
}

func CreateStack(client *cfn.Client, stackName string, fileContent string) (*cfn.DescribeStacksOutput, error) {
	templateBody := fileContent
	outpt, err := createStackAndWait(client, stackName, templateBody)
	if err != nil {
		return nil, err
	}
	return outpt, nil
}

func createStackAndWait(client *cfn.Client, name, stackBody string) (*cfn.DescribeStacksOutput, error) {
	creq := &cfn.CreateStackInput{
		StackName: aws.String(name),
		//OnFailure: aws.String("ROLLBACK"),
		TemplateBody: aws.String(stackBody),
	}

	resp, err := client.CreateStack(context.Background(), creq)
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
		switch string(resp.Stacks[0].StackStatus) {
		case "CREATE_COMPLETE":
			return resp, nil
		case "CREATE_FAILED":
			return nil, errors.New(*aws.String(*resp.Stacks[0].StackStatusReason))
		}
		log.Println("Waiting for stack creation: ", resp.Stacks[0].StackStatus)
		time.Sleep(3 * time.Second)
	}
}

func DeleteStack(svc *cfn.Client, name string) (*cfn.DescribeStacksOutput, error) {

	resp, err := deleteStackAndWait(svc, name)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func deleteStackAndWait(svc *cfn.Client, stackName string) (*cfn.DescribeStacksOutput, error) {
	dreq := &cfn.DeleteStackInput{
		StackName: aws.String(stackName),
	}
	_, err := svc.DeleteStack(context.Background(), dreq)
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

		switch string(resp.Stacks[0].StackStatus) {
		case "DELETE_COMPLETE":
			return resp, nil
		case "DELETE_FAILED":
			return nil, errors.New(*aws.String(*resp.Stacks[0].StackStatusReason))
		}
		log.Println("Waiting for stack deletion: ", resp.Stacks[0].StackStatus)
		time.Sleep(3 * time.Second)
	}
}

func UpdateStack(svc *cfn.Client, stackName string, templateBody string) (*cfn.DescribeStacksOutput, error) {
	outpt, err := updateStackAndWait(svc, stackName, templateBody)
	if err != nil {
		return nil, err
	}
	return outpt, nil
}

func updateStackAndWait(svc *cfn.Client, stackName, stackBody string) (*cfn.DescribeStacksOutput, error) {

	input := &cfn.UpdateStackInput{
		StackName:    aws.String(stackName),
		TemplateBody: aws.String(stackBody),
	}

	updateOutput, err := svc.UpdateStack(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("Error updating cloudformation stack: %v", err)
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
		statusString := string(resp.Stacks[0].StackStatus)
		switch statusString {
		case "UPDATE_COMPLETE":
			return resp, nil
		case "UPDATE_FAILED", "UPDATE_ROLLBACK_COMPLETE", "UPDATE_ROLLBACK_FAILED":
			errMsg := fmt.Sprintf("Stack status: %s : %s", statusString, *resp.Stacks[0].StackStatusReason)
			return nil, errors.New(errMsg)
		}
		log.Println("Waiting for stack update: ", resp.Stacks[0].StackStatus)
		time.Sleep(3 * time.Second)
	}
}

func ValidateTemplate(svc *cfn.Client, template string) (bool, error) {
	input := &cfn.ValidateTemplateInput{
		TemplateBody: aws.String(template),
	}

	_, err := svc.ValidateTemplate(context.Background(), input)

	if err != nil {
		return false, fmt.Errorf("Invalid cloudformation stack: %v", err)
	}

	return true, nil
}
