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

type ProjectsApi interface {

	/*
		AddUserToProject Add One MongoDB Cloud User to One Project

		[experimental] Adds one MongoDB Cloud user to the specified project. If the MongoDB Cloud user is not a member of the project's organization, then the user must accept their invitation to the organization to access information within the specified project. If the MongoDB Cloud User is already a member of the project's organization, then they will be added to the project immediately and an invitation will not be returned by this resource. To use this resource, the requesting API Key must have the Group User Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return AddUserToProjectApiRequest
	*/
	AddUserToProject(ctx context.Context, groupId string, groupInvitationRequest *GroupInvitationRequest) AddUserToProjectApiRequest
	/*
		AddUserToProject Add One MongoDB Cloud User to One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AddUserToProjectApiParams - Parameters for the request
		@return AddUserToProjectApiRequest
	*/
	AddUserToProjectWithParams(ctx context.Context, args *AddUserToProjectApiParams) AddUserToProjectApiRequest

	// Method available only for mocking purposes
	AddUserToProjectExecute(r AddUserToProjectApiRequest) (*OrganizationInvitation, *http.Response, error)

	/*
		CreateProject Create One Project

		Creates one project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting API Key must have the Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return CreateProjectApiRequest
	*/
	CreateProject(ctx context.Context, group *Group) CreateProjectApiRequest
	/*
		CreateProject Create One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateProjectApiParams - Parameters for the request
		@return CreateProjectApiRequest
	*/
	CreateProjectWithParams(ctx context.Context, args *CreateProjectApiParams) CreateProjectApiRequest

	// Method available only for mocking purposes
	CreateProjectExecute(r CreateProjectApiRequest) (*Group, *http.Response, error)

	/*
		CreateProjectInvitation Invite One MongoDB Cloud User to Join One Project

		Invites one MongoDB Cloud user to join the specified project. The MongoDB Cloud user must accept the invitation to access information within the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return CreateProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	CreateProjectInvitation(ctx context.Context, groupId string, groupInvitationRequest *GroupInvitationRequest) CreateProjectInvitationApiRequest
	/*
		CreateProjectInvitation Invite One MongoDB Cloud User to Join One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateProjectInvitationApiParams - Parameters for the request
		@return CreateProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	CreateProjectInvitationWithParams(ctx context.Context, args *CreateProjectInvitationApiParams) CreateProjectInvitationApiRequest

	// Method available only for mocking purposes
	CreateProjectInvitationExecute(r CreateProjectInvitationApiRequest) (*GroupInvitation, *http.Response, error)

	/*
		DeleteProject Remove One Project

		Removes the specified project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. You can delete a project only if there are no Online Archives for the clusters in the project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeleteProjectApiRequest
	*/
	DeleteProject(ctx context.Context, groupId string) DeleteProjectApiRequest
	/*
		DeleteProject Remove One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteProjectApiParams - Parameters for the request
		@return DeleteProjectApiRequest
	*/
	DeleteProjectWithParams(ctx context.Context, args *DeleteProjectApiParams) DeleteProjectApiRequest

	// Method available only for mocking purposes
	DeleteProjectExecute(r DeleteProjectApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		DeleteProjectInvitation Cancel One Project Invitation

		Cancels one pending invitation sent to the specified MongoDB Cloud user to join a project. You can't cancel an invitation that the user accepted. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
		@return DeleteProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	DeleteProjectInvitation(ctx context.Context, groupId string, invitationId string) DeleteProjectInvitationApiRequest
	/*
		DeleteProjectInvitation Cancel One Project Invitation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteProjectInvitationApiParams - Parameters for the request
		@return DeleteProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	DeleteProjectInvitationWithParams(ctx context.Context, args *DeleteProjectInvitationApiParams) DeleteProjectInvitationApiRequest

	// Method available only for mocking purposes
	DeleteProjectInvitationExecute(r DeleteProjectInvitationApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		DeleteProjectLimit Remove One Project Limit

		[experimental] Removes the specified project limit. Depending on the limit, Atlas either resets the limit to its default value or removes the limit entirely. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param limitName Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project | 25 | 90 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Serivce Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return DeleteProjectLimitApiRequest
	*/
	DeleteProjectLimit(ctx context.Context, limitName string, groupId string) DeleteProjectLimitApiRequest
	/*
		DeleteProjectLimit Remove One Project Limit


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteProjectLimitApiParams - Parameters for the request
		@return DeleteProjectLimitApiRequest
	*/
	DeleteProjectLimitWithParams(ctx context.Context, args *DeleteProjectLimitApiParams) DeleteProjectLimitApiRequest

	// Method available only for mocking purposes
	DeleteProjectLimitExecute(r DeleteProjectLimitApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetProject Return One Project

		Returns details about the specified project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetProjectApiRequest
	*/
	GetProject(ctx context.Context, groupId string) GetProjectApiRequest
	/*
		GetProject Return One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectApiParams - Parameters for the request
		@return GetProjectApiRequest
	*/
	GetProjectWithParams(ctx context.Context, args *GetProjectApiParams) GetProjectApiRequest

	// Method available only for mocking purposes
	GetProjectExecute(r GetProjectApiRequest) (*Group, *http.Response, error)

	/*
		GetProjectByName Return One Project using Its Name

		Returns details about the specified project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupName Human-readable label that identifies this project.
		@return GetProjectByNameApiRequest
	*/
	GetProjectByName(ctx context.Context, groupName string) GetProjectByNameApiRequest
	/*
		GetProjectByName Return One Project using Its Name


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectByNameApiParams - Parameters for the request
		@return GetProjectByNameApiRequest
	*/
	GetProjectByNameWithParams(ctx context.Context, args *GetProjectByNameApiParams) GetProjectByNameApiRequest

	// Method available only for mocking purposes
	GetProjectByNameExecute(r GetProjectByNameApiRequest) (*Group, *http.Response, error)

	/*
		GetProjectInvitation Return One Project Invitation

		Returns the details of one pending invitation to the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
		@return GetProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	GetProjectInvitation(ctx context.Context, groupId string, invitationId string) GetProjectInvitationApiRequest
	/*
		GetProjectInvitation Return One Project Invitation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectInvitationApiParams - Parameters for the request
		@return GetProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	GetProjectInvitationWithParams(ctx context.Context, args *GetProjectInvitationApiParams) GetProjectInvitationApiRequest

	// Method available only for mocking purposes
	GetProjectInvitationExecute(r GetProjectInvitationApiRequest) (*GroupInvitation, *http.Response, error)

	/*
		GetProjectLimit Return One Limit for One Project

		[experimental] Returns the specified limit for the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param limitName Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project | 25 | 90 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Serivce Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetProjectLimitApiRequest
	*/
	GetProjectLimit(ctx context.Context, limitName string, groupId string) GetProjectLimitApiRequest
	/*
		GetProjectLimit Return One Limit for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectLimitApiParams - Parameters for the request
		@return GetProjectLimitApiRequest
	*/
	GetProjectLimitWithParams(ctx context.Context, args *GetProjectLimitApiParams) GetProjectLimitApiRequest

	// Method available only for mocking purposes
	GetProjectLimitExecute(r GetProjectLimitApiRequest) (*DataFederationLimit, *http.Response, error)

	/*
		GetProjectSettings Return One Project Settings

		Returns details about the specified project's settings. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return GetProjectSettingsApiRequest
	*/
	GetProjectSettings(ctx context.Context, groupId string) GetProjectSettingsApiRequest
	/*
		GetProjectSettings Return One Project Settings


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetProjectSettingsApiParams - Parameters for the request
		@return GetProjectSettingsApiRequest
	*/
	GetProjectSettingsWithParams(ctx context.Context, args *GetProjectSettingsApiParams) GetProjectSettingsApiRequest

	// Method available only for mocking purposes
	GetProjectSettingsExecute(r GetProjectSettingsApiRequest) (*GroupSettings, *http.Response, error)

	/*
		ListProjectInvitations Return All Project Invitations

		Returns all pending invitations to the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListProjectInvitationsApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	ListProjectInvitations(ctx context.Context, groupId string) ListProjectInvitationsApiRequest
	/*
		ListProjectInvitations Return All Project Invitations


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectInvitationsApiParams - Parameters for the request
		@return ListProjectInvitationsApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	ListProjectInvitationsWithParams(ctx context.Context, args *ListProjectInvitationsApiParams) ListProjectInvitationsApiRequest

	// Method available only for mocking purposes
	ListProjectInvitationsExecute(r ListProjectInvitationsApiRequest) ([]GroupInvitation, *http.Response, error)

	/*
		ListProjectLimits Return All Limits for One Project

		[experimental] Returns all the limits for the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListProjectLimitsApiRequest
	*/
	ListProjectLimits(ctx context.Context, groupId string) ListProjectLimitsApiRequest
	/*
		ListProjectLimits Return All Limits for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectLimitsApiParams - Parameters for the request
		@return ListProjectLimitsApiRequest
	*/
	ListProjectLimitsWithParams(ctx context.Context, args *ListProjectLimitsApiParams) ListProjectLimitsApiRequest

	// Method available only for mocking purposes
	ListProjectLimitsExecute(r ListProjectLimitsApiRequest) ([]DataFederationLimit, *http.Response, error)

	/*
		ListProjectUsers Return All Users in One Project

		Returns details about all users in the specified project. Users belong to an organization. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListProjectUsersApiRequest
	*/
	ListProjectUsers(ctx context.Context, groupId string) ListProjectUsersApiRequest
	/*
		ListProjectUsers Return All Users in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectUsersApiParams - Parameters for the request
		@return ListProjectUsersApiRequest
	*/
	ListProjectUsersWithParams(ctx context.Context, args *ListProjectUsersApiParams) ListProjectUsersApiRequest

	// Method available only for mocking purposes
	ListProjectUsersExecute(r ListProjectUsersApiRequest) (*PaginatedAppUser, *http.Response, error)

	/*
		ListProjects Return All Projects

		Returns details about all projects. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting API Key must have the Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ListProjectsApiRequest
	*/
	ListProjects(ctx context.Context) ListProjectsApiRequest
	/*
		ListProjects Return All Projects


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListProjectsApiParams - Parameters for the request
		@return ListProjectsApiRequest
	*/
	ListProjectsWithParams(ctx context.Context, args *ListProjectsApiParams) ListProjectsApiRequest

	// Method available only for mocking purposes
	ListProjectsExecute(r ListProjectsApiRequest) (*PaginatedAtlasGroup, *http.Response, error)

	/*
		RemoveProjectUser Remove One User from One Project

		Removes the specified user from the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param userId Unique 24-hexadecimal string that identifies MongoDB Cloud user you want to remove from the specified project. To return a application user's ID using their application username, use the Get All application users in One Project endpoint.
		@return RemoveProjectUserApiRequest
	*/
	RemoveProjectUser(ctx context.Context, groupId string, userId string) RemoveProjectUserApiRequest
	/*
		RemoveProjectUser Remove One User from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RemoveProjectUserApiParams - Parameters for the request
		@return RemoveProjectUserApiRequest
	*/
	RemoveProjectUserWithParams(ctx context.Context, args *RemoveProjectUserApiParams) RemoveProjectUserApiRequest

	// Method available only for mocking purposes
	RemoveProjectUserExecute(r RemoveProjectUserApiRequest) (*http.Response, error)

	/*
		ReturnAllIPAddresses Return All IP Addresses for One Project

		[experimental] Returns all IP addresses for this project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ReturnAllIPAddressesApiRequest
	*/
	ReturnAllIPAddresses(ctx context.Context, groupId string) ReturnAllIPAddressesApiRequest
	/*
		ReturnAllIPAddresses Return All IP Addresses for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ReturnAllIPAddressesApiParams - Parameters for the request
		@return ReturnAllIPAddressesApiRequest
	*/
	ReturnAllIPAddressesWithParams(ctx context.Context, args *ReturnAllIPAddressesApiParams) ReturnAllIPAddressesApiRequest

	// Method available only for mocking purposes
	ReturnAllIPAddressesExecute(r ReturnAllIPAddressesApiRequest) (*GroupIPAddresses, *http.Response, error)

	/*
			SetProjectLimit Set One Project Limit

			[experimental] Sets the specified project limit. To use this resource, the requesting API Key must have the Project Owner role.

		**NOTE**: Increasing the following configuration limits might lead to slower response times in the MongoDB Cloud UI or increased user management overhead leading to authentication or authorization re-architecture. If possible, we recommend that you create additional projects to gain access to more of these resources for a more sustainable growth pattern.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param limitName Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project | 25 | 90 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Serivce Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return SetProjectLimitApiRequest
	*/
	SetProjectLimit(ctx context.Context, limitName string, groupId string, dataFederationLimit *DataFederationLimit) SetProjectLimitApiRequest
	/*
		SetProjectLimit Set One Project Limit


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param SetProjectLimitApiParams - Parameters for the request
		@return SetProjectLimitApiRequest
	*/
	SetProjectLimitWithParams(ctx context.Context, args *SetProjectLimitApiParams) SetProjectLimitApiRequest

	// Method available only for mocking purposes
	SetProjectLimitExecute(r SetProjectLimitApiRequest) (*DataFederationLimit, *http.Response, error)

	/*
		UpdateProject Update One Project

		[experimental] Updates the human-readable label that identifies the specified project, or the tags associated with the project. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return UpdateProjectApiRequest
	*/
	UpdateProject(ctx context.Context, groupId string, groupUpdate *GroupUpdate) UpdateProjectApiRequest
	/*
		UpdateProject Update One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateProjectApiParams - Parameters for the request
		@return UpdateProjectApiRequest
	*/
	UpdateProjectWithParams(ctx context.Context, args *UpdateProjectApiParams) UpdateProjectApiRequest

	// Method available only for mocking purposes
	UpdateProjectExecute(r UpdateProjectApiRequest) (*Group, *http.Response, error)

	/*
		UpdateProjectInvitation Update One Project Invitation

		Updates the details of one pending invitation to the specified project. To specify which invitation to update, provide the username of the invited user. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return UpdateProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	UpdateProjectInvitation(ctx context.Context, groupId string, groupInvitationRequest *GroupInvitationRequest) UpdateProjectInvitationApiRequest
	/*
		UpdateProjectInvitation Update One Project Invitation


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateProjectInvitationApiParams - Parameters for the request
		@return UpdateProjectInvitationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	UpdateProjectInvitationWithParams(ctx context.Context, args *UpdateProjectInvitationApiParams) UpdateProjectInvitationApiRequest

	// Method available only for mocking purposes
	UpdateProjectInvitationExecute(r UpdateProjectInvitationApiRequest) (*GroupInvitation, *http.Response, error)

	/*
		UpdateProjectInvitationById Update One Project Invitation by Invitation ID

		Updates the details of one pending invitation to the specified project. To specify which invitation to update, provide the unique identification string for that invitation. Use the Return All Project Invitations endpoint to retrieve IDs for all pending project invitations. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
		@return UpdateProjectInvitationByIdApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	UpdateProjectInvitationById(ctx context.Context, groupId string, invitationId string, groupInvitationUpdateRequest *GroupInvitationUpdateRequest) UpdateProjectInvitationByIdApiRequest
	/*
		UpdateProjectInvitationById Update One Project Invitation by Invitation ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateProjectInvitationByIdApiParams - Parameters for the request
		@return UpdateProjectInvitationByIdApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ProjectsApi
	*/
	UpdateProjectInvitationByIdWithParams(ctx context.Context, args *UpdateProjectInvitationByIdApiParams) UpdateProjectInvitationByIdApiRequest

	// Method available only for mocking purposes
	UpdateProjectInvitationByIdExecute(r UpdateProjectInvitationByIdApiRequest) (*GroupInvitation, *http.Response, error)

	/*
		UpdateProjectRoles Update Project Roles for One MongoDB Cloud User

		[experimental] Updates the roles of the specified user in the specified project. To specify the user to update, provide the unique 24-hexadecimal digit string that identifies the user in the specified project. To use this resource, the requesting API Key must have the Group User Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param userId Unique 24-hexadecimal digit string that identifies the user to modify.
		@return UpdateProjectRolesApiRequest
	*/
	UpdateProjectRoles(ctx context.Context, groupId string, userId string, updateGroupRolesForUser *UpdateGroupRolesForUser) UpdateProjectRolesApiRequest
	/*
		UpdateProjectRoles Update Project Roles for One MongoDB Cloud User


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateProjectRolesApiParams - Parameters for the request
		@return UpdateProjectRolesApiRequest
	*/
	UpdateProjectRolesWithParams(ctx context.Context, args *UpdateProjectRolesApiParams) UpdateProjectRolesApiRequest

	// Method available only for mocking purposes
	UpdateProjectRolesExecute(r UpdateProjectRolesApiRequest) (*UpdateGroupRolesForUser, *http.Response, error)

	/*
		UpdateProjectSettings Update One Project Settings

		Updates the settings of the specified project. You can update any of the options available. MongoDB cloud only updates the options provided in the request. To use this resource, the requesting API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return UpdateProjectSettingsApiRequest
	*/
	UpdateProjectSettings(ctx context.Context, groupId string, groupSettings *GroupSettings) UpdateProjectSettingsApiRequest
	/*
		UpdateProjectSettings Update One Project Settings


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateProjectSettingsApiParams - Parameters for the request
		@return UpdateProjectSettingsApiRequest
	*/
	UpdateProjectSettingsWithParams(ctx context.Context, args *UpdateProjectSettingsApiParams) UpdateProjectSettingsApiRequest

	// Method available only for mocking purposes
	UpdateProjectSettingsExecute(r UpdateProjectSettingsApiRequest) (*GroupSettings, *http.Response, error)
}

// ProjectsApiService ProjectsApi service
type ProjectsApiService service

type AddUserToProjectApiRequest struct {
	ctx                    context.Context
	ApiService             ProjectsApi
	groupId                string
	groupInvitationRequest *GroupInvitationRequest
}

type AddUserToProjectApiParams struct {
	GroupId                string
	GroupInvitationRequest *GroupInvitationRequest
}

func (a *ProjectsApiService) AddUserToProjectWithParams(ctx context.Context, args *AddUserToProjectApiParams) AddUserToProjectApiRequest {
	return AddUserToProjectApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                args.GroupId,
		groupInvitationRequest: args.GroupInvitationRequest,
	}
}

func (r AddUserToProjectApiRequest) Execute() (*OrganizationInvitation, *http.Response, error) {
	return r.ApiService.AddUserToProjectExecute(r)
}

/*
AddUserToProject Add One MongoDB Cloud User to One Project

[experimental] Adds one MongoDB Cloud user to the specified project. If the MongoDB Cloud user is not a member of the project's organization, then the user must accept their invitation to the organization to access information within the specified project. If the MongoDB Cloud User is already a member of the project's organization, then they will be added to the project immediately and an invitation will not be returned by this resource. To use this resource, the requesting API Key must have the Group User Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return AddUserToProjectApiRequest
*/
func (a *ProjectsApiService) AddUserToProject(ctx context.Context, groupId string, groupInvitationRequest *GroupInvitationRequest) AddUserToProjectApiRequest {
	return AddUserToProjectApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                groupId,
		groupInvitationRequest: groupInvitationRequest,
	}
}

