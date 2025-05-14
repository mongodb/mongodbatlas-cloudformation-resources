// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// CloudProviderAccessRole Cloud provider access role.
type CloudProviderAccessRole struct {
	// Human-readable label that identifies the cloud provider of the role.
	ProviderName string `json:"providerName"`
	// Amazon Resource Name that identifies the Amazon Web Services (AWS) user account that MongoDB Cloud uses when it assumes the Identity and Access Management (IAM) role.
	// Read only field.
	AtlasAWSAccountArn *string `json:"atlasAWSAccountArn,omitempty"`
	// Unique external ID that MongoDB Cloud uses when it assumes the IAM role in your Amazon Web Services (AWS) account.
	// Read only field.
	AtlasAssumedRoleExternalId *string `json:"atlasAssumedRoleExternalId,omitempty"`
	// Date and time when someone authorized this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	AuthorizedDate *time.Time `json:"authorizedDate,omitempty"`
	// Date and time when someone created this role for the specified cloud service provider. This parameter expresses its value in the ISO 8601 timestamp format in UTC.  Alternatively: Date and time when this Azure Service Principal was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// List that contains application features associated with this Amazon Web Services (AWS) Identity and Access Management (IAM) role.  Alternatively: List that contains application features associated with this Azure Service Principal.
	// Read only field.
	FeatureUsages *[]CloudProviderAccessFeatureUsage `json:"featureUsages,omitempty"`
	// Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that MongoDB Cloud assumes when it accesses resources in your AWS account.
	IamAssumedRoleArn *string `json:"iamAssumedRoleArn,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the role.
	// Read only field.
	RoleId *string `json:"roleId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Azure Service Principal in Atlas.
	// Read only field.
	Id *string `json:"_id,omitempty"`
	// Azure Active Directory Application ID of Atlas.
	AtlasAzureAppId *string `json:"atlasAzureAppId,omitempty"`
	// Date and time when this Azure Service Principal was last updated. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
	// UUID string that identifies the Azure Service Principal.
	ServicePrincipalId *string `json:"servicePrincipalId,omitempty"`
	// UUID String that identifies the Azure Active Directory Tenant ID.
	TenantId *string `json:"tenantId,omitempty"`
}

// NewCloudProviderAccessRole instantiates a new CloudProviderAccessRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessRole(providerName string) *CloudProviderAccessRole {
	this := CloudProviderAccessRole{}
	this.ProviderName = providerName
	return &this
}

// NewCloudProviderAccessRoleWithDefaults instantiates a new CloudProviderAccessRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessRoleWithDefaults() *CloudProviderAccessRole {
	this := CloudProviderAccessRole{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *CloudProviderAccessRole) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *CloudProviderAccessRole) SetProviderName(v string) {
	o.ProviderName = v
}

// GetAtlasAWSAccountArn returns the AtlasAWSAccountArn field value if set, zero value otherwise
func (o *CloudProviderAccessRole) GetAtlasAWSAccountArn() string {
	if o == nil || IsNil(o.AtlasAWSAccountArn) {
		var ret string
		return ret
	}
	return *o.AtlasAWSAccountArn
}

// GetAtlasAWSAccountArnOk returns a tuple with the AtlasAWSAccountArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetAtlasAWSAccountArnOk() (*string, bool) {
	if o == nil || IsNil(o.AtlasAWSAccountArn) {
		return nil, false
	}

	return o.AtlasAWSAccountArn, true
}

// HasAtlasAWSAccountArn returns a boolean if a field has been set.
func (o *CloudProviderAccessRole) HasAtlasAWSAccountArn() bool {
	if o != nil && !IsNil(o.AtlasAWSAccountArn) {
		return true
	}

	return false
}

// SetAtlasAWSAccountArn gets a reference to the given string and assigns it to the AtlasAWSAccountArn field.
func (o *CloudProviderAccessRole) SetAtlasAWSAccountArn(v string) {
	o.AtlasAWSAccountArn = &v
}

