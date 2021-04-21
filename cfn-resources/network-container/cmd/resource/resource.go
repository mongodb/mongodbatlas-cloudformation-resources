package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"go.mongodb.org/atlas/mongodbatlas"
	"log"
)

const (
	defaultProviderName = "AWS"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Not Supported",
		ResourceModel:   currentModel,
	}, nil
	/*
			client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
			if err != nil {
				return handler.ProgressEvent{}, err
			}

			projectID := currentModel.ProjectId
			providerName := currentModel.ProviderName
			containerRequest := &mongodbatlas.Container{}

			if projectID == nil || *projectID == "" {
				return handler.ProgressEvent{}, fmt.Errorf("error creating network container: `ProjectID` must be set")
			}
			if providerName == nil || *providerName == "" {
				aws := defaultProviderName
				providerName = &aws
			}
			regionName := currentModel.RegionName
			if regionName == nil || *regionName == "" {
				return handler.ProgressEvent{}, fmt.Errorf("`error creating network container: RegionName` must be set")
			}
			containerRequest.RegionName = *regionName
			containerRequest.ProviderName = *providerName
			CIDR := currentModel.AtlasCIDRBlock
			if CIDR == nil || *CIDR == "" {
				return handler.ProgressEvent{}, fmt.Errorf("error creating network container: `AtlasCIDRBlock` must be set")
			}
			containerRequest.AtlasCIDRBlock = *CIDR
			containerResponse, res, err := client.Containers.Create(context.Background(), *projectID, containerRequest)
			if err != nil {
		        if res.StatusCode == 409 {
		            log.Printf("Container already exists for this group. Try return existing container. err: %v", err)
		            containers, _, err2 := client.Containers.ListAll(context.Background(), *projectID, nil)
		            if err2 != nil {
		                log.Printf("Error Containers.ListAll err:%v",err)
		                return handler.ProgressEvent{}, fmt.Errorf("error Containers.ListAll err:%v", err)
		            }
		            log.Printf("containers:%v",containers)
		            first := containers[0]
		            log.Printf("Will return reference to first container: first:%+v",first)
			        currentModel.Id = &first.ID
		        } else {
				    return handler.ProgressEvent{}, fmt.Errorf("error creating network container: %s", err)
		        }

			} else {
			    currentModel.Id = &containerResponse.ID
		    }

		    log.Printf("Create about to return this --->> currentModel:%+v",currentModel)
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Create complete",
				ResourceModel:   currentModel,
			}, nil
	*/
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Not Supported",
		ResourceModel:   currentModel,
	}, nil
	/*
		client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)

		if err != nil {
			return handler.ProgressEvent{}, err
		}

		projectID := *currentModel.ProjectId
		containerID := *currentModel.Id

		containerResponse, _, err := client.Containers.Get(context.Background(), projectID, containerID)

		if err != nil {
			return handler.ProgressEvent{}, fmt.Errorf("error reading container with id(project: %s, container: %s): %s", projectID, containerID, err)
		}

		currentModel.RegionName = &containerResponse.RegionName
		currentModel.Provisioned = containerResponse.Provisioned
		currentModel.VpcId = &containerResponse.VPCID
		currentModel.AtlasCIDRBlock = &containerResponse.AtlasCIDRBlock

		return handler.ProgressEvent{
			OperationStatus: handler.Success,
			Message:         "Read Complete",
			ResourceModel:   currentModel,
		}, nil
	*/
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Not Supported",
		ResourceModel:   currentModel,
	}, nil
	/*
			client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)

			if err != nil {
				return handler.ProgressEvent{}, err
			}

			projectId := *currentModel.ProjectId
			containerId := *currentModel.Id
			containerRequest := &mongodbatlas.Container{}
			providerName := currentModel.ProviderName
			if providerName == nil || *providerName == "" {
				aws := defaultProviderName
				providerName = &aws
			}
			CIDR := currentModel.AtlasCIDRBlock
			if CIDR != nil {
				containerRequest.AtlasCIDRBlock = *CIDR
			}
			containerRequest.ProviderName = *providerName
			containerRequest.RegionName = *currentModel.RegionName
			containerResponse, _, err := client.Containers.Update(context.Background(), projectId, containerId, containerRequest)
			if err != nil {
				return handler.ProgressEvent{}, fmt.Errorf("error updating container with id(project: %s, container: %s): %s", projectId, containerRequest, err)
			}

			currentModel.Id = &containerResponse.ID
		    log.Printf("Create network container - Id:%v",currentModel.Id)

			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Update Complete",
				ResourceModel:   currentModel,
			}, nil
	*/
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	log.Printf("Delete currentModel:%+v", currentModel)
	projectId := *currentModel.ProjectId
	containerId := *currentModel.Id

	_, err = client.Containers.Delete(context.Background(), projectId, containerId)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting container with id(project: %s, container: %s): %s", projectId, containerId, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("List currentModel:%+v", currentModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectId := *currentModel.ProjectId
	providerName := currentModel.ProviderName
	if providerName == nil || *providerName == "" {
		aws := defaultProviderName
		providerName = &aws
	}
	log.Printf("projectId:%v", projectId)
	log.Printf("providerName:%v", providerName)
	containerRequest := &mongodbatlas.ContainersListOptions{
		ProviderName: *providerName,
		ListOptions:  mongodbatlas.ListOptions{},
	}
	log.Printf("List projectId:%v, containerRequest:%v", projectId, containerRequest)
	containerResponse, _, err := client.Containers.List(context.TODO(), projectId, containerRequest)
	if err != nil {
		log.Printf("Error %v", err)
		return handler.ProgressEvent{}, err
	}
	log.Printf("containerResponse:%v", containerResponse)

	var models []Model
	for _, container := range containerResponse {
		var model Model
		model.RegionName = &container.RegionName
		model.Provisioned = container.Provisioned
		model.Id = &container.ID
		model.VpcId = &container.VPCID
		model.AtlasCIDRBlock = &container.AtlasCIDRBlock

		models = append(models, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   models,
	}, nil
}
