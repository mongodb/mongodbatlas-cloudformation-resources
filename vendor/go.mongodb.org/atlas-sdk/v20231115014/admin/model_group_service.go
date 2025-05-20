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

// GroupService List of IP addresses in a project categorized by services.
type GroupService struct {
	// IP addresses of clusters.
	// Read only field.
	Clusters *[]ClusterIPAddresses `json:"clusters,omitempty"`
}

// NewGroupService instantiates a new GroupService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupService() *GroupService {
	this := GroupService{}
	return &this
}

// NewGroupServiceWithDefaults instantiates a new GroupService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupServiceWithDefaults() *GroupService {
	this := GroupService{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise
func (o *GroupService) GetClusters() []ClusterIPAddresses {
	if o == nil || IsNil(o.Clusters) {
		var ret []ClusterIPAddresses
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupService) GetClustersOk() (*[]ClusterIPAddresses, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}

	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *GroupService) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []ClusterIPAddresses and assigns it to the Clusters field.
func (o *GroupService) SetClusters(v []ClusterIPAddresses) {
	o.Clusters = &v
}

func (o GroupService) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
