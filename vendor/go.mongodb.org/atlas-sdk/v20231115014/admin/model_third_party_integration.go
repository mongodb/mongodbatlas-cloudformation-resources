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

// ThirdPartyIntegration Collection of settings that describe third-party integrations.
type ThirdPartyIntegration struct {
	// Integration id.
	Id *string `json:"id,omitempty"`
	// Integration type  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.  Alternatively: Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
	// Key that allows MongoDB Cloud to access your Datadog account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.  Alternatively: Key that allows MongoDB Cloud to access your Opsgenie account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.  Alternatively: Key that allows MongoDB Cloud to access your VictorOps account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	ApiKey *string `json:"apiKey,omitempty"`
	// Two-letter code that indicates which regional URL MongoDB uses to access the Datadog API.  To learn more about Datadog's regions, see <a href=\"https://docs.datadoghq.com/getting_started/site/\" target=\"_blank\" rel=\"noopener noreferrer\">Datadog Sites</a>.  Alternatively: Two-letter code that indicates which regional URL MongoDB uses to access the Opsgenie API.  Alternatively: PagerDuty region that indicates the API Uniform Resource Locator (URL) to use.
	Region *string `json:"region,omitempty"`
	// Endpoint web address of the Microsoft Teams webhook to which MongoDB Cloud sends notifications.  **NOTE**: When you view or edit the alert for a Microsoft Teams notification, the URL appears partially redacted.
	MicrosoftTeamsWebhookUrl *string `json:"microsoftTeamsWebhookUrl,omitempty"`
	// Unique 40-hexadecimal digit string that identifies your New Relic account.
	AccountId *string `json:"accountId,omitempty"`
	// Unique 40-hexadecimal digit string that identifies your New Relic license.  **IMPORTANT**: Effective Wednesday, June 16th, 2021, New Relic no longer supports the plugin-based integration with MongoDB. We do not recommend that you sign up for the plugin-based integration. To learn more, see the <a href=\"https://discuss.newrelic.com/t/new-relic-plugin-eol-wednesday-june-16th-2021/127267\" target=\"_blank\">New Relic Plugin EOL Statement</a> Consider configuring an alternative monitoring integration before June 16th to maintain visibility into your MongoDB deployments.
	LicenseKey *string `json:"licenseKey,omitempty"`
	// Query key used to access your New Relic account.
	ReadToken *string `json:"readToken,omitempty"`
	// Insert key associated with your New Relic account.
	WriteToken *string `json:"writeToken,omitempty"`
	// Service key associated with your PagerDuty account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.
	ServiceKey *string `json:"serviceKey,omitempty"`
	// Flag that indicates whether someone has activated the Prometheus integration.
	Enabled *bool `json:"enabled,omitempty"`
	// Password needed to allow MongoDB Cloud to access your Prometheus account.
	// Write only field.
	Password *string `json:"password,omitempty"`
	// Desired method to discover the Prometheus service.
	ServiceDiscovery *string `json:"serviceDiscovery,omitempty"`
	// Human-readable label that identifies your Prometheus incoming webhook.
	Username *string `json:"username,omitempty"`
	// Key that allows MongoDB Cloud to access your Slack account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.  **IMPORTANT**: Slack integrations now use the OAuth2 verification method and must  be initially configured, or updated from a legacy integration, through the Atlas  third-party service integrations page. Legacy tokens will soon no longer be  supported.
	ApiToken *string `json:"apiToken,omitempty"`
	// Name of the Slack channel to which MongoDB Cloud sends alert notifications.
	ChannelName *string `json:"channelName,omitempty"`
	// Human-readable label that identifies your Slack team. Set this parameter when you configure a legacy Slack integration.
	TeamName *string `json:"teamName,omitempty"`
	// Routing key associated with your Splunk On-Call account.
	RoutingKey *string `json:"routingKey,omitempty"`
	// An optional field returned if your webhook is configured with a secret.  **NOTE**: When you view or edit the alert for a webhook notification, the secret appears completely redacted.
	Secret *string `json:"secret,omitempty"`
	// Endpoint web address to which MongoDB Cloud sends notifications.  **NOTE**: When you view or edit the alert for a webhook notification, the URL appears partially redacted.
	Url *string `json:"url,omitempty"`
}

// NewThirdPartyIntegration instantiates a new ThirdPartyIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyIntegration() *ThirdPartyIntegration {
	this := ThirdPartyIntegration{}
	return &this
}

