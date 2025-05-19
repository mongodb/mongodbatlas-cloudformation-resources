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

// StreamsKafkaAuthentication User credentials required to connect to a Kafka Cluster. Includes the authentication type, as well as the parameters for that authentication mode.
type StreamsKafkaAuthentication struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Style of authentication. Can be one of PLAIN, SCRAM-256, or SCRAM-512.
	Mechanism *string `json:"mechanism,omitempty"`
	// Password of the account to connect to the Kafka cluster.
	// Write only field.
	Password *string `json:"password,omitempty"`
	// Username of the account to connect to the Kafka cluster.
	Username *string `json:"username,omitempty"`
}

// NewStreamsKafkaAuthentication instantiates a new StreamsKafkaAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsKafkaAuthentication() *StreamsKafkaAuthentication {
	this := StreamsKafkaAuthentication{}
	return &this
}

// NewStreamsKafkaAuthenticationWithDefaults instantiates a new StreamsKafkaAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsKafkaAuthenticationWithDefaults() *StreamsKafkaAuthentication {
	this := StreamsKafkaAuthentication{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsKafkaAuthentication) SetLinks(v []Link) {
	o.Links = &v
}

// GetMechanism returns the Mechanism field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetMechanism() string {
	if o == nil || IsNil(o.Mechanism) {
		var ret string
		return ret
	}
	return *o.Mechanism
}

// GetMechanismOk returns a tuple with the Mechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetMechanismOk() (*string, bool) {
	if o == nil || IsNil(o.Mechanism) {
		return nil, false
	}

	return o.Mechanism, true
}

// HasMechanism returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasMechanism() bool {
	if o != nil && !IsNil(o.Mechanism) {
		return true
	}

	return false
}

// SetMechanism gets a reference to the given string and assigns it to the Mechanism field.
func (o *StreamsKafkaAuthentication) SetMechanism(v string) {
	o.Mechanism = &v
}

// GetPassword returns the Password field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}

	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *StreamsKafkaAuthentication) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise
func (o *StreamsKafkaAuthentication) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}

	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *StreamsKafkaAuthentication) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *StreamsKafkaAuthentication) SetUsername(v string) {
	o.Username = &v
}

func (o StreamsKafkaAuthentication) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o StreamsKafkaAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mechanism) {
		toSerialize["mechanism"] = o.Mechanism
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}
