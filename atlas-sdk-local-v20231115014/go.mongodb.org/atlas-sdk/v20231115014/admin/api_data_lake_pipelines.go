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
	"time"
)

type DataLakePipelinesApi interface {

	/*
		CreatePipeline Create One Data Lake Pipeline

		Creates one Data Lake Pipeline.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return CreatePipelineApiRequest
	*/
	CreatePipeline(ctx context.Context, groupId string, dataLakeIngestionPipeline *DataLakeIngestionPipeline) CreatePipelineApiRequest
	/*
		CreatePipeline Create One Data Lake Pipeline


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreatePipelineApiParams - Parameters for the request
		@return CreatePipelineApiRequest
	*/
	CreatePipelineWithParams(ctx context.Context, args *CreatePipelineApiParams) CreatePipelineApiRequest

	// Method available only for mocking purposes
	CreatePipelineExecute(r CreatePipelineApiRequest) (*DataLakeIngestionPipeline, *http.Response, error)

	/*
		DeletePipeline Remove One Data Lake Pipeline

		Removes one Data Lake Pipeline.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
		@return DeletePipelineApiRequest
	*/
	DeletePipeline(ctx context.Context, groupId string, pipelineName string) DeletePipelineApiRequest
	/*
		DeletePipeline Remove One Data Lake Pipeline


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeletePipelineApiParams - Parameters for the request
		@return DeletePipelineApiRequest
	*/
	DeletePipelineWithParams(ctx context.Context, args *DeletePipelineApiParams) DeletePipelineApiRequest

	// Method available only for mocking purposes
	DeletePipelineExecute(r DeletePipelineApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		DeletePipelineRunDataset Delete Pipeline Run Dataset

		Deletes dataset that Atlas generated during the specified pipeline run.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
		@param pipelineRunId Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.
		@return DeletePipelineRunDatasetApiRequest
	*/
	DeletePipelineRunDataset(ctx context.Context, groupId string, pipelineName string, pipelineRunId string) DeletePipelineRunDatasetApiRequest
	/*
		DeletePipelineRunDataset Delete Pipeline Run Dataset


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeletePipelineRunDatasetApiParams - Parameters for the request
		@return DeletePipelineRunDatasetApiRequest
	*/
	DeletePipelineRunDatasetWithParams(ctx context.Context, args *DeletePipelineRunDatasetApiParams) DeletePipelineRunDatasetApiRequest

	// Method available only for mocking purposes
	DeletePipelineRunDatasetExecute(r DeletePipelineRunDatasetApiRequest) (map[string]interface{}, *http.Response, error)

	/*
		GetPipeline Return One Data Lake Pipeline

		Returns the details of one Data Lake Pipeline within the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
		@return GetPipelineApiRequest
	*/
	GetPipeline(ctx context.Context, groupId string, pipelineName string) GetPipelineApiRequest
	/*
		GetPipeline Return One Data Lake Pipeline


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetPipelineApiParams - Parameters for the request
		@return GetPipelineApiRequest
	*/
	GetPipelineWithParams(ctx context.Context, args *GetPipelineApiParams) GetPipelineApiRequest

	// Method available only for mocking purposes
	GetPipelineExecute(r GetPipelineApiRequest) (*DataLakeIngestionPipeline, *http.Response, error)

	/*
		GetPipelineRun Return One Data Lake Pipeline Run

		Returns the details of one Data Lake Pipeline run within the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
		@param pipelineRunId Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.
		@return GetPipelineRunApiRequest
	*/
	GetPipelineRun(ctx context.Context, groupId string, pipelineName string, pipelineRunId string) GetPipelineRunApiRequest
	/*
		GetPipelineRun Return One Data Lake Pipeline Run


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetPipelineRunApiParams - Parameters for the request
		@return GetPipelineRunApiRequest
	*/
	GetPipelineRunWithParams(ctx context.Context, args *GetPipelineRunApiParams) GetPipelineRunApiRequest

	// Method available only for mocking purposes
	GetPipelineRunExecute(r GetPipelineRunApiRequest) (*IngestionPipelineRun, *http.Response, error)

	/*
		ListPipelineRuns Return All Data Lake Pipeline Runs from One Project

		Returns a list of past Data Lake Pipeline runs. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
		@return ListPipelineRunsApiRequest
	*/
	ListPipelineRuns(ctx context.Context, groupId string, pipelineName string) ListPipelineRunsApiRequest
	/*
		ListPipelineRuns Return All Data Lake Pipeline Runs from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListPipelineRunsApiParams - Parameters for the request
		@return ListPipelineRunsApiRequest
	*/
	ListPipelineRunsWithParams(ctx context.Context, args *ListPipelineRunsApiParams) ListPipelineRunsApiRequest

	// Method available only for mocking purposes
	ListPipelineRunsExecute(r ListPipelineRunsApiRequest) (*PaginatedPipelineRun, *http.Response, error)

	/*
		ListPipelineSchedules Return Available Ingestion Schedules for One Data Lake Pipeline

		Returns a list of backup schedule policy items that you can use as a Data Lake Pipeline source. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
		@return ListPipelineSchedulesApiRequest
	*/
	ListPipelineSchedules(ctx context.Context, groupId string, pipelineName string) ListPipelineSchedulesApiRequest
	/*
		ListPipelineSchedules Return Available Ingestion Schedules for One Data Lake Pipeline


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListPipelineSchedulesApiParams - Parameters for the request
		@return ListPipelineSchedulesApiRequest
	*/
	ListPipelineSchedulesWithParams(ctx context.Context, args *ListPipelineSchedulesApiParams) ListPipelineSchedulesApiRequest

	// Method available only for mocking purposes
	ListPipelineSchedulesExecute(r ListPipelineSchedulesApiRequest) ([]DiskBackupApiPolicyItem, *http.Response, error)

	/*
		ListPipelineSnapshots Return Available Backup Snapshots for One Data Lake Pipeline

		Returns a list of backup snapshots that you can use to trigger an on demand pipeline run. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
		@return ListPipelineSnapshotsApiRequest
	*/
	ListPipelineSnapshots(ctx context.Context, groupId string, pipelineName string) ListPipelineSnapshotsApiRequest
	/*
		ListPipelineSnapshots Return Available Backup Snapshots for One Data Lake Pipeline


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListPipelineSnapshotsApiParams - Parameters for the request
		@return ListPipelineSnapshotsApiRequest
	*/
	ListPipelineSnapshotsWithParams(ctx context.Context, args *ListPipelineSnapshotsApiParams) ListPipelineSnapshotsApiRequest

	// Method available only for mocking purposes
	ListPipelineSnapshotsExecute(r ListPipelineSnapshotsApiRequest) (*PaginatedBackupSnapshot, *http.Response, error)

	/*
		ListPipelines Return All Data Lake Pipelines from One Project

		Returns a list of Data Lake Pipelines. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListPipelinesApiRequest
	*/
	ListPipelines(ctx context.Context, groupId string) ListPipelinesApiRequest
	/*
		ListPipelines Return All Data Lake Pipelines from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListPipelinesApiParams - Parameters for the request
		@return ListPipelinesApiRequest
	*/
	ListPipelinesWithParams(ctx context.Context, args *ListPipelinesApiParams) ListPipelinesApiRequest

	// Method available only for mocking purposes
	ListPipelinesExecute(r ListPipelinesApiRequest) ([]DataLakeIngestionPipeline, *http.Response, error)

	/*
		PausePipeline Pause One Data Lake Pipeline

		Pauses ingestion for a Data Lake Pipeline within the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
		@return PausePipelineApiRequest
	*/
	PausePipeline(ctx context.Context, groupId string, pipelineName string) PausePipelineApiRequest
	/*
		PausePipeline Pause One Data Lake Pipeline


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param PausePipelineApiParams - Parameters for the request
		@return PausePipelineApiRequest
	*/
	PausePipelineWithParams(ctx context.Context, args *PausePipelineApiParams) PausePipelineApiRequest

	// Method available only for mocking purposes
	PausePipelineExecute(r PausePipelineApiRequest) (*DataLakeIngestionPipeline, *http.Response, error)

	/*
		ResumePipeline Resume One Data Lake Pipeline

		Resumes ingestion for a Data Lake Pipeline within the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
		@return ResumePipelineApiRequest
	*/
	ResumePipeline(ctx context.Context, groupId string, pipelineName string) ResumePipelineApiRequest
	/*
		ResumePipeline Resume One Data Lake Pipeline


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ResumePipelineApiParams - Parameters for the request
		@return ResumePipelineApiRequest
	*/
	ResumePipelineWithParams(ctx context.Context, args *ResumePipelineApiParams) ResumePipelineApiRequest

	// Method available only for mocking purposes
	ResumePipelineExecute(r ResumePipelineApiRequest) (*DataLakeIngestionPipeline, *http.Response, error)

	/*
		TriggerSnapshotIngestion Trigger on demand snapshot ingestion

		Triggers a Data Lake Pipeline ingestion of a specified snapshot.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
		@return TriggerSnapshotIngestionApiRequest
	*/
	TriggerSnapshotIngestion(ctx context.Context, groupId string, pipelineName string, triggerIngestionPipelineRequest *TriggerIngestionPipelineRequest) TriggerSnapshotIngestionApiRequest
	/*
		TriggerSnapshotIngestion Trigger on demand snapshot ingestion


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param TriggerSnapshotIngestionApiParams - Parameters for the request
		@return TriggerSnapshotIngestionApiRequest
	*/
	TriggerSnapshotIngestionWithParams(ctx context.Context, args *TriggerSnapshotIngestionApiParams) TriggerSnapshotIngestionApiRequest

	// Method available only for mocking purposes
	TriggerSnapshotIngestionExecute(r TriggerSnapshotIngestionApiRequest) (*IngestionPipelineRun, *http.Response, error)

	/*
		UpdatePipeline Update One Data Lake Pipeline

		Updates one Data Lake Pipeline.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
		@return UpdatePipelineApiRequest
	*/
	UpdatePipeline(ctx context.Context, groupId string, pipelineName string, dataLakeIngestionPipeline *DataLakeIngestionPipeline) UpdatePipelineApiRequest
	/*
		UpdatePipeline Update One Data Lake Pipeline


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdatePipelineApiParams - Parameters for the request
		@return UpdatePipelineApiRequest
	*/
	UpdatePipelineWithParams(ctx context.Context, args *UpdatePipelineApiParams) UpdatePipelineApiRequest

	// Method available only for mocking purposes
	UpdatePipelineExecute(r UpdatePipelineApiRequest) (*DataLakeIngestionPipeline, *http.Response, error)
}

// DataLakePipelinesApiService DataLakePipelinesApi service
type DataLakePipelinesApiService service

type CreatePipelineApiRequest struct {
	ctx                       context.Context
	ApiService                DataLakePipelinesApi
	groupId                   string
	dataLakeIngestionPipeline *DataLakeIngestionPipeline
}

type CreatePipelineApiParams struct {
	GroupId                   string
	DataLakeIngestionPipeline *DataLakeIngestionPipeline
}

func (a *DataLakePipelinesApiService) CreatePipelineWithParams(ctx context.Context, args *CreatePipelineApiParams) CreatePipelineApiRequest {
	return CreatePipelineApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   args.GroupId,
		dataLakeIngestionPipeline: args.DataLakeIngestionPipeline,
	}
}

