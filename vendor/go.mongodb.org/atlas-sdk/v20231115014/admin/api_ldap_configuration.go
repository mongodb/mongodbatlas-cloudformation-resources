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

type LDAPConfigurationApi interface {

	/*
		DeleteLDAPConfiguration Remove the Current LDAP User to DN Mapping

		Removes the current LDAP Distinguished Name mapping captured in the ``userToDNMapping`` document from the LDAP configuration for the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeleteLDAPConfigurationApiRequest
	*/
	DeleteLDAPConfiguration(ctx context.Context, groupId string) DeleteLDAPConfigurationApiRequest
	/*
		DeleteLDAPConfiguration Remove the Current LDAP User to DN Mapping


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteLDAPConfigurationApiParams - Parameters for the request
		@return DeleteLDAPConfigurationApiRequest
	*/
	DeleteLDAPConfigurationWithParams(ctx context.Context, args *DeleteLDAPConfigurationApiParams) DeleteLDAPConfigurationApiRequest

	// Method available only for mocking purposes
	DeleteLDAPConfigurationExecute(r DeleteLDAPConfigurationApiRequest) (*UserSecurity, *http.Response, error)

	/*
		GetLDAPConfiguration Return the Current LDAP or X.509 Configuration

		Returns the current LDAP configuration for the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetLDAPConfigurationApiRequest
	*/
	GetLDAPConfiguration(ctx context.Context, groupId string) GetLDAPConfigurationApiRequest
	/*
		GetLDAPConfiguration Return the Current LDAP or X.509 Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetLDAPConfigurationApiParams - Parameters for the request
		@return GetLDAPConfigurationApiRequest
	*/
	GetLDAPConfigurationWithParams(ctx context.Context, args *GetLDAPConfigurationApiParams) GetLDAPConfigurationApiRequest

	// Method available only for mocking purposes
	GetLDAPConfigurationExecute(r GetLDAPConfigurationApiRequest) (*UserSecurity, *http.Response, error)

	/*
		GetLDAPConfigurationStatus Return the Status of One Verify LDAP Configuration Request

		Returns the status of one request to verify one LDAP configuration for the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param requestId Unique string that identifies the request to verify an <abbr title=\"Lightweight Directory Access Protocol\">LDAP</abbr> configuration.
		@return GetLDAPConfigurationStatusApiRequest
	*/
	GetLDAPConfigurationStatus(ctx context.Context, groupId string, requestId string) GetLDAPConfigurationStatusApiRequest
	/*
		GetLDAPConfigurationStatus Return the Status of One Verify LDAP Configuration Request


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetLDAPConfigurationStatusApiParams - Parameters for the request
		@return GetLDAPConfigurationStatusApiRequest
	*/
	GetLDAPConfigurationStatusWithParams(ctx context.Context, args *GetLDAPConfigurationStatusApiParams) GetLDAPConfigurationStatusApiRequest

	// Method available only for mocking purposes
	GetLDAPConfigurationStatusExecute(r GetLDAPConfigurationStatusApiRequest) (*LDAPVerifyConnectivityJobRequest, *http.Response, error)

	/*
			SaveLDAPConfiguration Edit the LDAP or X.509 Configuration

			Edits the LDAP configuration for the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		Updating this configuration triggers a rolling restart of the database.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return SaveLDAPConfigurationApiRequest
	*/
	SaveLDAPConfiguration(ctx context.Context, groupId string, userSecurity *UserSecurity) SaveLDAPConfigurationApiRequest
	/*
		SaveLDAPConfiguration Edit the LDAP or X.509 Configuration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param SaveLDAPConfigurationApiParams - Parameters for the request
		@return SaveLDAPConfigurationApiRequest
	*/
	SaveLDAPConfigurationWithParams(ctx context.Context, args *SaveLDAPConfigurationApiParams) SaveLDAPConfigurationApiRequest

	// Method available only for mocking purposes
	SaveLDAPConfigurationExecute(r SaveLDAPConfigurationApiRequest) (*UserSecurity, *http.Response, error)

	/*
		VerifyLDAPConfiguration Verify the LDAP Configuration in One Project

		Verifies the LDAP configuration for the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return VerifyLDAPConfigurationApiRequest
	*/
	VerifyLDAPConfiguration(ctx context.Context, groupId string, lDAPVerifyConnectivityJobRequestParams *LDAPVerifyConnectivityJobRequestParams) VerifyLDAPConfigurationApiRequest
	/*
		VerifyLDAPConfiguration Verify the LDAP Configuration in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param VerifyLDAPConfigurationApiParams - Parameters for the request
		@return VerifyLDAPConfigurationApiRequest
	*/
	VerifyLDAPConfigurationWithParams(ctx context.Context, args *VerifyLDAPConfigurationApiParams) VerifyLDAPConfigurationApiRequest

	// Method available only for mocking purposes
	VerifyLDAPConfigurationExecute(r VerifyLDAPConfigurationApiRequest) (*LDAPVerifyConnectivityJobRequest, *http.Response, error)
}

