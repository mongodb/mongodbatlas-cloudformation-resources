// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// CloudSearchMetrics struct for CloudSearchMetrics
type CloudSearchMetrics struct {
	// Unique 24-hexadecimal digit string that identifies the project.
	// Read only field.
	GroupId string `json:"groupId"`
	// List that contains all host compute, memory, and storage utilization dedicated to Atlas Search when MongoDB Atlas received this request.
	// Read only field.
	HardwareMetrics *[]FTSMetric `json:"hardwareMetrics,omitempty"`
	// List that contains all performance and utilization measurements that Atlas Search index performed by the time MongoDB Atlas received this request.
	// Read only field.
	IndexMetrics *[]FTSMetric `json:"indexMetrics,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Hostname and port that identifies the process.
	// Read only field.
	ProcessId string `json:"processId"`
	// List that contains all available Atlas Search status metrics when MongoDB Atlas received this request.
	// Read only field.
	StatusMetrics *[]FTSMetric `json:"statusMetrics,omitempty"`
}

// NewCloudSearchMetrics instantiates a new CloudSearchMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSearchMetrics(groupId string, processId string) *CloudSearchMetrics {
	this := CloudSearchMetrics{}
	this.GroupId = groupId
	this.ProcessId = processId
	return &this
}

// NewCloudSearchMetricsWithDefaults instantiates a new CloudSearchMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSearchMetricsWithDefaults() *CloudSearchMetrics {
	this := CloudSearchMetrics{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *CloudSearchMetrics) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *CloudSearchMetrics) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *CloudSearchMetrics) SetGroupId(v string) {
	o.GroupId = v
}

// GetHardwareMetrics returns the HardwareMetrics field value if set, zero value otherwise
func (o *CloudSearchMetrics) GetHardwareMetrics() []FTSMetric {
	if o == nil || IsNil(o.HardwareMetrics) {
		var ret []FTSMetric
		return ret
	}
	return *o.HardwareMetrics
}

// GetHardwareMetricsOk returns a tuple with the HardwareMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSearchMetrics) GetHardwareMetricsOk() (*[]FTSMetric, bool) {
	if o == nil || IsNil(o.HardwareMetrics) {
		return nil, false
	}

	return o.HardwareMetrics, true
}

// HasHardwareMetrics returns a boolean if a field has been set.
func (o *CloudSearchMetrics) HasHardwareMetrics() bool {
	if o != nil && !IsNil(o.HardwareMetrics) {
		return true
	}

	return false
}

// SetHardwareMetrics gets a reference to the given []FTSMetric and assigns it to the HardwareMetrics field.
func (o *CloudSearchMetrics) SetHardwareMetrics(v []FTSMetric) {
	o.HardwareMetrics = &v
}

// GetIndexMetrics returns the IndexMetrics field value if set, zero value otherwise
func (o *CloudSearchMetrics) GetIndexMetrics() []FTSMetric {
	if o == nil || IsNil(o.IndexMetrics) {
		var ret []FTSMetric
		return ret
	}
	return *o.IndexMetrics
}

// GetIndexMetricsOk returns a tuple with the IndexMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSearchMetrics) GetIndexMetricsOk() (*[]FTSMetric, bool) {
	if o == nil || IsNil(o.IndexMetrics) {
		return nil, false
	}

	return o.IndexMetrics, true
}

// HasIndexMetrics returns a boolean if a field has been set.
func (o *CloudSearchMetrics) HasIndexMetrics() bool {
	if o != nil && !IsNil(o.IndexMetrics) {
		return true
	}

	return false
}

// SetIndexMetrics gets a reference to the given []FTSMetric and assigns it to the IndexMetrics field.
func (o *CloudSearchMetrics) SetIndexMetrics(v []FTSMetric) {
	o.IndexMetrics = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *CloudSearchMetrics) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSearchMetrics) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CloudSearchMetrics) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *CloudSearchMetrics) SetLinks(v []Link) {
	o.Links = &v
}

// GetProcessId returns the ProcessId field value
func (o *CloudSearchMetrics) GetProcessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value
// and a boolean to check if the value has been set.
func (o *CloudSearchMetrics) GetProcessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessId, true
}

// SetProcessId sets field value
func (o *CloudSearchMetrics) SetProcessId(v string) {
	o.ProcessId = v
}

// GetStatusMetrics returns the StatusMetrics field value if set, zero value otherwise
func (o *CloudSearchMetrics) GetStatusMetrics() []FTSMetric {
	if o == nil || IsNil(o.StatusMetrics) {
		var ret []FTSMetric
		return ret
	}
	return *o.StatusMetrics
}

// GetStatusMetricsOk returns a tuple with the StatusMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSearchMetrics) GetStatusMetricsOk() (*[]FTSMetric, bool) {
	if o == nil || IsNil(o.StatusMetrics) {
		return nil, false
	}

	return o.StatusMetrics, true
}

// HasStatusMetrics returns a boolean if a field has been set.
func (o *CloudSearchMetrics) HasStatusMetrics() bool {
	if o != nil && !IsNil(o.StatusMetrics) {
		return true
	}

	return false
}

// SetStatusMetrics gets a reference to the given []FTSMetric and assigns it to the StatusMetrics field.
func (o *CloudSearchMetrics) SetStatusMetrics(v []FTSMetric) {
	o.StatusMetrics = &v
}

func (o CloudSearchMetrics) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudSearchMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