func (r CreatePipelineApiRequest) Execute() (*DataLakeIngestionPipeline, *http.Response, error) {
	return r.ApiService.CreatePipelineExecute(r)
}

/*
CreatePipeline Create One Data Lake Pipeline

Creates one Data Lake Pipeline.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreatePipelineApiRequest
*/
func (a *DataLakePipelinesApiService) CreatePipeline(ctx context.Context, groupId string, dataLakeIngestionPipeline *DataLakeIngestionPipeline) CreatePipelineApiRequest {
	return CreatePipelineApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   groupId,
		dataLakeIngestionPipeline: dataLakeIngestionPipeline,
	}
}

// Execute executes the request
//
//	@return DataLakeIngestionPipeline
func (a *DataLakePipelinesApiService) CreatePipelineExecute(r CreatePipelineApiRequest) (*DataLakeIngestionPipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataLakeIngestionPipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.CreatePipeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataLakeIngestionPipeline == nil {
		return localVarReturnValue, nil, reportError("dataLakeIngestionPipeline is required and must be specified")
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
	localVarPostBody = r.dataLakeIngestionPipeline
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

type DeletePipelineApiRequest struct {
	ctx          context.Context
	ApiService   DataLakePipelinesApi
	groupId      string
	pipelineName string
}

type DeletePipelineApiParams struct {
	GroupId      string
	PipelineName string
}

func (a *DataLakePipelinesApiService) DeletePipelineWithParams(ctx context.Context, args *DeletePipelineApiParams) DeletePipelineApiRequest {
	return DeletePipelineApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		pipelineName: args.PipelineName,
	}
}

func (r DeletePipelineApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeletePipelineExecute(r)
}

/*
DeletePipeline Remove One Data Lake Pipeline

Removes one Data Lake Pipeline.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
	@return DeletePipelineApiRequest
*/
func (a *DataLakePipelinesApiService) DeletePipeline(ctx context.Context, groupId string, pipelineName string) DeletePipelineApiRequest {
	return DeletePipelineApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		pipelineName: pipelineName,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *DataLakePipelinesApiService) DeletePipelineExecute(r DeletePipelineApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.DeletePipeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineName"+"}", url.PathEscape(parameterValueToString(r.pipelineName, "pipelineName")), -1)

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

type DeletePipelineRunDatasetApiRequest struct {
	ctx           context.Context
	ApiService    DataLakePipelinesApi
	groupId       string
	pipelineName  string
	pipelineRunId string
}

type DeletePipelineRunDatasetApiParams struct {
	GroupId       string
	PipelineName  string
	PipelineRunId string
}

func (a *DataLakePipelinesApiService) DeletePipelineRunDatasetWithParams(ctx context.Context, args *DeletePipelineRunDatasetApiParams) DeletePipelineRunDatasetApiRequest {
	return DeletePipelineRunDatasetApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		pipelineName:  args.PipelineName,
		pipelineRunId: args.PipelineRunId,
	}
}

func (r DeletePipelineRunDatasetApiRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeletePipelineRunDatasetExecute(r)
}

/*
DeletePipelineRunDataset Delete Pipeline Run Dataset

Deletes dataset that Atlas generated during the specified pipeline run.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
	@param pipelineRunId Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.
	@return DeletePipelineRunDatasetApiRequest
*/
func (a *DataLakePipelinesApiService) DeletePipelineRunDataset(ctx context.Context, groupId string, pipelineName string, pipelineRunId string) DeletePipelineRunDatasetApiRequest {
	return DeletePipelineRunDatasetApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		pipelineName:  pipelineName,
		pipelineRunId: pipelineRunId,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *DataLakePipelinesApiService) DeletePipelineRunDatasetExecute(r DeletePipelineRunDatasetApiRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.DeletePipelineRunDataset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineName"+"}", url.PathEscape(parameterValueToString(r.pipelineName, "pipelineName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineRunId"+"}", url.PathEscape(parameterValueToString(r.pipelineRunId, "pipelineRunId")), -1)

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

type GetPipelineApiRequest struct {
	ctx          context.Context
	ApiService   DataLakePipelinesApi
	groupId      string
	pipelineName string
}

type GetPipelineApiParams struct {
	GroupId      string
	PipelineName string
}

