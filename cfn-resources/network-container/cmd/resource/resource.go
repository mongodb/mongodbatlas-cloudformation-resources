package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progress_event"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
	"log"
	"net/http"
)

const (
	defaultProviderName = "AWS"
)

var CreateRequiredFields = []string{"ProjectId", "RegionName", "ApiKeys.PublicKey", "ApiKeys.PrivateKey", "AtlasCidrBlock"}
var ReadRequiredFields = []string{"ProjectId", "Id", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var UpdateRequiredFields = []string{"ProjectId", "Id", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var DeleteRequiredFields = []string{"ProjectId", "Id", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var ListRequiredFields = []string{"ProjectId", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}

// function to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := currentModel.ProjectId

	containerRequest := &mongodbatlas.Container{}

	if projectID == nil || *projectID == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network container: `ProjectID` must be set")
	}

	regionName := currentModel.RegionName
	if regionName == nil || *regionName == "" {
		return handler.ProgressEvent{}, fmt.Errorf("`error creating network container: RegionName` must be set")
	}
	containerRequest.RegionName = *regionName
	containerRequest.ProviderName = defaultProviderName
	CIDR := currentModel.AtlasCidrBlock
	if CIDR == nil || *CIDR == "" {
		return handler.ProgressEvent{}, fmt.Errorf("error creating network container: `AtlasCidrBlock` must be set")
	}
	containerRequest.AtlasCIDRBlock = *CIDR
	containerResponse, res, err := client.Containers.Create(context.Background(), *projectID, containerRequest)
	if err != nil {
		if res.StatusCode == http.StatusConflict {
			log.Printf("Container already exists for this group. Try return existing container. err: %v", err)
			containers, _, err2 := client.Containers.ListAll(context.Background(), *projectID, nil)
			if err2 != nil {
				log.Printf("Error Containers.ListAll err:%v", err)
				return handler.ProgressEvent{}, fmt.Errorf("error Containers.ListAll err:%v", err)
			}
			log.Printf("containers:%v", containers)
			first := containers[0]
			log.Printf("Will return reference to first container: first:%+v", first)
			currentModel.Id = &first.ID
		} else {
			return handler.ProgressEvent{}, fmt.Errorf("error creating network container: %s", err)
		}

	} else {
		currentModel.Id = &containerResponse.ID
	}

	log.Printf("Create about to return this --->> currentModel:%+v", currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectID := *currentModel.ProjectId
	containerID := *currentModel.Id

	containerResponse, response, err := client.Containers.Get(context.Background(), projectID, containerID)

	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response), nil
	}

	currentModel.RegionName = &containerResponse.RegionName
	currentModel.Provisioned = containerResponse.Provisioned
	currentModel.VpcId = &containerResponse.VPCID
	currentModel.AtlasCidrBlock = &containerResponse.AtlasCIDRBlock

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil

}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectId := *currentModel.ProjectId
	containerId := *currentModel.Id
	containerRequest := &mongodbatlas.Container{}

	CIDR := currentModel.AtlasCidrBlock
	if CIDR != nil {
		containerRequest.AtlasCIDRBlock = *CIDR
	}
	containerRequest.ProviderName = defaultProviderName
	containerRequest.RegionName = *currentModel.RegionName
	containerResponse, _, err := client.Containers.Update(context.Background(), projectId, containerId, containerRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error updating container with id(project: %s, container: %s): %s", projectId, containerRequest, err)
	}

	currentModel.Id = &containerResponse.ID
	log.Printf("Create network container - Id:%v", currentModel.Id)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	spew.Dump(currentModel)
	log.Printf("Delete currentModel:%+v", currentModel)
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	log.Printf("Delete currentModel:%+v", currentModel)
	projectId := *currentModel.ProjectId
	containerId := *currentModel.Id

	response, err := client.Containers.Delete(context.Background(), projectId, containerId)

	if err != nil {
		return progress_events.GetFailedEventByResponse(fmt.Sprintf("Error getting resource : %s", err.Error()),
			response.Response), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("List currentModel:%+v", currentModel)

	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectId := *currentModel.ProjectId

	log.Printf("projectId:%v", projectId)
	containerRequest := &mongodbatlas.ContainersListOptions{
		ProviderName: defaultProviderName,
		ListOptions:  mongodbatlas.ListOptions{},
	}
	log.Printf("List projectId:%v, containerRequest:%v", projectId, containerRequest)
	containerResponse, _, err := client.Containers.List(context.TODO(), projectId, containerRequest)
	if err != nil {
		log.Printf("Error %v", err)
		return handler.ProgressEvent{}, err
	}

	log.Printf("containerResponse:%v", containerResponse)
	spew.Dump(containerResponse)

	mm := make([]interface{}, 0)
	for _, container := range containerResponse {
		var m Model
		m.completeByConnection(container)
		mm = append(mm, m)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  mm,
	}, nil
}

func (m *Model) completeByConnection(c mongodbatlas.Container) {
	m.RegionName = &c.RegionName
	m.Provisioned = c.Provisioned
	m.Id = &c.ID
	m.VpcId = &c.VPCID
	m.AtlasCidrBlock = &c.AtlasCIDRBlock
}
