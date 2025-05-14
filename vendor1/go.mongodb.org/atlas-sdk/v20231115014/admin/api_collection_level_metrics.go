// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type CollectionLevelMetricsApi interface {

	/*
		GetCollStatsLatencyNamespaceClusterMeasurements Return Cluster-Level Query Latency

		[experimental] Get a list of the Coll Stats Latency cluster-level measurements for the given namespace.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster to retrieve metrics for.
		@param clusterView Human-readable label that identifies the cluster topology to retrieve metrics for.
		@param databaseName Human-readable label that identifies the database.
		@param collectionName Human-readable label that identifies the collection.
		@return GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest
	*/
	GetCollStatsLatencyNamespaceClusterMeasurements(ctx context.Context, groupId string, clusterName string, clusterView string, databaseName string, collectionName string) GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest
	/*
		GetCollStatsLatencyNamespaceClusterMeasurements Return Cluster-Level Query Latency


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetCollStatsLatencyNamespaceClusterMeasurementsApiParams - Parameters for the request
		@return GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest
	*/
	GetCollStatsLatencyNamespaceClusterMeasurementsWithParams(ctx context.Context, args *GetCollStatsLatencyNamespaceClusterMeasurementsApiParams) GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest

	// Method available only for mocking purposes
	GetCollStatsLatencyNamespaceClusterMeasurementsExecute(r GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest) (*MeasurementsCollStatsLatencyCluster, *http.Response, error)

	/*
		GetCollStatsLatencyNamespaceHostMeasurements Return Host-Level Query Latency

		[experimental] Get a list of the Coll Stats Latency process-level measurements for the given namespace

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param processId Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
		@param databaseName Human-readable label that identifies the database.
		@param collectionName Human-readable label that identifies the collection.
		@return GetCollStatsLatencyNamespaceHostMeasurementsApiRequest
	*/
	GetCollStatsLatencyNamespaceHostMeasurements(ctx context.Context, groupId string, processId string, databaseName string, collectionName string) GetCollStatsLatencyNamespaceHostMeasurementsApiRequest
	/*
		GetCollStatsLatencyNamespaceHostMeasurements Return Host-Level Query Latency


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetCollStatsLatencyNamespaceHostMeasurementsApiParams - Parameters for the request
		@return GetCollStatsLatencyNamespaceHostMeasurementsApiRequest
	*/
	GetCollStatsLatencyNamespaceHostMeasurementsWithParams(ctx context.Context, args *GetCollStatsLatencyNamespaceHostMeasurementsApiParams) GetCollStatsLatencyNamespaceHostMeasurementsApiRequest

	// Method available only for mocking purposes
	GetCollStatsLatencyNamespaceHostMeasurementsExecute(r GetCollStatsLatencyNamespaceHostMeasurementsApiRequest) (*MeasurementsCollStatsLatencyHost, *http.Response, error)

	/*
		GetCollStatsLatencyNamespaceMetrics Return all metric names

		[experimental] Returns all available Coll Stats Latency metric names and their respective units for the specified project at the time of request.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetCollStatsLatencyNamespaceMetricsApiRequest
	*/
	GetCollStatsLatencyNamespaceMetrics(ctx context.Context, groupId string) GetCollStatsLatencyNamespaceMetricsApiRequest
	/*
		GetCollStatsLatencyNamespaceMetrics Return all metric names


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetCollStatsLatencyNamespaceMetricsApiParams - Parameters for the request
		@return GetCollStatsLatencyNamespaceMetricsApiRequest
	*/
	GetCollStatsLatencyNamespaceMetricsWithParams(ctx context.Context, args *GetCollStatsLatencyNamespaceMetricsApiParams) GetCollStatsLatencyNamespaceMetricsApiRequest

	// Method available only for mocking purposes
	GetCollStatsLatencyNamespaceMetricsExecute(r GetCollStatsLatencyNamespaceMetricsApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetCollStatsLatencyNamespacesForCluster Return Ranked Namespaces from a Cluster

		[experimental] Return the subset of namespaces from the given cluster sorted by highest total execution time (descending) within the given time window.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster to pin namespaces to.
		@param clusterView Human-readable label that identifies the cluster topology to retrieve metrics for.
		@return GetCollStatsLatencyNamespacesForClusterApiRequest
	*/
	GetCollStatsLatencyNamespacesForCluster(ctx context.Context, groupId string, clusterName string, clusterView string) GetCollStatsLatencyNamespacesForClusterApiRequest
	/*
		GetCollStatsLatencyNamespacesForCluster Return Ranked Namespaces from a Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetCollStatsLatencyNamespacesForClusterApiParams - Parameters for the request
		@return GetCollStatsLatencyNamespacesForClusterApiRequest
	*/
	GetCollStatsLatencyNamespacesForClusterWithParams(ctx context.Context, args *GetCollStatsLatencyNamespacesForClusterApiParams) GetCollStatsLatencyNamespacesForClusterApiRequest

	// Method available only for mocking purposes
	GetCollStatsLatencyNamespacesForClusterExecute(r GetCollStatsLatencyNamespacesForClusterApiRequest) (*CollStatsRankedNamespaces, *http.Response, error)

	/*
		GetCollStatsLatencyNamespacesForHost Return Ranked Namespaces from a Host

		[experimental] Return the subset of namespaces from the given process ranked by highest total execution time (descending) within the given time window.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param processId Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
		@return GetCollStatsLatencyNamespacesForHostApiRequest
	*/
	GetCollStatsLatencyNamespacesForHost(ctx context.Context, groupId string, processId string) GetCollStatsLatencyNamespacesForHostApiRequest
	/*
		GetCollStatsLatencyNamespacesForHost Return Ranked Namespaces from a Host


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetCollStatsLatencyNamespacesForHostApiParams - Parameters for the request
		@return GetCollStatsLatencyNamespacesForHostApiRequest
	*/
	GetCollStatsLatencyNamespacesForHostWithParams(ctx context.Context, args *GetCollStatsLatencyNamespacesForHostApiParams) GetCollStatsLatencyNamespacesForHostApiRequest

	// Method available only for mocking purposes
	GetCollStatsLatencyNamespacesForHostExecute(r GetCollStatsLatencyNamespacesForHostApiRequest) (*CollStatsRankedNamespaces, *http.Response, error)

	/*
		GetPinnedNamespaces Return Pinned Namespaces

		[experimental] Returns a list of given cluster's pinned namespaces, a set of namespaces manually selected by users to collect query latency metrics on.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster to retrieve pinned namespaces for.
		@return GetPinnedNamespacesApiRequest
	*/
	GetPinnedNamespaces(ctx context.Context, groupId string, clusterName string) GetPinnedNamespacesApiRequest
	/*
		GetPinnedNamespaces Return Pinned Namespaces


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetPinnedNamespacesApiParams - Parameters for the request
		@return GetPinnedNamespacesApiRequest
	*/
	GetPinnedNamespacesWithParams(ctx context.Context, args *GetPinnedNamespacesApiParams) GetPinnedNamespacesApiRequest

	// Method available only for mocking purposes
	GetPinnedNamespacesExecute(r GetPinnedNamespacesApiRequest) (*PinnedNamespaces, *http.Response, error)

	/*
		PinNamespacesPatch Add Pinned Namespaces

		[experimental] Add provided list of namespaces to existing pinned namespaces list for collection-level latency metrics collection for the given Group and Cluster

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster to pin namespaces to.
		@return PinNamespacesPatchApiRequest
	*/
	PinNamespacesPatch(ctx context.Context, groupId string, clusterName string, namespacesRequest *NamespacesRequest) PinNamespacesPatchApiRequest
	/*
		PinNamespacesPatch Add Pinned Namespaces


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param PinNamespacesPatchApiParams - Parameters for the request
		@return PinNamespacesPatchApiRequest
	*/
	PinNamespacesPatchWithParams(ctx context.Context, args *PinNamespacesPatchApiParams) PinNamespacesPatchApiRequest

	// Method available only for mocking purposes
	PinNamespacesPatchExecute(r PinNamespacesPatchApiRequest) (*PinnedNamespaces, *http.Response, error)

	/*
		PinNamespacesPut Pin Namespaces

		[experimental] Pin provided list of namespaces for collection-level latency metrics collection for the given Group and Cluster. This initializes a pinned namespaces list or replaces any existing pinned namespaces list for the Group and Cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster to pin namespaces to.
		@return PinNamespacesPutApiRequest
	*/
	PinNamespacesPut(ctx context.Context, groupId string, clusterName string, namespacesRequest *NamespacesRequest) PinNamespacesPutApiRequest
	/*
		PinNamespacesPut Pin Namespaces


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param PinNamespacesPutApiParams - Parameters for the request
		@return PinNamespacesPutApiRequest
	*/
	PinNamespacesPutWithParams(ctx context.Context, args *PinNamespacesPutApiParams) PinNamespacesPutApiRequest

	// Method available only for mocking purposes
	PinNamespacesPutExecute(r PinNamespacesPutApiRequest) (*PinnedNamespaces, *http.Response, error)

	/*
		UnpinNamespaces Unpin namespaces

		[experimental] Unpin provided list of namespaces for collection-level latency metrics collection for the given Group and Cluster

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster to unpin namespaces from.
		@return UnpinNamespacesApiRequest
	*/
	UnpinNamespaces(ctx context.Context, groupId string, clusterName string, namespacesRequest *NamespacesRequest) UnpinNamespacesApiRequest
	/*
		UnpinNamespaces Unpin namespaces


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UnpinNamespacesApiParams - Parameters for the request
		@return UnpinNamespacesApiRequest
	*/
	UnpinNamespacesWithParams(ctx context.Context, args *UnpinNamespacesApiParams) UnpinNamespacesApiRequest

	// Method available only for mocking purposes
	UnpinNamespacesExecute(r UnpinNamespacesApiRequest) (*PinnedNamespaces, *http.Response, error)
}

