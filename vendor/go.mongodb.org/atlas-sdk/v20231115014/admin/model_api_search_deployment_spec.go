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

// ApiSearchDeploymentSpec struct for ApiSearchDeploymentSpec
type ApiSearchDeploymentSpec struct {
	// Hardware specification for the Search Node instance sizes.
	InstanceSize string `json:"instanceSize"`
	// Number of Search Nodes in the cluster.
	NodeCount int `json:"nodeCount"`
}

// NewApiSearchDeploymentSpec instantiates a new ApiSearchDeploymentSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSearchDeploymentSpec(instanceSize string, nodeCount int) *ApiSearchDeploymentSpec {
	this := ApiSearchDeploymentSpec{}
	this.InstanceSize = instanceSize
	this.NodeCount = nodeCount
	return &this
}

// NewApiSearchDeploymentSpecWithDefaults instantiates a new ApiSearchDeploymentSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSearchDeploymentSpecWithDefaults() *ApiSearchDeploymentSpec {
	this := ApiSearchDeploymentSpec{}
	return &this
}

// GetInstanceSize returns the InstanceSize field value
func (o *ApiSearchDeploymentSpec) GetInstanceSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentSpec) GetInstanceSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceSize, true
}

// SetInstanceSize sets field value
func (o *ApiSearchDeploymentSpec) SetInstanceSize(v string) {
	o.InstanceSize = v
}

// GetNodeCount returns the NodeCount field value
func (o *ApiSearchDeploymentSpec) GetNodeCount() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentSpec) GetNodeCountOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeCount, true
}

// SetNodeCount sets field value
func (o *ApiSearchDeploymentSpec) SetNodeCount(v int) {
	o.NodeCount = v
}

func (o ApiSearchDeploymentSpec) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ApiSearchDeploymentSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceSize"] = o.InstanceSize
	toSerialize["nodeCount"] = o.NodeCount
	return toSerialize, nil
}
