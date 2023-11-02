package resource

import (
	"reflect"
	"testing"
)

func TestConvertToAnySliceEmptyList(t *testing.T) {
	var input []string
	var expected []interface{}

	result, err := convertToAnySlice(input)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result does not match the expected value. Got: %v, Expected: %v", result, expected)
	}
}

func TestConvertToAnySliceValidJSON(t *testing.T) {
	input := []string{
		`{"type": "icuNormalize"}`,
	}
	expected := []interface{}{
		map[string]interface{}{"type": "icuNormalize"},
	}

	result, err := convertToAnySlice(input)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result does not match the expected value. Got: %v, Expected: %v", result, expected)
	}
}

func TestConvertToAnySliceInvalidJSON(t *testing.T) {
	input := []string{
		`this is an invalid json`,
	}
	var expected []interface{}

	result, err := convertToAnySlice(input)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Result does not match the expected value. Got: %v, Expected: %v", result, expected)
	}
}
