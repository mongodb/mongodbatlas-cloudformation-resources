// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ClusterFreeAutoScaling Range of instance sizes to which your cluster can scale.
type ClusterFreeAutoScaling struct {
	// Collection of settings that configures how a cluster might scale its cluster tier and whether the cluster can scale down.
	Compute *string `json:"compute,omitempty"`
}

// NewClusterFreeAutoScaling instantiates a new ClusterFreeAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterFreeAutoScaling() *ClusterFreeAutoScaling {
	this := ClusterFreeAutoScaling{}
	return &this
}

// NewClusterFreeAutoScalingWithDefaults instantiates a new ClusterFreeAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterFreeAutoScalingWithDefaults() *ClusterFreeAutoScaling {
	this := ClusterFreeAutoScaling{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise
func (o *ClusterFreeAutoScaling) GetCompute() string {
	if o == nil || IsNil(o.Compute) {
		var ret string
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFreeAutoScaling) GetComputeOk() (*string, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}

	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *ClusterFreeAutoScaling) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given string and assigns it to the Compute field.
func (o *ClusterFreeAutoScaling) SetCompute(v string) {
	o.Compute = &v
}

func (o ClusterFreeAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterFreeAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	return toSerialize, nil
}
