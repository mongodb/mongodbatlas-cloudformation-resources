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
	"crypto/rand"
	"log"
	"math/big"
	"os"
	"os/exec"
)

func GetRandNum() *big.Int {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(10000))
	return randInt
}

func RunCleanupScript(rctx ResourceContext) {
	output, err := exec.Command("/bin/sh", "../cleanup_cfn.sh").CombinedOutput()
	if err != nil {
		log.Println("Error when running cleanup script. Output:")
		log.Println(string(output))
		log.Printf("Got command status: %s\n", err.Error())
		return
	}
	log.Printf("E2E test resource type %s successfully de-registered from private registry!\n", rctx.ResourceTypeNameForE2e)
}

func PublishToPrivateRegistry(rctx ResourceContext) {
	os.Setenv("RESOURCE_TYPE_NAME", rctx.ResourceTypeName)
	os.Setenv("RESOURCE_TYPE_NAME_FOR_E2E", rctx.ResourceTypeNameForE2e)
	os.Setenv("E2E_RAND_SUFFIX", rctx.E2eRandSuffix)
	os.Setenv("RESOURCE_DIRECTORY_NAME", rctx.ResourceDirectory)

	// create a new resourceType and submit to private registry
	output, err := exec.Command("/bin/sh", "../publish_cfn_to_registry.sh").CombinedOutput()
	if err != nil {
		log.Println("Error when running command. Output:")
		log.Println(string(output))
		log.Printf("Got command status: %s\n", err.Error())
		return
	}
	log.Printf("New E2E test resource type %s successfully published to private registry!\n", rctx.ResourceTypeNameForE2e)
}
