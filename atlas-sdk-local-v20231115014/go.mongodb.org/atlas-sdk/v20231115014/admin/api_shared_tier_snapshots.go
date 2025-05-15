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

type SharedTierSnapshotsApi interface {

	/*
		DownloadSharedClusterBackup Download One M2 or M5 Cluster Snapshot

		[experimental] Requests one snapshot for the specified shared cluster. This resource returns a `snapshotURL` that you can use to download the snapshot. This `snapshotURL` remains active for four hours after you make the request. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clusterName Human-readable label that identifies the cluster.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DownloadSharedClusterBackupApiRequest
	*/
	DownloadSharedClusterBackup(ctx context.Context, clusterName string, groupId string, tenantRestore *TenantRestore) DownloadSharedClusterBackupApiRequest
	/*
		DownloadSharedClusterBackup Download One M2 or M5 Cluster Snapshot


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DownloadSharedClusterBackupApiParams - Parameters for the request
		@return DownloadSharedClusterBackupApiRequest
	*/
	DownloadSharedClusterBackupWithParams(ctx context.Context, args *DownloadSharedClusterBackupApiParams) DownloadSharedClusterBackupApiRequest

	// Method available only for mocking purposes
	DownloadSharedClusterBackupExecute(r DownloadSharedClusterBackupApiRequest) (*TenantRestore, *http.Response, error)

	/*
		GetSharedClusterBackup Return One Snapshot for One M2 or M5 Cluster

		[experimental] Returns details for one snapshot for the specified shared cluster. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
		@return GetSharedClusterBackupApiRequest
	*/
	GetSharedClusterBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) GetSharedClusterBackupApiRequest
	/*
		GetSharedClusterBackup Return One Snapshot for One M2 or M5 Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetSharedClusterBackupApiParams - Parameters for the request
		@return GetSharedClusterBackupApiRequest
	*/
	GetSharedClusterBackupWithParams(ctx context.Context, args *GetSharedClusterBackupApiParams) GetSharedClusterBackupApiRequest

	// Method available only for mocking purposes
	GetSharedClusterBackupExecute(r GetSharedClusterBackupApiRequest) (*BackupTenantSnapshot, *http.Response, error)

	/*
		ListSharedClusterBackups Return All Snapshots for One M2 or M5 Cluster

		[experimental] Returns details for all snapshots for the specified shared cluster. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return ListSharedClusterBackupsApiRequest
	*/
	ListSharedClusterBackups(ctx context.Context, groupId string, clusterName string) ListSharedClusterBackupsApiRequest
	/*
		ListSharedClusterBackups Return All Snapshots for One M2 or M5 Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListSharedClusterBackupsApiParams - Parameters for the request
		@return ListSharedClusterBackupsApiRequest
	*/
	ListSharedClusterBackupsWithParams(ctx context.Context, args *ListSharedClusterBackupsApiParams) ListSharedClusterBackupsApiRequest

	// Method available only for mocking purposes
	ListSharedClusterBackupsExecute(r ListSharedClusterBackupsApiRequest) (*PaginatedTenantSnapshot, *http.Response, error)
}

// SharedTierSnapshotsApiService SharedTierSnapshotsApi service
type SharedTierSnapshotsApiService service

type DownloadSharedClusterBackupApiRequest struct {
	ctx           context.Context
	ApiService    SharedTierSnapshotsApi
	clusterName   string
	groupId       string
	tenantRestore *TenantRestore
}

type DownloadSharedClusterBackupApiParams struct {
	ClusterName   string
	GroupId       string
	TenantRestore *TenantRestore
}

func (a *SharedTierSnapshotsApiService) DownloadSharedClusterBackupWithParams(ctx context.Context, args *DownloadSharedClusterBackupApiParams) DownloadSharedClusterBackupApiRequest {
	return DownloadSharedClusterBackupApiRequest{
		ApiService:    a,
		ctx:           ctx,
		clusterName:   args.ClusterName,
		groupId:       args.GroupId,
		tenantRestore: args.TenantRestore,
	}
}

