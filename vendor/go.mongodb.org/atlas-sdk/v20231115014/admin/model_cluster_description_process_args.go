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

// ClusterDescriptionProcessArgs struct for ClusterDescriptionProcessArgs
type ClusterDescriptionProcessArgs struct {
	// Number of threads on the source shard and the receiving shard for chunk migration. The number of threads should not exceed the half the total number of CPU cores in the sharded cluster.
	ChunkMigrationConcurrency *int `json:"chunkMigrationConcurrency,omitempty"`
	// Default level of acknowledgment requested from MongoDB for read operations set for this cluster.  MongoDB 4.4 clusters default to `available`. MongoDB 5.0 and later clusters default to `local`.
	DefaultReadConcern *string `json:"defaultReadConcern,omitempty"`
	// Default level of acknowledgment requested from MongoDB for write operations when none is specified by the driver.
	DefaultWriteConcern *string `json:"defaultWriteConcern,omitempty"`
	// Flag that indicates whether you can insert or update documents where all indexed entries don't exceed 1024 bytes. If you set this to false, [mongod](https://docs.mongodb.com/upcoming/reference/program/mongod/#mongodb-binary-bin.mongod) writes documents that exceed this limit but doesn't index them. This parameter has been removed as of [MongoDB 4.4](https://www.mongodb.com/docs/manual/reference/parameters/#mongodb-parameter-param.failIndexKeyTooLong).
	// Deprecated
	FailIndexKeyTooLong *bool `json:"failIndexKeyTooLong,omitempty"`
	// Flag that indicates whether the cluster allows execution of operations that perform server-side executions of JavaScript.
	JavascriptEnabled *bool `json:"javascriptEnabled,omitempty"`
	// Minimum Transport Layer Security (TLS) version that the cluster accepts for incoming connections. Clusters using TLS 1.0 or 1.1 should consider setting TLS 1.2 as the minimum TLS protocol version.
	MinimumEnabledTlsProtocol *string `json:"minimumEnabledTlsProtocol,omitempty"`
	// Flag that indicates whether the cluster disables executing any query that requires a collection scan to return results.
	NoTableScan *bool `json:"noTableScan,omitempty"`
	// Minimum retention window for cluster's oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates.
	OplogMinRetentionHours *float64 `json:"oplogMinRetentionHours,omitempty"`
	// Storage limit of cluster's oplog expressed in megabytes. A value of null indicates that the cluster uses the default oplog size that MongoDB Cloud calculates.
	OplogSizeMB *int `json:"oplogSizeMB,omitempty"`
	// May be set to 1 (disabled) or 3 (enabled). When set to 3, Atlas will include redacted and anonymized $queryStats output in MongoDB logs. $queryStats output does not contain literals or field values. Enabling this setting might impact the performance of your cluster.
	QueryStatsLogVerbosity *int `json:"queryStatsLogVerbosity,omitempty"`
	// Interval in seconds at which the mongosqld process re-samples data to create its relational schema.
	SampleRefreshIntervalBIConnector *int `json:"sampleRefreshIntervalBIConnector,omitempty"`
	// Number of documents per database to sample when gathering schema information.
	SampleSizeBIConnector *int `json:"sampleSizeBIConnector,omitempty"`
	// Lifetime, in seconds, of multi-document transactions. Atlas considers the transactions that exceed this limit as expired and so aborts them through a periodic cleanup process.
	TransactionLifetimeLimitSeconds *int64 `json:"transactionLifetimeLimitSeconds,omitempty"`
}

// NewClusterDescriptionProcessArgs instantiates a new ClusterDescriptionProcessArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDescriptionProcessArgs() *ClusterDescriptionProcessArgs {
	this := ClusterDescriptionProcessArgs{}
	var defaultReadConcern string = "available"
	this.DefaultReadConcern = &defaultReadConcern
	var failIndexKeyTooLong bool = true
	this.FailIndexKeyTooLong = &failIndexKeyTooLong
	var sampleRefreshIntervalBIConnector int = 0
	this.SampleRefreshIntervalBIConnector = &sampleRefreshIntervalBIConnector
	return &this
}

