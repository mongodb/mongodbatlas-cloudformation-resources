package util

import (
	"go.mongodb.org/atlas/mongodbatlas"

	"log"
	"testing"
	//"fmt"
	"flag"
	"os"
)

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

func TestMongo(t *testing.T) {
	log.Println("mongodb_test log start")
	flag.Parse()
	t.Run("test test", func(t *testing.T) {
		atlasClient, err := setupAtlasClient()
		if err != nil {
			t.Errorf("error should be nill, got = %v", err.Error())
		}
		if atlasClient == nil {
			t.Error("atlas client is not expected to be null")
		}
	})

}