// CollectionLevelMetricsApiService CollectionLevelMetricsApi service
type CollectionLevelMetricsApiService service

type GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest struct {
	ctx            context.Context
	ApiService     CollectionLevelMetricsApi
	groupId        string
	clusterName    string
	clusterView    string
	databaseName   string
	collectionName string
	metrics        *[]string
	start          *time.Time
	end            *time.Time
	period         *string
}

type GetCollStatsLatencyNamespaceClusterMeasurementsApiParams struct {
	GroupId        string
	ClusterName    string
	ClusterView    string
	DatabaseName   string
	CollectionName string
	Metrics        *[]string
	Start          *time.Time
	End            *time.Time
	Period         *string
}

func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespaceClusterMeasurementsWithParams(ctx context.Context, args *GetCollStatsLatencyNamespaceClusterMeasurementsApiParams) GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest {
	return GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		clusterName:    args.ClusterName,
		clusterView:    args.ClusterView,
		databaseName:   args.DatabaseName,
		collectionName: args.CollectionName,
		metrics:        args.Metrics,
		start:          args.Start,
		end:            args.End,
		period:         args.Period,
	}
}

// List that contains the metrics that you want to retrieve for the associated data series. If you don&#39;t set this parameter, this resource returns data series for all Coll Stats Latency metrics.
func (r GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest) Metrics(metrics []string) GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest {
	r.metrics = &metrics
	return r
}

// Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest) Start(start time.Time) GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest {
	r.start = &start
	return r
}

// Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest) End(end time.Time) GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest {
	r.end = &end
	return r
}

// Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.
func (r GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest) Period(period string) GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest {
	r.period = &period
	return r
}

func (r GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest) Execute() (*MeasurementsCollStatsLatencyCluster, *http.Response, error) {
	return r.ApiService.GetCollStatsLatencyNamespaceClusterMeasurementsExecute(r)
}

/*
GetCollStatsLatencyNamespaceClusterMeasurements Return Cluster-Level Query Latency

[experimental] Get a list of the Coll Stats Latency cluster-level measurements for the given namespace.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster to retrieve metrics for.
	@param clusterView Human-readable label that identifies the cluster topology to retrieve metrics for.
	@param databaseName Human-readable label that identifies the database.
	@param collectionName Human-readable label that identifies the collection.
	@return GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest
*/
func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespaceClusterMeasurements(ctx context.Context, groupId string, clusterName string, clusterView string, databaseName string, collectionName string) GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest {
	return GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		clusterName:    clusterName,
		clusterView:    clusterView,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

// Execute executes the request
//
//	@return MeasurementsCollStatsLatencyCluster
func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespaceClusterMeasurementsExecute(r GetCollStatsLatencyNamespaceClusterMeasurementsApiRequest) (*MeasurementsCollStatsLatencyCluster, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MeasurementsCollStatsLatencyCluster
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionLevelMetricsApiService.GetCollStatsLatencyNamespaceClusterMeasurements")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/{clusterView}/{databaseName}/{collectionName}/collStats/measurements"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterView"+"}", url.PathEscape(parameterValueToString(r.clusterView, "clusterView")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(parameterValueToString(r.databaseName, "databaseName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"collectionName"+"}", url.PathEscape(parameterValueToString(r.collectionName, "collectionName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.metrics != nil {
		t := *r.metrics
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", t, "multi")

	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetCollStatsLatencyNamespaceHostMeasurementsApiRequest struct {
	ctx            context.Context
	ApiService     CollectionLevelMetricsApi
	groupId        string
	processId      string
	databaseName   string
	collectionName string
	metrics        *[]string
	start          *time.Time
	end            *time.Time
	period         *string
}

type GetCollStatsLatencyNamespaceHostMeasurementsApiParams struct {
	GroupId        string
	ProcessId      string
	DatabaseName   string
	CollectionName string
	Metrics        *[]string
	Start          *time.Time
	End            *time.Time
	Period         *string
}

func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespaceHostMeasurementsWithParams(ctx context.Context, args *GetCollStatsLatencyNamespaceHostMeasurementsApiParams) GetCollStatsLatencyNamespaceHostMeasurementsApiRequest {
	return GetCollStatsLatencyNamespaceHostMeasurementsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		processId:      args.ProcessId,
		databaseName:   args.DatabaseName,
		collectionName: args.CollectionName,
		metrics:        args.Metrics,
		start:          args.Start,
		end:            args.End,
		period:         args.Period,
	}
}

// List that contains the metrics that you want to retrieve for the associated data series. If you don&#39;t set this parameter, this resource returns data series for all Coll Stats Latency metrics.
func (r GetCollStatsLatencyNamespaceHostMeasurementsApiRequest) Metrics(metrics []string) GetCollStatsLatencyNamespaceHostMeasurementsApiRequest {
	r.metrics = &metrics
	return r
}

// Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetCollStatsLatencyNamespaceHostMeasurementsApiRequest) Start(start time.Time) GetCollStatsLatencyNamespaceHostMeasurementsApiRequest {
	r.start = &start
	return r
}

// Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetCollStatsLatencyNamespaceHostMeasurementsApiRequest) End(end time.Time) GetCollStatsLatencyNamespaceHostMeasurementsApiRequest {
	r.end = &end
	return r
}

// Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.
func (r GetCollStatsLatencyNamespaceHostMeasurementsApiRequest) Period(period string) GetCollStatsLatencyNamespaceHostMeasurementsApiRequest {
	r.period = &period
	return r
}

func (r GetCollStatsLatencyNamespaceHostMeasurementsApiRequest) Execute() (*MeasurementsCollStatsLatencyHost, *http.Response, error) {
	return r.ApiService.GetCollStatsLatencyNamespaceHostMeasurementsExecute(r)
}

/*
GetCollStatsLatencyNamespaceHostMeasurements Return Host-Level Query Latency

[experimental] Get a list of the Coll Stats Latency process-level measurements for the given namespace

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param processId Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
	@param databaseName Human-readable label that identifies the database.
	@param collectionName Human-readable label that identifies the collection.
	@return GetCollStatsLatencyNamespaceHostMeasurementsApiRequest
*/
func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespaceHostMeasurements(ctx context.Context, groupId string, processId string, databaseName string, collectionName string) GetCollStatsLatencyNamespaceHostMeasurementsApiRequest {
	return GetCollStatsLatencyNamespaceHostMeasurementsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		processId:      processId,
		databaseName:   databaseName,
		collectionName: collectionName,
	}
}

// Execute executes the request
//
//	@return MeasurementsCollStatsLatencyHost
func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespaceHostMeasurementsExecute(r GetCollStatsLatencyNamespaceHostMeasurementsApiRequest) (*MeasurementsCollStatsLatencyHost, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MeasurementsCollStatsLatencyHost
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionLevelMetricsApiService.GetCollStatsLatencyNamespaceHostMeasurements")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}/{databaseName}/{collectionName}/collStats/measurements"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(parameterValueToString(r.processId, "processId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(parameterValueToString(r.databaseName, "databaseName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"collectionName"+"}", url.PathEscape(parameterValueToString(r.collectionName, "collectionName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.metrics != nil {
		t := *r.metrics
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", t, "multi")

	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetCollStatsLatencyNamespaceMetricsApiRequest struct {
	ctx        context.Context
	ApiService CollectionLevelMetricsApi
	groupId    string
}

type GetCollStatsLatencyNamespaceMetricsApiParams struct {
	GroupId string
}

func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespaceMetricsWithParams(ctx context.Context, args *GetCollStatsLatencyNamespaceMetricsApiParams) GetCollStatsLatencyNamespaceMetricsApiRequest {
	return GetCollStatsLatencyNamespaceMetricsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetCollStatsLatencyNamespaceMetricsApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetCollStatsLatencyNamespaceMetricsExecute(r)
}

/*
GetCollStatsLatencyNamespaceMetrics Return all metric names

[experimental] Returns all available Coll Stats Latency metric names and their respective units for the specified project at the time of request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetCollStatsLatencyNamespaceMetricsApiRequest
*/
func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespaceMetrics(ctx context.Context, groupId string) GetCollStatsLatencyNamespaceMetricsApiRequest {
	return GetCollStatsLatencyNamespaceMetricsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespaceMetricsExecute(r GetCollStatsLatencyNamespaceMetricsApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionLevelMetricsApiService.GetCollStatsLatencyNamespaceMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/collStats/metrics"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetCollStatsLatencyNamespacesForClusterApiRequest struct {
	ctx         context.Context
	ApiService  CollectionLevelMetricsApi
	groupId     string
	clusterName string
	clusterView string
	start       *time.Time
	end         *time.Time
	period      *string
}

type GetCollStatsLatencyNamespacesForClusterApiParams struct {
	GroupId     string
	ClusterName string
	ClusterView string
	Start       *time.Time
	End         *time.Time
	Period      *string
}

func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespacesForClusterWithParams(ctx context.Context, args *GetCollStatsLatencyNamespacesForClusterApiParams) GetCollStatsLatencyNamespacesForClusterApiRequest {
	return GetCollStatsLatencyNamespacesForClusterApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		clusterView: args.ClusterView,
		start:       args.Start,
		end:         args.End,
		period:      args.Period,
	}
}

// Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetCollStatsLatencyNamespacesForClusterApiRequest) Start(start time.Time) GetCollStatsLatencyNamespacesForClusterApiRequest {
	r.start = &start
	return r
}

// Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetCollStatsLatencyNamespacesForClusterApiRequest) End(end time.Time) GetCollStatsLatencyNamespacesForClusterApiRequest {
	r.end = &end
	return r
}

// Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.
func (r GetCollStatsLatencyNamespacesForClusterApiRequest) Period(period string) GetCollStatsLatencyNamespacesForClusterApiRequest {
	r.period = &period
	return r
}

func (r GetCollStatsLatencyNamespacesForClusterApiRequest) Execute() (*CollStatsRankedNamespaces, *http.Response, error) {
	return r.ApiService.GetCollStatsLatencyNamespacesForClusterExecute(r)
}

/*
GetCollStatsLatencyNamespacesForCluster Return Ranked Namespaces from a Cluster

[experimental] Return the subset of namespaces from the given cluster sorted by highest total execution time (descending) within the given time window.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster to pin namespaces to.
	@param clusterView Human-readable label that identifies the cluster topology to retrieve metrics for.
	@return GetCollStatsLatencyNamespacesForClusterApiRequest
*/
func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespacesForCluster(ctx context.Context, groupId string, clusterName string, clusterView string) GetCollStatsLatencyNamespacesForClusterApiRequest {
	return GetCollStatsLatencyNamespacesForClusterApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		clusterView: clusterView,
	}
}

// Execute executes the request
//
//	@return CollStatsRankedNamespaces
func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespacesForClusterExecute(r GetCollStatsLatencyNamespacesForClusterApiRequest) (*CollStatsRankedNamespaces, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollStatsRankedNamespaces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionLevelMetricsApiService.GetCollStatsLatencyNamespacesForCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/{clusterView}/collStats/namespaces"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterView"+"}", url.PathEscape(parameterValueToString(r.clusterView, "clusterView")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetCollStatsLatencyNamespacesForHostApiRequest struct {
	ctx        context.Context
	ApiService CollectionLevelMetricsApi
	groupId    string
	processId  string
	start      *time.Time
	end        *time.Time
	period     *string
}

type GetCollStatsLatencyNamespacesForHostApiParams struct {
	GroupId   string
	ProcessId string
	Start     *time.Time
	End       *time.Time
	Period    *string
}

func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespacesForHostWithParams(ctx context.Context, args *GetCollStatsLatencyNamespacesForHostApiParams) GetCollStatsLatencyNamespacesForHostApiRequest {
	return GetCollStatsLatencyNamespacesForHostApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		processId:  args.ProcessId,
		start:      args.Start,
		end:        args.End,
		period:     args.Period,
	}
}

// Date and time when MongoDB Cloud begins reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetCollStatsLatencyNamespacesForHostApiRequest) Start(start time.Time) GetCollStatsLatencyNamespacesForHostApiRequest {
	r.start = &start
	return r
}

// Date and time when MongoDB Cloud stops reporting the metrics. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Include this parameter when you do not set **period**.
func (r GetCollStatsLatencyNamespacesForHostApiRequest) End(end time.Time) GetCollStatsLatencyNamespacesForHostApiRequest {
	r.end = &end
	return r
}

// Duration over which Atlas reports the metrics. This parameter expresses its value in the ISO 8601 duration format in UTC. Include this parameter when you do not set **start** and **end**.
func (r GetCollStatsLatencyNamespacesForHostApiRequest) Period(period string) GetCollStatsLatencyNamespacesForHostApiRequest {
	r.period = &period
	return r
}

func (r GetCollStatsLatencyNamespacesForHostApiRequest) Execute() (*CollStatsRankedNamespaces, *http.Response, error) {
	return r.ApiService.GetCollStatsLatencyNamespacesForHostExecute(r)
}

/*
GetCollStatsLatencyNamespacesForHost Return Ranked Namespaces from a Host

[experimental] Return the subset of namespaces from the given process ranked by highest total execution time (descending) within the given time window.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param processId Combination of hostname and IANA port that serves the MongoDB process. The host must be the hostname, fully qualified domain name (FQDN), or Internet Protocol address (IPv4 or IPv6) of the host that runs the MongoDB process (mongod or mongos). The port must be the IANA port on which the MongoDB process listens for requests.
	@return GetCollStatsLatencyNamespacesForHostApiRequest
*/
func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespacesForHost(ctx context.Context, groupId string, processId string) GetCollStatsLatencyNamespacesForHostApiRequest {
	return GetCollStatsLatencyNamespacesForHostApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		processId:  processId,
	}
}

// Execute executes the request
//
//	@return CollStatsRankedNamespaces
func (a *CollectionLevelMetricsApiService) GetCollStatsLatencyNamespacesForHostExecute(r GetCollStatsLatencyNamespacesForHostApiRequest) (*CollStatsRankedNamespaces, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollStatsRankedNamespaces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionLevelMetricsApiService.GetCollStatsLatencyNamespacesForHost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}/collStats/namespaces"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(parameterValueToString(r.processId, "processId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetPinnedNamespacesApiRequest struct {
	ctx         context.Context
	ApiService  CollectionLevelMetricsApi
	groupId     string
	clusterName string
}

type GetPinnedNamespacesApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *CollectionLevelMetricsApiService) GetPinnedNamespacesWithParams(ctx context.Context, args *GetPinnedNamespacesApiParams) GetPinnedNamespacesApiRequest {
	return GetPinnedNamespacesApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetPinnedNamespacesApiRequest) Execute() (*PinnedNamespaces, *http.Response, error) {
	return r.ApiService.GetPinnedNamespacesExecute(r)
}

