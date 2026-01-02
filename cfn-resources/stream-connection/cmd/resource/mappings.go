// Copyright 2024 MongoDB Inc
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
	admin20250312010 "go.mongodb.org/atlas-sdk/v20250312010/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

func GetStreamConnectionModel(streamsConn *admin20250312010.StreamsConnection, currentModel *Model) *Model {
	var model *Model

	if currentModel != nil {
		// Use currentModel directly to preserve all fields including primary identifier fields
		// (ProjectId, WorkspaceName/InstanceName, Profile, ConnectionName)
		model = currentModel
		// Ensure WorkspaceName is set for primary identifier (required by CFN)
		// If InstanceName is provided but WorkspaceName is not, set WorkspaceName from InstanceName
		if model.WorkspaceName == nil && model.InstanceName != nil && *model.InstanceName != "" {
			model.WorkspaceName = model.InstanceName
		}
	} else {
		// Create new model only when currentModel is nil (e.g., in List operations)
		model = new(Model)
	}

	model.ConnectionName = streamsConn.Name
	model.Type = streamsConn.Type
	model.ClusterName = streamsConn.ClusterName
	if streamsConn.ClusterGroupId != nil {
		model.ClusterProjectId = streamsConn.ClusterGroupId
	}
	model.BootstrapServers = streamsConn.BootstrapServers
	if streamsConn.Url != nil {
		model.Url = streamsConn.Url
	}

	model.DbRoleToExecute = NewModelDBRoleToExecute(streamsConn.DbRoleToExecute)

	model.Authentication = NewModelAuthentication(streamsConn.Authentication, currentModel)

	model.Security = NewModelSecurity(streamsConn.Security)

	if streamsConn.Config != nil {
		model.Config = *streamsConn.Config
	}

	if streamsConn.Headers != nil {
		model.Headers = *streamsConn.Headers
	}

	// Networking
	if streamsConn.Networking != nil && streamsConn.Networking.Access != nil {
		model.Networking = &Networking{
			Access: &Access{
				Type:         streamsConn.Networking.Access.Type,
				ConnectionId: streamsConn.Networking.Access.ConnectionId,
			},
		}
	}

	// AWS
	if streamsConn.Aws != nil {
		model.Aws = &Aws{
			RoleArn: streamsConn.Aws.RoleArn,
		}
	}

	return model
}

func NewModelDBRoleToExecute(dbRole *admin20250312010.DBRoleToExecute) *DBRoleToExecute {
	if dbRole == nil {
		return nil
	}

	return &DBRoleToExecute{
		Role: dbRole.Role,
		Type: dbRole.Type,
	}
}

func NewModelAuthentication(authentication *admin20250312010.StreamsKafkaAuthentication, currentModel *Model) *StreamsKafkaAuthentication {
	if authentication == nil {
		return nil
	}

	authModel := &StreamsKafkaAuthentication{
		Mechanism:                 authentication.Mechanism,
		Method:                    authentication.Method,
		Username:                  authentication.Username,
		TokenEndpointUrl:          authentication.TokenEndpointUrl,
		ClientId:                  authentication.ClientId,
		Scope:                     authentication.Scope,
		SaslOauthbearerExtensions: authentication.SaslOauthbearerExtensions,
		// Note: Password and ClientSecret are write-only fields and should NOT be set here
		// They are only used during Create/Update operations, never returned in Read responses
	}

	return authModel
}

func NewModelSecurity(security *admin20250312010.StreamsKafkaSecurity) *StreamsKafkaSecurity {
	if security == nil {
		return nil
	}

	return &StreamsKafkaSecurity{
		BrokerPublicCertificate: security.BrokerPublicCertificate,
		Protocol:                security.Protocol,
	}
}

func newStreamConnectionReq(model *Model) *admin20250312010.StreamsConnection {
	streamConnReq := admin20250312010.StreamsConnection{
		Name: model.ConnectionName,
		Type: model.Type,
	}

	typeStr := util.SafeString(streamConnReq.Type)

	// Cluster type specific fields
	if typeStr == ClusterConnectionType {
		streamConnReq.ClusterName = model.ClusterName
		if model.ClusterProjectId != nil {
			streamConnReq.ClusterGroupId = model.ClusterProjectId
		}
		streamConnReq.DbRoleToExecute = NewDBRoleToExecute(model.DbRoleToExecute)
	}

	// Kafka type specific fields
	if typeStr == KafkaConnectionType {
		streamConnReq.BootstrapServers = model.BootstrapServers
		streamConnReq.Security = newStreamsKafkaSecurity(model.Security)
		streamConnReq.Authentication = newStreamsKafkaAuthentication(model.Authentication)

		if model.Config != nil {
			streamConnReq.Config = &model.Config
		}

		// Networking for Kafka
		if model.Networking != nil && model.Networking.Access != nil {
			streamConnReq.Networking = &admin20250312010.StreamsKafkaNetworking{
				Access: &admin20250312010.StreamsKafkaNetworkingAccess{
					Type:         model.Networking.Access.Type,
					ConnectionId: model.Networking.Access.ConnectionId,
				},
			}
		}
	}

	// AWS Lambda type specific fields
	if typeStr == AWSLambdaType {
		if model.Aws != nil {
			streamConnReq.Aws = &admin20250312010.StreamsAWSConnectionConfig{
				RoleArn: model.Aws.RoleArn,
			}
		}
	}

	// HTTPS type specific fields
	if typeStr == HTTPSType {
		streamConnReq.Url = model.Url
		if model.Headers != nil {
			streamConnReq.Headers = &model.Headers
		}
	}

	return &streamConnReq
}

func NewDBRoleToExecute(dbRoleToExecuteModel *DBRoleToExecute) *admin20250312010.DBRoleToExecute {
	if dbRoleToExecuteModel == nil {
		return nil
	}

	return &admin20250312010.DBRoleToExecute{
		Role: dbRoleToExecuteModel.Role,
		Type: dbRoleToExecuteModel.Type,
	}
}

func newStreamsKafkaSecurity(securityModel *StreamsKafkaSecurity) *admin20250312010.StreamsKafkaSecurity {
	if securityModel == nil {
		return nil
	}

	return &admin20250312010.StreamsKafkaSecurity{
		BrokerPublicCertificate: securityModel.BrokerPublicCertificate,
		Protocol:                securityModel.Protocol,
	}
}

func newStreamsKafkaAuthentication(authenticationModel *StreamsKafkaAuthentication) *admin20250312010.StreamsKafkaAuthentication {
	if authenticationModel == nil {
		return nil
	}

	return &admin20250312010.StreamsKafkaAuthentication{
		Mechanism:                 authenticationModel.Mechanism,
		Method:                    authenticationModel.Method,
		Username:                  authenticationModel.Username,
		Password:                  authenticationModel.Password,
		TokenEndpointUrl:          authenticationModel.TokenEndpointUrl,
		ClientId:                  authenticationModel.ClientId,
		ClientSecret:              authenticationModel.ClientSecret,
		Scope:                     authenticationModel.Scope,
		SaslOauthbearerExtensions: authenticationModel.SaslOauthbearerExtensions,
	}
}
