package util

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"log"
	"os"
	"strings"

	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/logging"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"go.mongodb.org/atlas/mongodbatlas"
)

const Version = "beta"

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

func CreateMongoDBClient(publicKey, privateKey string) (*mongodbatlas.Client, error) {
	// setup a transport to handle digest
	log.Printf("CreateMongoDBClient--- publicKey:%s", publicKey)
	transport := digest.NewTransport(publicKey, privateKey)

	// initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, err
	}

	// Initialize the MongoDB Atlas API Client.
	atlas := mongodbatlas.NewClient(client)
	atlas.UserAgent = "mongodbatlas-cloudformation-resources/" + Version
	return atlas, nil
}

const (
	EnvLogLevel = "LOG_LEVEL"
	Debug       = "debug"
)

var defaultLogLevel = "warning"

// defaultLogLevel can be set during compile time with an ld flag to enable
// more verbose logging.
// For example,
// env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="-s -w -X \
// 'github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug'" -tags="$(tags)" -o bin/handler cmd/main.go
func getLogLevel() logger.Level {
	levelString, exists := os.LookupEnv(EnvLogLevel)
	if !exists {
		_, _ = logger.Warnf("getLogLevel() Environment variable %s not found. Set it in template.yaml (defaultLogLevel=%s)", EnvLogLevel, defaultLogLevel)
		levelString = defaultLogLevel
	}
	switch levelString {
	case Debug:
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
func CreateSSMClient(session *session.Session) (*ssm.SSM, error) {
	ssmCli := ssm.New(session)
	return ssmCli, nil
}
func Put(keyID, keyValue, prefix string, session *session.Session) (*ssm.PutParameterOutput, error) {
	ssmClient, err := CreateSSMClient(session)
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

func Delete(keyID, prefix string, session *session.Session) (*ssm.DeleteParameterOutput, error) {
	ssmClient, err := CreateSSMClient(session)
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

func Get(keyID, prefix string, session *session.Session) string {
	ssmClient, err := CreateSSMClient(session)
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
	// moreover changing this might cause polution in parameter store -  be sure you know what you are doing
	return fmt.Sprintf("%s-%s", storePrefix, keyID)
}
