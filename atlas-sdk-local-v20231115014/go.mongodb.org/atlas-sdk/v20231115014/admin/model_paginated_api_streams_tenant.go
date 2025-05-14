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

// PaginatedApiStreamsTenant struct for PaginatedApiStreamsTenant
type PaginatedApiStreamsTenant struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	// Read only field.
	Results *[]StreamsTenant `json:"results,omitempty"`
	// Total number of documents available. MongoDB Cloud omits this value if `includeCount` is set to `false`.
	// Read only field.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedApiStreamsTenant instantiates a new PaginatedApiStreamsTenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedApiStreamsTenant() *PaginatedApiStreamsTenant {
	this := PaginatedApiStreamsTenant{}
	return &this
}

// NewPaginatedApiStreamsTenantWithDefaults instantiates a new PaginatedApiStreamsTenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedApiStreamsTenantWithDefaults() *PaginatedApiStreamsTenant {
	this := PaginatedApiStreamsTenant{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *PaginatedApiStreamsTenant) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiStreamsTenant) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedApiStreamsTenant) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedApiStreamsTenant) SetLinks(v []Link) {
	o.Links = &v
}

// GetResults returns the Results field value if set, zero value otherwise
func (o *PaginatedApiStreamsTenant) GetResults() []StreamsTenant {
	if o == nil || IsNil(o.Results) {
		var ret []StreamsTenant
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiStreamsTenant) GetResultsOk() (*[]StreamsTenant, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}

	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedApiStreamsTenant) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []StreamsTenant and assigns it to the Results field.
func (o *PaginatedApiStreamsTenant) SetResults(v []StreamsTenant) {
	o.Results = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise
func (o *PaginatedApiStreamsTenant) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedApiStreamsTenant) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}

	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedApiStreamsTenant) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedApiStreamsTenant) SetTotalCount(v int) {
	o.TotalCount = &v
}

func (o PaginatedApiStreamsTenant) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedApiStreamsTenant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
