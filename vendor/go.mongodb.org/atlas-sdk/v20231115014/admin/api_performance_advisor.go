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
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type PerformanceAdvisorApi interface {

	/*
		DisableSlowOperationThresholding Disable Managed Slow Operation Threshold

		Disables the slow operation threshold that MongoDB Cloud calculated for the specified project. The threshold determines which operations the Performance Advisor and Query Profiler considers slow. When enabled, MongoDB Cloud uses the average execution time for operations on your cluster to determine slow-running queries. As a result, the threshold is more pertinent to your cluster workload. The slow operation threshold is enabled by default for dedicated clusters (M10+). When disabled, MongoDB Cloud considers any operation that takes longer than 100 milliseconds to be slow. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DisableSlowOperationThresholdingApiRequest
	*/
	DisableSlowOperationThresholding(ctx context.Context, groupId string) DisableSlowOperationThresholdingApiRequest
	/*
		DisableSlowOperationThresholding Disable Managed Slow Operation Threshold


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DisableSlowOperationThresholdingApiParams - Parameters for the request
		@return DisableSlowOperationThresholdingApiRequest
	*/
	DisableSlowOperationThresholdingWithParams(ctx context.Context, args *DisableSlowOperationThresholdingApiParams) DisableSlowOperationThresholdingApiRequest

	// Method available only for mocking purposes
	DisableSlowOperationThresholdingExecute(r DisableSlowOperationThresholdingApiRequest) (*http.Response, error)

	/*
		EnableSlowOperationThresholding Enable Managed Slow Operation Threshold

		Enables MongoDB Cloud to use its slow operation threshold for the specified project. The threshold determines which operations the Performance Advisor and Query Profiler considers slow. When enabled, MongoDB Cloud uses the average execution time for operations on your cluster to determine slow-running queries. As a result, the threshold is more pertinent to your cluster workload. The slow operation threshold is enabled by default for dedicated clusters (M10+). When disabled, MongoDB Cloud considers any operation that takes longer than 100 milliseconds to be slow. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return EnableSlowOperationThresholdingApiRequest
	*/
	EnableSlowOperationThresholding(ctx context.Context, groupId string) EnableSlowOperationThresholdingApiRequest
	/*
		EnableSlowOperationThresholding Enable Managed Slow Operation Threshold


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param EnableSlowOperationThresholdingApiParams - Parameters for the request
		@return EnableSlowOperationThresholdingApiRequest
	*/
	EnableSlowOperationThresholdingWithParams(ctx context.Context, args *EnableSlowOperationThresholdingApiParams) EnableSlowOperationThresholdingApiRequest

	// Method available only for mocking purposes
	EnableSlowOperationThresholdingExecute(r EnableSlowOperationThresholdingApiRequest) (*http.Response, error)

	/*
		GetServerlessAutoIndexing Return Serverless Auto Indexing Enabled

		[experimental] Get whether the Serverless Auto Indexing feature is enabled.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return GetServerlessAutoIndexingApiRequest
	*/
	GetServerlessAutoIndexing(ctx context.Context, groupId string, clusterName string) GetServerlessAutoIndexingApiRequest
	/*
		GetServerlessAutoIndexing Return Serverless Auto Indexing Enabled


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetServerlessAutoIndexingApiParams - Parameters for the request
		@return GetServerlessAutoIndexingApiRequest
	*/
	GetServerlessAutoIndexingWithParams(ctx context.Context, args *GetServerlessAutoIndexingApiParams) GetServerlessAutoIndexingApiRequest

	// Method available only for mocking purposes
	GetServerlessAutoIndexingExecute(r GetServerlessAutoIndexingApiRequest) (bool, *http.Response, error)

	/*
		ListSlowQueries Return Slow Queries

		Returns log lines for slow queries that the Performance Advisor and Query Profiler identified. The Performance Advisor monitors queries that MongoDB considers slow and suggests new indexes to improve query performance. MongoDB Cloud bases the threshold for slow queries on the average time of operations on your cluster. This enables workload-relevant recommendations. To use this resource, the requesting API Key must have the Project Data Access Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param processId Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
		@return ListSlowQueriesApiRequest
	*/
	ListSlowQueries(ctx context.Context, groupId string, processId string) ListSlowQueriesApiRequest
	/*
		ListSlowQueries Return Slow Queries


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListSlowQueriesApiParams - Parameters for the request
		@return ListSlowQueriesApiRequest
	*/
	ListSlowQueriesWithParams(ctx context.Context, args *ListSlowQueriesApiParams) ListSlowQueriesApiRequest

	// Method available only for mocking purposes
	ListSlowQueriesExecute(r ListSlowQueriesApiRequest) (*PerformanceAdvisorSlowQueryList, *http.Response, error)

	/*
		ListSlowQueryNamespaces Return All Namespaces for One Host

		Returns up to 20 namespaces for collections experiencing slow queries on the specified host. If you specify a secondary member of a replica set that hasn't received any database read operations, the endpoint doesn't return any namespaces. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param processId Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
		@return ListSlowQueryNamespacesApiRequest
	*/
	ListSlowQueryNamespaces(ctx context.Context, groupId string, processId string) ListSlowQueryNamespacesApiRequest
	/*
		ListSlowQueryNamespaces Return All Namespaces for One Host


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListSlowQueryNamespacesApiParams - Parameters for the request
		@return ListSlowQueryNamespacesApiRequest
	*/
	ListSlowQueryNamespacesWithParams(ctx context.Context, args *ListSlowQueryNamespacesApiParams) ListSlowQueryNamespacesApiRequest

	// Method available only for mocking purposes
	ListSlowQueryNamespacesExecute(r ListSlowQueryNamespacesApiRequest) (*Namespaces, *http.Response, error)

	/*
		ListSuggestedIndexes Return Suggested Indexes

		Returns the indexes that the Performance Advisor suggests. The Performance Advisor monitors queries that MongoDB considers slow and suggests new indexes to improve query performance. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param processId Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
		@return ListSuggestedIndexesApiRequest
	*/
	ListSuggestedIndexes(ctx context.Context, groupId string, processId string) ListSuggestedIndexesApiRequest
	/*
		ListSuggestedIndexes Return Suggested Indexes


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListSuggestedIndexesApiParams - Parameters for the request
		@return ListSuggestedIndexesApiRequest
	*/
	ListSuggestedIndexesWithParams(ctx context.Context, args *ListSuggestedIndexesApiParams) ListSuggestedIndexesApiRequest

	// Method available only for mocking purposes
	ListSuggestedIndexesExecute(r ListSuggestedIndexesApiRequest) (*PerformanceAdvisorResponse, *http.Response, error)

	/*
		SetServerlessAutoIndexing Set Serverless Auto Indexing

		[experimental] Set whether the Serverless Auto Indexing feature is enabled.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return SetServerlessAutoIndexingApiRequest
	*/
	SetServerlessAutoIndexing(ctx context.Context, groupId string, clusterName string) SetServerlessAutoIndexingApiRequest
	/*
		SetServerlessAutoIndexing Set Serverless Auto Indexing


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param SetServerlessAutoIndexingApiParams - Parameters for the request
		@return SetServerlessAutoIndexingApiRequest
	*/
	SetServerlessAutoIndexingWithParams(ctx context.Context, args *SetServerlessAutoIndexingApiParams) SetServerlessAutoIndexingApiRequest

	// Method available only for mocking purposes
	SetServerlessAutoIndexingExecute(r SetServerlessAutoIndexingApiRequest) (map[string]interface{}, *http.Response, error)
}

