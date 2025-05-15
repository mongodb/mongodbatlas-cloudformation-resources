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
	"reflect"
	"strings"
)

type ClustersApi interface {

	/*
		CreateCluster Create One Cluster from One Project

		Creates one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can create multi-cloud clusters. Each project supports up to 25 database deployments. To use this resource, the requesting API Key must have the Project Owner role. This feature is not available for serverless clusters. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return CreateClusterApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	CreateCluster(ctx context.Context, groupId string, advancedClusterDescription *AdvancedClusterDescription) CreateClusterApiRequest
	/*
		CreateCluster Create One Cluster from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateClusterApiParams - Parameters for the request
		@return CreateClusterApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	CreateClusterWithParams(ctx context.Context, args *CreateClusterApiParams) CreateClusterApiRequest

	// Method available only for mocking purposes
	CreateClusterExecute(r CreateClusterApiRequest) (*AdvancedClusterDescription, *http.Response, error)

	/*
		DeleteCluster Remove One Cluster from One Project

		Removes one cluster from the specified project. The cluster must have termination protection disabled in order to be deleted. To use this resource, the requesting API Key must have the Project Owner role. This feature is not available for serverless clusters. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return DeleteClusterApiRequest
	*/
	DeleteCluster(ctx context.Context, groupId string, clusterName string) DeleteClusterApiRequest
	/*
		DeleteCluster Remove One Cluster from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteClusterApiParams - Parameters for the request
		@return DeleteClusterApiRequest
	*/
	DeleteClusterWithParams(ctx context.Context, args *DeleteClusterApiParams) DeleteClusterApiRequest

	// Method available only for mocking purposes
	DeleteClusterExecute(r DeleteClusterApiRequest) (*http.Response, error)

	/*
		GetCluster Return One Cluster from One Project

		Returns the details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. The response includes multi-cloud clusters. To use this resource, the requesting API Key must have the Project Read Only role. This feature is not available for serverless clusters. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies this cluster.
		@return GetClusterApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	GetCluster(ctx context.Context, groupId string, clusterName string) GetClusterApiRequest
	/*
		GetCluster Return One Cluster from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetClusterApiParams - Parameters for the request
		@return GetClusterApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	GetClusterWithParams(ctx context.Context, args *GetClusterApiParams) GetClusterApiRequest

	// Method available only for mocking purposes
	GetClusterExecute(r GetClusterApiRequest) (*AdvancedClusterDescription, *http.Response, error)

	/*
		GetClusterAdvancedConfiguration Return One Advanced Configuration Options for One Cluster

		Returns the advanced configuration details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. Advanced configuration details include the read/write concern, index and oplog limits, and other database settings. This feature isn't available for `M0` free clusters, `M2` and `M5` shared-tier clusters, or serverless clusters. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return GetClusterAdvancedConfigurationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	GetClusterAdvancedConfiguration(ctx context.Context, groupId string, clusterName string) GetClusterAdvancedConfigurationApiRequest
	/*
		GetClusterAdvancedConfiguration Return One Advanced Configuration Options for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetClusterAdvancedConfigurationApiParams - Parameters for the request
		@return GetClusterAdvancedConfigurationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	GetClusterAdvancedConfigurationWithParams(ctx context.Context, args *GetClusterAdvancedConfigurationApiParams) GetClusterAdvancedConfigurationApiRequest

	// Method available only for mocking purposes
	GetClusterAdvancedConfigurationExecute(r GetClusterAdvancedConfigurationApiRequest) (*ClusterDescriptionProcessArgs, *http.Response, error)

	/*
		GetClusterStatus Return Status of All Cluster Operations

		[experimental] Returns the status of all changes that you made to the specified cluster in the specified project. Use this resource to check the progress MongoDB Cloud has made in processing your changes. The response does not include the deployment of new dedicated clusters. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return GetClusterStatusApiRequest
	*/
	GetClusterStatus(ctx context.Context, groupId string, clusterName string) GetClusterStatusApiRequest
	/*
		GetClusterStatus Return Status of All Cluster Operations


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetClusterStatusApiParams - Parameters for the request
		@return GetClusterStatusApiRequest
	*/
	GetClusterStatusWithParams(ctx context.Context, args *GetClusterStatusApiParams) GetClusterStatusApiRequest

	// Method available only for mocking purposes
	GetClusterStatusExecute(r GetClusterStatusApiRequest) (*ClusterStatus, *http.Response, error)

	/*
		GetSampleDatasetLoadStatus Check Status of Cluster Sample Dataset Request

		Checks the progress of loading the sample dataset into one cluster. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param sampleDatasetId Unique 24-hexadecimal digit string that identifies the loaded sample dataset.
		@return GetSampleDatasetLoadStatusApiRequest
	*/
	GetSampleDatasetLoadStatus(ctx context.Context, groupId string, sampleDatasetId string) GetSampleDatasetLoadStatusApiRequest
	/*
		GetSampleDatasetLoadStatus Check Status of Cluster Sample Dataset Request


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetSampleDatasetLoadStatusApiParams - Parameters for the request
		@return GetSampleDatasetLoadStatusApiRequest
	*/
	GetSampleDatasetLoadStatusWithParams(ctx context.Context, args *GetSampleDatasetLoadStatusApiParams) GetSampleDatasetLoadStatusApiRequest

	// Method available only for mocking purposes
	GetSampleDatasetLoadStatusExecute(r GetSampleDatasetLoadStatusApiRequest) (*SampleDatasetStatus, *http.Response, error)

	/*
		ListCloudProviderRegions Return All Cloud Provider Regions

		Returns the list of regions available for the specified cloud provider at the specified tier. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListCloudProviderRegionsApiRequest
	*/
	ListCloudProviderRegions(ctx context.Context, groupId string) ListCloudProviderRegionsApiRequest
	/*
		ListCloudProviderRegions Return All Cloud Provider Regions


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListCloudProviderRegionsApiParams - Parameters for the request
		@return ListCloudProviderRegionsApiRequest
	*/
	ListCloudProviderRegionsWithParams(ctx context.Context, args *ListCloudProviderRegionsApiParams) ListCloudProviderRegionsApiRequest

	// Method available only for mocking purposes
	ListCloudProviderRegionsExecute(r ListCloudProviderRegionsApiRequest) (*PaginatedApiAtlasProviderRegions, *http.Response, error)

	/*
		ListClusters Return All Clusters in One Project

		Returns the details for all clusters in the specific project to which you have access. Clusters contain a group of hosts that maintain the same data set. The response includes multi-cloud clusters. To use this resource, the requesting API Key must have the Project Read Only role. This feature is not available for serverless clusters. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListClustersApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	ListClusters(ctx context.Context, groupId string) ListClustersApiRequest
	/*
		ListClusters Return All Clusters in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListClustersApiParams - Parameters for the request
		@return ListClustersApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	ListClustersWithParams(ctx context.Context, args *ListClustersApiParams) ListClustersApiRequest

	// Method available only for mocking purposes
	ListClustersExecute(r ListClustersApiRequest) (*PaginatedAdvancedClusterDescription, *http.Response, error)

	/*
		ListClustersForAllProjects Return All Authorized Clusters in All Projects

		[experimental] Returns the details for all clusters in all projects to which you have access. Clusters contain a group of hosts that maintain the same data set. The response does not include multi-cloud clusters. To use this resource, the requesting API Key can have any cluster-level role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ListClustersForAllProjectsApiRequest
	*/
	ListClustersForAllProjects(ctx context.Context) ListClustersForAllProjectsApiRequest
	/*
		ListClustersForAllProjects Return All Authorized Clusters in All Projects


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListClustersForAllProjectsApiParams - Parameters for the request
		@return ListClustersForAllProjectsApiRequest
	*/
	ListClustersForAllProjectsWithParams(ctx context.Context, args *ListClustersForAllProjectsApiParams) ListClustersForAllProjectsApiRequest

	// Method available only for mocking purposes
	ListClustersForAllProjectsExecute(r ListClustersForAllProjectsApiRequest) (*PaginatedOrgGroup, *http.Response, error)

	/*
		LoadSampleDataset Load Sample Dataset Request into Cluster

		Requests loading the MongoDB sample dataset into the specified cluster. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the cluster into which you load the sample dataset.
		@return LoadSampleDatasetApiRequest
	*/
	LoadSampleDataset(ctx context.Context, groupId string, name string) LoadSampleDatasetApiRequest
	/*
		LoadSampleDataset Load Sample Dataset Request into Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param LoadSampleDatasetApiParams - Parameters for the request
		@return LoadSampleDatasetApiRequest
	*/
	LoadSampleDatasetWithParams(ctx context.Context, args *LoadSampleDatasetApiParams) LoadSampleDatasetApiRequest

	// Method available only for mocking purposes
	LoadSampleDatasetExecute(r LoadSampleDatasetApiRequest) (*SampleDatasetStatus, *http.Response, error)

	/*
		TestFailover Test Failover for One Cluster

		Starts a failover test for the specified cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. A failover test checks how MongoDB Cloud handles the failure of the cluster's primary node. During the test, MongoDB Cloud shuts down the primary node and elects a new primary. To use this resource, the requesting API Key must have the Project Cluster Manager role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return TestFailoverApiRequest
	*/
	TestFailover(ctx context.Context, groupId string, clusterName string) TestFailoverApiRequest
	/*
		TestFailover Test Failover for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param TestFailoverApiParams - Parameters for the request
		@return TestFailoverApiRequest
	*/
	TestFailoverWithParams(ctx context.Context, args *TestFailoverApiParams) TestFailoverApiRequest

	// Method available only for mocking purposes
	TestFailoverExecute(r TestFailoverApiRequest) (*http.Response, error)

	/*
		UpdateCluster Modify One Cluster from One Project

		Updates the details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can update multi-cloud clusters. To update a cluster's termination protection, the requesting API Key must have the Project Owner role. For all other updates, the requesting API Key must have the Project Cluster Manager role. You can't modify a paused cluster (`paused : true`). You must call this endpoint to set `paused : false`. After this endpoint responds with `paused : false`, you can call it again with the changes you want to make to the cluster. This feature is not available for serverless clusters. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return UpdateClusterApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	UpdateCluster(ctx context.Context, groupId string, clusterName string, advancedClusterDescription *AdvancedClusterDescription) UpdateClusterApiRequest
	/*
		UpdateCluster Modify One Cluster from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateClusterApiParams - Parameters for the request
		@return UpdateClusterApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	UpdateClusterWithParams(ctx context.Context, args *UpdateClusterApiParams) UpdateClusterApiRequest

	// Method available only for mocking purposes
	UpdateClusterExecute(r UpdateClusterApiRequest) (*AdvancedClusterDescription, *http.Response, error)

	/*
		UpdateClusterAdvancedConfiguration Update Advanced Configuration Options for One Cluster

		Updates the advanced configuration details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. Advanced configuration details include the read/write concern, index and oplog limits, and other database settings. To use this resource, the requesting API Key must have the Project Cluster Manager role. This feature isn't available for `M0` free clusters, `M2` and `M5` shared-tier clusters, or serverless clusters.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return UpdateClusterAdvancedConfigurationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	UpdateClusterAdvancedConfiguration(ctx context.Context, groupId string, clusterName string, clusterDescriptionProcessArgs *ClusterDescriptionProcessArgs) UpdateClusterAdvancedConfigurationApiRequest
	/*
		UpdateClusterAdvancedConfiguration Update Advanced Configuration Options for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateClusterAdvancedConfigurationApiParams - Parameters for the request
		@return UpdateClusterAdvancedConfigurationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	UpdateClusterAdvancedConfigurationWithParams(ctx context.Context, args *UpdateClusterAdvancedConfigurationApiParams) UpdateClusterAdvancedConfigurationApiRequest

	// Method available only for mocking purposes
	UpdateClusterAdvancedConfigurationExecute(r UpdateClusterAdvancedConfigurationApiRequest) (*ClusterDescriptionProcessArgs, *http.Response, error)

	/*
		UpgradeSharedCluster Upgrade One Shared-tier Cluster

		[experimental] Upgrades a shared-tier cluster in the specified project. To use this resource, the requesting API key must have the Project Cluster Manager role. Each project supports up to 25 clusters.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return UpgradeSharedClusterApiRequest
	*/
	UpgradeSharedCluster(ctx context.Context, groupId string, legacyAtlasTenantClusterUpgradeRequest *LegacyAtlasTenantClusterUpgradeRequest) UpgradeSharedClusterApiRequest
	/*
		UpgradeSharedCluster Upgrade One Shared-tier Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpgradeSharedClusterApiParams - Parameters for the request
		@return UpgradeSharedClusterApiRequest
	*/
	UpgradeSharedClusterWithParams(ctx context.Context, args *UpgradeSharedClusterApiParams) UpgradeSharedClusterApiRequest

	// Method available only for mocking purposes
	UpgradeSharedClusterExecute(r UpgradeSharedClusterApiRequest) (*LegacyAtlasCluster, *http.Response, error)

	/*
		UpgradeSharedClusterToServerless Upgrades One Shared-Tier Cluster to the Serverless Instance

		[experimental] Upgrades a shared-tier cluster to a serverless instance in the specified project. To use this resource, the requesting API key must have the Project Cluster Manager role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return UpgradeSharedClusterToServerlessApiRequest
	*/
	UpgradeSharedClusterToServerless(ctx context.Context, groupId string, serverlessInstanceDescription *ServerlessInstanceDescription) UpgradeSharedClusterToServerlessApiRequest
	/*
		UpgradeSharedClusterToServerless Upgrades One Shared-Tier Cluster to the Serverless Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpgradeSharedClusterToServerlessApiParams - Parameters for the request
		@return UpgradeSharedClusterToServerlessApiRequest
	*/
	UpgradeSharedClusterToServerlessWithParams(ctx context.Context, args *UpgradeSharedClusterToServerlessApiParams) UpgradeSharedClusterToServerlessApiRequest

	// Method available only for mocking purposes
	UpgradeSharedClusterToServerlessExecute(r UpgradeSharedClusterToServerlessApiRequest) (*ServerlessInstanceDescription, *http.Response, error)
}

