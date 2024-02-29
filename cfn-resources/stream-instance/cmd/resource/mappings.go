package resource

import "go.mongodb.org/atlas-sdk/v20231115007/admin"

func NewStreamsTenant(model *Model) *admin.StreamsTenant {
	dataProcessRegion := *model.DataProcessRegion
	streamConfig := model.StreamConfig
	streamTenant := &admin.StreamsTenant{
		Name:    model.InstanceName,
		GroupId: model.ProjectId,
		DataProcessRegion: &admin.StreamsDataProcessRegion{
			CloudProvider: *dataProcessRegion.CloudProvider,
			Region:        *dataProcessRegion.Region,
		},
	}
	if streamConfig != nil {
		streamTenant.StreamConfig = &admin.StreamConfig{
			Tier: streamConfig.Tier,
		}
	}
	return streamTenant
}

func NewModelDataRegion(dataProcessRegion *admin.StreamsDataProcessRegion) *StreamsDataProcessRegion {
	return &StreamsDataProcessRegion{
		CloudProvider: &dataProcessRegion.CloudProvider,
		Region:        &dataProcessRegion.Region,
	}
}

func NewModelStreamConfig(streamConfig *admin.StreamConfig) *StreamConfig {
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
		Password:  authentication.Password,
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
