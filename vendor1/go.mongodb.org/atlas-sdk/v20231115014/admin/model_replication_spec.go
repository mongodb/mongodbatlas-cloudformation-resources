// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ReplicationSpec Details that explain how MongoDB Cloud replicates data on the specified MongoDB database.
type ReplicationSpec struct {
	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a Multi-Cloud Cluster. If you include existing zones in the request, you must specify this parameter. If you add a new zone to an existing Multi-Cloud Cluster, you may specify this parameter. The request deletes any existing zones in the Multi-Cloud Cluster that you exclude from the request.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// Positive integer that specifies the number of shards to deploy in each specified zone. If you set this value to `1` and **clusterType** is `SHARDED`, MongoDB Cloud deploys a single-shard sharded cluster. Don't create a sharded cluster with a single shard for production environments. Single-shard sharded clusters don't provide the same benefits as multi-shard configurations.   If you are upgrading a replica set to a sharded cluster, you cannot increase the number of shards in the same update request.  You should wait until after the cluster has completed upgrading to sharded and you have reconnected all application clients to the MongoDB router before adding additional shards. Otherwise, your data might become inconsistent once MongoDB Cloud begins distributing data across shards.
	NumShards *int `json:"numShards,omitempty"`
	// Hardware specifications for nodes set for a given region. Each **regionConfigs** object describes the region's priority in elections and the number and type of MongoDB nodes that MongoDB Cloud deploys to the region. Each **regionConfigs** object must have either an **analyticsSpecs** object, **electableSpecs** object, or **readOnlySpecs** object. Tenant clusters only require **electableSpecs. Dedicated** clusters can specify any of these specifications, but must have at least one **electableSpecs** object within a **replicationSpec**. Every hardware specification must use the same **instanceSize**.  **Example:**  If you set `\"replicationSpecs[n].regionConfigs[m].analyticsSpecs.instanceSize\" : \"M30\"`, set `\"replicationSpecs[n].regionConfigs[m].electableSpecs.instanceSize\" : `\"M30\"` if you have electable nodes and `\"replicationSpecs[n].regionConfigs[m].readOnlySpecs.instanceSize\" : `\"M30\"` if you have read-only nodes.
	RegionConfigs *[]CloudRegionConfig `json:"regionConfigs,omitempty"`
	// Human-readable label that identifies the zone in a Global Cluster. Provide this value only if `\"clusterType\" : \"GEOSHARDED\"`.
	ZoneName *string `json:"zoneName,omitempty"`
}

// NewReplicationSpec instantiates a new ReplicationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationSpec() *ReplicationSpec {
	this := ReplicationSpec{}
	return &this
}

// NewReplicationSpecWithDefaults instantiates a new ReplicationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationSpecWithDefaults() *ReplicationSpec {
	this := ReplicationSpec{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *ReplicationSpec) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationSpec) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReplicationSpec) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReplicationSpec) SetId(v string) {
	o.Id = &v
}

// GetNumShards returns the NumShards field value if set, zero value otherwise
func (o *ReplicationSpec) GetNumShards() int {
	if o == nil || IsNil(o.NumShards) {
		var ret int
		return ret
	}
	return *o.NumShards
}

// GetNumShardsOk returns a tuple with the NumShards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationSpec) GetNumShardsOk() (*int, bool) {
	if o == nil || IsNil(o.NumShards) {
		return nil, false
	}

	return o.NumShards, true
}

// HasNumShards returns a boolean if a field has been set.
func (o *ReplicationSpec) HasNumShards() bool {
	if o != nil && !IsNil(o.NumShards) {
		return true
	}

	return false
}

// SetNumShards gets a reference to the given int and assigns it to the NumShards field.
func (o *ReplicationSpec) SetNumShards(v int) {
	o.NumShards = &v
}

// GetRegionConfigs returns the RegionConfigs field value if set, zero value otherwise
func (o *ReplicationSpec) GetRegionConfigs() []CloudRegionConfig {
	if o == nil || IsNil(o.RegionConfigs) {
		var ret []CloudRegionConfig
		return ret
	}
	return *o.RegionConfigs
}

// GetRegionConfigsOk returns a tuple with the RegionConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationSpec) GetRegionConfigsOk() (*[]CloudRegionConfig, bool) {
	if o == nil || IsNil(o.RegionConfigs) {
		return nil, false
	}

	return o.RegionConfigs, true
}

// HasRegionConfigs returns a boolean if a field has been set.
func (o *ReplicationSpec) HasRegionConfigs() bool {
	if o != nil && !IsNil(o.RegionConfigs) {
		return true
	}

	return false
}

// SetRegionConfigs gets a reference to the given []CloudRegionConfig and assigns it to the RegionConfigs field.
func (o *ReplicationSpec) SetRegionConfigs(v []CloudRegionConfig) {
	o.RegionConfigs = &v
}

// GetZoneName returns the ZoneName field value if set, zero value otherwise
func (o *ReplicationSpec) GetZoneName() string {
	if o == nil || IsNil(o.ZoneName) {
		var ret string
		return ret
	}
	return *o.ZoneName
}

// GetZoneNameOk returns a tuple with the ZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationSpec) GetZoneNameOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneName) {
		return nil, false
	}

	return o.ZoneName, true
}

// HasZoneName returns a boolean if a field has been set.
func (o *ReplicationSpec) HasZoneName() bool {
	if o != nil && !IsNil(o.ZoneName) {
		return true
	}

	return false
}

// SetZoneName gets a reference to the given string and assigns it to the ZoneName field.
func (o *ReplicationSpec) SetZoneName(v string) {
	o.ZoneName = &v
}

func (o ReplicationSpec) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ReplicationSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumShards) {
		toSerialize["numShards"] = o.NumShards
	}
	if !IsNil(o.RegionConfigs) {
		toSerialize["regionConfigs"] = o.RegionConfigs
	}
	if !IsNil(o.ZoneName) {
		toSerialize["zoneName"] = o.ZoneName
	}
	return toSerialize, nil
}
