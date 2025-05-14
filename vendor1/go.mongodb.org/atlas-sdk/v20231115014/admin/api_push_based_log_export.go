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

type PushBasedLogExportApi interface {

	/*
		CreatePushBasedLogConfiguration Enable the push-based log export feature for a project

		[experimental] Configures the project level settings for the push-based log export feature.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return CreatePushBasedLogConfigurationApiRequest
	*/
	CreatePushBasedLogConfiguration(ctx context.Context, groupId string, createPushBasedLogExportProjectRequest *CreatePushBasedLogExportProjectRequest) CreatePushBasedLogConfigurationApiRequest
	/*
		CreatePushBasedLogConfiguration Enable the push-based log export feature for a project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreatePushBasedLogConfigurationApiParams - Parameters for the request
		@return CreatePushBasedLogConfigurationApiRequest
	*/
	CreatePushBasedLogConfigurationWithParams(ctx context.Context, args *CreatePushBasedLogConfigurationApiParams) CreatePushBasedLogConfigurationApiRequest

	// Method available only for mocking purposes
	CreatePushBasedLogConfigurationExecute(r CreatePushBasedLogConfigurationApiRequest) (*http.Response, error)

	/*
		DeletePushBasedLogConfiguration Disable the push-based log export feature for a project

		[experimental] Disables the push-based log export feature by resetting the project level settings to its default configuration.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeletePushBasedLogConfigurationApiRequest
	*/
	DeletePushBasedLogConfiguration(ctx context.Context, groupId string) DeletePushBasedLogConfigurationApiRequest
	/*
		DeletePushBasedLogConfiguration Disable the push-based log export feature for a project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeletePushBasedLogConfigurationApiParams - Parameters for the request
		@return DeletePushBasedLogConfigurationApiRequest
	*/
	DeletePushBasedLogConfigurationWithParams(ctx context.Context, args *DeletePushBasedLogConfigurationApiParams) DeletePushBasedLogConfigurationApiRequest

	// Method available only for mocking purposes
	DeletePushBasedLogConfigurationExecute(r DeletePushBasedLogConfigurationApiRequest) (*http.Response, error)

	/*
		GetPushBasedLogConfiguration Get the push-based log export configuration for a project

		[experimental] Fetches the current project level settings for the push-based log export feature.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetPushBasedLogConfigurationApiRequest
	*/
	GetPushBasedLogConfiguration(ctx context.Context, groupId string) GetPushBasedLogConfigurationApiRequest
	/*
		GetPushBasedLogConfiguration Get the push-based log export configuration for a project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetPushBasedLogConfigurationApiParams - Parameters for the request
		@return GetPushBasedLogConfigurationApiRequest
	*/
	GetPushBasedLogConfigurationWithParams(ctx context.Context, args *GetPushBasedLogConfigurationApiParams) GetPushBasedLogConfigurationApiRequest

	// Method available only for mocking purposes
	GetPushBasedLogConfigurationExecute(r GetPushBasedLogConfigurationApiRequest) (*PushBasedLogExportProject, *http.Response, error)

	/*
		UpdatePushBasedLogConfiguration Update the push-based log export feature for a project

		[experimental] Updates the project level settings for the push-based log export feature.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return UpdatePushBasedLogConfigurationApiRequest
	*/
	UpdatePushBasedLogConfiguration(ctx context.Context, groupId string, pushBasedLogExportProject *PushBasedLogExportProject) UpdatePushBasedLogConfigurationApiRequest
	/*
		UpdatePushBasedLogConfiguration Update the push-based log export feature for a project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdatePushBasedLogConfigurationApiParams - Parameters for the request
		@return UpdatePushBasedLogConfigurationApiRequest
	*/
	UpdatePushBasedLogConfigurationWithParams(ctx context.Context, args *UpdatePushBasedLogConfigurationApiParams) UpdatePushBasedLogConfigurationApiRequest

	// Method available only for mocking purposes
	UpdatePushBasedLogConfigurationExecute(r UpdatePushBasedLogConfigurationApiRequest) (*http.Response, error)
}

// PushBasedLogExportApiService PushBasedLogExportApi service
type PushBasedLogExportApiService service

type CreatePushBasedLogConfigurationApiRequest struct {
	ctx                                    context.Context
	ApiService                             PushBasedLogExportApi
	groupId                                string
	createPushBasedLogExportProjectRequest *CreatePushBasedLogExportProjectRequest
}

type CreatePushBasedLogConfigurationApiParams struct {
	GroupId                                string
	CreatePushBasedLogExportProjectRequest *CreatePushBasedLogExportProjectRequest
}

func (a *PushBasedLogExportApiService) CreatePushBasedLogConfigurationWithParams(ctx context.Context, args *CreatePushBasedLogConfigurationApiParams) CreatePushBasedLogConfigurationApiRequest {
	return CreatePushBasedLogConfigurationApiRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		groupId:                                args.GroupId,
		createPushBasedLogExportProjectRequest: args.CreatePushBasedLogExportProjectRequest,
	}
}

func (r CreatePushBasedLogConfigurationApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreatePushBasedLogConfigurationExecute(r)
}

