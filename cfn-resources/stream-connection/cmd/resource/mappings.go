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
	"go.mongodb.org/atlas-sdk/v20231115014/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

func GetStreamConnectionModel(streamsConn *admin.StreamsConnection, currentModel *Model) *Model {
	model := new(Model)

	if currentModel != nil {
		model = currentModel
	}

	model.ConnectionName = streamsConn.Name
	model.Type = streamsConn.Type
	model.ClusterName = streamsConn.ClusterName
	model.BootstrapServers = streamsConn.BootstrapServers

	model.DbRoleToExecute = NewModelDBRoleToExecute(streamsConn.DbRoleToExecute)

	model.Authentication = NewModelAuthentication(streamsConn.Authentication)

	model.Security = NewModelSecurity(streamsConn.Security)

	if streamsConn.Config != nil {
		model.Config = *streamsConn.Config
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

func NewModelAuthentication(authentication *admin.StreamsKafkaAuthentication) *StreamsKafkaAuthentication {
	if authentication == nil {
		return nil
	}

	return &StreamsKafkaAuthentication{
		Mechanism: authentication.Mechanism,
		Password:  authentication.Password,
		Username:  authentication.Username,
	}
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

	if util.SafeString(streamConnReq.Type) == ClusterConnectionType {
		streamConnReq.ClusterName = model.ClusterName
		streamConnReq.DbRoleToExecute = NewDBRoleToExecute(model.DbRoleToExecute)
	}

	if util.SafeString(streamConnReq.Type) == KafkaConnectionType {
		streamConnReq.BootstrapServers = model.BootstrapServers
		streamConnReq.Security = newStreamsKafkaSecurity(model.Security)
		streamConnReq.Authentication = newStreamsKafkaAuthentication(model.Authentication)

		if model.Config != nil {
			streamConnReq.Config = &model.Config
		}
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
		Mechanism: authenticationModel.Mechanism,
		Password:  authenticationModel.Password,
		Username:  authenticationModel.Username,
	}
}
