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

// IndexOptions One or more settings that determine how the MongoDB Cloud creates this MongoDB index.
type IndexOptions struct {
	// Index version number applied to the 2dsphere index. MongoDB 3.2 and later use version 3. Use this option to override the default version number. This option applies to the **2dsphere** index type only.
	Var2dsphereIndexVersion *int `json:"2dsphereIndexVersion,omitempty"`
	// Flag that indicates whether MongoDB should build the index in the background. This applies to MongoDB databases running feature compatibility version 4.0 or earlier. MongoDB databases running FCV 4.2 or later build indexes using an optimized build process. This process holds the exclusive lock only at the beginning and end of the build process. The rest of the build process yields to interleaving read and write operations. MongoDB databases running FCV 4.2 or later ignore this option. This option applies to all index types.
	Background *bool `json:"background,omitempty"`
	// Number of precision applied to the stored geohash value of the location data. This option applies to the **2d** index type only.
	Bits *int `json:"bits,omitempty"`
	// Number of units within which to group the location values. You could group in the same bucket those location values within the specified number of units to each other. This option applies to the geoHaystack index type only.  MongoDB 5.0 removed geoHaystack Indexes and the `geoSearch` command.
	BucketSize *int `json:"bucketSize,omitempty"`
	// The columnstoreProjection document allows to include or exclude subschemas schema. One cannot combine inclusion and exclusion statements. Accordingly, the <value> can be either of the following: 1 or true to include the field and recursively all fields it is a prefix of in the index 0 or false to exclude the field and recursively all fields it is a prefix of from the index.
	ColumnstoreProjection *map[string]int `json:"columnstoreProjection,omitempty"`
	// Human language that determines the list of stop words and the rules for the stemmer and tokenizer. This option accepts the supported languages using its name in lowercase english or the ISO 639-2 code. If you set this parameter to `\"none\"`, then the text search uses simple tokenization with no list of stop words and no stemming. This option applies to the **text** index type only.
	DefaultLanguage *string `json:"default_language,omitempty"`
	// Number of seconds that MongoDB retains documents in a Time To Live (TTL) index.
	ExpireAfterSeconds *int `json:"expireAfterSeconds,omitempty"`
	// Flag that determines whether the index is hidden from the query planner. A hidden index is not evaluated as part of the query plan selection.
	Hidden *bool `json:"hidden,omitempty"`
	// Human-readable label that identifies the document parameter that contains the override language for the document. This option applies to the **text** index type only.
	LanguageOverride *string `json:"language_override,omitempty"`
	// Upper inclusive boundary to limit the longitude and latitude values. This option applies to the 2d index type only.
	Max *int `json:"max,omitempty"`
	// Lower inclusive boundary to limit the longitude and latitude values. This option applies to the 2d index type only.
	Min *int `json:"min,omitempty"`
	// Human-readable label that identifies this index. This option applies to all index types.
	Name *string `json:"name,omitempty"`
	// Rules that limit the documents that the index references to a filter expression. All MongoDB index types accept a **partialFilterExpression** option. **partialFilterExpression** can include following expressions:  - equality (`\"parameter\" : \"value\"` or using the `$eq` operator) - `\"$exists\": true` , maximum: `$gt`, `$gte`, `$lt`, `$lte` comparisons - `$type` - `$and` (top-level only)  This option applies to all index types.
	PartialFilterExpression map[string]interface{} `json:"partialFilterExpression,omitempty"`
	// Flag that indicates whether the index references documents that only have the specified parameter. These indexes use less space but behave differently in some situations like when sorting. The following index types default to sparse and ignore this option: `2dsphere`, `2d`, `geoHaystack`, `text`.  Compound indexes that includes one or more indexes with `2dsphere` keys alongside other key types, only the `2dsphere` index parameters determine which documents the index references. If you run MongoDB 3.2 or later, use partial indexes. This option applies to all index types.
	Sparse *bool `json:"sparse,omitempty"`
	// Storage engine set for the specific index. This value can be set only at creation. This option uses the following format: `\"storageEngine\" : { \"<storage-engine-name>\" : \"<options>\" }` MongoDB validates storage engine configuration options when creating indexes. To support replica sets with members with different storage engines, MongoDB logs these options to the oplog during replication. This option applies to all index types.
	StorageEngine map[string]interface{} `json:"storageEngine,omitempty"`
	// Version applied to this text index. MongoDB 3.2 and later use version `3`. Use this option to override the default version number. This option applies to the **text** index type only.
	TextIndexVersion *int `json:"textIndexVersion,omitempty"`
	// Flag that indicates whether this index can accept insertion or update of documents when the index key value matches an existing index key value. Set `\"unique\" : true` to set this index as unique. You can't set a hashed index to be unique. This option applies to all index types. This option is unsupported for rolling indexes.
	Unique *bool `json:"unique,omitempty"`
	// Relative importance to place upon provided index parameters. This object expresses this as key/value pairs of index parameter and weight to apply to that parameter. You can specify weights for some or all the indexed parameters. The weight must be an integer between 1 and 99,999. MongoDB 5.0 and later can apply **weights** to **text** indexes only.
	Weights map[string]interface{} `json:"weights,omitempty"`
}

