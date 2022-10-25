package validator_def

type ModelValidator struct{}

var CreateRequiredFields = []string{"ProjectId", "RegionName", "ApiKeys.PublicKey", "ApiKeys.PrivateKey", "AtlasCIDRBlock"}
var ReadRequiredFields = []string{"ProjectId", "Id", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var UpdateRequiredFields = []string{"ProjectId", "Id", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var DeleteRequiredFields = []string{"ProjectId", "Id", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var ListRequiredFields = []string{"ProjectId", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}

func (m ModelValidator) GetCreateFields() []string {
	return CreateRequiredFields
}
func (m ModelValidator) GetReadFields() []string {
	return ReadRequiredFields
}
func (m ModelValidator) GetUpdateFields() []string {
	return UpdateRequiredFields
}
func (m ModelValidator) GetDeleteFields() []string {
	return DeleteRequiredFields
}
func (m ModelValidator) GetListFields() []string {
	return ListRequiredFields
}
