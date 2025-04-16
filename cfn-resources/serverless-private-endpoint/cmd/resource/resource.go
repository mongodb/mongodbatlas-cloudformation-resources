// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go.mongodb.org/atlas-sdk/v20250312002/admin"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/serverless-private-endpoint/cmd/resource/enums"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	aws_utils "github.com/mongodb/mongodbatlas-cloudformation-resources/util/aws"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.InstanceName}
var ReadRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ID}
var DeleteRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ID, AwsPrivateEndpointMetaData}
var UpdateRequiredFields = []string{constants.ProjectID, constants.InstanceName, constants.ID}
var ListRequiredFields = []string{constants.ProjectID, constants.InstanceName}

const (
	id                         = "id"
	stateName                  = "StateName"
	endpointServiceName        = "endpoint_service_name"
	callbackDelayInSeconds     = 5
	AwsPrivateEndpointMetaData = "AwsPrivateEndpointMetaData"
)

func setup() {
	util.SetupLogger("mongodb-atlas-serverless-private-endpoint")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	errorProgressEvent := currentModel.validateAwsPrivateEndpointProperties()
	if errorProgressEvent != nil {
		return *errorProgressEvent, nil
	}

	if currentModel.ProviderName == nil {
		currentModel.ProviderName = aws.String(constants.AWS)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	status, pe := getProcessStatus(req)
	if pe != nil {
		return *pe, nil
	}

	switch status {
	case enums.Init:
		atlasPrivateEndpoint, errPe := createAtlasPrivateEndpoint(currentModel, client)
		if errPe != nil {
			return *errPe, nil
		}

		currentModel.completeWithAtlasModel(*atlasPrivateEndpoint)
		callbackContext := getCallbackContext(*currentModel.Id, currentModel.EndpointServiceName)
		callbackContext[stateName] = enums.CreatingPrivateEndpoint

		return progressevents.GetInProgressProgressEvent("Creating ", callbackContext, currentModel, callbackDelayInSeconds), nil
	case enums.CreatingPrivateEndpoint:
		progressEvent := validateCompletion(req, currentModel, client, enums.Reserved, constants.CREATE)
		if progressEvent.OperationStatus != handler.Success {
			progressEvent.CallbackContext = req.CallbackContext
			return progressEvent, nil
		}

		if !*currentModel.CreateAndAssignAWSPrivateEndpoint {
			return progressEvent, nil
		}

		currentModel = progressEvent.ResourceModel.(*Model)

		awsPrivateEndpoint, peErr := createAwsPrivateEndpoint(currentModel, req)
		if peErr != nil {
			return *peErr, nil
		}

		return assignAwsPrivateEndpoint(req, client, *awsPrivateEndpoint, currentModel), nil

	default:
		progressEvent := validateCompletion(req, currentModel, client, enums.Available, constants.CREATE)
		if progressEvent.OperationStatus == handler.InProgress {
			progressEvent.CallbackContext = req.CallbackContext
			progressEvent.CallbackContext[stateName] = enums.InitiatingPrivateEndpoint
		}

		return progressEvent, nil
	}
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	createServerlessPrivateEndpointRequest := client.AtlasSDK.ServerlessPrivateEndpointsApi.GetServerlessPrivateEndpoint(context.Background(),
		*currentModel.ProjectId, *currentModel.InstanceName, *currentModel.Id)
	serverlessPrivateEndpoint, response, err := createServerlessPrivateEndpointRequest.Execute()
	defer response.Body.Close()
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

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

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

	createServerlessPrivateEndpointRequest := client.AtlasSDK.ServerlessPrivateEndpointsApi.UpdateServerlessPrivateEndpoint(context.Background(),
		*currentModel.ProjectId, *currentModel.InstanceName, *currentModel.Id, &serverlessPrivateEndpointInput)
	serverlessPrivateEndpoint, response, err := createServerlessPrivateEndpointRequest.Execute()
	defer response.Body.Close()
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

	currentModel.validateAwsPrivateEndpointProperties()

	// Check if the model is valid for deletion
	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	createAndAssignAWSPrivateEndpoint, region := unmarshallAwsMetadata(*currentModel.AwsPrivateEndpointMetaData)
	if !isRequestInProgress(req) && createAndAssignAWSPrivateEndpoint {
		if region == nil {
			return progressevents.GetFailedEventByCode("Error deleting aws private Endpoint region is null", cloudformation.HandlerErrorCodeServiceInternalError), nil
		}
		errPe := deleteAwsPrivateEndpoint(currentModel, *region, client, req)
		if errPe != nil {
			return *errPe, nil
		}
	}

	deleteServerlessPrivateEndpointRequest := client.AtlasSDK.ServerlessPrivateEndpointsApi.DeleteServerlessPrivateEndpoint(context.Background(),
		*currentModel.ProjectId, *currentModel.InstanceName, *currentModel.Id)
	response, err := deleteServerlessPrivateEndpointRequest.Execute()
	defer response.Body.Close()
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

	return progressevents.GetInProgressProgressEvent("Create in progress", getCallbackContext(*currentModel.Id, aws.String("")), currentModel, callbackDelayInSeconds), nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	listServerlessPrivateEndpointRequest := client.AtlasSDK.ServerlessPrivateEndpointsApi.ListServerlessPrivateEndpoints(context.Background(),
		*currentModel.ProjectId, *currentModel.InstanceName)
	serverlessPrivateEndpoints, response, err := listServerlessPrivateEndpointRequest.Execute()
	defer response.Body.Close()
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("error getting Serverless Private Endpoint %s",
			err.Error()), response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Completed",
		ResourceModels:  ConvertListToModelList(serverlessPrivateEndpoints, currentModel.Profile, currentModel.ProjectId, currentModel.InstanceName)}, nil
}

