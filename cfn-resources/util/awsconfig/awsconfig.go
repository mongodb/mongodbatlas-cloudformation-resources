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

// Package awsconfig provides utilities for creating AWS SDK v2 configurations
// from CloudFormation handler requests.
package awsconfig

import (
	"context"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/aws"
)

// FromHandlerRequest creates an AWS SDK v2 config using the CloudFormation handler's SDK v1 session.
//
// The cloudformation-cli-go-plugin provides credentials via handler.Request.Session,
// which is an AWS SDK v1 session. This function bridges those credentials to SDK v2
// using a provider that fetches credentials on each AWS API call.
func FromHandlerRequest(req *handler.Request) aws.Config {
	return aws.Config{
		Region: aws.ToString(req.Session.Config.Region),
		Credentials: aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			v1Creds, err := req.Session.Config.Credentials.Get()
			if err != nil {
				return aws.Credentials{}, fmt.Errorf("failed to get credentials from CloudFormation handler session: %w", err)
			}
			return aws.Credentials{
				AccessKeyID:     v1Creds.AccessKeyID,
				SecretAccessKey: v1Creds.SecretAccessKey,
				SessionToken:    v1Creds.SessionToken,
			}, nil
		}),
	}
}