// PerformanceAdvisorApiService PerformanceAdvisorApi service
type PerformanceAdvisorApiService service

type DisableSlowOperationThresholdingApiRequest struct {
	ctx        context.Context
	ApiService PerformanceAdvisorApi
	groupId    string
}

type DisableSlowOperationThresholdingApiParams struct {
	GroupId string
}

func (a *PerformanceAdvisorApiService) DisableSlowOperationThresholdingWithParams(ctx context.Context, args *DisableSlowOperationThresholdingApiParams) DisableSlowOperationThresholdingApiRequest {
	return DisableSlowOperationThresholdingApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r DisableSlowOperationThresholdingApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DisableSlowOperationThresholdingExecute(r)
}

/*
DisableSlowOperationThresholding Disable Managed Slow Operation Threshold

Disables the slow operation threshold that MongoDB Cloud calculated for the specified project. The threshold determines which operations the Performance Advisor and Query Profiler considers slow. When enabled, MongoDB Cloud uses the average execution time for operations on your cluster to determine slow-running queries. As a result, the threshold is more pertinent to your cluster workload. The slow operation threshold is enabled by default for dedicated clusters (M10+). When disabled, MongoDB Cloud considers any operation that takes longer than 100 milliseconds to be slow. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DisableSlowOperationThresholdingApiRequest
*/
func (a *PerformanceAdvisorApiService) DisableSlowOperationThresholding(ctx context.Context, groupId string) DisableSlowOperationThresholdingApiRequest {
	return DisableSlowOperationThresholdingApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
func (a *PerformanceAdvisorApiService) DisableSlowOperationThresholdingExecute(r DisableSlowOperationThresholdingApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PerformanceAdvisorApiService.DisableSlowOperationThresholding")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/managedSlowMs/disable"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type EnableSlowOperationThresholdingApiRequest struct {
	ctx        context.Context
	ApiService PerformanceAdvisorApi
	groupId    string
}

type EnableSlowOperationThresholdingApiParams struct {
	GroupId string
}

func (a *PerformanceAdvisorApiService) EnableSlowOperationThresholdingWithParams(ctx context.Context, args *EnableSlowOperationThresholdingApiParams) EnableSlowOperationThresholdingApiRequest {
	return EnableSlowOperationThresholdingApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r EnableSlowOperationThresholdingApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnableSlowOperationThresholdingExecute(r)
}

/*
EnableSlowOperationThresholding Enable Managed Slow Operation Threshold

Enables MongoDB Cloud to use its slow operation threshold for the specified project. The threshold determines which operations the Performance Advisor and Query Profiler considers slow. When enabled, MongoDB Cloud uses the average execution time for operations on your cluster to determine slow-running queries. As a result, the threshold is more pertinent to your cluster workload. The slow operation threshold is enabled by default for dedicated clusters (M10+). When disabled, MongoDB Cloud considers any operation that takes longer than 100 milliseconds to be slow. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return EnableSlowOperationThresholdingApiRequest
*/
func (a *PerformanceAdvisorApiService) EnableSlowOperationThresholding(ctx context.Context, groupId string) EnableSlowOperationThresholdingApiRequest {
	return EnableSlowOperationThresholdingApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
func (a *PerformanceAdvisorApiService) EnableSlowOperationThresholdingExecute(r EnableSlowOperationThresholdingApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PerformanceAdvisorApiService.EnableSlowOperationThresholding")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/managedSlowMs/enable"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type GetServerlessAutoIndexingApiRequest struct {
	ctx         context.Context
	ApiService  PerformanceAdvisorApi
	groupId     string
	clusterName string
}

type GetServerlessAutoIndexingApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *PerformanceAdvisorApiService) GetServerlessAutoIndexingWithParams(ctx context.Context, args *GetServerlessAutoIndexingApiParams) GetServerlessAutoIndexingApiRequest {
	return GetServerlessAutoIndexingApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetServerlessAutoIndexingApiRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.GetServerlessAutoIndexingExecute(r)
}

/*
GetServerlessAutoIndexing Return Serverless Auto Indexing Enabled

[experimental] Get whether the Serverless Auto Indexing feature is enabled.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return GetServerlessAutoIndexingApiRequest
*/
func (a *PerformanceAdvisorApiService) GetServerlessAutoIndexing(ctx context.Context, groupId string, clusterName string) GetServerlessAutoIndexingApiRequest {
	return GetServerlessAutoIndexingApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return bool
func (a *PerformanceAdvisorApiService) GetServerlessAutoIndexingExecute(r GetServerlessAutoIndexingApiRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PerformanceAdvisorApiService.GetServerlessAutoIndexing")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{clusterName}/performanceAdvisor/autoIndexing"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

type ListSlowQueriesApiRequest struct {
	ctx        context.Context
	ApiService PerformanceAdvisorApi
	groupId    string
	processId  string
	duration   *int64
	namespaces *[]string
	nLogs      *int64
	since      *int64
}

type ListSlowQueriesApiParams struct {
	GroupId    string
	ProcessId  string
	Duration   *int64
	Namespaces *[]string
	NLogs      *int64
	Since      *int64
}

func (a *PerformanceAdvisorApiService) ListSlowQueriesWithParams(ctx context.Context, args *ListSlowQueriesApiParams) ListSlowQueriesApiRequest {
	return ListSlowQueriesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		processId:  args.ProcessId,
		duration:   args.Duration,
		namespaces: args.Namespaces,
		nLogs:      args.NLogs,
		since:      args.Since,
	}
}

// Length of time expressed during which the query finds slow queries among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours.
func (r ListSlowQueriesApiRequest) Duration(duration int64) ListSlowQueriesApiRequest {
	r.duration = &duration
	return r
}

// Namespaces from which to retrieve slow queries. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces.
func (r ListSlowQueriesApiRequest) Namespaces(namespaces []string) ListSlowQueriesApiRequest {
	r.namespaces = &namespaces
	return r
}

// Maximum number of lines from the log to return.
func (r ListSlowQueriesApiRequest) NLogs(nLogs int64) ListSlowQueriesApiRequest {
	r.nLogs = &nLogs
	return r
}

// Date and time from which the query retrieves the slow queries. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours.
func (r ListSlowQueriesApiRequest) Since(since int64) ListSlowQueriesApiRequest {
	r.since = &since
	return r
}

func (r ListSlowQueriesApiRequest) Execute() (*PerformanceAdvisorSlowQueryList, *http.Response, error) {
	return r.ApiService.ListSlowQueriesExecute(r)
}

/*
ListSlowQueries Return Slow Queries

Returns log lines for slow queries that the Performance Advisor and Query Profiler identified. The Performance Advisor monitors queries that MongoDB considers slow and suggests new indexes to improve query performance. MongoDB Cloud bases the threshold for slow queries on the average time of operations on your cluster. This enables workload-relevant recommendations. To use this resource, the requesting API Key must have the Project Data Access Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param processId Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	@return ListSlowQueriesApiRequest
*/
func (a *PerformanceAdvisorApiService) ListSlowQueries(ctx context.Context, groupId string, processId string) ListSlowQueriesApiRequest {
	return ListSlowQueriesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		processId:  processId,
	}
}

// Execute executes the request
//
//	@return PerformanceAdvisorSlowQueryList
func (a *PerformanceAdvisorApiService) ListSlowQueriesExecute(r ListSlowQueriesApiRequest) (*PerformanceAdvisorSlowQueryList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PerformanceAdvisorSlowQueryList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PerformanceAdvisorApiService.ListSlowQueries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/slowQueryLogs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(parameterValueToString(r.processId, "processId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.duration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "duration", r.duration, "")
	}
	if r.namespaces != nil {
		t := *r.namespaces
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespaces", t, "multi")

	}
	if r.nLogs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nLogs", r.nLogs, "")
	} else {
		var defaultValue int64 = 20000
		r.nLogs = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "nLogs", r.nLogs, "")
	}
	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

type ListSlowQueryNamespacesApiRequest struct {
	ctx        context.Context
	ApiService PerformanceAdvisorApi
	groupId    string
	processId  string
	duration   *int64
	since      *int64
}

type ListSlowQueryNamespacesApiParams struct {
	GroupId   string
	ProcessId string
	Duration  *int64
	Since     *int64
}

func (a *PerformanceAdvisorApiService) ListSlowQueryNamespacesWithParams(ctx context.Context, args *ListSlowQueryNamespacesApiParams) ListSlowQueryNamespacesApiRequest {
	return ListSlowQueryNamespacesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		processId:  args.ProcessId,
		duration:   args.Duration,
		since:      args.Since,
	}
}

// Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours.
func (r ListSlowQueryNamespacesApiRequest) Duration(duration int64) ListSlowQueryNamespacesApiRequest {
	r.duration = &duration
	return r
}

// Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours.
func (r ListSlowQueryNamespacesApiRequest) Since(since int64) ListSlowQueryNamespacesApiRequest {
	r.since = &since
	return r
}

func (r ListSlowQueryNamespacesApiRequest) Execute() (*Namespaces, *http.Response, error) {
	return r.ApiService.ListSlowQueryNamespacesExecute(r)
}

/*
ListSlowQueryNamespaces Return All Namespaces for One Host

Returns up to 20 namespaces for collections experiencing slow queries on the specified host. If you specify a secondary member of a replica set that hasn't received any database read operations, the endpoint doesn't return any namespaces. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param processId Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	@return ListSlowQueryNamespacesApiRequest
*/
func (a *PerformanceAdvisorApiService) ListSlowQueryNamespaces(ctx context.Context, groupId string, processId string) ListSlowQueryNamespacesApiRequest {
	return ListSlowQueryNamespacesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		processId:  processId,
	}
}

// Execute executes the request
//
//	@return Namespaces
func (a *PerformanceAdvisorApiService) ListSlowQueryNamespacesExecute(r ListSlowQueryNamespacesApiRequest) (*Namespaces, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Namespaces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PerformanceAdvisorApiService.ListSlowQueryNamespaces")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/namespaces"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(parameterValueToString(r.processId, "processId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.duration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "duration", r.duration, "")
	}
	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

type ListSuggestedIndexesApiRequest struct {
	ctx          context.Context
	ApiService   PerformanceAdvisorApi
	groupId      string
	processId    string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
	duration     *int64
	namespaces   *[]string
	nExamples    *int64
	nIndexes     *int64
	since        *int64
}

type ListSuggestedIndexesApiParams struct {
	GroupId      string
	ProcessId    string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
	Duration     *int64
	Namespaces   *[]string
	NExamples    *int64
	NIndexes     *int64
	Since        *int64
}

func (a *PerformanceAdvisorApiService) ListSuggestedIndexesWithParams(ctx context.Context, args *ListSuggestedIndexesApiParams) ListSuggestedIndexesApiRequest {
	return ListSuggestedIndexesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		processId:    args.ProcessId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		duration:     args.Duration,
		namespaces:   args.Namespaces,
		nExamples:    args.NExamples,
		nIndexes:     args.NIndexes,
		since:        args.Since,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListSuggestedIndexesApiRequest) IncludeCount(includeCount bool) ListSuggestedIndexesApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListSuggestedIndexesApiRequest) ItemsPerPage(itemsPerPage int) ListSuggestedIndexesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListSuggestedIndexesApiRequest) PageNum(pageNum int) ListSuggestedIndexesApiRequest {
	r.pageNum = &pageNum
	return r
}

// Length of time expressed during which the query finds suggested indexes among the managed namespaces in the cluster. This parameter expresses its value in milliseconds.  - If you don&#39;t specify the **since** parameter, the endpoint returns data covering the duration before the current time. - If you specify neither the **duration** nor **since** parameters, the endpoint returns data from the previous 24 hours.
func (r ListSuggestedIndexesApiRequest) Duration(duration int64) ListSuggestedIndexesApiRequest {
	r.duration = &duration
	return r
}

// Namespaces from which to retrieve suggested indexes. A namespace consists of one database and one collection resource written as &#x60;.&#x60;: &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. To include multiple namespaces, pass the parameter multiple times delimited with an ampersand (&#x60;&amp;&#x60;) between each namespace. Omit this parameter to return results for all namespaces.
func (r ListSuggestedIndexesApiRequest) Namespaces(namespaces []string) ListSuggestedIndexesApiRequest {
	r.namespaces = &namespaces
	return r
}

// Maximum number of example queries that benefit from the suggested index.
func (r ListSuggestedIndexesApiRequest) NExamples(nExamples int64) ListSuggestedIndexesApiRequest {
	r.nExamples = &nExamples
	return r
}

// Number that indicates the maximum indexes to suggest.
func (r ListSuggestedIndexesApiRequest) NIndexes(nIndexes int64) ListSuggestedIndexesApiRequest {
	r.nIndexes = &nIndexes
	return r
}

// Date and time from which the query retrieves the suggested indexes. This parameter expresses its value in the number of milliseconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).  - If you don&#39;t specify the **duration** parameter, the endpoint returns data covering from the **since** value and the current time. - If you specify neither the **duration** nor the **since** parameters, the endpoint returns data from the previous 24 hours.
func (r ListSuggestedIndexesApiRequest) Since(since int64) ListSuggestedIndexesApiRequest {
	r.since = &since
	return r
}

func (r ListSuggestedIndexesApiRequest) Execute() (*PerformanceAdvisorResponse, *http.Response, error) {
	return r.ApiService.ListSuggestedIndexesExecute(r)
}

/*
ListSuggestedIndexes Return Suggested Indexes

Returns the indexes that the Performance Advisor suggests. The Performance Advisor monitors queries that MongoDB considers slow and suggests new indexes to improve query performance. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param processId Combination of host and port that serves the MongoDB process. The host must be the hostname, FQDN, IPv4 address, or IPv6 address of the host that runs the MongoDB process (`mongod` or `mongos`). The port must be the IANA port on which the MongoDB process listens for requests.
	@return ListSuggestedIndexesApiRequest
*/
func (a *PerformanceAdvisorApiService) ListSuggestedIndexes(ctx context.Context, groupId string, processId string) ListSuggestedIndexesApiRequest {
	return ListSuggestedIndexesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		processId:  processId,
	}
}

