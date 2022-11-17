package validator_def

type ModelValidator struct{}

var CreateRequiredFields = []string{"GroupId", "Region", "ApiKeys.PublicKey", "ApiKeys.PrivateKey", "SubnetId", "VpcId"}
var ReadRequiredFields = []string{"GroupId", "Id", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{"GroupId", "Id", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var ListRequiredFields = []string{"GroupId", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}

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
