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

type LegacyBackupApi interface {

	/*
		CreateLegacyBackupRestoreJob Create One Legacy Backup Restore Job

		[experimental] Restores one legacy backup for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role. Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule). This endpoint doesn't support creating checkpoint restore jobs for sharded clusters, or creating restore jobs for queryable backup snapshots. If you create an automated restore job by specifying `delivery.methodName` of `AUTOMATED_RESTORE` in your request body, MongoDB Cloud removes all existing data on the target cluster prior to the restore.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
		@return CreateLegacyBackupRestoreJobApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	CreateLegacyBackupRestoreJob(ctx context.Context, groupId string, clusterName string, backupRestoreJob *BackupRestoreJob) CreateLegacyBackupRestoreJobApiRequest
	/*
		CreateLegacyBackupRestoreJob Create One Legacy Backup Restore Job


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateLegacyBackupRestoreJobApiParams - Parameters for the request
		@return CreateLegacyBackupRestoreJobApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	CreateLegacyBackupRestoreJobWithParams(ctx context.Context, args *CreateLegacyBackupRestoreJobApiParams) CreateLegacyBackupRestoreJobApiRequest

	// Method available only for mocking purposes
	CreateLegacyBackupRestoreJobExecute(r CreateLegacyBackupRestoreJobApiRequest) (*PaginatedRestoreJob, *http.Response, error)

	/*
		DeleteLegacySnapshot Remove One Legacy Backup Snapshot

		[experimental] Removes one legacy backup snapshot for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role. Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
		@return DeleteLegacySnapshotApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	DeleteLegacySnapshot(ctx context.Context, groupId string, clusterName string, snapshotId string) DeleteLegacySnapshotApiRequest
	/*
		DeleteLegacySnapshot Remove One Legacy Backup Snapshot


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteLegacySnapshotApiParams - Parameters for the request
		@return DeleteLegacySnapshotApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	DeleteLegacySnapshotWithParams(ctx context.Context, args *DeleteLegacySnapshotApiParams) DeleteLegacySnapshotApiRequest

	// Method available only for mocking purposes
	DeleteLegacySnapshotExecute(r DeleteLegacySnapshotApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetLegacyBackupCheckpoint Return One Legacy Backup Checkpoint

		[experimental] Returns one legacy backup checkpoint for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param checkpointId Unique 24-hexadecimal digit string that identifies the checkpoint.
		@param clusterName Human-readable label that identifies the cluster that contains the checkpoints that you want to return.
		@return GetLegacyBackupCheckpointApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	GetLegacyBackupCheckpoint(ctx context.Context, groupId string, checkpointId string, clusterName string) GetLegacyBackupCheckpointApiRequest
	/*
		GetLegacyBackupCheckpoint Return One Legacy Backup Checkpoint


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetLegacyBackupCheckpointApiParams - Parameters for the request
		@return GetLegacyBackupCheckpointApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	GetLegacyBackupCheckpointWithParams(ctx context.Context, args *GetLegacyBackupCheckpointApiParams) GetLegacyBackupCheckpointApiRequest

	// Method available only for mocking purposes
	GetLegacyBackupCheckpointExecute(r GetLegacyBackupCheckpointApiRequest) (*ApiAtlasCheckpoint, *http.Response, error)

	/*
			GetLegacyBackupRestoreJob Return One Legacy Backup Restore Job

			[experimental] Returns one legacy backup restore job for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		 Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
			@param jobId Unique 24-hexadecimal digit string that identifies the restore job.
			@return GetLegacyBackupRestoreJobApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	GetLegacyBackupRestoreJob(ctx context.Context, groupId string, clusterName string, jobId string) GetLegacyBackupRestoreJobApiRequest
	/*
		GetLegacyBackupRestoreJob Return One Legacy Backup Restore Job


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetLegacyBackupRestoreJobApiParams - Parameters for the request
		@return GetLegacyBackupRestoreJobApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	GetLegacyBackupRestoreJobWithParams(ctx context.Context, args *GetLegacyBackupRestoreJobApiParams) GetLegacyBackupRestoreJobApiRequest

	// Method available only for mocking purposes
	GetLegacyBackupRestoreJobExecute(r GetLegacyBackupRestoreJobApiRequest) (*BackupRestoreJob, *http.Response, error)

	/*
		GetLegacySnapshot Return One Legacy Backup Snapshot

		[experimental] Returns one legacy backup snapshot for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role. Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
		@return GetLegacySnapshotApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	GetLegacySnapshot(ctx context.Context, groupId string, clusterName string, snapshotId string) GetLegacySnapshotApiRequest
	/*
		GetLegacySnapshot Return One Legacy Backup Snapshot


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetLegacySnapshotApiParams - Parameters for the request
		@return GetLegacySnapshotApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	GetLegacySnapshotWithParams(ctx context.Context, args *GetLegacySnapshotApiParams) GetLegacySnapshotApiRequest

	// Method available only for mocking purposes
	GetLegacySnapshotExecute(r GetLegacySnapshotApiRequest) (*BackupSnapshot, *http.Response, error)

	/*
			GetLegacySnapshotSchedule Return One Snapshot Schedule

			[experimental] Returns the snapshot schedule for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		 Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
			@return GetLegacySnapshotScheduleApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	GetLegacySnapshotSchedule(ctx context.Context, groupId string, clusterName string) GetLegacySnapshotScheduleApiRequest
	/*
		GetLegacySnapshotSchedule Return One Snapshot Schedule


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetLegacySnapshotScheduleApiParams - Parameters for the request
		@return GetLegacySnapshotScheduleApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	GetLegacySnapshotScheduleWithParams(ctx context.Context, args *GetLegacySnapshotScheduleApiParams) GetLegacySnapshotScheduleApiRequest

	// Method available only for mocking purposes
	GetLegacySnapshotScheduleExecute(r GetLegacySnapshotScheduleApiRequest) (*ApiAtlasSnapshotSchedule, *http.Response, error)

	/*
		ListLegacyBackupCheckpoints Return All Legacy Backup Checkpoints

		[experimental] Returns all legacy backup checkpoints for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster that contains the checkpoints that you want to return.
		@return ListLegacyBackupCheckpointsApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	ListLegacyBackupCheckpoints(ctx context.Context, groupId string, clusterName string) ListLegacyBackupCheckpointsApiRequest
	/*
		ListLegacyBackupCheckpoints Return All Legacy Backup Checkpoints


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListLegacyBackupCheckpointsApiParams - Parameters for the request
		@return ListLegacyBackupCheckpointsApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	ListLegacyBackupCheckpointsWithParams(ctx context.Context, args *ListLegacyBackupCheckpointsApiParams) ListLegacyBackupCheckpointsApiRequest

	// Method available only for mocking purposes
	ListLegacyBackupCheckpointsExecute(r ListLegacyBackupCheckpointsApiRequest) (*PaginatedApiAtlasCheckpoint, *http.Response, error)

	/*
			ListLegacyBackupRestoreJobs Return All Legacy Backup Restore Jobs

			[experimental] Returns all legacy backup restore jobs for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		 Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule). If you use the `BATCH-ID` query parameter, you can retrieve all restore jobs in the specified batch. When creating a restore job for a sharded cluster, MongoDB Cloud creates a separate job for each shard, plus another for the config server. Each of those jobs are part of a batch. However, a batch can't include a restore job for a replica set.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
			@return ListLegacyBackupRestoreJobsApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	ListLegacyBackupRestoreJobs(ctx context.Context, groupId string, clusterName string) ListLegacyBackupRestoreJobsApiRequest
	/*
		ListLegacyBackupRestoreJobs Return All Legacy Backup Restore Jobs


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListLegacyBackupRestoreJobsApiParams - Parameters for the request
		@return ListLegacyBackupRestoreJobsApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	ListLegacyBackupRestoreJobsWithParams(ctx context.Context, args *ListLegacyBackupRestoreJobsApiParams) ListLegacyBackupRestoreJobsApiRequest

	// Method available only for mocking purposes
	ListLegacyBackupRestoreJobsExecute(r ListLegacyBackupRestoreJobsApiRequest) (*PaginatedRestoreJob, *http.Response, error)

	/*
		ListLegacySnapshots Return All Legacy Backup Snapshots

		[experimental] Returns all legacy backup snapshots for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role. Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return ListLegacySnapshotsApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	ListLegacySnapshots(ctx context.Context, groupId string, clusterName string) ListLegacySnapshotsApiRequest
	/*
		ListLegacySnapshots Return All Legacy Backup Snapshots


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListLegacySnapshotsApiParams - Parameters for the request
		@return ListLegacySnapshotsApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	ListLegacySnapshotsWithParams(ctx context.Context, args *ListLegacySnapshotsApiParams) ListLegacySnapshotsApiRequest

	// Method available only for mocking purposes
	ListLegacySnapshotsExecute(r ListLegacySnapshotsApiRequest) (*PaginatedSnapshot, *http.Response, error)

	/*
		UpdateLegacySnapshotRetention Change One Legacy Backup Snapshot Expiration

		[experimental] Changes the expiration date for one legacy backup snapshot for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role. Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
		@return UpdateLegacySnapshotRetentionApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	UpdateLegacySnapshotRetention(ctx context.Context, groupId string, clusterName string, snapshotId string, backupSnapshot *BackupSnapshot) UpdateLegacySnapshotRetentionApiRequest
	/*
		UpdateLegacySnapshotRetention Change One Legacy Backup Snapshot Expiration


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateLegacySnapshotRetentionApiParams - Parameters for the request
		@return UpdateLegacySnapshotRetentionApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	UpdateLegacySnapshotRetentionWithParams(ctx context.Context, args *UpdateLegacySnapshotRetentionApiParams) UpdateLegacySnapshotRetentionApiRequest

	// Method available only for mocking purposes
	UpdateLegacySnapshotRetentionExecute(r UpdateLegacySnapshotRetentionApiRequest) (*BackupSnapshot, *http.Response, error)

	/*
			UpdateLegacySnapshotSchedule Update Snapshot Schedule for One Cluster

			[experimental] Updates the snapshot schedule for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

		 Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
			@return UpdateLegacySnapshotScheduleApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	UpdateLegacySnapshotSchedule(ctx context.Context, groupId string, clusterName string, apiAtlasSnapshotSchedule *ApiAtlasSnapshotSchedule) UpdateLegacySnapshotScheduleApiRequest
	/*
		UpdateLegacySnapshotSchedule Update Snapshot Schedule for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateLegacySnapshotScheduleApiParams - Parameters for the request
		@return UpdateLegacySnapshotScheduleApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for LegacyBackupApi
	*/
	UpdateLegacySnapshotScheduleWithParams(ctx context.Context, args *UpdateLegacySnapshotScheduleApiParams) UpdateLegacySnapshotScheduleApiRequest

	// Method available only for mocking purposes
	UpdateLegacySnapshotScheduleExecute(r UpdateLegacySnapshotScheduleApiRequest) (*ApiAtlasSnapshotSchedule, *http.Response, error)
}