// ClustersApiService ClustersApi service
type ClustersApiService service

type CreateClusterApiRequest struct {
	ctx                        context.Context
	ApiService                 ClustersApi
	groupId                    string
	advancedClusterDescription *AdvancedClusterDescription
}

type CreateClusterApiParams struct {
	GroupId                    string
	AdvancedClusterDescription *AdvancedClusterDescription
}

func (a *ClustersApiService) CreateClusterWithParams(ctx context.Context, args *CreateClusterApiParams) CreateClusterApiRequest {
	return CreateClusterApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		advancedClusterDescription: args.AdvancedClusterDescription,
	}
}

func (r CreateClusterApiRequest) Execute() (*AdvancedClusterDescription, *http.Response, error) {
	return r.ApiService.CreateClusterExecute(r)
}

/*
CreateCluster Create One Cluster from One Project

Creates one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can create multi-cloud clusters. Each project supports up to 25 database deployments. To use this resource, the requesting API Key must have the Project Owner role. This feature is not available for serverless clusters. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateClusterApiRequest

Deprecated
*/
func (a *ClustersApiService) CreateCluster(ctx context.Context, groupId string, advancedClusterDescription *AdvancedClusterDescription) CreateClusterApiRequest {
	return CreateClusterApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		advancedClusterDescription: advancedClusterDescription,
	}
}

