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
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/logging"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/mongodb-forks/digest"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/version"
	atlasSDK "go.mongodb.org/atlas-sdk/v20230201002/admin"
	"go.mongodb.org/atlas/mongodbatlas"
	realmAuth "go.mongodb.org/realm/auth"
	"go.mongodb.org/realm/realm"
)

const (
	cfn         = "mongodbatlas-cloudformation-resources"
	envLogLevel = "LOG_LEVEL"
	debug       = "debug"
)

type MongoDBClient struct {
	Atlas   *mongodbatlas.Client
	AtlasV2 *atlasSDK.APIClient
}

var (
	toolName           = cfn
	defaultLogLevel    = "warning"
	userAgent          = fmt.Sprintf("%s/%s (%s;%s)", toolName, version.Version, runtime.GOOS, runtime.GOARCH)
	terraformUserAgent = "terraform-provider-mongodbatlas"
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

func GetRealmClient(ctx context.Context, req handler.Request, profileName *string) (*realm.Client, error) {
	p, err := profile.NewProfile(&req, profileName)
	if err != nil {
		return nil, err
	}

	optsRealm := []realm.ClientOpt{realm.SetUserAgent(terraformUserAgent)}
	authConfig := realmAuth.NewConfig(nil)
	token, err := authConfig.NewTokenFromCredentials(ctx, p.PublicKey, p.PrivateKey)
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

// createMongoDBClient creates a new Client using apikeys
func createMongoDBClient(publicKey, privateKey string) (*mongodbatlas.Client, error) {
	// setup a transport to handle digest
	log.Printf("CreateMongoDBClient--- publicKey:%s", publicKey)
	transport := digest.NewTransport(publicKey, privateKey)

	// initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, err
	}

	opts := []mongodbatlas.ClientOpt{mongodbatlas.SetUserAgent(userAgent)}
	if baseURL := os.Getenv("MONGODB_ATLAS_BASE_URL"); baseURL != "" {
		opts = append(opts, mongodbatlas.SetBaseURL(baseURL))
	}

	return mongodbatlas.New(client, opts...)
}

func NewMongoDBClient(req handler.Request, profileName *string) (*mongodbatlas.Client, *handler.ProgressEvent) {
	p, err := profile.NewProfile(&req, profileName)
	if err != nil {
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}
	}

	client, err := newHTTPClient(p)
	if err != nil {
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}
	}

	opts := []mongodbatlas.ClientOpt{mongodbatlas.SetUserAgent(userAgent)}
	if baseURL := p.NewBaseURL(); baseURL != "" {
		opts = append(opts, mongodbatlas.SetBaseURL(baseURL))
	}

	mongodbClient, err := mongodbatlas.New(client, opts...)
	if err != nil {
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}
	}

	return mongodbClient, nil
}

// NewAtlasClient func for creating atlas-go-sdk and mongodb-atlas-go client
func NewAtlasClient(req *handler.Request, profileName *string) (*MongoDBClient, *handler.ProgressEvent) {
	prof, err := profile.NewProfile(req, profileName)

	if err != nil {
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}
	}

	// setup a transport to handle digest
	transport := digest.NewTransport(prof.PublicKey, prof.PublicKey)

	// initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}
	}

	optsAtlas := []mongodbatlas.ClientOpt{mongodbatlas.SetUserAgent(userAgent)}
	if prof.BaseURL != "" {
		optsAtlas = append(optsAtlas, mongodbatlas.SetBaseURL(prof.BaseURL))
	}

	// Initialize the MongoDB Atlas API Client.
	atlasClient, err := mongodbatlas.New(client, optsAtlas...)
	if err != nil {
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}
	}

	// New SDK Client
	sdkV2Client, err := newSDKV2Client(client)
	if err != nil {
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}
	}

	clients := &MongoDBClient{
		Atlas:   atlasClient,
		AtlasV2: sdkV2Client,
	}

	return clients, nil
}

func newSDKV2Client(client *http.Client) (*atlasSDK.APIClient, error) {
	opts := []atlasSDK.ClientModifier{
		atlasSDK.UseHTTPClient(client),
		atlasSDK.UseUserAgent(userAgent),
		atlasSDK.UseDebug(false)}

	// Initialize the MongoDB Versioned Atlas Client.
	sdkV2, err := atlasSDK.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return sdkV2, nil
}

func newHTTPClient(p *profile.Profile) (*http.Client, error) {
	if p.AreKeysAvailable() {
		return nil, errors.New("PublicKey and PrivateKey cannot be empty")
	}

	t := digest.NewTransport(p.NewPublicKey(), p.NewPrivateKey())
	return t.Client()
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
	print("ANDREA util.get\n")
	print(*getParamOutput.Parameter.Value)
	print("\n")
	return *getParamOutput.Parameter.Value
}

func Pointer[T any](x T) *T {
	return &x
}

func buildKey(keyID, storePrefix string) string {
	// this is strictly coupled with permissions for handlers, changing this means changing permissions in handler
	// moreover changing this might cause pollution in parameter store -  be sure you know what you are doing
	return fmt.Sprintf("%s-%s", storePrefix, keyID)
}
