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

type ServerlessInstancesApi interface {

	/*
		CreateServerlessInstance Create One Serverless Instance in One Project

		Creates one serverless instance in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return CreateServerlessInstanceApiRequest
	*/
	CreateServerlessInstance(ctx context.Context, groupId string, serverlessInstanceDescriptionCreate *ServerlessInstanceDescriptionCreate) CreateServerlessInstanceApiRequest
	/*
		CreateServerlessInstance Create One Serverless Instance in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateServerlessInstanceApiParams - Parameters for the request
		@return CreateServerlessInstanceApiRequest
	*/
	CreateServerlessInstanceWithParams(ctx context.Context, args *CreateServerlessInstanceApiParams) CreateServerlessInstanceApiRequest

	// Method available only for mocking purposes
	CreateServerlessInstanceExecute(r CreateServerlessInstanceApiRequest) (*ServerlessInstanceDescription, *http.Response, error)

	/*
		DeleteServerlessInstance Remove One Serverless Instance from One Project

		Removes one serverless instance from the specified project. The serverless instance must have termination protection disabled in order to be deleted. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the serverless instance.
		@return DeleteServerlessInstanceApiRequest
	*/
	DeleteServerlessInstance(ctx context.Context, groupId string, name string) DeleteServerlessInstanceApiRequest
	/*
		DeleteServerlessInstance Remove One Serverless Instance from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteServerlessInstanceApiParams - Parameters for the request
		@return DeleteServerlessInstanceApiRequest
	*/
	DeleteServerlessInstanceWithParams(ctx context.Context, args *DeleteServerlessInstanceApiParams) DeleteServerlessInstanceApiRequest

	// Method available only for mocking purposes
	DeleteServerlessInstanceExecute(r DeleteServerlessInstanceApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetServerlessInstance Return One Serverless Instance from One Project

		Returns details for one serverless instance in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the serverless instance.
		@return GetServerlessInstanceApiRequest
	*/
	GetServerlessInstance(ctx context.Context, groupId string, name string) GetServerlessInstanceApiRequest
	/*
		GetServerlessInstance Return One Serverless Instance from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetServerlessInstanceApiParams - Parameters for the request
		@return GetServerlessInstanceApiRequest
	*/
	GetServerlessInstanceWithParams(ctx context.Context, args *GetServerlessInstanceApiParams) GetServerlessInstanceApiRequest

	// Method available only for mocking purposes
	GetServerlessInstanceExecute(r GetServerlessInstanceApiRequest) (*ServerlessInstanceDescription, *http.Response, error)

	/*
		ListServerlessInstances Return All Serverless Instances from One Project

		Returns details for all serverless instances in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListServerlessInstancesApiRequest
	*/
	ListServerlessInstances(ctx context.Context, groupId string) ListServerlessInstancesApiRequest
	/*
		ListServerlessInstances Return All Serverless Instances from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListServerlessInstancesApiParams - Parameters for the request
		@return ListServerlessInstancesApiRequest
	*/
	ListServerlessInstancesWithParams(ctx context.Context, args *ListServerlessInstancesApiParams) ListServerlessInstancesApiRequest

	// Method available only for mocking purposes
	ListServerlessInstancesExecute(r ListServerlessInstancesApiRequest) (*PaginatedServerlessInstanceDescription, *http.Response, error)

	/*
		UpdateServerlessInstance Update One Serverless Instance in One Project

		Updates one serverless instance in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the serverless instance.
		@return UpdateServerlessInstanceApiRequest
	*/
	UpdateServerlessInstance(ctx context.Context, groupId string, name string, serverlessInstanceDescriptionUpdate *ServerlessInstanceDescriptionUpdate) UpdateServerlessInstanceApiRequest
	/*
		UpdateServerlessInstance Update One Serverless Instance in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateServerlessInstanceApiParams - Parameters for the request
		@return UpdateServerlessInstanceApiRequest
	*/
	UpdateServerlessInstanceWithParams(ctx context.Context, args *UpdateServerlessInstanceApiParams) UpdateServerlessInstanceApiRequest

	// Method available only for mocking purposes
	UpdateServerlessInstanceExecute(r UpdateServerlessInstanceApiRequest) (*ServerlessInstanceDescription, *http.Response, error)
}

// ServerlessInstancesApiService ServerlessInstancesApi service
type ServerlessInstancesApiService service

type CreateServerlessInstanceApiRequest struct {
	ctx                                 context.Context
	ApiService                          ServerlessInstancesApi
	groupId                             string
	serverlessInstanceDescriptionCreate *ServerlessInstanceDescriptionCreate
}

type CreateServerlessInstanceApiParams struct {
	GroupId                             string
	ServerlessInstanceDescriptionCreate *ServerlessInstanceDescriptionCreate
}

func (a *ServerlessInstancesApiService) CreateServerlessInstanceWithParams(ctx context.Context, args *CreateServerlessInstanceApiParams) CreateServerlessInstanceApiRequest {
	return CreateServerlessInstanceApiRequest{
		ApiService:                          a,
		ctx:                                 ctx,
		groupId:                             args.GroupId,
		serverlessInstanceDescriptionCreate: args.ServerlessInstanceDescriptionCreate,
	}
}

func (r CreateServerlessInstanceApiRequest) Execute() (*ServerlessInstanceDescription, *http.Response, error) {
	return r.ApiService.CreateServerlessInstanceExecute(r)
}

/*
CreateServerlessInstance Create One Serverless Instance in One Project

Creates one serverless instance in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateServerlessInstanceApiRequest
*/
func (a *ServerlessInstancesApiService) CreateServerlessInstance(ctx context.Context, groupId string, serverlessInstanceDescriptionCreate *ServerlessInstanceDescriptionCreate) CreateServerlessInstanceApiRequest {
	return CreateServerlessInstanceApiRequest{
		ApiService:                          a,
		ctx:                                 ctx,
		groupId:                             groupId,
		serverlessInstanceDescriptionCreate: serverlessInstanceDescriptionCreate,
	}
}

// Execute executes the request
//
//	@return ServerlessInstanceDescription
func (a *ServerlessInstancesApiService) CreateServerlessInstanceExecute(r CreateServerlessInstanceApiRequest) (*ServerlessInstanceDescription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServerlessInstanceDescription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessInstancesApiService.CreateServerlessInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serverlessInstanceDescriptionCreate == nil {
		return localVarReturnValue, nil, reportError("serverlessInstanceDescriptionCreate is required and must be specified")
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
	localVarPostBody = r.serverlessInstanceDescriptionCreate
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

type DeleteServerlessInstanceApiRequest struct {
	ctx        context.Context
	ApiService ServerlessInstancesApi
	groupId    string
	name       string
}

type DeleteServerlessInstanceApiParams struct {
	GroupId string
	Name    string
}

func (a *ServerlessInstancesApiService) DeleteServerlessInstanceWithParams(ctx context.Context, args *DeleteServerlessInstanceApiParams) DeleteServerlessInstanceApiRequest {
	return DeleteServerlessInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		name:       args.Name,
	}
}

func (r DeleteServerlessInstanceApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteServerlessInstanceExecute(r)
}

/*
DeleteServerlessInstance Remove One Serverless Instance from One Project

Removes one serverless instance from the specified project. The serverless instance must have termination protection disabled in order to be deleted. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the serverless instance.
	@return DeleteServerlessInstanceApiRequest
*/
func (a *ServerlessInstancesApiService) DeleteServerlessInstance(ctx context.Context, groupId string, name string) DeleteServerlessInstanceApiRequest {
	return DeleteServerlessInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		name:       name,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *ServerlessInstancesApiService) DeleteServerlessInstanceExecute(r DeleteServerlessInstanceApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessInstancesApiService.DeleteServerlessInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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

type GetServerlessInstanceApiRequest struct {
	ctx        context.Context
	ApiService ServerlessInstancesApi
	groupId    string
	name       string
}

type GetServerlessInstanceApiParams struct {
	GroupId string
	Name    string
}

func (a *ServerlessInstancesApiService) GetServerlessInstanceWithParams(ctx context.Context, args *GetServerlessInstanceApiParams) GetServerlessInstanceApiRequest {
	return GetServerlessInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		name:       args.Name,
	}
}

func (r GetServerlessInstanceApiRequest) Execute() (*ServerlessInstanceDescription, *http.Response, error) {
	return r.ApiService.GetServerlessInstanceExecute(r)
}

/*
GetServerlessInstance Return One Serverless Instance from One Project

Returns details for one serverless instance in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the serverless instance.
	@return GetServerlessInstanceApiRequest
*/
func (a *ServerlessInstancesApiService) GetServerlessInstance(ctx context.Context, groupId string, name string) GetServerlessInstanceApiRequest {
	return GetServerlessInstanceApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		name:       name,
	}
}

// Execute executes the request
//
//	@return ServerlessInstanceDescription
func (a *ServerlessInstancesApiService) GetServerlessInstanceExecute(r GetServerlessInstanceApiRequest) (*ServerlessInstanceDescription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServerlessInstanceDescription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessInstancesApiService.GetServerlessInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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

type ListServerlessInstancesApiRequest struct {
	ctx          context.Context
	ApiService   ServerlessInstancesApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListServerlessInstancesApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ServerlessInstancesApiService) ListServerlessInstancesWithParams(ctx context.Context, args *ListServerlessInstancesApiParams) ListServerlessInstancesApiRequest {
	return ListServerlessInstancesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListServerlessInstancesApiRequest) IncludeCount(includeCount bool) ListServerlessInstancesApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListServerlessInstancesApiRequest) ItemsPerPage(itemsPerPage int) ListServerlessInstancesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListServerlessInstancesApiRequest) PageNum(pageNum int) ListServerlessInstancesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListServerlessInstancesApiRequest) Execute() (*PaginatedServerlessInstanceDescription, *http.Response, error) {
	return r.ApiService.ListServerlessInstancesExecute(r)
}

