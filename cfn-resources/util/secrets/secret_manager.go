// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package secrets

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/awsconfig"
)

func Create(req *handler.Request, secretName string, data interface{}, description *string) (name *string, arn *string, err error) {
	secretString, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	// Create service client value configured for credentials
	// from assumed role.
	cfg := awsconfig.FromHandlerRequest(req)
	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.CreateSecretInput{
		Description:  description,
		Name:         aws.String(secretName),
		SecretString: aws.String(string(secretString)),
	}

	result, err := svc.CreateSecret(context.Background(), input)
	if err != nil {
		// Print the error, cast err to awserr. Error to get the Code and
		// Message from an error.
		log.Printf("error create secret: %+v", err.Error())
		return nil, nil, err
	}
	log.Printf("Created secret result:%+v", result)
	return result.Name, result.ARN, nil
}

func PutSecret(req *handler.Request, secretName string, data interface{}, description *string) (name *string, arn *string, err error) {
	secretString, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	// Create service client value configured for credentials
	// from assumed role.
	cfg := awsconfig.FromHandlerRequest(req)
	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.PutSecretValueInput{
		SecretId:     aws.String(secretName),
		SecretString: aws.String(string(secretString)),
	}

	result, err := svc.PutSecretValue(context.Background(), input)
	if err != nil {
		// Print the error, cast err to awserr. Error to get the Code and
		// Message from an error.
		log.Printf("error during put secret: %+v", err.Error())
		return nil, nil, err
	}
	log.Printf("Created secret result:%+v", result)
	return result.Name, result.ARN, nil
}

func Get(req *handler.Request, secretName string) (name *string, arn *string, err error) {
	cfg := awsconfig.FromHandlerRequest(req)
	sm := secretsmanager.NewFromConfig(cfg)

	output, err := sm.GetSecretValue(context.Background(), &secretsmanager.GetSecretValueInput{SecretId: aws.String(secretName)})
	if err != nil {
		log.Printf("Error --- %v", err.Error())
		return nil, nil, err
	}

	return output.SecretString, output.ARN, nil
}

func Delete(req *handler.Request, secretName string) (err error) {
	cfg := awsconfig.FromHandlerRequest(req)
	sm := secretsmanager.NewFromConfig(cfg)

	_, err = sm.DeleteSecret(context.Background(), &secretsmanager.DeleteSecretInput{
		SecretId:                   aws.String(secretName),
		ForceDeleteWithoutRecovery: util.Pointer(true),
	})
	if err != nil {
		log.Printf("error delete secret: %v", err.Error())
		return err
	}
	return nil
}
