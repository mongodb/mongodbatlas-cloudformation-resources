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

// EndpointService struct for EndpointService
type EndpointService struct {
	// Cloud service provider that serves the requested endpoint service.
	// Read only field.
	CloudProvider string `json:"cloudProvider"`
	// Error message returned when requesting private connection resource. The resource returns `null` if the request succeeded.
	// Read only field.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Private Endpoint Service.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// Cloud provider region that manages this Private Endpoint Service.
	// Read only field.
	RegionName *string `json:"regionName,omitempty"`
	// State of the Private Endpoint Service connection when MongoDB Cloud received this request.
	// Read only field.
	Status *string `json:"status,omitempty"`
	// Unique string that identifies the Amazon Web Services (AWS) PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service.
	// Read only field.
	EndpointServiceName *string `json:"endpointServiceName,omitempty"`
	// List of strings that identify private endpoint interfaces applied to the specified project.
	// Read only field.
	InterfaceEndpoints *[]string `json:"interfaceEndpoints,omitempty"`
	// List of private endpoints assigned to this Azure Private Link Service.
	// Read only field.
	PrivateEndpoints *[]string `json:"privateEndpoints,omitempty"`
	// Unique string that identifies the Azure Private Link Service that MongoDB Cloud manages.
	// Read only field.
	PrivateLinkServiceName *string `json:"privateLinkServiceName,omitempty"`
	// Root-relative path that identifies of the Azure Private Link Service that MongoDB Cloud manages. Use this value to create a private endpoint connection to an Azure VNet.
	// Read only field.
	PrivateLinkServiceResourceId *string `json:"privateLinkServiceResourceId,omitempty"`
	// List of Google Cloud network endpoint groups that corresponds to the Private Service Connect endpoint service.
	EndpointGroupNames *[]string `json:"endpointGroupNames,omitempty"`
	// List of Uniform Resource Locators (URLs) that identifies endpoints that MongoDB Cloud can use to access one Google Cloud Service across a Google Cloud Virtual Private Connection (VPC) network.
	ServiceAttachmentNames *[]string `json:"serviceAttachmentNames,omitempty"`
}

// NewEndpointService instantiates a new EndpointService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointService(cloudProvider string) *EndpointService {
	this := EndpointService{}
	this.CloudProvider = cloudProvider
	return &this
}

// NewEndpointServiceWithDefaults instantiates a new EndpointService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointServiceWithDefaults() *EndpointService {
	this := EndpointService{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *EndpointService) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *EndpointService) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *EndpointService) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise
func (o *EndpointService) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointService) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}

	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *EndpointService) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *EndpointService) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *EndpointService) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointService) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EndpointService) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EndpointService) SetId(v string) {
	o.Id = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *EndpointService) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointService) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *EndpointService) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *EndpointService) SetRegionName(v string) {
	o.RegionName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *EndpointService) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointService) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EndpointService) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EndpointService) SetStatus(v string) {
	o.Status = &v
}

// GetEndpointServiceName returns the EndpointServiceName field value if set, zero value otherwise
func (o *EndpointService) GetEndpointServiceName() string {
	if o == nil || IsNil(o.EndpointServiceName) {
		var ret string
		return ret
	}
	return *o.EndpointServiceName
}

// GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointService) GetEndpointServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointServiceName) {
		return nil, false
	}

	return o.EndpointServiceName, true
}

// HasEndpointServiceName returns a boolean if a field has been set.
func (o *EndpointService) HasEndpointServiceName() bool {
	if o != nil && !IsNil(o.EndpointServiceName) {
		return true
	}

	return false
}

// SetEndpointServiceName gets a reference to the given string and assigns it to the EndpointServiceName field.
func (o *EndpointService) SetEndpointServiceName(v string) {
	o.EndpointServiceName = &v
}

// GetInterfaceEndpoints returns the InterfaceEndpoints field value if set, zero value otherwise
func (o *EndpointService) GetInterfaceEndpoints() []string {
	if o == nil || IsNil(o.InterfaceEndpoints) {
		var ret []string
		return ret
	}
	return *o.InterfaceEndpoints
}

// GetInterfaceEndpointsOk returns a tuple with the InterfaceEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointService) GetInterfaceEndpointsOk() (*[]string, bool) {
	if o == nil || IsNil(o.InterfaceEndpoints) {
		return nil, false
	}

	return o.InterfaceEndpoints, true
}

// HasInterfaceEndpoints returns a boolean if a field has been set.
func (o *EndpointService) HasInterfaceEndpoints() bool {
	if o != nil && !IsNil(o.InterfaceEndpoints) {
		return true
	}

	return false
}

// SetInterfaceEndpoints gets a reference to the given []string and assigns it to the InterfaceEndpoints field.
func (o *EndpointService) SetInterfaceEndpoints(v []string) {
	o.InterfaceEndpoints = &v
}

// GetPrivateEndpoints returns the PrivateEndpoints field value if set, zero value otherwise
func (o *EndpointService) GetPrivateEndpoints() []string {
	if o == nil || IsNil(o.PrivateEndpoints) {
		var ret []string
		return ret
	}
	return *o.PrivateEndpoints
}

