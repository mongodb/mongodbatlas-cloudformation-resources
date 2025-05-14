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

type GlobalClustersApi interface {

	/*
		CreateCustomZoneMapping Add One Entry to One Custom Zone Mapping

		[experimental] Creates one custom zone mapping for the specified global cluster. A custom zone mapping matches one ISO 3166-2 location code to a zone in your global cluster. By default, MongoDB Cloud maps each location code to the closest geographical zone. To use this resource, the requesting API Key must have the Project Owner role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies this cluster.
		@return CreateCustomZoneMappingApiRequest
	*/
	CreateCustomZoneMapping(ctx context.Context, groupId string, clusterName string, customZoneMappings *CustomZoneMappings) CreateCustomZoneMappingApiRequest
	/*
		CreateCustomZoneMapping Add One Entry to One Custom Zone Mapping


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateCustomZoneMappingApiParams - Parameters for the request
		@return CreateCustomZoneMappingApiRequest
	*/
	CreateCustomZoneMappingWithParams(ctx context.Context, args *CreateCustomZoneMappingApiParams) CreateCustomZoneMappingApiRequest

	// Method available only for mocking purposes
	CreateCustomZoneMappingExecute(r CreateCustomZoneMappingApiRequest) (*GeoSharding, *http.Response, error)

	/*
		CreateManagedNamespace Create One Managed Namespace in One Global Cluster

		[experimental] Creates one managed namespace within the specified global cluster. A managed namespace identifies a collection using the database name, the dot separator, and the collection name. To use this resource, the requesting API Key must have the Project Data Access Admin role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies this cluster.
		@return CreateManagedNamespaceApiRequest
	*/
	CreateManagedNamespace(ctx context.Context, groupId string, clusterName string, managedNamespace *ManagedNamespace) CreateManagedNamespaceApiRequest
	/*
		CreateManagedNamespace Create One Managed Namespace in One Global Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateManagedNamespaceApiParams - Parameters for the request
		@return CreateManagedNamespaceApiRequest
	*/
	CreateManagedNamespaceWithParams(ctx context.Context, args *CreateManagedNamespaceApiParams) CreateManagedNamespaceApiRequest

	// Method available only for mocking purposes
	CreateManagedNamespaceExecute(r CreateManagedNamespaceApiRequest) (*GeoSharding, *http.Response, error)

	/*
		DeleteAllCustomZoneMappings Remove All Custom Zone Mappings from One Global Cluster

		[experimental] Removes all custom zone mappings for the specified global cluster. A custom zone mapping matches one ISO 3166-2 location code to a zone in your global cluster. Removing the custom zone mappings restores the default mapping. By default, MongoDB Cloud maps each location code to the closest geographical zone. To use this resource, the requesting API Key must have the Project Owner role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies this cluster.
		@return DeleteAllCustomZoneMappingsApiRequest
	*/
	DeleteAllCustomZoneMappings(ctx context.Context, groupId string, clusterName string) DeleteAllCustomZoneMappingsApiRequest
	/*
		DeleteAllCustomZoneMappings Remove All Custom Zone Mappings from One Global Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteAllCustomZoneMappingsApiParams - Parameters for the request
		@return DeleteAllCustomZoneMappingsApiRequest
	*/
	DeleteAllCustomZoneMappingsWithParams(ctx context.Context, args *DeleteAllCustomZoneMappingsApiParams) DeleteAllCustomZoneMappingsApiRequest

	// Method available only for mocking purposes
	DeleteAllCustomZoneMappingsExecute(r DeleteAllCustomZoneMappingsApiRequest) (*GeoSharding, *http.Response, error)

	/*
		DeleteManagedNamespace Remove One Managed Namespace from One Global Cluster

		[experimental] Removes one managed namespace within the specified global cluster. A managed namespace identifies a collection using the database name, the dot separator, and the collection name. Deleting a managed namespace does not remove the associated collection or data. To use this resource, the requesting API Key must have the Project Data Access Admin role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param clusterName Human-readable label that identifies this cluster.
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeleteManagedNamespaceApiRequest
	*/
	DeleteManagedNamespace(ctx context.Context, clusterName string, groupId string) DeleteManagedNamespaceApiRequest
	/*
		DeleteManagedNamespace Remove One Managed Namespace from One Global Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteManagedNamespaceApiParams - Parameters for the request
		@return DeleteManagedNamespaceApiRequest
	*/
	DeleteManagedNamespaceWithParams(ctx context.Context, args *DeleteManagedNamespaceApiParams) DeleteManagedNamespaceApiRequest

	// Method available only for mocking purposes
	DeleteManagedNamespaceExecute(r DeleteManagedNamespaceApiRequest) (*GeoSharding, *http.Response, error)

	/*
		GetManagedNamespace Return One Managed Namespace in One Global Cluster

		Returns one managed namespace within the specified global cluster. A managed namespace identifies a collection using the database name, the dot separator, and the collection name. To use this resource, the requesting API Key must have the Project Read Only role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies this cluster.
		@return GetManagedNamespaceApiRequest
	*/
	GetManagedNamespace(ctx context.Context, groupId string, clusterName string) GetManagedNamespaceApiRequest
	/*
		GetManagedNamespace Return One Managed Namespace in One Global Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetManagedNamespaceApiParams - Parameters for the request
		@return GetManagedNamespaceApiRequest
	*/
	GetManagedNamespaceWithParams(ctx context.Context, args *GetManagedNamespaceApiParams) GetManagedNamespaceApiRequest

	// Method available only for mocking purposes
	GetManagedNamespaceExecute(r GetManagedNamespaceApiRequest) (*GeoSharding, *http.Response, error)
}