// Execute executes the request
//
//	@return AdvancedClusterDescription
//
// Deprecated
func (a *ClustersApiService) CreateClusterExecute(r CreateClusterApiRequest) (*AdvancedClusterDescription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AdvancedClusterDescription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.CreateCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advancedClusterDescription == nil {
		return localVarReturnValue, nil, reportError("advancedClusterDescription is required and must be specified")
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
	localVarPostBody = r.advancedClusterDescription
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

type DeleteClusterApiRequest struct {
	ctx           context.Context
	ApiService    ClustersApi
	groupId       string
	clusterName   string
	retainBackups *bool
}

type DeleteClusterApiParams struct {
	GroupId       string
	ClusterName   string
	RetainBackups *bool
}

func (a *ClustersApiService) DeleteClusterWithParams(ctx context.Context, args *DeleteClusterApiParams) DeleteClusterApiRequest {
	return DeleteClusterApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		clusterName:   args.ClusterName,
		retainBackups: args.RetainBackups,
	}
}

// Flag that indicates whether to retain backup snapshots for the deleted dedicated cluster.
func (r DeleteClusterApiRequest) RetainBackups(retainBackups bool) DeleteClusterApiRequest {
	r.retainBackups = &retainBackups
	return r
}

func (r DeleteClusterApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteClusterExecute(r)
}

/*
DeleteCluster Remove One Cluster from One Project

Removes one cluster from the specified project. The cluster must have termination protection disabled in order to be deleted. To use this resource, the requesting API Key must have the Project Owner role. This feature is not available for serverless clusters. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return DeleteClusterApiRequest
*/
func (a *ClustersApiService) DeleteCluster(ctx context.Context, groupId string, clusterName string) DeleteClusterApiRequest {
	return DeleteClusterApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
func (a *ClustersApiService) DeleteClusterExecute(r DeleteClusterApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.DeleteCluster")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.retainBackups != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retainBackups", r.retainBackups, "")
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

type GetClusterApiRequest struct {
	ctx         context.Context
	ApiService  ClustersApi
	groupId     string
	clusterName string
}

type GetClusterApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClustersApiService) GetClusterWithParams(ctx context.Context, args *GetClusterApiParams) GetClusterApiRequest {
	return GetClusterApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetClusterApiRequest) Execute() (*AdvancedClusterDescription, *http.Response, error) {
	return r.ApiService.GetClusterExecute(r)
}

/*
GetCluster Return One Cluster from One Project

Returns the details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. The response includes multi-cloud clusters. To use this resource, the requesting API Key must have the Project Read Only role. This feature is not available for serverless clusters. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies this cluster.
	@return GetClusterApiRequest

Deprecated
*/
func (a *ClustersApiService) GetCluster(ctx context.Context, groupId string, clusterName string) GetClusterApiRequest {
	return GetClusterApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return AdvancedClusterDescription
//
// Deprecated
func (a *ClustersApiService) GetClusterExecute(r GetClusterApiRequest) (*AdvancedClusterDescription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AdvancedClusterDescription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.GetCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}"
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

type GetClusterAdvancedConfigurationApiRequest struct {
	ctx         context.Context
	ApiService  ClustersApi
	groupId     string
	clusterName string
}

type GetClusterAdvancedConfigurationApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClustersApiService) GetClusterAdvancedConfigurationWithParams(ctx context.Context, args *GetClusterAdvancedConfigurationApiParams) GetClusterAdvancedConfigurationApiRequest {
	return GetClusterAdvancedConfigurationApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetClusterAdvancedConfigurationApiRequest) Execute() (*ClusterDescriptionProcessArgs, *http.Response, error) {
	return r.ApiService.GetClusterAdvancedConfigurationExecute(r)
}

/*
GetClusterAdvancedConfiguration Return One Advanced Configuration Options for One Cluster

Returns the advanced configuration details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. Advanced configuration details include the read/write concern, index and oplog limits, and other database settings. This feature isn't available for `M0` free clusters, `M2` and `M5` shared-tier clusters, or serverless clusters. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return GetClusterAdvancedConfigurationApiRequest

Deprecated
*/
func (a *ClustersApiService) GetClusterAdvancedConfiguration(ctx context.Context, groupId string, clusterName string) GetClusterAdvancedConfigurationApiRequest {
	return GetClusterAdvancedConfigurationApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return ClusterDescriptionProcessArgs
//
// Deprecated
func (a *ClustersApiService) GetClusterAdvancedConfigurationExecute(r GetClusterAdvancedConfigurationApiRequest) (*ClusterDescriptionProcessArgs, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterDescriptionProcessArgs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.GetClusterAdvancedConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs"
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

type GetClusterStatusApiRequest struct {
	ctx         context.Context
	ApiService  ClustersApi
	groupId     string
	clusterName string
}

type GetClusterStatusApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClustersApiService) GetClusterStatusWithParams(ctx context.Context, args *GetClusterStatusApiParams) GetClusterStatusApiRequest {
	return GetClusterStatusApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetClusterStatusApiRequest) Execute() (*ClusterStatus, *http.Response, error) {
	return r.ApiService.GetClusterStatusExecute(r)
}

/*
GetClusterStatus Return Status of All Cluster Operations

[experimental] Returns the status of all changes that you made to the specified cluster in the specified project. Use this resource to check the progress MongoDB Cloud has made in processing your changes. The response does not include the deployment of new dedicated clusters. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return GetClusterStatusApiRequest
*/
func (a *ClustersApiService) GetClusterStatus(ctx context.Context, groupId string, clusterName string) GetClusterStatusApiRequest {
	return GetClusterStatusApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return ClusterStatus
func (a *ClustersApiService) GetClusterStatusExecute(r GetClusterStatusApiRequest) (*ClusterStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.GetClusterStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/status"
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

type GetSampleDatasetLoadStatusApiRequest struct {
	ctx             context.Context
	ApiService      ClustersApi
	groupId         string
	sampleDatasetId string
}

type GetSampleDatasetLoadStatusApiParams struct {
	GroupId         string
	SampleDatasetId string
}

func (a *ClustersApiService) GetSampleDatasetLoadStatusWithParams(ctx context.Context, args *GetSampleDatasetLoadStatusApiParams) GetSampleDatasetLoadStatusApiRequest {
	return GetSampleDatasetLoadStatusApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         args.GroupId,
		sampleDatasetId: args.SampleDatasetId,
	}
}

func (r GetSampleDatasetLoadStatusApiRequest) Execute() (*SampleDatasetStatus, *http.Response, error) {
	return r.ApiService.GetSampleDatasetLoadStatusExecute(r)
}

/*
GetSampleDatasetLoadStatus Check Status of Cluster Sample Dataset Request

Checks the progress of loading the sample dataset into one cluster. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param sampleDatasetId Unique 24-hexadecimal digit string that identifies the loaded sample dataset.
	@return GetSampleDatasetLoadStatusApiRequest
*/
func (a *ClustersApiService) GetSampleDatasetLoadStatus(ctx context.Context, groupId string, sampleDatasetId string) GetSampleDatasetLoadStatusApiRequest {
	return GetSampleDatasetLoadStatusApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         groupId,
		sampleDatasetId: sampleDatasetId,
	}
}

// Execute executes the request
//
//	@return SampleDatasetStatus
func (a *ClustersApiService) GetSampleDatasetLoadStatusExecute(r GetSampleDatasetLoadStatusApiRequest) (*SampleDatasetStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SampleDatasetStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.GetSampleDatasetLoadStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{sampleDatasetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sampleDatasetId"+"}", url.PathEscape(parameterValueToString(r.sampleDatasetId, "sampleDatasetId")), -1)

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

type ListCloudProviderRegionsApiRequest struct {
	ctx          context.Context
	ApiService   ClustersApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
	providers    *[]string
	tier         *string
}

type ListCloudProviderRegionsApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
	Providers    *[]string
	Tier         *string
}

func (a *ClustersApiService) ListCloudProviderRegionsWithParams(ctx context.Context, args *ListCloudProviderRegionsApiParams) ListCloudProviderRegionsApiRequest {
	return ListCloudProviderRegionsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		providers:    args.Providers,
		tier:         args.Tier,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListCloudProviderRegionsApiRequest) IncludeCount(includeCount bool) ListCloudProviderRegionsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListCloudProviderRegionsApiRequest) ItemsPerPage(itemsPerPage int) ListCloudProviderRegionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListCloudProviderRegionsApiRequest) PageNum(pageNum int) ListCloudProviderRegionsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Cloud providers whose regions to retrieve. When you specify multiple providers, the response can return only tiers and regions that support multi-cloud clusters.
