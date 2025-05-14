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

type StreamsApi interface {

	/*
		CreateStreamConnection Create One Connection

		[experimental] Creates one connection for a stream instance in the specified project. To use this resource, the requesting API Key must have the Project Owner or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@return CreateStreamConnectionApiRequest
	*/
	CreateStreamConnection(ctx context.Context, groupId string, tenantName string, streamsConnection *StreamsConnection) CreateStreamConnectionApiRequest
	/*
		CreateStreamConnection Create One Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateStreamConnectionApiParams - Parameters for the request
		@return CreateStreamConnectionApiRequest
	*/
	CreateStreamConnectionWithParams(ctx context.Context, args *CreateStreamConnectionApiParams) CreateStreamConnectionApiRequest

	// Method available only for mocking purposes
	CreateStreamConnectionExecute(r CreateStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error)

	/*
		CreateStreamInstance Create One Stream Instance

		[experimental] Creates one stream instance in the specified project. To use this resource, the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return CreateStreamInstanceApiRequest
	*/
	CreateStreamInstance(ctx context.Context, groupId string, streamsTenant *StreamsTenant) CreateStreamInstanceApiRequest
	/*
		CreateStreamInstance Create One Stream Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateStreamInstanceApiParams - Parameters for the request
		@return CreateStreamInstanceApiRequest
	*/
	CreateStreamInstanceWithParams(ctx context.Context, args *CreateStreamInstanceApiParams) CreateStreamInstanceApiRequest

	// Method available only for mocking purposes
	CreateStreamInstanceExecute(r CreateStreamInstanceApiRequest) (*StreamsTenant, *http.Response, error)

	/*
		DeleteStreamConnection Delete One Stream Connection

		[experimental] Delete one connection of the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param connectionName Human-readable label that identifies the stream connection.
		@return DeleteStreamConnectionApiRequest
	*/
	DeleteStreamConnection(ctx context.Context, groupId string, tenantName string, connectionName string) DeleteStreamConnectionApiRequest
	/*
		DeleteStreamConnection Delete One Stream Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteStreamConnectionApiParams - Parameters for the request
		@return DeleteStreamConnectionApiRequest
	*/
	DeleteStreamConnectionWithParams(ctx context.Context, args *DeleteStreamConnectionApiParams) DeleteStreamConnectionApiRequest

	// Method available only for mocking purposes
	DeleteStreamConnectionExecute(r DeleteStreamConnectionApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		DeleteStreamInstance Delete One Stream Instance

		[experimental] Delete one stream instance in the specified project. To use this resource, the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance to delete.
		@return DeleteStreamInstanceApiRequest
	*/
	DeleteStreamInstance(ctx context.Context, groupId string, tenantName string) DeleteStreamInstanceApiRequest
	/*
		DeleteStreamInstance Delete One Stream Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteStreamInstanceApiParams - Parameters for the request
		@return DeleteStreamInstanceApiRequest
	*/
	DeleteStreamInstanceWithParams(ctx context.Context, args *DeleteStreamInstanceApiParams) DeleteStreamInstanceApiRequest

	// Method available only for mocking purposes
	DeleteStreamInstanceExecute(r DeleteStreamInstanceApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		DownloadStreamTenantAuditLogs Download Audit Logs for One Atlas Stream Processing Instance

		[experimental] Downloads the audit logs for the specified Atlas Streams Processing instance. By default, logs cover periods of 30 days. To use this resource, the requesting API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role. The API does not support direct calls with the json response schema. You must request a gzip response schema using an accept header of the format: "Accept: application/vnd.atlas.YYYY-MM-DD+gzip".

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@return DownloadStreamTenantAuditLogsApiRequest
	*/
	DownloadStreamTenantAuditLogs(ctx context.Context, groupId string, tenantName string) DownloadStreamTenantAuditLogsApiRequest
	/*
		DownloadStreamTenantAuditLogs Download Audit Logs for One Atlas Stream Processing Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DownloadStreamTenantAuditLogsApiParams - Parameters for the request
		@return DownloadStreamTenantAuditLogsApiRequest
	*/
	DownloadStreamTenantAuditLogsWithParams(ctx context.Context, args *DownloadStreamTenantAuditLogsApiParams) DownloadStreamTenantAuditLogsApiRequest

	// Method available only for mocking purposes
	DownloadStreamTenantAuditLogsExecute(r DownloadStreamTenantAuditLogsApiRequest) (io.ReadCloser, *http.Response, error)

	/*
		GetStreamConnection Return One Stream Connection

		[experimental] Returns the details of one stream connection within the specified stream instance. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance to return.
		@param connectionName Human-readable label that identifies the stream connection to return.
		@return GetStreamConnectionApiRequest
	*/
	GetStreamConnection(ctx context.Context, groupId string, tenantName string, connectionName string) GetStreamConnectionApiRequest
	/*
		GetStreamConnection Return One Stream Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetStreamConnectionApiParams - Parameters for the request
		@return GetStreamConnectionApiRequest
	*/
	GetStreamConnectionWithParams(ctx context.Context, args *GetStreamConnectionApiParams) GetStreamConnectionApiRequest

	// Method available only for mocking purposes
	GetStreamConnectionExecute(r GetStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error)

	/*
		GetStreamInstance Return One Stream Instance

		[experimental] Returns the details of one stream instance within the specified project. To use this resource, the requesting API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance to return.
		@return GetStreamInstanceApiRequest
	*/
	GetStreamInstance(ctx context.Context, groupId string, tenantName string) GetStreamInstanceApiRequest
	/*
		GetStreamInstance Return One Stream Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetStreamInstanceApiParams - Parameters for the request
		@return GetStreamInstanceApiRequest
	*/
	GetStreamInstanceWithParams(ctx context.Context, args *GetStreamInstanceApiParams) GetStreamInstanceApiRequest

	// Method available only for mocking purposes
	GetStreamInstanceExecute(r GetStreamInstanceApiRequest) (*StreamsTenant, *http.Response, error)

	/*
		ListStreamConnections Return All Connections Of The Stream Instances

		[experimental] Returns all connections of the stream instance for the specified project.To use this resource, the requesting API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@return ListStreamConnectionsApiRequest
	*/
	ListStreamConnections(ctx context.Context, groupId string, tenantName string) ListStreamConnectionsApiRequest
	/*
		ListStreamConnections Return All Connections Of The Stream Instances


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListStreamConnectionsApiParams - Parameters for the request
		@return ListStreamConnectionsApiRequest
	*/
	ListStreamConnectionsWithParams(ctx context.Context, args *ListStreamConnectionsApiParams) ListStreamConnectionsApiRequest

	// Method available only for mocking purposes
	ListStreamConnectionsExecute(r ListStreamConnectionsApiRequest) (*PaginatedApiStreamsConnection, *http.Response, error)

	/*
		ListStreamInstances Return All Project Stream Instances

		[experimental] Returns all stream instances for the specified project.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListStreamInstancesApiRequest
	*/
	ListStreamInstances(ctx context.Context, groupId string) ListStreamInstancesApiRequest
	/*
		ListStreamInstances Return All Project Stream Instances


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListStreamInstancesApiParams - Parameters for the request
		@return ListStreamInstancesApiRequest
	*/
	ListStreamInstancesWithParams(ctx context.Context, args *ListStreamInstancesApiParams) ListStreamInstancesApiRequest

	// Method available only for mocking purposes
	ListStreamInstancesExecute(r ListStreamInstancesApiRequest) (*PaginatedApiStreamsTenant, *http.Response, error)

	/*
		UpdateStreamConnection Update One Stream Connection

		[experimental] Update one connection for the specified stream instance in the specified project. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance.
		@param connectionName Human-readable label that identifies the stream connection.
		@return UpdateStreamConnectionApiRequest
	*/
	UpdateStreamConnection(ctx context.Context, groupId string, tenantName string, connectionName string, streamsConnection *StreamsConnection) UpdateStreamConnectionApiRequest
	/*
		UpdateStreamConnection Update One Stream Connection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateStreamConnectionApiParams - Parameters for the request
		@return UpdateStreamConnectionApiRequest
	*/
	UpdateStreamConnectionWithParams(ctx context.Context, args *UpdateStreamConnectionApiParams) UpdateStreamConnectionApiRequest

	// Method available only for mocking purposes
	UpdateStreamConnectionExecute(r UpdateStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error)

	/*
		UpdateStreamInstance Update One Stream Instance

		[experimental] Update one stream instance in the specified project. To use this resource, the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param tenantName Human-readable label that identifies the stream instance to update.
		@return UpdateStreamInstanceApiRequest
	*/
	UpdateStreamInstance(ctx context.Context, groupId string, tenantName string, streamsDataProcessRegion *StreamsDataProcessRegion) UpdateStreamInstanceApiRequest
	/*
		UpdateStreamInstance Update One Stream Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateStreamInstanceApiParams - Parameters for the request
		@return UpdateStreamInstanceApiRequest
	*/
	UpdateStreamInstanceWithParams(ctx context.Context, args *UpdateStreamInstanceApiParams) UpdateStreamInstanceApiRequest

	// Method available only for mocking purposes
	UpdateStreamInstanceExecute(r UpdateStreamInstanceApiRequest) (*StreamsTenant, *http.Response, error)
}

// StreamsApiService StreamsApi service
type StreamsApiService service

type CreateStreamConnectionApiRequest struct {
	ctx               context.Context
	ApiService        StreamsApi
	groupId           string
	tenantName        string
	streamsConnection *StreamsConnection
}

type CreateStreamConnectionApiParams struct {
	GroupId           string
	TenantName        string
	StreamsConnection *StreamsConnection
}

func (a *StreamsApiService) CreateStreamConnectionWithParams(ctx context.Context, args *CreateStreamConnectionApiParams) CreateStreamConnectionApiRequest {
	return CreateStreamConnectionApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		tenantName:        args.TenantName,
		streamsConnection: args.StreamsConnection,
	}
}

func (r CreateStreamConnectionApiRequest) Execute() (*StreamsConnection, *http.Response, error) {
	return r.ApiService.CreateStreamConnectionExecute(r)
}

/*
CreateStreamConnection Create One Connection

[experimental] Creates one connection for a stream instance in the specified project. To use this resource, the requesting API Key must have the Project Owner or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@return CreateStreamConnectionApiRequest
*/
func (a *StreamsApiService) CreateStreamConnection(ctx context.Context, groupId string, tenantName string, streamsConnection *StreamsConnection) CreateStreamConnectionApiRequest {
	return CreateStreamConnectionApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		tenantName:        tenantName,
		streamsConnection: streamsConnection,
	}
}