/*
CreatePushBasedLogConfiguration Enable the push-based log export feature for a project

[experimental] Configures the project level settings for the push-based log export feature.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreatePushBasedLogConfigurationApiRequest
*/
func (a *PushBasedLogExportApiService) CreatePushBasedLogConfiguration(ctx context.Context, groupId string, createPushBasedLogExportProjectRequest *CreatePushBasedLogExportProjectRequest) CreatePushBasedLogConfigurationApiRequest {
	return CreatePushBasedLogConfigurationApiRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		groupId:                                groupId,
		createPushBasedLogExportProjectRequest: createPushBasedLogExportProjectRequest,
	}
}

// Execute executes the request
func (a *PushBasedLogExportApiService) CreatePushBasedLogConfigurationExecute(r CreatePushBasedLogConfigurationApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushBasedLogExportApiService.CreatePushBasedLogConfiguration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pushBasedLogExport"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createPushBasedLogExportProjectRequest == nil {
		return nil, reportError("createPushBasedLogExportProjectRequest is required and must be specified")
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
	localVarPostBody = r.createPushBasedLogExportProjectRequest
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

type DeletePushBasedLogConfigurationApiRequest struct {
	ctx        context.Context
	ApiService PushBasedLogExportApi
	groupId    string
}

type DeletePushBasedLogConfigurationApiParams struct {
	GroupId string
}

func (a *PushBasedLogExportApiService) DeletePushBasedLogConfigurationWithParams(ctx context.Context, args *DeletePushBasedLogConfigurationApiParams) DeletePushBasedLogConfigurationApiRequest {
	return DeletePushBasedLogConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r DeletePushBasedLogConfigurationApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePushBasedLogConfigurationExecute(r)
}

/*
DeletePushBasedLogConfiguration Disable the push-based log export feature for a project

[experimental] Disables the push-based log export feature by resetting the project level settings to its default configuration.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeletePushBasedLogConfigurationApiRequest
*/
func (a *PushBasedLogExportApiService) DeletePushBasedLogConfiguration(ctx context.Context, groupId string) DeletePushBasedLogConfigurationApiRequest {
	return DeletePushBasedLogConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
func (a *PushBasedLogExportApiService) DeletePushBasedLogConfigurationExecute(r DeletePushBasedLogConfigurationApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushBasedLogExportApiService.DeletePushBasedLogConfiguration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pushBasedLogExport"
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

type GetPushBasedLogConfigurationApiRequest struct {
	ctx        context.Context
	ApiService PushBasedLogExportApi
	groupId    string
}

type GetPushBasedLogConfigurationApiParams struct {
	GroupId string
}

func (a *PushBasedLogExportApiService) GetPushBasedLogConfigurationWithParams(ctx context.Context, args *GetPushBasedLogConfigurationApiParams) GetPushBasedLogConfigurationApiRequest {
	return GetPushBasedLogConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetPushBasedLogConfigurationApiRequest) Execute() (*PushBasedLogExportProject, *http.Response, error) {
	return r.ApiService.GetPushBasedLogConfigurationExecute(r)
}

/*
GetPushBasedLogConfiguration Get the push-based log export configuration for a project

[experimental] Fetches the current project level settings for the push-based log export feature.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetPushBasedLogConfigurationApiRequest
*/
func (a *PushBasedLogExportApiService) GetPushBasedLogConfiguration(ctx context.Context, groupId string) GetPushBasedLogConfigurationApiRequest {
	return GetPushBasedLogConfigurationApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return PushBasedLogExportProject
func (a *PushBasedLogExportApiService) GetPushBasedLogConfigurationExecute(r GetPushBasedLogConfigurationApiRequest) (*PushBasedLogExportProject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PushBasedLogExportProject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushBasedLogExportApiService.GetPushBasedLogConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pushBasedLogExport"
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

type UpdatePushBasedLogConfigurationApiRequest struct {
	ctx                       context.Context
	ApiService                PushBasedLogExportApi
	groupId                   string
	pushBasedLogExportProject *PushBasedLogExportProject
}

type UpdatePushBasedLogConfigurationApiParams struct {
	GroupId                   string
	PushBasedLogExportProject *PushBasedLogExportProject
}

func (a *PushBasedLogExportApiService) UpdatePushBasedLogConfigurationWithParams(ctx context.Context, args *UpdatePushBasedLogConfigurationApiParams) UpdatePushBasedLogConfigurationApiRequest {
	return UpdatePushBasedLogConfigurationApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   args.GroupId,
		pushBasedLogExportProject: args.PushBasedLogExportProject,
	}
}

func (r UpdatePushBasedLogConfigurationApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdatePushBasedLogConfigurationExecute(r)
}

/*
UpdatePushBasedLogConfiguration Update the push-based log export feature for a project

[experimental] Updates the project level settings for the push-based log export feature.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdatePushBasedLogConfigurationApiRequest
*/
func (a *PushBasedLogExportApiService) UpdatePushBasedLogConfiguration(ctx context.Context, groupId string, pushBasedLogExportProject *PushBasedLogExportProject) UpdatePushBasedLogConfigurationApiRequest {
	return UpdatePushBasedLogConfigurationApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   groupId,
		pushBasedLogExportProject: pushBasedLogExportProject,
	}
}

// Execute executes the request
func (a *PushBasedLogExportApiService) UpdatePushBasedLogConfigurationExecute(r UpdatePushBasedLogConfigurationApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PushBasedLogExportApiService.UpdatePushBasedLogConfiguration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pushBasedLogExport"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pushBasedLogExportProject == nil {
		return nil, reportError("pushBasedLogExportProject is required and must be specified")
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
	localVarPostBody = r.pushBasedLogExportProject
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
