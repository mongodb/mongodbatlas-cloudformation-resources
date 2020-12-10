package resource

import (
	"errors"
	"fmt"
	"log"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
)

// Handler is a container for the CRUDL actions exported by resources
type Handler struct{}

// Create wraps the related Create function exposed by the resource code
func (r *Handler) Create(req handler.Request) handler.ProgressEvent {
	return wrap(req, Create)
}

// Read wraps the related Read function exposed by the resource code
func (r *Handler) Read(req handler.Request) handler.ProgressEvent {
	return wrap(req, Read)
}

// Update wraps the related Update function exposed by the resource code
func (r *Handler) Update(req handler.Request) handler.ProgressEvent {
	return wrap(req, Update)
}

// Delete wraps the related Delete function exposed by the resource code
func (r *Handler) Delete(req handler.Request) handler.ProgressEvent {
	return wrap(req, Delete)
}

// List wraps the related List function exposed by the resource code
func (r *Handler) List(req handler.Request) handler.ProgressEvent {
	return wrap(req, List)
}

// main is the entry point of the application.
func main() {
	cfn.Start(&Handler{})
}

type handlerFunc func(handler.Request, *Model, *Model) (handler.ProgressEvent, error)

func wrap(req handler.Request, f handlerFunc) (response handler.ProgressEvent) {
	defer func() {
		// Catch any panics and return a failed ProgressEvent
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = errors.New(fmt.Sprint(r))
			}

			log.Printf("Trapped error in handler: %v", err)

			response = handler.NewFailedEvent(err)
		}
	}()

	// Populate the previous model
	prevModel := &Model{}
	if err := req.UnmarshalPrevious(prevModel); err != nil {
		log.Printf("Error unmarshaling prev model: %v", err)
		return handler.NewFailedEvent(err)
	}

	// Populate the current model
	currentModel := &Model{}
	if err := req.Unmarshal(currentModel); err != nil {
		log.Printf("Error unmarshaling model: %v", err)
		return handler.NewFailedEvent(err)
	}

	response, err := f(req, prevModel, currentModel)
	if err != nil {
		log.Printf("Error returned from handler function: %v", err)
		return handler.NewFailedEvent(err)
	}

	return response
}