// LegacyBackupApiService LegacyBackupApi service
type LegacyBackupApiService service

type CreateLegacyBackupRestoreJobApiRequest struct {
	ctx              context.Context
	ApiService       LegacyBackupApi
	groupId          string
	clusterName      string
	backupRestoreJob *BackupRestoreJob
}

type CreateLegacyBackupRestoreJobApiParams struct {
	GroupId          string
	ClusterName      string
	BackupRestoreJob *BackupRestoreJob
}

func (a *LegacyBackupApiService) CreateLegacyBackupRestoreJobWithParams(ctx context.Context, args *CreateLegacyBackupRestoreJobApiParams) CreateLegacyBackupRestoreJobApiRequest {
	return CreateLegacyBackupRestoreJobApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          args.GroupId,
		clusterName:      args.ClusterName,
		backupRestoreJob: args.BackupRestoreJob,
	}
}

func (r CreateLegacyBackupRestoreJobApiRequest) Execute() (*PaginatedRestoreJob, *http.Response, error) {
	return r.ApiService.CreateLegacyBackupRestoreJobExecute(r)
}

/*
CreateLegacyBackupRestoreJob Create One Legacy Backup Restore Job

[experimental] Restores one legacy backup for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role. Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule). This endpoint doesn't support creating checkpoint restore jobs for sharded clusters, or creating restore jobs for queryable backup snapshots. If you create an automated restore job by specifying `delivery.methodName` of `AUTOMATED_RESTORE` in your request body, MongoDB Cloud removes all existing data on the target cluster prior to the restore.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
	@return CreateLegacyBackupRestoreJobApiRequest

Deprecated
*/
func (a *LegacyBackupApiService) CreateLegacyBackupRestoreJob(ctx context.Context, groupId string, clusterName string, backupRestoreJob *BackupRestoreJob) CreateLegacyBackupRestoreJobApiRequest {
	return CreateLegacyBackupRestoreJobApiRequest{
		ApiService:       a,
		ctx:              ctx,
		groupId:          groupId,
		clusterName:      clusterName,
		backupRestoreJob: backupRestoreJob,
	}
}

