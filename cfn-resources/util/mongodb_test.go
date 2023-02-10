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
	"go.mongodb.org/atlas/mongodbatlas"

	"flag"
	"log"
	"os"
	"testing"
)

const (
	publicKeyEnv  = "ATLAS_PUBLIC_KEY"
	privateKeyEnv = "ATLAS_PRIVATE_KEY"
)

var (
	publicKey  = os.Getenv(publicKeyEnv)
	privateKey = os.Getenv(privateKeyEnv)
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
