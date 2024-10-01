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

import "go.mongodb.org/atlas-sdk/v20240805004/admin"

func NewStreamsTenant(model *Model) *admin.StreamsTenant {
	dataProcessRegion := *model.DataProcessRegion
	streamTenant := &admin.StreamsTenant{
		Name:    model.InstanceName,
		GroupId: model.ProjectId,
		DataProcessRegion: &admin.StreamsDataProcessRegion{
			CloudProvider: *dataProcessRegion.CloudProvider,
			Region:        *dataProcessRegion.Region,
		},
	}
	if streamConfig := model.StreamConfig; streamConfig != nil {
		if tier := streamConfig.Tier; tier != nil {
			streamTenant.StreamConfig = &admin.StreamConfig{
				Tier: streamConfig.Tier,
			}
		}
	}
	return streamTenant
}

func newModelDataRegion(dataProcessRegion *admin.StreamsDataProcessRegion) *StreamsDataProcessRegion {
	return &StreamsDataProcessRegion{
		CloudProvider: &dataProcessRegion.CloudProvider,
		Region:        &dataProcessRegion.Region,
	}
}

func newModelStreamConfig(streamConfig *admin.StreamConfig) *StreamConfig {
	return &StreamConfig{
		Tier: streamConfig.Tier,
	}
}

func newModelDBRoleToExecute(dbRole *admin.DBRoleToExecute) *DBRoleToExecute {
	return &DBRoleToExecute{
		Role: dbRole.Role,
		Type: dbRole.Type,
	}
}

func newModelAuthentication(authentication *admin.StreamsKafkaAuthentication) *StreamsKafkaAuthentication {
	return &StreamsKafkaAuthentication{
		Mechanism: authentication.Mechanism,
		Username:  authentication.Username,
	}
}

func newModelSecurity(security *admin.StreamsKafkaSecurity) *StreamsKafkaSecurity {
	return &StreamsKafkaSecurity{
		BrokerPublicCertificate: security.BrokerPublicCertificate,
		Protocol:                security.Protocol,
	}
}

func NewModelConnections(streamConfig *[]admin.StreamsConnection) []StreamsConnection {
	if streamConfig == nil || len(*streamConfig) == 0 {
		return nil
	}

	connections := make([]StreamsConnection, 0)
	for _, connection := range *streamConfig {
		modelConnection := StreamsConnection{
			Name: connection.Name,
			Type: connection.Type,
		}
		if connection.GetType() == Cluster {
			modelConnection.ClusterName = connection.ClusterName
			modelConnection.DbRoleToExecute = newModelDBRoleToExecute(connection.DbRoleToExecute)
		} else if connection.GetType() == Kafka {
			modelConnection.BootstrapServers = connection.BootstrapServers
			modelConnection.Authentication = newModelAuthentication(connection.Authentication)
			modelConnection.Security = newModelSecurity(connection.Security)
		}
		connections = append(connections, modelConnection)
	}
	return connections
}

func newCFNModelFromStreamInstance(prevModel *Model, streamTenant admin.StreamsTenant) *Model {
	return &Model{
		InstanceName:      streamTenant.Name,
		DataProcessRegion: newModelDataRegion(streamTenant.DataProcessRegion),
		StreamConfig:      newModelStreamConfig(streamTenant.StreamConfig),
		ProjectId:         streamTenant.GroupId,
		Id:                streamTenant.Id,
		Hostnames:         streamTenant.GetHostnames(),
		Profile:           prevModel.Profile,
		Connections:       NewModelConnections(streamTenant.Connections),
	}
}