// GlobalClustersApiService GlobalClustersApi service
type GlobalClustersApiService service

type CreateCustomZoneMappingApiRequest struct {
	ctx                context.Context
	ApiService         GlobalClustersApi
	groupId            string
	clusterName        string
	customZoneMappings *CustomZoneMappings
}

type CreateCustomZoneMappingApiParams struct {
	GroupId            string
	ClusterName        string
	CustomZoneMappings *CustomZoneMappings
}

func (a *GlobalClustersApiService) CreateCustomZoneMappingWithParams(ctx context.Context, args *CreateCustomZoneMappingApiParams) CreateCustomZoneMappingApiRequest {
	return CreateCustomZoneMappingApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		clusterName:        args.ClusterName,
		customZoneMappings: args.CustomZoneMappings,
	}
}

func (r CreateCustomZoneMappingApiRequest) Execute() (*GeoSharding, *http.Response, error) {
	return r.ApiService.CreateCustomZoneMappingExecute(r)
}

/*
CreateCustomZoneMapping Add One Entry to One Custom Zone Mapping

[experimental] Creates one custom zone mapping for the specified global cluster. A custom zone mapping matches one ISO 3166-2 location code to a zone in your global cluster. By default, MongoDB Cloud maps each location code to the closest geographical zone. To use this resource, the requesting API Key must have the Project Owner role. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies this cluster.
	@return CreateCustomZoneMappingApiRequest
*/
func (a *GlobalClustersApiService) CreateCustomZoneMapping(ctx context.Context, groupId string, clusterName string, customZoneMappings *CustomZoneMappings) CreateCustomZoneMappingApiRequest {
	return CreateCustomZoneMappingApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            groupId,
		clusterName:        clusterName,
		customZoneMappings: customZoneMappings,
	}
}

