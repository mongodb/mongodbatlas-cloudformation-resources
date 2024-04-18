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
