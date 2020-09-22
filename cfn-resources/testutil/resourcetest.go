package testutil 

import (
    "bytes"
	"fmt"
	"math/rand"
	"os"
    "path/filepath"
    "encoding/json"
    "gopkg.in/yaml.v2"
    "io/ioutil"
	"testing"
	"time"
    "log"
    "text/template"
	"github.com/Masterminds/sprig/v3"
)

const (
	publicKeyEnv  = "ATLAS_PUBLIC_KEY"
	privateKeyEnv = "ATLAS_PRIVATE_KEY"
	orgIDEnv      = "ATLAS_ORG_ID"
    testInputTemplate     = "./test_template.json"
    testInputTemplateData     = "./test_data.yaml"
)

var (
	publicKey  = os.Getenv(publicKeyEnv)
	privateKey = os.Getenv(privateKeyEnv)
	orgID      = os.Getenv(orgIDEnv)
)


func TestResource(resourcePath string, handler TestHandler, model interface{}, t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
    params := getTestParams(id, resourcePath)
    fmt.Printf("params---\n%v+\n",params)
    config := getTestPayloadFromTemplate(params, resourcePath, model)
    fmt.Printf("config---\n%s\n",config)
    testNameKey := fmt.Sprintf("%s-%d",resourcePath,id)
	Test(t, TestCase{
		Name:        testNameKey,
		TestHandler: handler,
		Steps: []TestStep{
			{
				Config: config,
                Operation: CreateOp,
				Check: ComposeTestCheckFunc(
					TestCheckResourceAttr("Name", testNameKey)),
			},
			{
				Config: config,
                Operation: DeleteOp,
				Check: ComposeTestCheckFunc(
					TestCheckResourceAttr("Name", testNameKey)),
			},
		},
	})
}

func getTestParams(id int, resourcePath string) *map[string]interface{} {
    params := make(map[string]interface{})


    tpl := filepath.Join(resourcePath, testInputTemplateData)
    log.Printf("getTestParams tpl=%s", tpl)
    file, err := os.Open(tpl)
    // open the file if it's there, ok for tests to not have input yaml for values
    if err == nil {
        d := yaml.NewDecoder(file)
        if err := d.Decode(&params); err != nil {
            panic(err)
        } 
    }
    defer file.Close()
    // fold in 'id' param
    params["id"] = id
    return &params
}

func getTestPayloadFromTemplate(params *map[string]interface{}, resourcePath string, model interface{}) string {

    // read resource_test.json which has canonical
    // json payload example for this resource.
    tpl := filepath.Join(resourcePath, testInputTemplate)
    log.Printf("getTestParams tpl=%s", tpl)
    text, err := ioutil.ReadFile(tpl)
	if err != nil {
	    panic(err)
	}
    fmt.Printf("text: %s", string(text))


    // Load and execute the template to inject run-time parameters, 
    // e.g. api keys -> this creates the final test payload
    t, err := template.
        New("atlas-cfn-test").
        Funcs(sprig.TxtFuncMap()).
        Funcs(template.FuncMap{
            "default": Template_dfault,
            "env":     Template_env,
        }).
        Parse(string(text))

    if err != nil {
        panic(err)
    }
    buf := new(bytes.Buffer)
    err = t.Execute(buf, &params)

    if err != nil {
        panic(err)
    }
    fmt.Printf("buf: %+v", buf)

    // Create config structure, parse the test json payload
    // to verify the sctructure is correct, fail if not.
    err = json.Unmarshal([]byte(buf.String()), model)
    if err != nil {
        panic(err)
    }
    return buf.String()
}