// Execute executes the request
//
//	@return OrganizationInvitation
func (a *ProjectsApiService) AddUserToProjectExecute(r AddUserToProjectApiRequest) (*OrganizationInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.AddUserToProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/access"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupInvitationRequest == nil {
		return localVarReturnValue, nil, reportError("groupInvitationRequest is required and must be specified")
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
	localVarPostBody = r.groupInvitationRequest
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

type CreateProjectApiRequest struct {
	ctx            context.Context
	ApiService     ProjectsApi
	group          *Group
	projectOwnerId *string
}

type CreateProjectApiParams struct {
	Group          *Group
	ProjectOwnerId *string
}

func (a *ProjectsApiService) CreateProjectWithParams(ctx context.Context, args *CreateProjectApiParams) CreateProjectApiRequest {
	return CreateProjectApiRequest{
		ApiService:     a,
		ctx:            ctx,
		group:          args.Group,
		projectOwnerId: args.ProjectOwnerId,
	}
}

// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user to whom to grant the Project Owner role on the specified project. If you set this parameter, it overrides the default value of the oldest Organization Owner.
func (r CreateProjectApiRequest) ProjectOwnerId(projectOwnerId string) CreateProjectApiRequest {
	r.projectOwnerId = &projectOwnerId
	return r
}

func (r CreateProjectApiRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.CreateProjectExecute(r)
}

/*
CreateProject Create One Project

Creates one project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting API Key must have the Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CreateProjectApiRequest
*/
func (a *ProjectsApiService) CreateProject(ctx context.Context, group *Group) CreateProjectApiRequest {
	return CreateProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		group:      group,
	}
}