// GetAtlasAssumedRoleExternalId returns the AtlasAssumedRoleExternalId field value if set, zero value otherwise
func (o *CloudProviderAccessRole) GetAtlasAssumedRoleExternalId() string {
	if o == nil || IsNil(o.AtlasAssumedRoleExternalId) {
		var ret string
		return ret
	}
	return *o.AtlasAssumedRoleExternalId
}

// GetAtlasAssumedRoleExternalIdOk returns a tuple with the AtlasAssumedRoleExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetAtlasAssumedRoleExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.AtlasAssumedRoleExternalId) {
		return nil, false
	}

	return o.AtlasAssumedRoleExternalId, true
}

// HasAtlasAssumedRoleExternalId returns a boolean if a field has been set.
func (o *CloudProviderAccessRole) HasAtlasAssumedRoleExternalId() bool {
	if o != nil && !IsNil(o.AtlasAssumedRoleExternalId) {
		return true
	}

	return false
}

// SetAtlasAssumedRoleExternalId gets a reference to the given string and assigns it to the AtlasAssumedRoleExternalId field.
func (o *CloudProviderAccessRole) SetAtlasAssumedRoleExternalId(v string) {
	o.AtlasAssumedRoleExternalId = &v
}

// GetAuthorizedDate returns the AuthorizedDate field value if set, zero value otherwise
func (o *CloudProviderAccessRole) GetAuthorizedDate() time.Time {
	if o == nil || IsNil(o.AuthorizedDate) {
		var ret time.Time
		return ret
	}
	return *o.AuthorizedDate
}

// GetAuthorizedDateOk returns a tuple with the AuthorizedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetAuthorizedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AuthorizedDate) {
		return nil, false
	}

	return o.AuthorizedDate, true
}

// HasAuthorizedDate returns a boolean if a field has been set.
func (o *CloudProviderAccessRole) HasAuthorizedDate() bool {
	if o != nil && !IsNil(o.AuthorizedDate) {
		return true
	}

	return false
}

// SetAuthorizedDate gets a reference to the given time.Time and assigns it to the AuthorizedDate field.
func (o *CloudProviderAccessRole) SetAuthorizedDate(v time.Time) {
	o.AuthorizedDate = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise
func (o *CloudProviderAccessRole) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}

	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *CloudProviderAccessRole) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *CloudProviderAccessRole) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetFeatureUsages returns the FeatureUsages field value if set, zero value otherwise
func (o *CloudProviderAccessRole) GetFeatureUsages() []CloudProviderAccessFeatureUsage {
	if o == nil || IsNil(o.FeatureUsages) {
		var ret []CloudProviderAccessFeatureUsage
		return ret
	}
	return *o.FeatureUsages
}

// GetFeatureUsagesOk returns a tuple with the FeatureUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetFeatureUsagesOk() (*[]CloudProviderAccessFeatureUsage, bool) {
	if o == nil || IsNil(o.FeatureUsages) {
		return nil, false
	}

	return o.FeatureUsages, true
}

// HasFeatureUsages returns a boolean if a field has been set.
func (o *CloudProviderAccessRole) HasFeatureUsages() bool {
	if o != nil && !IsNil(o.FeatureUsages) {
		return true
	}

	return false
}

// SetFeatureUsages gets a reference to the given []CloudProviderAccessFeatureUsage and assigns it to the FeatureUsages field.
func (o *CloudProviderAccessRole) SetFeatureUsages(v []CloudProviderAccessFeatureUsage) {
	o.FeatureUsages = &v
}

// GetIamAssumedRoleArn returns the IamAssumedRoleArn field value if set, zero value otherwise
func (o *CloudProviderAccessRole) GetIamAssumedRoleArn() string {
	if o == nil || IsNil(o.IamAssumedRoleArn) {
		var ret string
		return ret
	}
	return *o.IamAssumedRoleArn
}