// NewThirdPartyIntegrationWithDefaults instantiates a new ThirdPartyIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyIntegrationWithDefaults() *ThirdPartyIntegration {
	this := ThirdPartyIntegration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ThirdPartyIntegration) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ThirdPartyIntegration) SetType(v string) {
	o.Type = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}

	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *ThirdPartyIntegration) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetRegion returns the Region field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}

	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ThirdPartyIntegration) SetRegion(v string) {
	o.Region = &v
}

// GetMicrosoftTeamsWebhookUrl returns the MicrosoftTeamsWebhookUrl field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetMicrosoftTeamsWebhookUrl() string {
	if o == nil || IsNil(o.MicrosoftTeamsWebhookUrl) {
		var ret string
		return ret
	}
	return *o.MicrosoftTeamsWebhookUrl
}

// GetMicrosoftTeamsWebhookUrlOk returns a tuple with the MicrosoftTeamsWebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetMicrosoftTeamsWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MicrosoftTeamsWebhookUrl) {
		return nil, false
	}

	return o.MicrosoftTeamsWebhookUrl, true
}

// HasMicrosoftTeamsWebhookUrl returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasMicrosoftTeamsWebhookUrl() bool {
	if o != nil && !IsNil(o.MicrosoftTeamsWebhookUrl) {
		return true
	}

	return false
}

// SetMicrosoftTeamsWebhookUrl gets a reference to the given string and assigns it to the MicrosoftTeamsWebhookUrl field.
func (o *ThirdPartyIntegration) SetMicrosoftTeamsWebhookUrl(v string) {
	o.MicrosoftTeamsWebhookUrl = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}

	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ThirdPartyIntegration) SetAccountId(v string) {
	o.AccountId = &v
}

// GetLicenseKey returns the LicenseKey field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetLicenseKey() string {
	if o == nil || IsNil(o.LicenseKey) {
		var ret string
		return ret
	}
	return *o.LicenseKey
}

// GetLicenseKeyOk returns a tuple with the LicenseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetLicenseKeyOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseKey) {
		return nil, false
	}

	return o.LicenseKey, true
}

// HasLicenseKey returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasLicenseKey() bool {
	if o != nil && !IsNil(o.LicenseKey) {
		return true
	}

	return false
}

// SetLicenseKey gets a reference to the given string and assigns it to the LicenseKey field.
func (o *ThirdPartyIntegration) SetLicenseKey(v string) {
	o.LicenseKey = &v
}

// GetReadToken returns the ReadToken field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetReadToken() string {
	if o == nil || IsNil(o.ReadToken) {
		var ret string
		return ret
	}
	return *o.ReadToken
}

// GetReadTokenOk returns a tuple with the ReadToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetReadTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ReadToken) {
		return nil, false
	}

	return o.ReadToken, true
}

// HasReadToken returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasReadToken() bool {
	if o != nil && !IsNil(o.ReadToken) {
		return true
	}

	return false
}

// SetReadToken gets a reference to the given string and assigns it to the ReadToken field.
func (o *ThirdPartyIntegration) SetReadToken(v string) {
	o.ReadToken = &v
}

// GetWriteToken returns the WriteToken field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetWriteToken() string {
	if o == nil || IsNil(o.WriteToken) {
		var ret string
		return ret
	}
	return *o.WriteToken
}

// GetWriteTokenOk returns a tuple with the WriteToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetWriteTokenOk() (*string, bool) {
	if o == nil || IsNil(o.WriteToken) {
		return nil, false
	}

	return o.WriteToken, true
}

// HasWriteToken returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasWriteToken() bool {
	if o != nil && !IsNil(o.WriteToken) {
		return true
	}

	return false
}

// SetWriteToken gets a reference to the given string and assigns it to the WriteToken field.
func (o *ThirdPartyIntegration) SetWriteToken(v string) {
	o.WriteToken = &v
}

// GetServiceKey returns the ServiceKey field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetServiceKey() string {
	if o == nil || IsNil(o.ServiceKey) {
		var ret string
		return ret
	}
	return *o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetServiceKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceKey) {
		return nil, false
	}

	return o.ServiceKey, true
}

// HasServiceKey returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasServiceKey() bool {
	if o != nil && !IsNil(o.ServiceKey) {
		return true
	}

	return false
}

// SetServiceKey gets a reference to the given string and assigns it to the ServiceKey field.
func (o *ThirdPartyIntegration) SetServiceKey(v string) {
	o.ServiceKey = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ThirdPartyIntegration) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPassword returns the Password field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}

	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ThirdPartyIntegration) SetPassword(v string) {
	o.Password = &v
}