// Execute executes the request
//
//	@return Group
func (a *ProjectsApiService) CreateProjectExecute(r CreateProjectApiRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.CreateProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.group == nil {
		return localVarReturnValue, nil, reportError("group is required and must be specified")
	}

	if r.projectOwnerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "projectOwnerId", r.projectOwnerId, "")
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
	localVarPostBody = r.group
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

type CreateProjectInvitationApiRequest struct {
	ctx                    context.Context
	ApiService             ProjectsApi
	groupId                string
	groupInvitationRequest *GroupInvitationRequest
}

type CreateProjectInvitationApiParams struct {
	GroupId                string
	GroupInvitationRequest *GroupInvitationRequest
}

func (a *ProjectsApiService) CreateProjectInvitationWithParams(ctx context.Context, args *CreateProjectInvitationApiParams) CreateProjectInvitationApiRequest {
	return CreateProjectInvitationApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                args.GroupId,
		groupInvitationRequest: args.GroupInvitationRequest,
	}
}

func (r CreateProjectInvitationApiRequest) Execute() (*GroupInvitation, *http.Response, error) {
	return r.ApiService.CreateProjectInvitationExecute(r)
}

/*
CreateProjectInvitation Invite One MongoDB Cloud User to Join One Project

Invites one MongoDB Cloud user to join the specified project. The MongoDB Cloud user must accept the invitation to access information within the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateProjectInvitationApiRequest

Deprecated
*/
func (a *ProjectsApiService) CreateProjectInvitation(ctx context.Context, groupId string, groupInvitationRequest *GroupInvitationRequest) CreateProjectInvitationApiRequest {
	return CreateProjectInvitationApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                groupId,
		groupInvitationRequest: groupInvitationRequest,
	}
}

