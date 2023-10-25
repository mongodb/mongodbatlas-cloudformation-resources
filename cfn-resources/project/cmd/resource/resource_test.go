package resource

import (
	"reflect"
	"testing"
)

func TestGetChangeInAPIKeys_AddedButNotChangedOrRemoved(t *testing.T) {
	currentKeys := []ProjectApiKey{
		{Key: StringPtr("key1"), RoleNames: []string{"role1"}},
		{Key: StringPtr("key2"), RoleNames: []string{"role2"}},
	}
	previousKeys := []ProjectApiKey{
		{Key: StringPtr("key1"), RoleNames: []string{"role1"}},
	}

	newKeys, changedKeys, removeKeys := getChangeInAPIKeys(currentKeys, previousKeys)

	expectedNewKeys := []ProjectApiKey{
		{Key: StringPtr("key2"), RoleNames: []string{"role2"}},
	}

	if !reflect.DeepEqual(newKeys, expectedNewKeys) || len(changedKeys) > 0 || len(removeKeys) > 0 {
		t.Errorf("Test case failed. Expected newKeys to contain all added keys.")
	}
}

func TestGetChangeInAPIKeys_ChangedButNotAddedOrRemoved(t *testing.T) {
	currentKeys := []ProjectApiKey{
		{Key: StringPtr("key1"), RoleNames: []string{"role1"}},
		{Key: StringPtr("key2"), RoleNames: []string{"role2"}},
	}
	previousKeys := []ProjectApiKey{
		{Key: StringPtr("key1"), RoleNames: []string{"role1"}},
		{Key: StringPtr("key2"), RoleNames: []string{"role3"}},
	}

	newKeys, changedKeys, removeKeys := getChangeInAPIKeys(currentKeys, previousKeys)

	if len(newKeys) > 0 || len(removeKeys) > 0 {
		t.Errorf("Test case failed. No new or removed keys expected.")
	}

	expectedChangedKeys := []ProjectApiKey{
		{Key: StringPtr("key2"), RoleNames: []string{"role2"}},
	}

	if !reflect.DeepEqual(changedKeys, expectedChangedKeys) {
		t.Errorf("Test case failed. Expected changedKeys to contain keys with role changes.")
	}
}

func TestGetChangeInAPIKeys_RemovedButNotAddedOrChanged(t *testing.T) {
	currentKeys := []ProjectApiKey{
		{Key: StringPtr("key1"), RoleNames: []string{"role1"}},
	}
	previousKeys := []ProjectApiKey{
		{Key: StringPtr("key1"), RoleNames: []string{"role1"}},
		{Key: StringPtr("key2"), RoleNames: []string{"role2"}},
	}

	newKeys, changedKeys, removeKeys := getChangeInAPIKeys(currentKeys, previousKeys)

	if len(newKeys) > 0 || len(changedKeys) > 0 {
		t.Errorf("Test case failed. No new or changed keys expected.")
	}

	expectedRemoveKeys := []ProjectApiKey{
		{Key: StringPtr("key2"), RoleNames: []string{"role2"}},
	}

	if !reflect.DeepEqual(removeKeys, expectedRemoveKeys) {
		t.Errorf("Test case failed. Expected removeKeys to contain all previous keys.")
	}
}

func TestGetChangeInAPIKeys_AddedChangedRemovedMixed(t *testing.T) {
	previousKeys := []ProjectApiKey{
		{Key: StringPtr("key1"), RoleNames: []string{"role1"}},
		{Key: StringPtr("key2"), RoleNames: []string{"role2"}},
		{Key: StringPtr("key5"), RoleNames: []string{"role5"}},
		{Key: StringPtr("key6"), RoleNames: []string{"role6"}},
	}
	currentKeys := []ProjectApiKey{
		{Key: StringPtr("key1"), RoleNames: []string{"role9"}},
		{Key: StringPtr("key2"), RoleNames: []string{"role2"}},
		{Key: StringPtr("key3"), RoleNames: []string{"role3"}},
		{Key: StringPtr("key4"), RoleNames: []string{"role4"}},
	}

	newKeys, changedKeys, removeKeys := getChangeInAPIKeys(currentKeys, previousKeys)

	expectedNewKeys := []ProjectApiKey{
		{Key: StringPtr("key3"), RoleNames: []string{"role3"}},
		{Key: StringPtr("key4"), RoleNames: []string{"role4"}},
	}
	expectedChangedKeys := []ProjectApiKey{
		{Key: StringPtr("key1"), RoleNames: []string{"role9"}},
	}
	expectedRemoveKeys := []ProjectApiKey{
		{Key: StringPtr("key5"), RoleNames: []string{"role5"}},
		{Key: StringPtr("key6"), RoleNames: []string{"role6"}},
	}

	if !reflect.DeepEqual(newKeys, expectedNewKeys) || !reflect.DeepEqual(changedKeys, expectedChangedKeys) || !reflect.DeepEqual(removeKeys, expectedRemoveKeys) {
		t.Errorf("Test case failed. Unexpected results for mixed case.")
	}
}

func TestGetChangeInAPIKeys_NoneAddedChangedOrRemoved(t *testing.T) {
	currentKeys := []ProjectApiKey{
		{Key: StringPtr("key1"), RoleNames: []string{"role1"}},
		{Key: StringPtr("key2"), RoleNames: []string{"role2"}},
	}
	previousKeys := []ProjectApiKey{
		{Key: StringPtr("key1"), RoleNames: []string{"role1"}},
		{Key: StringPtr("key2"), RoleNames: []string{"role2"}},
	}

	newKeys, changedKeys, removeKeys := getChangeInAPIKeys(currentKeys, previousKeys)

	if len(newKeys) > 0 || len(changedKeys) > 0 || len(removeKeys) > 0 {
		t.Errorf("Test case failed. No new, changed, or removed keys expected.")
	}
}

func StringPtr(s string) *string {
	return &s
}
