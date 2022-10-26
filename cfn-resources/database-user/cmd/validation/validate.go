package validation

type ModelValidator struct{}

var CreateRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "DatabaseName", "ProjectId", "Roles", "Username"}
var ReadRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId", "DatabaseName", "Username"}
var UpdateRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "DatabaseName", "ProjectId", "Roles", "Username"}
var DeleteRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId", "DatabaseName", "Username"}
var ListRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId"}

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
