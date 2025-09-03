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
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"
	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"
	"go.mongodb.org/atlas-sdk/v20250312006/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/logging"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/mongodb-forks/digest"
	"github.com/mongodb-labs/go-client-mongodb-atlas-app-services/appservices"
	appServicesAuth "github.com/mongodb-labs/go-client-mongodb-atlas-app-services/auth"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/version"
)

const (
	cfn         = "mongodbatlas-cloudformation-resources"
	envLogLevel = "LOG_LEVEL"
	debug       = "debug"
)

type MongoDBClient struct {
	Atlas20231115002 *admin20231115002.APIClient
	Atlas20231115014 *admin20231115014.APIClient
	AtlasSDK         *admin.APIClient
	Config           *Config
}

type Config struct {
	AssumeRole         *AssumeRole
	PublicKey          string
	PrivateKey         string
	BaseURL            string
	AppServicesBaseURL string
	DebugClient        bool
}

type AssumeRole struct {
	Tags              map[string]string
	RoleARN           string
	ExternalID        string
	Policy            string
	SessionName       string
	SourceIdentity    string
	PolicyARNs        []string
	TransitiveTagKeys []string
	Duration          time.Duration
}

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

func GetAppServicesClient(ctx context.Context, req handler.Request, profileName *string) (*appservices.Client, error) {
	p, err := profile.NewProfile(&req, profileName, true)
	if err != nil {
		return nil, err
	}

	optsAppServices := []appservices.ClientOpt{appservices.SetUserAgent(userAgent)}
	authConfig := appServicesAuth.NewConfig(nil)
	token, err := authConfig.NewTokenFromCredentials(ctx, p.PublicKey, p.PrivateKey)
	if err != nil {
		return nil, err
	}

	clientAppServices := appServicesAuth.NewClient(appServicesAuth.BasicTokenSource(token))
	appServicesClient, err := appservices.New(clientAppServices, optsAppServices...)
	if err != nil {
		return nil, err
	}

	return appServicesClient, nil
}

func NewAtlasClient(req *handler.Request, profileName *string) (*MongoDBClient, *handler.ProgressEvent) {
	return newAtlasV2Client(req, profileName, true)
}

// NewAtlasClientRemovingProfilePrefix creates a client using the provided profileName but without adding the common prefix (`cfn/atlas/profile`)
// Only used for specific use case in organization resource.
func NewAtlasClientRemovingProfilePrefix(req *handler.Request, profileName *string) (*MongoDBClient, *handler.ProgressEvent) {
	return newAtlasV2Client(req, profileName, false)
}

func newAtlasV2Client(req *handler.Request, profileName *string, profileNamePrefixRequired bool) (*MongoDBClient, *handler.ProgressEvent) {
	prof, err := profile.NewProfile(req, profileName, profileNamePrefixRequired)

	if err != nil {
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}
	}

	// setup a transport to handle digest
	transport := digest.NewTransport(prof.PublicKey, prof.PrivateKey)

	// initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}
	}

	c := Config{BaseURL: prof.BaseURL, DebugClient: prof.UseDebug()}

	// new V2 version 20231115002 instance
	sdk20231115002Client, err := c.NewSDKv20231115002Client(client)
	if err != nil {
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}
	}

	// new V2 version 20231115014 instance
	sdk20231115014Client, err := c.NewSDKv20231115014Client(client)
	if err != nil {
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}
	}

	// latest V2 instance
	sdkV2LatestClient, err := c.NewSDKV2LatestClient(client)
	if err != nil {
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}
	}

	clients := &MongoDBClient{
		Atlas20231115002: sdk20231115002Client,
		Atlas20231115014: sdk20231115014Client,
		AtlasSDK:         sdkV2LatestClient,
		Config:           &c,
	}

	return clients, nil
}

