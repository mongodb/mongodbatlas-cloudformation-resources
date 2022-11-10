package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
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
		cloudformation.HandlerErrorCodeInvalidRequest)

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