// NewIndexOptions instantiates a new IndexOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexOptions() *IndexOptions {
	this := IndexOptions{}
	var var2dsphereIndexVersion int = 3
	this.Var2dsphereIndexVersion = &var2dsphereIndexVersion
	var background bool = false
	this.Background = &background
	var bits int = 26
	this.Bits = &bits
	var defaultLanguage string = "english"
	this.DefaultLanguage = &defaultLanguage
	var hidden bool = false
	this.Hidden = &hidden
	var languageOverride string = "language"
	this.LanguageOverride = &languageOverride
	var max int = 180
	this.Max = &max
	var min int = -180
	this.Min = &min
	var sparse bool = false
	this.Sparse = &sparse
	var textIndexVersion int = 3
	this.TextIndexVersion = &textIndexVersion
	var unique bool = false
	this.Unique = &unique
	return &this
}

// NewIndexOptionsWithDefaults instantiates a new IndexOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexOptionsWithDefaults() *IndexOptions {
	this := IndexOptions{}
	var var2dsphereIndexVersion int = 3
	this.Var2dsphereIndexVersion = &var2dsphereIndexVersion
	var background bool = false
	this.Background = &background
	var bits int = 26
	this.Bits = &bits
	var defaultLanguage string = "english"
	this.DefaultLanguage = &defaultLanguage
	var hidden bool = false
	this.Hidden = &hidden
	var languageOverride string = "language"
	this.LanguageOverride = &languageOverride
	var max int = 180
	this.Max = &max
	var min int = -180
	this.Min = &min
	var sparse bool = false
	this.Sparse = &sparse
	var textIndexVersion int = 3
	this.TextIndexVersion = &textIndexVersion
	var unique bool = false
	this.Unique = &unique
	return &this
}

// GetVar2dsphereIndexVersion returns the Var2dsphereIndexVersion field value if set, zero value otherwise
func (o *IndexOptions) GetVar2dsphereIndexVersion() int {
	if o == nil || IsNil(o.Var2dsphereIndexVersion) {
		var ret int
		return ret
	}
	return *o.Var2dsphereIndexVersion
}

// GetVar2dsphereIndexVersionOk returns a tuple with the Var2dsphereIndexVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetVar2dsphereIndexVersionOk() (*int, bool) {
	if o == nil || IsNil(o.Var2dsphereIndexVersion) {
		return nil, false
	}

	return o.Var2dsphereIndexVersion, true
}

// HasVar2dsphereIndexVersion returns a boolean if a field has been set.
func (o *IndexOptions) HasVar2dsphereIndexVersion() bool {
	if o != nil && !IsNil(o.Var2dsphereIndexVersion) {
		return true
	}

	return false
}

// SetVar2dsphereIndexVersion gets a reference to the given int and assigns it to the Var2dsphereIndexVersion field.
func (o *IndexOptions) SetVar2dsphereIndexVersion(v int) {
	o.Var2dsphereIndexVersion = &v
}

// GetBackground returns the Background field value if set, zero value otherwise
func (o *IndexOptions) GetBackground() bool {
	if o == nil || IsNil(o.Background) {
		var ret bool
		return ret
	}
	return *o.Background
}

// GetBackgroundOk returns a tuple with the Background field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetBackgroundOk() (*bool, bool) {
	if o == nil || IsNil(o.Background) {
		return nil, false
	}

	return o.Background, true
}

// HasBackground returns a boolean if a field has been set.
func (o *IndexOptions) HasBackground() bool {
	if o != nil && !IsNil(o.Background) {
		return true
	}

	return false
}

// SetBackground gets a reference to the given bool and assigns it to the Background field.
func (o *IndexOptions) SetBackground(v bool) {
	o.Background = &v
}

