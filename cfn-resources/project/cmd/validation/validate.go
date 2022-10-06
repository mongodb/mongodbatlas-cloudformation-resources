package validation

type ModelValidator struct{}

var CreateRequiredFields = []string{"Region", "ApiKeys.PublicKey", "ApiKeys.PrivateKey", "Name", "OrgId"}
var ReadRequiredFields = []string{"Id", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var UpdateRequiredFields = []string{"Id", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var DeleteRequiredFields = []string{"Region", "ApiKeys.PublicKey", "ApiKeys.PrivateKey", "Id"}
var ListRequiredFields = []string{"Region", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}

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
