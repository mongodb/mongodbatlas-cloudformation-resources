package resource

import "go.mongodb.org/atlas-sdk/v20240805004/admin"

func NewResourcePolicyCreateReq(currentModel *Model) *admin.ApiAtlasResourcePolicyCreate {
	return &admin.ApiAtlasResourcePolicyCreate{}
}

func NewResourcePolicyUpdateReq(currentModel *Model) *admin.ApiAtlasResourcePolicyEdit {
	return &admin.ApiAtlasResourcePolicyEdit{}
}

func GetResourcePolicyModel(resourcePolicyResp *admin.ApiAtlasResourcePolicy, currentModel *Model) *Model {
	return &Model{}
}
