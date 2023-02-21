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

package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"

	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/logging"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/version"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	cfn            = "mongodbatlas-cloudformation-resources"
	DefaultProfile = "default"
	envLogLevel    = "LOG_LEVEL"
	debug          = "debug"
	profileName    = "cfn/atlas/profile/%s"
)

var (
	toolName        = cfn
	defaultLogLevel = "warning"
	userAgent       = fmt.Sprintf("%s/%s (%s;%s)", toolName, version.Version, runtime.GOOS, runtime.GOARCH)
)

// EnsureAtlasRegion This takes either "us-east-1" or "US_EAST_1"
// and returns "US_EAST_1" -- i.e. a valid Atlas region
func EnsureAtlasRegion(region string) string {
	r := strings.ToUpper(strings.ReplaceAll(region, "-", "_"))
	log.Printf("EnsureAtlasRegion--- region:%s r:%s", region, r)
	return r
}

// EnsureAWSRegion This takes either "us-east-1" or "US_EAST_1"
// and returns "us-east-1" -- i.e. a valid AWS region
func EnsureAWSRegion(region string) string {
	r := strings.ToLower(strings.ReplaceAll(region, "_", "-"))
	log.Printf("EnsureAWSRegion--- region:%s r:%s", region, r)
	return r
}

// CreateMongoDBClient creates a new Client using apikeys
//
// Deprecated: In the future this function will be private, the NewMongoDBClient should be used instead.
func CreateMongoDBClient(publicKey, privateKey string) (*mongodbatlas.Client, error) {
	// setup a transport to handle digest
	log.Printf("CreateMongoDBClient--- publicKey:%s", publicKey)
	transport := digest.NewTransport(publicKey, privateKey)

	// initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, err
	}

	opts := []mongodbatlas.ClientOpt{mongodbatlas.SetUserAgent(userAgent)}
	if baseURL := os.Getenv("MONGODB_ATLAS_OPS_MANAGER_URL"); baseURL != "" {
		opts = append(opts, mongodbatlas.SetBaseURL(baseURL))
	}

	return mongodbatlas.New(client, opts...)
}

func NewMongoDBClient(req handler.Request, profile *string) (*mongodbatlas.Client, *handler.ProgressEvent) {
	profileInput := DefaultProfile
	if profile != nil {
		profileInput = *profile
	}

	print("ANDREA2")
	print("\n")
	print(profileInput)

	keys, handlerError := getAPIKeys(req, profileInput)
	if handlerError != nil {
		return nil, handlerError
	}

	// Create atlas client
	client, err := CreateMongoDBClient(keys.PublicKey, keys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf("Create - error: %+v", err)
		peErr := progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest)
		return nil, &peErr
	}

	return client, nil
}

func getAPIKeys(req handler.Request, profile string) (*DeploymentSecret, *handler.ProgressEvent) {
	key, err := GetAPIKeyFromDeploymentSecret(&req, fmt.Sprintf(profileName, profile))
	if err != nil {
		_, _ = logger.Warnf("Read - error: %+v", err)
		pe := handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message: fmt.Sprintf("Error getting API-key, API-key needs to be provided using an AWS secret,"+
				"  Ensure a secret named 'cfn/atlas/profile/%s' is created with the 'PublicKey' and 'PrivateKey' properties, error: %s",
				profile, err.Error()),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}
		return nil, &pe
	}

	return &key, nil
}

// defaultLogLevel can be set during compile time with an ld flag to enable
// more verbose logging.
// For example,
// env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="-s -w -X \
// 'github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug'" -tags="$(tags)" -o bin/handler cmd/main.go
func getLogLevel() logger.Level {
	levelString, exists := os.LookupEnv(envLogLevel)
	if !exists {
		_, _ = logger.Warnf("getLogLevel() Environment variable %s not found. Set it in template.yaml (defaultLogLevel=%s)", envLogLevel, defaultLogLevel)
		levelString = defaultLogLevel
	}
	switch levelString {
	case debug:
		return logger.DebugLevel
	default:
		return logger.WarningLevel
	}
}

// SetupLogger is called by each resource handler to centrally
// configure the logger level and properly connect to the cfn
// cloudwatch writer
func SetupLogger(loggerPrefix string) {
	logr := logging.New(loggerPrefix)
	logger.SetOutput(logr.Writer())
	logger.SetLevel(getLogLevel())
}

func ToStringMapE(ep any) (map[string]any, error) {
	var eMap map[string]any
	inrec, err := json.Marshal(ep)
	if err != nil {
		return eMap, err
	}
	err = json.Unmarshal(inrec, &eMap)
	if err != nil {
		return eMap, err
	}
	return eMap, nil
}

func CreateSSManagerClient(curSession *session.Session) (*ssm.SSM, error) {
	ssmCli := ssm.New(curSession)
	return ssmCli, nil
}
func PutKey(keyID, keyValue, prefix string, curSession *session.Session) (*ssm.PutParameterOutput, error) {
	ssmClient, err := CreateSSManagerClient(curSession)
	if err != nil {
		return nil, err
	}
	// transform api keys to json string
	parameterName := buildKey(keyID, prefix)
	parameterType := "SecureString"
	overwrite := true
	putParamOutput, err := ssmClient.PutParameter(&ssm.PutParameterInput{Name: &parameterName, Value: &keyValue, Type: &parameterType, Overwrite: &overwrite})
	if err != nil {
		return nil, err
	}

	return putParamOutput, nil
}

func DeleteKey(keyID, prefix string, curSession *session.Session) (*ssm.DeleteParameterOutput, error) {
	ssmClient, err := CreateSSManagerClient(curSession)
	if err != nil {
		return nil, err
	}
	parameterName := buildKey(keyID, prefix)

	deleteParamOutput, err := ssmClient.DeleteParameter(&ssm.DeleteParameterInput{Name: &parameterName})
	if err != nil {
		return nil, err
	}

	return deleteParamOutput, nil
}

func Get(keyID, prefix string, curSession *session.Session) string {
	ssmClient, err := CreateSSManagerClient(curSession)
	if err != nil {
		return ""
	}
	parameterName := buildKey(keyID, prefix)
	decrypt := true
	getParamOutput, err := ssmClient.GetParameter(&ssm.GetParameterInput{Name: &parameterName, WithDecryption: &decrypt})
	if err != nil {
		return ""
	}
	return *getParamOutput.Parameter.Value
}
func buildKey(keyID, storePrefix string) string {
	// this is strictly coupled with permissions for handlers, changing this means changing permissions in handler
	// moreover changing this might cause pollution in parameter store -  be sure you know what you are doing
	return fmt.Sprintf("%s-%s", storePrefix, keyID)
}
