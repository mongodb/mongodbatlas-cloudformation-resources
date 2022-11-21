package resource

import (
	"context"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.PubKey, constants.PvtKey, constants.Username}
var ReadRequiredFields = []string{constants.ProjectID, constants.ID, constants.PubKey, constants.PvtKey}
var UpdateRequiredFields = []string{constants.ProjectID, constants.ID, constants.PubKey, constants.PvtKey}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ID, constants.PubKey, constants.PvtKey}
var ListRequiredFields = []string{constants.PubKey, constants.PubKey}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-projectinvitation")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Create() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Warnf("Create - error: %+v", err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	invitationReq := &mongodbatlas.Invitation{
		Roles:    currentModel.Roles,
		Username: *currentModel.Username,
	}

	invitation, res, err := client.Projects.InviteUser(context.Background(), *currentModel.ProjectId, invitationReq)
	if err != nil {
		_, _ = log.Warnf("Create - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	currentModel.Id = &invitation.ID

	return Read(req, prevModel, currentModel)
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Read() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	invitation, res, err := client.Projects.Invitation(context.Background(), *currentModel.ProjectId, *currentModel.Id)
	if err != nil {
		_, _ = log.Warnf("Read - error: %+v", err)

		// if invitation already accepted
		if res.StatusCode == 404 {
			if alreadyAccepted, _ := validateProjectInvitationAlreadyAccepted(context.Background(), client, *currentModel.Username, *currentModel.ProjectId); alreadyAccepted {
				return progressevents.GetFailedEventByResponse("invitation has been already accepted", res.Response), nil
			}
		}

		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	readAtlasProjectInvitation(invitation, currentModel)
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Warnf("Update() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	invitationReq := &mongodbatlas.Invitation{
		Roles: currentModel.Roles,
	}

	_, res, err := client.Projects.UpdateInvitationByID(context.Background(), *currentModel.ProjectId, *currentModel.Id, invitationReq)
	if err != nil {
		_, _ = log.Warnf("Update - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	_, _ = log.Debugf("%s invitation updated", *currentModel.Id)

	// Response
	return Read(req, prevModel, currentModel)
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Delete() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Warnf("Delete - error: %+v", err)
		return handler.ProgressEvent{
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			Message:          err.Error(),
			OperationStatus:  handler.Failed,
		}, nil
	}

	res, err := client.Projects.DeleteInvitation(context.Background(), *currentModel.ProjectId, *currentModel.Id)
	if err != nil {
		_, _ = log.Warnf("Delete - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	_, _ = log.Debugf("deleted invitation with Id :%s", *currentModel.Id)

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   nil,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("List() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	if err != nil {
		_, _ = log.Warnf("List - error: %+v", err)
		return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
			cloudformation.HandlerErrorCodeInvalidRequest), nil
	}

	listOptions := &mongodbatlas.InvitationOptions{
		Username: *currentModel.Username,
	}
	invitations, res, err := client.Projects.Invitations(context.Background(), *currentModel.ProjectId, listOptions)
	if err != nil {
		_, _ = log.Warnf("List - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	var invites []interface{}
	// iterate invites
	for i := range invitations {
		invite := &Model{}
		readAtlasProjectInvitation(invitations[i], invite)
		invites = append(invites, invite)
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModels:  invites,
	}, nil
}

func readAtlasProjectInvitation(invitation *mongodbatlas.Invitation, currentModel *Model) {
	currentModel.Username = &invitation.Username
	currentModel.Id = &invitation.ID
	currentModel.Roles = invitation.Roles
	currentModel.ExpiresAt = &invitation.ExpiresAt
	currentModel.CreatedAt = &invitation.CreatedAt
	currentModel.InviterUsername = &invitation.InviterUsername
}

func validateProjectInvitationAlreadyAccepted(ctx context.Context, client *mongodbatlas.Client, username, projectId string) (bool, error) {
	user, _, err := client.AtlasUsers.GetByName(ctx, username)
	if err != nil {
		return false, err
	}
	for _, role := range user.Roles {
		if role.GroupID == projectId {
			return true, nil
		}
	}

	return false, nil
}
