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

// ApiSearchDeploymentRequest struct for ApiSearchDeploymentRequest
type ApiSearchDeploymentRequest struct {
	// List of settings that configure the Search Nodes for your cluster.
	Specs []ApiSearchDeploymentSpec `json:"specs"`
}

// NewApiSearchDeploymentRequest instantiates a new ApiSearchDeploymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSearchDeploymentRequest(specs []ApiSearchDeploymentSpec) *ApiSearchDeploymentRequest {
	this := ApiSearchDeploymentRequest{}
	this.Specs = specs
	return &this
}

// NewApiSearchDeploymentRequestWithDefaults instantiates a new ApiSearchDeploymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSearchDeploymentRequestWithDefaults() *ApiSearchDeploymentRequest {
	this := ApiSearchDeploymentRequest{}
	return &this
}

// GetSpecs returns the Specs field value
func (o *ApiSearchDeploymentRequest) GetSpecs() []ApiSearchDeploymentSpec {
	if o == nil {
		var ret []ApiSearchDeploymentSpec
		return ret
	}

	return o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value
// and a boolean to check if the value has been set.
func (o *ApiSearchDeploymentRequest) GetSpecsOk() (*[]ApiSearchDeploymentSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specs, true
}

// SetSpecs sets field value
func (o *ApiSearchDeploymentRequest) SetSpecs(v []ApiSearchDeploymentSpec) {
	o.Specs = v
}

func (o ApiSearchDeploymentRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ApiSearchDeploymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["specs"] = o.Specs
	return toSerialize, nil
}
