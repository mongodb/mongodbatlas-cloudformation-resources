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

type AtlasSearchApi interface {

	/*
		CreateAtlasSearchDeployment Create Search Nodes

		[experimental] Creates Search Nodes for the specified cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Label that identifies the cluster to create Search Nodes for.
		@return CreateAtlasSearchDeploymentApiRequest
	*/
	CreateAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string, apiSearchDeploymentRequest *ApiSearchDeploymentRequest) CreateAtlasSearchDeploymentApiRequest
	/*
		CreateAtlasSearchDeployment Create Search Nodes


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateAtlasSearchDeploymentApiParams - Parameters for the request
		@return CreateAtlasSearchDeploymentApiRequest
	*/
	CreateAtlasSearchDeploymentWithParams(ctx context.Context, args *CreateAtlasSearchDeploymentApiParams) CreateAtlasSearchDeploymentApiRequest

	// Method available only for mocking purposes
	CreateAtlasSearchDeploymentExecute(r CreateAtlasSearchDeploymentApiRequest) (*ApiSearchDeploymentResponse, *http.Response, error)

	/*
		CreateAtlasSearchIndex Create One Atlas Search Index

		Creates one Atlas Search index on the specified collection. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. Only clusters running MongoDB v4.2 or later can use Atlas Search. To use this resource, the requesting API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection on which to create an Atlas Search index.
		@return CreateAtlasSearchIndexApiRequest
	*/
	CreateAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, clusterSearchIndex *ClusterSearchIndex) CreateAtlasSearchIndexApiRequest
	/*
		CreateAtlasSearchIndex Create One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateAtlasSearchIndexApiParams - Parameters for the request
		@return CreateAtlasSearchIndexApiRequest
	*/
	CreateAtlasSearchIndexWithParams(ctx context.Context, args *CreateAtlasSearchIndexApiParams) CreateAtlasSearchIndexApiRequest

	// Method available only for mocking purposes
	CreateAtlasSearchIndexExecute(r CreateAtlasSearchIndexApiRequest) (*ClusterSearchIndex, *http.Response, error)

	/*
		DeleteAtlasSearchDeployment Delete Search Nodes

		[experimental] Deletes the Search Nodes for the specified cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Label that identifies the cluster to delete.
		@return DeleteAtlasSearchDeploymentApiRequest
	*/
	DeleteAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string) DeleteAtlasSearchDeploymentApiRequest
	/*
		DeleteAtlasSearchDeployment Delete Search Nodes


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteAtlasSearchDeploymentApiParams - Parameters for the request
		@return DeleteAtlasSearchDeploymentApiRequest
	*/
	DeleteAtlasSearchDeploymentWithParams(ctx context.Context, args *DeleteAtlasSearchDeploymentApiParams) DeleteAtlasSearchDeploymentApiRequest

	// Method available only for mocking purposes
	DeleteAtlasSearchDeploymentExecute(r DeleteAtlasSearchDeploymentApiRequest) (*http.Response, error)

	/*
		DeleteAtlasSearchIndex Remove One Atlas Search Index

		Removes one Atlas Search index that you identified with its unique ID. To use this resource, the requesting API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the database and collection with one or more Application Search indexes.
		@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
		@return DeleteAtlasSearchIndexApiRequest
	*/
	DeleteAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string) DeleteAtlasSearchIndexApiRequest
	/*
		DeleteAtlasSearchIndex Remove One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteAtlasSearchIndexApiParams - Parameters for the request
		@return DeleteAtlasSearchIndexApiRequest
	*/
	DeleteAtlasSearchIndexWithParams(ctx context.Context, args *DeleteAtlasSearchIndexApiParams) DeleteAtlasSearchIndexApiRequest

	// Method available only for mocking purposes
	DeleteAtlasSearchIndexExecute(r DeleteAtlasSearchIndexApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetAtlasSearchDeployment Return Search Nodes

		[experimental] Return the Search Nodes for the specified cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Label that identifies the cluster to return the Search Nodes for.
		@return GetAtlasSearchDeploymentApiRequest
	*/
	GetAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string) GetAtlasSearchDeploymentApiRequest
	/*
		GetAtlasSearchDeployment Return Search Nodes


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAtlasSearchDeploymentApiParams - Parameters for the request
		@return GetAtlasSearchDeploymentApiRequest
	*/
	GetAtlasSearchDeploymentWithParams(ctx context.Context, args *GetAtlasSearchDeploymentApiParams) GetAtlasSearchDeploymentApiRequest

	// Method available only for mocking purposes
	GetAtlasSearchDeploymentExecute(r GetAtlasSearchDeploymentApiRequest) (*ApiSearchDeploymentResponse, *http.Response, error)

	/*
		GetAtlasSearchIndex Return One Atlas Search Index

		Returns one Atlas Search index in the specified project. You identify this index using its unique ID. Atlas Search index contains the indexed fields and the analyzers used to create the index. To use this resource, the requesting API Key must have the Project Data Access Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
		@param indexId Unique 24-hexadecimal digit string that identifies the Application Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.
		@return GetAtlasSearchIndexApiRequest
	*/
	GetAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string) GetAtlasSearchIndexApiRequest
	/*
		GetAtlasSearchIndex Return One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAtlasSearchIndexApiParams - Parameters for the request
		@return GetAtlasSearchIndexApiRequest
	*/
	GetAtlasSearchIndexWithParams(ctx context.Context, args *GetAtlasSearchIndexApiParams) GetAtlasSearchIndexApiRequest

	// Method available only for mocking purposes
	GetAtlasSearchIndexExecute(r GetAtlasSearchIndexApiRequest) (*ClusterSearchIndex, *http.Response, error)

	/*
		ListAtlasSearchIndexes Return All Atlas Search Indexes for One Collection

		Returns all Atlas Search indexes on the specified collection. Atlas Search indexes contain the indexed fields and the analyzers used to create the indexes. To use this resource, the requesting API Key must have the Project Data Access Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
		@param collectionName Name of the collection that contains one or more Atlas Search indexes.
		@param databaseName Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.
		@return ListAtlasSearchIndexesApiRequest
	*/
	ListAtlasSearchIndexes(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string) ListAtlasSearchIndexesApiRequest
	/*
		ListAtlasSearchIndexes Return All Atlas Search Indexes for One Collection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAtlasSearchIndexesApiParams - Parameters for the request
		@return ListAtlasSearchIndexesApiRequest
	*/
	ListAtlasSearchIndexesWithParams(ctx context.Context, args *ListAtlasSearchIndexesApiParams) ListAtlasSearchIndexesApiRequest

	// Method available only for mocking purposes
	ListAtlasSearchIndexesExecute(r ListAtlasSearchIndexesApiRequest) ([]ClusterSearchIndex, *http.Response, error)

	/*
		UpdateAtlasSearchDeployment Update Search Nodes

		[experimental] Updates the Search Nodes for the specified cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Label that identifies the cluster to update the Search Nodes for.
		@return UpdateAtlasSearchDeploymentApiRequest
	*/
	UpdateAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string, apiSearchDeploymentRequest *ApiSearchDeploymentRequest) UpdateAtlasSearchDeploymentApiRequest
	/*
		UpdateAtlasSearchDeployment Update Search Nodes


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateAtlasSearchDeploymentApiParams - Parameters for the request
		@return UpdateAtlasSearchDeploymentApiRequest
	*/
	UpdateAtlasSearchDeploymentWithParams(ctx context.Context, args *UpdateAtlasSearchDeploymentApiParams) UpdateAtlasSearchDeploymentApiRequest

	// Method available only for mocking purposes
	UpdateAtlasSearchDeploymentExecute(r UpdateAtlasSearchDeploymentApiRequest) (*ApiSearchDeploymentResponse, *http.Response, error)

	/*
		UpdateAtlasSearchIndex Update One Atlas Search Index

		Updates one Atlas Search index that you identified with its unique ID. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. To use this resource, the requesting API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection whose Atlas Search index to update.
		@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
		@return UpdateAtlasSearchIndexApiRequest
	*/
	UpdateAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string, clusterSearchIndex *ClusterSearchIndex) UpdateAtlasSearchIndexApiRequest
	/*
		UpdateAtlasSearchIndex Update One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateAtlasSearchIndexApiParams - Parameters for the request
		@return UpdateAtlasSearchIndexApiRequest
	*/
	UpdateAtlasSearchIndexWithParams(ctx context.Context, args *UpdateAtlasSearchIndexApiParams) UpdateAtlasSearchIndexApiRequest

	// Method available only for mocking purposes
	UpdateAtlasSearchIndexExecute(r UpdateAtlasSearchIndexApiRequest) (*ClusterSearchIndex, *http.Response, error)
}

