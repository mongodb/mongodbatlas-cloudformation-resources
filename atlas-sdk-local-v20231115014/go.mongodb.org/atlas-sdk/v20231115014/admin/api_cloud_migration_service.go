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

type CloudMigrationServiceApi interface {

	/*
		CreateLinkToken Create One Link-Token

		Create one link-token that contains all the information required to complete the link. MongoDB Atlas uses the link-token for push live migrations only. Live migration (push) allows you to securely push data from Cloud Manager or Ops Manager into MongoDB Atlas. Your API Key must have the Organization Owner role to successfully call this resource.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return CreateLinkTokenApiRequest
	*/
	CreateLinkToken(ctx context.Context, orgId string, targetOrgRequest *TargetOrgRequest) CreateLinkTokenApiRequest
	/*
		CreateLinkToken Create One Link-Token


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateLinkTokenApiParams - Parameters for the request
		@return CreateLinkTokenApiRequest
	*/
	CreateLinkTokenWithParams(ctx context.Context, args *CreateLinkTokenApiParams) CreateLinkTokenApiRequest

	// Method available only for mocking purposes
	CreateLinkTokenExecute(r CreateLinkTokenApiRequest) (*TargetOrg, *http.Response, error)

	/*
			CreatePushMigration Migrate One Local Managed Cluster to MongoDB Atlas

			Migrate one cluster that Cloud or Ops Manager manages to MongoDB Atlas.

		 Please make sure to [validate](#tag/Cloud-Migration-Service/operation/validateMigration) your migration before initiating it.

		 You can use this API endpoint for push live migrations only. Your API Key must have the Organization Owner role to successfully call this resource.

		 **NOTE**: Migrating time-series collections is not yet supported on MongoDB 6.0 or higher. Migrations on MongoDB 6.0 or higher will skip any time-series collections on the source cluster.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return CreatePushMigrationApiRequest
	*/
	CreatePushMigration(ctx context.Context, groupId string, liveMigrationRequest *LiveMigrationRequest) CreatePushMigrationApiRequest
	/*
		CreatePushMigration Migrate One Local Managed Cluster to MongoDB Atlas


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreatePushMigrationApiParams - Parameters for the request
		@return CreatePushMigrationApiRequest
	*/
	CreatePushMigrationWithParams(ctx context.Context, args *CreatePushMigrationApiParams) CreatePushMigrationApiRequest

	// Method available only for mocking purposes
	CreatePushMigrationExecute(r CreatePushMigrationApiRequest) (*LiveMigrationResponse, *http.Response, error)

	/*
		CutoverMigration Cut Over the Migrated Cluster

		Cut over the migrated cluster to MongoDB Atlas. Confirm when the cut over completes. When the cut over completes, MongoDB Atlas completes the live migration process and stops synchronizing with the source cluster. Your API Key must have the Organization Owner role to successfully call this resource.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param liveMigrationId Unique 24-hexadecimal digit string that identifies the migration.
		@return CutoverMigrationApiRequest
	*/
	CutoverMigration(ctx context.Context, groupId string, liveMigrationId string) CutoverMigrationApiRequest
	/*
		CutoverMigration Cut Over the Migrated Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CutoverMigrationApiParams - Parameters for the request
		@return CutoverMigrationApiRequest
	*/
	CutoverMigrationWithParams(ctx context.Context, args *CutoverMigrationApiParams) CutoverMigrationApiRequest

	// Method available only for mocking purposes
	CutoverMigrationExecute(r CutoverMigrationApiRequest) (*http.Response, error)

	/*
		DeleteLinkToken Remove One Link-Token

		Remove one organization link and its associated public API key. MongoDB Atlas uses the link-token for push live migrations only. Live migrations (push) let you securely push data from Cloud Manager or Ops Manager into MongoDB Atlas. Your API Key must have the Organization Owner role to successfully call this resource.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return DeleteLinkTokenApiRequest
	*/
	DeleteLinkToken(ctx context.Context, orgId string) DeleteLinkTokenApiRequest
	/*
		DeleteLinkToken Remove One Link-Token


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteLinkTokenApiParams - Parameters for the request
		@return DeleteLinkTokenApiRequest
	*/
	DeleteLinkTokenWithParams(ctx context.Context, args *DeleteLinkTokenApiParams) DeleteLinkTokenApiRequest

	// Method available only for mocking purposes
	DeleteLinkTokenExecute(r DeleteLinkTokenApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetPushMigration Return One Migration Job

		Return details of one cluster migration job. Each push live migration job uses one migration host. Your API Key must have the Organization Member role to successfully call this resource.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param liveMigrationId Unique 24-hexadecimal digit string that identifies the migration.
		@return GetPushMigrationApiRequest
	*/
	GetPushMigration(ctx context.Context, groupId string, liveMigrationId string) GetPushMigrationApiRequest
	/*
		GetPushMigration Return One Migration Job


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetPushMigrationApiParams - Parameters for the request
		@return GetPushMigrationApiRequest
	*/
	GetPushMigrationWithParams(ctx context.Context, args *GetPushMigrationApiParams) GetPushMigrationApiRequest

	// Method available only for mocking purposes
	GetPushMigrationExecute(r GetPushMigrationApiRequest) (*LiveMigrationResponse, *http.Response, error)

	/*
		GetValidationStatus Return One Migration Validation Job

		Return the status of one migration validation job. Your API Key must have the Organization Owner role to successfully call this resource.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param validationId Unique 24-hexadecimal digit string that identifies the validation job.
		@return GetValidationStatusApiRequest
	*/
	GetValidationStatus(ctx context.Context, groupId string, validationId string) GetValidationStatusApiRequest
	/*
		GetValidationStatus Return One Migration Validation Job


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetValidationStatusApiParams - Parameters for the request
		@return GetValidationStatusApiRequest
	*/
	GetValidationStatusWithParams(ctx context.Context, args *GetValidationStatusApiParams) GetValidationStatusApiRequest

	// Method available only for mocking purposes
	GetValidationStatusExecute(r GetValidationStatusApiRequest) (*LiveImportValidation, *http.Response, error)

	/*
		ListSourceProjects Return All Projects Available for Migration

		[experimental] Return all projects that you can migrate to the specified organization.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListSourceProjectsApiRequest
	*/
	ListSourceProjects(ctx context.Context, orgId string) ListSourceProjectsApiRequest
	/*
		ListSourceProjects Return All Projects Available for Migration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListSourceProjectsApiParams - Parameters for the request
		@return ListSourceProjectsApiRequest
	*/
	ListSourceProjectsWithParams(ctx context.Context, args *ListSourceProjectsApiParams) ListSourceProjectsApiRequest

	// Method available only for mocking purposes
	ListSourceProjectsExecute(r ListSourceProjectsApiRequest) ([]LiveImportAvailableProject, *http.Response, error)

	/*
		ValidateMigration Validate One Migration Request

		Verifies whether the provided credentials, available disk space, MongoDB versions, and so on meet the requirements of the migration request. If the check passes, the migration can proceed. Your API Key must have the Organization Owner role to successfully call this resource.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ValidateMigrationApiRequest
	*/
	ValidateMigration(ctx context.Context, groupId string, liveMigrationRequest *LiveMigrationRequest) ValidateMigrationApiRequest
	/*
		ValidateMigration Validate One Migration Request


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ValidateMigrationApiParams - Parameters for the request
		@return ValidateMigrationApiRequest
	*/
	ValidateMigrationWithParams(ctx context.Context, args *ValidateMigrationApiParams) ValidateMigrationApiRequest

	// Method available only for mocking purposes
	ValidateMigrationExecute(r ValidateMigrationApiRequest) (*LiveImportValidation, *http.Response, error)
}

// CloudMigrationServiceApiService CloudMigrationServiceApi service
type CloudMigrationServiceApiService service

type CreateLinkTokenApiRequest struct {
	ctx              context.Context
	ApiService       CloudMigrationServiceApi
	orgId            string
	targetOrgRequest *TargetOrgRequest
}

type CreateLinkTokenApiParams struct {
	OrgId            string
	TargetOrgRequest *TargetOrgRequest
}

func (a *CloudMigrationServiceApiService) CreateLinkTokenWithParams(ctx context.Context, args *CreateLinkTokenApiParams) CreateLinkTokenApiRequest {
	return CreateLinkTokenApiRequest{
		ApiService:       a,
		ctx:              ctx,
		orgId:            args.OrgId,
		targetOrgRequest: args.TargetOrgRequest,
	}
}

func (r CreateLinkTokenApiRequest) Execute() (*TargetOrg, *http.Response, error) {
	return r.ApiService.CreateLinkTokenExecute(r)
}

/*
CreateLinkToken Create One Link-Token

Create one link-token that contains all the information required to complete the link. MongoDB Atlas uses the link-token for push live migrations only. Live migration (push) allows you to securely push data from Cloud Manager or Ops Manager into MongoDB Atlas. Your API Key must have the Organization Owner role to successfully call this resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateLinkTokenApiRequest
*/
func (a *CloudMigrationServiceApiService) CreateLinkToken(ctx context.Context, orgId string, targetOrgRequest *TargetOrgRequest) CreateLinkTokenApiRequest {
	return CreateLinkTokenApiRequest{
		ApiService:       a,
		ctx:              ctx,
		orgId:            orgId,
		targetOrgRequest: targetOrgRequest,
	}
}

// Execute executes the request
//
//	@return TargetOrg
func (a *CloudMigrationServiceApiService) CreateLinkTokenExecute(r CreateLinkTokenApiRequest) (*TargetOrg, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TargetOrg
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMigrationServiceApiService.CreateLinkToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.targetOrgRequest == nil {
		return localVarReturnValue, nil, reportError("targetOrgRequest is required and must be specified")
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
	localVarPostBody = r.targetOrgRequest
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

type CreatePushMigrationApiRequest struct {
	ctx                  context.Context
	ApiService           CloudMigrationServiceApi
	groupId              string
	liveMigrationRequest *LiveMigrationRequest
}

type CreatePushMigrationApiParams struct {
	GroupId              string
	LiveMigrationRequest *LiveMigrationRequest
}

func (a *CloudMigrationServiceApiService) CreatePushMigrationWithParams(ctx context.Context, args *CreatePushMigrationApiParams) CreatePushMigrationApiRequest {
	return CreatePushMigrationApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		groupId:              args.GroupId,
		liveMigrationRequest: args.LiveMigrationRequest,
	}
}

func (r CreatePushMigrationApiRequest) Execute() (*LiveMigrationResponse, *http.Response, error) {
	return r.ApiService.CreatePushMigrationExecute(r)
}

/*
CreatePushMigration Migrate One Local Managed Cluster to MongoDB Atlas

Migrate one cluster that Cloud or Ops Manager manages to MongoDB Atlas.

	Please make sure to [validate](#tag/Cloud-Migration-Service/operation/validateMigration) your migration before initiating it.

	You can use this API endpoint for push live migrations only. Your API Key must have the Organization Owner role to successfully call this resource.

	**NOTE**: Migrating time-series collections is not yet supported on MongoDB 6.0 or higher. Migrations on MongoDB 6.0 or higher will skip any time-series collections on the source cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreatePushMigrationApiRequest
*/
func (a *CloudMigrationServiceApiService) CreatePushMigration(ctx context.Context, groupId string, liveMigrationRequest *LiveMigrationRequest) CreatePushMigrationApiRequest {
	return CreatePushMigrationApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		groupId:              groupId,
		liveMigrationRequest: liveMigrationRequest,
	}
}

// Execute executes the request
//
//	@return LiveMigrationResponse
func (a *CloudMigrationServiceApiService) CreatePushMigrationExecute(r CreatePushMigrationApiRequest) (*LiveMigrationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LiveMigrationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMigrationServiceApiService.CreatePushMigration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/liveMigrations"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.liveMigrationRequest == nil {
		return localVarReturnValue, nil, reportError("liveMigrationRequest is required and must be specified")
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
	localVarPostBody = r.liveMigrationRequest
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

type CutoverMigrationApiRequest struct {
	ctx             context.Context
	ApiService      CloudMigrationServiceApi
	groupId         string
	liveMigrationId string
}

type CutoverMigrationApiParams struct {
	GroupId         string
	LiveMigrationId string
}

func (a *CloudMigrationServiceApiService) CutoverMigrationWithParams(ctx context.Context, args *CutoverMigrationApiParams) CutoverMigrationApiRequest {
	return CutoverMigrationApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         args.GroupId,
		liveMigrationId: args.LiveMigrationId,
	}
}

func (r CutoverMigrationApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.CutoverMigrationExecute(r)
}

/*
CutoverMigration Cut Over the Migrated Cluster

Cut over the migrated cluster to MongoDB Atlas. Confirm when the cut over completes. When the cut over completes, MongoDB Atlas completes the live migration process and stops synchronizing with the source cluster. Your API Key must have the Organization Owner role to successfully call this resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param liveMigrationId Unique 24-hexadecimal digit string that identifies the migration.
	@return CutoverMigrationApiRequest
*/
func (a *CloudMigrationServiceApiService) CutoverMigration(ctx context.Context, groupId string, liveMigrationId string) CutoverMigrationApiRequest {
	return CutoverMigrationApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         groupId,
		liveMigrationId: liveMigrationId,
	}
}

