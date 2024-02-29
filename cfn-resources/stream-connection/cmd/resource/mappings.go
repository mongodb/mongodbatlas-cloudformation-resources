package resource

import (
	"context"
	"encoding/json"

	"go.mongodb.org/atlas-sdk/v20231115007/admin"

	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
)

func GetStreamConnectionModel(streamsConn *admin.StreamsConnection, currentModel *Model) *Model {
	model := &Model{}

	if currentModel != nil {
		model = currentModel
	}

	jsonBytes, _ := json.MarshalIndent(streamsConn.DbRoleToExecute, "", "\t")
	log.Debugf("\n ------- 1. in newStreamConnectionModel DBRoleToExecute from API:  %s", string(jsonBytes))

	model.ConnectionName = streamsConn.Name
	model.Type = streamsConn.Type
	model.ClusterName = streamsConn.ClusterName
	model.BootstrapServers = streamsConn.BootstrapServers

	jsonBytes, _ = json.MarshalIndent(streamsConn.DbRoleToExecute, "", "\t")
	log.Debugf("\n ------- 2. in newStreamConnectionModel DBRoleToExecute from API:  %s", string(jsonBytes))

	jsonBytes, _ = json.MarshalIndent(streamsConn.DbRoleToExecute, "", "\t")
	log.Debugf("\n ------- 3. in newStreamConnectionModel DBRoleToExecute from API:  %s", string(jsonBytes))

	model.DbRoleToExecute = NewModelDBRoleToExecute(streamsConn.DbRoleToExecute)
	jsonBytes, _ = json.MarshalIndent(streamsConn.DbRoleToExecute, "", "\t")
	log.Debugf("\n ------- 4. in newStreamConnectionModel DBRoleToExecute from API:  %s", string(jsonBytes))

	model.Authentication = NewModelAuthentication(streamsConn.Authentication)
	jsonBytes, _ = json.MarshalIndent(streamsConn.Authentication, "", "\t")
	log.Debugf("\n ------- 5. in newStreamConnectionModel Authentication from API:  %s", string(jsonBytes))

	model.Security = NewModelSecurity(streamsConn.Security)
	jsonBytes, _ = json.MarshalIndent(streamsConn.Security, "", "\t")
	log.Debugf("\n ------- 6. in newStreamConnectionModel Security from API:  %s", string(jsonBytes))

	if streamsConn.Config != nil {
		model.Config = *streamsConn.Config
	}

	return model
}

func NewModelDBRoleToExecute(dbRole *admin.DBRoleToExecute) *DBRoleToExecute {
	if dbRole != nil {
		return &DBRoleToExecute{
			Role: dbRole.Role,
			Type: dbRole.Type,
		}
	}
	return nil
}

func NewModelAuthentication(authentication *admin.StreamsKafkaAuthentication) *StreamsKafkaAuthentication {
	if authentication != nil {
		return &StreamsKafkaAuthentication{
			Mechanism: authentication.Mechanism,
			Password:  authentication.Password,
			Username:  authentication.Username,
		}
	}
	return nil
}

func NewModelSecurity(security *admin.StreamsKafkaSecurity) *StreamsKafkaSecurity {
	if security != nil {
		return &StreamsKafkaSecurity{
			BrokerPublicCertificate: security.BrokerPublicCertificate,
			Protocol:                security.Protocol,
		}
	}
	return nil
}

func NewStreamConnectionReq(ctx context.Context, model *Model) *admin.StreamsConnection {
	streamConnReq := admin.StreamsConnection{
		Name: model.ConnectionName,
		Type: model.Type,
	}

	if *streamConnReq.Type == ClusterConnectionType {
		streamConnReq.ClusterName = model.ClusterName
		streamConnReq.DbRoleToExecute = NewDBRoleToExecute(model.DbRoleToExecute)
	}

	if *streamConnReq.Type == KafkaConnectionType {
		streamConnReq.BootstrapServers = model.BootstrapServers
		streamConnReq.Security = NewStreamsKafkaSecurity(model.Security)
		streamConnReq.Authentication = NewStreamsKafkaAuthentication(model.Authentication)

		if model.Config != nil {
			streamConnReq.Config = &model.Config
		}
	}

	return &streamConnReq
}

func NewDBRoleToExecute(dbRoleToExecuteModel *DBRoleToExecute) *admin.DBRoleToExecute {
	if dbRoleToExecuteModel != nil {
		return &admin.DBRoleToExecute{
			Role: dbRoleToExecuteModel.Role,
			Type: dbRoleToExecuteModel.Type,
		}
	}
	return nil
}

func NewStreamsKafkaSecurity(securityModel *StreamsKafkaSecurity) *admin.StreamsKafkaSecurity {
	if securityModel != nil {
		return &admin.StreamsKafkaSecurity{
			BrokerPublicCertificate: securityModel.BrokerPublicCertificate,
			Protocol:                securityModel.Protocol,
		}
	}
	return nil
}

func NewStreamsKafkaAuthentication(authenticationModel *StreamsKafkaAuthentication) *admin.StreamsKafkaAuthentication {
	if authenticationModel != nil {
		return &admin.StreamsKafkaAuthentication{
			Mechanism: authenticationModel.Mechanism,
			Password:  authenticationModel.Password,
			Username:  authenticationModel.Username,
		}
	}
	return nil
}