// AtlasSearchApiService AtlasSearchApi service
type AtlasSearchApiService service

type CreateAtlasSearchDeploymentApiRequest struct {
	ctx                        context.Context
	ApiService                 AtlasSearchApi
	groupId                    string
	clusterName                string
	apiSearchDeploymentRequest *ApiSearchDeploymentRequest
}

type CreateAtlasSearchDeploymentApiParams struct {
	GroupId                    string
	ClusterName                string
	ApiSearchDeploymentRequest *ApiSearchDeploymentRequest
}

func (a *AtlasSearchApiService) CreateAtlasSearchDeploymentWithParams(ctx context.Context, args *CreateAtlasSearchDeploymentApiParams) CreateAtlasSearchDeploymentApiRequest {
	return CreateAtlasSearchDeploymentApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		clusterName:                args.ClusterName,
		apiSearchDeploymentRequest: args.ApiSearchDeploymentRequest,
	}
}

func (r CreateAtlasSearchDeploymentApiRequest) Execute() (*ApiSearchDeploymentResponse, *http.Response, error) {
	return r.ApiService.CreateAtlasSearchDeploymentExecute(r)
}

/*
CreateAtlasSearchDeployment Create Search Nodes

[experimental] Creates Search Nodes for the specified cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Label that identifies the cluster to create Search Nodes for.
	@return CreateAtlasSearchDeploymentApiRequest
*/
func (a *AtlasSearchApiService) CreateAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string, apiSearchDeploymentRequest *ApiSearchDeploymentRequest) CreateAtlasSearchDeploymentApiRequest {
	return CreateAtlasSearchDeploymentApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		clusterName:                clusterName,
		apiSearchDeploymentRequest: apiSearchDeploymentRequest,
	}
}

