// Copyright 2024 MongoDB Inc
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