// Execute executes the request
//
//	@return PaginatedRestoreJob
//
// Deprecated
func (a *LegacyBackupApiService) CreateLegacyBackupRestoreJobExecute(r CreateLegacyBackupRestoreJobApiRequest) (*PaginatedRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupApiService.CreateLegacyBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.backupRestoreJob == nil {
		return localVarReturnValue, nil, reportError("backupRestoreJob is required and must be specified")
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
	localVarPostBody = r.backupRestoreJob
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

type DeleteLegacySnapshotApiRequest struct {
	ctx         context.Context
	ApiService  LegacyBackupApi
	groupId     string
	clusterName string
	snapshotId  string
}

type DeleteLegacySnapshotApiParams struct {
	GroupId     string
	ClusterName string
	SnapshotId  string
}

func (a *LegacyBackupApiService) DeleteLegacySnapshotWithParams(ctx context.Context, args *DeleteLegacySnapshotApiParams) DeleteLegacySnapshotApiRequest {
	return DeleteLegacySnapshotApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		snapshotId:  args.SnapshotId,
	}
}

func (r DeleteLegacySnapshotApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteLegacySnapshotExecute(r)
}

/*
DeleteLegacySnapshot Remove One Legacy Backup Snapshot

[experimental] Removes one legacy backup snapshot for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role. Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
	@return DeleteLegacySnapshotApiRequest

Deprecated
*/
func (a *LegacyBackupApiService) DeleteLegacySnapshot(ctx context.Context, groupId string, clusterName string, snapshotId string) DeleteLegacySnapshotApiRequest {
	return DeleteLegacySnapshotApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		snapshotId:  snapshotId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
//
// Deprecated
func (a *LegacyBackupApiService) DeleteLegacySnapshotExecute(r DeleteLegacySnapshotApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupApiService.DeleteLegacySnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

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

type GetLegacyBackupCheckpointApiRequest struct {
	ctx          context.Context
	ApiService   LegacyBackupApi
	groupId      string
	checkpointId string
	clusterName  string
}

type GetLegacyBackupCheckpointApiParams struct {
	GroupId      string
	CheckpointId string
	ClusterName  string
}

func (a *LegacyBackupApiService) GetLegacyBackupCheckpointWithParams(ctx context.Context, args *GetLegacyBackupCheckpointApiParams) GetLegacyBackupCheckpointApiRequest {
	return GetLegacyBackupCheckpointApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		checkpointId: args.CheckpointId,
		clusterName:  args.ClusterName,
	}
}

func (r GetLegacyBackupCheckpointApiRequest) Execute() (*ApiAtlasCheckpoint, *http.Response, error) {
	return r.ApiService.GetLegacyBackupCheckpointExecute(r)
}

/*
GetLegacyBackupCheckpoint Return One Legacy Backup Checkpoint

[experimental] Returns one legacy backup checkpoint for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param checkpointId Unique 24-hexadecimal digit string that identifies the checkpoint.
	@param clusterName Human-readable label that identifies the cluster that contains the checkpoints that you want to return.
	@return GetLegacyBackupCheckpointApiRequest

Deprecated
*/
func (a *LegacyBackupApiService) GetLegacyBackupCheckpoint(ctx context.Context, groupId string, checkpointId string, clusterName string) GetLegacyBackupCheckpointApiRequest {
	return GetLegacyBackupCheckpointApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		checkpointId: checkpointId,
		clusterName:  clusterName,
	}
}

// Execute executes the request
//
//	@return ApiAtlasCheckpoint
//
// Deprecated
func (a *LegacyBackupApiService) GetLegacyBackupCheckpointExecute(r GetLegacyBackupCheckpointApiRequest) (*ApiAtlasCheckpoint, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiAtlasCheckpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupApiService.GetLegacyBackupCheckpoint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints/{checkpointId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"checkpointId"+"}", url.PathEscape(parameterValueToString(r.checkpointId, "checkpointId")), -1)
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

type GetLegacyBackupRestoreJobApiRequest struct {
	ctx         context.Context
	ApiService  LegacyBackupApi
	groupId     string
	clusterName string
	jobId       string
}

type GetLegacyBackupRestoreJobApiParams struct {
	GroupId     string
	ClusterName string
	JobId       string
}

func (a *LegacyBackupApiService) GetLegacyBackupRestoreJobWithParams(ctx context.Context, args *GetLegacyBackupRestoreJobApiParams) GetLegacyBackupRestoreJobApiRequest {
	return GetLegacyBackupRestoreJobApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		jobId:       args.JobId,
	}
}

func (r GetLegacyBackupRestoreJobApiRequest) Execute() (*BackupRestoreJob, *http.Response, error) {
	return r.ApiService.GetLegacyBackupRestoreJobExecute(r)
}

/*
GetLegacyBackupRestoreJob Return One Legacy Backup Restore Job

[experimental] Returns one legacy backup restore job for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
	@param jobId Unique 24-hexadecimal digit string that identifies the restore job.
	@return GetLegacyBackupRestoreJobApiRequest

Deprecated
*/
func (a *LegacyBackupApiService) GetLegacyBackupRestoreJob(ctx context.Context, groupId string, clusterName string, jobId string) GetLegacyBackupRestoreJobApiRequest {
	return GetLegacyBackupRestoreJobApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		jobId:       jobId,
	}
}

