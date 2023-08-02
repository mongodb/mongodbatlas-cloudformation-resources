package resource

import (
	ctx "context"
	"fmt"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	atlasSDK "go.mongodb.org/atlas-sdk/v20230201002/admin"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.EndpointId}
var ReadRequiredFields = []string{constants.ProjectID, constants.EndpointId}
var DeleteRequiredFields = []string{constants.ProjectID, constants.EndpointId}
var ListRequiredFields = []string{constants.ProjectID}

const (
	AlreadyExists = "already exists"
)

func setup() {
	util.SetupLogger("mongodb-atlas-federated-query-limit")
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

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	readModel := Model{ProjectId: currentModel.ProjectId, EndpointId: currentModel.EndpointId}
	readResponse, err := readModel.getPrivateEndpoint(atlas)
	defer closeResponse(readResponse)
	if err == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          AlreadyExists,
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}

	response, err := createOrUpdate(currentModel, atlas)

	defer closeResponse(response)
	if err != nil {
		return handleError(response, err)
	}

	// Read endpoint
	response, err = currentModel.getPrivateEndpoint(atlas)
	defer closeResponse(response)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   currentModel}, nil
}

func createOrUpdate(currentModel *Model, atlas *util.MongoDBClient) (*http.Response, error) {
	provider := constants.AWS
	privateNetworkEndpointIdEntry := atlasSDK.PrivateNetworkEndpointIdEntry{
		EndpointId: *currentModel.EndpointId,
		Comment:    currentModel.Comment,
		Type:       currentModel.Type,
		Provider:   &provider,
	}
	createRequest := atlas.AtlasV2.DataFederationApi.CreateDataFederationPrivateEndpoint(
		ctx.Background(),
		*currentModel.ProjectId,
		&privateNetworkEndpointIdEntry,
	)
	_, response, err := createRequest.Execute()
	return response, err
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

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	response, err := currentModel.getPrivateEndpoint(atlas)

	defer closeResponse(response)
	if err != nil {
		return handleError(response, err)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
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

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	readModel := Model{ProjectId: currentModel.ProjectId, EndpointId: currentModel.EndpointId}
	readResponse, err := readModel.getPrivateEndpoint(atlas)
	defer closeResponse(readResponse)
	if err != nil {
		return handleError(readResponse, err)
	}
	response, err := createOrUpdate(currentModel, atlas)

	defer closeResponse(response)
	if err != nil {
		return handleError(response, err)
	}

	// Read endpoint
	response, err = currentModel.getPrivateEndpoint(atlas)
	defer closeResponse(response)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   currentModel}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	deleteRequest := atlas.AtlasV2.DataFederationApi.DeleteDataFederationPrivateEndpoint(
		ctx.Background(),
		*currentModel.ProjectId,
		*currentModel.EndpointId,
	)
	_, response, err := deleteRequest.Execute()

	//_, response, err := getFederatedQueryLimit(atlas, currentModel)
	defer closeResponse(response)
	if err != nil {
		return handleError(response, err)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
		ResourceModel:   nil}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	listRequest := atlas.AtlasV2.DataFederationApi.ListDataFederationPrivateEndpoints(
		ctx.Background(),
		*currentModel.ProjectId,
	)
	pe, response, err := listRequest.Execute()

	defer closeResponse(response)
	if err != nil {
		return handleError(response, err)
	}
	endpoints := make([]interface{}, len(pe.Results))
	for i := range pe.Results {
		endpoints[i] = Model{
			ProjectId:  currentModel.ProjectId,
			Profile:    currentModel.Profile,
			Comment:    pe.Results[i].Comment,
			Type:       pe.Results[i].Type,
			EndpointId: &pe.Results[i].EndpointId,
		}
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Completed",
		ResourceModels:  endpoints}, nil
}

func (model *Model) getPrivateEndpoint(atlas *util.MongoDBClient) (*http.Response, error) {
	readRequest := atlas.AtlasV2.DataFederationApi.GetDataFederationPrivateEndpoint(
		ctx.Background(),
		*model.ProjectId,
		*model.EndpointId,
	)
	paginatedPrivateNetworkEndpointIdEntry, response, err := readRequest.Execute()
	if err != nil {
		return response, err
	}
	model.readPrivateEndpoint(paginatedPrivateNetworkEndpointIdEntry)
	return response, err
}

func closeResponse(response *http.Response) {
	if response != nil {
		response.Body.Close()
	}
}

func (model *Model) readPrivateEndpoint(pe *atlasSDK.PrivateNetworkEndpointIdEntry) *Model {
	if pe == nil {
		return model
	}
	model.Comment = pe.Comment
	model.Type = pe.Type
	model.EndpointId = &pe.EndpointId
	return model
}
func (model *Model) readPrivateEndpoints(pe *atlasSDK.PaginatedPrivateNetworkEndpointIdEntry) []interface{} {
	if pe == nil {
		return nil
	}
	models := make([]interface{}, len(pe.Results))
	for i := range pe.Results {
		models[i] = &Model{
			ProjectId:  model.ProjectId,
			Profile:    model.Profile,
			Comment:    pe.Results[i].Comment,
			Type:       pe.Results[i].Type,
			EndpointId: &pe.Results[i].EndpointId,
		}
	}
	return models
}

func handleError(response *http.Response, err error) (handler.ProgressEvent, error) {
	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}
	return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error during execution : %s", err.Error()), response), nil
}