// GetIamAssumedRoleArnOk returns a tuple with the IamAssumedRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetIamAssumedRoleArnOk() (*string, bool) {
	if o == nil || IsNil(o.IamAssumedRoleArn) {
		return nil, false
	}

	return o.IamAssumedRoleArn, true
}

// HasIamAssumedRoleArn returns a boolean if a field has been set.
func (o *CloudProviderAccessRole) HasIamAssumedRoleArn() bool {
	if o != nil && !IsNil(o.IamAssumedRoleArn) {
		return true
	}

	return false
}

// SetIamAssumedRoleArn gets a reference to the given string and assigns it to the IamAssumedRoleArn field.
func (o *CloudProviderAccessRole) SetIamAssumedRoleArn(v string) {
	o.IamAssumedRoleArn = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise
func (o *CloudProviderAccessRole) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}

	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *CloudProviderAccessRole) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *CloudProviderAccessRole) SetRoleId(v string) {
	o.RoleId = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *CloudProviderAccessRole) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudProviderAccessRole) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudProviderAccessRole) SetId(v string) {
	o.Id = &v
}

// GetAtlasAzureAppId returns the AtlasAzureAppId field value if set, zero value otherwise
func (o *CloudProviderAccessRole) GetAtlasAzureAppId() string {
	if o == nil || IsNil(o.AtlasAzureAppId) {
		var ret string
		return ret
	}
	return *o.AtlasAzureAppId
}

// GetAtlasAzureAppIdOk returns a tuple with the AtlasAzureAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetAtlasAzureAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AtlasAzureAppId) {
		return nil, false
	}

	return o.AtlasAzureAppId, true
}

// HasAtlasAzureAppId returns a boolean if a field has been set.
func (o *CloudProviderAccessRole) HasAtlasAzureAppId() bool {
	if o != nil && !IsNil(o.AtlasAzureAppId) {
		return true
	}

	return false
}

// SetAtlasAzureAppId gets a reference to the given string and assigns it to the AtlasAzureAppId field.
func (o *CloudProviderAccessRole) SetAtlasAzureAppId(v string) {
	o.AtlasAzureAppId = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise
func (o *CloudProviderAccessRole) GetLastUpdatedDate() time.Time {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}

	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *CloudProviderAccessRole) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given time.Time and assigns it to the LastUpdatedDate field.
func (o *CloudProviderAccessRole) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = &v
}

// GetServicePrincipalId returns the ServicePrincipalId field value if set, zero value otherwise
func (o *CloudProviderAccessRole) GetServicePrincipalId() string {
	if o == nil || IsNil(o.ServicePrincipalId) {
		var ret string
		return ret
	}
	return *o.ServicePrincipalId
}

// GetServicePrincipalIdOk returns a tuple with the ServicePrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetServicePrincipalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePrincipalId) {
		return nil, false
	}

	return o.ServicePrincipalId, true
}

// HasServicePrincipalId returns a boolean if a field has been set.
func (o *CloudProviderAccessRole) HasServicePrincipalId() bool {
	if o != nil && !IsNil(o.ServicePrincipalId) {
		return true
	}

	return false
}

// SetServicePrincipalId gets a reference to the given string and assigns it to the ServicePrincipalId field.
func (o *CloudProviderAccessRole) SetServicePrincipalId(v string) {
	o.ServicePrincipalId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise
func (o *CloudProviderAccessRole) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessRole) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}

	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CloudProviderAccessRole) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CloudProviderAccessRole) SetTenantId(v string) {
	o.TenantId = &v
}

func (o CloudProviderAccessRole) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudProviderAccessRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["providerName"] = o.ProviderName
	if !IsNil(o.IamAssumedRoleArn) {
		toSerialize["iamAssumedRoleArn"] = o.IamAssumedRoleArn
	}
	if !IsNil(o.AtlasAzureAppId) {
		toSerialize["atlasAzureAppId"] = o.AtlasAzureAppId
	}
	if !IsNil(o.ServicePrincipalId) {
		toSerialize["servicePrincipalId"] = o.ServicePrincipalId
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	return toSerialize, nil
}