// NewClusterDescriptionProcessArgsWithDefaults instantiates a new ClusterDescriptionProcessArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDescriptionProcessArgsWithDefaults() *ClusterDescriptionProcessArgs {
	this := ClusterDescriptionProcessArgs{}
	var defaultReadConcern string = "available"
	this.DefaultReadConcern = &defaultReadConcern
	var failIndexKeyTooLong bool = true
	this.FailIndexKeyTooLong = &failIndexKeyTooLong
	var sampleRefreshIntervalBIConnector int = 0
	this.SampleRefreshIntervalBIConnector = &sampleRefreshIntervalBIConnector
	return &this
}

// GetChunkMigrationConcurrency returns the ChunkMigrationConcurrency field value if set, zero value otherwise
func (o *ClusterDescriptionProcessArgs) GetChunkMigrationConcurrency() int {
	if o == nil || IsNil(o.ChunkMigrationConcurrency) {
		var ret int
		return ret
	}
	return *o.ChunkMigrationConcurrency
}

// GetChunkMigrationConcurrencyOk returns a tuple with the ChunkMigrationConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionProcessArgs) GetChunkMigrationConcurrencyOk() (*int, bool) {
	if o == nil || IsNil(o.ChunkMigrationConcurrency) {
		return nil, false
	}

	return o.ChunkMigrationConcurrency, true
}

// HasChunkMigrationConcurrency returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasChunkMigrationConcurrency() bool {
	if o != nil && !IsNil(o.ChunkMigrationConcurrency) {
		return true
	}

	return false
}

// SetChunkMigrationConcurrency gets a reference to the given int and assigns it to the ChunkMigrationConcurrency field.
func (o *ClusterDescriptionProcessArgs) SetChunkMigrationConcurrency(v int) {
	o.ChunkMigrationConcurrency = &v
}

// GetDefaultReadConcern returns the DefaultReadConcern field value if set, zero value otherwise
func (o *ClusterDescriptionProcessArgs) GetDefaultReadConcern() string {
	if o == nil || IsNil(o.DefaultReadConcern) {
		var ret string
		return ret
	}
	return *o.DefaultReadConcern
}

// GetDefaultReadConcernOk returns a tuple with the DefaultReadConcern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionProcessArgs) GetDefaultReadConcernOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultReadConcern) {
		return nil, false
	}

	return o.DefaultReadConcern, true
}

// HasDefaultReadConcern returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasDefaultReadConcern() bool {
	if o != nil && !IsNil(o.DefaultReadConcern) {
		return true
	}

	return false
}

// SetDefaultReadConcern gets a reference to the given string and assigns it to the DefaultReadConcern field.
func (o *ClusterDescriptionProcessArgs) SetDefaultReadConcern(v string) {
	o.DefaultReadConcern = &v
}

// GetDefaultWriteConcern returns the DefaultWriteConcern field value if set, zero value otherwise
func (o *ClusterDescriptionProcessArgs) GetDefaultWriteConcern() string {
	if o == nil || IsNil(o.DefaultWriteConcern) {
		var ret string
		return ret
	}
	return *o.DefaultWriteConcern
}

// GetDefaultWriteConcernOk returns a tuple with the DefaultWriteConcern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionProcessArgs) GetDefaultWriteConcernOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultWriteConcern) {
		return nil, false
	}

	return o.DefaultWriteConcern, true
}

// HasDefaultWriteConcern returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasDefaultWriteConcern() bool {
	if o != nil && !IsNil(o.DefaultWriteConcern) {
		return true
	}

	return false
}

// SetDefaultWriteConcern gets a reference to the given string and assigns it to the DefaultWriteConcern field.
func (o *ClusterDescriptionProcessArgs) SetDefaultWriteConcern(v string) {
	o.DefaultWriteConcern = &v
}

// GetFailIndexKeyTooLong returns the FailIndexKeyTooLong field value if set, zero value otherwise
// Deprecated
func (o *ClusterDescriptionProcessArgs) GetFailIndexKeyTooLong() bool {
	if o == nil || IsNil(o.FailIndexKeyTooLong) {
		var ret bool
		return ret
	}
	return *o.FailIndexKeyTooLong
}

// GetFailIndexKeyTooLongOk returns a tuple with the FailIndexKeyTooLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ClusterDescriptionProcessArgs) GetFailIndexKeyTooLongOk() (*bool, bool) {
	if o == nil || IsNil(o.FailIndexKeyTooLong) {
		return nil, false
	}

	return o.FailIndexKeyTooLong, true
}