// Execute executes the request
//
//	@return GroupInvitation
//
// Deprecated
func (a *ProjectsApiService) CreateProjectInvitationExecute(r CreateProjectInvitationApiRequest) (*GroupInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.CreateProjectInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/invites"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupInvitationRequest == nil {
		return localVarReturnValue, nil, reportError("groupInvitationRequest is required and must be specified")
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
	localVarPostBody = r.groupInvitationRequest
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

type DeleteProjectApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
}

type DeleteProjectApiParams struct {
	GroupId string
}

func (a *ProjectsApiService) DeleteProjectWithParams(ctx context.Context, args *DeleteProjectApiParams) DeleteProjectApiRequest {
	return DeleteProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r DeleteProjectApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteProjectExecute(r)
}

/*
DeleteProject Remove One Project

Removes the specified project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. You can delete a project only if there are no Online Archives for the clusters in the project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeleteProjectApiRequest
*/
func (a *ProjectsApiService) DeleteProject(ctx context.Context, groupId string) DeleteProjectApiRequest {
	return DeleteProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *ProjectsApiService) DeleteProjectExecute(r DeleteProjectApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.DeleteProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}"
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

type DeleteProjectInvitationApiRequest struct {
	ctx          context.Context
	ApiService   ProjectsApi
	groupId      string
	invitationId string
}

type DeleteProjectInvitationApiParams struct {
	GroupId      string
	InvitationId string
}

func (a *ProjectsApiService) DeleteProjectInvitationWithParams(ctx context.Context, args *DeleteProjectInvitationApiParams) DeleteProjectInvitationApiRequest {
	return DeleteProjectInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		invitationId: args.InvitationId,
	}
}

func (r DeleteProjectInvitationApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteProjectInvitationExecute(r)
}

/*
DeleteProjectInvitation Cancel One Project Invitation

Cancels one pending invitation sent to the specified MongoDB Cloud user to join a project. You can't cancel an invitation that the user accepted. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
	@return DeleteProjectInvitationApiRequest

Deprecated
*/
func (a *ProjectsApiService) DeleteProjectInvitation(ctx context.Context, groupId string, invitationId string) DeleteProjectInvitationApiRequest {
	return DeleteProjectInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		invitationId: invitationId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
//
// Deprecated
func (a *ProjectsApiService) DeleteProjectInvitationExecute(r DeleteProjectInvitationApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.DeleteProjectInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/invites/{invitationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"invitationId"+"}", url.PathEscape(parameterValueToString(r.invitationId, "invitationId")), -1)

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

type DeleteProjectLimitApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	limitName  string
	groupId    string
}

type DeleteProjectLimitApiParams struct {
	LimitName string
	GroupId   string
}

func (a *ProjectsApiService) DeleteProjectLimitWithParams(ctx context.Context, args *DeleteProjectLimitApiParams) DeleteProjectLimitApiRequest {
	return DeleteProjectLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		limitName:  args.LimitName,
		groupId:    args.GroupId,
	}
}

func (r DeleteProjectLimitApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteProjectLimitExecute(r)
}

/*
DeleteProjectLimit Remove One Project Limit

[experimental] Removes the specified project limit. Depending on the limit, Atlas either resets the limit to its default value or removes the limit entirely. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param limitName Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project | 25 | 90 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Serivce Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return DeleteProjectLimitApiRequest
*/
func (a *ProjectsApiService) DeleteProjectLimit(ctx context.Context, limitName string, groupId string) DeleteProjectLimitApiRequest {
	return DeleteProjectLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		limitName:  limitName,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *ProjectsApiService) DeleteProjectLimitExecute(r DeleteProjectLimitApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.DeleteProjectLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/limits/{limitName}"
	localVarPath = strings.Replace(localVarPath, "{"+"limitName"+"}", url.PathEscape(parameterValueToString(r.limitName, "limitName")), -1)
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

type GetProjectApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
}

type GetProjectApiParams struct {
	GroupId string
}

func (a *ProjectsApiService) GetProjectWithParams(ctx context.Context, args *GetProjectApiParams) GetProjectApiRequest {
	return GetProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetProjectApiRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.GetProjectExecute(r)
}

/*
GetProject Return One Project

Returns details about the specified project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetProjectApiRequest
*/
func (a *ProjectsApiService) GetProject(ctx context.Context, groupId string) GetProjectApiRequest {
	return GetProjectApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return Group
func (a *ProjectsApiService) GetProjectExecute(r GetProjectApiRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.GetProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}"
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

type GetProjectByNameApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupName  string
}

type GetProjectByNameApiParams struct {
	GroupName string
}

