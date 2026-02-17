// Copyright 2025 MongoDB Inc
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

package resource

import (
	"context"
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"
)

func HandleCreate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()

	cluster, resp, err := client.AtlasSDK.ClustersApi.GetCluster(ctx, *model.ProjectId, *model.ClusterName).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe
	}

	if apiResp, ok := cluster.GetMongoDBEmployeeAccessGrantOk(); ok && apiResp != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "MongoDB employee access grant already exists for cluster: " + *model.ClusterName,
			HandlerErrorCode: string(types.HandlerErrorCodeAlreadyExists),
		}
	}

	expirationTime, err := time.Parse(time.RFC3339, *model.ExpirationTime)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Invalid expiration time format. Expected RFC3339 format (e.g., 2025-08-01T12:00:00Z)",
			HandlerErrorCode: "InvalidRequest",
		}
	}

	grantReq := &admin.EmployeeAccessGrant{
		GrantType:      *model.GrantType,
		ExpirationTime: expirationTime,
	}

	_, err = client.AtlasSDK.ClustersApi.GrantMongoEmployeeAccess(ctx, *model.ProjectId, *model.ClusterName, grantReq).Execute()
	if err != nil {
		return progressevent.GetFailedEventByCode("Error granting MongoDB employee access: "+err.Error(), string(types.HandlerErrorCodeInternalFailure))
	}

	return HandleRead(req, client, model)
}

func HandleRead(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()

	cluster, resp, err := client.AtlasSDK.ClustersApi.GetCluster(ctx, *model.ProjectId, *model.ClusterName).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe
	}

	apiResp, ok := cluster.GetMongoDBEmployeeAccessGrantOk()
	if !ok || apiResp == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          constants.ResourceNotFound,
			HandlerErrorCode: "NotFound",
		}
	}

	updateModel(model, apiResp)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.ReadComplete,
		ResourceModel:   model,
	}
}

func HandleUpdate(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()

	cluster, resp, err := client.AtlasSDK.ClustersApi.GetCluster(ctx, *model.ProjectId, *model.ClusterName).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe
	}

	if apiResp, ok := cluster.GetMongoDBEmployeeAccessGrantOk(); !ok || apiResp == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          constants.ResourceNotFound,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
		}
	}

	expirationTime, err := time.Parse(time.RFC3339, *model.ExpirationTime)
	if err != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Invalid expiration time format. Expected RFC3339 format (e.g., 2025-08-01T12:00:00Z)",
			HandlerErrorCode: "InvalidRequest",
		}
	}

	grantReq := &admin.EmployeeAccessGrant{
		GrantType:      *model.GrantType,
		ExpirationTime: expirationTime,
	}

	_, err = client.AtlasSDK.ClustersApi.GrantMongoEmployeeAccess(ctx, *model.ProjectId, *model.ClusterName, grantReq).Execute()
	if err != nil {
		return progressevent.GetFailedEventByCode("Error updating MongoDB employee access grant: "+err.Error(), string(types.HandlerErrorCodeInternalFailure))
	}

	return HandleRead(req, client, model)
}

func HandleDelete(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	ctx := context.Background()

	cluster, resp, err := client.AtlasSDK.ClustersApi.GetCluster(ctx, *model.ProjectId, *model.ClusterName).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe
	}

	if apiResp, ok := cluster.GetMongoDBEmployeeAccessGrantOk(); !ok || apiResp == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          constants.ResourceNotFound,
			HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
		}
	}

	resp, err = client.AtlasSDK.ClustersApi.RevokeMongoEmployeeAccess(ctx, *model.ProjectId, *model.ClusterName).Execute()
	if pe := util.HandleClusterError(err, resp); pe != nil {
		return *pe
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         constants.Complete,
	}
}

func HandleList(req *handler.Request, client *util.MongoDBClient, model *Model) handler.ProgressEvent {
	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          "List operation is not supported",
		HandlerErrorCode: string(types.HandlerErrorCodeNotFound),
	}
}