func (a *DataLakePipelinesApiService) GetPipelineWithParams(ctx context.Context, args *GetPipelineApiParams) GetPipelineApiRequest {
	return GetPipelineApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		pipelineName: args.PipelineName,
	}
}

func (r GetPipelineApiRequest) Execute() (*DataLakeIngestionPipeline, *http.Response, error) {
	return r.ApiService.GetPipelineExecute(r)
}

/*
GetPipeline Return One Data Lake Pipeline

Returns the details of one Data Lake Pipeline within the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
	@return GetPipelineApiRequest
*/
func (a *DataLakePipelinesApiService) GetPipeline(ctx context.Context, groupId string, pipelineName string) GetPipelineApiRequest {
	return GetPipelineApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		pipelineName: pipelineName,
	}
}

// Execute executes the request
//
//	@return DataLakeIngestionPipeline
func (a *DataLakePipelinesApiService) GetPipelineExecute(r GetPipelineApiRequest) (*DataLakeIngestionPipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataLakeIngestionPipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.GetPipeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineName"+"}", url.PathEscape(parameterValueToString(r.pipelineName, "pipelineName")), -1)

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

type GetPipelineRunApiRequest struct {
	ctx           context.Context
	ApiService    DataLakePipelinesApi
	groupId       string
	pipelineName  string
	pipelineRunId string
}

type GetPipelineRunApiParams struct {
	GroupId       string
	PipelineName  string
	PipelineRunId string
}

func (a *DataLakePipelinesApiService) GetPipelineRunWithParams(ctx context.Context, args *GetPipelineRunApiParams) GetPipelineRunApiRequest {
	return GetPipelineRunApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		pipelineName:  args.PipelineName,
		pipelineRunId: args.PipelineRunId,
	}
}

