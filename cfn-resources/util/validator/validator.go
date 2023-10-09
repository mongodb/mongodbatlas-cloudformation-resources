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
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
)

func ValidateModel(fields []string, model any) *handler.ProgressEvent {
	requiredFields := make([]string, 0)
	for _, field := range fields {
		if isEmptyField(model, field) {
			requiredFields = append(requiredFields, field)
		}
	}
	if len(requiredFields) == 0 {
		return nil
	}
	msg := fmt.Sprintf("These fields are required: %s", strings.Join(requiredFields, ", "))
	progressEvent := progressevent.GetFailedEventByCode(msg, cloudformation.HandlerErrorCodeInvalidRequest)
	return &progressEvent
}

func isEmptyField(model any, field string) bool {
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