// Execute executes the request
//
//	@return ApiSearchDeploymentResponse
func (a *AtlasSearchApiService) CreateAtlasSearchDeploymentExecute(r CreateAtlasSearchDeploymentApiRequest) (*ApiSearchDeploymentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiSearchDeploymentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.CreateAtlasSearchDeployment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiSearchDeploymentRequest == nil {
		return localVarReturnValue, nil, reportError("apiSearchDeploymentRequest is required and must be specified")
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
	localVarPostBody = r.apiSearchDeploymentRequest
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

type CreateAtlasSearchIndexApiRequest struct {
	ctx                context.Context
	ApiService         AtlasSearchApi
	groupId            string
	clusterName        string
	clusterSearchIndex *ClusterSearchIndex
}

type CreateAtlasSearchIndexApiParams struct {
	GroupId            string
	ClusterName        string
	ClusterSearchIndex *ClusterSearchIndex
}

func (a *AtlasSearchApiService) CreateAtlasSearchIndexWithParams(ctx context.Context, args *CreateAtlasSearchIndexApiParams) CreateAtlasSearchIndexApiRequest {
	return CreateAtlasSearchIndexApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		clusterName:        args.ClusterName,
		clusterSearchIndex: args.ClusterSearchIndex,
	}
}

func (r CreateAtlasSearchIndexApiRequest) Execute() (*ClusterSearchIndex, *http.Response, error) {
	return r.ApiService.CreateAtlasSearchIndexExecute(r)
}

/*
CreateAtlasSearchIndex Create One Atlas Search Index

Creates one Atlas Search index on the specified collection. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. Only clusters running MongoDB v4.2 or later can use Atlas Search. To use this resource, the requesting API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection on which to create an Atlas Search index.
	@return CreateAtlasSearchIndexApiRequest
*/
func (a *AtlasSearchApiService) CreateAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, clusterSearchIndex *ClusterSearchIndex) CreateAtlasSearchIndexApiRequest {
	return CreateAtlasSearchIndexApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            groupId,
		clusterName:        clusterName,
		clusterSearchIndex: clusterSearchIndex,
	}
}

