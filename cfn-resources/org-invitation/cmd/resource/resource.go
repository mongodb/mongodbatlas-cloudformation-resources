// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"context"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.OrgID, constants.Username}
var ReadRequiredFields = []string{constants.OrgID, constants.ID}
var UpdateRequiredFields = []string{constants.OrgID, constants.ID}
var DeleteRequiredFields = []string{constants.OrgID, constants.ID}
var ListRequiredFields = []string{constants.OrgID}

//var ListRequiredFields = []string{constants.PubKey, constants.PubKey}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-OrgInvitation")
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
	//client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	//if err != nil {
	//	_, _ = log.Warnf("Create - error: %+v", err)
	//	return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
	//		cloudformation.HandlerErrorCodeInvalidRequest), nil
	//}
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	invitationReq := &mongodbatlas.Invitation{
		TeamIDs:  currentModel.TeamIds,
		Roles:    currentModel.Roles,
		Username: *currentModel.Username,
	}
	invitation, res, err := client.Organizations.InviteUser(context.Background(), *currentModel.OrgId, invitationReq)
	if err != nil {
		_, _ = log.Warnf("Create - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	currentModel.Id = &invitation.ID

	if err != nil {
		_, _ = log.Warnf("Read - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
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
	//client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	//if err != nil {
	//	return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
	//		cloudformation.HandlerErrorCodeInvalidRequest), nil
	//}
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	invitation, res, err := client.Organizations.Invitation(context.Background(), *currentModel.OrgId, *currentModel.Id)
	if err != nil {
		_, _ = log.Warnf("Read - error: %+v", err)

		// if invitation already accepted
		if res.StatusCode == 404 {
			if alreadyAccepted, _ := validateOrgInvitationAlreadyAccepted(context.Background(), client, *currentModel.Username, *currentModel.OrgId); alreadyAccepted {
				return progressevents.GetFailedEventByResponse("invitation has been already accepted", res.Response), nil
			}
		}

		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	model := readAtlasOrgInvitation(invitation, currentModel)
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   model,
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
	//client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	//if err != nil {
	//	return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
	//		cloudformation.HandlerErrorCodeInvalidRequest), nil
	//}
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	invitationReq := &mongodbatlas.Invitation{
		TeamIDs: currentModel.TeamIds,
		Roles:   currentModel.Roles,
	}

	invitation, res, err := client.Organizations.UpdateInvitationByID(context.Background(), *currentModel.OrgId, *currentModel.Id, invitationReq)
	if err != nil {
		_, _ = log.Warnf("Update - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}
	_, _ = log.Debugf("%s invitation updated", *currentModel.Id)

	model := readAtlasOrgInvitation(invitation, currentModel)
	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   model,
	}, nil
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
	//client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	//if err != nil {
	//	_, _ = log.Warnf("Delete - error: %+v", err)
	//	return handler.ProgressEvent{
	//		HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
	//		Message:          err.Error(),
	//		OperationStatus:  handler.Failed,
	//	}, nil
	//}
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	res, err := client.Organizations.DeleteInvitation(context.Background(), *currentModel.OrgId, *currentModel.Id)
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
	//client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
	//if err != nil {
	//	_, _ = log.Warnf("List - error: %+v", err)
	//	return progressevents.GetFailedEventByCode(fmt.Sprintf("Failed to Create Client : %s", err.Error()),
	//		cloudformation.HandlerErrorCodeInvalidRequest), nil
	//}
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(util.DefaultProfile)
	}

	client, peErr := util.NewMongoDBClient(req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	listOptions := &mongodbatlas.InvitationOptions{
		Username: *currentModel.Username,
	}
	invitations, res, err := client.Organizations.Invitations(context.Background(), *currentModel.OrgId, listOptions)
	if err != nil {
		_, _ = log.Warnf("List - error: %+v", err)
		return progressevents.GetFailedEventByResponse(err.Error(), res.Response), nil
	}

	var invites []interface{}
	// iterate invites
	for i := range invitations {
		invite := &Model{}
		model := readAtlasOrgInvitation(invitations[i], invite)
		invites = append(invites, model)
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModels:  invites,
	}, nil
}

func readAtlasOrgInvitation(invitation *mongodbatlas.Invitation, currentModel *Model) (model *Model) {
	currentModel.Username = &invitation.Username
	currentModel.OrgId = &invitation.OrgID
	currentModel.Id = &invitation.ID
	currentModel.TeamIds = invitation.TeamIDs
	currentModel.Roles = invitation.Roles
	currentModel.ExpiresAt = &invitation.ExpiresAt
	currentModel.CreatedAt = &invitation.CreatedAt
	currentModel.InviterUsername = &invitation.InviterUsername
	return currentModel
}

func validateOrgInvitationAlreadyAccepted(ctx context.Context, client *mongodbatlas.Client, username, orgID string) (bool, error) {
	user, _, err := client.AtlasUsers.GetByName(ctx, username)
	if err != nil {
		return false, err
	}

	for _, role := range user.Roles {
		if role.OrgID == orgID {
			return true, nil
		}
	}
	return false, nil
}
