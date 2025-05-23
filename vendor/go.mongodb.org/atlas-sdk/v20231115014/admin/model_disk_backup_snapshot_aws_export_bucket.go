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

// DiskBackupSnapshotAWSExportBucket struct for DiskBackupSnapshotAWSExportBucket
type DiskBackupSnapshotAWSExportBucket struct {
	// Unique 24-hexadecimal character string that identifies the Amazon Web Services (AWS) Simple Storage Service (S3) export bucket.
	// Read only field.
	Id *string `json:"_id,omitempty"`
	// Human-readable label that identifies the AWS bucket that the role is authorized to access.
	BucketName *string `json:"bucketName,omitempty"`
	// Human-readable label that identifies the cloud provider that stores this snapshot.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// Unique 24-hexadecimal character string that identifies the <a href='https://www.mongodb.com/docs/atlas/security/set-up-unified-aws-access/' target='_blank'>Unified AWS Access role ID</a>  that MongoDB Cloud uses to access the AWS S3 bucket.
	IamRoleId *string `json:"iamRoleId,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
}

// NewDiskBackupSnapshotAWSExportBucket instantiates a new DiskBackupSnapshotAWSExportBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupSnapshotAWSExportBucket() *DiskBackupSnapshotAWSExportBucket {
	this := DiskBackupSnapshotAWSExportBucket{}
	return &this
}

// NewDiskBackupSnapshotAWSExportBucketWithDefaults instantiates a new DiskBackupSnapshotAWSExportBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupSnapshotAWSExportBucketWithDefaults() *DiskBackupSnapshotAWSExportBucket {
	this := DiskBackupSnapshotAWSExportBucket{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *DiskBackupSnapshotAWSExportBucket) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotAWSExportBucket) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiskBackupSnapshotAWSExportBucket) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DiskBackupSnapshotAWSExportBucket) SetId(v string) {
	o.Id = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise
func (o *DiskBackupSnapshotAWSExportBucket) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotAWSExportBucket) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}

	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *DiskBackupSnapshotAWSExportBucket) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *DiskBackupSnapshotAWSExportBucket) SetBucketName(v string) {
	o.BucketName = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *DiskBackupSnapshotAWSExportBucket) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotAWSExportBucket) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *DiskBackupSnapshotAWSExportBucket) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *DiskBackupSnapshotAWSExportBucket) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetIamRoleId returns the IamRoleId field value if set, zero value otherwise
func (o *DiskBackupSnapshotAWSExportBucket) GetIamRoleId() string {
	if o == nil || IsNil(o.IamRoleId) {
		var ret string
		return ret
	}
	return *o.IamRoleId
}

// GetIamRoleIdOk returns a tuple with the IamRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotAWSExportBucket) GetIamRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.IamRoleId) {
		return nil, false
	}

	return o.IamRoleId, true
}

// HasIamRoleId returns a boolean if a field has been set.
func (o *DiskBackupSnapshotAWSExportBucket) HasIamRoleId() bool {
	if o != nil && !IsNil(o.IamRoleId) {
		return true
	}

	return false
}

// SetIamRoleId gets a reference to the given string and assigns it to the IamRoleId field.
func (o *DiskBackupSnapshotAWSExportBucket) SetIamRoleId(v string) {
	o.IamRoleId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *DiskBackupSnapshotAWSExportBucket) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotAWSExportBucket) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DiskBackupSnapshotAWSExportBucket) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *DiskBackupSnapshotAWSExportBucket) SetLinks(v []Link) {
	o.Links = &v
}

func (o DiskBackupSnapshotAWSExportBucket) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DiskBackupSnapshotAWSExportBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BucketName) {
		toSerialize["bucketName"] = o.BucketName
	}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.IamRoleId) {
		toSerialize["iamRoleId"] = o.IamRoleId
	}
	return toSerialize, nil
}
