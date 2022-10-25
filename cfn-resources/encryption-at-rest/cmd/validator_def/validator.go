package validator_def

type ModelValidator struct{}

var CreateRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "AwsKms.RoleID", "AwsKms.CustomerMasterKeyID", "AwsKms.Region", "ProjectId"}
var ReadRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId"}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ProjectId"}
var ListRequiredFields = []string{}

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
