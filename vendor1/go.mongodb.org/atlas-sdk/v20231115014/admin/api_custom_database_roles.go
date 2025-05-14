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

type CustomDatabaseRolesApi interface {

	/*
		CreateCustomDatabaseRole Create One Custom Role

		Creates one custom role in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return CreateCustomDatabaseRoleApiRequest
	*/
	CreateCustomDatabaseRole(ctx context.Context, groupId string, userCustomDBRole *UserCustomDBRole) CreateCustomDatabaseRoleApiRequest
	/*
		CreateCustomDatabaseRole Create One Custom Role


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateCustomDatabaseRoleApiParams - Parameters for the request
		@return CreateCustomDatabaseRoleApiRequest
	*/
	CreateCustomDatabaseRoleWithParams(ctx context.Context, args *CreateCustomDatabaseRoleApiParams) CreateCustomDatabaseRoleApiRequest

	// Method available only for mocking purposes
	CreateCustomDatabaseRoleExecute(r CreateCustomDatabaseRoleApiRequest) (*UserCustomDBRole, *http.Response, error)

	/*
		DeleteCustomDatabaseRole Remove One Custom Role from One Project

		Removes one custom role from the specified project. You can't remove a custom role that would leave one or more child roles with no parent roles or actions. You also can't remove a custom role that would leave one or more database users without roles. To use this resource, the requesting API Key must have the Project Atlas Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param roleName Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.
		@return DeleteCustomDatabaseRoleApiRequest
	*/
	DeleteCustomDatabaseRole(ctx context.Context, groupId string, roleName string) DeleteCustomDatabaseRoleApiRequest
	/*
		DeleteCustomDatabaseRole Remove One Custom Role from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteCustomDatabaseRoleApiParams - Parameters for the request
		@return DeleteCustomDatabaseRoleApiRequest
	*/
	DeleteCustomDatabaseRoleWithParams(ctx context.Context, args *DeleteCustomDatabaseRoleApiParams) DeleteCustomDatabaseRoleApiRequest

	// Method available only for mocking purposes
	DeleteCustomDatabaseRoleExecute(r DeleteCustomDatabaseRoleApiRequest) (*http.Response, error)

	/*
		GetCustomDatabaseRole Return One Custom Role in One Project

		Returns one custom role for the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param roleName Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.
		@return GetCustomDatabaseRoleApiRequest
	*/
	GetCustomDatabaseRole(ctx context.Context, groupId string, roleName string) GetCustomDatabaseRoleApiRequest
	/*
		GetCustomDatabaseRole Return One Custom Role in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetCustomDatabaseRoleApiParams - Parameters for the request
		@return GetCustomDatabaseRoleApiRequest
	*/
	GetCustomDatabaseRoleWithParams(ctx context.Context, args *GetCustomDatabaseRoleApiParams) GetCustomDatabaseRoleApiRequest

	// Method available only for mocking purposes
	GetCustomDatabaseRoleExecute(r GetCustomDatabaseRoleApiRequest) (*UserCustomDBRole, *http.Response, error)

	/*
		ListCustomDatabaseRoles Return All Custom Roles in One Project

		Returns all custom roles for the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListCustomDatabaseRolesApiRequest
	*/
	ListCustomDatabaseRoles(ctx context.Context, groupId string) ListCustomDatabaseRolesApiRequest
	/*
		ListCustomDatabaseRoles Return All Custom Roles in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListCustomDatabaseRolesApiParams - Parameters for the request
		@return ListCustomDatabaseRolesApiRequest
	*/
	ListCustomDatabaseRolesWithParams(ctx context.Context, args *ListCustomDatabaseRolesApiParams) ListCustomDatabaseRolesApiRequest

	// Method available only for mocking purposes
	ListCustomDatabaseRolesExecute(r ListCustomDatabaseRolesApiRequest) ([]UserCustomDBRole, *http.Response, error)

	/*
		UpdateCustomDatabaseRole Update One Custom Role in One Project

		Updates one custom role in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param roleName Human-readable label that identifies the role for the request. This name must beunique for this custom role in this project.
		@return UpdateCustomDatabaseRoleApiRequest
	*/
	UpdateCustomDatabaseRole(ctx context.Context, groupId string, roleName string, updateCustomDBRole *UpdateCustomDBRole) UpdateCustomDatabaseRoleApiRequest
	/*
		UpdateCustomDatabaseRole Update One Custom Role in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateCustomDatabaseRoleApiParams - Parameters for the request
		@return UpdateCustomDatabaseRoleApiRequest
	*/
	UpdateCustomDatabaseRoleWithParams(ctx context.Context, args *UpdateCustomDatabaseRoleApiParams) UpdateCustomDatabaseRoleApiRequest

	// Method available only for mocking purposes
	UpdateCustomDatabaseRoleExecute(r UpdateCustomDatabaseRoleApiRequest) (*UserCustomDBRole, *http.Response, error)
}

