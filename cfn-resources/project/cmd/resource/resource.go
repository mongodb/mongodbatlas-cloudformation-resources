package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	matlasClient "go.mongodb.org/atlas/mongodbatlas"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.Printf("Create -- currentModel: %+v", currentModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	project, _, err := client.Projects.Create(context.Background(), &matlasClient.Project{
		Name:  *currentModel.Name,
		OrgID: *currentModel.OrgId,
	})
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating project: %s", err)
	}

	currentModel.Id = &project.ID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount

	/* TEMP DISABLE DEPOLOY SECRET FOR PROJECT
	      PENDING REVIEW - not sure need for this resource, yet.

	   // This is the intial call to Create, so inject a deployment
	   // secret for this resource in order to lookup progress properly
	   resourceID := util.NewResourceIdentifier("Project", project.ID, nil)
	   log.Printf("Created resourceID:%s",resourceID)
	   resourceProps := map[string]string{
	       "Name": project.Name,
	       "OrgID": project.OrgID,
	       "ResourceID": resourceID.String(),
	   }
	   secretName, err := util.CreateDeploymentSecret(&req, resourceID, *currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey, &resourceProps)
	   if err != nil {
	       log.Printf("Error - %+v",err)
	       return handler.ProgressEvent{}, err
	   }
	   log.Printf("Created deployment secret:%s this should be set to the resource priary key",secretName)

	*/

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	spew.Dump(currentModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	name := *currentModel.Name

	if len(name) > 0 {
		project, _, err := client.Projects.GetOneProjectByName(context.Background(), name)
		currentModel.Name = &project.Name
		currentModel.OrgId = &project.OrgID
		currentModel.Created = &project.Created
		currentModel.ClusterCount = &project.ClusterCount

		if err == nil {
			return handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Read Complete",
				ResourceModel:   currentModel,
			}, nil
		}
	}

	id := *currentModel.Id
	log.Printf("Looking for project: %s", id)
	project, _, err := client.Projects.GetOneProject(context.Background(), id)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading project with id(%s): %s", id, err)
	}

	currentModel.Name = &project.Name
	currentModel.OrgId = &project.OrgID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// no-op

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	log.Printf("Delet Project prevModel:%+v currentModel:%+v", prevModel, currentModel)

	var id string
	if currentModel.Id != nil {
		id = *currentModel.Id
	} else {
		name := *currentModel.Name
		if len(name) > 0 {
			log.Printf("Project id was nil, try lookup name:%s", name)
			project, _, err := client.Projects.GetOneProjectByName(context.Background(), name)
			if err != nil {
				return handler.ProgressEvent{}, err
			}
			log.Printf("Looked up project:%+v", project)
			id = project.ID
		} else {
			return handler.ProgressEvent{}, fmt.Errorf("Error deleting project. No Id or Name found currentModel:%+v)", currentModel)
		}
	}
	log.Printf("Deleting project with id(%s)", id)

	_, err = client.Projects.Delete(context.Background(), id)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting project with id(%s): %s", id, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	listOptions := &matlasClient.ListOptions{
		PageNum:      0,
		ItemsPerPage: 100,
	}
	projects, _, err := client.Projects.GetAllProjects(context.Background(), listOptions)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error retrieving projects: %s", err)
	}

	var models []Model
	for _, project := range projects.Results {
		var m Model
		m.Name = &project.Name
		m.OrgId = &project.OrgID
		m.Created = &project.Created
		m.ClusterCount = &project.ClusterCount
		m.Id = &project.ID

		models = append(models, m)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   models,
	}, nil
}
