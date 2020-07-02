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
	projectIDEnv  = "ATLAS_PROJECT_ID"
)

var (
	publicKey  = os.Getenv(publicKeyEnv)
	privateKey = os.Getenv(privateKeyEnv)
	projectID  = os.Getenv(projectIDEnv)
)

func Test_CreatProject(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int()
	config := getConfiguration(id, publicKey, privateKey, projectID)

	testutil.Test(t, testutil.TestCase{
		Name:        "Test Create",
		TestHandler: &Handler{},
		Steps: []testutil.TestStep{
			{
				Config: config,
				Check: testutil.ComposeTestCheckFunc(
					testutil.TestCheckResourceAttr("Name", fmt.Sprintf("test-acc-cfn-cluster-%d", id)),
				),
			},
		},
	})

}

func getConfiguration(id int, publicKey, privateKey, projectID string) string {
	return fmt.Sprintf(`{
	"ProjectID": "%s",
	"Name": "test-acc-cfn-cluster-%d",
	"NumShards": 1,
	"ProviderName": "AWS",
	"ReplicationFactor": 3,
	"ProviderBackupEnabled": false,
	"AutoScaling": {
		"DiskGBEnabled": false
	},
	"MongoDBVersion": "4.0",
	"ProviderSettings": {
		"EncryptEBSVolume": false,
		"InstanceSizeName": "M10",
		"RegionName": "EU_CENTRAL_1",
		"DiskIOPS": 100
	},
	"ApiKeys": {
		"PublicKey": "%s",
		"PrivateKey": "%s"
	}
}`, projectID, id, publicKey, privateKey)
}
