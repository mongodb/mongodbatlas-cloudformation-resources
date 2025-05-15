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

type CloudProviderAccessApi interface {

	/*
		AuthorizeCloudProviderAccessRole Authorize One Cloud Provider Access Role

		Grants access to the specified project for the specified access role. To use this resource, the requesting API Key must have the Project Owner role. This API endpoint is one step in a procedure to create unified access for MongoDB Cloud services.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param roleId Unique 24-hexadecimal digit string that identifies the role.
		@return AuthorizeCloudProviderAccessRoleApiRequest
	*/
	AuthorizeCloudProviderAccessRole(ctx context.Context, groupId string, roleId string, cloudProviderAccessRole *CloudProviderAccessRole) AuthorizeCloudProviderAccessRoleApiRequest
	/*
		AuthorizeCloudProviderAccessRole Authorize One Cloud Provider Access Role


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AuthorizeCloudProviderAccessRoleApiParams - Parameters for the request
		@return AuthorizeCloudProviderAccessRoleApiRequest
	*/
	AuthorizeCloudProviderAccessRoleWithParams(ctx context.Context, args *AuthorizeCloudProviderAccessRoleApiParams) AuthorizeCloudProviderAccessRoleApiRequest

	// Method available only for mocking purposes
	AuthorizeCloudProviderAccessRoleExecute(r AuthorizeCloudProviderAccessRoleApiRequest) (*CloudProviderAccessRole, *http.Response, error)

	/*
		CreateCloudProviderAccessRole Create One Cloud Provider Access Role

		Creates one access role for the specified cloud provider. Some MongoDB Cloud features use these cloud provider access roles for authentication. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return CreateCloudProviderAccessRoleApiRequest
	*/
	CreateCloudProviderAccessRole(ctx context.Context, groupId string, cloudProviderAccessRole *CloudProviderAccessRole) CreateCloudProviderAccessRoleApiRequest
	/*
		CreateCloudProviderAccessRole Create One Cloud Provider Access Role


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateCloudProviderAccessRoleApiParams - Parameters for the request
		@return CreateCloudProviderAccessRoleApiRequest
	*/
	CreateCloudProviderAccessRoleWithParams(ctx context.Context, args *CreateCloudProviderAccessRoleApiParams) CreateCloudProviderAccessRoleApiRequest

	// Method available only for mocking purposes
	CreateCloudProviderAccessRoleExecute(r CreateCloudProviderAccessRoleApiRequest) (*CloudProviderAccessRole, *http.Response, error)

	/*
		DeauthorizeCloudProviderAccessRole Deauthorize One Cloud Provider Access Role

		Revokes access to the specified project for the specified access role. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param cloudProvider Human-readable label that identifies the cloud provider of the role to deauthorize.
		@param roleId Unique 24-hexadecimal digit string that identifies the role.
		@return DeauthorizeCloudProviderAccessRoleApiRequest
	*/
	DeauthorizeCloudProviderAccessRole(ctx context.Context, groupId string, cloudProvider string, roleId string) DeauthorizeCloudProviderAccessRoleApiRequest
	/*
		DeauthorizeCloudProviderAccessRole Deauthorize One Cloud Provider Access Role


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeauthorizeCloudProviderAccessRoleApiParams - Parameters for the request
		@return DeauthorizeCloudProviderAccessRoleApiRequest
	*/
	DeauthorizeCloudProviderAccessRoleWithParams(ctx context.Context, args *DeauthorizeCloudProviderAccessRoleApiParams) DeauthorizeCloudProviderAccessRoleApiRequest

	// Method available only for mocking purposes
	DeauthorizeCloudProviderAccessRoleExecute(r DeauthorizeCloudProviderAccessRoleApiRequest) (*http.Response, error)

	/*
		GetCloudProviderAccessRole Return specified Cloud Provider Access Role

		[experimental] Returns the access role with the specified id and with access to the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param roleId Unique 24-hexadecimal digit string that identifies the role.
		@return GetCloudProviderAccessRoleApiRequest
	*/
	GetCloudProviderAccessRole(ctx context.Context, groupId string, roleId string) GetCloudProviderAccessRoleApiRequest
	/*
		GetCloudProviderAccessRole Return specified Cloud Provider Access Role


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetCloudProviderAccessRoleApiParams - Parameters for the request
		@return GetCloudProviderAccessRoleApiRequest
	*/
	GetCloudProviderAccessRoleWithParams(ctx context.Context, args *GetCloudProviderAccessRoleApiParams) GetCloudProviderAccessRoleApiRequest

	// Method available only for mocking purposes
	GetCloudProviderAccessRoleExecute(r GetCloudProviderAccessRoleApiRequest) (*CloudProviderAccessRole, *http.Response, error)

	/*
		ListCloudProviderAccessRoles Return All Cloud Provider Access Roles

		Returns all cloud provider access roles with access to the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListCloudProviderAccessRolesApiRequest
	*/
	ListCloudProviderAccessRoles(ctx context.Context, groupId string) ListCloudProviderAccessRolesApiRequest
	/*
		ListCloudProviderAccessRoles Return All Cloud Provider Access Roles


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListCloudProviderAccessRolesApiParams - Parameters for the request
		@return ListCloudProviderAccessRolesApiRequest
	*/
	ListCloudProviderAccessRolesWithParams(ctx context.Context, args *ListCloudProviderAccessRolesApiParams) ListCloudProviderAccessRolesApiRequest

	// Method available only for mocking purposes
	ListCloudProviderAccessRolesExecute(r ListCloudProviderAccessRolesApiRequest) (*CloudProviderAccessRoles, *http.Response, error)
}