// Execute executes the request
func (a *CloudMigrationServiceApiService) CutoverMigrationExecute(r CutoverMigrationApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMigrationServiceApiService.CutoverMigration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId}/cutover"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"liveMigrationId"+"}", url.PathEscape(parameterValueToString(r.liveMigrationId, "liveMigrationId")), -1)

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

type DeleteLinkTokenApiRequest struct {
	ctx        context.Context
	ApiService CloudMigrationServiceApi
	orgId      string
}

type DeleteLinkTokenApiParams struct {
	OrgId string
}

func (a *CloudMigrationServiceApiService) DeleteLinkTokenWithParams(ctx context.Context, args *DeleteLinkTokenApiParams) DeleteLinkTokenApiRequest {
	return DeleteLinkTokenApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
	}
}

func (r DeleteLinkTokenApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteLinkTokenExecute(r)
}

/*
DeleteLinkToken Remove One Link-Token

Remove one organization link and its associated public API key. MongoDB Atlas uses the link-token for push live migrations only. Live migrations (push) let you securely push data from Cloud Manager or Ops Manager into MongoDB Atlas. Your API Key must have the Organization Owner role to successfully call this resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return DeleteLinkTokenApiRequest
*/
func (a *CloudMigrationServiceApiService) DeleteLinkToken(ctx context.Context, orgId string) DeleteLinkTokenApiRequest {
	return DeleteLinkTokenApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *CloudMigrationServiceApiService) DeleteLinkTokenExecute(r DeleteLinkTokenApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMigrationServiceApiService.DeleteLinkToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/liveMigrations/linkTokens"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

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

type GetPushMigrationApiRequest struct {
	ctx             context.Context
	ApiService      CloudMigrationServiceApi
	groupId         string
	liveMigrationId string
}

type GetPushMigrationApiParams struct {
	GroupId         string
	LiveMigrationId string
}

func (a *CloudMigrationServiceApiService) GetPushMigrationWithParams(ctx context.Context, args *GetPushMigrationApiParams) GetPushMigrationApiRequest {
	return GetPushMigrationApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         args.GroupId,
		liveMigrationId: args.LiveMigrationId,
	}
}