// Execute executes the request
//
//	@return ClusterSearchIndex
func (a *AtlasSearchApiService) CreateAtlasSearchIndexExecute(r CreateAtlasSearchIndexApiRequest) (*ClusterSearchIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterSearchIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.CreateAtlasSearchIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clusterSearchIndex == nil {
		return localVarReturnValue, nil, reportError("clusterSearchIndex is required and must be specified")
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
	localVarPostBody = r.clusterSearchIndex
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

type DeleteAtlasSearchDeploymentApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
}

type DeleteAtlasSearchDeploymentApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *AtlasSearchApiService) DeleteAtlasSearchDeploymentWithParams(ctx context.Context, args *DeleteAtlasSearchDeploymentApiParams) DeleteAtlasSearchDeploymentApiRequest {
	return DeleteAtlasSearchDeploymentApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r DeleteAtlasSearchDeploymentApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAtlasSearchDeploymentExecute(r)
}

/*
DeleteAtlasSearchDeployment Delete Search Nodes

[experimental] Deletes the Search Nodes for the specified cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Label that identifies the cluster to delete.
	@return DeleteAtlasSearchDeploymentApiRequest
*/
func (a *AtlasSearchApiService) DeleteAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string) DeleteAtlasSearchDeploymentApiRequest {
	return DeleteAtlasSearchDeploymentApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
func (a *AtlasSearchApiService) DeleteAtlasSearchDeploymentExecute(r DeleteAtlasSearchDeploymentApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.DeleteAtlasSearchDeployment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

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

type DeleteAtlasSearchIndexApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
	indexId     string
}

type DeleteAtlasSearchIndexApiParams struct {
	GroupId     string
	ClusterName string
	IndexId     string
}

func (a *AtlasSearchApiService) DeleteAtlasSearchIndexWithParams(ctx context.Context, args *DeleteAtlasSearchIndexApiParams) DeleteAtlasSearchIndexApiRequest {
	return DeleteAtlasSearchIndexApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		indexId:     args.IndexId,
	}
}

func (r DeleteAtlasSearchIndexApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteAtlasSearchIndexExecute(r)
}

/*
DeleteAtlasSearchIndex Remove One Atlas Search Index

Removes one Atlas Search index that you identified with its unique ID. To use this resource, the requesting API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the database and collection with one or more Application Search indexes.
	@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
	@return DeleteAtlasSearchIndexApiRequest
*/
func (a *AtlasSearchApiService) DeleteAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string) DeleteAtlasSearchIndexApiRequest {
	return DeleteAtlasSearchIndexApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		indexId:     indexId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AtlasSearchApiService) DeleteAtlasSearchIndexExecute(r DeleteAtlasSearchIndexApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.DeleteAtlasSearchIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"indexId"+"}", url.PathEscape(parameterValueToString(r.indexId, "indexId")), -1)

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

type GetAtlasSearchDeploymentApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
}

type GetAtlasSearchDeploymentApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *AtlasSearchApiService) GetAtlasSearchDeploymentWithParams(ctx context.Context, args *GetAtlasSearchDeploymentApiParams) GetAtlasSearchDeploymentApiRequest {
	return GetAtlasSearchDeploymentApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetAtlasSearchDeploymentApiRequest) Execute() (*ApiSearchDeploymentResponse, *http.Response, error) {
	return r.ApiService.GetAtlasSearchDeploymentExecute(r)
}

