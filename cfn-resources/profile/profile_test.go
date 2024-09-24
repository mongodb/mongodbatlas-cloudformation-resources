package profile_test

import (
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/stretchr/testify/assert"
)

func Test_UseDebug(t *testing.T) {
	profileFalse := profile.Profile{}
	assert.False(t, profileFalse.UseDebug())
	t.Setenv("MONGODB_ATLAS_DEBUG", "true")
	assert.True(t, profileFalse.UseDebug())
	t.Setenv("MONGODB_ATLAS_DEBUG", "")
	assert.False(t, profileFalse.UseDebug())
	trueBool := true
	profileTrue := profile.Profile{DebugClient: &trueBool}
	assert.True(t, profileTrue.UseDebug())
}