func (r GetPushMigrationApiRequest) Execute() (*LiveMigrationResponse, *http.Response, error) {
	return r.ApiService.GetPushMigrationExecute(r)
}

/*
GetPushMigration Return One Migration Job

Return details of one cluster migration job. Each push live migration job uses one migration host. Your API Key must have the Organization Member role to successfully call this resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param liveMigrationId Unique 24-hexadecimal digit string that identifies the migration.
	@return GetPushMigrationApiRequest
*/
func (a *CloudMigrationServiceApiService) GetPushMigration(ctx context.Context, groupId string, liveMigrationId string) GetPushMigrationApiRequest {
	return GetPushMigrationApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         groupId,
		liveMigrationId: liveMigrationId,
	}
}

// Execute executes the request
//
//	@return LiveMigrationResponse
func (a *CloudMigrationServiceApiService) GetPushMigrationExecute(r GetPushMigrationApiRequest) (*LiveMigrationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LiveMigrationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMigrationServiceApiService.GetPushMigration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/liveMigrations/{liveMigrationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"liveMigrationId"+"}", url.PathEscape(parameterValueToString(r.liveMigrationId, "liveMigrationId")), -1)

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

type GetValidationStatusApiRequest struct {
	ctx          context.Context
	ApiService   CloudMigrationServiceApi
	groupId      string
	validationId string
}

type GetValidationStatusApiParams struct {
	GroupId      string
	ValidationId string
}

func (a *CloudMigrationServiceApiService) GetValidationStatusWithParams(ctx context.Context, args *GetValidationStatusApiParams) GetValidationStatusApiRequest {
	return GetValidationStatusApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		validationId: args.ValidationId,
	}
}

