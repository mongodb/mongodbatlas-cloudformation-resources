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

package resource_test

import (
	"testing"

	"go.mongodb.org/atlas-sdk/v20250312003/admin"

	"github.com/aws/smithy-go/ptr"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/resource-policy/cmd/resource"
	"github.com/stretchr/testify/assert"
)

func TestNewResourcePolicyCreateReq(t *testing.T) {
	var (
		body = "body"
		name = "name"
	)
	tests := []struct {
		input    *resource.Model
		expected *admin.ApiAtlasResourcePolicyCreate
		name     string
	}{
		{
			name:     "Nil Input",
			input:    nil,
			expected: nil,
		},
		{
			name: "Valid Input",
			input: &resource.Model{
				Name: ptr.String(name),
				Policies: []resource.ApiAtlasPolicy{
					{
						Body: ptr.String(body),
					},
				},
			},
			expected: &admin.ApiAtlasResourcePolicyCreate{
				Name: name,
				Policies: []admin.ApiAtlasPolicyCreate{
					{
						Body: body,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := resource.NewResourcePolicyCreateReq(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetResourcePolicyModel(t *testing.T) {
	var (
		id      = "id"
		name    = "name"
		orgID   = "orgID"
		version = "version"
	)
	tests := []struct {
		inputSDK   *admin.ApiAtlasResourcePolicy
		inputModel *resource.Model
		expected   *resource.Model
		name       string
	}{
		{
			name:       "Nil Input",
			inputSDK:   nil,
			inputModel: nil,
			expected:   new(resource.Model),
		},
		{
			name: "Valid Input",
			inputSDK: &admin.ApiAtlasResourcePolicy{
				Id:      ptr.String(id),
				Name:    ptr.String(name),
				OrgId:   ptr.String(orgID),
				Version: ptr.String(version),
			},
			expected: &resource.Model{
				Id:      ptr.String(id),
				Name:    ptr.String(name),
				OrgId:   ptr.String(orgID),
				Version: ptr.String(version),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := resource.GetResourcePolicyModel(tt.inputSDK, tt.inputModel)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestNewResourcePolicyUpdateReq(t *testing.T) {
	var (
		body = "body"
		name = "name"
	)
	tests := []struct {
		input    *resource.Model
		expected *admin.ApiAtlasResourcePolicyEdit
		name     string
	}{
		{
			name:     "Nil Input",
			input:    nil,
			expected: nil,
		},
		{
			name: "Valid Input",
			input: &resource.Model{
				Name: ptr.String(name),
				Policies: []resource.ApiAtlasPolicy{
					{
						Body: ptr.String(body),
					},
				},
			},
			expected: &admin.ApiAtlasResourcePolicyEdit{
				Name: ptr.String(name),
				Policies: &[]admin.ApiAtlasPolicyCreate{
					{
						Body: body,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := resource.NewResourcePolicyUpdateReq(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