// GetBits returns the Bits field value if set, zero value otherwise
func (o *IndexOptions) GetBits() int {
	if o == nil || IsNil(o.Bits) {
		var ret int
		return ret
	}
	return *o.Bits
}

// GetBitsOk returns a tuple with the Bits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetBitsOk() (*int, bool) {
	if o == nil || IsNil(o.Bits) {
		return nil, false
	}

	return o.Bits, true
}

// HasBits returns a boolean if a field has been set.
func (o *IndexOptions) HasBits() bool {
	if o != nil && !IsNil(o.Bits) {
		return true
	}

	return false
}

// SetBits gets a reference to the given int and assigns it to the Bits field.
func (o *IndexOptions) SetBits(v int) {
	o.Bits = &v
}

// GetBucketSize returns the BucketSize field value if set, zero value otherwise
func (o *IndexOptions) GetBucketSize() int {
	if o == nil || IsNil(o.BucketSize) {
		var ret int
		return ret
	}
	return *o.BucketSize
}

// GetBucketSizeOk returns a tuple with the BucketSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetBucketSizeOk() (*int, bool) {
	if o == nil || IsNil(o.BucketSize) {
		return nil, false
	}

	return o.BucketSize, true
}

// HasBucketSize returns a boolean if a field has been set.
func (o *IndexOptions) HasBucketSize() bool {
	if o != nil && !IsNil(o.BucketSize) {
		return true
	}

	return false
}

// SetBucketSize gets a reference to the given int and assigns it to the BucketSize field.
func (o *IndexOptions) SetBucketSize(v int) {
	o.BucketSize = &v
}

// GetColumnstoreProjection returns the ColumnstoreProjection field value if set, zero value otherwise
func (o *IndexOptions) GetColumnstoreProjection() map[string]int {
	if o == nil || IsNil(o.ColumnstoreProjection) {
		var ret map[string]int
		return ret
	}
	return *o.ColumnstoreProjection
}

// GetColumnstoreProjectionOk returns a tuple with the ColumnstoreProjection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetColumnstoreProjectionOk() (*map[string]int, bool) {
	if o == nil || IsNil(o.ColumnstoreProjection) {
		return nil, false
	}

	return o.ColumnstoreProjection, true
}

// HasColumnstoreProjection returns a boolean if a field has been set.
func (o *IndexOptions) HasColumnstoreProjection() bool {
	if o != nil && !IsNil(o.ColumnstoreProjection) {
		return true
	}

	return false
}

// SetColumnstoreProjection gets a reference to the given map[string]int and assigns it to the ColumnstoreProjection field.
func (o *IndexOptions) SetColumnstoreProjection(v map[string]int) {
	o.ColumnstoreProjection = &v
}

// GetDefaultLanguage returns the DefaultLanguage field value if set, zero value otherwise
func (o *IndexOptions) GetDefaultLanguage() string {
	if o == nil || IsNil(o.DefaultLanguage) {
		var ret string
		return ret
	}
	return *o.DefaultLanguage
}

// GetDefaultLanguageOk returns a tuple with the DefaultLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetDefaultLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultLanguage) {
		return nil, false
	}

	return o.DefaultLanguage, true
}

// HasDefaultLanguage returns a boolean if a field has been set.
func (o *IndexOptions) HasDefaultLanguage() bool {
	if o != nil && !IsNil(o.DefaultLanguage) {
		return true
	}

	return false
}

// SetDefaultLanguage gets a reference to the given string and assigns it to the DefaultLanguage field.
func (o *IndexOptions) SetDefaultLanguage(v string) {
	o.DefaultLanguage = &v
}

// GetExpireAfterSeconds returns the ExpireAfterSeconds field value if set, zero value otherwise
func (o *IndexOptions) GetExpireAfterSeconds() int {
	if o == nil || IsNil(o.ExpireAfterSeconds) {
		var ret int
		return ret
	}
	return *o.ExpireAfterSeconds
}

// GetExpireAfterSecondsOk returns a tuple with the ExpireAfterSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetExpireAfterSecondsOk() (*int, bool) {
	if o == nil || IsNil(o.ExpireAfterSeconds) {
		return nil, false
	}

	return o.ExpireAfterSeconds, true
}

// HasExpireAfterSeconds returns a boolean if a field has been set.
func (o *IndexOptions) HasExpireAfterSeconds() bool {
	if o != nil && !IsNil(o.ExpireAfterSeconds) {
		return true
	}

	return false
}