// Execute executes the request
//
//	@return BackupRestoreJob
//
// Deprecated
func (a *LegacyBackupApiService) GetLegacyBackupRestoreJobExecute(r GetLegacyBackupRestoreJobApiRequest) (*BackupRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackupRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupApiService.GetLegacyBackupRestoreJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs/{jobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

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

type GetLegacySnapshotApiRequest struct {
	ctx         context.Context
	ApiService  LegacyBackupApi
	groupId     string
	clusterName string
	snapshotId  string
}

type GetLegacySnapshotApiParams struct {
	GroupId     string
	ClusterName string
	SnapshotId  string
}

func (a *LegacyBackupApiService) GetLegacySnapshotWithParams(ctx context.Context, args *GetLegacySnapshotApiParams) GetLegacySnapshotApiRequest {
	return GetLegacySnapshotApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		snapshotId:  args.SnapshotId,
	}
}

func (r GetLegacySnapshotApiRequest) Execute() (*BackupSnapshot, *http.Response, error) {
	return r.ApiService.GetLegacySnapshotExecute(r)
}

/*
GetLegacySnapshot Return One Legacy Backup Snapshot

[experimental] Returns one legacy backup snapshot for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role. Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
	@return GetLegacySnapshotApiRequest

Deprecated
*/
func (a *LegacyBackupApiService) GetLegacySnapshot(ctx context.Context, groupId string, clusterName string, snapshotId string) GetLegacySnapshotApiRequest {
	return GetLegacySnapshotApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		snapshotId:  snapshotId,
	}
}

