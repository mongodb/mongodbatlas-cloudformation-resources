// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// CreateEndpointRequest struct for CreateEndpointRequest
type CreateEndpointRequest struct {
	// Unique string that identifies the private endpoint's network interface that someone added to this private endpoint service.
	// Write only field.
	Id *string `json:"id,omitempty"`
	// IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.
	PrivateEndpointIPAddress *string `json:"privateEndpointIPAddress,omitempty"`
	// Human-readable label that identifies a set of endpoints.
	// Write only field.
	EndpointGroupName *string `json:"endpointGroupName,omitempty"`
	// List of individual private endpoints that comprise this endpoint group.
	Endpoints *[]CreateGCPForwardingRuleRequest `json:"endpoints,omitempty"`
	// Unique string that identifies the Google Cloud project in which you created the endpoints.
	// Write only field.
	GcpProjectId *string `json:"gcpProjectId,omitempty"`
}

// NewCreateEndpointRequest instantiates a new CreateEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEndpointRequest() *CreateEndpointRequest {
	this := CreateEndpointRequest{}
	return &this
}

// NewCreateEndpointRequestWithDefaults instantiates a new CreateEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEndpointRequestWithDefaults() *CreateEndpointRequest {
	this := CreateEndpointRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *CreateEndpointRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEndpointRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateEndpointRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateEndpointRequest) SetId(v string) {
	o.Id = &v
}

// GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field value if set, zero value otherwise
func (o *CreateEndpointRequest) GetPrivateEndpointIPAddress() string {
	if o == nil || IsNil(o.PrivateEndpointIPAddress) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointIPAddress
}

// GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEndpointRequest) GetPrivateEndpointIPAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateEndpointIPAddress) {
		return nil, false
	}

	return o.PrivateEndpointIPAddress, true
}

// HasPrivateEndpointIPAddress returns a boolean if a field has been set.
func (o *CreateEndpointRequest) HasPrivateEndpointIPAddress() bool {
	if o != nil && !IsNil(o.PrivateEndpointIPAddress) {
		return true
	}

	return false
}

// SetPrivateEndpointIPAddress gets a reference to the given string and assigns it to the PrivateEndpointIPAddress field.
func (o *CreateEndpointRequest) SetPrivateEndpointIPAddress(v string) {
	o.PrivateEndpointIPAddress = &v
}

// GetEndpointGroupName returns the EndpointGroupName field value if set, zero value otherwise
func (o *CreateEndpointRequest) GetEndpointGroupName() string {
	if o == nil || IsNil(o.EndpointGroupName) {
		var ret string
		return ret
	}
	return *o.EndpointGroupName
}

// GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEndpointRequest) GetEndpointGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointGroupName) {
		return nil, false
	}

	return o.EndpointGroupName, true
}

// HasEndpointGroupName returns a boolean if a field has been set.
func (o *CreateEndpointRequest) HasEndpointGroupName() bool {
	if o != nil && !IsNil(o.EndpointGroupName) {
		return true
	}

	return false
}

// SetEndpointGroupName gets a reference to the given string and assigns it to the EndpointGroupName field.
func (o *CreateEndpointRequest) SetEndpointGroupName(v string) {
	o.EndpointGroupName = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise
func (o *CreateEndpointRequest) GetEndpoints() []CreateGCPForwardingRuleRequest {
	if o == nil || IsNil(o.Endpoints) {
		var ret []CreateGCPForwardingRuleRequest
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEndpointRequest) GetEndpointsOk() (*[]CreateGCPForwardingRuleRequest, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}

	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *CreateEndpointRequest) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []CreateGCPForwardingRuleRequest and assigns it to the Endpoints field.
func (o *CreateEndpointRequest) SetEndpoints(v []CreateGCPForwardingRuleRequest) {
	o.Endpoints = &v
}

// GetGcpProjectId returns the GcpProjectId field value if set, zero value otherwise
func (o *CreateEndpointRequest) GetGcpProjectId() string {
	if o == nil || IsNil(o.GcpProjectId) {
		var ret string
		return ret
	}
	return *o.GcpProjectId
}

// GetGcpProjectIdOk returns a tuple with the GcpProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEndpointRequest) GetGcpProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectId) {
		return nil, false
	}

	return o.GcpProjectId, true
}

// HasGcpProjectId returns a boolean if a field has been set.
func (o *CreateEndpointRequest) HasGcpProjectId() bool {
	if o != nil && !IsNil(o.GcpProjectId) {
		return true
	}

	return false
}

// SetGcpProjectId gets a reference to the given string and assigns it to the GcpProjectId field.
func (o *CreateEndpointRequest) SetGcpProjectId(v string) {
	o.GcpProjectId = &v
}

func (o CreateEndpointRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateEndpointRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PrivateEndpointIPAddress) {
		toSerialize["privateEndpointIPAddress"] = o.PrivateEndpointIPAddress
	}
	if !IsNil(o.EndpointGroupName) {
		toSerialize["endpointGroupName"] = o.EndpointGroupName
	}
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	if !IsNil(o.GcpProjectId) {
		toSerialize["gcpProjectId"] = o.GcpProjectId
	}
	return toSerialize, nil
}
