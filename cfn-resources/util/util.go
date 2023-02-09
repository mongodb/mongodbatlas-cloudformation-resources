package util

import (
	"encoding/json"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	cfnlog "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"log"
	"os"
	"strings"

	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/logging"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
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

	// Initialize the MongoDB Atlas API Client.
	atlas := mongodbatlas.NewClient(client)
	atlas.UserAgent = "mongodbatlas-cloudformation-resources/" + Version
	return atlas, nil
}

func NewMongoDBClient(req handler.Request) (*mongodbatlas.Client, *handler.ProgressEvent) {
	keys, handlerError := getApiKeys(req)
	if handlerError != nil {
		return nil, handlerError
	}

	// Create atlas client
	client, err := CreateMongoDBClient(keys.PublicKey, keys.PrivateKey)
	if err != nil {
		_, _ = cfnlog.Warnf("Create - error: %+v", err)
		peErr := progress_events.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest)
		return nil, &peErr
	}

	return client, nil
}

func getApiKeys(req handler.Request) (*DeploymentSecret, *handler.ProgressEvent) {
	key, err := GetAPIKeyFromDeploymentSecret(&req, "Atlas-Cloud-Formation-ApiKey-Secret")
	if err != nil {
		_, _ = cfnlog.Warnf("Read - error: %+v", err)
		pe := handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          fmt.Sprintf("Error getting api keyd secrets, the apikeys needs to be provided using aws secret, remember to validate if a secret named Atlas-Cloud-Formation-ApiKey-Secret is created with the PublicKey and PrivateKey properties, error: %s", err.Error()),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}
		return nil, &pe
	}

	return &key, nil
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