// CloudProviderAccessApiService CloudProviderAccessApi service
type CloudProviderAccessApiService service

type AuthorizeCloudProviderAccessRoleApiRequest struct {
	ctx                     context.Context
	ApiService              CloudProviderAccessApi
	groupId                 string
	roleId                  string
	cloudProviderAccessRole *CloudProviderAccessRole
}

type AuthorizeCloudProviderAccessRoleApiParams struct {
	GroupId                 string
	RoleId                  string
	CloudProviderAccessRole *CloudProviderAccessRole
}

func (a *CloudProviderAccessApiService) AuthorizeCloudProviderAccessRoleWithParams(ctx context.Context, args *AuthorizeCloudProviderAccessRoleApiParams) AuthorizeCloudProviderAccessRoleApiRequest {
	return AuthorizeCloudProviderAccessRoleApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 args.GroupId,
		roleId:                  args.RoleId,
		cloudProviderAccessRole: args.CloudProviderAccessRole,
	}
}

func (r AuthorizeCloudProviderAccessRoleApiRequest) Execute() (*CloudProviderAccessRole, *http.Response, error) {
	return r.ApiService.AuthorizeCloudProviderAccessRoleExecute(r)
}

/*
AuthorizeCloudProviderAccessRole Authorize One Cloud Provider Access Role

Grants access to the specified project for the specified access role. To use this resource, the requesting API Key must have the Project Owner role. This API endpoint is one step in a procedure to create unified access for MongoDB Cloud services.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param roleId Unique 24-hexadecimal digit string that identifies the role.
	@return AuthorizeCloudProviderAccessRoleApiRequest
*/
func (a *CloudProviderAccessApiService) AuthorizeCloudProviderAccessRole(ctx context.Context, groupId string, roleId string, cloudProviderAccessRole *CloudProviderAccessRole) AuthorizeCloudProviderAccessRoleApiRequest {
	return AuthorizeCloudProviderAccessRoleApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 groupId,
		roleId:                  roleId,
		cloudProviderAccessRole: cloudProviderAccessRole,
	}
}

