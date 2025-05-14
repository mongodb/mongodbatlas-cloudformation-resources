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

type OnlineArchiveApi interface {

	/*
		CreateOnlineArchive Create One Online Archive

		Creates one online archive. This archive stores data from one cluster within one project. To use this resource, the requesting API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster that contains the collection for which you want to create one online archive.
		@return CreateOnlineArchiveApiRequest
	*/
	CreateOnlineArchive(ctx context.Context, groupId string, clusterName string, backupOnlineArchiveCreate *BackupOnlineArchiveCreate) CreateOnlineArchiveApiRequest
	/*
		CreateOnlineArchive Create One Online Archive


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateOnlineArchiveApiParams - Parameters for the request
		@return CreateOnlineArchiveApiRequest
	*/
	CreateOnlineArchiveWithParams(ctx context.Context, args *CreateOnlineArchiveApiParams) CreateOnlineArchiveApiRequest

	// Method available only for mocking purposes
	CreateOnlineArchiveExecute(r CreateOnlineArchiveApiRequest) (*BackupOnlineArchive, *http.Response, error)

	/*
		DeleteOnlineArchive Remove One Online Archive

		Removes one online archive. This archive stores data from one cluster within one project. To use this resource, the requesting API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param archiveId Unique 24-hexadecimal digit string that identifies the online archive to delete.
		@param clusterName Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive.
		@return DeleteOnlineArchiveApiRequest
	*/
	DeleteOnlineArchive(ctx context.Context, groupId string, archiveId string, clusterName string) DeleteOnlineArchiveApiRequest
	/*
		DeleteOnlineArchive Remove One Online Archive


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteOnlineArchiveApiParams - Parameters for the request
		@return DeleteOnlineArchiveApiRequest
	*/
	DeleteOnlineArchiveWithParams(ctx context.Context, args *DeleteOnlineArchiveApiParams) DeleteOnlineArchiveApiRequest

	// Method available only for mocking purposes
	DeleteOnlineArchiveExecute(r DeleteOnlineArchiveApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		DownloadOnlineArchiveQueryLogs Download Online Archive Query Logs

		[experimental] Downloads query logs for the specified online archive. To use this resource, the requesting API Key must have the Project Data Access Read Only or higher role. The API does not support direct calls with the json response schema. You must request a gzip response schema using an accept header of the format: "Accept: application/vnd.atlas.YYYY-MM-DD+gzip".

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster that contains the collection for which you want to return the query logs from one online archive.
		@return DownloadOnlineArchiveQueryLogsApiRequest
	*/
	DownloadOnlineArchiveQueryLogs(ctx context.Context, groupId string, clusterName string) DownloadOnlineArchiveQueryLogsApiRequest
	/*
		DownloadOnlineArchiveQueryLogs Download Online Archive Query Logs


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DownloadOnlineArchiveQueryLogsApiParams - Parameters for the request
		@return DownloadOnlineArchiveQueryLogsApiRequest
	*/
	DownloadOnlineArchiveQueryLogsWithParams(ctx context.Context, args *DownloadOnlineArchiveQueryLogsApiParams) DownloadOnlineArchiveQueryLogsApiRequest

	// Method available only for mocking purposes
	DownloadOnlineArchiveQueryLogsExecute(r DownloadOnlineArchiveQueryLogsApiRequest) (io.ReadCloser, *http.Response, error)

	/*
		GetOnlineArchive Return One Online Archive

		Returns one online archive for one cluster. This archive stores data from one cluster within one project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param archiveId Unique 24-hexadecimal digit string that identifies the online archive to return.
		@param clusterName Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.
		@return GetOnlineArchiveApiRequest
	*/
	GetOnlineArchive(ctx context.Context, groupId string, archiveId string, clusterName string) GetOnlineArchiveApiRequest
	/*
		GetOnlineArchive Return One Online Archive


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetOnlineArchiveApiParams - Parameters for the request
		@return GetOnlineArchiveApiRequest
	*/
	GetOnlineArchiveWithParams(ctx context.Context, args *GetOnlineArchiveApiParams) GetOnlineArchiveApiRequest

	// Method available only for mocking purposes
	GetOnlineArchiveExecute(r GetOnlineArchiveApiRequest) (*BackupOnlineArchive, *http.Response, error)

	/*
		ListOnlineArchives Return All Online Archives for One Cluster

		Returns details of all online archives. This archive stores data from one cluster within one project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster that contains the collection for which you want to return the online archives.
		@return ListOnlineArchivesApiRequest
	*/
	ListOnlineArchives(ctx context.Context, groupId string, clusterName string) ListOnlineArchivesApiRequest
	/*
		ListOnlineArchives Return All Online Archives for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListOnlineArchivesApiParams - Parameters for the request
		@return ListOnlineArchivesApiRequest
	*/
	ListOnlineArchivesWithParams(ctx context.Context, args *ListOnlineArchivesApiParams) ListOnlineArchivesApiRequest

	// Method available only for mocking purposes
	ListOnlineArchivesExecute(r ListOnlineArchivesApiRequest) (*PaginatedOnlineArchive, *http.Response, error)

	/*
		UpdateOnlineArchive Update One Online Archive

		Updates, pauses, or resumes one online archive. This archive stores data from one cluster within one project. To use this resource, the requesting API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param archiveId Unique 24-hexadecimal digit string that identifies the online archive to update.
		@param clusterName Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.
		@return UpdateOnlineArchiveApiRequest
	*/
	UpdateOnlineArchive(ctx context.Context, groupId string, archiveId string, clusterName string, backupOnlineArchive *BackupOnlineArchive) UpdateOnlineArchiveApiRequest
	/*
		UpdateOnlineArchive Update One Online Archive


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateOnlineArchiveApiParams - Parameters for the request
		@return UpdateOnlineArchiveApiRequest
	*/
	UpdateOnlineArchiveWithParams(ctx context.Context, args *UpdateOnlineArchiveApiParams) UpdateOnlineArchiveApiRequest

	// Method available only for mocking purposes
	UpdateOnlineArchiveExecute(r UpdateOnlineArchiveApiRequest) (*BackupOnlineArchive, *http.Response, error)
}

// OnlineArchiveApiService OnlineArchiveApi service
type OnlineArchiveApiService service

type CreateOnlineArchiveApiRequest struct {
	ctx                       context.Context
	ApiService                OnlineArchiveApi
	groupId                   string
	clusterName               string
	backupOnlineArchiveCreate *BackupOnlineArchiveCreate
}

type CreateOnlineArchiveApiParams struct {
	GroupId                   string
	ClusterName               string
	BackupOnlineArchiveCreate *BackupOnlineArchiveCreate
}

func (a *OnlineArchiveApiService) CreateOnlineArchiveWithParams(ctx context.Context, args *CreateOnlineArchiveApiParams) CreateOnlineArchiveApiRequest {
	return CreateOnlineArchiveApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   args.GroupId,
		clusterName:               args.ClusterName,
		backupOnlineArchiveCreate: args.BackupOnlineArchiveCreate,
	}
}