func (r GetValidationStatusApiRequest) Execute() (*LiveImportValidation, *http.Response, error) {
	return r.ApiService.GetValidationStatusExecute(r)
}

/*
GetValidationStatus Return One Migration Validation Job

Return the status of one migration validation job. Your API Key must have the Organization Owner role to successfully call this resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param validationId Unique 24-hexadecimal digit string that identifies the validation job.
	@return GetValidationStatusApiRequest
*/
func (a *CloudMigrationServiceApiService) GetValidationStatus(ctx context.Context, groupId string, validationId string) GetValidationStatusApiRequest {
	return GetValidationStatusApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		validationId: validationId,
	}
}

// Execute executes the request
//
//	@return LiveImportValidation
func (a *CloudMigrationServiceApiService) GetValidationStatusExecute(r GetValidationStatusApiRequest) (*LiveImportValidation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LiveImportValidation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMigrationServiceApiService.GetValidationStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/liveMigrations/validate/{validationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"validationId"+"}", url.PathEscape(parameterValueToString(r.validationId, "validationId")), -1)

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

type ListSourceProjectsApiRequest struct {
	ctx        context.Context
	ApiService CloudMigrationServiceApi
	orgId      string
}

type ListSourceProjectsApiParams struct {
	OrgId string
}

func (a *CloudMigrationServiceApiService) ListSourceProjectsWithParams(ctx context.Context, args *ListSourceProjectsApiParams) ListSourceProjectsApiRequest {
	return ListSourceProjectsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
	}
}

