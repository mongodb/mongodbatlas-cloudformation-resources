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
	"sort"

	"go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func NewResourceTags(tags map[string]string) []admin.ResourceTag {
	sliceTags := make([]admin.ResourceTag, 0, len(tags))
	for _, k := range sortStringMapKeys(tags) {
		v := tags[k]
		tag := admin.NewResourceTag(k, v)
		sliceTags = append(sliceTags, *tag)
	}
	return sliceTags
}

func NewCfnTags(tags []admin.ResourceTag) map[string]string {
	mapTags := make(map[string]string, len(tags))
	for _, tag := range tags {
		mapTags[tag.Key] = tag.Value
	}
	return mapTags
}

// make test deterministic
func sortStringMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
