package main

import (
	"bytes"
	"encoding/json"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    //"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
	//mutils "github.com/mongodb/mongodbatlas-cloudformation-resources/util"
    "log"
    "os"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
)

func getMongoDBClientEnvAuth(mongodbURI string) (*mongo.Client, error) {

	envVariablesCredential := options.Credential{
		AuthMechanism: "MONGODB-AWS",
        AuthSource: "$external",
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodbURI), options.Client().SetAuth(envVariablesCredential))
	if err != nil {
		log.Fatal(err)
        return nil, err
	}
    log.Printf("setupMongoDBClient client:%v",client)
    err = client.Ping(context.TODO(), readpref.Primary())
    if err != nil {
        // TODO - Catch error here, if flag set to "automate" adding a dbuser 
        // for the current AWS IAM role - some assumed role for the lambda
        // create a user for short time active
        log.Printf("Got error connecting to driver: %+v",err)
        return nil, err
    }
    log.Println("setupMongoDBClient was able to ping primary")
    return client, nil
}
// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

    log.Println("===============================================")
    log.Printf("%+v",os.Environ())
    log.Println("===============================================")

    uri := os.Getenv("ATLAS_URI")
    log.Printf("uri:%s",uri)
    client, err := getMongoDBClientEnvAuth(uri)
    if err != nil {

        log.Printf("Error try connecting: err:%v",err)
    }
    log.Printf("client: %v",client)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