func (r ListSourceProjectsApiRequest) Execute() ([]LiveImportAvailableProject, *http.Response, error) {
	return r.ApiService.ListSourceProjectsExecute(r)
}

/*
ListSourceProjects Return All Projects Available for Migration

[experimental] Return all projects that you can migrate to the specified organization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListSourceProjectsApiRequest
*/
func (a *CloudMigrationServiceApiService) ListSourceProjects(ctx context.Context, orgId string) ListSourceProjectsApiRequest {
	return ListSourceProjectsApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return []LiveImportAvailableProject
func (a *CloudMigrationServiceApiService) ListSourceProjectsExecute(r ListSourceProjectsApiRequest) ([]LiveImportAvailableProject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []LiveImportAvailableProject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMigrationServiceApiService.ListSourceProjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/liveMigrations/availableProjects"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

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

type ValidateMigrationApiRequest struct {
	ctx                  context.Context
	ApiService           CloudMigrationServiceApi
	groupId              string
	liveMigrationRequest *LiveMigrationRequest
}

type ValidateMigrationApiParams struct {
	GroupId              string
	LiveMigrationRequest *LiveMigrationRequest
}

func (a *CloudMigrationServiceApiService) ValidateMigrationWithParams(ctx context.Context, args *ValidateMigrationApiParams) ValidateMigrationApiRequest {
	return ValidateMigrationApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		groupId:              args.GroupId,
		liveMigrationRequest: args.LiveMigrationRequest,
	}
}

func (r ValidateMigrationApiRequest) Execute() (*LiveImportValidation, *http.Response, error) {
	return r.ApiService.ValidateMigrationExecute(r)
}

/*
ValidateMigration Validate One Migration Request

Verifies whether the provided credentials, available disk space, MongoDB versions, and so on meet the requirements of the migration request. If the check passes, the migration can proceed. Your API Key must have the Organization Owner role to successfully call this resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ValidateMigrationApiRequest
*/
func (a *CloudMigrationServiceApiService) ValidateMigration(ctx context.Context, groupId string, liveMigrationRequest *LiveMigrationRequest) ValidateMigrationApiRequest {
	return ValidateMigrationApiRequest{
		ApiService:           a,
		ctx:                  ctx,
		groupId:              groupId,
		liveMigrationRequest: liveMigrationRequest,
	}
}

// Execute executes the request
//
//	@return LiveImportValidation
func (a *CloudMigrationServiceApiService) ValidateMigrationExecute(r ValidateMigrationApiRequest) (*LiveImportValidation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LiveImportValidation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudMigrationServiceApiService.ValidateMigration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/liveMigrations/validate"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.liveMigrationRequest == nil {
		return localVarReturnValue, nil, reportError("liveMigrationRequest is required and must be specified")
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
	localVarPostBody = r.liveMigrationRequest
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
