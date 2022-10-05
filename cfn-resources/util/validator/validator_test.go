package validator_test

import (
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	"testing"
)

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

type testModelValidator struct{}

func (m testModelValidator) GetCreateFields() []string {
	return []string{
		"FirstRequiredField",
		"SecondRequiredField",
		"ThirdRequiredField",
		"FourthRequiredField",
		"RequiredStruct",
		"RequiredStruct.PropertyTest",
	}
}
func (m testModelValidator) GetReadFields() []string {
	return nil
}
func (m testModelValidator) GetUpdateFields() []string {
	return nil
}
func (m testModelValidator) GetDeleteFields() []string {
	return nil
}
func (m testModelValidator) GetListFields() []string {
	return nil
}

func TestAllValidateRequiredFieldsEmpty(t *testing.T) {
	modelValidator := testModelValidator{}
	requiredStruct := RequiredStructProperty{}
	model := testModel{
		FirstRequiredField:  nil,
		SecondRequiredField: nil,
		ThirdRequiredField:  nil,
		FourthRequiredField: nil,
		NotRequiredField:    nil,
		RequiredStruct:      &requiredStruct,
	}
	progressEvent := validator.ValidateModel(constants.Create, modelValidator, &model)

	if progressEvent == nil {
		t.Errorf("Progress Event should not be nill")
	}

	expected := "The next fields are required FirstRequiredField SecondRequiredField ThirdRequiredField FourthRequiredField RequiredStruct.PropertyTest"

	if progressEvent.Message != expected {
		t.Errorf("Expectd = %s; got = %s", expected, progressEvent.Message)
	}
}

func TestSomeValidateRequiredFieldsEmpty(t *testing.T) {
	firstField := "a"
	secondField := 1
	thirdField := []string{"a", "b"}
	modelValidator := testModelValidator{}
	requiredStruct := RequiredStructProperty{PropertyTest: &firstField}
	model := testModel{
		FirstRequiredField:  &firstField,
		SecondRequiredField: &secondField,
		ThirdRequiredField:  thirdField,
		FourthRequiredField: nil,
		NotRequiredField:    nil,
		RequiredStruct:      &requiredStruct,
	}
	progressEvent := validator.ValidateModel(constants.Create, modelValidator, &model)

	if progressEvent == nil {
		t.Errorf("Progress Event should not be nill")
	}

	expected := "The next fields are required FourthRequiredField"

	if progressEvent.Message != expected {
		t.Errorf("Expectd = %s; got = %s", expected, progressEvent.Message)
	}
}

func TestNoneValidateRequiredFieldsEmpty(t *testing.T) {
	firstField := "a"
	secondField := 1
	thirdField := []string{"a", "b"}
	modelValidator := testModelValidator{}
	fourthField := true
	requiredStruct := RequiredStructProperty{PropertyTest: &firstField}
	model := testModel{
		FirstRequiredField:  &firstField,
		SecondRequiredField: &secondField,
		ThirdRequiredField:  thirdField,
		FourthRequiredField: &fourthField,
		NotRequiredField:    nil,
		RequiredStruct:      &requiredStruct,
	}
	progressEvent := validator.ValidateModel(constants.Create, modelValidator, &model)

	if progressEvent != nil {
		t.Errorf("Progress Event should be nil")
	}
}
