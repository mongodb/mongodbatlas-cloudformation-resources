// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// LiveMigrationRequest struct for LiveMigrationRequest
type LiveMigrationRequest struct {
	// Unique 24-hexadecimal digit string that identifies the migration request.
	// Read only field.
	Id          *string     `json:"_id,omitempty"`
	Destination Destination `json:"destination"`
	// Flag that indicates whether the migration process drops all collections from the destination cluster before the migration starts.
	// Write only field.
	DropEnabled bool `json:"dropEnabled"`
	// List of migration hosts used for this migration.
	MigrationHosts *[]string        `json:"migrationHosts,omitempty"`
	Sharding       *ShardingRequest `json:"sharding,omitempty"`
	Source         Source           `json:"source"`
}

// NewLiveMigrationRequest instantiates a new LiveMigrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveMigrationRequest(destination Destination, dropEnabled bool, source Source) *LiveMigrationRequest {
	this := LiveMigrationRequest{}
	this.Destination = destination
	this.DropEnabled = dropEnabled
	this.Source = source
	return &this
}

// NewLiveMigrationRequestWithDefaults instantiates a new LiveMigrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveMigrationRequestWithDefaults() *LiveMigrationRequest {
	this := LiveMigrationRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *LiveMigrationRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMigrationRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LiveMigrationRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LiveMigrationRequest) SetId(v string) {
	o.Id = &v
}

// GetDestination returns the Destination field value
func (o *LiveMigrationRequest) GetDestination() Destination {
	if o == nil {
		var ret Destination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *LiveMigrationRequest) GetDestinationOk() (*Destination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *LiveMigrationRequest) SetDestination(v Destination) {
	o.Destination = v
}

// GetDropEnabled returns the DropEnabled field value
func (o *LiveMigrationRequest) GetDropEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DropEnabled
}

// GetDropEnabledOk returns a tuple with the DropEnabled field value
// and a boolean to check if the value has been set.
func (o *LiveMigrationRequest) GetDropEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DropEnabled, true
}

// SetDropEnabled sets field value
func (o *LiveMigrationRequest) SetDropEnabled(v bool) {
	o.DropEnabled = v
}

// GetMigrationHosts returns the MigrationHosts field value if set, zero value otherwise
func (o *LiveMigrationRequest) GetMigrationHosts() []string {
	if o == nil || IsNil(o.MigrationHosts) {
		var ret []string
		return ret
	}
	return *o.MigrationHosts
}

// GetMigrationHostsOk returns a tuple with the MigrationHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMigrationRequest) GetMigrationHostsOk() (*[]string, bool) {
	if o == nil || IsNil(o.MigrationHosts) {
		return nil, false
	}

	return o.MigrationHosts, true
}

// HasMigrationHosts returns a boolean if a field has been set.
func (o *LiveMigrationRequest) HasMigrationHosts() bool {
	if o != nil && !IsNil(o.MigrationHosts) {
		return true
	}

	return false
}

// SetMigrationHosts gets a reference to the given []string and assigns it to the MigrationHosts field.
func (o *LiveMigrationRequest) SetMigrationHosts(v []string) {
	o.MigrationHosts = &v
}

// GetSharding returns the Sharding field value if set, zero value otherwise
func (o *LiveMigrationRequest) GetSharding() ShardingRequest {
	if o == nil || IsNil(o.Sharding) {
		var ret ShardingRequest
		return ret
	}
	return *o.Sharding
}

// GetShardingOk returns a tuple with the Sharding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMigrationRequest) GetShardingOk() (*ShardingRequest, bool) {
	if o == nil || IsNil(o.Sharding) {
		return nil, false
	}

	return o.Sharding, true
}

// HasSharding returns a boolean if a field has been set.
func (o *LiveMigrationRequest) HasSharding() bool {
	if o != nil && !IsNil(o.Sharding) {
		return true
	}

	return false
}

// SetSharding gets a reference to the given ShardingRequest and assigns it to the Sharding field.
func (o *LiveMigrationRequest) SetSharding(v ShardingRequest) {
	o.Sharding = &v
}

// GetSource returns the Source field value
func (o *LiveMigrationRequest) GetSource() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *LiveMigrationRequest) GetSourceOk() (*Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *LiveMigrationRequest) SetSource(v Source) {
	o.Source = v
}

func (o LiveMigrationRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o LiveMigrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination"] = o.Destination
	toSerialize["dropEnabled"] = o.DropEnabled
	if !IsNil(o.MigrationHosts) {
		toSerialize["migrationHosts"] = o.MigrationHosts
	}
	if !IsNil(o.Sharding) {
		toSerialize["sharding"] = o.Sharding
	}
	toSerialize["source"] = o.Source
	return toSerialize, nil
}