// Execute executes the request
//
//	@return CloudProviderAccessRole
func (a *CloudProviderAccessApiService) AuthorizeCloudProviderAccessRoleExecute(r AuthorizeCloudProviderAccessRoleApiRequest) (*CloudProviderAccessRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CloudProviderAccessRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudProviderAccessApiService.AuthorizeCloudProviderAccessRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudProviderAccessRole == nil {
		return localVarReturnValue, nil, reportError("cloudProviderAccessRole is required and must be specified")
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
	localVarPostBody = r.cloudProviderAccessRole
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

type CreateCloudProviderAccessRoleApiRequest struct {
	ctx                     context.Context
	ApiService              CloudProviderAccessApi
	groupId                 string
	cloudProviderAccessRole *CloudProviderAccessRole
}

type CreateCloudProviderAccessRoleApiParams struct {
	GroupId                 string
	CloudProviderAccessRole *CloudProviderAccessRole
}

func (a *CloudProviderAccessApiService) CreateCloudProviderAccessRoleWithParams(ctx context.Context, args *CreateCloudProviderAccessRoleApiParams) CreateCloudProviderAccessRoleApiRequest {
	return CreateCloudProviderAccessRoleApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 args.GroupId,
		cloudProviderAccessRole: args.CloudProviderAccessRole,
	}
}

func (r CreateCloudProviderAccessRoleApiRequest) Execute() (*CloudProviderAccessRole, *http.Response, error) {
	return r.ApiService.CreateCloudProviderAccessRoleExecute(r)
}

/*
CreateCloudProviderAccessRole Create One Cloud Provider Access Role

Creates one access role for the specified cloud provider. Some MongoDB Cloud features use these cloud provider access roles for authentication. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateCloudProviderAccessRoleApiRequest
*/
func (a *CloudProviderAccessApiService) CreateCloudProviderAccessRole(ctx context.Context, groupId string, cloudProviderAccessRole *CloudProviderAccessRole) CreateCloudProviderAccessRoleApiRequest {
	return CreateCloudProviderAccessRoleApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 groupId,
		cloudProviderAccessRole: cloudProviderAccessRole,
	}
}

// Execute executes the request
//
//	@return CloudProviderAccessRole
func (a *CloudProviderAccessApiService) CreateCloudProviderAccessRoleExecute(r CreateCloudProviderAccessRoleApiRequest) (*CloudProviderAccessRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CloudProviderAccessRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudProviderAccessApiService.CreateCloudProviderAccessRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/cloudProviderAccess"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudProviderAccessRole == nil {
		return localVarReturnValue, nil, reportError("cloudProviderAccessRole is required and must be specified")
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
	localVarPostBody = r.cloudProviderAccessRole
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

type DeauthorizeCloudProviderAccessRoleApiRequest struct {
	ctx           context.Context
	ApiService    CloudProviderAccessApi
	groupId       string
	cloudProvider string
	roleId        string
}

type DeauthorizeCloudProviderAccessRoleApiParams struct {
	GroupId       string
	CloudProvider string
	RoleId        string
}

func (a *CloudProviderAccessApiService) DeauthorizeCloudProviderAccessRoleWithParams(ctx context.Context, args *DeauthorizeCloudProviderAccessRoleApiParams) DeauthorizeCloudProviderAccessRoleApiRequest {
	return DeauthorizeCloudProviderAccessRoleApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		cloudProvider: args.CloudProvider,
		roleId:        args.RoleId,
	}
}

func (r DeauthorizeCloudProviderAccessRoleApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeauthorizeCloudProviderAccessRoleExecute(r)
}

/*
DeauthorizeCloudProviderAccessRole Deauthorize One Cloud Provider Access Role

Revokes access to the specified project for the specified access role. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param cloudProvider Human-readable label that identifies the cloud provider of the role to deauthorize.
	@param roleId Unique 24-hexadecimal digit string that identifies the role.
	@return DeauthorizeCloudProviderAccessRoleApiRequest
*/
func (a *CloudProviderAccessApiService) DeauthorizeCloudProviderAccessRole(ctx context.Context, groupId string, cloudProvider string, roleId string) DeauthorizeCloudProviderAccessRoleApiRequest {
	return DeauthorizeCloudProviderAccessRoleApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		cloudProvider: cloudProvider,
		roleId:        roleId,
	}
}

