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

package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

func main_sam() {
	// Create the sam command with the appropriate arguments
	cmd := exec.CommandContext(context.Background(), "sam", "local", "start-lambda", "--skip-pull-image")

	// Capture the output and error streams from the command
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Start the command and wait for it to finish
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}
	err = cmd.Wait()

	// Print the output and error streams from the command
	if stdout.Len() > 0 {
		fmt.Println("Command output:")
		fmt.Println(stdout.String())
	}
	if stderr.Len() > 0 {
		fmt.Println("Command error:")
		fmt.Println(stderr.String())
	}

	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}
}
