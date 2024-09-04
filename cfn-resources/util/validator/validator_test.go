// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	FourthRequiredField *bool
	RequiredStruct      *RequiredStructProperty
	NotRequiredField    *int
	ThirdRequiredField  []string
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

	expected := "The next fields are required FirstRequiredField SecondRequiredField ThirdRequiredField FourthRequiredField RequiredStruct.PropertyTest"
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

	expected := "The next fields are required FourthRequiredField"

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