func (r GetPipelineRunApiRequest) Execute() (*IngestionPipelineRun, *http.Response, error) {
	return r.ApiService.GetPipelineRunExecute(r)
}

/*
GetPipelineRun Return One Data Lake Pipeline Run

Returns the details of one Data Lake Pipeline run within the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
	@param pipelineRunId Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run.
	@return GetPipelineRunApiRequest
*/
func (a *DataLakePipelinesApiService) GetPipelineRun(ctx context.Context, groupId string, pipelineName string, pipelineRunId string) GetPipelineRunApiRequest {
	return GetPipelineRunApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		pipelineName:  pipelineName,
		pipelineRunId: pipelineRunId,
	}
}

// Execute executes the request
//
//	@return IngestionPipelineRun
func (a *DataLakePipelinesApiService) GetPipelineRunExecute(r GetPipelineRunApiRequest) (*IngestionPipelineRun, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IngestionPipelineRun
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.GetPipelineRun")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs/{pipelineRunId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineName"+"}", url.PathEscape(parameterValueToString(r.pipelineName, "pipelineName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineRunId"+"}", url.PathEscape(parameterValueToString(r.pipelineRunId, "pipelineRunId")), -1)

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

type ListPipelineRunsApiRequest struct {
	ctx           context.Context
	ApiService    DataLakePipelinesApi
	groupId       string
	pipelineName  string
	includeCount  *bool
	itemsPerPage  *int
	pageNum       *int
	createdBefore *time.Time
}

type ListPipelineRunsApiParams struct {
	GroupId       string
	PipelineName  string
	IncludeCount  *bool
	ItemsPerPage  *int
	PageNum       *int
	CreatedBefore *time.Time
}

func (a *DataLakePipelinesApiService) ListPipelineRunsWithParams(ctx context.Context, args *ListPipelineRunsApiParams) ListPipelineRunsApiRequest {
	return ListPipelineRunsApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		pipelineName:  args.PipelineName,
		includeCount:  args.IncludeCount,
		itemsPerPage:  args.ItemsPerPage,
		pageNum:       args.PageNum,
		createdBefore: args.CreatedBefore,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListPipelineRunsApiRequest) IncludeCount(includeCount bool) ListPipelineRunsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListPipelineRunsApiRequest) ItemsPerPage(itemsPerPage int) ListPipelineRunsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListPipelineRunsApiRequest) PageNum(pageNum int) ListPipelineRunsApiRequest {
	r.pageNum = &pageNum
	return r
}

// If specified, Atlas returns only Data Lake Pipeline runs initiated before this time and date.
func (r ListPipelineRunsApiRequest) CreatedBefore(createdBefore time.Time) ListPipelineRunsApiRequest {
	r.createdBefore = &createdBefore
	return r
}

func (r ListPipelineRunsApiRequest) Execute() (*PaginatedPipelineRun, *http.Response, error) {
	return r.ApiService.ListPipelineRunsExecute(r)
}

/*
ListPipelineRuns Return All Data Lake Pipeline Runs from One Project

Returns a list of past Data Lake Pipeline runs. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
	@return ListPipelineRunsApiRequest
*/
func (a *DataLakePipelinesApiService) ListPipelineRuns(ctx context.Context, groupId string, pipelineName string) ListPipelineRunsApiRequest {
	return ListPipelineRunsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		pipelineName: pipelineName,
	}
}

// Execute executes the request
//
//	@return PaginatedPipelineRun
func (a *DataLakePipelinesApiService) ListPipelineRunsExecute(r ListPipelineRunsApiRequest) (*PaginatedPipelineRun, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedPipelineRun
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.ListPipelineRuns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/runs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineName"+"}", url.PathEscape(parameterValueToString(r.pipelineName, "pipelineName")), -1)

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
	if r.createdBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdBefore", r.createdBefore, "")
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

type ListPipelineSchedulesApiRequest struct {
	ctx          context.Context
	ApiService   DataLakePipelinesApi
	groupId      string
	pipelineName string
}

type ListPipelineSchedulesApiParams struct {
	GroupId      string
	PipelineName string
}