func (r CreateOnlineArchiveApiRequest) Execute() (*BackupOnlineArchive, *http.Response, error) {
	return r.ApiService.CreateOnlineArchiveExecute(r)
}

/*
CreateOnlineArchive Create One Online Archive

Creates one online archive. This archive stores data from one cluster within one project. To use this resource, the requesting API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster that contains the collection for which you want to create one online archive.
	@return CreateOnlineArchiveApiRequest
*/
func (a *OnlineArchiveApiService) CreateOnlineArchive(ctx context.Context, groupId string, clusterName string, backupOnlineArchiveCreate *BackupOnlineArchiveCreate) CreateOnlineArchiveApiRequest {
	return CreateOnlineArchiveApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   groupId,
		clusterName:               clusterName,
		backupOnlineArchiveCreate: backupOnlineArchiveCreate,
	}
}

// Execute executes the request
//
//	@return BackupOnlineArchive
func (a *OnlineArchiveApiService) CreateOnlineArchiveExecute(r CreateOnlineArchiveApiRequest) (*BackupOnlineArchive, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackupOnlineArchive
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnlineArchiveApiService.CreateOnlineArchive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.backupOnlineArchiveCreate == nil {
		return localVarReturnValue, nil, reportError("backupOnlineArchiveCreate is required and must be specified")
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
	localVarPostBody = r.backupOnlineArchiveCreate
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

type DeleteOnlineArchiveApiRequest struct {
	ctx         context.Context
	ApiService  OnlineArchiveApi
	groupId     string
	archiveId   string
	clusterName string
}

type DeleteOnlineArchiveApiParams struct {
	GroupId     string
	ArchiveId   string
	ClusterName string
}

func (a *OnlineArchiveApiService) DeleteOnlineArchiveWithParams(ctx context.Context, args *DeleteOnlineArchiveApiParams) DeleteOnlineArchiveApiRequest {
	return DeleteOnlineArchiveApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		archiveId:   args.ArchiveId,
		clusterName: args.ClusterName,
	}
}

func (r DeleteOnlineArchiveApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteOnlineArchiveExecute(r)
}

/*
DeleteOnlineArchive Remove One Online Archive

Removes one online archive. This archive stores data from one cluster within one project. To use this resource, the requesting API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param archiveId Unique 24-hexadecimal digit string that identifies the online archive to delete.
	@param clusterName Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive.
	@return DeleteOnlineArchiveApiRequest
*/
func (a *OnlineArchiveApiService) DeleteOnlineArchive(ctx context.Context, groupId string, archiveId string, clusterName string) DeleteOnlineArchiveApiRequest {
	return DeleteOnlineArchiveApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		archiveId:   archiveId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *OnlineArchiveApiService) DeleteOnlineArchiveExecute(r DeleteOnlineArchiveApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnlineArchiveApiService.DeleteOnlineArchive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"archiveId"+"}", url.PathEscape(parameterValueToString(r.archiveId, "archiveId")), -1)
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

type DownloadOnlineArchiveQueryLogsApiRequest struct {
	ctx         context.Context
	ApiService  OnlineArchiveApi
	groupId     string
	clusterName string
	startDate   *int64
	endDate     *int64
	archiveOnly *bool
}

type DownloadOnlineArchiveQueryLogsApiParams struct {
	GroupId     string
	ClusterName string
	StartDate   *int64
	EndDate     *int64
	ArchiveOnly *bool
}

func (a *OnlineArchiveApiService) DownloadOnlineArchiveQueryLogsWithParams(ctx context.Context, args *DownloadOnlineArchiveQueryLogsApiParams) DownloadOnlineArchiveQueryLogsApiRequest {
	return DownloadOnlineArchiveQueryLogsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		startDate:   args.StartDate,
		endDate:     args.EndDate,
		archiveOnly: args.ArchiveOnly,
	}
}

