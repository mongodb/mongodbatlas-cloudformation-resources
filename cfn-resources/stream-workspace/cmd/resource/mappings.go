// Copyright 2026 MongoDB Inc
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

import "go.mongodb.org/atlas-sdk/v20250312012/admin"

func NewStreamWorkspaceCreateReq(model *Model) *admin.StreamsTenant {
	if model == nil {
		return nil
	}
	dataProcessRegion := *model.DataProcessRegion
	streamTenant := &admin.StreamsTenant{
		Name:    model.WorkspaceName,
		GroupId: model.ProjectId,
		DataProcessRegion: &admin.StreamsDataProcessRegion{
			CloudProvider: *dataProcessRegion.CloudProvider,
			Region:        *dataProcessRegion.Region,
		},
	}
	if streamConfig := model.StreamConfig; streamConfig != nil {
		streamTenant.StreamConfig = &admin.StreamConfig{}
		if tier := streamConfig.Tier; tier != nil {
			streamTenant.StreamConfig.Tier = tier
		}
		if maxTierSize := streamConfig.MaxTierSize; maxTierSize != nil {
			streamTenant.StreamConfig.MaxTierSize = maxTierSize
		}
	}
	return streamTenant
}

func NewStreamWorkspaceUpdateReq(model *Model) *admin.StreamsTenantUpdateRequest {
	if model == nil || model.DataProcessRegion == nil {
		return nil
	}
	dataProcessRegion := *model.DataProcessRegion
	if dataProcessRegion.Region == nil {
		return nil
	}
	// CloudFormation is AWS-only, so CloudProvider is always AWS
	cloudProvider := "AWS"
	return &admin.StreamsTenantUpdateRequest{
		CloudProvider: &cloudProvider,
		Region:        dataProcessRegion.Region,
	}
}

func newModelDataRegion(dataProcessRegion *admin.StreamsDataProcessRegion) *StreamsDataProcessRegion {
	return &StreamsDataProcessRegion{
		CloudProvider: &dataProcessRegion.CloudProvider,
		Region:        &dataProcessRegion.Region,
	}
}

func newModelStreamConfig(streamConfig *admin.StreamConfig) *StreamConfig {
	if streamConfig == nil {
		return nil
	}
	modelConfig := &StreamConfig{}
	if streamConfig.Tier != nil {
		modelConfig.Tier = streamConfig.Tier
	}
	if streamConfig.MaxTierSize != nil {
		modelConfig.MaxTierSize = streamConfig.MaxTierSize
	}
	return modelConfig
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
		if connection.GetType() == Kafka {
			modelConnection.BootstrapServers = connection.BootstrapServers
			modelConnection.Authentication = newModelAuthentication(connection.Authentication)
			modelConnection.Security = newModelSecurity(connection.Security)
		} else if connection.GetType() == Cluster {
			modelConnection.ClusterName = connection.ClusterName
			modelConnection.DbRoleToExecute = newModelDBRoleToExecute(connection.DbRoleToExecute)
		}
		connections = append(connections, modelConnection)
	}
	return connections
}

func GetStreamWorkspaceModel(streamTenant *admin.StreamsTenant, currentModel *Model) *Model {
	model := new(Model)

	if currentModel != nil {
		model = currentModel
	}

	if streamTenant != nil {
		model.WorkspaceName = streamTenant.Name
		model.DataProcessRegion = newModelDataRegion(streamTenant.DataProcessRegion)
		model.StreamConfig = newModelStreamConfig(streamTenant.StreamConfig)
		model.ProjectId = streamTenant.GroupId
		model.Id = streamTenant.Id
		model.Hostnames = streamTenant.GetHostnames()
		model.Connections = NewModelConnections(streamTenant.Connections)
	}

	return model
}