// GetServiceDiscovery returns the ServiceDiscovery field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetServiceDiscovery() string {
	if o == nil || IsNil(o.ServiceDiscovery) {
		var ret string
		return ret
	}
	return *o.ServiceDiscovery
}

// GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetServiceDiscoveryOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceDiscovery) {
		return nil, false
	}

	return o.ServiceDiscovery, true
}

// HasServiceDiscovery returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasServiceDiscovery() bool {
	if o != nil && !IsNil(o.ServiceDiscovery) {
		return true
	}

	return false
}

// SetServiceDiscovery gets a reference to the given string and assigns it to the ServiceDiscovery field.
func (o *ThirdPartyIntegration) SetServiceDiscovery(v string) {
	o.ServiceDiscovery = &v
}

// GetUsername returns the Username field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}

	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ThirdPartyIntegration) SetUsername(v string) {
	o.Username = &v
}

// GetApiToken returns the ApiToken field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetApiToken() string {
	if o == nil || IsNil(o.ApiToken) {
		var ret string
		return ret
	}
	return *o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetApiTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ApiToken) {
		return nil, false
	}

	return o.ApiToken, true
}

// HasApiToken returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasApiToken() bool {
	if o != nil && !IsNil(o.ApiToken) {
		return true
	}

	return false
}

// SetApiToken gets a reference to the given string and assigns it to the ApiToken field.
func (o *ThirdPartyIntegration) SetApiToken(v string) {
	o.ApiToken = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetChannelName() string {
	if o == nil || IsNil(o.ChannelName) {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetChannelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelName) {
		return nil, false
	}

	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasChannelName() bool {
	if o != nil && !IsNil(o.ChannelName) {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *ThirdPartyIntegration) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetTeamName() string {
	if o == nil || IsNil(o.TeamName) {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetTeamNameOk() (*string, bool) {
	if o == nil || IsNil(o.TeamName) {
		return nil, false
	}

	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasTeamName() bool {
	if o != nil && !IsNil(o.TeamName) {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *ThirdPartyIntegration) SetTeamName(v string) {
	o.TeamName = &v
}

// GetRoutingKey returns the RoutingKey field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetRoutingKey() string {
	if o == nil || IsNil(o.RoutingKey) {
		var ret string
		return ret
	}
	return *o.RoutingKey
}

// GetRoutingKeyOk returns a tuple with the RoutingKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetRoutingKeyOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingKey) {
		return nil, false
	}

	return o.RoutingKey, true
}

// HasRoutingKey returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasRoutingKey() bool {
	if o != nil && !IsNil(o.RoutingKey) {
		return true
	}

	return false
}

// SetRoutingKey gets a reference to the given string and assigns it to the RoutingKey field.
func (o *ThirdPartyIntegration) SetRoutingKey(v string) {
	o.RoutingKey = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}

	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ThirdPartyIntegration) SetSecret(v string) {
	o.Secret = &v
}

// GetUrl returns the Url field value if set, zero value otherwise
func (o *ThirdPartyIntegration) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyIntegration) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}

	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ThirdPartyIntegration) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ThirdPartyIntegration) SetUrl(v string) {
	o.Url = &v
}

func (o ThirdPartyIntegration) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ThirdPartyIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.MicrosoftTeamsWebhookUrl) {
		toSerialize["microsoftTeamsWebhookUrl"] = o.MicrosoftTeamsWebhookUrl
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.LicenseKey) {
		toSerialize["licenseKey"] = o.LicenseKey
	}
	if !IsNil(o.ReadToken) {
		toSerialize["readToken"] = o.ReadToken
	}
	if !IsNil(o.WriteToken) {
		toSerialize["writeToken"] = o.WriteToken
	}
	if !IsNil(o.ServiceKey) {
		toSerialize["serviceKey"] = o.ServiceKey
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.ServiceDiscovery) {
		toSerialize["serviceDiscovery"] = o.ServiceDiscovery
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.ApiToken) {
		toSerialize["apiToken"] = o.ApiToken
	}
	if !IsNil(o.ChannelName) {
		toSerialize["channelName"] = o.ChannelName
	}
	if !IsNil(o.TeamName) {
		toSerialize["teamName"] = o.TeamName
	}
	if !IsNil(o.RoutingKey) {
		toSerialize["routingKey"] = o.RoutingKey
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}
