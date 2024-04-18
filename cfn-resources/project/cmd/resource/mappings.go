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
