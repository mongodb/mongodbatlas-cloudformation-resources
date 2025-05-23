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
	"strings"
)

type SharedTierRestoreJobsApi interface {

	/*
		CreateSharedClusterBackupRestoreJob Create One Restore Job from One M2 or M5 Cluster

		[experimental] Restores the specified cluster. MongoDB Cloud limits which clusters can be the target clusters of a restore. The target cluster can't use encryption at rest, run a major release MongoDB version different than the snapshot, or receive client requests during restores. MongoDB Cloud deletes all existing data on the target cluster prior to the restore operation. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clusterName Human-readable label that identifies the cluster.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return CreateSharedClusterBackupRestoreJobApiRequest
	*/
	CreateSharedClusterBackupRestoreJob(ctx context.Context, clusterName string, groupId string, tenantRestore *TenantRestore) CreateSharedClusterBackupRestoreJobApiRequest
	/*
		CreateSharedClusterBackupRestoreJob Create One Restore Job from One M2 or M5 Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateSharedClusterBackupRestoreJobApiParams - Parameters for the request
		@return CreateSharedClusterBackupRestoreJobApiRequest
	*/
	CreateSharedClusterBackupRestoreJobWithParams(ctx context.Context, args *CreateSharedClusterBackupRestoreJobApiParams) CreateSharedClusterBackupRestoreJobApiRequest

	// Method available only for mocking purposes
	CreateSharedClusterBackupRestoreJobExecute(r CreateSharedClusterBackupRestoreJobApiRequest) (*TenantRestore, *http.Response, error)

	/*
		GetSharedClusterBackupRestoreJob Return One Restore Job for One M2 or M5 Cluster

		[experimental] Returns the specified restore job. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clusterName Human-readable label that identifies the cluster.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param restoreId Unique 24-hexadecimal digit string that identifies the restore job to return.
		@return GetSharedClusterBackupRestoreJobApiRequest
	*/
	GetSharedClusterBackupRestoreJob(ctx context.Context, clusterName string, groupId string, restoreId string) GetSharedClusterBackupRestoreJobApiRequest
	/*
		GetSharedClusterBackupRestoreJob Return One Restore Job for One M2 or M5 Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetSharedClusterBackupRestoreJobApiParams - Parameters for the request
		@return GetSharedClusterBackupRestoreJobApiRequest
	*/
	GetSharedClusterBackupRestoreJobWithParams(ctx context.Context, args *GetSharedClusterBackupRestoreJobApiParams) GetSharedClusterBackupRestoreJobApiRequest

	// Method available only for mocking purposes
	GetSharedClusterBackupRestoreJobExecute(r GetSharedClusterBackupRestoreJobApiRequest) (*TenantRestore, *http.Response, error)

	/*
		ListSharedClusterBackupRestoreJobs Return All Restore Jobs for One M2 or M5 Cluster

		[experimental] Returns all restore jobs for the specified M2 or M5 cluster. Restore jobs restore a cluster using a snapshot. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clusterName Human-readable label that identifies the cluster.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListSharedClusterBackupRestoreJobsApiRequest
	*/
	ListSharedClusterBackupRestoreJobs(ctx context.Context, clusterName string, groupId string) ListSharedClusterBackupRestoreJobsApiRequest
	/*
		ListSharedClusterBackupRestoreJobs Return All Restore Jobs for One M2 or M5 Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListSharedClusterBackupRestoreJobsApiParams - Parameters for the request
		@return ListSharedClusterBackupRestoreJobsApiRequest
	*/
	ListSharedClusterBackupRestoreJobsWithParams(ctx context.Context, args *ListSharedClusterBackupRestoreJobsApiParams) ListSharedClusterBackupRestoreJobsApiRequest

	// Method available only for mocking purposes
	ListSharedClusterBackupRestoreJobsExecute(r ListSharedClusterBackupRestoreJobsApiRequest) (*PaginatedTenantRestore, *http.Response, error)
}

// SharedTierRestoreJobsApiService SharedTierRestoreJobsApi service
type SharedTierRestoreJobsApiService service

type CreateSharedClusterBackupRestoreJobApiRequest struct {
	ctx           context.Context
	ApiService    SharedTierRestoreJobsApi
	clusterName   string
	groupId       string
	tenantRestore *TenantRestore
}

type CreateSharedClusterBackupRestoreJobApiParams struct {
	ClusterName   string
	GroupId       string
	TenantRestore *TenantRestore
}

func (a *SharedTierRestoreJobsApiService) CreateSharedClusterBackupRestoreJobWithParams(ctx context.Context, args *CreateSharedClusterBackupRestoreJobApiParams) CreateSharedClusterBackupRestoreJobApiRequest {
	return CreateSharedClusterBackupRestoreJobApiRequest{
		ApiService:    a,
		ctx:           ctx,
		clusterName:   args.ClusterName,
		groupId:       args.GroupId,
		tenantRestore: args.TenantRestore,
	}
}

func (r CreateSharedClusterBackupRestoreJobApiRequest) Execute() (*TenantRestore, *http.Response, error) {
	return r.ApiService.CreateSharedClusterBackupRestoreJobExecute(r)
}

