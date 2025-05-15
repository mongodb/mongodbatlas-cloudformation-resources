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

// CloudProviderAccessRoles struct for CloudProviderAccessRoles
type CloudProviderAccessRoles struct {
	// List that contains the Amazon Web Services (AWS) IAM roles registered and authorized with MongoDB Cloud.
	AwsIamRoles *[]CloudProviderAccessAWSIAMRole `json:"awsIamRoles,omitempty"`
}

// NewCloudProviderAccessRoles instantiates a new CloudProviderAccessRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessRoles() *CloudProviderAccessRoles {
	this := CloudProviderAccessRoles{}
	return &this
}

// NewCloudProviderAccessRolesWithDefaults instantiates a new CloudProviderAccessRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessRolesWithDefaults() *CloudProviderAccessRoles {
	this := CloudProviderAccessRoles{}
	return &this
}

// GetAwsIamRoles returns the AwsIamRoles field value if set, zero value otherwise
func (o *CloudProviderAccessRoles) GetAwsIamRoles() []CloudProviderAccessAWSIAMRole {
	if o == nil || IsNil(o.AwsIamRoles) {
		var ret []CloudProviderAccessAWSIAMRole
		return ret
	}
	return *o.AwsIamRoles
}

// GetAwsIamRolesOk returns a tuple with the AwsIamRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRoles) GetAwsIamRolesOk() (*[]CloudProviderAccessAWSIAMRole, bool) {
	if o == nil || IsNil(o.AwsIamRoles) {
		return nil, false
	}

	return o.AwsIamRoles, true
}

// HasAwsIamRoles returns a boolean if a field has been set.
func (o *CloudProviderAccessRoles) HasAwsIamRoles() bool {
	if o != nil && !IsNil(o.AwsIamRoles) {
		return true
	}

	return false
}

// SetAwsIamRoles gets a reference to the given []CloudProviderAccessAWSIAMRole and assigns it to the AwsIamRoles field.
func (o *CloudProviderAccessRoles) SetAwsIamRoles(v []CloudProviderAccessAWSIAMRole) {
	o.AwsIamRoles = &v
}

func (o CloudProviderAccessRoles) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudProviderAccessRoles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsIamRoles) {
		toSerialize["awsIamRoles"] = o.AwsIamRoles
	}
	return toSerialize, nil
}
