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

// ApiError struct for ApiError
type ApiError struct {
	// Describes the specific conditions or reasons that cause each type of error.
	Detail *string `json:"detail,omitempty"`
	// HTTP status code returned with this error.
	Error *int `json:"error,omitempty"`
	// Application error code returned with this error.
	ErrorCode *string `json:"errorCode,omitempty"`
	// Parameters used to give more information about the error.
	Parameters *[]interface{} `json:"parameters,omitempty"`
	// Application error message returned with this error.
	Reason *string `json:"reason,omitempty"`
}

// NewApiError instantiates a new ApiError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiError() *ApiError {
	this := ApiError{}
	return &this
}

// NewApiErrorWithDefaults instantiates a new ApiError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorWithDefaults() *ApiError {
	this := ApiError{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise
func (o *ApiError) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}

	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ApiError) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ApiError) SetDetail(v string) {
	o.Detail = &v
}

// GetError returns the Error field value if set, zero value otherwise
func (o *ApiError) GetError() int {
	if o == nil || IsNil(o.Error) {
		var ret int
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetErrorOk() (*int, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}

	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ApiError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given int and assigns it to the Error field.
func (o *ApiError) SetError(v int) {
	o.Error = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise
func (o *ApiError) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}

	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ApiError) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ApiError) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise
func (o *ApiError) GetParameters() []interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret []interface{}
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetParametersOk() (*[]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}

	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ApiError) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []interface{} and assigns it to the Parameters field.
func (o *ApiError) SetParameters(v []interface{}) {
	o.Parameters = &v
}

// GetReason returns the Reason field value if set, zero value otherwise
func (o *ApiError) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}

	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ApiError) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ApiError) SetReason(v string) {
	o.Reason = &v
}

func (o ApiError) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ApiError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}
