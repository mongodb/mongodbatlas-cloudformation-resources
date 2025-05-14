// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// GroupUpdate Request view to update the group.
type GroupUpdate struct {
	// Human-readable label that identifies the project included in the MongoDB Cloud organization.
	Name *string `json:"name,omitempty"`
	// List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the project.
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

// NewGroupUpdate instantiates a new GroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUpdate() *GroupUpdate {
	this := GroupUpdate{}
	return &this
}

// NewGroupUpdateWithDefaults instantiates a new GroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUpdateWithDefaults() *GroupUpdate {
	this := GroupUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise
func (o *GroupUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupUpdate) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise
func (o *GroupUpdate) GetTags() []ResourceTag {
	if o == nil || IsNil(o.Tags) {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdate) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}

	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GroupUpdate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *GroupUpdate) SetTags(v []ResourceTag) {
	o.Tags = &v
}

func (o GroupUpdate) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}
