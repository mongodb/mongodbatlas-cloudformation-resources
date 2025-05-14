// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// CustomZoneMappings struct for CustomZoneMappings
type CustomZoneMappings struct {
	// List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to the human-readable label for the desired custom zone. MongoDB Cloud maps the ISO 3166-1a2 code to the nearest geographical zone by default. Include this parameter to override the default mappings.  This parameter returns an empty object if no custom zones exist.
	CustomZoneMappings *[]ZoneMapping `json:"customZoneMappings,omitempty"`
}

// NewCustomZoneMappings instantiates a new CustomZoneMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomZoneMappings() *CustomZoneMappings {
	this := CustomZoneMappings{}
	return &this
}

// NewCustomZoneMappingsWithDefaults instantiates a new CustomZoneMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomZoneMappingsWithDefaults() *CustomZoneMappings {
	this := CustomZoneMappings{}
	return &this
}

// GetCustomZoneMappings returns the CustomZoneMappings field value if set, zero value otherwise
func (o *CustomZoneMappings) GetCustomZoneMappings() []ZoneMapping {
	if o == nil || IsNil(o.CustomZoneMappings) {
		var ret []ZoneMapping
		return ret
	}
	return *o.CustomZoneMappings
}

// GetCustomZoneMappingsOk returns a tuple with the CustomZoneMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomZoneMappings) GetCustomZoneMappingsOk() (*[]ZoneMapping, bool) {
	if o == nil || IsNil(o.CustomZoneMappings) {
		return nil, false
	}

	return o.CustomZoneMappings, true
}

// HasCustomZoneMappings returns a boolean if a field has been set.
func (o *CustomZoneMappings) HasCustomZoneMappings() bool {
	if o != nil && !IsNil(o.CustomZoneMappings) {
		return true
	}

	return false
}

// SetCustomZoneMappings gets a reference to the given []ZoneMapping and assigns it to the CustomZoneMappings field.
func (o *CustomZoneMappings) SetCustomZoneMappings(v []ZoneMapping) {
	o.CustomZoneMappings = &v
}

func (o CustomZoneMappings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CustomZoneMappings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomZoneMappings) {
		toSerialize["customZoneMappings"] = o.CustomZoneMappings
	}
	return toSerialize, nil
}