// HasFailIndexKeyTooLong returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasFailIndexKeyTooLong() bool {
	if o != nil && !IsNil(o.FailIndexKeyTooLong) {
		return true
	}

	return false
}

// SetFailIndexKeyTooLong gets a reference to the given bool and assigns it to the FailIndexKeyTooLong field.
// Deprecated
func (o *ClusterDescriptionProcessArgs) SetFailIndexKeyTooLong(v bool) {
	o.FailIndexKeyTooLong = &v
}

// GetJavascriptEnabled returns the JavascriptEnabled field value if set, zero value otherwise
func (o *ClusterDescriptionProcessArgs) GetJavascriptEnabled() bool {
	if o == nil || IsNil(o.JavascriptEnabled) {
		var ret bool
		return ret
	}
	return *o.JavascriptEnabled
}

// GetJavascriptEnabledOk returns a tuple with the JavascriptEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionProcessArgs) GetJavascriptEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.JavascriptEnabled) {
		return nil, false
	}

	return o.JavascriptEnabled, true
}

// HasJavascriptEnabled returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasJavascriptEnabled() bool {
	if o != nil && !IsNil(o.JavascriptEnabled) {
		return true
	}

	return false
}

// SetJavascriptEnabled gets a reference to the given bool and assigns it to the JavascriptEnabled field.
func (o *ClusterDescriptionProcessArgs) SetJavascriptEnabled(v bool) {
	o.JavascriptEnabled = &v
}

// GetMinimumEnabledTlsProtocol returns the MinimumEnabledTlsProtocol field value if set, zero value otherwise
func (o *ClusterDescriptionProcessArgs) GetMinimumEnabledTlsProtocol() string {
	if o == nil || IsNil(o.MinimumEnabledTlsProtocol) {
		var ret string
		return ret
	}
	return *o.MinimumEnabledTlsProtocol
}

// GetMinimumEnabledTlsProtocolOk returns a tuple with the MinimumEnabledTlsProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionProcessArgs) GetMinimumEnabledTlsProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumEnabledTlsProtocol) {
		return nil, false
	}

	return o.MinimumEnabledTlsProtocol, true
}

// HasMinimumEnabledTlsProtocol returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasMinimumEnabledTlsProtocol() bool {
	if o != nil && !IsNil(o.MinimumEnabledTlsProtocol) {
		return true
	}

	return false
}

// SetMinimumEnabledTlsProtocol gets a reference to the given string and assigns it to the MinimumEnabledTlsProtocol field.
func (o *ClusterDescriptionProcessArgs) SetMinimumEnabledTlsProtocol(v string) {
	o.MinimumEnabledTlsProtocol = &v
}

// GetNoTableScan returns the NoTableScan field value if set, zero value otherwise
func (o *ClusterDescriptionProcessArgs) GetNoTableScan() bool {
	if o == nil || IsNil(o.NoTableScan) {
		var ret bool
		return ret
	}
	return *o.NoTableScan
}

// GetNoTableScanOk returns a tuple with the NoTableScan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionProcessArgs) GetNoTableScanOk() (*bool, bool) {
	if o == nil || IsNil(o.NoTableScan) {
		return nil, false
	}

	return o.NoTableScan, true
}

// HasNoTableScan returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasNoTableScan() bool {
	if o != nil && !IsNil(o.NoTableScan) {
		return true
	}

	return false
}

// SetNoTableScan gets a reference to the given bool and assigns it to the NoTableScan field.
func (o *ClusterDescriptionProcessArgs) SetNoTableScan(v bool) {
	o.NoTableScan = &v
}

// GetOplogMinRetentionHours returns the OplogMinRetentionHours field value if set, zero value otherwise
func (o *ClusterDescriptionProcessArgs) GetOplogMinRetentionHours() float64 {
	if o == nil || IsNil(o.OplogMinRetentionHours) {
		var ret float64
		return ret
	}
	return *o.OplogMinRetentionHours
}

// GetOplogMinRetentionHoursOk returns a tuple with the OplogMinRetentionHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionProcessArgs) GetOplogMinRetentionHoursOk() (*float64, bool) {
	if o == nil || IsNil(o.OplogMinRetentionHours) {
		return nil, false
	}

	return o.OplogMinRetentionHours, true
}