func (r ListCloudProviderRegionsApiRequest) Providers(providers []string) ListCloudProviderRegionsApiRequest {
	r.providers = &providers
	return r
}

// Cluster tier for which to retrieve the regions.
func (r ListCloudProviderRegionsApiRequest) Tier(tier string) ListCloudProviderRegionsApiRequest {
	r.tier = &tier
	return r
}

func (r ListCloudProviderRegionsApiRequest) Execute() (*PaginatedApiAtlasProviderRegions, *http.Response, error) {
	return r.ApiService.ListCloudProviderRegionsExecute(r)
}

/*
ListCloudProviderRegions Return All Cloud Provider Regions

Returns the list of regions available for the specified cloud provider at the specified tier. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListCloudProviderRegionsApiRequest
*/
func (a *ClustersApiService) ListCloudProviderRegions(ctx context.Context, groupId string) ListCloudProviderRegionsApiRequest {
	return ListCloudProviderRegionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return PaginatedApiAtlasProviderRegions
func (a *ClustersApiService) ListCloudProviderRegionsExecute(r ListCloudProviderRegionsApiRequest) (*PaginatedApiAtlasProviderRegions, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAtlasProviderRegions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.ListCloudProviderRegions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/provider/regions"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	if r.providers != nil {
		t := *r.providers
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "providers", t, "multi")

	}
	if r.tier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tier", r.tier, "")
	}
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

