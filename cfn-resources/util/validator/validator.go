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

package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	cloudformationtypes "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	progressevents "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

func ValidateModel(fields []string, model interface{}) *handler.ProgressEvent {
	requiredFields := ""

	for _, field := range fields {
		if fieldIsEmpty(model, field) {
			requiredFields = fmt.Sprintf("%s %s", requiredFields, field)
		}
	}
	if requiredFields == "" {
		return nil
	}

	progressEvent := progressevents.GetFailedEventByCode(fmt.Sprintf("The next fields are required%s", requiredFields),
		string(cloudformationtypes.HandlerErrorCodeInvalidRequest))

	return &progressEvent
}

func fieldIsEmpty(model interface{}, field string) bool {
	var f reflect.Value
	if strings.Contains(field, ".") {
		fields := strings.Split(field, ".")
		r := reflect.ValueOf(model)

		for _, f := range fields {
			fmt.Println(f)
			baseProperty := reflect.Indirect(r).FieldByName(f)

			if baseProperty.IsNil() {
				return true
			}

			r = baseProperty
		}
		return false
	}
	r := reflect.ValueOf(model)
	f = reflect.Indirect(r).FieldByName(field)
	return f.IsNil()
}