// Execute executes the request
//
//	@return StreamsConnection
func (a *StreamsApiService) CreateStreamConnectionExecute(r CreateStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StreamsConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.CreateStreamConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.streamsConnection == nil {
		return localVarReturnValue, nil, reportError("streamsConnection is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-02-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.streamsConnection
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

type CreateStreamInstanceApiRequest struct {
	ctx           context.Context
	ApiService    StreamsApi
	groupId       string
	streamsTenant *StreamsTenant
}

type CreateStreamInstanceApiParams struct {
	GroupId       string
	StreamsTenant *StreamsTenant
}

func (a *StreamsApiService) CreateStreamInstanceWithParams(ctx context.Context, args *CreateStreamInstanceApiParams) CreateStreamInstanceApiRequest {
	return CreateStreamInstanceApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		streamsTenant: args.StreamsTenant,
	}
}

func (r CreateStreamInstanceApiRequest) Execute() (*StreamsTenant, *http.Response, error) {
	return r.ApiService.CreateStreamInstanceExecute(r)
}

/*
CreateStreamInstance Create One Stream Instance

[experimental] Creates one stream instance in the specified project. To use this resource, the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateStreamInstanceApiRequest
*/
func (a *StreamsApiService) CreateStreamInstance(ctx context.Context, groupId string, streamsTenant *StreamsTenant) CreateStreamInstanceApiRequest {
	return CreateStreamInstanceApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		streamsTenant: streamsTenant,
	}
}

// Execute executes the request
//
//	@return StreamsTenant
func (a *StreamsApiService) CreateStreamInstanceExecute(r CreateStreamInstanceApiRequest) (*StreamsTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StreamsTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.CreateStreamInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.streamsTenant == nil {
		return localVarReturnValue, nil, reportError("streamsTenant is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-02-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.streamsTenant
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

type DeleteStreamConnectionApiRequest struct {
	ctx            context.Context
	ApiService     StreamsApi
	groupId        string
	tenantName     string
	connectionName string
}

type DeleteStreamConnectionApiParams struct {
	GroupId        string
	TenantName     string
	ConnectionName string
}

func (a *StreamsApiService) DeleteStreamConnectionWithParams(ctx context.Context, args *DeleteStreamConnectionApiParams) DeleteStreamConnectionApiRequest {
	return DeleteStreamConnectionApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		tenantName:     args.TenantName,
		connectionName: args.ConnectionName,
	}
}

func (r DeleteStreamConnectionApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteStreamConnectionExecute(r)
}

/*
DeleteStreamConnection Delete One Stream Connection

[experimental] Delete one connection of the specified stream instance. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@param connectionName Human-readable label that identifies the stream connection.
	@return DeleteStreamConnectionApiRequest
*/
func (a *StreamsApiService) DeleteStreamConnection(ctx context.Context, groupId string, tenantName string, connectionName string) DeleteStreamConnectionApiRequest {
	return DeleteStreamConnectionApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		tenantName:     tenantName,
		connectionName: connectionName,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *StreamsApiService) DeleteStreamConnectionExecute(r DeleteStreamConnectionApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.DeleteStreamConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connectionName"+"}", url.PathEscape(parameterValueToString(r.connectionName, "connectionName")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

type DeleteStreamInstanceApiRequest struct {
	ctx        context.Context
	ApiService StreamsApi
	groupId    string
	tenantName string
}

type DeleteStreamInstanceApiParams struct {
	GroupId    string
	TenantName string
}

func (a *StreamsApiService) DeleteStreamInstanceWithParams(ctx context.Context, args *DeleteStreamInstanceApiParams) DeleteStreamInstanceApiRequest {
	return DeleteStreamInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		tenantName: args.TenantName,
	}
}

func (r DeleteStreamInstanceApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteStreamInstanceExecute(r)
}

/*
DeleteStreamInstance Delete One Stream Instance

[experimental] Delete one stream instance in the specified project. To use this resource, the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance to delete.
	@return DeleteStreamInstanceApiRequest
*/
func (a *StreamsApiService) DeleteStreamInstance(ctx context.Context, groupId string, tenantName string) DeleteStreamInstanceApiRequest {
	return DeleteStreamInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *StreamsApiService) DeleteStreamInstanceExecute(r DeleteStreamInstanceApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.DeleteStreamInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

type DownloadStreamTenantAuditLogsApiRequest struct {
	ctx        context.Context
	ApiService StreamsApi
	groupId    string
	tenantName string
	endDate    *int64
	startDate  *int64
}

type DownloadStreamTenantAuditLogsApiParams struct {
	GroupId    string
	TenantName string
	EndDate    *int64
	StartDate  *int64
}

func (a *StreamsApiService) DownloadStreamTenantAuditLogsWithParams(ctx context.Context, args *DownloadStreamTenantAuditLogsApiParams) DownloadStreamTenantAuditLogsApiRequest {
	return DownloadStreamTenantAuditLogsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		tenantName: args.TenantName,
		endDate:    args.EndDate,
		startDate:  args.StartDate,
	}
}

// Timestamp that specifies the end point for the range of log messages to download.  MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.
func (r DownloadStreamTenantAuditLogsApiRequest) EndDate(endDate int64) DownloadStreamTenantAuditLogsApiRequest {
	r.endDate = &endDate
	return r
}

// Timestamp that specifies the starting point for the range of log messages to download. MongoDB Cloud expresses this timestamp in the number of seconds that have elapsed since the UNIX epoch.
func (r DownloadStreamTenantAuditLogsApiRequest) StartDate(startDate int64) DownloadStreamTenantAuditLogsApiRequest {
	r.startDate = &startDate
	return r
}

func (r DownloadStreamTenantAuditLogsApiRequest) Execute() (io.ReadCloser, *http.Response, error) {
	return r.ApiService.DownloadStreamTenantAuditLogsExecute(r)
}

/*
DownloadStreamTenantAuditLogs Download Audit Logs for One Atlas Stream Processing Instance

[experimental] Downloads the audit logs for the specified Atlas Streams Processing instance. By default, logs cover periods of 30 days. To use this resource, the requesting API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role. The API does not support direct calls with the json response schema. You must request a gzip response schema using an accept header of the format: "Accept: application/vnd.atlas.YYYY-MM-DD+gzip".

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@return DownloadStreamTenantAuditLogsApiRequest
*/
func (a *StreamsApiService) DownloadStreamTenantAuditLogs(ctx context.Context, groupId string, tenantName string) DownloadStreamTenantAuditLogsApiRequest {
	return DownloadStreamTenantAuditLogsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// Execute executes the request
//
//	@return io.ReadCloser
func (a *StreamsApiService) DownloadStreamTenantAuditLogsExecute(r DownloadStreamTenantAuditLogsApiRequest) (io.ReadCloser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue io.ReadCloser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.DownloadStreamTenantAuditLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/auditLogs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+gzip", "application/json"}

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

type GetStreamConnectionApiRequest struct {
	ctx            context.Context
	ApiService     StreamsApi
	groupId        string
	tenantName     string
	connectionName string
}

type GetStreamConnectionApiParams struct {
	GroupId        string
	TenantName     string
	ConnectionName string
}

func (a *StreamsApiService) GetStreamConnectionWithParams(ctx context.Context, args *GetStreamConnectionApiParams) GetStreamConnectionApiRequest {
	return GetStreamConnectionApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		tenantName:     args.TenantName,
		connectionName: args.ConnectionName,
	}
}

func (r GetStreamConnectionApiRequest) Execute() (*StreamsConnection, *http.Response, error) {
	return r.ApiService.GetStreamConnectionExecute(r)
}

/*
GetStreamConnection Return One Stream Connection

[experimental] Returns the details of one stream connection within the specified stream instance. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance to return.
	@param connectionName Human-readable label that identifies the stream connection to return.
	@return GetStreamConnectionApiRequest
*/
func (a *StreamsApiService) GetStreamConnection(ctx context.Context, groupId string, tenantName string, connectionName string) GetStreamConnectionApiRequest {
	return GetStreamConnectionApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		tenantName:     tenantName,
		connectionName: connectionName,
	}
}

// Execute executes the request
//
//	@return StreamsConnection
func (a *StreamsApiService) GetStreamConnectionExecute(r GetStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StreamsConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.GetStreamConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connectionName"+"}", url.PathEscape(parameterValueToString(r.connectionName, "connectionName")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

type GetStreamInstanceApiRequest struct {
	ctx                context.Context
	ApiService         StreamsApi
	groupId            string
	tenantName         string
	includeConnections *bool
}

type GetStreamInstanceApiParams struct {
	GroupId            string
	TenantName         string
	IncludeConnections *bool
}

func (a *StreamsApiService) GetStreamInstanceWithParams(ctx context.Context, args *GetStreamInstanceApiParams) GetStreamInstanceApiRequest {
	return GetStreamInstanceApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		tenantName:         args.TenantName,
		includeConnections: args.IncludeConnections,
	}
}

// Flag to indicate whether connections information should be included in the stream instance.
func (r GetStreamInstanceApiRequest) IncludeConnections(includeConnections bool) GetStreamInstanceApiRequest {
	r.includeConnections = &includeConnections
	return r
}

func (r GetStreamInstanceApiRequest) Execute() (*StreamsTenant, *http.Response, error) {
	return r.ApiService.GetStreamInstanceExecute(r)
}

/*
GetStreamInstance Return One Stream Instance

[experimental] Returns the details of one stream instance within the specified project. To use this resource, the requesting API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance to return.
	@return GetStreamInstanceApiRequest
*/
func (a *StreamsApiService) GetStreamInstance(ctx context.Context, groupId string, tenantName string) GetStreamInstanceApiRequest {
	return GetStreamInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// Execute executes the request
//
//	@return StreamsTenant
func (a *StreamsApiService) GetStreamInstanceExecute(r GetStreamInstanceApiRequest) (*StreamsTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StreamsTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.GetStreamInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeConnections != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeConnections", r.includeConnections, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

type ListStreamConnectionsApiRequest struct {
	ctx          context.Context
	ApiService   StreamsApi
	groupId      string
	tenantName   string
	itemsPerPage *int
	pageNum      *int
}

type ListStreamConnectionsApiParams struct {
	GroupId      string
	TenantName   string
	ItemsPerPage *int
	PageNum      *int
}

func (a *StreamsApiService) ListStreamConnectionsWithParams(ctx context.Context, args *ListStreamConnectionsApiParams) ListStreamConnectionsApiRequest {
	return ListStreamConnectionsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		tenantName:   args.TenantName,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListStreamConnectionsApiRequest) ItemsPerPage(itemsPerPage int) ListStreamConnectionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListStreamConnectionsApiRequest) PageNum(pageNum int) ListStreamConnectionsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListStreamConnectionsApiRequest) Execute() (*PaginatedApiStreamsConnection, *http.Response, error) {
	return r.ApiService.ListStreamConnectionsExecute(r)
}

/*
ListStreamConnections Return All Connections Of The Stream Instances

[experimental] Returns all connections of the stream instance for the specified project.To use this resource, the requesting API Key must have the Project Data Access roles, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@return ListStreamConnectionsApiRequest
*/
func (a *StreamsApiService) ListStreamConnections(ctx context.Context, groupId string, tenantName string) ListStreamConnectionsApiRequest {
	return ListStreamConnectionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		tenantName: tenantName,
	}
}