// Date and time that specifies the starting point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).
func (r DownloadOnlineArchiveQueryLogsApiRequest) StartDate(startDate int64) DownloadOnlineArchiveQueryLogsApiRequest {
	r.startDate = &startDate
	return r
}

// Date and time that specifies the end point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).
func (r DownloadOnlineArchiveQueryLogsApiRequest) EndDate(endDate int64) DownloadOnlineArchiveQueryLogsApiRequest {
	r.endDate = &endDate
	return r
}

// Flag that indicates whether to download logs for queries against your online archive only or both your online archive and cluster.
func (r DownloadOnlineArchiveQueryLogsApiRequest) ArchiveOnly(archiveOnly bool) DownloadOnlineArchiveQueryLogsApiRequest {
	r.archiveOnly = &archiveOnly
	return r
}

func (r DownloadOnlineArchiveQueryLogsApiRequest) Execute() (io.ReadCloser, *http.Response, error) {
	return r.ApiService.DownloadOnlineArchiveQueryLogsExecute(r)
}

/*
DownloadOnlineArchiveQueryLogs Download Online Archive Query Logs

[experimental] Downloads query logs for the specified online archive. To use this resource, the requesting API Key must have the Project Data Access Read Only or higher role. The API does not support direct calls with the json response schema. You must request a gzip response schema using an accept header of the format: "Accept: application/vnd.atlas.YYYY-MM-DD+gzip".

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster that contains the collection for which you want to return the query logs from one online archive.
	@return DownloadOnlineArchiveQueryLogsApiRequest
*/
func (a *OnlineArchiveApiService) DownloadOnlineArchiveQueryLogs(ctx context.Context, groupId string, clusterName string) DownloadOnlineArchiveQueryLogsApiRequest {
	return DownloadOnlineArchiveQueryLogsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return io.ReadCloser
func (a *OnlineArchiveApiService) DownloadOnlineArchiveQueryLogsExecute(r DownloadOnlineArchiveQueryLogsApiRequest) (io.ReadCloser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue io.ReadCloser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnlineArchiveApiService.DownloadOnlineArchiveQueryLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/queryLogs.gz"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "")
	}
	if r.archiveOnly != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archiveOnly", r.archiveOnly, "")
	} else {
		var defaultValue bool = false
		r.archiveOnly = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "archiveOnly", r.archiveOnly, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+gzip", "application/json"}

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

type GetOnlineArchiveApiRequest struct {
	ctx         context.Context
	ApiService  OnlineArchiveApi
	groupId     string
	archiveId   string
	clusterName string
}

type GetOnlineArchiveApiParams struct {
	GroupId     string
	ArchiveId   string
	ClusterName string
}

func (a *OnlineArchiveApiService) GetOnlineArchiveWithParams(ctx context.Context, args *GetOnlineArchiveApiParams) GetOnlineArchiveApiRequest {
	return GetOnlineArchiveApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		archiveId:   args.ArchiveId,
		clusterName: args.ClusterName,
	}
}

func (r GetOnlineArchiveApiRequest) Execute() (*BackupOnlineArchive, *http.Response, error) {
	return r.ApiService.GetOnlineArchiveExecute(r)
}

