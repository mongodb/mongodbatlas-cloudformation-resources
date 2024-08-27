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

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"go.mongodb.org/atlas-sdk/v20240805001/admin"
)

var CreateRequiredFields = []string{constants.OrgID, constants.Username}
var ReadRequiredFields = []string{constants.OrgID, constants.ID}
var UpdateRequiredFields = []string{constants.OrgID, constants.ID}
var DeleteRequiredFields = []string{constants.OrgID, constants.ID}
var ListRequiredFields = []string{constants.OrgID}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-OrgInvitation")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	_, _ = log.Debugf("Create() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.Atlas20231115002

	invitationReq := &admin.OrganizationInvitationRequest{
		TeamIds:  currentModel.TeamIds,
		Roles:    currentModel.Roles,
		Username: currentModel.Username,
	}
	invitation, res, err := atlasV2.OrganizationsApi.CreateOrganizationInvitation(context.Background(), *currentModel.OrgId, invitationReq).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}
	currentModel.Id = invitation.Id

	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   currentModel,
	}, nil
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	_, _ = log.Debugf("Read() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.Atlas20231115002

	invitation, res, err := atlasV2.OrganizationsApi.GetOrganizationInvitation(context.Background(), *currentModel.OrgId, *currentModel.Id).Execute()
	if err != nil {
		_, _ = log.Debugf("Read - error: %+v", err)

		// if invitation already accepted
		if res.StatusCode == 404 {
			if alreadyAccepted, _ := validateOrgInvitationAlreadyAccepted(context.Background(), atlasV2, *currentModel.Username, *currentModel.OrgId); alreadyAccepted {
				return progressevent.GetFailedEventByResponse("invitation has been already accepted", res), nil
			}
		}

		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
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
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	_, _ = log.Warnf("Update() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.Atlas20231115002

	invitationReq := &admin.OrganizationInvitationUpdateRequest{
		TeamIds: currentModel.TeamIds,
		Roles:   currentModel.Roles,
	}

	invitation, res, err := atlasV2.OrganizationsApi.UpdateOrganizationInvitationById(context.Background(), *currentModel.OrgId, *currentModel.Id, invitationReq).Execute()

	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
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
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	_, _ = log.Debugf("Delete() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.Atlas20231115002

	_, res, err := atlasV2.OrganizationsApi.DeleteOrganizationInvitation(context.Background(), *currentModel.OrgId, *currentModel.Id).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
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
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)

	_, _ = log.Debugf("List() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	atlasV2 := client.Atlas20231115002

	invitations, res, err := atlasV2.OrganizationsApi.ListOrganizationInvitationsWithParams(context.Background(), &admin.ListOrganizationInvitationsApiParams{
		OrgId:    *currentModel.OrgId,
		Username: currentModel.Username,
	}).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), res), nil
	}

	var invites []interface{}
	// iterate invites
	for i := range invitations {
		invite := &Model{}
		model := readAtlasOrgInvitation(&invitations[i], invite)
		invites = append(invites, model)
	}

	// Response
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModels:  invites,
	}, nil
}

func readAtlasOrgInvitation(invitation *admin.OrganizationInvitation, currentModel *Model) (model *Model) {
	currentModel.Username = invitation.Username
	currentModel.OrgId = invitation.OrgId
	currentModel.Id = invitation.Id
	currentModel.TeamIds = invitation.TeamIds
	currentModel.Roles = invitation.Roles
	currentModel.ExpiresAt = util.TimePtrToStringPtr(invitation.ExpiresAt)
	currentModel.CreatedAt = util.TimePtrToStringPtr(invitation.CreatedAt)
	currentModel.InviterUsername = invitation.InviterUsername
	return currentModel
}

func validateOrgInvitationAlreadyAccepted(ctx context.Context, atlasV2 *admin.APIClient, username, orgID string) (bool, error) {
	user, _, err := atlasV2.MongoDBCloudUsersApi.GetUserByUsername(ctx, username).Execute()
	if err != nil {
		return false, err
	}

	for _, role := range user.Roles {
		if util.AreStringPtrEqual(role.OrgId, &orgID) {
			return true, nil
		}
	}
	return false, nil
}