func createAwsPrivateEndpoint(currentModel *Model, req handler.Request) (*aws_utils.PrivateEndpointOutput, *handler.ProgressEvent) {
	awsPrivateEndpointInput := aws_utils.PrivateEndpointInput{
		VpcID:     *currentModel.AwsPrivateEndpointConfigurationProperties.VpcId,
		SubnetIDs: currentModel.AwsPrivateEndpointConfigurationProperties.SubnetIds,
	}

	output, errpe := aws_utils.CreatePrivateEndpoint(req, *currentModel.EndpointServiceName,
		*currentModel.AwsPrivateEndpointConfigurationProperties.Region, []aws_utils.PrivateEndpointInput{awsPrivateEndpointInput})

	if errpe != nil {
		return nil, errpe
	}

	return &output[0], nil
}

func deleteAwsPrivateEndpoint(currentModel *Model, region string, client *util.MongoDBClient, req handler.Request) *handler.ProgressEvent {
	createServerlessPrivateEndpointRequest := client.AtlasSDK.ServerlessPrivateEndpointsApi.GetServerlessPrivateEndpoint(context.Background(),
		*currentModel.ProjectId, *currentModel.InstanceName, *currentModel.Id)
	serverlessPrivateEndpoint, response, err := createServerlessPrivateEndpointRequest.Execute()
	defer response.Body.Close()
	if err != nil {
		if isTenantPrivateEndpointNotFound(response) {
			pe := progressevents.GetFailedEventByCode(fmt.Sprintf("error getting Serverless Private Endpoint %s", err.Error()), cloudformation.HandlerErrorCodeNotFound)
			return &pe
		}
		pe := progressevents.GetFailedEventByResponse(fmt.Sprintf("error getting Serverless Private Endpoint %s",
			err.Error()), response)
		return &pe
	}

	interfaceEndpoint := []string{
		*serverlessPrivateEndpoint.CloudProviderEndpointId,
	}

	errorProgressEvent := aws_utils.DeletePrivateEndpoint(req, interfaceEndpoint, region)

	if errorProgressEvent != nil {
		return errorProgressEvent
	}

	return nil
}

func createAtlasPrivateEndpoint(currentModel *Model, client *util.MongoDBClient) (*admin.ServerlessTenantEndpoint, *handler.ProgressEvent) {
	serverlessPrivateEndpointInput := admin.ServerlessTenantCreateRequest{Comment: currentModel.Comment}

	createServerlessPrivateEndpointRequest := client.AtlasSDK.ServerlessPrivateEndpointsApi.CreateServerlessPrivateEndpoint(context.Background(),
		*currentModel.ProjectId, *currentModel.InstanceName, &serverlessPrivateEndpointInput)
	serverlessPrivateEndpoint, response, err := createServerlessPrivateEndpointRequest.Execute()
	defer response.Body.Close()
	if err != nil {
		errPe := progressevents.GetFailedEventByResponse(fmt.Sprintf("error creating Serverless Private Endpoint %s",
			err.Error()), response)
		return nil, &errPe
	}

	if serverlessPrivateEndpoint == nil {
		errPe := progressevents.GetFailedEventByCode(
			fmt.Sprintf("Error while trying to make api call, CreateServerlessPrivateEndpoint returned status %d, and the response is NULL",
				response.StatusCode),
			cloudformation.HandlerErrorCodeInternalFailure)
		return nil, &errPe
	}

	return serverlessPrivateEndpoint, nil
}