/*
GetOnlineArchive Return One Online Archive

Returns one online archive for one cluster. This archive stores data from one cluster within one project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param archiveId Unique 24-hexadecimal digit string that identifies the online archive to return.
	@param clusterName Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.
	@return GetOnlineArchiveApiRequest
*/
func (a *OnlineArchiveApiService) GetOnlineArchive(ctx context.Context, groupId string, archiveId string, clusterName string) GetOnlineArchiveApiRequest {
	return GetOnlineArchiveApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		archiveId:   archiveId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return BackupOnlineArchive
func (a *OnlineArchiveApiService) GetOnlineArchiveExecute(r GetOnlineArchiveApiRequest) (*BackupOnlineArchive, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackupOnlineArchive
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnlineArchiveApiService.GetOnlineArchive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"archiveId"+"}", url.PathEscape(parameterValueToString(r.archiveId, "archiveId")), -1)
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

type ListOnlineArchivesApiRequest struct {
	ctx          context.Context
	ApiService   OnlineArchiveApi
	groupId      string
	clusterName  string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListOnlineArchivesApiParams struct {
	GroupId      string
	ClusterName  string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *OnlineArchiveApiService) ListOnlineArchivesWithParams(ctx context.Context, args *ListOnlineArchivesApiParams) ListOnlineArchivesApiRequest {
	return ListOnlineArchivesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListOnlineArchivesApiRequest) IncludeCount(includeCount bool) ListOnlineArchivesApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListOnlineArchivesApiRequest) ItemsPerPage(itemsPerPage int) ListOnlineArchivesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListOnlineArchivesApiRequest) PageNum(pageNum int) ListOnlineArchivesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListOnlineArchivesApiRequest) Execute() (*PaginatedOnlineArchive, *http.Response, error) {
	return r.ApiService.ListOnlineArchivesExecute(r)
}

/*
ListOnlineArchives Return All Online Archives for One Cluster

Returns details of all online archives. This archive stores data from one cluster within one project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster that contains the collection for which you want to return the online archives.
	@return ListOnlineArchivesApiRequest
*/
func (a *OnlineArchiveApiService) ListOnlineArchives(ctx context.Context, groupId string, clusterName string) ListOnlineArchivesApiRequest {
	return ListOnlineArchivesApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PaginatedOnlineArchive
func (a *OnlineArchiveApiService) ListOnlineArchivesExecute(r ListOnlineArchivesApiRequest) (*PaginatedOnlineArchive, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedOnlineArchive
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnlineArchiveApiService.ListOnlineArchives")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

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

type UpdateOnlineArchiveApiRequest struct {
	ctx                 context.Context
	ApiService          OnlineArchiveApi
	groupId             string
	archiveId           string
	clusterName         string
	backupOnlineArchive *BackupOnlineArchive
}

type UpdateOnlineArchiveApiParams struct {
	GroupId             string
	ArchiveId           string
	ClusterName         string
	BackupOnlineArchive *BackupOnlineArchive
}

func (a *OnlineArchiveApiService) UpdateOnlineArchiveWithParams(ctx context.Context, args *UpdateOnlineArchiveApiParams) UpdateOnlineArchiveApiRequest {
	return UpdateOnlineArchiveApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		groupId:             args.GroupId,
		archiveId:           args.ArchiveId,
		clusterName:         args.ClusterName,
		backupOnlineArchive: args.BackupOnlineArchive,
	}
}

func (r UpdateOnlineArchiveApiRequest) Execute() (*BackupOnlineArchive, *http.Response, error) {
	return r.ApiService.UpdateOnlineArchiveExecute(r)
}

/*
UpdateOnlineArchive Update One Online Archive

Updates, pauses, or resumes one online archive. This archive stores data from one cluster within one project. To use this resource, the requesting API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param archiveId Unique 24-hexadecimal digit string that identifies the online archive to update.
	@param clusterName Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.
	@return UpdateOnlineArchiveApiRequest
*/
func (a *OnlineArchiveApiService) UpdateOnlineArchive(ctx context.Context, groupId string, archiveId string, clusterName string, backupOnlineArchive *BackupOnlineArchive) UpdateOnlineArchiveApiRequest {
	return UpdateOnlineArchiveApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		groupId:             groupId,
		archiveId:           archiveId,
		clusterName:         clusterName,
		backupOnlineArchive: backupOnlineArchive,
	}
}

// Execute executes the request
//
//	@return BackupOnlineArchive
func (a *OnlineArchiveApiService) UpdateOnlineArchiveExecute(r UpdateOnlineArchiveApiRequest) (*BackupOnlineArchive, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackupOnlineArchive
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnlineArchiveApiService.UpdateOnlineArchive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/onlineArchives/{archiveId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"archiveId"+"}", url.PathEscape(parameterValueToString(r.archiveId, "archiveId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.backupOnlineArchive == nil {
		return localVarReturnValue, nil, reportError("backupOnlineArchive is required and must be specified")
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
	localVarPostBody = r.backupOnlineArchive
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