func (c *Config) NewSDKv20231115002Client(client *http.Client) (*admin20231115002.APIClient, error) {
	opts := []admin20231115002.ClientModifier{
		admin20231115002.UseHTTPClient(client),
		admin20231115002.UseUserAgent(userAgent),
		admin20231115002.UseBaseURL(c.BaseURL),
		admin20231115002.UseDebug(c.DebugClient)}

	sdkV2, err := admin20231115002.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return sdkV2, nil
}

func (c *Config) NewSDKv20231115014Client(client *http.Client) (*admin20231115014.APIClient, error) {
	opts := []admin20231115014.ClientModifier{
		admin20231115014.UseHTTPClient(client),
		admin20231115014.UseUserAgent(userAgent),
		admin20231115014.UseBaseURL(c.BaseURL),
		admin20231115014.UseDebug(c.DebugClient)}

	sdkV2, err := admin20231115014.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return sdkV2, nil
}

func (c *Config) NewSDKV2LatestClient(client *http.Client) (*admin.APIClient, error) {
	opts := []admin.ClientModifier{
		admin.UseHTTPClient(client),
		admin.UseUserAgent(userAgent),
		admin.UseBaseURL(c.BaseURL),
		admin.UseDebug(c.DebugClient)}

	// Initialize the MongoDB Versioned Atlas Client.
	sdkV2, err := admin.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return sdkV2, nil
}

// defaultLogLevel can be set during compile time with an ld flag to enable
// more verbose logging.
// For example,
// env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="-s -w -X \
// 'github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug'" -tags="$(tags)" -o bin/bootstrap cmd/main.go
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

// Contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func SafeString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func SafeInt(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}

// TimePtrToStringPtr utility conversions that can potentially be defined in sdk
func TimePtrToStringPtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	res := TimeToString(*t)
	return &res
}

// TimeToString returns a RFC3339 (nano) date time string format.
// The resulting format is identical to the format returned by Atlas API, documented as ISO 8601 timestamp format in UTC.
// It also returns decimals in seconds (up to nanoseconds) if available.
// Example formats: "2023-07-18T16:12:23Z", "2023-07-18T16:12:23.456Z"
func TimeToString(t time.Time) string {
	return t.UTC().Format(time.RFC3339Nano)
}

// StringPtrToTimePtr parses a string with RFC3339 (nano) format and returns time.Time.
// It's the opposite function to TimeToString.
// Returns nil if date can't be parsed.
func StringPtrToTimePtr(p *string) *time.Time {
	if !IsStringPresent(p) {
		return nil
	}
	t, err := time.Parse(time.RFC3339Nano, *p)
	if err != nil {
		return nil
	}
	t = t.UTC()
	return &t
}

func StringToTime(t string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, t)
}

func IntPtr(i int) *int {
	return &i
}

func Int64PtrToIntPtr(i64 *int64) *int {
	if i64 == nil {
		return nil
	}

	i := int(*i64)
	return &i
}

func IsStringPresent(strPtr *string) bool {
	return strPtr != nil && *strPtr != ""
}

func AreStringPtrEqual(p1, p2 *string) bool {
	if p1 == nil {
		return p2 == nil
	}
	if p2 == nil {
		return false
	}
	return *p1 == *p2
}

// setDefaultProfileIfNotDefined can be called at the beginning of the CRUDL methods to set default profile if not defined
func SetDefaultProfileIfNotDefined(p **string) {
	if p != nil && !IsStringPresent(*p) {
		*p = aws.String(profile.DefaultProfile)
	}
}

func StrPtrToIntPtr(str *string) *int {
	if !IsStringPresent(str) {
		return nil
	}
	if val, err := strconv.Atoi(*str); err == nil {
		return &val
	}
	return nil
}

func IntPtrToStrPtr(i *int) *string {
	if i == nil {
		return nil
	}
	str := strconv.Itoa(*i)
	return &str
}

func StringPtr(v string) *string {
	if v != "" {
		return &v
	}
	return nil
}

func SameStringSliceWithoutOrder(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y]--
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}
