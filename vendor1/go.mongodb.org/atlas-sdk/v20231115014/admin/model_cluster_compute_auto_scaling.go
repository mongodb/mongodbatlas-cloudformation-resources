// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ClusterComputeAutoScaling Collection of settings that configures how a cluster might scale its cluster tier and whether the cluster can scale down. Cluster tier auto-scaling is unavailable for clusters using Low CPU or NVME storage classes.
type ClusterComputeAutoScaling struct {
	// Flag that indicates whether cluster tier auto-scaling is enabled. Set to `true` to enable cluster tier auto-scaling. If enabled, you must specify a value for **providerSettings.autoScaling.compute.maxInstanceSize** also. Set to `false` to disable cluster tier auto-scaling.
	Enabled *bool `json:"enabled,omitempty"`
	// Flag that indicates whether the cluster tier can scale down. This is required if **autoScaling.compute.enabled** is `true`. If you enable this option, specify a value for **providerSettings.autoScaling.compute.minInstanceSize**.
	ScaleDownEnabled *bool `json:"scaleDownEnabled,omitempty"`
}

// NewClusterComputeAutoScaling instantiates a new ClusterComputeAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterComputeAutoScaling() *ClusterComputeAutoScaling {
	this := ClusterComputeAutoScaling{}
	return &this
}

// NewClusterComputeAutoScalingWithDefaults instantiates a new ClusterComputeAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterComputeAutoScalingWithDefaults() *ClusterComputeAutoScaling {
	this := ClusterComputeAutoScaling{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *ClusterComputeAutoScaling) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterComputeAutoScaling) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ClusterComputeAutoScaling) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ClusterComputeAutoScaling) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetScaleDownEnabled returns the ScaleDownEnabled field value if set, zero value otherwise
func (o *ClusterComputeAutoScaling) GetScaleDownEnabled() bool {
	if o == nil || IsNil(o.ScaleDownEnabled) {
		var ret bool
		return ret
	}
	return *o.ScaleDownEnabled
}

// GetScaleDownEnabledOk returns a tuple with the ScaleDownEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterComputeAutoScaling) GetScaleDownEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ScaleDownEnabled) {
		return nil, false
	}

	return o.ScaleDownEnabled, true
}

// HasScaleDownEnabled returns a boolean if a field has been set.
func (o *ClusterComputeAutoScaling) HasScaleDownEnabled() bool {
	if o != nil && !IsNil(o.ScaleDownEnabled) {
		return true
	}

	return false
}

// SetScaleDownEnabled gets a reference to the given bool and assigns it to the ScaleDownEnabled field.
func (o *ClusterComputeAutoScaling) SetScaleDownEnabled(v bool) {
	o.ScaleDownEnabled = &v
}

func (o ClusterComputeAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterComputeAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ScaleDownEnabled) {
		toSerialize["scaleDownEnabled"] = o.ScaleDownEnabled
	}
	return toSerialize, nil
}
