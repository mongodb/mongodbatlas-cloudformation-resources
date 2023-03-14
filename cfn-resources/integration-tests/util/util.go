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
	"os"
	"runtime"
	"time"

	"github.com/mongodb-forks/digest"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/version"
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	realmAuth "go.mongodb.org/realm/auth"
	"go.mongodb.org/realm/realm"
)

const (
	cfnTool = "mongodbatlas-cloudformation-resources"
)

var (
	toolName           = cfnTool
	defaultLogLevel    = "warning"
	userAgent          = fmt.Sprintf("%s/%s (%s;%s)", toolName, version.Version, runtime.GOOS, runtime.GOARCH)
	terraformUserAgent = "terraform-provider-mongodbatlas"
)

type AtlasEnvOptions struct {
	OrgId      string
	PrivateKey string
	PublicKey  string
	BaseURL    string
}

func CreateStack(client *cfn.Client, name string, fileContent []byte) (*cfn.DescribeStacksOutput, error) {
	templateBody := string(fileContent)
	outpt, err := createStackAndWait(client, name, templateBody)
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
		fmt.Println("Waiting for stack creation: ", resp.Stacks[0].StackStatus)
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
		fmt.Println("Waiting for stack deletion: ", resp.Stacks[0].StackStatus)
		time.Sleep(3 * time.Second)
	}
}

func UpdateStack(svc *cfn.Client, stackName, templateBody string) (*cfn.DescribeStacksOutput, error) {
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

func ValidateTemplate(svc *cfn.Client, template string) (string, error) {
	input := &cfn.ValidateTemplateInput{
		TemplateBody: aws.String(template),
	}

	_, err := svc.ValidateTemplate(context.Background(), input)

	if err != nil {
		return "", fmt.Errorf("Invalid cloudformation stack: %v", err)
	}

	return "Template is valid", err
}

func GetRealmClient(ctx context.Context) (*realm.Client, error) {
	atlasEnv, err := getAtlasEnv()
	if err != nil {
		return nil, err
	}
	optsRealm := []realm.ClientOpt{realm.SetUserAgent(terraformUserAgent)}
	authConfig := realmAuth.NewConfig(nil)
	token, err := authConfig.NewTokenFromCredentials(ctx, atlasEnv.PublicKey, atlasEnv.PrivateKey)
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

func NewMongoDBClient() (atlasClient *mongodbatlas.Client, err error) {
	atlasEnv, err := getAtlasEnv()
	if err != nil {
		return nil, err
	}
	t := digest.NewTransport(atlasEnv.PublicKey, atlasEnv.PrivateKey)
	client, err := t.Client()

	opts := []mongodbatlas.ClientOpt{mongodbatlas.SetUserAgent(userAgent)}
	if baseURL := atlasEnv.BaseURL; baseURL != "" {
		opts = append(opts, mongodbatlas.SetBaseURL(baseURL))
	}

	mongodbClient, err := mongodbatlas.New(client, opts...)
	if err != nil {
		return nil, errors.New("unable to create Atlas client")
	}

	return mongodbClient, nil
}

func getAtlasEnv() (atlasEnvOpts *AtlasEnvOptions, err error) {
	orgId1 := os.Getenv("ATLAS_ORG_ID")
	fmt.Println(orgId1)
	orgId, OrgIdOk := os.LookupEnv("ATLAS_ORG_ID")
	publicKey, publicKeyOk := os.LookupEnv("ATLAS_PUBLIC_KEY")
	privateKey, privateKeyOk := os.LookupEnv("ATLAS_PRIVATE_KEY")
	baseUrl, baseUrlOk := os.LookupEnv("ATLAS_BASE_URL")

	if !OrgIdOk || !privateKeyOk || !publicKeyOk || !baseUrlOk {
		return nil, errors.New("please ensure following env variables are set: ATLAS_ORG_ID, ATLAS_PUBLIC_KEY, ATLAS_PRIVATE_KEY, ATLAS_BASE_URL")
	}

	return &AtlasEnvOptions{orgId, privateKey, publicKey, baseUrl}, nil
}

func GetAWSClient() (client *cfn.Client, err error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, errors.New("Error loading AWS configuration: " + err.Error())
	}
	fmt.Println("Loaded AWS configuration")
	return cfn.NewFromConfig(cfg), nil
}