// HasOplogMinRetentionHours returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasOplogMinRetentionHours() bool {
	if o != nil && !IsNil(o.OplogMinRetentionHours) {
		return true
	}

	return false
}

// SetOplogMinRetentionHours gets a reference to the given float64 and assigns it to the OplogMinRetentionHours field.
func (o *ClusterDescriptionProcessArgs) SetOplogMinRetentionHours(v float64) {
	o.OplogMinRetentionHours = &v
}

// GetOplogSizeMB returns the OplogSizeMB field value if set, zero value otherwise
func (o *ClusterDescriptionProcessArgs) GetOplogSizeMB() int {
	if o == nil || IsNil(o.OplogSizeMB) {
		var ret int
		return ret
	}
	return *o.OplogSizeMB
}

// GetOplogSizeMBOk returns a tuple with the OplogSizeMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionProcessArgs) GetOplogSizeMBOk() (*int, bool) {
	if o == nil || IsNil(o.OplogSizeMB) {
		return nil, false
	}

	return o.OplogSizeMB, true
}

// HasOplogSizeMB returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasOplogSizeMB() bool {
	if o != nil && !IsNil(o.OplogSizeMB) {
		return true
	}

	return false
}

// SetOplogSizeMB gets a reference to the given int and assigns it to the OplogSizeMB field.
func (o *ClusterDescriptionProcessArgs) SetOplogSizeMB(v int) {
	o.OplogSizeMB = &v
}

// GetQueryStatsLogVerbosity returns the QueryStatsLogVerbosity field value if set, zero value otherwise
func (o *ClusterDescriptionProcessArgs) GetQueryStatsLogVerbosity() int {
	if o == nil || IsNil(o.QueryStatsLogVerbosity) {
		var ret int
		return ret
	}
	return *o.QueryStatsLogVerbosity
}

// GetQueryStatsLogVerbosityOk returns a tuple with the QueryStatsLogVerbosity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionProcessArgs) GetQueryStatsLogVerbosityOk() (*int, bool) {
	if o == nil || IsNil(o.QueryStatsLogVerbosity) {
		return nil, false
	}

	return o.QueryStatsLogVerbosity, true
}

// HasQueryStatsLogVerbosity returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasQueryStatsLogVerbosity() bool {
	if o != nil && !IsNil(o.QueryStatsLogVerbosity) {
		return true
	}

	return false
}

// SetQueryStatsLogVerbosity gets a reference to the given int and assigns it to the QueryStatsLogVerbosity field.
func (o *ClusterDescriptionProcessArgs) SetQueryStatsLogVerbosity(v int) {
	o.QueryStatsLogVerbosity = &v
}

// GetSampleRefreshIntervalBIConnector returns the SampleRefreshIntervalBIConnector field value if set, zero value otherwise
func (o *ClusterDescriptionProcessArgs) GetSampleRefreshIntervalBIConnector() int {
	if o == nil || IsNil(o.SampleRefreshIntervalBIConnector) {
		var ret int
		return ret
	}
	return *o.SampleRefreshIntervalBIConnector
}

// GetSampleRefreshIntervalBIConnectorOk returns a tuple with the SampleRefreshIntervalBIConnector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionProcessArgs) GetSampleRefreshIntervalBIConnectorOk() (*int, bool) {
	if o == nil || IsNil(o.SampleRefreshIntervalBIConnector) {
		return nil, false
	}

	return o.SampleRefreshIntervalBIConnector, true
}

// HasSampleRefreshIntervalBIConnector returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasSampleRefreshIntervalBIConnector() bool {
	if o != nil && !IsNil(o.SampleRefreshIntervalBIConnector) {
		return true
	}

	return false
}

// SetSampleRefreshIntervalBIConnector gets a reference to the given int and assigns it to the SampleRefreshIntervalBIConnector field.
func (o *ClusterDescriptionProcessArgs) SetSampleRefreshIntervalBIConnector(v int) {
	o.SampleRefreshIntervalBIConnector = &v
}

// GetSampleSizeBIConnector returns the SampleSizeBIConnector field value if set, zero value otherwise
func (o *ClusterDescriptionProcessArgs) GetSampleSizeBIConnector() int {
	if o == nil || IsNil(o.SampleSizeBIConnector) {
		var ret int
		return ret
	}
	return *o.SampleSizeBIConnector
}

