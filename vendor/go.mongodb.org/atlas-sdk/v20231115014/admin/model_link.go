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

// Link struct for Link
type Link struct {
	// Uniform Resource Locator (URL) that points another API resource to which this response has some relationship. This URL often begins with `https://cloud.mongodb.com/api/atlas`.
	Href *string `json:"href,omitempty"`
	// Uniform Resource Locator (URL) that defines the semantic relationship between this resource and another API resource. This URL often begins with `https://cloud.mongodb.com/api/atlas`.
	Rel *string `json:"rel,omitempty"`
}

// NewLink instantiates a new Link object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLink() *Link {
	this := Link{}
	return &this
}

// NewLinkWithDefaults instantiates a new Link object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkWithDefaults() *Link {
	this := Link{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise
func (o *Link) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}

	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Link) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Link) SetHref(v string) {
	o.Href = &v
}

// GetRel returns the Rel field value if set, zero value otherwise
func (o *Link) GetRel() string {
	if o == nil || IsNil(o.Rel) {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetRelOk() (*string, bool) {
	if o == nil || IsNil(o.Rel) {
		return nil, false
	}

	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *Link) HasRel() bool {
	if o != nil && !IsNil(o.Rel) {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *Link) SetRel(v string) {
	o.Rel = &v
}

func (o Link) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o Link) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Rel) {
		toSerialize["rel"] = o.Rel
	}
	return toSerialize, nil
}
