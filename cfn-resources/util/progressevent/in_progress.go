package progressevents

import "github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"

func GetInProgressProgressEvent(message string, callBackContext map[string]interface{}, model interface{}, delaySeconds int64) handler.ProgressEvent {
	return handler.ProgressEvent{
		OperationStatus:      handler.InProgress,
		Message:              message,
		CallbackDelaySeconds: delaySeconds,
		ResourceModel:        model,
		CallbackContext:      callBackContext}
}
