// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// AuthFederationRoleMapping Mapping settings that link one IdP and MongoDB Cloud.
type AuthFederationRoleMapping struct {
	// Unique human-readable label that identifies the identity provider group to which this role mapping applies.
	ExternalGroupName string `json:"externalGroupName"`
	// Unique 24-hexadecimal digit string that identifies this role mapping.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// Atlas roles and the unique identifiers of the groups and organizations associated with each role.
	RoleAssignments *[]RoleAssignment `json:"roleAssignments,omitempty"`
}

// NewAuthFederationRoleMapping instantiates a new AuthFederationRoleMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthFederationRoleMapping(externalGroupName string) *AuthFederationRoleMapping {
	this := AuthFederationRoleMapping{}
	this.ExternalGroupName = externalGroupName
	return &this
}

// NewAuthFederationRoleMappingWithDefaults instantiates a new AuthFederationRoleMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthFederationRoleMappingWithDefaults() *AuthFederationRoleMapping {
	this := AuthFederationRoleMapping{}
	return &this
}

// GetExternalGroupName returns the ExternalGroupName field value
func (o *AuthFederationRoleMapping) GetExternalGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalGroupName
}

// GetExternalGroupNameOk returns a tuple with the ExternalGroupName field value
// and a boolean to check if the value has been set.
func (o *AuthFederationRoleMapping) GetExternalGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalGroupName, true
}

// SetExternalGroupName sets field value
func (o *AuthFederationRoleMapping) SetExternalGroupName(v string) {
	o.ExternalGroupName = v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *AuthFederationRoleMapping) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFederationRoleMapping) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthFederationRoleMapping) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthFederationRoleMapping) SetId(v string) {
	o.Id = &v
}

// GetRoleAssignments returns the RoleAssignments field value if set, zero value otherwise
func (o *AuthFederationRoleMapping) GetRoleAssignments() []RoleAssignment {
	if o == nil || IsNil(o.RoleAssignments) {
		var ret []RoleAssignment
		return ret
	}
	return *o.RoleAssignments
}

// GetRoleAssignmentsOk returns a tuple with the RoleAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthFederationRoleMapping) GetRoleAssignmentsOk() (*[]RoleAssignment, bool) {
	if o == nil || IsNil(o.RoleAssignments) {
		return nil, false
	}

	return o.RoleAssignments, true
}

// HasRoleAssignments returns a boolean if a field has been set.
func (o *AuthFederationRoleMapping) HasRoleAssignments() bool {
	if o != nil && !IsNil(o.RoleAssignments) {
		return true
	}

	return false
}

// SetRoleAssignments gets a reference to the given []RoleAssignment and assigns it to the RoleAssignments field.
func (o *AuthFederationRoleMapping) SetRoleAssignments(v []RoleAssignment) {
	o.RoleAssignments = &v
}

func (o AuthFederationRoleMapping) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AuthFederationRoleMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalGroupName"] = o.ExternalGroupName
	if !IsNil(o.RoleAssignments) {
		toSerialize["roleAssignments"] = o.RoleAssignments
	}
	return toSerialize, nil
}
