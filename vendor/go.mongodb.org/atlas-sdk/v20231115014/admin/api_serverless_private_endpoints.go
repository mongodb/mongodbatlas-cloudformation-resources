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

type ServerlessPrivateEndpointsApi interface {

	/*
			CreateServerlessPrivateEndpoint Create One Private Endpoint for One Serverless Instance

			[experimental] Creates one private endpoint for one serverless instance. To use this resource, the requesting API Key must have the Project Owner role.

		 A new endpoint won't be immediately available after creation.  Read the steps in the linked tutorial for detailed guidance.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param instanceName Human-readable label that identifies the serverless instance for which the tenant endpoint will be created.
			@return CreateServerlessPrivateEndpointApiRequest
	*/
	CreateServerlessPrivateEndpoint(ctx context.Context, groupId string, instanceName string, serverlessTenantCreateRequest *ServerlessTenantCreateRequest) CreateServerlessPrivateEndpointApiRequest
	/*
		CreateServerlessPrivateEndpoint Create One Private Endpoint for One Serverless Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateServerlessPrivateEndpointApiParams - Parameters for the request
		@return CreateServerlessPrivateEndpointApiRequest
	*/
	CreateServerlessPrivateEndpointWithParams(ctx context.Context, args *CreateServerlessPrivateEndpointApiParams) CreateServerlessPrivateEndpointApiRequest

	// Method available only for mocking purposes
	CreateServerlessPrivateEndpointExecute(r CreateServerlessPrivateEndpointApiRequest) (*ServerlessTenantEndpoint, *http.Response, error)

	/*
		DeleteServerlessPrivateEndpoint Remove One Private Endpoint for One Serverless Instance

		[experimental] Remove one private endpoint from one serverless instance. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param instanceName Human-readable label that identifies the serverless instance from which the tenant endpoint will be removed.
		@param endpointId Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be removed.
		@return DeleteServerlessPrivateEndpointApiRequest
	*/
	DeleteServerlessPrivateEndpoint(ctx context.Context, groupId string, instanceName string, endpointId string) DeleteServerlessPrivateEndpointApiRequest
	/*
		DeleteServerlessPrivateEndpoint Remove One Private Endpoint for One Serverless Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteServerlessPrivateEndpointApiParams - Parameters for the request
		@return DeleteServerlessPrivateEndpointApiRequest
	*/
	DeleteServerlessPrivateEndpointWithParams(ctx context.Context, args *DeleteServerlessPrivateEndpointApiParams) DeleteServerlessPrivateEndpointApiRequest

	// Method available only for mocking purposes
	DeleteServerlessPrivateEndpointExecute(r DeleteServerlessPrivateEndpointApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetServerlessPrivateEndpoint Return One Private Endpoint for One Serverless Instance

		[experimental] Return one private endpoint for one serverless instance. Identify this endpoint using its unique ID. You must have at least the Project Read Only role for the project to successfully call this resource.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param instanceName Human-readable label that identifies the serverless instance associated with the tenant endpoint.
		@param endpointId Unique 24-hexadecimal digit string that identifies the tenant endpoint.
		@return GetServerlessPrivateEndpointApiRequest
	*/
	GetServerlessPrivateEndpoint(ctx context.Context, groupId string, instanceName string, endpointId string) GetServerlessPrivateEndpointApiRequest
	/*
		GetServerlessPrivateEndpoint Return One Private Endpoint for One Serverless Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetServerlessPrivateEndpointApiParams - Parameters for the request
		@return GetServerlessPrivateEndpointApiRequest
	*/
	GetServerlessPrivateEndpointWithParams(ctx context.Context, args *GetServerlessPrivateEndpointApiParams) GetServerlessPrivateEndpointApiRequest

	// Method available only for mocking purposes
	GetServerlessPrivateEndpointExecute(r GetServerlessPrivateEndpointApiRequest) (*ServerlessTenantEndpoint, *http.Response, error)

	/*
		ListServerlessPrivateEndpoints Return All Private Endpoints for One Serverless Instance

		Returns all private endpoints for one serverless instance. You must have at least the Project Read Only role for the project to successfully call this resource.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param instanceName Human-readable label that identifies the serverless instance associated with the tenant endpoint.
		@return ListServerlessPrivateEndpointsApiRequest
	*/
	ListServerlessPrivateEndpoints(ctx context.Context, groupId string, instanceName string) ListServerlessPrivateEndpointsApiRequest
	/*
		ListServerlessPrivateEndpoints Return All Private Endpoints for One Serverless Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListServerlessPrivateEndpointsApiParams - Parameters for the request
		@return ListServerlessPrivateEndpointsApiRequest
	*/
	ListServerlessPrivateEndpointsWithParams(ctx context.Context, args *ListServerlessPrivateEndpointsApiParams) ListServerlessPrivateEndpointsApiRequest

	// Method available only for mocking purposes
	ListServerlessPrivateEndpointsExecute(r ListServerlessPrivateEndpointsApiRequest) ([]ServerlessTenantEndpoint, *http.Response, error)

	/*
		UpdateServerlessPrivateEndpoint Update One Private Endpoint for One Serverless Instance

		[experimental] Updates one private endpoint for one serverless instance. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param instanceName Human-readable label that identifies the serverless instance associated with the tenant endpoint that will be updated.
		@param endpointId Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be updated.
		@return UpdateServerlessPrivateEndpointApiRequest
	*/
	UpdateServerlessPrivateEndpoint(ctx context.Context, groupId string, instanceName string, endpointId string, serverlessTenantEndpointUpdate *ServerlessTenantEndpointUpdate) UpdateServerlessPrivateEndpointApiRequest
	/*
		UpdateServerlessPrivateEndpoint Update One Private Endpoint for One Serverless Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateServerlessPrivateEndpointApiParams - Parameters for the request
		@return UpdateServerlessPrivateEndpointApiRequest
	*/
	UpdateServerlessPrivateEndpointWithParams(ctx context.Context, args *UpdateServerlessPrivateEndpointApiParams) UpdateServerlessPrivateEndpointApiRequest

	// Method available only for mocking purposes
	UpdateServerlessPrivateEndpointExecute(r UpdateServerlessPrivateEndpointApiRequest) (*ServerlessTenantEndpoint, *http.Response, error)
}

// ServerlessPrivateEndpointsApiService ServerlessPrivateEndpointsApi service
type ServerlessPrivateEndpointsApiService service

type CreateServerlessPrivateEndpointApiRequest struct {
	ctx                           context.Context
	ApiService                    ServerlessPrivateEndpointsApi
	groupId                       string
	instanceName                  string
	serverlessTenantCreateRequest *ServerlessTenantCreateRequest
}

type CreateServerlessPrivateEndpointApiParams struct {
	GroupId                       string
	InstanceName                  string
	ServerlessTenantCreateRequest *ServerlessTenantCreateRequest
}

func (a *ServerlessPrivateEndpointsApiService) CreateServerlessPrivateEndpointWithParams(ctx context.Context, args *CreateServerlessPrivateEndpointApiParams) CreateServerlessPrivateEndpointApiRequest {
	return CreateServerlessPrivateEndpointApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		groupId:                       args.GroupId,
		instanceName:                  args.InstanceName,
		serverlessTenantCreateRequest: args.ServerlessTenantCreateRequest,
	}
}

