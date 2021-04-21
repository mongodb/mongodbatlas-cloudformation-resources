package resource

import (
	"github.com/mongodb/mongodbatlas-cloudformation-resources/testutil"
	"os"
	"testing"
)

func Test(t *testing.T) {
	dir, _ := os.Getwd()
	handler := &Handler{}
	model := &Model{}
	testutil.TestResource(dir, handler, model, t)
}
