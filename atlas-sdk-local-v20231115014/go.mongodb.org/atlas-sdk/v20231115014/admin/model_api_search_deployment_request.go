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