// SetExpireAfterSeconds gets a reference to the given int and assigns it to the ExpireAfterSeconds field.
func (o *IndexOptions) SetExpireAfterSeconds(v int) {
	o.ExpireAfterSeconds = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise
func (o *IndexOptions) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}

	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *IndexOptions) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *IndexOptions) SetHidden(v bool) {
	o.Hidden = &v
}

// GetLanguageOverride returns the LanguageOverride field value if set, zero value otherwise
func (o *IndexOptions) GetLanguageOverride() string {
	if o == nil || IsNil(o.LanguageOverride) {
		var ret string
		return ret
	}
	return *o.LanguageOverride
}

// GetLanguageOverrideOk returns a tuple with the LanguageOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetLanguageOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageOverride) {
		return nil, false
	}

	return o.LanguageOverride, true
}

// HasLanguageOverride returns a boolean if a field has been set.
func (o *IndexOptions) HasLanguageOverride() bool {
	if o != nil && !IsNil(o.LanguageOverride) {
		return true
	}

	return false
}

// SetLanguageOverride gets a reference to the given string and assigns it to the LanguageOverride field.
func (o *IndexOptions) SetLanguageOverride(v string) {
	o.LanguageOverride = &v
}

// GetMax returns the Max field value if set, zero value otherwise
func (o *IndexOptions) GetMax() int {
	if o == nil || IsNil(o.Max) {
		var ret int
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetMaxOk() (*int, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}

	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *IndexOptions) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int and assigns it to the Max field.
func (o *IndexOptions) SetMax(v int) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise
func (o *IndexOptions) GetMin() int {
	if o == nil || IsNil(o.Min) {
		var ret int
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetMinOk() (*int, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}

	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *IndexOptions) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int and assigns it to the Min field.
func (o *IndexOptions) SetMin(v int) {
	o.Min = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *IndexOptions) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IndexOptions) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IndexOptions) SetName(v string) {
	o.Name = &v
}

// GetPartialFilterExpression returns the PartialFilterExpression field value if set, zero value otherwise
func (o *IndexOptions) GetPartialFilterExpression() map[string]interface{} {
	if o == nil || IsNil(o.PartialFilterExpression) {
		var ret map[string]interface{}
		return ret
	}
	return o.PartialFilterExpression
}

// GetPartialFilterExpressionOk returns a tuple with the PartialFilterExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetPartialFilterExpressionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PartialFilterExpression) {
		return map[string]interface{}{}, false
	}

	return o.PartialFilterExpression, true
}

// HasPartialFilterExpression returns a boolean if a field has been set.
func (o *IndexOptions) HasPartialFilterExpression() bool {
	if o != nil && !IsNil(o.PartialFilterExpression) {
		return true
	}

	return false
}

// SetPartialFilterExpression gets a reference to the given map[string]interface{} and assigns it to the PartialFilterExpression field.
func (o *IndexOptions) SetPartialFilterExpression(v map[string]interface{}) {
	o.PartialFilterExpression = v
}

// GetSparse returns the Sparse field value if set, zero value otherwise
func (o *IndexOptions) GetSparse() bool {
	if o == nil || IsNil(o.Sparse) {
		var ret bool
		return ret
	}
	return *o.Sparse
}

// GetSparseOk returns a tuple with the Sparse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetSparseOk() (*bool, bool) {
	if o == nil || IsNil(o.Sparse) {
		return nil, false
	}

	return o.Sparse, true
}

// HasSparse returns a boolean if a field has been set.
func (o *IndexOptions) HasSparse() bool {
	if o != nil && !IsNil(o.Sparse) {
		return true
	}

	return false
}

// SetSparse gets a reference to the given bool and assigns it to the Sparse field.
func (o *IndexOptions) SetSparse(v bool) {
	o.Sparse = &v
}

// GetStorageEngine returns the StorageEngine field value if set, zero value otherwise
func (o *IndexOptions) GetStorageEngine() map[string]interface{} {
	if o == nil || IsNil(o.StorageEngine) {
		var ret map[string]interface{}
		return ret
	}
	return o.StorageEngine
}

// GetStorageEngineOk returns a tuple with the StorageEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetStorageEngineOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.StorageEngine) {
		return map[string]interface{}{}, false
	}

	return o.StorageEngine, true
}

// HasStorageEngine returns a boolean if a field has been set.
func (o *IndexOptions) HasStorageEngine() bool {
	if o != nil && !IsNil(o.StorageEngine) {
		return true
	}

	return false
}