func (r CreateServerlessPrivateEndpointApiRequest) Execute() (*ServerlessTenantEndpoint, *http.Response, error) {
	return r.ApiService.CreateServerlessPrivateEndpointExecute(r)
}

/*
CreateServerlessPrivateEndpoint Create One Private Endpoint for One Serverless Instance

[experimental] Creates one private endpoint for one serverless instance. To use this resource, the requesting API Key must have the Project Owner role.

	A new endpoint won't be immediately available after creation.  Read the steps in the linked tutorial for detailed guidance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param instanceName Human-readable label that identifies the serverless instance for which the tenant endpoint will be created.
	@return CreateServerlessPrivateEndpointApiRequest
*/
func (a *ServerlessPrivateEndpointsApiService) CreateServerlessPrivateEndpoint(ctx context.Context, groupId string, instanceName string, serverlessTenantCreateRequest *ServerlessTenantCreateRequest) CreateServerlessPrivateEndpointApiRequest {
	return CreateServerlessPrivateEndpointApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		groupId:                       groupId,
		instanceName:                  instanceName,
		serverlessTenantCreateRequest: serverlessTenantCreateRequest,
	}
}

// Execute executes the request
//
//	@return ServerlessTenantEndpoint
func (a *ServerlessPrivateEndpointsApiService) CreateServerlessPrivateEndpointExecute(r CreateServerlessPrivateEndpointApiRequest) (*ServerlessTenantEndpoint, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServerlessTenantEndpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessPrivateEndpointsApiService.CreateServerlessPrivateEndpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", url.PathEscape(parameterValueToString(r.instanceName, "instanceName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serverlessTenantCreateRequest == nil {
		return localVarReturnValue, nil, reportError("serverlessTenantCreateRequest is required and must be specified")
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
	localVarPostBody = r.serverlessTenantCreateRequest
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

type DeleteServerlessPrivateEndpointApiRequest struct {
	ctx          context.Context
	ApiService   ServerlessPrivateEndpointsApi
	groupId      string
	instanceName string
	endpointId   string
}

type DeleteServerlessPrivateEndpointApiParams struct {
	GroupId      string
	InstanceName string
	EndpointId   string
}

func (a *ServerlessPrivateEndpointsApiService) DeleteServerlessPrivateEndpointWithParams(ctx context.Context, args *DeleteServerlessPrivateEndpointApiParams) DeleteServerlessPrivateEndpointApiRequest {
	return DeleteServerlessPrivateEndpointApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		instanceName: args.InstanceName,
		endpointId:   args.EndpointId,
	}
}

func (r DeleteServerlessPrivateEndpointApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteServerlessPrivateEndpointExecute(r)
}

/*
DeleteServerlessPrivateEndpoint Remove One Private Endpoint for One Serverless Instance

[experimental] Remove one private endpoint from one serverless instance. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param instanceName Human-readable label that identifies the serverless instance from which the tenant endpoint will be removed.
	@param endpointId Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be removed.
	@return DeleteServerlessPrivateEndpointApiRequest
*/
func (a *ServerlessPrivateEndpointsApiService) DeleteServerlessPrivateEndpoint(ctx context.Context, groupId string, instanceName string, endpointId string) DeleteServerlessPrivateEndpointApiRequest {
	return DeleteServerlessPrivateEndpointApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		instanceName: instanceName,
		endpointId:   endpointId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *ServerlessPrivateEndpointsApiService) DeleteServerlessPrivateEndpointExecute(r DeleteServerlessPrivateEndpointApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessPrivateEndpointsApiService.DeleteServerlessPrivateEndpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", url.PathEscape(parameterValueToString(r.instanceName, "instanceName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(parameterValueToString(r.endpointId, "endpointId")), -1)

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

type GetServerlessPrivateEndpointApiRequest struct {
	ctx          context.Context
	ApiService   ServerlessPrivateEndpointsApi
	groupId      string
	instanceName string
	endpointId   string
}

type GetServerlessPrivateEndpointApiParams struct {
	GroupId      string
	InstanceName string
	EndpointId   string
}

func (a *ServerlessPrivateEndpointsApiService) GetServerlessPrivateEndpointWithParams(ctx context.Context, args *GetServerlessPrivateEndpointApiParams) GetServerlessPrivateEndpointApiRequest {
	return GetServerlessPrivateEndpointApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		instanceName: args.InstanceName,
		endpointId:   args.EndpointId,
	}
}