func assignAwsPrivateEndpoint(req handler.Request, client *util.MongoDBClient, awsPrivateEndpoint aws_utils.PrivateEndpointOutput, currentModel *Model) handler.ProgressEvent {
	serverlessPrivateEndpointInput := admin.ServerlessTenantEndpointUpdate{
		Comment:                 currentModel.Comment,
		ProviderName:            *currentModel.ProviderName,
		CloudProviderEndpointId: &awsPrivateEndpoint.InterfaceEndpointID,
	}

	createServerlessPrivateEndpointRequest := client.AtlasSDK.ServerlessPrivateEndpointsApi.UpdateServerlessPrivateEndpoint(context.Background(),
		*currentModel.ProjectId, *currentModel.InstanceName, *currentModel.Id, &serverlessPrivateEndpointInput)
	serverlessPrivateEndpoint, response, err := createServerlessPrivateEndpointRequest.Execute()
	defer response.Body.Close()

	if err != nil {
		if isTenantPrivateEndpointNotFound(response) {
			return progressevents.GetFailedEventByCode(fmt.Sprintf("error updating Serverless Private Endpoint %s", err.Error()), cloudformation.HandlerErrorCodeNotFound)
		}
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("error updating Serverless Private Endpoint %s",
			err.Error()), response)
	}

	if serverlessPrivateEndpoint == nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Error while trying to make api call, CreateServerlessPrivateEndpoint returned status %d, and the response is NULL", response.StatusCode),
			cloudformation.HandlerErrorCodeInternalFailure)
	}

	callbackContext := req.CallbackContext
	callbackContext[stateName] = enums.InitiatingPrivateEndpoint

	return progressevents.GetInProgressProgressEvent("Create in progress", callbackContext, currentModel, callbackDelayInSeconds)
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

func ConvertListToModelList(endpoints []admin.ServerlessTenantEndpoint, profileName, projectID, instanceName *string) []interface{} {
	models := make([]interface{}, 0)

	for _, endpoint := range endpoints {
		model := Model{
			InstanceName: instanceName,
			ProjectId:    projectID,
			Profile:      profileName,
		}

		model.completeWithAtlasModel(endpoint)
		models = append(models, model)
	}

	return models
}

func isRequestInProgress(req handler.Request) bool {
	_, idExists := req.CallbackContext[id]
	return idExists
}

func validateCompletion(req handler.Request, currentModel *Model, client *util.MongoDBClient, targetStatus enums.AtlasPrivateEndpointStatus, cfnFunction constants.CfnFunctions) handler.ProgressEvent {
	privateEndpointID := fmt.Sprint(req.CallbackContext[id])

	getServerlessPrivateEndpointRequest := client.AtlasSDK.ServerlessPrivateEndpointsApi.GetServerlessPrivateEndpoint(context.Background(),
		*currentModel.ProjectId, *currentModel.InstanceName, privateEndpointID)
	serverlessPrivateEndpoint, response, err := getServerlessPrivateEndpointRequest.Execute()
	defer response.Body.Close()
	if err != nil {
		return progressevents.GetFailedEventByResponse(fmt.Sprintf("%s: error Serverless Private Endpoint %s",
			string(cfnFunction),
			err.Error()), response)
	}

	if serverlessPrivateEndpoint == nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("%s: Error while trying to make api call, CreateServerlessPrivateEndpoint returned status %d, and the response is NULL",
			string(cfnFunction), response.StatusCode),
			cloudformation.HandlerErrorCodeInternalFailure)
	}

	if serverlessPrivateEndpoint.Status == nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("%s: Error while trying to get Serverless Private Endpoint, Private endpoint AtlasPrivateEndpointStatus is null", string(cfnFunction)),
			cloudformation.HandlerErrorCodeServiceInternalError)
	}

	switch *serverlessPrivateEndpoint.Status {
	case string(targetStatus):
		currentModel.completeWithAtlasModel(*serverlessPrivateEndpoint)
		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         fmt.Sprintf("%s Completed", string(cfnFunction)),
			ResourceModel:   currentModel}
	case string(enums.Failed):
		return progressevents.GetFailedEventByCode(fmt.Sprintf("%s : the serverless private endpoint is in a Failed AtlasPrivateEndpointStatus, error: %s", string(cfnFunction),
			*serverlessPrivateEndpoint.ErrorMessage), cloudformation.HandlerErrorCodeServiceInternalError)
	default:
		return progressevents.GetInProgressProgressEvent(fmt.Sprintf("%s in progress", string(cfnFunction)),
			getCallbackContext(privateEndpointID, serverlessPrivateEndpoint.EndpointServiceName), currentModel, callbackDelayInSeconds)
	}
}