// Execute executes the request
//
//	@return PerformanceAdvisorResponse
func (a *PerformanceAdvisorApiService) ListSuggestedIndexesExecute(r ListSuggestedIndexesApiRequest) (*PerformanceAdvisorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PerformanceAdvisorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PerformanceAdvisorApiService.ListSuggestedIndexes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/processes/{processId}/performanceAdvisor/suggestedIndexes"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"processId"+"}", url.PathEscape(parameterValueToString(r.processId, "processId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	if r.duration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "duration", r.duration, "")
	}
	if r.namespaces != nil {
		t := *r.namespaces
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespaces", t, "multi")

	}
	if r.nExamples != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nExamples", r.nExamples, "")
	} else {
		var defaultValue int64 = 5
		r.nExamples = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "nExamples", r.nExamples, "")
	}
	if r.nIndexes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nIndexes", r.nIndexes, "")
	}
	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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

type SetServerlessAutoIndexingApiRequest struct {
	ctx         context.Context
	ApiService  PerformanceAdvisorApi
	groupId     string
	clusterName string
	enable      *bool
}

type SetServerlessAutoIndexingApiParams struct {
	GroupId     string
	ClusterName string
	Enable      *bool
}

func (a *PerformanceAdvisorApiService) SetServerlessAutoIndexingWithParams(ctx context.Context, args *SetServerlessAutoIndexingApiParams) SetServerlessAutoIndexingApiRequest {
	return SetServerlessAutoIndexingApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		enable:      args.Enable,
	}
}

// Value that we want to set for the Serverless Auto Indexing toggle.
func (r SetServerlessAutoIndexingApiRequest) Enable(enable bool) SetServerlessAutoIndexingApiRequest {
	r.enable = &enable
	return r
}

func (r SetServerlessAutoIndexingApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.SetServerlessAutoIndexingExecute(r)
}

/*
SetServerlessAutoIndexing Set Serverless Auto Indexing

[experimental] Set whether the Serverless Auto Indexing feature is enabled.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return SetServerlessAutoIndexingApiRequest
*/
func (a *PerformanceAdvisorApiService) SetServerlessAutoIndexing(ctx context.Context, groupId string, clusterName string) SetServerlessAutoIndexingApiRequest {
	return SetServerlessAutoIndexingApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *PerformanceAdvisorApiService) SetServerlessAutoIndexingExecute(r SetServerlessAutoIndexingApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PerformanceAdvisorApiService.SetServerlessAutoIndexing")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{clusterName}/performanceAdvisor/autoIndexing"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.enable == nil {
		return localVarReturnValue, nil, reportError("enable is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "enable", r.enable, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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