func (r GetServerlessPrivateEndpointApiRequest) Execute() (*ServerlessTenantEndpoint, *http.Response, error) {
	return r.ApiService.GetServerlessPrivateEndpointExecute(r)
}

/*
GetServerlessPrivateEndpoint Return One Private Endpoint for One Serverless Instance

[experimental] Return one private endpoint for one serverless instance. Identify this endpoint using its unique ID. You must have at least the Project Read Only role for the project to successfully call this resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param instanceName Human-readable label that identifies the serverless instance associated with the tenant endpoint.
	@param endpointId Unique 24-hexadecimal digit string that identifies the tenant endpoint.
	@return GetServerlessPrivateEndpointApiRequest
*/
func (a *ServerlessPrivateEndpointsApiService) GetServerlessPrivateEndpoint(ctx context.Context, groupId string, instanceName string, endpointId string) GetServerlessPrivateEndpointApiRequest {
	return GetServerlessPrivateEndpointApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		instanceName: instanceName,
		endpointId:   endpointId,
	}
}

// Execute executes the request
//
//	@return ServerlessTenantEndpoint
func (a *ServerlessPrivateEndpointsApiService) GetServerlessPrivateEndpointExecute(r GetServerlessPrivateEndpointApiRequest) (*ServerlessTenantEndpoint, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServerlessTenantEndpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessPrivateEndpointsApiService.GetServerlessPrivateEndpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", url.PathEscape(parameterValueToString(r.instanceName, "instanceName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(parameterValueToString(r.endpointId, "endpointId")), -1)

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

type ListServerlessPrivateEndpointsApiRequest struct {
	ctx          context.Context
	ApiService   ServerlessPrivateEndpointsApi
	groupId      string
	instanceName string
}

type ListServerlessPrivateEndpointsApiParams struct {
	GroupId      string
	InstanceName string
}

func (a *ServerlessPrivateEndpointsApiService) ListServerlessPrivateEndpointsWithParams(ctx context.Context, args *ListServerlessPrivateEndpointsApiParams) ListServerlessPrivateEndpointsApiRequest {
	return ListServerlessPrivateEndpointsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		instanceName: args.InstanceName,
	}
}