// Execute executes the request
//
//	@return BackupSnapshot
//
// Deprecated
func (a *LegacyBackupApiService) GetLegacySnapshotExecute(r GetLegacySnapshotApiRequest) (*BackupSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackupSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupApiService.GetLegacySnapshot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

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

type GetLegacySnapshotScheduleApiRequest struct {
	ctx         context.Context
	ApiService  LegacyBackupApi
	groupId     string
	clusterName string
}

type GetLegacySnapshotScheduleApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *LegacyBackupApiService) GetLegacySnapshotScheduleWithParams(ctx context.Context, args *GetLegacySnapshotScheduleApiParams) GetLegacySnapshotScheduleApiRequest {
	return GetLegacySnapshotScheduleApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetLegacySnapshotScheduleApiRequest) Execute() (*ApiAtlasSnapshotSchedule, *http.Response, error) {
	return r.ApiService.GetLegacySnapshotScheduleExecute(r)
}

/*
GetLegacySnapshotSchedule Return One Snapshot Schedule

[experimental] Returns the snapshot schedule for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
	@return GetLegacySnapshotScheduleApiRequest

Deprecated
*/
func (a *LegacyBackupApiService) GetLegacySnapshotSchedule(ctx context.Context, groupId string, clusterName string) GetLegacySnapshotScheduleApiRequest {
	return GetLegacySnapshotScheduleApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return ApiAtlasSnapshotSchedule
//
// Deprecated
func (a *LegacyBackupApiService) GetLegacySnapshotScheduleExecute(r GetLegacySnapshotScheduleApiRequest) (*ApiAtlasSnapshotSchedule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiAtlasSnapshotSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupApiService.GetLegacySnapshotSchedule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule"
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

type ListLegacyBackupCheckpointsApiRequest struct {
	ctx          context.Context
	ApiService   LegacyBackupApi
	groupId      string
	clusterName  string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListLegacyBackupCheckpointsApiParams struct {
	GroupId      string
	ClusterName  string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *LegacyBackupApiService) ListLegacyBackupCheckpointsWithParams(ctx context.Context, args *ListLegacyBackupCheckpointsApiParams) ListLegacyBackupCheckpointsApiRequest {
	return ListLegacyBackupCheckpointsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListLegacyBackupCheckpointsApiRequest) IncludeCount(includeCount bool) ListLegacyBackupCheckpointsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListLegacyBackupCheckpointsApiRequest) ItemsPerPage(itemsPerPage int) ListLegacyBackupCheckpointsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListLegacyBackupCheckpointsApiRequest) PageNum(pageNum int) ListLegacyBackupCheckpointsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListLegacyBackupCheckpointsApiRequest) Execute() (*PaginatedApiAtlasCheckpoint, *http.Response, error) {
	return r.ApiService.ListLegacyBackupCheckpointsExecute(r)
}