// Execute executes the request
//
//	@return PaginatedApiStreamsConnection
func (a *StreamsApiService) ListStreamConnectionsExecute(r ListStreamConnectionsApiRequest) (*PaginatedApiStreamsConnection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedApiStreamsConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.ListStreamConnections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

type ListStreamInstancesApiRequest struct {
	ctx          context.Context
	ApiService   StreamsApi
	groupId      string
	itemsPerPage *int
	pageNum      *int
}

type ListStreamInstancesApiParams struct {
	GroupId      string
	ItemsPerPage *int
	PageNum      *int
}

func (a *StreamsApiService) ListStreamInstancesWithParams(ctx context.Context, args *ListStreamInstancesApiParams) ListStreamInstancesApiRequest {
	return ListStreamInstancesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListStreamInstancesApiRequest) ItemsPerPage(itemsPerPage int) ListStreamInstancesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListStreamInstancesApiRequest) PageNum(pageNum int) ListStreamInstancesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListStreamInstancesApiRequest) Execute() (*PaginatedApiStreamsTenant, *http.Response, error) {
	return r.ApiService.ListStreamInstancesExecute(r)
}

/*
ListStreamInstances Return All Project Stream Instances

[experimental] Returns all stream instances for the specified project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListStreamInstancesApiRequest
*/
func (a *StreamsApiService) ListStreamInstances(ctx context.Context, groupId string) ListStreamInstancesApiRequest {
	return ListStreamInstancesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return PaginatedApiStreamsTenant
func (a *StreamsApiService) ListStreamInstancesExecute(r ListStreamInstancesApiRequest) (*PaginatedApiStreamsTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedApiStreamsTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.ListStreamInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

type UpdateStreamConnectionApiRequest struct {
	ctx               context.Context
	ApiService        StreamsApi
	groupId           string
	tenantName        string
	connectionName    string
	streamsConnection *StreamsConnection
}

type UpdateStreamConnectionApiParams struct {
	GroupId           string
	TenantName        string
	ConnectionName    string
	StreamsConnection *StreamsConnection
}

func (a *StreamsApiService) UpdateStreamConnectionWithParams(ctx context.Context, args *UpdateStreamConnectionApiParams) UpdateStreamConnectionApiRequest {
	return UpdateStreamConnectionApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           args.GroupId,
		tenantName:        args.TenantName,
		connectionName:    args.ConnectionName,
		streamsConnection: args.StreamsConnection,
	}
}

func (r UpdateStreamConnectionApiRequest) Execute() (*StreamsConnection, *http.Response, error) {
	return r.ApiService.UpdateStreamConnectionExecute(r)
}

/*
UpdateStreamConnection Update One Stream Connection

[experimental] Update one connection for the specified stream instance in the specified project. To use this resource, the requesting API Key must have the Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance.
	@param connectionName Human-readable label that identifies the stream connection.
	@return UpdateStreamConnectionApiRequest
*/
func (a *StreamsApiService) UpdateStreamConnection(ctx context.Context, groupId string, tenantName string, connectionName string, streamsConnection *StreamsConnection) UpdateStreamConnectionApiRequest {
	return UpdateStreamConnectionApiRequest{
		ApiService:        a,
		ctx:               ctx,
		groupId:           groupId,
		tenantName:        tenantName,
		connectionName:    connectionName,
		streamsConnection: streamsConnection,
	}
}

// Execute executes the request
//
//	@return StreamsConnection
func (a *StreamsApiService) UpdateStreamConnectionExecute(r UpdateStreamConnectionApiRequest) (*StreamsConnection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StreamsConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.UpdateStreamConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connectionName"+"}", url.PathEscape(parameterValueToString(r.connectionName, "connectionName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.streamsConnection == nil {
		return localVarReturnValue, nil, reportError("streamsConnection is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-02-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.streamsConnection
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

type UpdateStreamInstanceApiRequest struct {
	ctx                      context.Context
	ApiService               StreamsApi
	groupId                  string
	tenantName               string
	streamsDataProcessRegion *StreamsDataProcessRegion
}

type UpdateStreamInstanceApiParams struct {
	GroupId                  string
	TenantName               string
	StreamsDataProcessRegion *StreamsDataProcessRegion
}

func (a *StreamsApiService) UpdateStreamInstanceWithParams(ctx context.Context, args *UpdateStreamInstanceApiParams) UpdateStreamInstanceApiRequest {
	return UpdateStreamInstanceApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  args.GroupId,
		tenantName:               args.TenantName,
		streamsDataProcessRegion: args.StreamsDataProcessRegion,
	}
}

func (r UpdateStreamInstanceApiRequest) Execute() (*StreamsTenant, *http.Response, error) {
	return r.ApiService.UpdateStreamInstanceExecute(r)
}

/*
UpdateStreamInstance Update One Stream Instance

[experimental] Update one stream instance in the specified project. To use this resource, the requesting API Key must have the Project Data Access Admin role, Project Owner role or Project Stream Processing Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param tenantName Human-readable label that identifies the stream instance to update.
	@return UpdateStreamInstanceApiRequest
*/
func (a *StreamsApiService) UpdateStreamInstance(ctx context.Context, groupId string, tenantName string, streamsDataProcessRegion *StreamsDataProcessRegion) UpdateStreamInstanceApiRequest {
	return UpdateStreamInstanceApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  groupId,
		tenantName:               tenantName,
		streamsDataProcessRegion: streamsDataProcessRegion,
	}
}

// Execute executes the request
//
//	@return StreamsTenant
func (a *StreamsApiService) UpdateStreamInstanceExecute(r UpdateStreamInstanceApiRequest) (*StreamsTenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StreamsTenant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StreamsApiService.UpdateStreamInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/streams/{tenantName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tenantName"+"}", url.PathEscape(parameterValueToString(r.tenantName, "tenantName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.streamsDataProcessRegion == nil {
		return localVarReturnValue, nil, reportError("streamsDataProcessRegion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-02-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.streamsDataProcessRegion
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
