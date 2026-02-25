// Copyright 2026 MongoDB Inc
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
	"errors"
	"fmt"
	"sort"

	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

func GetGroupServiceAccountModel(account *admin.GroupServiceAccount, currentModel *Model) *Model {
	model := new(Model)

	if currentModel != nil {
		model = currentModel
	}
	if account != nil {
		if currentModel != nil {
			model.ProjectId = currentModel.ProjectId
			model.Profile = currentModel.Profile
		}
		model.Name = account.Name
		model.Description = account.Description
		if account.Roles != nil {
			roles := *account.Roles
			// Preserve order from currentModel if it exists (required for CFN contract tests)
			// Otherwise preserve API response order
			if currentModel != nil && currentModel.Roles != nil && len(currentModel.Roles) > 0 {
				model.Roles = currentModel.Roles
			} else {
				model.Roles = roles
			}
		}
		model.ClientId = account.ClientId
		model.CreatedAt = util.TimePtrToStringPtr(account.CreatedAt)

		if account.Secrets != nil {
			model.Secrets = make([]SecretDefinition, len(*account.Secrets))
			for i, s := range *account.Secrets {
				createdAt := s.CreatedAt
				expiresAt := s.ExpiresAt
				model.Secrets[i] = SecretDefinition{
					Id:                &s.Id,
					CreatedAt:         util.TimePtrToStringPtr(&createdAt),
					ExpiresAt:         util.TimePtrToStringPtr(&expiresAt),
					LastUsedAt:        util.TimePtrToStringPtr(s.LastUsedAt),
					MaskedSecretValue: s.MaskedSecretValue,
					Secret:            s.Secret,
				}
			}
		}
	}
	return model
}

func NewGroupServiceAccountCreateReq(model *Model) (*admin.GroupServiceAccountRequest, error) {
	if model == nil {
		return nil, errors.New("model is nil")
	}
	if model.SecretExpiresAfterHours == nil {
		return nil, fmt.Errorf("SecretExpiresAfterHours is required")
	}
	secretExpiresAfterHours := *model.SecretExpiresAfterHours
	roles := make([]string, len(model.Roles))
	copy(roles, model.Roles)
	sort.Strings(roles)
	return &admin.GroupServiceAccountRequest{
		Name:                    *model.Name,
		Description:             *model.Description,
		Roles:                   roles,
		SecretExpiresAfterHours: secretExpiresAfterHours,
	}, nil
}

func NewGroupServiceAccountUpdateReq(model *Model) (*admin.GroupServiceAccountUpdateRequest, error) {
	if model == nil {
		return nil, errors.New("model is nil")
	}
	var roles *[]string
	if len(model.Roles) > 0 {
		sortedRoles := make([]string, len(model.Roles))
		copy(sortedRoles, model.Roles)
		sort.Strings(sortedRoles)
		roles = &sortedRoles
	}
	return &admin.GroupServiceAccountUpdateRequest{
		Name:        model.Name,
		Description: model.Description,
		Roles:       roles,
	}, nil
}
