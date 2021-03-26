package resource

import (
	"context"
	"fmt"
    "log"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
    "github.com/aws/aws-sdk-go/service/cloudformation"
	matlasClient "go.mongodb.org/atlas/mongodbatlas"
    "github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)


// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.Printf("Create log.Printf-- currentModel: %+v", *currentModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{
            OperationStatus: handler.Failed,
            Message: err.Error(),
            HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}
	project, _, err := client.Projects.Create(context.Background(), &matlasClient.Project{
		Name:  *currentModel.Name,
		OrgID: *currentModel.OrgId,
	})
	if err != nil {
		//return handler.ProgressEvent{}, fmt.Errorf("error creating project: %s", err)
        log.Printf("Create - error: %+v", err)
        // TODO- Should detect and return HandlerErrorCodeAlreadyExists
        return handler.ProgressEvent{
            OperationStatus: handler.Failed,
            Message: "Resource Not Found",
            HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil

	}

	currentModel.Id = &project.ID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	//spew.Dump(currentModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		//return handler.ProgressEvent{}, err
        return handler.ProgressEvent{
            Message: err.Error(),
            HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, err
	}

	name := *currentModel.Name

	if len(name) > 0 {
		project, _, err := client.Projects.GetOneProjectByName(context.Background(), name)
		if err != nil {
            return handler.ProgressEvent{
                OperationStatus: handler.Failed,
                Message: "Resource Not Found",
                HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
        }
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
		//return handler.ProgressEvent{}, fmt.Errorf(
        log.Printf("Read - error reading project with id(%s): %s", id, err)
        return handler.ProgressEvent{
            OperationStatus: handler.Failed,
            Message: "Resource Not Found",
            HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
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
		OperationStatus: handler.Failed,
		Message:         "Update Not Supported",
		ResourceModel:   nil,
        HandlerErrorCode: cloudformation.HandlerErrorCodeNotUpdatable,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return handler.ProgressEvent{}, err

	}
	log.Printf("Delete Project prevModel:%+v currentModel:%+v", *prevModel, *currentModel)
	//spew.Dump(prevModel)
    log.Println("----------------------------")
	//spew.Dump(currentModel)

	var id string
	if currentModel.Id != nil {
		id = *currentModel.Id
	}

    if len(id)==0 {
		name := *currentModel.Name
		if len(name) > 0 {
			log.Printf("Project id was nil, try lookup name:%s", name)
			project, _, err := client.Projects.GetOneProjectByName(context.Background(), name)
			if err != nil {
				return handler.ProgressEvent{
                    Message: err.Error(),
                    HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, err
			}
			log.Printf("Looked up project:%+v", project)
			id = project.ID
		} else {
            err := fmt.Errorf("@@@@Error deleting project. No Id or Name found currentModel:%+v)", currentModel)
            return handler.ProgressEvent{
                Message: err.Error(),
                HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, err
		}
	}
	log.Printf("Deleting project with id(%s)", id)

	_, err = client.Projects.Delete(context.Background(), id)
	if err != nil {
		//return handler.ProgressEvent{}, fmt.Errorf("####error deleting project with id(%s): %s", id, err)
        log.Printf("####error deleting project with id(%s): %s", id, err)

        return handler.ProgressEvent{
            OperationStatus: handler.Failed,
            Message: "Resource Not Found",
            HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		//ResourceModel:   currentModel,
		ResourceModel:   nil,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.Printf("List.Project prevModel:%+v currentModel:%+v", prevModel, currentModel)
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

    // Initialize like this in case no results will pass empty array
    mm := []interface{} {}
	for _, project := range projects.Results {
		var m Model
		m.Name = &project.Name
		m.OrgId = &project.OrgID
		m.Created = &project.Created
		m.ClusterCount = &project.ClusterCount
		m.Id = &project.ID
		mm = append(mm, m)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  mm,
	}, nil
}
