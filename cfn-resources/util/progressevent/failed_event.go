package progressevents

import (
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

func getHandlerErrorCode(response *http.Response) string {
	switch response.StatusCode {
	case http.StatusBadRequest:
		return cloudformation.HandlerErrorCodeInvalidRequest
	case http.StatusNotFound:
		return cloudformation.HandlerErrorCodeNotFound
	case http.StatusInternalServerError:
		return cloudformation.HandlerErrorCodeServiceInternalError
	case http.StatusPaymentRequired, http.StatusUnauthorized:
		return cloudformation.HandlerErrorCodeAccessDenied
	default:
		return cloudformation.HandlerErrorCodeInternalFailure
	}
}

func GetFailedEventByResponse(message string, response *http.Response) handler.ProgressEvent {
	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          message,
		HandlerErrorCode: getHandlerErrorCode(response)}
}

func GetFailedEventByCode(message, handlerErrorCode string) handler.ProgressEvent {
	return handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          message,
		HandlerErrorCode: handlerErrorCode}
}
