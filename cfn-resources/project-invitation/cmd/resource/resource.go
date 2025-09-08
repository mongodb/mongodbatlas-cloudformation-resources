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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	admin20231115002 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

var CreateRequiredFields = []string{constants.ProjectID, constants.Username}
var ReadRequiredFields = []string{constants.ProjectID, constants.ID}
var UpdateRequiredFields = []string{constants.ProjectID, constants.ID}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ID}
var ListRequiredFields = []string{constants.ProjectID}

func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-project-invitation")
}

func validateProjectInvitationAlreadyAccepted(ctx context.Context, client *util.MongoDBClient, username, projectID string) (bool, error) {
	user, _, err := client.Atlas20231115002.MongoDBCloudUsersApi.GetUserByUsername(ctx, username).Execute()
	if err != nil {
		return false, err
	}
	for _, role := range user.Roles {
		if util.AreStringPtrEqual(role.GroupId, &projectID) {
			return true, nil
		}
	}

	return false, nil
}

func invitationAPIRequestToModel(currentModel *Model, invitation *admin20231115002.GroupInvitation) Model {
	out := Model{
		Profile:         currentModel.Profile,
		ProjectId:       currentModel.ProjectId,
		Username:        invitation.Username,
		Id:              invitation.Id,
		Roles:           invitation.Roles,
		ExpiresAt:       util.TimePtrToStringPtr(invitation.ExpiresAt),
		CreatedAt:       util.TimePtrToStringPtr(invitation.CreatedAt),
		InviterUsername: invitation.InviterUsername,
	}

	return out
}
