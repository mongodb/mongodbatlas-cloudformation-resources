package resource

import (
	"context"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"github.com/openlyinc/pointy"
	log "github.com/sirupsen/logrus"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.Criteria, constants.CriteriaType, constants.PubKey, constants.PvtKey}
var ReadRequiredFields = []string{constants.ProjectID, constants.ArchiveId, constants.ClusterName, constants.PubKey, constants.PvtKey}
var UpdateRequiredFields = []string{constants.ProjectID, constants.ArchiveId, constants.ClusterName, constants.Criteria, constants.PubKey, constants.PvtKey}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ArchiveId, constants.ClusterName, constants.PubKey, constants.PvtKey}
var ListRequiredFields = []string{constants.PubKey, constants.PvtKey, constants.ProjectID, constants.ClusterName}

var requiredCriteriaType = map[string][]string{
	"DATE":   {"DateField", "ExpireAfterDays"},
	"CUSTOM": {"Query"},
}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-onlinearchives")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Create() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	inputRequest := mapToArchivePayload(currentModel)
	outputRequest, res, err := client.OnlineArchives.Create(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, &inputRequest)

	if err != nil {
		log.Debugf("Create - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	spew.Dump(outputRequest)
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func mapToArchivePayload(currentModel *Model) mongodbatlas.OnlineArchive {
	requestInput := mongodbatlas.OnlineArchive{
		DBName:   *currentModel.DbName,
		CollName: *currentModel.CollName,
	}
	requestInput.Criteria = mapCriteria(currentModel)
	return m
}

func mapCriteria(currentModel *Model) (*mongodbatlas.OnlineArchiveCriteria, error) {
	criteriaModel := *currentModel.Criteria

	criteriaInput := &mongodbatlas.OnlineArchiveCriteria{
		Type: *criteriaModel.Type,
	}
	if criteriaInput.Type == "DATE" {
		requiredInputs := requiredCriteriaType[criteriaInput.Type]
		// Validation
		criteriaInputDate := validateModel(requiredInputs, currentModel)
		if criteriaInputDate != nil {
			return *criteriaInputDate, nil
		}
		criteriaInput.DateField = *criteriaModel.DateField
		conversion := *criteriaModel.ExpireAfterDays
		criteriaInput.ExpireAfterDays = pointy.Float64(float64(conversion))

	}
	if criteriaInput.Type == "CUSTOM" {
		criteriaInput.Query = *criteriaModel.Query
	}
	return criteriaInput
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Read() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Read - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	/*
	   Considerable params from currentModel:
	   GroupId, ArchiveId, ClusterName, ...
	*/
	/*
	    // Pseudocode:
	    res , resModel, err := client.Onlinearchives.Read(context.Background(),&mongodbatlas.Onlinearchives{
	   	GroupId:currentModel.GroupId,
	   	ArchiveId:currentModel.ArchiveId,
	   	ClusterName:currentModel.ClusterName,
	   })

	*/

	if err != nil {
		log.Debugf("Read - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Update() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Update - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	/*
	   Considerable params from currentModel:
	   Criteria, ClusterName, ItemsPerPage, Results, Links, ArchiveId, DbName, CollName, Id, TotalCount, IncludeCount, PageNum, CollectionType, State, Type, Schedule, PartitionFields, GroupId, ApiKeys, ...
	*/
	/*
	    // Pseudocode:
	    res , resModel, err := client.Onlinearchives.Update(context.Background(),&mongodbatlas.Onlinearchives{
	   	Criteria:currentModel.Criteria,
	   	ClusterName:currentModel.ClusterName,
	   	ItemsPerPage:currentModel.ItemsPerPage,
	   	Results:currentModel.Results,
	   	Links:currentModel.Links,
	   	ArchiveId:currentModel.ArchiveId,
	   	DbName:currentModel.DbName,
	   	CollName:currentModel.CollName,
	   	Id:currentModel.Id,
	   	TotalCount:currentModel.TotalCount,
	   	IncludeCount:currentModel.IncludeCount,
	   	PageNum:currentModel.PageNum,
	   	CollectionType:currentModel.CollectionType,
	   	State:currentModel.State,
	   	Type:currentModel.Type,
	   	Schedule:currentModel.Schedule,
	   	PartitionFields:currentModel.PartitionFields,
	   	GroupId:currentModel.GroupId,
	   	ApiKeys:currentModel.ApiKeys,
	   })

	*/

	if err != nil {
		log.Debugf("Update - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("Delete() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("Delete - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	/*
	   Considerable params from currentModel:
	   ArchiveId, ClusterName, ...
	*/
	/*
	    // Pseudocode:
	    res , resModel, err := client.Onlinearchives.Delete(context.Background(),&mongodbatlas.Onlinearchives{
	   	ArchiveId:currentModel.ArchiveId,
	   	ClusterName:currentModel.ClusterName,
	   })

	*/

	if err != nil {
		log.Debugf("Delete - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	log.Debugf("List() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		log.Debugf("List - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}
	var res *mongodbatlas.Response

	//
	/*
	    // Pseudocode:
	    res , resModel, err := client.Onlinearchives.List(context.Background(),&mongodbatlas.Onlinearchives{
	   })

	*/

	if err != nil {
		log.Debugf("List - error: %+v", err)
		return progress_events.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	log.Debugf("Atlas Client %v", client)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}
