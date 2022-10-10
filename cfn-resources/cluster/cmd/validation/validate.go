package validation

type ModelValidator struct{}

var CreateRequiredFields = []string{"Region", "ApiKeys.PublicKey", "ApiKeys.PrivateKey", "Name", "ProjectId"}
var ReadRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "Name", "ProjectId"}
var UpdateRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var DeleteRequiredFields = []string{"Region", "ApiKeys.PublicKey", "ApiKeys.PrivateKey", "Name", "ProjectId"}
var ListRequiredFields = []string{"Region", "ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId"}

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
