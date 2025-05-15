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

// NetworkPermissionEntryStatus struct for NetworkPermissionEntryStatus
type NetworkPermissionEntryStatus struct {
	// State of the access list entry when MongoDB Cloud made this request.  | Status | Activity | |---|---| | `ACTIVE` | This access list entry applies to all relevant cloud providers. | | `PENDING` | MongoDB Cloud has started to add access list entry. This access list entry may not apply to all cloud providers at the time of this request. | | `FAILED` | MongoDB Cloud didn't succeed in adding this access list entry. |
	// Read only field.
	STATUS string `json:"STATUS"`
}

// NewNetworkPermissionEntryStatus instantiates a new NetworkPermissionEntryStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPermissionEntryStatus(sTATUS string) *NetworkPermissionEntryStatus {
	this := NetworkPermissionEntryStatus{}
	this.STATUS = sTATUS
	return &this
}

// NewNetworkPermissionEntryStatusWithDefaults instantiates a new NetworkPermissionEntryStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPermissionEntryStatusWithDefaults() *NetworkPermissionEntryStatus {
	this := NetworkPermissionEntryStatus{}
	return &this
}

// GetSTATUS returns the STATUS field value
func (o *NetworkPermissionEntryStatus) GetSTATUS() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value
// and a boolean to check if the value has been set.
func (o *NetworkPermissionEntryStatus) GetSTATUSOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.STATUS, true
}

// SetSTATUS sets field value
func (o *NetworkPermissionEntryStatus) SetSTATUS(v string) {
	o.STATUS = v
}

func (o NetworkPermissionEntryStatus) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o NetworkPermissionEntryStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
