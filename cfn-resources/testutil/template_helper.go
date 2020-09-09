package testutil

import ( 
    "reflect"
    "os"
    "fmt"
)

// template function which takes the name of an environment
// variable and return the value in the system from os.LookupEnv
// if any.
func Template_env(envvar string) (string, error) {
    value, found := os.LookupEnv(envvar)
    if !found {
        return "", fmt.Errorf("env %s NOT FOUND",envvar)
    }
    return value, nil
}

// custom default function to fix Sprig's stupidity with booleans
func Template_dfault(d interface{}, given ...interface{}) interface{} {
	if Template_empty(given) || Template_empty(given[0]) {
		return d
	}
	return given[0]
}

// custom empty function to fix Sprig's stupidity with booleans
func Template_empty(given interface{}) bool {
	g := reflect.ValueOf(given)
	if !g.IsValid() {
		return true
	}

	// Basically adapted from text/template.isTrue
	switch g.Kind() {
	default:
		return g.IsNil()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return g.Len() == 0
	case reflect.Bool:
		// return !g.Bool()
		return false // bool can NEVER be empty!
	case reflect.Complex64, reflect.Complex128:
		return g.Complex() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return g.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return g.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return g.Float() == 0
	case reflect.Struct:
		return false
	}
}