type ListClustersApiRequest struct {
	ctx                               context.Context
	ApiService                        ClustersApi
	groupId                           string
	includeCount                      *bool
	itemsPerPage                      *int
	pageNum                           *int
	includeDeletedWithRetainedBackups *bool
}

type ListClustersApiParams struct {
	GroupId                           string
	IncludeCount                      *bool
	ItemsPerPage                      *int
	PageNum                           *int
	IncludeDeletedWithRetainedBackups *bool
}

func (a *ClustersApiService) ListClustersWithParams(ctx context.Context, args *ListClustersApiParams) ListClustersApiRequest {
	return ListClustersApiRequest{
		ApiService:                        a,
		ctx:                               ctx,
		groupId:                           args.GroupId,
		includeCount:                      args.IncludeCount,
		itemsPerPage:                      args.ItemsPerPage,
		pageNum:                           args.PageNum,
		includeDeletedWithRetainedBackups: args.IncludeDeletedWithRetainedBackups,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListClustersApiRequest) IncludeCount(includeCount bool) ListClustersApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListClustersApiRequest) ItemsPerPage(itemsPerPage int) ListClustersApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListClustersApiRequest) PageNum(pageNum int) ListClustersApiRequest {
	r.pageNum = &pageNum
	return r
}

