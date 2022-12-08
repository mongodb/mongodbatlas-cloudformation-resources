package resource

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
)

var RequiredFields = []string{constants.PubKey, constants.PvtKey, constants.GroupID, constants.PvtKey, constants.EndpointID, constants.Type}
var ListRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.GroupID}

// function to validate inputs to all actions
func validateRequest(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	validationError := validateRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	ctx := context.Background()
	cm := mongodbatlas.PrivateLinkEndpointDataLake{
		Provider:   *currentModel.Provider,
		Type:       *currentModel.Type,
		EndpointID: *currentModel.EndpointId,
		Comment:    aws.StringValue(currentModel.Comment),
	}
	_, resp, err := client.DataLakes.CreatePrivateLinkEndpoint(ctx, *currentModel.GroupId, &cm)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create  Private Link ADL",
		ResourceModel:   currentModel,
	}
	return event, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if currentModel.EndpointId == nil {
		return progressevents.GetFailedEventByResponse("required field missing. Resource not found", &http.Response{
			StatusCode: http.StatusNotFound,
		}), nil
	}
	validationError := validateRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	ctx := context.Background()
	dlEndpoint, resp, err := client.DataLakes.GetPrivateLinkEndpoint(ctx, *currentModel.GroupId, *currentModel.EndpointId)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}

	currentModel.Comment = &dlEndpoint.Comment
	currentModel.Type = &dlEndpoint.Type
	currentModel.Provider = &dlEndpoint.Provider
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Private Link ADL",
		ResourceModel:   currentModel,
	}
	return event, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	validationError := validateRequest(RequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	ctx := context.Background()
	resp, err := client.DataLakes.DeletePrivateLinkEndpoint(ctx, *currentModel.GroupId, *currentModel.EndpointId)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "delete data lake endpoint",
	}
	return event, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	validationError := validateRequest(ListRequiredFields, currentModel)
	if validationError != nil {
		return *validationError, nil
	}
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error creating mongoDB client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}
	ctx := context.Background()
	list, resp, err := client.DataLakes.ListPrivateLinkEndpoint(ctx, *currentModel.GroupId)
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), resp.Response), nil
	}
	models := make([]any, 0, len(list.Results))
	for _, v := range list.Results {
		models = append(models, &Model{
			GroupId:    currentModel.GroupId,
			ApiKeys:    currentModel.ApiKeys,
			Comment:    &v.Comment,
			EndpointId: &v.EndpointID,
			Provider:   &v.Provider,
			Type:       &v.Type,
		})
	}
	event := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "list data lake endpoints",
		ResourceModels:  models,
	}
	return event, nil
}
