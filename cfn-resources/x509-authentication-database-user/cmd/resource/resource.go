package resource

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cast"

	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/openlyinc/pointy"
	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.UserId}
var ReadRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID}

// ValidateRequest function to validate the request
func ValidateRequest(requiredFields []string, currentModel *Model) (handler.ProgressEvent, *mongodbatlas.Client, error) {
	// Validate required fields are empty or nil
	if modelValidation := validateModel(requiredFields, currentModel); modelValidation != nil {
		return *modelValidation, nil, errors.New("required field not found")
	}
	// Validate API Keys
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = logger.Warnf(constants.ErrorCreateMongoClient, err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil, err
	}
	return handler.ProgressEvent{}, client, nil
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("Create - creating MongoDB X509 Authentication for DB User:%+v", currentModel)

	// Validate required fields in the request
	event, client, err := ValidateRequest(CreateRequiredFields, currentModel)
	if err != nil {
		return event, err
	}
	_, _ = logger.Debug("Creating MongoDB X509 Authentication for DB User starts")

	// progress callback setup
	if _, ok := req.CallbackContext["status"]; ok {
		sid := req.CallbackContext["ProjectId"].(string)
		currentModel.ProjectId = &sid
		return validateProgress(client, currentModel)
	}

	// Create Atlas API Request Object
	projectID := *currentModel.ProjectId
	userName := *currentModel.UserName
	expirationMonths := *currentModel.MonthsUntilExpiration
	// create new user certificate
	if expirationMonths > 0 {
		_, _ = logger.Debug("Creating User Certificate")
		res, _, err := client.X509AuthDBUsers.CreateUserCertificate(context.Background(), projectID, userName, expirationMonths)
		if err != nil {
			_, _ = logger.Warnf("error creating MongoDB X509 Authentication for DB User(%s) in the project(%s): %v", userName, projectID, err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
		if res != nil {
			currentModel.CustomerX509 = &CustomerX509{
				Cas: pointy.String(res.Certificate),
			}
		}
	} else { // save customer provided certificate
		_, _ = logger.Debug("Save Custom Certificate DB User starts")
		customerX509Cas := *currentModel.CustomerX509.Cas
		_, _, err := client.X509AuthDBUsers.SaveConfiguration(context.Background(), projectID, &mongodbatlas.CustomerX509{Cas: customerX509Cas})
		if err != nil {
			_, _ = logger.Warnf("error creating Customer X509 Authentication in the project(%s): %s", projectID, err)
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
		}
	}
	// track progress
	event = handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Created  Certificate  for DB User ",
		ResourceModel:   currentModel,
	}
	return event, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("Read - X509 certificates for Request() :%+v", currentModel)
	// Validate required fields in the request
	event, client, err := ValidateRequest(ReadRequiredFields, currentModel)
	if err != nil {
		return event, err
	}

	_, _ = logger.Debug("Read - X509 Certificates starts ")
	// check if certificate is enabled
	if !isEnabled(client, currentModel) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "config is not available",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	readModel, err := ReadUserX509Certificate(client, currentModel)
	if err != nil {
		_, _ = logger.Warnf("error reading MongoDB X509 Authentication for DB Users(%s) in the project(%s): %s", *currentModel.UserName, *currentModel.ProjectId, err)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   readModel,
	}, nil
}

// ReadUserX509Certificate Read handles the Read event from the Cloudformation service.
func ReadUserX509Certificate(client *mongodbatlas.Client, currentModel *Model) (*Model, error) {
	setup() // logger setup

	_, _ = logger.Debugf("Read - X509 certificates  function starts %v", currentModel)

	// Create Atlas API Request Object
	projectID := *currentModel.ProjectId
	userName := *currentModel.UserName
	params := &mongodbatlas.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}
	// API call to read all user certificates
	certificates, resp, err := client.X509AuthDBUsers.GetUserCertificates(context.Background(), projectID, userName, params)
	if err != nil {
		_, _ = logger.Warnf("create - error: %+v", err)
		return nil, fmt.Errorf("error reading all MongoDB X509 certificates for DB Users(%s) in the project(%s): %s", *currentModel.UserName, projectID, err)
	}
	currentModel.Links = flattenLinks(resp.Links)
	flattenCertificates(certificates, currentModel)

	// API call to get currently configured certificate
	certificate, _, err := client.X509AuthDBUsers.GetCurrentX509Conf(context.Background(), projectID)
	_, _ = logger.Debugf("Read - X509 Certificates starts : %+v ", certificate)
	if err != nil {
		_, _ = logger.Warnf("error reading MongoDB X509 Authentication for DB Users(%s) in the project(%s): %s", *currentModel.UserName, projectID, err)
		return nil, fmt.Errorf("error reading MongoDB X509 Authentication for DB Users(%s) in the project(%s): %s", *currentModel.UserName, projectID, err)
	} else if certificate != nil {
		currentModel.CustomerX509 = &CustomerX509{
			Cas: &certificate.Cas,
		}
	}
	return currentModel, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("not implemented")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup() // logger setup

	_, _ = logger.Debugf("Delete - X509 Certificates  for Request() currentModel:%+v", currentModel)

	// Validate required fields in the request
	event, client, err := ValidateRequest(CreateRequiredFields, currentModel)
	if err != nil {
		return event, err
	}

	// Create Atlas API Request Object
	if !isEnabled(client, currentModel) {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "config is not available",
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	_, _ = logger.Debug("Delete - X509 Certificates  starts ")
	// API call
	projectID := *currentModel.ProjectId
	_, err = client.X509AuthDBUsers.DisableCustomerX509(context.Background(), projectID)
	if err != nil {
		_, _ = logger.Warnf("error deleting Customer X509 Authentication in the project(%s): %s", projectID, *currentModel.UserName)
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Unable to Delete",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInternalFailure,
		}, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// Not implemented, return an empty handler.ProgressEvent
	// and an error
	return handler.ProgressEvent{}, errors.New("not implemented")
}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-x509-authentication-db-user")
}
func flattenLinks(linksResult []*mongodbatlas.Link) []Links {
	if linksResult != nil {
		links := make([]Links, 0)
		for _, link := range linksResult {
			var lin Links
			lin.Href = &link.Href
			lin.Rel = &link.Rel
			links = append(links, lin)
		}
		return links
	}
	return nil
}
func flattenCertificates(userCertificates []mongodbatlas.UserCertificate, currentModel *Model) *Model {
	if userCertificates != nil {
		certificates := make([]Certificate, 0)
		for _, v := range userCertificates {
			role := Certificate{
				Id:        pointy.String(cast.ToString(v.ID)),
				CreatedAt: pointy.String(v.CreatedAt),
				GroupId:   pointy.String(v.GroupID),
				NotAfter:  pointy.String(v.NotAfter),
				Subject:   pointy.String(v.Subject),
			}
			certificates = append(certificates, role)
		}
		currentModel.Results = certificates
		currentModel.TotalCount = pointy.Int(len(userCertificates))
	}
	return currentModel
}

