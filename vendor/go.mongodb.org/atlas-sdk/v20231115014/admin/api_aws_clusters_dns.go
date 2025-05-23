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

type AWSClustersDNSApi interface {

	/*
		GetAWSCustomDNS Return One Custom DNS Configuration for Atlas Clusters on AWS

		Returns the custom DNS configuration for AWS clusters in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetAWSCustomDNSApiRequest
	*/
	GetAWSCustomDNS(ctx context.Context, groupId string) GetAWSCustomDNSApiRequest
	/*
		GetAWSCustomDNS Return One Custom DNS Configuration for Atlas Clusters on AWS


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAWSCustomDNSApiParams - Parameters for the request
		@return GetAWSCustomDNSApiRequest
	*/
	GetAWSCustomDNSWithParams(ctx context.Context, args *GetAWSCustomDNSApiParams) GetAWSCustomDNSApiRequest

	// Method available only for mocking purposes
	GetAWSCustomDNSExecute(r GetAWSCustomDNSApiRequest) (*AWSCustomDNSEnabled, *http.Response, error)

	/*
		ToggleAWSCustomDNS Toggle State of One Custom DNS Configuration for Atlas Clusters on AWS

		Enables or disables the custom DNS configuration for AWS clusters in the specified project. Enable custom DNS if you use AWS VPC peering and use your own DNS servers. To use this resource, the requesting API Key must have the Project Atlas Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ToggleAWSCustomDNSApiRequest
	*/
	ToggleAWSCustomDNS(ctx context.Context, groupId string, aWSCustomDNSEnabled *AWSCustomDNSEnabled) ToggleAWSCustomDNSApiRequest
	/*
		ToggleAWSCustomDNS Toggle State of One Custom DNS Configuration for Atlas Clusters on AWS


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ToggleAWSCustomDNSApiParams - Parameters for the request
		@return ToggleAWSCustomDNSApiRequest
	*/
	ToggleAWSCustomDNSWithParams(ctx context.Context, args *ToggleAWSCustomDNSApiParams) ToggleAWSCustomDNSApiRequest

	// Method available only for mocking purposes
	ToggleAWSCustomDNSExecute(r ToggleAWSCustomDNSApiRequest) (*AWSCustomDNSEnabled, *http.Response, error)
}

// AWSClustersDNSApiService AWSClustersDNSApi service
type AWSClustersDNSApiService service

type GetAWSCustomDNSApiRequest struct {
	ctx        context.Context
	ApiService AWSClustersDNSApi
	groupId    string
}

type GetAWSCustomDNSApiParams struct {
	GroupId string
}

func (a *AWSClustersDNSApiService) GetAWSCustomDNSWithParams(ctx context.Context, args *GetAWSCustomDNSApiParams) GetAWSCustomDNSApiRequest {
	return GetAWSCustomDNSApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetAWSCustomDNSApiRequest) Execute() (*AWSCustomDNSEnabled, *http.Response, error) {
	return r.ApiService.GetAWSCustomDNSExecute(r)
}

/*
GetAWSCustomDNS Return One Custom DNS Configuration for Atlas Clusters on AWS

Returns the custom DNS configuration for AWS clusters in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetAWSCustomDNSApiRequest
*/
func (a *AWSClustersDNSApiService) GetAWSCustomDNS(ctx context.Context, groupId string) GetAWSCustomDNSApiRequest {
	return GetAWSCustomDNSApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return AWSCustomDNSEnabled
func (a *AWSClustersDNSApiService) GetAWSCustomDNSExecute(r GetAWSCustomDNSApiRequest) (*AWSCustomDNSEnabled, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AWSCustomDNSEnabled
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AWSClustersDNSApiService.GetAWSCustomDNS")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/awsCustomDNS"
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

type ToggleAWSCustomDNSApiRequest struct {
	ctx                 context.Context
	ApiService          AWSClustersDNSApi
	groupId             string
	aWSCustomDNSEnabled *AWSCustomDNSEnabled
}

type ToggleAWSCustomDNSApiParams struct {
	GroupId             string
	AWSCustomDNSEnabled *AWSCustomDNSEnabled
}

func (a *AWSClustersDNSApiService) ToggleAWSCustomDNSWithParams(ctx context.Context, args *ToggleAWSCustomDNSApiParams) ToggleAWSCustomDNSApiRequest {
	return ToggleAWSCustomDNSApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		groupId:             args.GroupId,
		aWSCustomDNSEnabled: args.AWSCustomDNSEnabled,
	}
}

func (r ToggleAWSCustomDNSApiRequest) Execute() (*AWSCustomDNSEnabled, *http.Response, error) {
	return r.ApiService.ToggleAWSCustomDNSExecute(r)
}

/*
ToggleAWSCustomDNS Toggle State of One Custom DNS Configuration for Atlas Clusters on AWS

Enables or disables the custom DNS configuration for AWS clusters in the specified project. Enable custom DNS if you use AWS VPC peering and use your own DNS servers. To use this resource, the requesting API Key must have the Project Atlas Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ToggleAWSCustomDNSApiRequest
*/
func (a *AWSClustersDNSApiService) ToggleAWSCustomDNS(ctx context.Context, groupId string, aWSCustomDNSEnabled *AWSCustomDNSEnabled) ToggleAWSCustomDNSApiRequest {
	return ToggleAWSCustomDNSApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		groupId:             groupId,
		aWSCustomDNSEnabled: aWSCustomDNSEnabled,
	}
}

// Execute executes the request
//
//	@return AWSCustomDNSEnabled
func (a *AWSClustersDNSApiService) ToggleAWSCustomDNSExecute(r ToggleAWSCustomDNSApiRequest) (*AWSCustomDNSEnabled, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AWSCustomDNSEnabled
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AWSClustersDNSApiService.ToggleAWSCustomDNS")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/awsCustomDNS"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aWSCustomDNSEnabled == nil {
		return localVarReturnValue, nil, reportError("aWSCustomDNSEnabled is required and must be specified")
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
	localVarPostBody = r.aWSCustomDNSEnabled
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