// Execute executes the request
//
//	@return GeoSharding
func (a *GlobalClustersApiService) CreateCustomZoneMappingExecute(r CreateCustomZoneMappingApiRequest) (*GeoSharding, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GeoSharding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalClustersApiService.CreateCustomZoneMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customZoneMappings == nil {
		return localVarReturnValue, nil, reportError("customZoneMappings is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-02-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.customZoneMappings
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

type CreateManagedNamespaceApiRequest struct {
	ctx              context.Context
	ApiService       GlobalClustersApi
	groupId          string
	clusterName      string
	managedNamespace *ManagedNamespace
}

type CreateManagedNamespaceApiParams struct {
	GroupId          string
	ClusterName      string
	ManagedNamespace *ManagedNamespace
}

func (a *GlobalClustersApiService) CreateManagedNamespaceWithParams(ctx context.Context, args *CreateManagedNamespaceApiParams) CreateManagedNamespaceApiRequest {
	return CreateManagedNamespaceApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          args.GroupId,
		clusterName:      args.ClusterName,
		managedNamespace: args.ManagedNamespace,
	}
}

func (r CreateManagedNamespaceApiRequest) Execute() (*GeoSharding, *http.Response, error) {
	return r.ApiService.CreateManagedNamespaceExecute(r)
}

/*
CreateManagedNamespace Create One Managed Namespace in One Global Cluster

[experimental] Creates one managed namespace within the specified global cluster. A managed namespace identifies a collection using the database name, the dot separator, and the collection name. To use this resource, the requesting API Key must have the Project Data Access Admin role. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies this cluster.
	@return CreateManagedNamespaceApiRequest
*/
func (a *GlobalClustersApiService) CreateManagedNamespace(ctx context.Context, groupId string, clusterName string, managedNamespace *ManagedNamespace) CreateManagedNamespaceApiRequest {
	return CreateManagedNamespaceApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		clusterName:      clusterName,
		managedNamespace: managedNamespace,
	}
}

// Execute executes the request
//
//	@return GeoSharding
func (a *GlobalClustersApiService) CreateManagedNamespaceExecute(r CreateManagedNamespaceApiRequest) (*GeoSharding, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GeoSharding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalClustersApiService.CreateManagedNamespace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.managedNamespace == nil {
		return localVarReturnValue, nil, reportError("managedNamespace is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-02-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.managedNamespace
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

type DeleteAllCustomZoneMappingsApiRequest struct {
	ctx         context.Context
	ApiService  GlobalClustersApi
	groupId     string
	clusterName string
}

type DeleteAllCustomZoneMappingsApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *GlobalClustersApiService) DeleteAllCustomZoneMappingsWithParams(ctx context.Context, args *DeleteAllCustomZoneMappingsApiParams) DeleteAllCustomZoneMappingsApiRequest {
	return DeleteAllCustomZoneMappingsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r DeleteAllCustomZoneMappingsApiRequest) Execute() (*GeoSharding, *http.Response, error) {
	return r.ApiService.DeleteAllCustomZoneMappingsExecute(r)
}

/*
DeleteAllCustomZoneMappings Remove All Custom Zone Mappings from One Global Cluster

[experimental] Removes all custom zone mappings for the specified global cluster. A custom zone mapping matches one ISO 3166-2 location code to a zone in your global cluster. Removing the custom zone mappings restores the default mapping. By default, MongoDB Cloud maps each location code to the closest geographical zone. To use this resource, the requesting API Key must have the Project Owner role. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies this cluster.
	@return DeleteAllCustomZoneMappingsApiRequest
*/
func (a *GlobalClustersApiService) DeleteAllCustomZoneMappings(ctx context.Context, groupId string, clusterName string) DeleteAllCustomZoneMappingsApiRequest {
	return DeleteAllCustomZoneMappingsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return GeoSharding
func (a *GlobalClustersApiService) DeleteAllCustomZoneMappingsExecute(r DeleteAllCustomZoneMappingsApiRequest) (*GeoSharding, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GeoSharding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalClustersApiService.DeleteAllCustomZoneMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/customZoneMapping"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

type DeleteManagedNamespaceApiRequest struct {
	ctx         context.Context
	ApiService  GlobalClustersApi
	clusterName string
	groupId     string
	db          *string
	collection  *string
}

type DeleteManagedNamespaceApiParams struct {
	ClusterName string
	GroupId     string
	Db          *string
	Collection  *string
}

func (a *GlobalClustersApiService) DeleteManagedNamespaceWithParams(ctx context.Context, args *DeleteManagedNamespaceApiParams) DeleteManagedNamespaceApiRequest {
	return DeleteManagedNamespaceApiRequest{
		ApiService:  a,
		ctx:         ctx,
		clusterName: args.ClusterName,
		groupId:     args.GroupId,
		db:          args.Db,
		collection:  args.Collection,
	}
}

// Human-readable label that identifies the database that contains the collection.
func (r DeleteManagedNamespaceApiRequest) Db(db string) DeleteManagedNamespaceApiRequest {
	r.db = &db
	return r
}

// Human-readable label that identifies the collection associated with the managed namespace.
func (r DeleteManagedNamespaceApiRequest) Collection(collection string) DeleteManagedNamespaceApiRequest {
	r.collection = &collection
	return r
}

func (r DeleteManagedNamespaceApiRequest) Execute() (*GeoSharding, *http.Response, error) {
	return r.ApiService.DeleteManagedNamespaceExecute(r)
}

/*
DeleteManagedNamespace Remove One Managed Namespace from One Global Cluster

[experimental] Removes one managed namespace within the specified global cluster. A managed namespace identifies a collection using the database name, the dot separator, and the collection name. Deleting a managed namespace does not remove the associated collection or data. To use this resource, the requesting API Key must have the Project Data Access Admin role. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clusterName Human-readable label that identifies this cluster.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeleteManagedNamespaceApiRequest
*/
func (a *GlobalClustersApiService) DeleteManagedNamespace(ctx context.Context, clusterName string, groupId string) DeleteManagedNamespaceApiRequest {
	return DeleteManagedNamespaceApiRequest{
		ApiService:  a,
		ctx:         ctx,
		clusterName: clusterName,
		groupId:     groupId,
	}
}

// Execute executes the request
//
//	@return GeoSharding
func (a *GlobalClustersApiService) DeleteManagedNamespaceExecute(r DeleteManagedNamespaceApiRequest) (*GeoSharding, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GeoSharding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalClustersApiService.DeleteManagedNamespace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites/managedNamespaces"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.db != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "db", r.db, "")
	}
	if r.collection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "collection", r.collection, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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

type GetManagedNamespaceApiRequest struct {
	ctx         context.Context
	ApiService  GlobalClustersApi
	groupId     string
	clusterName string
}

type GetManagedNamespaceApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *GlobalClustersApiService) GetManagedNamespaceWithParams(ctx context.Context, args *GetManagedNamespaceApiParams) GetManagedNamespaceApiRequest {
	return GetManagedNamespaceApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetManagedNamespaceApiRequest) Execute() (*GeoSharding, *http.Response, error) {
	return r.ApiService.GetManagedNamespaceExecute(r)
}

/*
GetManagedNamespace Return One Managed Namespace in One Global Cluster

Returns one managed namespace within the specified global cluster. A managed namespace identifies a collection using the database name, the dot separator, and the collection name. To use this resource, the requesting API Key must have the Project Read Only role. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies this cluster.
	@return GetManagedNamespaceApiRequest
*/
func (a *GlobalClustersApiService) GetManagedNamespace(ctx context.Context, groupId string, clusterName string) GetManagedNamespaceApiRequest {
	return GetManagedNamespaceApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return GeoSharding
func (a *GlobalClustersApiService) GetManagedNamespaceExecute(r GetManagedNamespaceApiRequest) (*GeoSharding, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GeoSharding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalClustersApiService.GetManagedNamespace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/globalWrites"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json", "application/json"}

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