/*
GetPinnedNamespaces Return Pinned Namespaces

[experimental] Returns a list of given cluster's pinned namespaces, a set of namespaces manually selected by users to collect query latency metrics on.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster to retrieve pinned namespaces for.
	@return GetPinnedNamespacesApiRequest
*/
func (a *CollectionLevelMetricsApiService) GetPinnedNamespaces(ctx context.Context, groupId string, clusterName string) GetPinnedNamespacesApiRequest {
	return GetPinnedNamespacesApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PinnedNamespaces
func (a *CollectionLevelMetricsApiService) GetPinnedNamespacesExecute(r GetPinnedNamespacesApiRequest) (*PinnedNamespaces, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PinnedNamespaces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionLevelMetricsApiService.GetPinnedNamespaces")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/pinned"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PinNamespacesPatchApiRequest struct {
	ctx               context.Context
	ApiService        CollectionLevelMetricsApi
	groupId           string
	clusterName       string
	namespacesRequest *NamespacesRequest
}

type PinNamespacesPatchApiParams struct {
	GroupId           string
	ClusterName       string
	NamespacesRequest *NamespacesRequest
}

func (a *CollectionLevelMetricsApiService) PinNamespacesPatchWithParams(ctx context.Context, args *PinNamespacesPatchApiParams) PinNamespacesPatchApiRequest {
	return PinNamespacesPatchApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		clusterName:       args.ClusterName,
		namespacesRequest: args.NamespacesRequest,
	}
}

func (r PinNamespacesPatchApiRequest) Execute() (*PinnedNamespaces, *http.Response, error) {
	return r.ApiService.PinNamespacesPatchExecute(r)
}

/*
PinNamespacesPatch Add Pinned Namespaces

[experimental] Add provided list of namespaces to existing pinned namespaces list for collection-level latency metrics collection for the given Group and Cluster

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster to pin namespaces to.
	@return PinNamespacesPatchApiRequest
*/
func (a *CollectionLevelMetricsApiService) PinNamespacesPatch(ctx context.Context, groupId string, clusterName string, namespacesRequest *NamespacesRequest) PinNamespacesPatchApiRequest {
	return PinNamespacesPatchApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		clusterName:       clusterName,
		namespacesRequest: namespacesRequest,
	}
}

