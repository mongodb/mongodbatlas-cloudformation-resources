// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ShardingRequest Document that configures sharding on the destination cluster when migrating from a replica set source to a sharded cluster destination on MongoDB 6.0 or higher. If you don't wish to shard any collections on the destination cluster, leave this empty.
type ShardingRequest struct {
	// Flag that lets the migration create supporting indexes for the shard keys, if none exists, as the destination cluster also needs compatible indexes for the specified shard keys.
	// Write only field.
	CreateSupportingIndexes bool `json:"createSupportingIndexes"`
	// List of shard configurations to shard destination collections. Atlas shards only those collections that you include in the sharding entries array.
	// Write only field.
	ShardingEntries *[]ShardEntry `json:"shardingEntries,omitempty"`
}

// NewShardingRequest instantiates a new ShardingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardingRequest(createSupportingIndexes bool) *ShardingRequest {
	this := ShardingRequest{}
	this.CreateSupportingIndexes = createSupportingIndexes
	return &this
}

// NewShardingRequestWithDefaults instantiates a new ShardingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardingRequestWithDefaults() *ShardingRequest {
	this := ShardingRequest{}
	return &this
}

// GetCreateSupportingIndexes returns the CreateSupportingIndexes field value
func (o *ShardingRequest) GetCreateSupportingIndexes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CreateSupportingIndexes
}

// GetCreateSupportingIndexesOk returns a tuple with the CreateSupportingIndexes field value
// and a boolean to check if the value has been set.
func (o *ShardingRequest) GetCreateSupportingIndexesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateSupportingIndexes, true
}

// SetCreateSupportingIndexes sets field value
func (o *ShardingRequest) SetCreateSupportingIndexes(v bool) {
	o.CreateSupportingIndexes = v
}

// GetShardingEntries returns the ShardingEntries field value if set, zero value otherwise
func (o *ShardingRequest) GetShardingEntries() []ShardEntry {
	if o == nil || IsNil(o.ShardingEntries) {
		var ret []ShardEntry
		return ret
	}
	return *o.ShardingEntries
}

// GetShardingEntriesOk returns a tuple with the ShardingEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRequest) GetShardingEntriesOk() (*[]ShardEntry, bool) {
	if o == nil || IsNil(o.ShardingEntries) {
		return nil, false
	}

	return o.ShardingEntries, true
}

// HasShardingEntries returns a boolean if a field has been set.
func (o *ShardingRequest) HasShardingEntries() bool {
	if o != nil && !IsNil(o.ShardingEntries) {
		return true
	}

	return false
}

// SetShardingEntries gets a reference to the given []ShardEntry and assigns it to the ShardingEntries field.
func (o *ShardingRequest) SetShardingEntries(v []ShardEntry) {
	o.ShardingEntries = &v
}

func (o ShardingRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ShardingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createSupportingIndexes"] = o.CreateSupportingIndexes
	if !IsNil(o.ShardingEntries) {
		toSerialize["shardingEntries"] = o.ShardingEntries
	}
	return toSerialize, nil
}