func getCallbackContext(privateEndpointID string, serviceName *string) map[string]interface{} {
	callbackContext := map[string]interface{}{
		id: privateEndpointID,
	}
	if serviceName != nil {
		callbackContext[endpointServiceName] = *serviceName
	}

	return callbackContext
}

func (currentModel *Model) completeWithAtlasModel(atlasModel admin.ServerlessTenantEndpoint) {
	currentModel.Id = atlasModel.Id
	currentModel.Status = atlasModel.Status
	currentModel.ProviderName = atlasModel.ProviderName
	currentModel.ErrorMessage = atlasModel.ErrorMessage
	currentModel.EndpointServiceName = atlasModel.EndpointServiceName
	currentModel.completeAwsPrivateEndpointMetaData()
}

func (currentModel *Model) validateAwsPrivateEndpointProperties() *handler.ProgressEvent {
	if currentModel.CreateAndAssignAWSPrivateEndpoint == nil {
		currentModel.CreateAndAssignAWSPrivateEndpoint = aws.Bool(false)
		return nil
	}

	if *currentModel.CreateAndAssignAWSPrivateEndpoint {
		if currentModel.AwsPrivateEndpointConfigurationProperties == nil {
			pe := progressevents.GetFailedEventByCode(
				"Validation failed: AwsPrivateEndpointConfigurationProperties must be present when CreateAndAssignAWSPrivateEndpoint is true",
				cloudformation.HandlerErrorCodeInvalidRequest)
			return &pe
		}

		if currentModel.AwsPrivateEndpointConfigurationProperties.VpcId == nil {
			pe := progressevents.GetFailedEventByCode("Validation failed: VpcId must be present when CreateAndAssignAWSPrivateEndpoint is true", cloudformation.HandlerErrorCodeInvalidRequest)
			return &pe
		}

		if currentModel.AwsPrivateEndpointConfigurationProperties.Region == nil {
			pe := progressevents.GetFailedEventByCode("Validation failed: REgion must be present when CreateAndAssignAWSPrivateEndpoint is true", cloudformation.HandlerErrorCodeInvalidRequest)
			return &pe
		}

		if len(currentModel.AwsPrivateEndpointConfigurationProperties.SubnetIds) == 0 {
			pe := progressevents.GetFailedEventByCode("Validation failed: SubnetIds must be present when CreateAndAssignAWSPrivateEndpoint is true", cloudformation.HandlerErrorCodeInvalidRequest)
			return &pe
		}
	}

	return nil
}

func getProcessStatus(req handler.Request) (enums.EventStatus, *handler.ProgressEvent) {
	callback := req.CallbackContext[stateName]
	if callback == nil {
		return enums.Init, nil
	}

	eventStatus, err := enums.ParseEventStatus(fmt.Sprintf("%v", callback))

	if err != nil {
		pe := progressevents.GetFailedEventByCode(fmt.Sprintf("Error parsing callback status : %s", err.Error()),
			cloudformation.HandlerErrorCodeServiceInternalError)
		return "", &pe
	}

	return eventStatus, nil
}

func unmarshallAwsMetadata(input string) (createAwsPrivateEndpoint bool, region *string) {
	parts := strings.Split(input, "/")
	if len(parts) != 2 {
		return false, nil
	}

	boolValue := parts[0] == "true"
	if parts[1] == "" {
		return boolValue, nil
	}

	stringValue := parts[1]
	return boolValue, &stringValue
}

func (currentModel *Model) completeAwsPrivateEndpointMetaData() {
	result := fmt.Sprintf("%v/", currentModel.CreateAndAssignAWSPrivateEndpoint != nil && *currentModel.CreateAndAssignAWSPrivateEndpoint)
	if currentModel.AwsPrivateEndpointConfigurationProperties != nil && currentModel.AwsPrivateEndpointConfigurationProperties.Region != nil {
		result += *currentModel.AwsPrivateEndpointConfigurationProperties.Region
	}
	currentModel.AwsPrivateEndpointMetaData = &result
}
