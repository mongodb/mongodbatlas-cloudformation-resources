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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	log "github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	_, _ = log.Debugf("Create() currentModel:%+v", currentModel)

	errValidation := validateModel(CreateRequiredFields, currentModel)
	if errValidation != nil {
		return *errValidation, nil
	}

	if !util.IsStringPresent(currentModel.Profile) {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	invitationReq := &admin20231115002.GroupInvitationRequest{
		Roles:    currentModel.Roles,
		Username: currentModel.Username,
	}

	invitation, res, err := client.Atlas20231115002.ProjectsApi.CreateProjectInvitation(context.Background(), *currentModel.ProjectId, invitationReq).Execute()
	if err != nil {
		return progressevents.GetFailedEventByResponse(err.Error(), res), nil
	}
	currentModel.Id = invitation.Id

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		ResourceModel:   invitationAPIRequestToModel(currentModel, invitation),
	}, nil
}
