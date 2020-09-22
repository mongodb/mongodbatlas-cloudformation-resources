package resource

import (
    "os"
    "testing"
    "github.com/mongodb/mongodbatlas-cloudformation-resources/testutil"
)

func Test(t *testing.T) {
    dir, _ := os.Getwd()
    handler := &Handler{}
    model := &Model{}
    testutil.TestResource( dir, handler, model, t )
}

