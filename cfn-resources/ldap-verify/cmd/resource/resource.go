package resource

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.GroupID, "BindUsername", "BindPassword", "Hostname", "Port", constants.PvtKey, constants.PubKey}
var ReadRequiredFields = []string{constants.GroupID, "RequestId", constants.PvtKey, constants.PubKey}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{constants.GroupID, constants.PvtKey, constants.PubKey}
var ListRequiredFields []string

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-LDAPVerify")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	if req.CallbackContext != nil {
		return validateProgress(client, currentModel, req), nil
	}

	ldapReq := currentModel.GetAtlasModel()

	LDAPConfigResponse, res, err := client.LDAPConfigurations.Verify(context.Background(), *currentModel.GroupId, ldapReq)
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.CompleteByResponse(*LDAPConfigResponse)

	/*
		Uncomment this once we have a way to test with an LDAP server

		return handler.ProgressEvent{

			OperationStatus: handler.InProgress,
			Message:         "Create in progress",
			ResourceModel:   currentModel,
			CallbackContext: map[string]interface{}{
				"RequestId": currentModel.RequestId,
			},
		}, nil
	*/
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create successfully",
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Read() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	var res *mongodbatlas.Response

	LDAPConfigResponse, res, err := client.LDAPConfigurations.GetStatus(context.Background(), *currentModel.GroupId, *currentModel.RequestId)
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	currentModel.CompleteByResponse(*LDAPConfigResponse)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Validation
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progress_events.GetFailedEventByCode(err.Error(), cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	_, res, err := client.LDAPConfigurations.GetStatus(context.Background(), *currentModel.GroupId, *currentModel.RequestId)
	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error deleting resource : %s", err.Error()),
			res.Response), nil
	}

	ldapReq := &mongodbatlas.LDAP{
		Hostname:     "-",
		Port:         1111,
		BindPassword: "-",
		BindUsername: "-",
	}

	_, res, err = client.LDAPConfigurations.Verify(context.Background(), *currentModel.GroupId, ldapReq)
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete completed successfully",
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func (m *Model) GetAtlasModel() *mongodbatlas.LDAP {
	ldap := &mongodbatlas.LDAP{
		Hostname:     *m.Hostname,
		Port:         *m.Port,
		BindPassword: *m.BindPassword,
		BindUsername: *m.BindUsername,
	}

	if m.AuthzQueryTemplate != nil {
		ldap.AuthzQueryTemplate = *m.AuthzQueryTemplate
	}

	if m.CaCertificate != nil {
		ldap.CaCertificate = *m.CaCertificate
	}

	return ldap
}

func (m *Model) CompleteByResponse(resp mongodbatlas.LDAPConfiguration) {
	m.RequestId = &resp.RequestID

	mapping := make([]Validation, len(resp.Validations))

	for i := range resp.Validations {
		validation := Validation{
			Status:         &resp.Validations[i].Status,
			ValidationType: &resp.Validations[i].ValidationType,
		}
		mapping[i] = validation
	}

	m.Validations = mapping
}

func validateProgress(client *mongodbatlas.Client, model *Model, req handler.Request) handler.ProgressEvent {
	requestId := req.CallbackContext["RequestId"].(string)

	LDAPConfigResponse, res, err := client.LDAPConfigurations.GetStatus(context.Background(), *model.GroupId, requestId)
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response)
	}

	switch LDAPConfigResponse.Status {
	case "PENDING":
		{
			return handler.ProgressEvent{
				OperationStatus: handler.InProgress,
				Message:         "Create in progress",
				ResourceModel:   model,
				CallbackContext: map[string]interface{}{
					"RequestId": requestId,
				},
			}
		}
	case "SUCCESS":
		{
			model.CompleteByResponse(*LDAPConfigResponse)
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Create successfully",
				ResourceModel:   model,
			}
		}
	default:
		{
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         getFailedMessage(*LDAPConfigResponse),
			}
		}
	}
}

func getFailedMessage(configuration mongodbatlas.LDAPConfiguration) string {
	for _, i := range configuration.Validations {
		if i.Status == "FAIL" {
			return fmt.Sprintf("Faild in validation %s", i.ValidationType)
		}
	}

	return ""
}
