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
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"go.mongodb.org/atlas-sdk/v20250312013/admin"
)

func UpdateModelFromSecret(model *Model, secret *admin.ServiceAccountSecret) {
	if secret == nil {
		return
	}
	model.SecretId = &secret.Id
	model.Secret = secret.Secret
	model.MaskedSecretValue = secret.MaskedSecretValue
	model.CreatedAt = util.TimePtrToStringPtr(&secret.CreatedAt)
	model.ExpiresAt = util.TimePtrToStringPtr(&secret.ExpiresAt)
	model.LastUsedAt = util.TimePtrToStringPtr(secret.LastUsedAt)
}