// CustomDatabaseRolesApiService CustomDatabaseRolesApi service
type CustomDatabaseRolesApiService service

type CreateCustomDatabaseRoleApiRequest struct {
	ctx              context.Context
	ApiService       CustomDatabaseRolesApi
	groupId          string
	userCustomDBRole *UserCustomDBRole
}

type CreateCustomDatabaseRoleApiParams struct {
	GroupId          string
	UserCustomDBRole *UserCustomDBRole
}

func (a *CustomDatabaseRolesApiService) CreateCustomDatabaseRoleWithParams(ctx context.Context, args *CreateCustomDatabaseRoleApiParams) CreateCustomDatabaseRoleApiRequest {
	return CreateCustomDatabaseRoleApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          args.GroupId,
		userCustomDBRole: args.UserCustomDBRole,
	}
}

func (r CreateCustomDatabaseRoleApiRequest) Execute() (*UserCustomDBRole, *http.Response, error) {
	return r.ApiService.CreateCustomDatabaseRoleExecute(r)
}

/*
CreateCustomDatabaseRole Create One Custom Role

Creates one custom role in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateCustomDatabaseRoleApiRequest
*/
func (a *CustomDatabaseRolesApiService) CreateCustomDatabaseRole(ctx context.Context, groupId string, userCustomDBRole *UserCustomDBRole) CreateCustomDatabaseRoleApiRequest {
	return CreateCustomDatabaseRoleApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		userCustomDBRole: userCustomDBRole,
	}
}

// Execute executes the request
//
//	@return UserCustomDBRole
func (a *CustomDatabaseRolesApiService) CreateCustomDatabaseRoleExecute(r CreateCustomDatabaseRoleApiRequest) (*UserCustomDBRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserCustomDBRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDatabaseRolesApiService.CreateCustomDatabaseRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/customDBRoles/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userCustomDBRole == nil {
		return localVarReturnValue, nil, reportError("userCustomDBRole is required and must be specified")
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
	localVarPostBody = r.userCustomDBRole
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

type DeleteCustomDatabaseRoleApiRequest struct {
	ctx        context.Context
	ApiService CustomDatabaseRolesApi
	groupId    string
	roleName   string
}

type DeleteCustomDatabaseRoleApiParams struct {
	GroupId  string
	RoleName string
}

func (a *CustomDatabaseRolesApiService) DeleteCustomDatabaseRoleWithParams(ctx context.Context, args *DeleteCustomDatabaseRoleApiParams) DeleteCustomDatabaseRoleApiRequest {
	return DeleteCustomDatabaseRoleApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		roleName:   args.RoleName,
	}
}

func (r DeleteCustomDatabaseRoleApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCustomDatabaseRoleExecute(r)
}

/*
DeleteCustomDatabaseRole Remove One Custom Role from One Project

Removes one custom role from the specified project. You can't remove a custom role that would leave one or more child roles with no parent roles or actions. You also can't remove a custom role that would leave one or more database users without roles. To use this resource, the requesting API Key must have the Project Atlas Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param roleName Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.
	@return DeleteCustomDatabaseRoleApiRequest
*/
func (a *CustomDatabaseRolesApiService) DeleteCustomDatabaseRole(ctx context.Context, groupId string, roleName string) DeleteCustomDatabaseRoleApiRequest {
	return DeleteCustomDatabaseRoleApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		roleName:   roleName,
	}
}

// Execute executes the request
func (a *CustomDatabaseRolesApiService) DeleteCustomDatabaseRoleExecute(r DeleteCustomDatabaseRoleApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDatabaseRolesApiService.DeleteCustomDatabaseRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleName"+"}", url.PathEscape(parameterValueToString(r.roleName, "roleName")), -1)

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

type GetCustomDatabaseRoleApiRequest struct {
	ctx        context.Context
	ApiService CustomDatabaseRolesApi
	groupId    string
	roleName   string
}

type GetCustomDatabaseRoleApiParams struct {
	GroupId  string
	RoleName string
}

func (a *CustomDatabaseRolesApiService) GetCustomDatabaseRoleWithParams(ctx context.Context, args *GetCustomDatabaseRoleApiParams) GetCustomDatabaseRoleApiRequest {
	return GetCustomDatabaseRoleApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		roleName:   args.RoleName,
	}
}

