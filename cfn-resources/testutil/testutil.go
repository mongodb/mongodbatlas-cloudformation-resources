package testutil

type TestOperation int

type TestCase struct {
	Name        string
	Steps       []TestStep
	TestHandler TestHandler
}

type TestStep struct {
	Config    string
	Check     TestCheckFunc
	Operation TestOperation
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