func (r ListServerlessPrivateEndpointsApiRequest) Execute() ([]ServerlessTenantEndpoint, *http.Response, error) {
	return r.ApiService.ListServerlessPrivateEndpointsExecute(r)
}

/*
ListServerlessPrivateEndpoints Return All Private Endpoints for One Serverless Instance

Returns all private endpoints for one serverless instance. You must have at least the Project Read Only role for the project to successfully call this resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param instanceName Human-readable label that identifies the serverless instance associated with the tenant endpoint.
	@return ListServerlessPrivateEndpointsApiRequest
*/
func (a *ServerlessPrivateEndpointsApiService) ListServerlessPrivateEndpoints(ctx context.Context, groupId string, instanceName string) ListServerlessPrivateEndpointsApiRequest {
	return ListServerlessPrivateEndpointsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		instanceName: instanceName,
	}
}

// Execute executes the request
//
//	@return []ServerlessTenantEndpoint
func (a *ServerlessPrivateEndpointsApiService) ListServerlessPrivateEndpointsExecute(r ListServerlessPrivateEndpointsApiRequest) ([]ServerlessTenantEndpoint, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ServerlessTenantEndpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessPrivateEndpointsApiService.ListServerlessPrivateEndpoints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", url.PathEscape(parameterValueToString(r.instanceName, "instanceName")), -1)

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

type UpdateServerlessPrivateEndpointApiRequest struct {
	ctx                            context.Context
	ApiService                     ServerlessPrivateEndpointsApi
	groupId                        string
	instanceName                   string
	endpointId                     string
	serverlessTenantEndpointUpdate *ServerlessTenantEndpointUpdate
}

type UpdateServerlessPrivateEndpointApiParams struct {
	GroupId                        string
	InstanceName                   string
	EndpointId                     string
	ServerlessTenantEndpointUpdate *ServerlessTenantEndpointUpdate
}

func (a *ServerlessPrivateEndpointsApiService) UpdateServerlessPrivateEndpointWithParams(ctx context.Context, args *UpdateServerlessPrivateEndpointApiParams) UpdateServerlessPrivateEndpointApiRequest {
	return UpdateServerlessPrivateEndpointApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        args.GroupId,
		instanceName:                   args.InstanceName,
		endpointId:                     args.EndpointId,
		serverlessTenantEndpointUpdate: args.ServerlessTenantEndpointUpdate,
	}
}

func (r UpdateServerlessPrivateEndpointApiRequest) Execute() (*ServerlessTenantEndpoint, *http.Response, error) {
	return r.ApiService.UpdateServerlessPrivateEndpointExecute(r)
}

/*
UpdateServerlessPrivateEndpoint Update One Private Endpoint for One Serverless Instance

[experimental] Updates one private endpoint for one serverless instance. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param instanceName Human-readable label that identifies the serverless instance associated with the tenant endpoint that will be updated.
	@param endpointId Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be updated.
	@return UpdateServerlessPrivateEndpointApiRequest
*/
func (a *ServerlessPrivateEndpointsApiService) UpdateServerlessPrivateEndpoint(ctx context.Context, groupId string, instanceName string, endpointId string, serverlessTenantEndpointUpdate *ServerlessTenantEndpointUpdate) UpdateServerlessPrivateEndpointApiRequest {
	return UpdateServerlessPrivateEndpointApiRequest{
		ApiService:                     a,
		ctx:                            ctx,
		groupId:                        groupId,
		instanceName:                   instanceName,
		endpointId:                     endpointId,
		serverlessTenantEndpointUpdate: serverlessTenantEndpointUpdate,
	}
}

// Execute executes the request
//
//	@return ServerlessTenantEndpoint
func (a *ServerlessPrivateEndpointsApiService) UpdateServerlessPrivateEndpointExecute(r UpdateServerlessPrivateEndpointApiRequest) (*ServerlessTenantEndpoint, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServerlessTenantEndpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerlessPrivateEndpointsApiService.UpdateServerlessPrivateEndpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/privateEndpoint/serverless/instance/{instanceName}/endpoint/{endpointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", url.PathEscape(parameterValueToString(r.instanceName, "instanceName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endpointId"+"}", url.PathEscape(parameterValueToString(r.endpointId, "endpointId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serverlessTenantEndpointUpdate == nil {
		return localVarReturnValue, nil, reportError("serverlessTenantEndpointUpdate is required and must be specified")
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
	localVarPostBody = r.serverlessTenantEndpointUpdate
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