// function to track snapshot creation status
func validateProgress(client *mongodbatlas.Client, currentModel *Model) (handler.ProgressEvent, error) {
	projectID := *currentModel.ProjectId
	isReady, state, err := certificateIsReady(client, projectID)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	if !isReady {
		p := handler.NewProgressEvent()
		p.ResourceModel = currentModel
		p.OperationStatus = handler.InProgress
		p.CallbackDelaySeconds = 10
		p.Message = "Pending"
		p.CallbackContext = map[string]interface{}{
			"status":    state,
			"ProjectId": *currentModel.ProjectId,
		}
		return p, nil
	}

	p := handler.NewProgressEvent()
	p.ResourceModel = currentModel
	p.OperationStatus = handler.Success
	p.Message = "Complete"
	return p, nil
}

// Read handles the Read event from the Cloudformation service.
func isEnabled(client *mongodbatlas.Client, currentModel *Model) bool {
	setup() // logger setup

	_, _ = logger.Debugf("Read - X509 certificates for Request() :%+v", currentModel)
	projectID := *currentModel.ProjectId

	certificate, _, err := client.X509AuthDBUsers.GetCurrentX509Conf(context.Background(), projectID)
	_, _ = logger.Debugf("Read - X509 Certificates starts : %+v ", certificate)
	if err != nil {
		_, _ = logger.Warnf("error reading MongoDB X509 Authentication for DB Users(%s) in the project(%s): %s", *currentModel.UserName, projectID, err)
		return false
	} else if certificate != nil && certificate.Cas != "" {
		return true
	}

	return false
}

// function to check if snapshot already exist in atlas
func certificateIsReady(client *mongodbatlas.Client, projectID string) (isExist bool, groupID string, errMsg error) {
	certificate, resp, err := client.X509AuthDBUsers.GetCurrentX509Conf(context.Background(), projectID)
	if err != nil {
		if certificate == nil && resp == nil {
			return false, "", err
		}
		if resp != nil && resp.StatusCode == 404 {
			return true, "deleted", nil
		}
		return false, "", err
	}
	return resp.StatusCode == constants.Success, "completed", nil
}
