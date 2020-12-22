package util

import (
	"go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"context"
	"log"
	"testing"
	//"fmt"
	"flag"
	"os"
)

var bgCtx = context.Background()

const (
	publicKeyEnv  = "ATLAS_PUBLIC_KEY"
	privateKeyEnv = "ATLAS_PRIVATE_KEY"
	//orgIDEnv      = "ATLAS_ORG_ID"
)

var (
	publicKey  = os.Getenv(publicKeyEnv)
	privateKey = os.Getenv(privateKeyEnv)
	//orgID      = os.Getenv(orgIDEnv)
)

func setupAtlasClient() (*mongodbatlas.Client, error) {
	client, err := CreateMongoDBClient(publicKey, privateKey)
	if err != nil {
		return nil, err
	}
	return client, nil
}
func testSetupMongoDBClient(opts ...*options.ClientOptions) (*mongo.Client, error) {
	if len(opts) == 0 {
		opts = append(opts, options.Client().ApplyURI("mongodb://localhost:27017"))
	}
	return mongo.NewClient(opts...)
}

var projectID string
var clusterName string

func init() {
	flag.StringVar(&projectID, "projectID", "", "Atlas Project ID")
	flag.StringVar(&clusterName, "clusterName", "", "Cluster Name")
}
func TestMongo(t *testing.T) {
	log.Println("mongodb_test log start")
	log.Printf("projectID=%v clusterName=%v", projectID, clusterName)
	flag.Parse()
	t.Run("test test", func(t *testing.T) {
		atlasClient, err := setupAtlasClient()
		if err != nil {
			panic(err)
		}
		log.Printf("Did it work? atlasClient: %#+v", atlasClient)
		dbs, err := ListDatabaseNames(atlasClient, &projectID, &clusterName)
		if err != nil {
			panic(err)
		}
		log.Printf("dbs:%+v", dbs)

	})

}
