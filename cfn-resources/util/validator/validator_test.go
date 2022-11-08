package validator_test

import (
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
)

var CreateRequiredFields = []string{""}
var ReadRequiredFields = []string{"ProjectId", "ClusterName", "Id", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var UpdateRequiredFields []string
var DeleteRequiredFields = []string{"ProjectId", "ClusterName", "Id", "ApiKeys.PublicKey", "ApiKeys.PrivateKey"}
var ListRequiredFields = []string{"ApiKeys.PublicKey", "ApiKeys.PrivateKey", "ClusterName", "ProjectId"}

type testModel struct {
	FirstRequiredField  *string
	SecondRequiredField *int
	ThirdRequiredField  []string
	FourthRequiredField *bool
	RequiredStruct      *RequiredStructProperty
	NotRequiredField    *int
}

type RequiredStructProperty struct {
	PropertyTest *string
}

func TestAllValidateRequiredFieldsEmpty(t *testing.T) {
	requiredStruct := RequiredStructProperty{}
	fields := []string{"FirstRequiredField", "SecondRequiredField", "ThirdRequiredField", "FourthRequiredField", "RequiredStruct.PropertyTest"}
	model := testModel{
		FirstRequiredField:  nil,
		SecondRequiredField: nil,
		ThirdRequiredField:  nil,
		FourthRequiredField: nil,
		NotRequiredField:    nil,
		RequiredStruct:      &requiredStruct,
	}
	progressEvent := validator.ValidateModel(fields, &model)

	expected := "The following fields are required FirstRequiredField SecondRequiredField ThirdRequiredField FourthRequiredField RequiredStruct.PropertyTest"
	if progressEvent != nil && progressEvent.Message != expected {
		t.Errorf("Expectd = %s; got = %s", expected, progressEvent.Message)
	}
}

func TestSomeValidateRequiredFieldsEmpty(t *testing.T) {
	firstField := "a"
	secondField := 1
	thirdField := []string{"a", "b"}
	fields := []string{"FirstRequiredField", "SecondRequiredField", "FourthRequiredField"}
	requiredStruct := RequiredStructProperty{PropertyTest: &firstField}
	model := testModel{
		FirstRequiredField:  &firstField,
		SecondRequiredField: &secondField,
		ThirdRequiredField:  thirdField,
		FourthRequiredField: nil,
		NotRequiredField:    nil,
		RequiredStruct:      &requiredStruct,
	}
	progressEvent := validator.ValidateModel(fields, &model)

	expected := "The following fields are required FourthRequiredField"

	if progressEvent != nil && progressEvent.Message != expected {
		t.Errorf("Expectd = %s; got = %s", expected, progressEvent.Message)
	}
}

func TestNoneValidateRequiredFieldsEmpty(t *testing.T) {
	firstField := "a"
	secondField := 1
	thirdField := []string{"a", "b"}
	fourthField := true
	var fields []string

	requiredStruct := RequiredStructProperty{PropertyTest: &firstField}
	model := testModel{
		FirstRequiredField:  &firstField,
		SecondRequiredField: &secondField,
		ThirdRequiredField:  thirdField,
		FourthRequiredField: &fourthField,
		NotRequiredField:    nil,
		RequiredStruct:      &requiredStruct,
	}
	progressEvent := validator.ValidateModel(fields, &model)

	if progressEvent != nil {
		t.Errorf("Progress Event should be nil")
	}
}
