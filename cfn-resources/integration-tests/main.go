// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package itests

/*
this would require the resource to be published to private registry already
*/
func main() {

	//readFromJsonTemplate("../../examples/trigger/trigger.json")
	// Load the AWS SDK configuration
	//cfg, err := config.LoadDefaultConfig(context.Background())
	//if err != nil {
	//	fmt.Println("Error loading configuration:", err)
	//	return
	//}
	//fmt.Println("Loaded AWS configuration")
	//// Create a new CloudFormation client
	//cfnClient := cloudformation.NewFromConfig(cfg)
	//
	////sess := session.Must(session.NewSessionWithOptions(session.Options{
	////	SharedConfigState: session.SharedConfigEnable,
	////}))
	//
	//// svc := cloudformation.New(sess)
	//fileContent, err := os.ReadFile("../../examples/trigger/trigger.json")
	//if err != nil {
	//	fmt.Println("Error loading configuration:", err)
	//	return
	//}
	//
	//// create stack
	//stackName := "stack-aws-sdk-am"
	//output, err := util.CreateStack(cfnClient, stackName, fileContent)
	//if err != nil {
	//	fmt.Println("Error creating stack:", err)
	//	return
	//}
	//fmt.Println("Stack created")
	//fmt.Println("stack output is:  ", *aws.String(*output.Stacks[0].Outputs[0].OutputKey))
	//
	//triggerId := *aws.String(*output.Stacks[0].Outputs[0].OutputValue) // only one output in this case (Id)
	//fmt.Println("created stack output triggerID:", triggerId)
	//
	////if successfully created, call atlas client and check if trigger was created
	//ctx := context.Background()
	////profileName := "default"
	//client, err := util.GetRealmClient(ctx)
	//trigger, resp, err := client.EventTriggers.Get(ctx, "63f4df9e1c744217893c19f7", "64020111d61a66df7ea2c72b", triggerId)
	//if err != nil {
	//	fmt.Println("received trigger passed", err)
	//	return
	//}
	//// assert success response that resource exists
	//// can assert attribute values if required
	//
	//// delete stack
	//err = util.DeleteStack(cfnClient, stackName)
	//
}
