package resource

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/testutil"
)

const (
	publicKeyEnv  = "ATLAS_PUBLIC_KEY"
	privateKeyEnv = "ATLAS_PRIVATE_KEY"
	orgIDEnv      = "ATLAS_ORG_ID"
)

var (
	publicKey  = os.Getenv(publicKeyEnv)
	privateKey = os.Getenv(privateKeyEnv)
	orgID      = os.Getenv(orgIDEnv)
)

func Test_CreatProject(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	config := getConfiguration(id, publicKey, privateKey, orgID)

	testutil.Test(t, testutil.TestCase{
		Name:        "Test Create",
		TestHandler: &Handler{},
		Steps: []testutil.TestStep{
			{
				Config: config,
				Check: testutil.ComposeTestCheckFunc(
					testutil.TestCheckResourceAttr("Name", fmt.Sprintf("test-acc-cfn-%d", id)),
				),
			},
		},
	})

}

func getConfiguration(id int, publicKey, privateKey, orgID string) string {
	return fmt.Sprintf(`{
"Name": "test-acc-cfn-%d",
"Orgid": "%s",
"ApiKeys": {
	"PublicKey": "%s",
	"PrivateKey": "%s"
}}`, id, orgID, publicKey, privateKey)
}
