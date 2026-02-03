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
	"sort"

	"go.mongodb.org/atlas-sdk/v20250312013/admin"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

func GetOrgServiceAccountModel(account *admin.OrgServiceAccount, currentModel *Model) *Model {
	model := new(Model)

	if currentModel != nil {
		model = currentModel
	}
	if account != nil {
		if currentModel != nil {
			model.OrgId = currentModel.OrgId
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
			model.Secrets = make([]Secret, len(*account.Secrets))
			for i, s := range *account.Secrets {
				createdAt := s.CreatedAt
				expiresAt := s.ExpiresAt
				model.Secrets[i] = Secret{
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

func NewOrgServiceAccountCreateReq(model *Model) *admin.OrgServiceAccountRequest {
	if model == nil {
		return nil
	}
	if model.SecretExpiresAfterHours == nil {
		return nil
	}
	secretExpiresAfterHours := *model.SecretExpiresAfterHours
	roles := make([]string, len(model.Roles))
	copy(roles, model.Roles)
	sort.Strings(roles)
	return &admin.OrgServiceAccountRequest{
		Name:                    *model.Name,
		Description:             *model.Description,
		Roles:                   roles,
		SecretExpiresAfterHours: secretExpiresAfterHours,
	}
}

func NewOrgServiceAccountUpdateReq(model *Model) *admin.OrgServiceAccountUpdateRequest {
	if model == nil {
		return nil
	}
	var roles *[]string
	if len(model.Roles) > 0 {
		sortedRoles := make([]string, len(model.Roles))
		copy(sortedRoles, model.Roles)
		sort.Strings(sortedRoles)
		roles = &sortedRoles
	}
	return &admin.OrgServiceAccountUpdateRequest{
		Name:        model.Name,
		Description: model.Description,
		Roles:       roles,
	}
}