func (a *ProjectsApiService) GetProjectByNameWithParams(ctx context.Context, args *GetProjectByNameApiParams) GetProjectByNameApiRequest {
	return GetProjectByNameApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupName:  args.GroupName,
	}
}

func (r GetProjectByNameApiRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.GetProjectByNameExecute(r)
}

/*
GetProjectByName Return One Project using Its Name

Returns details about the specified project. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupName Human-readable label that identifies this project.
	@return GetProjectByNameApiRequest
*/
func (a *ProjectsApiService) GetProjectByName(ctx context.Context, groupName string) GetProjectByNameApiRequest {
	return GetProjectByNameApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupName:  groupName,
	}
}

// Execute executes the request
//
//	@return Group
func (a *ProjectsApiService) GetProjectByNameExecute(r GetProjectByNameApiRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.GetProjectByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/byName/{groupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupName"+"}", url.PathEscape(parameterValueToString(r.groupName, "groupName")), -1)

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

type GetProjectInvitationApiRequest struct {
	ctx          context.Context
	ApiService   ProjectsApi
	groupId      string
	invitationId string
}

type GetProjectInvitationApiParams struct {
	GroupId      string
	InvitationId string
}

func (a *ProjectsApiService) GetProjectInvitationWithParams(ctx context.Context, args *GetProjectInvitationApiParams) GetProjectInvitationApiRequest {
	return GetProjectInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		invitationId: args.InvitationId,
	}
}

func (r GetProjectInvitationApiRequest) Execute() (*GroupInvitation, *http.Response, error) {
	return r.ApiService.GetProjectInvitationExecute(r)
}

/*
GetProjectInvitation Return One Project Invitation

Returns the details of one pending invitation to the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
	@return GetProjectInvitationApiRequest

Deprecated
*/
func (a *ProjectsApiService) GetProjectInvitation(ctx context.Context, groupId string, invitationId string) GetProjectInvitationApiRequest {
	return GetProjectInvitationApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		invitationId: invitationId,
	}
}

// Execute executes the request
//
//	@return GroupInvitation
//
// Deprecated
func (a *ProjectsApiService) GetProjectInvitationExecute(r GetProjectInvitationApiRequest) (*GroupInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.GetProjectInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/invites/{invitationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"invitationId"+"}", url.PathEscape(parameterValueToString(r.invitationId, "invitationId")), -1)

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

type GetProjectLimitApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	limitName  string
	groupId    string
}

type GetProjectLimitApiParams struct {
	LimitName string
	GroupId   string
}

func (a *ProjectsApiService) GetProjectLimitWithParams(ctx context.Context, args *GetProjectLimitApiParams) GetProjectLimitApiRequest {
	return GetProjectLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		limitName:  args.LimitName,
		groupId:    args.GroupId,
	}
}

func (r GetProjectLimitApiRequest) Execute() (*DataFederationLimit, *http.Response, error) {
	return r.ApiService.GetProjectLimitExecute(r)
}

