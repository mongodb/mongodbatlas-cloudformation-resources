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

package constants_test

import (
	"testing"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/private-endpoint/cmd/constants"
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