/*
CreateSharedClusterBackupRestoreJob Create One Restore Job from One M2 or M5 Cluster

[experimental] Restores the specified cluster. MongoDB Cloud limits which clusters can be the target clusters of a restore. The target cluster can't use encryption at rest, run a major release MongoDB version different than the snapshot, or receive client requests during restores. MongoDB Cloud deletes all existing data on the target cluster prior to the restore operation. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clusterName Human-readable label that identifies the cluster.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateSharedClusterBackupRestoreJobApiRequest
*/
func (a *SharedTierRestoreJobsApiService) CreateSharedClusterBackupRestoreJob(ctx context.Context, clusterName string, groupId string, tenantRestore *TenantRestore) CreateSharedClusterBackupRestoreJobApiRequest {
	return CreateSharedClusterBackupRestoreJobApiRequest{
		ApiService:    a,
		ctx:           ctx,
		clusterName:   clusterName,
		groupId:       groupId,
		tenantRestore: tenantRestore,
	}
}

// Execute executes the request
//
//	@return TenantRestore
func (a *SharedTierRestoreJobsApiService) CreateSharedClusterBackupRestoreJobExecute(r CreateSharedClusterBackupRestoreJobApiRequest) (*TenantRestore, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TenantRestore
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierRestoreJobsApiService.CreateSharedClusterBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restore"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantRestore == nil {
		return localVarReturnValue, nil, reportError("tenantRestore is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

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
	// body params
	localVarPostBody = r.tenantRestore
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

type GetSharedClusterBackupRestoreJobApiRequest struct {
	ctx         context.Context
	ApiService  SharedTierRestoreJobsApi
	clusterName string
	groupId     string
	restoreId   string
}

type GetSharedClusterBackupRestoreJobApiParams struct {
	ClusterName string
	GroupId     string
	RestoreId   string
}

func (a *SharedTierRestoreJobsApiService) GetSharedClusterBackupRestoreJobWithParams(ctx context.Context, args *GetSharedClusterBackupRestoreJobApiParams) GetSharedClusterBackupRestoreJobApiRequest {
	return GetSharedClusterBackupRestoreJobApiRequest{
		ApiService:  a,
		ctx:         ctx,
		clusterName: args.ClusterName,
		groupId:     args.GroupId,
		restoreId:   args.RestoreId,
	}
}

func (r GetSharedClusterBackupRestoreJobApiRequest) Execute() (*TenantRestore, *http.Response, error) {
	return r.ApiService.GetSharedClusterBackupRestoreJobExecute(r)
}

/*
GetSharedClusterBackupRestoreJob Return One Restore Job for One M2 or M5 Cluster

[experimental] Returns the specified restore job. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clusterName Human-readable label that identifies the cluster.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param restoreId Unique 24-hexadecimal digit string that identifies the restore job to return.
	@return GetSharedClusterBackupRestoreJobApiRequest
*/
func (a *SharedTierRestoreJobsApiService) GetSharedClusterBackupRestoreJob(ctx context.Context, clusterName string, groupId string, restoreId string) GetSharedClusterBackupRestoreJobApiRequest {
	return GetSharedClusterBackupRestoreJobApiRequest{
		ApiService:  a,
		ctx:         ctx,
		clusterName: clusterName,
		groupId:     groupId,
		restoreId:   restoreId,
	}
}

// Execute executes the request
//
//	@return TenantRestore
func (a *SharedTierRestoreJobsApiService) GetSharedClusterBackupRestoreJobExecute(r GetSharedClusterBackupRestoreJobApiRequest) (*TenantRestore, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TenantRestore
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierRestoreJobsApiService.GetSharedClusterBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores/{restoreId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restoreId"+"}", url.PathEscape(parameterValueToString(r.restoreId, "restoreId")), -1)

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

type ListSharedClusterBackupRestoreJobsApiRequest struct {
	ctx         context.Context
	ApiService  SharedTierRestoreJobsApi
	clusterName string
	groupId     string
}

type ListSharedClusterBackupRestoreJobsApiParams struct {
	ClusterName string
	GroupId     string
}

func (a *SharedTierRestoreJobsApiService) ListSharedClusterBackupRestoreJobsWithParams(ctx context.Context, args *ListSharedClusterBackupRestoreJobsApiParams) ListSharedClusterBackupRestoreJobsApiRequest {
	return ListSharedClusterBackupRestoreJobsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		clusterName: args.ClusterName,
		groupId:     args.GroupId,
	}
}

func (r ListSharedClusterBackupRestoreJobsApiRequest) Execute() (*PaginatedTenantRestore, *http.Response, error) {
	return r.ApiService.ListSharedClusterBackupRestoreJobsExecute(r)
}

/*
ListSharedClusterBackupRestoreJobs Return All Restore Jobs for One M2 or M5 Cluster

[experimental] Returns all restore jobs for the specified M2 or M5 cluster. Restore jobs restore a cluster using a snapshot. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clusterName Human-readable label that identifies the cluster.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListSharedClusterBackupRestoreJobsApiRequest
*/
func (a *SharedTierRestoreJobsApiService) ListSharedClusterBackupRestoreJobs(ctx context.Context, clusterName string, groupId string) ListSharedClusterBackupRestoreJobsApiRequest {
	return ListSharedClusterBackupRestoreJobsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		clusterName: clusterName,
		groupId:     groupId,
	}
}

// Execute executes the request
//
//	@return PaginatedTenantRestore
func (a *SharedTierRestoreJobsApiService) ListSharedClusterBackupRestoreJobsExecute(r ListSharedClusterBackupRestoreJobsApiRequest) (*PaginatedTenantRestore, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedTenantRestore
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierRestoreJobsApiService.ListSharedClusterBackupRestoreJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
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
