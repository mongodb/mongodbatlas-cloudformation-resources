// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// GreaterThanRawThreshold A Limit that triggers an alert when greater than a number.
type GreaterThanRawThreshold struct {
	// Comparison operator to apply when checking the current metric value.
	Operator *string `json:"operator,omitempty"`
	// Value of metric that, when exceeded, triggers an alert.
	Threshold *int `json:"threshold,omitempty"`
	// Element used to express the quantity. This can be an element of time, storage capacity, and the like.
	Units *string `json:"units,omitempty"`
}

// NewGreaterThanRawThreshold instantiates a new GreaterThanRawThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGreaterThanRawThreshold() *GreaterThanRawThreshold {
	this := GreaterThanRawThreshold{}
	var units string = "RAW"
	this.Units = &units
	return &this
}

// NewGreaterThanRawThresholdWithDefaults instantiates a new GreaterThanRawThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGreaterThanRawThresholdWithDefaults() *GreaterThanRawThreshold {
	this := GreaterThanRawThreshold{}
	var units string = "RAW"
	this.Units = &units
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise
func (o *GreaterThanRawThreshold) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanRawThreshold) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}

	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *GreaterThanRawThreshold) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *GreaterThanRawThreshold) SetOperator(v string) {
	o.Operator = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise
func (o *GreaterThanRawThreshold) GetThreshold() int {
	if o == nil || IsNil(o.Threshold) {
		var ret int
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanRawThreshold) GetThresholdOk() (*int, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}

	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *GreaterThanRawThreshold) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int and assigns it to the Threshold field.
func (o *GreaterThanRawThreshold) SetThreshold(v int) {
	o.Threshold = &v
}

// GetUnits returns the Units field value if set, zero value otherwise
func (o *GreaterThanRawThreshold) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GreaterThanRawThreshold) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}

	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *GreaterThanRawThreshold) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *GreaterThanRawThreshold) SetUnits(v string) {
	o.Units = &v
}

func (o GreaterThanRawThreshold) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GreaterThanRawThreshold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	return toSerialize, nil
}