// GetSampleSizeBIConnectorOk returns a tuple with the SampleSizeBIConnector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionProcessArgs) GetSampleSizeBIConnectorOk() (*int, bool) {
	if o == nil || IsNil(o.SampleSizeBIConnector) {
		return nil, false
	}

	return o.SampleSizeBIConnector, true
}

// HasSampleSizeBIConnector returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasSampleSizeBIConnector() bool {
	if o != nil && !IsNil(o.SampleSizeBIConnector) {
		return true
	}

	return false
}

// SetSampleSizeBIConnector gets a reference to the given int and assigns it to the SampleSizeBIConnector field.
func (o *ClusterDescriptionProcessArgs) SetSampleSizeBIConnector(v int) {
	o.SampleSizeBIConnector = &v
}

// GetTransactionLifetimeLimitSeconds returns the TransactionLifetimeLimitSeconds field value if set, zero value otherwise
func (o *ClusterDescriptionProcessArgs) GetTransactionLifetimeLimitSeconds() int64 {
	if o == nil || IsNil(o.TransactionLifetimeLimitSeconds) {
		var ret int64
		return ret
	}
	return *o.TransactionLifetimeLimitSeconds
}

// GetTransactionLifetimeLimitSecondsOk returns a tuple with the TransactionLifetimeLimitSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionProcessArgs) GetTransactionLifetimeLimitSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactionLifetimeLimitSeconds) {
		return nil, false
	}

	return o.TransactionLifetimeLimitSeconds, true
}

// HasTransactionLifetimeLimitSeconds returns a boolean if a field has been set.
func (o *ClusterDescriptionProcessArgs) HasTransactionLifetimeLimitSeconds() bool {
	if o != nil && !IsNil(o.TransactionLifetimeLimitSeconds) {
		return true
	}

	return false
}

// SetTransactionLifetimeLimitSeconds gets a reference to the given int64 and assigns it to the TransactionLifetimeLimitSeconds field.
func (o *ClusterDescriptionProcessArgs) SetTransactionLifetimeLimitSeconds(v int64) {
	o.TransactionLifetimeLimitSeconds = &v
}

func (o ClusterDescriptionProcessArgs) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterDescriptionProcessArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChunkMigrationConcurrency) {
		toSerialize["chunkMigrationConcurrency"] = o.ChunkMigrationConcurrency
	}
	if !IsNil(o.DefaultReadConcern) {
		toSerialize["defaultReadConcern"] = o.DefaultReadConcern
	}
	if !IsNil(o.DefaultWriteConcern) {
		toSerialize["defaultWriteConcern"] = o.DefaultWriteConcern
	}
	if !IsNil(o.FailIndexKeyTooLong) {
		toSerialize["failIndexKeyTooLong"] = o.FailIndexKeyTooLong
	}
	if !IsNil(o.JavascriptEnabled) {
		toSerialize["javascriptEnabled"] = o.JavascriptEnabled
	}
	if !IsNil(o.MinimumEnabledTlsProtocol) {
		toSerialize["minimumEnabledTlsProtocol"] = o.MinimumEnabledTlsProtocol
	}
	if !IsNil(o.NoTableScan) {
		toSerialize["noTableScan"] = o.NoTableScan
	}
	if !IsNil(o.OplogMinRetentionHours) {
		toSerialize["oplogMinRetentionHours"] = o.OplogMinRetentionHours
	}
	if !IsNil(o.OplogSizeMB) {
		toSerialize["oplogSizeMB"] = o.OplogSizeMB
	}
	if !IsNil(o.QueryStatsLogVerbosity) {
		toSerialize["queryStatsLogVerbosity"] = o.QueryStatsLogVerbosity
	}
	if !IsNil(o.SampleRefreshIntervalBIConnector) {
		toSerialize["sampleRefreshIntervalBIConnector"] = o.SampleRefreshIntervalBIConnector
	}
	if !IsNil(o.SampleSizeBIConnector) {
		toSerialize["sampleSizeBIConnector"] = o.SampleSizeBIConnector
	}
	if !IsNil(o.TransactionLifetimeLimitSeconds) {
		toSerialize["transactionLifetimeLimitSeconds"] = o.TransactionLifetimeLimitSeconds
	}
	return toSerialize, nil
}