/*
ListServerlessInstances Return All Serverless Instances from One Project

Returns details for all serverless instances in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListServerlessInstancesApiRequest
*/
func (a *ServerlessInstancesApiService) ListServerlessInstances(ctx context.Context, groupId string) ListServerlessInstancesApiRequest {
	return ListServerlessInstancesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return PaginatedServerlessInstanceDescription
func (a *ServerlessInstancesApiService) ListServerlessInstancesExecute(r ListServerlessInstancesApiRequest) (*PaginatedServerlessInstanceDescription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedServerlessInstanceDescription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessInstancesApiService.ListServerlessInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type UpdateServerlessInstanceApiRequest struct {
	ctx                                 context.Context
	ApiService                          ServerlessInstancesApi
	groupId                             string
	name                                string
	serverlessInstanceDescriptionUpdate *ServerlessInstanceDescriptionUpdate
}

type UpdateServerlessInstanceApiParams struct {
	GroupId                             string
	Name                                string
	ServerlessInstanceDescriptionUpdate *ServerlessInstanceDescriptionUpdate
}

func (a *ServerlessInstancesApiService) UpdateServerlessInstanceWithParams(ctx context.Context, args *UpdateServerlessInstanceApiParams) UpdateServerlessInstanceApiRequest {
	return UpdateServerlessInstanceApiRequest{
		ApiService:                          a,
		ctx:                                 ctx,
		groupId:                             args.GroupId,
		name:                                args.Name,
		serverlessInstanceDescriptionUpdate: args.ServerlessInstanceDescriptionUpdate,
	}
}

func (r UpdateServerlessInstanceApiRequest) Execute() (*ServerlessInstanceDescription, *http.Response, error) {
	return r.ApiService.UpdateServerlessInstanceExecute(r)
}

/*
UpdateServerlessInstance Update One Serverless Instance in One Project

Updates one serverless instance in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the serverless instance.
	@return UpdateServerlessInstanceApiRequest
*/
func (a *ServerlessInstancesApiService) UpdateServerlessInstance(ctx context.Context, groupId string, name string, serverlessInstanceDescriptionUpdate *ServerlessInstanceDescriptionUpdate) UpdateServerlessInstanceApiRequest {
	return UpdateServerlessInstanceApiRequest{
		ApiService:                          a,
		ctx:                                 ctx,
		groupId:                             groupId,
		name:                                name,
		serverlessInstanceDescriptionUpdate: serverlessInstanceDescriptionUpdate,
	}
}

// Execute executes the request
//
//	@return ServerlessInstanceDescription
func (a *ServerlessInstancesApiService) UpdateServerlessInstanceExecute(r UpdateServerlessInstanceApiRequest) (*ServerlessInstanceDescription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServerlessInstanceDescription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessInstancesApiService.UpdateServerlessInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/serverless/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serverlessInstanceDescriptionUpdate == nil {
		return localVarReturnValue, nil, reportError("serverlessInstanceDescriptionUpdate is required and must be specified")
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
	localVarPostBody = r.serverlessInstanceDescriptionUpdate
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
