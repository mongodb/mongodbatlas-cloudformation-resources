package util

import (
	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/logging"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/atlas/mongodbatlas"
	"os"
	"strings"
)

const (
	Version = "beta"
)

// This takes either "us-east-1" or "US_EAST_1"
// and returns "US_EAST_1" -- i.e. a valid Atlas region
func EnsureAtlasRegion(region string) string {
	r := strings.ToUpper(strings.Replace(string(region), "-", "_", -1))
	log.Printf("EnsureAtlasRegion--- region:%s r:%s", region, r)
	return r
}

// This takes either "us-east-1" or "US_EAST_1"
// and returns "us-east-1" -- i.e. a valid AWS region
func EnsureAWSRegion(region string) string {
	r := strings.ToLower(strings.Replace(string(region), "_", "-", -1))
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

	//Initialize the MongoDB Atlas API Client.
	atlas := mongodbatlas.NewClient(client)
	atlas.UserAgent = "mongodbatlas-cloudformation-resources/" + Version
	return atlas, nil
}

const EnvLogLevel = "LOG_LEVEL"

// defaultLogLevel can be set during compile time with an ld flag to enable
// more verbose logging.
// For example,
// env GOOS=$(goos) CGO_ENABLED=$(cgo) GOARCH=$(goarch) go build -ldflags="-s -w -X 'github.com/mongodb/mongodbatlas-cloudformation-resources/util.defaultLogLevel=debug'" -tags="$(tags)" -o bin/handler cmd/main.go
var defaultLogLevel = "info"

func getLogLevel() log.Level {
	levelString, exists := os.LookupEnv(EnvLogLevel)
	if !exists {
		log.Errorf("getLogLevel() Environment variable '%s' not found. Set it in template.yaml (defaultLogLevel=%v)", EnvLogLevel, defaultLogLevel)
		levelString = defaultLogLevel
	}

	level, err := log.ParseLevel(levelString)
	if err != nil {
		log.Errorf("error parsing %s: %v", EnvLogLevel, err)
		level, err = log.ParseLevel(defaultLogLevel)
		return level
	}
	log.Printf("getLogLevel() levelString=%s level=%v", levelString, level)
	return level
}

// SetupLogger is called by each resource handler to centrally
// configure the log level and properly connect to the cfn
// cloudwatch writer
func SetupLogger(loggerPrefix string) {
	logger := logging.New(loggerPrefix)
	log.SetOutput(logger.Writer())
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(getLogLevel())
	log.Info("INFO setLogger")
	log.Debug("DEBUG setLogger")
}
