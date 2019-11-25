package resource

import (
	"context"
	"fmt"
	"log"

	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"

	"github.com/aws-cloudformation/aws-cloudformation-rpdk-go-plugin/cfn/encoding"
	"github.com/aws-cloudformation/aws-cloudformation-rpdk-go-plugin/cfn/handler"
	matlasClient "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func createMongoDBClient(publicKey, privateKey string) (*matlasClient.Client, error) {
	// setup a transport to handle digest
	transport := digest.NewTransport(publicKey, privateKey)

	// initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, err
	}

	//Initialize the MongoDB Atlas API Client.
	return matlasClient.NewClient(client), nil

}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := createMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	project, _, err := client.Projects.Create(context.Background(), &matlasClient.Project{
		Name:  *currentModel.Name.Value(),
		OrgID: *currentModel.OrgId.Value(),
	})
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating project: %s", err)
	}

	currentModel.Id = encoding.NewString(project.ID)
	currentModel.Created = encoding.NewString(project.Created)
	currentModel.ClusterCount = encoding.NewInt(int64(project.ClusterCount))

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := createMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	id := *currentModel.Id.Value()
	log.Printf("Looking for project: %s", id)

	project, _, err := client.Projects.GetOneProject(context.Background(), id)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading project with id(%s): %s", id, err)
	}

	currentModel.Name = encoding.NewString(project.Name)
	currentModel.OrgId = encoding.NewString(project.OrgID)
	currentModel.Created = encoding.NewString(project.Created)
	currentModel.ClusterCount = encoding.NewInt(int64(project.ClusterCount))

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

	client, err := createMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	id := *currentModel.Id.Value()
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
	client, err := createMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projects, _, err := client.Projects.GetAllProjects(context.Background())
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error retrieving projects: %s", err)
	}

	var models []Model
	for _, project := range projects.Results {
		var m Model
		m.Name = encoding.NewString(project.Name)
		m.OrgId = encoding.NewString(project.OrgID)
		m.Created = encoding.NewString(project.Created)
		m.ClusterCount = encoding.NewInt(int64(project.ClusterCount))
		m.Id = encoding.NewString(project.ID)

		models = append(models, m)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   models,
	}, nil
}
