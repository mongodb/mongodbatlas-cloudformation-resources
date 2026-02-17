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
	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

func GetStreamConnectionModel(streamsConn *admin.StreamsConnection, currentModel *Model) *Model {
	model := new(Model)

	if currentModel != nil {
		model.ProjectId = currentModel.ProjectId
		model.Profile = currentModel.Profile
		model.WorkspaceName = currentModel.WorkspaceName
		model.InstanceName = currentModel.InstanceName
	}

	model.ConnectionName = streamsConn.Name
	model.Type = streamsConn.Type
	model.ClusterName = streamsConn.ClusterName
	model.ClusterProjectId = streamsConn.ClusterGroupId
	model.BootstrapServers = streamsConn.BootstrapServers
	model.Url = streamsConn.Url

	model.DbRoleToExecute = NewModelDBRoleToExecute(streamsConn.DbRoleToExecute)

	model.Authentication = NewModelAuthentication(streamsConn.Authentication, currentModel)

	model.Security = NewModelSecurity(streamsConn.Security)

	if streamsConn.Config != nil {
		model.Config = *streamsConn.Config
	}

	if streamsConn.Headers != nil {
		model.Headers = *streamsConn.Headers
	}

	if streamsConn.Networking != nil && streamsConn.Networking.Access != nil {
		model.Networking = &Networking{
			Access: &Access{
				Type:         streamsConn.Networking.Access.Type,
				ConnectionId: streamsConn.Networking.Access.ConnectionId,
				Name:         streamsConn.Networking.Access.Name,
				TgwRouteId:   streamsConn.Networking.Access.TgwRouteId,
			},
		}
	}

	if streamsConn.Aws != nil {
		model.Aws = &Aws{
			RoleArn: streamsConn.Aws.RoleArn,
		}
	}

	// Schema Registry fields
	if streamsConn.Provider != nil {
		model.Provider = streamsConn.Provider
	} else if currentModel != nil {
		model.Provider = currentModel.Provider
	}

	if streamsConn.SchemaRegistryAuthentication != nil {
		model.SchemaRegistryAuthentication = &SchemaRegistryAuthentication{
			Type:     &streamsConn.SchemaRegistryAuthentication.Type,
			Username: streamsConn.SchemaRegistryAuthentication.Username,
		}
	} else if currentModel != nil && currentModel.SchemaRegistryAuthentication != nil {
		model.SchemaRegistryAuthentication = &SchemaRegistryAuthentication{
			Type:     currentModel.SchemaRegistryAuthentication.Type,
			Username: currentModel.SchemaRegistryAuthentication.Username,
		}
	}

	if streamsConn.SchemaRegistryUrls != nil {
		model.SchemaRegistryUrls = *streamsConn.SchemaRegistryUrls
	} else if currentModel != nil {
		model.SchemaRegistryUrls = currentModel.SchemaRegistryUrls
	}

	return model
}

func NewModelDBRoleToExecute(dbRole *admin.DBRoleToExecute) *DBRoleToExecute {
	if dbRole == nil {
		return nil
	}

	return &DBRoleToExecute{
		Role: dbRole.Role,
		Type: dbRole.Type,
	}
}

func NewModelAuthentication(authentication *admin.StreamsKafkaAuthentication, currentModel *Model) *StreamsKafkaAuthentication {
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
	}

	return authModel
}

func NewModelSecurity(security *admin.StreamsKafkaSecurity) *StreamsKafkaSecurity {
	if security == nil {
		return nil
	}

	return &StreamsKafkaSecurity{
		BrokerPublicCertificate: security.BrokerPublicCertificate,
		Protocol:                security.Protocol,
	}
}

