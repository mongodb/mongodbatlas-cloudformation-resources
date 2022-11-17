package constants_test

import (
	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/constants"
	"testing"
)

func TestParseEventStatusSuccess(t *testing.T) {
	testStatus := "CREATING_PRIVATE_ENDPOINT_SERVICE"

	status, err := constants.ParseEventStatus(testStatus)

	if err != nil {
		t.Errorf("Error %s", err.Error())
	}

	if status != constants.CreatingPrivateEndpointService {
		t.Errorf("Unexpected Status , expected :%s , found: %s", constants.CreatingPrivateEndpointService, status)
	}
}

func TestParseEventStatusWithInvalidInput(t *testing.T) {
	testStatus := "adsfsadgsadgsadgsdag"

	_, err := constants.ParseEventStatus(testStatus)

	if err == nil {
		t.Errorf("error Should not be nill")
	}
}
