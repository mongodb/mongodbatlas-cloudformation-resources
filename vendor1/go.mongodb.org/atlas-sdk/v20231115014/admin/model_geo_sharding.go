// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// GeoSharding struct for GeoSharding
type GeoSharding struct {
	// List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.  This parameter returns an empty object if no custom zones exist.
	// Read only field.
	CustomZoneMapping *map[string]string `json:"customZoneMapping,omitempty"`
	// List that contains a namespace for a Global Cluster. MongoDB Cloud manages this cluster.
	// Read only field.
	ManagedNamespaces *[]ManagedNamespaces `json:"managedNamespaces,omitempty"`
	// Boolean that controls which management mode the Global Cluster is operating under. If this parameter is true Self-Managed Sharding is enabled and users are in control of the zone sharding within the Global Cluster. If this parameter is false Atlas-Managed Sharding is enabled and Atlas is control of zone sharding within the Global Cluster.
	// Read only field.
	SelfManagedSharding *bool `json:"selfManagedSharding,omitempty"`
}

// NewGeoSharding instantiates a new GeoSharding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoSharding() *GeoSharding {
	this := GeoSharding{}
	return &this
}

// NewGeoShardingWithDefaults instantiates a new GeoSharding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoShardingWithDefaults() *GeoSharding {
	this := GeoSharding{}
	return &this
}

// GetCustomZoneMapping returns the CustomZoneMapping field value if set, zero value otherwise
func (o *GeoSharding) GetCustomZoneMapping() map[string]string {
	if o == nil || IsNil(o.CustomZoneMapping) {
		var ret map[string]string
		return ret
	}
	return *o.CustomZoneMapping
}

// GetCustomZoneMappingOk returns a tuple with the CustomZoneMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoSharding) GetCustomZoneMappingOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomZoneMapping) {
		return nil, false
	}

	return o.CustomZoneMapping, true
}

// HasCustomZoneMapping returns a boolean if a field has been set.
func (o *GeoSharding) HasCustomZoneMapping() bool {
	if o != nil && !IsNil(o.CustomZoneMapping) {
		return true
	}

	return false
}

// SetCustomZoneMapping gets a reference to the given map[string]string and assigns it to the CustomZoneMapping field.
func (o *GeoSharding) SetCustomZoneMapping(v map[string]string) {
	o.CustomZoneMapping = &v
}

// GetManagedNamespaces returns the ManagedNamespaces field value if set, zero value otherwise
func (o *GeoSharding) GetManagedNamespaces() []ManagedNamespaces {
	if o == nil || IsNil(o.ManagedNamespaces) {
		var ret []ManagedNamespaces
		return ret
	}
	return *o.ManagedNamespaces
}

// GetManagedNamespacesOk returns a tuple with the ManagedNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoSharding) GetManagedNamespacesOk() (*[]ManagedNamespaces, bool) {
	if o == nil || IsNil(o.ManagedNamespaces) {
		return nil, false
	}

	return o.ManagedNamespaces, true
}

// HasManagedNamespaces returns a boolean if a field has been set.
func (o *GeoSharding) HasManagedNamespaces() bool {
	if o != nil && !IsNil(o.ManagedNamespaces) {
		return true
	}

	return false
}

// SetManagedNamespaces gets a reference to the given []ManagedNamespaces and assigns it to the ManagedNamespaces field.
func (o *GeoSharding) SetManagedNamespaces(v []ManagedNamespaces) {
	o.ManagedNamespaces = &v
}

// GetSelfManagedSharding returns the SelfManagedSharding field value if set, zero value otherwise
func (o *GeoSharding) GetSelfManagedSharding() bool {
	if o == nil || IsNil(o.SelfManagedSharding) {
		var ret bool
		return ret
	}
	return *o.SelfManagedSharding
}

// GetSelfManagedShardingOk returns a tuple with the SelfManagedSharding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoSharding) GetSelfManagedShardingOk() (*bool, bool) {
	if o == nil || IsNil(o.SelfManagedSharding) {
		return nil, false
	}

	return o.SelfManagedSharding, true
}

// HasSelfManagedSharding returns a boolean if a field has been set.
func (o *GeoSharding) HasSelfManagedSharding() bool {
	if o != nil && !IsNil(o.SelfManagedSharding) {
		return true
	}

	return false
}

// SetSelfManagedSharding gets a reference to the given bool and assigns it to the SelfManagedSharding field.
func (o *GeoSharding) SetSelfManagedSharding(v bool) {
	o.SelfManagedSharding = &v
}

func (o GeoSharding) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GeoSharding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