// Flag that indicates whether to return Clusters with retain backups.
func (r ListClustersApiRequest) IncludeDeletedWithRetainedBackups(includeDeletedWithRetainedBackups bool) ListClustersApiRequest {
	r.includeDeletedWithRetainedBackups = &includeDeletedWithRetainedBackups
	return r
}

func (r ListClustersApiRequest) Execute() (*PaginatedAdvancedClusterDescription, *http.Response, error) {
	return r.ApiService.ListClustersExecute(r)
}

/*
ListClusters Return All Clusters in One Project

Returns the details for all clusters in the specific project to which you have access. Clusters contain a group of hosts that maintain the same data set. The response includes multi-cloud clusters. To use this resource, the requesting API Key must have the Project Read Only role. This feature is not available for serverless clusters. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListClustersApiRequest

Deprecated
*/
func (a *ClustersApiService) ListClusters(ctx context.Context, groupId string) ListClustersApiRequest {
	return ListClustersApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return PaginatedAdvancedClusterDescription
//
// Deprecated
func (a *ClustersApiService) ListClustersExecute(r ListClustersApiRequest) (*PaginatedAdvancedClusterDescription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedAdvancedClusterDescription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.ListClusters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	if r.includeDeletedWithRetainedBackups != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeDeletedWithRetainedBackups", r.includeDeletedWithRetainedBackups, "")
	} else {
		var defaultValue bool = false
		r.includeDeletedWithRetainedBackups = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeDeletedWithRetainedBackups", r.includeDeletedWithRetainedBackups, "")
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

type ListClustersForAllProjectsApiRequest struct {
	ctx          context.Context
	ApiService   ClustersApi
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListClustersForAllProjectsApiParams struct {
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ClustersApiService) ListClustersForAllProjectsWithParams(ctx context.Context, args *ListClustersForAllProjectsApiParams) ListClustersForAllProjectsApiRequest {
	return ListClustersForAllProjectsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListClustersForAllProjectsApiRequest) IncludeCount(includeCount bool) ListClustersForAllProjectsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListClustersForAllProjectsApiRequest) ItemsPerPage(itemsPerPage int) ListClustersForAllProjectsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListClustersForAllProjectsApiRequest) PageNum(pageNum int) ListClustersForAllProjectsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListClustersForAllProjectsApiRequest) Execute() (*PaginatedOrgGroup, *http.Response, error) {
	return r.ApiService.ListClustersForAllProjectsExecute(r)
}

/*
ListClustersForAllProjects Return All Authorized Clusters in All Projects

[experimental] Returns the details for all clusters in all projects to which you have access. Clusters contain a group of hosts that maintain the same data set. The response does not include multi-cloud clusters. To use this resource, the requesting API Key can have any cluster-level role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ListClustersForAllProjectsApiRequest
*/
func (a *ClustersApiService) ListClustersForAllProjects(ctx context.Context) ListClustersForAllProjectsApiRequest {
	return ListClustersForAllProjectsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PaginatedOrgGroup
func (a *ClustersApiService) ListClustersForAllProjectsExecute(r ListClustersForAllProjectsApiRequest) (*PaginatedOrgGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedOrgGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.ListClustersForAllProjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/clusters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
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

type LoadSampleDatasetApiRequest struct {
	ctx        context.Context
	ApiService ClustersApi
	groupId    string
	name       string
}

type LoadSampleDatasetApiParams struct {
	GroupId string
	Name    string
}

func (a *ClustersApiService) LoadSampleDatasetWithParams(ctx context.Context, args *LoadSampleDatasetApiParams) LoadSampleDatasetApiRequest {
	return LoadSampleDatasetApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		name:       args.Name,
	}
}

func (r LoadSampleDatasetApiRequest) Execute() (*SampleDatasetStatus, *http.Response, error) {
	return r.ApiService.LoadSampleDatasetExecute(r)
}

/*
LoadSampleDataset Load Sample Dataset Request into Cluster

Requests loading the MongoDB sample dataset into the specified cluster. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the cluster into which you load the sample dataset.
	@return LoadSampleDatasetApiRequest
*/
func (a *ClustersApiService) LoadSampleDataset(ctx context.Context, groupId string, name string) LoadSampleDatasetApiRequest {
	return LoadSampleDatasetApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		name:       name,
	}
}

