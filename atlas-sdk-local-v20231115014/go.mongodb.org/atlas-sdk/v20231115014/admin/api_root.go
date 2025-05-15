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
)

type RootApi interface {

	/*
		GetSystemStatus Return the status of this MongoDB application

		[experimental] This resource returns information about the MongoDB application along with API key meta data.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return GetSystemStatusApiRequest
	*/
	GetSystemStatus(ctx context.Context) GetSystemStatusApiRequest
	/*
		GetSystemStatus Return the status of this MongoDB application


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetSystemStatusApiParams - Parameters for the request
		@return GetSystemStatusApiRequest
	*/
	GetSystemStatusWithParams(ctx context.Context, args *GetSystemStatusApiParams) GetSystemStatusApiRequest

	// Method available only for mocking purposes
	GetSystemStatusExecute(r GetSystemStatusApiRequest) (*SystemStatus, *http.Response, error)

	/*
		ReturnAllControlPlaneIPAddresses Return All Control Plane IP Addresses

		[experimental] Returns all control plane IP addresses.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ReturnAllControlPlaneIPAddressesApiRequest
	*/
	ReturnAllControlPlaneIPAddresses(ctx context.Context) ReturnAllControlPlaneIPAddressesApiRequest
	/*
		ReturnAllControlPlaneIPAddresses Return All Control Plane IP Addresses


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ReturnAllControlPlaneIPAddressesApiParams - Parameters for the request
		@return ReturnAllControlPlaneIPAddressesApiRequest
	*/
	ReturnAllControlPlaneIPAddressesWithParams(ctx context.Context, args *ReturnAllControlPlaneIPAddressesApiParams) ReturnAllControlPlaneIPAddressesApiRequest

	// Method available only for mocking purposes
	ReturnAllControlPlaneIPAddressesExecute(r ReturnAllControlPlaneIPAddressesApiRequest) (*ControlPlaneIPAddresses, *http.Response, error)
}

// RootApiService RootApi service
type RootApiService service

type GetSystemStatusApiRequest struct {
	ctx        context.Context
	ApiService RootApi
}

type GetSystemStatusApiParams struct {
}

func (a *RootApiService) GetSystemStatusWithParams(ctx context.Context, args *GetSystemStatusApiParams) GetSystemStatusApiRequest {
	return GetSystemStatusApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (r GetSystemStatusApiRequest) Execute() (*SystemStatus, *http.Response, error) {
	return r.ApiService.GetSystemStatusExecute(r)
}

/*
GetSystemStatus Return the status of this MongoDB application

[experimental] This resource returns information about the MongoDB application along with API key meta data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GetSystemStatusApiRequest
*/
func (a *RootApiService) GetSystemStatus(ctx context.Context) GetSystemStatusApiRequest {
	return GetSystemStatusApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SystemStatus
func (a *RootApiService) GetSystemStatusExecute(r GetSystemStatusApiRequest) (*SystemStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SystemStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RootApiService.GetSystemStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2"

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

type ReturnAllControlPlaneIPAddressesApiRequest struct {
	ctx        context.Context
	ApiService RootApi
}

type ReturnAllControlPlaneIPAddressesApiParams struct {
}

func (a *RootApiService) ReturnAllControlPlaneIPAddressesWithParams(ctx context.Context, args *ReturnAllControlPlaneIPAddressesApiParams) ReturnAllControlPlaneIPAddressesApiRequest {
	return ReturnAllControlPlaneIPAddressesApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (r ReturnAllControlPlaneIPAddressesApiRequest) Execute() (*ControlPlaneIPAddresses, *http.Response, error) {
	return r.ApiService.ReturnAllControlPlaneIPAddressesExecute(r)
}

/*
ReturnAllControlPlaneIPAddresses Return All Control Plane IP Addresses

[experimental] Returns all control plane IP addresses.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReturnAllControlPlaneIPAddressesApiRequest
*/
func (a *RootApiService) ReturnAllControlPlaneIPAddresses(ctx context.Context) ReturnAllControlPlaneIPAddressesApiRequest {
	return ReturnAllControlPlaneIPAddressesApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ControlPlaneIPAddresses
func (a *RootApiService) ReturnAllControlPlaneIPAddressesExecute(r ReturnAllControlPlaneIPAddressesApiRequest) (*ControlPlaneIPAddresses, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ControlPlaneIPAddresses
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RootApiService.ReturnAllControlPlaneIPAddresses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/unauth/controlPlaneIPAddresses"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-11-15+json", "application/json"}

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
