// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// CreatePushBasedLogExportProjectRequest struct for CreatePushBasedLogExportProjectRequest
type CreatePushBasedLogExportProjectRequest struct {
	// The name of the bucket to which the agent will send the logs to.
	BucketName string `json:"bucketName"`
	// ID of the AWS IAM role that will be used to write to the S3 bucket.
	IamRoleId string `json:"iamRoleId"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// S3 directory in which vector will write to in order to store the logs. An empty string denotes the root directory.
	PrefixPath string `json:"prefixPath"`
}

// NewCreatePushBasedLogExportProjectRequest instantiates a new CreatePushBasedLogExportProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePushBasedLogExportProjectRequest(bucketName string, iamRoleId string, prefixPath string) *CreatePushBasedLogExportProjectRequest {
	this := CreatePushBasedLogExportProjectRequest{}
	this.BucketName = bucketName
	this.IamRoleId = iamRoleId
	this.PrefixPath = prefixPath
	return &this
}

// NewCreatePushBasedLogExportProjectRequestWithDefaults instantiates a new CreatePushBasedLogExportProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePushBasedLogExportProjectRequestWithDefaults() *CreatePushBasedLogExportProjectRequest {
	this := CreatePushBasedLogExportProjectRequest{}
	return &this
}

// GetBucketName returns the BucketName field value
func (o *CreatePushBasedLogExportProjectRequest) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *CreatePushBasedLogExportProjectRequest) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *CreatePushBasedLogExportProjectRequest) SetBucketName(v string) {
	o.BucketName = v
}

// GetIamRoleId returns the IamRoleId field value
func (o *CreatePushBasedLogExportProjectRequest) GetIamRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IamRoleId
}

// GetIamRoleIdOk returns a tuple with the IamRoleId field value
// and a boolean to check if the value has been set.
func (o *CreatePushBasedLogExportProjectRequest) GetIamRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IamRoleId, true
}

// SetIamRoleId sets field value
func (o *CreatePushBasedLogExportProjectRequest) SetIamRoleId(v string) {
	o.IamRoleId = v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *CreatePushBasedLogExportProjectRequest) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePushBasedLogExportProjectRequest) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreatePushBasedLogExportProjectRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *CreatePushBasedLogExportProjectRequest) SetLinks(v []Link) {
	o.Links = &v
}

// GetPrefixPath returns the PrefixPath field value
func (o *CreatePushBasedLogExportProjectRequest) GetPrefixPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrefixPath
}

// GetPrefixPathOk returns a tuple with the PrefixPath field value
// and a boolean to check if the value has been set.
func (o *CreatePushBasedLogExportProjectRequest) GetPrefixPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixPath, true
}

// SetPrefixPath sets field value
func (o *CreatePushBasedLogExportProjectRequest) SetPrefixPath(v string) {
	o.PrefixPath = v
}

func (o CreatePushBasedLogExportProjectRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreatePushBasedLogExportProjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucketName"] = o.BucketName
	toSerialize["iamRoleId"] = o.IamRoleId
	toSerialize["prefixPath"] = o.PrefixPath
	return toSerialize, nil
}