// Execute executes the request
//
//	@return PinnedNamespaces
func (a *CollectionLevelMetricsApiService) PinNamespacesPatchExecute(r PinNamespacesPatchApiRequest) (*PinnedNamespaces, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PinnedNamespaces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionLevelMetricsApiService.PinNamespacesPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/pinned"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.namespacesRequest == nil {
		return localVarReturnValue, nil, reportError("namespacesRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-11-15+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.namespacesRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PinNamespacesPutApiRequest struct {
	ctx               context.Context
	ApiService        CollectionLevelMetricsApi
	groupId           string
	clusterName       string
	namespacesRequest *NamespacesRequest
}

type PinNamespacesPutApiParams struct {
	GroupId           string
	ClusterName       string
	NamespacesRequest *NamespacesRequest
}

func (a *CollectionLevelMetricsApiService) PinNamespacesPutWithParams(ctx context.Context, args *PinNamespacesPutApiParams) PinNamespacesPutApiRequest {
	return PinNamespacesPutApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		clusterName:       args.ClusterName,
		namespacesRequest: args.NamespacesRequest,
	}
}

func (r PinNamespacesPutApiRequest) Execute() (*PinnedNamespaces, *http.Response, error) {
	return r.ApiService.PinNamespacesPutExecute(r)
}

