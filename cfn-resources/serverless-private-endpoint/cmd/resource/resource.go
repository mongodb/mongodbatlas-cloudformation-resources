package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/serverless-private-endpoint/cmd/resource/enums"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.InstanceName}
var ReadRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ID}
var DeleteRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ID}
var UpdateRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ID}
var ListRequiredFields []string

const (
	id                     = "id"
	callbackDelayInSeconds = 5
)

func setup() {
	util.SetupLogger("mongodb-atlas-serverless-private-endpoint")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if isRequestInProgress(req) {
		return validateCompletion(req, currentModel, client, enums.Reserved, constants.CREATE), nil
	}

	serverlessPrivateEndpointInput := admin.ServerlessTenantCreateRequest{Comment: currentModel.Comment}

	createServerlessPrivateEndpointRequest := client.AtlasV2.ServerlessPrivateEndpointsApi.CreateServerlessPrivateEndpoint(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, &serverlessPrivateEndpointInput)
	serverlessPrivateEndpoint, response, err := createServerlessPrivateEndpointRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("error creating Serverless Private Endpoint %s",
			err.Error()), response), nil
	}

	if serverlessPrivateEndpoint == nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error while trying to make api call, CreateServerlessPrivateEndpoint returned status %d, and the response is NULL", response.StatusCode),
			cloudformation.HandlerErrorCodeInternalFailure), nil
	}
	currentModel.completeWithAtlasModel(*serverlessPrivateEndpoint)

	callbackContext := map[string]interface{}{
		id: currentModel.Id,
	}

	return progressevents.GetInProgressProgressEvent("Create in progress", callbackContext, currentModel, callbackDelayInSeconds), nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	createServerlessPrivateEndpointRequest := client.AtlasV2.ServerlessPrivateEndpointsApi.GetServerlessPrivateEndpoint(context.Background(),
		*currentModel.ProjectId, *currentModel.InstanceName, *currentModel.Id)
	serverlessPrivateEndpoint, response, err := createServerlessPrivateEndpointRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		if isTenantPrivateEndpointNotFound(response) {
			return progressevents.GetFailedEventByCode(fmt.Sprintf("error getting Serverless Private Endpoint %s", err.Error()), cloudformation.HandlerErrorCodeNotFound), nil
		}
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("error getting Serverless Private Endpoint %s",
			err.Error()), response), nil
	}

	if serverlessPrivateEndpoint == nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error while trying to make api call, CreateServerlessPrivateEndpoint returned status %d, and the response is NULL", response.StatusCode),
			cloudformation.HandlerErrorCodeInternalFailure), nil
	}

	currentModel.completeWithAtlasModel(*serverlessPrivateEndpoint)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   currentModel}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	if currentModel.ProviderName == nil {
		currentModel.ProviderName = aws.String(constants.AWS)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	if isRequestInProgress(req) {
		return validateCompletion(req, currentModel, client, enums.Available, constants.UPDATE), nil
	}

	serverlessPrivateEndpointInput := admin.ServerlessTenantEndpointUpdate{
		Comment:                  currentModel.Comment,
		ProviderName:             *currentModel.ProviderName,
		CloudProviderEndpointId:  currentModel.CloudProviderEndpointId,
		PrivateEndpointIpAddress: currentModel.PrivateEndpointIpAddress,
	}

	createServerlessPrivateEndpointRequest := client.AtlasV2.ServerlessPrivateEndpointsApi.UpdateServerlessPrivateEndpoint(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, *currentModel.Id, &serverlessPrivateEndpointInput)
	serverlessPrivateEndpoint, response, err := createServerlessPrivateEndpointRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		if isTenantPrivateEndpointNotFound(response) {
			return progressevents.GetFailedEventByCode(fmt.Sprintf("error updating Serverless Private Endpoint %s", err.Error()), cloudformation.HandlerErrorCodeNotFound), nil
		}
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("error updating Serverless Private Endpoint %s",
			err.Error()), response), nil
	}

	if serverlessPrivateEndpoint == nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error while trying to make api call, CreateServerlessPrivateEndpoint returned status %d, and the response is NULL", response.StatusCode),
			cloudformation.HandlerErrorCodeInternalFailure), nil
	}
	currentModel.completeWithAtlasModel(*serverlessPrivateEndpoint)

	if currentModel.PrivateEndpointIpAddress == nil && currentModel.CloudProviderEndpointId == nil {
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Create Completed",
			ResourceModel:   currentModel}, nil
	}

	callbackContext := map[string]interface{}{
		id: currentModel.Id,
	}

	return progressevents.GetInProgressProgressEvent("Update in progress", callbackContext, currentModel, callbackDelayInSeconds), nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	// Check if the model is valid for deletion
	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	// Make the API call to delete the serverless private endpoint
	deleteServerlessPrivateEndpointRequest := client.AtlasV2.ServerlessPrivateEndpointsApi.DeleteServerlessPrivateEndpoint(context.Background(),
		*currentModel.ProjectId, *currentModel.InstanceName, *currentModel.Id)
	_, response, err := deleteServerlessPrivateEndpointRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		if isTenantPrivateEndpointNotFound(response) {
			if isRequestInProgress(req) {
				return handler.ProgressEvent{
					OperationStatus: handler.Success,
					Message:         fmt.Sprintf("%s Completed", string(constants.DELETE)),
				}, nil
			}
			return progressevents.GetFailedEventByCode(fmt.Sprintf("error deleting Serverless Private Endpoint %s", err.Error()), cloudformation.HandlerErrorCodeNotFound), nil
		}
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("error deleting Serverless Private Endpoint %s", err.Error()), response), nil
	}

	return progressevents.GetInProgressProgressEvent("Create in progress", getCallbackContext(*currentModel.Id), currentModel, callbackDelayInSeconds), nil
}