func (r GetCustomDatabaseRoleApiRequest) Execute() (*UserCustomDBRole, *http.Response, error) {
	return r.ApiService.GetCustomDatabaseRoleExecute(r)
}

/*
GetCustomDatabaseRole Return One Custom Role in One Project

Returns one custom role for the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param roleName Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.
	@return GetCustomDatabaseRoleApiRequest
*/
func (a *CustomDatabaseRolesApiService) GetCustomDatabaseRole(ctx context.Context, groupId string, roleName string) GetCustomDatabaseRoleApiRequest {
	return GetCustomDatabaseRoleApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		roleName:   roleName,
	}
}

// Execute executes the request
//
//	@return UserCustomDBRole
func (a *CustomDatabaseRolesApiService) GetCustomDatabaseRoleExecute(r GetCustomDatabaseRoleApiRequest) (*UserCustomDBRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserCustomDBRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDatabaseRolesApiService.GetCustomDatabaseRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleName"+"}", url.PathEscape(parameterValueToString(r.roleName, "roleName")), -1)

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

type ListCustomDatabaseRolesApiRequest struct {
	ctx        context.Context
	ApiService CustomDatabaseRolesApi
	groupId    string
}

type ListCustomDatabaseRolesApiParams struct {
	GroupId string
}

func (a *CustomDatabaseRolesApiService) ListCustomDatabaseRolesWithParams(ctx context.Context, args *ListCustomDatabaseRolesApiParams) ListCustomDatabaseRolesApiRequest {
	return ListCustomDatabaseRolesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r ListCustomDatabaseRolesApiRequest) Execute() ([]UserCustomDBRole, *http.Response, error) {
	return r.ApiService.ListCustomDatabaseRolesExecute(r)
}

/*
ListCustomDatabaseRoles Return All Custom Roles in One Project

Returns all custom roles for the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListCustomDatabaseRolesApiRequest
*/
func (a *CustomDatabaseRolesApiService) ListCustomDatabaseRoles(ctx context.Context, groupId string) ListCustomDatabaseRolesApiRequest {
	return ListCustomDatabaseRolesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return []UserCustomDBRole
func (a *CustomDatabaseRolesApiService) ListCustomDatabaseRolesExecute(r ListCustomDatabaseRolesApiRequest) ([]UserCustomDBRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []UserCustomDBRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDatabaseRolesApiService.ListCustomDatabaseRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/customDBRoles/roles"
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

type UpdateCustomDatabaseRoleApiRequest struct {
	ctx                context.Context
	ApiService         CustomDatabaseRolesApi
	groupId            string
	roleName           string
	updateCustomDBRole *UpdateCustomDBRole
}

type UpdateCustomDatabaseRoleApiParams struct {
	GroupId            string
	RoleName           string
	UpdateCustomDBRole *UpdateCustomDBRole
}

func (a *CustomDatabaseRolesApiService) UpdateCustomDatabaseRoleWithParams(ctx context.Context, args *UpdateCustomDatabaseRoleApiParams) UpdateCustomDatabaseRoleApiRequest {
	return UpdateCustomDatabaseRoleApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		roleName:           args.RoleName,
		updateCustomDBRole: args.UpdateCustomDBRole,
	}
}

func (r UpdateCustomDatabaseRoleApiRequest) Execute() (*UserCustomDBRole, *http.Response, error) {
	return r.ApiService.UpdateCustomDatabaseRoleExecute(r)
}

/*
UpdateCustomDatabaseRole Update One Custom Role in One Project

Updates one custom role in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param roleName Human-readable label that identifies the role for the request. This name must beunique for this custom role in this project.
	@return UpdateCustomDatabaseRoleApiRequest
*/
func (a *CustomDatabaseRolesApiService) UpdateCustomDatabaseRole(ctx context.Context, groupId string, roleName string, updateCustomDBRole *UpdateCustomDBRole) UpdateCustomDatabaseRoleApiRequest {
	return UpdateCustomDatabaseRoleApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            groupId,
		roleName:           roleName,
		updateCustomDBRole: updateCustomDBRole,
	}
}

// Execute executes the request
//
//	@return UserCustomDBRole
func (a *CustomDatabaseRolesApiService) UpdateCustomDatabaseRoleExecute(r UpdateCustomDatabaseRoleApiRequest) (*UserCustomDBRole, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserCustomDBRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDatabaseRolesApiService.UpdateCustomDatabaseRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/customDBRoles/roles/{roleName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleName"+"}", url.PathEscape(parameterValueToString(r.roleName, "roleName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateCustomDBRole == nil {
		return localVarReturnValue, nil, reportError("updateCustomDBRole is required and must be specified")
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
	localVarPostBody = r.updateCustomDBRole
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
