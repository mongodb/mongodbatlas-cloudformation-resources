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

package main

import (
	"log"
	"os"

	"github.com/nsf/jsondiff"
)

// CompareJsonFiles Compares the JSON Content in given Files
func CompareJSONFiles(resourceName, existingFilePath, latestFilePath string) (diffJSON string, err error) {
	existingAPIContent, err := os.ReadFile(existingFilePath)
	if err != nil {
		return
	}

	latestAPIContent, err := os.ReadFile(latestFilePath)
	if err != nil {
		return
	}

	differences, diffJSON := jsondiff.Compare(existingAPIContent, latestAPIContent, &jsondiff.Options{SkipMatches: true})
	if differences > 0 {
		log.Printf("found difference in %s schema ", resourceName)
		err = os.WriteFile(resourceName+"-"+diffFile, []byte(diffJSON), 0600)
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}
