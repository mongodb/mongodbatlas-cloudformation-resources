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

package utility

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"os/exec"
	"testing"

	cfn "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"go.mongodb.org/atlas/mongodbatlas"
)

func GetRandNum() *big.Int {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(10000))
	return randInt
}

func FailNowIfError(t *testing.T, msgf string, err error) {
	if err != nil {
		t.Errorf(msgf, err.Error())
		t.FailNow()
	}
}

func RunCleanupScript(t *testing.T, rctx ResourceContext) {
	t.Setenv("RESOURCE_TYPE_NAME_FOR_E2E", rctx.ResourceTypeNameForE2e)
	t.Setenv("RESOURCE_DIRECTORY_NAME", rctx.ResourceDirectory)

	output, err := runShScript(t, "../utility/cleanup_cfn.sh")
	FailNowIfError(t, fmt.Sprintf("Error when executing cleanup script. Output: %s\n", output)+"%v", err)

	if err != nil {
		return
	}
	t.Logf("E2E test resource type %s successfully de-registered from private registry!\n", rctx.ResourceTypeNameForE2e)
}

func PublishToPrivateRegistry(t *testing.T, rctx ResourceContext) {
	t.Setenv("RESOURCE_TYPE_NAME", rctx.ResourceTypeName)
	t.Setenv("RESOURCE_TYPE_NAME_FOR_E2E", rctx.ResourceTypeNameForE2e)
	t.Setenv("E2E_RAND_SUFFIX", rctx.E2eRandSuffix)
	t.Setenv("RESOURCE_DIRECTORY_NAME", rctx.ResourceDirectory)
	t.Cleanup(func() {
		RunCleanupScript(t, rctx)
	})

	output, err := runShScript(t, "../utility/publish_cfn_to_registry.sh")
	FailNowIfError(t, fmt.Sprintf("Error when executing publishing script. Output: %s\n", output)+"%v", err)
	if err != nil {
		return
	}
	t.Logf("New E2E test resource type %s successfully published to private registry!\n", rctx.ResourceTypeNameForE2e)
}

func runShScript(t *testing.T, path string) ([]byte, error) {
	cmd := exec.Command(path)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	// Read the output from pipes
	output, _ := readOutput(stdout)
	errorOutput, _ := readOutput(stderr)

	err = cmd.Wait()

	t.Logf("runShScript Output: %v", string(output))
	t.Logf("runShScript Error: %v", string(errorOutput))

	return output, err
}

func readOutput(pipe io.Reader) ([]byte, error) {
	output, err := io.ReadAll(pipe)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func NewClients(t *testing.T) (cfnClient *cfn.Client, atlasClient *mongodbatlas.Client) {
	t.Helper()

	t.Log("Setting clients")
	atlasClient, err := NewMongoDBClient()
	FailNowIfError(t, "Unable to create atlas client: %v", err)

	cfnClient, err = NewCFNClient()
	FailNowIfError(t, "Unable to create AWS client, please check AWS config is correctly setup: %v", err)

	return cfnClient, atlasClient
}
