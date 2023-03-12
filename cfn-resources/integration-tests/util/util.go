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
package util

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	realmAuth "go.mongodb.org/realm/auth"
	"go.mongodb.org/realm/realm"
)

var (
	defaultLogLevel = "warning"
	//userAgent          = fmt.Sprintf("%s/%s (%s;%s)", toolName, version.Version, runtime.GOOS, runtime.GOARCH)
	terraformUserAgent = "terraform-provider-mongodbatlas"
)

func CreateStack(client *cloudformation.Client, name string, fileContent []byte) (*cloudformation.DescribeStacksOutput, error) {
	templateBody := string(fileContent)
	outpt, err := createStackAndWait(client, name, templateBody)
	if err != nil {
		return nil, err
	}
	return outpt, nil
}

func createStackAndWait(client *cloudformation.Client, name, stackBody string) (*cloudformation.DescribeStacksOutput, error) {
	creq := &cloudformation.CreateStackInput{
		StackName: aws.String(name),
		OnFailure: types.OnFailureRollback,
		//Capabilities: []*string{aws.String(cloudformation.CapabilityCapabilityIam)},
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

func DeleteStack(svc *cloudformation.Client, name string) error {
	dreq := &cloudformation.DeleteStackInput{
		StackName: aws.String(name),
	}
	_, err := svc.DeleteStack(context.Background(), dreq)
	return err
}

func validateTemplate(svc *cloudformation.Client, template string) (string, error) {

	input := &cloudformation.ValidateTemplateInput{
		TemplateBody: aws.String(template),
	}

	_, err := svc.ValidateTemplate(context.Background(), input)

	if err != nil {
		return "", fmt.Errorf("Invalid cloudformation stack: %v", err)
	}

	return "Template is valid", err
}

//func updateStack(svc *cloudformation.Client, stackName, stackBody string) (string, error) {
//
//	input := &cloudformation.UpdateStackInput{
//		Capabilities: []*string{aws.String(cloudformation.CapabilityCapabilityIam)},
//		StackName:    aws.String(stackName),
//		TemplateBody: aws.String(stackBody),
//	}
//
//	updateOutput, err := svc.UpdateStack(input)
//
//	if err != nil {
//		return "", fmt.Errorf("Error updating cloudformation stack: %v", err)
//	}
//
//	return updateOutput.String(), waitForStackUpdateComplete(svc, *updateOutput.StackId)
//}
//
//func waitForStackUpdateComplete(svc *cloudformation.CloudFormation, stackID string) error {
//	req := cloudformation.DescribeStacksInput{
//		StackName: aws.String(stackID),
//	}
//	for {
//		resp, err := svc.DescribeStacks(&req)
//		if err != nil {
//			return err
//		}
//		if len(resp.Stacks) == 0 {
//			return fmt.Errorf("stack not found")
//		}
//		statusString := aws.StringValue(resp.Stacks[0].StackStatus)
//		switch statusString {
//		case cloudformation.ResourceStatusUpdateComplete:
//			return nil
//		case cloudformation.ResourceStatusUpdateFailed, cloudformation.StackStatusUpdateRollbackComplete, cloudformation.StackStatusUpdateRollbackFailed:
//			errMsg := fmt.Sprintf("Stack status: %s : %s", statusString, aws.StringValue(resp.Stacks[0].StackStatusReason))
//			return errors.New(errMsg)
//		}
//		time.Sleep(3 * time.Second)
//	}
//}

func waitForStackCreateComplete(svc *cloudformation.Client, stackID string) (*cloudformation.DescribeStacksOutput, error) {
	req := cloudformation.DescribeStacksInput{
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
		time.Sleep(3 * time.Second)
	}
}

//func getStackResources(svc *cloudformation.CloudFormation, stackID string) ([]cloudformation.StackResourceSummary, error) {
//	resources := make([]cloudformation.StackResourceSummary, 0)
//	req := cloudformation.ListStackResourcesInput{
//		StackName: aws.String(stackID),
//	}
//	for {
//		resp, err := svc.ListStackResources(&req)
//		if err != nil {
//			return nil, err
//		}
//		for _, s := range resp.StackResourceSummaries {
//			resources = append(resources, *s)
//		}
//		req.NextToken = resp.NextToken
//		if aws.StringValue(req.NextToken) == "" {
//			break
//		}
//	}
//	return resources, nil
//}
//

func GetRealmClient(ctx context.Context) (*realm.Client, error) {
	//p, err := profile.NewProfile(&req, profileName)
	//if err != nil {
	//	return nil, err
	//}

	optsRealm := []realm.ClientOpt{realm.SetUserAgent(terraformUserAgent)}
	authConfig := realmAuth.NewConfig(nil)
	token, err := authConfig.NewTokenFromCredentials(ctx, "..", "....")
	if err != nil {
		return nil, err
	}

	clientRealm := realmAuth.NewClient(realmAuth.BasicTokenSource(token))
	realmClient, err := realm.New(clientRealm, optsRealm...)
	if err != nil {
		return nil, err
	}

	return realmClient, nil
}
