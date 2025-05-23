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

// GoogleCloudKMS Details that define the configuration of Encryption at Rest using Google Cloud Key Management Service (KMS).
type GoogleCloudKMS struct {
	// Flag that indicates whether someone enabled encryption at rest for the specified  project. To disable encryption at rest using customer key management and remove the configuration details, pass only this parameter with a value of `false`.
	Enabled *bool `json:"enabled,omitempty"`
	// Resource path that displays the key version resource ID for your Google Cloud KMS.
	KeyVersionResourceID *string `json:"keyVersionResourceID,omitempty"`
	// JavaScript Object Notation (JSON) object that contains the Google Cloud Key Management Service (KMS). Format the JSON as a string and not as an object.
	// Write only field.
	ServiceAccountKey *string `json:"serviceAccountKey,omitempty"`
	// Flag that indicates whether the Google Cloud Key Management Service (KMS) encryption key can encrypt and decrypt data.
	// Read only field.
	Valid *bool `json:"valid,omitempty"`
}

// NewGoogleCloudKMS instantiates a new GoogleCloudKMS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCloudKMS() *GoogleCloudKMS {
	this := GoogleCloudKMS{}
	return &this
}

// NewGoogleCloudKMSWithDefaults instantiates a new GoogleCloudKMS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudKMSWithDefaults() *GoogleCloudKMS {
	this := GoogleCloudKMS{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *GoogleCloudKMS) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudKMS) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GoogleCloudKMS) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GoogleCloudKMS) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetKeyVersionResourceID returns the KeyVersionResourceID field value if set, zero value otherwise
func (o *GoogleCloudKMS) GetKeyVersionResourceID() string {
	if o == nil || IsNil(o.KeyVersionResourceID) {
		var ret string
		return ret
	}
	return *o.KeyVersionResourceID
}

// GetKeyVersionResourceIDOk returns a tuple with the KeyVersionResourceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudKMS) GetKeyVersionResourceIDOk() (*string, bool) {
	if o == nil || IsNil(o.KeyVersionResourceID) {
		return nil, false
	}

	return o.KeyVersionResourceID, true
}

// HasKeyVersionResourceID returns a boolean if a field has been set.
func (o *GoogleCloudKMS) HasKeyVersionResourceID() bool {
	if o != nil && !IsNil(o.KeyVersionResourceID) {
		return true
	}

	return false
}

// SetKeyVersionResourceID gets a reference to the given string and assigns it to the KeyVersionResourceID field.
func (o *GoogleCloudKMS) SetKeyVersionResourceID(v string) {
	o.KeyVersionResourceID = &v
}

// GetServiceAccountKey returns the ServiceAccountKey field value if set, zero value otherwise
func (o *GoogleCloudKMS) GetServiceAccountKey() string {
	if o == nil || IsNil(o.ServiceAccountKey) {
		var ret string
		return ret
	}
	return *o.ServiceAccountKey
}

// GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudKMS) GetServiceAccountKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAccountKey) {
		return nil, false
	}

	return o.ServiceAccountKey, true
}

// HasServiceAccountKey returns a boolean if a field has been set.
func (o *GoogleCloudKMS) HasServiceAccountKey() bool {
	if o != nil && !IsNil(o.ServiceAccountKey) {
		return true
	}

	return false
}

// SetServiceAccountKey gets a reference to the given string and assigns it to the ServiceAccountKey field.
func (o *GoogleCloudKMS) SetServiceAccountKey(v string) {
	o.ServiceAccountKey = &v
}

// GetValid returns the Valid field value if set, zero value otherwise
func (o *GoogleCloudKMS) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudKMS) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}

	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *GoogleCloudKMS) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *GoogleCloudKMS) SetValid(v bool) {
	o.Valid = &v
}

func (o GoogleCloudKMS) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GoogleCloudKMS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.KeyVersionResourceID) {
		toSerialize["keyVersionResourceID"] = o.KeyVersionResourceID
	}
	if !IsNil(o.ServiceAccountKey) {
		toSerialize["serviceAccountKey"] = o.ServiceAccountKey
	}
	return toSerialize, nil
}