/*
GetAtlasSearchDeployment Return Search Nodes

[experimental] Return the Search Nodes for the specified cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Label that identifies the cluster to return the Search Nodes for.
	@return GetAtlasSearchDeploymentApiRequest
*/
func (a *AtlasSearchApiService) GetAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string) GetAtlasSearchDeploymentApiRequest {
	return GetAtlasSearchDeploymentApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return ApiSearchDeploymentResponse
func (a *AtlasSearchApiService) GetAtlasSearchDeploymentExecute(r GetAtlasSearchDeploymentApiRequest) (*ApiSearchDeploymentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiSearchDeploymentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.GetAtlasSearchDeployment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

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

type GetAtlasSearchIndexApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
	indexId     string
}

type GetAtlasSearchIndexApiParams struct {
	GroupId     string
	ClusterName string
	IndexId     string
}

func (a *AtlasSearchApiService) GetAtlasSearchIndexWithParams(ctx context.Context, args *GetAtlasSearchIndexApiParams) GetAtlasSearchIndexApiRequest {
	return GetAtlasSearchIndexApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		indexId:     args.IndexId,
	}
}

func (r GetAtlasSearchIndexApiRequest) Execute() (*ClusterSearchIndex, *http.Response, error) {
	return r.ApiService.GetAtlasSearchIndexExecute(r)
}

/*
GetAtlasSearchIndex Return One Atlas Search Index

Returns one Atlas Search index in the specified project. You identify this index using its unique ID. Atlas Search index contains the indexed fields and the analyzers used to create the index. To use this resource, the requesting API Key must have the Project Data Access Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
	@param indexId Unique 24-hexadecimal digit string that identifies the Application Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.
	@return GetAtlasSearchIndexApiRequest
*/
func (a *AtlasSearchApiService) GetAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string) GetAtlasSearchIndexApiRequest {
	return GetAtlasSearchIndexApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		indexId:     indexId,
	}
}

// Execute executes the request
//
//	@return ClusterSearchIndex
func (a *AtlasSearchApiService) GetAtlasSearchIndexExecute(r GetAtlasSearchIndexApiRequest) (*ClusterSearchIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterSearchIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.GetAtlasSearchIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"indexId"+"}", url.PathEscape(parameterValueToString(r.indexId, "indexId")), -1)

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

type ListAtlasSearchIndexesApiRequest struct {
	ctx            context.Context
	ApiService     AtlasSearchApi
	groupId        string
	clusterName    string
	collectionName string
	databaseName   string
}

type ListAtlasSearchIndexesApiParams struct {
	GroupId        string
	ClusterName    string
	CollectionName string
	DatabaseName   string
}

func (a *AtlasSearchApiService) ListAtlasSearchIndexesWithParams(ctx context.Context, args *ListAtlasSearchIndexesApiParams) ListAtlasSearchIndexesApiRequest {
	return ListAtlasSearchIndexesApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		clusterName:    args.ClusterName,
		collectionName: args.CollectionName,
		databaseName:   args.DatabaseName,
	}
}

func (r ListAtlasSearchIndexesApiRequest) Execute() ([]ClusterSearchIndex, *http.Response, error) {
	return r.ApiService.ListAtlasSearchIndexesExecute(r)
}

/*
ListAtlasSearchIndexes Return All Atlas Search Indexes for One Collection

Returns all Atlas Search indexes on the specified collection. Atlas Search indexes contain the indexed fields and the analyzers used to create the indexes. To use this resource, the requesting API Key must have the Project Data Access Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
	@param collectionName Name of the collection that contains one or more Atlas Search indexes.
	@param databaseName Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.
	@return ListAtlasSearchIndexesApiRequest
*/
func (a *AtlasSearchApiService) ListAtlasSearchIndexes(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string) ListAtlasSearchIndexesApiRequest {
	return ListAtlasSearchIndexesApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		clusterName:    clusterName,
		collectionName: collectionName,
		databaseName:   databaseName,
	}
}

