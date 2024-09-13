package resource

import (
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"go.mongodb.org/atlas-sdk/v20240805004/admin"
)

func NewResourcePolicyCreateReq(model *Model) *admin.ApiAtlasResourcePolicyCreate {
	if model == nil {
		return nil
	}
	return &admin.ApiAtlasResourcePolicyCreate{
		Name:     *model.Name,
		Policies: *modelPoliciesToSDKPolicies(model.Policies),
	}
}

func modelPoliciesToSDKPolicies(policy []ApiAtlasPolicy) *[]admin.ApiAtlasPolicyCreate {
	policiesCreate := make([]admin.ApiAtlasPolicyCreate, len(policy))
	for i, v := range policy {
		policiesCreate[i].Body = *v.Body
	}
	return &policiesCreate
}

func sdkPoliciesToModelPolicies(policies *[]admin.ApiAtlasPolicy) []ApiAtlasPolicy {
	if policies == nil {
		return nil
	}
	sdkPolicies := make([]ApiAtlasPolicy, len(*policies))
	for i, v := range sdkPolicies {
		sdkPolicies[i].Body = v.Body
	}
	return sdkPolicies
}

func NewResourcePolicyUpdateReq(model *Model) *admin.ApiAtlasResourcePolicyEdit {
	if model == nil {
		return nil
	}
	return &admin.ApiAtlasResourcePolicyEdit{
		Name:     model.Name,
		Policies: modelPoliciesToSDKPolicies(model.Policies),
	}
}

func GetResourcePolicyModel(resourcePolicyResp *admin.ApiAtlasResourcePolicy, currentModel *Model) *Model {
	model := new(Model)

	if currentModel != nil {
		model = currentModel
	}
	model.CreatedByUser = newAPIAtlasUserMetadata(resourcePolicyResp.CreatedByUser)
	model.CreatedDate = util.TimePtrToStringPtr(resourcePolicyResp.CreatedDate)
	model.Id = resourcePolicyResp.Id
	model.LastUpdatedByUser = newAPIAtlasUserMetadata(resourcePolicyResp.LastUpdatedByUser)
	model.LastUpdatedDate = util.TimePtrToStringPtr(resourcePolicyResp.LastUpdatedDate)
	model.Name = resourcePolicyResp.Name
	model.OrgId = resourcePolicyResp.OrgId
	model.Policies = sdkPoliciesToModelPolicies(resourcePolicyResp.Policies)
	model.ResourcePolicyId = resourcePolicyResp.Id
	model.Version = resourcePolicyResp.Version
	return model
}

func newAPIAtlasUserMetadata(userMetadata *admin.ApiAtlasUserMetadata) *ApiAtlasUserMetadata {
	return &ApiAtlasUserMetadata{
		Id:   userMetadata.Id,
		Name: userMetadata.Name,
	}
}
