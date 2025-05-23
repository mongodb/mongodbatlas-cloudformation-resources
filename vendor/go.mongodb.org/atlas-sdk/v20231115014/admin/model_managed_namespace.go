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

// ManagedNamespace struct for ManagedNamespace
type ManagedNamespace struct {
	// Human-readable label of the collection to manage for this Global Cluster.
	Collection *string `json:"collection,omitempty"`
	// Database parameter used to divide the *collection* into shards. Global clusters require a compound shard key. This compound shard key combines the location parameter and the user-selected custom key.
	CustomShardKey *string `json:"customShardKey,omitempty"`
	// Human-readable label of the database to manage for this Global Cluster.
	Db *string `json:"db,omitempty"`
	// Flag that indicates whether someone hashed the custom shard key. If this parameter returns `false`, this cluster uses ranged sharding.
	IsCustomShardKeyHashed *bool `json:"isCustomShardKeyHashed,omitempty"`
	// Flag that indicates whether the underlying index enforces unique values.
	IsShardKeyUnique *bool `json:"isShardKeyUnique,omitempty"`
	// Minimum number of chunks to create initially when sharding an empty collection with a hashed shard key.
	NumInitialChunks *int64 `json:"numInitialChunks,omitempty"`
	// Flag that indicates whether MongoDB Cloud should create and distribute initial chunks for an empty or non-existing collection. MongoDB Cloud distributes data based on the defined zones and zone ranges for the collection.
	PresplitHashedZones *bool `json:"presplitHashedZones,omitempty"`
}

// NewManagedNamespace instantiates a new ManagedNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedNamespace() *ManagedNamespace {
	this := ManagedNamespace{}
	var isCustomShardKeyHashed bool = false
	this.IsCustomShardKeyHashed = &isCustomShardKeyHashed
	var isShardKeyUnique bool = false
	this.IsShardKeyUnique = &isShardKeyUnique
	var presplitHashedZones bool = false
	this.PresplitHashedZones = &presplitHashedZones
	return &this
}

// NewManagedNamespaceWithDefaults instantiates a new ManagedNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedNamespaceWithDefaults() *ManagedNamespace {
	this := ManagedNamespace{}
	var isCustomShardKeyHashed bool = false
	this.IsCustomShardKeyHashed = &isCustomShardKeyHashed
	var isShardKeyUnique bool = false
	this.IsShardKeyUnique = &isShardKeyUnique
	var presplitHashedZones bool = false
	this.PresplitHashedZones = &presplitHashedZones
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise
func (o *ManagedNamespace) GetCollection() string {
	if o == nil || IsNil(o.Collection) {
		var ret string
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNamespace) GetCollectionOk() (*string, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}

	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *ManagedNamespace) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given string and assigns it to the Collection field.
func (o *ManagedNamespace) SetCollection(v string) {
	o.Collection = &v
}

// GetCustomShardKey returns the CustomShardKey field value if set, zero value otherwise
func (o *ManagedNamespace) GetCustomShardKey() string {
	if o == nil || IsNil(o.CustomShardKey) {
		var ret string
		return ret
	}
	return *o.CustomShardKey
}

// GetCustomShardKeyOk returns a tuple with the CustomShardKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNamespace) GetCustomShardKeyOk() (*string, bool) {
	if o == nil || IsNil(o.CustomShardKey) {
		return nil, false
	}

	return o.CustomShardKey, true
}

// HasCustomShardKey returns a boolean if a field has been set.
func (o *ManagedNamespace) HasCustomShardKey() bool {
	if o != nil && !IsNil(o.CustomShardKey) {
		return true
	}

	return false
}

// SetCustomShardKey gets a reference to the given string and assigns it to the CustomShardKey field.
func (o *ManagedNamespace) SetCustomShardKey(v string) {
	o.CustomShardKey = &v
}

// GetDb returns the Db field value if set, zero value otherwise
func (o *ManagedNamespace) GetDb() string {
	if o == nil || IsNil(o.Db) {
		var ret string
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNamespace) GetDbOk() (*string, bool) {
	if o == nil || IsNil(o.Db) {
		return nil, false
	}

	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *ManagedNamespace) HasDb() bool {
	if o != nil && !IsNil(o.Db) {
		return true
	}

	return false
}

// SetDb gets a reference to the given string and assigns it to the Db field.
func (o *ManagedNamespace) SetDb(v string) {
	o.Db = &v
}

// GetIsCustomShardKeyHashed returns the IsCustomShardKeyHashed field value if set, zero value otherwise
func (o *ManagedNamespace) GetIsCustomShardKeyHashed() bool {
	if o == nil || IsNil(o.IsCustomShardKeyHashed) {
		var ret bool
		return ret
	}
	return *o.IsCustomShardKeyHashed
}

// GetIsCustomShardKeyHashedOk returns a tuple with the IsCustomShardKeyHashed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNamespace) GetIsCustomShardKeyHashedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCustomShardKeyHashed) {
		return nil, false
	}

	return o.IsCustomShardKeyHashed, true
}