// GetPrivateEndpointsOk returns a tuple with the PrivateEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointService) GetPrivateEndpointsOk() (*[]string, bool) {
	if o == nil || IsNil(o.PrivateEndpoints) {
		return nil, false
	}

	return o.PrivateEndpoints, true
}

// HasPrivateEndpoints returns a boolean if a field has been set.
func (o *EndpointService) HasPrivateEndpoints() bool {
	if o != nil && !IsNil(o.PrivateEndpoints) {
		return true
	}

	return false
}

// SetPrivateEndpoints gets a reference to the given []string and assigns it to the PrivateEndpoints field.
func (o *EndpointService) SetPrivateEndpoints(v []string) {
	o.PrivateEndpoints = &v
}

// GetPrivateLinkServiceName returns the PrivateLinkServiceName field value if set, zero value otherwise
func (o *EndpointService) GetPrivateLinkServiceName() string {
	if o == nil || IsNil(o.PrivateLinkServiceName) {
		var ret string
		return ret
	}
	return *o.PrivateLinkServiceName
}

// GetPrivateLinkServiceNameOk returns a tuple with the PrivateLinkServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointService) GetPrivateLinkServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateLinkServiceName) {
		return nil, false
	}

	return o.PrivateLinkServiceName, true
}

// HasPrivateLinkServiceName returns a boolean if a field has been set.
func (o *EndpointService) HasPrivateLinkServiceName() bool {
	if o != nil && !IsNil(o.PrivateLinkServiceName) {
		return true
	}

	return false
}

// SetPrivateLinkServiceName gets a reference to the given string and assigns it to the PrivateLinkServiceName field.
func (o *EndpointService) SetPrivateLinkServiceName(v string) {
	o.PrivateLinkServiceName = &v
}

// GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field value if set, zero value otherwise
func (o *EndpointService) GetPrivateLinkServiceResourceId() string {
	if o == nil || IsNil(o.PrivateLinkServiceResourceId) {
		var ret string
		return ret
	}
	return *o.PrivateLinkServiceResourceId
}

// GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointService) GetPrivateLinkServiceResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateLinkServiceResourceId) {
		return nil, false
	}

	return o.PrivateLinkServiceResourceId, true
}

// HasPrivateLinkServiceResourceId returns a boolean if a field has been set.
func (o *EndpointService) HasPrivateLinkServiceResourceId() bool {
	if o != nil && !IsNil(o.PrivateLinkServiceResourceId) {
		return true
	}

	return false
}

// SetPrivateLinkServiceResourceId gets a reference to the given string and assigns it to the PrivateLinkServiceResourceId field.
func (o *EndpointService) SetPrivateLinkServiceResourceId(v string) {
	o.PrivateLinkServiceResourceId = &v
}

// GetEndpointGroupNames returns the EndpointGroupNames field value if set, zero value otherwise
func (o *EndpointService) GetEndpointGroupNames() []string {
	if o == nil || IsNil(o.EndpointGroupNames) {
		var ret []string
		return ret
	}
	return *o.EndpointGroupNames
}

// GetEndpointGroupNamesOk returns a tuple with the EndpointGroupNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointService) GetEndpointGroupNamesOk() (*[]string, bool) {
	if o == nil || IsNil(o.EndpointGroupNames) {
		return nil, false
	}

	return o.EndpointGroupNames, true
}

// HasEndpointGroupNames returns a boolean if a field has been set.
func (o *EndpointService) HasEndpointGroupNames() bool {
	if o != nil && !IsNil(o.EndpointGroupNames) {
		return true
	}

	return false
}

// SetEndpointGroupNames gets a reference to the given []string and assigns it to the EndpointGroupNames field.
func (o *EndpointService) SetEndpointGroupNames(v []string) {
	o.EndpointGroupNames = &v
}

// GetServiceAttachmentNames returns the ServiceAttachmentNames field value if set, zero value otherwise
func (o *EndpointService) GetServiceAttachmentNames() []string {
	if o == nil || IsNil(o.ServiceAttachmentNames) {
		var ret []string
		return ret
	}
	return *o.ServiceAttachmentNames
}

// GetServiceAttachmentNamesOk returns a tuple with the ServiceAttachmentNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointService) GetServiceAttachmentNamesOk() (*[]string, bool) {
	if o == nil || IsNil(o.ServiceAttachmentNames) {
		return nil, false
	}

	return o.ServiceAttachmentNames, true
}

// HasServiceAttachmentNames returns a boolean if a field has been set.
func (o *EndpointService) HasServiceAttachmentNames() bool {
	if o != nil && !IsNil(o.ServiceAttachmentNames) {
		return true
	}

	return false
}

// SetServiceAttachmentNames gets a reference to the given []string and assigns it to the ServiceAttachmentNames field.
func (o *EndpointService) SetServiceAttachmentNames(v []string) {
	o.ServiceAttachmentNames = &v
}

func (o EndpointService) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o EndpointService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointGroupNames) {
		toSerialize["endpointGroupNames"] = o.EndpointGroupNames
	}
	if !IsNil(o.ServiceAttachmentNames) {
		toSerialize["serviceAttachmentNames"] = o.ServiceAttachmentNames
	}
	return toSerialize, nil
}
