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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetTypeNamesFromRpdkJSONFiles() []string {
	var typeNames []string

	_ = filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == ".rpdk-config" {
			typeName, err := extractTypeNameFromFile(path)
			if err != nil {
				fmt.Printf("Error extracting typeName from %s: %v\n", path, err)
				return nil
			}
			typeNames = append(typeNames, typeName)
		}

		return nil
	})

	return typeNames
}

func extractTypeNameFromFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	var config map[string]interface{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return "", err
	}

	typeName, ok := config["typeName"].(string)
	if !ok {
		return "", fmt.Errorf("typeName not found or not a string in %s", filePath)
	}

	return typeName, nil
}