// Execute executes the request
func (a *CloudProviderAccessApiService) DeauthorizeCloudProviderAccessRoleExecute(r DeauthorizeCloudProviderAccessRoleApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudProviderAccessApiService.DeauthorizeCloudProviderAccessRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/cloudProviderAccess/{cloudProvider}/{roleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cloudProvider"+"}", url.PathEscape(parameterValueToString(r.cloudProvider, "cloudProvider")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

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

type GetCloudProviderAccessRoleApiRequest struct {
	ctx        context.Context
	ApiService CloudProviderAccessApi
	groupId    string
	roleId     string
}

type GetCloudProviderAccessRoleApiParams struct {
	GroupId string
	RoleId  string
}

func (a *CloudProviderAccessApiService) GetCloudProviderAccessRoleWithParams(ctx context.Context, args *GetCloudProviderAccessRoleApiParams) GetCloudProviderAccessRoleApiRequest {
	return GetCloudProviderAccessRoleApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		roleId:     args.RoleId,
	}
}

func (r GetCloudProviderAccessRoleApiRequest) Execute() (*CloudProviderAccessRole, *http.Response, error) {
	return r.ApiService.GetCloudProviderAccessRoleExecute(r)
}

/*
GetCloudProviderAccessRole Return specified Cloud Provider Access Role

[experimental] Returns the access role with the specified id and with access to the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param roleId Unique 24-hexadecimal digit string that identifies the role.
	@return GetCloudProviderAccessRoleApiRequest
*/
func (a *CloudProviderAccessApiService) GetCloudProviderAccessRole(ctx context.Context, groupId string, roleId string) GetCloudProviderAccessRoleApiRequest {
	return GetCloudProviderAccessRoleApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		roleId:     roleId,
	}
}

// Execute executes the request
//
//	@return CloudProviderAccessRole
func (a *CloudProviderAccessApiService) GetCloudProviderAccessRoleExecute(r GetCloudProviderAccessRoleApiRequest) (*CloudProviderAccessRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CloudProviderAccessRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudProviderAccessApiService.GetCloudProviderAccessRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/cloudProviderAccess/{roleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

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

type ListCloudProviderAccessRolesApiRequest struct {
	ctx        context.Context
	ApiService CloudProviderAccessApi
	groupId    string
}

type ListCloudProviderAccessRolesApiParams struct {
	GroupId string
}

func (a *CloudProviderAccessApiService) ListCloudProviderAccessRolesWithParams(ctx context.Context, args *ListCloudProviderAccessRolesApiParams) ListCloudProviderAccessRolesApiRequest {
	return ListCloudProviderAccessRolesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r ListCloudProviderAccessRolesApiRequest) Execute() (*CloudProviderAccessRoles, *http.Response, error) {
	return r.ApiService.ListCloudProviderAccessRolesExecute(r)
}

/*
ListCloudProviderAccessRoles Return All Cloud Provider Access Roles

Returns all cloud provider access roles with access to the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListCloudProviderAccessRolesApiRequest
*/
func (a *CloudProviderAccessApiService) ListCloudProviderAccessRoles(ctx context.Context, groupId string) ListCloudProviderAccessRolesApiRequest {
	return ListCloudProviderAccessRolesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return CloudProviderAccessRoles
func (a *CloudProviderAccessApiService) ListCloudProviderAccessRolesExecute(r ListCloudProviderAccessRolesApiRequest) (*CloudProviderAccessRoles, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CloudProviderAccessRoles
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudProviderAccessApiService.ListCloudProviderAccessRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/cloudProviderAccess"
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
