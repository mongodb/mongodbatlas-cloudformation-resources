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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
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

func Test_NewBaseURL(t *testing.T) {
	tests := []struct {
		name     string
		profile  profile.Profile
		envURL   string
		expected string
	}{
		{
			name:     "empty profile returns empty string",
			profile:  profile.Profile{},
			expected: "",
		},
		{
			name:     "explicit BaseURL is returned",
			profile:  profile.Profile{BaseURL: "https://custom.example.com/"},
			expected: "https://custom.example.com/",
		},
		{
			name:     "IsMongoDBGovCloud true returns gov URL",
			profile:  profile.Profile{IsMongoDBGovCloud: util.Pointer(true)},
			expected: profile.GovCloudBaseURL,
		},
		{
			name:     "IsMongoDBGovCloud false returns empty string",
			profile:  profile.Profile{IsMongoDBGovCloud: util.Pointer(false)},
			expected: "",
		},
		{
			name:     "BaseURL takes precedence over IsMongoDBGovCloud",
			profile:  profile.Profile{BaseURL: "https://custom.example.com/", IsMongoDBGovCloud: util.Pointer(true)},
			expected: "https://custom.example.com/",
		},
		{
			name:     "env var takes precedence over BaseURL",
			profile:  profile.Profile{BaseURL: "https://custom.example.com/"},
			envURL:   "https://env.example.com/",
			expected: "https://env.example.com/",
		},
		{
			name:     "env var takes precedence over IsMongoDBGovCloud",
			profile:  profile.Profile{IsMongoDBGovCloud: util.Pointer(true)},
			envURL:   "https://env.example.com/",
			expected: "https://env.example.com/",
		},
		{
			name:     "env var takes precedence over both BaseURL and IsMongoDBGovCloud",
			profile:  profile.Profile{BaseURL: "https://custom.example.com/", IsMongoDBGovCloud: util.Pointer(true)},
			envURL:   "https://env.example.com/",
			expected: "https://env.example.com/",
		},
		{
			name:     "IsMongoDBGovCloud nil returns empty string",
			profile:  profile.Profile{IsMongoDBGovCloud: nil},
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.envURL != "" {
				t.Setenv("MONGODB_ATLAS_BASE_URL", tc.envURL)
			} else {
				t.Setenv("MONGODB_ATLAS_BASE_URL", "")
			}
			result := tc.profile.NewBaseURL()
			assert.Equal(t, tc.expected, result)
		})
	}
}