/*
ListLegacyBackupCheckpoints Return All Legacy Backup Checkpoints

[experimental] Returns all legacy backup checkpoints for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster that contains the checkpoints that you want to return.
	@return ListLegacyBackupCheckpointsApiRequest

Deprecated
*/
func (a *LegacyBackupApiService) ListLegacyBackupCheckpoints(ctx context.Context, groupId string, clusterName string) ListLegacyBackupCheckpointsApiRequest {
	return ListLegacyBackupCheckpointsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PaginatedApiAtlasCheckpoint
//
// Deprecated
func (a *LegacyBackupApiService) ListLegacyBackupCheckpointsExecute(r ListLegacyBackupCheckpointsApiRequest) (*PaginatedApiAtlasCheckpoint, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAtlasCheckpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupApiService.ListLegacyBackupCheckpoints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backupCheckpoints"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

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

type ListLegacyBackupRestoreJobsApiRequest struct {
	ctx          context.Context
	ApiService   LegacyBackupApi
	groupId      string
	clusterName  string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
	batchId      *string
}

type ListLegacyBackupRestoreJobsApiParams struct {
	GroupId      string
	ClusterName  string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
	BatchId      *string
}

func (a *LegacyBackupApiService) ListLegacyBackupRestoreJobsWithParams(ctx context.Context, args *ListLegacyBackupRestoreJobsApiParams) ListLegacyBackupRestoreJobsApiRequest {
	return ListLegacyBackupRestoreJobsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		batchId:      args.BatchId,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListLegacyBackupRestoreJobsApiRequest) IncludeCount(includeCount bool) ListLegacyBackupRestoreJobsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListLegacyBackupRestoreJobsApiRequest) ItemsPerPage(itemsPerPage int) ListLegacyBackupRestoreJobsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListLegacyBackupRestoreJobsApiRequest) PageNum(pageNum int) ListLegacyBackupRestoreJobsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Unique 24-hexadecimal digit string that identifies the batch of restore jobs to return. Timestamp in ISO 8601 date and time format in UTC when creating a restore job for a sharded cluster, Application creates a separate job for each shard, plus another for the config host. Each of these jobs comprise one batch. A restore job for a replica set can&#39;t be part of a batch.
func (r ListLegacyBackupRestoreJobsApiRequest) BatchId(batchId string) ListLegacyBackupRestoreJobsApiRequest {
	r.batchId = &batchId
	return r
}

func (r ListLegacyBackupRestoreJobsApiRequest) Execute() (*PaginatedRestoreJob, *http.Response, error) {
	return r.ApiService.ListLegacyBackupRestoreJobsExecute(r)
}