// Execute executes the request
//
//	@return []ClusterSearchIndex
func (a *AtlasSearchApiService) ListAtlasSearchIndexesExecute(r ListAtlasSearchIndexesApiRequest) ([]ClusterSearchIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ClusterSearchIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.ListAtlasSearchIndexes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"collectionName"+"}", url.PathEscape(parameterValueToString(r.collectionName, "collectionName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(parameterValueToString(r.databaseName, "databaseName")), -1)

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

type UpdateAtlasSearchDeploymentApiRequest struct {
	ctx                        context.Context
	ApiService                 AtlasSearchApi
	groupId                    string
	clusterName                string
	apiSearchDeploymentRequest *ApiSearchDeploymentRequest
}

type UpdateAtlasSearchDeploymentApiParams struct {
	GroupId                    string
	ClusterName                string
	ApiSearchDeploymentRequest *ApiSearchDeploymentRequest
}

func (a *AtlasSearchApiService) UpdateAtlasSearchDeploymentWithParams(ctx context.Context, args *UpdateAtlasSearchDeploymentApiParams) UpdateAtlasSearchDeploymentApiRequest {
	return UpdateAtlasSearchDeploymentApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		clusterName:                args.ClusterName,
		apiSearchDeploymentRequest: args.ApiSearchDeploymentRequest,
	}
}

func (r UpdateAtlasSearchDeploymentApiRequest) Execute() (*ApiSearchDeploymentResponse, *http.Response, error) {
	return r.ApiService.UpdateAtlasSearchDeploymentExecute(r)
}

/*
UpdateAtlasSearchDeployment Update Search Nodes

[experimental] Updates the Search Nodes for the specified cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Label that identifies the cluster to update the Search Nodes for.
	@return UpdateAtlasSearchDeploymentApiRequest
*/
func (a *AtlasSearchApiService) UpdateAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string, apiSearchDeploymentRequest *ApiSearchDeploymentRequest) UpdateAtlasSearchDeploymentApiRequest {
	return UpdateAtlasSearchDeploymentApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		clusterName:                clusterName,
		apiSearchDeploymentRequest: apiSearchDeploymentRequest,
	}
}

// Execute executes the request
//
//	@return ApiSearchDeploymentResponse
func (a *AtlasSearchApiService) UpdateAtlasSearchDeploymentExecute(r UpdateAtlasSearchDeploymentApiRequest) (*ApiSearchDeploymentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiSearchDeploymentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.UpdateAtlasSearchDeployment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiSearchDeploymentRequest == nil {
		return localVarReturnValue, nil, reportError("apiSearchDeploymentRequest is required and must be specified")
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
	localVarPostBody = r.apiSearchDeploymentRequest
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

type UpdateAtlasSearchIndexApiRequest struct {
	ctx                context.Context
	ApiService         AtlasSearchApi
	groupId            string
	clusterName        string
	indexId            string
	clusterSearchIndex *ClusterSearchIndex
}

type UpdateAtlasSearchIndexApiParams struct {
	GroupId            string
	ClusterName        string
	IndexId            string
	ClusterSearchIndex *ClusterSearchIndex
}

func (a *AtlasSearchApiService) UpdateAtlasSearchIndexWithParams(ctx context.Context, args *UpdateAtlasSearchIndexApiParams) UpdateAtlasSearchIndexApiRequest {
	return UpdateAtlasSearchIndexApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		clusterName:        args.ClusterName,
		indexId:            args.IndexId,
		clusterSearchIndex: args.ClusterSearchIndex,
	}
}

func (r UpdateAtlasSearchIndexApiRequest) Execute() (*ClusterSearchIndex, *http.Response, error) {
	return r.ApiService.UpdateAtlasSearchIndexExecute(r)
}

/*
UpdateAtlasSearchIndex Update One Atlas Search Index

Updates one Atlas Search index that you identified with its unique ID. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. To use this resource, the requesting API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection whose Atlas Search index to update.
	@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
	@return UpdateAtlasSearchIndexApiRequest
*/
func (a *AtlasSearchApiService) UpdateAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string, clusterSearchIndex *ClusterSearchIndex) UpdateAtlasSearchIndexApiRequest {
	return UpdateAtlasSearchIndexApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            groupId,
		clusterName:        clusterName,
		indexId:            indexId,
		clusterSearchIndex: clusterSearchIndex,
	}
}

// Execute executes the request
//
//	@return ClusterSearchIndex
func (a *AtlasSearchApiService) UpdateAtlasSearchIndexExecute(r UpdateAtlasSearchIndexApiRequest) (*ClusterSearchIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterSearchIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.UpdateAtlasSearchIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"indexId"+"}", url.PathEscape(parameterValueToString(r.indexId, "indexId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clusterSearchIndex == nil {
		return localVarReturnValue, nil, reportError("clusterSearchIndex is required and must be specified")
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
	localVarPostBody = r.clusterSearchIndex
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
