// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// CreateAtlasOrganizationApiKey Details of the programmatic API key to be created.
type CreateAtlasOrganizationApiKey struct {
	// Purpose or explanation provided when someone created this organization API key.
	Desc string `json:"desc"`
	// List of roles to grant this API key. If you provide this list, provide a minimum of one role and ensure each role applies to this organization.
	Roles []string `json:"roles"`
}

// NewCreateAtlasOrganizationApiKey instantiates a new CreateAtlasOrganizationApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAtlasOrganizationApiKey(desc string, roles []string) *CreateAtlasOrganizationApiKey {
	this := CreateAtlasOrganizationApiKey{}
	this.Desc = desc
	this.Roles = roles
	return &this
}

// NewCreateAtlasOrganizationApiKeyWithDefaults instantiates a new CreateAtlasOrganizationApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAtlasOrganizationApiKeyWithDefaults() *CreateAtlasOrganizationApiKey {
	this := CreateAtlasOrganizationApiKey{}
	return &this
}

// GetDesc returns the Desc field value
func (o *CreateAtlasOrganizationApiKey) GetDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value
// and a boolean to check if the value has been set.
func (o *CreateAtlasOrganizationApiKey) GetDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Desc, true
}

// SetDesc sets field value
func (o *CreateAtlasOrganizationApiKey) SetDesc(v string) {
	o.Desc = v
}

// GetRoles returns the Roles field value
func (o *CreateAtlasOrganizationApiKey) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *CreateAtlasOrganizationApiKey) GetRolesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *CreateAtlasOrganizationApiKey) SetRoles(v []string) {
	o.Roles = v
}

func (o CreateAtlasOrganizationApiKey) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CreateAtlasOrganizationApiKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["desc"] = o.Desc
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}