/*
ListLegacyBackupRestoreJobs Return All Legacy Backup Restore Jobs

[experimental] Returns all legacy backup restore jobs for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule). If you use the `BATCH-ID` query parameter, you can retrieve all restore jobs in the specified batch. When creating a restore job for a sharded cluster, MongoDB Cloud creates a separate job for each shard, plus another for the config server. Each of those jobs are part of a batch. However, a batch can't include a restore job for a replica set.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
	@return ListLegacyBackupRestoreJobsApiRequest

Deprecated
*/
func (a *LegacyBackupApiService) ListLegacyBackupRestoreJobs(ctx context.Context, groupId string, clusterName string) ListLegacyBackupRestoreJobsApiRequest {
	return ListLegacyBackupRestoreJobsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PaginatedRestoreJob
//
// Deprecated
func (a *LegacyBackupApiService) ListLegacyBackupRestoreJobsExecute(r ListLegacyBackupRestoreJobsApiRequest) (*PaginatedRestoreJob, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedRestoreJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupApiService.ListLegacyBackupRestoreJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restoreJobs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

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
	if r.batchId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "batchId", r.batchId, "")
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

type ListLegacySnapshotsApiRequest struct {
	ctx          context.Context
	ApiService   LegacyBackupApi
	groupId      string
	clusterName  string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
	completed    *string
}

type ListLegacySnapshotsApiParams struct {
	GroupId      string
	ClusterName  string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
	Completed    *string
}

func (a *LegacyBackupApiService) ListLegacySnapshotsWithParams(ctx context.Context, args *ListLegacySnapshotsApiParams) ListLegacySnapshotsApiRequest {
	return ListLegacySnapshotsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		clusterName:  args.ClusterName,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		completed:    args.Completed,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListLegacySnapshotsApiRequest) IncludeCount(includeCount bool) ListLegacySnapshotsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListLegacySnapshotsApiRequest) ItemsPerPage(itemsPerPage int) ListLegacySnapshotsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListLegacySnapshotsApiRequest) PageNum(pageNum int) ListLegacySnapshotsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Human-readable label that specifies whether to return only completed, incomplete, or all snapshots. By default, MongoDB Cloud only returns completed snapshots.
func (r ListLegacySnapshotsApiRequest) Completed(completed string) ListLegacySnapshotsApiRequest {
	r.completed = &completed
	return r
}

func (r ListLegacySnapshotsApiRequest) Execute() (*PaginatedSnapshot, *http.Response, error) {
	return r.ApiService.ListLegacySnapshotsExecute(r)
}

/*
ListLegacySnapshots Return All Legacy Backup Snapshots

[experimental] Returns all legacy backup snapshots for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Read Only role. Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return ListLegacySnapshotsApiRequest

Deprecated
*/
func (a *LegacyBackupApiService) ListLegacySnapshots(ctx context.Context, groupId string, clusterName string) ListLegacySnapshotsApiRequest {
	return ListLegacySnapshotsApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//
//	@return PaginatedSnapshot
//
// Deprecated
func (a *LegacyBackupApiService) ListLegacySnapshotsExecute(r ListLegacySnapshotsApiRequest) (*PaginatedSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupApiService.ListLegacySnapshots")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

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
	if r.completed != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "completed", r.completed, "")
	} else {
		var defaultValue string = "true"
		r.completed = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "completed", r.completed, "")
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

type UpdateLegacySnapshotRetentionApiRequest struct {
	ctx            context.Context
	ApiService     LegacyBackupApi
	groupId        string
	clusterName    string
	snapshotId     string
	backupSnapshot *BackupSnapshot
}

type UpdateLegacySnapshotRetentionApiParams struct {
	GroupId        string
	ClusterName    string
	SnapshotId     string
	BackupSnapshot *BackupSnapshot
}

func (a *LegacyBackupApiService) UpdateLegacySnapshotRetentionWithParams(ctx context.Context, args *UpdateLegacySnapshotRetentionApiParams) UpdateLegacySnapshotRetentionApiRequest {
	return UpdateLegacySnapshotRetentionApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		clusterName:    args.ClusterName,
		snapshotId:     args.SnapshotId,
		backupSnapshot: args.BackupSnapshot,
	}
}

func (r UpdateLegacySnapshotRetentionApiRequest) Execute() (*BackupSnapshot, *http.Response, error) {
	return r.ApiService.UpdateLegacySnapshotRetentionExecute(r)
}

