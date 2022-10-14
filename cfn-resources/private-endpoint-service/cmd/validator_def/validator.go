package validator_def

type ModelValidator struct{}

var CreateRequiredFields = []string{"GroupId", "EndpointServiceId", "InterfaceEndpointConnectionId", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var ReadRequiredFields = []string{"GroupId", "EndpointServiceId", "InterfaceEndpointConnectionId", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{"GroupId", "EndpointServiceId", "InterfaceEndpointConnectionId", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var ListRequiredFields []string

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