/*
GetProjectLimit Return One Limit for One Project

[experimental] Returns the specified limit for the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param limitName Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project | 25 | 90 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Serivce Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetProjectLimitApiRequest
*/
func (a *ProjectsApiService) GetProjectLimit(ctx context.Context, limitName string, groupId string) GetProjectLimitApiRequest {
	return GetProjectLimitApiRequest{
		ApiService: a,
		ctx:        ctx,
		limitName:  limitName,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return DataFederationLimit
func (a *ProjectsApiService) GetProjectLimitExecute(r GetProjectLimitApiRequest) (*DataFederationLimit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataFederationLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.GetProjectLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/limits/{limitName}"
	localVarPath = strings.Replace(localVarPath, "{"+"limitName"+"}", url.PathEscape(parameterValueToString(r.limitName, "limitName")), -1)
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

type GetProjectSettingsApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
}

type GetProjectSettingsApiParams struct {
	GroupId string
}

func (a *ProjectsApiService) GetProjectSettingsWithParams(ctx context.Context, args *GetProjectSettingsApiParams) GetProjectSettingsApiRequest {
	return GetProjectSettingsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r GetProjectSettingsApiRequest) Execute() (*GroupSettings, *http.Response, error) {
	return r.ApiService.GetProjectSettingsExecute(r)
}

/*
GetProjectSettings Return One Project Settings

Returns details about the specified project's settings. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return GetProjectSettingsApiRequest
*/
func (a *ProjectsApiService) GetProjectSettings(ctx context.Context, groupId string) GetProjectSettingsApiRequest {
	return GetProjectSettingsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return GroupSettings
func (a *ProjectsApiService) GetProjectSettingsExecute(r GetProjectSettingsApiRequest) (*GroupSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.GetProjectSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/settings"
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

type ListProjectInvitationsApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
	username   *string
}

type ListProjectInvitationsApiParams struct {
	GroupId  string
	Username *string
}

func (a *ProjectsApiService) ListProjectInvitationsWithParams(ctx context.Context, args *ListProjectInvitationsApiParams) ListProjectInvitationsApiRequest {
	return ListProjectInvitationsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		username:   args.Username,
	}
}

// Email address of the user account invited to this project.
func (r ListProjectInvitationsApiRequest) Username(username string) ListProjectInvitationsApiRequest {
	r.username = &username
	return r
}

func (r ListProjectInvitationsApiRequest) Execute() ([]GroupInvitation, *http.Response, error) {
	return r.ApiService.ListProjectInvitationsExecute(r)
}

/*
ListProjectInvitations Return All Project Invitations

Returns all pending invitations to the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListProjectInvitationsApiRequest

Deprecated
*/
func (a *ProjectsApiService) ListProjectInvitations(ctx context.Context, groupId string) ListProjectInvitationsApiRequest {
	return ListProjectInvitationsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return []GroupInvitation
//
// Deprecated
func (a *ProjectsApiService) ListProjectInvitationsExecute(r ListProjectInvitationsApiRequest) ([]GroupInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []GroupInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.ListProjectInvitations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/invites"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.username != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "username", r.username, "")
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

type ListProjectLimitsApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
}

type ListProjectLimitsApiParams struct {
	GroupId string
}

func (a *ProjectsApiService) ListProjectLimitsWithParams(ctx context.Context, args *ListProjectLimitsApiParams) ListProjectLimitsApiRequest {
	return ListProjectLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r ListProjectLimitsApiRequest) Execute() ([]DataFederationLimit, *http.Response, error) {
	return r.ApiService.ListProjectLimitsExecute(r)
}

/*
ListProjectLimits Return All Limits for One Project

[experimental] Returns all the limits for the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListProjectLimitsApiRequest
*/
func (a *ProjectsApiService) ListProjectLimits(ctx context.Context, groupId string) ListProjectLimitsApiRequest {
	return ListProjectLimitsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return []DataFederationLimit
func (a *ProjectsApiService) ListProjectLimitsExecute(r ListProjectLimitsApiRequest) ([]DataFederationLimit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []DataFederationLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.ListProjectLimits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/limits"
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

type ListProjectUsersApiRequest struct {
	ctx             context.Context
	ApiService      ProjectsApi
	groupId         string
	includeCount    *bool
	itemsPerPage    *int
	pageNum         *int
	flattenTeams    *bool
	includeOrgUsers *bool
}

type ListProjectUsersApiParams struct {
	GroupId         string
	IncludeCount    *bool
	ItemsPerPage    *int
	PageNum         *int
	FlattenTeams    *bool
	IncludeOrgUsers *bool
}

func (a *ProjectsApiService) ListProjectUsersWithParams(ctx context.Context, args *ListProjectUsersApiParams) ListProjectUsersApiRequest {
	return ListProjectUsersApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         args.GroupId,
		includeCount:    args.IncludeCount,
		itemsPerPage:    args.ItemsPerPage,
		pageNum:         args.PageNum,
		flattenTeams:    args.FlattenTeams,
		includeOrgUsers: args.IncludeOrgUsers,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListProjectUsersApiRequest) IncludeCount(includeCount bool) ListProjectUsersApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListProjectUsersApiRequest) ItemsPerPage(itemsPerPage int) ListProjectUsersApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListProjectUsersApiRequest) PageNum(pageNum int) ListProjectUsersApiRequest {
	r.pageNum = &pageNum
	return r
}

// Flag that indicates whether the returned list should include users who belong to a team with a role in this project. You might not have assigned the individual users a role in this project. If &#x60;\&quot;flattenTeams\&quot; : false&#x60;, this resource returns only users with a role in the project.  If &#x60;\&quot;flattenTeams\&quot; : true&#x60;, this resource returns both users with roles in the project and users who belong to teams with roles in the project.
func (r ListProjectUsersApiRequest) FlattenTeams(flattenTeams bool) ListProjectUsersApiRequest {
	r.flattenTeams = &flattenTeams
	return r
}

// Flag that indicates whether the returned list should include users with implicit access to the project, the Organization Owner or Organization Read Only role. You might not have assigned the individual users a role in this project. If &#x60;\&quot;includeOrgUsers\&quot;: false&#x60;, this resource returns only users with a role in the project. If &#x60;\&quot;includeOrgUsers\&quot;: true&#x60;, this resource returns both users with roles in the project and users who have implicit access to the project through their organization role.
func (r ListProjectUsersApiRequest) IncludeOrgUsers(includeOrgUsers bool) ListProjectUsersApiRequest {
	r.includeOrgUsers = &includeOrgUsers
	return r
}

func (r ListProjectUsersApiRequest) Execute() (*PaginatedAppUser, *http.Response, error) {
	return r.ApiService.ListProjectUsersExecute(r)
}

/*
ListProjectUsers Return All Users in One Project

Returns details about all users in the specified project. Users belong to an organization. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListProjectUsersApiRequest
*/
func (a *ProjectsApiService) ListProjectUsers(ctx context.Context, groupId string) ListProjectUsersApiRequest {
	return ListProjectUsersApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return PaginatedAppUser
func (a *ProjectsApiService) ListProjectUsersExecute(r ListProjectUsersApiRequest) (*PaginatedAppUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedAppUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.ListProjectUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/users"
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
	if r.flattenTeams != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "flattenTeams", r.flattenTeams, "")
	} else {
		var defaultValue bool = false
		r.flattenTeams = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "flattenTeams", r.flattenTeams, "")
	}
	if r.includeOrgUsers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeOrgUsers", r.includeOrgUsers, "")
	} else {
		var defaultValue bool = false
		r.includeOrgUsers = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeOrgUsers", r.includeOrgUsers, "")
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

type ListProjectsApiRequest struct {
	ctx          context.Context
	ApiService   ProjectsApi
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListProjectsApiParams struct {
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ProjectsApiService) ListProjectsWithParams(ctx context.Context, args *ListProjectsApiParams) ListProjectsApiRequest {
	return ListProjectsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListProjectsApiRequest) IncludeCount(includeCount bool) ListProjectsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListProjectsApiRequest) ItemsPerPage(itemsPerPage int) ListProjectsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListProjectsApiRequest) PageNum(pageNum int) ListProjectsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListProjectsApiRequest) Execute() (*PaginatedAtlasGroup, *http.Response, error) {
	return r.ApiService.ListProjectsExecute(r)
}

/*
ListProjects Return All Projects

Returns details about all projects. Projects group clusters into logical collections that support an application environment, workload, or both. Each project can have its own users, teams, security, tags, and alert settings. To use this resource, the requesting API Key must have the Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ListProjectsApiRequest
*/
func (a *ProjectsApiService) ListProjects(ctx context.Context) ListProjectsApiRequest {
	return ListProjectsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PaginatedAtlasGroup
func (a *ProjectsApiService) ListProjectsExecute(r ListProjectsApiRequest) (*PaginatedAtlasGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedAtlasGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.ListProjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups"

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

type RemoveProjectUserApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
	userId     string
}

type RemoveProjectUserApiParams struct {
	GroupId string
	UserId  string
}

func (a *ProjectsApiService) RemoveProjectUserWithParams(ctx context.Context, args *RemoveProjectUserApiParams) RemoveProjectUserApiRequest {
	return RemoveProjectUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		userId:     args.UserId,
	}
}

func (r RemoveProjectUserApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveProjectUserExecute(r)
}

/*
RemoveProjectUser Remove One User from One Project

Removes the specified user from the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param userId Unique 24-hexadecimal string that identifies MongoDB Cloud user you want to remove from the specified project. To return a application user's ID using their application username, use the Get All application users in One Project endpoint.
	@return RemoveProjectUserApiRequest
*/
func (a *ProjectsApiService) RemoveProjectUser(ctx context.Context, groupId string, userId string) RemoveProjectUserApiRequest {
	return RemoveProjectUserApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		userId:     userId,
	}
}