func isTenantPrivateEndpointNotFound(response *http.Response) bool {
	type ErrorResponse struct {
		ErrorCode string `json:"errorCode"`
	}

	var errResponse ErrorResponse
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&errResponse)
	if err != nil {
		return false
	}

	return errResponse.ErrorCode == "TENANT_PRIVATE_ENDPOINT_NOT_FOUND"
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	listServerlessPrivateEndpointRequest := client.AtlasV2.ServerlessPrivateEndpointsApi.ListServerlessPrivateEndpoints(context.Background(),
		*currentModel.ProjectId, *currentModel.InstanceName)
	serverlessPrivateEndpoints, response, err := listServerlessPrivateEndpointRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("error getting Serverless Private Endpoint %s",
			err.Error()), response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Completed",
		ResourceModels:  ConvertListToModelList(serverlessPrivateEndpoints, currentModel.Profile)}, nil
}

func ConvertListToModelList(endpoints []admin.ServerlessTenantEndpoint, profile *string) []interface{} {
	var models []interface{}

	for _, endpoint := range endpoints {
		model := Model{
			Id:                       endpoint.Id,
			Comment:                  endpoint.Comment,
			EndpointServiceName:      endpoint.EndpointServiceName,
			ErrorMessage:             endpoint.ErrorMessage,
			ProviderName:             endpoint.ProviderName,
			Status:                   endpoint.Status,
			CloudProviderEndpointId:  endpoint.CloudProviderEndpointId,
			PrivateEndpointIpAddress: endpoint.PrivateEndpointIpAddress,
			Profile:                  profile,
		}
		models = append(models, model)
	}

	return models
}

func isRequestInProgress(req handler.Request) bool {
	_, idExists := req.CallbackContext[id]
	return idExists
}

func validateCompletion(req handler.Request, currentModel *Model, client *util.MongoDBClient, targetStatus enums.Status, cfnFunction constants.CfnFunctions) handler.ProgressEvent {
	privateEndpointID := fmt.Sprint(req.CallbackContext[id])

	getServerlessPrivateEndpointRequest := client.AtlasV2.ServerlessPrivateEndpointsApi.GetServerlessPrivateEndpoint(context.Background(), *currentModel.ProjectId, *currentModel.InstanceName, privateEndpointID)
	serverlessPrivateEndpoint, response, err := getServerlessPrivateEndpointRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("%s: error Serverless Private Endpoint %s",
			string(cfnFunction),
			err.Error()), response)
	}

	if serverlessPrivateEndpoint == nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("%s: Error while trying to make api call, CreateServerlessPrivateEndpoint returned status %d, and the response is NULL", response.StatusCode),
			cloudformation.HandlerErrorCodeInternalFailure)
	}

	if serverlessPrivateEndpoint.Status == nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("%s: Error while trying to get Serverless Private Endpoint, Private endpoint Status is null", string(cfnFunction)),
			cloudformation.HandlerErrorCodeServiceInternalError)
	}

	if *serverlessPrivateEndpoint.Status == string(targetStatus) {
		currentModel.completeWithAtlasModel(*serverlessPrivateEndpoint)
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         fmt.Sprintf("%s Completed", string(cfnFunction)),
			ResourceModel:   currentModel}
	} else if *serverlessPrivateEndpoint.Status == string(enums.Failed) {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("%s : the serverless private endpoint is in a Failed Status, error: %s", string(cfnFunction), *serverlessPrivateEndpoint.ErrorMessage), cloudformation.HandlerErrorCodeServiceInternalError)
	} else {
		return progressevents.GetInProgressProgressEvent(fmt.Sprintf("%s in progress", string(cfnFunction)), getCallbackContext(privateEndpointID), currentModel, callbackDelayInSeconds)
	}
}

func getCallbackContext(privateEndpointID string) map[string]interface{} {
	return map[string]interface{}{
		id: privateEndpointID,
	}
}

func (currentModel *Model) completeWithAtlasModel(atlasModel admin.ServerlessTenantEndpoint) {
	currentModel.Id = atlasModel.Id
	currentModel.Status = atlasModel.Status
	currentModel.ProviderName = atlasModel.ProviderName
	currentModel.ErrorMessage = atlasModel.ErrorMessage
	currentModel.EndpointServiceName = atlasModel.EndpointServiceName
}

func closeResponse(response *http.Response) {
	if response != nil {
		response.Body.Close()
	}
}
