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

package resource_test

import (
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/project/cmd/resource"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func TestNewResourceTags(t *testing.T) {
	testCases := map[string]struct {
		input  map[string]string
		output []admin.ResourceTag
	}{
		"empty":    {map[string]string{}, []admin.ResourceTag{}},
		"single":   {map[string]string{"key": "value"}, []admin.ResourceTag{*admin.NewResourceTag("key", "value")}},
		"multiple": {map[string]string{"k1": "v1", "k2": "v2"}, []admin.ResourceTag{*admin.NewResourceTag("k1", "v1"), *admin.NewResourceTag("k2", "v2")}},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.output, resource.NewResourceTags(tc.input))
			assert.Equal(t, tc.input, resource.NewCfnTags(tc.output))
		})
	}
}