func (a *DataLakePipelinesApiService) ListPipelineSchedulesWithParams(ctx context.Context, args *ListPipelineSchedulesApiParams) ListPipelineSchedulesApiRequest {
	return ListPipelineSchedulesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		pipelineName: args.PipelineName,
	}
}

func (r ListPipelineSchedulesApiRequest) Execute() ([]DiskBackupApiPolicyItem, *http.Response, error) {
	return r.ApiService.ListPipelineSchedulesExecute(r)
}

/*
ListPipelineSchedules Return Available Ingestion Schedules for One Data Lake Pipeline

Returns a list of backup schedule policy items that you can use as a Data Lake Pipeline source. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
	@return ListPipelineSchedulesApiRequest
*/
func (a *DataLakePipelinesApiService) ListPipelineSchedules(ctx context.Context, groupId string, pipelineName string) ListPipelineSchedulesApiRequest {
	return ListPipelineSchedulesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		pipelineName: pipelineName,
	}
}

// Execute executes the request
//
//	@return []DiskBackupApiPolicyItem
func (a *DataLakePipelinesApiService) ListPipelineSchedulesExecute(r ListPipelineSchedulesApiRequest) ([]DiskBackupApiPolicyItem, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []DiskBackupApiPolicyItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.ListPipelineSchedules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSchedules"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineName"+"}", url.PathEscape(parameterValueToString(r.pipelineName, "pipelineName")), -1)

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

type ListPipelineSnapshotsApiRequest struct {
	ctx            context.Context
	ApiService     DataLakePipelinesApi
	groupId        string
	pipelineName   string
	includeCount   *bool
	itemsPerPage   *int
	pageNum        *int
	completedAfter *time.Time
}

type ListPipelineSnapshotsApiParams struct {
	GroupId        string
	PipelineName   string
	IncludeCount   *bool
	ItemsPerPage   *int
	PageNum        *int
	CompletedAfter *time.Time
}

func (a *DataLakePipelinesApiService) ListPipelineSnapshotsWithParams(ctx context.Context, args *ListPipelineSnapshotsApiParams) ListPipelineSnapshotsApiRequest {
	return ListPipelineSnapshotsApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		pipelineName:   args.PipelineName,
		includeCount:   args.IncludeCount,
		itemsPerPage:   args.ItemsPerPage,
		pageNum:        args.PageNum,
		completedAfter: args.CompletedAfter,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListPipelineSnapshotsApiRequest) IncludeCount(includeCount bool) ListPipelineSnapshotsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListPipelineSnapshotsApiRequest) ItemsPerPage(itemsPerPage int) ListPipelineSnapshotsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListPipelineSnapshotsApiRequest) PageNum(pageNum int) ListPipelineSnapshotsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Date and time after which MongoDB Cloud created the snapshot. If specified, MongoDB Cloud returns available backup snapshots created after this time and date only. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
func (r ListPipelineSnapshotsApiRequest) CompletedAfter(completedAfter time.Time) ListPipelineSnapshotsApiRequest {
	r.completedAfter = &completedAfter
	return r
}

func (r ListPipelineSnapshotsApiRequest) Execute() (*PaginatedBackupSnapshot, *http.Response, error) {
	return r.ApiService.ListPipelineSnapshotsExecute(r)
}

/*
ListPipelineSnapshots Return Available Backup Snapshots for One Data Lake Pipeline

Returns a list of backup snapshots that you can use to trigger an on demand pipeline run. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
	@return ListPipelineSnapshotsApiRequest
*/
func (a *DataLakePipelinesApiService) ListPipelineSnapshots(ctx context.Context, groupId string, pipelineName string) ListPipelineSnapshotsApiRequest {
	return ListPipelineSnapshotsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		pipelineName: pipelineName,
	}
}

// Execute executes the request
//
//	@return PaginatedBackupSnapshot
func (a *DataLakePipelinesApiService) ListPipelineSnapshotsExecute(r ListPipelineSnapshotsApiRequest) (*PaginatedBackupSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedBackupSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.ListPipelineSnapshots")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/availableSnapshots"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineName"+"}", url.PathEscape(parameterValueToString(r.pipelineName, "pipelineName")), -1)

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
	if r.completedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "completedAfter", r.completedAfter, "")
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

type ListPipelinesApiRequest struct {
	ctx        context.Context
	ApiService DataLakePipelinesApi
	groupId    string
}

type ListPipelinesApiParams struct {
	GroupId string
}