// Execute executes the request
//
//	@return SampleDatasetStatus
func (a *ClustersApiService) LoadSampleDatasetExecute(r LoadSampleDatasetApiRequest) (*SampleDatasetStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SampleDatasetStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.LoadSampleDataset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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

type TestFailoverApiRequest struct {
	ctx         context.Context
	ApiService  ClustersApi
	groupId     string
	clusterName string
}

type TestFailoverApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClustersApiService) TestFailoverWithParams(ctx context.Context, args *TestFailoverApiParams) TestFailoverApiRequest {
	return TestFailoverApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r TestFailoverApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.TestFailoverExecute(r)
}

/*
TestFailover Test Failover for One Cluster

Starts a failover test for the specified cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. A failover test checks how MongoDB Cloud handles the failure of the cluster's primary node. During the test, MongoDB Cloud shuts down the primary node and elects a new primary. To use this resource, the requesting API Key must have the Project Cluster Manager role. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return TestFailoverApiRequest
*/
func (a *ClustersApiService) TestFailover(ctx context.Context, groupId string, clusterName string) TestFailoverApiRequest {
	return TestFailoverApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
func (a *ClustersApiService) TestFailoverExecute(r TestFailoverApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.TestFailover")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restartPrimaries"
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

type UpdateClusterApiRequest struct {
	ctx                        context.Context
	ApiService                 ClustersApi
	groupId                    string
	clusterName                string
	advancedClusterDescription *AdvancedClusterDescription
}

type UpdateClusterApiParams struct {
	GroupId                    string
	ClusterName                string
	AdvancedClusterDescription *AdvancedClusterDescription
}

func (a *ClustersApiService) UpdateClusterWithParams(ctx context.Context, args *UpdateClusterApiParams) UpdateClusterApiRequest {
	return UpdateClusterApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		clusterName:                args.ClusterName,
		advancedClusterDescription: args.AdvancedClusterDescription,
	}
}

func (r UpdateClusterApiRequest) Execute() (*AdvancedClusterDescription, *http.Response, error) {
	return r.ApiService.UpdateClusterExecute(r)
}

/*
UpdateCluster Modify One Cluster from One Project

Updates the details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can update multi-cloud clusters. To update a cluster's termination protection, the requesting API Key must have the Project Owner role. For all other updates, the requesting API Key must have the Project Cluster Manager role. You can't modify a paused cluster (`paused : true`). You must call this endpoint to set `paused : false`. After this endpoint responds with `paused : false`, you can call it again with the changes you want to make to the cluster. This feature is not available for serverless clusters. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return UpdateClusterApiRequest

Deprecated
*/
func (a *ClustersApiService) UpdateCluster(ctx context.Context, groupId string, clusterName string, advancedClusterDescription *AdvancedClusterDescription) UpdateClusterApiRequest {
	return UpdateClusterApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		clusterName:                clusterName,
		advancedClusterDescription: advancedClusterDescription,
	}
}

// Execute executes the request
//
//	@return AdvancedClusterDescription
//
// Deprecated
func (a *ClustersApiService) UpdateClusterExecute(r UpdateClusterApiRequest) (*AdvancedClusterDescription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AdvancedClusterDescription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.UpdateCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advancedClusterDescription == nil {
		return localVarReturnValue, nil, reportError("advancedClusterDescription is required and must be specified")
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
	localVarPostBody = r.advancedClusterDescription
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

type UpdateClusterAdvancedConfigurationApiRequest struct {
	ctx                           context.Context
	ApiService                    ClustersApi
	groupId                       string
	clusterName                   string
	clusterDescriptionProcessArgs *ClusterDescriptionProcessArgs
}

type UpdateClusterAdvancedConfigurationApiParams struct {
	GroupId                       string
	ClusterName                   string
	ClusterDescriptionProcessArgs *ClusterDescriptionProcessArgs
}

func (a *ClustersApiService) UpdateClusterAdvancedConfigurationWithParams(ctx context.Context, args *UpdateClusterAdvancedConfigurationApiParams) UpdateClusterAdvancedConfigurationApiRequest {
	return UpdateClusterAdvancedConfigurationApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		groupId:                       args.GroupId,
		clusterName:                   args.ClusterName,
		clusterDescriptionProcessArgs: args.ClusterDescriptionProcessArgs,
	}
}

func (r UpdateClusterAdvancedConfigurationApiRequest) Execute() (*ClusterDescriptionProcessArgs, *http.Response, error) {
	return r.ApiService.UpdateClusterAdvancedConfigurationExecute(r)
}

/*
UpdateClusterAdvancedConfiguration Update Advanced Configuration Options for One Cluster

Updates the advanced configuration details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. Advanced configuration details include the read/write concern, index and oplog limits, and other database settings. To use this resource, the requesting API Key must have the Project Cluster Manager role. This feature isn't available for `M0` free clusters, `M2` and `M5` shared-tier clusters, or serverless clusters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return UpdateClusterAdvancedConfigurationApiRequest

Deprecated
*/
func (a *ClustersApiService) UpdateClusterAdvancedConfiguration(ctx context.Context, groupId string, clusterName string, clusterDescriptionProcessArgs *ClusterDescriptionProcessArgs) UpdateClusterAdvancedConfigurationApiRequest {
	return UpdateClusterAdvancedConfigurationApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		groupId:                       groupId,
		clusterName:                   clusterName,
		clusterDescriptionProcessArgs: clusterDescriptionProcessArgs,
	}
}

