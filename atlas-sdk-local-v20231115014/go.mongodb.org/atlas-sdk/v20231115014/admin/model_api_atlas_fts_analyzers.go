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

// ApiAtlasFTSAnalyzers Settings that describe one Atlas Search custom analyzer.
type ApiAtlasFTSAnalyzers struct {
	// Filters that examine text one character at a time and perform filtering operations.
	CharFilters *[]interface{} `json:"charFilters,omitempty"`
	// Human-readable name that identifies the custom analyzer. Names must be unique within an index, and must not start with any of the following strings: - `lucene.` - `builtin.` - `mongodb.`
	Name string `json:"name"`
	// Filter that performs operations such as:  - Stemming, which reduces related words, such as \"talking\", \"talked\", and \"talks\" to their root word \"talk\".  - Redaction, the removal of sensitive information from public documents.
	TokenFilters *[]interface{}                `json:"tokenFilters,omitempty"`
	Tokenizer    ApiAtlasFTSAnalyzersTokenizer `json:"tokenizer"`
}

// NewApiAtlasFTSAnalyzers instantiates a new ApiAtlasFTSAnalyzers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasFTSAnalyzers(name string, tokenizer ApiAtlasFTSAnalyzersTokenizer) *ApiAtlasFTSAnalyzers {
	this := ApiAtlasFTSAnalyzers{}
	this.Name = name
	this.Tokenizer = tokenizer
	return &this
}

// NewApiAtlasFTSAnalyzersWithDefaults instantiates a new ApiAtlasFTSAnalyzers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasFTSAnalyzersWithDefaults() *ApiAtlasFTSAnalyzers {
	this := ApiAtlasFTSAnalyzers{}
	return &this
}

// GetCharFilters returns the CharFilters field value if set, zero value otherwise
func (o *ApiAtlasFTSAnalyzers) GetCharFilters() []interface{} {
	if o == nil || IsNil(o.CharFilters) {
		var ret []interface{}
		return ret
	}
	return *o.CharFilters
}

// GetCharFiltersOk returns a tuple with the CharFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFTSAnalyzers) GetCharFiltersOk() (*[]interface{}, bool) {
	if o == nil || IsNil(o.CharFilters) {
		return nil, false
	}

	return o.CharFilters, true
}

// HasCharFilters returns a boolean if a field has been set.
func (o *ApiAtlasFTSAnalyzers) HasCharFilters() bool {
	if o != nil && !IsNil(o.CharFilters) {
		return true
	}

	return false
}

// SetCharFilters gets a reference to the given []interface{} and assigns it to the CharFilters field.
func (o *ApiAtlasFTSAnalyzers) SetCharFilters(v []interface{}) {
	o.CharFilters = &v
}

// GetName returns the Name field value
func (o *ApiAtlasFTSAnalyzers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasFTSAnalyzers) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiAtlasFTSAnalyzers) SetName(v string) {
	o.Name = v
}

// GetTokenFilters returns the TokenFilters field value if set, zero value otherwise
func (o *ApiAtlasFTSAnalyzers) GetTokenFilters() []interface{} {
	if o == nil || IsNil(o.TokenFilters) {
		var ret []interface{}
		return ret
	}
	return *o.TokenFilters
}

// GetTokenFiltersOk returns a tuple with the TokenFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasFTSAnalyzers) GetTokenFiltersOk() (*[]interface{}, bool) {
	if o == nil || IsNil(o.TokenFilters) {
		return nil, false
	}

	return o.TokenFilters, true
}

// HasTokenFilters returns a boolean if a field has been set.
func (o *ApiAtlasFTSAnalyzers) HasTokenFilters() bool {
	if o != nil && !IsNil(o.TokenFilters) {
		return true
	}

	return false
}

// SetTokenFilters gets a reference to the given []interface{} and assigns it to the TokenFilters field.
func (o *ApiAtlasFTSAnalyzers) SetTokenFilters(v []interface{}) {
	o.TokenFilters = &v
}

// GetTokenizer returns the Tokenizer field value
func (o *ApiAtlasFTSAnalyzers) GetTokenizer() ApiAtlasFTSAnalyzersTokenizer {
	if o == nil {
		var ret ApiAtlasFTSAnalyzersTokenizer
		return ret
	}

	return o.Tokenizer
}

// GetTokenizerOk returns a tuple with the Tokenizer field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasFTSAnalyzers) GetTokenizerOk() (*ApiAtlasFTSAnalyzersTokenizer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tokenizer, true
}

// SetTokenizer sets field value
func (o *ApiAtlasFTSAnalyzers) SetTokenizer(v ApiAtlasFTSAnalyzersTokenizer) {
	o.Tokenizer = v
}

func (o ApiAtlasFTSAnalyzers) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ApiAtlasFTSAnalyzers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CharFilters) {
		toSerialize["charFilters"] = o.CharFilters
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.TokenFilters) {
		toSerialize["tokenFilters"] = o.TokenFilters
	}
	toSerialize["tokenizer"] = o.Tokenizer
	return toSerialize, nil
}