// LDAPConfigurationApiService LDAPConfigurationApi service
type LDAPConfigurationApiService service

type DeleteLDAPConfigurationApiRequest struct {
	ctx        context.Context
	ApiService LDAPConfigurationApi
	groupId    string
}

type DeleteLDAPConfigurationApiParams struct {
	GroupId string
}

func (a *LDAPConfigurationApiService) DeleteLDAPConfigurationWithParams(ctx context.Context, args *DeleteLDAPConfigurationApiParams) DeleteLDAPConfigurationApiRequest {
	return DeleteLDAPConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r DeleteLDAPConfigurationApiRequest) Execute() (*UserSecurity, *http.Response, error) {
	return r.ApiService.DeleteLDAPConfigurationExecute(r)
}

/*
DeleteLDAPConfiguration Remove the Current LDAP User to DN Mapping

Removes the current LDAP Distinguished Name mapping captured in the “userToDNMapping“ document from the LDAP configuration for the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeleteLDAPConfigurationApiRequest
*/
func (a *LDAPConfigurationApiService) DeleteLDAPConfiguration(ctx context.Context, groupId string) DeleteLDAPConfigurationApiRequest {
	return DeleteLDAPConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return UserSecurity
func (a *LDAPConfigurationApiService) DeleteLDAPConfigurationExecute(r DeleteLDAPConfigurationApiRequest) (*UserSecurity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserSecurity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPConfigurationApiService.DeleteLDAPConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/userSecurity/ldap/userToDNMapping"
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

type GetLDAPConfigurationApiRequest struct {
	ctx        context.Context
	ApiService LDAPConfigurationApi
	groupId    string
}

type GetLDAPConfigurationApiParams struct {
	GroupId string
}

func (a *LDAPConfigurationApiService) GetLDAPConfigurationWithParams(ctx context.Context, args *GetLDAPConfigurationApiParams) GetLDAPConfigurationApiRequest {
	return GetLDAPConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetLDAPConfigurationApiRequest) Execute() (*UserSecurity, *http.Response, error) {
	return r.ApiService.GetLDAPConfigurationExecute(r)
}

/*
GetLDAPConfiguration Return the Current LDAP or X.509 Configuration

Returns the current LDAP configuration for the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetLDAPConfigurationApiRequest
*/
func (a *LDAPConfigurationApiService) GetLDAPConfiguration(ctx context.Context, groupId string) GetLDAPConfigurationApiRequest {
	return GetLDAPConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return UserSecurity
func (a *LDAPConfigurationApiService) GetLDAPConfigurationExecute(r GetLDAPConfigurationApiRequest) (*UserSecurity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserSecurity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPConfigurationApiService.GetLDAPConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/userSecurity"
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

type GetLDAPConfigurationStatusApiRequest struct {
	ctx        context.Context
	ApiService LDAPConfigurationApi
	groupId    string
	requestId  string
}

type GetLDAPConfigurationStatusApiParams struct {
	GroupId   string
	RequestId string
}

func (a *LDAPConfigurationApiService) GetLDAPConfigurationStatusWithParams(ctx context.Context, args *GetLDAPConfigurationStatusApiParams) GetLDAPConfigurationStatusApiRequest {
	return GetLDAPConfigurationStatusApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		requestId:  args.RequestId,
	}
}

func (r GetLDAPConfigurationStatusApiRequest) Execute() (*LDAPVerifyConnectivityJobRequest, *http.Response, error) {
	return r.ApiService.GetLDAPConfigurationStatusExecute(r)
}

/*
GetLDAPConfigurationStatus Return the Status of One Verify LDAP Configuration Request

Returns the status of one request to verify one LDAP configuration for the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param requestId Unique string that identifies the request to verify an <abbr title=\"Lightweight Directory Access Protocol\">LDAP</abbr> configuration.
	@return GetLDAPConfigurationStatusApiRequest
*/
func (a *LDAPConfigurationApiService) GetLDAPConfigurationStatus(ctx context.Context, groupId string, requestId string) GetLDAPConfigurationStatusApiRequest {
	return GetLDAPConfigurationStatusApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		requestId:  requestId,
	}
}

// Execute executes the request
//
//	@return LDAPVerifyConnectivityJobRequest
func (a *LDAPConfigurationApiService) GetLDAPConfigurationStatusExecute(r GetLDAPConfigurationStatusApiRequest) (*LDAPVerifyConnectivityJobRequest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LDAPVerifyConnectivityJobRequest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPConfigurationApiService.GetLDAPConfigurationStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

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

type SaveLDAPConfigurationApiRequest struct {
	ctx          context.Context
	ApiService   LDAPConfigurationApi
	groupId      string
	userSecurity *UserSecurity
}

type SaveLDAPConfigurationApiParams struct {
	GroupId      string
	UserSecurity *UserSecurity
}

func (a *LDAPConfigurationApiService) SaveLDAPConfigurationWithParams(ctx context.Context, args *SaveLDAPConfigurationApiParams) SaveLDAPConfigurationApiRequest {
	return SaveLDAPConfigurationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		userSecurity: args.UserSecurity,
	}
}

func (r SaveLDAPConfigurationApiRequest) Execute() (*UserSecurity, *http.Response, error) {
	return r.ApiService.SaveLDAPConfigurationExecute(r)
}

/*
SaveLDAPConfiguration Edit the LDAP or X.509 Configuration

Edits the LDAP configuration for the specified project. To use this resource, the requesting API Key must have the Project Owner role.

Updating this configuration triggers a rolling restart of the database.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return SaveLDAPConfigurationApiRequest
*/
func (a *LDAPConfigurationApiService) SaveLDAPConfiguration(ctx context.Context, groupId string, userSecurity *UserSecurity) SaveLDAPConfigurationApiRequest {
	return SaveLDAPConfigurationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		userSecurity: userSecurity,
	}
}

// Execute executes the request
//
//	@return UserSecurity
func (a *LDAPConfigurationApiService) SaveLDAPConfigurationExecute(r SaveLDAPConfigurationApiRequest) (*UserSecurity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserSecurity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPConfigurationApiService.SaveLDAPConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/userSecurity"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userSecurity == nil {
		return localVarReturnValue, nil, reportError("userSecurity is required and must be specified")
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
	localVarPostBody = r.userSecurity
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

type VerifyLDAPConfigurationApiRequest struct {
	ctx                                    context.Context
	ApiService                             LDAPConfigurationApi
	groupId                                string
	lDAPVerifyConnectivityJobRequestParams *LDAPVerifyConnectivityJobRequestParams
}

type VerifyLDAPConfigurationApiParams struct {
	GroupId                                string
	LDAPVerifyConnectivityJobRequestParams *LDAPVerifyConnectivityJobRequestParams
}

func (a *LDAPConfigurationApiService) VerifyLDAPConfigurationWithParams(ctx context.Context, args *VerifyLDAPConfigurationApiParams) VerifyLDAPConfigurationApiRequest {
	return VerifyLDAPConfigurationApiRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		groupId:                                args.GroupId,
		lDAPVerifyConnectivityJobRequestParams: args.LDAPVerifyConnectivityJobRequestParams,
	}
}

func (r VerifyLDAPConfigurationApiRequest) Execute() (*LDAPVerifyConnectivityJobRequest, *http.Response, error) {
	return r.ApiService.VerifyLDAPConfigurationExecute(r)
}

/*
VerifyLDAPConfiguration Verify the LDAP Configuration in One Project

Verifies the LDAP configuration for the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return VerifyLDAPConfigurationApiRequest
*/
func (a *LDAPConfigurationApiService) VerifyLDAPConfiguration(ctx context.Context, groupId string, lDAPVerifyConnectivityJobRequestParams *LDAPVerifyConnectivityJobRequestParams) VerifyLDAPConfigurationApiRequest {
	return VerifyLDAPConfigurationApiRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		groupId:                                groupId,
		lDAPVerifyConnectivityJobRequestParams: lDAPVerifyConnectivityJobRequestParams,
	}
}

// Execute executes the request
//
//	@return LDAPVerifyConnectivityJobRequest
func (a *LDAPConfigurationApiService) VerifyLDAPConfigurationExecute(r VerifyLDAPConfigurationApiRequest) (*LDAPVerifyConnectivityJobRequest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LDAPVerifyConnectivityJobRequest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LDAPConfigurationApiService.VerifyLDAPConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/userSecurity/ldap/verify"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPVerifyConnectivityJobRequestParams == nil {
		return localVarReturnValue, nil, reportError("lDAPVerifyConnectivityJobRequestParams is required and must be specified")
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
	localVarPostBody = r.lDAPVerifyConnectivityJobRequestParams
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