// Execute executes the request
//
//	@return ClusterDescriptionProcessArgs
//
// Deprecated
func (a *ClustersApiService) UpdateClusterAdvancedConfigurationExecute(r UpdateClusterAdvancedConfigurationApiRequest) (*ClusterDescriptionProcessArgs, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClusterDescriptionProcessArgs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.UpdateClusterAdvancedConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clusterDescriptionProcessArgs == nil {
		return localVarReturnValue, nil, reportError("clusterDescriptionProcessArgs is required and must be specified")
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
	localVarPostBody = r.clusterDescriptionProcessArgs
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

type UpgradeSharedClusterApiRequest struct {
	ctx                                    context.Context
	ApiService                             ClustersApi
	groupId                                string
	legacyAtlasTenantClusterUpgradeRequest *LegacyAtlasTenantClusterUpgradeRequest
}

type UpgradeSharedClusterApiParams struct {
	GroupId                                string
	LegacyAtlasTenantClusterUpgradeRequest *LegacyAtlasTenantClusterUpgradeRequest
}

func (a *ClustersApiService) UpgradeSharedClusterWithParams(ctx context.Context, args *UpgradeSharedClusterApiParams) UpgradeSharedClusterApiRequest {
	return UpgradeSharedClusterApiRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		groupId:                                args.GroupId,
		legacyAtlasTenantClusterUpgradeRequest: args.LegacyAtlasTenantClusterUpgradeRequest,
	}
}

func (r UpgradeSharedClusterApiRequest) Execute() (*LegacyAtlasCluster, *http.Response, error) {
	return r.ApiService.UpgradeSharedClusterExecute(r)
}

/*
UpgradeSharedCluster Upgrade One Shared-tier Cluster

[experimental] Upgrades a shared-tier cluster in the specified project. To use this resource, the requesting API key must have the Project Cluster Manager role. Each project supports up to 25 clusters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpgradeSharedClusterApiRequest
*/
func (a *ClustersApiService) UpgradeSharedCluster(ctx context.Context, groupId string, legacyAtlasTenantClusterUpgradeRequest *LegacyAtlasTenantClusterUpgradeRequest) UpgradeSharedClusterApiRequest {
	return UpgradeSharedClusterApiRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		groupId:                                groupId,
		legacyAtlasTenantClusterUpgradeRequest: legacyAtlasTenantClusterUpgradeRequest,
	}
}

// Execute executes the request
//
//	@return LegacyAtlasCluster
func (a *ClustersApiService) UpgradeSharedClusterExecute(r UpgradeSharedClusterApiRequest) (*LegacyAtlasCluster, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LegacyAtlasCluster
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.UpgradeSharedCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/tenantUpgrade"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.legacyAtlasTenantClusterUpgradeRequest == nil {
		return localVarReturnValue, nil, reportError("legacyAtlasTenantClusterUpgradeRequest is required and must be specified")
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
	localVarPostBody = r.legacyAtlasTenantClusterUpgradeRequest
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

type UpgradeSharedClusterToServerlessApiRequest struct {
	ctx                           context.Context
	ApiService                    ClustersApi
	groupId                       string
	serverlessInstanceDescription *ServerlessInstanceDescription
}

type UpgradeSharedClusterToServerlessApiParams struct {
	GroupId                       string
	ServerlessInstanceDescription *ServerlessInstanceDescription
}

func (a *ClustersApiService) UpgradeSharedClusterToServerlessWithParams(ctx context.Context, args *UpgradeSharedClusterToServerlessApiParams) UpgradeSharedClusterToServerlessApiRequest {
	return UpgradeSharedClusterToServerlessApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		groupId:                       args.GroupId,
		serverlessInstanceDescription: args.ServerlessInstanceDescription,
	}
}

func (r UpgradeSharedClusterToServerlessApiRequest) Execute() (*ServerlessInstanceDescription, *http.Response, error) {
	return r.ApiService.UpgradeSharedClusterToServerlessExecute(r)
}

/*
UpgradeSharedClusterToServerless Upgrades One Shared-Tier Cluster to the Serverless Instance

[experimental] Upgrades a shared-tier cluster to a serverless instance in the specified project. To use this resource, the requesting API key must have the Project Cluster Manager role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpgradeSharedClusterToServerlessApiRequest
*/
func (a *ClustersApiService) UpgradeSharedClusterToServerless(ctx context.Context, groupId string, serverlessInstanceDescription *ServerlessInstanceDescription) UpgradeSharedClusterToServerlessApiRequest {
	return UpgradeSharedClusterToServerlessApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		groupId:                       groupId,
		serverlessInstanceDescription: serverlessInstanceDescription,
	}
}

// Execute executes the request
//
//	@return ServerlessInstanceDescription
func (a *ClustersApiService) UpgradeSharedClusterToServerlessExecute(r UpgradeSharedClusterToServerlessApiRequest) (*ServerlessInstanceDescription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ServerlessInstanceDescription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.UpgradeSharedClusterToServerless")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/tenantUpgradeToServerless"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serverlessInstanceDescription == nil {
		return localVarReturnValue, nil, reportError("serverlessInstanceDescription is required and must be specified")
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
	localVarPostBody = r.serverlessInstanceDescription
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