/*
PinNamespacesPut Pin Namespaces

[experimental] Pin provided list of namespaces for collection-level latency metrics collection for the given Group and Cluster. This initializes a pinned namespaces list or replaces any existing pinned namespaces list for the Group and Cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster to pin namespaces to.
	@return PinNamespacesPutApiRequest
*/
func (a *CollectionLevelMetricsApiService) PinNamespacesPut(ctx context.Context, groupId string, clusterName string, namespacesRequest *NamespacesRequest) PinNamespacesPutApiRequest {
	return PinNamespacesPutApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		clusterName:       clusterName,
		namespacesRequest: namespacesRequest,
	}
}

// Execute executes the request
//
//	@return PinnedNamespaces
func (a *CollectionLevelMetricsApiService) PinNamespacesPutExecute(r PinNamespacesPutApiRequest) (*PinnedNamespaces, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PinnedNamespaces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionLevelMetricsApiService.PinNamespacesPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/pinned"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.namespacesRequest == nil {
		return localVarReturnValue, nil, reportError("namespacesRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-11-15+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.namespacesRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UnpinNamespacesApiRequest struct {
	ctx               context.Context
	ApiService        CollectionLevelMetricsApi
	groupId           string
	clusterName       string
	namespacesRequest *NamespacesRequest
}

type UnpinNamespacesApiParams struct {
	GroupId           string
	ClusterName       string
	NamespacesRequest *NamespacesRequest
}

func (a *CollectionLevelMetricsApiService) UnpinNamespacesWithParams(ctx context.Context, args *UnpinNamespacesApiParams) UnpinNamespacesApiRequest {
	return UnpinNamespacesApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		clusterName:       args.ClusterName,
		namespacesRequest: args.NamespacesRequest,
	}
}

func (r UnpinNamespacesApiRequest) Execute() (*PinnedNamespaces, *http.Response, error) {
	return r.ApiService.UnpinNamespacesExecute(r)
}

/*
UnpinNamespaces Unpin namespaces

[experimental] Unpin provided list of namespaces for collection-level latency metrics collection for the given Group and Cluster

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster to unpin namespaces from.
	@return UnpinNamespacesApiRequest
*/
func (a *CollectionLevelMetricsApiService) UnpinNamespaces(ctx context.Context, groupId string, clusterName string, namespacesRequest *NamespacesRequest) UnpinNamespacesApiRequest {
	return UnpinNamespacesApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		clusterName:       clusterName,
		namespacesRequest: namespacesRequest,
	}
}

// Execute executes the request
//
//	@return PinnedNamespaces
func (a *CollectionLevelMetricsApiService) UnpinNamespacesExecute(r UnpinNamespacesApiRequest) (*PinnedNamespaces, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PinnedNamespaces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollectionLevelMetricsApiService.UnpinNamespaces")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/collStats/unpin"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.namespacesRequest == nil {
		return localVarReturnValue, nil, reportError("namespacesRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-11-15+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.namespacesRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
