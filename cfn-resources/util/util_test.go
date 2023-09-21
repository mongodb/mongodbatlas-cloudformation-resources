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

package util_test

import (
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

var (
	empty    = ""
	oneBlank = " "
	str      = "text"
)

func TestIsStringPresent(t *testing.T) {
	tests := []struct {
		strPtr   *string
		expected bool
	}{
		{nil, false},
		{&empty, false},
		{&oneBlank, true},
		{&str, true},
	}
	for _, test := range tests {
		if resp := util.IsStringPresent(test.strPtr); resp != test.expected {
			t.Errorf("IsStringPresent(%v) = %v; want %v", test.strPtr, resp, test.expected)
		}
	}
}

func TestAreStringPtrEqual(t *testing.T) {
	tests := []struct {
		strPtr1  *string
		strPtr2  *string
		expected bool
	}{
		{nil, &empty, false},
		{&empty, nil, false},
		{&oneBlank, &empty, false},
		{&empty, &oneBlank, false},
		{nil, nil, true},
		{&empty, &empty, true},
		{&oneBlank, &oneBlank, true},
		{&str, &str, true},
	}
	for _, test := range tests {
		if resp := util.AreStringPtrEqual(test.strPtr1, test.strPtr2); resp != test.expected {
			t.Errorf("AreStringPtrEqual(%v, %v) = %v; want %v", test.strPtr1, test.strPtr2, resp, test.expected)
		}
	}
}
