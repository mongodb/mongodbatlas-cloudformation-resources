package testutil

import (
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
)

type Operation func(req handler.Request) handler.ProgressEvent

type TestHandler interface {
	Create(req handler.Request) handler.ProgressEvent
	Read(req handler.Request) handler.ProgressEvent
	Update(req handler.Request) handler.ProgressEvent
	Delete(req handler.Request) handler.ProgressEvent
	List(req handler.Request) handler.ProgressEvent
}
