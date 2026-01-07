// Copyright 2025 MongoDB Inc
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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/org-service-account/cmd/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate_MissingRequiredFields(t *testing.T) {
	req := handler.Request{}
	prevModel := &resource.Model{}
	model := &resource.Model{
		// Missing required fields: OrgId, Name, Description, Roles
	}

	event, err := resource.Create(req, prevModel, model)

	require.NoError(t, err)
	assert.Equal(t, handler.Failed, event.OperationStatus)
}

func TestRead_MissingRequiredFields(t *testing.T) {
	req := handler.Request{}
	prevModel := &resource.Model{}
	model := &resource.Model{
		// Missing required fields: OrgId, ClientId
	}

	event, err := resource.Read(req, prevModel, model)

	require.NoError(t, err)
	assert.Equal(t, handler.Failed, event.OperationStatus)
}

func TestUpdate_MissingRequiredFields(t *testing.T) {
	req := handler.Request{}
	prevModel := &resource.Model{}
	model := &resource.Model{
		// Missing required fields: OrgId, ClientId
	}

	event, err := resource.Update(req, prevModel, model)

	require.NoError(t, err)
	assert.Equal(t, handler.Failed, event.OperationStatus)
}

func TestDelete_MissingRequiredFields(t *testing.T) {
	req := handler.Request{}
	prevModel := &resource.Model{}
	model := &resource.Model{
		// Missing required fields: OrgId, ClientId
	}

	event, err := resource.Delete(req, prevModel, model)

	require.NoError(t, err)
	assert.Equal(t, handler.Failed, event.OperationStatus)
}

func TestList_MissingRequiredFields(t *testing.T) {
	req := handler.Request{}
	prevModel := &resource.Model{}
	model := &resource.Model{
		// Missing required field: OrgId
	}

	event, err := resource.List(req, prevModel, model)

	require.NoError(t, err)
	assert.Equal(t, handler.Failed, event.OperationStatus)
}
