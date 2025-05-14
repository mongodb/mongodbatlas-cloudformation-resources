// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ApiKey Details contained in one API key.
type ApiKey struct {
	// List of network addresses granted access to this API using this API key.
	// Read only field.
	AccessList *[]AccessListItem `json:"accessList,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this organization API key.
	// Read only field.
	Id string `json:"id"`
	// Public API key value set for the specified organization API key.
	// Read only field.
	PublicKey string `json:"publicKey"`
	// List that contains roles that the API key needs to have. All roles you provide must be valid for the specified project or organization. Each request must include a minimum of one valid role. The resource returns all project and organization roles assigned to the Cloud user.
	// Read only field.
	Roles *[]CloudAccessRoleAssignment `json:"roles,omitempty"`
}

// NewApiKey instantiates a new ApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKey(id string, publicKey string) *ApiKey {
	this := ApiKey{}
	this.Id = id
	this.PublicKey = publicKey
	return &this
}

// NewApiKeyWithDefaults instantiates a new ApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyWithDefaults() *ApiKey {
	this := ApiKey{}
	return &this
}

// GetAccessList returns the AccessList field value if set, zero value otherwise
func (o *ApiKey) GetAccessList() []AccessListItem {
	if o == nil || IsNil(o.AccessList) {
		var ret []AccessListItem
		return ret
	}
	return *o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetAccessListOk() (*[]AccessListItem, bool) {
	if o == nil || IsNil(o.AccessList) {
		return nil, false
	}

	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *ApiKey) HasAccessList() bool {
	if o != nil && !IsNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given []AccessListItem and assigns it to the AccessList field.
func (o *ApiKey) SetAccessList(v []AccessListItem) {
	o.AccessList = &v
}

// GetId returns the Id field value
func (o *ApiKey) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiKey) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiKey) SetId(v string) {
	o.Id = v
}

// GetPublicKey returns the PublicKey field value
func (o *ApiKey) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *ApiKey) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *ApiKey) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *ApiKey) GetRoles() []CloudAccessRoleAssignment {
	if o == nil || IsNil(o.Roles) {
		var ret []CloudAccessRoleAssignment
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetRolesOk() (*[]CloudAccessRoleAssignment, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApiKey) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []CloudAccessRoleAssignment and assigns it to the Roles field.
func (o *ApiKey) SetRoles(v []CloudAccessRoleAssignment) {
	o.Roles = &v
}

func (o ApiKey) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ApiKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