func newStreamConnectionReq(model *Model) *admin.StreamsConnection {
	streamConnReq := admin.StreamsConnection{
		Name: model.ConnectionName,
		Type: model.Type,
	}

	typeStr := util.SafeString(streamConnReq.Type)

	if typeStr == ClusterConnectionType {
		streamConnReq.ClusterName = model.ClusterName
		if model.ClusterProjectId != nil {
			streamConnReq.ClusterGroupId = model.ClusterProjectId
		}
		streamConnReq.DbRoleToExecute = NewDBRoleToExecute(model.DbRoleToExecute)
	}

	if typeStr == KafkaConnectionType {
		streamConnReq.BootstrapServers = model.BootstrapServers
		streamConnReq.Security = newStreamsKafkaSecurity(model.Security)
		streamConnReq.Authentication = newStreamsKafkaAuthentication(model.Authentication)

		if model.Config != nil {
			streamConnReq.Config = &model.Config
		}

		if model.Networking != nil && model.Networking.Access != nil {
			streamConnReq.Networking = &admin.StreamsKafkaNetworking{
				Access: &admin.StreamsKafkaNetworkingAccess{
					Type:         model.Networking.Access.Type,
					ConnectionId: model.Networking.Access.ConnectionId,
					Name:         model.Networking.Access.Name,
					TgwRouteId:   model.Networking.Access.TgwRouteId,
				},
			}
		}
	}

	if typeStr == AWSLambdaType {
		if model.Aws != nil {
			streamConnReq.Aws = &admin.StreamsAWSConnectionConfig{
				RoleArn:    model.Aws.RoleArn,
				TestBucket: model.Aws.TestBucket,
			}
		}
	}

	if typeStr == HTTPSType {
		streamConnReq.Url = model.Url
		if model.Headers != nil {
			streamConnReq.Headers = &model.Headers
		}
	}

	// Schema Registry fields
	streamConnReq.Provider = model.Provider
	if model.SchemaRegistryAuthentication != nil {
		streamConnReq.SchemaRegistryAuthentication = &admin.SchemaRegistryAuthentication{
			Type:     util.SafeString(model.SchemaRegistryAuthentication.Type),
			Username: model.SchemaRegistryAuthentication.Username,
			Password: model.SchemaRegistryAuthentication.Password,
		}
	}
	if len(model.SchemaRegistryUrls) > 0 {
		streamConnReq.SchemaRegistryUrls = &model.SchemaRegistryUrls
	}

	return &streamConnReq
}

func NewDBRoleToExecute(dbRoleToExecuteModel *DBRoleToExecute) *admin.DBRoleToExecute {
	if dbRoleToExecuteModel == nil {
		return nil
	}

	return &admin.DBRoleToExecute{
		Role: dbRoleToExecuteModel.Role,
		Type: dbRoleToExecuteModel.Type,
	}
}

func newStreamsKafkaSecurity(securityModel *StreamsKafkaSecurity) *admin.StreamsKafkaSecurity {
	if securityModel == nil {
		return nil
	}

	return &admin.StreamsKafkaSecurity{
		BrokerPublicCertificate: securityModel.BrokerPublicCertificate,
		Protocol:                securityModel.Protocol,
	}
}

func newStreamsKafkaAuthentication(authenticationModel *StreamsKafkaAuthentication) *admin.StreamsKafkaAuthentication {
	if authenticationModel == nil {
		return nil
	}

	return &admin.StreamsKafkaAuthentication{
		Mechanism:                 authenticationModel.Mechanism,
		Method:                    authenticationModel.Method,
		Username:                  authenticationModel.Username,
		Password:                  authenticationModel.Password,
		TokenEndpointUrl:          authenticationModel.TokenEndpointUrl,
		ClientId:                  authenticationModel.ClientId,
		ClientSecret:              authenticationModel.ClientSecret,
		Scope:                     authenticationModel.Scope,
		SaslOauthbearerExtensions: authenticationModel.SaslOauthbearerExtensions,
		SslCertificate:            authenticationModel.SslCertificate,
		SslKey:                    authenticationModel.SslKey,
		SslKeyPassword:            authenticationModel.SslKeyPassword,
	}
}