func (a *DataLakePipelinesApiService) ListPipelinesWithParams(ctx context.Context, args *ListPipelinesApiParams) ListPipelinesApiRequest {
	return ListPipelinesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
	}
}

func (r ListPipelinesApiRequest) Execute() ([]DataLakeIngestionPipeline, *http.Response, error) {
	return r.ApiService.ListPipelinesExecute(r)
}

/*
ListPipelines Return All Data Lake Pipelines from One Project

Returns a list of Data Lake Pipelines. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListPipelinesApiRequest
*/
func (a *DataLakePipelinesApiService) ListPipelines(ctx context.Context, groupId string) ListPipelinesApiRequest {
	return ListPipelinesApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// Execute executes the request
//
//	@return []DataLakeIngestionPipeline
func (a *DataLakePipelinesApiService) ListPipelinesExecute(r ListPipelinesApiRequest) ([]DataLakeIngestionPipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []DataLakeIngestionPipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.ListPipelines")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines"
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

type PausePipelineApiRequest struct {
	ctx          context.Context
	ApiService   DataLakePipelinesApi
	groupId      string
	pipelineName string
}

type PausePipelineApiParams struct {
	GroupId      string
	PipelineName string
}

func (a *DataLakePipelinesApiService) PausePipelineWithParams(ctx context.Context, args *PausePipelineApiParams) PausePipelineApiRequest {
	return PausePipelineApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		pipelineName: args.PipelineName,
	}
}

func (r PausePipelineApiRequest) Execute() (*DataLakeIngestionPipeline, *http.Response, error) {
	return r.ApiService.PausePipelineExecute(r)
}

/*
PausePipeline Pause One Data Lake Pipeline

Pauses ingestion for a Data Lake Pipeline within the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
	@return PausePipelineApiRequest
*/
func (a *DataLakePipelinesApiService) PausePipeline(ctx context.Context, groupId string, pipelineName string) PausePipelineApiRequest {
	return PausePipelineApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		pipelineName: pipelineName,
	}
}

// Execute executes the request
//
//	@return DataLakeIngestionPipeline
func (a *DataLakePipelinesApiService) PausePipelineExecute(r PausePipelineApiRequest) (*DataLakeIngestionPipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataLakeIngestionPipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.PausePipeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/pause"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineName"+"}", url.PathEscape(parameterValueToString(r.pipelineName, "pipelineName")), -1)

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

type ResumePipelineApiRequest struct {
	ctx          context.Context
	ApiService   DataLakePipelinesApi
	groupId      string
	pipelineName string
}

type ResumePipelineApiParams struct {
	GroupId      string
	PipelineName string
}

func (a *DataLakePipelinesApiService) ResumePipelineWithParams(ctx context.Context, args *ResumePipelineApiParams) ResumePipelineApiRequest {
	return ResumePipelineApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		pipelineName: args.PipelineName,
	}
}

func (r ResumePipelineApiRequest) Execute() (*DataLakeIngestionPipeline, *http.Response, error) {
	return r.ApiService.ResumePipelineExecute(r)
}

/*
ResumePipeline Resume One Data Lake Pipeline

Resumes ingestion for a Data Lake Pipeline within the specified project. To use this resource, the requesting API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
	@return ResumePipelineApiRequest
*/
func (a *DataLakePipelinesApiService) ResumePipeline(ctx context.Context, groupId string, pipelineName string) ResumePipelineApiRequest {
	return ResumePipelineApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      groupId,
		pipelineName: pipelineName,
	}
}

// Execute executes the request
//
//	@return DataLakeIngestionPipeline
func (a *DataLakePipelinesApiService) ResumePipelineExecute(r ResumePipelineApiRequest) (*DataLakeIngestionPipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataLakeIngestionPipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.ResumePipeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/resume"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineName"+"}", url.PathEscape(parameterValueToString(r.pipelineName, "pipelineName")), -1)

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

type TriggerSnapshotIngestionApiRequest struct {
	ctx                             context.Context
	ApiService                      DataLakePipelinesApi
	groupId                         string
	pipelineName                    string
	triggerIngestionPipelineRequest *TriggerIngestionPipelineRequest
}

type TriggerSnapshotIngestionApiParams struct {
	GroupId                         string
	PipelineName                    string
	TriggerIngestionPipelineRequest *TriggerIngestionPipelineRequest
}

