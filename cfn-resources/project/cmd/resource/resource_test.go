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

package resource_test

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/project/cmd/resource"
)

func TestGetChangeInAPIKeys_AddedButNotChangedOrRemoved(t *testing.T) {
	currentKeys := []resource.ProjectApiKey{
		{Key: aws.String("key1"), RoleNames: []string{"role1"}},
		{Key: aws.String("key2"), RoleNames: []string{"role2"}},
	}
	previousKeys := []resource.ProjectApiKey{
		{Key: aws.String("key1"), RoleNames: []string{"role1"}},
	}

	newKeys, changedKeys, removeKeys := resource.GetChangeInAPIKeys(currentKeys, previousKeys)

	expectedNewKeys := []resource.ProjectApiKey{
		{Key: aws.String("key2"), RoleNames: []string{"role2"}},
	}

	if !reflect.DeepEqual(newKeys, expectedNewKeys) || len(changedKeys) > 0 || len(removeKeys) > 0 {
		t.Errorf("Test case failed. Expected newKeys to contain all added keys.")
	}
}

func TestGetChangeInAPIKeys_ChangedButNotAddedOrRemoved(t *testing.T) {
	currentKeys := []resource.ProjectApiKey{
		{Key: aws.String("key1"), RoleNames: []string{"role1", "role1b"}},
		{Key: aws.String("key2"), RoleNames: []string{"role2"}},
	}
	previousKeys := []resource.ProjectApiKey{
		{Key: aws.String("key1"), RoleNames: []string{"role1b", "role1"}},
		{Key: aws.String("key2"), RoleNames: []string{"role3", "role2"}},
	}

	newKeys, changedKeys, removeKeys := resource.GetChangeInAPIKeys(currentKeys, previousKeys)

	if len(newKeys) > 0 || len(removeKeys) > 0 {
		t.Errorf("Test case failed. No new or removed keys expected.")
	}

	expectedChangedKeys := []resource.ProjectApiKey{
		{Key: aws.String("key2"), RoleNames: []string{"role2"}},
	}

	if !reflect.DeepEqual(changedKeys, expectedChangedKeys) {
		t.Errorf("Test case failed. Expected changedKeys to contain keys with role changes.")
	}
}

func TestGetChangeInAPIKeys_RemovedButNotAddedOrChanged(t *testing.T) {
	currentKeys := []resource.ProjectApiKey{
		{Key: aws.String("key1"), RoleNames: []string{"role1"}},
	}
	previousKeys := []resource.ProjectApiKey{
		{Key: aws.String("key1"), RoleNames: []string{"role1"}},
		{Key: aws.String("key2"), RoleNames: []string{"role2"}},
	}

	newKeys, changedKeys, removeKeys := resource.GetChangeInAPIKeys(currentKeys, previousKeys)

	if len(newKeys) > 0 || len(changedKeys) > 0 {
		t.Errorf("Test case failed. No new or changed keys expected.")
	}

	expectedRemoveKeys := []resource.ProjectApiKey{
		{Key: aws.String("key2"), RoleNames: []string{"role2"}},
	}

	if !reflect.DeepEqual(removeKeys, expectedRemoveKeys) {
		t.Errorf("Test case failed. Expected removeKeys to contain all previous keys.")
	}
}

func TestGetChangeInAPIKeys_AddedChangedRemovedMixed(t *testing.T) {
	previousKeys := []resource.ProjectApiKey{
		{Key: aws.String("key1"), RoleNames: []string{"role1"}},
		{Key: aws.String("key2"), RoleNames: []string{"role2"}},
		{Key: aws.String("key5"), RoleNames: []string{"role5"}},
		{Key: aws.String("key6"), RoleNames: []string{"role6"}},
	}
	currentKeys := []resource.ProjectApiKey{
		{Key: aws.String("key1"), RoleNames: []string{"role9"}},
		{Key: aws.String("key2"), RoleNames: []string{"role2"}},
		{Key: aws.String("key3"), RoleNames: []string{"role3"}},
		{Key: aws.String("key4"), RoleNames: []string{"role4"}},
	}

	newKeys, changedKeys, removeKeys := resource.GetChangeInAPIKeys(currentKeys, previousKeys)

	expectedNewKeys := []resource.ProjectApiKey{
		{Key: aws.String("key3"), RoleNames: []string{"role3"}},
		{Key: aws.String("key4"), RoleNames: []string{"role4"}},
	}
	expectedChangedKeys := []resource.ProjectApiKey{
		{Key: aws.String("key1"), RoleNames: []string{"role9"}},
	}
	expectedRemoveKeys := []resource.ProjectApiKey{
		{Key: aws.String("key5"), RoleNames: []string{"role5"}},
		{Key: aws.String("key6"), RoleNames: []string{"role6"}},
	}

	if !reflect.DeepEqual(newKeys, expectedNewKeys) || !reflect.DeepEqual(changedKeys, expectedChangedKeys) || !reflect.DeepEqual(removeKeys, expectedRemoveKeys) {
		t.Errorf("Test case failed. Unexpected results for mixed case.")
	}
}

func TestGetChangeInAPIKeys_NoneAddedChangedOrRemoved(t *testing.T) {
	currentKeys := []resource.ProjectApiKey{
		{Key: aws.String("key1"), RoleNames: []string{"role1"}},
		{Key: aws.String("key2"), RoleNames: []string{"role2"}},
	}
	previousKeys := []resource.ProjectApiKey{
		{Key: aws.String("key1"), RoleNames: []string{"role1"}},
		{Key: aws.String("key2"), RoleNames: []string{"role2"}},
	}

	newKeys, changedKeys, removeKeys := resource.GetChangeInAPIKeys(currentKeys, previousKeys)

	if len(newKeys) > 0 || len(changedKeys) > 0 || len(removeKeys) > 0 {
		t.Errorf("Test case failed. No new, changed, or removed keys expected.")
	}
}
