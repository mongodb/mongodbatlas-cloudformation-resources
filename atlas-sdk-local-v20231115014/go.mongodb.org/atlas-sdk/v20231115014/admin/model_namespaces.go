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

// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// Namespaces struct for Namespaces
type Namespaces struct {
	// List that contains each combination of database, collection, and type on the specified host.
	// Read only field.
	Namespaces *[]NamespaceObj `json:"namespaces,omitempty"`
}

// NewNamespaces instantiates a new Namespaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamespaces() *Namespaces {
	this := Namespaces{}
	return &this
}

// NewNamespacesWithDefaults instantiates a new Namespaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamespacesWithDefaults() *Namespaces {
	this := Namespaces{}
	return &this
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise
func (o *Namespaces) GetNamespaces() []NamespaceObj {
	if o == nil || IsNil(o.Namespaces) {
		var ret []NamespaceObj
		return ret
	}
	return *o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Namespaces) GetNamespacesOk() (*[]NamespaceObj, bool) {
	if o == nil || IsNil(o.Namespaces) {
		return nil, false
	}

	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *Namespaces) HasNamespaces() bool {
	if o != nil && !IsNil(o.Namespaces) {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []NamespaceObj and assigns it to the Namespaces field.
func (o *Namespaces) SetNamespaces(v []NamespaceObj) {
	o.Namespaces = &v
}

func (o Namespaces) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o Namespaces) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
