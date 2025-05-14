// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// PaginatedDatabase struct for PaginatedDatabase
type PaginatedDatabase struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// List of returned documents that MongoDB Cloud providers when completing this request.
	// Read only field.
	Results *[]MesurementsDatabase `json:"results,omitempty"`
	// Total number of documents available. MongoDB Cloud omits this value if `includeCount` is set to `false`.
	// Read only field.
	TotalCount *int `json:"totalCount,omitempty"`
}

// NewPaginatedDatabase instantiates a new PaginatedDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedDatabase() *PaginatedDatabase {
	this := PaginatedDatabase{}
	return &this
}

// NewPaginatedDatabaseWithDefaults instantiates a new PaginatedDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedDatabaseWithDefaults() *PaginatedDatabase {
	this := PaginatedDatabase{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *PaginatedDatabase) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedDatabase) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginatedDatabase) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *PaginatedDatabase) SetLinks(v []Link) {
	o.Links = &v
}

// GetResults returns the Results field value if set, zero value otherwise
func (o *PaginatedDatabase) GetResults() []MesurementsDatabase {
	if o == nil || IsNil(o.Results) {
		var ret []MesurementsDatabase
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedDatabase) GetResultsOk() (*[]MesurementsDatabase, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}

	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedDatabase) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MesurementsDatabase and assigns it to the Results field.
func (o *PaginatedDatabase) SetResults(v []MesurementsDatabase) {
	o.Results = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise
func (o *PaginatedDatabase) GetTotalCount() int {
	if o == nil || IsNil(o.TotalCount) {
		var ret int
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedDatabase) GetTotalCountOk() (*int, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}

	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PaginatedDatabase) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int and assigns it to the TotalCount field.
func (o *PaginatedDatabase) SetTotalCount(v int) {
	o.TotalCount = &v
}

func (o PaginatedDatabase) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PaginatedDatabase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
