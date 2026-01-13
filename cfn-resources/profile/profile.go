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

package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/awsconfig"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
)

const (
	DefaultProfile = "default"
)

type Profile struct {
	DebugClient *bool  `json:"DebugClient,omitempty"`
	PublicKey   string `json:"PublicKey"`
	PrivateKey  string `json:"PrivateKey"`
	BaseURL     string `json:"BaseUrl,omitempty"`
}

func NewProfile(req *handler.Request, profileName *string, prefixRequired bool) (*Profile, error) {
	if profileName == nil || *profileName == "" {
		profileName = aws.String(DefaultProfile)
	}

	// When migrating to AWS SDK v2, we can't use config.LoadDefaultConfig() directly in CloudFormation resource handlers.
	// The cloudformation-cli-go-plugin provides credentials via handler.Request.Session, which is an AWS SDK v1 session.
	// These credentials have the permissions defined in our resource execution roles (e.g., Secrets Manager access).
	// Using LoadDefaultConfig() would use the Lambda's base execution role instead, which lacks these permissions.
	// See: https://github.com/aws-cloudformation/cloudformation-cli-go-plugin/issues/237
	cfg := awsconfig.FromHandlerRequest(req)
	secretsManagerClient := secretsmanager.NewFromConfig(cfg)
	secretID := *profileName
	if prefixRequired {
		secretID = SecretNameWithPrefix(*profileName)
	}
	resp, err := secretsManagerClient.GetSecretValue(context.Background(), &secretsmanager.GetSecretValueInput{SecretId: &secretID})
	if err != nil {
		return nil, err
	}

	profile := new(Profile)
	err = json.Unmarshal([]byte(*resp.SecretString), &profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (p *Profile) NewBaseURL() string {
	if baseURL := os.Getenv("MONGODB_ATLAS_BASE_URL"); baseURL != "" {
		return baseURL
	}

	return p.BaseURL
}

func (p *Profile) NewPublicKey() string {
	if k := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY"); k != "" {
		return k
	}

	return p.PublicKey
}

func (p *Profile) NewPrivateKey() string {
	if k := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY"); k != "" {
		return k
	}

	return p.PrivateKey
}

func (p *Profile) AreKeysAvailable() bool {
	return p.NewPublicKey() == "" || p.PrivateKey == ""
}

func (p *Profile) UseDebug() bool {
	if debug := os.Getenv("MONGODB_ATLAS_DEBUG"); debug != "" {
		return debug == "true"
	}
	if p.DebugClient != nil {
		return *p.DebugClient
	}
	return false
}

func SecretNameWithPrefix(name string) string {
	return fmt.Sprintf("%s/%s", constants.ProfileNamePrefix, name)
}
