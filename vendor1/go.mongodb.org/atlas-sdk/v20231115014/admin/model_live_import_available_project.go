// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// LiveImportAvailableProject struct for LiveImportAvailableProject
type LiveImportAvailableProject struct {
	// List of clusters that can be migrated to MongoDB Cloud.
	Deployments *[]AvailableClustersDeployment `json:"deployments,omitempty"`
	// Hostname of MongoDB Agent list that you configured to perform a migration.
	MigrationHosts *[]string `json:"migrationHosts,omitempty"`
	// Human-readable label that identifies this project.
	// Read only field.
	Name string `json:"name"`
	// Unique 24-hexadecimal digit string that identifies the project to be migrated.
	// Read only field.
	ProjectId string `json:"projectId"`
}

// NewLiveImportAvailableProject instantiates a new LiveImportAvailableProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveImportAvailableProject(name string, projectId string) *LiveImportAvailableProject {
	this := LiveImportAvailableProject{}
	this.Name = name
	this.ProjectId = projectId
	return &this
}

// NewLiveImportAvailableProjectWithDefaults instantiates a new LiveImportAvailableProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveImportAvailableProjectWithDefaults() *LiveImportAvailableProject {
	this := LiveImportAvailableProject{}
	return &this
}

// GetDeployments returns the Deployments field value if set, zero value otherwise
func (o *LiveImportAvailableProject) GetDeployments() []AvailableClustersDeployment {
	if o == nil || IsNil(o.Deployments) {
		var ret []AvailableClustersDeployment
		return ret
	}
	return *o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveImportAvailableProject) GetDeploymentsOk() (*[]AvailableClustersDeployment, bool) {
	if o == nil || IsNil(o.Deployments) {
		return nil, false
	}

	return o.Deployments, true
}

// HasDeployments returns a boolean if a field has been set.
func (o *LiveImportAvailableProject) HasDeployments() bool {
	if o != nil && !IsNil(o.Deployments) {
		return true
	}

	return false
}

// SetDeployments gets a reference to the given []AvailableClustersDeployment and assigns it to the Deployments field.
func (o *LiveImportAvailableProject) SetDeployments(v []AvailableClustersDeployment) {
	o.Deployments = &v
}

// GetMigrationHosts returns the MigrationHosts field value if set, zero value otherwise
func (o *LiveImportAvailableProject) GetMigrationHosts() []string {
	if o == nil || IsNil(o.MigrationHosts) {
		var ret []string
		return ret
	}
	return *o.MigrationHosts
}

// GetMigrationHostsOk returns a tuple with the MigrationHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveImportAvailableProject) GetMigrationHostsOk() (*[]string, bool) {
	if o == nil || IsNil(o.MigrationHosts) {
		return nil, false
	}

	return o.MigrationHosts, true
}

// HasMigrationHosts returns a boolean if a field has been set.
func (o *LiveImportAvailableProject) HasMigrationHosts() bool {
	if o != nil && !IsNil(o.MigrationHosts) {
		return true
	}

	return false
}

// SetMigrationHosts gets a reference to the given []string and assigns it to the MigrationHosts field.
func (o *LiveImportAvailableProject) SetMigrationHosts(v []string) {
	o.MigrationHosts = &v
}

// GetName returns the Name field value
func (o *LiveImportAvailableProject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LiveImportAvailableProject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LiveImportAvailableProject) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value
func (o *LiveImportAvailableProject) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *LiveImportAvailableProject) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *LiveImportAvailableProject) SetProjectId(v string) {
	o.ProjectId = v
}

func (o LiveImportAvailableProject) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o LiveImportAvailableProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deployments) {
		toSerialize["deployments"] = o.Deployments
	}
	if !IsNil(o.MigrationHosts) {
		toSerialize["migrationHosts"] = o.MigrationHosts
	}
	return toSerialize, nil
}