// SetStorageEngine gets a reference to the given map[string]interface{} and assigns it to the StorageEngine field.
func (o *IndexOptions) SetStorageEngine(v map[string]interface{}) {
	o.StorageEngine = v
}

// GetTextIndexVersion returns the TextIndexVersion field value if set, zero value otherwise
func (o *IndexOptions) GetTextIndexVersion() int {
	if o == nil || IsNil(o.TextIndexVersion) {
		var ret int
		return ret
	}
	return *o.TextIndexVersion
}

// GetTextIndexVersionOk returns a tuple with the TextIndexVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetTextIndexVersionOk() (*int, bool) {
	if o == nil || IsNil(o.TextIndexVersion) {
		return nil, false
	}

	return o.TextIndexVersion, true
}

// HasTextIndexVersion returns a boolean if a field has been set.
func (o *IndexOptions) HasTextIndexVersion() bool {
	if o != nil && !IsNil(o.TextIndexVersion) {
		return true
	}

	return false
}

// SetTextIndexVersion gets a reference to the given int and assigns it to the TextIndexVersion field.
func (o *IndexOptions) SetTextIndexVersion(v int) {
	o.TextIndexVersion = &v
}

// GetUnique returns the Unique field value if set, zero value otherwise
func (o *IndexOptions) GetUnique() bool {
	if o == nil || IsNil(o.Unique) {
		var ret bool
		return ret
	}
	return *o.Unique
}

// GetUniqueOk returns a tuple with the Unique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetUniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.Unique) {
		return nil, false
	}

	return o.Unique, true
}

// HasUnique returns a boolean if a field has been set.
func (o *IndexOptions) HasUnique() bool {
	if o != nil && !IsNil(o.Unique) {
		return true
	}

	return false
}

// SetUnique gets a reference to the given bool and assigns it to the Unique field.
func (o *IndexOptions) SetUnique(v bool) {
	o.Unique = &v
}

// GetWeights returns the Weights field value if set, zero value otherwise
func (o *IndexOptions) GetWeights() map[string]interface{} {
	if o == nil || IsNil(o.Weights) {
		var ret map[string]interface{}
		return ret
	}
	return o.Weights
}

// GetWeightsOk returns a tuple with the Weights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetWeightsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Weights) {
		return map[string]interface{}{}, false
	}

	return o.Weights, true
}

// HasWeights returns a boolean if a field has been set.
func (o *IndexOptions) HasWeights() bool {
	if o != nil && !IsNil(o.Weights) {
		return true
	}

	return false
}

// SetWeights gets a reference to the given map[string]interface{} and assigns it to the Weights field.
func (o *IndexOptions) SetWeights(v map[string]interface{}) {
	o.Weights = v
}

func (o IndexOptions) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o IndexOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var2dsphereIndexVersion) {
		toSerialize["2dsphereIndexVersion"] = o.Var2dsphereIndexVersion
	}
	if !IsNil(o.Background) {
		toSerialize["background"] = o.Background
	}
	if !IsNil(o.Bits) {
		toSerialize["bits"] = o.Bits
	}
	if !IsNil(o.BucketSize) {
		toSerialize["bucketSize"] = o.BucketSize
	}
	if !IsNil(o.ColumnstoreProjection) {
		toSerialize["columnstoreProjection"] = o.ColumnstoreProjection
	}
	if !IsNil(o.DefaultLanguage) {
		toSerialize["default_language"] = o.DefaultLanguage
	}
	if !IsNil(o.ExpireAfterSeconds) {
		toSerialize["expireAfterSeconds"] = o.ExpireAfterSeconds
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.LanguageOverride) {
		toSerialize["language_override"] = o.LanguageOverride
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PartialFilterExpression) {
		toSerialize["partialFilterExpression"] = o.PartialFilterExpression
	}
	if !IsNil(o.Sparse) {
		toSerialize["sparse"] = o.Sparse
	}
	if !IsNil(o.StorageEngine) {
		toSerialize["storageEngine"] = o.StorageEngine
	}
	if !IsNil(o.TextIndexVersion) {
		toSerialize["textIndexVersion"] = o.TextIndexVersion
	}
	if !IsNil(o.Unique) {
		toSerialize["unique"] = o.Unique
	}
	if !IsNil(o.Weights) {
		toSerialize["weights"] = o.Weights
	}
	return toSerialize, nil
}
