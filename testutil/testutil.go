package testutil

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws/session"
)

type TestCase struct {
	Name        string
	Steps       []TestStep
	TestHandler TestHandler
}

type TestStep struct {
	Config string
	Check  TestCheckFunc
}

type TestCheckFunc func(model interface{}) error

// TestT is the interface used to handle the test lifecycle of a test.
//
// Users should just use a *testing.T object, which implements this.
type TestT interface {
	Error(args ...interface{})
	Fatal(args ...interface{})
	Skip(args ...interface{})
	Name() string
	Parallel()
}

func ComposeTestCheckFunc(fs ...TestCheckFunc) TestCheckFunc {
	return func(model interface{}) error {
		for i, f := range fs {
			if err := f(model); err != nil {
				return fmt.Errorf("Check %d/%d error: %s", i+1, len(fs), err)
			}
		}
		return nil
	}
}

func TestCheckResourceAttr(key string, value interface{}) TestCheckFunc {
	return func(model interface{}) error {
		data, err := json.Marshal(model)
		if err != nil {
			return err
		}

		d := map[string]interface{}{}
		err = json.Unmarshal(data, &d)
		if err != nil {
			return err
		}

		v, ok := d[key]
		if !ok {
			return fmt.Errorf("Key %s not found", key)
		}
		if v != value {
			return fmt.Errorf("Given value %v is different from configuration value %v", value, v)
		}

		return nil
	}
}

func Test(t TestT, ts TestCase) {
	log.Printf("[INFO] Running case %s\n", ts.Name)

	var model interface{}
	for i, test := range ts.Steps {
		log.Printf("[DEBUG] Test: Executing step %d", i)
		if model == nil {
			data := []byte(test.Config)
			req := handler.NewRequest("id", map[string]interface{}{}, &session.Session{}, nil, data)
			h := ts.TestHandler.Create(req)
			if h.OperationStatus != handler.Success {
				t.Error(fmt.Sprintf("[ERROR] Expected %s Got Operation Status %s", h.OperationStatus, handler.Success))
				break
			}
			if h.OperationStatus == handler.InProgress {
				h2, err := waitForSuccess(h, test, ts.TestHandler.Update)
				if err != nil {
					t.Error(fmt.Sprintf("Error while waiting operation to conclude %s", err))
					break
				}
				h = h2
			}
			model = h.ResourceModel
		}
		if _, err := runTestStepChecks(model, test); err != nil {
			t.Error(fmt.Sprintf("Error in check: %s", err))
		}

	}
	if model != nil {
		log.Printf("[WARN] Test: Executing Delete step")
		data, err := json.Marshal(model)
		if err != nil {
			t.Error(fmt.Sprintf("[ERROR] Test: Error marshaling resource %s", err))
			return
		}
		req := handler.NewRequest("id", map[string]interface{}{}, &session.Session{}, nil, data)
		h := ts.TestHandler.Delete(req)
		if h.OperationStatus != handler.Success {
			t.Error(fmt.Sprintf("[ERROR] Expected %s Got Operation Status %s", h.OperationStatus, handler.Success))
		}
	}
}

func runTestStepChecks(model interface{}, step TestStep) (interface{}, error) {
	if step.Check != nil {
		if err := step.Check(model); err != nil {
			return model, err
		}
	}
	return model, nil
}

func waitForSuccess(h handler.ProgressEvent, test TestStep, op Operation) (handler.ProgressEvent, error) {
	for h.OperationStatus != handler.Success {
		d, err := time.ParseDuration(fmt.Sprintf("%ds", h.CallbackDelaySeconds))
		if err != nil {
			return h, fmt.Errorf("Failed to get duration: %s", err)
		}

		time.Sleep(d)

		data, err := json.Marshal(h.ResourceModel)
		if err != nil {
			return h, err
		}
		ctx := h.CallbackContext
		req := handler.NewRequest("id", ctx, &session.Session{}, nil, data)
		h = op(req)
		if h.OperationStatus == handler.Failed {
			return h, fmt.Errorf("Failed performing operation: %s", h.Message)
		}
	}
	return h, nil
}