/*
UpdateLegacySnapshotRetention Change One Legacy Backup Snapshot Expiration

[experimental] Changes the expiration date for one legacy backup snapshot for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role. Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
	@return UpdateLegacySnapshotRetentionApiRequest

Deprecated
*/
func (a *LegacyBackupApiService) UpdateLegacySnapshotRetention(ctx context.Context, groupId string, clusterName string, snapshotId string, backupSnapshot *BackupSnapshot) UpdateLegacySnapshotRetentionApiRequest {
	return UpdateLegacySnapshotRetentionApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		clusterName:    clusterName,
		snapshotId:     snapshotId,
		backupSnapshot: backupSnapshot,
	}
}

// Execute executes the request
//
//	@return BackupSnapshot
//
// Deprecated
func (a *LegacyBackupApiService) UpdateLegacySnapshotRetentionExecute(r UpdateLegacySnapshotRetentionApiRequest) (*BackupSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BackupSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupApiService.UpdateLegacySnapshotRetention")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterValueToString(r.snapshotId, "snapshotId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.backupSnapshot == nil {
		return localVarReturnValue, nil, reportError("backupSnapshot is required and must be specified")
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
	localVarPostBody = r.backupSnapshot
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

type UpdateLegacySnapshotScheduleApiRequest struct {
	ctx                      context.Context
	ApiService               LegacyBackupApi
	groupId                  string
	clusterName              string
	apiAtlasSnapshotSchedule *ApiAtlasSnapshotSchedule
}

type UpdateLegacySnapshotScheduleApiParams struct {
	GroupId                  string
	ClusterName              string
	ApiAtlasSnapshotSchedule *ApiAtlasSnapshotSchedule
}

func (a *LegacyBackupApiService) UpdateLegacySnapshotScheduleWithParams(ctx context.Context, args *UpdateLegacySnapshotScheduleApiParams) UpdateLegacySnapshotScheduleApiRequest {
	return UpdateLegacySnapshotScheduleApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  args.GroupId,
		clusterName:              args.ClusterName,
		apiAtlasSnapshotSchedule: args.ApiAtlasSnapshotSchedule,
	}
}

func (r UpdateLegacySnapshotScheduleApiRequest) Execute() (*ApiAtlasSnapshotSchedule, *http.Response, error) {
	return r.ApiService.UpdateLegacySnapshotScheduleExecute(r)
}

/*
UpdateLegacySnapshotSchedule Update Snapshot Schedule for One Cluster

[experimental] Updates the snapshot schedule for one cluster in the specified project. To use this resource, the requesting API Key must have the Project Owner role.

	Effective 23 March 2020, all new clusters can use only Cloud Backups. When you upgrade to 4.2, your backup system upgrades to cloud backup if it is currently set to legacy backup. After this upgrade, all your existing legacy backup snapshots remain available. They expire over time in accordance with your retention policy. Your backup policy resets to the default schedule. If you had a custom backup policy in place with legacy backups, you must re-create it with the procedure outlined in the [Cloud Backup documentation](https://www.mongodb.com/docs/atlas/backup/cloud-backup/scheduling/#std-label-cloud-provider-backup-schedule).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster with the snapshot you want to return.
	@return UpdateLegacySnapshotScheduleApiRequest

Deprecated
*/
func (a *LegacyBackupApiService) UpdateLegacySnapshotSchedule(ctx context.Context, groupId string, clusterName string, apiAtlasSnapshotSchedule *ApiAtlasSnapshotSchedule) UpdateLegacySnapshotScheduleApiRequest {
	return UpdateLegacySnapshotScheduleApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  groupId,
		clusterName:              clusterName,
		apiAtlasSnapshotSchedule: apiAtlasSnapshotSchedule,
	}
}

// Execute executes the request
//
//	@return ApiAtlasSnapshotSchedule
//
// Deprecated
func (a *LegacyBackupApiService) UpdateLegacySnapshotScheduleExecute(r UpdateLegacySnapshotScheduleApiRequest) (*ApiAtlasSnapshotSchedule, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiAtlasSnapshotSchedule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyBackupApiService.UpdateLegacySnapshotSchedule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/snapshotSchedule"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterValueToString(r.clusterName, "clusterName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiAtlasSnapshotSchedule == nil {
		return localVarReturnValue, nil, reportError("apiAtlasSnapshotSchedule is required and must be specified")
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
	localVarPostBody = r.apiAtlasSnapshotSchedule
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
