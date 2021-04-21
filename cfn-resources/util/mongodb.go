package util

import (
	"context"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	//"time"
)

func GetMongoDBClientEnvAuth(mongodbURI string) (*mongo.Client, error) {

	envVariablesCredential := options.Credential{
		AuthMechanism: "MONGODB-AWS",
		AuthSource:    "$external",
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodbURI), options.Client().SetAuth(envVariablesCredential))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("setupMongoDBClient client:%v", client)
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		// TODO - Catch error here, if flag set to "automate" adding a dbuser
		// for the current AWS IAM role - some assumed role for the lambda
		// create a user for short time active
		log.Printf("Got error connecting to driver: %+v", err)
		return nil, err
	}
	log.Println("setupMongoDBClient was able to ping primary")
	return client, nil
}

func getMongoDBClientAWSAuth(mongodbURI string, req *handler.Request, roleToAssumeArn string) (*mongo.Client, error) {

	// Create the credentials from AssumeRoleProvider to assume the role
	// referenced by the "myRoleARN" ARN.
	creds := stscreds.NewCredentials(req.Session, roleToAssumeArn)
	credValues, err := creds.Get()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Need to fetch this from the result?
	log.Printf("stscreds credValues:%+v", credValues)
	assumeRoleCredential := options.Credential{
		AuthMechanism: "MONGODB-AWS",
		Username:      credValues.AccessKeyID,
		Password:      credValues.SecretAccessKey,
		AuthMechanismProperties: map[string]string{
			"AWS_SESSION_TOKEN": credValues.SessionToken,
		},
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodbURI), options.Client().SetAuth(assumeRoleCredential))

	/*
			envVariablesCredential := options.Credential{
				AuthMechanism: "MONGODB-AWS",
		        AuthSource: "$external",
			}
			opts = append(opts, options.Client().SetAuth(envVariablesCredential))
			client, err := mongo.Connect(context.TODO(), opts...)
	*/
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("setupMongoDBClient client:%v", client)
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		// TODO - Catch error here, if flag set to "automate" adding a dbuser
		// for the current AWS IAM role - some assumed role for the lambda
		// create a user for short time active
		log.Printf("Got error connecting to driver: %+v", err)
		return nil, err
	}
	log.Println("setupMongoDBClient was able to ping primary")
	return client, nil
}

/*
This function will take an 'active' Atlas client connection
and use it to run a db command. This works by creating a temporary
db-user to connect to the cluster

*/
func ListDatabaseNamesByClusterName(client *mongodbatlas.Client, req *handler.Request, projectID string, clusterName string, roleToAssumeArn string) ([]string, error) {
	log.Printf("ListDatabaseNames projectID:%s, clusterName:%s", projectID, clusterName)
	var databases []string

	// Lookup the mongodb+srv from the atlas client and clustername
	cluster, _, err := client.Clusters.Get(context.TODO(), projectID, clusterName)
	if err != nil {
		return databases, err
	}
	log.Printf("ListDatabaseNames - cluster lookedup SrvAddress:%s", cluster.SrvAddress)
	return ListDatabaseNames(req, cluster.SrvAddress, roleToAssumeArn)
}

func ListDatabaseNames(req *handler.Request, srvAddress string, roleToAssumeArn string) ([]string, error) {
	log.Printf("ListDatabaseNames req:%v, srvAddress:%v roleToAssumeArn:%v", req, srvAddress, roleToAssumeArn)
	var databases []string
	mongoClient, err := getMongoDBClientAWSAuth(srvAddress, req, roleToAssumeArn)
	if err != nil {
		return databases, err
	}
	log.Printf("ListDatabases mongoClient:%+v", mongoClient)
	databases, err = mongoClient.ListDatabaseNames(context.TODO(), bson.D{})
	if err != nil {
		return databases, err
	}
	return databases, nil
}