func (r DownloadSharedClusterBackupApiRequest) Execute() (*TenantRestore, *http.Response, error) {
	return r.ApiService.DownloadSharedClusterBackupExecute(r)
}

/*
DownloadSharedClusterBackup Download One M2 or M5 Cluster Snapshot

[experimental] Requests one snapshot for the specified shared cluster. This resource returns a `snapshotURL` that you can use to download the snapshot. This `snapshotURL` remains active for four hours after you make the request. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clusterName Human-readable label that identifies the cluster.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DownloadSharedClusterBackupApiRequest
*/
func (a *SharedTierSnapshotsApiService) DownloadSharedClusterBackup(ctx context.Context, clusterName string, groupId string, tenantRestore *TenantRestore) DownloadSharedClusterBackupApiRequest {
	return DownloadSharedClusterBackupApiRequest{
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
func (a *SharedTierSnapshotsApiService) DownloadSharedClusterBackupExecute(r DownloadSharedClusterBackupApiRequest) (*TenantRestore, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TenantRestore
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierSnapshotsApiService.DownloadSharedClusterBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/download"
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

type GetSharedClusterBackupApiRequest struct {
	ctx         context.Context
	ApiService  SharedTierSnapshotsApi
	groupId     string
	clusterName string
	snapshotId  string
}

type GetSharedClusterBackupApiParams struct {
	GroupId     string
	ClusterName string
	SnapshotId  string
}

func (a *SharedTierSnapshotsApiService) GetSharedClusterBackupWithParams(ctx context.Context, args *GetSharedClusterBackupApiParams) GetSharedClusterBackupApiRequest {
	return GetSharedClusterBackupApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		snapshotId:  args.SnapshotId,
	}
}

func (r GetSharedClusterBackupApiRequest) Execute() (*BackupTenantSnapshot, *http.Response, error) {
	return r.ApiService.GetSharedClusterBackupExecute(r)
}

/*
GetSharedClusterBackup Return One Snapshot for One M2 or M5 Cluster

[experimental] Returns details for one snapshot for the specified shared cluster. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
	@return GetSharedClusterBackupApiRequest
*/
func (a *SharedTierSnapshotsApiService) GetSharedClusterBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) GetSharedClusterBackupApiRequest {
	return GetSharedClusterBackupApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		snapshotId:  snapshotId,
	}
}

// Execute executes the request
//
//	@return BackupTenantSnapshot
func (a *SharedTierSnapshotsApiService) GetSharedClusterBackupExecute(r GetSharedClusterBackupApiRequest) (*BackupTenantSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackupTenantSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierSnapshotsApiService.GetSharedClusterBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

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

type ListSharedClusterBackupsApiRequest struct {
	ctx         context.Context
	ApiService  SharedTierSnapshotsApi
	groupId     string
	clusterName string
}

type ListSharedClusterBackupsApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *SharedTierSnapshotsApiService) ListSharedClusterBackupsWithParams(ctx context.Context, args *ListSharedClusterBackupsApiParams) ListSharedClusterBackupsApiRequest {
	return ListSharedClusterBackupsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r ListSharedClusterBackupsApiRequest) Execute() (*PaginatedTenantSnapshot, *http.Response, error) {
	return r.ApiService.ListSharedClusterBackupsExecute(r)
}

/*
ListSharedClusterBackups Return All Snapshots for One M2 or M5 Cluster

[experimental] Returns details for all snapshots for the specified shared cluster. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return ListSharedClusterBackupsApiRequest
*/
func (a *SharedTierSnapshotsApiService) ListSharedClusterBackups(ctx context.Context, groupId string, clusterName string) ListSharedClusterBackupsApiRequest {
	return ListSharedClusterBackupsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PaginatedTenantSnapshot
func (a *SharedTierSnapshotsApiService) ListSharedClusterBackupsExecute(r ListSharedClusterBackupsApiRequest) (*PaginatedTenantSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedTenantSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierSnapshotsApiService.ListSharedClusterBackups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots"
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
