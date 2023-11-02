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

package resource

import (
	"reflect"
	"testing"
)

func TestConvertToAnySliceEmptyList(t *testing.T) {
	var input []string
	var expected []interface{}

	result, err := convertToAnySlice(input)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result does not match the expected value. Got: %v, Expected: %v", result, expected)
	}
}

func TestConvertToAnySliceValidJSON(t *testing.T) {
	input := []string{
		`{"type": "icuNormalize"}`,
	}
	expected := []interface{}{
		map[string]interface{}{"type": "icuNormalize"},
	}

	result, err := convertToAnySlice(input)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result does not match the expected value. Got: %v, Expected: %v", result, expected)
	}
}

func TestConvertToAnySliceInvalidJSON(t *testing.T) {
	input := []string{
		`this is an invalid json`,
	}
	var expected []interface{}

	result, err := convertToAnySlice(input)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result does not match the expected value. Got: %v, Expected: %v", result, expected)
	}
}