func (a *DataLakePipelinesApiService) TriggerSnapshotIngestionWithParams(ctx context.Context, args *TriggerSnapshotIngestionApiParams) TriggerSnapshotIngestionApiRequest {
	return TriggerSnapshotIngestionApiRequest{
		ApiService:                      a,
		ctx:                             ctx,
		groupId:                         args.GroupId,
		pipelineName:                    args.PipelineName,
		triggerIngestionPipelineRequest: args.TriggerIngestionPipelineRequest,
	}
}

func (r TriggerSnapshotIngestionApiRequest) Execute() (*IngestionPipelineRun, *http.Response, error) {
	return r.ApiService.TriggerSnapshotIngestionExecute(r)
}

/*
TriggerSnapshotIngestion Trigger on demand snapshot ingestion

Triggers a Data Lake Pipeline ingestion of a specified snapshot.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
	@return TriggerSnapshotIngestionApiRequest
*/
func (a *DataLakePipelinesApiService) TriggerSnapshotIngestion(ctx context.Context, groupId string, pipelineName string, triggerIngestionPipelineRequest *TriggerIngestionPipelineRequest) TriggerSnapshotIngestionApiRequest {
	return TriggerSnapshotIngestionApiRequest{
		ApiService:                      a,
		ctx:                             ctx,
		groupId:                         groupId,
		pipelineName:                    pipelineName,
		triggerIngestionPipelineRequest: triggerIngestionPipelineRequest,
	}
}

// Execute executes the request
//
//	@return IngestionPipelineRun
func (a *DataLakePipelinesApiService) TriggerSnapshotIngestionExecute(r TriggerSnapshotIngestionApiRequest) (*IngestionPipelineRun, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IngestionPipelineRun
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.TriggerSnapshotIngestion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}/trigger"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineName"+"}", url.PathEscape(parameterValueToString(r.pipelineName, "pipelineName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.triggerIngestionPipelineRequest == nil {
		return localVarReturnValue, nil, reportError("triggerIngestionPipelineRequest is required and must be specified")
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
	localVarPostBody = r.triggerIngestionPipelineRequest
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

type UpdatePipelineApiRequest struct {
	ctx                       context.Context
	ApiService                DataLakePipelinesApi
	groupId                   string
	pipelineName              string
	dataLakeIngestionPipeline *DataLakeIngestionPipeline
}

type UpdatePipelineApiParams struct {
	GroupId                   string
	PipelineName              string
	DataLakeIngestionPipeline *DataLakeIngestionPipeline
}

func (a *DataLakePipelinesApiService) UpdatePipelineWithParams(ctx context.Context, args *UpdatePipelineApiParams) UpdatePipelineApiRequest {
	return UpdatePipelineApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   args.GroupId,
		pipelineName:              args.PipelineName,
		dataLakeIngestionPipeline: args.DataLakeIngestionPipeline,
	}
}

func (r UpdatePipelineApiRequest) Execute() (*DataLakeIngestionPipeline, *http.Response, error) {
	return r.ApiService.UpdatePipelineExecute(r)
}

/*
UpdatePipeline Update One Data Lake Pipeline

Updates one Data Lake Pipeline.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param pipelineName Human-readable label that identifies the Data Lake Pipeline.
	@return UpdatePipelineApiRequest
*/
func (a *DataLakePipelinesApiService) UpdatePipeline(ctx context.Context, groupId string, pipelineName string, dataLakeIngestionPipeline *DataLakeIngestionPipeline) UpdatePipelineApiRequest {
	return UpdatePipelineApiRequest{
		ApiService:                a,
		ctx:                       ctx,
		groupId:                   groupId,
		pipelineName:              pipelineName,
		dataLakeIngestionPipeline: dataLakeIngestionPipeline,
	}
}

// Execute executes the request
//
//	@return DataLakeIngestionPipeline
func (a *DataLakePipelinesApiService) UpdatePipelineExecute(r UpdatePipelineApiRequest) (*DataLakeIngestionPipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DataLakeIngestionPipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataLakePipelinesApiService.UpdatePipeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/pipelines/{pipelineName}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineName"+"}", url.PathEscape(parameterValueToString(r.pipelineName, "pipelineName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dataLakeIngestionPipeline == nil {
		return localVarReturnValue, nil, reportError("dataLakeIngestionPipeline is required and must be specified")
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
	localVarPostBody = r.dataLakeIngestionPipeline
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
