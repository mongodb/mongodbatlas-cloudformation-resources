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
	"time"
)

// X509CertificateUpdate struct for X509CertificateUpdate
type X509CertificateUpdate struct {
	// Certificate content.
	Content *string `json:"content,omitempty"`
	// Latest date that the certificate is valid.
	NotAfter *time.Time `json:"notAfter,omitempty"`
	// Earliest date that the certificate is valid.
	NotBefore *time.Time `json:"notBefore,omitempty"`
}

// NewX509CertificateUpdate instantiates a new X509CertificateUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX509CertificateUpdate() *X509CertificateUpdate {
	this := X509CertificateUpdate{}
	return &this
}

// NewX509CertificateUpdateWithDefaults instantiates a new X509CertificateUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX509CertificateUpdateWithDefaults() *X509CertificateUpdate {
	this := X509CertificateUpdate{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise
func (o *X509CertificateUpdate) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateUpdate) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}

	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *X509CertificateUpdate) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *X509CertificateUpdate) SetContent(v string) {
	o.Content = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise
func (o *X509CertificateUpdate) GetNotAfter() time.Time {
	if o == nil || IsNil(o.NotAfter) {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateUpdate) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotAfter) {
		return nil, false
	}

	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *X509CertificateUpdate) HasNotAfter() bool {
	if o != nil && !IsNil(o.NotAfter) {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *X509CertificateUpdate) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise
func (o *X509CertificateUpdate) GetNotBefore() time.Time {
	if o == nil || IsNil(o.NotBefore) {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateUpdate) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotBefore) {
		return nil, false
	}

	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *X509CertificateUpdate) HasNotBefore() bool {
	if o != nil && !IsNil(o.NotBefore) {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *X509CertificateUpdate) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

func (o X509CertificateUpdate) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o X509CertificateUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.NotAfter) {
		toSerialize["notAfter"] = o.NotAfter
	}
	if !IsNil(o.NotBefore) {
		toSerialize["notBefore"] = o.NotBefore
	}
	return toSerialize, nil
}
