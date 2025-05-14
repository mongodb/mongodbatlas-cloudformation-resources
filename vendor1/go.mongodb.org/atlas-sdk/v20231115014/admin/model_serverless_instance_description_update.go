// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ServerlessInstanceDescriptionUpdate Settings that you can update when you request a serverless cluster update.
type ServerlessInstanceDescriptionUpdate struct {
	ServerlessBackupOptions *ClusterServerlessBackupOptions `json:"serverlessBackupOptions,omitempty"`
	// List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the serverless instance.
	Tags *[]ResourceTag `json:"tags,omitempty"`
	// Flag that indicates whether termination protection is enabled on the serverless instance. If set to `true`, MongoDB Cloud won't delete the serverless instance. If set to `false`, MongoDB Cloud will delete the serverless instance.
	TerminationProtectionEnabled *bool `json:"terminationProtectionEnabled,omitempty"`
}

// NewServerlessInstanceDescriptionUpdate instantiates a new ServerlessInstanceDescriptionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessInstanceDescriptionUpdate() *ServerlessInstanceDescriptionUpdate {
	this := ServerlessInstanceDescriptionUpdate{}
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	return &this
}

// NewServerlessInstanceDescriptionUpdateWithDefaults instantiates a new ServerlessInstanceDescriptionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessInstanceDescriptionUpdateWithDefaults() *ServerlessInstanceDescriptionUpdate {
	this := ServerlessInstanceDescriptionUpdate{}
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	return &this
}

// GetServerlessBackupOptions returns the ServerlessBackupOptions field value if set, zero value otherwise
func (o *ServerlessInstanceDescriptionUpdate) GetServerlessBackupOptions() ClusterServerlessBackupOptions {
	if o == nil || IsNil(o.ServerlessBackupOptions) {
		var ret ClusterServerlessBackupOptions
		return ret
	}
	return *o.ServerlessBackupOptions
}

// GetServerlessBackupOptionsOk returns a tuple with the ServerlessBackupOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionUpdate) GetServerlessBackupOptionsOk() (*ClusterServerlessBackupOptions, bool) {
	if o == nil || IsNil(o.ServerlessBackupOptions) {
		return nil, false
	}

	return o.ServerlessBackupOptions, true
}

// HasServerlessBackupOptions returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionUpdate) HasServerlessBackupOptions() bool {
	if o != nil && !IsNil(o.ServerlessBackupOptions) {
		return true
	}

	return false
}

// SetServerlessBackupOptions gets a reference to the given ClusterServerlessBackupOptions and assigns it to the ServerlessBackupOptions field.
func (o *ServerlessInstanceDescriptionUpdate) SetServerlessBackupOptions(v ClusterServerlessBackupOptions) {
	o.ServerlessBackupOptions = &v
}

// GetTags returns the Tags field value if set, zero value otherwise
func (o *ServerlessInstanceDescriptionUpdate) GetTags() []ResourceTag {
	if o == nil || IsNil(o.Tags) {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionUpdate) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}

	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionUpdate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *ServerlessInstanceDescriptionUpdate) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field value if set, zero value otherwise
func (o *ServerlessInstanceDescriptionUpdate) GetTerminationProtectionEnabled() bool {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.TerminationProtectionEnabled
}

// GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessInstanceDescriptionUpdate) GetTerminationProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		return nil, false
	}

	return o.TerminationProtectionEnabled, true
}

// HasTerminationProtectionEnabled returns a boolean if a field has been set.
func (o *ServerlessInstanceDescriptionUpdate) HasTerminationProtectionEnabled() bool {
	if o != nil && !IsNil(o.TerminationProtectionEnabled) {
		return true
	}

	return false
}

// SetTerminationProtectionEnabled gets a reference to the given bool and assigns it to the TerminationProtectionEnabled field.
func (o *ServerlessInstanceDescriptionUpdate) SetTerminationProtectionEnabled(v bool) {
	o.TerminationProtectionEnabled = &v
}

func (o ServerlessInstanceDescriptionUpdate) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServerlessInstanceDescriptionUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerlessBackupOptions) {
		toSerialize["serverlessBackupOptions"] = o.ServerlessBackupOptions
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TerminationProtectionEnabled) {
		toSerialize["terminationProtectionEnabled"] = o.TerminationProtectionEnabled
	}
	return toSerialize, nil
}
