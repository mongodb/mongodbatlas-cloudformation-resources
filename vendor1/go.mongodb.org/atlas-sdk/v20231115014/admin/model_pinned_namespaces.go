// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// PinnedNamespaces Pinned namespaces view for cluster
type PinnedNamespaces struct {
	// Unique 24-hexadecimal digit string that identifies the request cluster.
	// Read only field.
	ClusterId *string `json:"clusterId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the request project.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// List of all pinned namespaces.
	// Read only field.
	PinnedNamespaces *[]string `json:"pinnedNamespaces,omitempty"`
}

// NewPinnedNamespaces instantiates a new PinnedNamespaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPinnedNamespaces() *PinnedNamespaces {
	this := PinnedNamespaces{}
	return &this
}

// NewPinnedNamespacesWithDefaults instantiates a new PinnedNamespaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinnedNamespacesWithDefaults() *PinnedNamespaces {
	this := PinnedNamespaces{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise
func (o *PinnedNamespaces) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinnedNamespaces) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}

	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *PinnedNamespaces) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *PinnedNamespaces) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *PinnedNamespaces) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinnedNamespaces) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *PinnedNamespaces) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *PinnedNamespaces) SetGroupId(v string) {
	o.GroupId = &v
}

// GetPinnedNamespaces returns the PinnedNamespaces field value if set, zero value otherwise
func (o *PinnedNamespaces) GetPinnedNamespaces() []string {
	if o == nil || IsNil(o.PinnedNamespaces) {
		var ret []string
		return ret
	}
	return *o.PinnedNamespaces
}

// GetPinnedNamespacesOk returns a tuple with the PinnedNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinnedNamespaces) GetPinnedNamespacesOk() (*[]string, bool) {
	if o == nil || IsNil(o.PinnedNamespaces) {
		return nil, false
	}

	return o.PinnedNamespaces, true
}

// HasPinnedNamespaces returns a boolean if a field has been set.
func (o *PinnedNamespaces) HasPinnedNamespaces() bool {
	if o != nil && !IsNil(o.PinnedNamespaces) {
		return true
	}

	return false
}

// SetPinnedNamespaces gets a reference to the given []string and assigns it to the PinnedNamespaces field.
func (o *PinnedNamespaces) SetPinnedNamespaces(v []string) {
	o.PinnedNamespaces = &v
}

func (o PinnedNamespaces) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PinnedNamespaces) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
