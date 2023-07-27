// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"io/ioutil"
	"log"
	"os"

	//"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/federated-query-limit/cmd/resource"
)

// Handler is a container for the CRUDL actions exported by resources
type Handler struct{}

// Create wraps the related Create function exposed by the resource code
func (r *Handler) Create(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.Create)
}

// Read wraps the related Read function exposed by the resource code
func (r *Handler) Read(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.Read)
}

// Update wraps the related Update function exposed by the resource code
func (r *Handler) Update(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.Update)
}

// Delete wraps the related Delete function exposed by the resource code
func (r *Handler) Delete(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.Delete)
}

// List wraps the related List function exposed by the resource code
func (r *Handler) List(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.List)
}

// main is the entry point of the application.
//func main() {
//	cfn.Start(&Handler{})
//}

func main() {
	//cfn.Start(&Handler{})
	h := &Handler{}
	dir := "/federated-query-limit/test/federated-query-limit.sample.json"

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
	dir = path + dir
	data, err := ioutil.ReadFile(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	sess, err := session.NewSession(&aws.Config{Region: aws.String("ap-northeast-2")})
	if err != nil {
		fmt.Println(err)
		return
	}
	req := handler.NewRequest("id", map[string]interface{}{}, handler.RequestContext{}, sess, nil, data, data)
	h.Create(req)

}

type handlerFunc func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error)

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
	prevModel := &resource.Model{}
	if err := req.UnmarshalPrevious(prevModel); err != nil {
		log.Printf("Error unmarshaling prev model: %v", err)
		return handler.NewFailedEvent(err)
	}

	// Populate the current model
	currentModel := &resource.Model{}
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