// Execute executes the request
func (a *ProjectsApiService) RemoveProjectUserExecute(r RemoveProjectUserApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.RemoveProjectUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type ReturnAllIPAddressesApiRequest struct {
	ctx        context.Context
	ApiService ProjectsApi
	groupId    string
}

type ReturnAllIPAddressesApiParams struct {
	GroupId string
}

func (a *ProjectsApiService) ReturnAllIPAddressesWithParams(ctx context.Context, args *ReturnAllIPAddressesApiParams) ReturnAllIPAddressesApiRequest {
	return ReturnAllIPAddressesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r ReturnAllIPAddressesApiRequest) Execute() (*GroupIPAddresses, *http.Response, error) {
	return r.ApiService.ReturnAllIPAddressesExecute(r)
}

/*
ReturnAllIPAddresses Return All IP Addresses for One Project

[experimental] Returns all IP addresses for this project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ReturnAllIPAddressesApiRequest
*/
func (a *ProjectsApiService) ReturnAllIPAddresses(ctx context.Context, groupId string) ReturnAllIPAddressesApiRequest {
	return ReturnAllIPAddressesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return GroupIPAddresses
func (a *ProjectsApiService) ReturnAllIPAddressesExecute(r ReturnAllIPAddressesApiRequest) (*GroupIPAddresses, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupIPAddresses
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.ReturnAllIPAddresses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/ipAddresses"
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

type SetProjectLimitApiRequest struct {
	ctx                 context.Context
	ApiService          ProjectsApi
	limitName           string
	groupId             string
	dataFederationLimit *DataFederationLimit
}

type SetProjectLimitApiParams struct {
	LimitName           string
	GroupId             string
	DataFederationLimit *DataFederationLimit
}

func (a *ProjectsApiService) SetProjectLimitWithParams(ctx context.Context, args *SetProjectLimitApiParams) SetProjectLimitApiRequest {
	return SetProjectLimitApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		limitName:           args.LimitName,
		groupId:             args.GroupId,
		dataFederationLimit: args.DataFederationLimit,
	}
}

func (r SetProjectLimitApiRequest) Execute() (*DataFederationLimit, *http.Response, error) {
	return r.ApiService.SetProjectLimitExecute(r)
}

/*
SetProjectLimit Set One Project Limit

[experimental] Sets the specified project limit. To use this resource, the requesting API Key must have the Project Owner role.

**NOTE**: Increasing the following configuration limits might lead to slower response times in the MongoDB Cloud UI or increased user management overhead leading to authentication or authorization re-architecture. If possible, we recommend that you create additional projects to gain access to more of these resources for a more sustainable growth pattern.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param limitName Human-readable label that identifies this project limit.  | Limit Name | Description | Default | API Override Limit | | --- | --- | --- | --- | | atlas.project.deployment.clusters | Limit on the number of clusters in this project | 25 | 90 | | atlas.project.deployment.nodesPerPrivateLinkRegion | Limit on the number of nodes per Private Link region in this project | 50 | 90 | | atlas.project.security.databaseAccess.customRoles | Limit on the number of custom roles in this project | 100 | 1400 | | atlas.project.security.databaseAccess.users | Limit on the number of database users in this project | 100 | 900 | | atlas.project.security.networkAccess.crossRegionEntries | Limit on the number of cross-region network access entries in this project | 40 | 220 | | atlas.project.security.networkAccess.entries | Limit on the number of network access entries in this project | 200 | 20 | | dataFederation.bytesProcessed.query | Limit on the number of bytes processed during a single Data Federation query | N/A | N/A | | dataFederation.bytesProcessed.daily | Limit on the number of bytes processed across all Data Federation tenants for the current day | N/A | N/A | | dataFederation.bytesProcessed.weekly | Limit on the number of bytes processed across all Data Federation tenants for the current week | N/A | N/A | | dataFederation.bytesProcessed.monthly | Limit on the number of bytes processed across all Data Federation tenants for the current month | N/A | N/A | | atlas.project.deployment.privateServiceConnectionsPerRegionGroup | Number of Private Serivce Connections per Region Group | 50 | 100| | atlas.project.deployment.privateServiceConnectionsSubnetMask | Subnet mask for GCP PSC Networks. Has lower limit of 20. | 27 | 27|
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return SetProjectLimitApiRequest
*/
func (a *ProjectsApiService) SetProjectLimit(ctx context.Context, limitName string, groupId string, dataFederationLimit *DataFederationLimit) SetProjectLimitApiRequest {
	return SetProjectLimitApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		limitName:           limitName,
		groupId:             groupId,
		dataFederationLimit: dataFederationLimit,
	}
}

// Execute executes the request
//
//	@return DataFederationLimit
func (a *ProjectsApiService) SetProjectLimitExecute(r SetProjectLimitApiRequest) (*DataFederationLimit, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataFederationLimit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.SetProjectLimit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/limits/{limitName}"
	localVarPath = strings.Replace(localVarPath, "{"+"limitName"+"}", url.PathEscape(parameterValueToString(r.limitName, "limitName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataFederationLimit == nil {
		return localVarReturnValue, nil, reportError("dataFederationLimit is required and must be specified")
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
	localVarPostBody = r.dataFederationLimit
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

type UpdateProjectApiRequest struct {
	ctx         context.Context
	ApiService  ProjectsApi
	groupId     string
	groupUpdate *GroupUpdate
}

type UpdateProjectApiParams struct {
	GroupId     string
	GroupUpdate *GroupUpdate
}

func (a *ProjectsApiService) UpdateProjectWithParams(ctx context.Context, args *UpdateProjectApiParams) UpdateProjectApiRequest {
	return UpdateProjectApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		groupUpdate: args.GroupUpdate,
	}
}

func (r UpdateProjectApiRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.UpdateProjectExecute(r)
}

/*
UpdateProject Update One Project

[experimental] Updates the human-readable label that identifies the specified project, or the tags associated with the project. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateProjectApiRequest
*/
func (a *ProjectsApiService) UpdateProject(ctx context.Context, groupId string, groupUpdate *GroupUpdate) UpdateProjectApiRequest {
	return UpdateProjectApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		groupUpdate: groupUpdate,
	}
}

// Execute executes the request
//
//	@return Group
func (a *ProjectsApiService) UpdateProjectExecute(r UpdateProjectApiRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.UpdateProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupUpdate == nil {
		return localVarReturnValue, nil, reportError("groupUpdate is required and must be specified")
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
	localVarPostBody = r.groupUpdate
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

type UpdateProjectInvitationApiRequest struct {
	ctx                    context.Context
	ApiService             ProjectsApi
	groupId                string
	groupInvitationRequest *GroupInvitationRequest
}

type UpdateProjectInvitationApiParams struct {
	GroupId                string
	GroupInvitationRequest *GroupInvitationRequest
}

func (a *ProjectsApiService) UpdateProjectInvitationWithParams(ctx context.Context, args *UpdateProjectInvitationApiParams) UpdateProjectInvitationApiRequest {
	return UpdateProjectInvitationApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                args.GroupId,
		groupInvitationRequest: args.GroupInvitationRequest,
	}
}

func (r UpdateProjectInvitationApiRequest) Execute() (*GroupInvitation, *http.Response, error) {
	return r.ApiService.UpdateProjectInvitationExecute(r)
}

/*
UpdateProjectInvitation Update One Project Invitation

Updates the details of one pending invitation to the specified project. To specify which invitation to update, provide the username of the invited user. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateProjectInvitationApiRequest

Deprecated
*/
func (a *ProjectsApiService) UpdateProjectInvitation(ctx context.Context, groupId string, groupInvitationRequest *GroupInvitationRequest) UpdateProjectInvitationApiRequest {
	return UpdateProjectInvitationApiRequest{
		ApiService:             a,
		ctx:                    ctx,
		groupId:                groupId,
		groupInvitationRequest: groupInvitationRequest,
	}
}

// Execute executes the request
//
//	@return GroupInvitation
//
// Deprecated
func (a *ProjectsApiService) UpdateProjectInvitationExecute(r UpdateProjectInvitationApiRequest) (*GroupInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.UpdateProjectInvitation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/invites"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupInvitationRequest == nil {
		return localVarReturnValue, nil, reportError("groupInvitationRequest is required and must be specified")
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
	localVarPostBody = r.groupInvitationRequest
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

type UpdateProjectInvitationByIdApiRequest struct {
	ctx                          context.Context
	ApiService                   ProjectsApi
	groupId                      string
	invitationId                 string
	groupInvitationUpdateRequest *GroupInvitationUpdateRequest
}

type UpdateProjectInvitationByIdApiParams struct {
	GroupId                      string
	InvitationId                 string
	GroupInvitationUpdateRequest *GroupInvitationUpdateRequest
}

func (a *ProjectsApiService) UpdateProjectInvitationByIdWithParams(ctx context.Context, args *UpdateProjectInvitationByIdApiParams) UpdateProjectInvitationByIdApiRequest {
	return UpdateProjectInvitationByIdApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		groupId:                      args.GroupId,
		invitationId:                 args.InvitationId,
		groupInvitationUpdateRequest: args.GroupInvitationUpdateRequest,
	}
}

func (r UpdateProjectInvitationByIdApiRequest) Execute() (*GroupInvitation, *http.Response, error) {
	return r.ApiService.UpdateProjectInvitationByIdExecute(r)
}

/*
UpdateProjectInvitationById Update One Project Invitation by Invitation ID

Updates the details of one pending invitation to the specified project. To specify which invitation to update, provide the unique identification string for that invitation. Use the Return All Project Invitations endpoint to retrieve IDs for all pending project invitations. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param invitationId Unique 24-hexadecimal digit string that identifies the invitation.
	@return UpdateProjectInvitationByIdApiRequest

Deprecated
*/
func (a *ProjectsApiService) UpdateProjectInvitationById(ctx context.Context, groupId string, invitationId string, groupInvitationUpdateRequest *GroupInvitationUpdateRequest) UpdateProjectInvitationByIdApiRequest {
	return UpdateProjectInvitationByIdApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		groupId:                      groupId,
		invitationId:                 invitationId,
		groupInvitationUpdateRequest: groupInvitationUpdateRequest,
	}
}

// Execute executes the request
//
//	@return GroupInvitation
//
// Deprecated
func (a *ProjectsApiService) UpdateProjectInvitationByIdExecute(r UpdateProjectInvitationByIdApiRequest) (*GroupInvitation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.UpdateProjectInvitationById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/invites/{invitationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"invitationId"+"}", url.PathEscape(parameterValueToString(r.invitationId, "invitationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupInvitationUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("groupInvitationUpdateRequest is required and must be specified")
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
	localVarPostBody = r.groupInvitationUpdateRequest
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

type UpdateProjectRolesApiRequest struct {
	ctx                     context.Context
	ApiService              ProjectsApi
	groupId                 string
	userId                  string
	updateGroupRolesForUser *UpdateGroupRolesForUser
}

type UpdateProjectRolesApiParams struct {
	GroupId                 string
	UserId                  string
	UpdateGroupRolesForUser *UpdateGroupRolesForUser
}

func (a *ProjectsApiService) UpdateProjectRolesWithParams(ctx context.Context, args *UpdateProjectRolesApiParams) UpdateProjectRolesApiRequest {
	return UpdateProjectRolesApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 args.GroupId,
		userId:                  args.UserId,
		updateGroupRolesForUser: args.UpdateGroupRolesForUser,
	}
}

func (r UpdateProjectRolesApiRequest) Execute() (*UpdateGroupRolesForUser, *http.Response, error) {
	return r.ApiService.UpdateProjectRolesExecute(r)
}

/*
UpdateProjectRoles Update Project Roles for One MongoDB Cloud User

[experimental] Updates the roles of the specified user in the specified project. To specify the user to update, provide the unique 24-hexadecimal digit string that identifies the user in the specified project. To use this resource, the requesting API Key must have the Group User Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param userId Unique 24-hexadecimal digit string that identifies the user to modify.
	@return UpdateProjectRolesApiRequest
*/
func (a *ProjectsApiService) UpdateProjectRoles(ctx context.Context, groupId string, userId string, updateGroupRolesForUser *UpdateGroupRolesForUser) UpdateProjectRolesApiRequest {
	return UpdateProjectRolesApiRequest{
		ApiService:              a,
		ctx:                     ctx,
		groupId:                 groupId,
		userId:                  userId,
		updateGroupRolesForUser: updateGroupRolesForUser,
	}
}

// Execute executes the request
//
//	@return UpdateGroupRolesForUser
func (a *ProjectsApiService) UpdateProjectRolesExecute(r UpdateProjectRolesApiRequest) (*UpdateGroupRolesForUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UpdateGroupRolesForUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.UpdateProjectRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/users/{userId}/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateGroupRolesForUser == nil {
		return localVarReturnValue, nil, reportError("updateGroupRolesForUser is required and must be specified")
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
	localVarPostBody = r.updateGroupRolesForUser
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

type UpdateProjectSettingsApiRequest struct {
	ctx           context.Context
	ApiService    ProjectsApi
	groupId       string
	groupSettings *GroupSettings
}

type UpdateProjectSettingsApiParams struct {
	GroupId       string
	GroupSettings *GroupSettings
}

func (a *ProjectsApiService) UpdateProjectSettingsWithParams(ctx context.Context, args *UpdateProjectSettingsApiParams) UpdateProjectSettingsApiRequest {
	return UpdateProjectSettingsApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		groupSettings: args.GroupSettings,
	}
}

func (r UpdateProjectSettingsApiRequest) Execute() (*GroupSettings, *http.Response, error) {
	return r.ApiService.UpdateProjectSettingsExecute(r)
}

/*
UpdateProjectSettings Update One Project Settings

Updates the settings of the specified project. You can update any of the options available. MongoDB cloud only updates the options provided in the request. To use this resource, the requesting API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpdateProjectSettingsApiRequest
*/
func (a *ProjectsApiService) UpdateProjectSettings(ctx context.Context, groupId string, groupSettings *GroupSettings) UpdateProjectSettingsApiRequest {
	return UpdateProjectSettingsApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		groupSettings: groupSettings,
	}
}

// Execute executes the request
//
//	@return GroupSettings
func (a *ProjectsApiService) UpdateProjectSettingsExecute(r UpdateProjectSettingsApiRequest) (*GroupSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectsApiService.UpdateProjectSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupSettings == nil {
		return localVarReturnValue, nil, reportError("groupSettings is required and must be specified")
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
	localVarPostBody = r.groupSettings
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
