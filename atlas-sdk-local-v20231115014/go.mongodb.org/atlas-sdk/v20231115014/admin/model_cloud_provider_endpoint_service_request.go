// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// CloudProviderEndpointServiceRequest struct for CloudProviderEndpointServiceRequest
type CloudProviderEndpointServiceRequest struct {
	// Human-readable label that identifies the cloud service provider for which you want to create the private endpoint service.
	// Write only field.
	ProviderName string `json:"providerName"`
	// Cloud provider region in which you want to create the private endpoint service. Regions accepted as values differ for [Amazon Web Services](https://docs.atlas.mongodb.com/reference/amazon-aws/), [Google Cloud Platform](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Microsoft Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
	// Write only field.
	Region string `json:"region"`
}

// NewCloudProviderEndpointServiceRequest instantiates a new CloudProviderEndpointServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderEndpointServiceRequest(providerName string, region string) *CloudProviderEndpointServiceRequest {
	this := CloudProviderEndpointServiceRequest{}
	this.ProviderName = providerName
	this.Region = region
	return &this
}

// NewCloudProviderEndpointServiceRequestWithDefaults instantiates a new CloudProviderEndpointServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderEndpointServiceRequestWithDefaults() *CloudProviderEndpointServiceRequest {
	this := CloudProviderEndpointServiceRequest{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *CloudProviderEndpointServiceRequest) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *CloudProviderEndpointServiceRequest) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *CloudProviderEndpointServiceRequest) SetProviderName(v string) {
	o.ProviderName = v
}

// GetRegion returns the Region field value
func (o *CloudProviderEndpointServiceRequest) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CloudProviderEndpointServiceRequest) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CloudProviderEndpointServiceRequest) SetRegion(v string) {
	o.Region = v
}

func (o CloudProviderEndpointServiceRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudProviderEndpointServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["providerName"] = o.ProviderName
	toSerialize["region"] = o.Region
	return toSerialize, nil
}