// HasIsCustomShardKeyHashed returns a boolean if a field has been set.
func (o *ManagedNamespace) HasIsCustomShardKeyHashed() bool {
	if o != nil && !IsNil(o.IsCustomShardKeyHashed) {
		return true
	}

	return false
}

// SetIsCustomShardKeyHashed gets a reference to the given bool and assigns it to the IsCustomShardKeyHashed field.
func (o *ManagedNamespace) SetIsCustomShardKeyHashed(v bool) {
	o.IsCustomShardKeyHashed = &v
}

// GetIsShardKeyUnique returns the IsShardKeyUnique field value if set, zero value otherwise
func (o *ManagedNamespace) GetIsShardKeyUnique() bool {
	if o == nil || IsNil(o.IsShardKeyUnique) {
		var ret bool
		return ret
	}
	return *o.IsShardKeyUnique
}

// GetIsShardKeyUniqueOk returns a tuple with the IsShardKeyUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNamespace) GetIsShardKeyUniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShardKeyUnique) {
		return nil, false
	}

	return o.IsShardKeyUnique, true
}

// HasIsShardKeyUnique returns a boolean if a field has been set.
func (o *ManagedNamespace) HasIsShardKeyUnique() bool {
	if o != nil && !IsNil(o.IsShardKeyUnique) {
		return true
	}

	return false
}

// SetIsShardKeyUnique gets a reference to the given bool and assigns it to the IsShardKeyUnique field.
func (o *ManagedNamespace) SetIsShardKeyUnique(v bool) {
	o.IsShardKeyUnique = &v
}

// GetNumInitialChunks returns the NumInitialChunks field value if set, zero value otherwise
func (o *ManagedNamespace) GetNumInitialChunks() int64 {
	if o == nil || IsNil(o.NumInitialChunks) {
		var ret int64
		return ret
	}
	return *o.NumInitialChunks
}

// GetNumInitialChunksOk returns a tuple with the NumInitialChunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNamespace) GetNumInitialChunksOk() (*int64, bool) {
	if o == nil || IsNil(o.NumInitialChunks) {
		return nil, false
	}

	return o.NumInitialChunks, true
}

// HasNumInitialChunks returns a boolean if a field has been set.
func (o *ManagedNamespace) HasNumInitialChunks() bool {
	if o != nil && !IsNil(o.NumInitialChunks) {
		return true
	}

	return false
}

// SetNumInitialChunks gets a reference to the given int64 and assigns it to the NumInitialChunks field.
func (o *ManagedNamespace) SetNumInitialChunks(v int64) {
	o.NumInitialChunks = &v
}

// GetPresplitHashedZones returns the PresplitHashedZones field value if set, zero value otherwise
func (o *ManagedNamespace) GetPresplitHashedZones() bool {
	if o == nil || IsNil(o.PresplitHashedZones) {
		var ret bool
		return ret
	}
	return *o.PresplitHashedZones
}

// GetPresplitHashedZonesOk returns a tuple with the PresplitHashedZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNamespace) GetPresplitHashedZonesOk() (*bool, bool) {
	if o == nil || IsNil(o.PresplitHashedZones) {
		return nil, false
	}

	return o.PresplitHashedZones, true
}

// HasPresplitHashedZones returns a boolean if a field has been set.
func (o *ManagedNamespace) HasPresplitHashedZones() bool {
	if o != nil && !IsNil(o.PresplitHashedZones) {
		return true
	}

	return false
}

// SetPresplitHashedZones gets a reference to the given bool and assigns it to the PresplitHashedZones field.
func (o *ManagedNamespace) SetPresplitHashedZones(v bool) {
	o.PresplitHashedZones = &v
}

func (o ManagedNamespace) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ManagedNamespace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.CustomShardKey) {
		toSerialize["customShardKey"] = o.CustomShardKey
	}
	if !IsNil(o.Db) {
		toSerialize["db"] = o.Db
	}
	if !IsNil(o.IsCustomShardKeyHashed) {
		toSerialize["isCustomShardKeyHashed"] = o.IsCustomShardKeyHashed
	}
	if !IsNil(o.IsShardKeyUnique) {
		toSerialize["isShardKeyUnique"] = o.IsShardKeyUnique
	}
	if !IsNil(o.NumInitialChunks) {
		toSerialize["numInitialChunks"] = o.NumInitialChunks
	}
	if !IsNil(o.PresplitHashedZones) {
		toSerialize["presplitHashedZones"] = o.PresplitHashedZones
	}
	return toSerialize, nil
}
